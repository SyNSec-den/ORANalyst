//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:19
// Package stats is for collecting and reporting various network and RPC stats.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:19
// This package is for monitoring purpose only. All fields are read-only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:19
// All APIs are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
package stats

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:22
)

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc/metadata"
)

// RPCStats contains stats information about RPCs.
type RPCStats interface {
	isRPCStats()
	// IsClient returns true if this RPCStats is from client side.
	IsClient() bool
}

// Begin contains stats when an RPC attempt begins.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:39
// FailFast is only valid if this Begin is from client side.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:41
type Begin struct {
	// Client is true if this Begin is from client side.
	Client	bool
	// BeginTime is the time when the RPC attempt begins.
	BeginTime	time.Time
	// FailFast indicates if this RPC is failfast.
	FailFast	bool
	// IsClientStream indicates whether the RPC is a client streaming RPC.
	IsClientStream	bool
	// IsServerStream indicates whether the RPC is a server streaming RPC.
	IsServerStream	bool
	// IsTransparentRetryAttempt indicates whether this attempt was initiated
	// due to transparently retrying a previous attempt.
	IsTransparentRetryAttempt	bool
}

// IsClient indicates if the stats information is from client side.
func (s *Begin) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:58
	_go_fuzz_dep_.CoverTab[76020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:58
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:58
	// _ = "end of CoverTab[76020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:58
}

func (s *Begin) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76021]++; // _ = "end of CoverTab[76021]" }

// InPayload contains the information for an incoming payload.
type InPayload struct {
	// Client is true if this InPayload is from client side.
	Client	bool
	// Payload is the payload with original type.
	Payload	interface{}
	// Data is the serialized message payload.
	Data	[]byte

	// Length is the size of the uncompressed payload data. Does not include any
	// framing (gRPC or HTTP/2).
	Length	int
	// CompressedLength is the size of the compressed payload data. Does not
	// include any framing (gRPC or HTTP/2). Same as Length if compression not
	// enabled.
	CompressedLength	int
	// WireLength is the size of the compressed payload data plus gRPC framing.
	// Does not include HTTP/2 framing.
	WireLength	int

	// RecvTime is the time when the payload is received.
	RecvTime	time.Time
}

// IsClient indicates if the stats information is from client side.
func (s *InPayload) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:87
	_go_fuzz_dep_.CoverTab[76022]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:87
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:87
	// _ = "end of CoverTab[76022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:87
}

func (s *InPayload) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76023]++; // _ = "end of CoverTab[76023]" }

// InHeader contains stats when a header is received.
type InHeader struct {
	// Client is true if this InHeader is from client side.
	Client	bool
	// WireLength is the wire length of header.
	WireLength	int
	// Compression is the compression algorithm used for the RPC.
	Compression	string
	// Header contains the header metadata received.
	Header	metadata.MD

	// The following fields are valid only if Client is false.
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod	string
	// RemoteAddr is the remote address of the corresponding connection.
	RemoteAddr	net.Addr
	// LocalAddr is the local address of the corresponding connection.
	LocalAddr	net.Addr
}

// IsClient indicates if the stats information is from client side.
func (s *InHeader) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:112
	_go_fuzz_dep_.CoverTab[76024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:112
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:112
	// _ = "end of CoverTab[76024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:112
}

func (s *InHeader) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76025]++; // _ = "end of CoverTab[76025]" }

// InTrailer contains stats when a trailer is received.
type InTrailer struct {
	// Client is true if this InTrailer is from client side.
	Client	bool
	// WireLength is the wire length of trailer.
	WireLength	int
	// Trailer contains the trailer metadata received from the server. This
	// field is only valid if this InTrailer is from the client side.
	Trailer	metadata.MD
}

// IsClient indicates if the stats information is from client side.
func (s *InTrailer) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:128
	_go_fuzz_dep_.CoverTab[76026]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:128
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:128
	// _ = "end of CoverTab[76026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:128
}

func (s *InTrailer) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76027]++; // _ = "end of CoverTab[76027]" }

// OutPayload contains the information for an outgoing payload.
type OutPayload struct {
	// Client is true if this OutPayload is from client side.
	Client	bool
	// Payload is the payload with original type.
	Payload	interface{}
	// Data is the serialized message payload.
	Data	[]byte
	// Length is the size of the uncompressed payload data. Does not include any
	// framing (gRPC or HTTP/2).
	Length	int
	// CompressedLength is the size of the compressed payload data. Does not
	// include any framing (gRPC or HTTP/2). Same as Length if compression not
	// enabled.
	CompressedLength	int
	// WireLength is the size of the compressed payload data plus gRPC framing.
	// Does not include HTTP/2 framing.
	WireLength	int
	// SentTime is the time when the payload is sent.
	SentTime	time.Time
}

