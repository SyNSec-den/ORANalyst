// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// Package cryptobyte contains types that help with parsing and constructing
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// length-prefixed, binary messages, including ASN.1 DER. (The asn1 subpackage
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// contains useful ASN.1 constants.)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// The String type is for parsing. It wraps a []byte slice and provides helper
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// functions for consuming structures, value by value.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// The Builder type is for constructing messages. It providers helper functions
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// for appending values and also for appending length-prefixed submessages –
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// without having to worry about calculating the length prefix ahead of time.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// See the documentation and examples for the Builder and String types to get
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:5
// started.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
package cryptobyte

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:18
)

// String represents a string of bytes. It provides methods for parsing
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:20
// fixed-length and length-prefixed values from it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:22
type String []byte

// read advances a String by n bytes and returns them. If less than n bytes
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:24
// remain, it returns nil.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:26
func (s *String) read(n int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:26
	_go_fuzz_dep_.CoverTab[8785]++
										if len(*s) < n || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:27
		_go_fuzz_dep_.CoverTab[8787]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:27
		return n < 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:27
		// _ = "end of CoverTab[8787]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:27
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:27
		_go_fuzz_dep_.CoverTab[8788]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:28
		// _ = "end of CoverTab[8788]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:29
		_go_fuzz_dep_.CoverTab[8789]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:29
		// _ = "end of CoverTab[8789]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:29
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:29
	// _ = "end of CoverTab[8785]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:29
	_go_fuzz_dep_.CoverTab[8786]++
										v := (*s)[:n]
										*s = (*s)[n:]
										return v
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:32
	// _ = "end of CoverTab[8786]"
}

// Skip advances the String by n byte and reports whether it was successful.
func (s *String) Skip(n int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:36
	_go_fuzz_dep_.CoverTab[8790]++
										return s.read(n) != nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:37
	// _ = "end of CoverTab[8790]"
}

// ReadUint8 decodes an 8-bit value into out and advances over it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:40
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:42
func (s *String) ReadUint8(out *uint8) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:42
	_go_fuzz_dep_.CoverTab[8791]++
										v := s.read(1)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:44
		_go_fuzz_dep_.CoverTab[8793]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:45
		// _ = "end of CoverTab[8793]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:46
		_go_fuzz_dep_.CoverTab[8794]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:46
		// _ = "end of CoverTab[8794]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:46
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:46
	// _ = "end of CoverTab[8791]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:46
	_go_fuzz_dep_.CoverTab[8792]++
										*out = uint8(v[0])
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:48
	// _ = "end of CoverTab[8792]"
}

// ReadUint16 decodes a big-endian, 16-bit value into out and advances over it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:51
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:53
func (s *String) ReadUint16(out *uint16) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:53
	_go_fuzz_dep_.CoverTab[8795]++
										v := s.read(2)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:55
		_go_fuzz_dep_.CoverTab[8797]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:56
		// _ = "end of CoverTab[8797]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:57
		_go_fuzz_dep_.CoverTab[8798]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:57
		// _ = "end of CoverTab[8798]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:57
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:57
	// _ = "end of CoverTab[8795]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:57
	_go_fuzz_dep_.CoverTab[8796]++
										*out = uint16(v[0])<<8 | uint16(v[1])
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:59
	// _ = "end of CoverTab[8796]"
}

// ReadUint24 decodes a big-endian, 24-bit value into out and advances over it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:62
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:64
func (s *String) ReadUint24(out *uint32) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:64
	_go_fuzz_dep_.CoverTab[8799]++
										v := s.read(3)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:66
		_go_fuzz_dep_.CoverTab[8801]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:67
		// _ = "end of CoverTab[8801]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:68
		_go_fuzz_dep_.CoverTab[8802]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:68
		// _ = "end of CoverTab[8802]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:68
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:68
	// _ = "end of CoverTab[8799]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:68
	_go_fuzz_dep_.CoverTab[8800]++
										*out = uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2])
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:70
	// _ = "end of CoverTab[8800]"
}

// ReadUint32 decodes a big-endian, 32-bit value into out and advances over it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:73
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:75
func (s *String) ReadUint32(out *uint32) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:75
	_go_fuzz_dep_.CoverTab[8803]++
										v := s.read(4)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:77
		_go_fuzz_dep_.CoverTab[8805]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:78
		// _ = "end of CoverTab[8805]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:79
		_go_fuzz_dep_.CoverTab[8806]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:79
		// _ = "end of CoverTab[8806]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:79
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:79
	// _ = "end of CoverTab[8803]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:79
	_go_fuzz_dep_.CoverTab[8804]++
										*out = uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3])
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:81
	// _ = "end of CoverTab[8804]"
}

// ReadUint64 decodes a big-endian, 64-bit value into out and advances over it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:84
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:86
func (s *String) ReadUint64(out *uint64) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:86
	_go_fuzz_dep_.CoverTab[8807]++
										v := s.read(8)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:88
		_go_fuzz_dep_.CoverTab[8809]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:89
		// _ = "end of CoverTab[8809]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:90
		_go_fuzz_dep_.CoverTab[8810]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:90
		// _ = "end of CoverTab[8810]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:90
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:90
	// _ = "end of CoverTab[8807]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:90
	_go_fuzz_dep_.CoverTab[8808]++
										*out = uint64(v[0])<<56 | uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 | uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 | uint64(v[7])
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:92
	// _ = "end of CoverTab[8808]"
}

