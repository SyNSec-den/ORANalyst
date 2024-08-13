// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Floats.

//line /usr/local/go/src/math/big/floatmarsh.go:7
package big

//line /usr/local/go/src/math/big/floatmarsh.go:7
import (
//line /usr/local/go/src/math/big/floatmarsh.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/floatmarsh.go:7
)
//line /usr/local/go/src/math/big/floatmarsh.go:7
import (
//line /usr/local/go/src/math/big/floatmarsh.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/floatmarsh.go:7
)

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Gob codec version. Permits backward-compatible changes to the encoding.
const floatGobVersion byte = 1

// GobEncode implements the gob.GobEncoder interface.
//line /usr/local/go/src/math/big/floatmarsh.go:18
// The Float value and all its attributes (precision,
//line /usr/local/go/src/math/big/floatmarsh.go:18
// rounding mode, accuracy) are marshaled.
//line /usr/local/go/src/math/big/floatmarsh.go:21
func (x *Float) GobEncode() ([]byte, error) {
//line /usr/local/go/src/math/big/floatmarsh.go:21
	_go_fuzz_dep_.CoverTab[4906]++
							if x == nil {
//line /usr/local/go/src/math/big/floatmarsh.go:22
		_go_fuzz_dep_.CoverTab[4911]++
								return nil, nil
//line /usr/local/go/src/math/big/floatmarsh.go:23
		// _ = "end of CoverTab[4911]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:24
		_go_fuzz_dep_.CoverTab[4912]++
//line /usr/local/go/src/math/big/floatmarsh.go:24
		// _ = "end of CoverTab[4912]"
//line /usr/local/go/src/math/big/floatmarsh.go:24
	}
//line /usr/local/go/src/math/big/floatmarsh.go:24
	// _ = "end of CoverTab[4906]"
//line /usr/local/go/src/math/big/floatmarsh.go:24
	_go_fuzz_dep_.CoverTab[4907]++

//line /usr/local/go/src/math/big/floatmarsh.go:27
	sz := 1 + 1 + 4
	n := 0
	if x.form == finite {
//line /usr/local/go/src/math/big/floatmarsh.go:29
		_go_fuzz_dep_.CoverTab[4913]++

								n = int((x.prec + (_W - 1)) / _W)

//line /usr/local/go/src/math/big/floatmarsh.go:37
		if len(x.mant) < n {
//line /usr/local/go/src/math/big/floatmarsh.go:37
			_go_fuzz_dep_.CoverTab[4915]++
									n = len(x.mant)
//line /usr/local/go/src/math/big/floatmarsh.go:38
			// _ = "end of CoverTab[4915]"
		} else {
//line /usr/local/go/src/math/big/floatmarsh.go:39
			_go_fuzz_dep_.CoverTab[4916]++
//line /usr/local/go/src/math/big/floatmarsh.go:39
			// _ = "end of CoverTab[4916]"
//line /usr/local/go/src/math/big/floatmarsh.go:39
		}
//line /usr/local/go/src/math/big/floatmarsh.go:39
		// _ = "end of CoverTab[4913]"
//line /usr/local/go/src/math/big/floatmarsh.go:39
		_go_fuzz_dep_.CoverTab[4914]++

								sz += 4 + n*_S
//line /usr/local/go/src/math/big/floatmarsh.go:41
		// _ = "end of CoverTab[4914]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:42
		_go_fuzz_dep_.CoverTab[4917]++
//line /usr/local/go/src/math/big/floatmarsh.go:42
		// _ = "end of CoverTab[4917]"
//line /usr/local/go/src/math/big/floatmarsh.go:42
	}
//line /usr/local/go/src/math/big/floatmarsh.go:42
	// _ = "end of CoverTab[4907]"
//line /usr/local/go/src/math/big/floatmarsh.go:42
	_go_fuzz_dep_.CoverTab[4908]++
							buf := make([]byte, sz)

							buf[0] = floatGobVersion
							b := byte(x.mode&7)<<5 | byte((x.acc+1)&3)<<3 | byte(x.form&3)<<1
							if x.neg {
//line /usr/local/go/src/math/big/floatmarsh.go:47
		_go_fuzz_dep_.CoverTab[4918]++
								b |= 1
//line /usr/local/go/src/math/big/floatmarsh.go:48
		// _ = "end of CoverTab[4918]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:49
		_go_fuzz_dep_.CoverTab[4919]++
//line /usr/local/go/src/math/big/floatmarsh.go:49
		// _ = "end of CoverTab[4919]"
//line /usr/local/go/src/math/big/floatmarsh.go:49
	}
//line /usr/local/go/src/math/big/floatmarsh.go:49
	// _ = "end of CoverTab[4908]"
//line /usr/local/go/src/math/big/floatmarsh.go:49
	_go_fuzz_dep_.CoverTab[4909]++
							buf[1] = b
							binary.BigEndian.PutUint32(buf[2:], x.prec)

							if x.form == finite {
//line /usr/local/go/src/math/big/floatmarsh.go:53
		_go_fuzz_dep_.CoverTab[4920]++
								binary.BigEndian.PutUint32(buf[6:], uint32(x.exp))
								x.mant[len(x.mant)-n:].bytes(buf[10:])
//line /usr/local/go/src/math/big/floatmarsh.go:55
		// _ = "end of CoverTab[4920]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:56
		_go_fuzz_dep_.CoverTab[4921]++
//line /usr/local/go/src/math/big/floatmarsh.go:56
		// _ = "end of CoverTab[4921]"
//line /usr/local/go/src/math/big/floatmarsh.go:56
	}
//line /usr/local/go/src/math/big/floatmarsh.go:56
	// _ = "end of CoverTab[4909]"
//line /usr/local/go/src/math/big/floatmarsh.go:56
	_go_fuzz_dep_.CoverTab[4910]++

							return buf, nil
//line /usr/local/go/src/math/big/floatmarsh.go:58
	// _ = "end of CoverTab[4910]"
}

