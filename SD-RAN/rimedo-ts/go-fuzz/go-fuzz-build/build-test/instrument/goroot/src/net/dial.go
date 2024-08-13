// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/dial.go:5
package net

//line /usr/local/go/src/net/dial.go:5
import (
//line /usr/local/go/src/net/dial.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/dial.go:5
)
//line /usr/local/go/src/net/dial.go:5
import (
//line /usr/local/go/src/net/dial.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/dial.go:5
)

import (
	"context"
	"internal/nettrace"
	"syscall"
	"time"
)

// defaultTCPKeepAlive is a default constant value for TCPKeepAlive times
//line /usr/local/go/src/net/dial.go:14
// See golang.org/issue/31510
//line /usr/local/go/src/net/dial.go:16
const (
	defaultTCPKeepAlive = 15 * time.Second
)

// A Dialer contains options for connecting to an address.
//line /usr/local/go/src/net/dial.go:20
//
//line /usr/local/go/src/net/dial.go:20
// The zero value for each field is equivalent to dialing
//line /usr/local/go/src/net/dial.go:20
// without that option. Dialing with the zero value of Dialer
//line /usr/local/go/src/net/dial.go:20
// is therefore equivalent to just calling the Dial function.
//line /usr/local/go/src/net/dial.go:20
//
//line /usr/local/go/src/net/dial.go:20
// It is safe to call Dialer's methods concurrently.
//line /usr/local/go/src/net/dial.go:27
type Dialer struct {
	// Timeout is the maximum amount of time a dial will wait for
	// a connect to complete. If Deadline is also set, it may fail
	// earlier.
	//
	// The default is no timeout.
	//
	// When using TCP and dialing a host name with multiple IP
	// addresses, the timeout may be divided between them.
	//
	// With or without a timeout, the operating system may impose
	// its own earlier timeout. For instance, TCP timeouts are
	// often around 3 minutes.
	Timeout	time.Duration

	// Deadline is the absolute point in time after which dials
	// will fail. If Timeout is set, it may fail earlier.
	// Zero means no deadline, or dependent on the operating system
	// as with the Timeout option.
	Deadline	time.Time

	// LocalAddr is the local address to use when dialing an
	// address. The address must be of a compatible type for the
	// network being dialed.
	// If nil, a local address is automatically chosen.
	LocalAddr	Addr

	// DualStack previously enabled RFC 6555 Fast Fallback
	// support, also known as "Happy Eyeballs", in which IPv4 is
	// tried soon if IPv6 appears to be misconfigured and
	// hanging.
	//
	// Deprecated: Fast Fallback is enabled by default. To
	// disable, set FallbackDelay to a negative value.
	DualStack	bool

	// FallbackDelay specifies the length of time to wait before
	// spawning a RFC 6555 Fast Fallback connection. That is, this
	// is the amount of time to wait for IPv6 to succeed before
	// assuming that IPv6 is misconfigured and falling back to
	// IPv4.
	//
	// If zero, a default delay of 300ms is used.
	// A negative value disables Fast Fallback support.
	FallbackDelay	time.Duration

	// KeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	// If zero, keep-alive probes are sent with a default value
	// (currently 15 seconds), if supported by the protocol and operating
	// system. Network protocols or operating systems that do
	// not support keep-alives ignore this field.
	// If negative, keep-alive probes are disabled.
	KeepAlive	time.Duration

	// Resolver optionally specifies an alternate resolver to use.
	Resolver	*Resolver

	// Cancel is an optional channel whose closure indicates that
	// the dial should be canceled. Not all types of dials support
	// cancellation.
	//
	// Deprecated: Use DialContext instead.
	Cancel	<-chan struct{}

	// If Control is not nil, it is called after creating the network
	// connection but before actually dialing.
	//
	// Network and address parameters passed to Control method are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the Control function to be called with "tcp4" or "tcp6".
	//
	// Control is ignored if ControlContext is not nil.
	Control	func(network, address string, c syscall.RawConn) error

	// If ControlContext is not nil, it is called after creating the network
	// connection but before actually dialing.
	//
	// Network and address parameters passed to Control method are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the Control function to be called with "tcp4" or "tcp6".
	//
	// If ControlContext is not nil, Control is ignored.
	ControlContext	func(ctx context.Context, network, address string, c syscall.RawConn) error
}

func (d *Dialer) dualStack() bool {
//line /usr/local/go/src/net/dial.go:113
	_go_fuzz_dep_.CoverTab[4602]++
//line /usr/local/go/src/net/dial.go:113
	return d.FallbackDelay >= 0
//line /usr/local/go/src/net/dial.go:113
	// _ = "end of CoverTab[4602]"
//line /usr/local/go/src/net/dial.go:113
}

func minNonzeroTime(a, b time.Time) time.Time {
//line /usr/local/go/src/net/dial.go:115
	_go_fuzz_dep_.CoverTab[4603]++
						if a.IsZero() {
//line /usr/local/go/src/net/dial.go:116
		_go_fuzz_dep_.CoverTab[4606]++
							return b
//line /usr/local/go/src/net/dial.go:117
		// _ = "end of CoverTab[4606]"
	} else {
//line /usr/local/go/src/net/dial.go:118
		_go_fuzz_dep_.CoverTab[4607]++
//line /usr/local/go/src/net/dial.go:118
		// _ = "end of CoverTab[4607]"
//line /usr/local/go/src/net/dial.go:118
	}
//line /usr/local/go/src/net/dial.go:118
	// _ = "end of CoverTab[4603]"
//line /usr/local/go/src/net/dial.go:118
	_go_fuzz_dep_.CoverTab[4604]++
						if b.IsZero() || func() bool {
//line /usr/local/go/src/net/dial.go:119
		_go_fuzz_dep_.CoverTab[4608]++
//line /usr/local/go/src/net/dial.go:119
		return a.Before(b)
//line /usr/local/go/src/net/dial.go:119
		// _ = "end of CoverTab[4608]"
//line /usr/local/go/src/net/dial.go:119
	}() {
//line /usr/local/go/src/net/dial.go:119
		_go_fuzz_dep_.CoverTab[4609]++
							return a
//line /usr/local/go/src/net/dial.go:120
		// _ = "end of CoverTab[4609]"
	} else {
//line /usr/local/go/src/net/dial.go:121
		_go_fuzz_dep_.CoverTab[4610]++
//line /usr/local/go/src/net/dial.go:121
		// _ = "end of CoverTab[4610]"
//line /usr/local/go/src/net/dial.go:121
	}
//line /usr/local/go/src/net/dial.go:121
	// _ = "end of CoverTab[4604]"
//line /usr/local/go/src/net/dial.go:121
	_go_fuzz_dep_.CoverTab[4605]++
						return b
//line /usr/local/go/src/net/dial.go:122
	// _ = "end of CoverTab[4605]"
}

// deadline returns the earliest of:
//line /usr/local/go/src/net/dial.go:125
//   - now+Timeout
//line /usr/local/go/src/net/dial.go:125
//   - d.Deadline
//line /usr/local/go/src/net/dial.go:125
//   - the context's deadline
//line /usr/local/go/src/net/dial.go:125
//
//line /usr/local/go/src/net/dial.go:125
// Or zero, if none of Timeout, Deadline, or context's deadline is set.
//line /usr/local/go/src/net/dial.go:131
func (d *Dialer) deadline(ctx context.Context, now time.Time) (earliest time.Time) {
//line /usr/local/go/src/net/dial.go:131
	_go_fuzz_dep_.CoverTab[4611]++
						if d.Timeout != 0 {
//line /usr/local/go/src/net/dial.go:132
		_go_fuzz_dep_.CoverTab[4614]++
							earliest = now.Add(d.Timeout)
//line /usr/local/go/src/net/dial.go:133
		// _ = "end of CoverTab[4614]"
	} else {
//line /usr/local/go/src/net/dial.go:134
		_go_fuzz_dep_.CoverTab[4615]++
//line /usr/local/go/src/net/dial.go:134
		// _ = "end of CoverTab[4615]"
//line /usr/local/go/src/net/dial.go:134
	}
//line /usr/local/go/src/net/dial.go:134
	// _ = "end of CoverTab[4611]"
//line /usr/local/go/src/net/dial.go:134
	_go_fuzz_dep_.CoverTab[4612]++
						if d, ok := ctx.Deadline(); ok {
//line /usr/local/go/src/net/dial.go:135
		_go_fuzz_dep_.CoverTab[4616]++
							earliest = minNonzeroTime(earliest, d)
//line /usr/local/go/src/net/dial.go:136
		// _ = "end of CoverTab[4616]"
	} else {
//line /usr/local/go/src/net/dial.go:137
		_go_fuzz_dep_.CoverTab[4617]++
//line /usr/local/go/src/net/dial.go:137
		// _ = "end of CoverTab[4617]"
//line /usr/local/go/src/net/dial.go:137
	}
//line /usr/local/go/src/net/dial.go:137
	// _ = "end of CoverTab[4612]"
//line /usr/local/go/src/net/dial.go:137
	_go_fuzz_dep_.CoverTab[4613]++
						return minNonzeroTime(earliest, d.Deadline)
//line /usr/local/go/src/net/dial.go:138
	// _ = "end of CoverTab[4613]"
}

func (d *Dialer) resolver() *Resolver {
//line /usr/local/go/src/net/dial.go:141
	_go_fuzz_dep_.CoverTab[4618]++
						if d.Resolver != nil {
//line /usr/local/go/src/net/dial.go:142
		_go_fuzz_dep_.CoverTab[4620]++
							return d.Resolver
//line /usr/local/go/src/net/dial.go:143
		// _ = "end of CoverTab[4620]"
	} else {
//line /usr/local/go/src/net/dial.go:144
		_go_fuzz_dep_.CoverTab[4621]++
//line /usr/local/go/src/net/dial.go:144
		// _ = "end of CoverTab[4621]"
//line /usr/local/go/src/net/dial.go:144
	}
//line /usr/local/go/src/net/dial.go:144
	// _ = "end of CoverTab[4618]"
//line /usr/local/go/src/net/dial.go:144
	_go_fuzz_dep_.CoverTab[4619]++
						return DefaultResolver
//line /usr/local/go/src/net/dial.go:145
	// _ = "end of CoverTab[4619]"
}

