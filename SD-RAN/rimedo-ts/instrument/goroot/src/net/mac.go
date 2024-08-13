// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/mac.go:5
package net

//line /usr/local/go/src/net/mac.go:5
import (
//line /usr/local/go/src/net/mac.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/mac.go:5
)
//line /usr/local/go/src/net/mac.go:5
import (
//line /usr/local/go/src/net/mac.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/mac.go:5
)

const hexDigit = "0123456789abcdef"

// A HardwareAddr represents a physical hardware address.
type HardwareAddr []byte

func (a HardwareAddr) String() string {
//line /usr/local/go/src/net/mac.go:12
	_go_fuzz_dep_.CoverTab[15457]++
					if len(a) == 0 {
//line /usr/local/go/src/net/mac.go:13
		_go_fuzz_dep_.CoverTab[15460]++
						return ""
//line /usr/local/go/src/net/mac.go:14
		// _ = "end of CoverTab[15460]"
	} else {
//line /usr/local/go/src/net/mac.go:15
		_go_fuzz_dep_.CoverTab[15461]++
//line /usr/local/go/src/net/mac.go:15
		// _ = "end of CoverTab[15461]"
//line /usr/local/go/src/net/mac.go:15
	}
//line /usr/local/go/src/net/mac.go:15
	// _ = "end of CoverTab[15457]"
//line /usr/local/go/src/net/mac.go:15
	_go_fuzz_dep_.CoverTab[15458]++
					buf := make([]byte, 0, len(a)*3-1)
					for i, b := range a {
//line /usr/local/go/src/net/mac.go:17
		_go_fuzz_dep_.CoverTab[15462]++
						if i > 0 {
//line /usr/local/go/src/net/mac.go:18
			_go_fuzz_dep_.CoverTab[15464]++
							buf = append(buf, ':')
//line /usr/local/go/src/net/mac.go:19
			// _ = "end of CoverTab[15464]"
		} else {
//line /usr/local/go/src/net/mac.go:20
			_go_fuzz_dep_.CoverTab[15465]++
//line /usr/local/go/src/net/mac.go:20
			// _ = "end of CoverTab[15465]"
//line /usr/local/go/src/net/mac.go:20
		}
//line /usr/local/go/src/net/mac.go:20
		// _ = "end of CoverTab[15462]"
//line /usr/local/go/src/net/mac.go:20
		_go_fuzz_dep_.CoverTab[15463]++
						buf = append(buf, hexDigit[b>>4])
						buf = append(buf, hexDigit[b&0xF])
//line /usr/local/go/src/net/mac.go:22
		// _ = "end of CoverTab[15463]"
	}
//line /usr/local/go/src/net/mac.go:23
	// _ = "end of CoverTab[15458]"
//line /usr/local/go/src/net/mac.go:23
	_go_fuzz_dep_.CoverTab[15459]++
					return string(buf)
//line /usr/local/go/src/net/mac.go:24
	// _ = "end of CoverTab[15459]"
}

// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
//line /usr/local/go/src/net/mac.go:27
// IP over InfiniBand link-layer address using one of the following formats:
//line /usr/local/go/src/net/mac.go:27
//
//line /usr/local/go/src/net/mac.go:27
//	00:00:5e:00:53:01
//line /usr/local/go/src/net/mac.go:27
//	02:00:5e:10:00:00:00:01
//line /usr/local/go/src/net/mac.go:27
//	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
//line /usr/local/go/src/net/mac.go:27
//	00-00-5e-00-53-01
//line /usr/local/go/src/net/mac.go:27
//	02-00-5e-10-00-00-00-01
//line /usr/local/go/src/net/mac.go:27
//	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
//line /usr/local/go/src/net/mac.go:27
//	0000.5e00.5301
//line /usr/local/go/src/net/mac.go:27
//	0200.5e10.0000.0001
//line /usr/local/go/src/net/mac.go:27
//	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
//line /usr/local/go/src/net/mac.go:39
func ParseMAC(s string) (hw HardwareAddr, err error) {
//line /usr/local/go/src/net/mac.go:39
	_go_fuzz_dep_.CoverTab[15466]++
					if len(s) < 14 {
//line /usr/local/go/src/net/mac.go:40
		_go_fuzz_dep_.CoverTab[15469]++
						goto error
//line /usr/local/go/src/net/mac.go:41
		// _ = "end of CoverTab[15469]"
	} else {
//line /usr/local/go/src/net/mac.go:42
		_go_fuzz_dep_.CoverTab[15470]++
//line /usr/local/go/src/net/mac.go:42
		// _ = "end of CoverTab[15470]"
//line /usr/local/go/src/net/mac.go:42
	}
//line /usr/local/go/src/net/mac.go:42
	// _ = "end of CoverTab[15466]"
//line /usr/local/go/src/net/mac.go:42
	_go_fuzz_dep_.CoverTab[15467]++

					if s[2] == ':' || func() bool {
//line /usr/local/go/src/net/mac.go:44
		_go_fuzz_dep_.CoverTab[15471]++
//line /usr/local/go/src/net/mac.go:44
		return s[2] == '-'
//line /usr/local/go/src/net/mac.go:44
		// _ = "end of CoverTab[15471]"
//line /usr/local/go/src/net/mac.go:44
	}() {
//line /usr/local/go/src/net/mac.go:44
		_go_fuzz_dep_.CoverTab[15472]++
						if (len(s)+1)%3 != 0 {
//line /usr/local/go/src/net/mac.go:45
			_go_fuzz_dep_.CoverTab[15475]++
							goto error
//line /usr/local/go/src/net/mac.go:46
			// _ = "end of CoverTab[15475]"
		} else {
//line /usr/local/go/src/net/mac.go:47
			_go_fuzz_dep_.CoverTab[15476]++
//line /usr/local/go/src/net/mac.go:47
			// _ = "end of CoverTab[15476]"
//line /usr/local/go/src/net/mac.go:47
		}
//line /usr/local/go/src/net/mac.go:47
		// _ = "end of CoverTab[15472]"
//line /usr/local/go/src/net/mac.go:47
		_go_fuzz_dep_.CoverTab[15473]++
						n := (len(s) + 1) / 3
						if n != 6 && func() bool {
//line /usr/local/go/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[15477]++
//line /usr/local/go/src/net/mac.go:49
			return n != 8
//line /usr/local/go/src/net/mac.go:49
			// _ = "end of CoverTab[15477]"
//line /usr/local/go/src/net/mac.go:49
		}() && func() bool {
//line /usr/local/go/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[15478]++
//line /usr/local/go/src/net/mac.go:49
			return n != 20
//line /usr/local/go/src/net/mac.go:49
			// _ = "end of CoverTab[15478]"
//line /usr/local/go/src/net/mac.go:49
		}() {
//line /usr/local/go/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[15479]++
							goto error
//line /usr/local/go/src/net/mac.go:50
			// _ = "end of CoverTab[15479]"
		} else {
//line /usr/local/go/src/net/mac.go:51
			_go_fuzz_dep_.CoverTab[15480]++
//line /usr/local/go/src/net/mac.go:51
			// _ = "end of CoverTab[15480]"
//line /usr/local/go/src/net/mac.go:51
		}
//line /usr/local/go/src/net/mac.go:51
		// _ = "end of CoverTab[15473]"
//line /usr/local/go/src/net/mac.go:51
		_go_fuzz_dep_.CoverTab[15474]++
						hw = make(HardwareAddr, n)
						for x, i := 0, 0; i < n; i++ {
//line /usr/local/go/src/net/mac.go:53
			_go_fuzz_dep_.CoverTab[15481]++
							var ok bool
							if hw[i], ok = xtoi2(s[x:], s[2]); !ok {
//line /usr/local/go/src/net/mac.go:55
				_go_fuzz_dep_.CoverTab[15483]++
								goto error
//line /usr/local/go/src/net/mac.go:56
				// _ = "end of CoverTab[15483]"
			} else {
//line /usr/local/go/src/net/mac.go:57
				_go_fuzz_dep_.CoverTab[15484]++
//line /usr/local/go/src/net/mac.go:57
				// _ = "end of CoverTab[15484]"
//line /usr/local/go/src/net/mac.go:57
			}
//line /usr/local/go/src/net/mac.go:57
			// _ = "end of CoverTab[15481]"
//line /usr/local/go/src/net/mac.go:57
			_go_fuzz_dep_.CoverTab[15482]++
							x += 3
//line /usr/local/go/src/net/mac.go:58
			// _ = "end of CoverTab[15482]"
		}
//line /usr/local/go/src/net/mac.go:59
		// _ = "end of CoverTab[15474]"
	} else {
//line /usr/local/go/src/net/mac.go:60
		_go_fuzz_dep_.CoverTab[15485]++
//line /usr/local/go/src/net/mac.go:60
		if s[4] == '.' {
//line /usr/local/go/src/net/mac.go:60
			_go_fuzz_dep_.CoverTab[15486]++
							if (len(s)+1)%5 != 0 {
//line /usr/local/go/src/net/mac.go:61
				_go_fuzz_dep_.CoverTab[15489]++
								goto error
//line /usr/local/go/src/net/mac.go:62
				// _ = "end of CoverTab[15489]"
			} else {
//line /usr/local/go/src/net/mac.go:63
				_go_fuzz_dep_.CoverTab[15490]++
//line /usr/local/go/src/net/mac.go:63
				// _ = "end of CoverTab[15490]"
//line /usr/local/go/src/net/mac.go:63
			}
//line /usr/local/go/src/net/mac.go:63
			// _ = "end of CoverTab[15486]"
//line /usr/local/go/src/net/mac.go:63
			_go_fuzz_dep_.CoverTab[15487]++
							n := 2 * (len(s) + 1) / 5
							if n != 6 && func() bool {
//line /usr/local/go/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[15491]++
//line /usr/local/go/src/net/mac.go:65
				return n != 8
//line /usr/local/go/src/net/mac.go:65
				// _ = "end of CoverTab[15491]"
//line /usr/local/go/src/net/mac.go:65
			}() && func() bool {
//line /usr/local/go/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[15492]++
//line /usr/local/go/src/net/mac.go:65
				return n != 20
//line /usr/local/go/src/net/mac.go:65
				// _ = "end of CoverTab[15492]"
//line /usr/local/go/src/net/mac.go:65
			}() {
//line /usr/local/go/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[15493]++
								goto error
//line /usr/local/go/src/net/mac.go:66
				// _ = "end of CoverTab[15493]"
			} else {
//line /usr/local/go/src/net/mac.go:67
				_go_fuzz_dep_.CoverTab[15494]++
//line /usr/local/go/src/net/mac.go:67
				// _ = "end of CoverTab[15494]"
//line /usr/local/go/src/net/mac.go:67
			}
//line /usr/local/go/src/net/mac.go:67
			// _ = "end of CoverTab[15487]"
//line /usr/local/go/src/net/mac.go:67
			_go_fuzz_dep_.CoverTab[15488]++
							hw = make(HardwareAddr, n)
							for x, i := 0, 0; i < n; i += 2 {
//line /usr/local/go/src/net/mac.go:69
				_go_fuzz_dep_.CoverTab[15495]++
								var ok bool
								if hw[i], ok = xtoi2(s[x:x+2], 0); !ok {
//line /usr/local/go/src/net/mac.go:71
					_go_fuzz_dep_.CoverTab[15498]++
									goto error
//line /usr/local/go/src/net/mac.go:72
					// _ = "end of CoverTab[15498]"
				} else {
//line /usr/local/go/src/net/mac.go:73
					_go_fuzz_dep_.CoverTab[15499]++
//line /usr/local/go/src/net/mac.go:73
					// _ = "end of CoverTab[15499]"
//line /usr/local/go/src/net/mac.go:73
				}
//line /usr/local/go/src/net/mac.go:73
				// _ = "end of CoverTab[15495]"
//line /usr/local/go/src/net/mac.go:73
				_go_fuzz_dep_.CoverTab[15496]++
								if hw[i+1], ok = xtoi2(s[x+2:], s[4]); !ok {
//line /usr/local/go/src/net/mac.go:74
					_go_fuzz_dep_.CoverTab[15500]++
									goto error
//line /usr/local/go/src/net/mac.go:75
					// _ = "end of CoverTab[15500]"
				} else {
//line /usr/local/go/src/net/mac.go:76
					_go_fuzz_dep_.CoverTab[15501]++
//line /usr/local/go/src/net/mac.go:76
					// _ = "end of CoverTab[15501]"
//line /usr/local/go/src/net/mac.go:76
				}
//line /usr/local/go/src/net/mac.go:76
				// _ = "end of CoverTab[15496]"
//line /usr/local/go/src/net/mac.go:76
				_go_fuzz_dep_.CoverTab[15497]++
								x += 5
//line /usr/local/go/src/net/mac.go:77
				// _ = "end of CoverTab[15497]"
			}
//line /usr/local/go/src/net/mac.go:78
			// _ = "end of CoverTab[15488]"
		} else {
//line /usr/local/go/src/net/mac.go:79
			_go_fuzz_dep_.CoverTab[15502]++
							goto error
//line /usr/local/go/src/net/mac.go:80
			// _ = "end of CoverTab[15502]"
		}
//line /usr/local/go/src/net/mac.go:81
		// _ = "end of CoverTab[15485]"
//line /usr/local/go/src/net/mac.go:81
	}
//line /usr/local/go/src/net/mac.go:81
	// _ = "end of CoverTab[15467]"
//line /usr/local/go/src/net/mac.go:81
	_go_fuzz_dep_.CoverTab[15468]++
					return hw, nil

error:
					return nil, &AddrError{Err: "invalid MAC address", Addr: s}
//line /usr/local/go/src/net/mac.go:85
	// _ = "end of CoverTab[15468]"
}

//line /usr/local/go/src/net/mac.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/mac.go:86
var _ = _go_fuzz_dep_.CoverTab
