// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:5
)

const (
	prime3bytes	= 506832829
	prime4bytes	= 2654435761
	prime5bytes	= 889523592379
	prime6bytes	= 227718039650203
	prime7bytes	= 58295818150454627
	prime8bytes	= 0xcf1bbcdcb7a56463
)

// hashLen returns a hash of the lowest mls bytes of with length output bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:16
// mls must be >=3 and <=8. Any other value will return hash for 4 bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:16
// length should always be < 32.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:16
// Preferably length and mls should be a constant for inlining.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:20
func hashLen(u uint64, length, mls uint8) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:20
	_go_fuzz_dep_.CoverTab[94833]++
												switch mls {
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:22
		_go_fuzz_dep_.CoverTab[94834]++
													return (uint32(u<<8) * prime3bytes) >> (32 - length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:23
		// _ = "end of CoverTab[94834]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:24
		_go_fuzz_dep_.CoverTab[94835]++
													return uint32(((u << (64 - 40)) * prime5bytes) >> (64 - length))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:25
		// _ = "end of CoverTab[94835]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:26
		_go_fuzz_dep_.CoverTab[94836]++
													return uint32(((u << (64 - 48)) * prime6bytes) >> (64 - length))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:27
		// _ = "end of CoverTab[94836]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:28
		_go_fuzz_dep_.CoverTab[94837]++
													return uint32(((u << (64 - 56)) * prime7bytes) >> (64 - length))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:29
		// _ = "end of CoverTab[94837]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:30
		_go_fuzz_dep_.CoverTab[94838]++
													return uint32((u * prime8bytes) >> (64 - length))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:31
		// _ = "end of CoverTab[94838]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:32
		_go_fuzz_dep_.CoverTab[94839]++
													return (uint32(u) * prime4bytes) >> (32 - length)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:33
		// _ = "end of CoverTab[94839]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:34
	// _ = "end of CoverTab[94833]"
}

// hash3 returns the hash of the lower 3 bytes of u to fit in a hash table with h bits.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:37
// Preferably h should be a constant and should always be <32.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:39
func hash3(u uint32, h uint8) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:39
	_go_fuzz_dep_.CoverTab[94840]++
												return ((u << (32 - 24)) * prime3bytes) >> ((32 - h) & 31)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:40
	// _ = "end of CoverTab[94840]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/hash.go:41
var _ = _go_fuzz_dep_.CoverTab