// partialDeadline returns the deadline to use for a single address,
//line /usr/local/go/src/net/dial.go:148
// when multiple addresses are pending.
//line /usr/local/go/src/net/dial.go:150
func partialDeadline(now, deadline time.Time, addrsRemaining int) (time.Time, error) {
//line /usr/local/go/src/net/dial.go:150
	_go_fuzz_dep_.CoverTab[4622]++
						if deadline.IsZero() {
//line /usr/local/go/src/net/dial.go:151
		_go_fuzz_dep_.CoverTab[4626]++
							return deadline, nil
//line /usr/local/go/src/net/dial.go:152
		// _ = "end of CoverTab[4626]"
	} else {
//line /usr/local/go/src/net/dial.go:153
		_go_fuzz_dep_.CoverTab[4627]++
//line /usr/local/go/src/net/dial.go:153
		// _ = "end of CoverTab[4627]"
//line /usr/local/go/src/net/dial.go:153
	}
//line /usr/local/go/src/net/dial.go:153
	// _ = "end of CoverTab[4622]"
//line /usr/local/go/src/net/dial.go:153
	_go_fuzz_dep_.CoverTab[4623]++
						timeRemaining := deadline.Sub(now)
						if timeRemaining <= 0 {
//line /usr/local/go/src/net/dial.go:155
		_go_fuzz_dep_.CoverTab[4628]++
							return time.Time{}, errTimeout
//line /usr/local/go/src/net/dial.go:156
		// _ = "end of CoverTab[4628]"
	} else {
//line /usr/local/go/src/net/dial.go:157
		_go_fuzz_dep_.CoverTab[4629]++
//line /usr/local/go/src/net/dial.go:157
		// _ = "end of CoverTab[4629]"
//line /usr/local/go/src/net/dial.go:157
	}
//line /usr/local/go/src/net/dial.go:157
	// _ = "end of CoverTab[4623]"
//line /usr/local/go/src/net/dial.go:157
	_go_fuzz_dep_.CoverTab[4624]++

						timeout := timeRemaining / time.Duration(addrsRemaining)
	// If the time per address is too short, steal from the end of the list.
	const saneMinimum = 2 * time.Second
	if timeout < saneMinimum {
//line /usr/local/go/src/net/dial.go:162
		_go_fuzz_dep_.CoverTab[4630]++
							if timeRemaining < saneMinimum {
//line /usr/local/go/src/net/dial.go:163
			_go_fuzz_dep_.CoverTab[4631]++
								timeout = timeRemaining
//line /usr/local/go/src/net/dial.go:164
			// _ = "end of CoverTab[4631]"
		} else {
//line /usr/local/go/src/net/dial.go:165
			_go_fuzz_dep_.CoverTab[4632]++
								timeout = saneMinimum
//line /usr/local/go/src/net/dial.go:166
			// _ = "end of CoverTab[4632]"
		}
//line /usr/local/go/src/net/dial.go:167
		// _ = "end of CoverTab[4630]"
	} else {
//line /usr/local/go/src/net/dial.go:168
		_go_fuzz_dep_.CoverTab[4633]++
//line /usr/local/go/src/net/dial.go:168
		// _ = "end of CoverTab[4633]"
//line /usr/local/go/src/net/dial.go:168
	}
//line /usr/local/go/src/net/dial.go:168
	// _ = "end of CoverTab[4624]"
//line /usr/local/go/src/net/dial.go:168
	_go_fuzz_dep_.CoverTab[4625]++
						return now.Add(timeout), nil
//line /usr/local/go/src/net/dial.go:169
	// _ = "end of CoverTab[4625]"
}

func (d *Dialer) fallbackDelay() time.Duration {
//line /usr/local/go/src/net/dial.go:172
	_go_fuzz_dep_.CoverTab[4634]++
						if d.FallbackDelay > 0 {
//line /usr/local/go/src/net/dial.go:173
		_go_fuzz_dep_.CoverTab[4635]++
							return d.FallbackDelay
//line /usr/local/go/src/net/dial.go:174
		// _ = "end of CoverTab[4635]"
	} else {
//line /usr/local/go/src/net/dial.go:175
		_go_fuzz_dep_.CoverTab[4636]++
							return 300 * time.Millisecond
//line /usr/local/go/src/net/dial.go:176
		// _ = "end of CoverTab[4636]"
	}
//line /usr/local/go/src/net/dial.go:177
	// _ = "end of CoverTab[4634]"
}

func parseNetwork(ctx context.Context, network string, needsProto bool) (afnet string, proto int, err error) {
//line /usr/local/go/src/net/dial.go:180
	_go_fuzz_dep_.CoverTab[4637]++
						i := last(network, ':')
						if i < 0 {
//line /usr/local/go/src/net/dial.go:182
		_go_fuzz_dep_.CoverTab[4640]++
							switch network {
		case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/dial.go:184
			_go_fuzz_dep_.CoverTab[4642]++
//line /usr/local/go/src/net/dial.go:184
			// _ = "end of CoverTab[4642]"
		case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/dial.go:185
			_go_fuzz_dep_.CoverTab[4643]++
//line /usr/local/go/src/net/dial.go:185
			// _ = "end of CoverTab[4643]"
		case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/dial.go:186
			_go_fuzz_dep_.CoverTab[4644]++
								if needsProto {
//line /usr/local/go/src/net/dial.go:187
				_go_fuzz_dep_.CoverTab[4647]++
									return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:188
				// _ = "end of CoverTab[4647]"
			} else {
//line /usr/local/go/src/net/dial.go:189
				_go_fuzz_dep_.CoverTab[4648]++
//line /usr/local/go/src/net/dial.go:189
				// _ = "end of CoverTab[4648]"
//line /usr/local/go/src/net/dial.go:189
			}
//line /usr/local/go/src/net/dial.go:189
			// _ = "end of CoverTab[4644]"
		case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/dial.go:190
			_go_fuzz_dep_.CoverTab[4645]++
//line /usr/local/go/src/net/dial.go:190
			// _ = "end of CoverTab[4645]"
		default:
//line /usr/local/go/src/net/dial.go:191
			_go_fuzz_dep_.CoverTab[4646]++
								return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:192
			// _ = "end of CoverTab[4646]"
		}
//line /usr/local/go/src/net/dial.go:193
		// _ = "end of CoverTab[4640]"
//line /usr/local/go/src/net/dial.go:193
		_go_fuzz_dep_.CoverTab[4641]++
							return network, 0, nil
//line /usr/local/go/src/net/dial.go:194
		// _ = "end of CoverTab[4641]"
	} else {
//line /usr/local/go/src/net/dial.go:195
		_go_fuzz_dep_.CoverTab[4649]++
//line /usr/local/go/src/net/dial.go:195
		// _ = "end of CoverTab[4649]"
//line /usr/local/go/src/net/dial.go:195
	}
//line /usr/local/go/src/net/dial.go:195
	// _ = "end of CoverTab[4637]"
//line /usr/local/go/src/net/dial.go:195
	_go_fuzz_dep_.CoverTab[4638]++
						afnet = network[:i]
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/dial.go:198
		_go_fuzz_dep_.CoverTab[4650]++
							protostr := network[i+1:]
							proto, i, ok := dtoi(protostr)
							if !ok || func() bool {
//line /usr/local/go/src/net/dial.go:201
			_go_fuzz_dep_.CoverTab[4653]++
//line /usr/local/go/src/net/dial.go:201
			return i != len(protostr)
//line /usr/local/go/src/net/dial.go:201
			// _ = "end of CoverTab[4653]"
//line /usr/local/go/src/net/dial.go:201
		}() {
//line /usr/local/go/src/net/dial.go:201
			_go_fuzz_dep_.CoverTab[4654]++
								proto, err = lookupProtocol(ctx, protostr)
								if err != nil {
//line /usr/local/go/src/net/dial.go:203
				_go_fuzz_dep_.CoverTab[4655]++
									return "", 0, err
//line /usr/local/go/src/net/dial.go:204
				// _ = "end of CoverTab[4655]"
			} else {
//line /usr/local/go/src/net/dial.go:205
				_go_fuzz_dep_.CoverTab[4656]++
//line /usr/local/go/src/net/dial.go:205
				// _ = "end of CoverTab[4656]"
//line /usr/local/go/src/net/dial.go:205
			}
//line /usr/local/go/src/net/dial.go:205
			// _ = "end of CoverTab[4654]"
		} else {
//line /usr/local/go/src/net/dial.go:206
			_go_fuzz_dep_.CoverTab[4657]++
//line /usr/local/go/src/net/dial.go:206
			// _ = "end of CoverTab[4657]"
//line /usr/local/go/src/net/dial.go:206
		}
//line /usr/local/go/src/net/dial.go:206
		// _ = "end of CoverTab[4650]"
//line /usr/local/go/src/net/dial.go:206
		_go_fuzz_dep_.CoverTab[4651]++
							return afnet, proto, nil
//line /usr/local/go/src/net/dial.go:207
		// _ = "end of CoverTab[4651]"
//line /usr/local/go/src/net/dial.go:207
	default:
//line /usr/local/go/src/net/dial.go:207
		_go_fuzz_dep_.CoverTab[4652]++
//line /usr/local/go/src/net/dial.go:207
		// _ = "end of CoverTab[4652]"
	}
//line /usr/local/go/src/net/dial.go:208
	// _ = "end of CoverTab[4638]"
//line /usr/local/go/src/net/dial.go:208
	_go_fuzz_dep_.CoverTab[4639]++
						return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:209
	// _ = "end of CoverTab[4639]"
}

