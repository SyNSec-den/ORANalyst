// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/mac.go:5
package net

//line /snap/go/10455/src/net/mac.go:5
import (
//line /snap/go/10455/src/net/mac.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/mac.go:5
)
//line /snap/go/10455/src/net/mac.go:5
import (
//line /snap/go/10455/src/net/mac.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/mac.go:5
)

const hexDigit = "0123456789abcdef"

// A HardwareAddr represents a physical hardware address.
type HardwareAddr []byte

func (a HardwareAddr) String() string {
//line /snap/go/10455/src/net/mac.go:12
	_go_fuzz_dep_.CoverTab[7319]++
						if len(a) == 0 {
//line /snap/go/10455/src/net/mac.go:13
		_go_fuzz_dep_.CoverTab[529190]++
//line /snap/go/10455/src/net/mac.go:13
		_go_fuzz_dep_.CoverTab[7322]++
							return ""
//line /snap/go/10455/src/net/mac.go:14
		// _ = "end of CoverTab[7322]"
	} else {
//line /snap/go/10455/src/net/mac.go:15
		_go_fuzz_dep_.CoverTab[529191]++
//line /snap/go/10455/src/net/mac.go:15
		_go_fuzz_dep_.CoverTab[7323]++
//line /snap/go/10455/src/net/mac.go:15
		// _ = "end of CoverTab[7323]"
//line /snap/go/10455/src/net/mac.go:15
	}
//line /snap/go/10455/src/net/mac.go:15
	// _ = "end of CoverTab[7319]"
//line /snap/go/10455/src/net/mac.go:15
	_go_fuzz_dep_.CoverTab[7320]++
						buf := make([]byte, 0, len(a)*3-1)
//line /snap/go/10455/src/net/mac.go:16
	_go_fuzz_dep_.CoverTab[786717] = 0
						for i, b := range a {
//line /snap/go/10455/src/net/mac.go:17
		if _go_fuzz_dep_.CoverTab[786717] == 0 {
//line /snap/go/10455/src/net/mac.go:17
			_go_fuzz_dep_.CoverTab[529214]++
//line /snap/go/10455/src/net/mac.go:17
		} else {
//line /snap/go/10455/src/net/mac.go:17
			_go_fuzz_dep_.CoverTab[529215]++
//line /snap/go/10455/src/net/mac.go:17
		}
//line /snap/go/10455/src/net/mac.go:17
		_go_fuzz_dep_.CoverTab[786717] = 1
//line /snap/go/10455/src/net/mac.go:17
		_go_fuzz_dep_.CoverTab[7324]++
							if i > 0 {
//line /snap/go/10455/src/net/mac.go:18
			_go_fuzz_dep_.CoverTab[529192]++
//line /snap/go/10455/src/net/mac.go:18
			_go_fuzz_dep_.CoverTab[7326]++
								buf = append(buf, ':')
//line /snap/go/10455/src/net/mac.go:19
			// _ = "end of CoverTab[7326]"
		} else {
//line /snap/go/10455/src/net/mac.go:20
			_go_fuzz_dep_.CoverTab[529193]++
//line /snap/go/10455/src/net/mac.go:20
			_go_fuzz_dep_.CoverTab[7327]++
//line /snap/go/10455/src/net/mac.go:20
			// _ = "end of CoverTab[7327]"
//line /snap/go/10455/src/net/mac.go:20
		}
//line /snap/go/10455/src/net/mac.go:20
		// _ = "end of CoverTab[7324]"
//line /snap/go/10455/src/net/mac.go:20
		_go_fuzz_dep_.CoverTab[7325]++
							buf = append(buf, hexDigit[b>>4])
							buf = append(buf, hexDigit[b&0xF])
//line /snap/go/10455/src/net/mac.go:22
		// _ = "end of CoverTab[7325]"
	}
//line /snap/go/10455/src/net/mac.go:23
	if _go_fuzz_dep_.CoverTab[786717] == 0 {
//line /snap/go/10455/src/net/mac.go:23
		_go_fuzz_dep_.CoverTab[529216]++
//line /snap/go/10455/src/net/mac.go:23
	} else {
//line /snap/go/10455/src/net/mac.go:23
		_go_fuzz_dep_.CoverTab[529217]++
//line /snap/go/10455/src/net/mac.go:23
	}
//line /snap/go/10455/src/net/mac.go:23
	// _ = "end of CoverTab[7320]"
//line /snap/go/10455/src/net/mac.go:23
	_go_fuzz_dep_.CoverTab[7321]++
						return string(buf)
//line /snap/go/10455/src/net/mac.go:24
	// _ = "end of CoverTab[7321]"
}

// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
//line /snap/go/10455/src/net/mac.go:27
// IP over InfiniBand link-layer address using one of the following formats:
//line /snap/go/10455/src/net/mac.go:27
//
//line /snap/go/10455/src/net/mac.go:27
//	00:00:5e:00:53:01
//line /snap/go/10455/src/net/mac.go:27
//	02:00:5e:10:00:00:00:01
//line /snap/go/10455/src/net/mac.go:27
//	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
//line /snap/go/10455/src/net/mac.go:27
//	00-00-5e-00-53-01
//line /snap/go/10455/src/net/mac.go:27
//	02-00-5e-10-00-00-00-01
//line /snap/go/10455/src/net/mac.go:27
//	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
//line /snap/go/10455/src/net/mac.go:27
//	0000.5e00.5301
//line /snap/go/10455/src/net/mac.go:27
//	0200.5e10.0000.0001
//line /snap/go/10455/src/net/mac.go:27
//	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
//line /snap/go/10455/src/net/mac.go:39
func ParseMAC(s string) (hw HardwareAddr, err error) {
//line /snap/go/10455/src/net/mac.go:39
	_go_fuzz_dep_.CoverTab[7328]++
						if len(s) < 14 {
//line /snap/go/10455/src/net/mac.go:40
		_go_fuzz_dep_.CoverTab[529194]++
//line /snap/go/10455/src/net/mac.go:40
		_go_fuzz_dep_.CoverTab[7331]++
							goto error
//line /snap/go/10455/src/net/mac.go:41
		// _ = "end of CoverTab[7331]"
	} else {
//line /snap/go/10455/src/net/mac.go:42
		_go_fuzz_dep_.CoverTab[529195]++
//line /snap/go/10455/src/net/mac.go:42
		_go_fuzz_dep_.CoverTab[7332]++
//line /snap/go/10455/src/net/mac.go:42
		// _ = "end of CoverTab[7332]"
//line /snap/go/10455/src/net/mac.go:42
	}
//line /snap/go/10455/src/net/mac.go:42
	// _ = "end of CoverTab[7328]"
//line /snap/go/10455/src/net/mac.go:42
	_go_fuzz_dep_.CoverTab[7329]++

						if s[2] == ':' || func() bool {
//line /snap/go/10455/src/net/mac.go:44
		_go_fuzz_dep_.CoverTab[7333]++
//line /snap/go/10455/src/net/mac.go:44
		return s[2] == '-'
//line /snap/go/10455/src/net/mac.go:44
		// _ = "end of CoverTab[7333]"
//line /snap/go/10455/src/net/mac.go:44
	}() {
//line /snap/go/10455/src/net/mac.go:44
		_go_fuzz_dep_.CoverTab[529196]++
//line /snap/go/10455/src/net/mac.go:44
		_go_fuzz_dep_.CoverTab[7334]++
							if (len(s)+1)%3 != 0 {
//line /snap/go/10455/src/net/mac.go:45
			_go_fuzz_dep_.CoverTab[529198]++
//line /snap/go/10455/src/net/mac.go:45
			_go_fuzz_dep_.CoverTab[7337]++
								goto error
//line /snap/go/10455/src/net/mac.go:46
			// _ = "end of CoverTab[7337]"
		} else {
//line /snap/go/10455/src/net/mac.go:47
			_go_fuzz_dep_.CoverTab[529199]++
//line /snap/go/10455/src/net/mac.go:47
			_go_fuzz_dep_.CoverTab[7338]++
//line /snap/go/10455/src/net/mac.go:47
			// _ = "end of CoverTab[7338]"
//line /snap/go/10455/src/net/mac.go:47
		}
//line /snap/go/10455/src/net/mac.go:47
		// _ = "end of CoverTab[7334]"
//line /snap/go/10455/src/net/mac.go:47
		_go_fuzz_dep_.CoverTab[7335]++
							n := (len(s) + 1) / 3
							if n != 6 && func() bool {
//line /snap/go/10455/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[7339]++
//line /snap/go/10455/src/net/mac.go:49
			return n != 8
//line /snap/go/10455/src/net/mac.go:49
			// _ = "end of CoverTab[7339]"
//line /snap/go/10455/src/net/mac.go:49
		}() && func() bool {
//line /snap/go/10455/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[7340]++
//line /snap/go/10455/src/net/mac.go:49
			return n != 20
//line /snap/go/10455/src/net/mac.go:49
			// _ = "end of CoverTab[7340]"
//line /snap/go/10455/src/net/mac.go:49
		}() {
//line /snap/go/10455/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[529200]++
//line /snap/go/10455/src/net/mac.go:49
			_go_fuzz_dep_.CoverTab[7341]++
								goto error
//line /snap/go/10455/src/net/mac.go:50
			// _ = "end of CoverTab[7341]"
		} else {
//line /snap/go/10455/src/net/mac.go:51
			_go_fuzz_dep_.CoverTab[529201]++
//line /snap/go/10455/src/net/mac.go:51
			_go_fuzz_dep_.CoverTab[7342]++
//line /snap/go/10455/src/net/mac.go:51
			// _ = "end of CoverTab[7342]"
//line /snap/go/10455/src/net/mac.go:51
		}
//line /snap/go/10455/src/net/mac.go:51
		// _ = "end of CoverTab[7335]"
//line /snap/go/10455/src/net/mac.go:51
		_go_fuzz_dep_.CoverTab[7336]++
							hw = make(HardwareAddr, n)
//line /snap/go/10455/src/net/mac.go:52
		_go_fuzz_dep_.CoverTab[786718] = 0
							for x, i := 0, 0; i < n; i++ {
//line /snap/go/10455/src/net/mac.go:53
			if _go_fuzz_dep_.CoverTab[786718] == 0 {
//line /snap/go/10455/src/net/mac.go:53
				_go_fuzz_dep_.CoverTab[529218]++
//line /snap/go/10455/src/net/mac.go:53
			} else {
//line /snap/go/10455/src/net/mac.go:53
				_go_fuzz_dep_.CoverTab[529219]++
//line /snap/go/10455/src/net/mac.go:53
			}
//line /snap/go/10455/src/net/mac.go:53
			_go_fuzz_dep_.CoverTab[786718] = 1
//line /snap/go/10455/src/net/mac.go:53
			_go_fuzz_dep_.CoverTab[7343]++
								var ok bool
								if hw[i], ok = xtoi2(s[x:], s[2]); !ok {
//line /snap/go/10455/src/net/mac.go:55
				_go_fuzz_dep_.CoverTab[529202]++
//line /snap/go/10455/src/net/mac.go:55
				_go_fuzz_dep_.CoverTab[7345]++
									goto error
//line /snap/go/10455/src/net/mac.go:56
				// _ = "end of CoverTab[7345]"
			} else {
//line /snap/go/10455/src/net/mac.go:57
				_go_fuzz_dep_.CoverTab[529203]++
//line /snap/go/10455/src/net/mac.go:57
				_go_fuzz_dep_.CoverTab[7346]++
//line /snap/go/10455/src/net/mac.go:57
				// _ = "end of CoverTab[7346]"
//line /snap/go/10455/src/net/mac.go:57
			}
//line /snap/go/10455/src/net/mac.go:57
			// _ = "end of CoverTab[7343]"
//line /snap/go/10455/src/net/mac.go:57
			_go_fuzz_dep_.CoverTab[7344]++
								x += 3
//line /snap/go/10455/src/net/mac.go:58
			// _ = "end of CoverTab[7344]"
		}
//line /snap/go/10455/src/net/mac.go:59
		if _go_fuzz_dep_.CoverTab[786718] == 0 {
//line /snap/go/10455/src/net/mac.go:59
			_go_fuzz_dep_.CoverTab[529220]++
//line /snap/go/10455/src/net/mac.go:59
		} else {
//line /snap/go/10455/src/net/mac.go:59
			_go_fuzz_dep_.CoverTab[529221]++
//line /snap/go/10455/src/net/mac.go:59
		}
//line /snap/go/10455/src/net/mac.go:59
		// _ = "end of CoverTab[7336]"
	} else {
//line /snap/go/10455/src/net/mac.go:60
		_go_fuzz_dep_.CoverTab[529197]++
//line /snap/go/10455/src/net/mac.go:60
		_go_fuzz_dep_.CoverTab[7347]++
//line /snap/go/10455/src/net/mac.go:60
		if s[4] == '.' {
//line /snap/go/10455/src/net/mac.go:60
			_go_fuzz_dep_.CoverTab[529204]++
//line /snap/go/10455/src/net/mac.go:60
			_go_fuzz_dep_.CoverTab[7348]++
								if (len(s)+1)%5 != 0 {
//line /snap/go/10455/src/net/mac.go:61
				_go_fuzz_dep_.CoverTab[529206]++
//line /snap/go/10455/src/net/mac.go:61
				_go_fuzz_dep_.CoverTab[7351]++
									goto error
//line /snap/go/10455/src/net/mac.go:62
				// _ = "end of CoverTab[7351]"
			} else {
//line /snap/go/10455/src/net/mac.go:63
				_go_fuzz_dep_.CoverTab[529207]++
//line /snap/go/10455/src/net/mac.go:63
				_go_fuzz_dep_.CoverTab[7352]++
//line /snap/go/10455/src/net/mac.go:63
				// _ = "end of CoverTab[7352]"
//line /snap/go/10455/src/net/mac.go:63
			}
//line /snap/go/10455/src/net/mac.go:63
			// _ = "end of CoverTab[7348]"
//line /snap/go/10455/src/net/mac.go:63
			_go_fuzz_dep_.CoverTab[7349]++
								n := 2 * (len(s) + 1) / 5
								if n != 6 && func() bool {
//line /snap/go/10455/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[7353]++
//line /snap/go/10455/src/net/mac.go:65
				return n != 8
//line /snap/go/10455/src/net/mac.go:65
				// _ = "end of CoverTab[7353]"
//line /snap/go/10455/src/net/mac.go:65
			}() && func() bool {
//line /snap/go/10455/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[7354]++
//line /snap/go/10455/src/net/mac.go:65
				return n != 20
//line /snap/go/10455/src/net/mac.go:65
				// _ = "end of CoverTab[7354]"
//line /snap/go/10455/src/net/mac.go:65
			}() {
//line /snap/go/10455/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[529208]++
//line /snap/go/10455/src/net/mac.go:65
				_go_fuzz_dep_.CoverTab[7355]++
									goto error
//line /snap/go/10455/src/net/mac.go:66
				// _ = "end of CoverTab[7355]"
			} else {
//line /snap/go/10455/src/net/mac.go:67
				_go_fuzz_dep_.CoverTab[529209]++
//line /snap/go/10455/src/net/mac.go:67
				_go_fuzz_dep_.CoverTab[7356]++
//line /snap/go/10455/src/net/mac.go:67
				// _ = "end of CoverTab[7356]"
//line /snap/go/10455/src/net/mac.go:67
			}
//line /snap/go/10455/src/net/mac.go:67
			// _ = "end of CoverTab[7349]"
//line /snap/go/10455/src/net/mac.go:67
			_go_fuzz_dep_.CoverTab[7350]++
								hw = make(HardwareAddr, n)
//line /snap/go/10455/src/net/mac.go:68
			_go_fuzz_dep_.CoverTab[786719] = 0
								for x, i := 0, 0; i < n; i += 2 {
//line /snap/go/10455/src/net/mac.go:69
				if _go_fuzz_dep_.CoverTab[786719] == 0 {
//line /snap/go/10455/src/net/mac.go:69
					_go_fuzz_dep_.CoverTab[529222]++
//line /snap/go/10455/src/net/mac.go:69
				} else {
//line /snap/go/10455/src/net/mac.go:69
					_go_fuzz_dep_.CoverTab[529223]++
//line /snap/go/10455/src/net/mac.go:69
				}
//line /snap/go/10455/src/net/mac.go:69
				_go_fuzz_dep_.CoverTab[786719] = 1
//line /snap/go/10455/src/net/mac.go:69
				_go_fuzz_dep_.CoverTab[7357]++
									var ok bool
									if hw[i], ok = xtoi2(s[x:x+2], 0); !ok {
//line /snap/go/10455/src/net/mac.go:71
					_go_fuzz_dep_.CoverTab[529210]++
//line /snap/go/10455/src/net/mac.go:71
					_go_fuzz_dep_.CoverTab[7360]++
										goto error
//line /snap/go/10455/src/net/mac.go:72
					// _ = "end of CoverTab[7360]"
				} else {
//line /snap/go/10455/src/net/mac.go:73
					_go_fuzz_dep_.CoverTab[529211]++
//line /snap/go/10455/src/net/mac.go:73
					_go_fuzz_dep_.CoverTab[7361]++
//line /snap/go/10455/src/net/mac.go:73
					// _ = "end of CoverTab[7361]"
//line /snap/go/10455/src/net/mac.go:73
				}
//line /snap/go/10455/src/net/mac.go:73
				// _ = "end of CoverTab[7357]"
//line /snap/go/10455/src/net/mac.go:73
				_go_fuzz_dep_.CoverTab[7358]++
									if hw[i+1], ok = xtoi2(s[x+2:], s[4]); !ok {
//line /snap/go/10455/src/net/mac.go:74
					_go_fuzz_dep_.CoverTab[529212]++
//line /snap/go/10455/src/net/mac.go:74
					_go_fuzz_dep_.CoverTab[7362]++
										goto error
//line /snap/go/10455/src/net/mac.go:75
					// _ = "end of CoverTab[7362]"
				} else {
//line /snap/go/10455/src/net/mac.go:76
					_go_fuzz_dep_.CoverTab[529213]++
//line /snap/go/10455/src/net/mac.go:76
					_go_fuzz_dep_.CoverTab[7363]++
//line /snap/go/10455/src/net/mac.go:76
					// _ = "end of CoverTab[7363]"
//line /snap/go/10455/src/net/mac.go:76
				}
//line /snap/go/10455/src/net/mac.go:76
				// _ = "end of CoverTab[7358]"
//line /snap/go/10455/src/net/mac.go:76
				_go_fuzz_dep_.CoverTab[7359]++
									x += 5
//line /snap/go/10455/src/net/mac.go:77
				// _ = "end of CoverTab[7359]"
			}
//line /snap/go/10455/src/net/mac.go:78
			if _go_fuzz_dep_.CoverTab[786719] == 0 {
//line /snap/go/10455/src/net/mac.go:78
				_go_fuzz_dep_.CoverTab[529224]++
//line /snap/go/10455/src/net/mac.go:78
			} else {
//line /snap/go/10455/src/net/mac.go:78
				_go_fuzz_dep_.CoverTab[529225]++
//line /snap/go/10455/src/net/mac.go:78
			}
//line /snap/go/10455/src/net/mac.go:78
			// _ = "end of CoverTab[7350]"
		} else {
//line /snap/go/10455/src/net/mac.go:79
			_go_fuzz_dep_.CoverTab[529205]++
//line /snap/go/10455/src/net/mac.go:79
			_go_fuzz_dep_.CoverTab[7364]++
								goto error
//line /snap/go/10455/src/net/mac.go:80
			// _ = "end of CoverTab[7364]"
		}
//line /snap/go/10455/src/net/mac.go:81
		// _ = "end of CoverTab[7347]"
//line /snap/go/10455/src/net/mac.go:81
	}
//line /snap/go/10455/src/net/mac.go:81
	// _ = "end of CoverTab[7329]"
//line /snap/go/10455/src/net/mac.go:81
	_go_fuzz_dep_.CoverTab[7330]++
						return hw, nil

error:
						return nil, &AddrError{Err: "invalid MAC address", Addr: s}
//line /snap/go/10455/src/net/mac.go:85
	// _ = "end of CoverTab[7330]"
}

//line /snap/go/10455/src/net/mac.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/mac.go:86
var _ = _go_fuzz_dep_.CoverTab
