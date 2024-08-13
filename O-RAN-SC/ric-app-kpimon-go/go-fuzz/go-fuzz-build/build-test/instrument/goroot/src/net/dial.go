// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/dial.go:5
package net

//line /snap/go/10455/src/net/dial.go:5
import (
//line /snap/go/10455/src/net/dial.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/dial.go:5
)
//line /snap/go/10455/src/net/dial.go:5
import (
//line /snap/go/10455/src/net/dial.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/dial.go:5
)

import (
	"context"
	"internal/godebug"
	"internal/nettrace"
	"syscall"
	"time"
)

const (
	// defaultTCPKeepAlive is a default constant value for TCPKeepAlive times
	// See go.dev/issue/31510
	defaultTCPKeepAlive	= 15 * time.Second

	// For the moment, MultiPath TCP is not used by default
	// See go.dev/issue/56539
	defaultMPTCPEnabled	= false
)

var multipathtcp = godebug.New("multipathtcp")

// mptcpStatus is a tristate for Multipath TCP, see go.dev/issue/56539
type mptcpStatus uint8

const (
	// The value 0 is the system default, linked to defaultMPTCPEnabled
	mptcpUseDefault	mptcpStatus	= iota
	mptcpEnabled
	mptcpDisabled
)

func (m *mptcpStatus) get() bool {
//line /snap/go/10455/src/net/dial.go:37
	_go_fuzz_dep_.CoverTab[4947]++
						switch *m {
	case mptcpEnabled:
//line /snap/go/10455/src/net/dial.go:39
		_go_fuzz_dep_.CoverTab[527675]++
//line /snap/go/10455/src/net/dial.go:39
		_go_fuzz_dep_.CoverTab[4950]++
							return true
//line /snap/go/10455/src/net/dial.go:40
		// _ = "end of CoverTab[4950]"
	case mptcpDisabled:
//line /snap/go/10455/src/net/dial.go:41
		_go_fuzz_dep_.CoverTab[527676]++
//line /snap/go/10455/src/net/dial.go:41
		_go_fuzz_dep_.CoverTab[4951]++
							return false
//line /snap/go/10455/src/net/dial.go:42
		// _ = "end of CoverTab[4951]"
//line /snap/go/10455/src/net/dial.go:42
	default:
//line /snap/go/10455/src/net/dial.go:42
		_go_fuzz_dep_.CoverTab[527677]++
//line /snap/go/10455/src/net/dial.go:42
		_go_fuzz_dep_.CoverTab[4952]++
//line /snap/go/10455/src/net/dial.go:42
		// _ = "end of CoverTab[4952]"
	}
//line /snap/go/10455/src/net/dial.go:43
	// _ = "end of CoverTab[4947]"
//line /snap/go/10455/src/net/dial.go:43
	_go_fuzz_dep_.CoverTab[4948]++

//line /snap/go/10455/src/net/dial.go:46
	if multipathtcp.Value() == "1" {
//line /snap/go/10455/src/net/dial.go:46
		_go_fuzz_dep_.CoverTab[527678]++
//line /snap/go/10455/src/net/dial.go:46
		_go_fuzz_dep_.CoverTab[4953]++
							multipathtcp.IncNonDefault()

							return true
//line /snap/go/10455/src/net/dial.go:49
		// _ = "end of CoverTab[4953]"
	} else {
//line /snap/go/10455/src/net/dial.go:50
		_go_fuzz_dep_.CoverTab[527679]++
//line /snap/go/10455/src/net/dial.go:50
		_go_fuzz_dep_.CoverTab[4954]++
//line /snap/go/10455/src/net/dial.go:50
		// _ = "end of CoverTab[4954]"
//line /snap/go/10455/src/net/dial.go:50
	}
//line /snap/go/10455/src/net/dial.go:50
	// _ = "end of CoverTab[4948]"
//line /snap/go/10455/src/net/dial.go:50
	_go_fuzz_dep_.CoverTab[4949]++

						return defaultMPTCPEnabled
//line /snap/go/10455/src/net/dial.go:52
	// _ = "end of CoverTab[4949]"
}

func (m *mptcpStatus) set(use bool) {
//line /snap/go/10455/src/net/dial.go:55
	_go_fuzz_dep_.CoverTab[4955]++
						if use {
//line /snap/go/10455/src/net/dial.go:56
		_go_fuzz_dep_.CoverTab[527680]++
//line /snap/go/10455/src/net/dial.go:56
		_go_fuzz_dep_.CoverTab[4956]++
							*m = mptcpEnabled
//line /snap/go/10455/src/net/dial.go:57
		// _ = "end of CoverTab[4956]"
	} else {
//line /snap/go/10455/src/net/dial.go:58
		_go_fuzz_dep_.CoverTab[527681]++
//line /snap/go/10455/src/net/dial.go:58
		_go_fuzz_dep_.CoverTab[4957]++
							*m = mptcpDisabled
//line /snap/go/10455/src/net/dial.go:59
		// _ = "end of CoverTab[4957]"
	}
//line /snap/go/10455/src/net/dial.go:60
	// _ = "end of CoverTab[4955]"
}

// A Dialer contains options for connecting to an address.
//line /snap/go/10455/src/net/dial.go:63
//
//line /snap/go/10455/src/net/dial.go:63
// The zero value for each field is equivalent to dialing
//line /snap/go/10455/src/net/dial.go:63
// without that option. Dialing with the zero value of Dialer
//line /snap/go/10455/src/net/dial.go:63
// is therefore equivalent to just calling the Dial function.
//line /snap/go/10455/src/net/dial.go:63
//
//line /snap/go/10455/src/net/dial.go:63
// It is safe to call Dialer's methods concurrently.
//line /snap/go/10455/src/net/dial.go:70
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
	// Network and address parameters passed to Control function are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the Control function to be called with "tcp4" or "tcp6".
	//
	// Control is ignored if ControlContext is not nil.
	Control	func(network, address string, c syscall.RawConn) error

	// If ControlContext is not nil, it is called after creating the network
	// connection but before actually dialing.
	//
	// Network and address parameters passed to ControlContext function are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the ControlContext function to be called with "tcp4" or "tcp6".
	//
	// If ControlContext is not nil, Control is ignored.
	ControlContext	func(ctx context.Context, network, address string, c syscall.RawConn) error

	// If mptcpStatus is set to a value allowing Multipath TCP (MPTCP) to be
	// used, any call to Dial with "tcp(4|6)" as network will use MPTCP if
	// supported by the operating system.
	mptcpStatus	mptcpStatus
}

func (d *Dialer) dualStack() bool {
//line /snap/go/10455/src/net/dial.go:161
	_go_fuzz_dep_.CoverTab[4958]++
//line /snap/go/10455/src/net/dial.go:161
	return d.FallbackDelay >= 0
//line /snap/go/10455/src/net/dial.go:161
	// _ = "end of CoverTab[4958]"
//line /snap/go/10455/src/net/dial.go:161
}

func minNonzeroTime(a, b time.Time) time.Time {
//line /snap/go/10455/src/net/dial.go:163
	_go_fuzz_dep_.CoverTab[4959]++
						if a.IsZero() {
//line /snap/go/10455/src/net/dial.go:164
		_go_fuzz_dep_.CoverTab[527682]++
//line /snap/go/10455/src/net/dial.go:164
		_go_fuzz_dep_.CoverTab[4962]++
							return b
//line /snap/go/10455/src/net/dial.go:165
		// _ = "end of CoverTab[4962]"
	} else {
//line /snap/go/10455/src/net/dial.go:166
		_go_fuzz_dep_.CoverTab[527683]++
//line /snap/go/10455/src/net/dial.go:166
		_go_fuzz_dep_.CoverTab[4963]++
//line /snap/go/10455/src/net/dial.go:166
		// _ = "end of CoverTab[4963]"
//line /snap/go/10455/src/net/dial.go:166
	}
//line /snap/go/10455/src/net/dial.go:166
	// _ = "end of CoverTab[4959]"
//line /snap/go/10455/src/net/dial.go:166
	_go_fuzz_dep_.CoverTab[4960]++
						if b.IsZero() || func() bool {
//line /snap/go/10455/src/net/dial.go:167
		_go_fuzz_dep_.CoverTab[4964]++
//line /snap/go/10455/src/net/dial.go:167
		return a.Before(b)
//line /snap/go/10455/src/net/dial.go:167
		// _ = "end of CoverTab[4964]"
//line /snap/go/10455/src/net/dial.go:167
	}() {
//line /snap/go/10455/src/net/dial.go:167
		_go_fuzz_dep_.CoverTab[527684]++
//line /snap/go/10455/src/net/dial.go:167
		_go_fuzz_dep_.CoverTab[4965]++
							return a
//line /snap/go/10455/src/net/dial.go:168
		// _ = "end of CoverTab[4965]"
	} else {
//line /snap/go/10455/src/net/dial.go:169
		_go_fuzz_dep_.CoverTab[527685]++
//line /snap/go/10455/src/net/dial.go:169
		_go_fuzz_dep_.CoverTab[4966]++
//line /snap/go/10455/src/net/dial.go:169
		// _ = "end of CoverTab[4966]"
//line /snap/go/10455/src/net/dial.go:169
	}
//line /snap/go/10455/src/net/dial.go:169
	// _ = "end of CoverTab[4960]"
//line /snap/go/10455/src/net/dial.go:169
	_go_fuzz_dep_.CoverTab[4961]++
						return b
//line /snap/go/10455/src/net/dial.go:170
	// _ = "end of CoverTab[4961]"
}

// deadline returns the earliest of:
//line /snap/go/10455/src/net/dial.go:173
//   - now+Timeout
//line /snap/go/10455/src/net/dial.go:173
//   - d.Deadline
//line /snap/go/10455/src/net/dial.go:173
//   - the context's deadline
//line /snap/go/10455/src/net/dial.go:173
//
//line /snap/go/10455/src/net/dial.go:173
// Or zero, if none of Timeout, Deadline, or context's deadline is set.
//line /snap/go/10455/src/net/dial.go:179
func (d *Dialer) deadline(ctx context.Context, now time.Time) (earliest time.Time) {
//line /snap/go/10455/src/net/dial.go:179
	_go_fuzz_dep_.CoverTab[4967]++
						if d.Timeout != 0 {
//line /snap/go/10455/src/net/dial.go:180
		_go_fuzz_dep_.CoverTab[527686]++
//line /snap/go/10455/src/net/dial.go:180
		_go_fuzz_dep_.CoverTab[4970]++
							earliest = now.Add(d.Timeout)
//line /snap/go/10455/src/net/dial.go:181
		// _ = "end of CoverTab[4970]"
	} else {
//line /snap/go/10455/src/net/dial.go:182
		_go_fuzz_dep_.CoverTab[527687]++
//line /snap/go/10455/src/net/dial.go:182
		_go_fuzz_dep_.CoverTab[4971]++
//line /snap/go/10455/src/net/dial.go:182
		// _ = "end of CoverTab[4971]"
//line /snap/go/10455/src/net/dial.go:182
	}
//line /snap/go/10455/src/net/dial.go:182
	// _ = "end of CoverTab[4967]"
//line /snap/go/10455/src/net/dial.go:182
	_go_fuzz_dep_.CoverTab[4968]++
						if d, ok := ctx.Deadline(); ok {
//line /snap/go/10455/src/net/dial.go:183
		_go_fuzz_dep_.CoverTab[527688]++
//line /snap/go/10455/src/net/dial.go:183
		_go_fuzz_dep_.CoverTab[4972]++
							earliest = minNonzeroTime(earliest, d)
//line /snap/go/10455/src/net/dial.go:184
		// _ = "end of CoverTab[4972]"
	} else {
//line /snap/go/10455/src/net/dial.go:185
		_go_fuzz_dep_.CoverTab[527689]++
//line /snap/go/10455/src/net/dial.go:185
		_go_fuzz_dep_.CoverTab[4973]++
//line /snap/go/10455/src/net/dial.go:185
		// _ = "end of CoverTab[4973]"
//line /snap/go/10455/src/net/dial.go:185
	}
//line /snap/go/10455/src/net/dial.go:185
	// _ = "end of CoverTab[4968]"
//line /snap/go/10455/src/net/dial.go:185
	_go_fuzz_dep_.CoverTab[4969]++
						return minNonzeroTime(earliest, d.Deadline)
//line /snap/go/10455/src/net/dial.go:186
	// _ = "end of CoverTab[4969]"
}

func (d *Dialer) resolver() *Resolver {
//line /snap/go/10455/src/net/dial.go:189
	_go_fuzz_dep_.CoverTab[4974]++
						if d.Resolver != nil {
//line /snap/go/10455/src/net/dial.go:190
		_go_fuzz_dep_.CoverTab[527690]++
//line /snap/go/10455/src/net/dial.go:190
		_go_fuzz_dep_.CoverTab[4976]++
							return d.Resolver
//line /snap/go/10455/src/net/dial.go:191
		// _ = "end of CoverTab[4976]"
	} else {
//line /snap/go/10455/src/net/dial.go:192
		_go_fuzz_dep_.CoverTab[527691]++
//line /snap/go/10455/src/net/dial.go:192
		_go_fuzz_dep_.CoverTab[4977]++
//line /snap/go/10455/src/net/dial.go:192
		// _ = "end of CoverTab[4977]"
//line /snap/go/10455/src/net/dial.go:192
	}
//line /snap/go/10455/src/net/dial.go:192
	// _ = "end of CoverTab[4974]"
//line /snap/go/10455/src/net/dial.go:192
	_go_fuzz_dep_.CoverTab[4975]++
						return DefaultResolver
//line /snap/go/10455/src/net/dial.go:193
	// _ = "end of CoverTab[4975]"
}

