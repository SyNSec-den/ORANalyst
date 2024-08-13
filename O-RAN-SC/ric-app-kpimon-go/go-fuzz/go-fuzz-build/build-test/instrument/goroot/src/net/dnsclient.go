// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/dnsclient.go:5
package net

//line /snap/go/10455/src/net/dnsclient.go:5
import (
//line /snap/go/10455/src/net/dnsclient.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/dnsclient.go:5
)
//line /snap/go/10455/src/net/dnsclient.go:5
import (
//line /snap/go/10455/src/net/dnsclient.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/dnsclient.go:5
)

import (
	"internal/bytealg"
	"internal/itoa"
	"sort"

	"golang.org/x/net/dns/dnsmessage"
)

// provided by runtime
func fastrandu() uint

func randInt() int {
//line /snap/go/10455/src/net/dnsclient.go:18
	_go_fuzz_dep_.CoverTab[5207]++
						return int(fastrandu() >> 1)
//line /snap/go/10455/src/net/dnsclient.go:19
	// _ = "end of CoverTab[5207]"
}

func randIntn(n int) int {
//line /snap/go/10455/src/net/dnsclient.go:22
	_go_fuzz_dep_.CoverTab[5208]++
						return randInt() % n
//line /snap/go/10455/src/net/dnsclient.go:23
	// _ = "end of CoverTab[5208]"
}

// reverseaddr returns the in-addr.arpa. or ip6.arpa. hostname of the IP
//line /snap/go/10455/src/net/dnsclient.go:26
// address addr suitable for rDNS (PTR) record lookup or an error if it fails
//line /snap/go/10455/src/net/dnsclient.go:26
// to parse the IP address.
//line /snap/go/10455/src/net/dnsclient.go:29
func reverseaddr(addr string) (arpa string, err error) {
//line /snap/go/10455/src/net/dnsclient.go:29
	_go_fuzz_dep_.CoverTab[5209]++
						ip := ParseIP(addr)
						if ip == nil {
//line /snap/go/10455/src/net/dnsclient.go:31
		_go_fuzz_dep_.CoverTab[527831]++
//line /snap/go/10455/src/net/dnsclient.go:31
		_go_fuzz_dep_.CoverTab[5213]++
							return "", &DNSError{Err: "unrecognized address", Name: addr}
//line /snap/go/10455/src/net/dnsclient.go:32
		// _ = "end of CoverTab[5213]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:33
		_go_fuzz_dep_.CoverTab[527832]++
//line /snap/go/10455/src/net/dnsclient.go:33
		_go_fuzz_dep_.CoverTab[5214]++
//line /snap/go/10455/src/net/dnsclient.go:33
		// _ = "end of CoverTab[5214]"
//line /snap/go/10455/src/net/dnsclient.go:33
	}
//line /snap/go/10455/src/net/dnsclient.go:33
	// _ = "end of CoverTab[5209]"
//line /snap/go/10455/src/net/dnsclient.go:33
	_go_fuzz_dep_.CoverTab[5210]++
						if ip.To4() != nil {
//line /snap/go/10455/src/net/dnsclient.go:34
		_go_fuzz_dep_.CoverTab[527833]++
//line /snap/go/10455/src/net/dnsclient.go:34
		_go_fuzz_dep_.CoverTab[5215]++
							return itoa.Uitoa(uint(ip[15])) + "." + itoa.Uitoa(uint(ip[14])) + "." + itoa.Uitoa(uint(ip[13])) + "." + itoa.Uitoa(uint(ip[12])) + ".in-addr.arpa.", nil
//line /snap/go/10455/src/net/dnsclient.go:35
		// _ = "end of CoverTab[5215]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:36
		_go_fuzz_dep_.CoverTab[527834]++
//line /snap/go/10455/src/net/dnsclient.go:36
		_go_fuzz_dep_.CoverTab[5216]++
//line /snap/go/10455/src/net/dnsclient.go:36
		// _ = "end of CoverTab[5216]"
//line /snap/go/10455/src/net/dnsclient.go:36
	}
//line /snap/go/10455/src/net/dnsclient.go:36
	// _ = "end of CoverTab[5210]"
//line /snap/go/10455/src/net/dnsclient.go:36
	_go_fuzz_dep_.CoverTab[5211]++

						buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))
//line /snap/go/10455/src/net/dnsclient.go:38
	_go_fuzz_dep_.CoverTab[786652] = 0

						for i := len(ip) - 1; i >= 0; i-- {
//line /snap/go/10455/src/net/dnsclient.go:40
		if _go_fuzz_dep_.CoverTab[786652] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:40
			_go_fuzz_dep_.CoverTab[527868]++
//line /snap/go/10455/src/net/dnsclient.go:40
		} else {
//line /snap/go/10455/src/net/dnsclient.go:40
			_go_fuzz_dep_.CoverTab[527869]++
//line /snap/go/10455/src/net/dnsclient.go:40
		}
//line /snap/go/10455/src/net/dnsclient.go:40
		_go_fuzz_dep_.CoverTab[786652] = 1
//line /snap/go/10455/src/net/dnsclient.go:40
		_go_fuzz_dep_.CoverTab[5217]++
							v := ip[i]
							buf = append(buf, hexDigit[v&0xF],
			'.',
			hexDigit[v>>4],
			'.')
//line /snap/go/10455/src/net/dnsclient.go:45
		// _ = "end of CoverTab[5217]"
	}
//line /snap/go/10455/src/net/dnsclient.go:46
	if _go_fuzz_dep_.CoverTab[786652] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:46
		_go_fuzz_dep_.CoverTab[527870]++
//line /snap/go/10455/src/net/dnsclient.go:46
	} else {
//line /snap/go/10455/src/net/dnsclient.go:46
		_go_fuzz_dep_.CoverTab[527871]++
//line /snap/go/10455/src/net/dnsclient.go:46
	}