// IsClient indicates if this stats information is from client side.
func (s *OutPayload) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:155
	_go_fuzz_dep_.CoverTab[76028]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:155
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:155
	// _ = "end of CoverTab[76028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:155
}

func (s *OutPayload) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76029]++; // _ = "end of CoverTab[76029]" }

// OutHeader contains stats when a header is sent.
type OutHeader struct {
	// Client is true if this OutHeader is from client side.
	Client	bool
	// Compression is the compression algorithm used for the RPC.
	Compression	string
	// Header contains the header metadata sent.
	Header	metadata.MD

	// The following fields are valid only if Client is true.
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod	string
	// RemoteAddr is the remote address of the corresponding connection.
	RemoteAddr	net.Addr
	// LocalAddr is the local address of the corresponding connection.
	LocalAddr	net.Addr
}

// IsClient indicates if this stats information is from client side.
func (s *OutHeader) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:178
	_go_fuzz_dep_.CoverTab[76030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:178
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:178
	// _ = "end of CoverTab[76030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:178
}

func (s *OutHeader) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76031]++; // _ = "end of CoverTab[76031]" }

// OutTrailer contains stats when a trailer is sent.
type OutTrailer struct {
	// Client is true if this OutTrailer is from client side.
	Client	bool
	// WireLength is the wire length of trailer.
	//
	// Deprecated: This field is never set. The length is not known when this message is
	// emitted because the trailer fields are compressed with hpack after that.
	WireLength	int
	// Trailer contains the trailer metadata sent to the client. This
	// field is only valid if this OutTrailer is from the server side.
	Trailer	metadata.MD
}

// IsClient indicates if this stats information is from client side.
func (s *OutTrailer) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:197
	_go_fuzz_dep_.CoverTab[76032]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:197
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:197
	// _ = "end of CoverTab[76032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:197
}

func (s *OutTrailer) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76033]++; // _ = "end of CoverTab[76033]" }

// End contains stats when an RPC ends.
type End struct {
	// Client is true if this End is from client side.
	Client	bool
	// BeginTime is the time when the RPC began.
	BeginTime	time.Time
	// EndTime is the time when the RPC ends.
	EndTime	time.Time
	// Trailer contains the trailer metadata received from the server. This
	// field is only valid if this End is from the client side.
	// Deprecated: use Trailer in InTrailer instead.
	Trailer	metadata.MD
	// Error is the error the RPC ended with. It is an error generated from
	// status.Status and can be converted back to status.Status using
	// status.FromError if non-nil.
	Error	error
}

// IsClient indicates if this is from client side.
func (s *End) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:220
	_go_fuzz_dep_.CoverTab[76034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:220
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:220
	// _ = "end of CoverTab[76034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:220
}

func (s *End) isRPCStats()	{ _go_fuzz_dep_.CoverTab[76035]++; // _ = "end of CoverTab[76035]" }

// ConnStats contains stats information about connections.
type ConnStats interface {
	isConnStats()
	// IsClient returns true if this ConnStats is from client side.
	IsClient() bool
}

// ConnBegin contains the stats of a connection when it is established.
type ConnBegin struct {
	// Client is true if this ConnBegin is from client side.
	Client bool
}

// IsClient indicates if this is from client side.
func (s *ConnBegin) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:238
	_go_fuzz_dep_.CoverTab[76036]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:238
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:238
	// _ = "end of CoverTab[76036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:238
}