// partialDeadline returns the deadline to use for a single address,
//line /snap/go/10455/src/net/dial.go:196
// when multiple addresses are pending.
//line /snap/go/10455/src/net/dial.go:198
func partialDeadline(now, deadline time.Time, addrsRemaining int) (time.Time, error) {
//line /snap/go/10455/src/net/dial.go:198
	_go_fuzz_dep_.CoverTab[4978]++
						if deadline.IsZero() {
//line /snap/go/10455/src/net/dial.go:199
		_go_fuzz_dep_.CoverTab[527692]++
//line /snap/go/10455/src/net/dial.go:199
		_go_fuzz_dep_.CoverTab[4982]++
							return deadline, nil
//line /snap/go/10455/src/net/dial.go:200
		// _ = "end of CoverTab[4982]"
	} else {
//line /snap/go/10455/src/net/dial.go:201
		_go_fuzz_dep_.CoverTab[527693]++
//line /snap/go/10455/src/net/dial.go:201
		_go_fuzz_dep_.CoverTab[4983]++
//line /snap/go/10455/src/net/dial.go:201
		// _ = "end of CoverTab[4983]"
//line /snap/go/10455/src/net/dial.go:201
	}
//line /snap/go/10455/src/net/dial.go:201
	// _ = "end of CoverTab[4978]"
//line /snap/go/10455/src/net/dial.go:201
	_go_fuzz_dep_.CoverTab[4979]++
						timeRemaining := deadline.Sub(now)
						if timeRemaining <= 0 {
//line /snap/go/10455/src/net/dial.go:203
		_go_fuzz_dep_.CoverTab[527694]++
//line /snap/go/10455/src/net/dial.go:203
		_go_fuzz_dep_.CoverTab[4984]++
							return time.Time{}, errTimeout
//line /snap/go/10455/src/net/dial.go:204
		// _ = "end of CoverTab[4984]"
	} else {
//line /snap/go/10455/src/net/dial.go:205
		_go_fuzz_dep_.CoverTab[527695]++
//line /snap/go/10455/src/net/dial.go:205
		_go_fuzz_dep_.CoverTab[4985]++
//line /snap/go/10455/src/net/dial.go:205
		// _ = "end of CoverTab[4985]"
//line /snap/go/10455/src/net/dial.go:205
	}
//line /snap/go/10455/src/net/dial.go:205
	// _ = "end of CoverTab[4979]"
//line /snap/go/10455/src/net/dial.go:205
	_go_fuzz_dep_.CoverTab[4980]++

						timeout := timeRemaining / time.Duration(addrsRemaining)
	// If the time per address is too short, steal from the end of the list.
	const saneMinimum = 2 * time.Second
	if timeout < saneMinimum {
//line /snap/go/10455/src/net/dial.go:210
		_go_fuzz_dep_.CoverTab[527696]++
//line /snap/go/10455/src/net/dial.go:210
		_go_fuzz_dep_.CoverTab[4986]++
							if timeRemaining < saneMinimum {
//line /snap/go/10455/src/net/dial.go:211
			_go_fuzz_dep_.CoverTab[527698]++
//line /snap/go/10455/src/net/dial.go:211
			_go_fuzz_dep_.CoverTab[4987]++
								timeout = timeRemaining
//line /snap/go/10455/src/net/dial.go:212
			// _ = "end of CoverTab[4987]"
		} else {
//line /snap/go/10455/src/net/dial.go:213
			_go_fuzz_dep_.CoverTab[527699]++
//line /snap/go/10455/src/net/dial.go:213
			_go_fuzz_dep_.CoverTab[4988]++
								timeout = saneMinimum
//line /snap/go/10455/src/net/dial.go:214
			// _ = "end of CoverTab[4988]"
		}
//line /snap/go/10455/src/net/dial.go:215
		// _ = "end of CoverTab[4986]"
	} else {
//line /snap/go/10455/src/net/dial.go:216
		_go_fuzz_dep_.CoverTab[527697]++
//line /snap/go/10455/src/net/dial.go:216
		_go_fuzz_dep_.CoverTab[4989]++
//line /snap/go/10455/src/net/dial.go:216
		// _ = "end of CoverTab[4989]"
//line /snap/go/10455/src/net/dial.go:216
	}
//line /snap/go/10455/src/net/dial.go:216
	// _ = "end of CoverTab[4980]"
//line /snap/go/10455/src/net/dial.go:216
	_go_fuzz_dep_.CoverTab[4981]++
						return now.Add(timeout), nil
//line /snap/go/10455/src/net/dial.go:217
	// _ = "end of CoverTab[4981]"
}

func (d *Dialer) fallbackDelay() time.Duration {
//line /snap/go/10455/src/net/dial.go:220
	_go_fuzz_dep_.CoverTab[4990]++
						if d.FallbackDelay > 0 {
//line /snap/go/10455/src/net/dial.go:221
		_go_fuzz_dep_.CoverTab[527700]++
//line /snap/go/10455/src/net/dial.go:221
		_go_fuzz_dep_.CoverTab[4991]++
							return d.FallbackDelay
//line /snap/go/10455/src/net/dial.go:222
		// _ = "end of CoverTab[4991]"
	} else {
//line /snap/go/10455/src/net/dial.go:223
		_go_fuzz_dep_.CoverTab[527701]++
//line /snap/go/10455/src/net/dial.go:223
		_go_fuzz_dep_.CoverTab[4992]++
							return 300 * time.Millisecond
//line /snap/go/10455/src/net/dial.go:224
		// _ = "end of CoverTab[4992]"
	}
//line /snap/go/10455/src/net/dial.go:225
	// _ = "end of CoverTab[4990]"
}

func parseNetwork(ctx context.Context, network string, needsProto bool) (afnet string, proto int, err error) {
//line /snap/go/10455/src/net/dial.go:228
	_go_fuzz_dep_.CoverTab[4993]++
						i := last(network, ':')
						if i < 0 {
//line /snap/go/10455/src/net/dial.go:230
		_go_fuzz_dep_.CoverTab[527702]++
//line /snap/go/10455/src/net/dial.go:230
		_go_fuzz_dep_.CoverTab[4996]++
							switch network {
		case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/dial.go:232
			_go_fuzz_dep_.CoverTab[527704]++
//line /snap/go/10455/src/net/dial.go:232
			_go_fuzz_dep_.CoverTab[4998]++
//line /snap/go/10455/src/net/dial.go:232
			// _ = "end of CoverTab[4998]"
		case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/dial.go:233
			_go_fuzz_dep_.CoverTab[527705]++
//line /snap/go/10455/src/net/dial.go:233
			_go_fuzz_dep_.CoverTab[4999]++
//line /snap/go/10455/src/net/dial.go:233
			// _ = "end of CoverTab[4999]"
		case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/dial.go:234
			_go_fuzz_dep_.CoverTab[527706]++
//line /snap/go/10455/src/net/dial.go:234
			_go_fuzz_dep_.CoverTab[5000]++
								if needsProto {
//line /snap/go/10455/src/net/dial.go:235
				_go_fuzz_dep_.CoverTab[527709]++
//line /snap/go/10455/src/net/dial.go:235
				_go_fuzz_dep_.CoverTab[5003]++
									return "", 0, UnknownNetworkError(network)
//line /snap/go/10455/src/net/dial.go:236
				// _ = "end of CoverTab[5003]"
			} else {
//line /snap/go/10455/src/net/dial.go:237
				_go_fuzz_dep_.CoverTab[527710]++
//line /snap/go/10455/src/net/dial.go:237
				_go_fuzz_dep_.CoverTab[5004]++
//line /snap/go/10455/src/net/dial.go:237
				// _ = "end of CoverTab[5004]"
//line /snap/go/10455/src/net/dial.go:237
			}
//line /snap/go/10455/src/net/dial.go:237
			// _ = "end of CoverTab[5000]"
		case "unix", "unixgram", "unixpacket":
//line /snap/go/10455/src/net/dial.go:238
			_go_fuzz_dep_.CoverTab[527707]++
//line /snap/go/10455/src/net/dial.go:238
			_go_fuzz_dep_.CoverTab[5001]++
//line /snap/go/10455/src/net/dial.go:238
			// _ = "end of CoverTab[5001]"
		default:
//line /snap/go/10455/src/net/dial.go:239
			_go_fuzz_dep_.CoverTab[527708]++
//line /snap/go/10455/src/net/dial.go:239
			_go_fuzz_dep_.CoverTab[5002]++
								return "", 0, UnknownNetworkError(network)
//line /snap/go/10455/src/net/dial.go:240
			// _ = "end of CoverTab[5002]"
		}
//line /snap/go/10455/src/net/dial.go:241
		// _ = "end of CoverTab[4996]"
//line /snap/go/10455/src/net/dial.go:241
		_go_fuzz_dep_.CoverTab[4997]++
							return network, 0, nil
//line /snap/go/10455/src/net/dial.go:242
		// _ = "end of CoverTab[4997]"
	} else {
//line /snap/go/10455/src/net/dial.go:243
		_go_fuzz_dep_.CoverTab[527703]++
//line /snap/go/10455/src/net/dial.go:243
		_go_fuzz_dep_.CoverTab[5005]++
//line /snap/go/10455/src/net/dial.go:243
		// _ = "end of CoverTab[5005]"
//line /snap/go/10455/src/net/dial.go:243
	}
//line /snap/go/10455/src/net/dial.go:243
	// _ = "end of CoverTab[4993]"
//line /snap/go/10455/src/net/dial.go:243
	_go_fuzz_dep_.CoverTab[4994]++
						afnet = network[:i]
						switch afnet {
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/dial.go:246
		_go_fuzz_dep_.CoverTab[527711]++
//line /snap/go/10455/src/net/dial.go:246
		_go_fuzz_dep_.CoverTab[5006]++
							protostr := network[i+1:]
							proto, i, ok := dtoi(protostr)
							if !ok || func() bool {
//line /snap/go/10455/src/net/dial.go:249
			_go_fuzz_dep_.CoverTab[5009]++
//line /snap/go/10455/src/net/dial.go:249
			return i != len(protostr)
//line /snap/go/10455/src/net/dial.go:249
			// _ = "end of CoverTab[5009]"
//line /snap/go/10455/src/net/dial.go:249
		}() {
//line /snap/go/10455/src/net/dial.go:249
			_go_fuzz_dep_.CoverTab[527713]++
//line /snap/go/10455/src/net/dial.go:249
			_go_fuzz_dep_.CoverTab[5010]++
								proto, err = lookupProtocol(ctx, protostr)
								if err != nil {
//line /snap/go/10455/src/net/dial.go:251
				_go_fuzz_dep_.CoverTab[527715]++
//line /snap/go/10455/src/net/dial.go:251
				_go_fuzz_dep_.CoverTab[5011]++
									return "", 0, err
//line /snap/go/10455/src/net/dial.go:252
				// _ = "end of CoverTab[5011]"
			} else {
//line /snap/go/10455/src/net/dial.go:253
				_go_fuzz_dep_.CoverTab[527716]++
//line /snap/go/10455/src/net/dial.go:253
				_go_fuzz_dep_.CoverTab[5012]++
//line /snap/go/10455/src/net/dial.go:253
				// _ = "end of CoverTab[5012]"
//line /snap/go/10455/src/net/dial.go:253
			}
//line /snap/go/10455/src/net/dial.go:253
			// _ = "end of CoverTab[5010]"
		} else {
//line /snap/go/10455/src/net/dial.go:254
			_go_fuzz_dep_.CoverTab[527714]++
//line /snap/go/10455/src/net/dial.go:254
			_go_fuzz_dep_.CoverTab[5013]++
//line /snap/go/10455/src/net/dial.go:254
			// _ = "end of CoverTab[5013]"
//line /snap/go/10455/src/net/dial.go:254
		}
//line /snap/go/10455/src/net/dial.go:254
		// _ = "end of CoverTab[5006]"
//line /snap/go/10455/src/net/dial.go:254
		_go_fuzz_dep_.CoverTab[5007]++
							return afnet, proto, nil
//line /snap/go/10455/src/net/dial.go:255
		// _ = "end of CoverTab[5007]"
//line /snap/go/10455/src/net/dial.go:255
	default:
//line /snap/go/10455/src/net/dial.go:255
		_go_fuzz_dep_.CoverTab[527712]++
//line /snap/go/10455/src/net/dial.go:255
		_go_fuzz_dep_.CoverTab[5008]++
//line /snap/go/10455/src/net/dial.go:255
		// _ = "end of CoverTab[5008]"
	}
//line /snap/go/10455/src/net/dial.go:256
	// _ = "end of CoverTab[4994]"
//line /snap/go/10455/src/net/dial.go:256
	_go_fuzz_dep_.CoverTab[4995]++
						return "", 0, UnknownNetworkError(network)
//line /snap/go/10455/src/net/dial.go:257
	// _ = "end of CoverTab[4995]"
}

