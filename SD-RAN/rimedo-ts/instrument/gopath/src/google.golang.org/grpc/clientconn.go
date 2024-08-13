//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:19
)

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/backoff"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	iresolver "google.golang.org/grpc/internal/resolver"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	"google.golang.org/grpc/status"

	_ "google.golang.org/grpc/balancer/roundrobin"			// To register roundrobin.
	_ "google.golang.org/grpc/internal/resolver/dns"		// To register dns resolver.
	_ "google.golang.org/grpc/internal/resolver/passthrough"	// To register passthrough resolver.
	_ "google.golang.org/grpc/internal/resolver/unix"		// To register unix resolver.
)

const (
	// minimum time to give a connection to complete
	minConnectTimeout	= 20 * time.Second
	// must match grpclbName in grpclb/grpclb.go
	grpclbName	= "grpclb"
)

var (
	// ErrClientConnClosing indicates that the operation is illegal because
	// the ClientConn is closing.
	//
	// Deprecated: this error should not be relied upon by users; use the status
	// code of Canceled instead.
	ErrClientConnClosing	= status.Error(codes.Canceled, "grpc: the client connection is closing")
	// errConnDrain indicates that the connection starts to be drained and does not accept any new RPCs.
	errConnDrain	= errors.New("grpc: the connection is drained")
	// errConnClosing indicates that the connection is closing.
	errConnClosing	= errors.New("grpc: the connection is closing")
	// invalidDefaultServiceConfigErrPrefix is used to prefix the json parsing error for the default
	// service config.
	invalidDefaultServiceConfigErrPrefix	= "grpc: the provided default service config is invalid"
)

// The following errors are returned from Dial and DialContext
var (
	// errNoTransportSecurity indicates that there is no transport security
	// being set for ClientConn. Users should either set one or explicitly
	// call WithInsecure DialOption to disable security.
	errNoTransportSecurity	= errors.New("grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials)")
	// errTransportCredsAndBundle indicates that creds bundle is used together
	// with other individual Transport Credentials.
	errTransportCredsAndBundle	= errors.New("grpc: credentials.Bundle may not be used with individual TransportCredentials")
	// errNoTransportCredsInBundle indicated that the configured creds bundle
	// returned a transport credentials which was nil.
	errNoTransportCredsInBundle	= errors.New("grpc: credentials.Bundle must return non-nil transport credentials")
	// errTransportCredentialsMissing indicates that users want to transmit
	// security information (e.g., OAuth2 token) which requires secure
	// connection on an insecure connection.
	errTransportCredentialsMissing	= errors.New("grpc: the credentials require transport level security (use grpc.WithTransportCredentials() to set)")
)

const (
	defaultClientMaxReceiveMessageSize	= 1024 * 1024 * 4
	defaultClientMaxSendMessageSize		= math.MaxInt32
	// http2IOBufSize specifies the buffer size for sending frames.
	defaultWriteBufSize	= 32 * 1024
	defaultReadBufSize	= 32 * 1024
)

// Dial creates a client connection to the given target.
func Dial(target string, opts ...DialOption) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:104
	_go_fuzz_dep_.CoverTab[78855]++
											return DialContext(context.Background(), target, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:105
	// _ = "end of CoverTab[78855]"
}

type defaultConfigSelector struct {
	sc *ServiceConfig
}

