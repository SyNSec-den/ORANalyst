// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/dnsclient.go:5
package net

//line /usr/local/go/src/net/dnsclient.go:5
import (
//line /usr/local/go/src/net/dnsclient.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/dnsclient.go:5
)
//line /usr/local/go/src/net/dnsclient.go:5
import (
//line /usr/local/go/src/net/dnsclient.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/dnsclient.go:5
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
//line /usr/local/go/src/net/dnsclient.go:18
	_go_fuzz_dep_.CoverTab[4843]++
						return int(fastrandu() >> 1)
//line /usr/local/go/src/net/dnsclient.go:19
	// _ = "end of CoverTab[4843]"
}

func randIntn(n int) int {
//line /usr/local/go/src/net/dnsclient.go:22
	_go_fuzz_dep_.CoverTab[4844]++
						return randInt() % n
//line /usr/local/go/src/net/dnsclient.go:23
	// _ = "end of CoverTab[4844]"
}

// reverseaddr returns the in-addr.arpa. or ip6.arpa. hostname of the IP
//line /usr/local/go/src/net/dnsclient.go:26
// address addr suitable for rDNS (PTR) record lookup or an error if it fails
//line /usr/local/go/src/net/dnsclient.go:26
// to parse the IP address.
//line /usr/local/go/src/net/dnsclient.go:29
func reverseaddr(addr string) (arpa string, err error) {
//line /usr/local/go/src/net/dnsclient.go:29
	_go_fuzz_dep_.CoverTab[4845]++
						ip := ParseIP(addr)
						if ip == nil {
//line /usr/local/go/src/net/dnsclient.go:31
		_go_fuzz_dep_.CoverTab[4849]++
							return "", &DNSError{Err: "unrecognized address", Name: addr}
//line /usr/local/go/src/net/dnsclient.go:32
		// _ = "end of CoverTab[4849]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:33
		_go_fuzz_dep_.CoverTab[4850]++
//line /usr/local/go/src/net/dnsclient.go:33
		// _ = "end of CoverTab[4850]"
//line /usr/local/go/src/net/dnsclient.go:33
	}
//line /usr/local/go/src/net/dnsclient.go:33
	// _ = "end of CoverTab[4845]"
//line /usr/local/go/src/net/dnsclient.go:33
	_go_fuzz_dep_.CoverTab[4846]++
						if ip.To4() != nil {
//line /usr/local/go/src/net/dnsclient.go:34
		_go_fuzz_dep_.CoverTab[4851]++
							return itoa.Uitoa(uint(ip[15])) + "." + itoa.Uitoa(uint(ip[14])) + "." + itoa.Uitoa(uint(ip[13])) + "." + itoa.Uitoa(uint(ip[12])) + ".in-addr.arpa.", nil
//line /usr/local/go/src/net/dnsclient.go:35
		// _ = "end of CoverTab[4851]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:36
		_go_fuzz_dep_.CoverTab[4852]++
//line /usr/local/go/src/net/dnsclient.go:36
		// _ = "end of CoverTab[4852]"
//line /usr/local/go/src/net/dnsclient.go:36
	}
//line /usr/local/go/src/net/dnsclient.go:36
	// _ = "end of CoverTab[4846]"
//line /usr/local/go/src/net/dnsclient.go:36
	_go_fuzz_dep_.CoverTab[4847]++

						buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))

						for i := len(ip) - 1; i >= 0; i-- {
//line /usr/local/go/src/net/dnsclient.go:40
		_go_fuzz_dep_.CoverTab[4853]++
							v := ip[i]
							buf = append(buf, hexDigit[v&0xF],
			'.',
			hexDigit[v>>4],
			'.')
//line /usr/local/go/src/net/dnsclient.go:45
		// _ = "end of CoverTab[4853]"
	}
//line /usr/local/go/src/net/dnsclient.go:46
	// _ = "end of CoverTab[4847]"
//line /usr/local/go/src/net/dnsclient.go:46
	_go_fuzz_dep_.CoverTab[4848]++

						buf = append(buf, "ip6.arpa."...)
						return string(buf), nil
