//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:19
)

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/channelz"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/internal"
	internalbackoff "google.golang.org/grpc/internal/backoff"
	"google.golang.org/grpc/internal/binarylog"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/stats"
)

func init() {
	internal.AddGlobalDialOptions = func(opt ...DialOption) {
		globalDialOptions = append(globalDialOptions, opt...)
	}
	internal.ClearGlobalDialOptions = func() {
		globalDialOptions = nil
	}
	internal.WithBinaryLogger = withBinaryLogger
	internal.JoinDialOptions = newJoinDialOption
	internal.DisableGlobalDialOptions = newDisableGlobalDialOptions
}

// dialOptions configure a Dial call. dialOptions are set by the DialOption
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:51
// values passed to Dial.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:53
type dialOptions struct {
	unaryInt	UnaryClientInterceptor
	streamInt	StreamClientInterceptor

	chainUnaryInts	[]UnaryClientInterceptor
	chainStreamInts	[]StreamClientInterceptor

	cp				Compressor
	dc				Decompressor
	bs				internalbackoff.Strategy
	block				bool
	returnLastError			bool
	timeout				time.Duration
	scChan				<-chan ServiceConfig
	authority			string
	binaryLogger			binarylog.Logger
	copts				transport.ConnectOptions
	callOptions			[]CallOption
	channelzParentID		*channelz.Identifier
	disableServiceConfig		bool
	disableRetry			bool
	disableHealthCheck		bool
	healthCheckFunc			internal.HealthChecker
	minConnectTimeout		func() time.Duration
	defaultServiceConfig		*ServiceConfig	// defaultServiceConfig is parsed from defaultServiceConfigRawJSON.
	defaultServiceConfigRawJSON	*string
	resolvers			[]resolver.Builder
}

// DialOption configures how we set up the connection.
type DialOption interface {
	apply(*dialOptions)
}

var globalDialOptions []DialOption

// EmptyDialOption does not alter the dial configuration. It can be embedded in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
// another structure to build custom dial options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:89
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:96
type EmptyDialOption struct{}

func (EmptyDialOption) apply(*dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:98
	_go_fuzz_dep_.CoverTab[79375]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:98
	// _ = "end of CoverTab[79375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:98
}

type disableGlobalDialOptions struct{}

func (disableGlobalDialOptions) apply(*dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:102
	_go_fuzz_dep_.CoverTab[79376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:102
	// _ = "end of CoverTab[79376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:102
}

// newDisableGlobalDialOptions returns a DialOption that prevents the ClientConn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:104
// from applying the global DialOptions (set via AddGlobalDialOptions).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:106
func newDisableGlobalDialOptions() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:106
	_go_fuzz_dep_.CoverTab[79377]++
											return &disableGlobalDialOptions{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:107
	// _ = "end of CoverTab[79377]"
}

// funcDialOption wraps a function that modifies dialOptions into an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:110
// implementation of the DialOption interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:112
type funcDialOption struct {
	f func(*dialOptions)
}

func (fdo *funcDialOption) apply(do *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:116
	_go_fuzz_dep_.CoverTab[79378]++
											fdo.f(do)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:117
	// _ = "end of CoverTab[79378]"
}

func newFuncDialOption(f func(*dialOptions)) *funcDialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:120
	_go_fuzz_dep_.CoverTab[79379]++
											return &funcDialOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:123
	// _ = "end of CoverTab[79379]"
}

type joinDialOption struct {
	opts []DialOption
}

func (jdo *joinDialOption) apply(do *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:130
	_go_fuzz_dep_.CoverTab[79380]++
											for _, opt := range jdo.opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:131
		_go_fuzz_dep_.CoverTab[79381]++
												opt.apply(do)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:132
		// _ = "end of CoverTab[79381]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:133
	// _ = "end of CoverTab[79380]"
}

func newJoinDialOption(opts ...DialOption) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:136
	_go_fuzz_dep_.CoverTab[79382]++
											return &joinDialOption{opts: opts}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:137
	// _ = "end of CoverTab[79382]"
}