func (s *ConnBegin) isConnStats()	{ _go_fuzz_dep_.CoverTab[76037]++; // _ = "end of CoverTab[76037]" }

// ConnEnd contains the stats of a connection when it ends.
type ConnEnd struct {
	// Client is true if this ConnEnd is from client side.
	Client bool
}

// IsClient indicates if this is from client side.
func (s *ConnEnd) IsClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:249
	_go_fuzz_dep_.CoverTab[76038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:249
	return s.Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:249
	// _ = "end of CoverTab[76038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:249
}

func (s *ConnEnd) isConnStats()	{ _go_fuzz_dep_.CoverTab[76039]++; // _ = "end of CoverTab[76039]" }

type incomingTagsKey struct{}
type outgoingTagsKey struct{}

// SetTags attaches stats tagging data to the context, which will be sent in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// the outgoing RPC with the header grpc-tags-bin.  Subsequent calls to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// SetTags will overwrite the values from earlier calls.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// NOTE: this is provided only for backward compatibility with existing clients
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// and will likely be removed in an upcoming release.  New uses should transmit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// this type of data using metadata with a different, non-reserved (i.e. does
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:256
// not begin with "grpc-") header name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:264
func SetTags(ctx context.Context, b []byte) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:264
	_go_fuzz_dep_.CoverTab[76040]++
											return context.WithValue(ctx, outgoingTagsKey{}, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:265
	// _ = "end of CoverTab[76040]"
}

// Tags returns the tags from the context for the inbound RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:268
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:268
// NOTE: this is provided only for backward compatibility with existing clients
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:268
// and will likely be removed in an upcoming release.  New uses should transmit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:268
// this type of data using metadata with a different, non-reserved (i.e. does
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:268
// not begin with "grpc-") header name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:274
func Tags(ctx context.Context) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:274
	_go_fuzz_dep_.CoverTab[76041]++
											b, _ := ctx.Value(incomingTagsKey{}).([]byte)
											return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:276
	// _ = "end of CoverTab[76041]"
}

// SetIncomingTags attaches stats tagging data to the context, to be read by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:279
// the application (not sent in outgoing RPCs).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:279
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:279
// This is intended for gRPC-internal use ONLY.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:283
func SetIncomingTags(ctx context.Context, b []byte) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:283
	_go_fuzz_dep_.CoverTab[76042]++
											return context.WithValue(ctx, incomingTagsKey{}, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:284
	// _ = "end of CoverTab[76042]"
}

// OutgoingTags returns the tags from the context for the outbound RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:287
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:287
// This is intended for gRPC-internal use ONLY.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:290
func OutgoingTags(ctx context.Context) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:290
	_go_fuzz_dep_.CoverTab[76043]++
											b, _ := ctx.Value(outgoingTagsKey{}).([]byte)
											return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:292
	// _ = "end of CoverTab[76043]"
}

type incomingTraceKey struct{}
type outgoingTraceKey struct{}

// SetTrace attaches stats tagging data to the context, which will be sent in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// the outgoing RPC with the header grpc-trace-bin.  Subsequent calls to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// SetTrace will overwrite the values from earlier calls.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// NOTE: this is provided only for backward compatibility with existing clients
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// and will likely be removed in an upcoming release.  New uses should transmit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// this type of data using metadata with a different, non-reserved (i.e. does
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:298
// not begin with "grpc-") header name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:306
func SetTrace(ctx context.Context, b []byte) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:306
	_go_fuzz_dep_.CoverTab[76044]++
											return context.WithValue(ctx, outgoingTraceKey{}, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:307
	// _ = "end of CoverTab[76044]"
}

// Trace returns the trace from the context for the inbound RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:310
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:310
// NOTE: this is provided only for backward compatibility with existing clients
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:310
// and will likely be removed in an upcoming release.  New uses should transmit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:310
// this type of data using metadata with a different, non-reserved (i.e. does
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:310
// not begin with "grpc-") header name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:316
func Trace(ctx context.Context) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:316
	_go_fuzz_dep_.CoverTab[76045]++
											b, _ := ctx.Value(incomingTraceKey{}).([]byte)
											return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:318
	// _ = "end of CoverTab[76045]"
}

// SetIncomingTrace attaches stats tagging data to the context, to be read by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:321
// the application (not sent in outgoing RPCs).  It is intended for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:321
// gRPC-internal use.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:324
func SetIncomingTrace(ctx context.Context, b []byte) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:324
	_go_fuzz_dep_.CoverTab[76046]++
											return context.WithValue(ctx, incomingTraceKey{}, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:325
	// _ = "end of CoverTab[76046]"
}

// OutgoingTrace returns the trace from the context for the outbound RPC.  It is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:328
// intended for gRPC-internal use.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:330
func OutgoingTrace(ctx context.Context) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:330
	_go_fuzz_dep_.CoverTab[76047]++
											b, _ := ctx.Value(outgoingTraceKey{}).([]byte)
											return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:332
	// _ = "end of CoverTab[76047]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:333
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/stats/stats.go:333
var _ = _go_fuzz_dep_.CoverTab
