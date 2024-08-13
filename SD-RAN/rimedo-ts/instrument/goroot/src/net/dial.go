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
	_go_fuzz_dep_.CoverTab[12992]++
//line /usr/local/go/src/net/dial.go:113
	return d.FallbackDelay >= 0
//line /usr/local/go/src/net/dial.go:113
	// _ = "end of CoverTab[12992]"
//line /usr/local/go/src/net/dial.go:113
}

func minNonzeroTime(a, b time.Time) time.Time {
//line /usr/local/go/src/net/dial.go:115
	_go_fuzz_dep_.CoverTab[12993]++
						if a.IsZero() {
//line /usr/local/go/src/net/dial.go:116
		_go_fuzz_dep_.CoverTab[12996]++
							return b
//line /usr/local/go/src/net/dial.go:117
		// _ = "end of CoverTab[12996]"
	} else {
//line /usr/local/go/src/net/dial.go:118
		_go_fuzz_dep_.CoverTab[12997]++
//line /usr/local/go/src/net/dial.go:118
		// _ = "end of CoverTab[12997]"
//line /usr/local/go/src/net/dial.go:118
	}
//line /usr/local/go/src/net/dial.go:118
	// _ = "end of CoverTab[12993]"
//line /usr/local/go/src/net/dial.go:118
	_go_fuzz_dep_.CoverTab[12994]++
						if b.IsZero() || func() bool {
//line /usr/local/go/src/net/dial.go:119
		_go_fuzz_dep_.CoverTab[12998]++
//line /usr/local/go/src/net/dial.go:119
		return a.Before(b)
//line /usr/local/go/src/net/dial.go:119
		// _ = "end of CoverTab[12998]"
//line /usr/local/go/src/net/dial.go:119
	}() {
//line /usr/local/go/src/net/dial.go:119
		_go_fuzz_dep_.CoverTab[12999]++
							return a
//line /usr/local/go/src/net/dial.go:120
		// _ = "end of CoverTab[12999]"
	} else {
//line /usr/local/go/src/net/dial.go:121
		_go_fuzz_dep_.CoverTab[13000]++
//line /usr/local/go/src/net/dial.go:121
		// _ = "end of CoverTab[13000]"
//line /usr/local/go/src/net/dial.go:121
	}
//line /usr/local/go/src/net/dial.go:121
	// _ = "end of CoverTab[12994]"
//line /usr/local/go/src/net/dial.go:121
	_go_fuzz_dep_.CoverTab[12995]++
						return b
//line /usr/local/go/src/net/dial.go:122
	// _ = "end of CoverTab[12995]"
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
	_go_fuzz_dep_.CoverTab[13001]++
						if d.Timeout != 0 {
//line /usr/local/go/src/net/dial.go:132
		_go_fuzz_dep_.CoverTab[13004]++
							earliest = now.Add(d.Timeout)
//line /usr/local/go/src/net/dial.go:133
		// _ = "end of CoverTab[13004]"
	} else {
//line /usr/local/go/src/net/dial.go:134
		_go_fuzz_dep_.CoverTab[13005]++
//line /usr/local/go/src/net/dial.go:134
		// _ = "end of CoverTab[13005]"
//line /usr/local/go/src/net/dial.go:134
	}
//line /usr/local/go/src/net/dial.go:134
	// _ = "end of CoverTab[13001]"
//line /usr/local/go/src/net/dial.go:134
	_go_fuzz_dep_.CoverTab[13002]++
						if d, ok := ctx.Deadline(); ok {
//line /usr/local/go/src/net/dial.go:135
		_go_fuzz_dep_.CoverTab[13006]++
							earliest = minNonzeroTime(earliest, d)
//line /usr/local/go/src/net/dial.go:136
		// _ = "end of CoverTab[13006]"
	} else {
//line /usr/local/go/src/net/dial.go:137
		_go_fuzz_dep_.CoverTab[13007]++
//line /usr/local/go/src/net/dial.go:137
		// _ = "end of CoverTab[13007]"
//line /usr/local/go/src/net/dial.go:137
	}
//line /usr/local/go/src/net/dial.go:137
	// _ = "end of CoverTab[13002]"
//line /usr/local/go/src/net/dial.go:137
	_go_fuzz_dep_.CoverTab[13003]++
						return minNonzeroTime(earliest, d.Deadline)
//line /usr/local/go/src/net/dial.go:138
	// _ = "end of CoverTab[13003]"
}

func (d *Dialer) resolver() *Resolver {
//line /usr/local/go/src/net/dial.go:141
	_go_fuzz_dep_.CoverTab[13008]++
						if d.Resolver != nil {
//line /usr/local/go/src/net/dial.go:142
		_go_fuzz_dep_.CoverTab[13010]++
							return d.Resolver
//line /usr/local/go/src/net/dial.go:143
		// _ = "end of CoverTab[13010]"
	} else {
//line /usr/local/go/src/net/dial.go:144
		_go_fuzz_dep_.CoverTab[13011]++
//line /usr/local/go/src/net/dial.go:144
		// _ = "end of CoverTab[13011]"
//line /usr/local/go/src/net/dial.go:144
	}
//line /usr/local/go/src/net/dial.go:144
	// _ = "end of CoverTab[13008]"
//line /usr/local/go/src/net/dial.go:144
	_go_fuzz_dep_.CoverTab[13009]++
						return DefaultResolver
//line /usr/local/go/src/net/dial.go:145
	// _ = "end of CoverTab[13009]"
}

// partialDeadline returns the deadline to use for a single address,
//line /usr/local/go/src/net/dial.go:148
// when multiple addresses are pending.
//line /usr/local/go/src/net/dial.go:150
func partialDeadline(now, deadline time.Time, addrsRemaining int) (time.Time, error) {
//line /usr/local/go/src/net/dial.go:150
	_go_fuzz_dep_.CoverTab[13012]++
						if deadline.IsZero() {
//line /usr/local/go/src/net/dial.go:151
		_go_fuzz_dep_.CoverTab[13016]++
							return deadline, nil
//line /usr/local/go/src/net/dial.go:152
		// _ = "end of CoverTab[13016]"
	} else {
//line /usr/local/go/src/net/dial.go:153
		_go_fuzz_dep_.CoverTab[13017]++
//line /usr/local/go/src/net/dial.go:153
		// _ = "end of CoverTab[13017]"
//line /usr/local/go/src/net/dial.go:153
	}
//line /usr/local/go/src/net/dial.go:153
	// _ = "end of CoverTab[13012]"
//line /usr/local/go/src/net/dial.go:153
	_go_fuzz_dep_.CoverTab[13013]++
						timeRemaining := deadline.Sub(now)
						if timeRemaining <= 0 {
//line /usr/local/go/src/net/dial.go:155
		_go_fuzz_dep_.CoverTab[13018]++
							return time.Time{}, errTimeout
//line /usr/local/go/src/net/dial.go:156
		// _ = "end of CoverTab[13018]"
	} else {
//line /usr/local/go/src/net/dial.go:157
		_go_fuzz_dep_.CoverTab[13019]++
//line /usr/local/go/src/net/dial.go:157
		// _ = "end of CoverTab[13019]"
//line /usr/local/go/src/net/dial.go:157
	}
//line /usr/local/go/src/net/dial.go:157
	// _ = "end of CoverTab[13013]"
//line /usr/local/go/src/net/dial.go:157
	_go_fuzz_dep_.CoverTab[13014]++

						timeout := timeRemaining / time.Duration(addrsRemaining)
	// If the time per address is too short, steal from the end of the list.
	const saneMinimum = 2 * time.Second
	if timeout < saneMinimum {
//line /usr/local/go/src/net/dial.go:162
		_go_fuzz_dep_.CoverTab[13020]++
							if timeRemaining < saneMinimum {
//line /usr/local/go/src/net/dial.go:163
			_go_fuzz_dep_.CoverTab[13021]++
								timeout = timeRemaining
//line /usr/local/go/src/net/dial.go:164
			// _ = "end of CoverTab[13021]"
		} else {
//line /usr/local/go/src/net/dial.go:165
			_go_fuzz_dep_.CoverTab[13022]++
								timeout = saneMinimum
//line /usr/local/go/src/net/dial.go:166
			// _ = "end of CoverTab[13022]"
		}
//line /usr/local/go/src/net/dial.go:167
		// _ = "end of CoverTab[13020]"
	} else {
//line /usr/local/go/src/net/dial.go:168
		_go_fuzz_dep_.CoverTab[13023]++
//line /usr/local/go/src/net/dial.go:168
		// _ = "end of CoverTab[13023]"
//line /usr/local/go/src/net/dial.go:168
	}
//line /usr/local/go/src/net/dial.go:168
	// _ = "end of CoverTab[13014]"
//line /usr/local/go/src/net/dial.go:168
	_go_fuzz_dep_.CoverTab[13015]++
						return now.Add(timeout), nil
//line /usr/local/go/src/net/dial.go:169
	// _ = "end of CoverTab[13015]"
}

