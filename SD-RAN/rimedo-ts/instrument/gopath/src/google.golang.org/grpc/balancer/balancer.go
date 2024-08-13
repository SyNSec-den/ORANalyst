//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:19
// Package balancer defines APIs for load balancing in gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
package balancer

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:21
)

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"strings"

	"google.golang.org/grpc/channelz"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

var (
	// m is a map from name to balancer builder.
	m = make(map[string]Builder)
)

// Register registers the balancer builder to the balancer map. b.Name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// (lowercased) will be used as the name registered with this builder.  If the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// Builder implements ConfigParser, ParseConfig will be called when new service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// configs are received by the resolver, and the result will be provided to the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// Balancer in UpdateClientConnState.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// NOTE: this function must only be called during initialization time (i.e. in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// an init() function), and is not thread-safe. If multiple Balancers are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:44
// registered with the same name, the one registered last will take effect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:53
func Register(b Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:53
	_go_fuzz_dep_.CoverTab[67403]++
												m[strings.ToLower(b.Name())] = b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:54
	// _ = "end of CoverTab[67403]"
}

// unregisterForTesting deletes the balancer with the given name from the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:57
// balancer map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:57
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:57
// This function is not thread-safe.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:61
func unregisterForTesting(name string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:61
	_go_fuzz_dep_.CoverTab[67404]++
												delete(m, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:62
	// _ = "end of CoverTab[67404]"
}

func init() {
	internal.BalancerUnregister = unregisterForTesting
}

// Get returns the resolver builder registered with the given name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:69
// Note that the compare is done in a case-insensitive fashion.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:69
// If no builder is register with the name, nil will be returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:72
func Get(name string) Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:72
	_go_fuzz_dep_.CoverTab[67405]++
												if b, ok := m[strings.ToLower(name)]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:73
		_go_fuzz_dep_.CoverTab[67407]++
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:74
		// _ = "end of CoverTab[67407]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:75
		_go_fuzz_dep_.CoverTab[67408]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:75
		// _ = "end of CoverTab[67408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:75
	// _ = "end of CoverTab[67405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:75
	_go_fuzz_dep_.CoverTab[67406]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:76
	// _ = "end of CoverTab[67406]"
}

// A SubConn represents a single connection to a gRPC backend service.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// Each SubConn contains a list of addresses.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// All SubConns start in IDLE, and will not try to connect. To trigger the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// connecting, Balancers must call Connect.  If a connection re-enters IDLE,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// Balancers must call Connect again to trigger a new connection attempt.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// gRPC will try to connect to the addresses in sequence, and stop trying the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// remainder once the first connection is successful. If an attempt to connect
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// to all addresses encounters an error, the SubConn will enter
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// TRANSIENT_FAILURE for a backoff period, and then transition to IDLE.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// Once established, if a connection is lost, the SubConn will transition
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// directly to IDLE.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// This interface is to be implemented by gRPC. Users should not need their own
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// implementation of this interface. For situations like testing, any
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// implementations should embed this interface. This allows gRPC to add new
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:79
// methods to this interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:99
type SubConn interface {
	// UpdateAddresses updates the addresses used in this SubConn.
	// gRPC checks if currently-connected address is still in the new list.
	// If it's in the list, the connection will be kept.
	// If it's not in the list, the connection will gracefully closed, and
	// a new connection will be created.
	//
	// This will trigger a state transition for the SubConn.
	//
	// Deprecated: This method is now part of the ClientConn interface and will
	// eventually be removed from here.
	UpdateAddresses([]resolver.Address)
	// Connect starts the connecting for this SubConn.
	Connect()
	// GetOrBuildProducer returns a reference to the existing Producer for this
	// ProducerBuilder in this SubConn, or, if one does not currently exist,
	// creates a new one and returns it.  Returns a close function which must
	// be called when the Producer is no longer needed.
	GetOrBuildProducer(ProducerBuilder) (p Producer, close func())
}

