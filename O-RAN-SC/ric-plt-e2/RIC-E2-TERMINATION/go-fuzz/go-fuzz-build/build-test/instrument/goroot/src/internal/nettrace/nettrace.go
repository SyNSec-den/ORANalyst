// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/nettrace/nettrace.go:5
// Package nettrace contains internal hooks for tracing activity in
//line /usr/local/go/src/internal/nettrace/nettrace.go:5
// the net package. This package is purely internal for use by the
//line /usr/local/go/src/internal/nettrace/nettrace.go:5
// net/http/httptrace package and has no stable API exposed to end
//line /usr/local/go/src/internal/nettrace/nettrace.go:5
// users.
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
package nettrace

//line /usr/local/go/src/internal/nettrace/nettrace.go:9
import (
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
)
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
import (
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/nettrace/nettrace.go:9
)

// TraceKey is a context.Context Value key. Its associated value should
//line /usr/local/go/src/internal/nettrace/nettrace.go:11
// be a *Trace struct.
//line /usr/local/go/src/internal/nettrace/nettrace.go:13
type TraceKey struct{}

// LookupIPAltResolverKey is a context.Context Value key used by tests to
//line /usr/local/go/src/internal/nettrace/nettrace.go:15
// specify an alternate resolver func.
//line /usr/local/go/src/internal/nettrace/nettrace.go:15
// It is not exposed to outsider users. (But see issue 12503)
//line /usr/local/go/src/internal/nettrace/nettrace.go:15
// The value should be the same type as lookupIP:
//line /usr/local/go/src/internal/nettrace/nettrace.go:15
//
//line /usr/local/go/src/internal/nettrace/nettrace.go:15
//	func lookupIP(ctx context.Context, host string) ([]IPAddr, error)
//line /usr/local/go/src/internal/nettrace/nettrace.go:21
type LookupIPAltResolverKey struct{}

// Trace contains a set of hooks for tracing events within
//line /usr/local/go/src/internal/nettrace/nettrace.go:23
// the net package. Any specific hook may be nil.
//line /usr/local/go/src/internal/nettrace/nettrace.go:25
type Trace struct {
	// DNSStart is called with the hostname of a DNS lookup
	// before it begins.
	DNSStart	func(name string)

	// DNSDone is called after a DNS lookup completes (or fails).
	// The coalesced parameter is whether singleflight de-duped
	// the call. The addrs are of type net.IPAddr but can't
	// actually be for circular dependency reasons.
	DNSDone	func(netIPs []any, coalesced bool, err error)

	// ConnectStart is called before a Dial, excluding Dials made
	// during DNS lookups. In the case of DualStack (Happy Eyeballs)
	// dialing, this may be called multiple times, from multiple
	// goroutines.
	ConnectStart	func(network, addr string)

	// ConnectStart is called after a Dial with the results, excluding
	// Dials made during DNS lookups. It may also be called multiple
	// times, like ConnectStart.
	ConnectDone	func(network, addr string, err error)
}

//line /usr/local/go/src/internal/nettrace/nettrace.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/nettrace/nettrace.go:46
var _ = _go_fuzz_dep_.CoverTab