//line /usr/local/go/src/net/dnsclient.go:49
	// _ = "end of CoverTab[4848]"
}

func equalASCIIName(x, y dnsmessage.Name) bool {
//line /usr/local/go/src/net/dnsclient.go:52
	_go_fuzz_dep_.CoverTab[4854]++
						if x.Length != y.Length {
//line /usr/local/go/src/net/dnsclient.go:53
		_go_fuzz_dep_.CoverTab[4857]++
							return false
//line /usr/local/go/src/net/dnsclient.go:54
		// _ = "end of CoverTab[4857]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:55
		_go_fuzz_dep_.CoverTab[4858]++
//line /usr/local/go/src/net/dnsclient.go:55
		// _ = "end of CoverTab[4858]"
//line /usr/local/go/src/net/dnsclient.go:55
	}
//line /usr/local/go/src/net/dnsclient.go:55
	// _ = "end of CoverTab[4854]"
//line /usr/local/go/src/net/dnsclient.go:55
	_go_fuzz_dep_.CoverTab[4855]++
						for i := 0; i < int(x.Length); i++ {
//line /usr/local/go/src/net/dnsclient.go:56
		_go_fuzz_dep_.CoverTab[4859]++
							a := x.Data[i]
							b := y.Data[i]
							if 'A' <= a && func() bool {
//line /usr/local/go/src/net/dnsclient.go:59
			_go_fuzz_dep_.CoverTab[4862]++
//line /usr/local/go/src/net/dnsclient.go:59
			return a <= 'Z'
//line /usr/local/go/src/net/dnsclient.go:59
			// _ = "end of CoverTab[4862]"
//line /usr/local/go/src/net/dnsclient.go:59
		}() {
//line /usr/local/go/src/net/dnsclient.go:59
			_go_fuzz_dep_.CoverTab[4863]++
								a += 0x20
//line /usr/local/go/src/net/dnsclient.go:60
			// _ = "end of CoverTab[4863]"
		} else {
//line /usr/local/go/src/net/dnsclient.go:61
			_go_fuzz_dep_.CoverTab[4864]++
//line /usr/local/go/src/net/dnsclient.go:61
			// _ = "end of CoverTab[4864]"
//line /usr/local/go/src/net/dnsclient.go:61
		}
//line /usr/local/go/src/net/dnsclient.go:61
		// _ = "end of CoverTab[4859]"
//line /usr/local/go/src/net/dnsclient.go:61
		_go_fuzz_dep_.CoverTab[4860]++
							if 'A' <= b && func() bool {
//line /usr/local/go/src/net/dnsclient.go:62
			_go_fuzz_dep_.CoverTab[4865]++
//line /usr/local/go/src/net/dnsclient.go:62
			return b <= 'Z'
//line /usr/local/go/src/net/dnsclient.go:62
			// _ = "end of CoverTab[4865]"
//line /usr/local/go/src/net/dnsclient.go:62
		}() {
//line /usr/local/go/src/net/dnsclient.go:62
			_go_fuzz_dep_.CoverTab[4866]++
								b += 0x20
//line /usr/local/go/src/net/dnsclient.go:63
			// _ = "end of CoverTab[4866]"
		} else {
//line /usr/local/go/src/net/dnsclient.go:64
			_go_fuzz_dep_.CoverTab[4867]++
//line /usr/local/go/src/net/dnsclient.go:64
			// _ = "end of CoverTab[4867]"
//line /usr/local/go/src/net/dnsclient.go:64
		}
//line /usr/local/go/src/net/dnsclient.go:64
		// _ = "end of CoverTab[4860]"
//line /usr/local/go/src/net/dnsclient.go:64
		_go_fuzz_dep_.CoverTab[4861]++
							if a != b {
//line /usr/local/go/src/net/dnsclient.go:65
			_go_fuzz_dep_.CoverTab[4868]++
								return false
//line /usr/local/go/src/net/dnsclient.go:66
			// _ = "end of CoverTab[4868]"
		} else {
//line /usr/local/go/src/net/dnsclient.go:67
			_go_fuzz_dep_.CoverTab[4869]++
//line /usr/local/go/src/net/dnsclient.go:67
			// _ = "end of CoverTab[4869]"
//line /usr/local/go/src/net/dnsclient.go:67
		}
//line /usr/local/go/src/net/dnsclient.go:67
		// _ = "end of CoverTab[4861]"
	}
//line /usr/local/go/src/net/dnsclient.go:68
	// _ = "end of CoverTab[4855]"
//line /usr/local/go/src/net/dnsclient.go:68
	_go_fuzz_dep_.CoverTab[4856]++
						return true
//line /usr/local/go/src/net/dnsclient.go:69
	// _ = "end of CoverTab[4856]"
}