// resolveAddrList resolves addr using hint and returns a list of
//line /snap/go/10455/src/net/dial.go:260
// addresses. The result contains at least one address when error is
//line /snap/go/10455/src/net/dial.go:260
// nil.
//line /snap/go/10455/src/net/dial.go:263
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, addr string, hint Addr) (addrList, error) {
//line /snap/go/10455/src/net/dial.go:263
	_go_fuzz_dep_.CoverTab[5014]++
						afnet, _, err := parseNetwork(ctx, network, true)
						if err != nil {
//line /snap/go/10455/src/net/dial.go:265
		_go_fuzz_dep_.CoverTab[527717]++
//line /snap/go/10455/src/net/dial.go:265
		_go_fuzz_dep_.CoverTab[5022]++
							return nil, err
//line /snap/go/10455/src/net/dial.go:266
		// _ = "end of CoverTab[5022]"
	} else {
//line /snap/go/10455/src/net/dial.go:267
		_go_fuzz_dep_.CoverTab[527718]++
//line /snap/go/10455/src/net/dial.go:267
		_go_fuzz_dep_.CoverTab[5023]++
//line /snap/go/10455/src/net/dial.go:267
		// _ = "end of CoverTab[5023]"
//line /snap/go/10455/src/net/dial.go:267
	}
//line /snap/go/10455/src/net/dial.go:267
	// _ = "end of CoverTab[5014]"
//line /snap/go/10455/src/net/dial.go:267
	_go_fuzz_dep_.CoverTab[5015]++
						if op == "dial" && func() bool {
//line /snap/go/10455/src/net/dial.go:268
		_go_fuzz_dep_.CoverTab[5024]++
//line /snap/go/10455/src/net/dial.go:268
		return addr == ""
//line /snap/go/10455/src/net/dial.go:268
		// _ = "end of CoverTab[5024]"
//line /snap/go/10455/src/net/dial.go:268
	}() {
//line /snap/go/10455/src/net/dial.go:268
		_go_fuzz_dep_.CoverTab[527719]++
//line /snap/go/10455/src/net/dial.go:268
		_go_fuzz_dep_.CoverTab[5025]++
							return nil, errMissingAddress
//line /snap/go/10455/src/net/dial.go:269
		// _ = "end of CoverTab[5025]"
	} else {
//line /snap/go/10455/src/net/dial.go:270
		_go_fuzz_dep_.CoverTab[527720]++
//line /snap/go/10455/src/net/dial.go:270
		_go_fuzz_dep_.CoverTab[5026]++
//line /snap/go/10455/src/net/dial.go:270
		// _ = "end of CoverTab[5026]"
//line /snap/go/10455/src/net/dial.go:270
	}
//line /snap/go/10455/src/net/dial.go:270
	// _ = "end of CoverTab[5015]"
//line /snap/go/10455/src/net/dial.go:270
	_go_fuzz_dep_.CoverTab[5016]++
						switch afnet {
	case "unix", "unixgram", "unixpacket":
//line /snap/go/10455/src/net/dial.go:272
		_go_fuzz_dep_.CoverTab[527721]++
//line /snap/go/10455/src/net/dial.go:272
		_go_fuzz_dep_.CoverTab[5027]++
							addr, err := ResolveUnixAddr(afnet, addr)
							if err != nil {
//line /snap/go/10455/src/net/dial.go:274
			_go_fuzz_dep_.CoverTab[527723]++
//line /snap/go/10455/src/net/dial.go:274
			_go_fuzz_dep_.CoverTab[5031]++
								return nil, err
//line /snap/go/10455/src/net/dial.go:275
			// _ = "end of CoverTab[5031]"
		} else {
//line /snap/go/10455/src/net/dial.go:276
			_go_fuzz_dep_.CoverTab[527724]++
//line /snap/go/10455/src/net/dial.go:276
			_go_fuzz_dep_.CoverTab[5032]++
//line /snap/go/10455/src/net/dial.go:276
			// _ = "end of CoverTab[5032]"
//line /snap/go/10455/src/net/dial.go:276
		}
//line /snap/go/10455/src/net/dial.go:276
		// _ = "end of CoverTab[5027]"
//line /snap/go/10455/src/net/dial.go:276
		_go_fuzz_dep_.CoverTab[5028]++
							if op == "dial" && func() bool {
//line /snap/go/10455/src/net/dial.go:277
			_go_fuzz_dep_.CoverTab[5033]++
//line /snap/go/10455/src/net/dial.go:277
			return hint != nil
//line /snap/go/10455/src/net/dial.go:277
			// _ = "end of CoverTab[5033]"
//line /snap/go/10455/src/net/dial.go:277
		}() && func() bool {
//line /snap/go/10455/src/net/dial.go:277
			_go_fuzz_dep_.CoverTab[5034]++
//line /snap/go/10455/src/net/dial.go:277
			return addr.Network() != hint.Network()
//line /snap/go/10455/src/net/dial.go:277
			// _ = "end of CoverTab[5034]"
//line /snap/go/10455/src/net/dial.go:277
		}() {
//line /snap/go/10455/src/net/dial.go:277
			_go_fuzz_dep_.CoverTab[527725]++
//line /snap/go/10455/src/net/dial.go:277
			_go_fuzz_dep_.CoverTab[5035]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /snap/go/10455/src/net/dial.go:278
			// _ = "end of CoverTab[5035]"
		} else {
//line /snap/go/10455/src/net/dial.go:279
			_go_fuzz_dep_.CoverTab[527726]++
//line /snap/go/10455/src/net/dial.go:279
			_go_fuzz_dep_.CoverTab[5036]++
//line /snap/go/10455/src/net/dial.go:279
			// _ = "end of CoverTab[5036]"
//line /snap/go/10455/src/net/dial.go:279
		}
//line /snap/go/10455/src/net/dial.go:279
		// _ = "end of CoverTab[5028]"
//line /snap/go/10455/src/net/dial.go:279
		_go_fuzz_dep_.CoverTab[5029]++
							return addrList{addr}, nil
//line /snap/go/10455/src/net/dial.go:280
		// _ = "end of CoverTab[5029]"
//line /snap/go/10455/src/net/dial.go:280
	default:
//line /snap/go/10455/src/net/dial.go:280
		_go_fuzz_dep_.CoverTab[527722]++
//line /snap/go/10455/src/net/dial.go:280
		_go_fuzz_dep_.CoverTab[5030]++
//line /snap/go/10455/src/net/dial.go:280
		// _ = "end of CoverTab[5030]"
	}
//line /snap/go/10455/src/net/dial.go:281
	// _ = "end of CoverTab[5016]"
//line /snap/go/10455/src/net/dial.go:281
	_go_fuzz_dep_.CoverTab[5017]++
						addrs, err := r.internetAddrList(ctx, afnet, addr)
						if err != nil || func() bool {
//line /snap/go/10455/src/net/dial.go:283
		_go_fuzz_dep_.CoverTab[5037]++
//line /snap/go/10455/src/net/dial.go:283
		return op != "dial"
//line /snap/go/10455/src/net/dial.go:283
		// _ = "end of CoverTab[5037]"
//line /snap/go/10455/src/net/dial.go:283
	}() || func() bool {
//line /snap/go/10455/src/net/dial.go:283
		_go_fuzz_dep_.CoverTab[5038]++
//line /snap/go/10455/src/net/dial.go:283
		return hint == nil
//line /snap/go/10455/src/net/dial.go:283
		// _ = "end of CoverTab[5038]"
//line /snap/go/10455/src/net/dial.go:283
	}() {
//line /snap/go/10455/src/net/dial.go:283
		_go_fuzz_dep_.CoverTab[527727]++
//line /snap/go/10455/src/net/dial.go:283
		_go_fuzz_dep_.CoverTab[5039]++
							return addrs, err
//line /snap/go/10455/src/net/dial.go:284
		// _ = "end of CoverTab[5039]"
	} else {
//line /snap/go/10455/src/net/dial.go:285
		_go_fuzz_dep_.CoverTab[527728]++
//line /snap/go/10455/src/net/dial.go:285
		_go_fuzz_dep_.CoverTab[5040]++
//line /snap/go/10455/src/net/dial.go:285
		// _ = "end of CoverTab[5040]"
//line /snap/go/10455/src/net/dial.go:285
	}
//line /snap/go/10455/src/net/dial.go:285
	// _ = "end of CoverTab[5017]"
//line /snap/go/10455/src/net/dial.go:285
	_go_fuzz_dep_.CoverTab[5018]++
						var (
		tcp		*TCPAddr
		udp		*UDPAddr
		ip		*IPAddr
		wildcard	bool
	)
	switch hint := hint.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/dial.go:293
		_go_fuzz_dep_.CoverTab[527729]++
//line /snap/go/10455/src/net/dial.go:293
		_go_fuzz_dep_.CoverTab[5041]++
							tcp = hint
							wildcard = tcp.isWildcard()
//line /snap/go/10455/src/net/dial.go:295
		// _ = "end of CoverTab[5041]"
	case *UDPAddr:
//line /snap/go/10455/src/net/dial.go:296
		_go_fuzz_dep_.CoverTab[527730]++
//line /snap/go/10455/src/net/dial.go:296
		_go_fuzz_dep_.CoverTab[5042]++
							udp = hint
							wildcard = udp.isWildcard()
//line /snap/go/10455/src/net/dial.go:298
		// _ = "end of CoverTab[5042]"
	case *IPAddr:
//line /snap/go/10455/src/net/dial.go:299
		_go_fuzz_dep_.CoverTab[527731]++
//line /snap/go/10455/src/net/dial.go:299
		_go_fuzz_dep_.CoverTab[5043]++
							ip = hint
							wildcard = ip.isWildcard()
//line /snap/go/10455/src/net/dial.go:301
		// _ = "end of CoverTab[5043]"
	}
//line /snap/go/10455/src/net/dial.go:302
	// _ = "end of CoverTab[5018]"
//line /snap/go/10455/src/net/dial.go:302
	_go_fuzz_dep_.CoverTab[5019]++
						naddrs := addrs[:0]
//line /snap/go/10455/src/net/dial.go:303
	_go_fuzz_dep_.CoverTab[786649] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/dial.go:304
		if _go_fuzz_dep_.CoverTab[786649] == 0 {
//line /snap/go/10455/src/net/dial.go:304
			_go_fuzz_dep_.CoverTab[527819]++
//line /snap/go/10455/src/net/dial.go:304
		} else {
//line /snap/go/10455/src/net/dial.go:304
			_go_fuzz_dep_.CoverTab[527820]++
//line /snap/go/10455/src/net/dial.go:304
		}
//line /snap/go/10455/src/net/dial.go:304
		_go_fuzz_dep_.CoverTab[786649] = 1
//line /snap/go/10455/src/net/dial.go:304
		_go_fuzz_dep_.CoverTab[5044]++
							if addr.Network() != hint.Network() {
//line /snap/go/10455/src/net/dial.go:305
			_go_fuzz_dep_.CoverTab[527732]++
//line /snap/go/10455/src/net/dial.go:305
			_go_fuzz_dep_.CoverTab[5046]++
								return nil, &AddrError{Err: "mismatched local address type", Addr: hint.String()}
//line /snap/go/10455/src/net/dial.go:306
			// _ = "end of CoverTab[5046]"
		} else {
//line /snap/go/10455/src/net/dial.go:307
			_go_fuzz_dep_.CoverTab[527733]++
//line /snap/go/10455/src/net/dial.go:307
			_go_fuzz_dep_.CoverTab[5047]++
//line /snap/go/10455/src/net/dial.go:307
			// _ = "end of CoverTab[5047]"
//line /snap/go/10455/src/net/dial.go:307
		}
//line /snap/go/10455/src/net/dial.go:307
		// _ = "end of CoverTab[5044]"