//line /snap/go/10455/src/net/dnsclient.go:46
	// _ = "end of CoverTab[5211]"
//line /snap/go/10455/src/net/dnsclient.go:46
	_go_fuzz_dep_.CoverTab[5212]++

						buf = append(buf, "ip6.arpa."...)
						return string(buf), nil
//line /snap/go/10455/src/net/dnsclient.go:49
	// _ = "end of CoverTab[5212]"
}

func equalASCIIName(x, y dnsmessage.Name) bool {
//line /snap/go/10455/src/net/dnsclient.go:52
	_go_fuzz_dep_.CoverTab[5218]++
						if x.Length != y.Length {
//line /snap/go/10455/src/net/dnsclient.go:53
		_go_fuzz_dep_.CoverTab[527835]++
//line /snap/go/10455/src/net/dnsclient.go:53
		_go_fuzz_dep_.CoverTab[5221]++
							return false
//line /snap/go/10455/src/net/dnsclient.go:54
		// _ = "end of CoverTab[5221]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:55
		_go_fuzz_dep_.CoverTab[527836]++
//line /snap/go/10455/src/net/dnsclient.go:55
		_go_fuzz_dep_.CoverTab[5222]++
//line /snap/go/10455/src/net/dnsclient.go:55
		// _ = "end of CoverTab[5222]"
//line /snap/go/10455/src/net/dnsclient.go:55
	}
//line /snap/go/10455/src/net/dnsclient.go:55
	// _ = "end of CoverTab[5218]"
//line /snap/go/10455/src/net/dnsclient.go:55
	_go_fuzz_dep_.CoverTab[5219]++
//line /snap/go/10455/src/net/dnsclient.go:55
	_go_fuzz_dep_.CoverTab[786653] = 0
						for i := 0; i < int(x.Length); i++ {
//line /snap/go/10455/src/net/dnsclient.go:56
		if _go_fuzz_dep_.CoverTab[786653] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:56
			_go_fuzz_dep_.CoverTab[527872]++
//line /snap/go/10455/src/net/dnsclient.go:56
		} else {
//line /snap/go/10455/src/net/dnsclient.go:56
			_go_fuzz_dep_.CoverTab[527873]++
//line /snap/go/10455/src/net/dnsclient.go:56
		}
//line /snap/go/10455/src/net/dnsclient.go:56
		_go_fuzz_dep_.CoverTab[786653] = 1
//line /snap/go/10455/src/net/dnsclient.go:56
		_go_fuzz_dep_.CoverTab[5223]++
							a := x.Data[i]
							b := y.Data[i]
							if 'A' <= a && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:59
			_go_fuzz_dep_.CoverTab[5226]++
//line /snap/go/10455/src/net/dnsclient.go:59
			return a <= 'Z'
//line /snap/go/10455/src/net/dnsclient.go:59
			// _ = "end of CoverTab[5226]"
//line /snap/go/10455/src/net/dnsclient.go:59
		}() {
//line /snap/go/10455/src/net/dnsclient.go:59
			_go_fuzz_dep_.CoverTab[527837]++
//line /snap/go/10455/src/net/dnsclient.go:59
			_go_fuzz_dep_.CoverTab[5227]++
								a += 0x20
//line /snap/go/10455/src/net/dnsclient.go:60
			// _ = "end of CoverTab[5227]"
		} else {
//line /snap/go/10455/src/net/dnsclient.go:61
			_go_fuzz_dep_.CoverTab[527838]++
//line /snap/go/10455/src/net/dnsclient.go:61
			_go_fuzz_dep_.CoverTab[5228]++
//line /snap/go/10455/src/net/dnsclient.go:61
			// _ = "end of CoverTab[5228]"
//line /snap/go/10455/src/net/dnsclient.go:61
		}
//line /snap/go/10455/src/net/dnsclient.go:61
		// _ = "end of CoverTab[5223]"
//line /snap/go/10455/src/net/dnsclient.go:61
		_go_fuzz_dep_.CoverTab[5224]++
							if 'A' <= b && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:62
			_go_fuzz_dep_.CoverTab[5229]++
//line /snap/go/10455/src/net/dnsclient.go:62
			return b <= 'Z'
//line /snap/go/10455/src/net/dnsclient.go:62
			// _ = "end of CoverTab[5229]"
//line /snap/go/10455/src/net/dnsclient.go:62
		}() {
//line /snap/go/10455/src/net/dnsclient.go:62
			_go_fuzz_dep_.CoverTab[527839]++
//line /snap/go/10455/src/net/dnsclient.go:62
			_go_fuzz_dep_.CoverTab[5230]++
								b += 0x20
//line /snap/go/10455/src/net/dnsclient.go:63
			// _ = "end of CoverTab[5230]"
		} else {
//line /snap/go/10455/src/net/dnsclient.go:64
			_go_fuzz_dep_.CoverTab[527840]++
//line /snap/go/10455/src/net/dnsclient.go:64
			_go_fuzz_dep_.CoverTab[5231]++
//line /snap/go/10455/src/net/dnsclient.go:64
			// _ = "end of CoverTab[5231]"
//line /snap/go/10455/src/net/dnsclient.go:64
		}
//line /snap/go/10455/src/net/dnsclient.go:64
		// _ = "end of CoverTab[5224]"
//line /snap/go/10455/src/net/dnsclient.go:64
		_go_fuzz_dep_.CoverTab[5225]++
							if a != b {
//line /snap/go/10455/src/net/dnsclient.go:65
			_go_fuzz_dep_.CoverTab[527841]++
//line /snap/go/10455/src/net/dnsclient.go:65
			_go_fuzz_dep_.CoverTab[5232]++
								return false
//line /snap/go/10455/src/net/dnsclient.go:66
			// _ = "end of CoverTab[5232]"
		} else {
//line /snap/go/10455/src/net/dnsclient.go:67
			_go_fuzz_dep_.CoverTab[527842]++
//line /snap/go/10455/src/net/dnsclient.go:67
			_go_fuzz_dep_.CoverTab[5233]++
//line /snap/go/10455/src/net/dnsclient.go:67
			// _ = "end of CoverTab[5233]"
//line /snap/go/10455/src/net/dnsclient.go:67
		}
