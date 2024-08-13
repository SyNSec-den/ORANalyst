// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/dnsconfig.go:5
package net

//line /snap/go/10455/src/net/dnsconfig.go:5
import (
//line /snap/go/10455/src/net/dnsconfig.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/dnsconfig.go:5
)
//line /snap/go/10455/src/net/dnsconfig.go:5
import (
//line /snap/go/10455/src/net/dnsconfig.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/dnsconfig.go:5
)

import (
	"os"
	"sync/atomic"
	"time"
)

var (
	defaultNS	= []string{"127.0.0.1:53", "[::1]:53"}
	getHostname	= os.Hostname	// variable for testing
)

type dnsConfig struct {
	servers		[]string	// server addresses (in host:port form) to use
	search		[]string	// rooted suffixes to append to local name
	ndots		int		// number of dots in name to trigger absolute lookup
	timeout		time.Duration	// wait before giving up on a query, including retries
	attempts	int		// lost packets before giving up on server
	rotate		bool		// round robin among servers
	unknownOpt	bool		// anything unknown was encountered
	lookup		[]string	// OpenBSD top-level database "lookup" order
	err		error		// any error that occurs during open of resolv.conf
	mtime		time.Time	// time of resolv.conf modification
	soffset		uint32		// used by serverOffset
	singleRequest	bool		// use sequential A and AAAA queries instead of parallel queries
	useTCP		bool		// force usage of TCP for DNS resolutions
	trustAD		bool		// add AD flag to queries
	noReload	bool		// do not check for config file updates
}

// serverOffset returns an offset that can be used to determine
//line /snap/go/10455/src/net/dnsconfig.go:36
// indices of servers in c.servers when making queries.
//line /snap/go/10455/src/net/dnsconfig.go:36
// When the rotate option is enabled, this offset increases.
//line /snap/go/10455/src/net/dnsconfig.go:36
// Otherwise it is always 0.
//line /snap/go/10455/src/net/dnsconfig.go:40
func (c *dnsConfig) serverOffset() uint32 {
//line /snap/go/10455/src/net/dnsconfig.go:40
	_go_fuzz_dep_.CoverTab[5720]++
						if c.rotate {
//line /snap/go/10455/src/net/dnsconfig.go:41
		_go_fuzz_dep_.CoverTab[528185]++
//line /snap/go/10455/src/net/dnsconfig.go:41
		_go_fuzz_dep_.CoverTab[5722]++
							return atomic.AddUint32(&c.soffset, 1) - 1
//line /snap/go/10455/src/net/dnsconfig.go:42
		// _ = "end of CoverTab[5722]"
	} else {
//line /snap/go/10455/src/net/dnsconfig.go:43
		_go_fuzz_dep_.CoverTab[528186]++
//line /snap/go/10455/src/net/dnsconfig.go:43
		_go_fuzz_dep_.CoverTab[5723]++
//line /snap/go/10455/src/net/dnsconfig.go:43
		// _ = "end of CoverTab[5723]"
//line /snap/go/10455/src/net/dnsconfig.go:43
	}
//line /snap/go/10455/src/net/dnsconfig.go:43
	// _ = "end of CoverTab[5720]"
//line /snap/go/10455/src/net/dnsconfig.go:43
	_go_fuzz_dep_.CoverTab[5721]++
						return 0
//line /snap/go/10455/src/net/dnsconfig.go:44
	// _ = "end of CoverTab[5721]"
}

//line /snap/go/10455/src/net/dnsconfig.go:45
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/dnsconfig.go:45
var _ = _go_fuzz_dep_.CoverTab
