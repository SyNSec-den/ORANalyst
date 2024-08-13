//go:build (amd64 || arm64) && !appengine && gc && !purego && !noasm
// +build amd64 arm64
// +build !appengine
// +build gc
// +build !purego
// +build !noasm

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
package xxhash

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:8
)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:12
//go:noescape
func Sum64(b []byte) uint64

//go:noescape
func writeBlocks(d *Digest, b []byte) int

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash_asm.go:16
var _ = _go_fuzz_dep_.CoverTab