// isDomainName checks if a string is a presentation-format domain name
//line /usr/local/go/src/net/dnsclient.go:72
// (currently restricted to hostname-compatible "preferred name" LDH labels and
//line /usr/local/go/src/net/dnsclient.go:72
// SRV-like "underscore labels"; see golang.org/issue/12421).
//line /usr/local/go/src/net/dnsclient.go:75
func isDomainName(s string) bool {
//line /usr/local/go/src/net/dnsclient.go:75
	_go_fuzz_dep_.CoverTab[4870]++

						if s == "." {
//line /usr/local/go/src/net/dnsclient.go:77
		_go_fuzz_dep_.CoverTab[4875]++
							return true
//line /usr/local/go/src/net/dnsclient.go:78
		// _ = "end of CoverTab[4875]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:79
		_go_fuzz_dep_.CoverTab[4876]++
//line /usr/local/go/src/net/dnsclient.go:79
		// _ = "end of CoverTab[4876]"
//line /usr/local/go/src/net/dnsclient.go:79
	}
//line /usr/local/go/src/net/dnsclient.go:79
	// _ = "end of CoverTab[4870]"
//line /usr/local/go/src/net/dnsclient.go:79
	_go_fuzz_dep_.CoverTab[4871]++

//line /usr/local/go/src/net/dnsclient.go:89
	l := len(s)
	if l == 0 || func() bool {
//line /usr/local/go/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[4877]++
//line /usr/local/go/src/net/dnsclient.go:90
		return l > 254
//line /usr/local/go/src/net/dnsclient.go:90
		// _ = "end of CoverTab[4877]"
//line /usr/local/go/src/net/dnsclient.go:90
	}() || func() bool {
//line /usr/local/go/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[4878]++
//line /usr/local/go/src/net/dnsclient.go:90
		return l == 254 && func() bool {
//line /usr/local/go/src/net/dnsclient.go:90
			_go_fuzz_dep_.CoverTab[4879]++
//line /usr/local/go/src/net/dnsclient.go:90
			return s[l-1] != '.'
//line /usr/local/go/src/net/dnsclient.go:90
			// _ = "end of CoverTab[4879]"
//line /usr/local/go/src/net/dnsclient.go:90
		}()
//line /usr/local/go/src/net/dnsclient.go:90
		// _ = "end of CoverTab[4878]"
//line /usr/local/go/src/net/dnsclient.go:90
	}() {
//line /usr/local/go/src/net/dnsclient.go:90
		_go_fuzz_dep_.CoverTab[4880]++
							return false
//line /usr/local/go/src/net/dnsclient.go:91
		// _ = "end of CoverTab[4880]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:92
		_go_fuzz_dep_.CoverTab[4881]++
//line /usr/local/go/src/net/dnsclient.go:92
		// _ = "end of CoverTab[4881]"
//line /usr/local/go/src/net/dnsclient.go:92
	}
//line /usr/local/go/src/net/dnsclient.go:92
	// _ = "end of CoverTab[4871]"
//line /usr/local/go/src/net/dnsclient.go:92
	_go_fuzz_dep_.CoverTab[4872]++

						last := byte('.')
						nonNumeric := false
						partlen := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/dnsclient.go:97
		_go_fuzz_dep_.CoverTab[4882]++
							c := s[i]
							switch {
		default:
//line /usr/local/go/src/net/dnsclient.go:100
			_go_fuzz_dep_.CoverTab[4884]++
								return false
//line /usr/local/go/src/net/dnsclient.go:101
			// _ = "end of CoverTab[4884]"
		case 'a' <= c && func() bool {
//line /usr/local/go/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[4892]++
//line /usr/local/go/src/net/dnsclient.go:102
			return c <= 'z'
//line /usr/local/go/src/net/dnsclient.go:102
			// _ = "end of CoverTab[4892]"
//line /usr/local/go/src/net/dnsclient.go:102
		}() || func() bool {
//line /usr/local/go/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[4893]++
//line /usr/local/go/src/net/dnsclient.go:102
			return 'A' <= c && func() bool {
//line /usr/local/go/src/net/dnsclient.go:102
				_go_fuzz_dep_.CoverTab[4894]++
//line /usr/local/go/src/net/dnsclient.go:102
				return c <= 'Z'
//line /usr/local/go/src/net/dnsclient.go:102
				// _ = "end of CoverTab[4894]"
//line /usr/local/go/src/net/dnsclient.go:102
			}()
//line /usr/local/go/src/net/dnsclient.go:102
			// _ = "end of CoverTab[4893]"
//line /usr/local/go/src/net/dnsclient.go:102
		}() || func() bool {
//line /usr/local/go/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[4895]++
//line /usr/local/go/src/net/dnsclient.go:102
			return c == '_'
//line /usr/local/go/src/net/dnsclient.go:102
			// _ = "end of CoverTab[4895]"
//line /usr/local/go/src/net/dnsclient.go:102
		}():
//line /usr/local/go/src/net/dnsclient.go:102
			_go_fuzz_dep_.CoverTab[4885]++
								nonNumeric = true
								partlen++
//line /usr/local/go/src/net/dnsclient.go:104
			// _ = "end of CoverTab[4885]"
		case '0' <= c && func() bool {
//line /usr/local/go/src/net/dnsclient.go:105
			_go_fuzz_dep_.CoverTab[4896]++
//line /usr/local/go/src/net/dnsclient.go:105
			return c <= '9'
//line /usr/local/go/src/net/dnsclient.go:105
			// _ = "end of CoverTab[4896]"
//line /usr/local/go/src/net/dnsclient.go:105
		}():
//line /usr/local/go/src/net/dnsclient.go:105
			_go_fuzz_dep_.CoverTab[4886]++

								partlen++
//line /usr/local/go/src/net/dnsclient.go:107
			// _ = "end of CoverTab[4886]"
		case c == '-':
//line /usr/local/go/src/net/dnsclient.go:108
			_go_fuzz_dep_.CoverTab[4887]++

								if last == '.' {
//line /usr/local/go/src/net/dnsclient.go:110
				_go_fuzz_dep_.CoverTab[4897]++
									return false
//line /usr/local/go/src/net/dnsclient.go:111
				// _ = "end of CoverTab[4897]"
			} else {
//line /usr/local/go/src/net/dnsclient.go:112
				_go_fuzz_dep_.CoverTab[4898]++
//line /usr/local/go/src/net/dnsclient.go:112
				// _ = "end of CoverTab[4898]"
//line /usr/local/go/src/net/dnsclient.go:112
			}
//line /usr/local/go/src/net/dnsclient.go:112
			// _ = "end of CoverTab[4887]"
//line /usr/local/go/src/net/dnsclient.go:112
			_go_fuzz_dep_.CoverTab[4888]++
								partlen++
								nonNumeric = true
//line /usr/local/go/src/net/dnsclient.go:114
			// _ = "end of CoverTab[4888]"
		case c == '.':
//line /usr/local/go/src/net/dnsclient.go:115
			_go_fuzz_dep_.CoverTab[4889]++

								if last == '.' || func() bool {
//line /usr/local/go/src/net/dnsclient.go:117
				_go_fuzz_dep_.CoverTab[4899]++
//line /usr/local/go/src/net/dnsclient.go:117
				return last == '-'
//line /usr/local/go/src/net/dnsclient.go:117
				// _ = "end of CoverTab[4899]"
//line /usr/local/go/src/net/dnsclient.go:117
			}() {
//line /usr/local/go/src/net/dnsclient.go:117
				_go_fuzz_dep_.CoverTab[4900]++
									return false
//line /usr/local/go/src/net/dnsclient.go:118
				// _ = "end of CoverTab[4900]"
			} else {
//line /usr/local/go/src/net/dnsclient.go:119
				_go_fuzz_dep_.CoverTab[4901]++
//line /usr/local/go/src/net/dnsclient.go:119
				// _ = "end of CoverTab[4901]"
//line /usr/local/go/src/net/dnsclient.go:119
			}
//line /usr/local/go/src/net/dnsclient.go:119
			// _ = "end of CoverTab[4889]"
//line /usr/local/go/src/net/dnsclient.go:119
			_go_fuzz_dep_.CoverTab[4890]++
								if partlen > 63 || func() bool {
//line /usr/local/go/src/net/dnsclient.go:120
				_go_fuzz_dep_.CoverTab[4902]++
//line /usr/local/go/src/net/dnsclient.go:120
				return partlen == 0
//line /usr/local/go/src/net/dnsclient.go:120
				// _ = "end of CoverTab[4902]"
//line /usr/local/go/src/net/dnsclient.go:120
			}() {
//line /usr/local/go/src/net/dnsclient.go:120
				_go_fuzz_dep_.CoverTab[4903]++
									return false
//line /usr/local/go/src/net/dnsclient.go:121
				// _ = "end of CoverTab[4903]"
			} else {
//line /usr/local/go/src/net/dnsclient.go:122
				_go_fuzz_dep_.CoverTab[4904]++
//line /usr/local/go/src/net/dnsclient.go:122
				// _ = "end of CoverTab[4904]"
//line /usr/local/go/src/net/dnsclient.go:122
			}
//line /usr/local/go/src/net/dnsclient.go:122
			// _ = "end of CoverTab[4890]"
//line /usr/local/go/src/net/dnsclient.go:122
			_go_fuzz_dep_.CoverTab[4891]++
								partlen = 0
//line /usr/local/go/src/net/dnsclient.go:123
			// _ = "end of CoverTab[4891]"
		}
//line /usr/local/go/src/net/dnsclient.go:124
		// _ = "end of CoverTab[4882]"
//line /usr/local/go/src/net/dnsclient.go:124
		_go_fuzz_dep_.CoverTab[4883]++
							last = c
//line /usr/local/go/src/net/dnsclient.go:125
		// _ = "end of CoverTab[4883]"
	}