//line /snap/go/10455/src/net/dnsclient.go:67
		// _ = "end of CoverTab[5225]"
	}
//line /snap/go/10455/src/net/dnsclient.go:68
	if _go_fuzz_dep_.CoverTab[786653] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:68
		_go_fuzz_dep_.CoverTab[527874]++
//line /snap/go/10455/src/net/dnsclient.go:68
	} else {
//line /snap/go/10455/src/net/dnsclient.go:68
		_go_fuzz_dep_.CoverTab[527875]++
//line /snap/go/10455/src/net/dnsclient.go:68
	}
//line /snap/go/10455/src/net/dnsclient.go:68
	// _ = "end of CoverTab[5219]"
//line /snap/go/10455/src/net/dnsclient.go:68
	_go_fuzz_dep_.CoverTab[5220]++
						return true
//line /snap/go/10455/src/net/dnsclient.go:69
	// _ = "end of CoverTab[5220]"
}

// isDomainName checks if a string is a presentation-format domain name
//line /snap/go/10455/src/net/dnsclient.go:72
// (currently restricted to hostname-compatible "preferred name" LDH labels and
//line /snap/go/10455/src/net/dnsclient.go:72
// SRV-like "underscore labels"; see golang.org/issue/12421).
//line /snap/go/10455/src/net/dnsclient.go:75
func isDomainName(s string) bool {
//line /snap/go/10455/src/net/dnsclient.go:75
	_go_fuzz_dep_.CoverTab[5234]++

						if s == "." {
//line /snap/go/10455/src/net/dnsclient.go:77
		_go_fuzz_dep_.CoverTab[527843]++
//line /snap/go/10455/src/net/dnsclient.go:77
		_go_fuzz_dep_.CoverTab[5239]++
							return true
//line /snap/go/10455/src/net/dnsclient.go:78
		// _ = "end of CoverTab[5239]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:79
		_go_fuzz_dep_.CoverTab[527844]++
//line /snap/go/10455/src/net/dnsclient.go:79
		_go_fuzz_dep_.CoverTab[5240]++
//line /snap/go/10455/src/net/dnsclient.go:79
		// _ = "end of CoverTab[5240]"
//line /snap/go/10455/src/net/dnsclient.go:79
	}
//line /snap/go/10455/src/net/dnsclient.go:79
	// _ = "end of CoverTab[5234]"
//line /snap/go/10455/src/net/dnsclient.go:79
	_go_fuzz_dep_.CoverTab[5235]++

//line /snap/go/10455/src/net/dnsclient.go:89
	l := len(s)
	if l == 0 || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[5241]++
//line /snap/go/10455/src/net/dnsclient.go:90
		return l > 254
//line /snap/go/10455/src/net/dnsclient.go:90
		// _ = "end of CoverTab[5241]"
//line /snap/go/10455/src/net/dnsclient.go:90
	}() || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[5242]++
//line /snap/go/10455/src/net/dnsclient.go:90
		return l == 254 && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:90
			_go_fuzz_dep_.CoverTab[5243]++
//line /snap/go/10455/src/net/dnsclient.go:90
			return s[l-1] != '.'
//line /snap/go/10455/src/net/dnsclient.go:90
			// _ = "end of CoverTab[5243]"
//line /snap/go/10455/src/net/dnsclient.go:90
		}()
//line /snap/go/10455/src/net/dnsclient.go:90
		// _ = "end of CoverTab[5242]"
//line /snap/go/10455/src/net/dnsclient.go:90
	}() {
//line /snap/go/10455/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[527845]++
//line /snap/go/10455/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[5244]++
							return false
//line /snap/go/10455/src/net/dnsclient.go:91
		// _ = "end of CoverTab[5244]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:92
		_go_fuzz_dep_.CoverTab[527846]++
//line /snap/go/10455/src/net/dnsclient.go:92
		_go_fuzz_dep_.CoverTab[5245]++
//line /snap/go/10455/src/net/dnsclient.go:92
		// _ = "end of CoverTab[5245]"
//line /snap/go/10455/src/net/dnsclient.go:92
	}
//line /snap/go/10455/src/net/dnsclient.go:92
	// _ = "end of CoverTab[5235]"
//line /snap/go/10455/src/net/dnsclient.go:92
	_go_fuzz_dep_.CoverTab[5236]++

						last := byte('.')
						nonNumeric := false
						partlen := 0
//line /snap/go/10455/src/net/dnsclient.go:96
	_go_fuzz_dep_.CoverTab[786654] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/dnsclient.go:97
		if _go_fuzz_dep_.CoverTab[786654] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:97
			_go_fuzz_dep_.CoverTab[527876]++
//line /snap/go/10455/src/net/dnsclient.go:97
		} else {
//line /snap/go/10455/src/net/dnsclient.go:97
			_go_fuzz_dep_.CoverTab[527877]++
//line /snap/go/10455/src/net/dnsclient.go:97
		}
//line /snap/go/10455/src/net/dnsclient.go:97
		_go_fuzz_dep_.CoverTab[786654] = 1