//line /snap/go/10455/src/net/dial.go:307
		_go_fuzz_dep_.CoverTab[5045]++
							switch addr := addr.(type) {
		case *TCPAddr:
//line /snap/go/10455/src/net/dial.go:309
			_go_fuzz_dep_.CoverTab[527734]++
//line /snap/go/10455/src/net/dial.go:309
			_go_fuzz_dep_.CoverTab[5048]++
								if !wildcard && func() bool {
//line /snap/go/10455/src/net/dial.go:310
				_go_fuzz_dep_.CoverTab[5054]++
//line /snap/go/10455/src/net/dial.go:310
				return !addr.isWildcard()
//line /snap/go/10455/src/net/dial.go:310
				// _ = "end of CoverTab[5054]"
//line /snap/go/10455/src/net/dial.go:310
			}() && func() bool {
//line /snap/go/10455/src/net/dial.go:310
				_go_fuzz_dep_.CoverTab[5055]++
//line /snap/go/10455/src/net/dial.go:310
				return !addr.IP.matchAddrFamily(tcp.IP)
//line /snap/go/10455/src/net/dial.go:310
				// _ = "end of CoverTab[5055]"
//line /snap/go/10455/src/net/dial.go:310
			}() {
//line /snap/go/10455/src/net/dial.go:310
				_go_fuzz_dep_.CoverTab[527737]++
//line /snap/go/10455/src/net/dial.go:310
				_go_fuzz_dep_.CoverTab[5056]++
									continue
//line /snap/go/10455/src/net/dial.go:311
				// _ = "end of CoverTab[5056]"
			} else {
//line /snap/go/10455/src/net/dial.go:312
				_go_fuzz_dep_.CoverTab[527738]++
//line /snap/go/10455/src/net/dial.go:312
				_go_fuzz_dep_.CoverTab[5057]++
//line /snap/go/10455/src/net/dial.go:312
				// _ = "end of CoverTab[5057]"
//line /snap/go/10455/src/net/dial.go:312
			}
//line /snap/go/10455/src/net/dial.go:312
			// _ = "end of CoverTab[5048]"
//line /snap/go/10455/src/net/dial.go:312
			_go_fuzz_dep_.CoverTab[5049]++
								naddrs = append(naddrs, addr)
//line /snap/go/10455/src/net/dial.go:313
			// _ = "end of CoverTab[5049]"
		case *UDPAddr:
//line /snap/go/10455/src/net/dial.go:314
			_go_fuzz_dep_.CoverTab[527735]++
//line /snap/go/10455/src/net/dial.go:314
			_go_fuzz_dep_.CoverTab[5050]++
								if !wildcard && func() bool {
//line /snap/go/10455/src/net/dial.go:315
				_go_fuzz_dep_.CoverTab[5058]++
//line /snap/go/10455/src/net/dial.go:315
				return !addr.isWildcard()
//line /snap/go/10455/src/net/dial.go:315
				// _ = "end of CoverTab[5058]"
//line /snap/go/10455/src/net/dial.go:315
			}() && func() bool {
//line /snap/go/10455/src/net/dial.go:315
				_go_fuzz_dep_.CoverTab[5059]++
//line /snap/go/10455/src/net/dial.go:315
				return !addr.IP.matchAddrFamily(udp.IP)
//line /snap/go/10455/src/net/dial.go:315
				// _ = "end of CoverTab[5059]"
//line /snap/go/10455/src/net/dial.go:315
			}() {
//line /snap/go/10455/src/net/dial.go:315
				_go_fuzz_dep_.CoverTab[527739]++
//line /snap/go/10455/src/net/dial.go:315
				_go_fuzz_dep_.CoverTab[5060]++
									continue
//line /snap/go/10455/src/net/dial.go:316
				// _ = "end of CoverTab[5060]"
			} else {
//line /snap/go/10455/src/net/dial.go:317
				_go_fuzz_dep_.CoverTab[527740]++
//line /snap/go/10455/src/net/dial.go:317
				_go_fuzz_dep_.CoverTab[5061]++
//line /snap/go/10455/src/net/dial.go:317
				// _ = "end of CoverTab[5061]"
//line /snap/go/10455/src/net/dial.go:317
			}
//line /snap/go/10455/src/net/dial.go:317
			// _ = "end of CoverTab[5050]"
//line /snap/go/10455/src/net/dial.go:317
			_go_fuzz_dep_.CoverTab[5051]++
								naddrs = append(naddrs, addr)
//line /snap/go/10455/src/net/dial.go:318
			// _ = "end of CoverTab[5051]"
		case *IPAddr:
//line /snap/go/10455/src/net/dial.go:319
			_go_fuzz_dep_.CoverTab[527736]++
//line /snap/go/10455/src/net/dial.go:319
			_go_fuzz_dep_.CoverTab[5052]++
								if !wildcard && func() bool {
//line /snap/go/10455/src/net/dial.go:320
				_go_fuzz_dep_.CoverTab[5062]++
//line /snap/go/10455/src/net/dial.go:320
				return !addr.isWildcard()
//line /snap/go/10455/src/net/dial.go:320
				// _ = "end of CoverTab[5062]"
//line /snap/go/10455/src/net/dial.go:320
			}() && func() bool {
//line /snap/go/10455/src/net/dial.go:320
				_go_fuzz_dep_.CoverTab[5063]++
//line /snap/go/10455/src/net/dial.go:320
				return !addr.IP.matchAddrFamily(ip.IP)
//line /snap/go/10455/src/net/dial.go:320
				// _ = "end of CoverTab[5063]"
//line /snap/go/10455/src/net/dial.go:320
			}() {
//line /snap/go/10455/src/net/dial.go:320
				_go_fuzz_dep_.CoverTab[527741]++
//line /snap/go/10455/src/net/dial.go:320
				_go_fuzz_dep_.CoverTab[5064]++
									continue
//line /snap/go/10455/src/net/dial.go:321
				// _ = "end of CoverTab[5064]"
			} else {
//line /snap/go/10455/src/net/dial.go:322
				_go_fuzz_dep_.CoverTab[527742]++
//line /snap/go/10455/src/net/dial.go:322
				_go_fuzz_dep_.CoverTab[5065]++
//line /snap/go/10455/src/net/dial.go:322
				// _ = "end of CoverTab[5065]"
//line /snap/go/10455/src/net/dial.go:322
			}
//line /snap/go/10455/src/net/dial.go:322
			// _ = "end of CoverTab[5052]"
//line /snap/go/10455/src/net/dial.go:322
			_go_fuzz_dep_.CoverTab[5053]++
								naddrs = append(naddrs, addr)
//line /snap/go/10455/src/net/dial.go:323
			// _ = "end of CoverTab[5053]"
		}
//line /snap/go/10455/src/net/dial.go:324
		// _ = "end of CoverTab[5045]"
	}
//line /snap/go/10455/src/net/dial.go:325
	if _go_fuzz_dep_.CoverTab[786649] == 0 {
//line /snap/go/10455/src/net/dial.go:325
		_go_fuzz_dep_.CoverTab[527821]++
//line /snap/go/10455/src/net/dial.go:325
	} else {
//line /snap/go/10455/src/net/dial.go:325
		_go_fuzz_dep_.CoverTab[527822]++
//line /snap/go/10455/src/net/dial.go:325
	}
//line /snap/go/10455/src/net/dial.go:325
	// _ = "end of CoverTab[5019]"
//line /snap/go/10455/src/net/dial.go:325
	_go_fuzz_dep_.CoverTab[5020]++
						if len(naddrs) == 0 {
//line /snap/go/10455/src/net/dial.go:326
		_go_fuzz_dep_.CoverTab[527743]++
//line /snap/go/10455/src/net/dial.go:326
		_go_fuzz_dep_.CoverTab[5066]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: hint.String()}
//line /snap/go/10455/src/net/dial.go:327
		// _ = "end of CoverTab[5066]"
	} else {
//line /snap/go/10455/src/net/dial.go:328
		_go_fuzz_dep_.CoverTab[527744]++
//line /snap/go/10455/src/net/dial.go:328
		_go_fuzz_dep_.CoverTab[5067]++
//line /snap/go/10455/src/net/dial.go:328
		// _ = "end of CoverTab[5067]"
//line /snap/go/10455/src/net/dial.go:328
	}
//line /snap/go/10455/src/net/dial.go:328
	// _ = "end of CoverTab[5020]"
//line /snap/go/10455/src/net/dial.go:328
	_go_fuzz_dep_.CoverTab[5021]++
						return naddrs, nil
//line /snap/go/10455/src/net/dial.go:329
	// _ = "end of CoverTab[5021]"
}

// MultipathTCP reports whether MPTCP will be used.
//line /snap/go/10455/src/net/dial.go:332
//
//line /snap/go/10455/src/net/dial.go:332
// This method doesn't check if MPTCP is supported by the operating
//line /snap/go/10455/src/net/dial.go:332
// system or not.
//line /snap/go/10455/src/net/dial.go:336
func (d *Dialer) MultipathTCP() bool {
//line /snap/go/10455/src/net/dial.go:336
	_go_fuzz_dep_.CoverTab[5068]++
						return d.mptcpStatus.get()
//line /snap/go/10455/src/net/dial.go:337
	// _ = "end of CoverTab[5068]"
}

// SetMultipathTCP directs the Dial methods to use, or not use, MPTCP,
//line /snap/go/10455/src/net/dial.go:340
// if supported by the operating system. This method overrides the
//line /snap/go/10455/src/net/dial.go:340
// system default and the GODEBUG=multipathtcp=... setting if any.
//line /snap/go/10455/src/net/dial.go:340
//
//line /snap/go/10455/src/net/dial.go:340
// If MPTCP is not available on the host or not supported by the server,
//line /snap/go/10455/src/net/dial.go:340
// the Dial methods will fall back to TCP.
//line /snap/go/10455/src/net/dial.go:346
func (d *Dialer) SetMultipathTCP(use bool) {
//line /snap/go/10455/src/net/dial.go:346
	_go_fuzz_dep_.CoverTab[5069]++
						d.mptcpStatus.set(use)
//line /snap/go/10455/src/net/dial.go:347
	// _ = "end of CoverTab[5069]"
}

// Dial connects to the address on the named network.
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
//line /snap/go/10455/src/net/dial.go:350
// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
//line /snap/go/10455/src/net/dial.go:350
// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
//line /snap/go/10455/src/net/dial.go:350
// "unixpacket".
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// For TCP and UDP networks, the address has the form "host:port".
//line /snap/go/10455/src/net/dial.go:350
// The host must be a literal IP address, or a host name that can be
//line /snap/go/10455/src/net/dial.go:350
// resolved to IP addresses.
//line /snap/go/10455/src/net/dial.go:350
// The port must be a literal port number or a service name.
//line /snap/go/10455/src/net/dial.go:350
// If the host is a literal IPv6 address it must be enclosed in square
//line /snap/go/10455/src/net/dial.go:350
// brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
//line /snap/go/10455/src/net/dial.go:350
// The zone specifies the scope of the literal IPv6 address as defined
//line /snap/go/10455/src/net/dial.go:350
// in RFC 4007.
//line /snap/go/10455/src/net/dial.go:350
// The functions JoinHostPort and SplitHostPort manipulate a pair of
//line /snap/go/10455/src/net/dial.go:350
// host and port in this form.
//line /snap/go/10455/src/net/dial.go:350
// When using TCP, and the host resolves to multiple IP addresses,
//line /snap/go/10455/src/net/dial.go:350
// Dial will try each IP address in order until one succeeds.
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// Examples:
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
//	Dial("tcp", "golang.org:http")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("tcp", "192.0.2.1:http")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("tcp", "198.51.100.1:80")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("udp", "[2001:db8::1]:domain")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("udp", "[fe80::1%lo0]:53")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("tcp", ":80")
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// For IP networks, the network must be "ip", "ip4" or "ip6" followed
//line /snap/go/10455/src/net/dial.go:350
// by a colon and a literal protocol number or a protocol name, and
//line /snap/go/10455/src/net/dial.go:350
// the address has the form "host". The host must be a literal IP
//line /snap/go/10455/src/net/dial.go:350
// address or a literal IPv6 address with zone.
//line /snap/go/10455/src/net/dial.go:350
// It depends on each operating system how the operating system
//line /snap/go/10455/src/net/dial.go:350
// behaves with a non-well known protocol number such as "0" or "255".
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// Examples:
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
//	Dial("ip4:1", "192.0.2.1")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("ip6:ipv6-icmp", "2001:db8::1")
//line /snap/go/10455/src/net/dial.go:350
//	Dial("ip6:58", "fe80::1%lo0")
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// For TCP, UDP and IP networks, if the host is empty or a literal
//line /snap/go/10455/src/net/dial.go:350
// unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
//line /snap/go/10455/src/net/dial.go:350
// TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
//line /snap/go/10455/src/net/dial.go:350
// assumed.
//line /snap/go/10455/src/net/dial.go:350
//
//line /snap/go/10455/src/net/dial.go:350
// For Unix networks, the address must be a file system path.
//line /snap/go/10455/src/net/dial.go:398
func Dial(network, address string) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:398
	_go_fuzz_dep_.CoverTab[5070]++
						var d Dialer
						return d.Dial(network, address)
//line /snap/go/10455/src/net/dial.go:400
	// _ = "end of CoverTab[5070]"
}

// DialTimeout acts like Dial but takes a timeout.
//line /snap/go/10455/src/net/dial.go:403
//
//line /snap/go/10455/src/net/dial.go:403
// The timeout includes name resolution, if required.
//line /snap/go/10455/src/net/dial.go:403
// When using TCP, and the host in the address parameter resolves to
//line /snap/go/10455/src/net/dial.go:403
// multiple IP addresses, the timeout is spread over each consecutive
//line /snap/go/10455/src/net/dial.go:403
// dial, such that each is given an appropriate fraction of the time
//line /snap/go/10455/src/net/dial.go:403
// to connect.
//line /snap/go/10455/src/net/dial.go:403
//
//line /snap/go/10455/src/net/dial.go:403
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/dial.go:403
// parameters.
//line /snap/go/10455/src/net/dial.go:413
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:413
	_go_fuzz_dep_.CoverTab[5071]++
						d := Dialer{Timeout: timeout}
						return d.Dial(network, address)
//line /snap/go/10455/src/net/dial.go:415
	// _ = "end of CoverTab[5071]"
}

// sysDialer contains a Dial's parameters and configuration.
type sysDialer struct {
	Dialer
	network, address	string
	testHookDialTCP		func(ctx context.Context, net string, laddr, raddr *TCPAddr) (*TCPConn, error)
}

// Dial connects to the address on the named network.
//line /snap/go/10455/src/net/dial.go:425
//
//line /snap/go/10455/src/net/dial.go:425
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/dial.go:425
// parameters.
//line /snap/go/10455/src/net/dial.go:425
//
//line /snap/go/10455/src/net/dial.go:425
// Dial uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/dial.go:425
// DialContext.
//line /snap/go/10455/src/net/dial.go:432
func (d *Dialer) Dial(network, address string) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:432
	_go_fuzz_dep_.CoverTab[5072]++
						return d.DialContext(context.Background(), network, address)
//line /snap/go/10455/src/net/dial.go:433
	// _ = "end of CoverTab[5072]"
}

// DialContext connects to the address on the named network using
//line /snap/go/10455/src/net/dial.go:436
// the provided context.
//line /snap/go/10455/src/net/dial.go:436
//
//line /snap/go/10455/src/net/dial.go:436
// The provided Context must be non-nil. If the context expires before
//line /snap/go/10455/src/net/dial.go:436
// the connection is complete, an error is returned. Once successfully
//line /snap/go/10455/src/net/dial.go:436
// connected, any expiration of the context will not affect the
//line /snap/go/10455/src/net/dial.go:436
// connection.
//line /snap/go/10455/src/net/dial.go:436
//
//line /snap/go/10455/src/net/dial.go:436
// When using TCP, and the host in the address parameter resolves to multiple
//line /snap/go/10455/src/net/dial.go:436
// network addresses, any dial timeout (from d.Timeout or ctx) is spread
//line /snap/go/10455/src/net/dial.go:436
// over each consecutive dial, such that each is given an appropriate
//line /snap/go/10455/src/net/dial.go:436
// fraction of the time to connect.
//line /snap/go/10455/src/net/dial.go:436
// For example, if a host has 4 IP addresses and the timeout is 1 minute,
//line /snap/go/10455/src/net/dial.go:436
// the connect to each single address will be given 15 seconds to complete
//line /snap/go/10455/src/net/dial.go:436
// before trying the next one.
//line /snap/go/10455/src/net/dial.go:436
//
//line /snap/go/10455/src/net/dial.go:436
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/dial.go:436
// parameters.
//line /snap/go/10455/src/net/dial.go:454
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:454
	_go_fuzz_dep_.CoverTab[5073]++
						if ctx == nil {
//line /snap/go/10455/src/net/dial.go:455
		_go_fuzz_dep_.CoverTab[527745]++
//line /snap/go/10455/src/net/dial.go:455
		_go_fuzz_dep_.CoverTab[5080]++
							panic("nil context")
//line /snap/go/10455/src/net/dial.go:456
		// _ = "end of CoverTab[5080]"
	} else {
//line /snap/go/10455/src/net/dial.go:457
		_go_fuzz_dep_.CoverTab[527746]++
//line /snap/go/10455/src/net/dial.go:457
		_go_fuzz_dep_.CoverTab[5081]++
//line /snap/go/10455/src/net/dial.go:457
		// _ = "end of CoverTab[5081]"
//line /snap/go/10455/src/net/dial.go:457
	}