// WithWriteBufferSize determines how much data can be batched before doing a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// write on the wire. The corresponding memory allocation for this buffer will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// be twice the size to keep syscalls low. The default value for this buffer is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// 32KB.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// Zero or negative values will disable the write buffer such that each write
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// will be on underlying connection. Note: A Send call may not directly
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:140
// translate to a write.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:148
func WithWriteBufferSize(s int) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:148
	_go_fuzz_dep_.CoverTab[79383]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:149
		_go_fuzz_dep_.CoverTab[79384]++
												o.copts.WriteBufferSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:150
		// _ = "end of CoverTab[79384]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:151
	// _ = "end of CoverTab[79383]"
}

// WithReadBufferSize lets you set the size of read buffer, this determines how
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:154
// much data can be read at most for each read syscall.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:154
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:154
// The default value for this buffer is 32KB. Zero or negative values will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:154
// disable read buffer for a connection so data framer can access the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:154
// underlying conn directly.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:160
func WithReadBufferSize(s int) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:160
	_go_fuzz_dep_.CoverTab[79385]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:161
		_go_fuzz_dep_.CoverTab[79386]++
												o.copts.ReadBufferSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:162
		// _ = "end of CoverTab[79386]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:163
	// _ = "end of CoverTab[79385]"
}

// WithInitialWindowSize returns a DialOption which sets the value for initial
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:166
// window size on a stream. The lower bound for window size is 64K and any value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:166
// smaller than that will be ignored.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:169
func WithInitialWindowSize(s int32) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:169
	_go_fuzz_dep_.CoverTab[79387]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:170
		_go_fuzz_dep_.CoverTab[79388]++
												o.copts.InitialWindowSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:171
		// _ = "end of CoverTab[79388]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:172
	// _ = "end of CoverTab[79387]"
}

// WithInitialConnWindowSize returns a DialOption which sets the value for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:175
// initial window size on a connection. The lower bound for window size is 64K
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:175
// and any value smaller than that will be ignored.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:178
func WithInitialConnWindowSize(s int32) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:178
	_go_fuzz_dep_.CoverTab[79389]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:179
		_go_fuzz_dep_.CoverTab[79390]++
												o.copts.InitialConnWindowSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:180
		// _ = "end of CoverTab[79390]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:181
	// _ = "end of CoverTab[79389]"
}

// WithMaxMsgSize returns a DialOption which sets the maximum message size the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:184
// client can receive.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:184
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:184
// Deprecated: use WithDefaultCallOptions(MaxCallRecvMsgSize(s)) instead.  Will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:184
// be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:189
func WithMaxMsgSize(s int) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:189
	_go_fuzz_dep_.CoverTab[79391]++
											return WithDefaultCallOptions(MaxCallRecvMsgSize(s))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:190
	// _ = "end of CoverTab[79391]"
}

// WithDefaultCallOptions returns a DialOption which sets the default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:193
// CallOptions for calls over the connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:195
func WithDefaultCallOptions(cos ...CallOption) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:195
	_go_fuzz_dep_.CoverTab[79392]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:196
		_go_fuzz_dep_.CoverTab[79393]++
												o.callOptions = append(o.callOptions, cos...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:197
		// _ = "end of CoverTab[79393]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:198
	// _ = "end of CoverTab[79392]"
}

// WithCodec returns a DialOption which sets a codec for message marshaling and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:201
// unmarshaling.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:201
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:201
// Deprecated: use WithDefaultCallOptions(ForceCodec(_)) instead.  Will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:201
// supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:206
func WithCodec(c Codec) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:206
	_go_fuzz_dep_.CoverTab[79394]++
											return WithDefaultCallOptions(CallCustomCodec(c))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:207
	// _ = "end of CoverTab[79394]"
}

// WithCompressor returns a DialOption which sets a Compressor to use for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:210
// message compression. It has lower priority than the compressor set by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:210
// UseCompressor CallOption.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:210
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:210
// Deprecated: use UseCompressor instead.  Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:215
func WithCompressor(cp Compressor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:215
	_go_fuzz_dep_.CoverTab[79395]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:216
		_go_fuzz_dep_.CoverTab[79396]++
												o.cp = cp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:217
		// _ = "end of CoverTab[79396]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:218
	// _ = "end of CoverTab[79395]"
}