// NewSubConnOptions contains options to create new SubConn.
type NewSubConnOptions struct {
	// CredsBundle is the credentials bundle that will be used in the created
	// SubConn. If it's nil, the original creds from grpc DialOptions will be
	// used.
	//
	// Deprecated: Use the Attributes field in resolver.Address to pass
	// arbitrary data to the credential handshaker.
	CredsBundle	credentials.Bundle
	// HealthCheckEnabled indicates whether health check service should be
	// enabled on this SubConn
	HealthCheckEnabled	bool
}

// State contains the balancer's state relevant to the gRPC ClientConn.
type State struct {
	// State contains the connectivity state of the balancer, which is used to
	// determine the state of the ClientConn.
	ConnectivityState	connectivity.State
	// Picker is used to choose connections (SubConns) for RPCs.
	Picker	Picker
}

// ClientConn represents a gRPC ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:143
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:143
// This interface is to be implemented by gRPC. Users should not need a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:143
// brand new implementation of this interface. For the situations like
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:143
// testing, the new implementation should embed this interface. This allows
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:143
// gRPC to add new methods to this interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:149
type ClientConn interface {
	// NewSubConn is called by balancer to create a new SubConn.
	// It doesn't block and wait for the connections to be established.
	// Behaviors of the SubConn can be controlled by options.
	NewSubConn([]resolver.Address, NewSubConnOptions) (SubConn, error)
	// RemoveSubConn removes the SubConn from ClientConn.
	// The SubConn will be shutdown.
	RemoveSubConn(SubConn)
	// UpdateAddresses updates the addresses used in the passed in SubConn.
	// gRPC checks if the currently connected address is still in the new list.
	// If so, the connection will be kept. Else, the connection will be
	// gracefully closed, and a new connection will be created.
	//
	// This will trigger a state transition for the SubConn.
	UpdateAddresses(SubConn, []resolver.Address)

	// UpdateState notifies gRPC that the balancer's internal state has
	// changed.
	//
	// gRPC will update the connectivity state of the ClientConn, and will call
	// Pick on the new Picker to pick new SubConns.
	UpdateState(State)

	// ResolveNow is called by balancer to notify gRPC to do a name resolving.
	ResolveNow(resolver.ResolveNowOptions)

	// Target returns the dial target for this ClientConn.
	//
	// Deprecated: Use the Target field in the BuildOptions instead.
	Target() string
}

// BuildOptions contains additional information for Build.
type BuildOptions struct {
	// DialCreds is the transport credentials to use when communicating with a
	// remote load balancer server. Balancer implementations which do not
	// communicate with a remote load balancer server can ignore this field.
	DialCreds	credentials.TransportCredentials
	// CredsBundle is the credentials bundle to use when communicating with a
	// remote load balancer server. Balancer implementations which do not
	// communicate with a remote load balancer server can ignore this field.
	CredsBundle	credentials.Bundle
	// Dialer is the custom dialer to use when communicating with a remote load
	// balancer server. Balancer implementations which do not communicate with a
	// remote load balancer server can ignore this field.
	Dialer	func(context.Context, string) (net.Conn, error)
	// Authority is the server name to use as part of the authentication
	// handshake when communicating with a remote load balancer server. Balancer
	// implementations which do not communicate with a remote load balancer
	// server can ignore this field.
	Authority	string
	// ChannelzParentID is the parent ClientConn's channelz ID.
	ChannelzParentID	*channelz.Identifier
	// CustomUserAgent is the custom user agent set on the parent ClientConn.
	// The balancer should set the same custom user agent if it creates a
	// ClientConn.
	CustomUserAgent	string
	// Target contains the parsed address info of the dial target. It is the
	// same resolver.Target as passed to the resolver. See the documentation for
	// the resolver.Target type for details about what it contains.
	Target	resolver.Target
}

// Builder creates a balancer.
type Builder interface {
	// Build creates a new balancer with the ClientConn.
	Build(cc ClientConn, opts BuildOptions) Balancer
	// Name returns the name of balancers built by this builder.
	// It will be used to pick balancers (for example in service config).
	Name() string
}