//line /usr/local/go/src/net/dnsclient.go:126
	// _ = "end of CoverTab[4872]"
//line /usr/local/go/src/net/dnsclient.go:126
	_go_fuzz_dep_.CoverTab[4873]++
						if last == '-' || func() bool {
//line /usr/local/go/src/net/dnsclient.go:127
		_go_fuzz_dep_.CoverTab[4905]++
//line /usr/local/go/src/net/dnsclient.go:127
		return partlen > 63
//line /usr/local/go/src/net/dnsclient.go:127
		// _ = "end of CoverTab[4905]"
//line /usr/local/go/src/net/dnsclient.go:127
	}() {
//line /usr/local/go/src/net/dnsclient.go:127
		_go_fuzz_dep_.CoverTab[4906]++
							return false
//line /usr/local/go/src/net/dnsclient.go:128
		// _ = "end of CoverTab[4906]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:129
		_go_fuzz_dep_.CoverTab[4907]++
//line /usr/local/go/src/net/dnsclient.go:129
		// _ = "end of CoverTab[4907]"
//line /usr/local/go/src/net/dnsclient.go:129
	}
//line /usr/local/go/src/net/dnsclient.go:129
	// _ = "end of CoverTab[4873]"
//line /usr/local/go/src/net/dnsclient.go:129
	_go_fuzz_dep_.CoverTab[4874]++

						return nonNumeric