// WithDecompressor returns a DialOption which sets a Decompressor to use for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// incoming message decompression.  If incoming response messages are encoded
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// using the decompressor's Type(), it will be used.  Otherwise, the message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// encoding will be used to look up the compressor registered via
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// encoding.RegisterCompressor, which will then be used to decompress the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// message.  If no compressor is registered for the encoding, an Unimplemented
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// status error will be returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// Deprecated: use encoding.RegisterCompressor instead.  Will be supported
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:221
// throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:231
func WithDecompressor(dc Decompressor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:231
	_go_fuzz_dep_.CoverTab[79397]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:232
		_go_fuzz_dep_.CoverTab[79398]++
												o.dc = dc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:233
		// _ = "end of CoverTab[79398]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:234
	// _ = "end of CoverTab[79397]"
}

// WithServiceConfig returns a DialOption which has a channel to read the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
// service configuration.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
// Deprecated: service config should be received through name resolver or via
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
// WithDefaultServiceConfig, as specified at
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
// https://github.com/grpc/grpc/blob/master/doc/service_config.md.  Will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:237
// removed in a future 1.x release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:244
func WithServiceConfig(c <-chan ServiceConfig) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:244
	_go_fuzz_dep_.CoverTab[79399]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:245
		_go_fuzz_dep_.CoverTab[79400]++
												o.scChan = c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:246
		// _ = "end of CoverTab[79400]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:247
	// _ = "end of CoverTab[79399]"
}

// WithConnectParams configures the ClientConn to use the provided ConnectParams
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// for creating and maintaining connections to servers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// The backoff configuration specified as part of the ConnectParams overrides
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// all defaults specified in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md. Consider
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// using the backoff.DefaultConfig as a base, in cases where you want to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:250
// override only a subset of the backoff configuration.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:258
func WithConnectParams(p ConnectParams) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:258
	_go_fuzz_dep_.CoverTab[79401]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:259
		_go_fuzz_dep_.CoverTab[79402]++
												o.bs = internalbackoff.Exponential{Config: p.Backoff}
												o.minConnectTimeout = func() time.Duration {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:261
			_go_fuzz_dep_.CoverTab[79403]++
													return p.MinConnectTimeout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:262
			// _ = "end of CoverTab[79403]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:263
		// _ = "end of CoverTab[79402]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:264
	// _ = "end of CoverTab[79401]"
}

// WithBackoffMaxDelay configures the dialer to use the provided maximum delay
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:267
// when backing off after failed connection attempts.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:267
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:267
// Deprecated: use WithConnectParams instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:271
func WithBackoffMaxDelay(md time.Duration) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:271
	_go_fuzz_dep_.CoverTab[79404]++
											return WithBackoffConfig(BackoffConfig{MaxDelay: md})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:272
	// _ = "end of CoverTab[79404]"
}

// WithBackoffConfig configures the dialer to use the provided backoff
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:275
// parameters after connection failures.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:275
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:275
// Deprecated: use WithConnectParams instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:279
func WithBackoffConfig(b BackoffConfig) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:279
	_go_fuzz_dep_.CoverTab[79405]++
											bc := backoff.DefaultConfig
											bc.MaxDelay = b.MaxDelay
											return withBackoff(internalbackoff.Exponential{Config: bc})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:282
	// _ = "end of CoverTab[79405]"
}

// withBackoff sets the backoff strategy used for connectRetryNum after a failed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:285
// connection attempt.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:285
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:285
// This can be exported if arbitrary backoff strategies are allowed by gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:289
func withBackoff(bs internalbackoff.Strategy) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:289
	_go_fuzz_dep_.CoverTab[79406]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:290
		_go_fuzz_dep_.CoverTab[79407]++
												o.bs = bs
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:291
		// _ = "end of CoverTab[79407]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:292
	// _ = "end of CoverTab[79406]"
}

// WithBlock returns a DialOption which makes callers of Dial block until the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:295
// underlying connection is up. Without this, Dial returns immediately and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:295
// connecting the server happens in background.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:298
func WithBlock() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:298
	_go_fuzz_dep_.CoverTab[79408]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:299
		_go_fuzz_dep_.CoverTab[79409]++
												o.block = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:300
		// _ = "end of CoverTab[79409]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:301
	// _ = "end of CoverTab[79408]"
}

// WithReturnConnectionError returns a DialOption which makes the client connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// return a string containing both the last connection error that occurred and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// the context.DeadlineExceeded error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// Implies WithBlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:304
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:313
func WithReturnConnectionError() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:313
	_go_fuzz_dep_.CoverTab[79410]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:314
		_go_fuzz_dep_.CoverTab[79411]++
												o.block = true
												o.returnLastError = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:316
		// _ = "end of CoverTab[79411]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:317
	// _ = "end of CoverTab[79410]"
}