// resolveAddrList resolves addr using hint and returns a list of
//line /usr/local/go/src/net/dial.go:212
// addresses. The result contains at least one address when error is
//line /usr/local/go/src/net/dial.go:212
// nil.
//line /usr/local/go/src/net/dial.go:215
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, addr string, hint Addr) (addrList, error) {
//line /usr/local/go/src/net/dial.go:215
	_go_fuzz_dep_.CoverTab[4658]++
						afnet, _, err := parseNetwork(ctx, network, true)
						if err != nil {
//line /usr/local/go/src/net/dial.go:217
		_go_fuzz_dep_.CoverTab[4666]++
							return nil, err
//line /usr/local/go/src/net/dial.go:218
		// _ = "end of CoverTab[4666]"
	} else {
//line /usr/local/go/src/net/dial.go:219
		_go_fuzz_dep_.CoverTab[4667]++
//line /usr/local/go/src/net/dial.go:219
		// _ = "end of CoverTab[4667]"
//line /usr/local/go/src/net/dial.go:219
	}
//line /usr/local/go/src/net/dial.go:219
	// _ = "end of CoverTab[4658]"
//line /usr/local/go/src/net/dial.go:219
	_go_fuzz_dep_.CoverTab[4659]++
						if op == "dial" && func() bool {
//line /usr/local/go/src/net/dial.go:220
		_go_fuzz_dep_.CoverTab[4668]++
//line /usr/local/go/src/net/dial.go:220
		return addr == ""
//line /usr/local/go/src/net/dial.go:220
		// _ = "end of CoverTab[4668]"
//line /usr/local/go/src/net/dial.go:220
	}() {
//line /usr/local/go/src/net/dial.go:220
		_go_fuzz_dep_.CoverTab[4669]++
							return nil, errMissingAddress
//line /usr/local/go/src/net/dial.go:221
		// _ = "end of CoverTab[4669]"
	} else {
//line /usr/local/go/src/net/dial.go:222
		_go_fuzz_dep_.CoverTab[4670]++
//line /usr/local/go/src/net/dial.go:222
		// _ = "end of CoverTab[4670]"
//line /usr/local/go/src/net/dial.go:222
	}
//line /usr/local/go/src/net/dial.go:222
	// _ = "end of CoverTab[4659]"
//line /usr/local/go/src/net/dial.go:222
	_go_fuzz_dep_.CoverTab[4660]++
						switch afnet {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/dial.go:224
		_go_fuzz_dep_.CoverTab[4671]++
							addr, err := ResolveUnixAddr(afnet, addr)
							if err != nil {
//line /usr/local/go/src/net/dial.go:226
			_go_fuzz_dep_.CoverTab[4675]++
								return nil, err
//line /usr/local/go/src/net/dial.go:227
			// _ = "end of CoverTab[4675]"
		} else {
//line /usr/local/go/src/net/dial.go:228
			_go_fuzz_dep_.CoverTab[4676]++
//line /usr/local/go/src/net/dial.go:228
			// _ = "end of CoverTab[4676]"
//line /usr/local/go/src/net/dial.go:228
		}
//line /usr/local/go/src/net/dial.go:228
		// _ = "end of CoverTab[4671]"
//line /usr/local/go/src/net/dial.go:228
		_go_fuzz_dep_.CoverTab[4672]++
							if op == "dial" && func() bool {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[4677]++
//line /usr/local/go/src/net/dial.go:229
			return hint != nil
//line /usr/local/go/src/net/dial.go:229
			// _ = "end of CoverTab[4677]"
//line /usr/local/go/src/net/dial.go:229
		}() && func() bool {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[4678]++
//line /usr/local/go/src/net/dial.go:229
			return addr.Network() != hint.Network()
//line /usr/local/go/src/net/dial.go:229
			// _ = "end of CoverTab[4678]"
//line /usr/local/go/src/net/dial.go:229
		}() {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[4679]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:230
			// _ = "end of CoverTab[4679]"
		} else {
//line /usr/local/go/src/net/dial.go:231
			_go_fuzz_dep_.CoverTab[4680]++
//line /usr/local/go/src/net/dial.go:231
			// _ = "end of CoverTab[4680]"
//line /usr/local/go/src/net/dial.go:231
		}
//line /usr/local/go/src/net/dial.go:231
		// _ = "end of CoverTab[4672]"
//line /usr/local/go/src/net/dial.go:231
		_go_fuzz_dep_.CoverTab[4673]++
							return addrList{addr}, nil
//line /usr/local/go/src/net/dial.go:232
		// _ = "end of CoverTab[4673]"
//line /usr/local/go/src/net/dial.go:232
	default:
//line /usr/local/go/src/net/dial.go:232
		_go_fuzz_dep_.CoverTab[4674]++
//line /usr/local/go/src/net/dial.go:232
		// _ = "end of CoverTab[4674]"
	}
//line /usr/local/go/src/net/dial.go:233
	// _ = "end of CoverTab[4660]"
//line /usr/local/go/src/net/dial.go:233
	_go_fuzz_dep_.CoverTab[4661]++
						addrs, err := r.internetAddrList(ctx, afnet, addr)
						if err != nil || func() bool {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[4681]++
//line /usr/local/go/src/net/dial.go:235
		return op != "dial"
//line /usr/local/go/src/net/dial.go:235
		// _ = "end of CoverTab[4681]"
//line /usr/local/go/src/net/dial.go:235
	}() || func() bool {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[4682]++
//line /usr/local/go/src/net/dial.go:235
		return hint == nil
//line /usr/local/go/src/net/dial.go:235
		// _ = "end of CoverTab[4682]"
//line /usr/local/go/src/net/dial.go:235
	}() {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[4683]++
							return addrs, err
//line /usr/local/go/src/net/dial.go:236
		// _ = "end of CoverTab[4683]"
	} else {
//line /usr/local/go/src/net/dial.go:237
		_go_fuzz_dep_.CoverTab[4684]++
//line /usr/local/go/src/net/dial.go:237
		// _ = "end of CoverTab[4684]"
//line /usr/local/go/src/net/dial.go:237
	}
//line /usr/local/go/src/net/dial.go:237
	// _ = "end of CoverTab[4661]"
//line /usr/local/go/src/net/dial.go:237
	_go_fuzz_dep_.CoverTab[4662]++
						var (
		tcp		*TCPAddr
		udp		*UDPAddr
		ip		*IPAddr
		wildcard	bool
	)
	switch hint := hint.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/dial.go:245
		_go_fuzz_dep_.CoverTab[4685]++
							tcp = hint
							wildcard = tcp.isWildcard()
//line /usr/local/go/src/net/dial.go:247
		// _ = "end of CoverTab[4685]"
	case *UDPAddr:
//line /usr/local/go/src/net/dial.go:248
		_go_fuzz_dep_.CoverTab[4686]++
							udp = hint
							wildcard = udp.isWildcard()
//line /usr/local/go/src/net/dial.go:250
		// _ = "end of CoverTab[4686]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:251
		_go_fuzz_dep_.CoverTab[4687]++
							ip = hint
							wildcard = ip.isWildcard()
//line /usr/local/go/src/net/dial.go:253
		// _ = "end of CoverTab[4687]"
	}
//line /usr/local/go/src/net/dial.go:254
	// _ = "end of CoverTab[4662]"
//line /usr/local/go/src/net/dial.go:254
	_go_fuzz_dep_.CoverTab[4663]++
						naddrs := addrs[:0]
						for _, addr := range addrs {
//line /usr/local/go/src/net/dial.go:256
		_go_fuzz_dep_.CoverTab[4688]++
							if addr.Network() != hint.Network() {
//line /usr/local/go/src/net/dial.go:257
			_go_fuzz_dep_.CoverTab[4690]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:258
			// _ = "end of CoverTab[4690]"
		} else {
//line /usr/local/go/src/net/dial.go:259
			_go_fuzz_dep_.CoverTab[4691]++
//line /usr/local/go/src/net/dial.go:259
			// _ = "end of CoverTab[4691]"
//line /usr/local/go/src/net/dial.go:259
		}
//line /usr/local/go/src/net/dial.go:259
		// _ = "end of CoverTab[4688]"
//line /usr/local/go/src/net/dial.go:259
		_go_fuzz_dep_.CoverTab[4689]++
							switch addr := addr.(type) {
		case *TCPAddr:
//line /usr/local/go/src/net/dial.go:261
			_go_fuzz_dep_.CoverTab[4692]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[4698]++
//line /usr/local/go/src/net/dial.go:262
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:262
				// _ = "end of CoverTab[4698]"
//line /usr/local/go/src/net/dial.go:262
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[4699]++
//line /usr/local/go/src/net/dial.go:262
				return !addr.IP.matchAddrFamily(tcp.IP)
//line /usr/local/go/src/net/dial.go:262
				// _ = "end of CoverTab[4699]"
//line /usr/local/go/src/net/dial.go:262
			}() {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[4700]++
									continue
//line /usr/local/go/src/net/dial.go:263
				// _ = "end of CoverTab[4700]"
			} else {
//line /usr/local/go/src/net/dial.go:264
				_go_fuzz_dep_.CoverTab[4701]++
//line /usr/local/go/src/net/dial.go:264
				// _ = "end of CoverTab[4701]"
//line /usr/local/go/src/net/dial.go:264
			}
//line /usr/local/go/src/net/dial.go:264
			// _ = "end of CoverTab[4692]"
//line /usr/local/go/src/net/dial.go:264
			_go_fuzz_dep_.CoverTab[4693]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:265
			// _ = "end of CoverTab[4693]"
		case *UDPAddr:
//line /usr/local/go/src/net/dial.go:266
			_go_fuzz_dep_.CoverTab[4694]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[4702]++
//line /usr/local/go/src/net/dial.go:267
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:267
				// _ = "end of CoverTab[4702]"
//line /usr/local/go/src/net/dial.go:267
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[4703]++
//line /usr/local/go/src/net/dial.go:267
				return !addr.IP.matchAddrFamily(udp.IP)
//line /usr/local/go/src/net/dial.go:267
				// _ = "end of CoverTab[4703]"
//line /usr/local/go/src/net/dial.go:267
			}() {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[4704]++
									continue
//line /usr/local/go/src/net/dial.go:268
				// _ = "end of CoverTab[4704]"
			} else {
//line /usr/local/go/src/net/dial.go:269
				_go_fuzz_dep_.CoverTab[4705]++
//line /usr/local/go/src/net/dial.go:269
				// _ = "end of CoverTab[4705]"
//line /usr/local/go/src/net/dial.go:269
			}
//line /usr/local/go/src/net/dial.go:269
			// _ = "end of CoverTab[4694]"
