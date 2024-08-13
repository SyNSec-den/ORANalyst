// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
package uri

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:15
)

// Options URI options for creating a new URI
type Options struct {
	scheme		string
	opaque		string		// encoded opaque data
	user		*Userinfo	// username and password information
	host		string		// host or host:port
	path		string		// path (relative paths may omit leading slash)
	rawPath		string		// encoded path hint (see EscapedPath method)
	forceQuery	bool		// append a query ('?') even if RawQuery is empty
	rawQuery	string		// encoded query values, without '?'
	fragment	string		// fragment for references, without '#'
	rawFragment	string		// encoded fragment hint (see EscapedFragment method)
}

// Option URI option
type Option func(options *Options)

// NewURI creates a new URI
func NewURI(opts ...Option) *URI {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:35
	_go_fuzz_dep_.CoverTab[183056]++
													URIOptions := &Options{}
													for _, option := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:37
		_go_fuzz_dep_.CoverTab[183058]++
														option(URIOptions)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:38
		// _ = "end of CoverTab[183058]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:39
	// _ = "end of CoverTab[183056]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:39
	_go_fuzz_dep_.CoverTab[183057]++

													return &URI{
		Scheme:		URIOptions.scheme,
		Opaque:		URIOptions.opaque,
		User:		URIOptions.user,
		Path:		URIOptions.path,
		RawPath:	URIOptions.rawPath,
		ForceQuery:	URIOptions.forceQuery,
		RawQuery:	URIOptions.rawQuery,
		Fragment:	URIOptions.fragment,
		RawFragment:	URIOptions.rawFragment,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:51
	// _ = "end of CoverTab[183057]"

}

// WithScheme sets URI scheme
func WithScheme(scheme string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:56
	_go_fuzz_dep_.CoverTab[183059]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:57
		_go_fuzz_dep_.CoverTab[183060]++
														options.scheme = scheme
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:58
		// _ = "end of CoverTab[183060]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:60
	// _ = "end of CoverTab[183059]"
}

// WithOpaque  sets URI opaque
func WithOpaque(opaque string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:64
	_go_fuzz_dep_.CoverTab[183061]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:65
		_go_fuzz_dep_.CoverTab[183062]++
														options.opaque = opaque
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:66
		// _ = "end of CoverTab[183062]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:68
	// _ = "end of CoverTab[183061]"
}

// WithUser  sets URI user information
func WithUser(user *Userinfo) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:72
	_go_fuzz_dep_.CoverTab[183063]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:73
		_go_fuzz_dep_.CoverTab[183064]++
														options.user = user
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:74
		// _ = "end of CoverTab[183064]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:76
	// _ = "end of CoverTab[183063]"
}

// WithHost  sets URI host information
func WithHost(host string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:80
	_go_fuzz_dep_.CoverTab[183065]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:81
		_go_fuzz_dep_.CoverTab[183066]++
														options.host = host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:82
		// _ = "end of CoverTab[183066]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:84
	// _ = "end of CoverTab[183065]"
}

// WithPath  sets URI path information
func WithPath(path string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:88
	_go_fuzz_dep_.CoverTab[183067]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:89
		_go_fuzz_dep_.CoverTab[183068]++
														options.path = path
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:90
		// _ = "end of CoverTab[183068]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:92
	// _ = "end of CoverTab[183067]"
}

// WithRawPath  sets URI raw path information
func WithRawPath(rawPath string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:96
	_go_fuzz_dep_.CoverTab[183069]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:97
		_go_fuzz_dep_.CoverTab[183070]++
														options.rawPath = rawPath
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:98
		// _ = "end of CoverTab[183070]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:100
	// _ = "end of CoverTab[183069]"
}

// WithForceQuery  sets URI force query information
func WithForceQuery(forceQuery bool) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:104
	_go_fuzz_dep_.CoverTab[183071]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:105
		_go_fuzz_dep_.CoverTab[183072]++
														options.forceQuery = forceQuery
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:106
		// _ = "end of CoverTab[183072]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:108
	// _ = "end of CoverTab[183071]"
}

// WithRawQuery  sets URI raw query information
func WithRawQuery(rawQuery string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:112
	_go_fuzz_dep_.CoverTab[183073]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:113
		_go_fuzz_dep_.CoverTab[183074]++
														options.rawQuery = rawQuery
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:114
		// _ = "end of CoverTab[183074]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:116
	// _ = "end of CoverTab[183073]"
}

// WithFragment  sets URI fragment information
func WithFragment(fragment string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:120
	_go_fuzz_dep_.CoverTab[183075]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:121
		_go_fuzz_dep_.CoverTab[183076]++
														options.fragment = fragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:122
		// _ = "end of CoverTab[183076]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:124
	// _ = "end of CoverTab[183075]"
}

// WithRawFragment  sets URI raw fragment information
func WithRawFragment(rawFragment string) func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:128
	_go_fuzz_dep_.CoverTab[183077]++
													return func(options *Options) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:129
		_go_fuzz_dep_.CoverTab[183078]++
														options.rawFragment = rawFragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:130
		// _ = "end of CoverTab[183078]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:132
	// _ = "end of CoverTab[183077]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:133
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/options.go:133
var _ = _go_fuzz_dep_.CoverTab