func (d *Dialer) fallbackDelay() time.Duration {
//line /usr/local/go/src/net/dial.go:172
	_go_fuzz_dep_.CoverTab[13024]++
						if d.FallbackDelay > 0 {
//line /usr/local/go/src/net/dial.go:173
		_go_fuzz_dep_.CoverTab[13025]++
							return d.FallbackDelay
//line /usr/local/go/src/net/dial.go:174
		// _ = "end of CoverTab[13025]"
	} else {
//line /usr/local/go/src/net/dial.go:175
		_go_fuzz_dep_.CoverTab[13026]++
							return 300 * time.Millisecond
//line /usr/local/go/src/net/dial.go:176
		// _ = "end of CoverTab[13026]"
	}
//line /usr/local/go/src/net/dial.go:177
	// _ = "end of CoverTab[13024]"
}

func parseNetwork(ctx context.Context, network string, needsProto bool) (afnet string, proto int, err error) {
//line /usr/local/go/src/net/dial.go:180
	_go_fuzz_dep_.CoverTab[13027]++
						i := last(network, ':')
						if i < 0 {
//line /usr/local/go/src/net/dial.go:182
		_go_fuzz_dep_.CoverTab[13030]++
							switch network {
		case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/dial.go:184
			_go_fuzz_dep_.CoverTab[13032]++
//line /usr/local/go/src/net/dial.go:184
			// _ = "end of CoverTab[13032]"
		case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/dial.go:185
			_go_fuzz_dep_.CoverTab[13033]++
//line /usr/local/go/src/net/dial.go:185
			// _ = "end of CoverTab[13033]"
		case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/dial.go:186
			_go_fuzz_dep_.CoverTab[13034]++
								if needsProto {
//line /usr/local/go/src/net/dial.go:187
				_go_fuzz_dep_.CoverTab[13037]++
									return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:188
				// _ = "end of CoverTab[13037]"
			} else {
//line /usr/local/go/src/net/dial.go:189
				_go_fuzz_dep_.CoverTab[13038]++
//line /usr/local/go/src/net/dial.go:189
				// _ = "end of CoverTab[13038]"
//line /usr/local/go/src/net/dial.go:189
			}
//line /usr/local/go/src/net/dial.go:189
			// _ = "end of CoverTab[13034]"
		case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/dial.go:190
			_go_fuzz_dep_.CoverTab[13035]++
//line /usr/local/go/src/net/dial.go:190
			// _ = "end of CoverTab[13035]"
		default:
//line /usr/local/go/src/net/dial.go:191
			_go_fuzz_dep_.CoverTab[13036]++
								return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:192
			// _ = "end of CoverTab[13036]"
		}
//line /usr/local/go/src/net/dial.go:193
		// _ = "end of CoverTab[13030]"
//line /usr/local/go/src/net/dial.go:193
		_go_fuzz_dep_.CoverTab[13031]++
							return network, 0, nil
//line /usr/local/go/src/net/dial.go:194
		// _ = "end of CoverTab[13031]"
	} else {
//line /usr/local/go/src/net/dial.go:195
		_go_fuzz_dep_.CoverTab[13039]++
//line /usr/local/go/src/net/dial.go:195
		// _ = "end of CoverTab[13039]"
//line /usr/local/go/src/net/dial.go:195
	}
//line /usr/local/go/src/net/dial.go:195
	// _ = "end of CoverTab[13027]"
//line /usr/local/go/src/net/dial.go:195
	_go_fuzz_dep_.CoverTab[13028]++
						afnet = network[:i]
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/dial.go:198
		_go_fuzz_dep_.CoverTab[13040]++
							protostr := network[i+1:]
							proto, i, ok := dtoi(protostr)
							if !ok || func() bool {
//line /usr/local/go/src/net/dial.go:201
			_go_fuzz_dep_.CoverTab[13043]++
//line /usr/local/go/src/net/dial.go:201
			return i != len(protostr)
//line /usr/local/go/src/net/dial.go:201
			// _ = "end of CoverTab[13043]"
//line /usr/local/go/src/net/dial.go:201
		}() {
//line /usr/local/go/src/net/dial.go:201
			_go_fuzz_dep_.CoverTab[13044]++
								proto, err = lookupProtocol(ctx, protostr)
								if err != nil {
//line /usr/local/go/src/net/dial.go:203
				_go_fuzz_dep_.CoverTab[13045]++
									return "", 0, err
//line /usr/local/go/src/net/dial.go:204
				// _ = "end of CoverTab[13045]"
			} else {
//line /usr/local/go/src/net/dial.go:205
				_go_fuzz_dep_.CoverTab[13046]++
//line /usr/local/go/src/net/dial.go:205
				// _ = "end of CoverTab[13046]"
//line /usr/local/go/src/net/dial.go:205
			}
//line /usr/local/go/src/net/dial.go:205
			// _ = "end of CoverTab[13044]"
		} else {
//line /usr/local/go/src/net/dial.go:206
			_go_fuzz_dep_.CoverTab[13047]++
//line /usr/local/go/src/net/dial.go:206
			// _ = "end of CoverTab[13047]"
//line /usr/local/go/src/net/dial.go:206
		}
//line /usr/local/go/src/net/dial.go:206
		// _ = "end of CoverTab[13040]"
//line /usr/local/go/src/net/dial.go:206
		_go_fuzz_dep_.CoverTab[13041]++
							return afnet, proto, nil
//line /usr/local/go/src/net/dial.go:207
		// _ = "end of CoverTab[13041]"
//line /usr/local/go/src/net/dial.go:207
	default:
//line /usr/local/go/src/net/dial.go:207
		_go_fuzz_dep_.CoverTab[13042]++
//line /usr/local/go/src/net/dial.go:207
		// _ = "end of CoverTab[13042]"
	}
//line /usr/local/go/src/net/dial.go:208
	// _ = "end of CoverTab[13028]"
//line /usr/local/go/src/net/dial.go:208
	_go_fuzz_dep_.CoverTab[13029]++
						return "", 0, UnknownNetworkError(network)
//line /usr/local/go/src/net/dial.go:209
	// _ = "end of CoverTab[13029]"
}