//line /snap/go/10455/src/net/dnsclient.go:97
		_go_fuzz_dep_.CoverTab[5246]++
							c := s[i]
							switch {
		default:
//line /snap/go/10455/src/net/dnsclient.go:100
			_go_fuzz_dep_.CoverTab[527847]++
//line /snap/go/10455/src/net/dnsclient.go:100
			_go_fuzz_dep_.CoverTab[5248]++
								return false
//line /snap/go/10455/src/net/dnsclient.go:101
			// _ = "end of CoverTab[5248]"
		case 'a' <= c && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[5256]++
//line /snap/go/10455/src/net/dnsclient.go:102
			return c <= 'z'
//line /snap/go/10455/src/net/dnsclient.go:102
			// _ = "end of CoverTab[5256]"
//line /snap/go/10455/src/net/dnsclient.go:102
		}() || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[5257]++
//line /snap/go/10455/src/net/dnsclient.go:102
			return 'A' <= c && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:102
				_go_fuzz_dep_.CoverTab[5258]++
//line /snap/go/10455/src/net/dnsclient.go:102
				return c <= 'Z'
//line /snap/go/10455/src/net/dnsclient.go:102
				// _ = "end of CoverTab[5258]"
//line /snap/go/10455/src/net/dnsclient.go:102
			}()
//line /snap/go/10455/src/net/dnsclient.go:102
			// _ = "end of CoverTab[5257]"
//line /snap/go/10455/src/net/dnsclient.go:102
		}() || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[5259]++
//line /snap/go/10455/src/net/dnsclient.go:102
			return c == '_'
//line /snap/go/10455/src/net/dnsclient.go:102
			// _ = "end of CoverTab[5259]"
//line /snap/go/10455/src/net/dnsclient.go:102
		}():
//line /snap/go/10455/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[527848]++
//line /snap/go/10455/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[5249]++
								nonNumeric = true
								partlen++
//line /snap/go/10455/src/net/dnsclient.go:104
			// _ = "end of CoverTab[5249]"
		case '0' <= c && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:105
			_go_fuzz_dep_.CoverTab[5260]++
//line /snap/go/10455/src/net/dnsclient.go:105
			return c <= '9'
//line /snap/go/10455/src/net/dnsclient.go:105
			// _ = "end of CoverTab[5260]"
//line /snap/go/10455/src/net/dnsclient.go:105
		}():
//line /snap/go/10455/src/net/dnsclient.go:105
			_go_fuzz_dep_.CoverTab[527849]++
//line /snap/go/10455/src/net/dnsclient.go:105
			_go_fuzz_dep_.CoverTab[5250]++

								partlen++
//line /snap/go/10455/src/net/dnsclient.go:107
			// _ = "end of CoverTab[5250]"
		case c == '-':
//line /snap/go/10455/src/net/dnsclient.go:108
			_go_fuzz_dep_.CoverTab[527850]++
//line /snap/go/10455/src/net/dnsclient.go:108
			_go_fuzz_dep_.CoverTab[5251]++

								if last == '.' {
//line /snap/go/10455/src/net/dnsclient.go:110
				_go_fuzz_dep_.CoverTab[527852]++
//line /snap/go/10455/src/net/dnsclient.go:110
				_go_fuzz_dep_.CoverTab[5261]++
									return false
//line /snap/go/10455/src/net/dnsclient.go:111
				// _ = "end of CoverTab[5261]"
			} else {
//line /snap/go/10455/src/net/dnsclient.go:112
				_go_fuzz_dep_.CoverTab[527853]++
//line /snap/go/10455/src/net/dnsclient.go:112
				_go_fuzz_dep_.CoverTab[5262]++
//line /snap/go/10455/src/net/dnsclient.go:112
				// _ = "end of CoverTab[5262]"
//line /snap/go/10455/src/net/dnsclient.go:112
			}
//line /snap/go/10455/src/net/dnsclient.go:112
			// _ = "end of CoverTab[5251]"
//line /snap/go/10455/src/net/dnsclient.go:112
			_go_fuzz_dep_.CoverTab[5252]++
								partlen++
								nonNumeric = true
//line /snap/go/10455/src/net/dnsclient.go:114
			// _ = "end of CoverTab[5252]"
		case c == '.':
//line /snap/go/10455/src/net/dnsclient.go:115
			_go_fuzz_dep_.CoverTab[527851]++
//line /snap/go/10455/src/net/dnsclient.go:115
			_go_fuzz_dep_.CoverTab[5253]++

								if last == '.' || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:117
				_go_fuzz_dep_.CoverTab[5263]++
//line /snap/go/10455/src/net/dnsclient.go:117
				return last == '-'
//line /snap/go/10455/src/net/dnsclient.go:117
				// _ = "end of CoverTab[5263]"
//line /snap/go/10455/src/net/dnsclient.go:117
			}() {
//line /snap/go/10455/src/net/dnsclient.go:117
				_go_fuzz_dep_.CoverTab[527854]++
//line /snap/go/10455/src/net/dnsclient.go:117
				_go_fuzz_dep_.CoverTab[5264]++
									return false
//line /snap/go/10455/src/net/dnsclient.go:118
				// _ = "end of CoverTab[5264]"
			} else {
//line /snap/go/10455/src/net/dnsclient.go:119
				_go_fuzz_dep_.CoverTab[527855]++
//line /snap/go/10455/src/net/dnsclient.go:119
				_go_fuzz_dep_.CoverTab[5265]++
//line /snap/go/10455/src/net/dnsclient.go:119
				// _ = "end of CoverTab[5265]"
//line /snap/go/10455/src/net/dnsclient.go:119
			}
//line /snap/go/10455/src/net/dnsclient.go:119
			// _ = "end of CoverTab[5253]"
//line /snap/go/10455/src/net/dnsclient.go:119
			_go_fuzz_dep_.CoverTab[5254]++
								if partlen > 63 || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:120
				_go_fuzz_dep_.CoverTab[5266]++
//line /snap/go/10455/src/net/dnsclient.go:120
				return partlen == 0
//line /snap/go/10455/src/net/dnsclient.go:120
				// _ = "end of CoverTab[5266]"
