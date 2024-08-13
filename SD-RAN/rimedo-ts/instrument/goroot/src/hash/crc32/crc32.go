// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/hash/crc32/crc32.go:5
// Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32,
//line /usr/local/go/src/hash/crc32/crc32.go:5
// checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for
//line /usr/local/go/src/hash/crc32/crc32.go:5
// information.
//line /usr/local/go/src/hash/crc32/crc32.go:5
//
//line /usr/local/go/src/hash/crc32/crc32.go:5
// Polynomials are represented in LSB-first form also known as reversed representation.
//line /usr/local/go/src/hash/crc32/crc32.go:5
//
//line /usr/local/go/src/hash/crc32/crc32.go:5
// See https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials
//line /usr/local/go/src/hash/crc32/crc32.go:5
// for information.
//line /usr/local/go/src/hash/crc32/crc32.go:13
package crc32

//line /usr/local/go/src/hash/crc32/crc32.go:13
import (
//line /usr/local/go/src/hash/crc32/crc32.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/hash/crc32/crc32.go:13
)
//line /usr/local/go/src/hash/crc32/crc32.go:13
import (
//line /usr/local/go/src/hash/crc32/crc32.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/hash/crc32/crc32.go:13
)

import (
	"errors"
	"hash"
	"sync"
	"sync/atomic"
)

// The size of a CRC-32 checksum in bytes.
const Size = 4

// Predefined polynomials.
const (
	// IEEE is by far and away the most common CRC-32 polynomial.
	// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
	IEEE	= 0xedb88320

	// Castagnoli's polynomial, used in iSCSI.
	// Has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/26.231911
	Castagnoli	= 0x82f63b78

	// Koopman's polynomial.
	// Also has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman	= 0xeb31d82e
)

// Table is a 256-word table representing the polynomial for efficient processing.
type Table [256]uint32

//line /usr/local/go/src/hash/crc32/crc32.go:73
// castagnoliTable points to a lazily initialized Table for the Castagnoli
//line /usr/local/go/src/hash/crc32/crc32.go:73
// polynomial. MakeTable will always return this value when asked to make a
//line /usr/local/go/src/hash/crc32/crc32.go:73
// Castagnoli table so we can compare against it to find when the caller is
//line /usr/local/go/src/hash/crc32/crc32.go:73
// using this polynomial.
//line /usr/local/go/src/hash/crc32/crc32.go:77
var castagnoliTable *Table
var castagnoliTable8 *slicing8Table
var updateCastagnoli func(crc uint32, p []byte) uint32
var castagnoliOnce sync.Once
var haveCastagnoli atomic.Bool

func castagnoliInit() {
//line /usr/local/go/src/hash/crc32/crc32.go:83
	_go_fuzz_dep_.CoverTab[26539]++
							castagnoliTable = simpleMakeTable(Castagnoli)

							if archAvailableCastagnoli() {
//line /usr/local/go/src/hash/crc32/crc32.go:86
		_go_fuzz_dep_.CoverTab[26541]++
								archInitCastagnoli()
								updateCastagnoli = archUpdateCastagnoli
//line /usr/local/go/src/hash/crc32/crc32.go:88
		// _ = "end of CoverTab[26541]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:89
		_go_fuzz_dep_.CoverTab[26542]++

								castagnoliTable8 = slicingMakeTable(Castagnoli)
								updateCastagnoli = func(crc uint32, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:92
			_go_fuzz_dep_.CoverTab[26543]++
									return slicingUpdate(crc, castagnoliTable8, p)
//line /usr/local/go/src/hash/crc32/crc32.go:93
			// _ = "end of CoverTab[26543]"
		}
//line /usr/local/go/src/hash/crc32/crc32.go:94
		// _ = "end of CoverTab[26542]"
	}
//line /usr/local/go/src/hash/crc32/crc32.go:95
	// _ = "end of CoverTab[26539]"
//line /usr/local/go/src/hash/crc32/crc32.go:95
	_go_fuzz_dep_.CoverTab[26540]++

							haveCastagnoli.Store(true)
//line /usr/local/go/src/hash/crc32/crc32.go:97
	// _ = "end of CoverTab[26540]"
}

// IEEETable is the table for the IEEE polynomial.
var IEEETable = simpleMakeTable(IEEE)

// ieeeTable8 is the slicing8Table for IEEE
var ieeeTable8 *slicing8Table
var updateIEEE func(crc uint32, p []byte) uint32
var ieeeOnce sync.Once

