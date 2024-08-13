//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:19
// Package resolver defines APIs for name resolution in gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
package resolver

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:21
)

import (
	"context"
	"net"
	"net/url"
	"strings"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/pretty"
	"google.golang.org/grpc/serviceconfig"
)

var (
	// m is a map from scheme to resolver builder.
	m	= make(map[string]Builder)
	// defaultScheme is the default scheme to use.
	defaultScheme	= "passthrough"
)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// Register registers the resolver builder to the resolver map. b.Scheme will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// be used as the scheme registered with this builder. The registry is case
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// sensitive, and schemes should not contain any uppercase characters.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// NOTE: this function must only be called during initialization time (i.e. in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// an init() function), and is not thread-safe. If multiple Resolvers are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:44
// registered with the same name, the one registered last will take effect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:51
func Register(b Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:51
	_go_fuzz_dep_.CoverTab[67310]++
												m[b.Scheme()] = b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:52
	// _ = "end of CoverTab[67310]"
}

// Get returns the resolver builder registered with the given scheme.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:55
// If no builder is register with the scheme, nil will be returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:58
func Get(scheme string) Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:58
	_go_fuzz_dep_.CoverTab[67311]++
												if b, ok := m[scheme]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:59
		_go_fuzz_dep_.CoverTab[67313]++
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:60
		// _ = "end of CoverTab[67313]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:61
		_go_fuzz_dep_.CoverTab[67314]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:61
		// _ = "end of CoverTab[67314]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:61
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:61
	// _ = "end of CoverTab[67311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:61
	_go_fuzz_dep_.CoverTab[67312]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:62
	// _ = "end of CoverTab[67312]"
}

// SetDefaultScheme sets the default scheme that will be used. The default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:65
// default scheme is "passthrough".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:65
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:65
// NOTE: this function must only be called during initialization time (i.e. in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:65
// an init() function), and is not thread-safe. The scheme set last overrides
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:65
// previously set values.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:71
func SetDefaultScheme(scheme string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:71
	_go_fuzz_dep_.CoverTab[67315]++
												defaultScheme = scheme
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:72
	// _ = "end of CoverTab[67315]"
}

// GetDefaultScheme gets the default scheme that will be used.
func GetDefaultScheme() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:76
	_go_fuzz_dep_.CoverTab[67316]++
												return defaultScheme
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:77
	// _ = "end of CoverTab[67316]"
}

// AddressType indicates the address type returned by name resolution.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:80
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:80
// Deprecated: use Attributes in Address instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:83
type AddressType uint8

const (
	// Backend indicates the address is for a backend server.
	//
	// Deprecated: use Attributes in Address instead.
	Backend	AddressType	= iota
	// GRPCLB indicates the address is for a grpclb load balancer.
	//
	// Deprecated: to select the GRPCLB load balancing policy, use a service
	// config with a corresponding loadBalancingConfig.  To supply balancer
	// addresses to the GRPCLB load balancing policy, set State.Attributes
	// using balancer/grpclb/state.Set.
	GRPCLB
)

// Address represents a server the client connects to.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:99
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:99
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:99
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:99
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:99
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:105
type Address struct {
	// Addr is the server address on which a connection will be established.
	Addr	string

	// ServerName is the name of this address.
	// If non-empty, the ServerName is used as the transport certification authority for
	// the address, instead of the hostname from the Dial target string. In most cases,
	// this should not be set.
	//
	// If Type is GRPCLB, ServerName should be the name of the remote load
	// balancer, not the name of the backend.
	//
	// WARNING: ServerName must only be populated with trusted values. It
	// is insecure to populate it with data from untrusted inputs since untrusted
	// values could be used to bypass the authority checks performed by TLS.
	ServerName	string

	// Attributes contains arbitrary data about this address intended for
	// consumption by the SubConn.
	Attributes	*attributes.Attributes

	// BalancerAttributes contains arbitrary data about this address intended
	// for consumption by the LB policy.  These attribes do not affect SubConn
	// creation, connection establishment, handshaking, etc.
	BalancerAttributes	*attributes.Attributes

	// Type is the type of this address.
	//
	// Deprecated: use Attributes instead.
	Type	AddressType

	// Metadata is the information associated with Addr, which may be used
	// to make load balancing decision.
	//
	// Deprecated: use Attributes instead.
	Metadata	interface{}
}