//line /usr/local/go/src/net/dial.go:269
			_go_fuzz_dep_.CoverTab[4695]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:270
			// _ = "end of CoverTab[4695]"
		case *IPAddr:
//line /usr/local/go/src/net/dial.go:271
			_go_fuzz_dep_.CoverTab[4696]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[4706]++
//line /usr/local/go/src/net/dial.go:272
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:272
				// _ = "end of CoverTab[4706]"
//line /usr/local/go/src/net/dial.go:272
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[4707]++
//line /usr/local/go/src/net/dial.go:272
				return !addr.IP.matchAddrFamily(ip.IP)
//line /usr/local/go/src/net/dial.go:272
				// _ = "end of CoverTab[4707]"
//line /usr/local/go/src/net/dial.go:272
			}() {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[4708]++
									continue
//line /usr/local/go/src/net/dial.go:273
				// _ = "end of CoverTab[4708]"
			} else {
//line /usr/local/go/src/net/dial.go:274
				_go_fuzz_dep_.CoverTab[4709]++
//line /usr/local/go/src/net/dial.go:274
				// _ = "end of CoverTab[4709]"
//line /usr/local/go/src/net/dial.go:274
			}
//line /usr/local/go/src/net/dial.go:274
			// _ = "end of CoverTab[4696]"
//line /usr/local/go/src/net/dial.go:274
			_go_fuzz_dep_.CoverTab[4697]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:275
			// _ = "end of CoverTab[4697]"
		}
//line /usr/local/go/src/net/dial.go:276
		// _ = "end of CoverTab[4689]"
	}
//line /usr/local/go/src/net/dial.go:277
	// _ = "end of CoverTab[4663]"
//line /usr/local/go/src/net/dial.go:277
	_go_fuzz_dep_.CoverTab[4664]++
						if len(naddrs) == 0 {
//line /usr/local/go/src/net/dial.go:278
		_go_fuzz_dep_.CoverTab[4710]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:279
		// _ = "end of CoverTab[4710]"
	} else {
//line /usr/local/go/src/net/dial.go:280
		_go_fuzz_dep_.CoverTab[4711]++
//line /usr/local/go/src/net/dial.go:280
		// _ = "end of CoverTab[4711]"
//line /usr/local/go/src/net/dial.go:280
	}
//line /usr/local/go/src/net/dial.go:280
	// _ = "end of CoverTab[4664]"
//line /usr/local/go/src/net/dial.go:280
	_go_fuzz_dep_.CoverTab[4665]++
						return naddrs, nil
//line /usr/local/go/src/net/dial.go:281
	// _ = "end of CoverTab[4665]"
}

// Dial connects to the address on the named network.
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
//line /usr/local/go/src/net/dial.go:284
// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
//line /usr/local/go/src/net/dial.go:284
// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
//line /usr/local/go/src/net/dial.go:284
// "unixpacket".
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// For TCP and UDP networks, the address has the form "host:port".
//line /usr/local/go/src/net/dial.go:284
// The host must be a literal IP address, or a host name that can be
//line /usr/local/go/src/net/dial.go:284
// resolved to IP addresses.
//line /usr/local/go/src/net/dial.go:284
// The port must be a literal port number or a service name.
//line /usr/local/go/src/net/dial.go:284
// If the host is a literal IPv6 address it must be enclosed in square
//line /usr/local/go/src/net/dial.go:284
// brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
//line /usr/local/go/src/net/dial.go:284
// The zone specifies the scope of the literal IPv6 address as defined
//line /usr/local/go/src/net/dial.go:284
// in RFC 4007.
//line /usr/local/go/src/net/dial.go:284
// The functions JoinHostPort and SplitHostPort manipulate a pair of
//line /usr/local/go/src/net/dial.go:284
// host and port in this form.
//line /usr/local/go/src/net/dial.go:284
// When using TCP, and the host resolves to multiple IP addresses,
//line /usr/local/go/src/net/dial.go:284
// Dial will try each IP address in order until one succeeds.
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// Examples:
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
//	Dial("tcp", "golang.org:http")
//line /usr/local/go/src/net/dial.go:284
//	Dial("tcp", "192.0.2.1:http")
//line /usr/local/go/src/net/dial.go:284
//	Dial("tcp", "198.51.100.1:80")
//line /usr/local/go/src/net/dial.go:284
//	Dial("udp", "[2001:db8::1]:domain")
//line /usr/local/go/src/net/dial.go:284
//	Dial("udp", "[fe80::1%lo0]:53")
//line /usr/local/go/src/net/dial.go:284
//	Dial("tcp", ":80")
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// For IP networks, the network must be "ip", "ip4" or "ip6" followed
//line /usr/local/go/src/net/dial.go:284
// by a colon and a literal protocol number or a protocol name, and
//line /usr/local/go/src/net/dial.go:284
// the address has the form "host". The host must be a literal IP
//line /usr/local/go/src/net/dial.go:284
// address or a literal IPv6 address with zone.
//line /usr/local/go/src/net/dial.go:284
// It depends on each operating system how the operating system
//line /usr/local/go/src/net/dial.go:284
// behaves with a non-well known protocol number such as "0" or "255".
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// Examples:
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
//	Dial("ip4:1", "192.0.2.1")
//line /usr/local/go/src/net/dial.go:284
//	Dial("ip6:ipv6-icmp", "2001:db8::1")
//line /usr/local/go/src/net/dial.go:284
//	Dial("ip6:58", "fe80::1%lo0")
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// For TCP, UDP and IP networks, if the host is empty or a literal
//line /usr/local/go/src/net/dial.go:284
// unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
//line /usr/local/go/src/net/dial.go:284
// TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
//line /usr/local/go/src/net/dial.go:284
// assumed.
//line /usr/local/go/src/net/dial.go:284
//
//line /usr/local/go/src/net/dial.go:284
// For Unix networks, the address must be a file system path.
//line /usr/local/go/src/net/dial.go:332
func Dial(network, address string) (Conn, error) {
//line /usr/local/go/src/net/dial.go:332
	_go_fuzz_dep_.CoverTab[4712]++
						var d Dialer
						return d.Dial(network, address)
//line /usr/local/go/src/net/dial.go:334
	// _ = "end of CoverTab[4712]"
}

// DialTimeout acts like Dial but takes a timeout.
//line /usr/local/go/src/net/dial.go:337
//
//line /usr/local/go/src/net/dial.go:337
// The timeout includes name resolution, if required.
//line /usr/local/go/src/net/dial.go:337
// When using TCP, and the host in the address parameter resolves to
//line /usr/local/go/src/net/dial.go:337
// multiple IP addresses, the timeout is spread over each consecutive
//line /usr/local/go/src/net/dial.go:337
// dial, such that each is given an appropriate fraction of the time
//line /usr/local/go/src/net/dial.go:337
// to connect.
//line /usr/local/go/src/net/dial.go:337
//
//line /usr/local/go/src/net/dial.go:337
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/dial.go:337
// parameters.
//line /usr/local/go/src/net/dial.go:347
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) {
//line /usr/local/go/src/net/dial.go:347
	_go_fuzz_dep_.CoverTab[4713]++
						d := Dialer{Timeout: timeout}
						return d.Dial(network, address)
//line /usr/local/go/src/net/dial.go:349
	// _ = "end of CoverTab[4713]"
}

// sysDialer contains a Dial's parameters and configuration.
type sysDialer struct {
	Dialer
	network, address	string
	testHookDialTCP		func(ctx context.Context, net string, laddr, raddr *TCPAddr) (*TCPConn, error)
}