// resolveAddrList resolves addr using hint and returns a list of
//line /usr/local/go/src/net/dial.go:212
// addresses. The result contains at least one address when error is
//line /usr/local/go/src/net/dial.go:212
// nil.
//line /usr/local/go/src/net/dial.go:215
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, addr string, hint Addr) (addrList, error) {
//line /usr/local/go/src/net/dial.go:215
	_go_fuzz_dep_.CoverTab[13048]++
						afnet, _, err := parseNetwork(ctx, network, true)
						if err != nil {
//line /usr/local/go/src/net/dial.go:217
		_go_fuzz_dep_.CoverTab[13056]++
							return nil, err
//line /usr/local/go/src/net/dial.go:218
		// _ = "end of CoverTab[13056]"
	} else {
//line /usr/local/go/src/net/dial.go:219
		_go_fuzz_dep_.CoverTab[13057]++
//line /usr/local/go/src/net/dial.go:219
		// _ = "end of CoverTab[13057]"
//line /usr/local/go/src/net/dial.go:219
	}
//line /usr/local/go/src/net/dial.go:219
	// _ = "end of CoverTab[13048]"
//line /usr/local/go/src/net/dial.go:219
	_go_fuzz_dep_.CoverTab[13049]++
						if op == "dial" && func() bool {
//line /usr/local/go/src/net/dial.go:220
		_go_fuzz_dep_.CoverTab[13058]++
//line /usr/local/go/src/net/dial.go:220
		return addr == ""
//line /usr/local/go/src/net/dial.go:220
		// _ = "end of CoverTab[13058]"
//line /usr/local/go/src/net/dial.go:220
	}() {
//line /usr/local/go/src/net/dial.go:220
		_go_fuzz_dep_.CoverTab[13059]++
							return nil, errMissingAddress
//line /usr/local/go/src/net/dial.go:221
		// _ = "end of CoverTab[13059]"
	} else {
//line /usr/local/go/src/net/dial.go:222
		_go_fuzz_dep_.CoverTab[13060]++
//line /usr/local/go/src/net/dial.go:222
		// _ = "end of CoverTab[13060]"
//line /usr/local/go/src/net/dial.go:222
	}
//line /usr/local/go/src/net/dial.go:222
	// _ = "end of CoverTab[13049]"
//line /usr/local/go/src/net/dial.go:222
	_go_fuzz_dep_.CoverTab[13050]++
						switch afnet {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/dial.go:224
		_go_fuzz_dep_.CoverTab[13061]++
							addr, err := ResolveUnixAddr(afnet, addr)
							if err != nil {
//line /usr/local/go/src/net/dial.go:226
			_go_fuzz_dep_.CoverTab[13065]++
								return nil, err
//line /usr/local/go/src/net/dial.go:227
			// _ = "end of CoverTab[13065]"
		} else {
//line /usr/local/go/src/net/dial.go:228
			_go_fuzz_dep_.CoverTab[13066]++
//line /usr/local/go/src/net/dial.go:228
			// _ = "end of CoverTab[13066]"
//line /usr/local/go/src/net/dial.go:228
		}
//line /usr/local/go/src/net/dial.go:228
		// _ = "end of CoverTab[13061]"
//line /usr/local/go/src/net/dial.go:228
		_go_fuzz_dep_.CoverTab[13062]++
							if op == "dial" && func() bool {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[13067]++
//line /usr/local/go/src/net/dial.go:229
			return hint != nil
//line /usr/local/go/src/net/dial.go:229
			// _ = "end of CoverTab[13067]"
//line /usr/local/go/src/net/dial.go:229
		}() && func() bool {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[13068]++
//line /usr/local/go/src/net/dial.go:229
			return addr.Network() != hint.Network()
//line /usr/local/go/src/net/dial.go:229
			// _ = "end of CoverTab[13068]"
//line /usr/local/go/src/net/dial.go:229
		}() {
//line /usr/local/go/src/net/dial.go:229
			_go_fuzz_dep_.CoverTab[13069]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:230
			// _ = "end of CoverTab[13069]"
		} else {
//line /usr/local/go/src/net/dial.go:231
			_go_fuzz_dep_.CoverTab[13070]++
//line /usr/local/go/src/net/dial.go:231
			// _ = "end of CoverTab[13070]"
//line /usr/local/go/src/net/dial.go:231
		}
//line /usr/local/go/src/net/dial.go:231
		// _ = "end of CoverTab[13062]"
//line /usr/local/go/src/net/dial.go:231
		_go_fuzz_dep_.CoverTab[13063]++
							return addrList{addr}, nil
//line /usr/local/go/src/net/dial.go:232
		// _ = "end of CoverTab[13063]"
//line /usr/local/go/src/net/dial.go:232
	default:
//line /usr/local/go/src/net/dial.go:232
		_go_fuzz_dep_.CoverTab[13064]++
//line /usr/local/go/src/net/dial.go:232
		// _ = "end of CoverTab[13064]"
	}
//line /usr/local/go/src/net/dial.go:233
	// _ = "end of CoverTab[13050]"
//line /usr/local/go/src/net/dial.go:233
	_go_fuzz_dep_.CoverTab[13051]++
						addrs, err := r.internetAddrList(ctx, afnet, addr)
						if err != nil || func() bool {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[13071]++
//line /usr/local/go/src/net/dial.go:235
		return op != "dial"
//line /usr/local/go/src/net/dial.go:235
		// _ = "end of CoverTab[13071]"
//line /usr/local/go/src/net/dial.go:235
	}() || func() bool {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[13072]++
//line /usr/local/go/src/net/dial.go:235
		return hint == nil
//line /usr/local/go/src/net/dial.go:235
		// _ = "end of CoverTab[13072]"
//line /usr/local/go/src/net/dial.go:235
	}() {
//line /usr/local/go/src/net/dial.go:235
		_go_fuzz_dep_.CoverTab[13073]++
							return addrs, err
//line /usr/local/go/src/net/dial.go:236
		// _ = "end of CoverTab[13073]"
	} else {
//line /usr/local/go/src/net/dial.go:237
		_go_fuzz_dep_.CoverTab[13074]++
//line /usr/local/go/src/net/dial.go:237
		// _ = "end of CoverTab[13074]"
//line /usr/local/go/src/net/dial.go:237
	}
//line /usr/local/go/src/net/dial.go:237
	// _ = "end of CoverTab[13051]"
//line /usr/local/go/src/net/dial.go:237
	_go_fuzz_dep_.CoverTab[13052]++
						var (
		tcp		*TCPAddr
		udp		*UDPAddr
		ip		*IPAddr
		wildcard	bool
	)
	switch hint := hint.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/dial.go:245
		_go_fuzz_dep_.CoverTab[13075]++
							tcp = hint
							wildcard = tcp.isWildcard()
//line /usr/local/go/src/net/dial.go:247
		// _ = "end of CoverTab[13075]"
	case *UDPAddr:
//line /usr/local/go/src/net/dial.go:248
		_go_fuzz_dep_.CoverTab[13076]++
							udp = hint
							wildcard = udp.isWildcard()
//line /usr/local/go/src/net/dial.go:250
		// _ = "end of CoverTab[13076]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:251
		_go_fuzz_dep_.CoverTab[13077]++
							ip = hint
							wildcard = ip.isWildcard()
//line /usr/local/go/src/net/dial.go:253
		// _ = "end of CoverTab[13077]"
	}
//line /usr/local/go/src/net/dial.go:254
	// _ = "end of CoverTab[13052]"
