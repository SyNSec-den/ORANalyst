// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/port.go:5
package net

//line /usr/local/go/src/net/port.go:5
import (
//line /usr/local/go/src/net/port.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/port.go:5
)
//line /usr/local/go/src/net/port.go:5
import (
//line /usr/local/go/src/net/port.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/port.go:5
)

// parsePort parses service as a decimal integer and returns the
//line /usr/local/go/src/net/port.go:7
// corresponding value as port. It is the caller's responsibility to
//line /usr/local/go/src/net/port.go:7
// parse service as a non-decimal integer when needsLookup is true.
//line /usr/local/go/src/net/port.go:7
//
//line /usr/local/go/src/net/port.go:7
// Some system resolvers will return a valid port number when given a number
//line /usr/local/go/src/net/port.go:7
// over 65536 (see https://golang.org/issues/11715). Alas, the parser
//line /usr/local/go/src/net/port.go:7
// can't bail early on numbers > 65536. Therefore reasonably large/small
//line /usr/local/go/src/net/port.go:7
// numbers are parsed in full and rejected if invalid.
//line /usr/local/go/src/net/port.go:15
func parsePort(service string) (port int, needsLookup bool) {
//line /usr/local/go/src/net/port.go:15
	_go_fuzz_dep_.CoverTab[16008]++
						if service == "" {
//line /usr/local/go/src/net/port.go:16
		_go_fuzz_dep_.CoverTab[16014]++

//line /usr/local/go/src/net/port.go:19
		return 0, false
//line /usr/local/go/src/net/port.go:19
		// _ = "end of CoverTab[16014]"
	} else {
//line /usr/local/go/src/net/port.go:20
		_go_fuzz_dep_.CoverTab[16015]++
//line /usr/local/go/src/net/port.go:20
		// _ = "end of CoverTab[16015]"
//line /usr/local/go/src/net/port.go:20
	}
//line /usr/local/go/src/net/port.go:20
	// _ = "end of CoverTab[16008]"
//line /usr/local/go/src/net/port.go:20
	_go_fuzz_dep_.CoverTab[16009]++
						const (
		max	= uint32(1<<32 - 1)
		cutoff	= uint32(1 << 30)
	)
	neg := false
	if service[0] == '+' {
//line /usr/local/go/src/net/port.go:26
		_go_fuzz_dep_.CoverTab[16016]++
							service = service[1:]
//line /usr/local/go/src/net/port.go:27
		// _ = "end of CoverTab[16016]"
	} else {
//line /usr/local/go/src/net/port.go:28
		_go_fuzz_dep_.CoverTab[16017]++
//line /usr/local/go/src/net/port.go:28
		if service[0] == '-' {
//line /usr/local/go/src/net/port.go:28
			_go_fuzz_dep_.CoverTab[16018]++
								neg = true
								service = service[1:]
//line /usr/local/go/src/net/port.go:30
			// _ = "end of CoverTab[16018]"
		} else {
//line /usr/local/go/src/net/port.go:31
			_go_fuzz_dep_.CoverTab[16019]++
//line /usr/local/go/src/net/port.go:31
			// _ = "end of CoverTab[16019]"
//line /usr/local/go/src/net/port.go:31
		}
//line /usr/local/go/src/net/port.go:31
		// _ = "end of CoverTab[16017]"
//line /usr/local/go/src/net/port.go:31
	}
//line /usr/local/go/src/net/port.go:31
	// _ = "end of CoverTab[16009]"
//line /usr/local/go/src/net/port.go:31
	_go_fuzz_dep_.CoverTab[16010]++
						var n uint32
						for _, d := range service {
//line /usr/local/go/src/net/port.go:33
		_go_fuzz_dep_.CoverTab[16020]++
							if '0' <= d && func() bool {
//line /usr/local/go/src/net/port.go:34
			_go_fuzz_dep_.CoverTab[16024]++
//line /usr/local/go/src/net/port.go:34
			return d <= '9'
//line /usr/local/go/src/net/port.go:34
			// _ = "end of CoverTab[16024]"
//line /usr/local/go/src/net/port.go:34
		}() {
//line /usr/local/go/src/net/port.go:34
			_go_fuzz_dep_.CoverTab[16025]++
								d -= '0'
//line /usr/local/go/src/net/port.go:35
			// _ = "end of CoverTab[16025]"
		} else {
//line /usr/local/go/src/net/port.go:36
			_go_fuzz_dep_.CoverTab[16026]++
								return 0, true
//line /usr/local/go/src/net/port.go:37
			// _ = "end of CoverTab[16026]"
		}
//line /usr/local/go/src/net/port.go:38
		// _ = "end of CoverTab[16020]"
//line /usr/local/go/src/net/port.go:38
		_go_fuzz_dep_.CoverTab[16021]++
							if n >= cutoff {
//line /usr/local/go/src/net/port.go:39
			_go_fuzz_dep_.CoverTab[16027]++
								n = max
								break
//line /usr/local/go/src/net/port.go:41
			// _ = "end of CoverTab[16027]"
		} else {
//line /usr/local/go/src/net/port.go:42
			_go_fuzz_dep_.CoverTab[16028]++
//line /usr/local/go/src/net/port.go:42
			// _ = "end of CoverTab[16028]"
//line /usr/local/go/src/net/port.go:42
		}
//line /usr/local/go/src/net/port.go:42
		// _ = "end of CoverTab[16021]"
//line /usr/local/go/src/net/port.go:42
		_go_fuzz_dep_.CoverTab[16022]++
							n *= 10
							nn := n + uint32(d)
							if nn < n || func() bool {
//line /usr/local/go/src/net/port.go:45
			_go_fuzz_dep_.CoverTab[16029]++
//line /usr/local/go/src/net/port.go:45
			return nn > max
//line /usr/local/go/src/net/port.go:45
			// _ = "end of CoverTab[16029]"
//line /usr/local/go/src/net/port.go:45
		}() {
//line /usr/local/go/src/net/port.go:45
			_go_fuzz_dep_.CoverTab[16030]++
								n = max
								break
//line /usr/local/go/src/net/port.go:47
			// _ = "end of CoverTab[16030]"
		} else {
//line /usr/local/go/src/net/port.go:48
			_go_fuzz_dep_.CoverTab[16031]++
//line /usr/local/go/src/net/port.go:48
			// _ = "end of CoverTab[16031]"
//line /usr/local/go/src/net/port.go:48
		}
//line /usr/local/go/src/net/port.go:48
		// _ = "end of CoverTab[16022]"
//line /usr/local/go/src/net/port.go:48
		_go_fuzz_dep_.CoverTab[16023]++
							n = nn
//line /usr/local/go/src/net/port.go:49
		// _ = "end of CoverTab[16023]"
	}
//line /usr/local/go/src/net/port.go:50
	// _ = "end of CoverTab[16010]"
//line /usr/local/go/src/net/port.go:50
	_go_fuzz_dep_.CoverTab[16011]++
						if !neg && func() bool {
//line /usr/local/go/src/net/port.go:51
		_go_fuzz_dep_.CoverTab[16032]++
//line /usr/local/go/src/net/port.go:51
		return n >= cutoff
//line /usr/local/go/src/net/port.go:51
		// _ = "end of CoverTab[16032]"
//line /usr/local/go/src/net/port.go:51
	}() {
//line /usr/local/go/src/net/port.go:51
		_go_fuzz_dep_.CoverTab[16033]++
							port = int(cutoff - 1)
//line /usr/local/go/src/net/port.go:52
		// _ = "end of CoverTab[16033]"
	} else {
//line /usr/local/go/src/net/port.go:53
		_go_fuzz_dep_.CoverTab[16034]++
//line /usr/local/go/src/net/port.go:53
		if neg && func() bool {
//line /usr/local/go/src/net/port.go:53
			_go_fuzz_dep_.CoverTab[16035]++
//line /usr/local/go/src/net/port.go:53
			return n > cutoff
//line /usr/local/go/src/net/port.go:53
			// _ = "end of CoverTab[16035]"
//line /usr/local/go/src/net/port.go:53
		}() {
//line /usr/local/go/src/net/port.go:53
			_go_fuzz_dep_.CoverTab[16036]++
								port = int(cutoff)
//line /usr/local/go/src/net/port.go:54
			// _ = "end of CoverTab[16036]"
		} else {
//line /usr/local/go/src/net/port.go:55
			_go_fuzz_dep_.CoverTab[16037]++
								port = int(n)
//line /usr/local/go/src/net/port.go:56
			// _ = "end of CoverTab[16037]"
		}
//line /usr/local/go/src/net/port.go:57
		// _ = "end of CoverTab[16034]"
//line /usr/local/go/src/net/port.go:57
	}
//line /usr/local/go/src/net/port.go:57
	// _ = "end of CoverTab[16011]"
//line /usr/local/go/src/net/port.go:57
	_go_fuzz_dep_.CoverTab[16012]++
						if neg {
//line /usr/local/go/src/net/port.go:58
		_go_fuzz_dep_.CoverTab[16038]++
							port = -port
//line /usr/local/go/src/net/port.go:59
		// _ = "end of CoverTab[16038]"
	} else {
//line /usr/local/go/src/net/port.go:60
		_go_fuzz_dep_.CoverTab[16039]++
//line /usr/local/go/src/net/port.go:60
		// _ = "end of CoverTab[16039]"
//line /usr/local/go/src/net/port.go:60
	}
//line /usr/local/go/src/net/port.go:60
	// _ = "end of CoverTab[16012]"
//line /usr/local/go/src/net/port.go:60
	_go_fuzz_dep_.CoverTab[16013]++
						return port, false
//line /usr/local/go/src/net/port.go:61
	// _ = "end of CoverTab[16013]"
}

//line /usr/local/go/src/net/port.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/port.go:62
var _ = _go_fuzz_dep_.CoverTab