// WithInsecure returns a DialOption which disables transport security for this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// ClientConn. Under the hood, it uses insecure.NewCredentials().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// Note that using this DialOption with per-RPC credentials (through
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// WithCredentialsBundle or WithPerRPCCredentials) which require transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// security is incompatible and will cause grpc.Dial() to fail.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// Deprecated: use WithTransportCredentials and insecure.NewCredentials()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:320
// instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:329
func WithInsecure() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:329
	_go_fuzz_dep_.CoverTab[79412]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:330
		_go_fuzz_dep_.CoverTab[79413]++
												o.copts.TransportCredentials = insecure.NewCredentials()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:331
		// _ = "end of CoverTab[79413]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:332
	// _ = "end of CoverTab[79412]"
}

// WithNoProxy returns a DialOption which disables the use of proxies for this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
// ClientConn. This is ignored if WithDialer or WithContextDialer are used.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:335
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:342
func WithNoProxy() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:342
	_go_fuzz_dep_.CoverTab[79414]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:343
		_go_fuzz_dep_.CoverTab[79415]++
												o.copts.UseProxy = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:344
		// _ = "end of CoverTab[79415]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:345
	// _ = "end of CoverTab[79414]"
}

// WithTransportCredentials returns a DialOption which configures a connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:348
// level security credentials (e.g., TLS/SSL). This should not be used together
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:348
// with WithCredentialsBundle.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:351
func WithTransportCredentials(creds credentials.TransportCredentials) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:351
	_go_fuzz_dep_.CoverTab[79416]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:352
		_go_fuzz_dep_.CoverTab[79417]++
												o.copts.TransportCredentials = creds
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:353
		// _ = "end of CoverTab[79417]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:354
	// _ = "end of CoverTab[79416]"
}

// WithPerRPCCredentials returns a DialOption which sets credentials and places
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:357
// auth state on each outbound RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:359
func WithPerRPCCredentials(creds credentials.PerRPCCredentials) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:359
	_go_fuzz_dep_.CoverTab[79418]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:360
		_go_fuzz_dep_.CoverTab[79419]++
												o.copts.PerRPCCredentials = append(o.copts.PerRPCCredentials, creds)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:361
		// _ = "end of CoverTab[79419]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:362
	// _ = "end of CoverTab[79418]"
}

// WithCredentialsBundle returns a DialOption to set a credentials bundle for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
// the ClientConn.WithCreds. This should not be used together with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
// WithTransportCredentials.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:365
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:373
func WithCredentialsBundle(b credentials.Bundle) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:373
	_go_fuzz_dep_.CoverTab[79420]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:374
		_go_fuzz_dep_.CoverTab[79421]++
												o.copts.CredsBundle = b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:375
		// _ = "end of CoverTab[79421]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:376
	// _ = "end of CoverTab[79420]"
}

// WithTimeout returns a DialOption that configures a timeout for dialing a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:379
// ClientConn initially. This is valid if and only if WithBlock() is present.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:379
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:379
// Deprecated: use DialContext instead of Dial and context.WithTimeout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:379
// instead.  Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:384
func WithTimeout(d time.Duration) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:384
	_go_fuzz_dep_.CoverTab[79422]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:385
		_go_fuzz_dep_.CoverTab[79423]++
												o.timeout = d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:386
		// _ = "end of CoverTab[79423]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:387
	// _ = "end of CoverTab[79422]"
}

// WithContextDialer returns a DialOption that sets a dialer to create
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:390
// connections. If FailOnNonTempDialError() is set to true, and an error is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:390
// returned by f, gRPC checks the error's Temporary() method to decide if it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:390
// should try to reconnect to the network address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:394
func WithContextDialer(f func(context.Context, string) (net.Conn, error)) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:394
	_go_fuzz_dep_.CoverTab[79424]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:395
		_go_fuzz_dep_.CoverTab[79425]++
												o.copts.Dialer = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:396
		// _ = "end of CoverTab[79425]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:397
	// _ = "end of CoverTab[79424]"
}

func init() {
	internal.WithHealthCheckFunc = withHealthCheckFunc
}