func (dcs *defaultConfigSelector) SelectConfig(rpcInfo iresolver.RPCInfo) (*iresolver.RPCConfig, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:112
	_go_fuzz_dep_.CoverTab[78856]++
											return &iresolver.RPCConfig{
		Context:	rpcInfo.Context,
		MethodConfig:	getMethodConfig(dcs.sc, rpcInfo.Method),
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:116
	// _ = "end of CoverTab[78856]"
}

// DialContext creates a client connection to the given target. By default, it's
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// a non-blocking dial (the function won't wait for connections to be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// established, and connecting happens in the background). To make it a blocking
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// dial, use WithBlock() dial option.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// In the non-blocking case, the ctx does not act against the connection. It
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// only controls the setup steps.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// In the blocking case, ctx can be used to cancel or expire the pending
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// connection. Once this function returns, the cancellation and expiration of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// ctx will be noop. Users should call ClientConn.Close to terminate all the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// pending operations after this function returns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// The target name syntax is defined in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// https://github.com/grpc/grpc/blob/master/doc/naming.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:119
// e.g. to use dns resolver, a "dns:///" prefix should be applied to the target.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:135
func DialContext(ctx context.Context, target string, opts ...DialOption) (conn *ClientConn, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:135
	_go_fuzz_dep_.CoverTab[78857]++
											cc := &ClientConn{
		target:			target,
		csMgr:			&connectivityStateManager{},
		conns:			make(map[*addrConn]struct{}),
		dopts:			defaultDialOptions(),
		blockingpicker:		newPickerWrapper(),
		czData:			new(channelzData),
		firstResolveEvent:	grpcsync.NewEvent(),
	}
	cc.retryThrottler.Store((*retryThrottler)(nil))
	cc.safeConfigSelector.UpdateConfigSelector(&defaultConfigSelector{nil})
	cc.ctx, cc.cancel = context.WithCancel(context.Background())

	disableGlobalOpts := false
	for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:150
		_go_fuzz_dep_.CoverTab[78881]++
												if _, ok := opt.(*disableGlobalDialOptions); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:151
			_go_fuzz_dep_.CoverTab[78882]++
													disableGlobalOpts = true
													break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:153
			// _ = "end of CoverTab[78882]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:154
			_go_fuzz_dep_.CoverTab[78883]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:154
			// _ = "end of CoverTab[78883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:154
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:154
		// _ = "end of CoverTab[78881]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:155
	// _ = "end of CoverTab[78857]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:155
	_go_fuzz_dep_.CoverTab[78858]++

											if !disableGlobalOpts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:157
		_go_fuzz_dep_.CoverTab[78884]++
												for _, opt := range globalDialOptions {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:158
			_go_fuzz_dep_.CoverTab[78885]++
													opt.apply(&cc.dopts)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:159
			// _ = "end of CoverTab[78885]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:160
		// _ = "end of CoverTab[78884]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:161
		_go_fuzz_dep_.CoverTab[78886]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:161
		// _ = "end of CoverTab[78886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:161
	// _ = "end of CoverTab[78858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:161
	_go_fuzz_dep_.CoverTab[78859]++

											for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:163
		_go_fuzz_dep_.CoverTab[78887]++
												opt.apply(&cc.dopts)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:164
		// _ = "end of CoverTab[78887]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:165
	// _ = "end of CoverTab[78859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:165
	_go_fuzz_dep_.CoverTab[78860]++

											chainUnaryClientInterceptors(cc)
											chainStreamClientInterceptors(cc)

											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:170
		_go_fuzz_dep_.CoverTab[78888]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:171
			_go_fuzz_dep_.CoverTab[78889]++
													cc.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:172
			// _ = "end of CoverTab[78889]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:173
			_go_fuzz_dep_.CoverTab[78890]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:173
			// _ = "end of CoverTab[78890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:173
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:173
		// _ = "end of CoverTab[78888]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:174
	// _ = "end of CoverTab[78860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:174
	_go_fuzz_dep_.CoverTab[78861]++

											pid := cc.dopts.channelzParentID
											cc.channelzID = channelz.RegisterChannel(&channelzChannel{cc}, pid, target)
											ted := &channelz.TraceEventDesc{
		Desc:		"Channel created",
		Severity:	channelz.CtInfo,
	}
	if cc.dopts.channelzParentID != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:182
		_go_fuzz_dep_.CoverTab[78891]++
												ted.Parent = &channelz.TraceEventDesc{
			Desc:		fmt.Sprintf("Nested Channel(id:%d) created", cc.channelzID.Int()),
			Severity:	channelz.CtInfo,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:186
		// _ = "end of CoverTab[78891]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:187
		_go_fuzz_dep_.CoverTab[78892]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:187
		// _ = "end of CoverTab[78892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:187
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:187
	// _ = "end of CoverTab[78861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:187
	_go_fuzz_dep_.CoverTab[78862]++
											channelz.AddTraceEvent(logger, cc.channelzID, 1, ted)
											cc.csMgr.channelzID = cc.channelzID

											if cc.dopts.copts.TransportCredentials == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:191
		_go_fuzz_dep_.CoverTab[78893]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:191
		return cc.dopts.copts.CredsBundle == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:191
		// _ = "end of CoverTab[78893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:191
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:191
		_go_fuzz_dep_.CoverTab[78894]++
												return nil, errNoTransportSecurity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:192
		// _ = "end of CoverTab[78894]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:193
		_go_fuzz_dep_.CoverTab[78895]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:193
		// _ = "end of CoverTab[78895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:193
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:193
	// _ = "end of CoverTab[78862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:193
	_go_fuzz_dep_.CoverTab[78863]++
											if cc.dopts.copts.TransportCredentials != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:194
		_go_fuzz_dep_.CoverTab[78896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:194
		return cc.dopts.copts.CredsBundle != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:194
		// _ = "end of CoverTab[78896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:194
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:194
		_go_fuzz_dep_.CoverTab[78897]++
												return nil, errTransportCredsAndBundle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:195
		// _ = "end of CoverTab[78897]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:196
		_go_fuzz_dep_.CoverTab[78898]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:196
		// _ = "end of CoverTab[78898]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:196
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:196
	// _ = "end of CoverTab[78863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:196
	_go_fuzz_dep_.CoverTab[78864]++
											if cc.dopts.copts.CredsBundle != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:197
		_go_fuzz_dep_.CoverTab[78899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:197
		return cc.dopts.copts.CredsBundle.TransportCredentials() == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:197
		// _ = "end of CoverTab[78899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:197
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:197
		_go_fuzz_dep_.CoverTab[78900]++
												return nil, errNoTransportCredsInBundle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:198
		// _ = "end of CoverTab[78900]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:199
		_go_fuzz_dep_.CoverTab[78901]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:199
		// _ = "end of CoverTab[78901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:199
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:199
	// _ = "end of CoverTab[78864]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:199
	_go_fuzz_dep_.CoverTab[78865]++
											transportCreds := cc.dopts.copts.TransportCredentials
											if transportCreds == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:201
		_go_fuzz_dep_.CoverTab[78902]++
												transportCreds = cc.dopts.copts.CredsBundle.TransportCredentials()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:202
		// _ = "end of CoverTab[78902]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:203
		_go_fuzz_dep_.CoverTab[78903]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:203
		// _ = "end of CoverTab[78903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:203
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:203
	// _ = "end of CoverTab[78865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:203
	_go_fuzz_dep_.CoverTab[78866]++
											if transportCreds.Info().SecurityProtocol == "insecure" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:204
		_go_fuzz_dep_.CoverTab[78904]++
												for _, cd := range cc.dopts.copts.PerRPCCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:205
			_go_fuzz_dep_.CoverTab[78905]++
													if cd.RequireTransportSecurity() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:206
				_go_fuzz_dep_.CoverTab[78906]++
														return nil, errTransportCredentialsMissing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:207
				// _ = "end of CoverTab[78906]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:208
				_go_fuzz_dep_.CoverTab[78907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:208
				// _ = "end of CoverTab[78907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:208
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:208
			// _ = "end of CoverTab[78905]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:209
		// _ = "end of CoverTab[78904]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:210
		_go_fuzz_dep_.CoverTab[78908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:210
		// _ = "end of CoverTab[78908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:210
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:210
	// _ = "end of CoverTab[78866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:210
	_go_fuzz_dep_.CoverTab[78867]++

											if cc.dopts.defaultServiceConfigRawJSON != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:212
		_go_fuzz_dep_.CoverTab[78909]++
												scpr := parseServiceConfig(*cc.dopts.defaultServiceConfigRawJSON)
												if scpr.Err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:214
			_go_fuzz_dep_.CoverTab[78911]++
													return nil, fmt.Errorf("%s: %v", invalidDefaultServiceConfigErrPrefix, scpr.Err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:215
			// _ = "end of CoverTab[78911]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:216
			_go_fuzz_dep_.CoverTab[78912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:216
			// _ = "end of CoverTab[78912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:216
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:216
		// _ = "end of CoverTab[78909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:216
		_go_fuzz_dep_.CoverTab[78910]++
												cc.dopts.defaultServiceConfig, _ = scpr.Config.(*ServiceConfig)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:217
		// _ = "end of CoverTab[78910]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:218
		_go_fuzz_dep_.CoverTab[78913]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:218
		// _ = "end of CoverTab[78913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:218
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:218
	// _ = "end of CoverTab[78867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:218
	_go_fuzz_dep_.CoverTab[78868]++
											cc.mkp = cc.dopts.copts.KeepaliveParams

											if cc.dopts.copts.UserAgent != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:221
		_go_fuzz_dep_.CoverTab[78914]++
												cc.dopts.copts.UserAgent += " " + grpcUA
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:222
		// _ = "end of CoverTab[78914]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:223
		_go_fuzz_dep_.CoverTab[78915]++
												cc.dopts.copts.UserAgent = grpcUA
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:224
		// _ = "end of CoverTab[78915]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:225
	// _ = "end of CoverTab[78868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:225
	_go_fuzz_dep_.CoverTab[78869]++

											if cc.dopts.timeout > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:227
		_go_fuzz_dep_.CoverTab[78916]++
												var cancel context.CancelFunc
												ctx, cancel = context.WithTimeout(ctx, cc.dopts.timeout)
												defer cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:230
		// _ = "end of CoverTab[78916]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:231
		_go_fuzz_dep_.CoverTab[78917]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:231
		// _ = "end of CoverTab[78917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:231
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:231
	// _ = "end of CoverTab[78869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:231
	_go_fuzz_dep_.CoverTab[78870]++
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:232
		_go_fuzz_dep_.CoverTab[78918]++
												select {
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:234
			_go_fuzz_dep_.CoverTab[78919]++
													switch {
			case ctx.Err() == err:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:236
				_go_fuzz_dep_.CoverTab[78921]++
														conn = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:237
				// _ = "end of CoverTab[78921]"
			case err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:238
				_go_fuzz_dep_.CoverTab[78924]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:238
				return !cc.dopts.returnLastError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:238
				// _ = "end of CoverTab[78924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:238
			}():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:238
				_go_fuzz_dep_.CoverTab[78922]++
														conn, err = nil, ctx.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:239
				// _ = "end of CoverTab[78922]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:240
				_go_fuzz_dep_.CoverTab[78923]++
														conn, err = nil, fmt.Errorf("%v: %v", ctx.Err(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:241
				// _ = "end of CoverTab[78923]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:242
			// _ = "end of CoverTab[78919]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:243
			_go_fuzz_dep_.CoverTab[78920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:243
			// _ = "end of CoverTab[78920]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:244
		// _ = "end of CoverTab[78918]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:245
	// _ = "end of CoverTab[78870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:245
	_go_fuzz_dep_.CoverTab[78871]++

											scSet := false
											if cc.dopts.scChan != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:248
		_go_fuzz_dep_.CoverTab[78925]++

												select {
		case sc, ok := <-cc.dopts.scChan:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:251
			_go_fuzz_dep_.CoverTab[78926]++
													if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:252
				_go_fuzz_dep_.CoverTab[78928]++
														cc.sc = &sc
														cc.safeConfigSelector.UpdateConfigSelector(&defaultConfigSelector{&sc})
														scSet = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:255
				// _ = "end of CoverTab[78928]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:256
				_go_fuzz_dep_.CoverTab[78929]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:256
				// _ = "end of CoverTab[78929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:256
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:256
			// _ = "end of CoverTab[78926]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:257
			_go_fuzz_dep_.CoverTab[78927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:257
			// _ = "end of CoverTab[78927]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:258
		// _ = "end of CoverTab[78925]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:259
		_go_fuzz_dep_.CoverTab[78930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:259
		// _ = "end of CoverTab[78930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:259
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:259
	// _ = "end of CoverTab[78871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:259
	_go_fuzz_dep_.CoverTab[78872]++
											if cc.dopts.bs == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:260
		_go_fuzz_dep_.CoverTab[78931]++
												cc.dopts.bs = backoff.DefaultExponential
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:261
		// _ = "end of CoverTab[78931]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:262
		_go_fuzz_dep_.CoverTab[78932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:262
		// _ = "end of CoverTab[78932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:262
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:262
	// _ = "end of CoverTab[78872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:262
	_go_fuzz_dep_.CoverTab[78873]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:265
	resolverBuilder, err := cc.parseTargetAndFindResolver()
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:266
		_go_fuzz_dep_.CoverTab[78933]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:267
		// _ = "end of CoverTab[78933]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:268
		_go_fuzz_dep_.CoverTab[78934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:268
		// _ = "end of CoverTab[78934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:268
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:268
	// _ = "end of CoverTab[78873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:268
	_go_fuzz_dep_.CoverTab[78874]++
											cc.authority, err = determineAuthority(cc.parsedTarget.Endpoint(), cc.target, cc.dopts)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:270
		_go_fuzz_dep_.CoverTab[78935]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:271
		// _ = "end of CoverTab[78935]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:272
		_go_fuzz_dep_.CoverTab[78936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:272
		// _ = "end of CoverTab[78936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:272
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:272
	// _ = "end of CoverTab[78874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:272
	_go_fuzz_dep_.CoverTab[78875]++
											channelz.Infof(logger, cc.channelzID, "Channel authority set to %q", cc.authority)

											if cc.dopts.scChan != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:275
		_go_fuzz_dep_.CoverTab[78937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:275
		return !scSet
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:275
		// _ = "end of CoverTab[78937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:275
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:275
		_go_fuzz_dep_.CoverTab[78938]++

												select {
		case sc, ok := <-cc.dopts.scChan:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:278
			_go_fuzz_dep_.CoverTab[78939]++
													if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:279
				_go_fuzz_dep_.CoverTab[78941]++
														cc.sc = &sc
														cc.safeConfigSelector.UpdateConfigSelector(&defaultConfigSelector{&sc})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:281
				// _ = "end of CoverTab[78941]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:282
				_go_fuzz_dep_.CoverTab[78942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:282
				// _ = "end of CoverTab[78942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:282
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:282
			// _ = "end of CoverTab[78939]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:283
			_go_fuzz_dep_.CoverTab[78940]++
													return nil, ctx.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:284
			// _ = "end of CoverTab[78940]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:285
		// _ = "end of CoverTab[78938]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:286
		_go_fuzz_dep_.CoverTab[78943]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:286
		// _ = "end of CoverTab[78943]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:286
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:286
	// _ = "end of CoverTab[78875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:286
	_go_fuzz_dep_.CoverTab[78876]++
											if cc.dopts.scChan != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:287
		_go_fuzz_dep_.CoverTab[78944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:287
		_curRoutineNum88_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:287
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum88_)
												go cc.scWatcher()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:288
		// _ = "end of CoverTab[78944]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:289
		_go_fuzz_dep_.CoverTab[78945]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:289
		// _ = "end of CoverTab[78945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:289
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:289
	// _ = "end of CoverTab[78876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:289
	_go_fuzz_dep_.CoverTab[78877]++

											var credsClone credentials.TransportCredentials
											if creds := cc.dopts.copts.TransportCredentials; creds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:292
		_go_fuzz_dep_.CoverTab[78946]++
												credsClone = creds.Clone()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:293
		// _ = "end of CoverTab[78946]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:294
		_go_fuzz_dep_.CoverTab[78947]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:294
		// _ = "end of CoverTab[78947]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:294
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:294
	// _ = "end of CoverTab[78877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:294
	_go_fuzz_dep_.CoverTab[78878]++
											cc.balancerWrapper = newCCBalancerWrapper(cc, balancer.BuildOptions{
		DialCreds:		credsClone,
		CredsBundle:		cc.dopts.copts.CredsBundle,
		Dialer:			cc.dopts.copts.Dialer,
		Authority:		cc.authority,
		CustomUserAgent:	cc.dopts.copts.UserAgent,
		ChannelzParentID:	cc.channelzID,
		Target:			cc.parsedTarget,
	})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:306
	rWrapper, err := newCCResolverWrapper(cc, resolverBuilder)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:307
		_go_fuzz_dep_.CoverTab[78948]++
												return nil, fmt.Errorf("failed to build resolver: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:308
		// _ = "end of CoverTab[78948]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:309
		_go_fuzz_dep_.CoverTab[78949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:309
		// _ = "end of CoverTab[78949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:309
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:309
	// _ = "end of CoverTab[78878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:309
	_go_fuzz_dep_.CoverTab[78879]++
											cc.mu.Lock()
											cc.resolverWrapper = rWrapper
											cc.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:315
	if cc.dopts.block {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:315
		_go_fuzz_dep_.CoverTab[78950]++
												for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:316
			_go_fuzz_dep_.CoverTab[78951]++
													cc.Connect()
													s := cc.GetState()
													if s == connectivity.Ready {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:319
				_go_fuzz_dep_.CoverTab[78953]++
														break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:320
				// _ = "end of CoverTab[78953]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
				_go_fuzz_dep_.CoverTab[78954]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
				if cc.dopts.copts.FailOnNonTempDialError && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
					_go_fuzz_dep_.CoverTab[78955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
					return s == connectivity.TransientFailure
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
					// _ = "end of CoverTab[78955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:321
					_go_fuzz_dep_.CoverTab[78956]++
															if err = cc.connectionError(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:322
						_go_fuzz_dep_.CoverTab[78957]++
																terr, ok := err.(interface {
							Temporary() bool
						})
						if ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:326
							_go_fuzz_dep_.CoverTab[78958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:326
							return !terr.Temporary()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:326
							// _ = "end of CoverTab[78958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:326
						}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:326
							_go_fuzz_dep_.CoverTab[78959]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:327
							// _ = "end of CoverTab[78959]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:328
							_go_fuzz_dep_.CoverTab[78960]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:328
							// _ = "end of CoverTab[78960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:328
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:328
						// _ = "end of CoverTab[78957]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:329
						_go_fuzz_dep_.CoverTab[78961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:329
						// _ = "end of CoverTab[78961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:329
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:329
					// _ = "end of CoverTab[78956]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
					_go_fuzz_dep_.CoverTab[78962]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
					// _ = "end of CoverTab[78962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
				// _ = "end of CoverTab[78954]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
			// _ = "end of CoverTab[78951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:330
			_go_fuzz_dep_.CoverTab[78952]++
													if !cc.WaitForStateChange(ctx, s) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:331
				_go_fuzz_dep_.CoverTab[78963]++

														if err = cc.connectionError(); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:333
					_go_fuzz_dep_.CoverTab[78965]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:333
					return cc.dopts.returnLastError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:333
					// _ = "end of CoverTab[78965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:333
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:333
					_go_fuzz_dep_.CoverTab[78966]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:334
					// _ = "end of CoverTab[78966]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:335
					_go_fuzz_dep_.CoverTab[78967]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:335
					// _ = "end of CoverTab[78967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:335
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:335
				// _ = "end of CoverTab[78963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:335
				_go_fuzz_dep_.CoverTab[78964]++
														return nil, ctx.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:336
				// _ = "end of CoverTab[78964]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:337
				_go_fuzz_dep_.CoverTab[78968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:337
				// _ = "end of CoverTab[78968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:337
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:337
			// _ = "end of CoverTab[78952]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:338
		// _ = "end of CoverTab[78950]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:339
		_go_fuzz_dep_.CoverTab[78969]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:339
		// _ = "end of CoverTab[78969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:339
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:339
	// _ = "end of CoverTab[78879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:339
	_go_fuzz_dep_.CoverTab[78880]++

											return cc, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:341
	// _ = "end of CoverTab[78880]"
}

// chainUnaryClientInterceptors chains all unary client interceptors into one.
func chainUnaryClientInterceptors(cc *ClientConn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:345
	_go_fuzz_dep_.CoverTab[78970]++
											interceptors := cc.dopts.chainUnaryInts

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:349
	if cc.dopts.unaryInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:349
		_go_fuzz_dep_.CoverTab[78973]++
												interceptors = append([]UnaryClientInterceptor{cc.dopts.unaryInt}, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:350
		// _ = "end of CoverTab[78973]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:351
		_go_fuzz_dep_.CoverTab[78974]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:351
		// _ = "end of CoverTab[78974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:351
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:351
	// _ = "end of CoverTab[78970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:351
	_go_fuzz_dep_.CoverTab[78971]++
											var chainedInt UnaryClientInterceptor
											if len(interceptors) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:353
		_go_fuzz_dep_.CoverTab[78975]++
												chainedInt = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:354
		// _ = "end of CoverTab[78975]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:355
		_go_fuzz_dep_.CoverTab[78976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:355
		if len(interceptors) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:355
			_go_fuzz_dep_.CoverTab[78977]++
													chainedInt = interceptors[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:356
			// _ = "end of CoverTab[78977]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:357
			_go_fuzz_dep_.CoverTab[78978]++
													chainedInt = func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:358
				_go_fuzz_dep_.CoverTab[78979]++
														return interceptors[0](ctx, method, req, reply, cc, getChainUnaryInvoker(interceptors, 0, invoker), opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:359
				// _ = "end of CoverTab[78979]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:360
			// _ = "end of CoverTab[78978]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:361
		// _ = "end of CoverTab[78976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:361
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:361
	// _ = "end of CoverTab[78971]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:361
	_go_fuzz_dep_.CoverTab[78972]++
											cc.dopts.unaryInt = chainedInt
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:362
	// _ = "end of CoverTab[78972]"
}

// getChainUnaryInvoker recursively generate the chained unary invoker.
func getChainUnaryInvoker(interceptors []UnaryClientInterceptor, curr int, finalInvoker UnaryInvoker) UnaryInvoker {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:366
	_go_fuzz_dep_.CoverTab[78980]++
											if curr == len(interceptors)-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:367
		_go_fuzz_dep_.CoverTab[78982]++
												return finalInvoker
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:368
		// _ = "end of CoverTab[78982]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:369
		_go_fuzz_dep_.CoverTab[78983]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:369
		// _ = "end of CoverTab[78983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:369
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:369
	// _ = "end of CoverTab[78980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:369
	_go_fuzz_dep_.CoverTab[78981]++
											return func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:370
		_go_fuzz_dep_.CoverTab[78984]++
												return interceptors[curr+1](ctx, method, req, reply, cc, getChainUnaryInvoker(interceptors, curr+1, finalInvoker), opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:371
		// _ = "end of CoverTab[78984]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:372
	// _ = "end of CoverTab[78981]"
}

// chainStreamClientInterceptors chains all stream client interceptors into one.
func chainStreamClientInterceptors(cc *ClientConn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:376
	_go_fuzz_dep_.CoverTab[78985]++
											interceptors := cc.dopts.chainStreamInts

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:380
	if cc.dopts.streamInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:380
		_go_fuzz_dep_.CoverTab[78988]++
												interceptors = append([]StreamClientInterceptor{cc.dopts.streamInt}, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:381
		// _ = "end of CoverTab[78988]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:382
		_go_fuzz_dep_.CoverTab[78989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:382
		// _ = "end of CoverTab[78989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:382
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:382
	// _ = "end of CoverTab[78985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:382
	_go_fuzz_dep_.CoverTab[78986]++
											var chainedInt StreamClientInterceptor
											if len(interceptors) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:384
		_go_fuzz_dep_.CoverTab[78990]++
												chainedInt = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:385
		// _ = "end of CoverTab[78990]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:386
		_go_fuzz_dep_.CoverTab[78991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:386
		if len(interceptors) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:386
			_go_fuzz_dep_.CoverTab[78992]++
													chainedInt = interceptors[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:387
			// _ = "end of CoverTab[78992]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:388
			_go_fuzz_dep_.CoverTab[78993]++
													chainedInt = func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:389
				_go_fuzz_dep_.CoverTab[78994]++
														return interceptors[0](ctx, desc, cc, method, getChainStreamer(interceptors, 0, streamer), opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:390
				// _ = "end of CoverTab[78994]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:391
			// _ = "end of CoverTab[78993]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:392
		// _ = "end of CoverTab[78991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:392
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:392
	// _ = "end of CoverTab[78986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:392
	_go_fuzz_dep_.CoverTab[78987]++
											cc.dopts.streamInt = chainedInt
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:393
	// _ = "end of CoverTab[78987]"
}

// getChainStreamer recursively generate the chained client stream constructor.
func getChainStreamer(interceptors []StreamClientInterceptor, curr int, finalStreamer Streamer) Streamer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:397
	_go_fuzz_dep_.CoverTab[78995]++
											if curr == len(interceptors)-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:398
		_go_fuzz_dep_.CoverTab[78997]++
												return finalStreamer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:399
		// _ = "end of CoverTab[78997]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:400
		_go_fuzz_dep_.CoverTab[78998]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:400
		// _ = "end of CoverTab[78998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:400
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:400
	// _ = "end of CoverTab[78995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:400
	_go_fuzz_dep_.CoverTab[78996]++
											return func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (ClientStream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:401
		_go_fuzz_dep_.CoverTab[78999]++
												return interceptors[curr+1](ctx, desc, cc, method, getChainStreamer(interceptors, curr+1, finalStreamer), opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:402
		// _ = "end of CoverTab[78999]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:403
	// _ = "end of CoverTab[78996]"
}

// connectivityStateManager keeps the connectivity.State of ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:406
// This struct will eventually be exported so the balancers can access it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:408
type connectivityStateManager struct {
	mu		sync.Mutex
	state		connectivity.State
	notifyChan	chan struct{}
	channelzID	*channelz.Identifier
}

// updateState updates the connectivity.State of ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:415
// If there's a change it notifies goroutines waiting on state change to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:415
// happen.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:418
func (csm *connectivityStateManager) updateState(state connectivity.State) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:418
	_go_fuzz_dep_.CoverTab[79000]++
											csm.mu.Lock()
											defer csm.mu.Unlock()
											if csm.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:421
		_go_fuzz_dep_.CoverTab[79003]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:422
		// _ = "end of CoverTab[79003]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:423
		_go_fuzz_dep_.CoverTab[79004]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:423
		// _ = "end of CoverTab[79004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:423
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:423
	// _ = "end of CoverTab[79000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:423
	_go_fuzz_dep_.CoverTab[79001]++
											if csm.state == state {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:424
		_go_fuzz_dep_.CoverTab[79005]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:425
		// _ = "end of CoverTab[79005]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:426
		_go_fuzz_dep_.CoverTab[79006]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:426
		// _ = "end of CoverTab[79006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:426
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:426
	// _ = "end of CoverTab[79001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:426
	_go_fuzz_dep_.CoverTab[79002]++
											csm.state = state
											channelz.Infof(logger, csm.channelzID, "Channel Connectivity change to %v", state)
											if csm.notifyChan != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:429
		_go_fuzz_dep_.CoverTab[79007]++

												close(csm.notifyChan)
												csm.notifyChan = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:432
		// _ = "end of CoverTab[79007]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:433
		_go_fuzz_dep_.CoverTab[79008]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:433
		// _ = "end of CoverTab[79008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:433
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:433
	// _ = "end of CoverTab[79002]"
}

func (csm *connectivityStateManager) getState() connectivity.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:436
	_go_fuzz_dep_.CoverTab[79009]++
											csm.mu.Lock()
											defer csm.mu.Unlock()
											return csm.state
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:439
	// _ = "end of CoverTab[79009]"
}

func (csm *connectivityStateManager) getNotifyChan() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:442
	_go_fuzz_dep_.CoverTab[79010]++
											csm.mu.Lock()
											defer csm.mu.Unlock()
											if csm.notifyChan == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:445
		_go_fuzz_dep_.CoverTab[79012]++
												csm.notifyChan = make(chan struct{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:446
		// _ = "end of CoverTab[79012]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:447
		_go_fuzz_dep_.CoverTab[79013]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:447
		// _ = "end of CoverTab[79013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:447
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:447
	// _ = "end of CoverTab[79010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:447
	_go_fuzz_dep_.CoverTab[79011]++
											return csm.notifyChan
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:448
	// _ = "end of CoverTab[79011]"
}

// ClientConnInterface defines the functions clients need to perform unary and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:451
// streaming RPCs.  It is implemented by *ClientConn, and is only intended to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:451
// be referenced by generated code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:454
type ClientConnInterface interface {
	// Invoke performs a unary RPC and returns after the response is received
	// into reply.
	Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...CallOption) error
	// NewStream begins a streaming RPC.
	NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error)
}

// Assert *ClientConn implements ClientConnInterface.
var _ ClientConnInterface = (*ClientConn)(nil)

// ClientConn represents a virtual connection to a conceptual endpoint, to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// perform RPCs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// A ClientConn is free to have zero or more actual connections to the endpoint
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// based on configuration, load, etc. It is also free to determine which actual
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// endpoints to use and may change it every RPC, permitting client-side load
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// balancing.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// A ClientConn encapsulates a range of functionality including name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// resolution, TCP connection establishment (with retries and backoff) and TLS
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// handshakes. It also handles errors on established connections by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:465
// re-resolving the name and reconnecting.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:477
type ClientConn struct {
	ctx	context.Context		// Initialized using the background context at dial time.
	cancel	context.CancelFunc	// Cancelled on close.

	// The following are initialized at dial time, and are read-only after that.
	target		string			// User's dial target.
	parsedTarget	resolver.Target		// See parseTargetAndFindResolver().
	authority	string			// See determineAuthority().
	dopts		dialOptions		// Default and user specified dial options.
	channelzID	*channelz.Identifier	// Channelz identifier for the channel.
	balancerWrapper	*ccBalancerWrapper	// Uses gracefulswitch.balancer underneath.

	// The following provide their own synchronization, and therefore don't
	// require cc.mu to be held to access them.
	csMgr			*connectivityStateManager
	blockingpicker		*pickerWrapper
	safeConfigSelector	iresolver.SafeConfigSelector
	czData			*channelzData
	retryThrottler		atomic.Value	// Updated from service config.

	// firstResolveEvent is used to track whether the name resolver sent us at
	// least one update. RPCs block on this event.
	firstResolveEvent	*grpcsync.Event

	// mu protects the following fields.
	// TODO: split mu so the same mutex isn't used for everything.
	mu		sync.RWMutex
	resolverWrapper	*ccResolverWrapper		// Initialized in Dial; cleared in Close.
	sc		*ServiceConfig			// Latest service config received from the resolver.
	conns		map[*addrConn]struct{}		// Set to nil on close.
	mkp		keepalive.ClientParameters	// May be updated upon receipt of a GoAway.

	lceMu			sync.Mutex	// protects lastConnectionError
	lastConnectionError	error
}

// WaitForStateChange waits until the connectivity.State of ClientConn changes from sourceState or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
// ctx expires. A true value is returned in former case and false in latter.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:513
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:520
func (cc *ClientConn) WaitForStateChange(ctx context.Context, sourceState connectivity.State) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:520
	_go_fuzz_dep_.CoverTab[79014]++
											ch := cc.csMgr.getNotifyChan()
											if cc.csMgr.getState() != sourceState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:522
		_go_fuzz_dep_.CoverTab[79016]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:523
		// _ = "end of CoverTab[79016]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:524
		_go_fuzz_dep_.CoverTab[79017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:524
		// _ = "end of CoverTab[79017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:524
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:524
	// _ = "end of CoverTab[79014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:524
	_go_fuzz_dep_.CoverTab[79015]++
											select {
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:526
		_go_fuzz_dep_.CoverTab[79018]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:527
		// _ = "end of CoverTab[79018]"
	case <-ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:528
		_go_fuzz_dep_.CoverTab[79019]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:529
		// _ = "end of CoverTab[79019]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:530
	// _ = "end of CoverTab[79015]"
}

// GetState returns the connectivity.State of ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:533
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:533
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:533
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:533
// Notice: This API is EXPERIMENTAL and may be changed or removed in a later
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:533
// release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:539
func (cc *ClientConn) GetState() connectivity.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:539
	_go_fuzz_dep_.CoverTab[79020]++
											return cc.csMgr.getState()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:540
	// _ = "end of CoverTab[79020]"
}

// Connect causes all subchannels in the ClientConn to attempt to connect if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
// the channel is idle.  Does not wait for the connection attempts to begin
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
// before returning.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
// Notice: This API is EXPERIMENTAL and may be changed or removed in a later
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:543
// release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:551
func (cc *ClientConn) Connect() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:551
	_go_fuzz_dep_.CoverTab[79021]++
											cc.balancerWrapper.exitIdle()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:552
	// _ = "end of CoverTab[79021]"
}

func (cc *ClientConn) scWatcher() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:555
	_go_fuzz_dep_.CoverTab[79022]++
											for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:556
		_go_fuzz_dep_.CoverTab[79023]++
												select {
		case sc, ok := <-cc.dopts.scChan:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:558
			_go_fuzz_dep_.CoverTab[79024]++
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:559
				_go_fuzz_dep_.CoverTab[79027]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:560
				// _ = "end of CoverTab[79027]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:561
				_go_fuzz_dep_.CoverTab[79028]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:561
				// _ = "end of CoverTab[79028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:561
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:561
			// _ = "end of CoverTab[79024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:561
			_go_fuzz_dep_.CoverTab[79025]++
													cc.mu.Lock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:565
			cc.sc = &sc
													cc.safeConfigSelector.UpdateConfigSelector(&defaultConfigSelector{&sc})
													cc.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:567
			// _ = "end of CoverTab[79025]"
		case <-cc.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:568
			_go_fuzz_dep_.CoverTab[79026]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:569
			// _ = "end of CoverTab[79026]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:570
		// _ = "end of CoverTab[79023]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:571
	// _ = "end of CoverTab[79022]"
}

// waitForResolvedAddrs blocks until the resolver has provided addresses or the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:574
// context expires.  Returns nil unless the context expires first; otherwise
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:574
// returns a status error based on the context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:577
func (cc *ClientConn) waitForResolvedAddrs(ctx context.Context) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:577
	_go_fuzz_dep_.CoverTab[79029]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:580
	if cc.firstResolveEvent.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:580
		_go_fuzz_dep_.CoverTab[79031]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:581
		// _ = "end of CoverTab[79031]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:582
		_go_fuzz_dep_.CoverTab[79032]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:582
		// _ = "end of CoverTab[79032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:582
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:582
	// _ = "end of CoverTab[79029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:582
	_go_fuzz_dep_.CoverTab[79030]++
											select {
	case <-cc.firstResolveEvent.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:584
		_go_fuzz_dep_.CoverTab[79033]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:585
		// _ = "end of CoverTab[79033]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:586
		_go_fuzz_dep_.CoverTab[79034]++
												return status.FromContextError(ctx.Err()).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:587
		// _ = "end of CoverTab[79034]"
	case <-cc.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:588
		_go_fuzz_dep_.CoverTab[79035]++
												return ErrClientConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:589
		// _ = "end of CoverTab[79035]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:590
	// _ = "end of CoverTab[79030]"
}

var emptyServiceConfig *ServiceConfig

func init() {
	cfg := parseServiceConfig("{}")
	if cfg.Err != nil {
		panic(fmt.Sprintf("impossible error parsing empty service config: %v", cfg.Err))
	}
	emptyServiceConfig = cfg.Config.(*ServiceConfig)
}

func (cc *ClientConn) maybeApplyDefaultServiceConfig(addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:603
	_go_fuzz_dep_.CoverTab[79036]++
											if cc.sc != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:604
		_go_fuzz_dep_.CoverTab[79038]++
												cc.applyServiceConfigAndBalancer(cc.sc, nil, addrs)
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:606
		// _ = "end of CoverTab[79038]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:607
		_go_fuzz_dep_.CoverTab[79039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:607
		// _ = "end of CoverTab[79039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:607
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:607
	// _ = "end of CoverTab[79036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:607
	_go_fuzz_dep_.CoverTab[79037]++
											if cc.dopts.defaultServiceConfig != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:608
		_go_fuzz_dep_.CoverTab[79040]++
												cc.applyServiceConfigAndBalancer(cc.dopts.defaultServiceConfig, &defaultConfigSelector{cc.dopts.defaultServiceConfig}, addrs)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:609
		// _ = "end of CoverTab[79040]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:610
		_go_fuzz_dep_.CoverTab[79041]++
												cc.applyServiceConfigAndBalancer(emptyServiceConfig, &defaultConfigSelector{emptyServiceConfig}, addrs)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:611
		// _ = "end of CoverTab[79041]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:612
	// _ = "end of CoverTab[79037]"
}

func (cc *ClientConn) updateResolverState(s resolver.State, err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:615
	_go_fuzz_dep_.CoverTab[79042]++
											defer cc.firstResolveEvent.Fire()
											cc.mu.Lock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:621
	if cc.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:621
		_go_fuzz_dep_.CoverTab[79048]++
												cc.mu.Unlock()
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:623
		// _ = "end of CoverTab[79048]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:624
		_go_fuzz_dep_.CoverTab[79049]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:624
		// _ = "end of CoverTab[79049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:624
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:624
	// _ = "end of CoverTab[79042]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:624
	_go_fuzz_dep_.CoverTab[79043]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:626
		_go_fuzz_dep_.CoverTab[79050]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:630
		cc.maybeApplyDefaultServiceConfig(nil)

												cc.balancerWrapper.resolverError(err)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:635
		cc.mu.Unlock()
												return balancer.ErrBadResolverState
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:636
		// _ = "end of CoverTab[79050]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:637
		_go_fuzz_dep_.CoverTab[79051]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:637
		// _ = "end of CoverTab[79051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:637
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:637
	// _ = "end of CoverTab[79043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:637
	_go_fuzz_dep_.CoverTab[79044]++

											var ret error
											if cc.dopts.disableServiceConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:640
		_go_fuzz_dep_.CoverTab[79052]++
												channelz.Infof(logger, cc.channelzID, "ignoring service config from resolver (%v) and applying the default because service config is disabled", s.ServiceConfig)
												cc.maybeApplyDefaultServiceConfig(s.Addresses)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:642
		// _ = "end of CoverTab[79052]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:643
		_go_fuzz_dep_.CoverTab[79053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:643
		if s.ServiceConfig == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:643
			_go_fuzz_dep_.CoverTab[79054]++
													cc.maybeApplyDefaultServiceConfig(s.Addresses)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:644
			// _ = "end of CoverTab[79054]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:647
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:647
			_go_fuzz_dep_.CoverTab[79055]++
													if sc, ok := s.ServiceConfig.Config.(*ServiceConfig); s.ServiceConfig.Err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:648
				_go_fuzz_dep_.CoverTab[79056]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:648
				return ok
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:648
				// _ = "end of CoverTab[79056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:648
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:648
				_go_fuzz_dep_.CoverTab[79057]++
														configSelector := iresolver.GetConfigSelector(s)
														if configSelector != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:650
					_go_fuzz_dep_.CoverTab[79059]++
															if len(s.ServiceConfig.Config.(*ServiceConfig).Methods) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:651
						_go_fuzz_dep_.CoverTab[79060]++
																channelz.Infof(logger, cc.channelzID, "method configs in service config will be ignored due to presence of config selector")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:652
						// _ = "end of CoverTab[79060]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:653
						_go_fuzz_dep_.CoverTab[79061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:653
						// _ = "end of CoverTab[79061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:653
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:653
					// _ = "end of CoverTab[79059]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:654
					_go_fuzz_dep_.CoverTab[79062]++
															configSelector = &defaultConfigSelector{sc}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:655
					// _ = "end of CoverTab[79062]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:656
				// _ = "end of CoverTab[79057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:656
				_go_fuzz_dep_.CoverTab[79058]++
														cc.applyServiceConfigAndBalancer(sc, configSelector, s.Addresses)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:657
				// _ = "end of CoverTab[79058]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:658
				_go_fuzz_dep_.CoverTab[79063]++
														ret = balancer.ErrBadResolverState
														if cc.sc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:660
					_go_fuzz_dep_.CoverTab[79064]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:663
					cc.applyFailingLB(s.ServiceConfig)
															cc.mu.Unlock()
															return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:665
					// _ = "end of CoverTab[79064]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:666
					_go_fuzz_dep_.CoverTab[79065]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:666
					// _ = "end of CoverTab[79065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:666
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:666
				// _ = "end of CoverTab[79063]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:667
			// _ = "end of CoverTab[79055]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:668
		// _ = "end of CoverTab[79053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:668
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:668
	// _ = "end of CoverTab[79044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:668
	_go_fuzz_dep_.CoverTab[79045]++

											var balCfg serviceconfig.LoadBalancingConfig
											if cc.sc != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:671
		_go_fuzz_dep_.CoverTab[79066]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:671
		return cc.sc.lbConfig != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:671
		// _ = "end of CoverTab[79066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:671
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:671
		_go_fuzz_dep_.CoverTab[79067]++
												balCfg = cc.sc.lbConfig.cfg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:672
		// _ = "end of CoverTab[79067]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:673
		_go_fuzz_dep_.CoverTab[79068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:673
		// _ = "end of CoverTab[79068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:673
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:673
	// _ = "end of CoverTab[79045]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:673
	_go_fuzz_dep_.CoverTab[79046]++
											bw := cc.balancerWrapper
											cc.mu.Unlock()

											uccsErr := bw.updateClientConnState(&balancer.ClientConnState{ResolverState: s, BalancerConfig: balCfg})
											if ret == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:678
		_go_fuzz_dep_.CoverTab[79069]++
												ret = uccsErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:679
		// _ = "end of CoverTab[79069]"

	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:681
		_go_fuzz_dep_.CoverTab[79070]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:681
		// _ = "end of CoverTab[79070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:681
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:681
	// _ = "end of CoverTab[79046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:681
	_go_fuzz_dep_.CoverTab[79047]++
											return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:682
	// _ = "end of CoverTab[79047]"
}

// applyFailingLB is akin to configuring an LB policy on the channel which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// always fails RPCs. Here, an actual LB policy is not configured, but an always
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// erroring picker is configured, which returns errors with information about
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// what was invalid in the received service config. A config selector with no
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// service config is configured, and the connectivity state of the channel is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// set to TransientFailure.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:685
// Caller must hold cc.mu.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:693
func (cc *ClientConn) applyFailingLB(sc *serviceconfig.ParseResult) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:693
	_go_fuzz_dep_.CoverTab[79071]++
											var err error
											if sc.Err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:695
		_go_fuzz_dep_.CoverTab[79073]++
												err = status.Errorf(codes.Unavailable, "error parsing service config: %v", sc.Err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:696
		// _ = "end of CoverTab[79073]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:697
		_go_fuzz_dep_.CoverTab[79074]++
												err = status.Errorf(codes.Unavailable, "illegal service config type: %T", sc.Config)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:698
		// _ = "end of CoverTab[79074]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:699
	// _ = "end of CoverTab[79071]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:699
	_go_fuzz_dep_.CoverTab[79072]++
											cc.safeConfigSelector.UpdateConfigSelector(&defaultConfigSelector{nil})
											cc.blockingpicker.updatePicker(base.NewErrPicker(err))
											cc.csMgr.updateState(connectivity.TransientFailure)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:702
	// _ = "end of CoverTab[79072]"
}

func (cc *ClientConn) handleSubConnStateChange(sc balancer.SubConn, s connectivity.State, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:705
	_go_fuzz_dep_.CoverTab[79075]++
											cc.balancerWrapper.updateSubConnState(sc, s, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:706
	// _ = "end of CoverTab[79075]"
}

// newAddrConn creates an addrConn for addrs and adds it to cc.conns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:709
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:709
// Caller needs to make sure len(addrs) > 0.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:712
func (cc *ClientConn) newAddrConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (*addrConn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:712
	_go_fuzz_dep_.CoverTab[79076]++
											ac := &addrConn{
		state:		connectivity.Idle,
		cc:		cc,
		addrs:		addrs,
		scopts:		opts,
		dopts:		cc.dopts,
		czData:		new(channelzData),
		resetBackoff:	make(chan struct{}),
	}
	ac.ctx, ac.cancel = context.WithCancel(cc.ctx)

	cc.mu.Lock()
	defer cc.mu.Unlock()
	if cc.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:726
		_go_fuzz_dep_.CoverTab[79079]++
												return nil, ErrClientConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:727
		// _ = "end of CoverTab[79079]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:728
		_go_fuzz_dep_.CoverTab[79080]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:728
		// _ = "end of CoverTab[79080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:728
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:728
	// _ = "end of CoverTab[79076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:728
	_go_fuzz_dep_.CoverTab[79077]++

											var err error
											ac.channelzID, err = channelz.RegisterSubChannel(ac, cc.channelzID, "")
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:732
		_go_fuzz_dep_.CoverTab[79081]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:733
		// _ = "end of CoverTab[79081]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:734
		_go_fuzz_dep_.CoverTab[79082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:734
		// _ = "end of CoverTab[79082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:734
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:734
	// _ = "end of CoverTab[79077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:734
	_go_fuzz_dep_.CoverTab[79078]++
											channelz.AddTraceEvent(logger, ac.channelzID, 0, &channelz.TraceEventDesc{
		Desc:		"Subchannel created",
		Severity:	channelz.CtInfo,
		Parent: &channelz.TraceEventDesc{
			Desc:		fmt.Sprintf("Subchannel(id:%d) created", ac.channelzID.Int()),
			Severity:	channelz.CtInfo,
		},
	})

											cc.conns[ac] = struct{}{}
											return ac, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:745
	// _ = "end of CoverTab[79078]"
}

// removeAddrConn removes the addrConn in the subConn from clientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:748
// It also tears down the ac with the given error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:750
func (cc *ClientConn) removeAddrConn(ac *addrConn, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:750
	_go_fuzz_dep_.CoverTab[79083]++
											cc.mu.Lock()
											if cc.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:752
		_go_fuzz_dep_.CoverTab[79085]++
												cc.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:754
		// _ = "end of CoverTab[79085]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:755
		_go_fuzz_dep_.CoverTab[79086]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:755
		// _ = "end of CoverTab[79086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:755
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:755
	// _ = "end of CoverTab[79083]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:755
	_go_fuzz_dep_.CoverTab[79084]++
											delete(cc.conns, ac)
											cc.mu.Unlock()
											ac.tearDown(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:758
	// _ = "end of CoverTab[79084]"
}

func (cc *ClientConn) channelzMetric() *channelz.ChannelInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:761
	_go_fuzz_dep_.CoverTab[79087]++
											return &channelz.ChannelInternalMetric{
		State:				cc.GetState(),
		Target:				cc.target,
		CallsStarted:			atomic.LoadInt64(&cc.czData.callsStarted),
		CallsSucceeded:			atomic.LoadInt64(&cc.czData.callsSucceeded),
		CallsFailed:			atomic.LoadInt64(&cc.czData.callsFailed),
		LastCallStartedTimestamp:	time.Unix(0, atomic.LoadInt64(&cc.czData.lastCallStartedTime)),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:769
	// _ = "end of CoverTab[79087]"
}

// Target returns the target string of the ClientConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:772
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:772
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:772
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:772
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:772
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:778
func (cc *ClientConn) Target() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:778
	_go_fuzz_dep_.CoverTab[79088]++
											return cc.target
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:779
	// _ = "end of CoverTab[79088]"
}

func (cc *ClientConn) incrCallsStarted() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:782
	_go_fuzz_dep_.CoverTab[79089]++
											atomic.AddInt64(&cc.czData.callsStarted, 1)
											atomic.StoreInt64(&cc.czData.lastCallStartedTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:784
	// _ = "end of CoverTab[79089]"
}

func (cc *ClientConn) incrCallsSucceeded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:787
	_go_fuzz_dep_.CoverTab[79090]++
											atomic.AddInt64(&cc.czData.callsSucceeded, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:788
	// _ = "end of CoverTab[79090]"
}

func (cc *ClientConn) incrCallsFailed() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:791
	_go_fuzz_dep_.CoverTab[79091]++
											atomic.AddInt64(&cc.czData.callsFailed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:792
	// _ = "end of CoverTab[79091]"
}

// connect starts creating a transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:795
// It does nothing if the ac is not IDLE.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:795
// TODO(bar) Move this to the addrConn section.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:798
func (ac *addrConn) connect() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:798
	_go_fuzz_dep_.CoverTab[79092]++
											ac.mu.Lock()
											if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:800
		_go_fuzz_dep_.CoverTab[79095]++
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:801
			_go_fuzz_dep_.CoverTab[79097]++
													logger.Infof("connect called on shutdown addrConn; ignoring.")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:802
			// _ = "end of CoverTab[79097]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:803
			_go_fuzz_dep_.CoverTab[79098]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:803
			// _ = "end of CoverTab[79098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:803
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:803
		// _ = "end of CoverTab[79095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:803
		_go_fuzz_dep_.CoverTab[79096]++
												ac.mu.Unlock()
												return errConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:805
		// _ = "end of CoverTab[79096]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:806
		_go_fuzz_dep_.CoverTab[79099]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:806
		// _ = "end of CoverTab[79099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:806
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:806
	// _ = "end of CoverTab[79092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:806
	_go_fuzz_dep_.CoverTab[79093]++
											if ac.state != connectivity.Idle {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:807
		_go_fuzz_dep_.CoverTab[79100]++
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:808
			_go_fuzz_dep_.CoverTab[79102]++
													logger.Infof("connect called on addrConn in non-idle state (%v); ignoring.", ac.state)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:809
			// _ = "end of CoverTab[79102]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:810
			_go_fuzz_dep_.CoverTab[79103]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:810
			// _ = "end of CoverTab[79103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:810
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:810
		// _ = "end of CoverTab[79100]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:810
		_go_fuzz_dep_.CoverTab[79101]++
												ac.mu.Unlock()
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:812
		// _ = "end of CoverTab[79101]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:813
		_go_fuzz_dep_.CoverTab[79104]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:813
		// _ = "end of CoverTab[79104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:813
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:813
	// _ = "end of CoverTab[79093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:813
	_go_fuzz_dep_.CoverTab[79094]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:816
	ac.updateConnectivityState(connectivity.Connecting, nil)
											ac.mu.Unlock()

											ac.resetTransport()
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:820
	// _ = "end of CoverTab[79094]"
}

func equalAddresses(a, b []resolver.Address) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:823
	_go_fuzz_dep_.CoverTab[79105]++
											if len(a) != len(b) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:824
		_go_fuzz_dep_.CoverTab[79108]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:825
		// _ = "end of CoverTab[79108]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:826
		_go_fuzz_dep_.CoverTab[79109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:826
		// _ = "end of CoverTab[79109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:826
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:826
	// _ = "end of CoverTab[79105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:826
	_go_fuzz_dep_.CoverTab[79106]++
											for i, v := range a {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:827
		_go_fuzz_dep_.CoverTab[79110]++
												if !v.Equal(b[i]) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:828
			_go_fuzz_dep_.CoverTab[79111]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:829
			// _ = "end of CoverTab[79111]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:830
			_go_fuzz_dep_.CoverTab[79112]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:830
			// _ = "end of CoverTab[79112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:830
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:830
		// _ = "end of CoverTab[79110]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:831
	// _ = "end of CoverTab[79106]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:831
	_go_fuzz_dep_.CoverTab[79107]++
											return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:832
	// _ = "end of CoverTab[79107]"
}

// tryUpdateAddrs tries to update ac.addrs with the new addresses list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// If ac is TransientFailure, it updates ac.addrs and returns true. The updated
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// addresses will be picked up by retry in the next iteration after backoff.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// If ac is Shutdown or Idle, it updates ac.addrs and returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// If the addresses is the same as the old list, it does nothing and returns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// If ac is Connecting, it returns false. The caller should tear down the ac and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// create a new one. Note that the backoff will be reset when this happens.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// If ac is Ready, it checks whether current connected address of ac is in the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
// new addrs list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//   - If true, it updates ac.addrs and returns true. The ac will keep using
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//     the existing connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:835
//   - If false, it does nothing and returns false.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:853
func (ac *addrConn) tryUpdateAddrs(addrs []resolver.Address) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:853
	_go_fuzz_dep_.CoverTab[79113]++
											ac.mu.Lock()
											defer ac.mu.Unlock()
											channelz.Infof(logger, ac.channelzID, "addrConn: tryUpdateAddrs curAddr: %v, addrs: %v", ac.curAddr, addrs)
											if ac.state == connectivity.Shutdown || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:857
		_go_fuzz_dep_.CoverTab[79119]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:857
		return ac.state == connectivity.TransientFailure
												// _ = "end of CoverTab[79119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:858
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:858
		_go_fuzz_dep_.CoverTab[79120]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:858
		return ac.state == connectivity.Idle
												// _ = "end of CoverTab[79120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:859
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:859
		_go_fuzz_dep_.CoverTab[79121]++
												ac.addrs = addrs
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:861
		// _ = "end of CoverTab[79121]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:862
		_go_fuzz_dep_.CoverTab[79122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:862
		// _ = "end of CoverTab[79122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:862
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:862
	// _ = "end of CoverTab[79113]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:862
	_go_fuzz_dep_.CoverTab[79114]++

											if equalAddresses(ac.addrs, addrs) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:864
		_go_fuzz_dep_.CoverTab[79123]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:865
		// _ = "end of CoverTab[79123]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:866
		_go_fuzz_dep_.CoverTab[79124]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:866
		// _ = "end of CoverTab[79124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:866
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:866
	// _ = "end of CoverTab[79114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:866
	_go_fuzz_dep_.CoverTab[79115]++

											if ac.state == connectivity.Connecting {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:868
		_go_fuzz_dep_.CoverTab[79125]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:869
		// _ = "end of CoverTab[79125]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:870
		_go_fuzz_dep_.CoverTab[79126]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:870
		// _ = "end of CoverTab[79126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:870
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:870
	// _ = "end of CoverTab[79115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:870
	_go_fuzz_dep_.CoverTab[79116]++

	// ac.state is Ready, try to find the connected address.
	var curAddrFound bool
	for _, a := range addrs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:874
		_go_fuzz_dep_.CoverTab[79127]++
												a.ServerName = ac.cc.getServerName(a)
												if reflect.DeepEqual(ac.curAddr, a) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:876
			_go_fuzz_dep_.CoverTab[79128]++
													curAddrFound = true
													break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:878
			// _ = "end of CoverTab[79128]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:879
			_go_fuzz_dep_.CoverTab[79129]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:879
			// _ = "end of CoverTab[79129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:879
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:879
		// _ = "end of CoverTab[79127]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:880
	// _ = "end of CoverTab[79116]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:880
	_go_fuzz_dep_.CoverTab[79117]++
											channelz.Infof(logger, ac.channelzID, "addrConn: tryUpdateAddrs curAddrFound: %v", curAddrFound)
											if curAddrFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:882
		_go_fuzz_dep_.CoverTab[79130]++
												ac.addrs = addrs
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:883
		// _ = "end of CoverTab[79130]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:884
		_go_fuzz_dep_.CoverTab[79131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:884
		// _ = "end of CoverTab[79131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:884
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:884
	// _ = "end of CoverTab[79117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:884
	_go_fuzz_dep_.CoverTab[79118]++

											return curAddrFound
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:886
	// _ = "end of CoverTab[79118]"
}

// getServerName determines the serverName to be used in the connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// handshake. The default value for the serverName is the authority on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// ClientConn, which either comes from the user's dial target or through an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// authority override specified using the WithAuthority dial option. Name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// resolvers can specify a per-address override for the serverName through the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// resolver.Address.ServerName field which is used only if the WithAuthority
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// dial option was not used. The rationale is that per-address authority
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// overrides specified by the name resolver can represent a security risk, while
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// an override specified by the user is more dependable since they probably know
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:889
// what they are doing.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:899
func (cc *ClientConn) getServerName(addr resolver.Address) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:899
	_go_fuzz_dep_.CoverTab[79132]++
											if cc.dopts.authority != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:900
		_go_fuzz_dep_.CoverTab[79135]++
												return cc.dopts.authority
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:901
		// _ = "end of CoverTab[79135]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:902
		_go_fuzz_dep_.CoverTab[79136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:902
		// _ = "end of CoverTab[79136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:902
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:902
	// _ = "end of CoverTab[79132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:902
	_go_fuzz_dep_.CoverTab[79133]++
											if addr.ServerName != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:903
		_go_fuzz_dep_.CoverTab[79137]++
												return addr.ServerName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:904
		// _ = "end of CoverTab[79137]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:905
		_go_fuzz_dep_.CoverTab[79138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:905
		// _ = "end of CoverTab[79138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:905
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:905
	// _ = "end of CoverTab[79133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:905
	_go_fuzz_dep_.CoverTab[79134]++
											return cc.authority
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:906
	// _ = "end of CoverTab[79134]"
}

func getMethodConfig(sc *ServiceConfig, method string) MethodConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:909
	_go_fuzz_dep_.CoverTab[79139]++
											if sc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:910
		_go_fuzz_dep_.CoverTab[79143]++
												return MethodConfig{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:911
		// _ = "end of CoverTab[79143]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:912
		_go_fuzz_dep_.CoverTab[79144]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:912
		// _ = "end of CoverTab[79144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:912
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:912
	// _ = "end of CoverTab[79139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:912
	_go_fuzz_dep_.CoverTab[79140]++
											if m, ok := sc.Methods[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:913
		_go_fuzz_dep_.CoverTab[79145]++
												return m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:914
		// _ = "end of CoverTab[79145]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:915
		_go_fuzz_dep_.CoverTab[79146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:915
		// _ = "end of CoverTab[79146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:915
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:915
	// _ = "end of CoverTab[79140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:915
	_go_fuzz_dep_.CoverTab[79141]++
											i := strings.LastIndex(method, "/")
											if m, ok := sc.Methods[method[:i+1]]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:917
		_go_fuzz_dep_.CoverTab[79147]++
												return m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:918
		// _ = "end of CoverTab[79147]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:919
		_go_fuzz_dep_.CoverTab[79148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:919
		// _ = "end of CoverTab[79148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:919
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:919
	// _ = "end of CoverTab[79141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:919
	_go_fuzz_dep_.CoverTab[79142]++
											return sc.Methods[""]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:920
	// _ = "end of CoverTab[79142]"
}

// GetMethodConfig gets the method config of the input method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// If there's an exact match for input method (i.e. /service/method), we return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// the corresponding MethodConfig.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// If there isn't an exact match for the input method, we look for the service's default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// config under the service (i.e /service/) and then for the default for all services (empty string).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// If there is a default MethodConfig for the service, we return it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:923
// Otherwise, we return an empty MethodConfig.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:931
func (cc *ClientConn) GetMethodConfig(method string) MethodConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:931
	_go_fuzz_dep_.CoverTab[79149]++

											cc.mu.RLock()
											defer cc.mu.RUnlock()
											return getMethodConfig(cc.sc, method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:935
	// _ = "end of CoverTab[79149]"
}

func (cc *ClientConn) healthCheckConfig() *healthCheckConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:938
	_go_fuzz_dep_.CoverTab[79150]++
											cc.mu.RLock()
											defer cc.mu.RUnlock()
											if cc.sc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:941
		_go_fuzz_dep_.CoverTab[79152]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:942
		// _ = "end of CoverTab[79152]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:943
		_go_fuzz_dep_.CoverTab[79153]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:943
		// _ = "end of CoverTab[79153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:943
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:943
	// _ = "end of CoverTab[79150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:943
	_go_fuzz_dep_.CoverTab[79151]++
											return cc.sc.healthCheckConfig
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:944
	// _ = "end of CoverTab[79151]"
}

func (cc *ClientConn) getTransport(ctx context.Context, failfast bool, method string) (transport.ClientTransport, balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:947
	_go_fuzz_dep_.CoverTab[79154]++
											return cc.blockingpicker.pick(ctx, failfast, balancer.PickInfo{
		Ctx:		ctx,
		FullMethodName:	method,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:951
	// _ = "end of CoverTab[79154]"
}

func (cc *ClientConn) applyServiceConfigAndBalancer(sc *ServiceConfig, configSelector iresolver.ConfigSelector, addrs []resolver.Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:954
	_go_fuzz_dep_.CoverTab[79155]++
											if sc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:955
		_go_fuzz_dep_.CoverTab[79160]++

												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:957
		// _ = "end of CoverTab[79160]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:958
		_go_fuzz_dep_.CoverTab[79161]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:958
		// _ = "end of CoverTab[79161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:958
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:958
	// _ = "end of CoverTab[79155]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:958
	_go_fuzz_dep_.CoverTab[79156]++
											cc.sc = sc
											if configSelector != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:960
		_go_fuzz_dep_.CoverTab[79162]++
												cc.safeConfigSelector.UpdateConfigSelector(configSelector)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:961
		// _ = "end of CoverTab[79162]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:962
		_go_fuzz_dep_.CoverTab[79163]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:962
		// _ = "end of CoverTab[79163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:962
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:962
	// _ = "end of CoverTab[79156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:962
	_go_fuzz_dep_.CoverTab[79157]++

											if cc.sc.retryThrottling != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:964
		_go_fuzz_dep_.CoverTab[79164]++
												newThrottler := &retryThrottler{
			tokens:	cc.sc.retryThrottling.MaxTokens,
			max:	cc.sc.retryThrottling.MaxTokens,
			thresh:	cc.sc.retryThrottling.MaxTokens / 2,
			ratio:	cc.sc.retryThrottling.TokenRatio,
		}
												cc.retryThrottler.Store(newThrottler)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:971
		// _ = "end of CoverTab[79164]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:972
		_go_fuzz_dep_.CoverTab[79165]++
												cc.retryThrottler.Store((*retryThrottler)(nil))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:973
		// _ = "end of CoverTab[79165]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:974
	// _ = "end of CoverTab[79157]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:974
	_go_fuzz_dep_.CoverTab[79158]++

											var newBalancerName string
											if cc.sc != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:977
		_go_fuzz_dep_.CoverTab[79166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:977
		return cc.sc.lbConfig != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:977
		// _ = "end of CoverTab[79166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:977
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:977
		_go_fuzz_dep_.CoverTab[79167]++
												newBalancerName = cc.sc.lbConfig.name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:978
		// _ = "end of CoverTab[79167]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:979
		_go_fuzz_dep_.CoverTab[79168]++
												var isGRPCLB bool
												for _, a := range addrs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:981
			_go_fuzz_dep_.CoverTab[79170]++
													if a.Type == resolver.GRPCLB {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:982
				_go_fuzz_dep_.CoverTab[79171]++
														isGRPCLB = true
														break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:984
				// _ = "end of CoverTab[79171]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:985
				_go_fuzz_dep_.CoverTab[79172]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:985
				// _ = "end of CoverTab[79172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:985
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:985
			// _ = "end of CoverTab[79170]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:986
		// _ = "end of CoverTab[79168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:986
		_go_fuzz_dep_.CoverTab[79169]++
												if isGRPCLB {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:987
			_go_fuzz_dep_.CoverTab[79173]++
													newBalancerName = grpclbName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:988
			// _ = "end of CoverTab[79173]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
			_go_fuzz_dep_.CoverTab[79174]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
			if cc.sc != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
				_go_fuzz_dep_.CoverTab[79175]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
				return cc.sc.LB != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
				// _ = "end of CoverTab[79175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:989
				_go_fuzz_dep_.CoverTab[79176]++
														newBalancerName = *cc.sc.LB
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:990
				// _ = "end of CoverTab[79176]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:991
				_go_fuzz_dep_.CoverTab[79177]++
														newBalancerName = PickFirstBalancerName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:992
				// _ = "end of CoverTab[79177]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:993
			// _ = "end of CoverTab[79174]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:993
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:993
		// _ = "end of CoverTab[79169]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:994
	// _ = "end of CoverTab[79158]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:994
	_go_fuzz_dep_.CoverTab[79159]++
											cc.balancerWrapper.switchTo(newBalancerName)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:995
	// _ = "end of CoverTab[79159]"
}

func (cc *ClientConn) resolveNow(o resolver.ResolveNowOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:998
	_go_fuzz_dep_.CoverTab[79178]++
											cc.mu.RLock()
											r := cc.resolverWrapper
											cc.mu.RUnlock()
											if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1002
		_go_fuzz_dep_.CoverTab[79180]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1003
		// _ = "end of CoverTab[79180]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
		_go_fuzz_dep_.CoverTab[79181]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
		// _ = "end of CoverTab[79181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
	// _ = "end of CoverTab[79178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
	_go_fuzz_dep_.CoverTab[79179]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
	_curRoutineNum89_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1004
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum89_)
											go r.resolveNow(o)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1005
	// _ = "end of CoverTab[79179]"
}

// ResetConnectBackoff wakes up all subchannels in transient failure and causes
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// them to attempt another connection immediately.  It also resets the backoff
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// times used for subsequent attempts regardless of the current state.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// In general, this function should not be used.  Typical service or network
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// outages result in a reasonable client reconnection strategy by default.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// However, if a previously unavailable network becomes available, this may be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// used to trigger an immediate reconnect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1008
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1021
func (cc *ClientConn) ResetConnectBackoff() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1021
	_go_fuzz_dep_.CoverTab[79182]++
											cc.mu.Lock()
											conns := cc.conns
											cc.mu.Unlock()
											for ac := range conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1025
		_go_fuzz_dep_.CoverTab[79183]++
												ac.resetConnectBackoff()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1026
		// _ = "end of CoverTab[79183]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1027
	// _ = "end of CoverTab[79182]"
}

// Close tears down the ClientConn and all underlying connections.
func (cc *ClientConn) Close() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1031
	_go_fuzz_dep_.CoverTab[79184]++
											defer cc.cancel()

											cc.mu.Lock()
											if cc.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1035
		_go_fuzz_dep_.CoverTab[79190]++
												cc.mu.Unlock()
												return ErrClientConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1037
		// _ = "end of CoverTab[79190]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1038
		_go_fuzz_dep_.CoverTab[79191]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1038
		// _ = "end of CoverTab[79191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1038
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1038
	// _ = "end of CoverTab[79184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1038
	_go_fuzz_dep_.CoverTab[79185]++
											conns := cc.conns
											cc.conns = nil
											cc.csMgr.updateState(connectivity.Shutdown)

											rWrapper := cc.resolverWrapper
											cc.resolverWrapper = nil
											bWrapper := cc.balancerWrapper
											cc.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1050
	cc.blockingpicker.close()
	if bWrapper != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1051
		_go_fuzz_dep_.CoverTab[79192]++
												bWrapper.close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1052
		// _ = "end of CoverTab[79192]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1053
		_go_fuzz_dep_.CoverTab[79193]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1053
		// _ = "end of CoverTab[79193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1053
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1053
	// _ = "end of CoverTab[79185]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1053
	_go_fuzz_dep_.CoverTab[79186]++
											if rWrapper != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1054
		_go_fuzz_dep_.CoverTab[79194]++
												rWrapper.close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1055
		// _ = "end of CoverTab[79194]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1056
		_go_fuzz_dep_.CoverTab[79195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1056
		// _ = "end of CoverTab[79195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1056
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1056
	// _ = "end of CoverTab[79186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1056
	_go_fuzz_dep_.CoverTab[79187]++

											for ac := range conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1058
		_go_fuzz_dep_.CoverTab[79196]++
												ac.tearDown(ErrClientConnClosing)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1059
		// _ = "end of CoverTab[79196]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1060
	// _ = "end of CoverTab[79187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1060
	_go_fuzz_dep_.CoverTab[79188]++
											ted := &channelz.TraceEventDesc{
		Desc:		"Channel deleted",
		Severity:	channelz.CtInfo,
	}
	if cc.dopts.channelzParentID != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1065
		_go_fuzz_dep_.CoverTab[79197]++
												ted.Parent = &channelz.TraceEventDesc{
			Desc:		fmt.Sprintf("Nested channel(id:%d) deleted", cc.channelzID.Int()),
			Severity:	channelz.CtInfo,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1069
		// _ = "end of CoverTab[79197]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1070
		_go_fuzz_dep_.CoverTab[79198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1070
		// _ = "end of CoverTab[79198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1070
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1070
	// _ = "end of CoverTab[79188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1070
	_go_fuzz_dep_.CoverTab[79189]++
											channelz.AddTraceEvent(logger, cc.channelzID, 0, ted)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1075
	channelz.RemoveEntry(cc.channelzID)

											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1077
	// _ = "end of CoverTab[79189]"
}

// addrConn is a network connection to a given address.
type addrConn struct {
	ctx	context.Context
	cancel	context.CancelFunc

	cc	*ClientConn
	dopts	dialOptions
	acbw	balancer.SubConn
	scopts	balancer.NewSubConnOptions

	// transport is set when there's a viable transport (note: ac state may not be READY as LB channel
	// health checking may require server to report healthy to set ac to READY), and is reset
	// to nil when the current transport should no longer be used to create a stream (e.g. after GoAway
	// is received, transport is closed, ac has been torn down).
	transport	transport.ClientTransport	// The current transport.

	mu	sync.Mutex
	curAddr	resolver.Address	// The current address.
	addrs	[]resolver.Address	// All addresses that the resolver resolved to.

	// Use updateConnectivityState for updating addrConn's connectivity state.
	state	connectivity.State

	backoffIdx	int	// Needs to be stateful for resetConnectBackoff.
	resetBackoff	chan struct{}

	channelzID	*channelz.Identifier
	czData		*channelzData
}

// Note: this requires a lock on ac.mu.
func (ac *addrConn) updateConnectivityState(s connectivity.State, lastErr error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1111
	_go_fuzz_dep_.CoverTab[79199]++
											if ac.state == s {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1112
		_go_fuzz_dep_.CoverTab[79202]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1113
		// _ = "end of CoverTab[79202]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1114
		_go_fuzz_dep_.CoverTab[79203]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1114
		// _ = "end of CoverTab[79203]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1114
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1114
	// _ = "end of CoverTab[79199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1114
	_go_fuzz_dep_.CoverTab[79200]++
											ac.state = s
											if lastErr == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1116
		_go_fuzz_dep_.CoverTab[79204]++
												channelz.Infof(logger, ac.channelzID, "Subchannel Connectivity change to %v", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1117
		// _ = "end of CoverTab[79204]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1118
		_go_fuzz_dep_.CoverTab[79205]++
												channelz.Infof(logger, ac.channelzID, "Subchannel Connectivity change to %v, last error: %s", s, lastErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1119
		// _ = "end of CoverTab[79205]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1120
	// _ = "end of CoverTab[79200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1120
	_go_fuzz_dep_.CoverTab[79201]++
											ac.cc.handleSubConnStateChange(ac.acbw, s, lastErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1121
	// _ = "end of CoverTab[79201]"
}

// adjustParams updates parameters used to create transports upon
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1124
// receiving a GoAway.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1126
func (ac *addrConn) adjustParams(r transport.GoAwayReason) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1126
	_go_fuzz_dep_.CoverTab[79206]++
											switch r {
	case transport.GoAwayTooManyPings:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1128
		_go_fuzz_dep_.CoverTab[79207]++
												v := 2 * ac.dopts.copts.KeepaliveParams.Time
												ac.cc.mu.Lock()
												if v > ac.cc.mkp.Time {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1131
			_go_fuzz_dep_.CoverTab[79210]++
													ac.cc.mkp.Time = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1132
			// _ = "end of CoverTab[79210]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1133
			_go_fuzz_dep_.CoverTab[79211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1133
			// _ = "end of CoverTab[79211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1133
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1133
		// _ = "end of CoverTab[79207]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1133
		_go_fuzz_dep_.CoverTab[79208]++
												ac.cc.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1134
		// _ = "end of CoverTab[79208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1134
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1134
		_go_fuzz_dep_.CoverTab[79209]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1134
		// _ = "end of CoverTab[79209]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1135
	// _ = "end of CoverTab[79206]"
}

func (ac *addrConn) resetTransport() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1138
	_go_fuzz_dep_.CoverTab[79212]++
											ac.mu.Lock()
											if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1140
		_go_fuzz_dep_.CoverTab[79217]++
												ac.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1142
		// _ = "end of CoverTab[79217]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1143
		_go_fuzz_dep_.CoverTab[79218]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1143
		// _ = "end of CoverTab[79218]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1143
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1143
	// _ = "end of CoverTab[79212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1143
	_go_fuzz_dep_.CoverTab[79213]++

											addrs := ac.addrs
											backoffFor := ac.dopts.bs.Backoff(ac.backoffIdx)

											dialDuration := minConnectTimeout
											if ac.dopts.minConnectTimeout != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1149
		_go_fuzz_dep_.CoverTab[79219]++
												dialDuration = ac.dopts.minConnectTimeout()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1150
		// _ = "end of CoverTab[79219]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1151
		_go_fuzz_dep_.CoverTab[79220]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1151
		// _ = "end of CoverTab[79220]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1151
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1151
	// _ = "end of CoverTab[79213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1151
	_go_fuzz_dep_.CoverTab[79214]++

											if dialDuration < backoffFor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1153
		_go_fuzz_dep_.CoverTab[79221]++

												dialDuration = backoffFor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1155
		// _ = "end of CoverTab[79221]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1156
		_go_fuzz_dep_.CoverTab[79222]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1156
		// _ = "end of CoverTab[79222]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1156
	// _ = "end of CoverTab[79214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1156
	_go_fuzz_dep_.CoverTab[79215]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1163
	connectDeadline := time.Now().Add(dialDuration)

	ac.updateConnectivityState(connectivity.Connecting, nil)
	ac.mu.Unlock()

	if err := ac.tryAllAddrs(addrs, connectDeadline); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1168
		_go_fuzz_dep_.CoverTab[79223]++
												ac.cc.resolveNow(resolver.ResolveNowOptions{})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1172
		ac.mu.Lock()
		if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1173
			_go_fuzz_dep_.CoverTab[79227]++
													ac.mu.Unlock()
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1175
			// _ = "end of CoverTab[79227]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1176
			_go_fuzz_dep_.CoverTab[79228]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1176
			// _ = "end of CoverTab[79228]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1176
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1176
		// _ = "end of CoverTab[79223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1176
		_go_fuzz_dep_.CoverTab[79224]++
												ac.updateConnectivityState(connectivity.TransientFailure, err)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1180
		b := ac.resetBackoff
		ac.mu.Unlock()

		timer := time.NewTimer(backoffFor)
		select {
		case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1185
			_go_fuzz_dep_.CoverTab[79229]++
													ac.mu.Lock()
													ac.backoffIdx++
													ac.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1188
			// _ = "end of CoverTab[79229]"
		case <-b:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1189
			_go_fuzz_dep_.CoverTab[79230]++
													timer.Stop()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1190
			// _ = "end of CoverTab[79230]"
		case <-ac.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1191
			_go_fuzz_dep_.CoverTab[79231]++
													timer.Stop()
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1193
			// _ = "end of CoverTab[79231]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1194
		// _ = "end of CoverTab[79224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1194
		_go_fuzz_dep_.CoverTab[79225]++

												ac.mu.Lock()
												if ac.state != connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1197
			_go_fuzz_dep_.CoverTab[79232]++
													ac.updateConnectivityState(connectivity.Idle, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1198
			// _ = "end of CoverTab[79232]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1199
			_go_fuzz_dep_.CoverTab[79233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1199
			// _ = "end of CoverTab[79233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1199
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1199
		// _ = "end of CoverTab[79225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1199
		_go_fuzz_dep_.CoverTab[79226]++
												ac.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1201
		// _ = "end of CoverTab[79226]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1202
		_go_fuzz_dep_.CoverTab[79234]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1202
		// _ = "end of CoverTab[79234]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1202
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1202
	// _ = "end of CoverTab[79215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1202
	_go_fuzz_dep_.CoverTab[79216]++

											ac.mu.Lock()
											ac.backoffIdx = 0
											ac.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1206
	// _ = "end of CoverTab[79216]"
}

// tryAllAddrs tries to creates a connection to the addresses, and stop when at
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1209
// the first successful one. It returns an error if no address was successfully
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1209
// connected, or updates ac appropriately with the new transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1212
func (ac *addrConn) tryAllAddrs(addrs []resolver.Address, connectDeadline time.Time) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1212
	_go_fuzz_dep_.CoverTab[79235]++
											var firstConnErr error
											for _, addr := range addrs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1214
		_go_fuzz_dep_.CoverTab[79237]++
												ac.mu.Lock()
												if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1216
			_go_fuzz_dep_.CoverTab[79242]++
													ac.mu.Unlock()
													return errConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1218
			// _ = "end of CoverTab[79242]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1219
			_go_fuzz_dep_.CoverTab[79243]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1219
			// _ = "end of CoverTab[79243]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1219
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1219
		// _ = "end of CoverTab[79237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1219
		_go_fuzz_dep_.CoverTab[79238]++

												ac.cc.mu.RLock()
												ac.dopts.copts.KeepaliveParams = ac.cc.mkp
												ac.cc.mu.RUnlock()

												copts := ac.dopts.copts
												if ac.scopts.CredsBundle != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1226
			_go_fuzz_dep_.CoverTab[79244]++
													copts.CredsBundle = ac.scopts.CredsBundle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1227
			// _ = "end of CoverTab[79244]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1228
			_go_fuzz_dep_.CoverTab[79245]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1228
			// _ = "end of CoverTab[79245]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1228
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1228
		// _ = "end of CoverTab[79238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1228
		_go_fuzz_dep_.CoverTab[79239]++
												ac.mu.Unlock()

												channelz.Infof(logger, ac.channelzID, "Subchannel picks a new address %q to connect", addr.Addr)

												err := ac.createTransport(addr, copts, connectDeadline)
												if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1234
			_go_fuzz_dep_.CoverTab[79246]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1235
			// _ = "end of CoverTab[79246]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1236
			_go_fuzz_dep_.CoverTab[79247]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1236
			// _ = "end of CoverTab[79247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1236
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1236
		// _ = "end of CoverTab[79239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1236
		_go_fuzz_dep_.CoverTab[79240]++
												if firstConnErr == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1237
			_go_fuzz_dep_.CoverTab[79248]++
													firstConnErr = err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1238
			// _ = "end of CoverTab[79248]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1239
			_go_fuzz_dep_.CoverTab[79249]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1239
			// _ = "end of CoverTab[79249]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1239
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1239
		// _ = "end of CoverTab[79240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1239
		_go_fuzz_dep_.CoverTab[79241]++
												ac.cc.updateConnectionError(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1240
		// _ = "end of CoverTab[79241]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1241
	// _ = "end of CoverTab[79235]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1241
	_go_fuzz_dep_.CoverTab[79236]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1244
	return firstConnErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1244
	// _ = "end of CoverTab[79236]"
}

// createTransport creates a connection to addr. It returns an error if the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1247
// address was not successfully connected, or updates ac appropriately with the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1247
// new transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1250
func (ac *addrConn) createTransport(addr resolver.Address, copts transport.ConnectOptions, connectDeadline time.Time) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1250
	_go_fuzz_dep_.CoverTab[79250]++
											addr.ServerName = ac.cc.getServerName(addr)
											hctx, hcancel := context.WithCancel(ac.ctx)

											onClose := func(r transport.GoAwayReason) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1254
		_go_fuzz_dep_.CoverTab[79255]++
												ac.mu.Lock()
												defer ac.mu.Unlock()

												ac.adjustParams(r)
												if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1259
			_go_fuzz_dep_.CoverTab[79258]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1263
			return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1263
			// _ = "end of CoverTab[79258]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1264
			_go_fuzz_dep_.CoverTab[79259]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1264
			// _ = "end of CoverTab[79259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1264
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1264
		// _ = "end of CoverTab[79255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1264
		_go_fuzz_dep_.CoverTab[79256]++
												hcancel()
												if ac.transport == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1266
			_go_fuzz_dep_.CoverTab[79260]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1271
			return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1271
			// _ = "end of CoverTab[79260]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1272
			_go_fuzz_dep_.CoverTab[79261]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1272
			// _ = "end of CoverTab[79261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1272
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1272
		// _ = "end of CoverTab[79256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1272
		_go_fuzz_dep_.CoverTab[79257]++
												ac.transport = nil

												ac.cc.resolveNow(resolver.ResolveNowOptions{})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1278
		ac.updateConnectivityState(connectivity.Idle, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1278
		// _ = "end of CoverTab[79257]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1279
	// _ = "end of CoverTab[79250]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1279
	_go_fuzz_dep_.CoverTab[79251]++

											connectCtx, cancel := context.WithDeadline(ac.ctx, connectDeadline)
											defer cancel()
											copts.ChannelzParentID = ac.channelzID

											newTr, err := transport.NewClientTransport(connectCtx, ac.cc.ctx, addr, copts, onClose)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1286
		_go_fuzz_dep_.CoverTab[79262]++
												if logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1287
			_go_fuzz_dep_.CoverTab[79264]++
													logger.Infof("Creating new client transport to %q: %v", addr, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1288
			// _ = "end of CoverTab[79264]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1289
			_go_fuzz_dep_.CoverTab[79265]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1289
			// _ = "end of CoverTab[79265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1289
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1289
		// _ = "end of CoverTab[79262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1289
		_go_fuzz_dep_.CoverTab[79263]++

												hcancel()
												channelz.Warningf(logger, ac.channelzID, "grpc: addrConn.createTransport failed to connect to %s. Err: %v", addr, err)
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1293
		// _ = "end of CoverTab[79263]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1294
		_go_fuzz_dep_.CoverTab[79266]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1294
		// _ = "end of CoverTab[79266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1294
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1294
	// _ = "end of CoverTab[79251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1294
	_go_fuzz_dep_.CoverTab[79252]++

											ac.mu.Lock()
											defer ac.mu.Unlock()
											if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1298
		_go_fuzz_dep_.CoverTab[79267]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1298
		_curRoutineNum90_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1298
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum90_)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1310
		go newTr.Close(transport.ErrConnClosing)
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1311
		// _ = "end of CoverTab[79267]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1312
		_go_fuzz_dep_.CoverTab[79268]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1312
		// _ = "end of CoverTab[79268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1312
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1312
	// _ = "end of CoverTab[79252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1312
	_go_fuzz_dep_.CoverTab[79253]++
											if hctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1313
		_go_fuzz_dep_.CoverTab[79269]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1317
		ac.updateConnectivityState(connectivity.Idle, nil)
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1318
		// _ = "end of CoverTab[79269]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1319
		_go_fuzz_dep_.CoverTab[79270]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1319
		// _ = "end of CoverTab[79270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1319
	// _ = "end of CoverTab[79253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1319
	_go_fuzz_dep_.CoverTab[79254]++
											ac.curAddr = addr
											ac.transport = newTr
											ac.startHealthCheck(hctx)
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1323
	// _ = "end of CoverTab[79254]"
}

// startHealthCheck starts the health checking stream (RPC) to watch the health
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// stats of this connection if health checking is requested and configured.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// LB channel health checking is enabled when all requirements below are met:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// 1. it is not disabled by the user with the WithDisableHealthCheck DialOption
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// 2. internal.HealthCheckFunc is set by importing the grpc/health package
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// 3. a service config with non-empty healthCheckConfig field is provided
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// 4. the load balancer requests it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// It sets addrConn to READY if the health checking stream is not started.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1326
// Caller must hold ac.mu.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1338
func (ac *addrConn) startHealthCheck(ctx context.Context) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1338
	_go_fuzz_dep_.CoverTab[79271]++
											var healthcheckManagingState bool
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1340
		_go_fuzz_dep_.CoverTab[79279]++
												if !healthcheckManagingState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1341
			_go_fuzz_dep_.CoverTab[79280]++
													ac.updateConnectivityState(connectivity.Ready, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1342
			// _ = "end of CoverTab[79280]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1343
			_go_fuzz_dep_.CoverTab[79281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1343
			// _ = "end of CoverTab[79281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1343
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1343
		// _ = "end of CoverTab[79279]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1344
	// _ = "end of CoverTab[79271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1344
	_go_fuzz_dep_.CoverTab[79272]++

											if ac.cc.dopts.disableHealthCheck {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1346
		_go_fuzz_dep_.CoverTab[79282]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1347
		// _ = "end of CoverTab[79282]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1348
		_go_fuzz_dep_.CoverTab[79283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1348
		// _ = "end of CoverTab[79283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1348
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1348
	// _ = "end of CoverTab[79272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1348
	_go_fuzz_dep_.CoverTab[79273]++
											healthCheckConfig := ac.cc.healthCheckConfig()
											if healthCheckConfig == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1350
		_go_fuzz_dep_.CoverTab[79284]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1351
		// _ = "end of CoverTab[79284]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1352
		_go_fuzz_dep_.CoverTab[79285]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1352
		// _ = "end of CoverTab[79285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1352
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1352
	// _ = "end of CoverTab[79273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1352
	_go_fuzz_dep_.CoverTab[79274]++
											if !ac.scopts.HealthCheckEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1353
		_go_fuzz_dep_.CoverTab[79286]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1354
		// _ = "end of CoverTab[79286]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1355
		_go_fuzz_dep_.CoverTab[79287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1355
		// _ = "end of CoverTab[79287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1355
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1355
	// _ = "end of CoverTab[79274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1355
	_go_fuzz_dep_.CoverTab[79275]++
											healthCheckFunc := ac.cc.dopts.healthCheckFunc
											if healthCheckFunc == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1357
		_go_fuzz_dep_.CoverTab[79288]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1361
		channelz.Error(logger, ac.channelzID, "Health check is requested but health check function is not set.")
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1362
		// _ = "end of CoverTab[79288]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1363
		_go_fuzz_dep_.CoverTab[79289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1363
		// _ = "end of CoverTab[79289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1363
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1363
	// _ = "end of CoverTab[79275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1363
	_go_fuzz_dep_.CoverTab[79276]++

											healthcheckManagingState = true

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1368
	currentTr := ac.transport
	newStream := func(method string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1369
		_go_fuzz_dep_.CoverTab[79290]++
												ac.mu.Lock()
												if ac.transport != currentTr {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1371
			_go_fuzz_dep_.CoverTab[79292]++
													ac.mu.Unlock()
													return nil, status.Error(codes.Canceled, "the provided transport is no longer valid to use")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1373
			// _ = "end of CoverTab[79292]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1374
			_go_fuzz_dep_.CoverTab[79293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1374
			// _ = "end of CoverTab[79293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1374
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1374
		// _ = "end of CoverTab[79290]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1374
		_go_fuzz_dep_.CoverTab[79291]++
												ac.mu.Unlock()
												return newNonRetryClientStream(ctx, &StreamDesc{ServerStreams: true}, method, currentTr, ac)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1376
		// _ = "end of CoverTab[79291]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1377
	// _ = "end of CoverTab[79276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1377
	_go_fuzz_dep_.CoverTab[79277]++
											setConnectivityState := func(s connectivity.State, lastErr error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1378
		_go_fuzz_dep_.CoverTab[79294]++
												ac.mu.Lock()
												defer ac.mu.Unlock()
												if ac.transport != currentTr {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1381
			_go_fuzz_dep_.CoverTab[79296]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1382
			// _ = "end of CoverTab[79296]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1383
			_go_fuzz_dep_.CoverTab[79297]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1383
			// _ = "end of CoverTab[79297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1383
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1383
		// _ = "end of CoverTab[79294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1383
		_go_fuzz_dep_.CoverTab[79295]++
												ac.updateConnectivityState(s, lastErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1384
		// _ = "end of CoverTab[79295]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1385
	// _ = "end of CoverTab[79277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1385
	_go_fuzz_dep_.CoverTab[79278]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1385
	_curRoutineNum91_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1385
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum91_)

											go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
		_go_fuzz_dep_.CoverTab[79298]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
			_go_fuzz_dep_.CoverTab[79299]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum91_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
			// _ = "end of CoverTab[79299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1387
		}()
												err := ac.cc.dopts.healthCheckFunc(ctx, newStream, setConnectivityState, healthCheckConfig.ServiceName)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1389
			_go_fuzz_dep_.CoverTab[79300]++
													if status.Code(err) == codes.Unimplemented {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1390
				_go_fuzz_dep_.CoverTab[79301]++
														channelz.Error(logger, ac.channelzID, "Subchannel health check is unimplemented at server side, thus health check is disabled")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1391
				// _ = "end of CoverTab[79301]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1392
				_go_fuzz_dep_.CoverTab[79302]++
														channelz.Errorf(logger, ac.channelzID, "Health checking failed: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1393
				// _ = "end of CoverTab[79302]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1394
			// _ = "end of CoverTab[79300]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1395
			_go_fuzz_dep_.CoverTab[79303]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1395
			// _ = "end of CoverTab[79303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1395
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1395
		// _ = "end of CoverTab[79298]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1396
	// _ = "end of CoverTab[79278]"
}

func (ac *addrConn) resetConnectBackoff() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1399
	_go_fuzz_dep_.CoverTab[79304]++
											ac.mu.Lock()
											close(ac.resetBackoff)
											ac.backoffIdx = 0
											ac.resetBackoff = make(chan struct{})
											ac.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1404
	// _ = "end of CoverTab[79304]"
}

// getReadyTransport returns the transport if ac's state is READY or nil if not.
func (ac *addrConn) getReadyTransport() transport.ClientTransport {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1408
	_go_fuzz_dep_.CoverTab[79305]++
											ac.mu.Lock()
											defer ac.mu.Unlock()
											if ac.state == connectivity.Ready {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1411
		_go_fuzz_dep_.CoverTab[79307]++
												return ac.transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1412
		// _ = "end of CoverTab[79307]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1413
		_go_fuzz_dep_.CoverTab[79308]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1413
		// _ = "end of CoverTab[79308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1413
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1413
	// _ = "end of CoverTab[79305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1413
	_go_fuzz_dep_.CoverTab[79306]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1414
	// _ = "end of CoverTab[79306]"
}

// tearDown starts to tear down the addrConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1417
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1417
// Note that tearDown doesn't remove ac from ac.cc.conns, so the addrConn struct
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1417
// will leak. In most cases, call cc.removeAddrConn() instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1421
func (ac *addrConn) tearDown(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1421
	_go_fuzz_dep_.CoverTab[79309]++
											ac.mu.Lock()
											if ac.state == connectivity.Shutdown {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1423
		_go_fuzz_dep_.CoverTab[79312]++
												ac.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1425
		// _ = "end of CoverTab[79312]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1426
		_go_fuzz_dep_.CoverTab[79313]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1426
		// _ = "end of CoverTab[79313]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1426
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1426
	// _ = "end of CoverTab[79309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1426
	_go_fuzz_dep_.CoverTab[79310]++
											curTr := ac.transport
											ac.transport = nil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1431
	ac.updateConnectivityState(connectivity.Shutdown, nil)
	ac.cancel()
	ac.curAddr = resolver.Address{}
	if err == errConnDrain && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1434
		_go_fuzz_dep_.CoverTab[79314]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1434
		return curTr != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1434
		// _ = "end of CoverTab[79314]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1434
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1434
		_go_fuzz_dep_.CoverTab[79315]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1440
		ac.mu.Unlock()
												curTr.GracefulClose()
												ac.mu.Lock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1442
		// _ = "end of CoverTab[79315]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1443
		_go_fuzz_dep_.CoverTab[79316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1443
		// _ = "end of CoverTab[79316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1443
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1443
	// _ = "end of CoverTab[79310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1443
	_go_fuzz_dep_.CoverTab[79311]++
											channelz.AddTraceEvent(logger, ac.channelzID, 0, &channelz.TraceEventDesc{
		Desc:		"Subchannel deleted",
		Severity:	channelz.CtInfo,
		Parent: &channelz.TraceEventDesc{
			Desc:		fmt.Sprintf("Subchannel(id:%d) deleted", ac.channelzID.Int()),
			Severity:	channelz.CtInfo,
		},
	})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1455
	channelz.RemoveEntry(ac.channelzID)
											ac.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1456
	// _ = "end of CoverTab[79311]"
}

func (ac *addrConn) getState() connectivity.State {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1459
	_go_fuzz_dep_.CoverTab[79317]++
											ac.mu.Lock()
											defer ac.mu.Unlock()
											return ac.state
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1462
	// _ = "end of CoverTab[79317]"
}

func (ac *addrConn) ChannelzMetric() *channelz.ChannelInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1465
	_go_fuzz_dep_.CoverTab[79318]++
											ac.mu.Lock()
											addr := ac.curAddr.Addr
											ac.mu.Unlock()
											return &channelz.ChannelInternalMetric{
		State:				ac.getState(),
		Target:				addr,
		CallsStarted:			atomic.LoadInt64(&ac.czData.callsStarted),
		CallsSucceeded:			atomic.LoadInt64(&ac.czData.callsSucceeded),
		CallsFailed:			atomic.LoadInt64(&ac.czData.callsFailed),
		LastCallStartedTimestamp:	time.Unix(0, atomic.LoadInt64(&ac.czData.lastCallStartedTime)),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1476
	// _ = "end of CoverTab[79318]"
}

func (ac *addrConn) incrCallsStarted() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1479
	_go_fuzz_dep_.CoverTab[79319]++
											atomic.AddInt64(&ac.czData.callsStarted, 1)
											atomic.StoreInt64(&ac.czData.lastCallStartedTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1481
	// _ = "end of CoverTab[79319]"
}

func (ac *addrConn) incrCallsSucceeded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1484
	_go_fuzz_dep_.CoverTab[79320]++
											atomic.AddInt64(&ac.czData.callsSucceeded, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1485
	// _ = "end of CoverTab[79320]"
}

func (ac *addrConn) incrCallsFailed() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1488
	_go_fuzz_dep_.CoverTab[79321]++
											atomic.AddInt64(&ac.czData.callsFailed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1489
	// _ = "end of CoverTab[79321]"
}

type retryThrottler struct {
	max	float64
	thresh	float64
	ratio	float64

	mu	sync.Mutex
	tokens	float64	// TODO(dfawley): replace with atomic and remove lock.
}

// throttle subtracts a retry token from the pool and returns whether a retry
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1501
// should be throttled (disallowed) based upon the retry throttling policy in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1501
// the service config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1504
func (rt *retryThrottler) throttle() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1504
	_go_fuzz_dep_.CoverTab[79322]++
											if rt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1505
		_go_fuzz_dep_.CoverTab[79325]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1506
		// _ = "end of CoverTab[79325]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1507
		_go_fuzz_dep_.CoverTab[79326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1507
		// _ = "end of CoverTab[79326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1507
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1507
	// _ = "end of CoverTab[79322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1507
	_go_fuzz_dep_.CoverTab[79323]++
											rt.mu.Lock()
											defer rt.mu.Unlock()
											rt.tokens--
											if rt.tokens < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1511
		_go_fuzz_dep_.CoverTab[79327]++
												rt.tokens = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1512
		// _ = "end of CoverTab[79327]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1513
		_go_fuzz_dep_.CoverTab[79328]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1513
		// _ = "end of CoverTab[79328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1513
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1513
	// _ = "end of CoverTab[79323]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1513
	_go_fuzz_dep_.CoverTab[79324]++
											return rt.tokens <= rt.thresh
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1514
	// _ = "end of CoverTab[79324]"
}

func (rt *retryThrottler) successfulRPC() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1517
	_go_fuzz_dep_.CoverTab[79329]++
											if rt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1518
		_go_fuzz_dep_.CoverTab[79331]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1519
		// _ = "end of CoverTab[79331]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1520
		_go_fuzz_dep_.CoverTab[79332]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1520
		// _ = "end of CoverTab[79332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1520
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1520
	// _ = "end of CoverTab[79329]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1520
	_go_fuzz_dep_.CoverTab[79330]++
											rt.mu.Lock()
											defer rt.mu.Unlock()
											rt.tokens += rt.ratio
											if rt.tokens > rt.max {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1524
		_go_fuzz_dep_.CoverTab[79333]++
												rt.tokens = rt.max
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1525
		// _ = "end of CoverTab[79333]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1526
		_go_fuzz_dep_.CoverTab[79334]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1526
		// _ = "end of CoverTab[79334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1526
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1526
	// _ = "end of CoverTab[79330]"
}

type channelzChannel struct {
	cc *ClientConn
}

func (c *channelzChannel) ChannelzMetric() *channelz.ChannelInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1533
	_go_fuzz_dep_.CoverTab[79335]++
											return c.cc.channelzMetric()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1534
	// _ = "end of CoverTab[79335]"
}

// ErrClientConnTimeout indicates that the ClientConn cannot establish the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1537
// underlying connections within the specified timeout.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1537
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1537
// Deprecated: This error is never returned by grpc and should not be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1537
// referenced by users.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1542
var ErrClientConnTimeout = errors.New("grpc: timed out when dialing")

// getResolver finds the scheme in the cc's resolvers or the global registry.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1544
// scheme should always be lowercase (typically by virtue of url.Parse()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1544
// performing proper RFC3986 behavior).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1547
func (cc *ClientConn) getResolver(scheme string) resolver.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1547
	_go_fuzz_dep_.CoverTab[79336]++
											for _, rb := range cc.dopts.resolvers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1548
		_go_fuzz_dep_.CoverTab[79338]++
												if scheme == rb.Scheme() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1549
			_go_fuzz_dep_.CoverTab[79339]++
													return rb
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1550
			// _ = "end of CoverTab[79339]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1551
			_go_fuzz_dep_.CoverTab[79340]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1551
			// _ = "end of CoverTab[79340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1551
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1551
		// _ = "end of CoverTab[79338]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1552
	// _ = "end of CoverTab[79336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1552
	_go_fuzz_dep_.CoverTab[79337]++
											return resolver.Get(scheme)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1553
	// _ = "end of CoverTab[79337]"
}

func (cc *ClientConn) updateConnectionError(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1556
	_go_fuzz_dep_.CoverTab[79341]++
											cc.lceMu.Lock()
											cc.lastConnectionError = err
											cc.lceMu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1559
	// _ = "end of CoverTab[79341]"
}

func (cc *ClientConn) connectionError() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1562
	_go_fuzz_dep_.CoverTab[79342]++
											cc.lceMu.Lock()
											defer cc.lceMu.Unlock()
											return cc.lastConnectionError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1565
	// _ = "end of CoverTab[79342]"
}

func (cc *ClientConn) parseTargetAndFindResolver() (resolver.Builder, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1568
	_go_fuzz_dep_.CoverTab[79343]++
											channelz.Infof(logger, cc.channelzID, "original dial target is: %q", cc.target)

											var rb resolver.Builder
											parsedTarget, err := parseTarget(cc.target)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1573
		_go_fuzz_dep_.CoverTab[79347]++
												channelz.Infof(logger, cc.channelzID, "dial target %q parse failed: %v", cc.target, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1574
		// _ = "end of CoverTab[79347]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1575
		_go_fuzz_dep_.CoverTab[79348]++
												channelz.Infof(logger, cc.channelzID, "parsed dial target is: %+v", parsedTarget)
												rb = cc.getResolver(parsedTarget.URL.Scheme)
												if rb != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1578
			_go_fuzz_dep_.CoverTab[79349]++
													cc.parsedTarget = parsedTarget
													return rb, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1580
			// _ = "end of CoverTab[79349]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1581
			_go_fuzz_dep_.CoverTab[79350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1581
			// _ = "end of CoverTab[79350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1581
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1581
		// _ = "end of CoverTab[79348]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1582
	// _ = "end of CoverTab[79343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1582
	_go_fuzz_dep_.CoverTab[79344]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1588
	defScheme := resolver.GetDefaultScheme()
	channelz.Infof(logger, cc.channelzID, "fallback to scheme %q", defScheme)
	canonicalTarget := defScheme + ":///" + cc.target

	parsedTarget, err = parseTarget(canonicalTarget)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1593
		_go_fuzz_dep_.CoverTab[79351]++
												channelz.Infof(logger, cc.channelzID, "dial target %q parse failed: %v", canonicalTarget, err)
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1595
		// _ = "end of CoverTab[79351]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1596
		_go_fuzz_dep_.CoverTab[79352]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1596
		// _ = "end of CoverTab[79352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1596
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1596
	// _ = "end of CoverTab[79344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1596
	_go_fuzz_dep_.CoverTab[79345]++
											channelz.Infof(logger, cc.channelzID, "parsed dial target is: %+v", parsedTarget)
											rb = cc.getResolver(parsedTarget.URL.Scheme)
											if rb == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1599
		_go_fuzz_dep_.CoverTab[79353]++
												return nil, fmt.Errorf("could not get resolver for default scheme: %q", parsedTarget.URL.Scheme)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1600
		// _ = "end of CoverTab[79353]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1601
		_go_fuzz_dep_.CoverTab[79354]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1601
		// _ = "end of CoverTab[79354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1601
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1601
	// _ = "end of CoverTab[79345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1601
	_go_fuzz_dep_.CoverTab[79346]++
											cc.parsedTarget = parsedTarget
											return rb, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1603
	// _ = "end of CoverTab[79346]"
}

// parseTarget uses RFC 3986 semantics to parse the given target into a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1606
// resolver.Target struct containing scheme, authority and url. Query
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1606
// params are stripped from the endpoint.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1609
func parseTarget(target string) (resolver.Target, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1609
	_go_fuzz_dep_.CoverTab[79355]++
											u, err := url.Parse(target)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1611
		_go_fuzz_dep_.CoverTab[79357]++
												return resolver.Target{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1612
		// _ = "end of CoverTab[79357]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1613
		_go_fuzz_dep_.CoverTab[79358]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1613
		// _ = "end of CoverTab[79358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1613
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1613
	// _ = "end of CoverTab[79355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1613
	_go_fuzz_dep_.CoverTab[79356]++

											return resolver.Target{
		Scheme:		u.Scheme,
		Authority:	u.Host,
		URL:		*u,
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1619
	// _ = "end of CoverTab[79356]"
}

// Determine channel authority. The order of precedence is as follows:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1622
// - user specified authority override using `WithAuthority` dial option
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1622
// - creds' notion of server name for the authentication handshake
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1622
// - endpoint from dial target of the form "scheme://[authority]/endpoint"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1626
func determineAuthority(endpoint, target string, dopts dialOptions) (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1626
	_go_fuzz_dep_.CoverTab[79359]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1637
	authorityFromCreds := ""
	if creds := dopts.copts.TransportCredentials; creds != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1638
		_go_fuzz_dep_.CoverTab[79362]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1638
		return creds.Info().ServerName != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1638
		// _ = "end of CoverTab[79362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1638
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1638
		_go_fuzz_dep_.CoverTab[79363]++
												authorityFromCreds = creds.Info().ServerName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1639
		// _ = "end of CoverTab[79363]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1640
		_go_fuzz_dep_.CoverTab[79364]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1640
		// _ = "end of CoverTab[79364]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1640
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1640
	// _ = "end of CoverTab[79359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1640
	_go_fuzz_dep_.CoverTab[79360]++
											authorityFromDialOption := dopts.authority
											if (authorityFromCreds != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		_go_fuzz_dep_.CoverTab[79365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		return authorityFromDialOption != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		// _ = "end of CoverTab[79365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		_go_fuzz_dep_.CoverTab[79366]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		return authorityFromCreds != authorityFromDialOption
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		// _ = "end of CoverTab[79366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1642
		_go_fuzz_dep_.CoverTab[79367]++
												return "", fmt.Errorf("ClientConn's authority from transport creds %q and dial option %q don't match", authorityFromCreds, authorityFromDialOption)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1643
		// _ = "end of CoverTab[79367]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1644
		_go_fuzz_dep_.CoverTab[79368]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1644
		// _ = "end of CoverTab[79368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1644
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1644
	// _ = "end of CoverTab[79360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1644
	_go_fuzz_dep_.CoverTab[79361]++

											switch {
	case authorityFromDialOption != "":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1647
		_go_fuzz_dep_.CoverTab[79369]++
												return authorityFromDialOption, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1648
		// _ = "end of CoverTab[79369]"
	case authorityFromCreds != "":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1649
		_go_fuzz_dep_.CoverTab[79370]++
												return authorityFromCreds, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1650
		// _ = "end of CoverTab[79370]"
	case strings.HasPrefix(target, "unix:") || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1651
		_go_fuzz_dep_.CoverTab[79374]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1651
		return strings.HasPrefix(target, "unix-abstract:")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1651
		// _ = "end of CoverTab[79374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1651
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1651
		_go_fuzz_dep_.CoverTab[79371]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1654
		return "localhost", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1654
		// _ = "end of CoverTab[79371]"
	case strings.HasPrefix(endpoint, ":"):
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1655
		_go_fuzz_dep_.CoverTab[79372]++
												return "localhost" + endpoint, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1656
		// _ = "end of CoverTab[79372]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1657
		_go_fuzz_dep_.CoverTab[79373]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1662
		return endpoint, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1662
		// _ = "end of CoverTab[79373]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1663
	// _ = "end of CoverTab[79361]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1664
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/clientconn.go:1664
var _ = _go_fuzz_dep_.CoverTab