//line /snap/go/10455/src/net/dial.go:457
	// _ = "end of CoverTab[5073]"
//line /snap/go/10455/src/net/dial.go:457
	_go_fuzz_dep_.CoverTab[5074]++
						deadline := d.deadline(ctx, time.Now())
						if !deadline.IsZero() {
//line /snap/go/10455/src/net/dial.go:459
		_go_fuzz_dep_.CoverTab[527747]++
//line /snap/go/10455/src/net/dial.go:459
		_go_fuzz_dep_.CoverTab[5082]++
							if d, ok := ctx.Deadline(); !ok || func() bool {
//line /snap/go/10455/src/net/dial.go:460
			_go_fuzz_dep_.CoverTab[5083]++
//line /snap/go/10455/src/net/dial.go:460
			return deadline.Before(d)
//line /snap/go/10455/src/net/dial.go:460
			// _ = "end of CoverTab[5083]"
//line /snap/go/10455/src/net/dial.go:460
		}() {
//line /snap/go/10455/src/net/dial.go:460
			_go_fuzz_dep_.CoverTab[527749]++
//line /snap/go/10455/src/net/dial.go:460
			_go_fuzz_dep_.CoverTab[5084]++
								subCtx, cancel := context.WithDeadline(ctx, deadline)
								defer cancel()
								ctx = subCtx
//line /snap/go/10455/src/net/dial.go:463
			// _ = "end of CoverTab[5084]"
		} else {
//line /snap/go/10455/src/net/dial.go:464
			_go_fuzz_dep_.CoverTab[527750]++
//line /snap/go/10455/src/net/dial.go:464
			_go_fuzz_dep_.CoverTab[5085]++
//line /snap/go/10455/src/net/dial.go:464
			// _ = "end of CoverTab[5085]"
//line /snap/go/10455/src/net/dial.go:464
		}
//line /snap/go/10455/src/net/dial.go:464
		// _ = "end of CoverTab[5082]"
	} else {
//line /snap/go/10455/src/net/dial.go:465
		_go_fuzz_dep_.CoverTab[527748]++
//line /snap/go/10455/src/net/dial.go:465
		_go_fuzz_dep_.CoverTab[5086]++
//line /snap/go/10455/src/net/dial.go:465
		// _ = "end of CoverTab[5086]"
//line /snap/go/10455/src/net/dial.go:465
	}
//line /snap/go/10455/src/net/dial.go:465
	// _ = "end of CoverTab[5074]"
//line /snap/go/10455/src/net/dial.go:465
	_go_fuzz_dep_.CoverTab[5075]++
						if oldCancel := d.Cancel; oldCancel != nil {
//line /snap/go/10455/src/net/dial.go:466
		_go_fuzz_dep_.CoverTab[527751]++
//line /snap/go/10455/src/net/dial.go:466
		_go_fuzz_dep_.CoverTab[5087]++
							subCtx, cancel := context.WithCancel(ctx)
							defer cancel()
//line /snap/go/10455/src/net/dial.go:468
		_curRoutineNum4_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/dial.go:468
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum4_)
							go func() {
//line /snap/go/10455/src/net/dial.go:469
			_go_fuzz_dep_.CoverTab[5089]++
//line /snap/go/10455/src/net/dial.go:469
			defer func() {
//line /snap/go/10455/src/net/dial.go:469
				_go_fuzz_dep_.CoverTab[5090]++
//line /snap/go/10455/src/net/dial.go:469
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum4_)
//line /snap/go/10455/src/net/dial.go:469
				// _ = "end of CoverTab[5090]"
//line /snap/go/10455/src/net/dial.go:469
			}()
								select {
			case <-oldCancel:
//line /snap/go/10455/src/net/dial.go:471
				_go_fuzz_dep_.CoverTab[5091]++
									cancel()
//line /snap/go/10455/src/net/dial.go:472
				// _ = "end of CoverTab[5091]"
			case <-subCtx.Done():
//line /snap/go/10455/src/net/dial.go:473
				_go_fuzz_dep_.CoverTab[5092]++
//line /snap/go/10455/src/net/dial.go:473
				// _ = "end of CoverTab[5092]"
			}
//line /snap/go/10455/src/net/dial.go:474
			// _ = "end of CoverTab[5089]"
		}()
//line /snap/go/10455/src/net/dial.go:475
		// _ = "end of CoverTab[5087]"
//line /snap/go/10455/src/net/dial.go:475
		_go_fuzz_dep_.CoverTab[5088]++
							ctx = subCtx
//line /snap/go/10455/src/net/dial.go:476
		// _ = "end of CoverTab[5088]"
	} else {
//line /snap/go/10455/src/net/dial.go:477
		_go_fuzz_dep_.CoverTab[527752]++
//line /snap/go/10455/src/net/dial.go:477
		_go_fuzz_dep_.CoverTab[5093]++
//line /snap/go/10455/src/net/dial.go:477
		// _ = "end of CoverTab[5093]"
//line /snap/go/10455/src/net/dial.go:477
	}
//line /snap/go/10455/src/net/dial.go:477
	// _ = "end of CoverTab[5075]"
//line /snap/go/10455/src/net/dial.go:477
	_go_fuzz_dep_.CoverTab[5076]++

//line /snap/go/10455/src/net/dial.go:480
	resolveCtx := ctx
	if trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace); trace != nil {
//line /snap/go/10455/src/net/dial.go:481
		_go_fuzz_dep_.CoverTab[527753]++
//line /snap/go/10455/src/net/dial.go:481
		_go_fuzz_dep_.CoverTab[5094]++
							shadow := *trace
							shadow.ConnectStart = nil
							shadow.ConnectDone = nil
							resolveCtx = context.WithValue(resolveCtx, nettrace.TraceKey{}, &shadow)
//line /snap/go/10455/src/net/dial.go:485
		// _ = "end of CoverTab[5094]"
	} else {
//line /snap/go/10455/src/net/dial.go:486
		_go_fuzz_dep_.CoverTab[527754]++
//line /snap/go/10455/src/net/dial.go:486
		_go_fuzz_dep_.CoverTab[5095]++
//line /snap/go/10455/src/net/dial.go:486
		// _ = "end of CoverTab[5095]"
//line /snap/go/10455/src/net/dial.go:486
	}
//line /snap/go/10455/src/net/dial.go:486
	// _ = "end of CoverTab[5076]"
//line /snap/go/10455/src/net/dial.go:486
	_go_fuzz_dep_.CoverTab[5077]++

						addrs, err := d.resolver().resolveAddrList(resolveCtx, "dial", network, address, d.LocalAddr)
						if err != nil {
//line /snap/go/10455/src/net/dial.go:489
		_go_fuzz_dep_.CoverTab[527755]++
//line /snap/go/10455/src/net/dial.go:489
		_go_fuzz_dep_.CoverTab[5096]++
							return nil, &OpError{Op: "dial", Net: network, Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/dial.go:490
		// _ = "end of CoverTab[5096]"
	} else {
//line /snap/go/10455/src/net/dial.go:491
		_go_fuzz_dep_.CoverTab[527756]++
//line /snap/go/10455/src/net/dial.go:491
		_go_fuzz_dep_.CoverTab[5097]++
//line /snap/go/10455/src/net/dial.go:491
		// _ = "end of CoverTab[5097]"
//line /snap/go/10455/src/net/dial.go:491
	}
//line /snap/go/10455/src/net/dial.go:491
	// _ = "end of CoverTab[5077]"
//line /snap/go/10455/src/net/dial.go:491
	_go_fuzz_dep_.CoverTab[5078]++

						sd := &sysDialer{
		Dialer:		*d,
		network:	network,
		address:	address,
	}

	var primaries, fallbacks addrList
	if d.dualStack() && func() bool {
//line /snap/go/10455/src/net/dial.go:500
		_go_fuzz_dep_.CoverTab[5098]++
//line /snap/go/10455/src/net/dial.go:500
		return network == "tcp"
//line /snap/go/10455/src/net/dial.go:500
		// _ = "end of CoverTab[5098]"
//line /snap/go/10455/src/net/dial.go:500
	}() {
//line /snap/go/10455/src/net/dial.go:500
		_go_fuzz_dep_.CoverTab[527757]++
//line /snap/go/10455/src/net/dial.go:500
		_go_fuzz_dep_.CoverTab[5099]++
							primaries, fallbacks = addrs.partition(isIPv4)
//line /snap/go/10455/src/net/dial.go:501
		// _ = "end of CoverTab[5099]"
	} else {
//line /snap/go/10455/src/net/dial.go:502
		_go_fuzz_dep_.CoverTab[527758]++
//line /snap/go/10455/src/net/dial.go:502
		_go_fuzz_dep_.CoverTab[5100]++
							primaries = addrs
//line /snap/go/10455/src/net/dial.go:503
		// _ = "end of CoverTab[5100]"
	}
//line /snap/go/10455/src/net/dial.go:504
	// _ = "end of CoverTab[5078]"
//line /snap/go/10455/src/net/dial.go:504
	_go_fuzz_dep_.CoverTab[5079]++

						return sd.dialParallel(ctx, primaries, fallbacks)
//line /snap/go/10455/src/net/dial.go:506
	// _ = "end of CoverTab[5079]"
}

// dialParallel races two copies of dialSerial, giving the first a
//line /snap/go/10455/src/net/dial.go:509
// head start. It returns the first established connection and
//line /snap/go/10455/src/net/dial.go:509
// closes the others. Otherwise it returns an error from the first
//line /snap/go/10455/src/net/dial.go:509
// primary address.
//line /snap/go/10455/src/net/dial.go:513
func (sd *sysDialer) dialParallel(ctx context.Context, primaries, fallbacks addrList) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:513
	_go_fuzz_dep_.CoverTab[5101]++
						if len(fallbacks) == 0 {
//line /snap/go/10455/src/net/dial.go:514
		_go_fuzz_dep_.CoverTab[527759]++
//line /snap/go/10455/src/net/dial.go:514
		_go_fuzz_dep_.CoverTab[5104]++
							return sd.dialSerial(ctx, primaries)
//line /snap/go/10455/src/net/dial.go:515
		// _ = "end of CoverTab[5104]"
	} else {
//line /snap/go/10455/src/net/dial.go:516
		_go_fuzz_dep_.CoverTab[527760]++
//line /snap/go/10455/src/net/dial.go:516
		_go_fuzz_dep_.CoverTab[5105]++
//line /snap/go/10455/src/net/dial.go:516
		// _ = "end of CoverTab[5105]"
//line /snap/go/10455/src/net/dial.go:516
	}
//line /snap/go/10455/src/net/dial.go:516
	// _ = "end of CoverTab[5101]"
//line /snap/go/10455/src/net/dial.go:516
	_go_fuzz_dep_.CoverTab[5102]++

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
//line /snap/go/10455/src/net/dial.go:529
		_go_fuzz_dep_.CoverTab[5106]++
							ras := primaries
							if !primary {
//line /snap/go/10455/src/net/dial.go:531
			_go_fuzz_dep_.CoverTab[527761]++
//line /snap/go/10455/src/net/dial.go:531
			_go_fuzz_dep_.CoverTab[5108]++
								ras = fallbacks
//line /snap/go/10455/src/net/dial.go:532
			// _ = "end of CoverTab[5108]"
		} else {
//line /snap/go/10455/src/net/dial.go:533
			_go_fuzz_dep_.CoverTab[527762]++
//line /snap/go/10455/src/net/dial.go:533
			_go_fuzz_dep_.CoverTab[5109]++
//line /snap/go/10455/src/net/dial.go:533
			// _ = "end of CoverTab[5109]"
//line /snap/go/10455/src/net/dial.go:533
		}
//line /snap/go/10455/src/net/dial.go:533
		// _ = "end of CoverTab[5106]"
//line /snap/go/10455/src/net/dial.go:533
		_go_fuzz_dep_.CoverTab[5107]++
							c, err := sd.dialSerial(ctx, ras)
							select {
		case results <- dialResult{Conn: c, error: err, primary: primary, done: true}:
//line /snap/go/10455/src/net/dial.go:536
			_go_fuzz_dep_.CoverTab[5110]++
//line /snap/go/10455/src/net/dial.go:536
			// _ = "end of CoverTab[5110]"
		case <-returned:
//line /snap/go/10455/src/net/dial.go:537
			_go_fuzz_dep_.CoverTab[5111]++
								if c != nil {
//line /snap/go/10455/src/net/dial.go:538
				_go_fuzz_dep_.CoverTab[527763]++
//line /snap/go/10455/src/net/dial.go:538
				_go_fuzz_dep_.CoverTab[5112]++
									c.Close()
//line /snap/go/10455/src/net/dial.go:539
				// _ = "end of CoverTab[5112]"
			} else {
//line /snap/go/10455/src/net/dial.go:540
				_go_fuzz_dep_.CoverTab[527764]++
//line /snap/go/10455/src/net/dial.go:540
				_go_fuzz_dep_.CoverTab[5113]++
//line /snap/go/10455/src/net/dial.go:540
				// _ = "end of CoverTab[5113]"
//line /snap/go/10455/src/net/dial.go:540
			}
//line /snap/go/10455/src/net/dial.go:540
			// _ = "end of CoverTab[5111]"
		}