//line /usr/local/go/src/net/dnsclient.go:131
	// _ = "end of CoverTab[4874]"
}

// absDomainName returns an absolute domain name which ends with a
//line /usr/local/go/src/net/dnsclient.go:134
// trailing dot to match pure Go reverse resolver and all other lookup
//line /usr/local/go/src/net/dnsclient.go:134
// routines.
//line /usr/local/go/src/net/dnsclient.go:134
// See golang.org/issue/12189.
//line /usr/local/go/src/net/dnsclient.go:134
// But we don't want to add dots for local names from /etc/hosts.
//line /usr/local/go/src/net/dnsclient.go:134
// It's hard to tell so we settle on the heuristic that names without dots
//line /usr/local/go/src/net/dnsclient.go:134
// (like "localhost" or "myhost") do not get trailing dots, but any other
//line /usr/local/go/src/net/dnsclient.go:134
// names do.
//line /usr/local/go/src/net/dnsclient.go:142
func absDomainName(s string) string {
//line /usr/local/go/src/net/dnsclient.go:142
	_go_fuzz_dep_.CoverTab[4908]++
						if bytealg.IndexByteString(s, '.') != -1 && func() bool {
//line /usr/local/go/src/net/dnsclient.go:143
		_go_fuzz_dep_.CoverTab[4910]++
//line /usr/local/go/src/net/dnsclient.go:143
		return s[len(s)-1] != '.'
//line /usr/local/go/src/net/dnsclient.go:143
		// _ = "end of CoverTab[4910]"
//line /usr/local/go/src/net/dnsclient.go:143
	}() {
//line /usr/local/go/src/net/dnsclient.go:143
		_go_fuzz_dep_.CoverTab[4911]++
							s += "."
//line /usr/local/go/src/net/dnsclient.go:144
		// _ = "end of CoverTab[4911]"
	} else {
//line /usr/local/go/src/net/dnsclient.go:145
		_go_fuzz_dep_.CoverTab[4912]++
//line /usr/local/go/src/net/dnsclient.go:145
		// _ = "end of CoverTab[4912]"
//line /usr/local/go/src/net/dnsclient.go:145
	}
//line /usr/local/go/src/net/dnsclient.go:145
	// _ = "end of CoverTab[4908]"
//line /usr/local/go/src/net/dnsclient.go:145
	_go_fuzz_dep_.CoverTab[4909]++
						return s
//line /usr/local/go/src/net/dnsclient.go:146
	// _ = "end of CoverTab[4909]"
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
//line /usr/local/go/src/net/dnsclient.go:160
	_go_fuzz_dep_.CoverTab[4913]++
//line /usr/local/go/src/net/dnsclient.go:160
	return len(s)
//line /usr/local/go/src/net/dnsclient.go:160
	// _ = "end of CoverTab[4913]"
//line /usr/local/go/src/net/dnsclient.go:160
}
func (s byPriorityWeight) Less(i, j int) bool {
//line /usr/local/go/src/net/dnsclient.go:161
	_go_fuzz_dep_.CoverTab[4914]++
						return s[i].Priority < s[j].Priority || func() bool {
//line /usr/local/go/src/net/dnsclient.go:162
		_go_fuzz_dep_.CoverTab[4915]++
//line /usr/local/go/src/net/dnsclient.go:162
		return (s[i].Priority == s[j].Priority && func() bool {
//line /usr/local/go/src/net/dnsclient.go:162
			_go_fuzz_dep_.CoverTab[4916]++
//line /usr/local/go/src/net/dnsclient.go:162
			return s[i].Weight < s[j].Weight
//line /usr/local/go/src/net/dnsclient.go:162
			// _ = "end of CoverTab[4916]"
//line /usr/local/go/src/net/dnsclient.go:162
		}())
//line /usr/local/go/src/net/dnsclient.go:162
		// _ = "end of CoverTab[4915]"
//line /usr/local/go/src/net/dnsclient.go:162
	}()
//line /usr/local/go/src/net/dnsclient.go:162
	// _ = "end of CoverTab[4914]"
}
func (s byPriorityWeight) Swap(i, j int) {
//line /usr/local/go/src/net/dnsclient.go:164
	_go_fuzz_dep_.CoverTab[4917]++
//line /usr/local/go/src/net/dnsclient.go:164
	s[i], s[j] = s[j], s[i]
//line /usr/local/go/src/net/dnsclient.go:164
	// _ = "end of CoverTab[4917]"
//line /usr/local/go/src/net/dnsclient.go:164
}