// ConfigParser parses load balancer configs.
type ConfigParser interface {
	// ParseConfig parses the JSON load balancer config provided into an
	// internal form or returns an error if the config is invalid.  For future
	// compatibility reasons, unknown fields in the config should be ignored.
	ParseConfig(LoadBalancingConfigJSON json.RawMessage) (serviceconfig.LoadBalancingConfig, error)
}

// PickInfo contains additional information for the Pick operation.
type PickInfo struct {
	// FullMethodName is the method name that NewClientStream() is called
	// with. The canonical format is /service/Method.
	FullMethodName	string
	// Ctx is the RPC's context, and may contain relevant RPC-level information
	// like the outgoing header metadata.
	Ctx	context.Context
}

// DoneInfo contains additional information for done.
type DoneInfo struct {
	// Err is the rpc error the RPC finished with. It could be nil.
	Err	error
	// Trailer contains the metadata from the RPC's trailer, if present.
	Trailer	metadata.MD
	// BytesSent indicates if any bytes have been sent to the server.
	BytesSent	bool
	// BytesReceived indicates if any byte has been received from the server.
	BytesReceived	bool
	// ServerLoad is the load received from server. It's usually sent as part of
	// trailing metadata.
	//
	// The only supported type now is *orca_v3.LoadReport.
	ServerLoad	interface{}
}

var (
	// ErrNoSubConnAvailable indicates no SubConn is available for pick().
	// gRPC will block the RPC until a new picker is available via UpdateState().
	ErrNoSubConnAvailable	= errors.New("no SubConn is available")
	// ErrTransientFailure indicates all SubConns are in TransientFailure.
	// WaitForReady RPCs will block, non-WaitForReady RPCs will fail.
	//
	// Deprecated: return an appropriate error based on the last resolution or
	// connection attempt instead.  The behavior is the same for any non-gRPC
	// status error.
	ErrTransientFailure	= errors.New("all SubConns are in TransientFailure")
)

// PickResult contains information related to a connection chosen for an RPC.
type PickResult struct {
	// SubConn is the connection to use for this pick, if its state is Ready.
	// If the state is not Ready, gRPC will block the RPC until a new Picker is
	// provided by the balancer (using ClientConn.UpdateState).  The SubConn
	// must be one returned by ClientConn.NewSubConn.
	SubConn	SubConn

	// Done is called when the RPC is completed.  If the SubConn is not ready,
	// this will be called with a nil parameter.  If the SubConn is not a valid
	// type, Done may not be called.  May be nil if the balancer does not wish
	// to be notified when the RPC completes.
	Done	func(DoneInfo)

	// Metadata provides a way for LB policies to inject arbitrary per-call
	// metadata. Any metadata returned here will be merged with existing
	// metadata added by the client application.
	//
	// LB policies with child policies are responsible for propagating metadata
	// injected by their children to the ClientConn, as part of Pick().
	Metatada	metadata.MD
}