//line /snap/go/10455/src/net/dnsclient.go:120
			}() {
//line /snap/go/10455/src/net/dnsclient.go:120
				_go_fuzz_dep_.CoverTab[527856]++
//line /snap/go/10455/src/net/dnsclient.go:120
				_go_fuzz_dep_.CoverTab[5267]++
									return false
//line /snap/go/10455/src/net/dnsclient.go:121
				// _ = "end of CoverTab[5267]"
			} else {
//line /snap/go/10455/src/net/dnsclient.go:122
				_go_fuzz_dep_.CoverTab[527857]++
//line /snap/go/10455/src/net/dnsclient.go:122
				_go_fuzz_dep_.CoverTab[5268]++
//line /snap/go/10455/src/net/dnsclient.go:122
				// _ = "end of CoverTab[5268]"
//line /snap/go/10455/src/net/dnsclient.go:122
			}
//line /snap/go/10455/src/net/dnsclient.go:122
			// _ = "end of CoverTab[5254]"
//line /snap/go/10455/src/net/dnsclient.go:122
			_go_fuzz_dep_.CoverTab[5255]++
								partlen = 0
//line /snap/go/10455/src/net/dnsclient.go:123
			// _ = "end of CoverTab[5255]"
		}
//line /snap/go/10455/src/net/dnsclient.go:124
		// _ = "end of CoverTab[5246]"
//line /snap/go/10455/src/net/dnsclient.go:124
		_go_fuzz_dep_.CoverTab[5247]++
							last = c
//line /snap/go/10455/src/net/dnsclient.go:125
		// _ = "end of CoverTab[5247]"
	}
//line /snap/go/10455/src/net/dnsclient.go:126
	if _go_fuzz_dep_.CoverTab[786654] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:126
		_go_fuzz_dep_.CoverTab[527878]++
//line /snap/go/10455/src/net/dnsclient.go:126
	} else {
//line /snap/go/10455/src/net/dnsclient.go:126
		_go_fuzz_dep_.CoverTab[527879]++
//line /snap/go/10455/src/net/dnsclient.go:126
	}
//line /snap/go/10455/src/net/dnsclient.go:126
	// _ = "end of CoverTab[5236]"
//line /snap/go/10455/src/net/dnsclient.go:126
	_go_fuzz_dep_.CoverTab[5237]++
						if last == '-' || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:127
		_go_fuzz_dep_.CoverTab[5269]++
//line /snap/go/10455/src/net/dnsclient.go:127
		return partlen > 63
//line /snap/go/10455/src/net/dnsclient.go:127
		// _ = "end of CoverTab[5269]"
//line /snap/go/10455/src/net/dnsclient.go:127
	}() {
//line /snap/go/10455/src/net/dnsclient.go:127
		_go_fuzz_dep_.CoverTab[527858]++
//line /snap/go/10455/src/net/dnsclient.go:127
		_go_fuzz_dep_.CoverTab[5270]++
							return false
//line /snap/go/10455/src/net/dnsclient.go:128
		// _ = "end of CoverTab[5270]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:129
		_go_fuzz_dep_.CoverTab[527859]++
//line /snap/go/10455/src/net/dnsclient.go:129
		_go_fuzz_dep_.CoverTab[5271]++
//line /snap/go/10455/src/net/dnsclient.go:129
		// _ = "end of CoverTab[5271]"
//line /snap/go/10455/src/net/dnsclient.go:129
	}
//line /snap/go/10455/src/net/dnsclient.go:129
	// _ = "end of CoverTab[5237]"
//line /snap/go/10455/src/net/dnsclient.go:129
	_go_fuzz_dep_.CoverTab[5238]++

						return nonNumeric
//line /snap/go/10455/src/net/dnsclient.go:131
	// _ = "end of CoverTab[5238]"
}

// absDomainName returns an absolute domain name which ends with a
//line /snap/go/10455/src/net/dnsclient.go:134
// trailing dot to match pure Go reverse resolver and all other lookup
//line /snap/go/10455/src/net/dnsclient.go:134
// routines.
//line /snap/go/10455/src/net/dnsclient.go:134
// See golang.org/issue/12189.
//line /snap/go/10455/src/net/dnsclient.go:134
// But we don't want to add dots for local names from /etc/hosts.
//line /snap/go/10455/src/net/dnsclient.go:134
// It's hard to tell so we settle on the heuristic that names without dots
//line /snap/go/10455/src/net/dnsclient.go:134
// (like "localhost" or "myhost") do not get trailing dots, but any other
//line /snap/go/10455/src/net/dnsclient.go:134
// names do.
//line /snap/go/10455/src/net/dnsclient.go:142
func absDomainName(s string) string {
//line /snap/go/10455/src/net/dnsclient.go:142
	_go_fuzz_dep_.CoverTab[5272]++
						if bytealg.IndexByteString(s, '.') != -1 && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:143
		_go_fuzz_dep_.CoverTab[5274]++
//line /snap/go/10455/src/net/dnsclient.go:143
		return s[len(s)-1] != '.'
//line /snap/go/10455/src/net/dnsclient.go:143
		// _ = "end of CoverTab[5274]"
//line /snap/go/10455/src/net/dnsclient.go:143
	}() {
//line /snap/go/10455/src/net/dnsclient.go:143
		_go_fuzz_dep_.CoverTab[527860]++
//line /snap/go/10455/src/net/dnsclient.go:143
		_go_fuzz_dep_.CoverTab[5275]++
							s += "."
//line /snap/go/10455/src/net/dnsclient.go:144
		// _ = "end of CoverTab[5275]"
	} else {
//line /snap/go/10455/src/net/dnsclient.go:145
		_go_fuzz_dep_.CoverTab[527861]++
//line /snap/go/10455/src/net/dnsclient.go:145
		_go_fuzz_dep_.CoverTab[5276]++
//line /snap/go/10455/src/net/dnsclient.go:145
		// _ = "end of CoverTab[5276]"
//line /snap/go/10455/src/net/dnsclient.go:145
	}