// Equal returns whether a and o are identical.  Metadata is compared directly,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:143
// not with any recursive introspection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:145
func (a Address) Equal(o Address) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:145
	_go_fuzz_dep_.CoverTab[67317]++
												return a.Addr == o.Addr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
		_go_fuzz_dep_.CoverTab[67318]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
		return a.ServerName == o.ServerName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
		// _ = "end of CoverTab[67318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
		_go_fuzz_dep_.CoverTab[67319]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:146
		return a.Attributes.Equal(o.Attributes)
													// _ = "end of CoverTab[67319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:147
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:147
		_go_fuzz_dep_.CoverTab[67320]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:147
		return a.BalancerAttributes.Equal(o.BalancerAttributes)
													// _ = "end of CoverTab[67320]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:148
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:148
		_go_fuzz_dep_.CoverTab[67321]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:148
		return a.Type == o.Type
													// _ = "end of CoverTab[67321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
		_go_fuzz_dep_.CoverTab[67322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
		return a.Metadata == o.Metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
		// _ = "end of CoverTab[67322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:149
	// _ = "end of CoverTab[67317]"
}

// String returns JSON formatted string representation of the address.
func (a Address) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:153
	_go_fuzz_dep_.CoverTab[67323]++
												return pretty.ToJSON(a)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:154
	// _ = "end of CoverTab[67323]"
}

// BuildOptions includes additional information for the builder to create
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:157
// the resolver.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:159
type BuildOptions struct {
	// DisableServiceConfig indicates whether a resolver implementation should
	// fetch service config data.
	DisableServiceConfig	bool
	// DialCreds is the transport credentials used by the ClientConn for
	// communicating with the target gRPC service (set via
	// WithTransportCredentials). In cases where a name resolution service
	// requires the same credentials, the resolver may use this field. In most
	// cases though, it is not appropriate, and this field may be ignored.
	DialCreds	credentials.TransportCredentials
	// CredsBundle is the credentials bundle used by the ClientConn for
	// communicating with the target gRPC service (set via
	// WithCredentialsBundle). In cases where a name resolution service
	// requires the same credentials, the resolver may use this field. In most
	// cases though, it is not appropriate, and this field may be ignored.
	CredsBundle	credentials.Bundle
	// Dialer is the custom dialer used by the ClientConn for dialling the
	// target gRPC service (set via WithDialer). In cases where a name
	// resolution service requires the same dialer, the resolver may use this
	// field. In most cases though, it is not appropriate, and this field may
	// be ignored.
	Dialer	func(context.Context, string) (net.Conn, error)
}

// State contains the current Resolver state relevant to the ClientConn.
type State struct {
	// Addresses is the latest set of resolved addresses for the target.
	Addresses	[]Address

	// ServiceConfig contains the result from parsing the latest service
	// config.  If it is nil, it indicates no service config is present or the
	// resolver does not provide service configs.
	ServiceConfig	*serviceconfig.ParseResult

	// Attributes contains arbitrary data about the resolver intended for
	// consumption by the load balancing policy.
	Attributes	*attributes.Attributes
}

// ClientConn contains the callbacks for resolver to notify any updates
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
// to the gRPC ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
// This interface is to be implemented by gRPC. Users should not need a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
// brand new implementation of this interface. For the situations like
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
// testing, the new implementation should embed this interface. This allows
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:198
// gRPC to add new methods to this interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:205
type ClientConn interface {
	// UpdateState updates the state of the ClientConn appropriately.
	//
	// If an error is returned, the resolver should try to resolve the
	// target again. The resolver should use a backoff timer to prevent
	// overloading the server with requests. If a resolver is certain that
	// reresolving will not change the result, e.g. because it is
	// a watch-based resolver, returned errors can be ignored.
	//
	// If the resolved State is the same as the last reported one, calling
	// UpdateState can be omitted.
	UpdateState(State) error
	// ReportError notifies the ClientConn that the Resolver encountered an
	// error.  The ClientConn will notify the load balancer and begin calling
	// ResolveNow on the Resolver with exponential backoff.
	ReportError(error)
	// NewAddress is called by resolver to notify ClientConn a new list
	// of resolved addresses.
	// The address list should be the complete list of resolved addresses.
	//
	// Deprecated: Use UpdateState instead.
	NewAddress(addresses []Address)
	// NewServiceConfig is called by resolver to notify ClientConn a new
	// service config. The service config should be provided as a json string.
	//
	// Deprecated: Use UpdateState instead.
	NewServiceConfig(serviceConfig string)
	// ParseServiceConfig parses the provided service config and returns an
	// object that provides the parsed config.
	ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult
}

