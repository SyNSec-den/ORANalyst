//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:19
)

import (
	"strings"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/pretty"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// ccResolverWrapper is a wrapper on top of cc for resolvers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:34
// It implements resolver.ClientConn interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:36
type ccResolverWrapper struct {
	cc		*ClientConn
	resolverMu	sync.Mutex
	resolver	resolver.Resolver
	done		*grpcsync.Event
	curState	resolver.State

	incomingMu	sync.Mutex	// Synchronizes all the incoming calls.
}

// newCCResolverWrapper uses the resolver.Builder to build a Resolver and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:46
// returns a ccResolverWrapper object which wraps the newly built resolver.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:48
func newCCResolverWrapper(cc *ClientConn, rb resolver.Builder) (*ccResolverWrapper, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:48
	_go_fuzz_dep_.CoverTab[79610]++
												ccr := &ccResolverWrapper{
		cc:	cc,
		done:	grpcsync.NewEvent(),
	}

	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:55
		_go_fuzz_dep_.CoverTab[79613]++
													credsClone = creds.Clone()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:56
		// _ = "end of CoverTab[79613]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:57
		_go_fuzz_dep_.CoverTab[79614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:57
		// _ = "end of CoverTab[79614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:57
	// _ = "end of CoverTab[79610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:57
	_go_fuzz_dep_.CoverTab[79611]++
												rbo := resolver.BuildOptions{
		DisableServiceConfig:	cc.dopts.disableServiceConfig,
		DialCreds:		credsClone,
		CredsBundle:		cc.dopts.copts.CredsBundle,
		Dialer:			cc.dopts.copts.Dialer,
	}

												var err error

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:70
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:73
		_go_fuzz_dep_.CoverTab[79615]++
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:74
		// _ = "end of CoverTab[79615]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:75
		_go_fuzz_dep_.CoverTab[79616]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:75
		// _ = "end of CoverTab[79616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:75
	// _ = "end of CoverTab[79611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:75
	_go_fuzz_dep_.CoverTab[79612]++
												return ccr, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:76
	// _ = "end of CoverTab[79612]"
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:79
	_go_fuzz_dep_.CoverTab[79617]++
												ccr.resolverMu.Lock()
												if !ccr.done.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:81
		_go_fuzz_dep_.CoverTab[79619]++
													ccr.resolver.ResolveNow(o)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:82
		// _ = "end of CoverTab[79619]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:83
		_go_fuzz_dep_.CoverTab[79620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:83
		// _ = "end of CoverTab[79620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:83
	// _ = "end of CoverTab[79617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:83
	_go_fuzz_dep_.CoverTab[79618]++
												ccr.resolverMu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:84
	// _ = "end of CoverTab[79618]"
}

func (ccr *ccResolverWrapper) close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:87
	_go_fuzz_dep_.CoverTab[79621]++
												ccr.resolverMu.Lock()
												ccr.resolver.Close()
												ccr.done.Fire()
												ccr.resolverMu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:91
	// _ = "end of CoverTab[79621]"
}

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:94
	_go_fuzz_dep_.CoverTab[79622]++
												ccr.incomingMu.Lock()
												defer ccr.incomingMu.Unlock()
												if ccr.done.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:97
		_go_fuzz_dep_.CoverTab[79625]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:98
		// _ = "end of CoverTab[79625]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:99
		_go_fuzz_dep_.CoverTab[79626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:99
		// _ = "end of CoverTab[79626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:99
	// _ = "end of CoverTab[79622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:99
	_go_fuzz_dep_.CoverTab[79623]++
												ccr.addChannelzTraceEvent(s)
												ccr.curState = s
												if err := ccr.cc.updateResolverState(ccr.curState, nil); err == balancer.ErrBadResolverState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:102
		_go_fuzz_dep_.CoverTab[79627]++
													return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:103
		// _ = "end of CoverTab[79627]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:104
		_go_fuzz_dep_.CoverTab[79628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:104
		// _ = "end of CoverTab[79628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:104
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:104
	// _ = "end of CoverTab[79623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:104
	_go_fuzz_dep_.CoverTab[79624]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:105
	// _ = "end of CoverTab[79624]"
}

func (ccr *ccResolverWrapper) ReportError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:108
	_go_fuzz_dep_.CoverTab[79629]++
												ccr.incomingMu.Lock()
												defer ccr.incomingMu.Unlock()
												if ccr.done.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:111
		_go_fuzz_dep_.CoverTab[79631]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:112
		// _ = "end of CoverTab[79631]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:113
		_go_fuzz_dep_.CoverTab[79632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:113
		// _ = "end of CoverTab[79632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:113
	// _ = "end of CoverTab[79629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:113
	_go_fuzz_dep_.CoverTab[79630]++
												channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: reporting error to cc: %v", err)
												ccr.cc.updateResolverState(resolver.State{}, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:115
	// _ = "end of CoverTab[79630]"
}

// NewAddress is called by the resolver implementation to send addresses to gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:119
	_go_fuzz_dep_.CoverTab[79633]++
												ccr.incomingMu.Lock()
												defer ccr.incomingMu.Unlock()
												if ccr.done.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:122
		_go_fuzz_dep_.CoverTab[79635]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:123
		// _ = "end of CoverTab[79635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:124
		_go_fuzz_dep_.CoverTab[79636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:124
		// _ = "end of CoverTab[79636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:124
	// _ = "end of CoverTab[79633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:124
	_go_fuzz_dep_.CoverTab[79634]++
												ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
												ccr.curState.Addresses = addrs
												ccr.cc.updateResolverState(ccr.curState, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:127
	// _ = "end of CoverTab[79634]"
}

// NewServiceConfig is called by the resolver implementation to send service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:130
// configs to gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:132
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:132
	_go_fuzz_dep_.CoverTab[79637]++
												ccr.incomingMu.Lock()
												defer ccr.incomingMu.Unlock()
												if ccr.done.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:135
		_go_fuzz_dep_.CoverTab[79641]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:136
		// _ = "end of CoverTab[79641]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:137
		_go_fuzz_dep_.CoverTab[79642]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:137
		// _ = "end of CoverTab[79642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:137
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:137
	// _ = "end of CoverTab[79637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:137
	_go_fuzz_dep_.CoverTab[79638]++
												channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: got new service config: %s", sc)
												if ccr.cc.dopts.disableServiceConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:139
		_go_fuzz_dep_.CoverTab[79643]++
													channelz.Info(logger, ccr.cc.channelzID, "Service config lookups disabled; ignoring config")
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:141
		// _ = "end of CoverTab[79643]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:142
		_go_fuzz_dep_.CoverTab[79644]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:142
		// _ = "end of CoverTab[79644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:142
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:142
	// _ = "end of CoverTab[79638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:142
	_go_fuzz_dep_.CoverTab[79639]++
												scpr := parseServiceConfig(sc)
												if scpr.Err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:144
		_go_fuzz_dep_.CoverTab[79645]++
													channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: error parsing service config: %v", scpr.Err)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:146
		// _ = "end of CoverTab[79645]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:147
		_go_fuzz_dep_.CoverTab[79646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:147
		// _ = "end of CoverTab[79646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:147
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:147
	// _ = "end of CoverTab[79639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:147
	_go_fuzz_dep_.CoverTab[79640]++
												ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
												ccr.curState.ServiceConfig = scpr
												ccr.cc.updateResolverState(ccr.curState, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:150
	// _ = "end of CoverTab[79640]"
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:153
	_go_fuzz_dep_.CoverTab[79647]++
												return parseServiceConfig(scJSON)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:154
	// _ = "end of CoverTab[79647]"
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:157
	_go_fuzz_dep_.CoverTab[79648]++
												var updates []string
												var oldSC, newSC *ServiceConfig
												var oldOK, newOK bool
												if ccr.curState.ServiceConfig != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:161
		_go_fuzz_dep_.CoverTab[79653]++
													oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:162
		// _ = "end of CoverTab[79653]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:163
		_go_fuzz_dep_.CoverTab[79654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:163
		// _ = "end of CoverTab[79654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:163
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:163
	// _ = "end of CoverTab[79648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:163
	_go_fuzz_dep_.CoverTab[79649]++
												if s.ServiceConfig != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:164
		_go_fuzz_dep_.CoverTab[79655]++
													newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:165
		// _ = "end of CoverTab[79655]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:166
		_go_fuzz_dep_.CoverTab[79656]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:166
		// _ = "end of CoverTab[79656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:166
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:166
	// _ = "end of CoverTab[79649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:166
	_go_fuzz_dep_.CoverTab[79650]++
												if oldOK != newOK || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		_go_fuzz_dep_.CoverTab[79657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		return (oldOK && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			_go_fuzz_dep_.CoverTab[79658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			return newOK
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			// _ = "end of CoverTab[79658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			_go_fuzz_dep_.CoverTab[79659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			return oldSC.rawJSONString != newSC.rawJSONString
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
			// _ = "end of CoverTab[79659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		// _ = "end of CoverTab[79657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:167
		_go_fuzz_dep_.CoverTab[79660]++
													updates = append(updates, "service config updated")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:168
		// _ = "end of CoverTab[79660]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:169
		_go_fuzz_dep_.CoverTab[79661]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:169
		// _ = "end of CoverTab[79661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:169
	// _ = "end of CoverTab[79650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:169
	_go_fuzz_dep_.CoverTab[79651]++
												if len(ccr.curState.Addresses) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:170
		_go_fuzz_dep_.CoverTab[79662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:170
		return len(s.Addresses) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:170
		// _ = "end of CoverTab[79662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:170
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:170
		_go_fuzz_dep_.CoverTab[79663]++
													updates = append(updates, "resolver returned an empty address list")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:171
		// _ = "end of CoverTab[79663]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
		_go_fuzz_dep_.CoverTab[79664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
		if len(ccr.curState.Addresses) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
			_go_fuzz_dep_.CoverTab[79665]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
			return len(s.Addresses) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
			// _ = "end of CoverTab[79665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:172
			_go_fuzz_dep_.CoverTab[79666]++
														updates = append(updates, "resolver returned new addresses")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:173
			// _ = "end of CoverTab[79666]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
			_go_fuzz_dep_.CoverTab[79667]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
			// _ = "end of CoverTab[79667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
		// _ = "end of CoverTab[79664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
	// _ = "end of CoverTab[79651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:174
	_go_fuzz_dep_.CoverTab[79652]++
												channelz.Infof(logger, ccr.cc.channelzID, "Resolver state updated: %s (%v)", pretty.ToJSON(s), strings.Join(updates, "; "))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:175
	// _ = "end of CoverTab[79652]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver_conn_wrapper.go:176
var _ = _go_fuzz_dep_.CoverTab