func ieeeInit() {
//line /usr/local/go/src/hash/crc32/crc32.go:108
	_go_fuzz_dep_.CoverTab[26544]++
							if archAvailableIEEE() {
//line /usr/local/go/src/hash/crc32/crc32.go:109
		_go_fuzz_dep_.CoverTab[26545]++
								archInitIEEE()
								updateIEEE = archUpdateIEEE
//line /usr/local/go/src/hash/crc32/crc32.go:111
		// _ = "end of CoverTab[26545]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:112
		_go_fuzz_dep_.CoverTab[26546]++

								ieeeTable8 = slicingMakeTable(IEEE)
								updateIEEE = func(crc uint32, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:115
			_go_fuzz_dep_.CoverTab[26547]++
									return slicingUpdate(crc, ieeeTable8, p)
//line /usr/local/go/src/hash/crc32/crc32.go:116
			// _ = "end of CoverTab[26547]"
		}
//line /usr/local/go/src/hash/crc32/crc32.go:117
		// _ = "end of CoverTab[26546]"
	}
//line /usr/local/go/src/hash/crc32/crc32.go:118
	// _ = "end of CoverTab[26544]"
}

// MakeTable returns a Table constructed from the specified polynomial.
//line /usr/local/go/src/hash/crc32/crc32.go:121
// The contents of this Table must not be modified.
//line /usr/local/go/src/hash/crc32/crc32.go:123
func MakeTable(poly uint32) *Table {
//line /usr/local/go/src/hash/crc32/crc32.go:123
	_go_fuzz_dep_.CoverTab[26548]++
							switch poly {
	case IEEE:
//line /usr/local/go/src/hash/crc32/crc32.go:125
		_go_fuzz_dep_.CoverTab[26549]++
								ieeeOnce.Do(ieeeInit)
								return IEEETable
//line /usr/local/go/src/hash/crc32/crc32.go:127
		// _ = "end of CoverTab[26549]"
	case Castagnoli:
//line /usr/local/go/src/hash/crc32/crc32.go:128
		_go_fuzz_dep_.CoverTab[26550]++
								castagnoliOnce.Do(castagnoliInit)
								return castagnoliTable
//line /usr/local/go/src/hash/crc32/crc32.go:130
		// _ = "end of CoverTab[26550]"
	default:
//line /usr/local/go/src/hash/crc32/crc32.go:131
		_go_fuzz_dep_.CoverTab[26551]++
								return simpleMakeTable(poly)
//line /usr/local/go/src/hash/crc32/crc32.go:132
		// _ = "end of CoverTab[26551]"
	}
//line /usr/local/go/src/hash/crc32/crc32.go:133
	// _ = "end of CoverTab[26548]"
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	crc	uint32
	tab	*Table
}

// New creates a new hash.Hash32 computing the CRC-32 checksum using the
//line /usr/local/go/src/hash/crc32/crc32.go:142
// polynomial represented by the Table. Its Sum method will lay the
//line /usr/local/go/src/hash/crc32/crc32.go:142
// value out in big-endian byte order. The returned Hash32 also
//line /usr/local/go/src/hash/crc32/crc32.go:142
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
//line /usr/local/go/src/hash/crc32/crc32.go:142
// marshal and unmarshal the internal state of the hash.
//line /usr/local/go/src/hash/crc32/crc32.go:147
func New(tab *Table) hash.Hash32 {
//line /usr/local/go/src/hash/crc32/crc32.go:147
	_go_fuzz_dep_.CoverTab[26552]++
							if tab == IEEETable {
//line /usr/local/go/src/hash/crc32/crc32.go:148
		_go_fuzz_dep_.CoverTab[26554]++
								ieeeOnce.Do(ieeeInit)
//line /usr/local/go/src/hash/crc32/crc32.go:149
		// _ = "end of CoverTab[26554]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:150
		_go_fuzz_dep_.CoverTab[26555]++
//line /usr/local/go/src/hash/crc32/crc32.go:150
		// _ = "end of CoverTab[26555]"
//line /usr/local/go/src/hash/crc32/crc32.go:150
	}
//line /usr/local/go/src/hash/crc32/crc32.go:150
	// _ = "end of CoverTab[26552]"
//line /usr/local/go/src/hash/crc32/crc32.go:150
	_go_fuzz_dep_.CoverTab[26553]++
							return &digest{0, tab}
//line /usr/local/go/src/hash/crc32/crc32.go:151
	// _ = "end of CoverTab[26553]"
}

// NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using
//line /usr/local/go/src/hash/crc32/crc32.go:154
// the IEEE polynomial. Its Sum method will lay the value out in
//line /usr/local/go/src/hash/crc32/crc32.go:154
// big-endian byte order. The returned Hash32 also implements
//line /usr/local/go/src/hash/crc32/crc32.go:154
// encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal
//line /usr/local/go/src/hash/crc32/crc32.go:154
// and unmarshal the internal state of the hash.
//line /usr/local/go/src/hash/crc32/crc32.go:159
func NewIEEE() hash.Hash32 {
//line /usr/local/go/src/hash/crc32/crc32.go:159
	_go_fuzz_dep_.CoverTab[26556]++
//line /usr/local/go/src/hash/crc32/crc32.go:159
	return New(IEEETable)
//line /usr/local/go/src/hash/crc32/crc32.go:159
	// _ = "end of CoverTab[26556]"
//line /usr/local/go/src/hash/crc32/crc32.go:159
}

func (d *digest) Size() int {
//line /usr/local/go/src/hash/crc32/crc32.go:161
	_go_fuzz_dep_.CoverTab[26557]++
//line /usr/local/go/src/hash/crc32/crc32.go:161
	return Size
//line /usr/local/go/src/hash/crc32/crc32.go:161
	// _ = "end of CoverTab[26557]"
//line /usr/local/go/src/hash/crc32/crc32.go:161
}

func (d *digest) BlockSize() int {
//line /usr/local/go/src/hash/crc32/crc32.go:163
	_go_fuzz_dep_.CoverTab[26558]++
//line /usr/local/go/src/hash/crc32/crc32.go:163
	return 1
//line /usr/local/go/src/hash/crc32/crc32.go:163
	// _ = "end of CoverTab[26558]"
//line /usr/local/go/src/hash/crc32/crc32.go:163
}

func (d *digest) Reset()	{ _go_fuzz_dep_.CoverTab[26559]++; d.crc = 0; // _ = "end of CoverTab[26559]" }

const (
	magic		= "crc\x01"
	marshaledSize	= len(magic) + 4 + 4
)

func (d *digest) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/crc32/crc32.go:172
	_go_fuzz_dep_.CoverTab[26560]++
							b := make([]byte, 0, marshaledSize)
							b = append(b, magic...)
							b = appendUint32(b, tableSum(d.tab))
							b = appendUint32(b, d.crc)
							return b, nil
//line /usr/local/go/src/hash/crc32/crc32.go:177
	// _ = "end of CoverTab[26560]"
}

func (d *digest) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/crc32/crc32.go:180
	_go_fuzz_dep_.CoverTab[26561]++
							if len(b) < len(magic) || func() bool {
//line /usr/local/go/src/hash/crc32/crc32.go:181
		_go_fuzz_dep_.CoverTab[26565]++
//line /usr/local/go/src/hash/crc32/crc32.go:181
		return string(b[:len(magic)]) != magic
//line /usr/local/go/src/hash/crc32/crc32.go:181
		// _ = "end of CoverTab[26565]"
//line /usr/local/go/src/hash/crc32/crc32.go:181
	}() {
//line /usr/local/go/src/hash/crc32/crc32.go:181
		_go_fuzz_dep_.CoverTab[26566]++
								return errors.New("hash/crc32: invalid hash state identifier")
//line /usr/local/go/src/hash/crc32/crc32.go:182
		// _ = "end of CoverTab[26566]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:183
		_go_fuzz_dep_.CoverTab[26567]++
//line /usr/local/go/src/hash/crc32/crc32.go:183
		// _ = "end of CoverTab[26567]"
//line /usr/local/go/src/hash/crc32/crc32.go:183
	}
//line /usr/local/go/src/hash/crc32/crc32.go:183
	// _ = "end of CoverTab[26561]"
//line /usr/local/go/src/hash/crc32/crc32.go:183
	_go_fuzz_dep_.CoverTab[26562]++
							if len(b) != marshaledSize {
//line /usr/local/go/src/hash/crc32/crc32.go:184
		_go_fuzz_dep_.CoverTab[26568]++
								return errors.New("hash/crc32: invalid hash state size")
//line /usr/local/go/src/hash/crc32/crc32.go:185
		// _ = "end of CoverTab[26568]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:186
		_go_fuzz_dep_.CoverTab[26569]++
//line /usr/local/go/src/hash/crc32/crc32.go:186
		// _ = "end of CoverTab[26569]"
//line /usr/local/go/src/hash/crc32/crc32.go:186
	}
