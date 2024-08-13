// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains CRC32 algorithms that are not specific to any architecture
// and don't use hardware acceleration.
//
// The simple (and slow) CRC32 implementation only uses a 256*4 bytes table.
//
// The slicing-by-8 algorithm is a faster implementation that uses a bigger
// table (8*256*4 bytes).

//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
package crc32

//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
import (
//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
)
//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
import (
//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:13
)

// simpleMakeTable allocates and constructs a Table for the specified
//line /usr/local/go/src/hash/crc32/crc32_generic.go:15
// polynomial. The table is suitable for use with the simple algorithm
//line /usr/local/go/src/hash/crc32/crc32_generic.go:15
// (simpleUpdate).
//line /usr/local/go/src/hash/crc32/crc32_generic.go:18
func simpleMakeTable(poly uint32) *Table {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:18
	_go_fuzz_dep_.CoverTab[26632]++
								t := new(Table)
								simplePopulateTable(poly, t)
								return t
//line /usr/local/go/src/hash/crc32/crc32_generic.go:21
	// _ = "end of CoverTab[26632]"
}

// simplePopulateTable constructs a Table for the specified polynomial, suitable
//line /usr/local/go/src/hash/crc32/crc32_generic.go:24
// for use with simpleUpdate.
//line /usr/local/go/src/hash/crc32/crc32_generic.go:26
func simplePopulateTable(poly uint32, t *Table) {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:26
	_go_fuzz_dep_.CoverTab[26633]++
								for i := 0; i < 256; i++ {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:27
		_go_fuzz_dep_.CoverTab[26634]++
									crc := uint32(i)
									for j := 0; j < 8; j++ {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:29
			_go_fuzz_dep_.CoverTab[26636]++
										if crc&1 == 1 {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:30
				_go_fuzz_dep_.CoverTab[26637]++
											crc = (crc >> 1) ^ poly
//line /usr/local/go/src/hash/crc32/crc32_generic.go:31
				// _ = "end of CoverTab[26637]"
			} else {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:32
				_go_fuzz_dep_.CoverTab[26638]++
											crc >>= 1
//line /usr/local/go/src/hash/crc32/crc32_generic.go:33
				// _ = "end of CoverTab[26638]"
			}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:34
			// _ = "end of CoverTab[26636]"
		}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:35
		// _ = "end of CoverTab[26634]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:35
		_go_fuzz_dep_.CoverTab[26635]++
									t[i] = crc
//line /usr/local/go/src/hash/crc32/crc32_generic.go:36
		// _ = "end of CoverTab[26635]"
	}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:37
	// _ = "end of CoverTab[26633]"
}

// simpleUpdate uses the simple algorithm to update the CRC, given a table that
//line /usr/local/go/src/hash/crc32/crc32_generic.go:40
// was previously computed using simpleMakeTable.
//line /usr/local/go/src/hash/crc32/crc32_generic.go:42
func simpleUpdate(crc uint32, tab *Table, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:42
	_go_fuzz_dep_.CoverTab[26639]++
								crc = ^crc
								for _, v := range p {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:44
		_go_fuzz_dep_.CoverTab[26641]++
									crc = tab[byte(crc)^v] ^ (crc >> 8)
//line /usr/local/go/src/hash/crc32/crc32_generic.go:45
		// _ = "end of CoverTab[26641]"
	}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:46
	// _ = "end of CoverTab[26639]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:46
	_go_fuzz_dep_.CoverTab[26640]++
								return ^crc
//line /usr/local/go/src/hash/crc32/crc32_generic.go:47
	// _ = "end of CoverTab[26640]"
}

// Use slicing-by-8 when payload >= this value.
const slicing8Cutoff = 16

// slicing8Table is array of 8 Tables, used by the slicing-by-8 algorithm.
type slicing8Table [8]Table