// Dial connects to the address on the named network.
//line /usr/local/go/src/net/dial.go:359
//
//line /usr/local/go/src/net/dial.go:359
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/dial.go:359
// parameters.
//line /usr/local/go/src/net/dial.go:359
//
//line /usr/local/go/src/net/dial.go:359
// Dial uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/dial.go:359
// DialContext.
//line /usr/local/go/src/net/dial.go:366
func (d *Dialer) Dial(network, address string) (Conn, error) {
//line /usr/local/go/src/net/dial.go:366
	_go_fuzz_dep_.CoverTab[4714]++
						return d.DialContext(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:367
	// _ = "end of CoverTab[4714]"
}

// DialContext connects to the address on the named network using
//line /usr/local/go/src/net/dial.go:370
// the provided context.
//line /usr/local/go/src/net/dial.go:370
//
//line /usr/local/go/src/net/dial.go:370
// The provided Context must be non-nil. If the context expires before
//line /usr/local/go/src/net/dial.go:370
// the connection is complete, an error is returned. Once successfully
//line /usr/local/go/src/net/dial.go:370
// connected, any expiration of the context will not affect the
//line /usr/local/go/src/net/dial.go:370
// connection.
//line /usr/local/go/src/net/dial.go:370
//
//line /usr/local/go/src/net/dial.go:370
// When using TCP, and the host in the address parameter resolves to multiple
//line /usr/local/go/src/net/dial.go:370
// network addresses, any dial timeout (from d.Timeout or ctx) is spread
//line /usr/local/go/src/net/dial.go:370
// over each consecutive dial, such that each is given an appropriate
//line /usr/local/go/src/net/dial.go:370
// fraction of the time to connect.
//line /usr/local/go/src/net/dial.go:370
// For example, if a host has 4 IP addresses and the timeout is 1 minute,
//line /usr/local/go/src/net/dial.go:370
// the connect to each single address will be given 15 seconds to complete
//line /usr/local/go/src/net/dial.go:370
// before trying the next one.
//line /usr/local/go/src/net/dial.go:370
//
//line /usr/local/go/src/net/dial.go:370
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/dial.go:370
// parameters.
//line /usr/local/go/src/net/dial.go:388
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error) {
//line /usr/local/go/src/net/dial.go:388
	_go_fuzz_dep_.CoverTab[4715]++
						if ctx == nil {
//line /usr/local/go/src/net/dial.go:389
		_go_fuzz_dep_.CoverTab[4722]++
							panic("nil context")
//line /usr/local/go/src/net/dial.go:390
		// _ = "end of CoverTab[4722]"
	} else {
//line /usr/local/go/src/net/dial.go:391
		_go_fuzz_dep_.CoverTab[4723]++
//line /usr/local/go/src/net/dial.go:391
		// _ = "end of CoverTab[4723]"
//line /usr/local/go/src/net/dial.go:391
	}
//line /usr/local/go/src/net/dial.go:391
	// _ = "end of CoverTab[4715]"
//line /usr/local/go/src/net/dial.go:391
	_go_fuzz_dep_.CoverTab[4716]++
						deadline := d.deadline(ctx, time.Now())
						if !deadline.IsZero() {
//line /usr/local/go/src/net/dial.go:393
		_go_fuzz_dep_.CoverTab[4724]++
							if d, ok := ctx.Deadline(); !ok || func() bool {
//line /usr/local/go/src/net/dial.go:394
			_go_fuzz_dep_.CoverTab[4725]++
//line /usr/local/go/src/net/dial.go:394
			return deadline.Before(d)
//line /usr/local/go/src/net/dial.go:394
			// _ = "end of CoverTab[4725]"
//line /usr/local/go/src/net/dial.go:394
		}() {
//line /usr/local/go/src/net/dial.go:394
			_go_fuzz_dep_.CoverTab[4726]++
								subCtx, cancel := context.WithDeadline(ctx, deadline)
								defer cancel()
								ctx = subCtx
//line /usr/local/go/src/net/dial.go:397
			// _ = "end of CoverTab[4726]"
		} else {
//line /usr/local/go/src/net/dial.go:398
			_go_fuzz_dep_.CoverTab[4727]++
//line /usr/local/go/src/net/dial.go:398
			// _ = "end of CoverTab[4727]"
//line /usr/local/go/src/net/dial.go:398
		}
//line /usr/local/go/src/net/dial.go:398
		// _ = "end of CoverTab[4724]"
	} else {
//line /usr/local/go/src/net/dial.go:399
		_go_fuzz_dep_.CoverTab[4728]++
//line /usr/local/go/src/net/dial.go:399
		// _ = "end of CoverTab[4728]"
//line /usr/local/go/src/net/dial.go:399
	}
//line /usr/local/go/src/net/dial.go:399
	// _ = "end of CoverTab[4716]"
//line /usr/local/go/src/net/dial.go:399
	_go_fuzz_dep_.CoverTab[4717]++
						if oldCancel := d.Cancel; oldCancel != nil {
//line /usr/local/go/src/net/dial.go:400
		_go_fuzz_dep_.CoverTab[4729]++
							subCtx, cancel := context.WithCancel(ctx)
							defer cancel()
//line /usr/local/go/src/net/dial.go:402
		_curRoutineNum6_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/dial.go:402
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum6_)
							go func() {
//line /usr/local/go/src/net/dial.go:403
			_go_fuzz_dep_.CoverTab[4731]++
//line /usr/local/go/src/net/dial.go:403
			defer func() {
//line /usr/local/go/src/net/dial.go:403
				_go_fuzz_dep_.CoverTab[4732]++
//line /usr/local/go/src/net/dial.go:403
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum6_)
//line /usr/local/go/src/net/dial.go:403
				// _ = "end of CoverTab[4732]"
//line /usr/local/go/src/net/dial.go:403
			}()
								select {
			case <-oldCancel:
//line /usr/local/go/src/net/dial.go:405
				_go_fuzz_dep_.CoverTab[4733]++
									cancel()
//line /usr/local/go/src/net/dial.go:406
				// _ = "end of CoverTab[4733]"
			case <-subCtx.Done():
//line /usr/local/go/src/net/dial.go:407
				_go_fuzz_dep_.CoverTab[4734]++
//line /usr/local/go/src/net/dial.go:407
				// _ = "end of CoverTab[4734]"
			}
//line /usr/local/go/src/net/dial.go:408
			// _ = "end of CoverTab[4731]"
		}()
//line /usr/local/go/src/net/dial.go:409
		// _ = "end of CoverTab[4729]"
//line /usr/local/go/src/net/dial.go:409
		_go_fuzz_dep_.CoverTab[4730]++
							ctx = subCtx
//line /usr/local/go/src/net/dial.go:410
		// _ = "end of CoverTab[4730]"
	} else {
//line /usr/local/go/src/net/dial.go:411
		_go_fuzz_dep_.CoverTab[4735]++
//line /usr/local/go/src/net/dial.go:411
		// _ = "end of CoverTab[4735]"
//line /usr/local/go/src/net/dial.go:411
	}
//line /usr/local/go/src/net/dial.go:411
	// _ = "end of CoverTab[4717]"
//line /usr/local/go/src/net/dial.go:411
	_go_fuzz_dep_.CoverTab[4718]++

//line /usr/local/go/src/net/dial.go:414
	resolveCtx := ctx
	if trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace); trace != nil {
//line /usr/local/go/src/net/dial.go:415
		_go_fuzz_dep_.CoverTab[4736]++
							shadow := *trace
							shadow.ConnectStart = nil
							shadow.ConnectDone = nil
							resolveCtx = context.WithValue(resolveCtx, nettrace.TraceKey{}, &shadow)
//line /usr/local/go/src/net/dial.go:419
		// _ = "end of CoverTab[4736]"
	} else {
//line /usr/local/go/src/net/dial.go:420
		_go_fuzz_dep_.CoverTab[4737]++
//line /usr/local/go/src/net/dial.go:420
		// _ = "end of CoverTab[4737]"
//line /usr/local/go/src/net/dial.go:420
	}
//line /usr/local/go/src/net/dial.go:420
	// _ = "end of CoverTab[4718]"
//line /usr/local/go/src/net/dial.go:420
	_go_fuzz_dep_.CoverTab[4719]++

						addrs, err := d.resolver().resolveAddrList(resolveCtx, "dial", network, address, d.LocalAddr)
						if err != nil {
//line /usr/local/go/src/net/dial.go:423
		_go_fuzz_dep_.CoverTab[4738]++
							return nil, &OpError{Op: "dial", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:424
		// _ = "end of CoverTab[4738]"
	} else {
//line /usr/local/go/src/net/dial.go:425
		_go_fuzz_dep_.CoverTab[4739]++
//line /usr/local/go/src/net/dial.go:425
		// _ = "end of CoverTab[4739]"
//line /usr/local/go/src/net/dial.go:425
	}
//line /usr/local/go/src/net/dial.go:425
	// _ = "end of CoverTab[4719]"
//line /usr/local/go/src/net/dial.go:425
	_go_fuzz_dep_.CoverTab[4720]++

						sd := &sysDialer{
		Dialer:		*d,
		network:	network,
		address:	address,
	}

	var primaries, fallbacks addrList
	if d.dualStack() && func() bool {
//line /usr/local/go/src/net/dial.go:434
		_go_fuzz_dep_.CoverTab[4740]++
//line /usr/local/go/src/net/dial.go:434
		return network == "tcp"
//line /usr/local/go/src/net/dial.go:434
		// _ = "end of CoverTab[4740]"
//line /usr/local/go/src/net/dial.go:434
	}() {
//line /usr/local/go/src/net/dial.go:434
		_go_fuzz_dep_.CoverTab[4741]++
							primaries, fallbacks = addrs.partition(isIPv4)
//line /usr/local/go/src/net/dial.go:435
		// _ = "end of CoverTab[4741]"
	} else {
//line /usr/local/go/src/net/dial.go:436
		_go_fuzz_dep_.CoverTab[4742]++
							primaries = addrs
//line /usr/local/go/src/net/dial.go:437
		// _ = "end of CoverTab[4742]"
	}
//line /usr/local/go/src/net/dial.go:438
	// _ = "end of CoverTab[4720]"
//line /usr/local/go/src/net/dial.go:438
	_go_fuzz_dep_.CoverTab[4721]++

						return sd.dialParallel(ctx, primaries, fallbacks)
//line /usr/local/go/src/net/dial.go:440
	// _ = "end of CoverTab[4721]"
}

// dialParallel races two copies of dialSerial, giving the first a
//line /usr/local/go/src/net/dial.go:443
// head start. It returns the first established connection and
//line /usr/local/go/src/net/dial.go:443
// closes the others. Otherwise it returns an error from the first
//line /usr/local/go/src/net/dial.go:443
// primary address.
//line /usr/local/go/src/net/dial.go:447
func (sd *sysDialer) dialParallel(ctx context.Context, primaries, fallbacks addrList) (Conn, error) {
//line /usr/local/go/src/net/dial.go:447
	_go_fuzz_dep_.CoverTab[4743]++
						if len(fallbacks) == 0 {
//line /usr/local/go/src/net/dial.go:448
		_go_fuzz_dep_.CoverTab[4746]++
							return sd.dialSerial(ctx, primaries)
//line /usr/local/go/src/net/dial.go:449
		// _ = "end of CoverTab[4746]"
	} else {
//line /usr/local/go/src/net/dial.go:450
		_go_fuzz_dep_.CoverTab[4747]++
//line /usr/local/go/src/net/dial.go:450
		// _ = "end of CoverTab[4747]"
//line /usr/local/go/src/net/dial.go:450
	}
//line /usr/local/go/src/net/dial.go:450
	// _ = "end of CoverTab[4743]"
//line /usr/local/go/src/net/dial.go:450
	_go_fuzz_dep_.CoverTab[4744]++

						returned := make(chan struct{})
						defer close(returned)

						type dialResult struct {
		Conn
		error
		primary	bool
		done	bool
	}
	results := make(chan dialResult)

	startRacer := func(ctx context.Context, primary bool) {
//line /usr/local/go/src/net/dial.go:463
		_go_fuzz_dep_.CoverTab[4748]++
							ras := primaries
							if !primary {
//line /usr/local/go/src/net/dial.go:465
			_go_fuzz_dep_.CoverTab[4750]++
								ras = fallbacks
//line /usr/local/go/src/net/dial.go:466
			// _ = "end of CoverTab[4750]"
		} else {
//line /usr/local/go/src/net/dial.go:467
			_go_fuzz_dep_.CoverTab[4751]++
//line /usr/local/go/src/net/dial.go:467
			// _ = "end of CoverTab[4751]"
//line /usr/local/go/src/net/dial.go:467
		}
//line /usr/local/go/src/net/dial.go:467
		// _ = "end of CoverTab[4748]"
//line /usr/local/go/src/net/dial.go:467
		_go_fuzz_dep_.CoverTab[4749]++
							c, err := sd.dialSerial(ctx, ras)
							select {
		case results <- dialResult{Conn: c, error: err, primary: primary, done: true}:
//line /usr/local/go/src/net/dial.go:470
			_go_fuzz_dep_.CoverTab[4752]++
//line /usr/local/go/src/net/dial.go:470
			// _ = "end of CoverTab[4752]"
		case <-returned:
//line /usr/local/go/src/net/dial.go:471
			_go_fuzz_dep_.CoverTab[4753]++
								if c != nil {
//line /usr/local/go/src/net/dial.go:472
				_go_fuzz_dep_.CoverTab[4754]++
									c.Close()
//line /usr/local/go/src/net/dial.go:473
				// _ = "end of CoverTab[4754]"
			} else {
//line /usr/local/go/src/net/dial.go:474
				_go_fuzz_dep_.CoverTab[4755]++
//line /usr/local/go/src/net/dial.go:474
				// _ = "end of CoverTab[4755]"
//line /usr/local/go/src/net/dial.go:474
			}
//line /usr/local/go/src/net/dial.go:474
			// _ = "end of CoverTab[4753]"
		}
//line /usr/local/go/src/net/dial.go:475
		// _ = "end of CoverTab[4749]"
	}
//line /usr/local/go/src/net/dial.go:476
	// _ = "end of CoverTab[4744]"
//line /usr/local/go/src/net/dial.go:476
	_go_fuzz_dep_.CoverTab[4745]++

						var primary, fallback dialResult

//line /usr/local/go/src/net/dial.go:481
	primaryCtx, primaryCancel := context.WithCancel(ctx)
						defer primaryCancel()
//line /usr/local/go/src/net/dial.go:482
	_curRoutineNum7_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/dial.go:482
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum7_)
						go func() {
//line /usr/local/go/src/net/dial.go:483
		_go_fuzz_dep_.CoverTab[4756]++
//line /usr/local/go/src/net/dial.go:483
		defer func() {
//line /usr/local/go/src/net/dial.go:483
			_go_fuzz_dep_.CoverTab[4757]++
//line /usr/local/go/src/net/dial.go:483
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum7_)
//line /usr/local/go/src/net/dial.go:483
			// _ = "end of CoverTab[4757]"
//line /usr/local/go/src/net/dial.go:483
		}()