//line /snap/go/10455/src/net/dial.go:541
		// _ = "end of CoverTab[5107]"
	}
//line /snap/go/10455/src/net/dial.go:542
	// _ = "end of CoverTab[5102]"
//line /snap/go/10455/src/net/dial.go:542
	_go_fuzz_dep_.CoverTab[5103]++

						var primary, fallback dialResult

//line /snap/go/10455/src/net/dial.go:547
	primaryCtx, primaryCancel := context.WithCancel(ctx)
						defer primaryCancel()
//line /snap/go/10455/src/net/dial.go:548
	_curRoutineNum5_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/dial.go:548
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum5_)
						go func() {
//line /snap/go/10455/src/net/dial.go:549
		_go_fuzz_dep_.CoverTab[5114]++
//line /snap/go/10455/src/net/dial.go:549
		defer func() {
//line /snap/go/10455/src/net/dial.go:549
			_go_fuzz_dep_.CoverTab[5115]++
//line /snap/go/10455/src/net/dial.go:549
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum5_)
//line /snap/go/10455/src/net/dial.go:549
			// _ = "end of CoverTab[5115]"
//line /snap/go/10455/src/net/dial.go:549
		}()
//line /snap/go/10455/src/net/dial.go:549
		startRacer(primaryCtx, true)
//line /snap/go/10455/src/net/dial.go:549
		// _ = "end of CoverTab[5114]"
//line /snap/go/10455/src/net/dial.go:549
	}()

//line /snap/go/10455/src/net/dial.go:552
	fallbackTimer := time.NewTimer(sd.fallbackDelay())
						defer fallbackTimer.Stop()
//line /snap/go/10455/src/net/dial.go:553
	_go_fuzz_dep_.CoverTab[786650] = 0

						for {
//line /snap/go/10455/src/net/dial.go:555
		if _go_fuzz_dep_.CoverTab[786650] == 0 {
//line /snap/go/10455/src/net/dial.go:555
			_go_fuzz_dep_.CoverTab[527823]++
//line /snap/go/10455/src/net/dial.go:555
		} else {
//line /snap/go/10455/src/net/dial.go:555
			_go_fuzz_dep_.CoverTab[527824]++
//line /snap/go/10455/src/net/dial.go:555
		}
//line /snap/go/10455/src/net/dial.go:555
		_go_fuzz_dep_.CoverTab[786650] = 1
//line /snap/go/10455/src/net/dial.go:555
		_go_fuzz_dep_.CoverTab[5116]++
							select {
		case <-fallbackTimer.C:
//line /snap/go/10455/src/net/dial.go:557
			_go_fuzz_dep_.CoverTab[5117]++
								fallbackCtx, fallbackCancel := context.WithCancel(ctx)
								defer fallbackCancel()
								go startRacer(fallbackCtx, false)
//line /snap/go/10455/src/net/dial.go:560
			// _ = "end of CoverTab[5117]"

		case res := <-results:
//line /snap/go/10455/src/net/dial.go:562
			_go_fuzz_dep_.CoverTab[5118]++
								if res.error == nil {
//line /snap/go/10455/src/net/dial.go:563
				_go_fuzz_dep_.CoverTab[527765]++
//line /snap/go/10455/src/net/dial.go:563
				_go_fuzz_dep_.CoverTab[5122]++
									return res.Conn, nil
//line /snap/go/10455/src/net/dial.go:564
				// _ = "end of CoverTab[5122]"
			} else {
//line /snap/go/10455/src/net/dial.go:565
				_go_fuzz_dep_.CoverTab[527766]++
//line /snap/go/10455/src/net/dial.go:565
				_go_fuzz_dep_.CoverTab[5123]++
//line /snap/go/10455/src/net/dial.go:565
				// _ = "end of CoverTab[5123]"
//line /snap/go/10455/src/net/dial.go:565
			}
//line /snap/go/10455/src/net/dial.go:565
			// _ = "end of CoverTab[5118]"
//line /snap/go/10455/src/net/dial.go:565
			_go_fuzz_dep_.CoverTab[5119]++
								if res.primary {
//line /snap/go/10455/src/net/dial.go:566
				_go_fuzz_dep_.CoverTab[527767]++
//line /snap/go/10455/src/net/dial.go:566
				_go_fuzz_dep_.CoverTab[5124]++
									primary = res
//line /snap/go/10455/src/net/dial.go:567
				// _ = "end of CoverTab[5124]"
			} else {
//line /snap/go/10455/src/net/dial.go:568
				_go_fuzz_dep_.CoverTab[527768]++
//line /snap/go/10455/src/net/dial.go:568
				_go_fuzz_dep_.CoverTab[5125]++
									fallback = res
//line /snap/go/10455/src/net/dial.go:569
				// _ = "end of CoverTab[5125]"
			}
//line /snap/go/10455/src/net/dial.go:570
			// _ = "end of CoverTab[5119]"
//line /snap/go/10455/src/net/dial.go:570
			_go_fuzz_dep_.CoverTab[5120]++
								if primary.done && func() bool {
//line /snap/go/10455/src/net/dial.go:571
				_go_fuzz_dep_.CoverTab[5126]++
//line /snap/go/10455/src/net/dial.go:571
				return fallback.done
//line /snap/go/10455/src/net/dial.go:571
				// _ = "end of CoverTab[5126]"
//line /snap/go/10455/src/net/dial.go:571
			}() {
//line /snap/go/10455/src/net/dial.go:571
				_go_fuzz_dep_.CoverTab[527769]++
//line /snap/go/10455/src/net/dial.go:571
				_go_fuzz_dep_.CoverTab[5127]++
									return nil, primary.error
//line /snap/go/10455/src/net/dial.go:572
				// _ = "end of CoverTab[5127]"
			} else {
//line /snap/go/10455/src/net/dial.go:573
				_go_fuzz_dep_.CoverTab[527770]++
//line /snap/go/10455/src/net/dial.go:573
				_go_fuzz_dep_.CoverTab[5128]++
//line /snap/go/10455/src/net/dial.go:573
				// _ = "end of CoverTab[5128]"
//line /snap/go/10455/src/net/dial.go:573
			}
//line /snap/go/10455/src/net/dial.go:573
			// _ = "end of CoverTab[5120]"
//line /snap/go/10455/src/net/dial.go:573
			_go_fuzz_dep_.CoverTab[5121]++
								if res.primary && func() bool {
//line /snap/go/10455/src/net/dial.go:574
				_go_fuzz_dep_.CoverTab[5129]++
//line /snap/go/10455/src/net/dial.go:574
				return fallbackTimer.Stop()
//line /snap/go/10455/src/net/dial.go:574
				// _ = "end of CoverTab[5129]"
//line /snap/go/10455/src/net/dial.go:574
			}() {
//line /snap/go/10455/src/net/dial.go:574
				_go_fuzz_dep_.CoverTab[527771]++
//line /snap/go/10455/src/net/dial.go:574
				_go_fuzz_dep_.CoverTab[5130]++

//line /snap/go/10455/src/net/dial.go:579
				fallbackTimer.Reset(0)
//line /snap/go/10455/src/net/dial.go:579
				// _ = "end of CoverTab[5130]"
			} else {
//line /snap/go/10455/src/net/dial.go:580
				_go_fuzz_dep_.CoverTab[527772]++
//line /snap/go/10455/src/net/dial.go:580
				_go_fuzz_dep_.CoverTab[5131]++
//line /snap/go/10455/src/net/dial.go:580
				// _ = "end of CoverTab[5131]"
//line /snap/go/10455/src/net/dial.go:580
			}
//line /snap/go/10455/src/net/dial.go:580
			// _ = "end of CoverTab[5121]"
		}
//line /snap/go/10455/src/net/dial.go:581
		// _ = "end of CoverTab[5116]"
	}
//line /snap/go/10455/src/net/dial.go:582
	// _ = "end of CoverTab[5103]"
}

// dialSerial connects to a list of addresses in sequence, returning
//line /snap/go/10455/src/net/dial.go:585
// either the first successful connection, or the first error.
//line /snap/go/10455/src/net/dial.go:587
func (sd *sysDialer) dialSerial(ctx context.Context, ras addrList) (Conn, error) {
//line /snap/go/10455/src/net/dial.go:587
	_go_fuzz_dep_.CoverTab[5132]++
						var firstErr error
//line /snap/go/10455/src/net/dial.go:588
	_go_fuzz_dep_. // The error from the first address is most relevant.
//line /snap/go/10455/src/net/dial.go:588
	CoverTab[786651] = 0

						for i, ra := range ras {
//line /snap/go/10455/src/net/dial.go:590
		if _go_fuzz_dep_.CoverTab[786651] == 0 {
//line /snap/go/10455/src/net/dial.go:590
			_go_fuzz_dep_.CoverTab[527827]++
//line /snap/go/10455/src/net/dial.go:590
		} else {
//line /snap/go/10455/src/net/dial.go:590
			_go_fuzz_dep_.CoverTab[527828]++
//line /snap/go/10455/src/net/dial.go:590
		}
//line /snap/go/10455/src/net/dial.go:590
		_go_fuzz_dep_.CoverTab[786651] = 1
//line /snap/go/10455/src/net/dial.go:590
		_go_fuzz_dep_.CoverTab[5135]++
							select {
		case <-ctx.Done():
//line /snap/go/10455/src/net/dial.go:592
			_go_fuzz_dep_.CoverTab[5139]++
								return nil, &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: mapErr(ctx.Err())}
//line /snap/go/10455/src/net/dial.go:593
			// _ = "end of CoverTab[5139]"
		default:
//line /snap/go/10455/src/net/dial.go:594
			_go_fuzz_dep_.CoverTab[5140]++
//line /snap/go/10455/src/net/dial.go:594
			// _ = "end of CoverTab[5140]"
		}
//line /snap/go/10455/src/net/dial.go:595
		// _ = "end of CoverTab[5135]"
//line /snap/go/10455/src/net/dial.go:595
		_go_fuzz_dep_.CoverTab[5136]++

							dialCtx := ctx
							if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /snap/go/10455/src/net/dial.go:598
			_go_fuzz_dep_.CoverTab[527773]++
//line /snap/go/10455/src/net/dial.go:598
			_go_fuzz_dep_.CoverTab[5141]++
								partialDeadline, err := partialDeadline(time.Now(), deadline, len(ras)-i)
								if err != nil {
//line /snap/go/10455/src/net/dial.go:600
				_go_fuzz_dep_.CoverTab[527775]++
//line /snap/go/10455/src/net/dial.go:600
				_go_fuzz_dep_.CoverTab[5143]++

									if firstErr == nil {
//line /snap/go/10455/src/net/dial.go:602
					_go_fuzz_dep_.CoverTab[527777]++
//line /snap/go/10455/src/net/dial.go:602
					_go_fuzz_dep_.CoverTab[5145]++
										firstErr = &OpError{Op: "dial", Net: sd.network, Source: sd.LocalAddr, Addr: ra, Err: err}
//line /snap/go/10455/src/net/dial.go:603
					// _ = "end of CoverTab[5145]"
				} else {
//line /snap/go/10455/src/net/dial.go:604
					_go_fuzz_dep_.CoverTab[527778]++
//line /snap/go/10455/src/net/dial.go:604
					_go_fuzz_dep_.CoverTab[5146]++
//line /snap/go/10455/src/net/dial.go:604
					// _ = "end of CoverTab[5146]"
//line /snap/go/10455/src/net/dial.go:604
				}
//line /snap/go/10455/src/net/dial.go:604
				// _ = "end of CoverTab[5143]"
//line /snap/go/10455/src/net/dial.go:604
				_go_fuzz_dep_.CoverTab[5144]++
									break
//line /snap/go/10455/src/net/dial.go:605
				// _ = "end of CoverTab[5144]"
			} else {
//line /snap/go/10455/src/net/dial.go:606
				_go_fuzz_dep_.CoverTab[527776]++
//line /snap/go/10455/src/net/dial.go:606
				_go_fuzz_dep_.CoverTab[5147]++
//line /snap/go/10455/src/net/dial.go:606
				// _ = "end of CoverTab[5147]"
//line /snap/go/10455/src/net/dial.go:606
			}
//line /snap/go/10455/src/net/dial.go:606
			// _ = "end of CoverTab[5141]"
//line /snap/go/10455/src/net/dial.go:606
			_go_fuzz_dep_.CoverTab[5142]++
								if partialDeadline.Before(deadline) {
//line /snap/go/10455/src/net/dial.go:607
				_go_fuzz_dep_.CoverTab[527779]++
//line /snap/go/10455/src/net/dial.go:607
				_go_fuzz_dep_.CoverTab[5148]++
									var cancel context.CancelFunc
									dialCtx, cancel = context.WithDeadline(ctx, partialDeadline)
									defer cancel()
//line /snap/go/10455/src/net/dial.go:610
				// _ = "end of CoverTab[5148]"
			} else {
//line /snap/go/10455/src/net/dial.go:611
				_go_fuzz_dep_.CoverTab[527780]++
//line /snap/go/10455/src/net/dial.go:611
				_go_fuzz_dep_.CoverTab[5149]++
//line /snap/go/10455/src/net/dial.go:611
				// _ = "end of CoverTab[5149]"
//line /snap/go/10455/src/net/dial.go:611
			}
//line /snap/go/10455/src/net/dial.go:611
			// _ = "end of CoverTab[5142]"
		} else {
//line /snap/go/10455/src/net/dial.go:612
			_go_fuzz_dep_.CoverTab[527774]++
//line /snap/go/10455/src/net/dial.go:612
			_go_fuzz_dep_.CoverTab[5150]++
//line /snap/go/10455/src/net/dial.go:612
			// _ = "end of CoverTab[5150]"
//line /snap/go/10455/src/net/dial.go:612
		}
//line /snap/go/10455/src/net/dial.go:612
		// _ = "end of CoverTab[5136]"