// TransientFailureError returns e.  It exists for backward compatibility and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:292
// will be deleted soon.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:292
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:292
// Deprecated: no longer necessary, picker errors are treated this way by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:292
// default.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:297
func TransientFailureError(e error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:297
	_go_fuzz_dep_.CoverTab[67409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:297
	return e
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:297
	// _ = "end of CoverTab[67409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:297
}

// Picker is used by gRPC to pick a SubConn to send an RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:299
// Balancer is expected to generate a new picker from its snapshot every time its
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:299
// internal state has changed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:299
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:299
// The pickers used by gRPC can be updated by ClientConn.UpdateState().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:304
type Picker interface {
	// Pick returns the connection to use for this RPC and related information.
	//
	// Pick should not block.  If the balancer needs to do I/O or any blocking
	// or time-consuming work to service this call, it should return
	// ErrNoSubConnAvailable, and the Pick call will be repeated by gRPC when
	// the Picker is updated (using ClientConn.UpdateState).
	//
	// If an error is returned:
	//
	// - If the error is ErrNoSubConnAvailable, gRPC will block until a new
	//   Picker is provided by the balancer (using ClientConn.UpdateState).
	//
	// - If the error is a status error (implemented by the grpc/status
	//   package), gRPC will terminate the RPC with the code and message
	//   provided.
	//
	// - For all other errors, wait for ready RPCs will wait, but non-wait for
	//   ready RPCs will be terminated with this error's Error() string and
	//   status code Unavailable.
	Pick(info PickInfo) (PickResult, error)
}

// Balancer takes input from gRPC, manages SubConns, and collects and aggregates
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
// the connectivity states.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
// It also generates and updates the Picker used by gRPC to pick SubConns for RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
// UpdateClientConnState, ResolverError, UpdateSubConnState, and Close are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
// guaranteed to be called synchronously from the same goroutine.  There's no
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:327
// guarantee on picker.Pick, it may be called anytime.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:335
type Balancer interface {
	// UpdateClientConnState is called by gRPC when the state of the ClientConn
	// changes.  If the error returned is ErrBadResolverState, the ClientConn
	// will begin calling ResolveNow on the active name resolver with
	// exponential backoff until a subsequent call to UpdateClientConnState
	// returns a nil error.  Any other errors are currently ignored.
	UpdateClientConnState(ClientConnState) error
	// ResolverError is called by gRPC when the name resolver reports an error.
	ResolverError(error)
	// UpdateSubConnState is called by gRPC when the state of a SubConn
	// changes.
	UpdateSubConnState(SubConn, SubConnState)
	// Close closes the balancer. The balancer is not required to call
	// ClientConn.RemoveSubConn for its existing SubConns.
	Close()
}

// ExitIdler is an optional interface for balancers to implement.  If
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
// implemented, ExitIdle will be called when ClientConn.Connect is called, if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
// the ClientConn is idle.  If unimplemented, ClientConn.Connect will cause
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
// all SubConns to connect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
// Notice: it will be required for all balancers to implement this in a future
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:352
// release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:359
type ExitIdler interface {
	// ExitIdle instructs the LB policy to reconnect to backends / exit the
	// IDLE state, if appropriate and possible.  Note that SubConns that enter
	// the IDLE state will not reconnect until SubConn.Connect is called.
	ExitIdle()
}

// SubConnState describes the state of a SubConn.
type SubConnState struct {
	// ConnectivityState is the connectivity state of the SubConn.
	ConnectivityState	connectivity.State
	// ConnectionError is set if the ConnectivityState is TransientFailure,
	// describing the reason the SubConn failed.  Otherwise, it is nil.
	ConnectionError	error
}

// ClientConnState describes the state of a ClientConn relevant to the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:375
// balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:377
type ClientConnState struct {
	ResolverState	resolver.State
	// The parsed load balancing configuration returned by the builder's
	// ParseConfig method, if implemented.
	BalancerConfig	serviceconfig.LoadBalancingConfig
}

// ErrBadResolverState may be returned by UpdateClientConnState to indicate a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:384
// problem with the provided name resolver data.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:386
var ErrBadResolverState = errors.New("bad resolver state")

// A ProducerBuilder is a simple constructor for a Producer.  It is used by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:388
// SubConn to create producers when needed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:390
type ProducerBuilder interface {
	// Build creates a Producer.  The first parameter is always a
	// grpc.ClientConnInterface (a type to allow creating RPCs/streams on the
	// associated SubConn), but is declared as interface{} to avoid a
	// dependency cycle.  Should also return a close function that will be
	// called when all references to the Producer have been given up.
	Build(grpcClientConnInterface interface{}) (p Producer, close func())
}

// A Producer is a type shared among potentially many consumers.  It is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:399
// associated with a SubConn, and an implementation will typically contain
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:399
// other methods to provide additional functionality, e.g. configuration or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:399
// subscription registration.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:403
type Producer interface {
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:404
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/balancer/balancer.go:404
var _ = _go_fuzz_dep_.CoverTab