// WithDialer returns a DialOption that specifies a function to use for dialing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
// network addresses. If FailOnNonTempDialError() is set to true, and an error
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
// is returned by f, gRPC checks the error's Temporary() method to decide if it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
// should try to reconnect to the network address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
// Deprecated: use WithContextDialer instead.  Will be supported throughout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:404
// 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:411
func WithDialer(f func(string, time.Duration) (net.Conn, error)) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:411
	_go_fuzz_dep_.CoverTab[79426]++
											return WithContextDialer(
		func(ctx context.Context, addr string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:413
			_go_fuzz_dep_.CoverTab[79427]++
													if deadline, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:414
				_go_fuzz_dep_.CoverTab[79429]++
														return f(addr, time.Until(deadline))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:415
				// _ = "end of CoverTab[79429]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:416
				_go_fuzz_dep_.CoverTab[79430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:416
				// _ = "end of CoverTab[79430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:416
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:416
			// _ = "end of CoverTab[79427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:416
			_go_fuzz_dep_.CoverTab[79428]++
													return f(addr, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:417
			// _ = "end of CoverTab[79428]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:418
	// _ = "end of CoverTab[79426]"
}

// WithStatsHandler returns a DialOption that specifies the stats handler for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:421
// all the RPCs and underlying network connections in this ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:423
func WithStatsHandler(h stats.Handler) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:423
	_go_fuzz_dep_.CoverTab[79431]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:424
		_go_fuzz_dep_.CoverTab[79432]++
												if h == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:425
			_go_fuzz_dep_.CoverTab[79434]++
													logger.Error("ignoring nil parameter in grpc.WithStatsHandler ClientOption")

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:429
			return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:429
			// _ = "end of CoverTab[79434]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:430
			_go_fuzz_dep_.CoverTab[79435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:430
			// _ = "end of CoverTab[79435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:430
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:430
		// _ = "end of CoverTab[79432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:430
		_go_fuzz_dep_.CoverTab[79433]++
												o.copts.StatsHandlers = append(o.copts.StatsHandlers, h)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:431
		// _ = "end of CoverTab[79433]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:432
	// _ = "end of CoverTab[79431]"
}

// withBinaryLogger returns a DialOption that specifies the binary logger for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:435
// this ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:437
func withBinaryLogger(bl binarylog.Logger) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:437
	_go_fuzz_dep_.CoverTab[79436]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:438
		_go_fuzz_dep_.CoverTab[79437]++
												o.binaryLogger = bl
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:439
		// _ = "end of CoverTab[79437]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:440
	// _ = "end of CoverTab[79436]"
}

// FailOnNonTempDialError returns a DialOption that specifies if gRPC fails on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// non-temporary dial errors. If f is true, and dialer returns a non-temporary
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// error, gRPC will fail the connection to the network address and won't try to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// reconnect. The default value of FailOnNonTempDialError is false.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// FailOnNonTempDialError only affects the initial dial, and does not do
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// anything useful unless you are also using WithBlock().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:443
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:455
func FailOnNonTempDialError(f bool) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:455
	_go_fuzz_dep_.CoverTab[79438]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:456
		_go_fuzz_dep_.CoverTab[79439]++
												o.copts.FailOnNonTempDialError = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:457
		// _ = "end of CoverTab[79439]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:458
	// _ = "end of CoverTab[79438]"
}

// WithUserAgent returns a DialOption that specifies a user agent string for all
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:461
// the RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:463
func WithUserAgent(s string) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:463
	_go_fuzz_dep_.CoverTab[79440]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:464
		_go_fuzz_dep_.CoverTab[79441]++
												o.copts.UserAgent = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:465
		// _ = "end of CoverTab[79441]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:466
	// _ = "end of CoverTab[79440]"
}

// WithKeepaliveParams returns a DialOption that specifies keepalive parameters
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:469
// for the client transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:471
func WithKeepaliveParams(kp keepalive.ClientParameters) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:471
	_go_fuzz_dep_.CoverTab[79442]++
											if kp.Time < internal.KeepaliveMinPingTime {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:472
		_go_fuzz_dep_.CoverTab[79444]++
												logger.Warningf("Adjusting keepalive ping interval to minimum period of %v", internal.KeepaliveMinPingTime)
												kp.Time = internal.KeepaliveMinPingTime
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:474
		// _ = "end of CoverTab[79444]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:475
		_go_fuzz_dep_.CoverTab[79445]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:475
		// _ = "end of CoverTab[79445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:475
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:475
	// _ = "end of CoverTab[79442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:475
	_go_fuzz_dep_.CoverTab[79443]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:476
		_go_fuzz_dep_.CoverTab[79446]++
												o.copts.KeepaliveParams = kp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:477
		// _ = "end of CoverTab[79446]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:478
	// _ = "end of CoverTab[79443]"
}