//line /usr/local/go/src/net/dial.go:254
	_go_fuzz_dep_.CoverTab[13053]++
						naddrs := addrs[:0]
						for _, addr := range addrs {
//line /usr/local/go/src/net/dial.go:256
		_go_fuzz_dep_.CoverTab[13078]++
							if addr.Network() != hint.Network() {
//line /usr/local/go/src/net/dial.go:257
			_go_fuzz_dep_.CoverTab[13080]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:258
			// _ = "end of CoverTab[13080]"
		} else {
//line /usr/local/go/src/net/dial.go:259
			_go_fuzz_dep_.CoverTab[13081]++
//line /usr/local/go/src/net/dial.go:259
			// _ = "end of CoverTab[13081]"
//line /usr/local/go/src/net/dial.go:259
		}
//line /usr/local/go/src/net/dial.go:259
		// _ = "end of CoverTab[13078]"
//line /usr/local/go/src/net/dial.go:259
		_go_fuzz_dep_.CoverTab[13079]++
							switch addr := addr.(type) {
		case *TCPAddr:
//line /usr/local/go/src/net/dial.go:261
			_go_fuzz_dep_.CoverTab[13082]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[13088]++
//line /usr/local/go/src/net/dial.go:262
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:262
				// _ = "end of CoverTab[13088]"
//line /usr/local/go/src/net/dial.go:262
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[13089]++
//line /usr/local/go/src/net/dial.go:262
				return !addr.IP.matchAddrFamily(tcp.IP)
//line /usr/local/go/src/net/dial.go:262
				// _ = "end of CoverTab[13089]"
//line /usr/local/go/src/net/dial.go:262
			}() {
//line /usr/local/go/src/net/dial.go:262
				_go_fuzz_dep_.CoverTab[13090]++
									continue
//line /usr/local/go/src/net/dial.go:263
				// _ = "end of CoverTab[13090]"
			} else {
//line /usr/local/go/src/net/dial.go:264
				_go_fuzz_dep_.CoverTab[13091]++
//line /usr/local/go/src/net/dial.go:264
				// _ = "end of CoverTab[13091]"
//line /usr/local/go/src/net/dial.go:264
			}
//line /usr/local/go/src/net/dial.go:264
			// _ = "end of CoverTab[13082]"
//line /usr/local/go/src/net/dial.go:264
			_go_fuzz_dep_.CoverTab[13083]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:265
			// _ = "end of CoverTab[13083]"
		case *UDPAddr:
//line /usr/local/go/src/net/dial.go:266
			_go_fuzz_dep_.CoverTab[13084]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[13092]++
//line /usr/local/go/src/net/dial.go:267
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:267
				// _ = "end of CoverTab[13092]"
//line /usr/local/go/src/net/dial.go:267
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[13093]++
//line /usr/local/go/src/net/dial.go:267
				return !addr.IP.matchAddrFamily(udp.IP)
//line /usr/local/go/src/net/dial.go:267
				// _ = "end of CoverTab[13093]"
//line /usr/local/go/src/net/dial.go:267
			}() {
//line /usr/local/go/src/net/dial.go:267
				_go_fuzz_dep_.CoverTab[13094]++
									continue
//line /usr/local/go/src/net/dial.go:268
				// _ = "end of CoverTab[13094]"
			} else {
//line /usr/local/go/src/net/dial.go:269
				_go_fuzz_dep_.CoverTab[13095]++
//line /usr/local/go/src/net/dial.go:269
				// _ = "end of CoverTab[13095]"
//line /usr/local/go/src/net/dial.go:269
			}
//line /usr/local/go/src/net/dial.go:269
			// _ = "end of CoverTab[13084]"
//line /usr/local/go/src/net/dial.go:269
			_go_fuzz_dep_.CoverTab[13085]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:270
			// _ = "end of CoverTab[13085]"
		case *IPAddr:
//line /usr/local/go/src/net/dial.go:271
			_go_fuzz_dep_.CoverTab[13086]++
								if !wildcard && func() bool {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[13096]++
//line /usr/local/go/src/net/dial.go:272
				return !addr.isWildcard()
//line /usr/local/go/src/net/dial.go:272
				// _ = "end of CoverTab[13096]"
//line /usr/local/go/src/net/dial.go:272
			}() && func() bool {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[13097]++
//line /usr/local/go/src/net/dial.go:272
				return !addr.IP.matchAddrFamily(ip.IP)
//line /usr/local/go/src/net/dial.go:272
				// _ = "end of CoverTab[13097]"
//line /usr/local/go/src/net/dial.go:272
			}() {
//line /usr/local/go/src/net/dial.go:272
				_go_fuzz_dep_.CoverTab[13098]++
									continue
//line /usr/local/go/src/net/dial.go:273
				// _ = "end of CoverTab[13098]"
			} else {
//line /usr/local/go/src/net/dial.go:274
				_go_fuzz_dep_.CoverTab[13099]++
//line /usr/local/go/src/net/dial.go:274
				// _ = "end of CoverTab[13099]"
//line /usr/local/go/src/net/dial.go:274
			}
//line /usr/local/go/src/net/dial.go:274
			// _ = "end of CoverTab[13086]"
//line /usr/local/go/src/net/dial.go:274
			_go_fuzz_dep_.CoverTab[13087]++
								naddrs = append(naddrs, addr)
//line /usr/local/go/src/net/dial.go:275
			// _ = "end of CoverTab[13087]"
		}
//line /usr/local/go/src/net/dial.go:276
		// _ = "end of CoverTab[13079]"
	}
//line /usr/local/go/src/net/dial.go:277
	// _ = "end of CoverTab[13053]"
//line /usr/local/go/src/net/dial.go:277
	_go_fuzz_dep_.CoverTab[13054]++
						if len(naddrs) == 0 {
//line /usr/local/go/src/net/dial.go:278
		_go_fuzz_dep_.CoverTab[13100]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: hint.String()}
//line /usr/local/go/src/net/dial.go:279
		// _ = "end of CoverTab[13100]"
	} else {
//line /usr/local/go/src/net/dial.go:280
		_go_fuzz_dep_.CoverTab[13101]++
//line /usr/local/go/src/net/dial.go:280
		// _ = "end of CoverTab[13101]"
//line /usr/local/go/src/net/dial.go:280
	}
//line /usr/local/go/src/net/dial.go:280
	// _ = "end of CoverTab[13054]"
//line /usr/local/go/src/net/dial.go:280
	_go_fuzz_dep_.CoverTab[13055]++
						return naddrs, nil
//line /usr/local/go/src/net/dial.go:281
	// _ = "end of CoverTab[13055]"
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
	_go_fuzz_dep_.CoverTab[13102]++
						var d Dialer
						return d.Dial(network, address)
//line /usr/local/go/src/net/dial.go:334
	// _ = "end of CoverTab[13102]"
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
	_go_fuzz_dep_.CoverTab[13103]++
						d := Dialer{Timeout: timeout}
						return d.Dial(network, address)
//line /usr/local/go/src/net/dial.go:349
	// _ = "end of CoverTab[13103]"
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
	_go_fuzz_dep_.CoverTab[13104]++
						return d.DialContext(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:367
	// _ = "end of CoverTab[13104]"
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
	_go_fuzz_dep_.CoverTab[13105]++
						if ctx == nil {
//line /usr/local/go/src/net/dial.go:389
		_go_fuzz_dep_.CoverTab[13112]++
							panic("nil context")
//line /usr/local/go/src/net/dial.go:390
		// _ = "end of CoverTab[13112]"
	} else {
//line /usr/local/go/src/net/dial.go:391
		_go_fuzz_dep_.CoverTab[13113]++
//line /usr/local/go/src/net/dial.go:391
		// _ = "end of CoverTab[13113]"
//line /usr/local/go/src/net/dial.go:391
	}
//line /usr/local/go/src/net/dial.go:391
	// _ = "end of CoverTab[13105]"
//line /usr/local/go/src/net/dial.go:391
	_go_fuzz_dep_.CoverTab[13106]++
						deadline := d.deadline(ctx, time.Now())
						if !deadline.IsZero() {
//line /usr/local/go/src/net/dial.go:393
		_go_fuzz_dep_.CoverTab[13114]++
							if d, ok := ctx.Deadline(); !ok || func() bool {
//line /usr/local/go/src/net/dial.go:394
			_go_fuzz_dep_.CoverTab[13115]++
//line /usr/local/go/src/net/dial.go:394
			return deadline.Before(d)
//line /usr/local/go/src/net/dial.go:394
			// _ = "end of CoverTab[13115]"
//line /usr/local/go/src/net/dial.go:394
		}() {
//line /usr/local/go/src/net/dial.go:394
			_go_fuzz_dep_.CoverTab[13116]++
								subCtx, cancel := context.WithDeadline(ctx, deadline)
								defer cancel()
								ctx = subCtx
//line /usr/local/go/src/net/dial.go:397
			// _ = "end of CoverTab[13116]"
		} else {
//line /usr/local/go/src/net/dial.go:398
			_go_fuzz_dep_.CoverTab[13117]++
//line /usr/local/go/src/net/dial.go:398
			// _ = "end of CoverTab[13117]"
//line /usr/local/go/src/net/dial.go:398
		}
//line /usr/local/go/src/net/dial.go:398
		// _ = "end of CoverTab[13114]"
	} else {
//line /usr/local/go/src/net/dial.go:399
		_go_fuzz_dep_.CoverTab[13118]++
//line /usr/local/go/src/net/dial.go:399
		// _ = "end of CoverTab[13118]"
//line /usr/local/go/src/net/dial.go:399
	}
//line /usr/local/go/src/net/dial.go:399
	// _ = "end of CoverTab[13106]"
//line /usr/local/go/src/net/dial.go:399
	_go_fuzz_dep_.CoverTab[13107]++
						if oldCancel := d.Cancel; oldCancel != nil {
//line /usr/local/go/src/net/dial.go:400
		_go_fuzz_dep_.CoverTab[13119]++
							subCtx, cancel := context.WithCancel(ctx)
							defer cancel()
//line /usr/local/go/src/net/dial.go:402
		_curRoutineNum6_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/dial.go:402
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum6_)
							go func() {
//line /usr/local/go/src/net/dial.go:403
			_go_fuzz_dep_.CoverTab[13121]++
//line /usr/local/go/src/net/dial.go:403
			defer func() {
//line /usr/local/go/src/net/dial.go:403
				_go_fuzz_dep_.CoverTab[13122]++
//line /usr/local/go/src/net/dial.go:403
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum6_)
//line /usr/local/go/src/net/dial.go:403
				// _ = "end of CoverTab[13122]"
//line /usr/local/go/src/net/dial.go:403
			}()
								select {
			case <-oldCancel:
//line /usr/local/go/src/net/dial.go:405
				_go_fuzz_dep_.CoverTab[13123]++
									cancel()
//line /usr/local/go/src/net/dial.go:406
				// _ = "end of CoverTab[13123]"
			case <-subCtx.Done():
//line /usr/local/go/src/net/dial.go:407
				_go_fuzz_dep_.CoverTab[13124]++
//line /usr/local/go/src/net/dial.go:407
				// _ = "end of CoverTab[13124]"
			}
//line /usr/local/go/src/net/dial.go:408
			// _ = "end of CoverTab[13121]"
		}()