// shuffleByWeight shuffles SRV records by weight using the algorithm
//line /usr/local/go/src/net/dnsclient.go:166
// described in RFC 2782.
//line /usr/local/go/src/net/dnsclient.go:168
func (addrs byPriorityWeight) shuffleByWeight() {
//line /usr/local/go/src/net/dnsclient.go:168
	_go_fuzz_dep_.CoverTab[4918]++
						sum := 0
						for _, addr := range addrs {
//line /usr/local/go/src/net/dnsclient.go:170
		_go_fuzz_dep_.CoverTab[4920]++
							sum += int(addr.Weight)
//line /usr/local/go/src/net/dnsclient.go:171
		// _ = "end of CoverTab[4920]"
	}
//line /usr/local/go/src/net/dnsclient.go:172
	// _ = "end of CoverTab[4918]"
//line /usr/local/go/src/net/dnsclient.go:172
	_go_fuzz_dep_.CoverTab[4919]++
						for sum > 0 && func() bool {
//line /usr/local/go/src/net/dnsclient.go:173
		_go_fuzz_dep_.CoverTab[4921]++
//line /usr/local/go/src/net/dnsclient.go:173
		return len(addrs) > 1
//line /usr/local/go/src/net/dnsclient.go:173
		// _ = "end of CoverTab[4921]"
//line /usr/local/go/src/net/dnsclient.go:173
	}() {
//line /usr/local/go/src/net/dnsclient.go:173
		_go_fuzz_dep_.CoverTab[4922]++
							s := 0
							n := randIntn(sum)
							for i := range addrs {
//line /usr/local/go/src/net/dnsclient.go:176
			_go_fuzz_dep_.CoverTab[4924]++
								s += int(addrs[i].Weight)
								if s > n {
//line /usr/local/go/src/net/dnsclient.go:178
				_go_fuzz_dep_.CoverTab[4925]++
									if i > 0 {
//line /usr/local/go/src/net/dnsclient.go:179
					_go_fuzz_dep_.CoverTab[4927]++
										addrs[0], addrs[i] = addrs[i], addrs[0]
//line /usr/local/go/src/net/dnsclient.go:180
					// _ = "end of CoverTab[4927]"
				} else {
//line /usr/local/go/src/net/dnsclient.go:181
					_go_fuzz_dep_.CoverTab[4928]++
//line /usr/local/go/src/net/dnsclient.go:181
					// _ = "end of CoverTab[4928]"
//line /usr/local/go/src/net/dnsclient.go:181
				}
//line /usr/local/go/src/net/dnsclient.go:181
				// _ = "end of CoverTab[4925]"
//line /usr/local/go/src/net/dnsclient.go:181
				_go_fuzz_dep_.CoverTab[4926]++
									break
//line /usr/local/go/src/net/dnsclient.go:182
				// _ = "end of CoverTab[4926]"
			} else {
//line /usr/local/go/src/net/dnsclient.go:183
				_go_fuzz_dep_.CoverTab[4929]++
//line /usr/local/go/src/net/dnsclient.go:183
				// _ = "end of CoverTab[4929]"
//line /usr/local/go/src/net/dnsclient.go:183
			}
//line /usr/local/go/src/net/dnsclient.go:183
			// _ = "end of CoverTab[4924]"
		}
//line /usr/local/go/src/net/dnsclient.go:184
		// _ = "end of CoverTab[4922]"
//line /usr/local/go/src/net/dnsclient.go:184
		_go_fuzz_dep_.CoverTab[4923]++
							sum -= int(addrs[0].Weight)
							addrs = addrs[1:]
//line /usr/local/go/src/net/dnsclient.go:186
		// _ = "end of CoverTab[4923]"
	}
//line /usr/local/go/src/net/dnsclient.go:187
	// _ = "end of CoverTab[4919]"
}

