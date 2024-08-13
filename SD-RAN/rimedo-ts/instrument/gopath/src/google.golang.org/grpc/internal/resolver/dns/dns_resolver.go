//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:19
// Package dns implements a dns resolver to be installed as the default resolver
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:19
// in grpc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
package dns

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:21
)

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	grpclbstate "google.golang.org/grpc/balancer/grpclb/state"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/backoff"
	"google.golang.org/grpc/internal/envconfig"
	"google.golang.org/grpc/internal/grpcrand"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// EnableSRVLookups controls whether the DNS resolver attempts to fetch gRPCLB
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:44
// addresses from SRV records.  Must not be changed after init time.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:46
var EnableSRVLookups = false

var logger = grpclog.Component("dns")

// Globals to stub out in tests. TODO: Perhaps these two can be combined into a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:50
// single variable for testing the resolver?
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:52
var (
	newTimer		= time.NewTimer
	newTimerDNSResRate	= time.NewTimer
)

func init() {
	resolver.Register(NewBuilder())
}

const (
	defaultPort		= "443"
	defaultDNSSvrPort	= "53"
	golang			= "GO"
	// txtPrefix is the prefix string to be prepended to the host name for txt record lookup.
	txtPrefix	= "_grpc_config."
	// In DNS, service config is encoded in a TXT record via the mechanism
	// described in RFC-1464 using the attribute name grpc_config.
	txtAttribute	= "grpc_config="
)

var (
	errMissingAddr	= errors.New("dns resolver: missing address")

	// Addresses ending with a colon that is supposed to be the separator
	// between host and port is not allowed.  E.g. "::" is a valid address as
	// it is an IPv6 address (host only) and "[::]:" is invalid as it ends with
	// a colon as the host and port separator
	errEndsWithColon	= errors.New("dns resolver: missing port after port-separator colon")
)

var (
	defaultResolver	netResolver	= net.DefaultResolver
	// To prevent excessive re-resolution, we enforce a rate limit on DNS
	// resolution requests.
	minDNSResRate	= 30 * time.Second
)

