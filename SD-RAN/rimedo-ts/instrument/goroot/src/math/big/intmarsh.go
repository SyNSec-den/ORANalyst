// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Ints.

//line /usr/local/go/src/math/big/intmarsh.go:7
package big

//line /usr/local/go/src/math/big/intmarsh.go:7
import (
//line /usr/local/go/src/math/big/intmarsh.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/intmarsh.go:7
)
//line /usr/local/go/src/math/big/intmarsh.go:7
import (
//line /usr/local/go/src/math/big/intmarsh.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/intmarsh.go:7
)

import (
	"bytes"
	"fmt"
)

// Gob codec version. Permits backward-compatible changes to the encoding.
const intGobVersion byte = 1

// GobEncode implements the gob.GobEncoder interface.
func (x *Int) GobEncode() ([]byte, error) {
//line /usr/local/go/src/math/big/intmarsh.go:18
	_go_fuzz_dep_.CoverTab[5644]++
							if x == nil {
//line /usr/local/go/src/math/big/intmarsh.go:19
		_go_fuzz_dep_.CoverTab[5647]++
								return nil, nil
//line /usr/local/go/src/math/big/intmarsh.go:20
		// _ = "end of CoverTab[5647]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:21
		_go_fuzz_dep_.CoverTab[5648]++
//line /usr/local/go/src/math/big/intmarsh.go:21
		// _ = "end of CoverTab[5648]"
//line /usr/local/go/src/math/big/intmarsh.go:21
	}
//line /usr/local/go/src/math/big/intmarsh.go:21
	// _ = "end of CoverTab[5644]"
//line /usr/local/go/src/math/big/intmarsh.go:21
	_go_fuzz_dep_.CoverTab[5645]++
							buf := make([]byte, 1+len(x.abs)*_S)
							i := x.abs.bytes(buf) - 1
							b := intGobVersion << 1
							if x.neg {
//line /usr/local/go/src/math/big/intmarsh.go:25
		_go_fuzz_dep_.CoverTab[5649]++
								b |= 1
//line /usr/local/go/src/math/big/intmarsh.go:26
		// _ = "end of CoverTab[5649]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:27
		_go_fuzz_dep_.CoverTab[5650]++
//line /usr/local/go/src/math/big/intmarsh.go:27
		// _ = "end of CoverTab[5650]"
//line /usr/local/go/src/math/big/intmarsh.go:27
	}
//line /usr/local/go/src/math/big/intmarsh.go:27
	// _ = "end of CoverTab[5645]"
//line /usr/local/go/src/math/big/intmarsh.go:27
	_go_fuzz_dep_.CoverTab[5646]++
							buf[i] = b
							return buf[i:], nil
//line /usr/local/go/src/math/big/intmarsh.go:29
	// _ = "end of CoverTab[5646]"
}

// GobDecode implements the gob.GobDecoder interface.
func (z *Int) GobDecode(buf []byte) error {
//line /usr/local/go/src/math/big/intmarsh.go:33
	_go_fuzz_dep_.CoverTab[5651]++
							if len(buf) == 0 {
//line /usr/local/go/src/math/big/intmarsh.go:34
		_go_fuzz_dep_.CoverTab[5654]++

								*z = Int{}
								return nil
//line /usr/local/go/src/math/big/intmarsh.go:37
		// _ = "end of CoverTab[5654]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:38
		_go_fuzz_dep_.CoverTab[5655]++
//line /usr/local/go/src/math/big/intmarsh.go:38
		// _ = "end of CoverTab[5655]"
//line /usr/local/go/src/math/big/intmarsh.go:38
	}
//line /usr/local/go/src/math/big/intmarsh.go:38
	// _ = "end of CoverTab[5651]"
//line /usr/local/go/src/math/big/intmarsh.go:38
	_go_fuzz_dep_.CoverTab[5652]++
							b := buf[0]
							if b>>1 != intGobVersion {
//line /usr/local/go/src/math/big/intmarsh.go:40
		_go_fuzz_dep_.CoverTab[5656]++
								return fmt.Errorf("Int.GobDecode: encoding version %d not supported", b>>1)
//line /usr/local/go/src/math/big/intmarsh.go:41
		// _ = "end of CoverTab[5656]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:42
		_go_fuzz_dep_.CoverTab[5657]++
//line /usr/local/go/src/math/big/intmarsh.go:42
		// _ = "end of CoverTab[5657]"
//line /usr/local/go/src/math/big/intmarsh.go:42
	}
//line /usr/local/go/src/math/big/intmarsh.go:42
	// _ = "end of CoverTab[5652]"
//line /usr/local/go/src/math/big/intmarsh.go:42
	_go_fuzz_dep_.CoverTab[5653]++
							z.neg = b&1 != 0
							z.abs = z.abs.setBytes(buf[1:])
							return nil
//line /usr/local/go/src/math/big/intmarsh.go:45
	// _ = "end of CoverTab[5653]"
}