//line /usr/local/go/src/net/dial.go:409
		// _ = "end of CoverTab[13119]"
//line /usr/local/go/src/net/dial.go:409
		_go_fuzz_dep_.CoverTab[13120]++
							ctx = subCtx
//line /usr/local/go/src/net/dial.go:410
		// _ = "end of CoverTab[13120]"
	} else {
//line /usr/local/go/src/net/dial.go:411
		_go_fuzz_dep_.CoverTab[13125]++
//line /usr/local/go/src/net/dial.go:411
		// _ = "end of CoverTab[13125]"
//line /usr/local/go/src/net/dial.go:411
	}
//line /usr/local/go/src/net/dial.go:411
	// _ = "end of CoverTab[13107]"
//line /usr/local/go/src/net/dial.go:411
	_go_fuzz_dep_.CoverTab[13108]++

//line /usr/local/go/src/net/dial.go:414
	resolveCtx := ctx
	if trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace); trace != nil {
//line /usr/local/go/src/net/dial.go:415
		_go_fuzz_dep_.CoverTab[13126]++
							shadow := *trace
							shadow.ConnectStart = nil
							shadow.ConnectDone = nil
							resolveCtx = context.WithValue(resolveCtx, nettrace.TraceKey{}, &shadow)
//line /usr/local/go/src/net/dial.go:419
		// _ = "end of CoverTab[13126]"
	} else {
//line /usr/local/go/src/net/dial.go:420
		_go_fuzz_dep_.CoverTab[13127]++
//line /usr/local/go/src/net/dial.go:420
		// _ = "end of CoverTab[13127]"
//line /usr/local/go/src/net/dial.go:420
	}
//line /usr/local/go/src/net/dial.go:420
	// _ = "end of CoverTab[13108]"
//line /usr/local/go/src/net/dial.go:420
	_go_fuzz_dep_.CoverTab[13109]++

						addrs, err := d.resolver().resolveAddrList(resolveCtx, "dial", network, address, d.LocalAddr)
						if err != nil {
//line /usr/local/go/src/net/dial.go:423
		_go_fuzz_dep_.CoverTab[13128]++
							return nil, &OpError{Op: "dial", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:424
		// _ = "end of CoverTab[13128]"
	} else {
//line /usr/local/go/src/net/dial.go:425
		_go_fuzz_dep_.CoverTab[13129]++
//line /usr/local/go/src/net/dial.go:425
		// _ = "end of CoverTab[13129]"
//line /usr/local/go/src/net/dial.go:425
	}
//line /usr/local/go/src/net/dial.go:425
	// _ = "end of CoverTab[13109]"
//line /usr/local/go/src/net/dial.go:425
	_go_fuzz_dep_.CoverTab[13110]++

						sd := &sysDialer{
		Dialer:		*d,
		network:	network,
		address:	address,
	}

	var primaries, fallbacks addrList
	if d.dualStack() && func() bool {
//line /usr/local/go/src/net/dial.go:434
		_go_fuzz_dep_.CoverTab[13130]++
//line /usr/local/go/src/net/dial.go:434
		return network == "tcp"
//line /usr/local/go/src/net/dial.go:434
		// _ = "end of CoverTab[13130]"
//line /usr/local/go/src/net/dial.go:434
	}() {
//line /usr/local/go/src/net/dial.go:434
		_go_fuzz_dep_.CoverTab[13131]++
							primaries, fallbacks = addrs.partition(isIPv4)
//line /usr/local/go/src/net/dial.go:435
		// _ = "end of CoverTab[13131]"
	} else {
//line /usr/local/go/src/net/dial.go:436
		_go_fuzz_dep_.CoverTab[13132]++
							primaries = addrs
//line /usr/local/go/src/net/dial.go:437
		// _ = "end of CoverTab[13132]"
	}
//line /usr/local/go/src/net/dial.go:438
	// _ = "end of CoverTab[13110]"
//line /usr/local/go/src/net/dial.go:438
	_go_fuzz_dep_.CoverTab[13111]++

						return sd.dialParallel(ctx, primaries, fallbacks)
//line /usr/local/go/src/net/dial.go:440
	// _ = "end of CoverTab[13111]"
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
	_go_fuzz_dep_.CoverTab[13133]++
						if len(fallbacks) == 0 {
//line /usr/local/go/src/net/dial.go:448
		_go_fuzz_dep_.CoverTab[13136]++
							return sd.dialSerial(ctx, primaries)
//line /usr/local/go/src/net/dial.go:449
		// _ = "end of CoverTab[13136]"
	} else {
//line /usr/local/go/src/net/dial.go:450
		_go_fuzz_dep_.CoverTab[13137]++
//line /usr/local/go/src/net/dial.go:450
		// _ = "end of CoverTab[13137]"
//line /usr/local/go/src/net/dial.go:450
	}
//line /usr/local/go/src/net/dial.go:450
	// _ = "end of CoverTab[13133]"
//line /usr/local/go/src/net/dial.go:450
	_go_fuzz_dep_.CoverTab[13134]++

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
		_go_fuzz_dep_.CoverTab[13138]++
							ras := primaries
							if !primary {
//line /usr/local/go/src/net/dial.go:465
			_go_fuzz_dep_.CoverTab[13140]++
								ras = fallbacks
//line /usr/local/go/src/net/dial.go:466
			// _ = "end of CoverTab[13140]"
		} else {
//line /usr/local/go/src/net/dial.go:467
			_go_fuzz_dep_.CoverTab[13141]++
//line /usr/local/go/src/net/dial.go:467
			// _ = "end of CoverTab[13141]"
//line /usr/local/go/src/net/dial.go:467
		}
//line /usr/local/go/src/net/dial.go:467
		// _ = "end of CoverTab[13138]"
//line /usr/local/go/src/net/dial.go:467
		_go_fuzz_dep_.CoverTab[13139]++
							c, err := sd.dialSerial(ctx, ras)
							select {
		case results <- dialResult{Conn: c, error: err, primary: primary, done: true}:
//line /usr/local/go/src/net/dial.go:470
			_go_fuzz_dep_.CoverTab[13142]++
//line /usr/local/go/src/net/dial.go:470
			// _ = "end of CoverTab[13142]"
		case <-returned:
//line /usr/local/go/src/net/dial.go:471
			_go_fuzz_dep_.CoverTab[13143]++
								if c != nil {
//line /usr/local/go/src/net/dial.go:472
				_go_fuzz_dep_.CoverTab[13144]++
									c.Close()
//line /usr/local/go/src/net/dial.go:473
				// _ = "end of CoverTab[13144]"
			} else {
//line /usr/local/go/src/net/dial.go:474
				_go_fuzz_dep_.CoverTab[13145]++
//line /usr/local/go/src/net/dial.go:474
				// _ = "end of CoverTab[13145]"
//line /usr/local/go/src/net/dial.go:474
			}
//line /usr/local/go/src/net/dial.go:474
			// _ = "end of CoverTab[13143]"
		}
//line /usr/local/go/src/net/dial.go:475
		// _ = "end of CoverTab[13139]"
	}