// sort reorders SRV records as specified in RFC 2782.
func (addrs byPriorityWeight) sort() {
//line /usr/local/go/src/net/dnsclient.go:191
	_go_fuzz_dep_.CoverTab[4930]++
						sort.Sort(addrs)
						i := 0
						for j := 1; j < len(addrs); j++ {
//line /usr/local/go/src/net/dnsclient.go:194
		_go_fuzz_dep_.CoverTab[4932]++
							if addrs[i].Priority != addrs[j].Priority {
//line /usr/local/go/src/net/dnsclient.go:195
			_go_fuzz_dep_.CoverTab[4933]++
								addrs[i:j].shuffleByWeight()
								i = j
//line /usr/local/go/src/net/dnsclient.go:197
			// _ = "end of CoverTab[4933]"
		} else {
//line /usr/local/go/src/net/dnsclient.go:198
			_go_fuzz_dep_.CoverTab[4934]++
//line /usr/local/go/src/net/dnsclient.go:198
			// _ = "end of CoverTab[4934]"
//line /usr/local/go/src/net/dnsclient.go:198
		}
//line /usr/local/go/src/net/dnsclient.go:198
		// _ = "end of CoverTab[4932]"
	}
//line /usr/local/go/src/net/dnsclient.go:199
	// _ = "end of CoverTab[4930]"
//line /usr/local/go/src/net/dnsclient.go:199
	_go_fuzz_dep_.CoverTab[4931]++
						addrs[i:].shuffleByWeight()
//line /usr/local/go/src/net/dnsclient.go:200
	// _ = "end of CoverTab[4931]"
}

