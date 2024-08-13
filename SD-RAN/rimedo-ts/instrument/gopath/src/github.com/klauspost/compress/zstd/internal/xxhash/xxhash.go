// Package xxhash implements the 64-bit variant of xxHash (XXH64) as described
// at http://cyan4973.github.io/xxHash/.
// THIS IS VENDORED: Go to github.com/cespare/xxhash for original package.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
package xxhash

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:5
)

import (
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	prime1	uint64	= 11400714785074694791
	prime2	uint64	= 14029467366897019727
	prime3	uint64	= 1609587929392839161
	prime4	uint64	= 9650029242287828579
	prime5	uint64	= 2870177450012600261
)

// NOTE(caleb): I'm using both consts and vars of the primes. Using consts where
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:21
// possible in the Go code is worth a small (but measurable) performance boost
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:21
// by avoiding some MOVQs. Vars are needed for the asm and also are useful for
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:21
// convenience in the Go code in a few places where we need to intentionally
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:21
// avoid constant arithmetic (e.g., v1 := prime1 + prime2 fails because the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:21
// result overflows a uint64).
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:27
var (
	prime1v	= prime1
	prime2v	= prime2
	prime3v	= prime3
	prime4v	= prime4
	prime5v	= prime5
)

// Digest implements hash.Hash64.
type Digest struct {
	v1	uint64
	v2	uint64
	v3	uint64
	v4	uint64
	total	uint64
	mem	[32]byte
	n	int	// how much of mem is used
}

// New creates a new Digest that computes the 64-bit xxHash algorithm.
func New() *Digest {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:47
	_go_fuzz_dep_.CoverTab[90659]++
														var d Digest
														d.Reset()
														return &d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:50
	// _ = "end of CoverTab[90659]"
}

// Reset clears the Digest's state so that it can be reused.
func (d *Digest) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:54
	_go_fuzz_dep_.CoverTab[90660]++
														d.v1 = prime1v + prime2
														d.v2 = prime2
														d.v3 = 0
														d.v4 = -prime1v
														d.total = 0
														d.n = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:60
	// _ = "end of CoverTab[90660]"
}

// Size always returns 8 bytes.
func (d *Digest) Size() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:64
	_go_fuzz_dep_.CoverTab[90661]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:64
	return 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:64
	// _ = "end of CoverTab[90661]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:64
}

// BlockSize always returns 32 bytes.
func (d *Digest) BlockSize() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:67
	_go_fuzz_dep_.CoverTab[90662]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:67
	return 32
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:67
	// _ = "end of CoverTab[90662]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:67
}

