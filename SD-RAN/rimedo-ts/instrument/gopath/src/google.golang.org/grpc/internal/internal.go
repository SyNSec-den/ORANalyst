//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:18
// Package internal contains gRPC-internal code, to avoid polluting
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:18
// the godoc of the top-level grpc package.  It must not import any grpc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:18
// symbols to avoid circular dependencies.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
package internal

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:21
)

import (
	"context"
	"time"

	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/serviceconfig"
)

var (
	// WithHealthCheckFunc is set by dialoptions.go
	WithHealthCheckFunc	interface{}	// func (HealthChecker) DialOption
	// HealthCheckFunc is used to provide client-side LB channel health checking
	HealthCheckFunc	HealthChecker
	// BalancerUnregister is exported by package balancer to unregister a balancer.
	BalancerUnregister	func(name string)
	// KeepaliveMinPingTime is the minimum ping interval.  This must be 10s by
	// default, but tests may wish to set it lower for convenience.
	KeepaliveMinPingTime	= 10 * time.Second
	// ParseServiceConfig parses a JSON representation of the service config.
	ParseServiceConfig	interface{}	// func(string) *serviceconfig.ParseResult
	// EqualServiceConfigForTesting is for testing service config generation and
	// parsing. Both a and b should be returned by ParseServiceConfig.
	// This function compares the config without rawJSON stripped, in case the
	// there's difference in white space.
	EqualServiceConfigForTesting	func(a, b serviceconfig.Config) bool
	// GetCertificateProviderBuilder returns the registered builder for the
	// given name. This is set by package certprovider for use from xDS
	// bootstrap code while parsing certificate provider configs in the
	// bootstrap file.
	GetCertificateProviderBuilder	interface{}	// func(string) certprovider.Builder
	// GetXDSHandshakeInfoForTesting returns a pointer to the xds.HandshakeInfo
	// stored in the passed in attributes. This is set by
	// credentials/xds/xds.go.
	GetXDSHandshakeInfoForTesting	interface{}	// func (*attributes.Attributes) *xds.HandshakeInfo
	// GetServerCredentials returns the transport credentials configured on a
	// gRPC server. An xDS-enabled server needs to know what type of credentials
	// is configured on the underlying gRPC server. This is set by server.go.
	GetServerCredentials	interface{}	// func (*grpc.Server) credentials.TransportCredentials
	// CanonicalString returns the canonical string of the code defined here:
	// https://github.com/grpc/grpc/blob/master/doc/statuscodes.md.
	CanonicalString	interface{}	// func (codes.Code) string
	// DrainServerTransports initiates a graceful close of existing connections
	// on a gRPC server accepted on the provided listener address. An
	// xDS-enabled server invokes this method on a grpc.Server when a particular
	// listener moves to "not-serving" mode.
	DrainServerTransports	interface{}	// func(*grpc.Server, string)
	// AddGlobalServerOptions adds an array of ServerOption that will be
	// effective globally for newly created servers. The priority will be: 1.
	// user-provided; 2. this method; 3. default values.
	AddGlobalServerOptions	interface{}	// func(opt ...ServerOption)
	// ClearGlobalServerOptions clears the array of extra ServerOption. This
	// method is useful in testing and benchmarking.
	ClearGlobalServerOptions	func()
	// AddGlobalDialOptions adds an array of DialOption that will be effective
	// globally for newly created client channels. The priority will be: 1.
	// user-provided; 2. this method; 3. default values.
	AddGlobalDialOptions	interface{}	// func(opt ...DialOption)
	// DisableGlobalDialOptions returns a DialOption that prevents the
	// ClientConn from applying the global DialOptions (set via
	// AddGlobalDialOptions).
	DisableGlobalDialOptions	interface{}	// func() grpc.DialOption
	// ClearGlobalDialOptions clears the array of extra DialOption. This
	// method is useful in testing and benchmarking.
	ClearGlobalDialOptions	func()
	// JoinDialOptions combines the dial options passed as arguments into a
	// single dial option.
	JoinDialOptions	interface{}	// func(...grpc.DialOption) grpc.DialOption
	// JoinServerOptions combines the server options passed as arguments into a
	// single server option.
	JoinServerOptions	interface{}	// func(...grpc.ServerOption) grpc.ServerOption

	// WithBinaryLogger returns a DialOption that specifies the binary logger
	// for a ClientConn.
	WithBinaryLogger	interface{}	// func(binarylog.Logger) grpc.DialOption
	// BinaryLogger returns a ServerOption that can set the binary logger for a
	// server.
	BinaryLogger	interface{}	// func(binarylog.Logger) grpc.ServerOption

	// NewXDSResolverWithConfigForTesting creates a new xds resolver builder using
	// the provided xds bootstrap config instead of the global configuration from
	// the supported environment variables.  The resolver.Builder is meant to be
	// used in conjunction with the grpc.WithResolvers DialOption.
	//
	// Testing Only
	//
	// This function should ONLY be used for testing and may not work with some
	// other features, including the CSDS service.
	NewXDSResolverWithConfigForTesting	interface{}	// func([]byte) (resolver.Builder, error)

	// RegisterRLSClusterSpecifierPluginForTesting registers the RLS Cluster
	// Specifier Plugin for testing purposes, regardless of the XDSRLS environment
	// variable.
	//
	// TODO: Remove this function once the RLS env var is removed.
	RegisterRLSClusterSpecifierPluginForTesting	func()

	// UnregisterRLSClusterSpecifierPluginForTesting unregisters the RLS Cluster
	// Specifier Plugin for testing purposes. This is needed because there is no way
	// to unregister the RLS Cluster Specifier Plugin after registering it solely
	// for testing purposes using RegisterRLSClusterSpecifierPluginForTesting().
	//
	// TODO: Remove this function once the RLS env var is removed.
	UnregisterRLSClusterSpecifierPluginForTesting	func()

	// RegisterRBACHTTPFilterForTesting registers the RBAC HTTP Filter for testing
	// purposes, regardless of the RBAC environment variable.
	//
	// TODO: Remove this function once the RBAC env var is removed.
	RegisterRBACHTTPFilterForTesting	func()

	// UnregisterRBACHTTPFilterForTesting unregisters the RBAC HTTP Filter for
	// testing purposes. This is needed because there is no way to unregister the
	// HTTP Filter after registering it solely for testing purposes using
	// RegisterRBACHTTPFilterForTesting().
	//
	// TODO: Remove this function once the RBAC env var is removed.
	UnregisterRBACHTTPFilterForTesting	func()

	// ORCAAllowAnyMinReportingInterval is for examples/orca use ONLY.
	ORCAAllowAnyMinReportingInterval	interface{}	// func(so *orca.ServiceOptions)
)