// WithUnaryInterceptor returns a DialOption that specifies the interceptor for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:481
// unary RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:483
func WithUnaryInterceptor(f UnaryClientInterceptor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:483
	_go_fuzz_dep_.CoverTab[79447]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:484
		_go_fuzz_dep_.CoverTab[79448]++
												o.unaryInt = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:485
		// _ = "end of CoverTab[79448]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:486
	// _ = "end of CoverTab[79447]"
}

// WithChainUnaryInterceptor returns a DialOption that specifies the chained
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:489
// interceptor for unary RPCs. The first interceptor will be the outer most,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:489
// while the last interceptor will be the inner most wrapper around the real call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:489
// All interceptors added by this method will be chained, and the interceptor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:489
// defined by WithUnaryInterceptor will always be prepended to the chain.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:494
func WithChainUnaryInterceptor(interceptors ...UnaryClientInterceptor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:494
	_go_fuzz_dep_.CoverTab[79449]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:495
		_go_fuzz_dep_.CoverTab[79450]++
												o.chainUnaryInts = append(o.chainUnaryInts, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:496
		// _ = "end of CoverTab[79450]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:497
	// _ = "end of CoverTab[79449]"
}

// WithStreamInterceptor returns a DialOption that specifies the interceptor for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:500
// streaming RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:502
func WithStreamInterceptor(f StreamClientInterceptor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:502
	_go_fuzz_dep_.CoverTab[79451]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:503
		_go_fuzz_dep_.CoverTab[79452]++
												o.streamInt = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:504
		// _ = "end of CoverTab[79452]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:505
	// _ = "end of CoverTab[79451]"
}

// WithChainStreamInterceptor returns a DialOption that specifies the chained
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:508
// interceptor for streaming RPCs. The first interceptor will be the outer most,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:508
// while the last interceptor will be the inner most wrapper around the real call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:508
// All interceptors added by this method will be chained, and the interceptor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:508
// defined by WithStreamInterceptor will always be prepended to the chain.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:513
func WithChainStreamInterceptor(interceptors ...StreamClientInterceptor) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:513
	_go_fuzz_dep_.CoverTab[79453]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:514
		_go_fuzz_dep_.CoverTab[79454]++
												o.chainStreamInts = append(o.chainStreamInts, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:515
		// _ = "end of CoverTab[79454]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:516
	// _ = "end of CoverTab[79453]"
}

// WithAuthority returns a DialOption that specifies the value to be used as the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:519
// :authority pseudo-header and as the server name in authentication handshake.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:521
func WithAuthority(a string) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:521
	_go_fuzz_dep_.CoverTab[79455]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:522
		_go_fuzz_dep_.CoverTab[79456]++
												o.authority = a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:523
		// _ = "end of CoverTab[79456]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:524
	// _ = "end of CoverTab[79455]"
}

// WithChannelzParentID returns a DialOption that specifies the channelz ID of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
// current ClientConn's parent. This function is used in nested channel creation
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
// (e.g. grpclb dial).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:527
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:535
func WithChannelzParentID(id *channelz.Identifier) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:535
	_go_fuzz_dep_.CoverTab[79457]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:536
		_go_fuzz_dep_.CoverTab[79458]++
												o.channelzParentID = id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:537
		// _ = "end of CoverTab[79458]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:538
	// _ = "end of CoverTab[79457]"
}

// WithDisableServiceConfig returns a DialOption that causes gRPC to ignore any
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:541
// service config provided by the resolver and provides a hint to the resolver
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:541
// to not fetch service configs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:541
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:541
// Note that this dial option only disables service config from resolver. If
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:541
// default service config is provided, gRPC will use the default service config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:547
func WithDisableServiceConfig() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:547
	_go_fuzz_dep_.CoverTab[79459]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:548
		_go_fuzz_dep_.CoverTab[79460]++
												o.disableServiceConfig = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:549
		// _ = "end of CoverTab[79460]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:550
	// _ = "end of CoverTab[79459]"
}