// Write adds more data to d. It always returns len(b), nil.
func (d *Digest) Write(b []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:70
	_go_fuzz_dep_.CoverTab[90663]++
														n = len(b)
														d.total += uint64(n)

														if d.n+n < 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:74
		_go_fuzz_dep_.CoverTab[90667]++

															copy(d.mem[d.n:], b)
															d.n += n
															return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:78
		// _ = "end of CoverTab[90667]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:79
		_go_fuzz_dep_.CoverTab[90668]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:79
		// _ = "end of CoverTab[90668]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:79
	// _ = "end of CoverTab[90663]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:79
	_go_fuzz_dep_.CoverTab[90664]++

														if d.n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:81
		_go_fuzz_dep_.CoverTab[90669]++

															copy(d.mem[d.n:], b)
															d.v1 = round(d.v1, u64(d.mem[0:8]))
															d.v2 = round(d.v2, u64(d.mem[8:16]))
															d.v3 = round(d.v3, u64(d.mem[16:24]))
															d.v4 = round(d.v4, u64(d.mem[24:32]))
															b = b[32-d.n:]
															d.n = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:89
		// _ = "end of CoverTab[90669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:90
		_go_fuzz_dep_.CoverTab[90670]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:90
		// _ = "end of CoverTab[90670]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:90
	// _ = "end of CoverTab[90664]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:90
	_go_fuzz_dep_.CoverTab[90665]++

														if len(b) >= 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:92
		_go_fuzz_dep_.CoverTab[90671]++

															nw := writeBlocks(d, b)
															b = b[nw:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:95
		// _ = "end of CoverTab[90671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:96
		_go_fuzz_dep_.CoverTab[90672]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:96
		// _ = "end of CoverTab[90672]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:96
	// _ = "end of CoverTab[90665]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:96
	_go_fuzz_dep_.CoverTab[90666]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:99
	copy(d.mem[:], b)
														d.n = len(b)

														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:102
	// _ = "end of CoverTab[90666]"
}

// Sum appends the current hash to b and returns the resulting slice.
func (d *Digest) Sum(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:106
	_go_fuzz_dep_.CoverTab[90673]++
														s := d.Sum64()
														return append(
		b,
		byte(s>>56),
		byte(s>>48),
		byte(s>>40),
		byte(s>>32),
		byte(s>>24),
		byte(s>>16),
		byte(s>>8),
		byte(s),
	)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:118
	// _ = "end of CoverTab[90673]"
}

// Sum64 returns the current hash.
func (d *Digest) Sum64() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:122
	_go_fuzz_dep_.CoverTab[90674]++
														var h uint64

														if d.total >= 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:125
		_go_fuzz_dep_.CoverTab[90679]++
															v1, v2, v3, v4 := d.v1, d.v2, d.v3, d.v4
															h = rol1(v1) + rol7(v2) + rol12(v3) + rol18(v4)
															h = mergeRound(h, v1)
															h = mergeRound(h, v2)
															h = mergeRound(h, v3)
															h = mergeRound(h, v4)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:131
		// _ = "end of CoverTab[90679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:132
		_go_fuzz_dep_.CoverTab[90680]++
															h = d.v3 + prime5
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:133
		// _ = "end of CoverTab[90680]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:134
	// _ = "end of CoverTab[90674]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:134
	_go_fuzz_dep_.CoverTab[90675]++

														h += d.total

														i, end := 0, d.n
														for ; i+8 <= end; i += 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:139
		_go_fuzz_dep_.CoverTab[90681]++
															k1 := round(0, u64(d.mem[i:i+8]))
															h ^= k1
															h = rol27(h)*prime1 + prime4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:142
		// _ = "end of CoverTab[90681]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:143
	// _ = "end of CoverTab[90675]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:143
	_go_fuzz_dep_.CoverTab[90676]++
														if i+4 <= end {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:144
		_go_fuzz_dep_.CoverTab[90682]++
															h ^= uint64(u32(d.mem[i:i+4])) * prime1
															h = rol23(h)*prime2 + prime3
															i += 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:147
		// _ = "end of CoverTab[90682]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:148
		_go_fuzz_dep_.CoverTab[90683]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:148
		// _ = "end of CoverTab[90683]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:148
	// _ = "end of CoverTab[90676]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:148
	_go_fuzz_dep_.CoverTab[90677]++
														for i < end {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:149
		_go_fuzz_dep_.CoverTab[90684]++
															h ^= uint64(d.mem[i]) * prime5
															h = rol11(h) * prime1
															i++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:152
		// _ = "end of CoverTab[90684]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:153
	// _ = "end of CoverTab[90677]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:153
	_go_fuzz_dep_.CoverTab[90678]++

														h ^= h >> 33
														h *= prime2
														h ^= h >> 29
														h *= prime3
														h ^= h >> 32

														return h
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:161
	// _ = "end of CoverTab[90678]"
}

const (
	magic		= "xxh\x06"
	marshaledSize	= len(magic) + 8*5 + 32
)

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (d *Digest) MarshalBinary() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:170
	_go_fuzz_dep_.CoverTab[90685]++
														b := make([]byte, 0, marshaledSize)
														b = append(b, magic...)
														b = appendUint64(b, d.v1)
														b = appendUint64(b, d.v2)
														b = appendUint64(b, d.v3)
														b = appendUint64(b, d.v4)
														b = appendUint64(b, d.total)
														b = append(b, d.mem[:d.n]...)
														b = b[:len(b)+len(d.mem)-d.n]
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:180
	// _ = "end of CoverTab[90685]"
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (d *Digest) UnmarshalBinary(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:184
	_go_fuzz_dep_.CoverTab[90686]++
														if len(b) < len(magic) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:185
		_go_fuzz_dep_.CoverTab[90689]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:185
		return string(b[:len(magic)]) != magic
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:185
		// _ = "end of CoverTab[90689]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:185
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:185
		_go_fuzz_dep_.CoverTab[90690]++
															return errors.New("xxhash: invalid hash state identifier")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:186
		// _ = "end of CoverTab[90690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:187
		_go_fuzz_dep_.CoverTab[90691]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:187
		// _ = "end of CoverTab[90691]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:187
	// _ = "end of CoverTab[90686]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:187
	_go_fuzz_dep_.CoverTab[90687]++
														if len(b) != marshaledSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:188
		_go_fuzz_dep_.CoverTab[90692]++
															return errors.New("xxhash: invalid hash state size")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:189
		// _ = "end of CoverTab[90692]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:190
		_go_fuzz_dep_.CoverTab[90693]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:190
		// _ = "end of CoverTab[90693]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:190
	// _ = "end of CoverTab[90687]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:190
	_go_fuzz_dep_.CoverTab[90688]++
														b = b[len(magic):]
														b, d.v1 = consumeUint64(b)
														b, d.v2 = consumeUint64(b)
														b, d.v3 = consumeUint64(b)
														b, d.v4 = consumeUint64(b)
														b, d.total = consumeUint64(b)
														copy(d.mem[:], b)
														d.n = int(d.total % uint64(len(d.mem)))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:199
	// _ = "end of CoverTab[90688]"
}

func appendUint64(b []byte, x uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:202
	_go_fuzz_dep_.CoverTab[90694]++
														var a [8]byte
														binary.LittleEndian.PutUint64(a[:], x)
														return append(b, a[:]...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:205
	// _ = "end of CoverTab[90694]"
}

func consumeUint64(b []byte) ([]byte, uint64) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:208
	_go_fuzz_dep_.CoverTab[90695]++
														x := u64(b)
														return b[8:], x
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:210
	// _ = "end of CoverTab[90695]"
}

func u64(b []byte) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:213
	_go_fuzz_dep_.CoverTab[90696]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:213
	return binary.LittleEndian.Uint64(b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:213
	// _ = "end of CoverTab[90696]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:213
}
func u32(b []byte) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:214
	_go_fuzz_dep_.CoverTab[90697]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:214
	return binary.LittleEndian.Uint32(b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:214
	// _ = "end of CoverTab[90697]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:214
}

func round(acc, input uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:216
	_go_fuzz_dep_.CoverTab[90698]++
														acc += input * prime2
														acc = rol31(acc)
														acc *= prime1
														return acc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:220
	// _ = "end of CoverTab[90698]"
}

func mergeRound(acc, val uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:223
	_go_fuzz_dep_.CoverTab[90699]++
														val = round(0, val)
														acc ^= val
														acc = acc*prime1 + prime4
														return acc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:227
	// _ = "end of CoverTab[90699]"
}

func rol1(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:230
	_go_fuzz_dep_.CoverTab[90700]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:230
	return bits.RotateLeft64(x, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:230
	// _ = "end of CoverTab[90700]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:230
}
func rol7(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:231
	_go_fuzz_dep_.CoverTab[90701]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:231
	return bits.RotateLeft64(x, 7)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:231
	// _ = "end of CoverTab[90701]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:231
}
func rol11(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:232
	_go_fuzz_dep_.CoverTab[90702]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:232
	return bits.RotateLeft64(x, 11)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:232
	// _ = "end of CoverTab[90702]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:232
}
func rol12(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:233
	_go_fuzz_dep_.CoverTab[90703]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:233
	return bits.RotateLeft64(x, 12)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:233
	// _ = "end of CoverTab[90703]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:233
}
func rol18(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:234
	_go_fuzz_dep_.CoverTab[90704]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:234
	return bits.RotateLeft64(x, 18)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:234
	// _ = "end of CoverTab[90704]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:234
}
func rol23(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:235
	_go_fuzz_dep_.CoverTab[90705]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:235
	return bits.RotateLeft64(x, 23)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:235
	// _ = "end of CoverTab[90705]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:235
}
func rol27(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:236
	_go_fuzz_dep_.CoverTab[90706]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:236
	return bits.RotateLeft64(x, 27)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:236
	// _ = "end of CoverTab[90706]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:236
}
func rol31(x uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
	_go_fuzz_dep_.CoverTab[90707]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
	return bits.RotateLeft64(x, 31)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
	// _ = "end of CoverTab[90707]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/internal/xxhash/xxhash.go:237
var _ = _go_fuzz_dep_.CoverTab
