//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:19
)

import (
	"context"
	"errors"
	"io"
	"math"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/trace"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/internal/balancerload"
	"google.golang.org/grpc/internal/binarylog"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcrand"
	"google.golang.org/grpc/internal/grpcutil"
	imetadata "google.golang.org/grpc/internal/metadata"
	iresolver "google.golang.org/grpc/internal/resolver"
	"google.golang.org/grpc/internal/serviceconfig"
	istatus "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

// StreamHandler defines the handler called by gRPC server to complete the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
// execution of a streaming RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
// If a StreamHandler returns an error, it should either be produced by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
// status package, or be one of the context errors. Otherwise, gRPC will use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
// codes.Unknown as the status code and err.Error() as the status message of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:50
// RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:57
type StreamHandler func(srv interface{}, stream ServerStream) error

// StreamDesc represents a streaming RPC service's method specification.  Used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:59
// on the server when registering services and on the client when initiating
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:59
// new streams.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:62
type StreamDesc struct {
	// StreamName and Handler are only used when registering handlers on a
	// server.
	StreamName	string		// the name of the method excluding the service
	Handler		StreamHandler	// the handler called for the method

	// ServerStreams and ClientStreams are used for registering handlers on a
	// server as well as defining RPC behavior when passed to NewClientStream
	// and ClientConn.NewStream.  At least one must be true.
	ServerStreams	bool	// indicates the server can perform streaming sends
	ClientStreams	bool	// indicates the client can perform streaming sends
}

// Stream defines the common interface a client or server stream has to satisfy.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:75
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:75
// Deprecated: See ClientStream and ServerStream documentation instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:78
type Stream interface {
	// Deprecated: See ClientStream and ServerStream documentation instead.
	Context() context.Context
	// Deprecated: See ClientStream and ServerStream documentation instead.
	SendMsg(m interface{}) error
	// Deprecated: See ClientStream and ServerStream documentation instead.
	RecvMsg(m interface{}) error
}

// ClientStream defines the client-side behavior of a streaming RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:87
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:87
// All errors returned from ClientStream methods are compatible with the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:87
// status package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:91
type ClientStream interface {
	// Header returns the header metadata received from the server if there
	// is any. It blocks if the metadata is not ready to read.
	Header() (metadata.MD, error)
	// Trailer returns the trailer metadata from the server, if there is any.
	// It must only be called after stream.CloseAndRecv has returned, or
	// stream.Recv has returned a non-nil error (including io.EOF).
	Trailer() metadata.MD
	// CloseSend closes the send direction of the stream. It closes the stream
	// when non-nil error is met. It is also not safe to call CloseSend
	// concurrently with SendMsg.
	CloseSend() error
	// Context returns the context for this stream.
	//
	// It should not be called until after Header or RecvMsg has returned. Once
	// called, subsequent client-side retries are disabled.
	Context() context.Context
	// SendMsg is generally called by generated code. On error, SendMsg aborts
	// the stream. If the error was generated by the client, the status is
	// returned directly; otherwise, io.EOF is returned and the status of
	// the stream may be discovered using RecvMsg.
	//
	// SendMsg blocks until:
	//   - There is sufficient flow control to schedule m with the transport, or
	//   - The stream is done, or
	//   - The stream breaks.
	//
	// SendMsg does not wait until the message is received by the server. An
	// untimely stream closure may result in lost messages. To ensure delivery,
	// users should ensure the RPC completed successfully using RecvMsg.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not safe
	// to call SendMsg on the same stream in different goroutines. It is also
	// not safe to call CloseSend concurrently with SendMsg.
	SendMsg(m interface{}) error
	// RecvMsg blocks until it receives a message into m or the stream is
	// done. It returns io.EOF when the stream completes successfully. On
	// any other error, the stream is aborted and the error contains the RPC
	// status.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not
	// safe to call RecvMsg on the same stream in different goroutines.
	RecvMsg(m interface{}) error
}

// NewStream creates a new Stream for the client side. This is typically
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
// called by generated code. ctx is used for the lifetime of the stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
// To ensure resources are not leaked due to the stream returned, one of the following
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
// actions must be performed:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//  1. Call Close on the ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//  2. Cancel the context provided.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//  3. Call RecvMsg until a non-nil error is returned. A protobuf-generated
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//     client-streaming RPC, for instance, might use the helper function
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//     CloseAndRecv (note that CloseSend does not Recv, therefore is not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//     guaranteed to release all resources).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//  4. Receive a non-nil, non-io.EOF error from Header or SendMsg.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
// If none of the above happen, a goroutine and a context will be leaked, and grpc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:138
// will not call the optionally-configured stats handler with a stats.End message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:154
func (cc *ClientConn) NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:154
	_go_fuzz_dep_.CoverTab[80727]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:157
	opts = combine(cc.dopts.callOptions, opts)

	if cc.dopts.streamInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:159
		_go_fuzz_dep_.CoverTab[80729]++
											return cc.dopts.streamInt(ctx, desc, cc, method, newClientStream, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:160
		// _ = "end of CoverTab[80729]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:161
		_go_fuzz_dep_.CoverTab[80730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:161
		// _ = "end of CoverTab[80730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:161
	// _ = "end of CoverTab[80727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:161
	_go_fuzz_dep_.CoverTab[80728]++
										return newClientStream(ctx, desc, cc, method, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:162
	// _ = "end of CoverTab[80728]"
}

// NewClientStream is a wrapper for ClientConn.NewStream.
func NewClientStream(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:166
	_go_fuzz_dep_.CoverTab[80731]++
										return cc.NewStream(ctx, desc, method, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:167
	// _ = "end of CoverTab[80731]"
}

func newClientStream(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (_ ClientStream, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:170
	_go_fuzz_dep_.CoverTab[80732]++
										if md, added, ok := metadata.FromOutgoingContextRaw(ctx); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:171
		_go_fuzz_dep_.CoverTab[80739]++

											if err := imetadata.Validate(md); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:173
			_go_fuzz_dep_.CoverTab[80741]++
												return nil, status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:174
			// _ = "end of CoverTab[80741]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:175
			_go_fuzz_dep_.CoverTab[80742]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:175
			// _ = "end of CoverTab[80742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:175
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:175
		// _ = "end of CoverTab[80739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:175
		_go_fuzz_dep_.CoverTab[80740]++

											for _, kvs := range added {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:177
			_go_fuzz_dep_.CoverTab[80743]++
												for i := 0; i < len(kvs); i += 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:178
				_go_fuzz_dep_.CoverTab[80744]++
													if err := imetadata.ValidatePair(kvs[i], kvs[i+1]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:179
					_go_fuzz_dep_.CoverTab[80745]++
														return nil, status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:180
					// _ = "end of CoverTab[80745]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:181
					_go_fuzz_dep_.CoverTab[80746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:181
					// _ = "end of CoverTab[80746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:181
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:181
				// _ = "end of CoverTab[80744]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:182
			// _ = "end of CoverTab[80743]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:183
		// _ = "end of CoverTab[80740]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:184
		_go_fuzz_dep_.CoverTab[80747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:184
		// _ = "end of CoverTab[80747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:184
	// _ = "end of CoverTab[80732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:184
	_go_fuzz_dep_.CoverTab[80733]++
										if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:185
		_go_fuzz_dep_.CoverTab[80748]++
											cc.incrCallsStarted()
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:187
			_go_fuzz_dep_.CoverTab[80749]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:188
				_go_fuzz_dep_.CoverTab[80750]++
													cc.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:189
				// _ = "end of CoverTab[80750]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:190
				_go_fuzz_dep_.CoverTab[80751]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:190
				// _ = "end of CoverTab[80751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:190
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:190
			// _ = "end of CoverTab[80749]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:191
		// _ = "end of CoverTab[80748]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:192
		_go_fuzz_dep_.CoverTab[80752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:192
		// _ = "end of CoverTab[80752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:192
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:192
	// _ = "end of CoverTab[80733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:192
	_go_fuzz_dep_.CoverTab[80734]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:195
	if err := cc.waitForResolvedAddrs(ctx); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:195
		_go_fuzz_dep_.CoverTab[80753]++
											return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:196
		// _ = "end of CoverTab[80753]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:197
		_go_fuzz_dep_.CoverTab[80754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:197
		// _ = "end of CoverTab[80754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:197
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:197
	// _ = "end of CoverTab[80734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:197
	_go_fuzz_dep_.CoverTab[80735]++

										var mc serviceconfig.MethodConfig
										var onCommit func()
										var newStream = func(ctx context.Context, done func()) (iresolver.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:201
		_go_fuzz_dep_.CoverTab[80755]++
											return newClientStreamWithParams(ctx, desc, cc, method, mc, onCommit, done, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:202
		// _ = "end of CoverTab[80755]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:203
	// _ = "end of CoverTab[80735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:203
	_go_fuzz_dep_.CoverTab[80736]++

										rpcInfo := iresolver.RPCInfo{Context: ctx, Method: method}
										rpcConfig, err := cc.safeConfigSelector.SelectConfig(rpcInfo)
										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:207
		_go_fuzz_dep_.CoverTab[80756]++
											if st, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:208
			_go_fuzz_dep_.CoverTab[80758]++

												if istatus.IsRestrictedControlPlaneCode(st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:210
				_go_fuzz_dep_.CoverTab[80760]++
													err = status.Errorf(codes.Internal, "config selector returned illegal status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:211
				// _ = "end of CoverTab[80760]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:212
				_go_fuzz_dep_.CoverTab[80761]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:212
				// _ = "end of CoverTab[80761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:212
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:212
			// _ = "end of CoverTab[80758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:212
			_go_fuzz_dep_.CoverTab[80759]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:213
			// _ = "end of CoverTab[80759]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:214
			_go_fuzz_dep_.CoverTab[80762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:214
			// _ = "end of CoverTab[80762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:214
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:214
		// _ = "end of CoverTab[80756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:214
		_go_fuzz_dep_.CoverTab[80757]++
											return nil, toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:215
		// _ = "end of CoverTab[80757]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:216
		_go_fuzz_dep_.CoverTab[80763]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:216
		// _ = "end of CoverTab[80763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:216
	// _ = "end of CoverTab[80736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:216
	_go_fuzz_dep_.CoverTab[80737]++

										if rpcConfig != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:218
		_go_fuzz_dep_.CoverTab[80764]++
											if rpcConfig.Context != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:219
			_go_fuzz_dep_.CoverTab[80766]++
												ctx = rpcConfig.Context
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:220
			// _ = "end of CoverTab[80766]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:221
			_go_fuzz_dep_.CoverTab[80767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:221
			// _ = "end of CoverTab[80767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:221
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:221
		// _ = "end of CoverTab[80764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:221
		_go_fuzz_dep_.CoverTab[80765]++
											mc = rpcConfig.MethodConfig
											onCommit = rpcConfig.OnCommitted
											if rpcConfig.Interceptor != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:224
			_go_fuzz_dep_.CoverTab[80768]++
												rpcInfo.Context = nil
												ns := newStream
												newStream = func(ctx context.Context, done func()) (iresolver.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:227
				_go_fuzz_dep_.CoverTab[80769]++
													cs, err := rpcConfig.Interceptor.NewStream(ctx, rpcInfo, done, ns)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:229
					_go_fuzz_dep_.CoverTab[80771]++
														return nil, toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:230
					// _ = "end of CoverTab[80771]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:231
					_go_fuzz_dep_.CoverTab[80772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:231
					// _ = "end of CoverTab[80772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:231
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:231
				// _ = "end of CoverTab[80769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:231
				_go_fuzz_dep_.CoverTab[80770]++
													return cs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:232
				// _ = "end of CoverTab[80770]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:233
			// _ = "end of CoverTab[80768]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:234
			_go_fuzz_dep_.CoverTab[80773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:234
			// _ = "end of CoverTab[80773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:234
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:234
		// _ = "end of CoverTab[80765]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:235
		_go_fuzz_dep_.CoverTab[80774]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:235
		// _ = "end of CoverTab[80774]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:235
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:235
	// _ = "end of CoverTab[80737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:235
	_go_fuzz_dep_.CoverTab[80738]++

										return newStream(ctx, func() { _go_fuzz_dep_.CoverTab[80775]++; // _ = "end of CoverTab[80775]" })
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:237
	// _ = "end of CoverTab[80738]"
}

func newClientStreamWithParams(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, mc serviceconfig.MethodConfig, onCommit, doneFunc func(), opts ...CallOption) (_ iresolver.ClientStream, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:240
	_go_fuzz_dep_.CoverTab[80776]++
										c := defaultCallInfo()
										if mc.WaitForReady != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:242
		_go_fuzz_dep_.CoverTab[80791]++
											c.failFast = !*mc.WaitForReady
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:243
		// _ = "end of CoverTab[80791]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:244
		_go_fuzz_dep_.CoverTab[80792]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:244
		// _ = "end of CoverTab[80792]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:244
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:244
	// _ = "end of CoverTab[80776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:244
	_go_fuzz_dep_.CoverTab[80777]++

	// Possible context leak:
	// The cancel function for the child context we create will only be called
	// when RecvMsg returns a non-nil error, if the ClientConn is closed, or if
	// an error is generated by SendMsg.
	// https://github.com/grpc/grpc-go/issues/1818.
	var cancel context.CancelFunc
	if mc.Timeout != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:252
		_go_fuzz_dep_.CoverTab[80793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:252
		return *mc.Timeout >= 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:252
		// _ = "end of CoverTab[80793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:252
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:252
		_go_fuzz_dep_.CoverTab[80794]++
											ctx, cancel = context.WithTimeout(ctx, *mc.Timeout)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:253
		// _ = "end of CoverTab[80794]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:254
		_go_fuzz_dep_.CoverTab[80795]++
											ctx, cancel = context.WithCancel(ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:255
		// _ = "end of CoverTab[80795]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:256
	// _ = "end of CoverTab[80777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:256
	_go_fuzz_dep_.CoverTab[80778]++
										defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:257
		_go_fuzz_dep_.CoverTab[80796]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:258
			_go_fuzz_dep_.CoverTab[80797]++
												cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:259
			// _ = "end of CoverTab[80797]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:260
			_go_fuzz_dep_.CoverTab[80798]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:260
			// _ = "end of CoverTab[80798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:260
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:260
		// _ = "end of CoverTab[80796]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:261
	// _ = "end of CoverTab[80778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:261
	_go_fuzz_dep_.CoverTab[80779]++

										for _, o := range opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:263
		_go_fuzz_dep_.CoverTab[80799]++
											if err := o.before(c); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:264
			_go_fuzz_dep_.CoverTab[80800]++
												return nil, toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:265
			// _ = "end of CoverTab[80800]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:266
			_go_fuzz_dep_.CoverTab[80801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:266
			// _ = "end of CoverTab[80801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:266
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:266
		// _ = "end of CoverTab[80799]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:267
	// _ = "end of CoverTab[80779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:267
	_go_fuzz_dep_.CoverTab[80780]++
										c.maxSendMessageSize = getMaxSize(mc.MaxReqSize, c.maxSendMessageSize, defaultClientMaxSendMessageSize)
										c.maxReceiveMessageSize = getMaxSize(mc.MaxRespSize, c.maxReceiveMessageSize, defaultClientMaxReceiveMessageSize)
										if err := setCallInfoCodec(c); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:270
		_go_fuzz_dep_.CoverTab[80802]++
											return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:271
		// _ = "end of CoverTab[80802]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:272
		_go_fuzz_dep_.CoverTab[80803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:272
		// _ = "end of CoverTab[80803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:272
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:272
	// _ = "end of CoverTab[80780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:272
	_go_fuzz_dep_.CoverTab[80781]++

										callHdr := &transport.CallHdr{
		Host:		cc.authority,
		Method:		method,
		ContentSubtype:	c.contentSubtype,
		DoneFunc:	doneFunc,
	}

	// Set our outgoing compression according to the UseCompressor CallOption, if
	// set.  In that case, also find the compressor from the encoding package.
	// Otherwise, use the compressor configured by the WithCompressor DialOption,
	// if set.
	var cp Compressor
	var comp encoding.Compressor
	if ct := c.compressorType; ct != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:287
		_go_fuzz_dep_.CoverTab[80804]++
											callHdr.SendCompress = ct
											if ct != encoding.Identity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:289
			_go_fuzz_dep_.CoverTab[80805]++
												comp = encoding.GetCompressor(ct)
												if comp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:291
				_go_fuzz_dep_.CoverTab[80806]++
													return nil, status.Errorf(codes.Internal, "grpc: Compressor is not installed for requested grpc-encoding %q", ct)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:292
				// _ = "end of CoverTab[80806]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:293
				_go_fuzz_dep_.CoverTab[80807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:293
				// _ = "end of CoverTab[80807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:293
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:293
			// _ = "end of CoverTab[80805]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:294
			_go_fuzz_dep_.CoverTab[80808]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:294
			// _ = "end of CoverTab[80808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:294
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:294
		// _ = "end of CoverTab[80804]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:295
		_go_fuzz_dep_.CoverTab[80809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:295
		if cc.dopts.cp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:295
			_go_fuzz_dep_.CoverTab[80810]++
												callHdr.SendCompress = cc.dopts.cp.Type()
												cp = cc.dopts.cp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:297
			// _ = "end of CoverTab[80810]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
			_go_fuzz_dep_.CoverTab[80811]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
			// _ = "end of CoverTab[80811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
		// _ = "end of CoverTab[80809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
	// _ = "end of CoverTab[80781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:298
	_go_fuzz_dep_.CoverTab[80782]++
										if c.creds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:299
		_go_fuzz_dep_.CoverTab[80812]++
											callHdr.Creds = c.creds
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:300
		// _ = "end of CoverTab[80812]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:301
		_go_fuzz_dep_.CoverTab[80813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:301
		// _ = "end of CoverTab[80813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:301
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:301
	// _ = "end of CoverTab[80782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:301
	_go_fuzz_dep_.CoverTab[80783]++

										cs := &clientStream{
		callHdr:	callHdr,
		ctx:		ctx,
		methodConfig:	&mc,
		opts:		opts,
		callInfo:	c,
		cc:		cc,
		desc:		desc,
		codec:		c.codec,
		cp:		cp,
		comp:		comp,
		cancel:		cancel,
		firstAttempt:	true,
		onCommit:	onCommit,
	}
	if !cc.dopts.disableRetry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:318
		_go_fuzz_dep_.CoverTab[80814]++
											cs.retryThrottler = cc.retryThrottler.Load().(*retryThrottler)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:319
		// _ = "end of CoverTab[80814]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:320
		_go_fuzz_dep_.CoverTab[80815]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:320
		// _ = "end of CoverTab[80815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:320
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:320
	// _ = "end of CoverTab[80783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:320
	_go_fuzz_dep_.CoverTab[80784]++
										if ml := binarylog.GetMethodLogger(method); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:321
		_go_fuzz_dep_.CoverTab[80816]++
											cs.binlogs = append(cs.binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:322
		// _ = "end of CoverTab[80816]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:323
		_go_fuzz_dep_.CoverTab[80817]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:323
		// _ = "end of CoverTab[80817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:323
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:323
	// _ = "end of CoverTab[80784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:323
	_go_fuzz_dep_.CoverTab[80785]++
										if cc.dopts.binaryLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:324
		_go_fuzz_dep_.CoverTab[80818]++
											if ml := cc.dopts.binaryLogger.GetMethodLogger(method); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:325
			_go_fuzz_dep_.CoverTab[80819]++
												cs.binlogs = append(cs.binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:326
			// _ = "end of CoverTab[80819]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:327
			_go_fuzz_dep_.CoverTab[80820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:327
			// _ = "end of CoverTab[80820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:327
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:327
		// _ = "end of CoverTab[80818]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:328
		_go_fuzz_dep_.CoverTab[80821]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:328
		// _ = "end of CoverTab[80821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:328
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:328
	// _ = "end of CoverTab[80785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:328
	_go_fuzz_dep_.CoverTab[80786]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:332
	op := func(a *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:332
		_go_fuzz_dep_.CoverTab[80822]++
											if err := a.getTransport(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:333
			_go_fuzz_dep_.CoverTab[80825]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:334
			// _ = "end of CoverTab[80825]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:335
			_go_fuzz_dep_.CoverTab[80826]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:335
			// _ = "end of CoverTab[80826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:335
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:335
		// _ = "end of CoverTab[80822]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:335
		_go_fuzz_dep_.CoverTab[80823]++
											if err := a.newStream(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:336
			_go_fuzz_dep_.CoverTab[80827]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:337
			// _ = "end of CoverTab[80827]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:338
			_go_fuzz_dep_.CoverTab[80828]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:338
			// _ = "end of CoverTab[80828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:338
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:338
		// _ = "end of CoverTab[80823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:338
		_go_fuzz_dep_.CoverTab[80824]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:342
		cs.attempt = a
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:343
		// _ = "end of CoverTab[80824]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:344
	// _ = "end of CoverTab[80786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:344
	_go_fuzz_dep_.CoverTab[80787]++
										if err := cs.withRetry(op, func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:345
		_go_fuzz_dep_.CoverTab[80829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:345
		cs.bufferForRetryLocked(0, op)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:345
		// _ = "end of CoverTab[80829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:345
	}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:345
		_go_fuzz_dep_.CoverTab[80830]++
											return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:346
		// _ = "end of CoverTab[80830]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:347
		_go_fuzz_dep_.CoverTab[80831]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:347
		// _ = "end of CoverTab[80831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:347
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:347
	// _ = "end of CoverTab[80787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:347
	_go_fuzz_dep_.CoverTab[80788]++

										if len(cs.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:349
		_go_fuzz_dep_.CoverTab[80832]++
											md, _ := metadata.FromOutgoingContext(ctx)
											logEntry := &binarylog.ClientHeader{
			OnClientSide:	true,
			Header:		md,
			MethodName:	method,
			Authority:	cs.cc.authority,
		}
		if deadline, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:357
			_go_fuzz_dep_.CoverTab[80834]++
												logEntry.Timeout = time.Until(deadline)
												if logEntry.Timeout < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:359
				_go_fuzz_dep_.CoverTab[80835]++
													logEntry.Timeout = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:360
				// _ = "end of CoverTab[80835]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:361
				_go_fuzz_dep_.CoverTab[80836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:361
				// _ = "end of CoverTab[80836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:361
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:361
			// _ = "end of CoverTab[80834]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:362
			_go_fuzz_dep_.CoverTab[80837]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:362
			// _ = "end of CoverTab[80837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:362
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:362
		// _ = "end of CoverTab[80832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:362
		_go_fuzz_dep_.CoverTab[80833]++
											for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:363
			_go_fuzz_dep_.CoverTab[80838]++
												binlog.Log(cs.ctx, logEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:364
			// _ = "end of CoverTab[80838]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:365
		// _ = "end of CoverTab[80833]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:366
		_go_fuzz_dep_.CoverTab[80839]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:366
		// _ = "end of CoverTab[80839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:366
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:366
	// _ = "end of CoverTab[80788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:366
	_go_fuzz_dep_.CoverTab[80789]++

										if desc != unaryStreamDesc {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:368
		_go_fuzz_dep_.CoverTab[80840]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:368
		_curRoutineNum97_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:368
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum97_)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
		go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
			_go_fuzz_dep_.CoverTab[80841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
			defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
				_go_fuzz_dep_.CoverTab[80842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum97_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
				// _ = "end of CoverTab[80842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:374
			}()
												select {
			case <-cc.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:376
				_go_fuzz_dep_.CoverTab[80843]++
													cs.finish(ErrClientConnClosing)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:377
				// _ = "end of CoverTab[80843]"
			case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:378
				_go_fuzz_dep_.CoverTab[80844]++
													cs.finish(toRPCErr(ctx.Err()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:379
				// _ = "end of CoverTab[80844]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:380
			// _ = "end of CoverTab[80841]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:381
		// _ = "end of CoverTab[80840]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:382
		_go_fuzz_dep_.CoverTab[80845]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:382
		// _ = "end of CoverTab[80845]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:382
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:382
	// _ = "end of CoverTab[80789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:382
	_go_fuzz_dep_.CoverTab[80790]++
										return cs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:383
	// _ = "end of CoverTab[80790]"
}

// newAttemptLocked creates a new csAttempt without a transport or stream.
func (cs *clientStream) newAttemptLocked(isTransparent bool) (*csAttempt, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:387
	_go_fuzz_dep_.CoverTab[80846]++
										if err := cs.ctx.Err(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:388
		_go_fuzz_dep_.CoverTab[80852]++
											return nil, toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:389
		// _ = "end of CoverTab[80852]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:390
		_go_fuzz_dep_.CoverTab[80853]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:390
		// _ = "end of CoverTab[80853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:390
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:390
	// _ = "end of CoverTab[80846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:390
	_go_fuzz_dep_.CoverTab[80847]++
										if err := cs.cc.ctx.Err(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:391
		_go_fuzz_dep_.CoverTab[80854]++
											return nil, ErrClientConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:392
		// _ = "end of CoverTab[80854]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:393
		_go_fuzz_dep_.CoverTab[80855]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:393
		// _ = "end of CoverTab[80855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:393
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:393
	// _ = "end of CoverTab[80847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:393
	_go_fuzz_dep_.CoverTab[80848]++

										ctx := newContextWithRPCInfo(cs.ctx, cs.callInfo.failFast, cs.callInfo.codec, cs.cp, cs.comp)
										method := cs.callHdr.Method
										var beginTime time.Time
										shs := cs.cc.dopts.copts.StatsHandlers
										for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:399
		_go_fuzz_dep_.CoverTab[80856]++
											ctx = sh.TagRPC(ctx, &stats.RPCTagInfo{FullMethodName: method, FailFast: cs.callInfo.failFast})
											beginTime = time.Now()
											begin := &stats.Begin{
			Client:				true,
			BeginTime:			beginTime,
			FailFast:			cs.callInfo.failFast,
			IsClientStream:			cs.desc.ClientStreams,
			IsServerStream:			cs.desc.ServerStreams,
			IsTransparentRetryAttempt:	isTransparent,
		}
											sh.HandleRPC(ctx, begin)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:410
		// _ = "end of CoverTab[80856]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:411
	// _ = "end of CoverTab[80848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:411
	_go_fuzz_dep_.CoverTab[80849]++

										var trInfo *traceInfo
										if EnableTracing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:414
		_go_fuzz_dep_.CoverTab[80857]++
											trInfo = &traceInfo{
			tr:	trace.New("grpc.Sent."+methodFamily(method), method),
			firstLine: firstLine{
				client: true,
			},
		}
		if deadline, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:421
			_go_fuzz_dep_.CoverTab[80859]++
												trInfo.firstLine.deadline = time.Until(deadline)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:422
			// _ = "end of CoverTab[80859]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:423
			_go_fuzz_dep_.CoverTab[80860]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:423
			// _ = "end of CoverTab[80860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:423
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:423
		// _ = "end of CoverTab[80857]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:423
		_go_fuzz_dep_.CoverTab[80858]++
											trInfo.tr.LazyLog(&trInfo.firstLine, false)
											ctx = trace.NewContext(ctx, trInfo.tr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:425
		// _ = "end of CoverTab[80858]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:426
		_go_fuzz_dep_.CoverTab[80861]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:426
		// _ = "end of CoverTab[80861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:426
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:426
	// _ = "end of CoverTab[80849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:426
	_go_fuzz_dep_.CoverTab[80850]++

										if cs.cc.parsedTarget.URL.Scheme == "xds" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:428
		_go_fuzz_dep_.CoverTab[80862]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:431
		ctx = grpcutil.WithExtraMetadata(ctx, metadata.Pairs(
			"content-type", grpcutil.ContentType(cs.callHdr.ContentSubtype),
		))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:433
		// _ = "end of CoverTab[80862]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:434
		_go_fuzz_dep_.CoverTab[80863]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:434
		// _ = "end of CoverTab[80863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:434
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:434
	// _ = "end of CoverTab[80850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:434
	_go_fuzz_dep_.CoverTab[80851]++

										return &csAttempt{
		ctx:		ctx,
		beginTime:	beginTime,
		cs:		cs,
		dc:		cs.cc.dopts.dc,
		statsHandlers:	shs,
		trInfo:		trInfo,
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:443
	// _ = "end of CoverTab[80851]"
}

func (a *csAttempt) getTransport() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:446
	_go_fuzz_dep_.CoverTab[80864]++
										cs := a.cs

										var err error
										a.t, a.pickResult, err = cs.cc.getTransport(a.ctx, cs.callInfo.failFast, cs.callHdr.Method)
										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:451
		_go_fuzz_dep_.CoverTab[80867]++
											if de, ok := err.(dropError); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:452
			_go_fuzz_dep_.CoverTab[80869]++
												err = de.error
												a.drop = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:454
			// _ = "end of CoverTab[80869]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:455
			_go_fuzz_dep_.CoverTab[80870]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:455
			// _ = "end of CoverTab[80870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:455
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:455
		// _ = "end of CoverTab[80867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:455
		_go_fuzz_dep_.CoverTab[80868]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:456
		// _ = "end of CoverTab[80868]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:457
		_go_fuzz_dep_.CoverTab[80871]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:457
		// _ = "end of CoverTab[80871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:457
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:457
	// _ = "end of CoverTab[80864]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:457
	_go_fuzz_dep_.CoverTab[80865]++
										if a.trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:458
		_go_fuzz_dep_.CoverTab[80872]++
											a.trInfo.firstLine.SetRemoteAddr(a.t.RemoteAddr())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:459
		// _ = "end of CoverTab[80872]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:460
		_go_fuzz_dep_.CoverTab[80873]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:460
		// _ = "end of CoverTab[80873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:460
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:460
	// _ = "end of CoverTab[80865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:460
	_go_fuzz_dep_.CoverTab[80866]++
										return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:461
	// _ = "end of CoverTab[80866]"
}

func (a *csAttempt) newStream() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:464
	_go_fuzz_dep_.CoverTab[80874]++
										cs := a.cs
										cs.callHdr.PreviousAttempts = cs.numRetries

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:472
	if a.pickResult.Metatada != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:472
		_go_fuzz_dep_.CoverTab[80877]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:481
		md, _ := metadata.FromOutgoingContext(a.ctx)
											md = metadata.Join(md, a.pickResult.Metatada)
											a.ctx = metadata.NewOutgoingContext(a.ctx, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:483
		// _ = "end of CoverTab[80877]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:484
		_go_fuzz_dep_.CoverTab[80878]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:484
		// _ = "end of CoverTab[80878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:484
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:484
	// _ = "end of CoverTab[80874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:484
	_go_fuzz_dep_.CoverTab[80875]++

										s, err := a.t.NewStream(a.ctx, cs.callHdr)
										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:487
		_go_fuzz_dep_.CoverTab[80879]++
											nse, ok := err.(*transport.NewStreamError)
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:489
			_go_fuzz_dep_.CoverTab[80882]++

												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:491
			// _ = "end of CoverTab[80882]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:492
			_go_fuzz_dep_.CoverTab[80883]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:492
			// _ = "end of CoverTab[80883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:492
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:492
		// _ = "end of CoverTab[80879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:492
		_go_fuzz_dep_.CoverTab[80880]++

											if nse.AllowTransparentRetry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:494
			_go_fuzz_dep_.CoverTab[80884]++
												a.allowTransparentRetry = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:495
			// _ = "end of CoverTab[80884]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:496
			_go_fuzz_dep_.CoverTab[80885]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:496
			// _ = "end of CoverTab[80885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:496
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:496
		// _ = "end of CoverTab[80880]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:496
		_go_fuzz_dep_.CoverTab[80881]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:499
		return toRPCErr(nse.Err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:499
		// _ = "end of CoverTab[80881]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:500
		_go_fuzz_dep_.CoverTab[80886]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:500
		// _ = "end of CoverTab[80886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:500
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:500
	// _ = "end of CoverTab[80875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:500
	_go_fuzz_dep_.CoverTab[80876]++
										a.s = s
										a.p = &parser{r: s}
										return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:503
	// _ = "end of CoverTab[80876]"
}

// clientStream implements a client side Stream.
type clientStream struct {
	callHdr		*transport.CallHdr
	opts		[]CallOption
	callInfo	*callInfo
	cc		*ClientConn
	desc		*StreamDesc

	codec	baseCodec
	cp	Compressor
	comp	encoding.Compressor

	cancel	context.CancelFunc	// cancels all attempts

	sentLast	bool	// sent an end stream

	methodConfig	*MethodConfig

	ctx	context.Context	// the application's context, wrapped by stats/tracing

	retryThrottler	*retryThrottler	// The throttler active when the RPC began.

	binlogs	[]binarylog.MethodLogger
	// serverHeaderBinlogged is a boolean for whether server header has been
	// logged. Server header will be logged when the first time one of those
	// happens: stream.Header(), stream.Recv().
	//
	// It's only read and used by Recv() and Header(), so it doesn't need to be
	// synchronized.
	serverHeaderBinlogged	bool

	mu			sync.Mutex
	firstAttempt		bool	// if true, transparent retry is valid
	numRetries		int	// exclusive of transparent retry attempt(s)
	numRetriesSincePushback	int	// retries since pushback; to reset backoff
	finished		bool	// TODO: replace with atomic cmpxchg or sync.Once?
	// attempt is the active client stream attempt.
	// The only place where it is written is the newAttemptLocked method and this method never writes nil.
	// So, attempt can be nil only inside newClientStream function when clientStream is first created.
	// One of the first things done after clientStream's creation, is to call newAttemptLocked which either
	// assigns a non nil value to the attempt or returns an error. If an error is returned from newAttemptLocked,
	// then newClientStream calls finish on the clientStream and returns. So, finish method is the only
	// place where we need to check if the attempt is nil.
	attempt	*csAttempt
	// TODO(hedging): hedging will have multiple attempts simultaneously.
	committed	bool	// active attempt committed for retry?
	onCommit	func()
	buffer		[]func(a *csAttempt) error	// operations to replay on retry
	bufferSize	int				// current size of buffer
}

// csAttempt implements a single transport stream attempt within a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:557
// clientStream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:559
type csAttempt struct {
	ctx		context.Context
	cs		*clientStream
	t		transport.ClientTransport
	s		*transport.Stream
	p		*parser
	pickResult	balancer.PickResult

	finished	bool
	dc		Decompressor
	decomp		encoding.Compressor
	decompSet	bool

	mu	sync.Mutex	// guards trInfo.tr
	// trInfo may be nil (if EnableTracing is false).
	// trInfo.tr is set when created (if EnableTracing is true),
	// and cleared when the finish method is called.
	trInfo	*traceInfo

	statsHandlers	[]stats.Handler
	beginTime	time.Time

	// set for newStream errors that may be transparently retried
	allowTransparentRetry	bool
	// set for pick errors that are returned as a status
	drop	bool
}

func (cs *clientStream) commitAttemptLocked() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:587
	_go_fuzz_dep_.CoverTab[80887]++
										if !cs.committed && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:588
		_go_fuzz_dep_.CoverTab[80889]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:588
		return cs.onCommit != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:588
		// _ = "end of CoverTab[80889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:588
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:588
		_go_fuzz_dep_.CoverTab[80890]++
											cs.onCommit()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:589
		// _ = "end of CoverTab[80890]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:590
		_go_fuzz_dep_.CoverTab[80891]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:590
		// _ = "end of CoverTab[80891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:590
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:590
	// _ = "end of CoverTab[80887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:590
	_go_fuzz_dep_.CoverTab[80888]++
										cs.committed = true
										cs.buffer = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:592
	// _ = "end of CoverTab[80888]"
}

func (cs *clientStream) commitAttempt() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:595
	_go_fuzz_dep_.CoverTab[80892]++
										cs.mu.Lock()
										cs.commitAttemptLocked()
										cs.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:598
	// _ = "end of CoverTab[80892]"
}

// shouldRetry returns nil if the RPC should be retried; otherwise it returns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:601
// the error that should be returned by the operation.  If the RPC should be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:601
// retried, the bool indicates whether it is being retried transparently.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:604
func (a *csAttempt) shouldRetry(err error) (bool, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:604
	_go_fuzz_dep_.CoverTab[80893]++
										cs := a.cs

										if cs.finished || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		_go_fuzz_dep_.CoverTab[80905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		return cs.committed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		// _ = "end of CoverTab[80905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		_go_fuzz_dep_.CoverTab[80906]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		return a.drop
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		// _ = "end of CoverTab[80906]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:607
		_go_fuzz_dep_.CoverTab[80907]++

											return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:609
		// _ = "end of CoverTab[80907]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:610
		_go_fuzz_dep_.CoverTab[80908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:610
		// _ = "end of CoverTab[80908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:610
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:610
	// _ = "end of CoverTab[80893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:610
	_go_fuzz_dep_.CoverTab[80894]++
										if a.s == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:611
		_go_fuzz_dep_.CoverTab[80909]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:611
		return a.allowTransparentRetry
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:611
		// _ = "end of CoverTab[80909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:611
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:611
		_go_fuzz_dep_.CoverTab[80910]++
											return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:612
		// _ = "end of CoverTab[80910]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:613
		_go_fuzz_dep_.CoverTab[80911]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:613
		// _ = "end of CoverTab[80911]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:613
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:613
	// _ = "end of CoverTab[80894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:613
	_go_fuzz_dep_.CoverTab[80895]++

										unprocessed := false
										if a.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:616
		_go_fuzz_dep_.CoverTab[80912]++
											<-a.s.Done()
											unprocessed = a.s.Unprocessed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:618
		// _ = "end of CoverTab[80912]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:619
		_go_fuzz_dep_.CoverTab[80913]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:619
		// _ = "end of CoverTab[80913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:619
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:619
	// _ = "end of CoverTab[80895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:619
	_go_fuzz_dep_.CoverTab[80896]++
										if cs.firstAttempt && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:620
		_go_fuzz_dep_.CoverTab[80914]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:620
		return unprocessed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:620
		// _ = "end of CoverTab[80914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:620
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:620
		_go_fuzz_dep_.CoverTab[80915]++

											return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:622
		// _ = "end of CoverTab[80915]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:623
		_go_fuzz_dep_.CoverTab[80916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:623
		// _ = "end of CoverTab[80916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:623
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:623
	// _ = "end of CoverTab[80896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:623
	_go_fuzz_dep_.CoverTab[80897]++
										if cs.cc.dopts.disableRetry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:624
		_go_fuzz_dep_.CoverTab[80917]++
											return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:625
		// _ = "end of CoverTab[80917]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:626
		_go_fuzz_dep_.CoverTab[80918]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:626
		// _ = "end of CoverTab[80918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:626
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:626
	// _ = "end of CoverTab[80897]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:626
	_go_fuzz_dep_.CoverTab[80898]++

										pushback := 0
										hasPushback := false
										if a.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:630
		_go_fuzz_dep_.CoverTab[80919]++
											if !a.s.TrailersOnly() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:631
			_go_fuzz_dep_.CoverTab[80921]++
												return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:632
			// _ = "end of CoverTab[80921]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:633
			_go_fuzz_dep_.CoverTab[80922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:633
			// _ = "end of CoverTab[80922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:633
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:633
		// _ = "end of CoverTab[80919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:633
		_go_fuzz_dep_.CoverTab[80920]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:637
		sps := a.s.Trailer()["grpc-retry-pushback-ms"]
		if len(sps) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:638
			_go_fuzz_dep_.CoverTab[80923]++
												var e error
												if pushback, e = strconv.Atoi(sps[0]); e != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:640
				_go_fuzz_dep_.CoverTab[80925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:640
				return pushback < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:640
				// _ = "end of CoverTab[80925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:640
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:640
				_go_fuzz_dep_.CoverTab[80926]++
													channelz.Infof(logger, cs.cc.channelzID, "Server retry pushback specified to abort (%q).", sps[0])
													cs.retryThrottler.throttle()
													return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:643
				// _ = "end of CoverTab[80926]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:644
				_go_fuzz_dep_.CoverTab[80927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:644
				// _ = "end of CoverTab[80927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:644
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:644
			// _ = "end of CoverTab[80923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:644
			_go_fuzz_dep_.CoverTab[80924]++
												hasPushback = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:645
			// _ = "end of CoverTab[80924]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:646
			_go_fuzz_dep_.CoverTab[80928]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:646
			if len(sps) > 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:646
				_go_fuzz_dep_.CoverTab[80929]++
													channelz.Warningf(logger, cs.cc.channelzID, "Server retry pushback specified multiple values (%q); not retrying.", sps)
													cs.retryThrottler.throttle()
													return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:649
				// _ = "end of CoverTab[80929]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
				_go_fuzz_dep_.CoverTab[80930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
				// _ = "end of CoverTab[80930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
			// _ = "end of CoverTab[80928]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:650
		// _ = "end of CoverTab[80920]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:651
		_go_fuzz_dep_.CoverTab[80931]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:651
		// _ = "end of CoverTab[80931]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:651
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:651
	// _ = "end of CoverTab[80898]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:651
	_go_fuzz_dep_.CoverTab[80899]++

										var code codes.Code
										if a.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:654
		_go_fuzz_dep_.CoverTab[80932]++
											code = a.s.Status().Code()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:655
		// _ = "end of CoverTab[80932]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:656
		_go_fuzz_dep_.CoverTab[80933]++
											code = status.Code(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:657
		// _ = "end of CoverTab[80933]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:658
	// _ = "end of CoverTab[80899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:658
	_go_fuzz_dep_.CoverTab[80900]++

										rp := cs.methodConfig.RetryPolicy
										if rp == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:661
		_go_fuzz_dep_.CoverTab[80934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:661
		return !rp.RetryableStatusCodes[code]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:661
		// _ = "end of CoverTab[80934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:661
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:661
		_go_fuzz_dep_.CoverTab[80935]++
											return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:662
		// _ = "end of CoverTab[80935]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:663
		_go_fuzz_dep_.CoverTab[80936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:663
		// _ = "end of CoverTab[80936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:663
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:663
	// _ = "end of CoverTab[80900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:663
	_go_fuzz_dep_.CoverTab[80901]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:667
	if cs.retryThrottler.throttle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:667
		_go_fuzz_dep_.CoverTab[80937]++
											return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:668
		// _ = "end of CoverTab[80937]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:669
		_go_fuzz_dep_.CoverTab[80938]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:669
		// _ = "end of CoverTab[80938]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:669
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:669
	// _ = "end of CoverTab[80901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:669
	_go_fuzz_dep_.CoverTab[80902]++
										if cs.numRetries+1 >= rp.MaxAttempts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:670
		_go_fuzz_dep_.CoverTab[80939]++
											return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:671
		// _ = "end of CoverTab[80939]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:672
		_go_fuzz_dep_.CoverTab[80940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:672
		// _ = "end of CoverTab[80940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:672
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:672
	// _ = "end of CoverTab[80902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:672
	_go_fuzz_dep_.CoverTab[80903]++

										var dur time.Duration
										if hasPushback {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:675
		_go_fuzz_dep_.CoverTab[80941]++
											dur = time.Millisecond * time.Duration(pushback)
											cs.numRetriesSincePushback = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:677
		// _ = "end of CoverTab[80941]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:678
		_go_fuzz_dep_.CoverTab[80942]++
											fact := math.Pow(rp.BackoffMultiplier, float64(cs.numRetriesSincePushback))
											cur := float64(rp.InitialBackoff) * fact
											if max := float64(rp.MaxBackoff); cur > max {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:681
			_go_fuzz_dep_.CoverTab[80944]++
												cur = max
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:682
			// _ = "end of CoverTab[80944]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:683
			_go_fuzz_dep_.CoverTab[80945]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:683
			// _ = "end of CoverTab[80945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:683
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:683
		// _ = "end of CoverTab[80942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:683
		_go_fuzz_dep_.CoverTab[80943]++
											dur = time.Duration(grpcrand.Int63n(int64(cur)))
											cs.numRetriesSincePushback++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:685
		// _ = "end of CoverTab[80943]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:686
	// _ = "end of CoverTab[80903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:686
	_go_fuzz_dep_.CoverTab[80904]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:690
	t := time.NewTimer(dur)
	select {
	case <-t.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:692
		_go_fuzz_dep_.CoverTab[80946]++
											cs.numRetries++
											return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:694
		// _ = "end of CoverTab[80946]"
	case <-cs.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:695
		_go_fuzz_dep_.CoverTab[80947]++
											t.Stop()
											return false, status.FromContextError(cs.ctx.Err()).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:697
		// _ = "end of CoverTab[80947]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:698
	// _ = "end of CoverTab[80904]"
}

// Returns nil if a retry was performed and succeeded; error otherwise.
func (cs *clientStream) retryLocked(attempt *csAttempt, lastErr error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:702
	_go_fuzz_dep_.CoverTab[80948]++
										for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:703
		_go_fuzz_dep_.CoverTab[80949]++
											attempt.finish(toRPCErr(lastErr))
											isTransparent, err := attempt.shouldRetry(lastErr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:706
			_go_fuzz_dep_.CoverTab[80952]++
												cs.commitAttemptLocked()
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:708
			// _ = "end of CoverTab[80952]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:709
			_go_fuzz_dep_.CoverTab[80953]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:709
			// _ = "end of CoverTab[80953]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:709
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:709
		// _ = "end of CoverTab[80949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:709
		_go_fuzz_dep_.CoverTab[80950]++
											cs.firstAttempt = false
											attempt, err = cs.newAttemptLocked(isTransparent)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:712
			_go_fuzz_dep_.CoverTab[80954]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:715
			return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:715
			// _ = "end of CoverTab[80954]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:716
			_go_fuzz_dep_.CoverTab[80955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:716
			// _ = "end of CoverTab[80955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:716
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:716
		// _ = "end of CoverTab[80950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:716
		_go_fuzz_dep_.CoverTab[80951]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:719
		if lastErr = cs.replayBufferLocked(attempt); lastErr == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:719
			_go_fuzz_dep_.CoverTab[80956]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:720
			// _ = "end of CoverTab[80956]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:721
			_go_fuzz_dep_.CoverTab[80957]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:721
			// _ = "end of CoverTab[80957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:721
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:721
		// _ = "end of CoverTab[80951]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:722
	// _ = "end of CoverTab[80948]"
}

func (cs *clientStream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:725
	_go_fuzz_dep_.CoverTab[80958]++
										cs.commitAttempt()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:729
	if cs.attempt.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:729
		_go_fuzz_dep_.CoverTab[80960]++
											return cs.attempt.s.Context()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:730
		// _ = "end of CoverTab[80960]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:731
		_go_fuzz_dep_.CoverTab[80961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:731
		// _ = "end of CoverTab[80961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:731
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:731
	// _ = "end of CoverTab[80958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:731
	_go_fuzz_dep_.CoverTab[80959]++
										return cs.ctx
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:732
	// _ = "end of CoverTab[80959]"
}

func (cs *clientStream) withRetry(op func(a *csAttempt) error, onSuccess func()) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:735
	_go_fuzz_dep_.CoverTab[80962]++
										cs.mu.Lock()
										for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:737
		_go_fuzz_dep_.CoverTab[80963]++
											if cs.committed {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:738
			_go_fuzz_dep_.CoverTab[80969]++
												cs.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:744
			return toRPCErr(op(cs.attempt))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:744
			// _ = "end of CoverTab[80969]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:745
			_go_fuzz_dep_.CoverTab[80970]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:745
			// _ = "end of CoverTab[80970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:745
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:745
		// _ = "end of CoverTab[80963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:745
		_go_fuzz_dep_.CoverTab[80964]++
											if len(cs.buffer) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:746
			_go_fuzz_dep_.CoverTab[80971]++
			// For the first op, which controls creation of the stream and
			// assigns cs.attempt, we need to create a new attempt inline
			// before executing the first op.  On subsequent ops, the attempt
			// is created immediately before replaying the ops.
			var err error
			if cs.attempt, err = cs.newAttemptLocked(false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:752
				_go_fuzz_dep_.CoverTab[80972]++
													cs.mu.Unlock()
													cs.finish(err)
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:755
				// _ = "end of CoverTab[80972]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:756
				_go_fuzz_dep_.CoverTab[80973]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:756
				// _ = "end of CoverTab[80973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:756
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:756
			// _ = "end of CoverTab[80971]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:757
			_go_fuzz_dep_.CoverTab[80974]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:757
			// _ = "end of CoverTab[80974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:757
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:757
		// _ = "end of CoverTab[80964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:757
		_go_fuzz_dep_.CoverTab[80965]++
											a := cs.attempt
											cs.mu.Unlock()
											err := op(a)
											cs.mu.Lock()
											if a != cs.attempt {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:762
			_go_fuzz_dep_.CoverTab[80975]++

												continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:764
			// _ = "end of CoverTab[80975]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:765
			_go_fuzz_dep_.CoverTab[80976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:765
			// _ = "end of CoverTab[80976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:765
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:765
		// _ = "end of CoverTab[80965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:765
		_go_fuzz_dep_.CoverTab[80966]++
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:766
			_go_fuzz_dep_.CoverTab[80977]++
												<-a.s.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:767
			// _ = "end of CoverTab[80977]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:768
			_go_fuzz_dep_.CoverTab[80978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:768
			// _ = "end of CoverTab[80978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:768
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:768
		// _ = "end of CoverTab[80966]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:768
		_go_fuzz_dep_.CoverTab[80967]++
											if err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
			_go_fuzz_dep_.CoverTab[80979]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
			return (err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
				_go_fuzz_dep_.CoverTab[80980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
				return a.s.Status().Code() == codes.OK
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
				// _ = "end of CoverTab[80980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
			// _ = "end of CoverTab[80979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:769
			_go_fuzz_dep_.CoverTab[80981]++
												onSuccess()
												cs.mu.Unlock()
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:772
			// _ = "end of CoverTab[80981]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:773
			_go_fuzz_dep_.CoverTab[80982]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:773
			// _ = "end of CoverTab[80982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:773
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:773
		// _ = "end of CoverTab[80967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:773
		_go_fuzz_dep_.CoverTab[80968]++
											if err := cs.retryLocked(a, err); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:774
			_go_fuzz_dep_.CoverTab[80983]++
												cs.mu.Unlock()
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:776
			// _ = "end of CoverTab[80983]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:777
			_go_fuzz_dep_.CoverTab[80984]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:777
			// _ = "end of CoverTab[80984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:777
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:777
		// _ = "end of CoverTab[80968]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:778
	// _ = "end of CoverTab[80962]"
}

func (cs *clientStream) Header() (metadata.MD, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:781
	_go_fuzz_dep_.CoverTab[80985]++
										var m metadata.MD
										noHeader := false
										err := cs.withRetry(func(a *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:784
		_go_fuzz_dep_.CoverTab[80989]++
											var err error
											m, err = a.s.Header()
											if err == transport.ErrNoHeaders {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:787
			_go_fuzz_dep_.CoverTab[80991]++
												noHeader = true
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:789
			// _ = "end of CoverTab[80991]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:790
			_go_fuzz_dep_.CoverTab[80992]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:790
			// _ = "end of CoverTab[80992]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:790
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:790
		// _ = "end of CoverTab[80989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:790
		_go_fuzz_dep_.CoverTab[80990]++
											return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:791
		// _ = "end of CoverTab[80990]"
	}, cs.commitAttemptLocked)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:792
	// _ = "end of CoverTab[80985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:792
	_go_fuzz_dep_.CoverTab[80986]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:794
		_go_fuzz_dep_.CoverTab[80993]++
											cs.finish(err)
											return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:796
		// _ = "end of CoverTab[80993]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:797
		_go_fuzz_dep_.CoverTab[80994]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:797
		// _ = "end of CoverTab[80994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:797
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:797
	// _ = "end of CoverTab[80986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:797
	_go_fuzz_dep_.CoverTab[80987]++

										if len(cs.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		_go_fuzz_dep_.CoverTab[80995]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		return !cs.serverHeaderBinlogged
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		// _ = "end of CoverTab[80995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		_go_fuzz_dep_.CoverTab[80996]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		return !noHeader
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		// _ = "end of CoverTab[80996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:799
		_go_fuzz_dep_.CoverTab[80997]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:802
		logEntry := &binarylog.ServerHeader{
			OnClientSide:	true,
			Header:		m,
			PeerAddr:	nil,
		}
		if peer, ok := peer.FromContext(cs.Context()); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:807
			_go_fuzz_dep_.CoverTab[80999]++
												logEntry.PeerAddr = peer.Addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:808
			// _ = "end of CoverTab[80999]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:809
			_go_fuzz_dep_.CoverTab[81000]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:809
			// _ = "end of CoverTab[81000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:809
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:809
		// _ = "end of CoverTab[80997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:809
		_go_fuzz_dep_.CoverTab[80998]++
											cs.serverHeaderBinlogged = true
											for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:811
			_go_fuzz_dep_.CoverTab[81001]++
												binlog.Log(cs.ctx, logEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:812
			// _ = "end of CoverTab[81001]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:813
		// _ = "end of CoverTab[80998]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:814
		_go_fuzz_dep_.CoverTab[81002]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:814
		// _ = "end of CoverTab[81002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:814
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:814
	// _ = "end of CoverTab[80987]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:814
	_go_fuzz_dep_.CoverTab[80988]++
										return m, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:815
	// _ = "end of CoverTab[80988]"
}

func (cs *clientStream) Trailer() metadata.MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:818
	_go_fuzz_dep_.CoverTab[81003]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:826
	cs.commitAttempt()
	if cs.attempt.s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:827
		_go_fuzz_dep_.CoverTab[81005]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:828
		// _ = "end of CoverTab[81005]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:829
		_go_fuzz_dep_.CoverTab[81006]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:829
		// _ = "end of CoverTab[81006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:829
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:829
	// _ = "end of CoverTab[81003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:829
	_go_fuzz_dep_.CoverTab[81004]++
										return cs.attempt.s.Trailer()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:830
	// _ = "end of CoverTab[81004]"
}

func (cs *clientStream) replayBufferLocked(attempt *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:833
	_go_fuzz_dep_.CoverTab[81007]++
										for _, f := range cs.buffer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:834
		_go_fuzz_dep_.CoverTab[81009]++
											if err := f(attempt); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:835
			_go_fuzz_dep_.CoverTab[81010]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:836
			// _ = "end of CoverTab[81010]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:837
			_go_fuzz_dep_.CoverTab[81011]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:837
			// _ = "end of CoverTab[81011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:837
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:837
		// _ = "end of CoverTab[81009]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:838
	// _ = "end of CoverTab[81007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:838
	_go_fuzz_dep_.CoverTab[81008]++
										return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:839
	// _ = "end of CoverTab[81008]"
}

func (cs *clientStream) bufferForRetryLocked(sz int, op func(a *csAttempt) error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:842
	_go_fuzz_dep_.CoverTab[81012]++

										if cs.committed {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:844
		_go_fuzz_dep_.CoverTab[81015]++
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:845
		// _ = "end of CoverTab[81015]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:846
		_go_fuzz_dep_.CoverTab[81016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:846
		// _ = "end of CoverTab[81016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:846
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:846
	// _ = "end of CoverTab[81012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:846
	_go_fuzz_dep_.CoverTab[81013]++
										cs.bufferSize += sz
										if cs.bufferSize > cs.callInfo.maxRetryRPCBufferSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:848
		_go_fuzz_dep_.CoverTab[81017]++
											cs.commitAttemptLocked()
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:850
		// _ = "end of CoverTab[81017]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:851
		_go_fuzz_dep_.CoverTab[81018]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:851
		// _ = "end of CoverTab[81018]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:851
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:851
	// _ = "end of CoverTab[81013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:851
	_go_fuzz_dep_.CoverTab[81014]++
										cs.buffer = append(cs.buffer, op)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:852
	// _ = "end of CoverTab[81014]"
}

func (cs *clientStream) SendMsg(m interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:855
	_go_fuzz_dep_.CoverTab[81019]++
										defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:856
		_go_fuzz_dep_.CoverTab[81028]++
											if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:857
			_go_fuzz_dep_.CoverTab[81029]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:857
			return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:857
			// _ = "end of CoverTab[81029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:857
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:857
			_go_fuzz_dep_.CoverTab[81030]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:863
			cs.finish(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:863
			// _ = "end of CoverTab[81030]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:864
			_go_fuzz_dep_.CoverTab[81031]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:864
			// _ = "end of CoverTab[81031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:864
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:864
		// _ = "end of CoverTab[81028]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:865
	// _ = "end of CoverTab[81019]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:865
	_go_fuzz_dep_.CoverTab[81020]++
										if cs.sentLast {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:866
		_go_fuzz_dep_.CoverTab[81032]++
											return status.Errorf(codes.Internal, "SendMsg called after CloseSend")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:867
		// _ = "end of CoverTab[81032]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:868
		_go_fuzz_dep_.CoverTab[81033]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:868
		// _ = "end of CoverTab[81033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:868
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:868
	// _ = "end of CoverTab[81020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:868
	_go_fuzz_dep_.CoverTab[81021]++
										if !cs.desc.ClientStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:869
		_go_fuzz_dep_.CoverTab[81034]++
											cs.sentLast = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:870
		// _ = "end of CoverTab[81034]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:871
		_go_fuzz_dep_.CoverTab[81035]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:871
		// _ = "end of CoverTab[81035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:871
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:871
	// _ = "end of CoverTab[81021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:871
	_go_fuzz_dep_.CoverTab[81022]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:874
	hdr, payload, data, err := prepareMsg(m, cs.codec, cs.cp, cs.comp)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:875
		_go_fuzz_dep_.CoverTab[81036]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:876
		// _ = "end of CoverTab[81036]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:877
		_go_fuzz_dep_.CoverTab[81037]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:877
		// _ = "end of CoverTab[81037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:877
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:877
	// _ = "end of CoverTab[81022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:877
	_go_fuzz_dep_.CoverTab[81023]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:880
	if len(payload) > *cs.callInfo.maxSendMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:880
		_go_fuzz_dep_.CoverTab[81038]++
											return status.Errorf(codes.ResourceExhausted, "trying to send message larger than max (%d vs. %d)", len(payload), *cs.callInfo.maxSendMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:881
		// _ = "end of CoverTab[81038]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:882
		_go_fuzz_dep_.CoverTab[81039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:882
		// _ = "end of CoverTab[81039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:882
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:882
	// _ = "end of CoverTab[81023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:882
	_go_fuzz_dep_.CoverTab[81024]++
										op := func(a *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:883
		_go_fuzz_dep_.CoverTab[81040]++
											return a.sendMsg(m, hdr, payload, data)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:884
		// _ = "end of CoverTab[81040]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:885
	// _ = "end of CoverTab[81024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:885
	_go_fuzz_dep_.CoverTab[81025]++
										err = cs.withRetry(op, func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
		_go_fuzz_dep_.CoverTab[81041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
		cs.bufferForRetryLocked(len(hdr)+len(payload), op)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
		// _ = "end of CoverTab[81041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
	// _ = "end of CoverTab[81025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:886
	_go_fuzz_dep_.CoverTab[81026]++
										if len(cs.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:887
		_go_fuzz_dep_.CoverTab[81042]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:887
		return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:887
		// _ = "end of CoverTab[81042]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:887
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:887
		_go_fuzz_dep_.CoverTab[81043]++
											cm := &binarylog.ClientMessage{
			OnClientSide:	true,
			Message:	data,
		}
		for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:892
			_go_fuzz_dep_.CoverTab[81044]++
												binlog.Log(cs.ctx, cm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:893
			// _ = "end of CoverTab[81044]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:894
		// _ = "end of CoverTab[81043]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:895
		_go_fuzz_dep_.CoverTab[81045]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:895
		// _ = "end of CoverTab[81045]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:895
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:895
	// _ = "end of CoverTab[81026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:895
	_go_fuzz_dep_.CoverTab[81027]++
										return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:896
	// _ = "end of CoverTab[81027]"
}

func (cs *clientStream) RecvMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:899
	_go_fuzz_dep_.CoverTab[81046]++
										if len(cs.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:900
		_go_fuzz_dep_.CoverTab[81052]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:900
		return !cs.serverHeaderBinlogged
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:900
		// _ = "end of CoverTab[81052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:900
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:900
		_go_fuzz_dep_.CoverTab[81053]++

											cs.Header()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:902
		// _ = "end of CoverTab[81053]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:903
		_go_fuzz_dep_.CoverTab[81054]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:903
		// _ = "end of CoverTab[81054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:903
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:903
	// _ = "end of CoverTab[81046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:903
	_go_fuzz_dep_.CoverTab[81047]++
										var recvInfo *payloadInfo
										if len(cs.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:905
		_go_fuzz_dep_.CoverTab[81055]++
											recvInfo = &payloadInfo{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:906
		// _ = "end of CoverTab[81055]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:907
		_go_fuzz_dep_.CoverTab[81056]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:907
		// _ = "end of CoverTab[81056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:907
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:907
	// _ = "end of CoverTab[81047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:907
	_go_fuzz_dep_.CoverTab[81048]++
										err := cs.withRetry(func(a *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:908
		_go_fuzz_dep_.CoverTab[81057]++
											return a.recvMsg(m, recvInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:909
		// _ = "end of CoverTab[81057]"
	}, cs.commitAttemptLocked)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:910
	// _ = "end of CoverTab[81048]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:910
	_go_fuzz_dep_.CoverTab[81049]++
										if len(cs.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:911
		_go_fuzz_dep_.CoverTab[81058]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:911
		return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:911
		// _ = "end of CoverTab[81058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:911
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:911
		_go_fuzz_dep_.CoverTab[81059]++
											sm := &binarylog.ServerMessage{
			OnClientSide:	true,
			Message:	recvInfo.uncompressedBytes,
		}
		for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:916
			_go_fuzz_dep_.CoverTab[81060]++
												binlog.Log(cs.ctx, sm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:917
			// _ = "end of CoverTab[81060]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:918
		// _ = "end of CoverTab[81059]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:919
		_go_fuzz_dep_.CoverTab[81061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:919
		// _ = "end of CoverTab[81061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:919
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:919
	// _ = "end of CoverTab[81049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:919
	_go_fuzz_dep_.CoverTab[81050]++
										if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:920
		_go_fuzz_dep_.CoverTab[81062]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:920
		return !cs.desc.ServerStreams
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:920
		// _ = "end of CoverTab[81062]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:920
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:920
		_go_fuzz_dep_.CoverTab[81063]++

											cs.finish(err)

											if len(cs.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:924
			_go_fuzz_dep_.CoverTab[81064]++

												logEntry := &binarylog.ServerTrailer{
				OnClientSide:	true,
				Trailer:	cs.Trailer(),
				Err:		err,
			}
			if logEntry.Err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:931
				_go_fuzz_dep_.CoverTab[81067]++
													logEntry.Err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:932
				// _ = "end of CoverTab[81067]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:933
				_go_fuzz_dep_.CoverTab[81068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:933
				// _ = "end of CoverTab[81068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:933
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:933
			// _ = "end of CoverTab[81064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:933
			_go_fuzz_dep_.CoverTab[81065]++
												if peer, ok := peer.FromContext(cs.Context()); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:934
				_go_fuzz_dep_.CoverTab[81069]++
													logEntry.PeerAddr = peer.Addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:935
				// _ = "end of CoverTab[81069]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:936
				_go_fuzz_dep_.CoverTab[81070]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:936
				// _ = "end of CoverTab[81070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:936
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:936
			// _ = "end of CoverTab[81065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:936
			_go_fuzz_dep_.CoverTab[81066]++
												for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:937
				_go_fuzz_dep_.CoverTab[81071]++
													binlog.Log(cs.ctx, logEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:938
				// _ = "end of CoverTab[81071]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:939
			// _ = "end of CoverTab[81066]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:940
			_go_fuzz_dep_.CoverTab[81072]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:940
			// _ = "end of CoverTab[81072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:940
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:940
		// _ = "end of CoverTab[81063]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:941
		_go_fuzz_dep_.CoverTab[81073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:941
		// _ = "end of CoverTab[81073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:941
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:941
	// _ = "end of CoverTab[81050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:941
	_go_fuzz_dep_.CoverTab[81051]++
										return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:942
	// _ = "end of CoverTab[81051]"
}

func (cs *clientStream) CloseSend() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:945
	_go_fuzz_dep_.CoverTab[81074]++
										if cs.sentLast {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:946
		_go_fuzz_dep_.CoverTab[81079]++

											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:948
		// _ = "end of CoverTab[81079]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:949
		_go_fuzz_dep_.CoverTab[81080]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:949
		// _ = "end of CoverTab[81080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:949
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:949
	// _ = "end of CoverTab[81074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:949
	_go_fuzz_dep_.CoverTab[81075]++
										cs.sentLast = true
										op := func(a *csAttempt) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:951
		_go_fuzz_dep_.CoverTab[81081]++
											a.t.Write(a.s, nil, nil, &transport.Options{Last: true})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:957
		return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:957
		// _ = "end of CoverTab[81081]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:958
	// _ = "end of CoverTab[81075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:958
	_go_fuzz_dep_.CoverTab[81076]++
										cs.withRetry(op, func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
		_go_fuzz_dep_.CoverTab[81082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
		cs.bufferForRetryLocked(0, op)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
		// _ = "end of CoverTab[81082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
	// _ = "end of CoverTab[81076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:959
	_go_fuzz_dep_.CoverTab[81077]++
										if len(cs.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:960
		_go_fuzz_dep_.CoverTab[81083]++
											chc := &binarylog.ClientHalfClose{
			OnClientSide: true,
		}
		for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:964
			_go_fuzz_dep_.CoverTab[81084]++
												binlog.Log(cs.ctx, chc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:965
			// _ = "end of CoverTab[81084]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:966
		// _ = "end of CoverTab[81083]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:967
		_go_fuzz_dep_.CoverTab[81085]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:967
		// _ = "end of CoverTab[81085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:967
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:967
	// _ = "end of CoverTab[81077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:967
	_go_fuzz_dep_.CoverTab[81078]++

										return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:969
	// _ = "end of CoverTab[81078]"
}

func (cs *clientStream) finish(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:972
	_go_fuzz_dep_.CoverTab[81086]++
										if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:973
		_go_fuzz_dep_.CoverTab[81094]++

											err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:975
		// _ = "end of CoverTab[81094]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:976
		_go_fuzz_dep_.CoverTab[81095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:976
		// _ = "end of CoverTab[81095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:976
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:976
	// _ = "end of CoverTab[81086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:976
	_go_fuzz_dep_.CoverTab[81087]++
										cs.mu.Lock()
										if cs.finished {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:978
		_go_fuzz_dep_.CoverTab[81096]++
											cs.mu.Unlock()
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:980
		// _ = "end of CoverTab[81096]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:981
		_go_fuzz_dep_.CoverTab[81097]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:981
		// _ = "end of CoverTab[81097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:981
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:981
	// _ = "end of CoverTab[81087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:981
	_go_fuzz_dep_.CoverTab[81088]++
										cs.finished = true
										for _, onFinish := range cs.callInfo.onFinish {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:983
		_go_fuzz_dep_.CoverTab[81098]++
											onFinish(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:984
		// _ = "end of CoverTab[81098]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:985
	// _ = "end of CoverTab[81088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:985
	_go_fuzz_dep_.CoverTab[81089]++
										cs.commitAttemptLocked()
										if cs.attempt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:987
		_go_fuzz_dep_.CoverTab[81099]++
											cs.attempt.finish(err)

											if cs.attempt.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:990
			_go_fuzz_dep_.CoverTab[81100]++
												for _, o := range cs.opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:991
				_go_fuzz_dep_.CoverTab[81101]++
													o.after(cs.callInfo, cs.attempt)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:992
				// _ = "end of CoverTab[81101]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:993
			// _ = "end of CoverTab[81100]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:994
			_go_fuzz_dep_.CoverTab[81102]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:994
			// _ = "end of CoverTab[81102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:994
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:994
		// _ = "end of CoverTab[81099]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:995
		_go_fuzz_dep_.CoverTab[81103]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:995
		// _ = "end of CoverTab[81103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:995
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:995
		// _ = "end of CoverTab[81089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:995
		_go_fuzz_dep_.CoverTab[81090]++
											cs.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
	if len(cs.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
		_go_fuzz_dep_.CoverTab[81104]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
		return status.Code(err) == codes.Canceled
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
		// _ = "end of CoverTab[81104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1002
		_go_fuzz_dep_.CoverTab[81105]++
												c := &binarylog.Cancel{
			OnClientSide: true,
		}
		for _, binlog := range cs.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1006
			_go_fuzz_dep_.CoverTab[81106]++
													binlog.Log(cs.ctx, c)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1007
			// _ = "end of CoverTab[81106]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1008
		// _ = "end of CoverTab[81105]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1009
		_go_fuzz_dep_.CoverTab[81107]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1009
		// _ = "end of CoverTab[81107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1009
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1009
	// _ = "end of CoverTab[81090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1009
	_go_fuzz_dep_.CoverTab[81091]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1010
		_go_fuzz_dep_.CoverTab[81108]++
												cs.retryThrottler.successfulRPC()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1011
		// _ = "end of CoverTab[81108]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1012
		_go_fuzz_dep_.CoverTab[81109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1012
		// _ = "end of CoverTab[81109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1012
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1012
	// _ = "end of CoverTab[81091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1012
	_go_fuzz_dep_.CoverTab[81092]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1013
		_go_fuzz_dep_.CoverTab[81110]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1014
			_go_fuzz_dep_.CoverTab[81111]++
													cs.cc.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1015
			// _ = "end of CoverTab[81111]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1016
			_go_fuzz_dep_.CoverTab[81112]++
													cs.cc.incrCallsSucceeded()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1017
			// _ = "end of CoverTab[81112]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1018
		// _ = "end of CoverTab[81110]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1019
		_go_fuzz_dep_.CoverTab[81113]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1019
		// _ = "end of CoverTab[81113]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1019
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1019
	// _ = "end of CoverTab[81092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1019
	_go_fuzz_dep_.CoverTab[81093]++
											cs.cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1020
	// _ = "end of CoverTab[81093]"
}

func (a *csAttempt) sendMsg(m interface{}, hdr, payld, data []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1023
	_go_fuzz_dep_.CoverTab[81114]++
											cs := a.cs
											if a.trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1025
		_go_fuzz_dep_.CoverTab[81119]++
												a.mu.Lock()
												if a.trInfo.tr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1027
			_go_fuzz_dep_.CoverTab[81121]++
													a.trInfo.tr.LazyLog(&payload{sent: true, msg: m}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1028
			// _ = "end of CoverTab[81121]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1029
			_go_fuzz_dep_.CoverTab[81122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1029
			// _ = "end of CoverTab[81122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1029
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1029
		// _ = "end of CoverTab[81119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1029
		_go_fuzz_dep_.CoverTab[81120]++
												a.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1030
		// _ = "end of CoverTab[81120]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1031
		_go_fuzz_dep_.CoverTab[81123]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1031
		// _ = "end of CoverTab[81123]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1031
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1031
	// _ = "end of CoverTab[81114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1031
	_go_fuzz_dep_.CoverTab[81115]++
											if err := a.t.Write(a.s, hdr, payld, &transport.Options{Last: !cs.desc.ClientStreams}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1032
		_go_fuzz_dep_.CoverTab[81124]++
												if !cs.desc.ClientStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1033
			_go_fuzz_dep_.CoverTab[81126]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1037
			return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1037
			// _ = "end of CoverTab[81126]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1038
			_go_fuzz_dep_.CoverTab[81127]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1038
			// _ = "end of CoverTab[81127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1038
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1038
		// _ = "end of CoverTab[81124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1038
		_go_fuzz_dep_.CoverTab[81125]++
												return io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1039
		// _ = "end of CoverTab[81125]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1040
		_go_fuzz_dep_.CoverTab[81128]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1040
		// _ = "end of CoverTab[81128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1040
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1040
	// _ = "end of CoverTab[81115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1040
	_go_fuzz_dep_.CoverTab[81116]++
											for _, sh := range a.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1041
		_go_fuzz_dep_.CoverTab[81129]++
												sh.HandleRPC(a.ctx, outPayload(true, m, data, payld, time.Now()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1042
		// _ = "end of CoverTab[81129]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1043
	// _ = "end of CoverTab[81116]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1043
	_go_fuzz_dep_.CoverTab[81117]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1044
		_go_fuzz_dep_.CoverTab[81130]++
												a.t.IncrMsgSent()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1045
		// _ = "end of CoverTab[81130]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1046
		_go_fuzz_dep_.CoverTab[81131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1046
		// _ = "end of CoverTab[81131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1046
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1046
	// _ = "end of CoverTab[81117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1046
	_go_fuzz_dep_.CoverTab[81118]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1047
	// _ = "end of CoverTab[81118]"
}

func (a *csAttempt) recvMsg(m interface{}, payInfo *payloadInfo) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1050
	_go_fuzz_dep_.CoverTab[81132]++
											cs := a.cs
											if len(a.statsHandlers) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1052
		_go_fuzz_dep_.CoverTab[81142]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1052
		return payInfo == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1052
		// _ = "end of CoverTab[81142]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1052
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1052
		_go_fuzz_dep_.CoverTab[81143]++
												payInfo = &payloadInfo{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1053
		// _ = "end of CoverTab[81143]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1054
		_go_fuzz_dep_.CoverTab[81144]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1054
		// _ = "end of CoverTab[81144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1054
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1054
	// _ = "end of CoverTab[81132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1054
	_go_fuzz_dep_.CoverTab[81133]++

											if !a.decompSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1056
		_go_fuzz_dep_.CoverTab[81145]++

												if ct := a.s.RecvCompress(); ct != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1058
			_go_fuzz_dep_.CoverTab[81147]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1058
			return ct != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1058
			// _ = "end of CoverTab[81147]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1058
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1058
			_go_fuzz_dep_.CoverTab[81148]++
													if a.dc == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1059
				_go_fuzz_dep_.CoverTab[81149]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1059
				return a.dc.Type() != ct
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1059
				// _ = "end of CoverTab[81149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1059
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1059
				_go_fuzz_dep_.CoverTab[81150]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1062
				a.dc = nil
														a.decomp = encoding.GetCompressor(ct)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1063
				// _ = "end of CoverTab[81150]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1064
				_go_fuzz_dep_.CoverTab[81151]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1064
				// _ = "end of CoverTab[81151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1064
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1064
			// _ = "end of CoverTab[81148]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1065
			_go_fuzz_dep_.CoverTab[81152]++

													a.dc = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1067
			// _ = "end of CoverTab[81152]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1068
		// _ = "end of CoverTab[81145]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1068
		_go_fuzz_dep_.CoverTab[81146]++

												a.decompSet = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1070
		// _ = "end of CoverTab[81146]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1071
		_go_fuzz_dep_.CoverTab[81153]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1071
		// _ = "end of CoverTab[81153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1071
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1071
	// _ = "end of CoverTab[81133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1071
	_go_fuzz_dep_.CoverTab[81134]++
											err = recv(a.p, cs.codec, a.s, a.dc, m, *cs.callInfo.maxReceiveMessageSize, payInfo, a.decomp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1073
		_go_fuzz_dep_.CoverTab[81154]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1074
			_go_fuzz_dep_.CoverTab[81156]++
													if statusErr := a.s.Status().Err(); statusErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1075
				_go_fuzz_dep_.CoverTab[81158]++
														return statusErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1076
				// _ = "end of CoverTab[81158]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1077
				_go_fuzz_dep_.CoverTab[81159]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1077
				// _ = "end of CoverTab[81159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1077
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1077
			// _ = "end of CoverTab[81156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1077
			_go_fuzz_dep_.CoverTab[81157]++
													return io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1078
			// _ = "end of CoverTab[81157]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1079
			_go_fuzz_dep_.CoverTab[81160]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1079
			// _ = "end of CoverTab[81160]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1079
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1079
		// _ = "end of CoverTab[81154]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1079
		_go_fuzz_dep_.CoverTab[81155]++

												return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1081
		// _ = "end of CoverTab[81155]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1082
		_go_fuzz_dep_.CoverTab[81161]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1082
		// _ = "end of CoverTab[81161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1082
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1082
	// _ = "end of CoverTab[81134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1082
	_go_fuzz_dep_.CoverTab[81135]++
											if a.trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1083
		_go_fuzz_dep_.CoverTab[81162]++
												a.mu.Lock()
												if a.trInfo.tr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1085
			_go_fuzz_dep_.CoverTab[81164]++
													a.trInfo.tr.LazyLog(&payload{sent: false, msg: m}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1086
			// _ = "end of CoverTab[81164]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1087
			_go_fuzz_dep_.CoverTab[81165]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1087
			// _ = "end of CoverTab[81165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1087
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1087
		// _ = "end of CoverTab[81162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1087
		_go_fuzz_dep_.CoverTab[81163]++
												a.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1088
		// _ = "end of CoverTab[81163]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1089
		_go_fuzz_dep_.CoverTab[81166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1089
		// _ = "end of CoverTab[81166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1089
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1089
	// _ = "end of CoverTab[81135]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1089
	_go_fuzz_dep_.CoverTab[81136]++
											for _, sh := range a.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1090
		_go_fuzz_dep_.CoverTab[81167]++
												sh.HandleRPC(a.ctx, &stats.InPayload{
			Client:		true,
			RecvTime:	time.Now(),
			Payload:	m,

			Data:			payInfo.uncompressedBytes,
			WireLength:		payInfo.compressedLength + headerLen,
			CompressedLength:	payInfo.compressedLength,
			Length:			len(payInfo.uncompressedBytes),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1100
		// _ = "end of CoverTab[81167]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1101
	// _ = "end of CoverTab[81136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1101
	_go_fuzz_dep_.CoverTab[81137]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1102
		_go_fuzz_dep_.CoverTab[81168]++
												a.t.IncrMsgRecv()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1103
		// _ = "end of CoverTab[81168]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1104
		_go_fuzz_dep_.CoverTab[81169]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1104
		// _ = "end of CoverTab[81169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1104
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1104
	// _ = "end of CoverTab[81137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1104
	_go_fuzz_dep_.CoverTab[81138]++
											if cs.desc.ServerStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1105
		_go_fuzz_dep_.CoverTab[81170]++

												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1107
		// _ = "end of CoverTab[81170]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1108
		_go_fuzz_dep_.CoverTab[81171]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1108
		// _ = "end of CoverTab[81171]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1108
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1108
	// _ = "end of CoverTab[81138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1108
	_go_fuzz_dep_.CoverTab[81139]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1111
	err = recv(a.p, cs.codec, a.s, a.dc, m, *cs.callInfo.maxReceiveMessageSize, nil, a.decomp)
	if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1112
		_go_fuzz_dep_.CoverTab[81172]++
												return toRPCErr(errors.New("grpc: client streaming protocol violation: get <nil>, want <EOF>"))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1113
		// _ = "end of CoverTab[81172]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1114
		_go_fuzz_dep_.CoverTab[81173]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1114
		// _ = "end of CoverTab[81173]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1114
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1114
	// _ = "end of CoverTab[81139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1114
	_go_fuzz_dep_.CoverTab[81140]++
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1115
		_go_fuzz_dep_.CoverTab[81174]++
												return a.s.Status().Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1116
		// _ = "end of CoverTab[81174]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1117
		_go_fuzz_dep_.CoverTab[81175]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1117
		// _ = "end of CoverTab[81175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1117
	// _ = "end of CoverTab[81140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1117
	_go_fuzz_dep_.CoverTab[81141]++
											return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1118
	// _ = "end of CoverTab[81141]"
}

func (a *csAttempt) finish(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1121
	_go_fuzz_dep_.CoverTab[81176]++
											a.mu.Lock()
											if a.finished {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1123
		_go_fuzz_dep_.CoverTab[81183]++
												a.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1125
		// _ = "end of CoverTab[81183]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1126
		_go_fuzz_dep_.CoverTab[81184]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1126
		// _ = "end of CoverTab[81184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1126
	// _ = "end of CoverTab[81176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1126
	_go_fuzz_dep_.CoverTab[81177]++
											a.finished = true
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1128
		_go_fuzz_dep_.CoverTab[81185]++

												err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1130
		// _ = "end of CoverTab[81185]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1131
		_go_fuzz_dep_.CoverTab[81186]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1131
		// _ = "end of CoverTab[81186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1131
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1131
	// _ = "end of CoverTab[81177]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1131
	_go_fuzz_dep_.CoverTab[81178]++
											var tr metadata.MD
											if a.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1133
		_go_fuzz_dep_.CoverTab[81187]++
												a.t.CloseStream(a.s, err)
												tr = a.s.Trailer()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1135
		// _ = "end of CoverTab[81187]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1136
		_go_fuzz_dep_.CoverTab[81188]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1136
		// _ = "end of CoverTab[81188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1136
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1136
	// _ = "end of CoverTab[81178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1136
	_go_fuzz_dep_.CoverTab[81179]++

											if a.pickResult.Done != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1138
		_go_fuzz_dep_.CoverTab[81189]++
												br := false
												if a.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1140
			_go_fuzz_dep_.CoverTab[81191]++
													br = a.s.BytesReceived()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1141
			// _ = "end of CoverTab[81191]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1142
			_go_fuzz_dep_.CoverTab[81192]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1142
			// _ = "end of CoverTab[81192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1142
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1142
		// _ = "end of CoverTab[81189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1142
		_go_fuzz_dep_.CoverTab[81190]++
												a.pickResult.Done(balancer.DoneInfo{
			Err:		err,
			Trailer:	tr,
			BytesSent:	a.s != nil,
			BytesReceived:	br,
			ServerLoad:	balancerload.Parse(tr),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1149
		// _ = "end of CoverTab[81190]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1150
		_go_fuzz_dep_.CoverTab[81193]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1150
		// _ = "end of CoverTab[81193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1150
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1150
	// _ = "end of CoverTab[81179]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1150
	_go_fuzz_dep_.CoverTab[81180]++
											for _, sh := range a.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1151
		_go_fuzz_dep_.CoverTab[81194]++
												end := &stats.End{
			Client:		true,
			BeginTime:	a.beginTime,
			EndTime:	time.Now(),
			Trailer:	tr,
			Error:		err,
		}
												sh.HandleRPC(a.ctx, end)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1159
		// _ = "end of CoverTab[81194]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1160
	// _ = "end of CoverTab[81180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1160
	_go_fuzz_dep_.CoverTab[81181]++
											if a.trInfo != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1161
		_go_fuzz_dep_.CoverTab[81195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1161
		return a.trInfo.tr != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1161
		// _ = "end of CoverTab[81195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1161
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1161
		_go_fuzz_dep_.CoverTab[81196]++
												if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1162
			_go_fuzz_dep_.CoverTab[81198]++
													a.trInfo.tr.LazyPrintf("RPC: [OK]")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1163
			// _ = "end of CoverTab[81198]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1164
			_go_fuzz_dep_.CoverTab[81199]++
													a.trInfo.tr.LazyPrintf("RPC: [%v]", err)
													a.trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1166
			// _ = "end of CoverTab[81199]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1167
		// _ = "end of CoverTab[81196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1167
		_go_fuzz_dep_.CoverTab[81197]++
												a.trInfo.tr.Finish()
												a.trInfo.tr = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1169
		// _ = "end of CoverTab[81197]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1170
		_go_fuzz_dep_.CoverTab[81200]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1170
		// _ = "end of CoverTab[81200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1170
	// _ = "end of CoverTab[81181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1170
	_go_fuzz_dep_.CoverTab[81182]++
											a.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1171
	// _ = "end of CoverTab[81182]"
}

// newClientStream creates a ClientStream with the specified transport, on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// given addrConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// It's expected that the given transport is either the same one in addrConn, or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// is already closed. To avoid race, transport is specified separately, instead
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// of using ac.transpot.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// Main difference between this and ClientConn.NewStream:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// - no retry
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// - no service config (or wait for service config)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1174
// - no tracing or stats
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1185
func newNonRetryClientStream(ctx context.Context, desc *StreamDesc, method string, t transport.ClientTransport, ac *addrConn, opts ...CallOption) (_ ClientStream, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1185
	_go_fuzz_dep_.CoverTab[81201]++
											if t == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1186
		_go_fuzz_dep_.CoverTab[81210]++

												return nil, errors.New("transport provided is nil")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1188
		// _ = "end of CoverTab[81210]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1189
		_go_fuzz_dep_.CoverTab[81211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1189
		// _ = "end of CoverTab[81211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1189
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1189
	// _ = "end of CoverTab[81201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1189
	_go_fuzz_dep_.CoverTab[81202]++

											c := &callInfo{}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1198
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1199
		_go_fuzz_dep_.CoverTab[81212]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1200
			_go_fuzz_dep_.CoverTab[81213]++
													cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1201
			// _ = "end of CoverTab[81213]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1202
			_go_fuzz_dep_.CoverTab[81214]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1202
			// _ = "end of CoverTab[81214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1202
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1202
		// _ = "end of CoverTab[81212]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1203
	// _ = "end of CoverTab[81202]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1203
	_go_fuzz_dep_.CoverTab[81203]++

											for _, o := range opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1205
		_go_fuzz_dep_.CoverTab[81215]++
												if err := o.before(c); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1206
			_go_fuzz_dep_.CoverTab[81216]++
													return nil, toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1207
			// _ = "end of CoverTab[81216]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1208
			_go_fuzz_dep_.CoverTab[81217]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1208
			// _ = "end of CoverTab[81217]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1208
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1208
		// _ = "end of CoverTab[81215]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1209
	// _ = "end of CoverTab[81203]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1209
	_go_fuzz_dep_.CoverTab[81204]++
											c.maxReceiveMessageSize = getMaxSize(nil, c.maxReceiveMessageSize, defaultClientMaxReceiveMessageSize)
											c.maxSendMessageSize = getMaxSize(nil, c.maxSendMessageSize, defaultServerMaxSendMessageSize)
											if err := setCallInfoCodec(c); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1212
		_go_fuzz_dep_.CoverTab[81218]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1213
		// _ = "end of CoverTab[81218]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1214
		_go_fuzz_dep_.CoverTab[81219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1214
		// _ = "end of CoverTab[81219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1214
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1214
	// _ = "end of CoverTab[81204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1214
	_go_fuzz_dep_.CoverTab[81205]++

											callHdr := &transport.CallHdr{
		Host:		ac.cc.authority,
		Method:		method,
		ContentSubtype:	c.contentSubtype,
	}

	// Set our outgoing compression according to the UseCompressor CallOption, if
	// set.  In that case, also find the compressor from the encoding package.
	// Otherwise, use the compressor configured by the WithCompressor DialOption,
	// if set.
	var cp Compressor
	var comp encoding.Compressor
	if ct := c.compressorType; ct != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1228
		_go_fuzz_dep_.CoverTab[81220]++
												callHdr.SendCompress = ct
												if ct != encoding.Identity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1230
			_go_fuzz_dep_.CoverTab[81221]++
													comp = encoding.GetCompressor(ct)
													if comp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1232
				_go_fuzz_dep_.CoverTab[81222]++
														return nil, status.Errorf(codes.Internal, "grpc: Compressor is not installed for requested grpc-encoding %q", ct)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1233
				// _ = "end of CoverTab[81222]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1234
				_go_fuzz_dep_.CoverTab[81223]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1234
				// _ = "end of CoverTab[81223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1234
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1234
			// _ = "end of CoverTab[81221]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1235
			_go_fuzz_dep_.CoverTab[81224]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1235
			// _ = "end of CoverTab[81224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1235
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1235
		// _ = "end of CoverTab[81220]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1236
		_go_fuzz_dep_.CoverTab[81225]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1236
		if ac.cc.dopts.cp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1236
			_go_fuzz_dep_.CoverTab[81226]++
													callHdr.SendCompress = ac.cc.dopts.cp.Type()
													cp = ac.cc.dopts.cp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1238
			// _ = "end of CoverTab[81226]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
			_go_fuzz_dep_.CoverTab[81227]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
			// _ = "end of CoverTab[81227]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
		// _ = "end of CoverTab[81225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
	// _ = "end of CoverTab[81205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1239
	_go_fuzz_dep_.CoverTab[81206]++
											if c.creds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1240
		_go_fuzz_dep_.CoverTab[81228]++
												callHdr.Creds = c.creds
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1241
		// _ = "end of CoverTab[81228]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1242
		_go_fuzz_dep_.CoverTab[81229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1242
		// _ = "end of CoverTab[81229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1242
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1242
	// _ = "end of CoverTab[81206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1242
	_go_fuzz_dep_.CoverTab[81207]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1245
	as := &addrConnStream{
		callHdr:	callHdr,
		ac:		ac,
		ctx:		ctx,
		cancel:		cancel,
		opts:		opts,
		callInfo:	c,
		desc:		desc,
		codec:		c.codec,
		cp:		cp,
		comp:		comp,
		t:		t,
	}

	s, err := as.t.NewStream(as.ctx, as.callHdr)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1260
		_go_fuzz_dep_.CoverTab[81230]++
												err = toRPCErr(err)
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1262
		// _ = "end of CoverTab[81230]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1263
		_go_fuzz_dep_.CoverTab[81231]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1263
		// _ = "end of CoverTab[81231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1263
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1263
	// _ = "end of CoverTab[81207]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1263
	_go_fuzz_dep_.CoverTab[81208]++
											as.s = s
											as.p = &parser{r: s}
											ac.incrCallsStarted()
											if desc != unaryStreamDesc {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1267
		_go_fuzz_dep_.CoverTab[81232]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1267
		_curRoutineNum98_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1267
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum98_)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
		go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
			_go_fuzz_dep_.CoverTab[81233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
			defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
				_go_fuzz_dep_.CoverTab[81234]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum98_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
				// _ = "end of CoverTab[81234]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1273
			}()
													select {
			case <-ac.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1275
				_go_fuzz_dep_.CoverTab[81235]++
														as.finish(status.Error(codes.Canceled, "grpc: the SubConn is closing"))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1276
				// _ = "end of CoverTab[81235]"
			case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1277
				_go_fuzz_dep_.CoverTab[81236]++
														as.finish(toRPCErr(ctx.Err()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1278
				// _ = "end of CoverTab[81236]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1279
			// _ = "end of CoverTab[81233]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1280
		// _ = "end of CoverTab[81232]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1281
		_go_fuzz_dep_.CoverTab[81237]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1281
		// _ = "end of CoverTab[81237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1281
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1281
	// _ = "end of CoverTab[81208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1281
	_go_fuzz_dep_.CoverTab[81209]++
											return as, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1282
	// _ = "end of CoverTab[81209]"
}

type addrConnStream struct {
	s		*transport.Stream
	ac		*addrConn
	callHdr		*transport.CallHdr
	cancel		context.CancelFunc
	opts		[]CallOption
	callInfo	*callInfo
	t		transport.ClientTransport
	ctx		context.Context
	sentLast	bool
	desc		*StreamDesc
	codec		baseCodec
	cp		Compressor
	comp		encoding.Compressor
	decompSet	bool
	dc		Decompressor
	decomp		encoding.Compressor
	p		*parser
	mu		sync.Mutex
	finished	bool
}

func (as *addrConnStream) Header() (metadata.MD, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1307
	_go_fuzz_dep_.CoverTab[81238]++
											m, err := as.s.Header()
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1309
		_go_fuzz_dep_.CoverTab[81240]++
												as.finish(toRPCErr(err))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1310
		// _ = "end of CoverTab[81240]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1311
		_go_fuzz_dep_.CoverTab[81241]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1311
		// _ = "end of CoverTab[81241]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1311
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1311
	// _ = "end of CoverTab[81238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1311
	_go_fuzz_dep_.CoverTab[81239]++
											return m, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1312
	// _ = "end of CoverTab[81239]"
}

func (as *addrConnStream) Trailer() metadata.MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1315
	_go_fuzz_dep_.CoverTab[81242]++
											return as.s.Trailer()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1316
	// _ = "end of CoverTab[81242]"
}

func (as *addrConnStream) CloseSend() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1319
	_go_fuzz_dep_.CoverTab[81243]++
											if as.sentLast {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1320
		_go_fuzz_dep_.CoverTab[81245]++

												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1322
		// _ = "end of CoverTab[81245]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1323
		_go_fuzz_dep_.CoverTab[81246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1323
		// _ = "end of CoverTab[81246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1323
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1323
	// _ = "end of CoverTab[81243]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1323
	_go_fuzz_dep_.CoverTab[81244]++
											as.sentLast = true

											as.t.Write(as.s, nil, nil, &transport.Options{Last: true})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1331
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1331
	// _ = "end of CoverTab[81244]"
}

func (as *addrConnStream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1334
	_go_fuzz_dep_.CoverTab[81247]++
											return as.s.Context()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1335
	// _ = "end of CoverTab[81247]"
}

func (as *addrConnStream) SendMsg(m interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1338
	_go_fuzz_dep_.CoverTab[81248]++
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1339
		_go_fuzz_dep_.CoverTab[81256]++
												if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1340
			_go_fuzz_dep_.CoverTab[81257]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1340
			return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1340
			// _ = "end of CoverTab[81257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1340
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1340
			_go_fuzz_dep_.CoverTab[81258]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1346
			as.finish(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1346
			// _ = "end of CoverTab[81258]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1347
			_go_fuzz_dep_.CoverTab[81259]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1347
			// _ = "end of CoverTab[81259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1347
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1347
		// _ = "end of CoverTab[81256]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1348
	// _ = "end of CoverTab[81248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1348
	_go_fuzz_dep_.CoverTab[81249]++
											if as.sentLast {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1349
		_go_fuzz_dep_.CoverTab[81260]++
												return status.Errorf(codes.Internal, "SendMsg called after CloseSend")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1350
		// _ = "end of CoverTab[81260]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1351
		_go_fuzz_dep_.CoverTab[81261]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1351
		// _ = "end of CoverTab[81261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1351
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1351
	// _ = "end of CoverTab[81249]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1351
	_go_fuzz_dep_.CoverTab[81250]++
											if !as.desc.ClientStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1352
		_go_fuzz_dep_.CoverTab[81262]++
												as.sentLast = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1353
		// _ = "end of CoverTab[81262]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1354
		_go_fuzz_dep_.CoverTab[81263]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1354
		// _ = "end of CoverTab[81263]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1354
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1354
	// _ = "end of CoverTab[81250]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1354
	_go_fuzz_dep_.CoverTab[81251]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1357
	hdr, payld, _, err := prepareMsg(m, as.codec, as.cp, as.comp)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1358
		_go_fuzz_dep_.CoverTab[81264]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1359
		// _ = "end of CoverTab[81264]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1360
		_go_fuzz_dep_.CoverTab[81265]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1360
		// _ = "end of CoverTab[81265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1360
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1360
	// _ = "end of CoverTab[81251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1360
	_go_fuzz_dep_.CoverTab[81252]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1363
	if len(payld) > *as.callInfo.maxSendMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1363
		_go_fuzz_dep_.CoverTab[81266]++
												return status.Errorf(codes.ResourceExhausted, "trying to send message larger than max (%d vs. %d)", len(payld), *as.callInfo.maxSendMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1364
		// _ = "end of CoverTab[81266]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1365
		_go_fuzz_dep_.CoverTab[81267]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1365
		// _ = "end of CoverTab[81267]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1365
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1365
	// _ = "end of CoverTab[81252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1365
	_go_fuzz_dep_.CoverTab[81253]++

											if err := as.t.Write(as.s, hdr, payld, &transport.Options{Last: !as.desc.ClientStreams}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1367
		_go_fuzz_dep_.CoverTab[81268]++
												if !as.desc.ClientStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1368
			_go_fuzz_dep_.CoverTab[81270]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1372
			return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1372
			// _ = "end of CoverTab[81270]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1373
			_go_fuzz_dep_.CoverTab[81271]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1373
			// _ = "end of CoverTab[81271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1373
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1373
		// _ = "end of CoverTab[81268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1373
		_go_fuzz_dep_.CoverTab[81269]++
												return io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1374
		// _ = "end of CoverTab[81269]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1375
		_go_fuzz_dep_.CoverTab[81272]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1375
		// _ = "end of CoverTab[81272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1375
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1375
	// _ = "end of CoverTab[81253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1375
	_go_fuzz_dep_.CoverTab[81254]++

											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1377
		_go_fuzz_dep_.CoverTab[81273]++
												as.t.IncrMsgSent()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1378
		// _ = "end of CoverTab[81273]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1379
		_go_fuzz_dep_.CoverTab[81274]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1379
		// _ = "end of CoverTab[81274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1379
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1379
	// _ = "end of CoverTab[81254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1379
	_go_fuzz_dep_.CoverTab[81255]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1380
	// _ = "end of CoverTab[81255]"
}

func (as *addrConnStream) RecvMsg(m interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1383
	_go_fuzz_dep_.CoverTab[81275]++
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1384
		_go_fuzz_dep_.CoverTab[81283]++
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1385
			_go_fuzz_dep_.CoverTab[81284]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1385
			return !as.desc.ServerStreams
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1385
			// _ = "end of CoverTab[81284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1385
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1385
			_go_fuzz_dep_.CoverTab[81285]++

													as.finish(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1387
			// _ = "end of CoverTab[81285]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1388
			_go_fuzz_dep_.CoverTab[81286]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1388
			// _ = "end of CoverTab[81286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1388
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1388
		// _ = "end of CoverTab[81283]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1389
	// _ = "end of CoverTab[81275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1389
	_go_fuzz_dep_.CoverTab[81276]++

											if !as.decompSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1391
		_go_fuzz_dep_.CoverTab[81287]++

												if ct := as.s.RecvCompress(); ct != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1393
			_go_fuzz_dep_.CoverTab[81289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1393
			return ct != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1393
			// _ = "end of CoverTab[81289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1393
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1393
			_go_fuzz_dep_.CoverTab[81290]++
													if as.dc == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1394
				_go_fuzz_dep_.CoverTab[81291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1394
				return as.dc.Type() != ct
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1394
				// _ = "end of CoverTab[81291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1394
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1394
				_go_fuzz_dep_.CoverTab[81292]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1397
				as.dc = nil
														as.decomp = encoding.GetCompressor(ct)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1398
				// _ = "end of CoverTab[81292]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1399
				_go_fuzz_dep_.CoverTab[81293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1399
				// _ = "end of CoverTab[81293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1399
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1399
			// _ = "end of CoverTab[81290]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1400
			_go_fuzz_dep_.CoverTab[81294]++

													as.dc = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1402
			// _ = "end of CoverTab[81294]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1403
		// _ = "end of CoverTab[81287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1403
		_go_fuzz_dep_.CoverTab[81288]++

												as.decompSet = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1405
		// _ = "end of CoverTab[81288]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1406
		_go_fuzz_dep_.CoverTab[81295]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1406
		// _ = "end of CoverTab[81295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1406
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1406
	// _ = "end of CoverTab[81276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1406
	_go_fuzz_dep_.CoverTab[81277]++
											err = recv(as.p, as.codec, as.s, as.dc, m, *as.callInfo.maxReceiveMessageSize, nil, as.decomp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1408
		_go_fuzz_dep_.CoverTab[81296]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1409
			_go_fuzz_dep_.CoverTab[81298]++
													if statusErr := as.s.Status().Err(); statusErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1410
				_go_fuzz_dep_.CoverTab[81300]++
														return statusErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1411
				// _ = "end of CoverTab[81300]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1412
				_go_fuzz_dep_.CoverTab[81301]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1412
				// _ = "end of CoverTab[81301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1412
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1412
			// _ = "end of CoverTab[81298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1412
			_go_fuzz_dep_.CoverTab[81299]++
													return io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1413
			// _ = "end of CoverTab[81299]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1414
			_go_fuzz_dep_.CoverTab[81302]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1414
			// _ = "end of CoverTab[81302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1414
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1414
		// _ = "end of CoverTab[81296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1414
		_go_fuzz_dep_.CoverTab[81297]++
												return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1415
		// _ = "end of CoverTab[81297]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1416
		_go_fuzz_dep_.CoverTab[81303]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1416
		// _ = "end of CoverTab[81303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1416
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1416
	// _ = "end of CoverTab[81277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1416
	_go_fuzz_dep_.CoverTab[81278]++

											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1418
		_go_fuzz_dep_.CoverTab[81304]++
												as.t.IncrMsgRecv()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1419
		// _ = "end of CoverTab[81304]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1420
		_go_fuzz_dep_.CoverTab[81305]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1420
		// _ = "end of CoverTab[81305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1420
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1420
	// _ = "end of CoverTab[81278]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1420
	_go_fuzz_dep_.CoverTab[81279]++
											if as.desc.ServerStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1421
		_go_fuzz_dep_.CoverTab[81306]++

												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1423
		// _ = "end of CoverTab[81306]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1424
		_go_fuzz_dep_.CoverTab[81307]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1424
		// _ = "end of CoverTab[81307]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1424
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1424
	// _ = "end of CoverTab[81279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1424
	_go_fuzz_dep_.CoverTab[81280]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1428
	err = recv(as.p, as.codec, as.s, as.dc, m, *as.callInfo.maxReceiveMessageSize, nil, as.decomp)
	if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1429
		_go_fuzz_dep_.CoverTab[81308]++
												return toRPCErr(errors.New("grpc: client streaming protocol violation: get <nil>, want <EOF>"))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1430
		// _ = "end of CoverTab[81308]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1431
		_go_fuzz_dep_.CoverTab[81309]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1431
		// _ = "end of CoverTab[81309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1431
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1431
	// _ = "end of CoverTab[81280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1431
	_go_fuzz_dep_.CoverTab[81281]++
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1432
		_go_fuzz_dep_.CoverTab[81310]++
												return as.s.Status().Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1433
		// _ = "end of CoverTab[81310]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1434
		_go_fuzz_dep_.CoverTab[81311]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1434
		// _ = "end of CoverTab[81311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1434
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1434
	// _ = "end of CoverTab[81281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1434
	_go_fuzz_dep_.CoverTab[81282]++
											return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1435
	// _ = "end of CoverTab[81282]"
}

func (as *addrConnStream) finish(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1438
	_go_fuzz_dep_.CoverTab[81312]++
											as.mu.Lock()
											if as.finished {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1440
		_go_fuzz_dep_.CoverTab[81317]++
												as.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1442
		// _ = "end of CoverTab[81317]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1443
		_go_fuzz_dep_.CoverTab[81318]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1443
		// _ = "end of CoverTab[81318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1443
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1443
	// _ = "end of CoverTab[81312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1443
	_go_fuzz_dep_.CoverTab[81313]++
											as.finished = true
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1445
		_go_fuzz_dep_.CoverTab[81319]++

												err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1447
		// _ = "end of CoverTab[81319]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1448
		_go_fuzz_dep_.CoverTab[81320]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1448
		// _ = "end of CoverTab[81320]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1448
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1448
	// _ = "end of CoverTab[81313]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1448
	_go_fuzz_dep_.CoverTab[81314]++
											if as.s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1449
		_go_fuzz_dep_.CoverTab[81321]++
												as.t.CloseStream(as.s, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1450
		// _ = "end of CoverTab[81321]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1451
		_go_fuzz_dep_.CoverTab[81322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1451
		// _ = "end of CoverTab[81322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1451
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1451
	// _ = "end of CoverTab[81314]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1451
	_go_fuzz_dep_.CoverTab[81315]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1453
		_go_fuzz_dep_.CoverTab[81323]++
												as.ac.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1454
		// _ = "end of CoverTab[81323]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1455
		_go_fuzz_dep_.CoverTab[81324]++
												as.ac.incrCallsSucceeded()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1456
		// _ = "end of CoverTab[81324]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1457
	// _ = "end of CoverTab[81315]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1457
	_go_fuzz_dep_.CoverTab[81316]++
											as.cancel()
											as.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1459
	// _ = "end of CoverTab[81316]"
}

// ServerStream defines the server-side behavior of a streaming RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1462
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1462
// Errors returned from ServerStream methods are compatible with the status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1462
// package.  However, the status code will often not match the RPC status as
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1462
// seen by the client application, and therefore, should not be relied upon for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1462
// this purpose.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1468
type ServerStream interface {
	// SetHeader sets the header metadata. It may be called multiple times.
	// When call multiple times, all the provided metadata will be merged.
	// All the metadata will be sent out when one of the following happens:
	//  - ServerStream.SendHeader() is called;
	//  - The first response is sent out;
	//  - An RPC status is sent out (error or success).
	SetHeader(metadata.MD) error
	// SendHeader sends the header metadata.
	// The provided md and headers set by SetHeader() will be sent.
	// It fails if called multiple times.
	SendHeader(metadata.MD) error
	// SetTrailer sets the trailer metadata which will be sent with the RPC status.
	// When called more than once, all the provided metadata will be merged.
	SetTrailer(metadata.MD)
	// Context returns the context for this stream.
	Context() context.Context
	// SendMsg sends a message. On error, SendMsg aborts the stream and the
	// error is returned directly.
	//
	// SendMsg blocks until:
	//   - There is sufficient flow control to schedule m with the transport, or
	//   - The stream is done, or
	//   - The stream breaks.
	//
	// SendMsg does not wait until the message is received by the client. An
	// untimely stream closure may result in lost messages.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not safe
	// to call SendMsg on the same stream in different goroutines.
	//
	// It is not safe to modify the message after calling SendMsg. Tracing
	// libraries and stats handlers may use the message lazily.
	SendMsg(m interface{}) error
	// RecvMsg blocks until it receives a message into m or the stream is
	// done. It returns io.EOF when the client has performed a CloseSend. On
	// any non-EOF error, the stream is aborted and the error contains the
	// RPC status.
	//
	// It is safe to have a goroutine calling SendMsg and another goroutine
	// calling RecvMsg on the same stream at the same time, but it is not
	// safe to call RecvMsg on the same stream in different goroutines.
	RecvMsg(m interface{}) error
}

// serverStream implements a server side Stream.
type serverStream struct {
	ctx	context.Context
	t	transport.ServerTransport
	s	*transport.Stream
	p	*parser
	codec	baseCodec

	cp	Compressor
	dc	Decompressor
	comp	encoding.Compressor
	decomp	encoding.Compressor

	sendCompressorName	string

	maxReceiveMessageSize	int
	maxSendMessageSize	int
	trInfo			*traceInfo

	statsHandler	[]stats.Handler

	binlogs	[]binarylog.MethodLogger
	// serverHeaderBinlogged indicates whether server header has been logged. It
	// will happen when one of the following two happens: stream.SendHeader(),
	// stream.Send().
	//
	// It's only checked in send and sendHeader, doesn't need to be
	// synchronized.
	serverHeaderBinlogged	bool

	mu	sync.Mutex	// protects trInfo.tr after the service handler runs.
}

func (ss *serverStream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1547
	_go_fuzz_dep_.CoverTab[81325]++
											return ss.ctx
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1548
	// _ = "end of CoverTab[81325]"
}

func (ss *serverStream) SetHeader(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1551
	_go_fuzz_dep_.CoverTab[81326]++
											if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1552
		_go_fuzz_dep_.CoverTab[81329]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1553
		// _ = "end of CoverTab[81329]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1554
		_go_fuzz_dep_.CoverTab[81330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1554
		// _ = "end of CoverTab[81330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1554
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1554
	// _ = "end of CoverTab[81326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1554
	_go_fuzz_dep_.CoverTab[81327]++
											err := imetadata.Validate(md)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1556
		_go_fuzz_dep_.CoverTab[81331]++
												return status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1557
		// _ = "end of CoverTab[81331]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1558
		_go_fuzz_dep_.CoverTab[81332]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1558
		// _ = "end of CoverTab[81332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1558
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1558
	// _ = "end of CoverTab[81327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1558
	_go_fuzz_dep_.CoverTab[81328]++
											return ss.s.SetHeader(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1559
	// _ = "end of CoverTab[81328]"
}

func (ss *serverStream) SendHeader(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1562
	_go_fuzz_dep_.CoverTab[81333]++
											err := imetadata.Validate(md)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1564
		_go_fuzz_dep_.CoverTab[81336]++
												return status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1565
		// _ = "end of CoverTab[81336]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1566
		_go_fuzz_dep_.CoverTab[81337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1566
		// _ = "end of CoverTab[81337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1566
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1566
	// _ = "end of CoverTab[81333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1566
	_go_fuzz_dep_.CoverTab[81334]++

											err = ss.t.WriteHeader(ss.s, md)
											if len(ss.binlogs) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1569
		_go_fuzz_dep_.CoverTab[81338]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1569
		return !ss.serverHeaderBinlogged
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1569
		// _ = "end of CoverTab[81338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1569
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1569
		_go_fuzz_dep_.CoverTab[81339]++
												h, _ := ss.s.Header()
												sh := &binarylog.ServerHeader{
			Header: h,
		}
		ss.serverHeaderBinlogged = true
		for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1575
			_go_fuzz_dep_.CoverTab[81340]++
													binlog.Log(ss.ctx, sh)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1576
			// _ = "end of CoverTab[81340]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1577
		// _ = "end of CoverTab[81339]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1578
		_go_fuzz_dep_.CoverTab[81341]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1578
		// _ = "end of CoverTab[81341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1578
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1578
	// _ = "end of CoverTab[81334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1578
	_go_fuzz_dep_.CoverTab[81335]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1579
	// _ = "end of CoverTab[81335]"
}

func (ss *serverStream) SetTrailer(md metadata.MD) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1582
	_go_fuzz_dep_.CoverTab[81342]++
											if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1583
		_go_fuzz_dep_.CoverTab[81345]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1584
		// _ = "end of CoverTab[81345]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1585
		_go_fuzz_dep_.CoverTab[81346]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1585
		// _ = "end of CoverTab[81346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1585
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1585
	// _ = "end of CoverTab[81342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1585
	_go_fuzz_dep_.CoverTab[81343]++
											if err := imetadata.Validate(md); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1586
		_go_fuzz_dep_.CoverTab[81347]++
												logger.Errorf("stream: failed to validate md when setting trailer, err: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1587
		// _ = "end of CoverTab[81347]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1588
		_go_fuzz_dep_.CoverTab[81348]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1588
		// _ = "end of CoverTab[81348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1588
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1588
	// _ = "end of CoverTab[81343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1588
	_go_fuzz_dep_.CoverTab[81344]++
											ss.s.SetTrailer(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1589
	// _ = "end of CoverTab[81344]"
}

func (ss *serverStream) SendMsg(m interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1592
	_go_fuzz_dep_.CoverTab[81349]++
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1593
		_go_fuzz_dep_.CoverTab[81357]++
												if ss.trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1594
			_go_fuzz_dep_.CoverTab[81360]++
													ss.mu.Lock()
													if ss.trInfo.tr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1596
				_go_fuzz_dep_.CoverTab[81362]++
														if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1597
					_go_fuzz_dep_.CoverTab[81363]++
															ss.trInfo.tr.LazyLog(&payload{sent: true, msg: m}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1598
					// _ = "end of CoverTab[81363]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1599
					_go_fuzz_dep_.CoverTab[81364]++
															ss.trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
															ss.trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1601
					// _ = "end of CoverTab[81364]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1602
				// _ = "end of CoverTab[81362]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1603
				_go_fuzz_dep_.CoverTab[81365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1603
				// _ = "end of CoverTab[81365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1603
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1603
			// _ = "end of CoverTab[81360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1603
			_go_fuzz_dep_.CoverTab[81361]++
													ss.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1604
			// _ = "end of CoverTab[81361]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1605
			_go_fuzz_dep_.CoverTab[81366]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1605
			// _ = "end of CoverTab[81366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1605
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1605
		// _ = "end of CoverTab[81357]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1605
		_go_fuzz_dep_.CoverTab[81358]++
												if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1606
			_go_fuzz_dep_.CoverTab[81367]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1606
			return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1606
			// _ = "end of CoverTab[81367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1606
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1606
			_go_fuzz_dep_.CoverTab[81368]++
													st, _ := status.FromError(toRPCErr(err))
													ss.t.WriteStatus(ss.s, st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1608
			// _ = "end of CoverTab[81368]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
			_go_fuzz_dep_.CoverTab[81369]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
			// _ = "end of CoverTab[81369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
		// _ = "end of CoverTab[81358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1615
		_go_fuzz_dep_.CoverTab[81359]++
												if channelz.IsOn() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1616
			_go_fuzz_dep_.CoverTab[81370]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1616
			return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1616
			// _ = "end of CoverTab[81370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1616
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1616
			_go_fuzz_dep_.CoverTab[81371]++
													ss.t.IncrMsgSent()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1617
			// _ = "end of CoverTab[81371]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1618
			_go_fuzz_dep_.CoverTab[81372]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1618
			// _ = "end of CoverTab[81372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1618
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1618
		// _ = "end of CoverTab[81359]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1619
	// _ = "end of CoverTab[81349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1619
	_go_fuzz_dep_.CoverTab[81350]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1623
	if sendCompressorsName := ss.s.SendCompress(); sendCompressorsName != ss.sendCompressorName {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1623
		_go_fuzz_dep_.CoverTab[81373]++
												ss.comp = encoding.GetCompressor(sendCompressorsName)
												ss.sendCompressorName = sendCompressorsName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1625
		// _ = "end of CoverTab[81373]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1626
		_go_fuzz_dep_.CoverTab[81374]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1626
		// _ = "end of CoverTab[81374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1626
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1626
	// _ = "end of CoverTab[81350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1626
	_go_fuzz_dep_.CoverTab[81351]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1629
	hdr, payload, data, err := prepareMsg(m, ss.codec, ss.cp, ss.comp)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1630
		_go_fuzz_dep_.CoverTab[81375]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1631
		// _ = "end of CoverTab[81375]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1632
		_go_fuzz_dep_.CoverTab[81376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1632
		// _ = "end of CoverTab[81376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1632
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1632
	// _ = "end of CoverTab[81351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1632
	_go_fuzz_dep_.CoverTab[81352]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1635
	if len(payload) > ss.maxSendMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1635
		_go_fuzz_dep_.CoverTab[81377]++
												return status.Errorf(codes.ResourceExhausted, "trying to send message larger than max (%d vs. %d)", len(payload), ss.maxSendMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1636
		// _ = "end of CoverTab[81377]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1637
		_go_fuzz_dep_.CoverTab[81378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1637
		// _ = "end of CoverTab[81378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1637
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1637
	// _ = "end of CoverTab[81352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1637
	_go_fuzz_dep_.CoverTab[81353]++
											if err := ss.t.Write(ss.s, hdr, payload, &transport.Options{Last: false}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1638
		_go_fuzz_dep_.CoverTab[81379]++
												return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1639
		// _ = "end of CoverTab[81379]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1640
		_go_fuzz_dep_.CoverTab[81380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1640
		// _ = "end of CoverTab[81380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1640
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1640
	// _ = "end of CoverTab[81353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1640
	_go_fuzz_dep_.CoverTab[81354]++
											if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1641
		_go_fuzz_dep_.CoverTab[81381]++
												if !ss.serverHeaderBinlogged {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1642
			_go_fuzz_dep_.CoverTab[81383]++
													h, _ := ss.s.Header()
													sh := &binarylog.ServerHeader{
				Header: h,
			}
			ss.serverHeaderBinlogged = true
			for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1648
				_go_fuzz_dep_.CoverTab[81384]++
														binlog.Log(ss.ctx, sh)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1649
				// _ = "end of CoverTab[81384]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1650
			// _ = "end of CoverTab[81383]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1651
			_go_fuzz_dep_.CoverTab[81385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1651
			// _ = "end of CoverTab[81385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1651
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1651
		// _ = "end of CoverTab[81381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1651
		_go_fuzz_dep_.CoverTab[81382]++
												sm := &binarylog.ServerMessage{
			Message: data,
		}
		for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1655
			_go_fuzz_dep_.CoverTab[81386]++
													binlog.Log(ss.ctx, sm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1656
			// _ = "end of CoverTab[81386]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1657
		// _ = "end of CoverTab[81382]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1658
		_go_fuzz_dep_.CoverTab[81387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1658
		// _ = "end of CoverTab[81387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1658
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1658
	// _ = "end of CoverTab[81354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1658
	_go_fuzz_dep_.CoverTab[81355]++
											if len(ss.statsHandler) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1659
		_go_fuzz_dep_.CoverTab[81388]++
												for _, sh := range ss.statsHandler {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1660
			_go_fuzz_dep_.CoverTab[81389]++
													sh.HandleRPC(ss.s.Context(), outPayload(false, m, data, payload, time.Now()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1661
			// _ = "end of CoverTab[81389]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1662
		// _ = "end of CoverTab[81388]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1663
		_go_fuzz_dep_.CoverTab[81390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1663
		// _ = "end of CoverTab[81390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1663
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1663
	// _ = "end of CoverTab[81355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1663
	_go_fuzz_dep_.CoverTab[81356]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1664
	// _ = "end of CoverTab[81356]"
}

func (ss *serverStream) RecvMsg(m interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1667
	_go_fuzz_dep_.CoverTab[81391]++
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1668
		_go_fuzz_dep_.CoverTab[81397]++
												if ss.trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1669
			_go_fuzz_dep_.CoverTab[81400]++
													ss.mu.Lock()
													if ss.trInfo.tr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1671
				_go_fuzz_dep_.CoverTab[81402]++
														if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1672
					_go_fuzz_dep_.CoverTab[81403]++
															ss.trInfo.tr.LazyLog(&payload{sent: false, msg: m}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1673
					// _ = "end of CoverTab[81403]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1674
					_go_fuzz_dep_.CoverTab[81404]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1674
					if err != io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1674
						_go_fuzz_dep_.CoverTab[81405]++
																ss.trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
																ss.trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1676
						// _ = "end of CoverTab[81405]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
						_go_fuzz_dep_.CoverTab[81406]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
						// _ = "end of CoverTab[81406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
					// _ = "end of CoverTab[81404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1677
				// _ = "end of CoverTab[81402]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1678
				_go_fuzz_dep_.CoverTab[81407]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1678
				// _ = "end of CoverTab[81407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1678
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1678
			// _ = "end of CoverTab[81400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1678
			_go_fuzz_dep_.CoverTab[81401]++
													ss.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1679
			// _ = "end of CoverTab[81401]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1680
			_go_fuzz_dep_.CoverTab[81408]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1680
			// _ = "end of CoverTab[81408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1680
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1680
		// _ = "end of CoverTab[81397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1680
		_go_fuzz_dep_.CoverTab[81398]++
												if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1681
			_go_fuzz_dep_.CoverTab[81409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1681
			return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1681
			// _ = "end of CoverTab[81409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1681
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1681
			_go_fuzz_dep_.CoverTab[81410]++
													st, _ := status.FromError(toRPCErr(err))
													ss.t.WriteStatus(ss.s, st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1683
			// _ = "end of CoverTab[81410]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
			_go_fuzz_dep_.CoverTab[81411]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
			// _ = "end of CoverTab[81411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
		// _ = "end of CoverTab[81398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1690
		_go_fuzz_dep_.CoverTab[81399]++
												if channelz.IsOn() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1691
			_go_fuzz_dep_.CoverTab[81412]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1691
			return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1691
			// _ = "end of CoverTab[81412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1691
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1691
			_go_fuzz_dep_.CoverTab[81413]++
													ss.t.IncrMsgRecv()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1692
			// _ = "end of CoverTab[81413]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1693
			_go_fuzz_dep_.CoverTab[81414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1693
			// _ = "end of CoverTab[81414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1693
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1693
		// _ = "end of CoverTab[81399]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1694
	// _ = "end of CoverTab[81391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1694
	_go_fuzz_dep_.CoverTab[81392]++
											var payInfo *payloadInfo
											if len(ss.statsHandler) != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1696
		_go_fuzz_dep_.CoverTab[81415]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1696
		return len(ss.binlogs) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1696
		// _ = "end of CoverTab[81415]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1696
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1696
		_go_fuzz_dep_.CoverTab[81416]++
												payInfo = &payloadInfo{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1697
		// _ = "end of CoverTab[81416]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1698
		_go_fuzz_dep_.CoverTab[81417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1698
		// _ = "end of CoverTab[81417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1698
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1698
	// _ = "end of CoverTab[81392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1698
	_go_fuzz_dep_.CoverTab[81393]++
											if err := recv(ss.p, ss.codec, ss.s, ss.dc, m, ss.maxReceiveMessageSize, payInfo, ss.decomp); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1699
		_go_fuzz_dep_.CoverTab[81418]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1700
			_go_fuzz_dep_.CoverTab[81421]++
													if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1701
				_go_fuzz_dep_.CoverTab[81423]++
														chc := &binarylog.ClientHalfClose{}
														for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1703
					_go_fuzz_dep_.CoverTab[81424]++
															binlog.Log(ss.ctx, chc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1704
					// _ = "end of CoverTab[81424]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1705
				// _ = "end of CoverTab[81423]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1706
				_go_fuzz_dep_.CoverTab[81425]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1706
				// _ = "end of CoverTab[81425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1706
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1706
			// _ = "end of CoverTab[81421]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1706
			_go_fuzz_dep_.CoverTab[81422]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1707
			// _ = "end of CoverTab[81422]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1708
			_go_fuzz_dep_.CoverTab[81426]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1708
			// _ = "end of CoverTab[81426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1708
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1708
		// _ = "end of CoverTab[81418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1708
		_go_fuzz_dep_.CoverTab[81419]++
												if err == io.ErrUnexpectedEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1709
			_go_fuzz_dep_.CoverTab[81427]++
													err = status.Errorf(codes.Internal, io.ErrUnexpectedEOF.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1710
			// _ = "end of CoverTab[81427]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1711
			_go_fuzz_dep_.CoverTab[81428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1711
			// _ = "end of CoverTab[81428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1711
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1711
		// _ = "end of CoverTab[81419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1711
		_go_fuzz_dep_.CoverTab[81420]++
												return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1712
		// _ = "end of CoverTab[81420]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1713
		_go_fuzz_dep_.CoverTab[81429]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1713
		// _ = "end of CoverTab[81429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1713
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1713
	// _ = "end of CoverTab[81393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1713
	_go_fuzz_dep_.CoverTab[81394]++
											if len(ss.statsHandler) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1714
		_go_fuzz_dep_.CoverTab[81430]++
												for _, sh := range ss.statsHandler {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1715
			_go_fuzz_dep_.CoverTab[81431]++
													sh.HandleRPC(ss.s.Context(), &stats.InPayload{
				RecvTime:	time.Now(),
				Payload:	m,

				Data:			payInfo.uncompressedBytes,
				Length:			len(payInfo.uncompressedBytes),
				WireLength:		payInfo.compressedLength + headerLen,
				CompressedLength:	payInfo.compressedLength,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1724
			// _ = "end of CoverTab[81431]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1725
		// _ = "end of CoverTab[81430]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1726
		_go_fuzz_dep_.CoverTab[81432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1726
		// _ = "end of CoverTab[81432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1726
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1726
	// _ = "end of CoverTab[81394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1726
	_go_fuzz_dep_.CoverTab[81395]++
											if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1727
		_go_fuzz_dep_.CoverTab[81433]++
												cm := &binarylog.ClientMessage{
			Message: payInfo.uncompressedBytes,
		}
		for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1731
			_go_fuzz_dep_.CoverTab[81434]++
													binlog.Log(ss.ctx, cm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1732
			// _ = "end of CoverTab[81434]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1733
		// _ = "end of CoverTab[81433]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1734
		_go_fuzz_dep_.CoverTab[81435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1734
		// _ = "end of CoverTab[81435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1734
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1734
	// _ = "end of CoverTab[81395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1734
	_go_fuzz_dep_.CoverTab[81396]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1735
	// _ = "end of CoverTab[81396]"
}

// MethodFromServerStream returns the method string for the input stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1738
// The returned string is in the format of "/service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1740
func MethodFromServerStream(stream ServerStream) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1740
	_go_fuzz_dep_.CoverTab[81436]++
											return Method(stream.Context())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1741
	// _ = "end of CoverTab[81436]"
}

// prepareMsg returns the hdr, payload and data
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1744
// using the compressors passed or using the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1744
// passed preparedmsg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1747
func prepareMsg(m interface{}, codec baseCodec, cp Compressor, comp encoding.Compressor) (hdr, payload, data []byte, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1747
	_go_fuzz_dep_.CoverTab[81437]++
											if preparedMsg, ok := m.(*PreparedMsg); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1748
		_go_fuzz_dep_.CoverTab[81441]++
												return preparedMsg.hdr, preparedMsg.payload, preparedMsg.encodedData, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1749
		// _ = "end of CoverTab[81441]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1750
		_go_fuzz_dep_.CoverTab[81442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1750
		// _ = "end of CoverTab[81442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1750
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1750
	// _ = "end of CoverTab[81437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1750
	_go_fuzz_dep_.CoverTab[81438]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1753
	data, err = encode(codec, m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1754
		_go_fuzz_dep_.CoverTab[81443]++
												return nil, nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1755
		// _ = "end of CoverTab[81443]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1756
		_go_fuzz_dep_.CoverTab[81444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1756
		// _ = "end of CoverTab[81444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1756
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1756
	// _ = "end of CoverTab[81438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1756
	_go_fuzz_dep_.CoverTab[81439]++
											compData, err := compress(data, cp, comp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1758
		_go_fuzz_dep_.CoverTab[81445]++
												return nil, nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1759
		// _ = "end of CoverTab[81445]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1760
		_go_fuzz_dep_.CoverTab[81446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1760
		// _ = "end of CoverTab[81446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1760
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1760
	// _ = "end of CoverTab[81439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1760
	_go_fuzz_dep_.CoverTab[81440]++
											hdr, payload = msgHeader(data, compData)
											return hdr, payload, data, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1762
	// _ = "end of CoverTab[81440]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1763
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stream.go:1763
var _ = _go_fuzz_dep_.CoverTab
