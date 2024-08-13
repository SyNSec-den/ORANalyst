//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:17
)

import (
	"crypto"
	"encoding/binary"
	"hash"
	"io"
)

type concatKDF struct {
	z, info	[]byte
	i	uint32
	cache	[]byte
	hasher	hash.Hash
}

// NewConcatKDF builds a KDF reader based on the given inputs.
func NewConcatKDF(hash crypto.Hash, z, algID, ptyUInfo, ptyVInfo, supPubInfo, supPrivInfo []byte) io.Reader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:34
	_go_fuzz_dep_.CoverTab[187387]++
												buffer := make([]byte, uint64(len(algID))+uint64(len(ptyUInfo))+uint64(len(ptyVInfo))+uint64(len(supPubInfo))+uint64(len(supPrivInfo)))
												n := 0
												n += copy(buffer, algID)
												n += copy(buffer[n:], ptyUInfo)
												n += copy(buffer[n:], ptyVInfo)
												n += copy(buffer[n:], supPubInfo)
												copy(buffer[n:], supPrivInfo)

												hasher := hash.New()

												return &concatKDF{
		z:	z,
		info:	buffer,
		hasher:	hasher,
		cache:	[]byte{},
		i:	1,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:51
	// _ = "end of CoverTab[187387]"
}

func (ctx *concatKDF) Read(out []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:54
	_go_fuzz_dep_.CoverTab[187388]++
												copied := copy(out, ctx.cache)
												ctx.cache = ctx.cache[copied:]

												for copied < len(out) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:58
		_go_fuzz_dep_.CoverTab[187390]++
													ctx.hasher.Reset()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:62
		_ = binary.Write(ctx.hasher, binary.BigEndian, ctx.i)
													_, _ = ctx.hasher.Write(ctx.z)
													_, _ = ctx.hasher.Write(ctx.info)

													hash := ctx.hasher.Sum(nil)
													chunkCopied := copy(out[copied:], hash)
													copied += chunkCopied
													ctx.cache = hash[chunkCopied:]

													ctx.i++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:71
		// _ = "end of CoverTab[187390]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:72
	// _ = "end of CoverTab[187388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:72
	_go_fuzz_dep_.CoverTab[187389]++

												return copied, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:74
	// _ = "end of CoverTab[187389]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:75
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/concat_kdf.go:75
var _ = _go_fuzz_dep_.CoverTab