//line /usr/local/go/src/net/dial.go:476
	// _ = "end of CoverTab[13134]"
//line /usr/local/go/src/net/dial.go:476
	_go_fuzz_dep_.CoverTab[13135]++

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
		_go_fuzz_dep_.CoverTab[13146]++
//line /usr/local/go/src/net/dial.go:483
		defer func() {
//line /usr/local/go/src/net/dial.go:483
			_go_fuzz_dep_.CoverTab[13147]++
//line /usr/local/go/src/net/dial.go:483
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum7_)
//line /usr/local/go/src/net/dial.go:483
			// _ = "end of CoverTab[13147]"
//line /usr/local/go/src/net/dial.go:483
		}()
//line /usr/local/go/src/net/dial.go:483
		startRacer(primaryCtx, true)
//line /usr/local/go/src/net/dial.go:483
		// _ = "end of CoverTab[13146]"
//line /usr/local/go/src/net/dial.go:483
	}()

//line /usr/local/go/src/net/dial.go:486
	fallbackTimer := time.NewTimer(sd.fallbackDelay())
	defer fallbackTimer.Stop()

	for {
//line /usr/local/go/src/net/dial.go:489
		_go_fuzz_dep_.CoverTab[13148]++
							select {
		case <-fallbackTimer.C:
//line /usr/local/go/src/net/dial.go:491
			_go_fuzz_dep_.CoverTab[13149]++
								fallbackCtx, fallbackCancel := context.WithCancel(ctx)
								defer fallbackCancel()
								go startRacer(fallbackCtx, false)
//line /usr/local/go/src/net/dial.go:494
			// _ = "end of CoverTab[13149]"

		case res := <-results:
//line /usr/local/go/src/net/dial.go:496
			_go_fuzz_dep_.CoverTab[13150]++
								if res.error == nil {
//line /usr/local/go/src/net/dial.go:497
				_go_fuzz_dep_.CoverTab[13154]++
									return res.Conn, nil
//line /usr/local/go/src/net/dial.go:498
				// _ = "end of CoverTab[13154]"
			} else {
//line /usr/local/go/src/net/dial.go:499
				_go_fuzz_dep_.CoverTab[13155]++
//line /usr/local/go/src/net/dial.go:499
				// _ = "end of CoverTab[13155]"
//line /usr/local/go/src/net/dial.go:499
			}
//line /usr/local/go/src/net/dial.go:499
			// _ = "end of CoverTab[13150]"
//line /usr/local/go/src/net/dial.go:499
			_go_fuzz_dep_.CoverTab[13151]++
								if res.primary {
//line /usr/local/go/src/net/dial.go:500
				_go_fuzz_dep_.CoverTab[13156]++
									primary = res
//line /usr/local/go/src/net/dial.go:501
				// _ = "end of CoverTab[13156]"
			} else {
//line /usr/local/go/src/net/dial.go:502
				_go_fuzz_dep_.CoverTab[13157]++
									fallback = res
//line /usr/local/go/src/net/dial.go:503
				// _ = "end of CoverTab[13157]"
			}
//line /usr/local/go/src/net/dial.go:504
			// _ = "end of CoverTab[13151]"
//line /usr/local/go/src/net/dial.go:504
			_go_fuzz_dep_.CoverTab[13152]++
								if primary.done && func() bool {
//line /usr/local/go/src/net/dial.go:505
				_go_fuzz_dep_.CoverTab[13158]++
//line /usr/local/go/src/net/dial.go:505
				return fallback.done
//line /usr/local/go/src/net/dial.go:505
				// _ = "end of CoverTab[13158]"
//line /usr/local/go/src/net/dial.go:505
			}() {
//line /usr/local/go/src/net/dial.go:505
				_go_fuzz_dep_.CoverTab[13159]++
									return nil, primary.error
//line /usr/local/go/src/net/dial.go:506
				// _ = "end of CoverTab[13159]"
			} else {
//line /usr/local/go/src/net/dial.go:507
				_go_fuzz_dep_.CoverTab[13160]++
//line /usr/local/go/src/net/dial.go:507
				// _ = "end of CoverTab[13160]"
//line /usr/local/go/src/net/dial.go:507
			}
//line /usr/local/go/src/net/dial.go:507
			// _ = "end of CoverTab[13152]"
//line /usr/local/go/src/net/dial.go:507
			_go_fuzz_dep_.CoverTab[13153]++
								if res.primary && func() bool {
//line /usr/local/go/src/net/dial.go:508
				_go_fuzz_dep_.CoverTab[13161]++
//line /usr/local/go/src/net/dial.go:508
				return fallbackTimer.Stop()
//line /usr/local/go/src/net/dial.go:508
				// _ = "end of CoverTab[13161]"
//line /usr/local/go/src/net/dial.go:508
			}() {
//line /usr/local/go/src/net/dial.go:508
				_go_fuzz_dep_.CoverTab[13162]++

//line /usr/local/go/src/net/dial.go:513
				fallbackTimer.Reset(0)
//line /usr/local/go/src/net/dial.go:513
				// _ = "end of CoverTab[13162]"
			} else {
//line /usr/local/go/src/net/dial.go:514
				_go_fuzz_dep_.CoverTab[13163]++
//line /usr/local/go/src/net/dial.go:514
				// _ = "end of CoverTab[13163]"
//line /usr/local/go/src/net/dial.go:514
			}
//line /usr/local/go/src/net/dial.go:514
			// _ = "end of CoverTab[13153]"
		}
//line /usr/local/go/src/net/dial.go:515
		// _ = "end of CoverTab[13148]"
	}
//line /usr/local/go/src/net/dial.go:516
	// _ = "end of CoverTab[13135]"
}

// dialSerial connects to a list of addresses in sequence, returning
//line /usr/local/go/src/net/dial.go:519
// either the first successful connection, or the first error.
//line /usr/local/go/src/net/dial.go:521
func (sd *sysDialer) dialSerial(ctx context.Context, ras addrList) (Conn, error) {
//line /usr/local/go/src/net/dial.go:521
	_go_fuzz_dep_.CoverTab[13164]++
						var firstErr error	// The error from the first address is most relevant.

						for i, ra := range ras {
//line /usr/local/go/src/net/dial.go:524
		_go_fuzz_dep_.CoverTab[13167]++
							select {
		case <-ctx.Done():
//line /usr/local/go/src/net/dial.go:526
			_go_fuzz_dep_.CoverTab[13171]++
								return nil, &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: mapErr(ctx.Err())}
//line /usr/local/go/src/net/dial.go:527
			// _ = "end of CoverTab[13171]"
		default:
//line /usr/local/go/src/net/dial.go:528
			_go_fuzz_dep_.CoverTab[13172]++
//line /usr/local/go/src/net/dial.go:528
			// _ = "end of CoverTab[13172]"
		}
//line /usr/local/go/src/net/dial.go:529
		// _ = "end of CoverTab[13167]"
//line /usr/local/go/src/net/dial.go:529
		_go_fuzz_dep_.CoverTab[13168]++

							dialCtx := ctx
							if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /usr/local/go/src/net/dial.go:532
			_go_fuzz_dep_.CoverTab[13173]++
								partialDeadline, err := partialDeadline(time.Now(), deadline, len(ras)-i)
								if err != nil {
//line /usr/local/go/src/net/dial.go:534
				_go_fuzz_dep_.CoverTab[13175]++

									if firstErr == nil {
//line /usr/local/go/src/net/dial.go:536
					_go_fuzz_dep_.CoverTab[13177]++
										firstErr = &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: err}
//line /usr/local/go/src/net/dial.go:537
					// _ = "end of CoverTab[13177]"
				} else {
//line /usr/local/go/src/net/dial.go:538
					_go_fuzz_dep_.CoverTab[13178]++
//line /usr/local/go/src/net/dial.go:538
					// _ = "end of CoverTab[13178]"
//line /usr/local/go/src/net/dial.go:538
				}
//line /usr/local/go/src/net/dial.go:538
				// _ = "end of CoverTab[13175]"
//line /usr/local/go/src/net/dial.go:538
				_go_fuzz_dep_.CoverTab[13176]++
									break
//line /usr/local/go/src/net/dial.go:539
				// _ = "end of CoverTab[13176]"
			} else {
//line /usr/local/go/src/net/dial.go:540
				_go_fuzz_dep_.CoverTab[13179]++
//line /usr/local/go/src/net/dial.go:540
				// _ = "end of CoverTab[13179]"
//line /usr/local/go/src/net/dial.go:540
			}