func (s *String) readUnsigned(out *uint32, length int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:95
	_go_fuzz_dep_.CoverTab[8811]++
										v := s.read(length)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:97
		_go_fuzz_dep_.CoverTab[8814]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:98
		// _ = "end of CoverTab[8814]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:99
		_go_fuzz_dep_.CoverTab[8815]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:99
		// _ = "end of CoverTab[8815]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:99
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:99
	// _ = "end of CoverTab[8811]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:99
	_go_fuzz_dep_.CoverTab[8812]++
										var result uint32
										for i := 0; i < length; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:101
		_go_fuzz_dep_.CoverTab[8816]++
											result <<= 8
											result |= uint32(v[i])
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:103
		// _ = "end of CoverTab[8816]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:104
	// _ = "end of CoverTab[8812]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:104
	_go_fuzz_dep_.CoverTab[8813]++
										*out = result
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:106
	// _ = "end of CoverTab[8813]"
}

func (s *String) readLengthPrefixed(lenLen int, outChild *String) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:109
	_go_fuzz_dep_.CoverTab[8817]++
										lenBytes := s.read(lenLen)
										if lenBytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:111
		_go_fuzz_dep_.CoverTab[8821]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:112
		// _ = "end of CoverTab[8821]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:113
		_go_fuzz_dep_.CoverTab[8822]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:113
		// _ = "end of CoverTab[8822]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:113
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:113
	// _ = "end of CoverTab[8817]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:113
	_go_fuzz_dep_.CoverTab[8818]++
										var length uint32
										for _, b := range lenBytes {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:115
		_go_fuzz_dep_.CoverTab[8823]++
											length = length << 8
											length = length | uint32(b)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:117
		// _ = "end of CoverTab[8823]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:118
	// _ = "end of CoverTab[8818]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:118
	_go_fuzz_dep_.CoverTab[8819]++
										v := s.read(int(length))
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:120
		_go_fuzz_dep_.CoverTab[8824]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:121
		// _ = "end of CoverTab[8824]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:122
		_go_fuzz_dep_.CoverTab[8825]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:122
		// _ = "end of CoverTab[8825]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:122
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:122
	// _ = "end of CoverTab[8819]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:122
	_go_fuzz_dep_.CoverTab[8820]++
										*outChild = v
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:124
	// _ = "end of CoverTab[8820]"
}

// ReadUint8LengthPrefixed reads the content of an 8-bit length-prefixed value
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:127
// into out and advances over it. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:129
func (s *String) ReadUint8LengthPrefixed(out *String) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:129
	_go_fuzz_dep_.CoverTab[8826]++
										return s.readLengthPrefixed(1, out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:130
	// _ = "end of CoverTab[8826]"
}

// ReadUint16LengthPrefixed reads the content of a big-endian, 16-bit
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:133
// length-prefixed value into out and advances over it. It reports whether the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:133
// read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:136
func (s *String) ReadUint16LengthPrefixed(out *String) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:136
	_go_fuzz_dep_.CoverTab[8827]++
										return s.readLengthPrefixed(2, out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:137
	// _ = "end of CoverTab[8827]"
}

// ReadUint24LengthPrefixed reads the content of a big-endian, 24-bit
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:140
// length-prefixed value into out and advances over it. It reports whether
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:140
// the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:143
func (s *String) ReadUint24LengthPrefixed(out *String) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:143
	_go_fuzz_dep_.CoverTab[8828]++
										return s.readLengthPrefixed(3, out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:144
	// _ = "end of CoverTab[8828]"
}

// ReadBytes reads n bytes into out and advances over them. It reports
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:147
// whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:149
func (s *String) ReadBytes(out *[]byte, n int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:149
	_go_fuzz_dep_.CoverTab[8829]++
										v := s.read(n)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:151
		_go_fuzz_dep_.CoverTab[8831]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:152
		// _ = "end of CoverTab[8831]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:153
		_go_fuzz_dep_.CoverTab[8832]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:153
		// _ = "end of CoverTab[8832]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:153
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:153
	// _ = "end of CoverTab[8829]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:153
	_go_fuzz_dep_.CoverTab[8830]++
										*out = v
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:155
	// _ = "end of CoverTab[8830]"
}

// CopyBytes copies len(out) bytes into out and advances over them. It reports
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:158
// whether the copy operation was successful
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:160
func (s *String) CopyBytes(out []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:160
	_go_fuzz_dep_.CoverTab[8833]++
										n := len(out)
										v := s.read(n)
										if v == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:163
		_go_fuzz_dep_.CoverTab[8835]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:164
		// _ = "end of CoverTab[8835]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:165
		_go_fuzz_dep_.CoverTab[8836]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:165
		// _ = "end of CoverTab[8836]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:165
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:165
	// _ = "end of CoverTab[8833]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:165
	_go_fuzz_dep_.CoverTab[8834]++
										return copy(out, v) == n
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:166
	// _ = "end of CoverTab[8834]"
}

// Empty reports whether the string does not contain any bytes.
func (s String) Empty() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:170
	_go_fuzz_dep_.CoverTab[8837]++
										return len(s) == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:171
	// _ = "end of CoverTab[8837]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:172
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/string.go:172
var _ = _go_fuzz_dep_.CoverTab
