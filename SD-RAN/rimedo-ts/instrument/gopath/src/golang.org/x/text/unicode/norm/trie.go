// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:5
)

type valueRange struct {
	value	uint16	// header: value:stride
	lo, hi	byte	// header: lo:n
}

type sparseBlocks struct {
	values	[]valueRange
	offset	[]uint16
}

var nfcSparse = sparseBlocks{
	values:	nfcSparseValues[:],
	offset:	nfcSparseOffset[:],
}

var nfkcSparse = sparseBlocks{
	values:	nfkcSparseValues[:],
	offset:	nfkcSparseOffset[:],
}

var (
	nfcData		= newNfcTrie(0)
	nfkcData	= newNfkcTrie(0)
)

// lookupValue determines the type of block n and looks up the value for b.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:32
// For n < t.cutoff, the block is a simple lookup table. Otherwise, the block
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:32
// is a list of ranges with an accompanying value. Given a matching range r,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:32
// the value for b is by r.value + (b - r.lo) * stride.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:36
func (t *sparseBlocks) lookup(n uint32, b byte) uint16 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:36
	_go_fuzz_dep_.CoverTab[71244]++
											offset := t.offset[n]
											header := t.values[offset]
											lo := offset + 1
											hi := lo + uint16(header.lo)
											for lo < hi {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:41
		_go_fuzz_dep_.CoverTab[71246]++
												m := lo + (hi-lo)/2
												r := t.values[m]
												if r.lo <= b && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:44
			_go_fuzz_dep_.CoverTab[71248]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:44
			return b <= r.hi
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:44
			// _ = "end of CoverTab[71248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:44
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:44
			_go_fuzz_dep_.CoverTab[71249]++
													return r.value + uint16(b-r.lo)*header.value
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:45
			// _ = "end of CoverTab[71249]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:46
			_go_fuzz_dep_.CoverTab[71250]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:46
			// _ = "end of CoverTab[71250]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:46
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:46
		// _ = "end of CoverTab[71246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:46
		_go_fuzz_dep_.CoverTab[71247]++
												if b < r.lo {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:47
			_go_fuzz_dep_.CoverTab[71251]++
													hi = m
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:48
			// _ = "end of CoverTab[71251]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:49
			_go_fuzz_dep_.CoverTab[71252]++
													lo = m + 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:50
			// _ = "end of CoverTab[71252]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:51
		// _ = "end of CoverTab[71247]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:52
	// _ = "end of CoverTab[71244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:52
	_go_fuzz_dep_.CoverTab[71245]++
											return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:53
	// _ = "end of CoverTab[71245]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/trie.go:54
var _ = _go_fuzz_dep_.CoverTab
