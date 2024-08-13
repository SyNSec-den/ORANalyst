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

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
package uri

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:19
)

import (
	"strings"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

const upperhex = "0123456789ABCDEF"

type encoding int

const (
	encodePath	encoding	= 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

// stringContainsCTLByte reports whether s contains any ASCII control character.
func stringContainsCTLByte(s string) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:42
	_go_fuzz_dep_.CoverTab[183460]++
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:43
		_go_fuzz_dep_.CoverTab[183462]++
														b := s[i]
														if b < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:45
			_go_fuzz_dep_.CoverTab[183463]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:45
			return b == 0x7f
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:45
			// _ = "end of CoverTab[183463]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:45
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:45
			_go_fuzz_dep_.CoverTab[183464]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:46
			// _ = "end of CoverTab[183464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:47
			_go_fuzz_dep_.CoverTab[183465]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:47
			// _ = "end of CoverTab[183465]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:47
		// _ = "end of CoverTab[183462]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:48
	// _ = "end of CoverTab[183460]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:48
	_go_fuzz_dep_.CoverTab[183461]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:49
	// _ = "end of CoverTab[183461]"
}

func ishex(c byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:52
	_go_fuzz_dep_.CoverTab[183466]++
													switch {
	case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:54
		_go_fuzz_dep_.CoverTab[183472]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:54
		return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:54
		// _ = "end of CoverTab[183472]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:54
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:54
		_go_fuzz_dep_.CoverTab[183468]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:55
		// _ = "end of CoverTab[183468]"
	case 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:56
		_go_fuzz_dep_.CoverTab[183473]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:56
		return c <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:56
		// _ = "end of CoverTab[183473]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:56
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:56
		_go_fuzz_dep_.CoverTab[183469]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:57
		// _ = "end of CoverTab[183469]"
	case 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:58
		_go_fuzz_dep_.CoverTab[183474]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:58
		return c <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:58
		// _ = "end of CoverTab[183474]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:58
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:58
		_go_fuzz_dep_.CoverTab[183470]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:59
		// _ = "end of CoverTab[183470]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:59
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:59
		_go_fuzz_dep_.CoverTab[183471]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:59
		// _ = "end of CoverTab[183471]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:60
	// _ = "end of CoverTab[183466]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:60
	_go_fuzz_dep_.CoverTab[183467]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:61
	// _ = "end of CoverTab[183467]"
}

func unhex(c byte) byte {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:64
	_go_fuzz_dep_.CoverTab[183475]++
													switch {
	case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:66
		_go_fuzz_dep_.CoverTab[183481]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:66
		return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:66
		// _ = "end of CoverTab[183481]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:66
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:66
		_go_fuzz_dep_.CoverTab[183477]++
														return c - '0'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:67
		// _ = "end of CoverTab[183477]"
	case 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:68
		_go_fuzz_dep_.CoverTab[183482]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:68
		return c <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:68
		// _ = "end of CoverTab[183482]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:68
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:68
		_go_fuzz_dep_.CoverTab[183478]++
														return c - 'a' + 10
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:69
		// _ = "end of CoverTab[183478]"
	case 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:70
		_go_fuzz_dep_.CoverTab[183483]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:70
		return c <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:70
		// _ = "end of CoverTab[183483]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:70
	}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:70
		_go_fuzz_dep_.CoverTab[183479]++
														return c - 'A' + 10
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:71
		// _ = "end of CoverTab[183479]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:71
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:71
		_go_fuzz_dep_.CoverTab[183480]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:71
		// _ = "end of CoverTab[183480]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:72
	// _ = "end of CoverTab[183475]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:72
	_go_fuzz_dep_.CoverTab[183476]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:73
	// _ = "end of CoverTab[183476]"
}

// Return true if the specified character should be escaped when
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:76
// appearing in a URI string, according to RFC 3986.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:76
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:76
// Please be informed that for now shouldEscape does not check all
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:76
// reserved characters correctly. See golang.org/issue/5684.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:81
func shouldEscape(c byte, mode encoding) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:81
	_go_fuzz_dep_.CoverTab[183484]++

													if 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		_go_fuzz_dep_.CoverTab[183489]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		return c <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		// _ = "end of CoverTab[183489]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		_go_fuzz_dep_.CoverTab[183490]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			_go_fuzz_dep_.CoverTab[183491]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			return c <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			// _ = "end of CoverTab[183491]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		// _ = "end of CoverTab[183490]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		_go_fuzz_dep_.CoverTab[183492]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		return '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			_go_fuzz_dep_.CoverTab[183493]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
			// _ = "end of CoverTab[183493]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		// _ = "end of CoverTab[183492]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:83
		_go_fuzz_dep_.CoverTab[183494]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:84
		// _ = "end of CoverTab[183494]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:85
		_go_fuzz_dep_.CoverTab[183495]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:85
		// _ = "end of CoverTab[183495]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:85
	// _ = "end of CoverTab[183484]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:85
	_go_fuzz_dep_.CoverTab[183485]++

													if mode == encodeHost || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:87
		_go_fuzz_dep_.CoverTab[183496]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:87
		return mode == encodeZone
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:87
		// _ = "end of CoverTab[183496]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:87
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:87
		_go_fuzz_dep_.CoverTab[183497]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:97
		switch c {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '[', ']', '<', '>', '"':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:98
			_go_fuzz_dep_.CoverTab[183498]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:99
			// _ = "end of CoverTab[183498]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:99
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:99
			_go_fuzz_dep_.CoverTab[183499]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:99
			// _ = "end of CoverTab[183499]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:100
		// _ = "end of CoverTab[183497]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:101
		_go_fuzz_dep_.CoverTab[183500]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:101
		// _ = "end of CoverTab[183500]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:101
	// _ = "end of CoverTab[183485]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:101
	_go_fuzz_dep_.CoverTab[183486]++

													switch c {
	case '-', '_', '.', '~':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:104
		_go_fuzz_dep_.CoverTab[183501]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:105
		// _ = "end of CoverTab[183501]"

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:107
		_go_fuzz_dep_.CoverTab[183502]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:110
		switch mode {
		case encodePath:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:111
			_go_fuzz_dep_.CoverTab[183504]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:116
			return c == '?'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:116
			// _ = "end of CoverTab[183504]"

		case encodePathSegment:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:118
			_go_fuzz_dep_.CoverTab[183505]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
			return c == '/' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				_go_fuzz_dep_.CoverTab[183510]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				return c == ';'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				// _ = "end of CoverTab[183510]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				_go_fuzz_dep_.CoverTab[183511]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				return c == ','
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				// _ = "end of CoverTab[183511]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				_go_fuzz_dep_.CoverTab[183512]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				return c == '?'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
				// _ = "end of CoverTab[183512]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
			}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:121
			// _ = "end of CoverTab[183505]"

		case encodeUserPassword:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:123
			_go_fuzz_dep_.CoverTab[183506]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
			return c == '@' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				_go_fuzz_dep_.CoverTab[183513]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				return c == '/'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				// _ = "end of CoverTab[183513]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				_go_fuzz_dep_.CoverTab[183514]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				return c == '?'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				// _ = "end of CoverTab[183514]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				_go_fuzz_dep_.CoverTab[183515]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				return c == ':'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
				// _ = "end of CoverTab[183515]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
			}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:128
			// _ = "end of CoverTab[183506]"

		case encodeQueryComponent:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:130
			_go_fuzz_dep_.CoverTab[183507]++

															return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:132
			// _ = "end of CoverTab[183507]"

		case encodeFragment:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:134
			_go_fuzz_dep_.CoverTab[183508]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:137
			return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:137
			// _ = "end of CoverTab[183508]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:137
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:137
			_go_fuzz_dep_.CoverTab[183509]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:137
			// _ = "end of CoverTab[183509]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:138
		// _ = "end of CoverTab[183502]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:138
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:138
		_go_fuzz_dep_.CoverTab[183503]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:138
		// _ = "end of CoverTab[183503]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:139
	// _ = "end of CoverTab[183486]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:139
	_go_fuzz_dep_.CoverTab[183487]++

													if mode == encodeFragment {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:141
		_go_fuzz_dep_.CoverTab[183516]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:148
		switch c {
		case '!', '(', ')', '*':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:149
			_go_fuzz_dep_.CoverTab[183517]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:150
			// _ = "end of CoverTab[183517]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:150
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:150
			_go_fuzz_dep_.CoverTab[183518]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:150
			// _ = "end of CoverTab[183518]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:151
		// _ = "end of CoverTab[183516]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:152
		_go_fuzz_dep_.CoverTab[183519]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:152
		// _ = "end of CoverTab[183519]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:152
	// _ = "end of CoverTab[183487]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:152
	_go_fuzz_dep_.CoverTab[183488]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:155
	return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:155
	// _ = "end of CoverTab[183488]"
}

// unescape unescapes a string; the mode specifies
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:158
// which section of the URI string is being unescaped.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:160
func unescape(s string, mode encoding) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:160
	_go_fuzz_dep_.CoverTab[183520]++

													n := 0
													hasPlus := false
													for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:164
		_go_fuzz_dep_.CoverTab[183524]++
														switch s[i] {
		case '%':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:166
			_go_fuzz_dep_.CoverTab[183525]++
															n++
															if i+2 >= len(s) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				_go_fuzz_dep_.CoverTab[183532]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				return !ishex(s[i+1])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				// _ = "end of CoverTab[183532]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				_go_fuzz_dep_.CoverTab[183533]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				return !ishex(s[i+2])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				// _ = "end of CoverTab[183533]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:168
				_go_fuzz_dep_.CoverTab[183534]++
																s = s[i:]
																if len(s) > 3 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:170
					_go_fuzz_dep_.CoverTab[183536]++
																	s = s[:3]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:171
					// _ = "end of CoverTab[183536]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:172
					_go_fuzz_dep_.CoverTab[183537]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:172
					// _ = "end of CoverTab[183537]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:172
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:172
				// _ = "end of CoverTab[183534]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:172
				_go_fuzz_dep_.CoverTab[183535]++
																return "", errors.NewInvalid("invalid URI escape ", s)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:173
				// _ = "end of CoverTab[183535]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:174
				_go_fuzz_dep_.CoverTab[183538]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:174
				// _ = "end of CoverTab[183538]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:174
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:174
			// _ = "end of CoverTab[183525]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:174
			_go_fuzz_dep_.CoverTab[183526]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
			if mode == encodeHost && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				_go_fuzz_dep_.CoverTab[183539]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				return unhex(s[i+1]) < 8
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				// _ = "end of CoverTab[183539]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				_go_fuzz_dep_.CoverTab[183540]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				return s[i:i+3] != "%25"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				// _ = "end of CoverTab[183540]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:181
				_go_fuzz_dep_.CoverTab[183541]++
																return "", errors.NewInvalid("invalid URI escape ", s[i:i+3])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:182
				// _ = "end of CoverTab[183541]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:183
				_go_fuzz_dep_.CoverTab[183542]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:183
				// _ = "end of CoverTab[183542]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:183
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:183
			// _ = "end of CoverTab[183526]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:183
			_go_fuzz_dep_.CoverTab[183527]++
															if mode == encodeZone {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:184
				_go_fuzz_dep_.CoverTab[183543]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:192
				v := unhex(s[i+1])<<4 | unhex(s[i+2])
				if s[i:i+3] != "%25" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					_go_fuzz_dep_.CoverTab[183544]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					return v != ' '
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					// _ = "end of CoverTab[183544]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					_go_fuzz_dep_.CoverTab[183545]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					return shouldEscape(v, encodeHost)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					// _ = "end of CoverTab[183545]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
				}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:193
					_go_fuzz_dep_.CoverTab[183546]++
																	return "", errors.NewInvalid("invalid URI escape ", s[i:i+3])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:194
					// _ = "end of CoverTab[183546]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:195
					_go_fuzz_dep_.CoverTab[183547]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:195
					// _ = "end of CoverTab[183547]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:195
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:195
				// _ = "end of CoverTab[183543]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:196
				_go_fuzz_dep_.CoverTab[183548]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:196
				// _ = "end of CoverTab[183548]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:196
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:196
			// _ = "end of CoverTab[183527]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:196
			_go_fuzz_dep_.CoverTab[183528]++
															i += 3
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:197
			// _ = "end of CoverTab[183528]"
		case '+':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:198
			_go_fuzz_dep_.CoverTab[183529]++
															hasPlus = mode == encodeQueryComponent
															i++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:200
			// _ = "end of CoverTab[183529]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:201
			_go_fuzz_dep_.CoverTab[183530]++
															if (mode == encodeHost || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				_go_fuzz_dep_.CoverTab[183549]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				return mode == encodeZone
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				// _ = "end of CoverTab[183549]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				_go_fuzz_dep_.CoverTab[183550]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				return s[i] < 0x80
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				// _ = "end of CoverTab[183550]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				_go_fuzz_dep_.CoverTab[183551]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				return shouldEscape(s[i], mode)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				// _ = "end of CoverTab[183551]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:202
				_go_fuzz_dep_.CoverTab[183552]++
																return "", errors.NewInvalid("Invalid character in hostname", s[i:i+1])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:203
				// _ = "end of CoverTab[183552]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:204
				_go_fuzz_dep_.CoverTab[183553]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:204
				// _ = "end of CoverTab[183553]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:204
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:204
			// _ = "end of CoverTab[183530]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:204
			_go_fuzz_dep_.CoverTab[183531]++
															i++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:205
			// _ = "end of CoverTab[183531]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:206
		// _ = "end of CoverTab[183524]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:207
	// _ = "end of CoverTab[183520]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:207
	_go_fuzz_dep_.CoverTab[183521]++

													if n == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:209
		_go_fuzz_dep_.CoverTab[183554]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:209
		return !hasPlus
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:209
		// _ = "end of CoverTab[183554]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:209
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:209
		_go_fuzz_dep_.CoverTab[183555]++
														return s, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:210
		// _ = "end of CoverTab[183555]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:211
		_go_fuzz_dep_.CoverTab[183556]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:211
		// _ = "end of CoverTab[183556]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:211
	// _ = "end of CoverTab[183521]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:211
	_go_fuzz_dep_.CoverTab[183522]++

													var t strings.Builder
													t.Grow(len(s) - 2*n)
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:215
		_go_fuzz_dep_.CoverTab[183557]++
														switch s[i] {
		case '%':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:217
			_go_fuzz_dep_.CoverTab[183558]++
															t.WriteByte(unhex(s[i+1])<<4 | unhex(s[i+2]))
															i += 2
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:219
			// _ = "end of CoverTab[183558]"
		case '+':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:220
			_go_fuzz_dep_.CoverTab[183559]++
															if mode == encodeQueryComponent {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:221
				_go_fuzz_dep_.CoverTab[183561]++
																t.WriteByte(' ')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:222
				// _ = "end of CoverTab[183561]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:223
				_go_fuzz_dep_.CoverTab[183562]++
																t.WriteByte('+')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:224
				// _ = "end of CoverTab[183562]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:225
			// _ = "end of CoverTab[183559]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:226
			_go_fuzz_dep_.CoverTab[183560]++
															t.WriteByte(s[i])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:227
			// _ = "end of CoverTab[183560]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:228
		// _ = "end of CoverTab[183557]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:229
	// _ = "end of CoverTab[183522]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:229
	_go_fuzz_dep_.CoverTab[183523]++
													return t.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:230
	// _ = "end of CoverTab[183523]"
}

func escape(s string, mode encoding) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:233
	_go_fuzz_dep_.CoverTab[183563]++
													spaceCount, hexCount := 0, 0
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:235
		_go_fuzz_dep_.CoverTab[183569]++
														c := s[i]
														if shouldEscape(c, mode) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:237
			_go_fuzz_dep_.CoverTab[183570]++
															if c == ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:238
				_go_fuzz_dep_.CoverTab[183571]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:238
				return mode == encodeQueryComponent
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:238
				// _ = "end of CoverTab[183571]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:238
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:238
				_go_fuzz_dep_.CoverTab[183572]++
																spaceCount++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:239
				// _ = "end of CoverTab[183572]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:240
				_go_fuzz_dep_.CoverTab[183573]++
																hexCount++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:241
				// _ = "end of CoverTab[183573]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:242
			// _ = "end of CoverTab[183570]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:243
			_go_fuzz_dep_.CoverTab[183574]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:243
			// _ = "end of CoverTab[183574]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:243
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:243
		// _ = "end of CoverTab[183569]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:244
	// _ = "end of CoverTab[183563]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:244
	_go_fuzz_dep_.CoverTab[183564]++

													if spaceCount == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:246
		_go_fuzz_dep_.CoverTab[183575]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:246
		return hexCount == 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:246
		// _ = "end of CoverTab[183575]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:246
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:246
		_go_fuzz_dep_.CoverTab[183576]++
														return s
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:247
		// _ = "end of CoverTab[183576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:248
		_go_fuzz_dep_.CoverTab[183577]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:248
		// _ = "end of CoverTab[183577]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:248
	// _ = "end of CoverTab[183564]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:248
	_go_fuzz_dep_.CoverTab[183565]++

													var buf [64]byte
													var t []byte

													required := len(s) + 2*hexCount
													if required <= len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:254
		_go_fuzz_dep_.CoverTab[183578]++
														t = buf[:required]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:255
		// _ = "end of CoverTab[183578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:256
		_go_fuzz_dep_.CoverTab[183579]++
														t = make([]byte, required)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:257
		// _ = "end of CoverTab[183579]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:258
	// _ = "end of CoverTab[183565]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:258
	_go_fuzz_dep_.CoverTab[183566]++

													if hexCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:260
		_go_fuzz_dep_.CoverTab[183580]++
														copy(t, s)
														for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:262
			_go_fuzz_dep_.CoverTab[183582]++
															if s[i] == ' ' {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:263
				_go_fuzz_dep_.CoverTab[183583]++
																t[i] = '+'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:264
				// _ = "end of CoverTab[183583]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:265
				_go_fuzz_dep_.CoverTab[183584]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:265
				// _ = "end of CoverTab[183584]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:265
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:265
			// _ = "end of CoverTab[183582]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:266
		// _ = "end of CoverTab[183580]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:266
		_go_fuzz_dep_.CoverTab[183581]++
														return string(t)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:267
		// _ = "end of CoverTab[183581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:268
		_go_fuzz_dep_.CoverTab[183585]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:268
		// _ = "end of CoverTab[183585]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:268
	// _ = "end of CoverTab[183566]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:268
	_go_fuzz_dep_.CoverTab[183567]++

													j := 0
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:271
		_go_fuzz_dep_.CoverTab[183586]++
														switch c := s[i]; {
		case c == ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:273
			_go_fuzz_dep_.CoverTab[183590]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:273
			return mode == encodeQueryComponent
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:273
			// _ = "end of CoverTab[183590]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:273
		}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:273
			_go_fuzz_dep_.CoverTab[183587]++
															t[j] = '+'
															j++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:275
			// _ = "end of CoverTab[183587]"
		case shouldEscape(c, mode):
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:276
			_go_fuzz_dep_.CoverTab[183588]++
															t[j] = '%'
															t[j+1] = upperhex[c>>4]
															t[j+2] = upperhex[c&15]
															j += 3
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:280
			// _ = "end of CoverTab[183588]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:281
			_go_fuzz_dep_.CoverTab[183589]++
															t[j] = s[i]
															j++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:283
			// _ = "end of CoverTab[183589]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:284
		// _ = "end of CoverTab[183586]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:285
	// _ = "end of CoverTab[183567]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:285
	_go_fuzz_dep_.CoverTab[183568]++
													return string(t)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:286
	// _ = "end of CoverTab[183568]"
}

// split slices s into two substrings separated by the first occurrence of
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:289
// sep. If cutc is true then sep is excluded from the second substring.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:289
// If sep does not occur in s then s and the empty string is returned.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:292
func split(s string, sep byte, cutc bool) (string, string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:292
	_go_fuzz_dep_.CoverTab[183591]++
													i := strings.IndexByte(s, sep)
													if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:294
		_go_fuzz_dep_.CoverTab[183594]++
														return s, ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:295
		// _ = "end of CoverTab[183594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:296
		_go_fuzz_dep_.CoverTab[183595]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:296
		// _ = "end of CoverTab[183595]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:296
	// _ = "end of CoverTab[183591]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:296
	_go_fuzz_dep_.CoverTab[183592]++
													if cutc {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:297
		_go_fuzz_dep_.CoverTab[183596]++
														return s[:i], s[i+1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:298
		// _ = "end of CoverTab[183596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:299
		_go_fuzz_dep_.CoverTab[183597]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:299
		// _ = "end of CoverTab[183597]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:299
	// _ = "end of CoverTab[183592]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:299
	_go_fuzz_dep_.CoverTab[183593]++
													return s[:i], s[i:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:300
	// _ = "end of CoverTab[183593]"
}

// validUserinfo reports whether s is a valid userinfo string per RFC 3986
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
// Section 3.2.1:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//	userinfo    = *( unreserved / pct-encoded / sub-delims / ":" )
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//	unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//	              / "*" / "+" / "," / ";" / "="
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:303
// It doesn't validate pct-encoded. The caller does that via func unescape.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:311
func validUserinfo(s string) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:311
	_go_fuzz_dep_.CoverTab[183598]++
													for _, r := range s {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:312
		_go_fuzz_dep_.CoverTab[183600]++
														if 'A' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:313
			_go_fuzz_dep_.CoverTab[183604]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:313
			return r <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:313
			// _ = "end of CoverTab[183604]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:313
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:313
			_go_fuzz_dep_.CoverTab[183605]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:314
			// _ = "end of CoverTab[183605]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:315
			_go_fuzz_dep_.CoverTab[183606]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:315
			// _ = "end of CoverTab[183606]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:315
		// _ = "end of CoverTab[183600]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:315
		_go_fuzz_dep_.CoverTab[183601]++
														if 'a' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:316
			_go_fuzz_dep_.CoverTab[183607]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:316
			return r <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:316
			// _ = "end of CoverTab[183607]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:316
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:316
			_go_fuzz_dep_.CoverTab[183608]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:317
			// _ = "end of CoverTab[183608]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:318
			_go_fuzz_dep_.CoverTab[183609]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:318
			// _ = "end of CoverTab[183609]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:318
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:318
		// _ = "end of CoverTab[183601]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:318
		_go_fuzz_dep_.CoverTab[183602]++
														if '0' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:319
			_go_fuzz_dep_.CoverTab[183610]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:319
			return r <= '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:319
			// _ = "end of CoverTab[183610]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:319
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:319
			_go_fuzz_dep_.CoverTab[183611]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:320
			// _ = "end of CoverTab[183611]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:321
			_go_fuzz_dep_.CoverTab[183612]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:321
			// _ = "end of CoverTab[183612]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:321
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:321
		// _ = "end of CoverTab[183602]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:321
		_go_fuzz_dep_.CoverTab[183603]++
														switch r {
		case '-', '.', '_', ':', '~', '!', '$', '&', '\'',
			'(', ')', '*', '+', ',', ';', '=', '%', '@':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:324
			_go_fuzz_dep_.CoverTab[183613]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:325
			// _ = "end of CoverTab[183613]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:326
			_go_fuzz_dep_.CoverTab[183614]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:327
			// _ = "end of CoverTab[183614]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:328
		// _ = "end of CoverTab[183603]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:329
	// _ = "end of CoverTab[183598]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:329
	_go_fuzz_dep_.CoverTab[183599]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:330
	// _ = "end of CoverTab[183599]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:331
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/utils.go:331
var _ = _go_fuzz_dep_.CoverTab
