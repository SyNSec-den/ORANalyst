// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Rats.

//line /usr/local/go/src/math/big/ratmarsh.go:7
package big

//line /usr/local/go/src/math/big/ratmarsh.go:7
import (
//line /usr/local/go/src/math/big/ratmarsh.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/ratmarsh.go:7
)
//line /usr/local/go/src/math/big/ratmarsh.go:7
import (
//line /usr/local/go/src/math/big/ratmarsh.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/ratmarsh.go:7
)

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// Gob codec version. Permits backward-compatible changes to the encoding.
const ratGobVersion byte = 1

// GobEncode implements the gob.GobEncoder interface.
func (x *Rat) GobEncode() ([]byte, error) {
//line /usr/local/go/src/math/big/ratmarsh.go:20
	_go_fuzz_dep_.CoverTab[6874]++
							if x == nil {
//line /usr/local/go/src/math/big/ratmarsh.go:21
		_go_fuzz_dep_.CoverTab[6878]++
								return nil, nil
//line /usr/local/go/src/math/big/ratmarsh.go:22
		// _ = "end of CoverTab[6878]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:23
		_go_fuzz_dep_.CoverTab[6879]++
//line /usr/local/go/src/math/big/ratmarsh.go:23
		// _ = "end of CoverTab[6879]"
//line /usr/local/go/src/math/big/ratmarsh.go:23
	}
//line /usr/local/go/src/math/big/ratmarsh.go:23
	// _ = "end of CoverTab[6874]"
//line /usr/local/go/src/math/big/ratmarsh.go:23
	_go_fuzz_dep_.CoverTab[6875]++
							buf := make([]byte, 1+4+(len(x.a.abs)+len(x.b.abs))*_S)
							i := x.b.abs.bytes(buf)
							j := x.a.abs.bytes(buf[:i])
							n := i - j
							if int(uint32(n)) != n {
//line /usr/local/go/src/math/big/ratmarsh.go:28
		_go_fuzz_dep_.CoverTab[6880]++

								return nil, errors.New("Rat.GobEncode: numerator too large")
//line /usr/local/go/src/math/big/ratmarsh.go:30
		// _ = "end of CoverTab[6880]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:31
		_go_fuzz_dep_.CoverTab[6881]++
//line /usr/local/go/src/math/big/ratmarsh.go:31
		// _ = "end of CoverTab[6881]"
//line /usr/local/go/src/math/big/ratmarsh.go:31
	}
//line /usr/local/go/src/math/big/ratmarsh.go:31
	// _ = "end of CoverTab[6875]"
//line /usr/local/go/src/math/big/ratmarsh.go:31
	_go_fuzz_dep_.CoverTab[6876]++
							binary.BigEndian.PutUint32(buf[j-4:j], uint32(n))
							j -= 1 + 4
							b := ratGobVersion << 1
							if x.a.neg {
//line /usr/local/go/src/math/big/ratmarsh.go:35
		_go_fuzz_dep_.CoverTab[6882]++
								b |= 1
//line /usr/local/go/src/math/big/ratmarsh.go:36
		// _ = "end of CoverTab[6882]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:37
		_go_fuzz_dep_.CoverTab[6883]++
//line /usr/local/go/src/math/big/ratmarsh.go:37
		// _ = "end of CoverTab[6883]"
//line /usr/local/go/src/math/big/ratmarsh.go:37
	}
//line /usr/local/go/src/math/big/ratmarsh.go:37
	// _ = "end of CoverTab[6876]"
//line /usr/local/go/src/math/big/ratmarsh.go:37
	_go_fuzz_dep_.CoverTab[6877]++
							buf[j] = b
							return buf[j:], nil
//line /usr/local/go/src/math/big/ratmarsh.go:39
	// _ = "end of CoverTab[6877]"
}