//line /snap/go/10455/src/net/dial.go:612
		_go_fuzz_dep_.CoverTab[5137]++

							c, err := sd.dialSingle(dialCtx, ra)
							if err == nil {
//line /snap/go/10455/src/net/dial.go:615
			_go_fuzz_dep_.CoverTab[527781]++
//line /snap/go/10455/src/net/dial.go:615
			_go_fuzz_dep_.CoverTab[5151]++
								return c, nil
//line /snap/go/10455/src/net/dial.go:616
			// _ = "end of CoverTab[5151]"
		} else {
//line /snap/go/10455/src/net/dial.go:617
			_go_fuzz_dep_.CoverTab[527782]++
//line /snap/go/10455/src/net/dial.go:617
			_go_fuzz_dep_.CoverTab[5152]++
//line /snap/go/10455/src/net/dial.go:617
			// _ = "end of CoverTab[5152]"
//line /snap/go/10455/src/net/dial.go:617
		}
//line /snap/go/10455/src/net/dial.go:617
		// _ = "end of CoverTab[5137]"
//line /snap/go/10455/src/net/dial.go:617
		_go_fuzz_dep_.CoverTab[5138]++
							if firstErr == nil {
//line /snap/go/10455/src/net/dial.go:618
			_go_fuzz_dep_.CoverTab[527783]++
//line /snap/go/10455/src/net/dial.go:618
			_go_fuzz_dep_.CoverTab[5153]++
								firstErr = err
//line /snap/go/10455/src/net/dial.go:619
			// _ = "end of CoverTab[5153]"
		} else {
//line /snap/go/10455/src/net/dial.go:620
			_go_fuzz_dep_.CoverTab[527784]++
//line /snap/go/10455/src/net/dial.go:620
			_go_fuzz_dep_.CoverTab[5154]++
//line /snap/go/10455/src/net/dial.go:620
			// _ = "end of CoverTab[5154]"
//line /snap/go/10455/src/net/dial.go:620
		}
//line /snap/go/10455/src/net/dial.go:620
		// _ = "end of CoverTab[5138]"
	}
//line /snap/go/10455/src/net/dial.go:621
	if _go_fuzz_dep_.CoverTab[786651] == 0 {
//line /snap/go/10455/src/net/dial.go:621
		_go_fuzz_dep_.CoverTab[527829]++
//line /snap/go/10455/src/net/dial.go:621
	} else {
//line /snap/go/10455/src/net/dial.go:621
		_go_fuzz_dep_.CoverTab[527830]++
//line /snap/go/10455/src/net/dial.go:621
	}
//line /snap/go/10455/src/net/dial.go:621
	// _ = "end of CoverTab[5132]"
//line /snap/go/10455/src/net/dial.go:621
	_go_fuzz_dep_.CoverTab[5133]++

						if firstErr == nil {
//line /snap/go/10455/src/net/dial.go:623
		_go_fuzz_dep_.CoverTab[527785]++
//line /snap/go/10455/src/net/dial.go:623
		_go_fuzz_dep_.CoverTab[5155]++
							firstErr = &OpError{Op: "dial", Net: sd.network, Source: nil, Addr: nil, Err: errMissingAddress}
//line /snap/go/10455/src/net/dial.go:624
		// _ = "end of CoverTab[5155]"
	} else {
//line /snap/go/10455/src/net/dial.go:625
		_go_fuzz_dep_.CoverTab[527786]++
//line /snap/go/10455/src/net/dial.go:625
		_go_fuzz_dep_.CoverTab[5156]++
//line /snap/go/10455/src/net/dial.go:625
		// _ = "end of CoverTab[5156]"
//line /snap/go/10455/src/net/dial.go:625
	}
//line /snap/go/10455/src/net/dial.go:625
	// _ = "end of CoverTab[5133]"
//line /snap/go/10455/src/net/dial.go:625
	_go_fuzz_dep_.CoverTab[5134]++
						return nil, firstErr
//line /snap/go/10455/src/net/dial.go:626
	// _ = "end of CoverTab[5134]"
}

// dialSingle attempts to establish and returns a single connection to
//line /snap/go/10455/src/net/dial.go:629
// the destination address.
//line /snap/go/10455/src/net/dial.go:631
func (sd *sysDialer) dialSingle(ctx context.Context, ra Addr) (c Conn, err error) {
//line /snap/go/10455/src/net/dial.go:631
	_go_fuzz_dep_.CoverTab[5157]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil {
//line /snap/go/10455/src/net/dial.go:633
		_go_fuzz_dep_.CoverTab[527787]++
//line /snap/go/10455/src/net/dial.go:633
		_go_fuzz_dep_.CoverTab[5161]++
							raStr := ra.String()
							if trace.ConnectStart != nil {
//line /snap/go/10455/src/net/dial.go:635
			_go_fuzz_dep_.CoverTab[527789]++
//line /snap/go/10455/src/net/dial.go:635
			_go_fuzz_dep_.CoverTab[5163]++
								trace.ConnectStart(sd.network, raStr)
//line /snap/go/10455/src/net/dial.go:636
			// _ = "end of CoverTab[5163]"
		} else {
//line /snap/go/10455/src/net/dial.go:637
			_go_fuzz_dep_.CoverTab[527790]++
//line /snap/go/10455/src/net/dial.go:637
			_go_fuzz_dep_.CoverTab[5164]++
//line /snap/go/10455/src/net/dial.go:637
			// _ = "end of CoverTab[5164]"
//line /snap/go/10455/src/net/dial.go:637
		}
//line /snap/go/10455/src/net/dial.go:637
		// _ = "end of CoverTab[5161]"
//line /snap/go/10455/src/net/dial.go:637
		_go_fuzz_dep_.CoverTab[5162]++
							if trace.ConnectDone != nil {
//line /snap/go/10455/src/net/dial.go:638
			_go_fuzz_dep_.CoverTab[527791]++
//line /snap/go/10455/src/net/dial.go:638
			_go_fuzz_dep_.CoverTab[5165]++
								defer func() {
//line /snap/go/10455/src/net/dial.go:639
				_go_fuzz_dep_.CoverTab[5166]++
//line /snap/go/10455/src/net/dial.go:639
				trace.ConnectDone(sd.network, raStr, err)
//line /snap/go/10455/src/net/dial.go:639
				// _ = "end of CoverTab[5166]"
//line /snap/go/10455/src/net/dial.go:639
			}()
//line /snap/go/10455/src/net/dial.go:639
			// _ = "end of CoverTab[5165]"
		} else {
//line /snap/go/10455/src/net/dial.go:640
			_go_fuzz_dep_.CoverTab[527792]++
//line /snap/go/10455/src/net/dial.go:640
			_go_fuzz_dep_.CoverTab[5167]++
//line /snap/go/10455/src/net/dial.go:640
			// _ = "end of CoverTab[5167]"
//line /snap/go/10455/src/net/dial.go:640
		}
//line /snap/go/10455/src/net/dial.go:640
		// _ = "end of CoverTab[5162]"
	} else {
//line /snap/go/10455/src/net/dial.go:641
		_go_fuzz_dep_.CoverTab[527788]++
//line /snap/go/10455/src/net/dial.go:641
		_go_fuzz_dep_.CoverTab[5168]++
//line /snap/go/10455/src/net/dial.go:641
		// _ = "end of CoverTab[5168]"
//line /snap/go/10455/src/net/dial.go:641
	}
//line /snap/go/10455/src/net/dial.go:641
	// _ = "end of CoverTab[5157]"
//line /snap/go/10455/src/net/dial.go:641
	_go_fuzz_dep_.CoverTab[5158]++
						la := sd.LocalAddr
						switch ra := ra.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/dial.go:644
		_go_fuzz_dep_.CoverTab[527793]++
//line /snap/go/10455/src/net/dial.go:644
		_go_fuzz_dep_.CoverTab[5169]++
							la, _ := la.(*TCPAddr)
							if sd.MultipathTCP() {
//line /snap/go/10455/src/net/dial.go:646
			_go_fuzz_dep_.CoverTab[527798]++
//line /snap/go/10455/src/net/dial.go:646
			_go_fuzz_dep_.CoverTab[5174]++
								c, err = sd.dialMPTCP(ctx, la, ra)
//line /snap/go/10455/src/net/dial.go:647
			// _ = "end of CoverTab[5174]"
		} else {
//line /snap/go/10455/src/net/dial.go:648
			_go_fuzz_dep_.CoverTab[527799]++
//line /snap/go/10455/src/net/dial.go:648
			_go_fuzz_dep_.CoverTab[5175]++
								c, err = sd.dialTCP(ctx, la, ra)
//line /snap/go/10455/src/net/dial.go:649
			// _ = "end of CoverTab[5175]"
		}
//line /snap/go/10455/src/net/dial.go:650
		// _ = "end of CoverTab[5169]"
	case *UDPAddr:
//line /snap/go/10455/src/net/dial.go:651
		_go_fuzz_dep_.CoverTab[527794]++
//line /snap/go/10455/src/net/dial.go:651
		_go_fuzz_dep_.CoverTab[5170]++
							la, _ := la.(*UDPAddr)
							c, err = sd.dialUDP(ctx, la, ra)
//line /snap/go/10455/src/net/dial.go:653
		// _ = "end of CoverTab[5170]"
	case *IPAddr:
//line /snap/go/10455/src/net/dial.go:654
		_go_fuzz_dep_.CoverTab[527795]++
//line /snap/go/10455/src/net/dial.go:654
		_go_fuzz_dep_.CoverTab[5171]++
							la, _ := la.(*IPAddr)
							c, err = sd.dialIP(ctx, la, ra)
//line /snap/go/10455/src/net/dial.go:656
		// _ = "end of CoverTab[5171]"
	case *UnixAddr:
//line /snap/go/10455/src/net/dial.go:657
		_go_fuzz_dep_.CoverTab[527796]++
//line /snap/go/10455/src/net/dial.go:657
		_go_fuzz_dep_.CoverTab[5172]++
							la, _ := la.(*UnixAddr)
							c, err = sd.dialUnix(ctx, la, ra)
//line /snap/go/10455/src/net/dial.go:659
		// _ = "end of CoverTab[5172]"
	default:
//line /snap/go/10455/src/net/dial.go:660
		_go_fuzz_dep_.CoverTab[527797]++
//line /snap/go/10455/src/net/dial.go:660
		_go_fuzz_dep_.CoverTab[5173]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: &AddrError{Err: "unexpected address type", Addr: sd.address}}
//line /snap/go/10455/src/net/dial.go:661
		// _ = "end of CoverTab[5173]"
	}
//line /snap/go/10455/src/net/dial.go:662
	// _ = "end of CoverTab[5158]"
//line /snap/go/10455/src/net/dial.go:662
	_go_fuzz_dep_.CoverTab[5159]++
						if err != nil {
//line /snap/go/10455/src/net/dial.go:663
		_go_fuzz_dep_.CoverTab[527800]++
//line /snap/go/10455/src/net/dial.go:663
		_go_fuzz_dep_.CoverTab[5176]++
							return nil, &OpError{Op: "dial", Net: sd.network, Source: la, Addr: ra, Err: err}
//line /snap/go/10455/src/net/dial.go:664
		// _ = "end of CoverTab[5176]"
	} else {
//line /snap/go/10455/src/net/dial.go:665
		_go_fuzz_dep_.CoverTab[527801]++
//line /snap/go/10455/src/net/dial.go:665
		_go_fuzz_dep_.CoverTab[5177]++
//line /snap/go/10455/src/net/dial.go:665
		// _ = "end of CoverTab[5177]"
//line /snap/go/10455/src/net/dial.go:665
	}
//line /snap/go/10455/src/net/dial.go:665
	// _ = "end of CoverTab[5159]"
//line /snap/go/10455/src/net/dial.go:665
	_go_fuzz_dep_.CoverTab[5160]++
						return c, nil
//line /snap/go/10455/src/net/dial.go:666
	// _ = "end of CoverTab[5160]"
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

	// If mptcpStatus is set to a value allowing Multipath TCP (MPTCP) to be
	// used, any call to Listen with "tcp(4|6)" as network will use MPTCP if
	// supported by the operating system.
	mptcpStatus	mptcpStatus
}

// MultipathTCP reports whether MPTCP will be used.
//line /snap/go/10455/src/net/dial.go:693
//
//line /snap/go/10455/src/net/dial.go:693
// This method doesn't check if MPTCP is supported by the operating
//line /snap/go/10455/src/net/dial.go:693
// system or not.
//line /snap/go/10455/src/net/dial.go:697
func (lc *ListenConfig) MultipathTCP() bool {
//line /snap/go/10455/src/net/dial.go:697
	_go_fuzz_dep_.CoverTab[5178]++
						return lc.mptcpStatus.get()
//line /snap/go/10455/src/net/dial.go:698
	// _ = "end of CoverTab[5178]"
}

// SetMultipathTCP directs the Listen method to use, or not use, MPTCP,
//line /snap/go/10455/src/net/dial.go:701
// if supported by the operating system. This method overrides the
//line /snap/go/10455/src/net/dial.go:701
// system default and the GODEBUG=multipathtcp=... setting if any.
//line /snap/go/10455/src/net/dial.go:701
//
//line /snap/go/10455/src/net/dial.go:701
// If MPTCP is not available on the host or not supported by the client,
//line /snap/go/10455/src/net/dial.go:701
// the Listen method will fall back to TCP.
//line /snap/go/10455/src/net/dial.go:707
func (lc *ListenConfig) SetMultipathTCP(use bool) {
//line /snap/go/10455/src/net/dial.go:707
	_go_fuzz_dep_.CoverTab[5179]++
						lc.mptcpStatus.set(use)
//line /snap/go/10455/src/net/dial.go:708
	// _ = "end of CoverTab[5179]"
}

