// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
package idna

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:7
)

// appendMapping appends the mapping for the respective rune. isMapped must be
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:9
// true. A mapping is a categorization of a rune as defined in UTS #46.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:11
func (c info) appendMapping(b []byte, s string) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:11
	_go_fuzz_dep_.CoverTab[71756]++
										index := int(c >> indexShift)
										if c&xorBit == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:13
		_go_fuzz_dep_.CoverTab[71759]++
											s := mappings[index:]
											return append(b, s[1:s[0]+1]...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:15
		// _ = "end of CoverTab[71759]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:16
		_go_fuzz_dep_.CoverTab[71760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:16
		// _ = "end of CoverTab[71760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:16
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:16
	// _ = "end of CoverTab[71756]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:16
	_go_fuzz_dep_.CoverTab[71757]++
										b = append(b, s...)
										if c&inlineXOR == inlineXOR {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:18
		_go_fuzz_dep_.CoverTab[71761]++

											b[len(b)-1] ^= byte(index)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:20
		// _ = "end of CoverTab[71761]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:21
		_go_fuzz_dep_.CoverTab[71762]++
											for p := len(b) - int(xorData[index]); p < len(b); p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:22
			_go_fuzz_dep_.CoverTab[71763]++
												index++
												b[p] ^= xorData[index]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:24
			// _ = "end of CoverTab[71763]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:25
		// _ = "end of CoverTab[71762]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:26
	// _ = "end of CoverTab[71757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:26
	_go_fuzz_dep_.CoverTab[71758]++
										return b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:27
	// _ = "end of CoverTab[71758]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:32
type valueRange struct {
	value	uint16	// header: value:stride
	lo, hi	byte	// header: lo:n
}

type sparseBlocks struct {
	values	[]valueRange
	offset	[]uint16
}

var idnaSparse = sparseBlocks{
	values:	idnaSparseValues[:],
	offset:	idnaSparseOffset[:],
}

// Don't use newIdnaTrie to avoid unconditional linking in of the table.
var trie = &idnaTrie{}

// lookup determines the type of block n and looks up the value for b.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:50
// For n < t.cutoff, the block is a simple lookup table. Otherwise, the block
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:50
// is a list of ranges with an accompanying value. Given a matching range r,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:50
// the value for b is by r.value + (b - r.lo) * stride.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:54
func (t *sparseBlocks) lookup(n uint32, b byte) uint16 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:54
	_go_fuzz_dep_.CoverTab[71764]++
										offset := t.offset[n]
										header := t.values[offset]
										lo := offset + 1
										hi := lo + uint16(header.lo)
										for lo < hi {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:59
		_go_fuzz_dep_.CoverTab[71766]++
											m := lo + (hi-lo)/2
											r := t.values[m]
											if r.lo <= b && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:62
			_go_fuzz_dep_.CoverTab[71768]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:62
			return b <= r.hi
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:62
			// _ = "end of CoverTab[71768]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:62
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:62
			_go_fuzz_dep_.CoverTab[71769]++
												return r.value + uint16(b-r.lo)*header.value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:63
			// _ = "end of CoverTab[71769]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:64
			_go_fuzz_dep_.CoverTab[71770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:64
			// _ = "end of CoverTab[71770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:64
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:64
		// _ = "end of CoverTab[71766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:64
		_go_fuzz_dep_.CoverTab[71767]++
											if b < r.lo {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:65
			_go_fuzz_dep_.CoverTab[71771]++
												hi = m
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:66
			// _ = "end of CoverTab[71771]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:67
			_go_fuzz_dep_.CoverTab[71772]++
												lo = m + 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:68
			// _ = "end of CoverTab[71772]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:69
		// _ = "end of CoverTab[71767]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:70
	// _ = "end of CoverTab[71764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:70
	_go_fuzz_dep_.CoverTab[71765]++
										return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:71
	// _ = "end of CoverTab[71765]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/idna/trie.go:72
var _ = _go_fuzz_dep_.CoverTab