//line /usr/local/go/src/net/dial.go:540
			// _ = "end of CoverTab[13173]"
//line /usr/local/go/src/net/dial.go:540
			_go_fuzz_dep_.CoverTab[13174]++
								if partialDeadline.Before(deadline) {
//line /usr/local/go/src/net/dial.go:541
				_go_fuzz_dep_.CoverTab[13180]++
									var cancel context.CancelFunc
									dialCtx, cancel = context.WithDeadline(ctx, partialDeadline)
									defer cancel()
//line /usr/local/go/src/net/dial.go:544
				// _ = "end of CoverTab[13180]"
			} else {
//line /usr/local/go/src/net/dial.go:545
				_go_fuzz_dep_.CoverTab[13181]++
//line /usr/local/go/src/net/dial.go:545
				// _ = "end of CoverTab[13181]"
//line /usr/local/go/src/net/dial.go:545
			}
//line /usr/local/go/src/net/dial.go:545
			// _ = "end of CoverTab[13174]"
		} else {
//line /usr/local/go/src/net/dial.go:546
			_go_fuzz_dep_.CoverTab[13182]++
//line /usr/local/go/src/net/dial.go:546
			// _ = "end of CoverTab[13182]"
//line /usr/local/go/src/net/dial.go:546
		}
//line /usr/local/go/src/net/dial.go:546
		// _ = "end of CoverTab[13168]"
//line /usr/local/go/src/net/dial.go:546
		_go_fuzz_dep_.CoverTab[13169]++

							c, err := sd.dialSingle(dialCtx, ra)
							if err == nil {
//line /usr/local/go/src/net/dial.go:549
			_go_fuzz_dep_.CoverTab[13183]++
								return c, nil
//line /usr/local/go/src/net/dial.go:550
			// _ = "end of CoverTab[13183]"
		} else {
//line /usr/local/go/src/net/dial.go:551
			_go_fuzz_dep_.CoverTab[13184]++
//line /usr/local/go/src/net/dial.go:551
			// _ = "end of CoverTab[13184]"
//line /usr/local/go/src/net/dial.go:551
		}
//line /usr/local/go/src/net/dial.go:551
		// _ = "end of CoverTab[13169]"
//line /usr/local/go/src/net/dial.go:551
		_go_fuzz_dep_.CoverTab[13170]++
							if firstErr == nil {
//line /usr/local/go/src/net/dial.go:552
			_go_fuzz_dep_.CoverTab[13185]++
								firstErr = err
//line /usr/local/go/src/net/dial.go:553
			// _ = "end of CoverTab[13185]"
		} else {
//line /usr/local/go/src/net/dial.go:554
			_go_fuzz_dep_.CoverTab[13186]++
//line /usr/local/go/src/net/dial.go:554
			// _ = "end of CoverTab[13186]"
//line /usr/local/go/src/net/dial.go:554
		}
//line /usr/local/go/src/net/dial.go:554
		// _ = "end of CoverTab[13170]"
	}
//line /usr/local/go/src/net/dial.go:555
	// _ = "end of CoverTab[13164]"
//line /usr/local/go/src/net/dial.go:555
	_go_fuzz_dep_.CoverTab[13165]++

						if firstErr == nil {
//line /usr/local/go/src/net/dial.go:557
		_go_fuzz_dep_.CoverTab[13187]++
							firstErr = &OpError{Op: "dial", Net: sd.network, Source: nil, Addr: nil, Err: errMissingAddress}
//line /usr/local/go/src/net/dial.go:558
		// _ = "end of CoverTab[13187]"
	} else {
//line /usr/local/go/src/net/dial.go:559
		_go_fuzz_dep_.CoverTab[13188]++
//line /usr/local/go/src/net/dial.go:559
		// _ = "end of CoverTab[13188]"
//line /usr/local/go/src/net/dial.go:559
	}
//line /usr/local/go/src/net/dial.go:559
	// _ = "end of CoverTab[13165]"
//line /usr/local/go/src/net/dial.go:559
	_go_fuzz_dep_.CoverTab[13166]++
						return nil, firstErr
//line /usr/local/go/src/net/dial.go:560
	// _ = "end of CoverTab[13166]"
}

// dialSingle attempts to establish and returns a single connection to
//line /usr/local/go/src/net/dial.go:563
// the destination address.
//line /usr/local/go/src/net/dial.go:565
func (sd *sysDialer) dialSingle(ctx context.Context, ra Addr) (c Conn, err error) {
//line /usr/local/go/src/net/dial.go:565
	_go_fuzz_dep_.CoverTab[13189]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil {
//line /usr/local/go/src/net/dial.go:567
		_go_fuzz_dep_.CoverTab[13193]++
							raStr := ra.String()
							if trace.ConnectStart != nil {
//line /usr/local/go/src/net/dial.go:569
			_go_fuzz_dep_.CoverTab[13195]++
								trace.ConnectStart(sd.network, raStr)
//line /usr/local/go/src/net/dial.go:570
			// _ = "end of CoverTab[13195]"
		} else {
//line /usr/local/go/src/net/dial.go:571
			_go_fuzz_dep_.CoverTab[13196]++
//line /usr/local/go/src/net/dial.go:571
			// _ = "end of CoverTab[13196]"
//line /usr/local/go/src/net/dial.go:571
		}
//line /usr/local/go/src/net/dial.go:571
		// _ = "end of CoverTab[13193]"
//line /usr/local/go/src/net/dial.go:571
		_go_fuzz_dep_.CoverTab[13194]++
							if trace.ConnectDone != nil {
//line /usr/local/go/src/net/dial.go:572
			_go_fuzz_dep_.CoverTab[13197]++
								defer func() {
//line /usr/local/go/src/net/dial.go:573
				_go_fuzz_dep_.CoverTab[13198]++
//line /usr/local/go/src/net/dial.go:573
				trace.ConnectDone(sd.network, raStr, err)
//line /usr/local/go/src/net/dial.go:573
				// _ = "end of CoverTab[13198]"
//line /usr/local/go/src/net/dial.go:573
			}()
//line /usr/local/go/src/net/dial.go:573
			// _ = "end of CoverTab[13197]"
		} else {
//line /usr/local/go/src/net/dial.go:574
			_go_fuzz_dep_.CoverTab[13199]++
//line /usr/local/go/src/net/dial.go:574
			// _ = "end of CoverTab[13199]"
//line /usr/local/go/src/net/dial.go:574
		}
//line /usr/local/go/src/net/dial.go:574
		// _ = "end of CoverTab[13194]"
	} else {
//line /usr/local/go/src/net/dial.go:575
		_go_fuzz_dep_.CoverTab[13200]++
//line /usr/local/go/src/net/dial.go:575
		// _ = "end of CoverTab[13200]"
//line /usr/local/go/src/net/dial.go:575
	}
//line /usr/local/go/src/net/dial.go:575
	// _ = "end of CoverTab[13189]"
//line /usr/local/go/src/net/dial.go:575
	_go_fuzz_dep_.CoverTab[13190]++
						la := sd.LocalAddr
						switch ra := ra.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/dial.go:578
		_go_fuzz_dep_.CoverTab[13201]++
							la, _ := la.(*TCPAddr)
							c, err = sd.dialTCP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:580
		// _ = "end of CoverTab[13201]"
	case *UDPAddr:
//line /usr/local/go/src/net/dial.go:581
		_go_fuzz_dep_.CoverTab[13202]++
							la, _ := la.(*UDPAddr)
							c, err = sd.dialUDP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:583
		// _ = "end of CoverTab[13202]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:584
		_go_fuzz_dep_.CoverTab[13203]++
							la, _ := la.(*IPAddr)
							c, err = sd.dialIP(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:586
		// _ = "end of CoverTab[13203]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:587
		_go_fuzz_dep_.CoverTab[13204]++
							la, _ := la.(*UnixAddr)
							c, err = sd.dialUnix(ctx, la, ra)
//line /usr/local/go/src/net/dial.go:589
		// _ = "end of CoverTab[13204]"
	default:
//line /usr/local/go/src/net/dial.go:590
		_go_fuzz_dep_.CoverTab[13205]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: &AddrError{Err: "unexpected address type", Addr: sd.address}}
//line /usr/local/go/src/net/dial.go:591
		// _ = "end of CoverTab[13205]"
	}