//line /usr/local/go/src/net/dial.go:483
		startRacer(primaryCtx, true)
//line /usr/local/go/src/net/dial.go:483
		// _ = "end of CoverTab[4756]"
//line /usr/local/go/src/net/dial.go:483
	}()

//line /usr/local/go/src/net/dial.go:486
	fallbackTimer := time.NewTimer(sd.fallbackDelay())
	defer fallbackTimer.Stop()

	for {
//line /usr/local/go/src/net/dial.go:489
		_go_fuzz_dep_.CoverTab[4758]++
							select {
		case <-fallbackTimer.C:
//line /usr/local/go/src/net/dial.go:491
			_go_fuzz_dep_.CoverTab[4759]++
								fallbackCtx, fallbackCancel := context.WithCancel(ctx)
								defer fallbackCancel()
								go startRacer(fallbackCtx, false)
//line /usr/local/go/src/net/dial.go:494
			// _ = "end of CoverTab[4759]"

		case res := <-results:
//line /usr/local/go/src/net/dial.go:496
			_go_fuzz_dep_.CoverTab[4760]++
								if res.error == nil {
//line /usr/local/go/src/net/dial.go:497
				_go_fuzz_dep_.CoverTab[4764]++
									return res.Conn, nil
//line /usr/local/go/src/net/dial.go:498
				// _ = "end of CoverTab[4764]"
			} else {
//line /usr/local/go/src/net/dial.go:499
				_go_fuzz_dep_.CoverTab[4765]++
//line /usr/local/go/src/net/dial.go:499
				// _ = "end of CoverTab[4765]"
//line /usr/local/go/src/net/dial.go:499
			}
//line /usr/local/go/src/net/dial.go:499
			// _ = "end of CoverTab[4760]"
//line /usr/local/go/src/net/dial.go:499
			_go_fuzz_dep_.CoverTab[4761]++
								if res.primary {
//line /usr/local/go/src/net/dial.go:500
				_go_fuzz_dep_.CoverTab[4766]++
									primary = res
//line /usr/local/go/src/net/dial.go:501
				// _ = "end of CoverTab[4766]"
			} else {
//line /usr/local/go/src/net/dial.go:502
				_go_fuzz_dep_.CoverTab[4767]++
									fallback = res
//line /usr/local/go/src/net/dial.go:503
				// _ = "end of CoverTab[4767]"
			}
//line /usr/local/go/src/net/dial.go:504
			// _ = "end of CoverTab[4761]"
//line /usr/local/go/src/net/dial.go:504
			_go_fuzz_dep_.CoverTab[4762]++
								if primary.done && func() bool {
//line /usr/local/go/src/net/dial.go:505
				_go_fuzz_dep_.CoverTab[4768]++
//line /usr/local/go/src/net/dial.go:505
				return fallback.done
//line /usr/local/go/src/net/dial.go:505
				// _ = "end of CoverTab[4768]"
//line /usr/local/go/src/net/dial.go:505
			}() {
//line /usr/local/go/src/net/dial.go:505
				_go_fuzz_dep_.CoverTab[4769]++
									return nil, primary.error
//line /usr/local/go/src/net/dial.go:506
				// _ = "end of CoverTab[4769]"
			} else {
//line /usr/local/go/src/net/dial.go:507
				_go_fuzz_dep_.CoverTab[4770]++
//line /usr/local/go/src/net/dial.go:507
				// _ = "end of CoverTab[4770]"
//line /usr/local/go/src/net/dial.go:507
			}
//line /usr/local/go/src/net/dial.go:507
			// _ = "end of CoverTab[4762]"
//line /usr/local/go/src/net/dial.go:507
			_go_fuzz_dep_.CoverTab[4763]++
								if res.primary && func() bool {
//line /usr/local/go/src/net/dial.go:508
				_go_fuzz_dep_.CoverTab[4771]++
//line /usr/local/go/src/net/dial.go:508
				return fallbackTimer.Stop()
//line /usr/local/go/src/net/dial.go:508
				// _ = "end of CoverTab[4771]"
//line /usr/local/go/src/net/dial.go:508
			}() {
//line /usr/local/go/src/net/dial.go:508
				_go_fuzz_dep_.CoverTab[4772]++

//line /usr/local/go/src/net/dial.go:513
				fallbackTimer.Reset(0)
//line /usr/local/go/src/net/dial.go:513
				// _ = "end of CoverTab[4772]"
			} else {
//line /usr/local/go/src/net/dial.go:514
				_go_fuzz_dep_.CoverTab[4773]++
//line /usr/local/go/src/net/dial.go:514
				// _ = "end of CoverTab[4773]"
//line /usr/local/go/src/net/dial.go:514
			}
//line /usr/local/go/src/net/dial.go:514
			// _ = "end of CoverTab[4763]"
		}
//line /usr/local/go/src/net/dial.go:515
		// _ = "end of CoverTab[4758]"
	}
//line /usr/local/go/src/net/dial.go:516
	// _ = "end of CoverTab[4745]"
}

// dialSerial connects to a list of addresses in sequence, returning
//line /usr/local/go/src/net/dial.go:519
// either the first successful connection, or the first error.
//line /usr/local/go/src/net/dial.go:521
func (sd *sysDialer) dialSerial(ctx context.Context, ras addrList) (Conn, error) {
//line /usr/local/go/src/net/dial.go:521
	_go_fuzz_dep_.CoverTab[4774]++
						var firstErr error	// The error from the first address is most relevant.

						for i, ra := range ras {
//line /usr/local/go/src/net/dial.go:524
		_go_fuzz_dep_.CoverTab[4777]++
							select {
		case <-ctx.Done():
//line /usr/local/go/src/net/dial.go:526
			_go_fuzz_dep_.CoverTab[4781]++
								return nil, &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: mapErr(ctx.Err())}
//line /usr/local/go/src/net/dial.go:527
			// _ = "end of CoverTab[4781]"
		default:
//line /usr/local/go/src/net/dial.go:528
			_go_fuzz_dep_.CoverTab[4782]++
//line /usr/local/go/src/net/dial.go:528
			// _ = "end of CoverTab[4782]"
		}
//line /usr/local/go/src/net/dial.go:529
		// _ = "end of CoverTab[4777]"
//line /usr/local/go/src/net/dial.go:529
		_go_fuzz_dep_.CoverTab[4778]++

							dialCtx := ctx
							if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /usr/local/go/src/net/dial.go:532
			_go_fuzz_dep_.CoverTab[4783]++
								partialDeadline, err := partialDeadline(time.Now(), deadline, len(ras)-i)
								if err != nil {
//line /usr/local/go/src/net/dial.go:534
				_go_fuzz_dep_.CoverTab[4785]++

									if firstErr == nil {
//line /usr/local/go/src/net/dial.go:536
					_go_fuzz_dep_.CoverTab[4787]++
										firstErr = &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: err}
//line /usr/local/go/src/net/dial.go:537
					// _ = "end of CoverTab[4787]"
				} else {
//line /usr/local/go/src/net/dial.go:538
					_go_fuzz_dep_.CoverTab[4788]++
//line /usr/local/go/src/net/dial.go:538
					// _ = "end of CoverTab[4788]"
//line /usr/local/go/src/net/dial.go:538
				}
//line /usr/local/go/src/net/dial.go:538
				// _ = "end of CoverTab[4785]"
//line /usr/local/go/src/net/dial.go:538
				_go_fuzz_dep_.CoverTab[4786]++
									break
//line /usr/local/go/src/net/dial.go:539
				// _ = "end of CoverTab[4786]"
			} else {
//line /usr/local/go/src/net/dial.go:540
				_go_fuzz_dep_.CoverTab[4789]++
//line /usr/local/go/src/net/dial.go:540
				// _ = "end of CoverTab[4789]"
//line /usr/local/go/src/net/dial.go:540
			}
//line /usr/local/go/src/net/dial.go:540
			// _ = "end of CoverTab[4783]"