var customAuthorityDialler = func(authority string) func(ctx context.Context, network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:89
	_go_fuzz_dep_.CoverTab[68993]++
														return func(ctx context.Context, network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:90
		_go_fuzz_dep_.CoverTab[68994]++
															var dialer net.Dialer
															return dialer.DialContext(ctx, network, authority)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:92
		// _ = "end of CoverTab[68994]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:93
	// _ = "end of CoverTab[68993]"
}

var customAuthorityResolver = func(authority string) (netResolver, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:96
	_go_fuzz_dep_.CoverTab[68995]++
														host, port, err := parseTarget(authority, defaultDNSSvrPort)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:98
		_go_fuzz_dep_.CoverTab[68997]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:99
		// _ = "end of CoverTab[68997]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:100
		_go_fuzz_dep_.CoverTab[68998]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:100
		// _ = "end of CoverTab[68998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:100
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:100
	// _ = "end of CoverTab[68995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:100
	_go_fuzz_dep_.CoverTab[68996]++

														authorityWithPort := net.JoinHostPort(host, port)

														return &net.Resolver{
		PreferGo:	true,
		Dial:		customAuthorityDialler(authorityWithPort),
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:107
	// _ = "end of CoverTab[68996]"
}

// NewBuilder creates a dnsBuilder which is used to factory DNS resolvers.
func NewBuilder() resolver.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:111
	_go_fuzz_dep_.CoverTab[68999]++
														return &dnsBuilder{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:112
	// _ = "end of CoverTab[68999]"
}

type dnsBuilder struct{}

// Build creates and starts a DNS resolver that watches the name resolution of the target.
func (b *dnsBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:118
	_go_fuzz_dep_.CoverTab[69000]++
														host, port, err := parseTarget(target.Endpoint(), defaultPort)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:120
		_go_fuzz_dep_.CoverTab[69004]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:121
		// _ = "end of CoverTab[69004]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:122
		_go_fuzz_dep_.CoverTab[69005]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:122
		// _ = "end of CoverTab[69005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:122
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:122
	// _ = "end of CoverTab[69000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:122
	_go_fuzz_dep_.CoverTab[69001]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:125
	if ipAddr, ok := formatIP(host); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:125
		_go_fuzz_dep_.CoverTab[69006]++
															addr := []resolver.Address{{Addr: ipAddr + ":" + port}}
															cc.UpdateState(resolver.State{Addresses: addr})
															return deadResolver{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:128
		// _ = "end of CoverTab[69006]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:129
		_go_fuzz_dep_.CoverTab[69007]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:129
		// _ = "end of CoverTab[69007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:129
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:129
	// _ = "end of CoverTab[69001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:129
	_go_fuzz_dep_.CoverTab[69002]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:132
	ctx, cancel := context.WithCancel(context.Background())
	d := &dnsResolver{
		host:			host,
		port:			port,
		ctx:			ctx,
		cancel:			cancel,
		cc:			cc,
		rn:			make(chan struct{}, 1),
		disableServiceConfig:	opts.DisableServiceConfig,
	}

	if target.URL.Host == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:143
		_go_fuzz_dep_.CoverTab[69008]++
															d.resolver = defaultResolver
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:144
		// _ = "end of CoverTab[69008]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:145
		_go_fuzz_dep_.CoverTab[69009]++
															d.resolver, err = customAuthorityResolver(target.URL.Host)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:147
			_go_fuzz_dep_.CoverTab[69010]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:148
			// _ = "end of CoverTab[69010]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:149
			_go_fuzz_dep_.CoverTab[69011]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:149
			// _ = "end of CoverTab[69011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:149
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:149
		// _ = "end of CoverTab[69009]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:150
	// _ = "end of CoverTab[69002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:150
	_go_fuzz_dep_.CoverTab[69003]++

														d.wg.Add(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:152
	_curRoutineNum53_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:152
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum53_)
														go d.watcher()
														return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:154
	// _ = "end of CoverTab[69003]"
}

// Scheme returns the naming scheme of this resolver builder, which is "dns".
func (b *dnsBuilder) Scheme() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:158
	_go_fuzz_dep_.CoverTab[69012]++
														return "dns"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:159
	// _ = "end of CoverTab[69012]"
}

type netResolver interface {
	LookupHost(ctx context.Context, host string) (addrs []string, err error)
	LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*net.SRV, err error)
	LookupTXT(ctx context.Context, name string) (txts []string, err error)
}

// deadResolver is a resolver that does nothing.
type deadResolver struct{}

func (deadResolver) ResolveNow(resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:171
	_go_fuzz_dep_.CoverTab[69013]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:171
	// _ = "end of CoverTab[69013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:171
}

func (deadResolver) Close()	{ _go_fuzz_dep_.CoverTab[69014]++; // _ = "end of CoverTab[69014]" }

// dnsResolver watches for the name resolution update for a non-IP target.
type dnsResolver struct {
	host		string
	port		string
	resolver	netResolver
	ctx		context.Context
	cancel		context.CancelFunc
	cc		resolver.ClientConn
	// rn channel is used by ResolveNow() to force an immediate resolution of the target.
	rn	chan struct{}
	// wg is used to enforce Close() to return after the watcher() goroutine has finished.
	// Otherwise, data race will be possible. [Race Example] in dns_resolver_test we
	// replace the real lookup functions with mocked ones to facilitate testing.
	// If Close() doesn't wait for watcher() goroutine finishes, race detector sometimes
	// will warns lookup (READ the lookup function pointers) inside watcher() goroutine
	// has data race with replaceNetFunc (WRITE the lookup function pointers).
	wg			sync.WaitGroup
	disableServiceConfig	bool
}

// ResolveNow invoke an immediate resolution of the target that this dnsResolver watches.
func (d *dnsResolver) ResolveNow(resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:196
	_go_fuzz_dep_.CoverTab[69015]++
														select {
	case d.rn <- struct{}{}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:198
		_go_fuzz_dep_.CoverTab[69016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:198
		// _ = "end of CoverTab[69016]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:199
		_go_fuzz_dep_.CoverTab[69017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:199
		// _ = "end of CoverTab[69017]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:200
	// _ = "end of CoverTab[69015]"
}

// Close closes the dnsResolver.
func (d *dnsResolver) Close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:204
	_go_fuzz_dep_.CoverTab[69018]++
														d.cancel()
														d.wg.Wait()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:206
	// _ = "end of CoverTab[69018]"
}

func (d *dnsResolver) watcher() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:209
	_go_fuzz_dep_.CoverTab[69019]++
														defer d.wg.Done()
														backoffIndex := 1
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:212
		_go_fuzz_dep_.CoverTab[69020]++
															state, err := d.lookup()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:214
			_go_fuzz_dep_.CoverTab[69023]++

																d.cc.ReportError(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:216
			// _ = "end of CoverTab[69023]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:217
			_go_fuzz_dep_.CoverTab[69024]++
																err = d.cc.UpdateState(*state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:218
			// _ = "end of CoverTab[69024]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:219
		// _ = "end of CoverTab[69020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:219
		_go_fuzz_dep_.CoverTab[69021]++

															var timer *time.Timer
															if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:222
			_go_fuzz_dep_.CoverTab[69025]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:225
			backoffIndex = 1
			timer = newTimerDNSResRate(minDNSResRate)
			select {
			case <-d.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:228
				_go_fuzz_dep_.CoverTab[69026]++
																	timer.Stop()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:230
				// _ = "end of CoverTab[69026]"
			case <-d.rn:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:231
				_go_fuzz_dep_.CoverTab[69027]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:231
				// _ = "end of CoverTab[69027]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:232
			// _ = "end of CoverTab[69025]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:233
			_go_fuzz_dep_.CoverTab[69028]++

																timer = newTimer(backoff.DefaultExponential.Backoff(backoffIndex))
																backoffIndex++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:236
			// _ = "end of CoverTab[69028]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:237
		// _ = "end of CoverTab[69021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:237
		_go_fuzz_dep_.CoverTab[69022]++
															select {
		case <-d.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:239
			_go_fuzz_dep_.CoverTab[69029]++
																timer.Stop()
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:241
			// _ = "end of CoverTab[69029]"
		case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:242
			_go_fuzz_dep_.CoverTab[69030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:242
			// _ = "end of CoverTab[69030]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:243
		// _ = "end of CoverTab[69022]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:244
	// _ = "end of CoverTab[69019]"
}

func (d *dnsResolver) lookupSRV() ([]resolver.Address, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:247
	_go_fuzz_dep_.CoverTab[69031]++
														if !EnableSRVLookups {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:248
		_go_fuzz_dep_.CoverTab[69035]++
															return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:249
		// _ = "end of CoverTab[69035]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:250
		_go_fuzz_dep_.CoverTab[69036]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:250
		// _ = "end of CoverTab[69036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:250
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:250
	// _ = "end of CoverTab[69031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:250
	_go_fuzz_dep_.CoverTab[69032]++
														var newAddrs []resolver.Address
														_, srvs, err := d.resolver.LookupSRV(d.ctx, "grpclb", "tcp", d.host)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:253
		_go_fuzz_dep_.CoverTab[69037]++
															err = handleDNSError(err, "SRV")
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:255
		// _ = "end of CoverTab[69037]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:256
		_go_fuzz_dep_.CoverTab[69038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:256
		// _ = "end of CoverTab[69038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:256
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:256
	// _ = "end of CoverTab[69032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:256
	_go_fuzz_dep_.CoverTab[69033]++
														for _, s := range srvs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:257
		_go_fuzz_dep_.CoverTab[69039]++
															lbAddrs, err := d.resolver.LookupHost(d.ctx, s.Target)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:259
			_go_fuzz_dep_.CoverTab[69041]++
																err = handleDNSError(err, "A")
																if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:261
				_go_fuzz_dep_.CoverTab[69043]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:264
				continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:264
				// _ = "end of CoverTab[69043]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:265
				_go_fuzz_dep_.CoverTab[69044]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:265
				// _ = "end of CoverTab[69044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:265
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:265
			// _ = "end of CoverTab[69041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:265
			_go_fuzz_dep_.CoverTab[69042]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:266
			// _ = "end of CoverTab[69042]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:267
			_go_fuzz_dep_.CoverTab[69045]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:267
			// _ = "end of CoverTab[69045]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:267
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:267
		// _ = "end of CoverTab[69039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:267
		_go_fuzz_dep_.CoverTab[69040]++
															for _, a := range lbAddrs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:268
			_go_fuzz_dep_.CoverTab[69046]++
																ip, ok := formatIP(a)
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:270
				_go_fuzz_dep_.CoverTab[69048]++
																	return nil, fmt.Errorf("dns: error parsing A record IP address %v", a)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:271
				// _ = "end of CoverTab[69048]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:272
				_go_fuzz_dep_.CoverTab[69049]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:272
				// _ = "end of CoverTab[69049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:272
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:272
			// _ = "end of CoverTab[69046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:272
			_go_fuzz_dep_.CoverTab[69047]++
																addr := ip + ":" + strconv.Itoa(int(s.Port))
																newAddrs = append(newAddrs, resolver.Address{Addr: addr, ServerName: s.Target})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:274
			// _ = "end of CoverTab[69047]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:275
		// _ = "end of CoverTab[69040]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:276
	// _ = "end of CoverTab[69033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:276
	_go_fuzz_dep_.CoverTab[69034]++
														return newAddrs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:277
	// _ = "end of CoverTab[69034]"
}

func handleDNSError(err error, lookupType string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:280
	_go_fuzz_dep_.CoverTab[69050]++
														if dnsErr, ok := err.(*net.DNSError); ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		_go_fuzz_dep_.CoverTab[69053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		return !dnsErr.IsTimeout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		// _ = "end of CoverTab[69053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		_go_fuzz_dep_.CoverTab[69054]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		return !dnsErr.IsTemporary
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		// _ = "end of CoverTab[69054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:281
		_go_fuzz_dep_.CoverTab[69055]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:285
		return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:285
		// _ = "end of CoverTab[69055]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:286
		_go_fuzz_dep_.CoverTab[69056]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:286
		// _ = "end of CoverTab[69056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:286
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:286
	// _ = "end of CoverTab[69050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:286
	_go_fuzz_dep_.CoverTab[69051]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:287
		_go_fuzz_dep_.CoverTab[69057]++
															err = fmt.Errorf("dns: %v record lookup error: %v", lookupType, err)
															logger.Info(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:289
		// _ = "end of CoverTab[69057]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:290
		_go_fuzz_dep_.CoverTab[69058]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:290
		// _ = "end of CoverTab[69058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:290
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:290
	// _ = "end of CoverTab[69051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:290
	_go_fuzz_dep_.CoverTab[69052]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:291
	// _ = "end of CoverTab[69052]"
}

func (d *dnsResolver) lookupTXT() *serviceconfig.ParseResult {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:294
	_go_fuzz_dep_.CoverTab[69059]++
														ss, err := d.resolver.LookupTXT(d.ctx, txtPrefix+d.host)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:296
		_go_fuzz_dep_.CoverTab[69063]++
															if envconfig.TXTErrIgnore {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:297
			_go_fuzz_dep_.CoverTab[69066]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:298
			// _ = "end of CoverTab[69066]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:299
			_go_fuzz_dep_.CoverTab[69067]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:299
			// _ = "end of CoverTab[69067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:299
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:299
		// _ = "end of CoverTab[69063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:299
		_go_fuzz_dep_.CoverTab[69064]++
															if err = handleDNSError(err, "TXT"); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:300
			_go_fuzz_dep_.CoverTab[69068]++
																return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:301
			// _ = "end of CoverTab[69068]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:302
			_go_fuzz_dep_.CoverTab[69069]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:302
			// _ = "end of CoverTab[69069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:302
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:302
		// _ = "end of CoverTab[69064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:302
		_go_fuzz_dep_.CoverTab[69065]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:303
		// _ = "end of CoverTab[69065]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:304
		_go_fuzz_dep_.CoverTab[69070]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:304
		// _ = "end of CoverTab[69070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:304
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:304
	// _ = "end of CoverTab[69059]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:304
	_go_fuzz_dep_.CoverTab[69060]++
														var res string
														for _, s := range ss {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:306
		_go_fuzz_dep_.CoverTab[69071]++
															res += s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:307
		// _ = "end of CoverTab[69071]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:308
	// _ = "end of CoverTab[69060]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:308
	_go_fuzz_dep_.CoverTab[69061]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:311
	if !strings.HasPrefix(res, txtAttribute) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:311
		_go_fuzz_dep_.CoverTab[69072]++
															logger.Warningf("dns: TXT record %v missing %v attribute", res, txtAttribute)

															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:314
		// _ = "end of CoverTab[69072]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:315
		_go_fuzz_dep_.CoverTab[69073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:315
		// _ = "end of CoverTab[69073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:315
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:315
	// _ = "end of CoverTab[69061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:315
	_go_fuzz_dep_.CoverTab[69062]++
														sc := canaryingSC(strings.TrimPrefix(res, txtAttribute))
														return d.cc.ParseServiceConfig(sc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:317
	// _ = "end of CoverTab[69062]"
}

func (d *dnsResolver) lookupHost() ([]resolver.Address, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:320
	_go_fuzz_dep_.CoverTab[69074]++
														addrs, err := d.resolver.LookupHost(d.ctx, d.host)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:322
		_go_fuzz_dep_.CoverTab[69077]++
															err = handleDNSError(err, "A")
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:324
		// _ = "end of CoverTab[69077]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:325
		_go_fuzz_dep_.CoverTab[69078]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:325
		// _ = "end of CoverTab[69078]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:325
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:325
	// _ = "end of CoverTab[69074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:325
	_go_fuzz_dep_.CoverTab[69075]++
														newAddrs := make([]resolver.Address, 0, len(addrs))
														for _, a := range addrs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:327
		_go_fuzz_dep_.CoverTab[69079]++
															ip, ok := formatIP(a)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:329
			_go_fuzz_dep_.CoverTab[69081]++
																return nil, fmt.Errorf("dns: error parsing A record IP address %v", a)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:330
			// _ = "end of CoverTab[69081]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:331
			_go_fuzz_dep_.CoverTab[69082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:331
			// _ = "end of CoverTab[69082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:331
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:331
		// _ = "end of CoverTab[69079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:331
		_go_fuzz_dep_.CoverTab[69080]++
															addr := ip + ":" + d.port
															newAddrs = append(newAddrs, resolver.Address{Addr: addr})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:333
		// _ = "end of CoverTab[69080]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:334
	// _ = "end of CoverTab[69075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:334
	_go_fuzz_dep_.CoverTab[69076]++
														return newAddrs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:335
	// _ = "end of CoverTab[69076]"
}

func (d *dnsResolver) lookup() (*resolver.State, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:338
	_go_fuzz_dep_.CoverTab[69083]++
														srv, srvErr := d.lookupSRV()
														addrs, hostErr := d.lookupHost()
														if hostErr != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
		_go_fuzz_dep_.CoverTab[69087]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
		return (srvErr != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
			_go_fuzz_dep_.CoverTab[69088]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
			return len(srv) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
			// _ = "end of CoverTab[69088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
		// _ = "end of CoverTab[69087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:341
		_go_fuzz_dep_.CoverTab[69089]++
															return nil, hostErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:342
		// _ = "end of CoverTab[69089]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:343
		_go_fuzz_dep_.CoverTab[69090]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:343
		// _ = "end of CoverTab[69090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:343
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:343
	// _ = "end of CoverTab[69083]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:343
	_go_fuzz_dep_.CoverTab[69084]++

														state := resolver.State{Addresses: addrs}
														if len(srv) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:346
		_go_fuzz_dep_.CoverTab[69091]++
															state = grpclbstate.Set(state, &grpclbstate.State{BalancerAddresses: srv})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:347
		// _ = "end of CoverTab[69091]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:348
		_go_fuzz_dep_.CoverTab[69092]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:348
		// _ = "end of CoverTab[69092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:348
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:348
	// _ = "end of CoverTab[69084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:348
	_go_fuzz_dep_.CoverTab[69085]++
														if !d.disableServiceConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:349
		_go_fuzz_dep_.CoverTab[69093]++
															state.ServiceConfig = d.lookupTXT()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:350
		// _ = "end of CoverTab[69093]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:351
		_go_fuzz_dep_.CoverTab[69094]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:351
		// _ = "end of CoverTab[69094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:351
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:351
	// _ = "end of CoverTab[69085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:351
	_go_fuzz_dep_.CoverTab[69086]++
														return &state, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:352
	// _ = "end of CoverTab[69086]"
}

// formatIP returns ok = false if addr is not a valid textual representation of an IP address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:355
// If addr is an IPv4 address, return the addr and ok = true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:355
// If addr is an IPv6 address, return the addr enclosed in square brackets and ok = true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:358
func formatIP(addr string) (addrIP string, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:358
	_go_fuzz_dep_.CoverTab[69095]++
														ip := net.ParseIP(addr)
														if ip == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:360
		_go_fuzz_dep_.CoverTab[69098]++
															return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:361
		// _ = "end of CoverTab[69098]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:362
		_go_fuzz_dep_.CoverTab[69099]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:362
		// _ = "end of CoverTab[69099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:362
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:362
	// _ = "end of CoverTab[69095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:362
	_go_fuzz_dep_.CoverTab[69096]++
														if ip.To4() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:363
		_go_fuzz_dep_.CoverTab[69100]++
															return addr, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:364
		// _ = "end of CoverTab[69100]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:365
		_go_fuzz_dep_.CoverTab[69101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:365
		// _ = "end of CoverTab[69101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:365
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:365
	// _ = "end of CoverTab[69096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:365
	_go_fuzz_dep_.CoverTab[69097]++
														return "[" + addr + "]", true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:366
	// _ = "end of CoverTab[69097]"
}

// parseTarget takes the user input target string and default port, returns formatted host and port info.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// If target doesn't specify a port, set the port to be the defaultPort.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// If target is in IPv6 format and host-name is enclosed in square brackets, brackets
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// are stripped when setting the host.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// examples:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// target: "www.google.com" defaultPort: "443" returns host: "www.google.com", port: "443"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// target: "ipv4-host:80" defaultPort: "443" returns host: "ipv4-host", port: "80"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// target: "[ipv6-host]" defaultPort: "443" returns host: "ipv6-host", port: "443"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:369
// target: ":80" defaultPort: "443" returns host: "localhost", port: "80"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:378
func parseTarget(target, defaultPort string) (host, port string, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:378
	_go_fuzz_dep_.CoverTab[69102]++
														if target == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:379
		_go_fuzz_dep_.CoverTab[69107]++
															return "", "", errMissingAddr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:380
		// _ = "end of CoverTab[69107]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:381
		_go_fuzz_dep_.CoverTab[69108]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:381
		// _ = "end of CoverTab[69108]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:381
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:381
	// _ = "end of CoverTab[69102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:381
	_go_fuzz_dep_.CoverTab[69103]++
														if ip := net.ParseIP(target); ip != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:382
		_go_fuzz_dep_.CoverTab[69109]++

															return target, defaultPort, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:384
		// _ = "end of CoverTab[69109]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:385
		_go_fuzz_dep_.CoverTab[69110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:385
		// _ = "end of CoverTab[69110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:385
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:385
	// _ = "end of CoverTab[69103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:385
	_go_fuzz_dep_.CoverTab[69104]++
														if host, port, err = net.SplitHostPort(target); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:386
		_go_fuzz_dep_.CoverTab[69111]++
															if port == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:387
			_go_fuzz_dep_.CoverTab[69114]++

																return "", "", errEndsWithColon
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:389
			// _ = "end of CoverTab[69114]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:390
			_go_fuzz_dep_.CoverTab[69115]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:390
			// _ = "end of CoverTab[69115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:390
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:390
		// _ = "end of CoverTab[69111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:390
		_go_fuzz_dep_.CoverTab[69112]++

															if host == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:392
			_go_fuzz_dep_.CoverTab[69116]++

																host = "localhost"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:394
			// _ = "end of CoverTab[69116]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:395
			_go_fuzz_dep_.CoverTab[69117]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:395
			// _ = "end of CoverTab[69117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:395
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:395
		// _ = "end of CoverTab[69112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:395
		_go_fuzz_dep_.CoverTab[69113]++
															return host, port, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:396
		// _ = "end of CoverTab[69113]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:397
		_go_fuzz_dep_.CoverTab[69118]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:397
		// _ = "end of CoverTab[69118]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:397
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:397
	// _ = "end of CoverTab[69104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:397
	_go_fuzz_dep_.CoverTab[69105]++
														if host, port, err = net.SplitHostPort(target + ":" + defaultPort); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:398
		_go_fuzz_dep_.CoverTab[69119]++

															return host, port, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:400
		// _ = "end of CoverTab[69119]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:401
		_go_fuzz_dep_.CoverTab[69120]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:401
		// _ = "end of CoverTab[69120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:401
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:401
	// _ = "end of CoverTab[69105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:401
	_go_fuzz_dep_.CoverTab[69106]++
														return "", "", fmt.Errorf("invalid target address %v, error info: %v", target, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:402
	// _ = "end of CoverTab[69106]"
}

type rawChoice struct {
	ClientLanguage	*[]string		`json:"clientLanguage,omitempty"`
	Percentage	*int			`json:"percentage,omitempty"`
	ClientHostName	*[]string		`json:"clientHostName,omitempty"`
	ServiceConfig	*json.RawMessage	`json:"serviceConfig,omitempty"`
}

func containsString(a *[]string, b string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:412
	_go_fuzz_dep_.CoverTab[69121]++
														if a == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:413
		_go_fuzz_dep_.CoverTab[69124]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:414
		// _ = "end of CoverTab[69124]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:415
		_go_fuzz_dep_.CoverTab[69125]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:415
		// _ = "end of CoverTab[69125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:415
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:415
	// _ = "end of CoverTab[69121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:415
	_go_fuzz_dep_.CoverTab[69122]++
														for _, c := range *a {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:416
		_go_fuzz_dep_.CoverTab[69126]++
															if c == b {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:417
			_go_fuzz_dep_.CoverTab[69127]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:418
			// _ = "end of CoverTab[69127]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:419
			_go_fuzz_dep_.CoverTab[69128]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:419
			// _ = "end of CoverTab[69128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:419
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:419
		// _ = "end of CoverTab[69126]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:420
	// _ = "end of CoverTab[69122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:420
	_go_fuzz_dep_.CoverTab[69123]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:421
	// _ = "end of CoverTab[69123]"
}

func chosenByPercentage(a *int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:424
	_go_fuzz_dep_.CoverTab[69129]++
														if a == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:425
		_go_fuzz_dep_.CoverTab[69131]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:426
		// _ = "end of CoverTab[69131]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:427
		_go_fuzz_dep_.CoverTab[69132]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:427
		// _ = "end of CoverTab[69132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:427
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:427
	// _ = "end of CoverTab[69129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:427
	_go_fuzz_dep_.CoverTab[69130]++
														return grpcrand.Intn(100)+1 <= *a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:428
	// _ = "end of CoverTab[69130]"
}

func canaryingSC(js string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:431
	_go_fuzz_dep_.CoverTab[69133]++
														if js == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:432
		_go_fuzz_dep_.CoverTab[69138]++
															return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:433
		// _ = "end of CoverTab[69138]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:434
		_go_fuzz_dep_.CoverTab[69139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:434
		// _ = "end of CoverTab[69139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:434
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:434
	// _ = "end of CoverTab[69133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:434
	_go_fuzz_dep_.CoverTab[69134]++
														var rcs []rawChoice
														err := json.Unmarshal([]byte(js), &rcs)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:437
		_go_fuzz_dep_.CoverTab[69140]++
															logger.Warningf("dns: error parsing service config json: %v", err)
															return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:439
		// _ = "end of CoverTab[69140]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:440
		_go_fuzz_dep_.CoverTab[69141]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:440
		// _ = "end of CoverTab[69141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:440
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:440
	// _ = "end of CoverTab[69134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:440
	_go_fuzz_dep_.CoverTab[69135]++
														cliHostname, err := os.Hostname()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:442
		_go_fuzz_dep_.CoverTab[69142]++
															logger.Warningf("dns: error getting client hostname: %v", err)
															return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:444
		// _ = "end of CoverTab[69142]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:445
		_go_fuzz_dep_.CoverTab[69143]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:445
		// _ = "end of CoverTab[69143]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:445
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:445
	// _ = "end of CoverTab[69135]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:445
	_go_fuzz_dep_.CoverTab[69136]++
														var sc string
														for _, c := range rcs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:447
		_go_fuzz_dep_.CoverTab[69144]++
															if !containsString(c.ClientLanguage, golang) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:448
			_go_fuzz_dep_.CoverTab[69146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:448
			return !chosenByPercentage(c.Percentage)
																// _ = "end of CoverTab[69146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:449
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:449
			_go_fuzz_dep_.CoverTab[69147]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:449
			return !containsString(c.ClientHostName, cliHostname)
																// _ = "end of CoverTab[69147]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:450
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:450
			_go_fuzz_dep_.CoverTab[69148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:450
			return c.ServiceConfig == nil
																// _ = "end of CoverTab[69148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:451
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:451
			_go_fuzz_dep_.CoverTab[69149]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:452
			// _ = "end of CoverTab[69149]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:453
			_go_fuzz_dep_.CoverTab[69150]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:453
			// _ = "end of CoverTab[69150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:453
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:453
		// _ = "end of CoverTab[69144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:453
		_go_fuzz_dep_.CoverTab[69145]++
															sc = string(*c.ServiceConfig)
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:455
		// _ = "end of CoverTab[69145]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:456
	// _ = "end of CoverTab[69136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:456
	_go_fuzz_dep_.CoverTab[69137]++
														return sc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:457
	// _ = "end of CoverTab[69137]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:458
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/resolver/dns/dns_resolver.go:458
var _ = _go_fuzz_dep_.CoverTab
