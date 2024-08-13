// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:5
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
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:32
// For n < t.cutoff, the block is a simple lookup table. Otherwise, the block
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:32
// is a list of ranges with an accompanying value. Given a matching range r,
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:32
// the value for b is by r.value + (b - r.lo) * stride.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:36
func (t *sparseBlocks) lookup(n uint32, b byte) uint16 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:36
	_go_fuzz_dep_.CoverTab[33963]++
										offset := t.offset[n]
										header := t.values[offset]
										lo := offset + 1
										hi := lo + uint16(header.lo)
										for lo < hi {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:41
		_go_fuzz_dep_.CoverTab[33965]++
											m := lo + (hi-lo)/2
											r := t.values[m]
											if r.lo <= b && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:44
			_go_fuzz_dep_.CoverTab[33967]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:44
			return b <= r.hi
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:44
			// _ = "end of CoverTab[33967]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:44
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:44
			_go_fuzz_dep_.CoverTab[33968]++
												return r.value + uint16(b-r.lo)*header.value
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:45
			// _ = "end of CoverTab[33968]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:46
			_go_fuzz_dep_.CoverTab[33969]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:46
			// _ = "end of CoverTab[33969]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:46
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:46
		// _ = "end of CoverTab[33965]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:46
		_go_fuzz_dep_.CoverTab[33966]++
											if b < r.lo {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:47
			_go_fuzz_dep_.CoverTab[33970]++
												hi = m
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:48
			// _ = "end of CoverTab[33970]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:49
			_go_fuzz_dep_.CoverTab[33971]++
												lo = m + 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:50
			// _ = "end of CoverTab[33971]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:51
		// _ = "end of CoverTab[33966]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:52
	// _ = "end of CoverTab[33963]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:52
	_go_fuzz_dep_.CoverTab[33964]++
										return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:53
	// _ = "end of CoverTab[33964]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/trie.go:54
var _ = _go_fuzz_dep_.CoverTab