//line /usr/local/go/src/net/dial.go:540
			_go_fuzz_dep_.CoverTab[4784]++
								if partialDeadline.Before(deadline) {
//line /usr/local/go/src/net/dial.go:541
				_go_fuzz_dep_.CoverTab[4790]++
									var cancel context.CancelFunc
									dialCtx, cancel = context.WithDeadline(ctx, partialDeadline)
									defer cancel()
//line /usr/local/go/src/net/dial.go:544
				// _ = "end of CoverTab[4790]"
			} else {
//line /usr/local/go/src/net/dial.go:545
				_go_fuzz_dep_.CoverTab[4791]++
//line /usr/local/go/src/net/dial.go:545
				// _ = "end of CoverTab[4791]"
//line /usr/local/go/src/net/dial.go:545
			}
//line /usr/local/go/src/net/dial.go:545
			// _ = "end of CoverTab[4784]"
		} else {
//line /usr/local/go/src/net/dial.go:546
			_go_fuzz_dep_.CoverTab[4792]++
//line /usr/local/go/src/net/dial.go:546
			// _ = "end of CoverTab[4792]"
//line /usr/local/go/src/net/dial.go:546
		}
//line /usr/local/go/src/net/dial.go:546
		// _ = "end of CoverTab[4778]"
//line /usr/local/go/src/net/dial.go:546
		_go_fuzz_dep_.CoverTab[4779]++

							c, err := sd.dialSingle(dialCtx, ra)
							if err == nil {
//line /usr/local/go/src/net/dial.go:549
			_go_fuzz_dep_.CoverTab[4793]++
								return c, nil
//line /usr/local/go/src/net/dial.go:550
			// _ = "end of CoverTab[4793]"
		} else {
//line /usr/local/go/src/net/dial.go:551
			_go_fuzz_dep_.CoverTab[4794]++
//line /usr/local/go/src/net/dial.go:551
			// _ = "end of CoverTab[4794]"
//line /usr/local/go/src/net/dial.go:551
		}
//line /usr/local/go/src/net/dial.go:551
		// _ = "end of CoverTab[4779]"
//line /usr/local/go/src/net/dial.go:551
		_go_fuzz_dep_.CoverTab[4780]++
							if firstErr == nil {
//line /usr/local/go/src/net/dial.go:552
			_go_fuzz_dep_.CoverTab[4795]++
								firstErr = err
//line /usr/local/go/src/net/dial.go:553
			// _ = "end of CoverTab[4795]"
		} else {
//line /usr/local/go/src/net/dial.go:554
			_go_fuzz_dep_.CoverTab[4796]++
//line /usr/local/go/src/net/dial.go:554
			// _ = "end of CoverTab[4796]"
//line /usr/local/go/src/net/dial.go:554
		}
//line /usr/local/go/src/net/dial.go:554
		// _ = "end of CoverTab[4780]"
	}
//line /usr/local/go/src/net/dial.go:555
	// _ = "end of CoverTab[4774]"
//line /usr/local/go/src/net/dial.go:555
	_go_fuzz_dep_.CoverTab[4775]++

						if firstErr == nil {
//line /usr/local/go/src/net/dial.go:557
		_go_fuzz_dep_.CoverTab[4797]++
							firstErr = &OpError{Op: "dial", Net: sd.network, Source: nil, Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/dial.go:558
		// _ = "end of CoverTab[4797]"
	} else {
//line /usr/local/go/src/net/dial.go:559
		_go_fuzz_dep_.CoverTab[4798]++
//line /usr/local/go/src/net/dial.go:559
		// _ = "end of CoverTab[4798]"
//line /usr/local/go/src/net/dial.go:559
	}
//line /usr/local/go/src/net/dial.go:559
	// _ = "end of CoverTab[4775]"
//line /usr/local/go/src/net/dial.go:559
	_go_fuzz_dep_.CoverTab[4776]++
						return nil, firstErr
//line /usr/local/go/src/net/dial.go:560
	// _ = "end of CoverTab[4776]"
}

// dialSingle attempts to establish and returns a single connection to
//line /usr/local/go/src/net/dial.go:563
// the destination address.
//line /usr/local/go/src/net/dial.go:565
func (sd *sysDialer) dialSingle(ctx context.Context, ra Addr) (c Conn, err error) {
//line /usr/local/go/src/net/dial.go:565
	_go_fuzz_dep_.CoverTab[4799]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil {
//line /usr/local/go/src/net/dial.go:567
		_go_fuzz_dep_.CoverTab[4803]++
							raStr := ra.String()
							if trace.ConnectStart != nil {
//line /usr/local/go/src/net/dial.go:569
			_go_fuzz_dep_.CoverTab[4805]++
								trace.ConnectStart(sd.network, raStr)
//line /usr/local/go/src/net/dial.go:570
			// _ = "end of CoverTab[4805]"
		} else {
//line /usr/local/go/src/net/dial.go:571
			_go_fuzz_dep_.CoverTab[4806]++
//line /usr/local/go/src/net/dial.go:571
			// _ = "end of CoverTab[4806]"
//line /usr/local/go/src/net/dial.go:571
		}
//line /usr/local/go/src/net/dial.go:571
		// _ = "end of CoverTab[4803]"
//line /usr/local/go/src/net/dial.go:571
		_go_fuzz_dep_.CoverTab[4804]++
							if trace.ConnectDone != nil {
//line /usr/local/go/src/net/dial.go:572
			_go_fuzz_dep_.CoverTab[4807]++
								defer func() {
//line /usr/local/go/src/net/dial.go:573
				_go_fuzz_dep_.CoverTab[4808]++
//line /usr/local/go/src/net/dial.go:573
				trace.ConnectDone(sd.network, raStr, err)
//line /usr/local/go/src/net/dial.go:573
				// _ = "end of CoverTab[4808]"
//line /usr/local/go/src/net/dial.go:573
			}()
//line /usr/local/go/src/net/dial.go:573
			// _ = "end of CoverTab[4807]"
		} else {
//line /usr/local/go/src/net/dial.go:574
			_go_fuzz_dep_.CoverTab[4809]++
//line /usr/local/go/src/net/dial.go:574
			// _ = "end of CoverTab[4809]"
//line /usr/local/go/src/net/dial.go:574
		}
//line /usr/local/go/src/net/dial.go:574
		// _ = "end of CoverTab[4804]"
	} else {
//line /usr/local/go/src/net/dial.go:575
		_go_fuzz_dep_.CoverTab[4810]++
//line /usr/local/go/src/net/dial.go:575
		// _ = "end of CoverTab[4810]"
//line /usr/local/go/src/net/dial.go:575
	}
//line /usr/local/go/src/net/dial.go:575
	// _ = "end of CoverTab[4799]"
//line /usr/local/go/src/net/dial.go:575
	_go_fuzz_dep_.CoverTab[4800]++
						la := sd.LocalAddr
						switch ra := ra.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/dial.go:578
		_go_fuzz_dep_.CoverTab[4811]++
							la, _ := la.(*TCPAddr)
							c, err = sd.dialTCP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:580
		// _ = "end of CoverTab[4811]"
	case *UDPAddr:
//line /usr/local/go/src/net/dial.go:581
		_go_fuzz_dep_.CoverTab[4812]++
							la, _ := la.(*UDPAddr)
							c, err = sd.dialUDP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:583
		// _ = "end of CoverTab[4812]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:584
		_go_fuzz_dep_.CoverTab[4813]++
							la, _ := la.(*IPAddr)
							c, err = sd.dialIP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:586
		// _ = "end of CoverTab[4813]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:587
		_go_fuzz_dep_.CoverTab[4814]++
							la, _ := la.(*UnixAddr)
							c, err = sd.dialUnix(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:589
		// _ = "end of CoverTab[4814]"
	default:
//line /usr/local/go/src/net/dial.go:590
		_go_fuzz_dep_.CoverTab[4815]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: &AddrError{Err: "unexpected address type", Addr: sd.address}}
//line /usr/local/go/src/net/dial.go:591
		// _ = "end of CoverTab[4815]"
	}
//line /usr/local/go/src/net/dial.go:592
	// _ = "end of CoverTab[4800]"
//line /usr/local/go/src/net/dial.go:592
	_go_fuzz_dep_.CoverTab[4801]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:593
		_go_fuzz_dep_.CoverTab[4816]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: err}
//line /usr/local/go/src/net/dial.go:594
		// _ = "end of CoverTab[4816]"
	} else {
//line /usr/local/go/src/net/dial.go:595
		_go_fuzz_dep_.CoverTab[4817]++
//line /usr/local/go/src/net/dial.go:595
		// _ = "end of CoverTab[4817]"
//line /usr/local/go/src/net/dial.go:595
	}
//line /usr/local/go/src/net/dial.go:595
	// _ = "end of CoverTab[4801]"
//line /usr/local/go/src/net/dial.go:595
	_go_fuzz_dep_.CoverTab[4802]++
						return c, nil
//line /usr/local/go/src/net/dial.go:596
	// _ = "end of CoverTab[4802]"
}

// ListenConfig contains options for listening to an address.
type ListenConfig struct {
	// If Control is not nil, it is called after creating the network
	// connection but before binding it to the operating system.
	//
	// Network and address parameters passed to Control method are not
	// necessarily the ones passed to Listen. For example, passing "tcp" to
	// Listen will cause the Control function to be called with "tcp4" or "tcp6".
	Control	func(network, address string, c syscall.RawConn) error

	// KeepAlive specifies the keep-alive period for network
	// connections accepted by this listener.
	// If zero, keep-alives are enabled if supported by the protocol
	// and operating system. Network protocols or operating systems
	// that do not support keep-alives ignore this field.
	// If negative, keep-alives are disabled.
	KeepAlive	time.Duration
}

// Listen announces on the local network address.
//line /usr/local/go/src/net/dial.go:618
//
//line /usr/local/go/src/net/dial.go:618
// See func Listen for a description of the network and address
//line /usr/local/go/src/net/dial.go:618
// parameters.
//line /usr/local/go/src/net/dial.go:622
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error) {
//line /usr/local/go/src/net/dial.go:622
	_go_fuzz_dep_.CoverTab[4818]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /usr/local/go/src/net/dial.go:624
		_go_fuzz_dep_.CoverTab[4822]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:625
		// _ = "end of CoverTab[4822]"
	} else {
//line /usr/local/go/src/net/dial.go:626
		_go_fuzz_dep_.CoverTab[4823]++
//line /usr/local/go/src/net/dial.go:626
		// _ = "end of CoverTab[4823]"
//line /usr/local/go/src/net/dial.go:626
	}