// GobDecode implements the gob.GobDecoder interface.
//line /usr/local/go/src/math/big/floatmarsh.go:61
// The result is rounded per the precision and rounding mode of
//line /usr/local/go/src/math/big/floatmarsh.go:61
// z unless z's precision is 0, in which case z is set exactly
//line /usr/local/go/src/math/big/floatmarsh.go:61
// to the decoded value.
//line /usr/local/go/src/math/big/floatmarsh.go:65
func (z *Float) GobDecode(buf []byte) error {
//line /usr/local/go/src/math/big/floatmarsh.go:65
	_go_fuzz_dep_.CoverTab[4922]++
							if len(buf) == 0 {
//line /usr/local/go/src/math/big/floatmarsh.go:66
		_go_fuzz_dep_.CoverTab[4928]++

								*z = Float{}
								return nil
//line /usr/local/go/src/math/big/floatmarsh.go:69
		// _ = "end of CoverTab[4928]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:70
		_go_fuzz_dep_.CoverTab[4929]++
//line /usr/local/go/src/math/big/floatmarsh.go:70
		// _ = "end of CoverTab[4929]"
//line /usr/local/go/src/math/big/floatmarsh.go:70
	}
//line /usr/local/go/src/math/big/floatmarsh.go:70
	// _ = "end of CoverTab[4922]"
//line /usr/local/go/src/math/big/floatmarsh.go:70
	_go_fuzz_dep_.CoverTab[4923]++
							if len(buf) < 6 {
//line /usr/local/go/src/math/big/floatmarsh.go:71
		_go_fuzz_dep_.CoverTab[4930]++
								return errors.New("Float.GobDecode: buffer too small")
//line /usr/local/go/src/math/big/floatmarsh.go:72
		// _ = "end of CoverTab[4930]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:73
		_go_fuzz_dep_.CoverTab[4931]++
//line /usr/local/go/src/math/big/floatmarsh.go:73
		// _ = "end of CoverTab[4931]"
//line /usr/local/go/src/math/big/floatmarsh.go:73
	}
//line /usr/local/go/src/math/big/floatmarsh.go:73
	// _ = "end of CoverTab[4923]"
//line /usr/local/go/src/math/big/floatmarsh.go:73
	_go_fuzz_dep_.CoverTab[4924]++

							if buf[0] != floatGobVersion {
//line /usr/local/go/src/math/big/floatmarsh.go:75
		_go_fuzz_dep_.CoverTab[4932]++
								return fmt.Errorf("Float.GobDecode: encoding version %d not supported", buf[0])
//line /usr/local/go/src/math/big/floatmarsh.go:76
		// _ = "end of CoverTab[4932]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:77
		_go_fuzz_dep_.CoverTab[4933]++
//line /usr/local/go/src/math/big/floatmarsh.go:77
		// _ = "end of CoverTab[4933]"
//line /usr/local/go/src/math/big/floatmarsh.go:77
	}
//line /usr/local/go/src/math/big/floatmarsh.go:77
	// _ = "end of CoverTab[4924]"
//line /usr/local/go/src/math/big/floatmarsh.go:77
	_go_fuzz_dep_.CoverTab[4925]++

							oldPrec := z.prec
							oldMode := z.mode

							b := buf[1]
							z.mode = RoundingMode((b >> 5) & 7)
							z.acc = Accuracy((b>>3)&3) - 1
							z.form = form((b >> 1) & 3)
							z.neg = b&1 != 0
							z.prec = binary.BigEndian.Uint32(buf[2:])

							if z.form == finite {
//line /usr/local/go/src/math/big/floatmarsh.go:89
		_go_fuzz_dep_.CoverTab[4934]++
								if len(buf) < 10 {
//line /usr/local/go/src/math/big/floatmarsh.go:90
			_go_fuzz_dep_.CoverTab[4936]++
									return errors.New("Float.GobDecode: buffer too small for finite form float")
//line /usr/local/go/src/math/big/floatmarsh.go:91
			// _ = "end of CoverTab[4936]"
		} else {
//line /usr/local/go/src/math/big/floatmarsh.go:92
			_go_fuzz_dep_.CoverTab[4937]++
//line /usr/local/go/src/math/big/floatmarsh.go:92
			// _ = "end of CoverTab[4937]"
//line /usr/local/go/src/math/big/floatmarsh.go:92
		}
//line /usr/local/go/src/math/big/floatmarsh.go:92
		// _ = "end of CoverTab[4934]"
//line /usr/local/go/src/math/big/floatmarsh.go:92
		_go_fuzz_dep_.CoverTab[4935]++
								z.exp = int32(binary.BigEndian.Uint32(buf[6:]))
								z.mant = z.mant.setBytes(buf[10:])
//line /usr/local/go/src/math/big/floatmarsh.go:94
		// _ = "end of CoverTab[4935]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:95
		_go_fuzz_dep_.CoverTab[4938]++
//line /usr/local/go/src/math/big/floatmarsh.go:95
		// _ = "end of CoverTab[4938]"
//line /usr/local/go/src/math/big/floatmarsh.go:95
	}
//line /usr/local/go/src/math/big/floatmarsh.go:95
	// _ = "end of CoverTab[4925]"
//line /usr/local/go/src/math/big/floatmarsh.go:95
	_go_fuzz_dep_.CoverTab[4926]++

							if oldPrec != 0 {
//line /usr/local/go/src/math/big/floatmarsh.go:97
		_go_fuzz_dep_.CoverTab[4939]++
								z.mode = oldMode
								z.SetPrec(uint(oldPrec))
//line /usr/local/go/src/math/big/floatmarsh.go:99
		// _ = "end of CoverTab[4939]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:100
		_go_fuzz_dep_.CoverTab[4940]++
//line /usr/local/go/src/math/big/floatmarsh.go:100
		// _ = "end of CoverTab[4940]"
//line /usr/local/go/src/math/big/floatmarsh.go:100
	}
//line /usr/local/go/src/math/big/floatmarsh.go:100
	// _ = "end of CoverTab[4926]"
//line /usr/local/go/src/math/big/floatmarsh.go:100
	_go_fuzz_dep_.CoverTab[4927]++

							return nil
//line /usr/local/go/src/math/big/floatmarsh.go:102
	// _ = "end of CoverTab[4927]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /usr/local/go/src/math/big/floatmarsh.go:105
// Only the Float value is marshaled (in full precision), other
//line /usr/local/go/src/math/big/floatmarsh.go:105
// attributes such as precision or accuracy are ignored.
//line /usr/local/go/src/math/big/floatmarsh.go:108
func (x *Float) MarshalText() (text []byte, err error) {
//line /usr/local/go/src/math/big/floatmarsh.go:108
	_go_fuzz_dep_.CoverTab[4941]++
							if x == nil {
//line /usr/local/go/src/math/big/floatmarsh.go:109
		_go_fuzz_dep_.CoverTab[4943]++
								return []byte("<nil>"), nil
//line /usr/local/go/src/math/big/floatmarsh.go:110
		// _ = "end of CoverTab[4943]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:111
		_go_fuzz_dep_.CoverTab[4944]++
//line /usr/local/go/src/math/big/floatmarsh.go:111
		// _ = "end of CoverTab[4944]"
//line /usr/local/go/src/math/big/floatmarsh.go:111
	}
//line /usr/local/go/src/math/big/floatmarsh.go:111
	// _ = "end of CoverTab[4941]"
//line /usr/local/go/src/math/big/floatmarsh.go:111
	_go_fuzz_dep_.CoverTab[4942]++
							var buf []byte
							return x.Append(buf, 'g', -1), nil
//line /usr/local/go/src/math/big/floatmarsh.go:113
	// _ = "end of CoverTab[4942]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /usr/local/go/src/math/big/floatmarsh.go:116
// The result is rounded per the precision and rounding mode of z.
//line /usr/local/go/src/math/big/floatmarsh.go:116
// If z's precision is 0, it is changed to 64 before rounding takes
//line /usr/local/go/src/math/big/floatmarsh.go:116
// effect.
//line /usr/local/go/src/math/big/floatmarsh.go:120
func (z *Float) UnmarshalText(text []byte) error {
//line /usr/local/go/src/math/big/floatmarsh.go:120
	_go_fuzz_dep_.CoverTab[4945]++

							_, _, err := z.Parse(string(text), 0)
							if err != nil {
//line /usr/local/go/src/math/big/floatmarsh.go:123
		_go_fuzz_dep_.CoverTab[4947]++
								err = fmt.Errorf("math/big: cannot unmarshal %q into a *big.Float (%v)", text, err)
//line /usr/local/go/src/math/big/floatmarsh.go:124
		// _ = "end of CoverTab[4947]"
	} else {
//line /usr/local/go/src/math/big/floatmarsh.go:125
		_go_fuzz_dep_.CoverTab[4948]++
//line /usr/local/go/src/math/big/floatmarsh.go:125
		// _ = "end of CoverTab[4948]"
//line /usr/local/go/src/math/big/floatmarsh.go:125
	}
//line /usr/local/go/src/math/big/floatmarsh.go:125
	// _ = "end of CoverTab[4945]"
//line /usr/local/go/src/math/big/floatmarsh.go:125
	_go_fuzz_dep_.CoverTab[4946]++
							return err
//line /usr/local/go/src/math/big/floatmarsh.go:126
	// _ = "end of CoverTab[4946]"
}

//line /usr/local/go/src/math/big/floatmarsh.go:127
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/floatmarsh.go:127
var _ = _go_fuzz_dep_.CoverTab