//line /usr/local/go/src/hash/crc32/crc32.go:186
	// _ = "end of CoverTab[26562]"
//line /usr/local/go/src/hash/crc32/crc32.go:186
	_go_fuzz_dep_.CoverTab[26563]++
							if tableSum(d.tab) != readUint32(b[4:]) {
//line /usr/local/go/src/hash/crc32/crc32.go:187
		_go_fuzz_dep_.CoverTab[26570]++
								return errors.New("hash/crc32: tables do not match")
//line /usr/local/go/src/hash/crc32/crc32.go:188
		// _ = "end of CoverTab[26570]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:189
		_go_fuzz_dep_.CoverTab[26571]++
//line /usr/local/go/src/hash/crc32/crc32.go:189
		// _ = "end of CoverTab[26571]"
//line /usr/local/go/src/hash/crc32/crc32.go:189
	}
//line /usr/local/go/src/hash/crc32/crc32.go:189
	// _ = "end of CoverTab[26563]"
//line /usr/local/go/src/hash/crc32/crc32.go:189
	_go_fuzz_dep_.CoverTab[26564]++
							d.crc = readUint32(b[8:])
							return nil
//line /usr/local/go/src/hash/crc32/crc32.go:191
	// _ = "end of CoverTab[26564]"
}

func appendUint32(b []byte, x uint32) []byte {
//line /usr/local/go/src/hash/crc32/crc32.go:194
	_go_fuzz_dep_.CoverTab[26572]++
							a := [4]byte{
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
							return append(b, a[:]...)
//line /usr/local/go/src/hash/crc32/crc32.go:201
	// _ = "end of CoverTab[26572]"
}

func readUint32(b []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:204
	_go_fuzz_dep_.CoverTab[26573]++
							_ = b[3]
							return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
//line /usr/local/go/src/hash/crc32/crc32.go:206
	// _ = "end of CoverTab[26573]"
}

func update(crc uint32, tab *Table, p []byte, checkInitIEEE bool) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:209
	_go_fuzz_dep_.CoverTab[26574]++
							switch {
	case haveCastagnoli.Load() && func() bool {
//line /usr/local/go/src/hash/crc32/crc32.go:211
		_go_fuzz_dep_.CoverTab[26579]++
//line /usr/local/go/src/hash/crc32/crc32.go:211
		return tab == castagnoliTable
//line /usr/local/go/src/hash/crc32/crc32.go:211
		// _ = "end of CoverTab[26579]"
//line /usr/local/go/src/hash/crc32/crc32.go:211
	}():
//line /usr/local/go/src/hash/crc32/crc32.go:211
		_go_fuzz_dep_.CoverTab[26575]++
								return updateCastagnoli(crc, p)
//line /usr/local/go/src/hash/crc32/crc32.go:212
		// _ = "end of CoverTab[26575]"
	case tab == IEEETable:
//line /usr/local/go/src/hash/crc32/crc32.go:213
		_go_fuzz_dep_.CoverTab[26576]++
								if checkInitIEEE {
//line /usr/local/go/src/hash/crc32/crc32.go:214
			_go_fuzz_dep_.CoverTab[26580]++
									ieeeOnce.Do(ieeeInit)
//line /usr/local/go/src/hash/crc32/crc32.go:215
			// _ = "end of CoverTab[26580]"
		} else {
//line /usr/local/go/src/hash/crc32/crc32.go:216
			_go_fuzz_dep_.CoverTab[26581]++
//line /usr/local/go/src/hash/crc32/crc32.go:216
			// _ = "end of CoverTab[26581]"
//line /usr/local/go/src/hash/crc32/crc32.go:216
		}
//line /usr/local/go/src/hash/crc32/crc32.go:216
		// _ = "end of CoverTab[26576]"
//line /usr/local/go/src/hash/crc32/crc32.go:216
		_go_fuzz_dep_.CoverTab[26577]++
								return updateIEEE(crc, p)
//line /usr/local/go/src/hash/crc32/crc32.go:217
		// _ = "end of CoverTab[26577]"
	default:
//line /usr/local/go/src/hash/crc32/crc32.go:218
		_go_fuzz_dep_.CoverTab[26578]++
								return simpleUpdate(crc, tab, p)
//line /usr/local/go/src/hash/crc32/crc32.go:219
		// _ = "end of CoverTab[26578]"
	}
//line /usr/local/go/src/hash/crc32/crc32.go:220
	// _ = "end of CoverTab[26574]"
}