// slicingMakeTable constructs a slicing8Table for the specified polynomial. The
//line /usr/local/go/src/hash/crc32/crc32_generic.go:56
// table is suitable for use with the slicing-by-8 algorithm (slicingUpdate).
//line /usr/local/go/src/hash/crc32/crc32_generic.go:58
func slicingMakeTable(poly uint32) *slicing8Table {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:58
	_go_fuzz_dep_.CoverTab[26642]++
								t := new(slicing8Table)
								simplePopulateTable(poly, &t[0])
								for i := 0; i < 256; i++ {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:61
		_go_fuzz_dep_.CoverTab[26644]++
									crc := t[0][i]
									for j := 1; j < 8; j++ {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:63
			_go_fuzz_dep_.CoverTab[26645]++
										crc = t[0][crc&0xFF] ^ (crc >> 8)
										t[j][i] = crc
//line /usr/local/go/src/hash/crc32/crc32_generic.go:65
			// _ = "end of CoverTab[26645]"
		}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:66
		// _ = "end of CoverTab[26644]"
	}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:67
	// _ = "end of CoverTab[26642]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:67
	_go_fuzz_dep_.CoverTab[26643]++
								return t
//line /usr/local/go/src/hash/crc32/crc32_generic.go:68
	// _ = "end of CoverTab[26643]"
}

// slicingUpdate uses the slicing-by-8 algorithm to update the CRC, given a
//line /usr/local/go/src/hash/crc32/crc32_generic.go:71
// table that was previously computed using slicingMakeTable.
//line /usr/local/go/src/hash/crc32/crc32_generic.go:73
func slicingUpdate(crc uint32, tab *slicing8Table, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:73
	_go_fuzz_dep_.CoverTab[26646]++
								if len(p) >= slicing8Cutoff {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:74
		_go_fuzz_dep_.CoverTab[26649]++
									crc = ^crc
									for len(p) > 8 {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:76
			_go_fuzz_dep_.CoverTab[26651]++
										crc ^= uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24
										crc = tab[0][p[7]] ^ tab[1][p[6]] ^ tab[2][p[5]] ^ tab[3][p[4]] ^
				tab[4][crc>>24] ^ tab[5][(crc>>16)&0xFF] ^
				tab[6][(crc>>8)&0xFF] ^ tab[7][crc&0xFF]
										p = p[8:]
//line /usr/local/go/src/hash/crc32/crc32_generic.go:81
			// _ = "end of CoverTab[26651]"
		}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:82
		// _ = "end of CoverTab[26649]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:82
		_go_fuzz_dep_.CoverTab[26650]++
									crc = ^crc
//line /usr/local/go/src/hash/crc32/crc32_generic.go:83
		// _ = "end of CoverTab[26650]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:84
		_go_fuzz_dep_.CoverTab[26652]++
//line /usr/local/go/src/hash/crc32/crc32_generic.go:84
		// _ = "end of CoverTab[26652]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:84
	}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:84
	// _ = "end of CoverTab[26646]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:84
	_go_fuzz_dep_.CoverTab[26647]++
								if len(p) == 0 {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:85
		_go_fuzz_dep_.CoverTab[26653]++
									return crc
//line /usr/local/go/src/hash/crc32/crc32_generic.go:86
		// _ = "end of CoverTab[26653]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32_generic.go:87
		_go_fuzz_dep_.CoverTab[26654]++
//line /usr/local/go/src/hash/crc32/crc32_generic.go:87
		// _ = "end of CoverTab[26654]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:87
	}
//line /usr/local/go/src/hash/crc32/crc32_generic.go:87
	// _ = "end of CoverTab[26647]"
//line /usr/local/go/src/hash/crc32/crc32_generic.go:87
	_go_fuzz_dep_.CoverTab[26648]++
								return simpleUpdate(crc, &tab[0], p)
//line /usr/local/go/src/hash/crc32/crc32_generic.go:88
	// _ = "end of CoverTab[26648]"
}

//line /usr/local/go/src/hash/crc32/crc32_generic.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/hash/crc32/crc32_generic.go:89
var _ = _go_fuzz_dep_.CoverTab