// WithDefaultServiceConfig returns a DialOption that configures the default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// service config, which will be used in cases where:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// 1. WithDisableServiceConfig is also used, or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// 2. The name resolver does not provide a service config or provides an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// invalid service config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// The parameter s is the JSON representation of the default service config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// For more information about service configs, see:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// https://github.com/grpc/grpc/blob/master/doc/service_config.md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// For a simple example of usage, see:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:553
// examples/features/load_balancing/client/main.go
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:566
func WithDefaultServiceConfig(s string) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:566
	_go_fuzz_dep_.CoverTab[79461]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:567
		_go_fuzz_dep_.CoverTab[79462]++
												o.defaultServiceConfigRawJSON = &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:568
		// _ = "end of CoverTab[79462]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:569
	// _ = "end of CoverTab[79461]"
}

// WithDisableRetry returns a DialOption that disables retries, even if the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:572
// service config enables them.  This does not impact transparent retries, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:572
// will happen automatically if no data is written to the wire or if the RPC is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:572
// unprocessed by the remote server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:576
func WithDisableRetry() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:576
	_go_fuzz_dep_.CoverTab[79463]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:577
		_go_fuzz_dep_.CoverTab[79464]++
												o.disableRetry = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:578
		// _ = "end of CoverTab[79464]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:579
	// _ = "end of CoverTab[79463]"
}

// WithMaxHeaderListSize returns a DialOption that specifies the maximum
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:582
// (uncompressed) size of header list that the client is prepared to accept.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:584
func WithMaxHeaderListSize(s uint32) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:584
	_go_fuzz_dep_.CoverTab[79465]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:585
		_go_fuzz_dep_.CoverTab[79466]++
												o.copts.MaxHeaderListSize = &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:586
		// _ = "end of CoverTab[79466]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:587
	// _ = "end of CoverTab[79465]"
}

// WithDisableHealthCheck disables the LB channel health checking for all
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
// SubConns of this ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:590
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:597
func WithDisableHealthCheck() DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:597
	_go_fuzz_dep_.CoverTab[79467]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:598
		_go_fuzz_dep_.CoverTab[79468]++
												o.disableHealthCheck = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:599
		// _ = "end of CoverTab[79468]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:600
	// _ = "end of CoverTab[79467]"
}

// withHealthCheckFunc replaces the default health check function with the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:603
// provided one. It makes tests easier to change the health check function.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:603
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:603
// For testing purpose only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:607
func withHealthCheckFunc(f internal.HealthChecker) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:607
	_go_fuzz_dep_.CoverTab[79469]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:608
		_go_fuzz_dep_.CoverTab[79470]++
												o.healthCheckFunc = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:609
		// _ = "end of CoverTab[79470]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:610
	// _ = "end of CoverTab[79469]"
}

func defaultDialOptions() dialOptions {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:613
	_go_fuzz_dep_.CoverTab[79471]++
											return dialOptions{
		healthCheckFunc:	internal.HealthCheckFunc,
		copts: transport.ConnectOptions{
			WriteBufferSize:	defaultWriteBufSize,
			ReadBufferSize:		defaultReadBufSize,
			UseProxy:		true,
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:621
	// _ = "end of CoverTab[79471]"
}

// withGetMinConnectDeadline specifies the function that clientconn uses to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:624
// get minConnectDeadline. This can be used to make connection attempts happen
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:624
// faster/slower.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:624
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:624
// For testing purpose only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:629
func withMinConnectDeadline(f func() time.Duration) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:629
	_go_fuzz_dep_.CoverTab[79472]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:630
		_go_fuzz_dep_.CoverTab[79473]++
												o.minConnectTimeout = f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:631
		// _ = "end of CoverTab[79473]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:632
	// _ = "end of CoverTab[79472]"
}

// WithResolvers allows a list of resolver implementations to be registered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// locally with the ClientConn without needing to be globally registered via
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// resolver.Register.  They will be matched against the scheme used for the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// current Dial only, and will take precedence over the global registry.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:635
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:644
func WithResolvers(rs ...resolver.Builder) DialOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:644
	_go_fuzz_dep_.CoverTab[79474]++
											return newFuncDialOption(func(o *dialOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:645
		_go_fuzz_dep_.CoverTab[79475]++
												o.resolvers = append(o.resolvers, rs...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:646
		// _ = "end of CoverTab[79475]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:647
	// _ = "end of CoverTab[79474]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:648
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/dialoptions.go:648
var _ = _go_fuzz_dep_.CoverTab