// MarshalText implements the encoding.TextMarshaler interface.
func (x *Int) MarshalText() (text []byte, err error) {
//line /usr/local/go/src/math/big/intmarsh.go:49
	_go_fuzz_dep_.CoverTab[5658]++
							if x == nil {
//line /usr/local/go/src/math/big/intmarsh.go:50
		_go_fuzz_dep_.CoverTab[5660]++
								return []byte("<nil>"), nil
//line /usr/local/go/src/math/big/intmarsh.go:51
		// _ = "end of CoverTab[5660]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:52
		_go_fuzz_dep_.CoverTab[5661]++
//line /usr/local/go/src/math/big/intmarsh.go:52
		// _ = "end of CoverTab[5661]"
//line /usr/local/go/src/math/big/intmarsh.go:52
	}
//line /usr/local/go/src/math/big/intmarsh.go:52
	// _ = "end of CoverTab[5658]"
//line /usr/local/go/src/math/big/intmarsh.go:52
	_go_fuzz_dep_.CoverTab[5659]++
							return x.abs.itoa(x.neg, 10), nil
//line /usr/local/go/src/math/big/intmarsh.go:53
	// _ = "end of CoverTab[5659]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (z *Int) UnmarshalText(text []byte) error {
//line /usr/local/go/src/math/big/intmarsh.go:57
	_go_fuzz_dep_.CoverTab[5662]++
							if _, ok := z.setFromScanner(bytes.NewReader(text), 0); !ok {
//line /usr/local/go/src/math/big/intmarsh.go:58
		_go_fuzz_dep_.CoverTab[5664]++
								return fmt.Errorf("math/big: cannot unmarshal %q into a *big.Int", text)
//line /usr/local/go/src/math/big/intmarsh.go:59
		// _ = "end of CoverTab[5664]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:60
		_go_fuzz_dep_.CoverTab[5665]++
//line /usr/local/go/src/math/big/intmarsh.go:60
		// _ = "end of CoverTab[5665]"
//line /usr/local/go/src/math/big/intmarsh.go:60
	}
//line /usr/local/go/src/math/big/intmarsh.go:60
	// _ = "end of CoverTab[5662]"
//line /usr/local/go/src/math/big/intmarsh.go:60
	_go_fuzz_dep_.CoverTab[5663]++
							return nil
//line /usr/local/go/src/math/big/intmarsh.go:61
	// _ = "end of CoverTab[5663]"
}

//line /usr/local/go/src/math/big/intmarsh.go:68
// MarshalJSON implements the json.Marshaler interface.
func (x *Int) MarshalJSON() ([]byte, error) {
//line /usr/local/go/src/math/big/intmarsh.go:69
	_go_fuzz_dep_.CoverTab[5666]++
							if x == nil {
//line /usr/local/go/src/math/big/intmarsh.go:70
		_go_fuzz_dep_.CoverTab[5668]++
								return []byte("null"), nil
//line /usr/local/go/src/math/big/intmarsh.go:71
		// _ = "end of CoverTab[5668]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:72
		_go_fuzz_dep_.CoverTab[5669]++
//line /usr/local/go/src/math/big/intmarsh.go:72
		// _ = "end of CoverTab[5669]"
//line /usr/local/go/src/math/big/intmarsh.go:72
	}
//line /usr/local/go/src/math/big/intmarsh.go:72
	// _ = "end of CoverTab[5666]"
//line /usr/local/go/src/math/big/intmarsh.go:72
	_go_fuzz_dep_.CoverTab[5667]++
							return x.abs.itoa(x.neg, 10), nil
//line /usr/local/go/src/math/big/intmarsh.go:73
	// _ = "end of CoverTab[5667]"
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (z *Int) UnmarshalJSON(text []byte) error {
//line /usr/local/go/src/math/big/intmarsh.go:77
	_go_fuzz_dep_.CoverTab[5670]++

							if string(text) == "null" {
//line /usr/local/go/src/math/big/intmarsh.go:79
		_go_fuzz_dep_.CoverTab[5672]++
								return nil
//line /usr/local/go/src/math/big/intmarsh.go:80
		// _ = "end of CoverTab[5672]"
	} else {
//line /usr/local/go/src/math/big/intmarsh.go:81
		_go_fuzz_dep_.CoverTab[5673]++
//line /usr/local/go/src/math/big/intmarsh.go:81
		// _ = "end of CoverTab[5673]"
//line /usr/local/go/src/math/big/intmarsh.go:81
	}
//line /usr/local/go/src/math/big/intmarsh.go:81
	// _ = "end of CoverTab[5670]"
//line /usr/local/go/src/math/big/intmarsh.go:81
	_go_fuzz_dep_.CoverTab[5671]++
							return z.UnmarshalText(text)
//line /usr/local/go/src/math/big/intmarsh.go:82
	// _ = "end of CoverTab[5671]"
}

//line /usr/local/go/src/math/big/intmarsh.go:83
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/intmarsh.go:83
var _ = _go_fuzz_dep_.CoverTab