//line /usr/local/go/src/net/dial.go:626
	// _ = "end of CoverTab[4818]"
//line /usr/local/go/src/net/dial.go:626
	_go_fuzz_dep_.CoverTab[4819]++
						sl := &sysListener{
		ListenConfig:	*lc,
		network:	network,
		address:	address,
	}
	var l Listener
	la := addrs.first(isIPv4)
	switch la := la.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/dial.go:635
		_go_fuzz_dep_.CoverTab[4824]++
							l, err = sl.listenTCP(ctx, la)
//line /usr/local/go/src/net/dial.go:636
		// _ = "end of CoverTab[4824]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:637
		_go_fuzz_dep_.CoverTab[4825]++
							l, err = sl.listenUnix(ctx, la)
//line /usr/local/go/src/net/dial.go:638
		// _ = "end of CoverTab[4825]"
	default:
//line /usr/local/go/src/net/dial.go:639
		_go_fuzz_dep_.CoverTab[4826]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /usr/local/go/src/net/dial.go:640
		// _ = "end of CoverTab[4826]"
	}
//line /usr/local/go/src/net/dial.go:641
	// _ = "end of CoverTab[4819]"
//line /usr/local/go/src/net/dial.go:641
	_go_fuzz_dep_.CoverTab[4820]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:642
		_go_fuzz_dep_.CoverTab[4827]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /usr/local/go/src/net/dial.go:643
		// _ = "end of CoverTab[4827]"
	} else {
//line /usr/local/go/src/net/dial.go:644
		_go_fuzz_dep_.CoverTab[4828]++
//line /usr/local/go/src/net/dial.go:644
		// _ = "end of CoverTab[4828]"
//line /usr/local/go/src/net/dial.go:644
	}
//line /usr/local/go/src/net/dial.go:644
	// _ = "end of CoverTab[4820]"
//line /usr/local/go/src/net/dial.go:644
	_go_fuzz_dep_.CoverTab[4821]++
						return l, nil
//line /usr/local/go/src/net/dial.go:645
	// _ = "end of CoverTab[4821]"
}

// ListenPacket announces on the local network address.
//line /usr/local/go/src/net/dial.go:648
//
//line /usr/local/go/src/net/dial.go:648
// See func ListenPacket for a description of the network and address
//line /usr/local/go/src/net/dial.go:648
// parameters.
//line /usr/local/go/src/net/dial.go:652
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error) {
//line /usr/local/go/src/net/dial.go:652
	_go_fuzz_dep_.CoverTab[4829]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /usr/local/go/src/net/dial.go:654
		_go_fuzz_dep_.CoverTab[4833]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:655
		// _ = "end of CoverTab[4833]"
	} else {
//line /usr/local/go/src/net/dial.go:656
		_go_fuzz_dep_.CoverTab[4834]++
//line /usr/local/go/src/net/dial.go:656
		// _ = "end of CoverTab[4834]"
//line /usr/local/go/src/net/dial.go:656
	}
//line /usr/local/go/src/net/dial.go:656
	// _ = "end of CoverTab[4829]"
//line /usr/local/go/src/net/dial.go:656
	_go_fuzz_dep_.CoverTab[4830]++
						sl := &sysListener{
		ListenConfig:	*lc,
		network:	network,
		address:	address,
	}
	var c PacketConn
	la := addrs.first(isIPv4)
	switch la := la.(type) {
	case *UDPAddr:
//line /usr/local/go/src/net/dial.go:665
		_go_fuzz_dep_.CoverTab[4835]++
							c, err = sl.listenUDP(ctx, la)
//line /usr/local/go/src/net/dial.go:666
		// _ = "end of CoverTab[4835]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:667
		_go_fuzz_dep_.CoverTab[4836]++
							c, err = sl.listenIP(ctx, la)
//line /usr/local/go/src/net/dial.go:668
		// _ = "end of CoverTab[4836]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:669
		_go_fuzz_dep_.CoverTab[4837]++
							c, err = sl.listenUnixgram(ctx, la)
//line /usr/local/go/src/net/dial.go:670
		// _ = "end of CoverTab[4837]"
	default:
//line /usr/local/go/src/net/dial.go:671
		_go_fuzz_dep_.CoverTab[4838]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /usr/local/go/src/net/dial.go:672
		// _ = "end of CoverTab[4838]"
	}
//line /usr/local/go/src/net/dial.go:673
	// _ = "end of CoverTab[4830]"
//line /usr/local/go/src/net/dial.go:673
	_go_fuzz_dep_.CoverTab[4831]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:674
		_go_fuzz_dep_.CoverTab[4839]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /usr/local/go/src/net/dial.go:675
		// _ = "end of CoverTab[4839]"
	} else {
//line /usr/local/go/src/net/dial.go:676
		_go_fuzz_dep_.CoverTab[4840]++
//line /usr/local/go/src/net/dial.go:676
		// _ = "end of CoverTab[4840]"
//line /usr/local/go/src/net/dial.go:676
	}
//line /usr/local/go/src/net/dial.go:676
	// _ = "end of CoverTab[4831]"
//line /usr/local/go/src/net/dial.go:676
	_go_fuzz_dep_.CoverTab[4832]++
						return c, nil
//line /usr/local/go/src/net/dial.go:677
	// _ = "end of CoverTab[4832]"
}

// sysListener contains a Listen's parameters and configuration.
type sysListener struct {
	ListenConfig
	network, address	string
}

// Listen announces on the local network address.
//line /usr/local/go/src/net/dial.go:686
//
//line /usr/local/go/src/net/dial.go:686
// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
//line /usr/local/go/src/net/dial.go:686
//
//line /usr/local/go/src/net/dial.go:686
// For TCP networks, if the host in the address parameter is empty or
//line /usr/local/go/src/net/dial.go:686
// a literal unspecified IP address, Listen listens on all available
//line /usr/local/go/src/net/dial.go:686
// unicast and anycast IP addresses of the local system.
//line /usr/local/go/src/net/dial.go:686
// To only use IPv4, use network "tcp4".
//line /usr/local/go/src/net/dial.go:686
// The address can use a host name, but this is not recommended,
//line /usr/local/go/src/net/dial.go:686
// because it will create a listener for at most one of the host's IP
//line /usr/local/go/src/net/dial.go:686
// addresses.
//line /usr/local/go/src/net/dial.go:686
// If the port in the address parameter is empty or "0", as in
//line /usr/local/go/src/net/dial.go:686
// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
//line /usr/local/go/src/net/dial.go:686
// The Addr method of Listener can be used to discover the chosen
//line /usr/local/go/src/net/dial.go:686
// port.
//line /usr/local/go/src/net/dial.go:686
//
//line /usr/local/go/src/net/dial.go:686
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/dial.go:686
// parameters.
//line /usr/local/go/src/net/dial.go:686
//
//line /usr/local/go/src/net/dial.go:686
// Listen uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/dial.go:686
// ListenConfig.Listen.
//line /usr/local/go/src/net/dial.go:707
func Listen(network, address string) (Listener, error) {
//line /usr/local/go/src/net/dial.go:707
	_go_fuzz_dep_.CoverTab[4841]++
						var lc ListenConfig
						return lc.Listen(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:709
	// _ = "end of CoverTab[4841]"
}

// ListenPacket announces on the local network address.
//line /usr/local/go/src/net/dial.go:712
//
//line /usr/local/go/src/net/dial.go:712
// The network must be "udp", "udp4", "udp6", "unixgram", or an IP
//line /usr/local/go/src/net/dial.go:712
// transport. The IP transports are "ip", "ip4", or "ip6" followed by
//line /usr/local/go/src/net/dial.go:712
// a colon and a literal protocol number or a protocol name, as in
//line /usr/local/go/src/net/dial.go:712
// "ip:1" or "ip:icmp".
//line /usr/local/go/src/net/dial.go:712
//
//line /usr/local/go/src/net/dial.go:712
// For UDP and IP networks, if the host in the address parameter is
//line /usr/local/go/src/net/dial.go:712
// empty or a literal unspecified IP address, ListenPacket listens on
//line /usr/local/go/src/net/dial.go:712
// all available IP addresses of the local system except multicast IP
//line /usr/local/go/src/net/dial.go:712
// addresses.
//line /usr/local/go/src/net/dial.go:712
// To only use IPv4, use network "udp4" or "ip4:proto".
//line /usr/local/go/src/net/dial.go:712
// The address can use a host name, but this is not recommended,
//line /usr/local/go/src/net/dial.go:712
// because it will create a listener for at most one of the host's IP
//line /usr/local/go/src/net/dial.go:712
// addresses.
//line /usr/local/go/src/net/dial.go:712
// If the port in the address parameter is empty or "0", as in
//line /usr/local/go/src/net/dial.go:712
// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
//line /usr/local/go/src/net/dial.go:712
// The LocalAddr method of PacketConn can be used to discover the
//line /usr/local/go/src/net/dial.go:712
// chosen port.
//line /usr/local/go/src/net/dial.go:712
//
//line /usr/local/go/src/net/dial.go:712
// See func Dial for a description of the network and address
//line /usr/local/go/src/net/dial.go:712
// parameters.
//line /usr/local/go/src/net/dial.go:712
//
//line /usr/local/go/src/net/dial.go:712
// ListenPacket uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/dial.go:712
// ListenConfig.ListenPacket.
//line /usr/local/go/src/net/dial.go:737
func ListenPacket(network, address string) (PacketConn, error) {
//line /usr/local/go/src/net/dial.go:737
	_go_fuzz_dep_.CoverTab[4842]++
						var lc ListenConfig
						return lc.ListenPacket(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:739
	// _ = "end of CoverTab[4842]"
}

//line /usr/local/go/src/net/dial.go:740
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dial.go:740
var _ = _go_fuzz_dep_.CoverTab