// Update returns the result of adding the bytes in p to the crc.
func Update(crc uint32, tab *Table, p []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:224
	_go_fuzz_dep_.CoverTab[26582]++

//line /usr/local/go/src/hash/crc32/crc32.go:227
	return update(crc, tab, p, true)
//line /usr/local/go/src/hash/crc32/crc32.go:227
	// _ = "end of CoverTab[26582]"
}

func (d *digest) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/hash/crc32/crc32.go:230
	_go_fuzz_dep_.CoverTab[26583]++

//line /usr/local/go/src/hash/crc32/crc32.go:233
	d.crc = update(d.crc, d.tab, p, false)
							return len(p), nil
//line /usr/local/go/src/hash/crc32/crc32.go:234
	// _ = "end of CoverTab[26583]"
}

func (d *digest) Sum32() uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:237
	_go_fuzz_dep_.CoverTab[26584]++
//line /usr/local/go/src/hash/crc32/crc32.go:237
	return d.crc
//line /usr/local/go/src/hash/crc32/crc32.go:237
	// _ = "end of CoverTab[26584]"
//line /usr/local/go/src/hash/crc32/crc32.go:237
}

func (d *digest) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/crc32/crc32.go:239
	_go_fuzz_dep_.CoverTab[26585]++
							s := d.Sum32()
							return append(in, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
//line /usr/local/go/src/hash/crc32/crc32.go:241
	// _ = "end of CoverTab[26585]"
}

// Checksum returns the CRC-32 checksum of data
//line /usr/local/go/src/hash/crc32/crc32.go:244
// using the polynomial represented by the Table.
//line /usr/local/go/src/hash/crc32/crc32.go:246
func Checksum(data []byte, tab *Table) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:246
	_go_fuzz_dep_.CoverTab[26586]++
//line /usr/local/go/src/hash/crc32/crc32.go:246
	return Update(0, tab, data)
//line /usr/local/go/src/hash/crc32/crc32.go:246
	// _ = "end of CoverTab[26586]"
//line /usr/local/go/src/hash/crc32/crc32.go:246
}

// ChecksumIEEE returns the CRC-32 checksum of data
//line /usr/local/go/src/hash/crc32/crc32.go:248
// using the IEEE polynomial.
//line /usr/local/go/src/hash/crc32/crc32.go:250
func ChecksumIEEE(data []byte) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:250
	_go_fuzz_dep_.CoverTab[26587]++
							ieeeOnce.Do(ieeeInit)
							return updateIEEE(0, data)
//line /usr/local/go/src/hash/crc32/crc32.go:252
	// _ = "end of CoverTab[26587]"
}

// tableSum returns the IEEE checksum of table t.
func tableSum(t *Table) uint32 {
//line /usr/local/go/src/hash/crc32/crc32.go:256
	_go_fuzz_dep_.CoverTab[26588]++
							var a [1024]byte
							b := a[:0]
							if t != nil {
//line /usr/local/go/src/hash/crc32/crc32.go:259
		_go_fuzz_dep_.CoverTab[26590]++
								for _, x := range t {
//line /usr/local/go/src/hash/crc32/crc32.go:260
			_go_fuzz_dep_.CoverTab[26591]++
									b = appendUint32(b, x)
//line /usr/local/go/src/hash/crc32/crc32.go:261
			// _ = "end of CoverTab[26591]"
		}
//line /usr/local/go/src/hash/crc32/crc32.go:262
		// _ = "end of CoverTab[26590]"
	} else {
//line /usr/local/go/src/hash/crc32/crc32.go:263
		_go_fuzz_dep_.CoverTab[26592]++
//line /usr/local/go/src/hash/crc32/crc32.go:263
		// _ = "end of CoverTab[26592]"
//line /usr/local/go/src/hash/crc32/crc32.go:263
	}
//line /usr/local/go/src/hash/crc32/crc32.go:263
	// _ = "end of CoverTab[26588]"
//line /usr/local/go/src/hash/crc32/crc32.go:263
	_go_fuzz_dep_.CoverTab[26589]++
							return ChecksumIEEE(b)
//line /usr/local/go/src/hash/crc32/crc32.go:264
	// _ = "end of CoverTab[26589]"
}

//line /usr/local/go/src/hash/crc32/crc32.go:265
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/hash/crc32/crc32.go:265
var _ = _go_fuzz_dep_.CoverTab