//line /snap/go/10455/src/net/dnsclient.go:145
	// _ = "end of CoverTab[5272]"
//line /snap/go/10455/src/net/dnsclient.go:145
	_go_fuzz_dep_.CoverTab[5273]++
						return s
//line /snap/go/10455/src/net/dnsclient.go:146
	// _ = "end of CoverTab[5273]"
}

// An SRV represents a single DNS SRV record.
type SRV struct {
	Target		string
	Port		uint16
	Priority	uint16
	Weight		uint16
}

// byPriorityWeight sorts SRV records by ascending priority and weight.
type byPriorityWeight []*SRV

func (s byPriorityWeight) Len() int {
//line /snap/go/10455/src/net/dnsclient.go:160
	_go_fuzz_dep_.CoverTab[5277]++
//line /snap/go/10455/src/net/dnsclient.go:160
	return len(s)
//line /snap/go/10455/src/net/dnsclient.go:160
	// _ = "end of CoverTab[5277]"
//line /snap/go/10455/src/net/dnsclient.go:160
}
func (s byPriorityWeight) Less(i, j int) bool {
//line /snap/go/10455/src/net/dnsclient.go:161
	_go_fuzz_dep_.CoverTab[5278]++
						return s[i].Priority < s[j].Priority || func() bool {
//line /snap/go/10455/src/net/dnsclient.go:162
		_go_fuzz_dep_.CoverTab[5279]++
//line /snap/go/10455/src/net/dnsclient.go:162
		return (s[i].Priority == s[j].Priority && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:162
			_go_fuzz_dep_.CoverTab[5280]++
//line /snap/go/10455/src/net/dnsclient.go:162
			return s[i].Weight < s[j].Weight
//line /snap/go/10455/src/net/dnsclient.go:162
			// _ = "end of CoverTab[5280]"
//line /snap/go/10455/src/net/dnsclient.go:162
		}())
//line /snap/go/10455/src/net/dnsclient.go:162
		// _ = "end of CoverTab[5279]"
//line /snap/go/10455/src/net/dnsclient.go:162
	}()
//line /snap/go/10455/src/net/dnsclient.go:162
	// _ = "end of CoverTab[5278]"
}
func (s byPriorityWeight) Swap(i, j int) {
//line /snap/go/10455/src/net/dnsclient.go:164
	_go_fuzz_dep_.CoverTab[5281]++
//line /snap/go/10455/src/net/dnsclient.go:164
	s[i], s[j] = s[j], s[i]
//line /snap/go/10455/src/net/dnsclient.go:164
	// _ = "end of CoverTab[5281]"
//line /snap/go/10455/src/net/dnsclient.go:164
}

// shuffleByWeight shuffles SRV records by weight using the algorithm
//line /snap/go/10455/src/net/dnsclient.go:166
// described in RFC 2782.
//line /snap/go/10455/src/net/dnsclient.go:168
func (addrs byPriorityWeight) shuffleByWeight() {
//line /snap/go/10455/src/net/dnsclient.go:168
	_go_fuzz_dep_.CoverTab[5282]++
						sum := 0
//line /snap/go/10455/src/net/dnsclient.go:169
	_go_fuzz_dep_.CoverTab[786655] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/dnsclient.go:170
		if _go_fuzz_dep_.CoverTab[786655] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:170
			_go_fuzz_dep_.CoverTab[527880]++
//line /snap/go/10455/src/net/dnsclient.go:170
		} else {
//line /snap/go/10455/src/net/dnsclient.go:170
			_go_fuzz_dep_.CoverTab[527881]++
//line /snap/go/10455/src/net/dnsclient.go:170
		}
//line /snap/go/10455/src/net/dnsclient.go:170
		_go_fuzz_dep_.CoverTab[786655] = 1
//line /snap/go/10455/src/net/dnsclient.go:170
		_go_fuzz_dep_.CoverTab[5284]++
							sum += int(addr.Weight)
//line /snap/go/10455/src/net/dnsclient.go:171
		// _ = "end of CoverTab[5284]"
	}
//line /snap/go/10455/src/net/dnsclient.go:172
	if _go_fuzz_dep_.CoverTab[786655] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:172
		_go_fuzz_dep_.CoverTab[527882]++
//line /snap/go/10455/src/net/dnsclient.go:172
	} else {
//line /snap/go/10455/src/net/dnsclient.go:172
		_go_fuzz_dep_.CoverTab[527883]++
//line /snap/go/10455/src/net/dnsclient.go:172
	}
//line /snap/go/10455/src/net/dnsclient.go:172
	// _ = "end of CoverTab[5282]"
//line /snap/go/10455/src/net/dnsclient.go:172
	_go_fuzz_dep_.CoverTab[5283]++
//line /snap/go/10455/src/net/dnsclient.go:172
	_go_fuzz_dep_.CoverTab[786656] = 0
						for sum > 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient.go:173
		_go_fuzz_dep_.CoverTab[5285]++
//line /snap/go/10455/src/net/dnsclient.go:173
		return len(addrs) > 1
//line /snap/go/10455/src/net/dnsclient.go:173
		// _ = "end of CoverTab[5285]"
//line /snap/go/10455/src/net/dnsclient.go:173
	}() {
//line /snap/go/10455/src/net/dnsclient.go:173
		if _go_fuzz_dep_.CoverTab[786656] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:173
			_go_fuzz_dep_.CoverTab[527884]++
//line /snap/go/10455/src/net/dnsclient.go:173
		} else {
//line /snap/go/10455/src/net/dnsclient.go:173
			_go_fuzz_dep_.CoverTab[527885]++
//line /snap/go/10455/src/net/dnsclient.go:173
		}