// HealthChecker defines the signature of the client-side LB channel health checking function.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
// The implementation is expected to create a health checking RPC stream by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
// calling newStream(), watch for the health status of serviceName, and report
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
// it's health back by calling setConnectivityState().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
// The health checking protocol is defined at:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:145
// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:153
type HealthChecker func(ctx context.Context, newStream func(string) (interface{}, error), setConnectivityState func(connectivity.State, error), serviceName string) error

const (
	// CredsBundleModeFallback switches GoogleDefaultCreds to fallback mode.
	CredsBundleModeFallback	= "fallback"
	// CredsBundleModeBalancer switches GoogleDefaultCreds to grpclb balancer
	// mode.
	CredsBundleModeBalancer	= "balancer"
	// CredsBundleModeBackendFromBalancer switches GoogleDefaultCreds to mode
	// that supports backend returned by grpclb balancer.
	CredsBundleModeBackendFromBalancer	= "backend-from-balancer"
)

// RLSLoadBalancingPolicyName is the name of the RLS LB policy.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:166
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:166
// It currently has an experimental suffix which would be removed once
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:166
// end-to-end testing of the policy is completed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:170
const RLSLoadBalancingPolicyName = "rls_experimental"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:170
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/internal.go:170
var _ = _go_fuzz_dep_.CoverTab