//line /usr/local/go/src/net/dial.go:592
	// _ = "end of CoverTab[13190]"
//line /usr/local/go/src/net/dial.go:592
	_go_fuzz_dep_.CoverTab[13191]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:593
		_go_fuzz_dep_.CoverTab[13206]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: err}
//line /usr/local/go/src/net/dial.go:594
		// _ = "end of CoverTab[13206]"
	} else {
//line /usr/local/go/src/net/dial.go:595
		_go_fuzz_dep_.CoverTab[13207]++
//line /usr/local/go/src/net/dial.go:595
		// _ = "end of CoverTab[13207]"
//line /usr/local/go/src/net/dial.go:595
	}
//line /usr/local/go/src/net/dial.go:595
	// _ = "end of CoverTab[13191]"
//line /usr/local/go/src/net/dial.go:595
	_go_fuzz_dep_.CoverTab[13192]++
						return c, nil
//line /usr/local/go/src/net/dial.go:596
	// _ = "end of CoverTab[13192]"
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
	_go_fuzz_dep_.CoverTab[13208]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /usr/local/go/src/net/dial.go:624
		_go_fuzz_dep_.CoverTab[13212]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:625
		// _ = "end of CoverTab[13212]"
	} else {
//line /usr/local/go/src/net/dial.go:626
		_go_fuzz_dep_.CoverTab[13213]++
//line /usr/local/go/src/net/dial.go:626
		// _ = "end of CoverTab[13213]"
//line /usr/local/go/src/net/dial.go:626
	}
//line /usr/local/go/src/net/dial.go:626
	// _ = "end of CoverTab[13208]"
//line /usr/local/go/src/net/dial.go:626
	_go_fuzz_dep_.CoverTab[13209]++
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
		_go_fuzz_dep_.CoverTab[13214]++
							l, err = sl.listenTCP(ctx, la)
//line /usr/local/go/src/net/dial.go:636
		// _ = "end of CoverTab[13214]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:637
		_go_fuzz_dep_.CoverTab[13215]++
							l, err = sl.listenUnix(ctx, la)
//line /usr/local/go/src/net/dial.go:638
		// _ = "end of CoverTab[13215]"
	default:
//line /usr/local/go/src/net/dial.go:639
		_go_fuzz_dep_.CoverTab[13216]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /usr/local/go/src/net/dial.go:640
		// _ = "end of CoverTab[13216]"
	}
//line /usr/local/go/src/net/dial.go:641
	// _ = "end of CoverTab[13209]"
//line /usr/local/go/src/net/dial.go:641
	_go_fuzz_dep_.CoverTab[13210]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:642
		_go_fuzz_dep_.CoverTab[13217]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /usr/local/go/src/net/dial.go:643
		// _ = "end of CoverTab[13217]"
	} else {
//line /usr/local/go/src/net/dial.go:644
		_go_fuzz_dep_.CoverTab[13218]++
//line /usr/local/go/src/net/dial.go:644
		// _ = "end of CoverTab[13218]"
//line /usr/local/go/src/net/dial.go:644
	}
//line /usr/local/go/src/net/dial.go:644
	// _ = "end of CoverTab[13210]"
//line /usr/local/go/src/net/dial.go:644
	_go_fuzz_dep_.CoverTab[13211]++
						return l, nil
//line /usr/local/go/src/net/dial.go:645
	// _ = "end of CoverTab[13211]"
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
	_go_fuzz_dep_.CoverTab[13219]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /usr/local/go/src/net/dial.go:654
		_go_fuzz_dep_.CoverTab[13223]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/dial.go:655
		// _ = "end of CoverTab[13223]"
	} else {
//line /usr/local/go/src/net/dial.go:656
		_go_fuzz_dep_.CoverTab[13224]++
//line /usr/local/go/src/net/dial.go:656
		// _ = "end of CoverTab[13224]"
//line /usr/local/go/src/net/dial.go:656
	}
//line /usr/local/go/src/net/dial.go:656
	// _ = "end of CoverTab[13219]"
//line /usr/local/go/src/net/dial.go:656
	_go_fuzz_dep_.CoverTab[13220]++
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
		_go_fuzz_dep_.CoverTab[13225]++
							c, err = sl.listenUDP(ctx, la)
//line /usr/local/go/src/net/dial.go:666
		// _ = "end of CoverTab[13225]"
	case *IPAddr:
//line /usr/local/go/src/net/dial.go:667
		_go_fuzz_dep_.CoverTab[13226]++
							c, err = sl.listenIP(ctx, la)
//line /usr/local/go/src/net/dial.go:668
		// _ = "end of CoverTab[13226]"
	case *UnixAddr:
//line /usr/local/go/src/net/dial.go:669
		_go_fuzz_dep_.CoverTab[13227]++
							c, err = sl.listenUnixgram(ctx, la)
//line /usr/local/go/src/net/dial.go:670
		// _ = "end of CoverTab[13227]"
	default:
//line /usr/local/go/src/net/dial.go:671
		_go_fuzz_dep_.CoverTab[13228]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /usr/local/go/src/net/dial.go:672
		// _ = "end of CoverTab[13228]"
	}
//line /usr/local/go/src/net/dial.go:673
	// _ = "end of CoverTab[13220]"
//line /usr/local/go/src/net/dial.go:673
	_go_fuzz_dep_.CoverTab[13221]++
						if err != nil {
//line /usr/local/go/src/net/dial.go:674
		_go_fuzz_dep_.CoverTab[13229]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /usr/local/go/src/net/dial.go:675
		// _ = "end of CoverTab[13229]"
	} else {
//line /usr/local/go/src/net/dial.go:676
		_go_fuzz_dep_.CoverTab[13230]++
//line /usr/local/go/src/net/dial.go:676
		// _ = "end of CoverTab[13230]"
//line /usr/local/go/src/net/dial.go:676
	}
//line /usr/local/go/src/net/dial.go:676
	// _ = "end of CoverTab[13221]"
//line /usr/local/go/src/net/dial.go:676
	_go_fuzz_dep_.CoverTab[13222]++
						return c, nil
//line /usr/local/go/src/net/dial.go:677
	// _ = "end of CoverTab[13222]"
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
	_go_fuzz_dep_.CoverTab[13231]++
						var lc ListenConfig
						return lc.Listen(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:709
	// _ = "end of CoverTab[13231]"
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
	_go_fuzz_dep_.CoverTab[13232]++
						var lc ListenConfig
						return lc.ListenPacket(context.Background(), network, address)
//line /usr/local/go/src/net/dial.go:739
	// _ = "end of CoverTab[13232]"
}

//line /usr/local/go/src/net/dial.go:740
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dial.go:740
var _ = _go_fuzz_dep_.CoverTab