// Target represents a target for gRPC, as specified in:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// https://github.com/grpc/grpc/blob/master/doc/naming.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// It is parsed from the target string that gets passed into Dial or DialContext
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// by the user. And gRPC passes it to the resolver and the balancer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// If the target follows the naming spec, and the parsed scheme is registered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// with gRPC, we will parse the target string according to the spec. If the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// target does not contain a scheme or if the parsed scheme is not registered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// (i.e. no corresponding resolver available to resolve the endpoint), we will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// apply the default scheme, and will attempt to reparse it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
// Examples:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//   - "dns://some_authority/foo.bar"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//     Target{Scheme: "dns", Authority: "some_authority", Endpoint: "foo.bar"}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//   - "foo.bar"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//     Target{Scheme: resolver.GetDefaultScheme(), Endpoint: "foo.bar"}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//   - "unknown_scheme://authority/endpoint"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:237
//     Target{Scheme: resolver.GetDefaultScheme(), Endpoint: "unknown_scheme://authority/endpoint"}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:256
type Target struct {
	// Deprecated: use URL.Scheme instead.
	Scheme	string
	// Deprecated: use URL.Host instead.
	Authority	string
	// URL contains the parsed dial target with an optional default scheme added
	// to it if the original dial target contained no scheme or contained an
	// unregistered scheme. Any query params specified in the original dial
	// target can be accessed from here.
	URL	url.URL
}

// Endpoint retrieves endpoint without leading "/" from either `URL.Path`
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:268
// or `URL.Opaque`. The latter is used when the former is empty.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:270
func (t Target) Endpoint() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:270
	_go_fuzz_dep_.CoverTab[67324]++
												endpoint := t.URL.Path
												if endpoint == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:272
		_go_fuzz_dep_.CoverTab[67326]++
													endpoint = t.URL.Opaque
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:273
		// _ = "end of CoverTab[67326]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:274
		_go_fuzz_dep_.CoverTab[67327]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:274
		// _ = "end of CoverTab[67327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:274
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:274
	// _ = "end of CoverTab[67324]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:274
	_go_fuzz_dep_.CoverTab[67325]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:283
	return strings.TrimPrefix(endpoint, "/")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:283
	// _ = "end of CoverTab[67325]"
}

// Builder creates a resolver that will be used to watch name resolution updates.
type Builder interface {
	// Build creates a new resolver for the given target.
	//
	// gRPC dial calls Build synchronously, and fails if the returned error is
	// not nil.
	Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)
	// Scheme returns the scheme supported by this resolver.  Scheme is defined
	// at https://github.com/grpc/grpc/blob/master/doc/naming.md.  The returned
	// string should not contain uppercase characters, as they will not match
	// the parsed target's scheme as defined in RFC 3986.
	Scheme() string
}

// ResolveNowOptions includes additional information for ResolveNow.
type ResolveNowOptions struct{}

// Resolver watches for the updates on the specified target.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:303
// Updates include address updates and service config updates.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:305
type Resolver interface {
	// ResolveNow will be called by gRPC to try to resolve the target name
	// again. It's just a hint, resolver can ignore this if it's not necessary.
	//
	// It could be called multiple times concurrently.
	ResolveNow(ResolveNowOptions)
	// Close closes the resolver.
	Close()
}

// UnregisterForTesting removes the resolver builder with the given scheme from the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:315
// resolver map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:315
// This function is for testing only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:318
func UnregisterForTesting(scheme string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:318
	_go_fuzz_dep_.CoverTab[67328]++
												delete(m, scheme)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:319
	// _ = "end of CoverTab[67328]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:320
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/resolver.go:320
var _ = _go_fuzz_dep_.CoverTab