// An MX represents a single DNS MX record.
type MX struct {
	Host	string
	Pref	uint16
}

// byPref implements sort.Interface to sort MX records by preference
type byPref []*MX

func (s byPref) Len() int {
//line /usr/local/go/src/net/dnsclient.go:212
	_go_fuzz_dep_.CoverTab[4935]++
//line /usr/local/go/src/net/dnsclient.go:212
	return len(s)
//line /usr/local/go/src/net/dnsclient.go:212
	// _ = "end of CoverTab[4935]"
//line /usr/local/go/src/net/dnsclient.go:212
}
func (s byPref) Less(i, j int) bool {
//line /usr/local/go/src/net/dnsclient.go:213
	_go_fuzz_dep_.CoverTab[4936]++
//line /usr/local/go/src/net/dnsclient.go:213
	return s[i].Pref < s[j].Pref
//line /usr/local/go/src/net/dnsclient.go:213
	// _ = "end of CoverTab[4936]"
//line /usr/local/go/src/net/dnsclient.go:213
}
func (s byPref) Swap(i, j int) {
//line /usr/local/go/src/net/dnsclient.go:214
	_go_fuzz_dep_.CoverTab[4937]++
//line /usr/local/go/src/net/dnsclient.go:214
	s[i], s[j] = s[j], s[i]
//line /usr/local/go/src/net/dnsclient.go:214
	// _ = "end of CoverTab[4937]"
//line /usr/local/go/src/net/dnsclient.go:214
}

// sort reorders MX records as specified in RFC 5321.
func (s byPref) sort() {
//line /usr/local/go/src/net/dnsclient.go:217
	_go_fuzz_dep_.CoverTab[4938]++
						for i := range s {
//line /usr/local/go/src/net/dnsclient.go:218
		_go_fuzz_dep_.CoverTab[4940]++
							j := randIntn(i + 1)
							s[i], s[j] = s[j], s[i]
//line /usr/local/go/src/net/dnsclient.go:220
		// _ = "end of CoverTab[4940]"
	}
//line /usr/local/go/src/net/dnsclient.go:221
	// _ = "end of CoverTab[4938]"
//line /usr/local/go/src/net/dnsclient.go:221
	_go_fuzz_dep_.CoverTab[4939]++
						sort.Sort(s)
//line /usr/local/go/src/net/dnsclient.go:222
	// _ = "end of CoverTab[4939]"
}

// An NS represents a single DNS NS record.
type NS struct {
	Host string
}

//line /usr/local/go/src/net/dnsclient.go:228
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dnsclient.go:228
var _ = _go_fuzz_dep_.CoverTab
