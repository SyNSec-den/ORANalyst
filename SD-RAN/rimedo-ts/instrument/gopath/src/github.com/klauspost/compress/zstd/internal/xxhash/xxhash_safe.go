//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
package xxhash

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:1
)

// Sum64String computes the 64-bit xxHash digest of s.
func Sum64String(s string) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:4
	_go_fuzz_dep_.CoverTab[90708]++
														return Sum64([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:5
	// _ = "end of CoverTab[90708]"
}

// WriteString adds more data to d. It always returns len(s), nil.
func (d *Digest) WriteString(s string) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:9
	_go_fuzz_dep_.CoverTab[90709]++
														return d.Write([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:10
	// _ = "end of CoverTab[90709]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_safe.go:11
var _ = _go_fuzz_dep_.CoverTab