// Listen announces on the local network address.
//line /snap/go/10455/src/net/dial.go:711
//
//line /snap/go/10455/src/net/dial.go:711
// See func Listen for a description of the network and address
//line /snap/go/10455/src/net/dial.go:711
// parameters.
//line /snap/go/10455/src/net/dial.go:715
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error) {
//line /snap/go/10455/src/net/dial.go:715
	_go_fuzz_dep_.CoverTab[5180]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /snap/go/10455/src/net/dial.go:717
		_go_fuzz_dep_.CoverTab[527802]++
//line /snap/go/10455/src/net/dial.go:717
		_go_fuzz_dep_.CoverTab[5184]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/dial.go:718
		// _ = "end of CoverTab[5184]"
	} else {
//line /snap/go/10455/src/net/dial.go:719
		_go_fuzz_dep_.CoverTab[527803]++
//line /snap/go/10455/src/net/dial.go:719
		_go_fuzz_dep_.CoverTab[5185]++
//line /snap/go/10455/src/net/dial.go:719
		// _ = "end of CoverTab[5185]"
//line /snap/go/10455/src/net/dial.go:719
	}
//line /snap/go/10455/src/net/dial.go:719
	// _ = "end of CoverTab[5180]"
//line /snap/go/10455/src/net/dial.go:719
	_go_fuzz_dep_.CoverTab[5181]++
						sl := &sysListener{
		ListenConfig:	*lc,
		network:	network,
		address:	address,
	}
	var l Listener
	la := addrs.first(isIPv4)
	switch la := la.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/dial.go:728
		_go_fuzz_dep_.CoverTab[527804]++
//line /snap/go/10455/src/net/dial.go:728
		_go_fuzz_dep_.CoverTab[5186]++
							if sl.MultipathTCP() {
//line /snap/go/10455/src/net/dial.go:729
			_go_fuzz_dep_.CoverTab[527807]++
//line /snap/go/10455/src/net/dial.go:729
			_go_fuzz_dep_.CoverTab[5189]++
								l, err = sl.listenMPTCP(ctx, la)
//line /snap/go/10455/src/net/dial.go:730
			// _ = "end of CoverTab[5189]"
		} else {
//line /snap/go/10455/src/net/dial.go:731
			_go_fuzz_dep_.CoverTab[527808]++
//line /snap/go/10455/src/net/dial.go:731
			_go_fuzz_dep_.CoverTab[5190]++
								l, err = sl.listenTCP(ctx, la)
//line /snap/go/10455/src/net/dial.go:732
			// _ = "end of CoverTab[5190]"
		}
//line /snap/go/10455/src/net/dial.go:733
		// _ = "end of CoverTab[5186]"
	case *UnixAddr:
//line /snap/go/10455/src/net/dial.go:734
		_go_fuzz_dep_.CoverTab[527805]++
//line /snap/go/10455/src/net/dial.go:734
		_go_fuzz_dep_.CoverTab[5187]++
							l, err = sl.listenUnix(ctx, la)
//line /snap/go/10455/src/net/dial.go:735
		// _ = "end of CoverTab[5187]"
	default:
//line /snap/go/10455/src/net/dial.go:736
		_go_fuzz_dep_.CoverTab[527806]++
//line /snap/go/10455/src/net/dial.go:736
		_go_fuzz_dep_.CoverTab[5188]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /snap/go/10455/src/net/dial.go:737
		// _ = "end of CoverTab[5188]"
	}
//line /snap/go/10455/src/net/dial.go:738
	// _ = "end of CoverTab[5181]"
//line /snap/go/10455/src/net/dial.go:738
	_go_fuzz_dep_.CoverTab[5182]++
						if err != nil {
//line /snap/go/10455/src/net/dial.go:739
		_go_fuzz_dep_.CoverTab[527809]++
//line /snap/go/10455/src/net/dial.go:739
		_go_fuzz_dep_.CoverTab[5191]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /snap/go/10455/src/net/dial.go:740
		// _ = "end of CoverTab[5191]"
	} else {
//line /snap/go/10455/src/net/dial.go:741
		_go_fuzz_dep_.CoverTab[527810]++
//line /snap/go/10455/src/net/dial.go:741
		_go_fuzz_dep_.CoverTab[5192]++
//line /snap/go/10455/src/net/dial.go:741
		// _ = "end of CoverTab[5192]"
//line /snap/go/10455/src/net/dial.go:741
	}
//line /snap/go/10455/src/net/dial.go:741
	// _ = "end of CoverTab[5182]"
//line /snap/go/10455/src/net/dial.go:741
	_go_fuzz_dep_.CoverTab[5183]++
						return l, nil
//line /snap/go/10455/src/net/dial.go:742
	// _ = "end of CoverTab[5183]"
}

// ListenPacket announces on the local network address.
//line /snap/go/10455/src/net/dial.go:745
//
//line /snap/go/10455/src/net/dial.go:745
// See func ListenPacket for a description of the network and address
//line /snap/go/10455/src/net/dial.go:745
// parameters.
//line /snap/go/10455/src/net/dial.go:749
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error) {
//line /snap/go/10455/src/net/dial.go:749
	_go_fuzz_dep_.CoverTab[5193]++
						addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
						if err != nil {
//line /snap/go/10455/src/net/dial.go:751
		_go_fuzz_dep_.CoverTab[527811]++
//line /snap/go/10455/src/net/dial.go:751
		_go_fuzz_dep_.CoverTab[5197]++
							return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/dial.go:752
		// _ = "end of CoverTab[5197]"
	} else {
//line /snap/go/10455/src/net/dial.go:753
		_go_fuzz_dep_.CoverTab[527812]++
//line /snap/go/10455/src/net/dial.go:753
		_go_fuzz_dep_.CoverTab[5198]++
//line /snap/go/10455/src/net/dial.go:753
		// _ = "end of CoverTab[5198]"
//line /snap/go/10455/src/net/dial.go:753
	}
//line /snap/go/10455/src/net/dial.go:753
	// _ = "end of CoverTab[5193]"
//line /snap/go/10455/src/net/dial.go:753
	_go_fuzz_dep_.CoverTab[5194]++
						sl := &sysListener{
		ListenConfig:	*lc,
		network:	network,
		address:	address,
	}
	var c PacketConn
	la := addrs.first(isIPv4)
	switch la := la.(type) {
	case *UDPAddr:
//line /snap/go/10455/src/net/dial.go:762
		_go_fuzz_dep_.CoverTab[527813]++
//line /snap/go/10455/src/net/dial.go:762
		_go_fuzz_dep_.CoverTab[5199]++
							c, err = sl.listenUDP(ctx, la)
//line /snap/go/10455/src/net/dial.go:763
		// _ = "end of CoverTab[5199]"
	case *IPAddr:
//line /snap/go/10455/src/net/dial.go:764
		_go_fuzz_dep_.CoverTab[527814]++
//line /snap/go/10455/src/net/dial.go:764
		_go_fuzz_dep_.CoverTab[5200]++
							c, err = sl.listenIP(ctx, la)
//line /snap/go/10455/src/net/dial.go:765
		// _ = "end of CoverTab[5200]"
	case *UnixAddr:
//line /snap/go/10455/src/net/dial.go:766
		_go_fuzz_dep_.CoverTab[527815]++
//line /snap/go/10455/src/net/dial.go:766
		_go_fuzz_dep_.CoverTab[5201]++
							c, err = sl.listenUnixgram(ctx, la)
//line /snap/go/10455/src/net/dial.go:767
		// _ = "end of CoverTab[5201]"
	default:
//line /snap/go/10455/src/net/dial.go:768
		_go_fuzz_dep_.CoverTab[527816]++
//line /snap/go/10455/src/net/dial.go:768
		_go_fuzz_dep_.CoverTab[5202]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
//line /snap/go/10455/src/net/dial.go:769
		// _ = "end of CoverTab[5202]"
	}
//line /snap/go/10455/src/net/dial.go:770
	// _ = "end of CoverTab[5194]"
//line /snap/go/10455/src/net/dial.go:770
	_go_fuzz_dep_.CoverTab[5195]++
						if err != nil {
//line /snap/go/10455/src/net/dial.go:771
		_go_fuzz_dep_.CoverTab[527817]++
//line /snap/go/10455/src/net/dial.go:771
		_go_fuzz_dep_.CoverTab[5203]++
							return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err}
//line /snap/go/10455/src/net/dial.go:772
		// _ = "end of CoverTab[5203]"
	} else {
//line /snap/go/10455/src/net/dial.go:773
		_go_fuzz_dep_.CoverTab[527818]++
//line /snap/go/10455/src/net/dial.go:773
		_go_fuzz_dep_.CoverTab[5204]++
//line /snap/go/10455/src/net/dial.go:773
		// _ = "end of CoverTab[5204]"
//line /snap/go/10455/src/net/dial.go:773
	}
//line /snap/go/10455/src/net/dial.go:773
	// _ = "end of CoverTab[5195]"
//line /snap/go/10455/src/net/dial.go:773
	_go_fuzz_dep_.CoverTab[5196]++
						return c, nil
//line /snap/go/10455/src/net/dial.go:774
	// _ = "end of CoverTab[5196]"
}

// sysListener contains a Listen's parameters and configuration.
type sysListener struct {
	ListenConfig
	network, address	string
}

// Listen announces on the local network address.
//line /snap/go/10455/src/net/dial.go:783
//
//line /snap/go/10455/src/net/dial.go:783
// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
//line /snap/go/10455/src/net/dial.go:783
//
//line /snap/go/10455/src/net/dial.go:783
// For TCP networks, if the host in the address parameter is empty or
//line /snap/go/10455/src/net/dial.go:783
// a literal unspecified IP address, Listen listens on all available
//line /snap/go/10455/src/net/dial.go:783
// unicast and anycast IP addresses of the local system.
//line /snap/go/10455/src/net/dial.go:783
// To only use IPv4, use network "tcp4".
//line /snap/go/10455/src/net/dial.go:783
// The address can use a host name, but this is not recommended,
//line /snap/go/10455/src/net/dial.go:783
// because it will create a listener for at most one of the host's IP
//line /snap/go/10455/src/net/dial.go:783
// addresses.
//line /snap/go/10455/src/net/dial.go:783
// If the port in the address parameter is empty or "0", as in
//line /snap/go/10455/src/net/dial.go:783
// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
//line /snap/go/10455/src/net/dial.go:783
// The Addr method of Listener can be used to discover the chosen
//line /snap/go/10455/src/net/dial.go:783
// port.
//line /snap/go/10455/src/net/dial.go:783
//
//line /snap/go/10455/src/net/dial.go:783
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/dial.go:783
// parameters.
//line /snap/go/10455/src/net/dial.go:783
//
//line /snap/go/10455/src/net/dial.go:783
// Listen uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/dial.go:783
// ListenConfig.Listen.
//line /snap/go/10455/src/net/dial.go:804
func Listen(network, address string) (Listener, error) {
//line /snap/go/10455/src/net/dial.go:804
	_go_fuzz_dep_.CoverTab[5205]++
						var lc ListenConfig
						return lc.Listen(context.Background(), network, address)
//line /snap/go/10455/src/net/dial.go:806
	// _ = "end of CoverTab[5205]"
}

// ListenPacket announces on the local network address.
//line /snap/go/10455/src/net/dial.go:809
//
//line /snap/go/10455/src/net/dial.go:809
// The network must be "udp", "udp4", "udp6", "unixgram", or an IP
//line /snap/go/10455/src/net/dial.go:809
// transport. The IP transports are "ip", "ip4", or "ip6" followed by
//line /snap/go/10455/src/net/dial.go:809
// a colon and a literal protocol number or a protocol name, as in
//line /snap/go/10455/src/net/dial.go:809
// "ip:1" or "ip:icmp".
//line /snap/go/10455/src/net/dial.go:809
//
//line /snap/go/10455/src/net/dial.go:809
// For UDP and IP networks, if the host in the address parameter is
//line /snap/go/10455/src/net/dial.go:809
// empty or a literal unspecified IP address, ListenPacket listens on
//line /snap/go/10455/src/net/dial.go:809
// all available IP addresses of the local system except multicast IP
//line /snap/go/10455/src/net/dial.go:809
// addresses.
//line /snap/go/10455/src/net/dial.go:809
// To only use IPv4, use network "udp4" or "ip4:proto".
//line /snap/go/10455/src/net/dial.go:809
// The address can use a host name, but this is not recommended,
//line /snap/go/10455/src/net/dial.go:809
// because it will create a listener for at most one of the host's IP
//line /snap/go/10455/src/net/dial.go:809
// addresses.
//line /snap/go/10455/src/net/dial.go:809
// If the port in the address parameter is empty or "0", as in
//line /snap/go/10455/src/net/dial.go:809
// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
//line /snap/go/10455/src/net/dial.go:809
// The LocalAddr method of PacketConn can be used to discover the
//line /snap/go/10455/src/net/dial.go:809
// chosen port.
//line /snap/go/10455/src/net/dial.go:809
//
//line /snap/go/10455/src/net/dial.go:809
// See func Dial for a description of the network and address
//line /snap/go/10455/src/net/dial.go:809
// parameters.
//line /snap/go/10455/src/net/dial.go:809
//
//line /snap/go/10455/src/net/dial.go:809
// ListenPacket uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/dial.go:809
// ListenConfig.ListenPacket.
//line /snap/go/10455/src/net/dial.go:834
func ListenPacket(network, address string) (PacketConn, error) {
//line /snap/go/10455/src/net/dial.go:834
	_go_fuzz_dep_.CoverTab[5206]++
						var lc ListenConfig
						return lc.ListenPacket(context.Background(), network, address)
//line /snap/go/10455/src/net/dial.go:836
	// _ = "end of CoverTab[5206]"
}

//line /snap/go/10455/src/net/dial.go:837
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/dial.go:837
var _ = _go_fuzz_dep_.CoverTab