//line /snap/go/10455/src/net/dnsclient.go:173
		_go_fuzz_dep_.CoverTab[786656] = 1
//line /snap/go/10455/src/net/dnsclient.go:173
		_go_fuzz_dep_.CoverTab[5286]++
							s := 0
							n := randIntn(sum)
//line /snap/go/10455/src/net/dnsclient.go:175
		_go_fuzz_dep_.CoverTab[786657] = 0
							for i := range addrs {
//line /snap/go/10455/src/net/dnsclient.go:176
			if _go_fuzz_dep_.CoverTab[786657] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:176
				_go_fuzz_dep_.CoverTab[527888]++
//line /snap/go/10455/src/net/dnsclient.go:176
			} else {
//line /snap/go/10455/src/net/dnsclient.go:176
				_go_fuzz_dep_.CoverTab[527889]++
//line /snap/go/10455/src/net/dnsclient.go:176
			}
//line /snap/go/10455/src/net/dnsclient.go:176
			_go_fuzz_dep_.CoverTab[786657] = 1
//line /snap/go/10455/src/net/dnsclient.go:176
			_go_fuzz_dep_.CoverTab[5288]++
								s += int(addrs[i].Weight)
								if s > n {
//line /snap/go/10455/src/net/dnsclient.go:178
				_go_fuzz_dep_.CoverTab[527862]++
//line /snap/go/10455/src/net/dnsclient.go:178
				_go_fuzz_dep_.CoverTab[5289]++
									if i > 0 {
//line /snap/go/10455/src/net/dnsclient.go:179
					_go_fuzz_dep_.CoverTab[527864]++
//line /snap/go/10455/src/net/dnsclient.go:179
					_go_fuzz_dep_.CoverTab[5291]++
										addrs[0], addrs[i] = addrs[i], addrs[0]
//line /snap/go/10455/src/net/dnsclient.go:180
					// _ = "end of CoverTab[5291]"
				} else {
//line /snap/go/10455/src/net/dnsclient.go:181
					_go_fuzz_dep_.CoverTab[527865]++
//line /snap/go/10455/src/net/dnsclient.go:181
					_go_fuzz_dep_.CoverTab[5292]++
//line /snap/go/10455/src/net/dnsclient.go:181
					// _ = "end of CoverTab[5292]"
//line /snap/go/10455/src/net/dnsclient.go:181
				}
//line /snap/go/10455/src/net/dnsclient.go:181
				// _ = "end of CoverTab[5289]"
//line /snap/go/10455/src/net/dnsclient.go:181
				_go_fuzz_dep_.CoverTab[5290]++
									break
//line /snap/go/10455/src/net/dnsclient.go:182
				// _ = "end of CoverTab[5290]"
			} else {
//line /snap/go/10455/src/net/dnsclient.go:183
				_go_fuzz_dep_.CoverTab[527863]++
//line /snap/go/10455/src/net/dnsclient.go:183
				_go_fuzz_dep_.CoverTab[5293]++
//line /snap/go/10455/src/net/dnsclient.go:183
				// _ = "end of CoverTab[5293]"
//line /snap/go/10455/src/net/dnsclient.go:183
			}
//line /snap/go/10455/src/net/dnsclient.go:183
			// _ = "end of CoverTab[5288]"
		}
//line /snap/go/10455/src/net/dnsclient.go:184
		if _go_fuzz_dep_.CoverTab[786657] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:184
			_go_fuzz_dep_.CoverTab[527890]++
//line /snap/go/10455/src/net/dnsclient.go:184
		} else {
//line /snap/go/10455/src/net/dnsclient.go:184
			_go_fuzz_dep_.CoverTab[527891]++
//line /snap/go/10455/src/net/dnsclient.go:184
		}
//line /snap/go/10455/src/net/dnsclient.go:184
		// _ = "end of CoverTab[5286]"
//line /snap/go/10455/src/net/dnsclient.go:184
		_go_fuzz_dep_.CoverTab[5287]++
							sum -= int(addrs[0].Weight)
							addrs = addrs[1:]
//line /snap/go/10455/src/net/dnsclient.go:186
		// _ = "end of CoverTab[5287]"
	}
//line /snap/go/10455/src/net/dnsclient.go:187
	if _go_fuzz_dep_.CoverTab[786656] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:187
		_go_fuzz_dep_.CoverTab[527886]++
//line /snap/go/10455/src/net/dnsclient.go:187
	} else {
//line /snap/go/10455/src/net/dnsclient.go:187
		_go_fuzz_dep_.CoverTab[527887]++
//line /snap/go/10455/src/net/dnsclient.go:187
	}
//line /snap/go/10455/src/net/dnsclient.go:187
	// _ = "end of CoverTab[5283]"
}

// sort reorders SRV records as specified in RFC 2782.
func (addrs byPriorityWeight) sort() {
//line /snap/go/10455/src/net/dnsclient.go:191
	_go_fuzz_dep_.CoverTab[5294]++
						sort.Sort(addrs)
						i := 0
//line /snap/go/10455/src/net/dnsclient.go:193
	_go_fuzz_dep_.CoverTab[786658] = 0
						for j := 1; j < len(addrs); j++ {
//line /snap/go/10455/src/net/dnsclient.go:194
		if _go_fuzz_dep_.CoverTab[786658] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:194
			_go_fuzz_dep_.CoverTab[527892]++
//line /snap/go/10455/src/net/dnsclient.go:194
		} else {
//line /snap/go/10455/src/net/dnsclient.go:194
			_go_fuzz_dep_.CoverTab[527893]++
//line /snap/go/10455/src/net/dnsclient.go:194
		}
//line /snap/go/10455/src/net/dnsclient.go:194
		_go_fuzz_dep_.CoverTab[786658] = 1
//line /snap/go/10455/src/net/dnsclient.go:194
		_go_fuzz_dep_.CoverTab[5296]++
							if addrs[i].Priority != addrs[j].Priority {
//line /snap/go/10455/src/net/dnsclient.go:195
			_go_fuzz_dep_.CoverTab[527866]++
//line /snap/go/10455/src/net/dnsclient.go:195
			_go_fuzz_dep_.CoverTab[5297]++
								addrs[i:j].shuffleByWeight()
								i = j
//line /snap/go/10455/src/net/dnsclient.go:197
			// _ = "end of CoverTab[5297]"
		} else {
//line /snap/go/10455/src/net/dnsclient.go:198
			_go_fuzz_dep_.CoverTab[527867]++
//line /snap/go/10455/src/net/dnsclient.go:198
			_go_fuzz_dep_.CoverTab[5298]++
//line /snap/go/10455/src/net/dnsclient.go:198
			// _ = "end of CoverTab[5298]"
//line /snap/go/10455/src/net/dnsclient.go:198
		}
//line /snap/go/10455/src/net/dnsclient.go:198
		// _ = "end of CoverTab[5296]"
	}
