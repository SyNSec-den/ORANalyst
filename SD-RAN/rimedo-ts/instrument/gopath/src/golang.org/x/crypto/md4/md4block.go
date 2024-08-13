// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MD4 block step.
// In its own file so that a faster assembly or C version
// can be substituted easily.

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
package md4

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:9
)

var shift1 = []uint{3, 7, 11, 19}
var shift2 = []uint{3, 5, 9, 13}
var shift3 = []uint{3, 9, 11, 15}

var xIndex2 = []uint{0, 4, 8, 12, 1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15}
var xIndex3 = []uint{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}

func _Block(dig *digest, p []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:18
	_go_fuzz_dep_.CoverTab[85788]++
														a := dig.s[0]
														b := dig.s[1]
														c := dig.s[2]
														d := dig.s[3]
														n := 0
														var X [16]uint32
														for len(p) >= _Chunk {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:25
		_go_fuzz_dep_.CoverTab[85790]++
															aa, bb, cc, dd := a, b, c, d

															j := 0
															for i := 0; i < 16; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:29
			_go_fuzz_dep_.CoverTab[85795]++
																X[i] = uint32(p[j]) | uint32(p[j+1])<<8 | uint32(p[j+2])<<16 | uint32(p[j+3])<<24
																j += 4
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:31
			// _ = "end of CoverTab[85795]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:32
		// _ = "end of CoverTab[85790]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:32
		_go_fuzz_dep_.CoverTab[85791]++

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:46
		for i := uint(0); i < 16; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:46
			_go_fuzz_dep_.CoverTab[85796]++
																x := i
																s := shift1[i%4]
																f := ((c ^ d) & b) ^ d
																a += f + X[x]
																a = a<<s | a>>(32-s)
																a, b, c, d = d, a, b, c
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:52
			// _ = "end of CoverTab[85796]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:53
		// _ = "end of CoverTab[85791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:53
		_go_fuzz_dep_.CoverTab[85792]++

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:56
		for i := uint(0); i < 16; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:56
			_go_fuzz_dep_.CoverTab[85797]++
																x := xIndex2[i]
																s := shift2[i%4]
																g := (b & c) | (b & d) | (c & d)
																a += g + X[x] + 0x5a827999
																a = a<<s | a>>(32-s)
																a, b, c, d = d, a, b, c
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:62
			// _ = "end of CoverTab[85797]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:63
		// _ = "end of CoverTab[85792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:63
		_go_fuzz_dep_.CoverTab[85793]++

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:66
		for i := uint(0); i < 16; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:66
			_go_fuzz_dep_.CoverTab[85798]++
																x := xIndex3[i]
																s := shift3[i%4]
																h := b ^ c ^ d
																a += h + X[x] + 0x6ed9eba1
																a = a<<s | a>>(32-s)
																a, b, c, d = d, a, b, c
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:72
			// _ = "end of CoverTab[85798]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:73
		// _ = "end of CoverTab[85793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:73
		_go_fuzz_dep_.CoverTab[85794]++

															a += aa
															b += bb
															c += cc
															d += dd

															p = p[_Chunk:]
															n += _Chunk
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:81
		// _ = "end of CoverTab[85794]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:82
	// _ = "end of CoverTab[85788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:82
	_go_fuzz_dep_.CoverTab[85789]++

														dig.s[0] = a
														dig.s[1] = b
														dig.s[2] = c
														dig.s[3] = d
														return n
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:88
	// _ = "end of CoverTab[85789]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4block.go:89
var _ = _go_fuzz_dep_.CoverTab
