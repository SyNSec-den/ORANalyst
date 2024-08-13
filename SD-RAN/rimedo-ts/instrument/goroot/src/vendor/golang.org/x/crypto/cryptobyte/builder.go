// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
package cryptobyte

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:5
)

import (
	"errors"
	"fmt"
)

// A Builder builds byte strings from fixed-length and length-prefixed values.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// Builders either allocate space as needed, or are ‘fixed’, which means that
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// they write into a given buffer and produce an error if it's exhausted.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// The zero value is a usable Builder that allocates space as needed.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// Simple values are marshaled and appended to a Builder using methods on the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// Builder. Length-prefixed values are marshaled by providing a
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// BuilderContinuation, which is a function that writes the inner contents of
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// the value to a given Builder. See the documentation for BuilderContinuation
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:12
// for details.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:23
type Builder struct {
	err		error
	result		[]byte
	fixedSize	bool
	child		*Builder
	offset		int
	pendingLenLen	int
	pendingIsASN1	bool
	inContinuation	*bool
}

// NewBuilder creates a Builder that appends its output to the given buffer.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:34
// Like append(), the slice will be reallocated if its capacity is exceeded.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:34
// Use Bytes to get the final buffer.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:37
func NewBuilder(buffer []byte) *Builder {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:37
	_go_fuzz_dep_.CoverTab[8676]++
										return &Builder{
		result: buffer,
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:40
	// _ = "end of CoverTab[8676]"
}

// NewFixedBuilder creates a Builder that appends its output into the given
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:43
// buffer. This builder does not reallocate the output buffer. Writes that
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:43
// would exceed the buffer's capacity are treated as an error.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:46
func NewFixedBuilder(buffer []byte) *Builder {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:46
	_go_fuzz_dep_.CoverTab[8677]++
										return &Builder{
		result:		buffer,
		fixedSize:	true,
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:50
	// _ = "end of CoverTab[8677]"
}

// SetError sets the value to be returned as the error from Bytes. Writes
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:53
// performed after calling SetError are ignored.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:55
func (b *Builder) SetError(err error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:55
	_go_fuzz_dep_.CoverTab[8678]++
										b.err = err
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:56
	// _ = "end of CoverTab[8678]"
}

// Bytes returns the bytes written by the builder or an error if one has
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:59
// occurred during building.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:61
func (b *Builder) Bytes() ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:61
	_go_fuzz_dep_.CoverTab[8679]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:62
		_go_fuzz_dep_.CoverTab[8681]++
											return nil, b.err
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:63
		// _ = "end of CoverTab[8681]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:64
		_go_fuzz_dep_.CoverTab[8682]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:64
		// _ = "end of CoverTab[8682]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:64
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:64
	// _ = "end of CoverTab[8679]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:64
	_go_fuzz_dep_.CoverTab[8680]++
										return b.result[b.offset:], nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:65
	// _ = "end of CoverTab[8680]"
}

// BytesOrPanic returns the bytes written by the builder or panics if an error
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:68
// has occurred during building.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:70
func (b *Builder) BytesOrPanic() []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:70
	_go_fuzz_dep_.CoverTab[8683]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:71
		_go_fuzz_dep_.CoverTab[8685]++
											panic(b.err)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:72
		// _ = "end of CoverTab[8685]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:73
		_go_fuzz_dep_.CoverTab[8686]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:73
		// _ = "end of CoverTab[8686]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:73
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:73
	// _ = "end of CoverTab[8683]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:73
	_go_fuzz_dep_.CoverTab[8684]++
										return b.result[b.offset:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:74
	// _ = "end of CoverTab[8684]"
}

// AddUint8 appends an 8-bit value to the byte string.
func (b *Builder) AddUint8(v uint8) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:78
	_go_fuzz_dep_.CoverTab[8687]++
										b.add(byte(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:79
	// _ = "end of CoverTab[8687]"
}

// AddUint16 appends a big-endian, 16-bit value to the byte string.
func (b *Builder) AddUint16(v uint16) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:83
	_go_fuzz_dep_.CoverTab[8688]++
										b.add(byte(v>>8), byte(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:84
	// _ = "end of CoverTab[8688]"
}

// AddUint24 appends a big-endian, 24-bit value to the byte string. The highest
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:87
// byte of the 32-bit input value is silently truncated.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:89
func (b *Builder) AddUint24(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:89
	_go_fuzz_dep_.CoverTab[8689]++
										b.add(byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:90
	// _ = "end of CoverTab[8689]"
}

// AddUint32 appends a big-endian, 32-bit value to the byte string.
func (b *Builder) AddUint32(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:94
	_go_fuzz_dep_.CoverTab[8690]++
										b.add(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:95
	// _ = "end of CoverTab[8690]"
}

// AddUint64 appends a big-endian, 64-bit value to the byte string.
func (b *Builder) AddUint64(v uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:99
	_go_fuzz_dep_.CoverTab[8691]++
										b.add(byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:100
	// _ = "end of CoverTab[8691]"
}

// AddBytes appends a sequence of bytes to the byte string.
func (b *Builder) AddBytes(v []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:104
	_go_fuzz_dep_.CoverTab[8692]++
										b.add(v...)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:105
	// _ = "end of CoverTab[8692]"
}

// BuilderContinuation is a continuation-passing interface for building
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// length-prefixed byte sequences. Builder methods for length-prefixed
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// sequences (AddUint8LengthPrefixed etc) will invoke the BuilderContinuation
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// supplied to them. The child builder passed to the continuation can be used
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// to build the content of the length-prefixed sequence. For example:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	parent := cryptobyte.NewBuilder()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	parent.AddUint8LengthPrefixed(func (child *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	  child.AddUint8(42)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	  child.AddUint8LengthPrefixed(func (grandchild *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	    grandchild.AddUint8(5)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	  })
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// It is an error to write more bytes to the child than allowed by the reserved
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// length prefix. After the continuation returns, the child must be considered
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// invalid, i.e. users must not store any copies or references of the child
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// that outlive the continuation.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// If the continuation panics with a value of type BuildError then the inner
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// error will be returned as the error from Bytes. If the child panics
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:108
// otherwise then Bytes will repanic with the same value.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:130
type BuilderContinuation func(child *Builder)

// BuildError wraps an error. If a BuilderContinuation panics with this value,
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:132
// the panic will be recovered and the inner error will be returned from
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:132
// Builder.Bytes.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:135
type BuildError struct {
	Err error
}

// AddUint8LengthPrefixed adds a 8-bit length-prefixed byte sequence.
func (b *Builder) AddUint8LengthPrefixed(f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:140
	_go_fuzz_dep_.CoverTab[8693]++
										b.addLengthPrefixed(1, false, f)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:141
	// _ = "end of CoverTab[8693]"
}

// AddUint16LengthPrefixed adds a big-endian, 16-bit length-prefixed byte sequence.
func (b *Builder) AddUint16LengthPrefixed(f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:145
	_go_fuzz_dep_.CoverTab[8694]++
										b.addLengthPrefixed(2, false, f)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:146
	// _ = "end of CoverTab[8694]"
}

// AddUint24LengthPrefixed adds a big-endian, 24-bit length-prefixed byte sequence.
func (b *Builder) AddUint24LengthPrefixed(f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:150
	_go_fuzz_dep_.CoverTab[8695]++
										b.addLengthPrefixed(3, false, f)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:151
	// _ = "end of CoverTab[8695]"
}

// AddUint32LengthPrefixed adds a big-endian, 32-bit length-prefixed byte sequence.
func (b *Builder) AddUint32LengthPrefixed(f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:155
	_go_fuzz_dep_.CoverTab[8696]++
										b.addLengthPrefixed(4, false, f)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:156
	// _ = "end of CoverTab[8696]"
}

func (b *Builder) callContinuation(f BuilderContinuation, arg *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:159
	_go_fuzz_dep_.CoverTab[8697]++
										if !*b.inContinuation {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:160
		_go_fuzz_dep_.CoverTab[8699]++
											*b.inContinuation = true

											defer func() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:163
			_go_fuzz_dep_.CoverTab[8700]++
												*b.inContinuation = false

												r := recover()
												if r == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:167
				_go_fuzz_dep_.CoverTab[8702]++
													return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:168
				// _ = "end of CoverTab[8702]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:169
				_go_fuzz_dep_.CoverTab[8703]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:169
				// _ = "end of CoverTab[8703]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:169
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:169
			// _ = "end of CoverTab[8700]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:169
			_go_fuzz_dep_.CoverTab[8701]++

												if buildError, ok := r.(BuildError); ok {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:171
				_go_fuzz_dep_.CoverTab[8704]++
													b.err = buildError.Err
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:172
				// _ = "end of CoverTab[8704]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:173
				_go_fuzz_dep_.CoverTab[8705]++
													panic(r)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:174
				// _ = "end of CoverTab[8705]"
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:175
			// _ = "end of CoverTab[8701]"
		}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:176
		// _ = "end of CoverTab[8699]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:177
		_go_fuzz_dep_.CoverTab[8706]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:177
		// _ = "end of CoverTab[8706]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:177
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:177
	// _ = "end of CoverTab[8697]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:177
	_go_fuzz_dep_.CoverTab[8698]++

										f(arg)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:179
	// _ = "end of CoverTab[8698]"
}

func (b *Builder) addLengthPrefixed(lenLen int, isASN1 bool, f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:182
	_go_fuzz_dep_.CoverTab[8707]++

										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:184
		_go_fuzz_dep_.CoverTab[8710]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:185
		// _ = "end of CoverTab[8710]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:186
		_go_fuzz_dep_.CoverTab[8711]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:186
		// _ = "end of CoverTab[8711]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:186
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:186
	// _ = "end of CoverTab[8707]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:186
	_go_fuzz_dep_.CoverTab[8708]++

										offset := len(b.result)
										b.add(make([]byte, lenLen)...)

										if b.inContinuation == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:191
		_go_fuzz_dep_.CoverTab[8712]++
											b.inContinuation = new(bool)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:192
		// _ = "end of CoverTab[8712]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:193
		_go_fuzz_dep_.CoverTab[8713]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:193
		// _ = "end of CoverTab[8713]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:193
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:193
	// _ = "end of CoverTab[8708]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:193
	_go_fuzz_dep_.CoverTab[8709]++

										b.child = &Builder{
		result:		b.result,
		fixedSize:	b.fixedSize,
		offset:		offset,
		pendingLenLen:	lenLen,
		pendingIsASN1:	isASN1,
		inContinuation:	b.inContinuation,
	}

	b.callContinuation(f, b.child)
	b.flushChild()
	if b.child != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:206
		_go_fuzz_dep_.CoverTab[8714]++
											panic("cryptobyte: internal error")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:207
		// _ = "end of CoverTab[8714]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:208
		_go_fuzz_dep_.CoverTab[8715]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:208
		// _ = "end of CoverTab[8715]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:208
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:208
	// _ = "end of CoverTab[8709]"
}

func (b *Builder) flushChild() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:211
	_go_fuzz_dep_.CoverTab[8716]++
										if b.child == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:212
		_go_fuzz_dep_.CoverTab[8724]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:213
		// _ = "end of CoverTab[8724]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:214
		_go_fuzz_dep_.CoverTab[8725]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:214
		// _ = "end of CoverTab[8725]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:214
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:214
	// _ = "end of CoverTab[8716]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:214
	_go_fuzz_dep_.CoverTab[8717]++
										b.child.flushChild()
										child := b.child
										b.child = nil

										if child.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:219
		_go_fuzz_dep_.CoverTab[8726]++
											b.err = child.err
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:221
		// _ = "end of CoverTab[8726]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:222
		_go_fuzz_dep_.CoverTab[8727]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:222
		// _ = "end of CoverTab[8727]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:222
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:222
	// _ = "end of CoverTab[8717]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:222
	_go_fuzz_dep_.CoverTab[8718]++

										length := len(child.result) - child.pendingLenLen - child.offset

										if length < 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:226
		_go_fuzz_dep_.CoverTab[8728]++
											panic("cryptobyte: internal error")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:227
		// _ = "end of CoverTab[8728]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:228
		_go_fuzz_dep_.CoverTab[8729]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:228
		// _ = "end of CoverTab[8729]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:228
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:228
	// _ = "end of CoverTab[8718]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:228
	_go_fuzz_dep_.CoverTab[8719]++

										if child.pendingIsASN1 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:230
		_go_fuzz_dep_.CoverTab[8730]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:234
		if child.pendingLenLen != 1 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:234
			_go_fuzz_dep_.CoverTab[8734]++
												panic("cryptobyte: internal error")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:235
			// _ = "end of CoverTab[8734]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:236
			_go_fuzz_dep_.CoverTab[8735]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:236
			// _ = "end of CoverTab[8735]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:236
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:236
		// _ = "end of CoverTab[8730]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:236
		_go_fuzz_dep_.CoverTab[8731]++
											var lenLen, lenByte uint8
											if int64(length) > 0xfffffffe {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:238
			_go_fuzz_dep_.CoverTab[8736]++
												b.err = errors.New("pending ASN.1 child too long")
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:240
			// _ = "end of CoverTab[8736]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:241
			_go_fuzz_dep_.CoverTab[8737]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:241
			if length > 0xffffff {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:241
				_go_fuzz_dep_.CoverTab[8738]++
													lenLen = 5
													lenByte = 0x80 | 4
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:243
				// _ = "end of CoverTab[8738]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:244
				_go_fuzz_dep_.CoverTab[8739]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:244
				if length > 0xffff {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:244
					_go_fuzz_dep_.CoverTab[8740]++
														lenLen = 4
														lenByte = 0x80 | 3
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:246
					// _ = "end of CoverTab[8740]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:247
					_go_fuzz_dep_.CoverTab[8741]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:247
					if length > 0xff {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:247
						_go_fuzz_dep_.CoverTab[8742]++
															lenLen = 3
															lenByte = 0x80 | 2
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:249
						// _ = "end of CoverTab[8742]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:250
						_go_fuzz_dep_.CoverTab[8743]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:250
						if length > 0x7f {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:250
							_go_fuzz_dep_.CoverTab[8744]++
																lenLen = 2
																lenByte = 0x80 | 1
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:252
							// _ = "end of CoverTab[8744]"
						} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:253
							_go_fuzz_dep_.CoverTab[8745]++
																lenLen = 1
																lenByte = uint8(length)
																length = 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:256
							// _ = "end of CoverTab[8745]"
						}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
						// _ = "end of CoverTab[8743]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
					}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
					// _ = "end of CoverTab[8741]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
				}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
				// _ = "end of CoverTab[8739]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
			// _ = "end of CoverTab[8737]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
		// _ = "end of CoverTab[8731]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:257
		_go_fuzz_dep_.CoverTab[8732]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:261
		child.result[child.offset] = lenByte
		extraBytes := int(lenLen - 1)
		if extraBytes != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:263
			_go_fuzz_dep_.CoverTab[8746]++
												child.add(make([]byte, extraBytes)...)
												childStart := child.offset + child.pendingLenLen
												copy(child.result[childStart+extraBytes:], child.result[childStart:])
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:266
			// _ = "end of CoverTab[8746]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:267
			_go_fuzz_dep_.CoverTab[8747]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:267
			// _ = "end of CoverTab[8747]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:267
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:267
		// _ = "end of CoverTab[8732]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:267
		_go_fuzz_dep_.CoverTab[8733]++
											child.offset++
											child.pendingLenLen = extraBytes
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:269
		// _ = "end of CoverTab[8733]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:270
		_go_fuzz_dep_.CoverTab[8748]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:270
		// _ = "end of CoverTab[8748]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:270
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:270
	// _ = "end of CoverTab[8719]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:270
	_go_fuzz_dep_.CoverTab[8720]++

										l := length
										for i := child.pendingLenLen - 1; i >= 0; i-- {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:273
		_go_fuzz_dep_.CoverTab[8749]++
											child.result[child.offset+i] = uint8(l)
											l >>= 8
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:275
		// _ = "end of CoverTab[8749]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:276
	// _ = "end of CoverTab[8720]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:276
	_go_fuzz_dep_.CoverTab[8721]++
										if l != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:277
		_go_fuzz_dep_.CoverTab[8750]++
											b.err = fmt.Errorf("cryptobyte: pending child length %d exceeds %d-byte length prefix", length, child.pendingLenLen)
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:279
		// _ = "end of CoverTab[8750]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:280
		_go_fuzz_dep_.CoverTab[8751]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:280
		// _ = "end of CoverTab[8751]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:280
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:280
	// _ = "end of CoverTab[8721]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:280
	_go_fuzz_dep_.CoverTab[8722]++

										if b.fixedSize && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:282
		_go_fuzz_dep_.CoverTab[8752]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:282
		return &b.result[0] != &child.result[0]
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:282
		// _ = "end of CoverTab[8752]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:282
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:282
		_go_fuzz_dep_.CoverTab[8753]++
											panic("cryptobyte: BuilderContinuation reallocated a fixed-size buffer")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:283
		// _ = "end of CoverTab[8753]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:284
		_go_fuzz_dep_.CoverTab[8754]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:284
		// _ = "end of CoverTab[8754]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:284
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:284
	// _ = "end of CoverTab[8722]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:284
	_go_fuzz_dep_.CoverTab[8723]++

										b.result = child.result
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:286
	// _ = "end of CoverTab[8723]"
}

func (b *Builder) add(bytes ...byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:289
	_go_fuzz_dep_.CoverTab[8755]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:290
		_go_fuzz_dep_.CoverTab[8760]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:291
		// _ = "end of CoverTab[8760]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:292
		_go_fuzz_dep_.CoverTab[8761]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:292
		// _ = "end of CoverTab[8761]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:292
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:292
	// _ = "end of CoverTab[8755]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:292
	_go_fuzz_dep_.CoverTab[8756]++
										if b.child != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:293
		_go_fuzz_dep_.CoverTab[8762]++
											panic("cryptobyte: attempted write while child is pending")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:294
		// _ = "end of CoverTab[8762]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:295
		_go_fuzz_dep_.CoverTab[8763]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:295
		// _ = "end of CoverTab[8763]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:295
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:295
	// _ = "end of CoverTab[8756]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:295
	_go_fuzz_dep_.CoverTab[8757]++
										if len(b.result)+len(bytes) < len(bytes) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:296
		_go_fuzz_dep_.CoverTab[8764]++
											b.err = errors.New("cryptobyte: length overflow")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:297
		// _ = "end of CoverTab[8764]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:298
		_go_fuzz_dep_.CoverTab[8765]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:298
		// _ = "end of CoverTab[8765]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:298
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:298
	// _ = "end of CoverTab[8757]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:298
	_go_fuzz_dep_.CoverTab[8758]++
										if b.fixedSize && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:299
		_go_fuzz_dep_.CoverTab[8766]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:299
		return len(b.result)+len(bytes) > cap(b.result)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:299
		// _ = "end of CoverTab[8766]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:299
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:299
		_go_fuzz_dep_.CoverTab[8767]++
											b.err = errors.New("cryptobyte: Builder is exceeding its fixed-size buffer")
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:301
		// _ = "end of CoverTab[8767]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:302
		_go_fuzz_dep_.CoverTab[8768]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:302
		// _ = "end of CoverTab[8768]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:302
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:302
	// _ = "end of CoverTab[8758]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:302
	_go_fuzz_dep_.CoverTab[8759]++
										b.result = append(b.result, bytes...)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:303
	// _ = "end of CoverTab[8759]"
}

// Unwrite rolls back n bytes written directly to the Builder. An attempt by a
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:306
// child builder passed to a continuation to unwrite bytes from its parent will
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:306
// panic.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:309
func (b *Builder) Unwrite(n int) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:309
	_go_fuzz_dep_.CoverTab[8769]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:310
		_go_fuzz_dep_.CoverTab[8774]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:311
		// _ = "end of CoverTab[8774]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:312
		_go_fuzz_dep_.CoverTab[8775]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:312
		// _ = "end of CoverTab[8775]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:312
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:312
	// _ = "end of CoverTab[8769]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:312
	_go_fuzz_dep_.CoverTab[8770]++
										if b.child != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:313
		_go_fuzz_dep_.CoverTab[8776]++
											panic("cryptobyte: attempted unwrite while child is pending")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:314
		// _ = "end of CoverTab[8776]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:315
		_go_fuzz_dep_.CoverTab[8777]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:315
		// _ = "end of CoverTab[8777]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:315
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:315
	// _ = "end of CoverTab[8770]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:315
	_go_fuzz_dep_.CoverTab[8771]++
										length := len(b.result) - b.pendingLenLen - b.offset
										if length < 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:317
		_go_fuzz_dep_.CoverTab[8778]++
											panic("cryptobyte: internal error")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:318
		// _ = "end of CoverTab[8778]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:319
		_go_fuzz_dep_.CoverTab[8779]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:319
		// _ = "end of CoverTab[8779]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:319
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:319
	// _ = "end of CoverTab[8771]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:319
	_go_fuzz_dep_.CoverTab[8772]++
										if n > length {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:320
		_go_fuzz_dep_.CoverTab[8780]++
											panic("cryptobyte: attempted to unwrite more than was written")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:321
		// _ = "end of CoverTab[8780]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:322
		_go_fuzz_dep_.CoverTab[8781]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:322
		// _ = "end of CoverTab[8781]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:322
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:322
	// _ = "end of CoverTab[8772]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:322
	_go_fuzz_dep_.CoverTab[8773]++
										b.result = b.result[:len(b.result)-n]
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:323
	// _ = "end of CoverTab[8773]"
}

// A MarshalingValue marshals itself into a Builder.
type MarshalingValue interface {
	// Marshal is called by Builder.AddValue. It receives a pointer to a builder
	// to marshal itself into. It may return an error that occurred during
	// marshaling, such as unset or invalid values.
	Marshal(b *Builder) error
}

// AddValue calls Marshal on v, passing a pointer to the builder to append to.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:334
// If Marshal returns an error, it is set on the Builder so that subsequent
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:334
// appends don't have an effect.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:337
func (b *Builder) AddValue(v MarshalingValue) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:337
	_go_fuzz_dep_.CoverTab[8782]++
										err := v.Marshal(b)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:339
		_go_fuzz_dep_.CoverTab[8783]++
											b.err = err
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:340
		// _ = "end of CoverTab[8783]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:341
		_go_fuzz_dep_.CoverTab[8784]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:341
		// _ = "end of CoverTab[8784]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:341
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:341
	// _ = "end of CoverTab[8782]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:342
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/builder.go:342
var _ = _go_fuzz_dep_.CoverTab