//line /snap/go/10455/src/net/dnsclient.go:199
	if _go_fuzz_dep_.CoverTab[786658] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:199
		_go_fuzz_dep_.CoverTab[527894]++
//line /snap/go/10455/src/net/dnsclient.go:199
	} else {
//line /snap/go/10455/src/net/dnsclient.go:199
		_go_fuzz_dep_.CoverTab[527895]++
//line /snap/go/10455/src/net/dnsclient.go:199
	}
//line /snap/go/10455/src/net/dnsclient.go:199
	// _ = "end of CoverTab[5294]"
//line /snap/go/10455/src/net/dnsclient.go:199
	_go_fuzz_dep_.CoverTab[5295]++
						addrs[i:].shuffleByWeight()
//line /snap/go/10455/src/net/dnsclient.go:200
	// _ = "end of CoverTab[5295]"
}

// An MX represents a single DNS MX record.
type MX struct {
	Host	string
	Pref	uint16
}

// byPref implements sort.Interface to sort MX records by preference
type byPref []*MX

func (s byPref) Len() int {
//line /snap/go/10455/src/net/dnsclient.go:212
	_go_fuzz_dep_.CoverTab[5299]++
//line /snap/go/10455/src/net/dnsclient.go:212
	return len(s)
//line /snap/go/10455/src/net/dnsclient.go:212
	// _ = "end of CoverTab[5299]"
//line /snap/go/10455/src/net/dnsclient.go:212
}
func (s byPref) Less(i, j int) bool {
//line /snap/go/10455/src/net/dnsclient.go:213
	_go_fuzz_dep_.CoverTab[5300]++
//line /snap/go/10455/src/net/dnsclient.go:213
	return s[i].Pref < s[j].Pref
//line /snap/go/10455/src/net/dnsclient.go:213
	// _ = "end of CoverTab[5300]"
//line /snap/go/10455/src/net/dnsclient.go:213
}
func (s byPref) Swap(i, j int) {
//line /snap/go/10455/src/net/dnsclient.go:214
	_go_fuzz_dep_.CoverTab[5301]++
//line /snap/go/10455/src/net/dnsclient.go:214
	s[i], s[j] = s[j], s[i]
//line /snap/go/10455/src/net/dnsclient.go:214
	// _ = "end of CoverTab[5301]"
//line /snap/go/10455/src/net/dnsclient.go:214
}

// sort reorders MX records as specified in RFC 5321.
func (s byPref) sort() {
//line /snap/go/10455/src/net/dnsclient.go:217
	_go_fuzz_dep_.CoverTab[5302]++
//line /snap/go/10455/src/net/dnsclient.go:217
	_go_fuzz_dep_.CoverTab[786659] = 0
						for i := range s {
//line /snap/go/10455/src/net/dnsclient.go:218
		if _go_fuzz_dep_.CoverTab[786659] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:218
			_go_fuzz_dep_.CoverTab[527896]++
//line /snap/go/10455/src/net/dnsclient.go:218
		} else {
//line /snap/go/10455/src/net/dnsclient.go:218
			_go_fuzz_dep_.CoverTab[527897]++
//line /snap/go/10455/src/net/dnsclient.go:218
		}
//line /snap/go/10455/src/net/dnsclient.go:218
		_go_fuzz_dep_.CoverTab[786659] = 1
//line /snap/go/10455/src/net/dnsclient.go:218
		_go_fuzz_dep_.CoverTab[5304]++
							j := randIntn(i + 1)
							s[i], s[j] = s[j], s[i]
//line /snap/go/10455/src/net/dnsclient.go:220
		// _ = "end of CoverTab[5304]"
	}
//line /snap/go/10455/src/net/dnsclient.go:221
	if _go_fuzz_dep_.CoverTab[786659] == 0 {
//line /snap/go/10455/src/net/dnsclient.go:221
		_go_fuzz_dep_.CoverTab[527898]++
//line /snap/go/10455/src/net/dnsclient.go:221
	} else {
//line /snap/go/10455/src/net/dnsclient.go:221
		_go_fuzz_dep_.CoverTab[527899]++
//line /snap/go/10455/src/net/dnsclient.go:221
	}
//line /snap/go/10455/src/net/dnsclient.go:221
	// _ = "end of CoverTab[5302]"
//line /snap/go/10455/src/net/dnsclient.go:221
	_go_fuzz_dep_.CoverTab[5303]++
						sort.Sort(s)
//line /snap/go/10455/src/net/dnsclient.go:222
	// _ = "end of CoverTab[5303]"
}

// An NS represents a single DNS NS record.
type NS struct {
	Host string
}

//line /snap/go/10455/src/net/dnsclient.go:228
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/dnsclient.go:228
var _ = _go_fuzz_dep_.CoverTab