// GobDecode implements the gob.GobDecoder interface.
func (z *Rat) GobDecode(buf []byte) error {
//line /usr/local/go/src/math/big/ratmarsh.go:43
	_go_fuzz_dep_.CoverTab[6884]++
							if len(buf) == 0 {
//line /usr/local/go/src/math/big/ratmarsh.go:44
		_go_fuzz_dep_.CoverTab[6890]++

								*z = Rat{}
								return nil
//line /usr/local/go/src/math/big/ratmarsh.go:47
		// _ = "end of CoverTab[6890]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:48
		_go_fuzz_dep_.CoverTab[6891]++
//line /usr/local/go/src/math/big/ratmarsh.go:48
		// _ = "end of CoverTab[6891]"
//line /usr/local/go/src/math/big/ratmarsh.go:48
	}
//line /usr/local/go/src/math/big/ratmarsh.go:48
	// _ = "end of CoverTab[6884]"
//line /usr/local/go/src/math/big/ratmarsh.go:48
	_go_fuzz_dep_.CoverTab[6885]++
							if len(buf) < 5 {
//line /usr/local/go/src/math/big/ratmarsh.go:49
		_go_fuzz_dep_.CoverTab[6892]++
								return errors.New("Rat.GobDecode: buffer too small")
//line /usr/local/go/src/math/big/ratmarsh.go:50
		// _ = "end of CoverTab[6892]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:51
		_go_fuzz_dep_.CoverTab[6893]++
//line /usr/local/go/src/math/big/ratmarsh.go:51
		// _ = "end of CoverTab[6893]"
//line /usr/local/go/src/math/big/ratmarsh.go:51
	}
//line /usr/local/go/src/math/big/ratmarsh.go:51
	// _ = "end of CoverTab[6885]"
//line /usr/local/go/src/math/big/ratmarsh.go:51
	_go_fuzz_dep_.CoverTab[6886]++
							b := buf[0]
							if b>>1 != ratGobVersion {
//line /usr/local/go/src/math/big/ratmarsh.go:53
		_go_fuzz_dep_.CoverTab[6894]++
								return fmt.Errorf("Rat.GobDecode: encoding version %d not supported", b>>1)
//line /usr/local/go/src/math/big/ratmarsh.go:54
		// _ = "end of CoverTab[6894]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:55
		_go_fuzz_dep_.CoverTab[6895]++
//line /usr/local/go/src/math/big/ratmarsh.go:55
		// _ = "end of CoverTab[6895]"
//line /usr/local/go/src/math/big/ratmarsh.go:55
	}
//line /usr/local/go/src/math/big/ratmarsh.go:55
	// _ = "end of CoverTab[6886]"
//line /usr/local/go/src/math/big/ratmarsh.go:55
	_go_fuzz_dep_.CoverTab[6887]++
							const j = 1 + 4
							ln := binary.BigEndian.Uint32(buf[j-4 : j])
							if uint64(ln) > math.MaxInt-j {
//line /usr/local/go/src/math/big/ratmarsh.go:58
		_go_fuzz_dep_.CoverTab[6896]++
								return errors.New("Rat.GobDecode: invalid length")
//line /usr/local/go/src/math/big/ratmarsh.go:59
		// _ = "end of CoverTab[6896]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:60
		_go_fuzz_dep_.CoverTab[6897]++
//line /usr/local/go/src/math/big/ratmarsh.go:60
		// _ = "end of CoverTab[6897]"
//line /usr/local/go/src/math/big/ratmarsh.go:60
	}
//line /usr/local/go/src/math/big/ratmarsh.go:60
	// _ = "end of CoverTab[6887]"
//line /usr/local/go/src/math/big/ratmarsh.go:60
	_go_fuzz_dep_.CoverTab[6888]++
							i := j + int(ln)
							if len(buf) < i {
//line /usr/local/go/src/math/big/ratmarsh.go:62
		_go_fuzz_dep_.CoverTab[6898]++
								return errors.New("Rat.GobDecode: buffer too small")
//line /usr/local/go/src/math/big/ratmarsh.go:63
		// _ = "end of CoverTab[6898]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:64
		_go_fuzz_dep_.CoverTab[6899]++
//line /usr/local/go/src/math/big/ratmarsh.go:64
		// _ = "end of CoverTab[6899]"
//line /usr/local/go/src/math/big/ratmarsh.go:64
	}
//line /usr/local/go/src/math/big/ratmarsh.go:64
	// _ = "end of CoverTab[6888]"
//line /usr/local/go/src/math/big/ratmarsh.go:64
	_go_fuzz_dep_.CoverTab[6889]++
							z.a.neg = b&1 != 0
							z.a.abs = z.a.abs.setBytes(buf[j:i])
							z.b.abs = z.b.abs.setBytes(buf[i:])
							return nil
//line /usr/local/go/src/math/big/ratmarsh.go:68
	// _ = "end of CoverTab[6889]"
}

// MarshalText implements the encoding.TextMarshaler interface.
func (x *Rat) MarshalText() (text []byte, err error) {
//line /usr/local/go/src/math/big/ratmarsh.go:72
	_go_fuzz_dep_.CoverTab[6900]++
							if x.IsInt() {
//line /usr/local/go/src/math/big/ratmarsh.go:73
		_go_fuzz_dep_.CoverTab[6902]++
								return x.a.MarshalText()
//line /usr/local/go/src/math/big/ratmarsh.go:74
		// _ = "end of CoverTab[6902]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:75
		_go_fuzz_dep_.CoverTab[6903]++
//line /usr/local/go/src/math/big/ratmarsh.go:75
		// _ = "end of CoverTab[6903]"
//line /usr/local/go/src/math/big/ratmarsh.go:75
	}
//line /usr/local/go/src/math/big/ratmarsh.go:75
	// _ = "end of CoverTab[6900]"
//line /usr/local/go/src/math/big/ratmarsh.go:75
	_go_fuzz_dep_.CoverTab[6901]++
							return x.marshal(), nil
//line /usr/local/go/src/math/big/ratmarsh.go:76
	// _ = "end of CoverTab[6901]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (z *Rat) UnmarshalText(text []byte) error {
//line /usr/local/go/src/math/big/ratmarsh.go:80
	_go_fuzz_dep_.CoverTab[6904]++

							if _, ok := z.SetString(string(text)); !ok {
//line /usr/local/go/src/math/big/ratmarsh.go:82
		_go_fuzz_dep_.CoverTab[6906]++
								return fmt.Errorf("math/big: cannot unmarshal %q into a *big.Rat", text)
//line /usr/local/go/src/math/big/ratmarsh.go:83
		// _ = "end of CoverTab[6906]"
	} else {
//line /usr/local/go/src/math/big/ratmarsh.go:84
		_go_fuzz_dep_.CoverTab[6907]++
//line /usr/local/go/src/math/big/ratmarsh.go:84
		// _ = "end of CoverTab[6907]"
//line /usr/local/go/src/math/big/ratmarsh.go:84
	}
//line /usr/local/go/src/math/big/ratmarsh.go:84
	// _ = "end of CoverTab[6904]"
//line /usr/local/go/src/math/big/ratmarsh.go:84
	_go_fuzz_dep_.CoverTab[6905]++
							return nil
//line /usr/local/go/src/math/big/ratmarsh.go:85
	// _ = "end of CoverTab[6905]"
}

//line /usr/local/go/src/math/big/ratmarsh.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/ratmarsh.go:86
var _ = _go_fuzz_dep_.CoverTab
