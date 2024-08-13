// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/port.go:5
package net

//line /snap/go/10455/src/net/port.go:5
import (
//line /snap/go/10455/src/net/port.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/port.go:5
)
//line /snap/go/10455/src/net/port.go:5
import (
//line /snap/go/10455/src/net/port.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/port.go:5
)

// parsePort parses service as a decimal integer and returns the
//line /snap/go/10455/src/net/port.go:7
// corresponding value as port. It is the caller's responsibility to
//line /snap/go/10455/src/net/port.go:7
// parse service as a non-decimal integer when needsLookup is true.
//line /snap/go/10455/src/net/port.go:7
//
//line /snap/go/10455/src/net/port.go:7
// Some system resolvers will return a valid port number when given a number
//line /snap/go/10455/src/net/port.go:7
// over 65536 (see https://golang.org/issues/11715). Alas, the parser
//line /snap/go/10455/src/net/port.go:7
// can't bail early on numbers > 65536. Therefore reasonably large/small
//line /snap/go/10455/src/net/port.go:7
// numbers are parsed in full and rejected if invalid.
//line /snap/go/10455/src/net/port.go:15
func parsePort(service string) (port int, needsLookup bool) {
//line /snap/go/10455/src/net/port.go:15
	_go_fuzz_dep_.CoverTab[7898]++
						if service == "" {
//line /snap/go/10455/src/net/port.go:16
		_go_fuzz_dep_.CoverTab[529540]++
//line /snap/go/10455/src/net/port.go:16
		_go_fuzz_dep_.CoverTab[7904]++

//line /snap/go/10455/src/net/port.go:19
		return 0, false
//line /snap/go/10455/src/net/port.go:19
		// _ = "end of CoverTab[7904]"
	} else {
//line /snap/go/10455/src/net/port.go:20
		_go_fuzz_dep_.CoverTab[529541]++
//line /snap/go/10455/src/net/port.go:20
		_go_fuzz_dep_.CoverTab[7905]++
//line /snap/go/10455/src/net/port.go:20
		// _ = "end of CoverTab[7905]"
//line /snap/go/10455/src/net/port.go:20
	}
//line /snap/go/10455/src/net/port.go:20
	// _ = "end of CoverTab[7898]"
//line /snap/go/10455/src/net/port.go:20
	_go_fuzz_dep_.CoverTab[7899]++
						const (
		max	= uint32(1<<32 - 1)
		cutoff	= uint32(1 << 30)
	)
	neg := false
	if service[0] == '+' {
//line /snap/go/10455/src/net/port.go:26
		_go_fuzz_dep_.CoverTab[529542]++
//line /snap/go/10455/src/net/port.go:26
		_go_fuzz_dep_.CoverTab[7906]++
							service = service[1:]
//line /snap/go/10455/src/net/port.go:27
		// _ = "end of CoverTab[7906]"
	} else {
//line /snap/go/10455/src/net/port.go:28
		_go_fuzz_dep_.CoverTab[529543]++
//line /snap/go/10455/src/net/port.go:28
		_go_fuzz_dep_.CoverTab[7907]++
//line /snap/go/10455/src/net/port.go:28
		if service[0] == '-' {
//line /snap/go/10455/src/net/port.go:28
			_go_fuzz_dep_.CoverTab[529544]++
//line /snap/go/10455/src/net/port.go:28
			_go_fuzz_dep_.CoverTab[7908]++
								neg = true
								service = service[1:]
//line /snap/go/10455/src/net/port.go:30
			// _ = "end of CoverTab[7908]"
		} else {
//line /snap/go/10455/src/net/port.go:31
			_go_fuzz_dep_.CoverTab[529545]++
//line /snap/go/10455/src/net/port.go:31
			_go_fuzz_dep_.CoverTab[7909]++
//line /snap/go/10455/src/net/port.go:31
			// _ = "end of CoverTab[7909]"
//line /snap/go/10455/src/net/port.go:31
		}
//line /snap/go/10455/src/net/port.go:31
		// _ = "end of CoverTab[7907]"
//line /snap/go/10455/src/net/port.go:31
	}
//line /snap/go/10455/src/net/port.go:31
	// _ = "end of CoverTab[7899]"
//line /snap/go/10455/src/net/port.go:31
	_go_fuzz_dep_.CoverTab[7900]++
						var n uint32
//line /snap/go/10455/src/net/port.go:32
	_go_fuzz_dep_.CoverTab[786741] = 0
						for _, d := range service {
//line /snap/go/10455/src/net/port.go:33
		if _go_fuzz_dep_.CoverTab[786741] == 0 {
//line /snap/go/10455/src/net/port.go:33
			_go_fuzz_dep_.CoverTab[529558]++
//line /snap/go/10455/src/net/port.go:33
		} else {
//line /snap/go/10455/src/net/port.go:33
			_go_fuzz_dep_.CoverTab[529559]++
//line /snap/go/10455/src/net/port.go:33
		}
//line /snap/go/10455/src/net/port.go:33
		_go_fuzz_dep_.CoverTab[786741] = 1
//line /snap/go/10455/src/net/port.go:33
		_go_fuzz_dep_.CoverTab[7910]++
							if '0' <= d && func() bool {
//line /snap/go/10455/src/net/port.go:34
			_go_fuzz_dep_.CoverTab[7914]++
//line /snap/go/10455/src/net/port.go:34
			return d <= '9'
//line /snap/go/10455/src/net/port.go:34
			// _ = "end of CoverTab[7914]"
//line /snap/go/10455/src/net/port.go:34
		}() {
//line /snap/go/10455/src/net/port.go:34
			_go_fuzz_dep_.CoverTab[529546]++
//line /snap/go/10455/src/net/port.go:34
			_go_fuzz_dep_.CoverTab[7915]++
								d -= '0'
//line /snap/go/10455/src/net/port.go:35
			// _ = "end of CoverTab[7915]"
		} else {
//line /snap/go/10455/src/net/port.go:36
			_go_fuzz_dep_.CoverTab[529547]++
//line /snap/go/10455/src/net/port.go:36
			_go_fuzz_dep_.CoverTab[7916]++
								return 0, true
//line /snap/go/10455/src/net/port.go:37
			// _ = "end of CoverTab[7916]"
		}
//line /snap/go/10455/src/net/port.go:38
		// _ = "end of CoverTab[7910]"
//line /snap/go/10455/src/net/port.go:38
		_go_fuzz_dep_.CoverTab[7911]++
							if n >= cutoff {
//line /snap/go/10455/src/net/port.go:39
			_go_fuzz_dep_.CoverTab[529548]++
//line /snap/go/10455/src/net/port.go:39
			_go_fuzz_dep_.CoverTab[7917]++
								n = max
								break
//line /snap/go/10455/src/net/port.go:41
			// _ = "end of CoverTab[7917]"
		} else {
//line /snap/go/10455/src/net/port.go:42
			_go_fuzz_dep_.CoverTab[529549]++
//line /snap/go/10455/src/net/port.go:42
			_go_fuzz_dep_.CoverTab[7918]++
//line /snap/go/10455/src/net/port.go:42
			// _ = "end of CoverTab[7918]"
//line /snap/go/10455/src/net/port.go:42
		}
//line /snap/go/10455/src/net/port.go:42
		// _ = "end of CoverTab[7911]"
//line /snap/go/10455/src/net/port.go:42
		_go_fuzz_dep_.CoverTab[7912]++
							n *= 10
							nn := n + uint32(d)
							if nn < n || func() bool {
//line /snap/go/10455/src/net/port.go:45
			_go_fuzz_dep_.CoverTab[7919]++
//line /snap/go/10455/src/net/port.go:45
			return nn > max
//line /snap/go/10455/src/net/port.go:45
			// _ = "end of CoverTab[7919]"
//line /snap/go/10455/src/net/port.go:45
		}() {
//line /snap/go/10455/src/net/port.go:45
			_go_fuzz_dep_.CoverTab[529550]++
//line /snap/go/10455/src/net/port.go:45
			_go_fuzz_dep_.CoverTab[7920]++
								n = max
								break
//line /snap/go/10455/src/net/port.go:47
			// _ = "end of CoverTab[7920]"
		} else {
//line /snap/go/10455/src/net/port.go:48
			_go_fuzz_dep_.CoverTab[529551]++
//line /snap/go/10455/src/net/port.go:48
			_go_fuzz_dep_.CoverTab[7921]++
//line /snap/go/10455/src/net/port.go:48
			// _ = "end of CoverTab[7921]"
//line /snap/go/10455/src/net/port.go:48
		}
//line /snap/go/10455/src/net/port.go:48
		// _ = "end of CoverTab[7912]"
//line /snap/go/10455/src/net/port.go:48
		_go_fuzz_dep_.CoverTab[7913]++
							n = nn
//line /snap/go/10455/src/net/port.go:49
		// _ = "end of CoverTab[7913]"
	}
//line /snap/go/10455/src/net/port.go:50
	if _go_fuzz_dep_.CoverTab[786741] == 0 {
//line /snap/go/10455/src/net/port.go:50
		_go_fuzz_dep_.CoverTab[529560]++
//line /snap/go/10455/src/net/port.go:50
	} else {
//line /snap/go/10455/src/net/port.go:50
		_go_fuzz_dep_.CoverTab[529561]++
//line /snap/go/10455/src/net/port.go:50
	}
//line /snap/go/10455/src/net/port.go:50
	// _ = "end of CoverTab[7900]"
//line /snap/go/10455/src/net/port.go:50
	_go_fuzz_dep_.CoverTab[7901]++
						if !neg && func() bool {
//line /snap/go/10455/src/net/port.go:51
		_go_fuzz_dep_.CoverTab[7922]++
//line /snap/go/10455/src/net/port.go:51
		return n >= cutoff
//line /snap/go/10455/src/net/port.go:51
		// _ = "end of CoverTab[7922]"
//line /snap/go/10455/src/net/port.go:51
	}() {
//line /snap/go/10455/src/net/port.go:51
		_go_fuzz_dep_.CoverTab[529552]++
//line /snap/go/10455/src/net/port.go:51
		_go_fuzz_dep_.CoverTab[7923]++
							port = int(cutoff - 1)
//line /snap/go/10455/src/net/port.go:52
		// _ = "end of CoverTab[7923]"
	} else {
//line /snap/go/10455/src/net/port.go:53
		_go_fuzz_dep_.CoverTab[529553]++
//line /snap/go/10455/src/net/port.go:53
		_go_fuzz_dep_.CoverTab[7924]++
//line /snap/go/10455/src/net/port.go:53
		if neg && func() bool {
//line /snap/go/10455/src/net/port.go:53
			_go_fuzz_dep_.CoverTab[7925]++
//line /snap/go/10455/src/net/port.go:53
			return n > cutoff
//line /snap/go/10455/src/net/port.go:53
			// _ = "end of CoverTab[7925]"
//line /snap/go/10455/src/net/port.go:53
		}() {
//line /snap/go/10455/src/net/port.go:53
			_go_fuzz_dep_.CoverTab[529554]++
//line /snap/go/10455/src/net/port.go:53
			_go_fuzz_dep_.CoverTab[7926]++
								port = int(cutoff)
//line /snap/go/10455/src/net/port.go:54
			// _ = "end of CoverTab[7926]"
		} else {
//line /snap/go/10455/src/net/port.go:55
			_go_fuzz_dep_.CoverTab[529555]++
//line /snap/go/10455/src/net/port.go:55
			_go_fuzz_dep_.CoverTab[7927]++
								port = int(n)
//line /snap/go/10455/src/net/port.go:56
			// _ = "end of CoverTab[7927]"
		}
//line /snap/go/10455/src/net/port.go:57
		// _ = "end of CoverTab[7924]"
//line /snap/go/10455/src/net/port.go:57
	}
//line /snap/go/10455/src/net/port.go:57
	// _ = "end of CoverTab[7901]"
//line /snap/go/10455/src/net/port.go:57
	_go_fuzz_dep_.CoverTab[7902]++
						if neg {
//line /snap/go/10455/src/net/port.go:58
		_go_fuzz_dep_.CoverTab[529556]++
//line /snap/go/10455/src/net/port.go:58
		_go_fuzz_dep_.CoverTab[7928]++
							port = -port
//line /snap/go/10455/src/net/port.go:59
		// _ = "end of CoverTab[7928]"
	} else {
//line /snap/go/10455/src/net/port.go:60
		_go_fuzz_dep_.CoverTab[529557]++
//line /snap/go/10455/src/net/port.go:60
		_go_fuzz_dep_.CoverTab[7929]++
//line /snap/go/10455/src/net/port.go:60
		// _ = "end of CoverTab[7929]"
//line /snap/go/10455/src/net/port.go:60
	}
//line /snap/go/10455/src/net/port.go:60
	// _ = "end of CoverTab[7902]"
//line /snap/go/10455/src/net/port.go:60
	_go_fuzz_dep_.CoverTab[7903]++
						return port, false
//line /snap/go/10455/src/net/port.go:61
	// _ = "end of CoverTab[7903]"
}

//line /snap/go/10455/src/net/port.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/port.go:62
var _ = _go_fuzz_dep_.CoverTab
