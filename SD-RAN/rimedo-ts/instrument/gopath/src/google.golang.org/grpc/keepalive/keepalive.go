//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:19
// Package keepalive defines configurable parameters for point-to-point
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:19
// healthcheck.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
package keepalive

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:21
)

import (
	"time"
)

// ClientParameters is used to set keepalive parameters on the client-side.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:27
// These configure how the client will actively probe to notice when a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:27
// connection is broken and send pings so intermediaries will be aware of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:27
// liveness of the connection. Make sure these parameters are set in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:27
// coordination with the keepalive policy on the server, as incompatible
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:27
// settings can result in closing of connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:33
type ClientParameters struct {
	// After a duration of this time if the client doesn't see any activity it
	// pings the server to see if the transport is still alive.
	// If set below 10s, a minimum value of 10s will be used instead.
	Time	time.Duration	// The current default value is infinity.
	// After having pinged for keepalive check, the client waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout	time.Duration	// The current default value is 20 seconds.
	// If true, client sends keepalive pings even with no active RPCs. If false,
	// when there are no active RPCs, Time and Timeout will be ignored and no
	// keepalive pings will be sent.
	PermitWithoutStream	bool	// false by default.
}

// ServerParameters is used to set keepalive and max-age parameters on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:48
// server-side.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:50
type ServerParameters struct {
	// MaxConnectionIdle is a duration for the amount of time after which an
	// idle connection would be closed by sending a GoAway. Idleness duration is
	// defined since the most recent time the number of outstanding RPCs became
	// zero or the connection establishment.
	MaxConnectionIdle	time.Duration	// The current default value is infinity.
	// MaxConnectionAge is a duration for the maximum amount of time a
	// connection may exist before it will be closed by sending a GoAway. A
	// random jitter of +/-10% will be added to MaxConnectionAge to spread out
	// connection storms.
	MaxConnectionAge	time.Duration	// The current default value is infinity.
	// MaxConnectionAgeGrace is an additive period after MaxConnectionAge after
	// which the connection will be forcibly closed.
	MaxConnectionAgeGrace	time.Duration	// The current default value is infinity.
	// After a duration of this time if the server doesn't see any activity it
	// pings the client to see if the transport is still alive.
	// If set below 1s, a minimum value of 1s will be used instead.
	Time	time.Duration	// The current default value is 2 hours.
	// After having pinged for keepalive check, the server waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout	time.Duration	// The current default value is 20 seconds.
}

// EnforcementPolicy is used to set keepalive enforcement policy on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:74
// server-side. Server will close connection with a client that violates this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:74
// policy.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:77
type EnforcementPolicy struct {
	// MinTime is the minimum amount of time a client should wait before sending
	// a keepalive ping.
	MinTime	time.Duration	// The current default value is 5 minutes.
	// If true, server allows keepalive pings even when there are no active
	// streams(RPCs). If false, and client sends ping when there are no active
	// streams, server will send GOAWAY and close the connection.
	PermitWithoutStream	bool	// false by default.
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/keepalive/keepalive.go:85
var _ = _go_fuzz_dep_.CoverTab
