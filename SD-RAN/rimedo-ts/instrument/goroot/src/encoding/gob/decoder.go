// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/gob/decoder.go:5
package gob

//line /usr/local/go/src/encoding/gob/decoder.go:5
import (
//line /usr/local/go/src/encoding/gob/decoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/decoder.go:5
)
//line /usr/local/go/src/encoding/gob/decoder.go:5
import (
//line /usr/local/go/src/encoding/gob/decoder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/decoder.go:5
)

import (
	"bufio"
	"errors"
	"internal/saferio"
	"io"
	"reflect"
	"sync"
)

// tooBig provides a sanity check for sizes; used in several places. Upper limit
//line /usr/local/go/src/encoding/gob/decoder.go:16
// of is 1GB on 32-bit systems, 8GB on 64-bit, allowing room to grow a little
//line /usr/local/go/src/encoding/gob/decoder.go:16
// without overflow.
//line /usr/local/go/src/encoding/gob/decoder.go:19
const tooBig = (1 << 30) << (^uint(0) >> 62)

// A Decoder manages the receipt of type and data information read from the
//line /usr/local/go/src/encoding/gob/decoder.go:21
// remote side of a connection.  It is safe for concurrent use by multiple
//line /usr/local/go/src/encoding/gob/decoder.go:21
// goroutines.
//line /usr/local/go/src/encoding/gob/decoder.go:21
//
//line /usr/local/go/src/encoding/gob/decoder.go:21
// The Decoder does only basic sanity checking on decoded input sizes,
//line /usr/local/go/src/encoding/gob/decoder.go:21
// and its limits are not configurable. Take caution when decoding gob data
//line /usr/local/go/src/encoding/gob/decoder.go:21
// from untrusted sources.
//line /usr/local/go/src/encoding/gob/decoder.go:28
type Decoder struct {
	mutex		sync.Mutex				// each item must be received atomically
	r		io.Reader				// source of the data
	buf		decBuffer				// buffer for more efficient i/o from r
	wireType	map[typeId]*wireType			// map from remote ID to local description
	decoderCache	map[reflect.Type]map[typeId]**decEngine	// cache of compiled engines
	ignorerCache	map[typeId]**decEngine			// ditto for ignored objects
	freeList	*decoderState				// list of free decoderStates; avoids reallocation
	countBuf	[]byte					// used for decoding integers while parsing messages
	err		error
}

// NewDecoder returns a new decoder that reads from the io.Reader.
//line /usr/local/go/src/encoding/gob/decoder.go:40
// If r does not also implement io.ByteReader, it will be wrapped in a
//line /usr/local/go/src/encoding/gob/decoder.go:40
// bufio.Reader.
//line /usr/local/go/src/encoding/gob/decoder.go:43
func NewDecoder(r io.Reader) *Decoder {
//line /usr/local/go/src/encoding/gob/decoder.go:43
	_go_fuzz_dep_.CoverTab[84544]++
							dec := new(Decoder)

							if _, ok := r.(io.ByteReader); !ok {
//line /usr/local/go/src/encoding/gob/decoder.go:46
		_go_fuzz_dep_.CoverTab[84546]++
								r = bufio.NewReader(r)
//line /usr/local/go/src/encoding/gob/decoder.go:47
		// _ = "end of CoverTab[84546]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:48
		_go_fuzz_dep_.CoverTab[84547]++
//line /usr/local/go/src/encoding/gob/decoder.go:48
		// _ = "end of CoverTab[84547]"
//line /usr/local/go/src/encoding/gob/decoder.go:48
	}
//line /usr/local/go/src/encoding/gob/decoder.go:48
	// _ = "end of CoverTab[84544]"
//line /usr/local/go/src/encoding/gob/decoder.go:48
	_go_fuzz_dep_.CoverTab[84545]++
							dec.r = r
							dec.wireType = make(map[typeId]*wireType)
							dec.decoderCache = make(map[reflect.Type]map[typeId]**decEngine)
							dec.ignorerCache = make(map[typeId]**decEngine)
							dec.countBuf = make([]byte, 9)

							return dec
//line /usr/local/go/src/encoding/gob/decoder.go:55
	// _ = "end of CoverTab[84545]"
}

// recvType loads the definition of a type.
func (dec *Decoder) recvType(id typeId) {
//line /usr/local/go/src/encoding/gob/decoder.go:59
	_go_fuzz_dep_.CoverTab[84548]++

							if id < firstUserId || func() bool {
//line /usr/local/go/src/encoding/gob/decoder.go:61
		_go_fuzz_dep_.CoverTab[84551]++
//line /usr/local/go/src/encoding/gob/decoder.go:61
		return dec.wireType[id] != nil
//line /usr/local/go/src/encoding/gob/decoder.go:61
		// _ = "end of CoverTab[84551]"
//line /usr/local/go/src/encoding/gob/decoder.go:61
	}() {
//line /usr/local/go/src/encoding/gob/decoder.go:61
		_go_fuzz_dep_.CoverTab[84552]++
								dec.err = errors.New("gob: duplicate type received")
								return
//line /usr/local/go/src/encoding/gob/decoder.go:63
		// _ = "end of CoverTab[84552]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:64
		_go_fuzz_dep_.CoverTab[84553]++
//line /usr/local/go/src/encoding/gob/decoder.go:64
		// _ = "end of CoverTab[84553]"
//line /usr/local/go/src/encoding/gob/decoder.go:64
	}
//line /usr/local/go/src/encoding/gob/decoder.go:64
	// _ = "end of CoverTab[84548]"
//line /usr/local/go/src/encoding/gob/decoder.go:64
	_go_fuzz_dep_.CoverTab[84549]++

//line /usr/local/go/src/encoding/gob/decoder.go:67
	wire := new(wireType)
	dec.decodeValue(tWireType, reflect.ValueOf(wire))
	if dec.err != nil {
//line /usr/local/go/src/encoding/gob/decoder.go:69
		_go_fuzz_dep_.CoverTab[84554]++
								return
//line /usr/local/go/src/encoding/gob/decoder.go:70
		// _ = "end of CoverTab[84554]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:71
		_go_fuzz_dep_.CoverTab[84555]++
//line /usr/local/go/src/encoding/gob/decoder.go:71
		// _ = "end of CoverTab[84555]"
//line /usr/local/go/src/encoding/gob/decoder.go:71
	}
//line /usr/local/go/src/encoding/gob/decoder.go:71
	// _ = "end of CoverTab[84549]"
//line /usr/local/go/src/encoding/gob/decoder.go:71
	_go_fuzz_dep_.CoverTab[84550]++

							dec.wireType[id] = wire
//line /usr/local/go/src/encoding/gob/decoder.go:73
	// _ = "end of CoverTab[84550]"
}

var errBadCount = errors.New("invalid message length")

// recvMessage reads the next count-delimited item from the input. It is the converse
//line /usr/local/go/src/encoding/gob/decoder.go:78
// of Encoder.writeMessage. It returns false on EOF or other error reading the message.
//line /usr/local/go/src/encoding/gob/decoder.go:80
func (dec *Decoder) recvMessage() bool {
//line /usr/local/go/src/encoding/gob/decoder.go:80
	_go_fuzz_dep_.CoverTab[84556]++

							nbytes, _, err := decodeUintReader(dec.r, dec.countBuf)
							if err != nil {
//line /usr/local/go/src/encoding/gob/decoder.go:83
		_go_fuzz_dep_.CoverTab[84559]++
								dec.err = err
								return false
//line /usr/local/go/src/encoding/gob/decoder.go:85
		// _ = "end of CoverTab[84559]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:86
		_go_fuzz_dep_.CoverTab[84560]++
//line /usr/local/go/src/encoding/gob/decoder.go:86
		// _ = "end of CoverTab[84560]"
//line /usr/local/go/src/encoding/gob/decoder.go:86
	}
//line /usr/local/go/src/encoding/gob/decoder.go:86
	// _ = "end of CoverTab[84556]"
//line /usr/local/go/src/encoding/gob/decoder.go:86
	_go_fuzz_dep_.CoverTab[84557]++
							if nbytes >= tooBig {
//line /usr/local/go/src/encoding/gob/decoder.go:87
		_go_fuzz_dep_.CoverTab[84561]++
								dec.err = errBadCount
								return false
//line /usr/local/go/src/encoding/gob/decoder.go:89
		// _ = "end of CoverTab[84561]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:90
		_go_fuzz_dep_.CoverTab[84562]++
//line /usr/local/go/src/encoding/gob/decoder.go:90
		// _ = "end of CoverTab[84562]"
//line /usr/local/go/src/encoding/gob/decoder.go:90
	}
//line /usr/local/go/src/encoding/gob/decoder.go:90
	// _ = "end of CoverTab[84557]"
//line /usr/local/go/src/encoding/gob/decoder.go:90
	_go_fuzz_dep_.CoverTab[84558]++
							dec.readMessage(int(nbytes))
							return dec.err == nil
//line /usr/local/go/src/encoding/gob/decoder.go:92
	// _ = "end of CoverTab[84558]"
}

// readMessage reads the next nbytes bytes from the input.
func (dec *Decoder) readMessage(nbytes int) {
//line /usr/local/go/src/encoding/gob/decoder.go:96
	_go_fuzz_dep_.CoverTab[84563]++
							if dec.buf.Len() != 0 {
//line /usr/local/go/src/encoding/gob/decoder.go:97
		_go_fuzz_dep_.CoverTab[84565]++

								panic("non-empty decoder buffer")
//line /usr/local/go/src/encoding/gob/decoder.go:99
		// _ = "end of CoverTab[84565]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:100
		_go_fuzz_dep_.CoverTab[84566]++
//line /usr/local/go/src/encoding/gob/decoder.go:100
		// _ = "end of CoverTab[84566]"
//line /usr/local/go/src/encoding/gob/decoder.go:100
	}
//line /usr/local/go/src/encoding/gob/decoder.go:100
	// _ = "end of CoverTab[84563]"
//line /usr/local/go/src/encoding/gob/decoder.go:100
	_go_fuzz_dep_.CoverTab[84564]++
	// Read the data
	var buf []byte
	buf, dec.err = saferio.ReadData(dec.r, uint64(nbytes))
	dec.buf.SetBytes(buf)
	if dec.err == io.EOF {
//line /usr/local/go/src/encoding/gob/decoder.go:105
		_go_fuzz_dep_.CoverTab[84567]++
								dec.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/gob/decoder.go:106
		// _ = "end of CoverTab[84567]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:107
		_go_fuzz_dep_.CoverTab[84568]++
//line /usr/local/go/src/encoding/gob/decoder.go:107
		// _ = "end of CoverTab[84568]"
//line /usr/local/go/src/encoding/gob/decoder.go:107
	}
//line /usr/local/go/src/encoding/gob/decoder.go:107
	// _ = "end of CoverTab[84564]"
}

// toInt turns an encoded uint64 into an int, according to the marshaling rules.
func toInt(x uint64) int64 {
//line /usr/local/go/src/encoding/gob/decoder.go:111
	_go_fuzz_dep_.CoverTab[84569]++
							i := int64(x >> 1)
							if x&1 != 0 {
//line /usr/local/go/src/encoding/gob/decoder.go:113
		_go_fuzz_dep_.CoverTab[84571]++
								i = ^i
//line /usr/local/go/src/encoding/gob/decoder.go:114
		// _ = "end of CoverTab[84571]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:115
		_go_fuzz_dep_.CoverTab[84572]++
//line /usr/local/go/src/encoding/gob/decoder.go:115
		// _ = "end of CoverTab[84572]"
//line /usr/local/go/src/encoding/gob/decoder.go:115
	}
//line /usr/local/go/src/encoding/gob/decoder.go:115
	// _ = "end of CoverTab[84569]"
//line /usr/local/go/src/encoding/gob/decoder.go:115
	_go_fuzz_dep_.CoverTab[84570]++
							return i
//line /usr/local/go/src/encoding/gob/decoder.go:116
	// _ = "end of CoverTab[84570]"
}

func (dec *Decoder) nextInt() int64 {
//line /usr/local/go/src/encoding/gob/decoder.go:119
	_go_fuzz_dep_.CoverTab[84573]++
							n, _, err := decodeUintReader(&dec.buf, dec.countBuf)
							if err != nil {
//line /usr/local/go/src/encoding/gob/decoder.go:121
		_go_fuzz_dep_.CoverTab[84575]++
								dec.err = err
//line /usr/local/go/src/encoding/gob/decoder.go:122
		// _ = "end of CoverTab[84575]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:123
		_go_fuzz_dep_.CoverTab[84576]++
//line /usr/local/go/src/encoding/gob/decoder.go:123
		// _ = "end of CoverTab[84576]"
//line /usr/local/go/src/encoding/gob/decoder.go:123
	}
//line /usr/local/go/src/encoding/gob/decoder.go:123
	// _ = "end of CoverTab[84573]"
//line /usr/local/go/src/encoding/gob/decoder.go:123
	_go_fuzz_dep_.CoverTab[84574]++
							return toInt(n)
//line /usr/local/go/src/encoding/gob/decoder.go:124
	// _ = "end of CoverTab[84574]"
}

func (dec *Decoder) nextUint() uint64 {
//line /usr/local/go/src/encoding/gob/decoder.go:127
	_go_fuzz_dep_.CoverTab[84577]++
							n, _, err := decodeUintReader(&dec.buf, dec.countBuf)
							if err != nil {
//line /usr/local/go/src/encoding/gob/decoder.go:129
		_go_fuzz_dep_.CoverTab[84579]++
								dec.err = err
//line /usr/local/go/src/encoding/gob/decoder.go:130
		// _ = "end of CoverTab[84579]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:131
		_go_fuzz_dep_.CoverTab[84580]++
//line /usr/local/go/src/encoding/gob/decoder.go:131
		// _ = "end of CoverTab[84580]"
//line /usr/local/go/src/encoding/gob/decoder.go:131
	}
//line /usr/local/go/src/encoding/gob/decoder.go:131
	// _ = "end of CoverTab[84577]"
//line /usr/local/go/src/encoding/gob/decoder.go:131
	_go_fuzz_dep_.CoverTab[84578]++
							return n
//line /usr/local/go/src/encoding/gob/decoder.go:132
	// _ = "end of CoverTab[84578]"
}

// decodeTypeSequence parses:
//line /usr/local/go/src/encoding/gob/decoder.go:135
// TypeSequence
//line /usr/local/go/src/encoding/gob/decoder.go:135
//
//line /usr/local/go/src/encoding/gob/decoder.go:135
//	(TypeDefinition DelimitedTypeDefinition*)?
//line /usr/local/go/src/encoding/gob/decoder.go:135
//
//line /usr/local/go/src/encoding/gob/decoder.go:135
// and returns the type id of the next value. It returns -1 at
//line /usr/local/go/src/encoding/gob/decoder.go:135
// EOF.  Upon return, the remainder of dec.buf is the value to be
//line /usr/local/go/src/encoding/gob/decoder.go:135
// decoded. If this is an interface value, it can be ignored by
//line /usr/local/go/src/encoding/gob/decoder.go:135
// resetting that buffer.
//line /usr/local/go/src/encoding/gob/decoder.go:144
func (dec *Decoder) decodeTypeSequence(isInterface bool) typeId {
//line /usr/local/go/src/encoding/gob/decoder.go:144
	_go_fuzz_dep_.CoverTab[84581]++
							firstMessage := true
							for dec.err == nil {
//line /usr/local/go/src/encoding/gob/decoder.go:146
		_go_fuzz_dep_.CoverTab[84583]++
								if dec.buf.Len() == 0 {
//line /usr/local/go/src/encoding/gob/decoder.go:147
			_go_fuzz_dep_.CoverTab[84588]++
									if !dec.recvMessage() {
//line /usr/local/go/src/encoding/gob/decoder.go:148
				_go_fuzz_dep_.CoverTab[84589]++

//line /usr/local/go/src/encoding/gob/decoder.go:153
				if !firstMessage && func() bool {
//line /usr/local/go/src/encoding/gob/decoder.go:153
					_go_fuzz_dep_.CoverTab[84591]++
//line /usr/local/go/src/encoding/gob/decoder.go:153
					return dec.err == io.EOF
//line /usr/local/go/src/encoding/gob/decoder.go:153
					// _ = "end of CoverTab[84591]"
//line /usr/local/go/src/encoding/gob/decoder.go:153
				}() {
//line /usr/local/go/src/encoding/gob/decoder.go:153
					_go_fuzz_dep_.CoverTab[84592]++
											dec.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/gob/decoder.go:154
					// _ = "end of CoverTab[84592]"
				} else {
//line /usr/local/go/src/encoding/gob/decoder.go:155
					_go_fuzz_dep_.CoverTab[84593]++
//line /usr/local/go/src/encoding/gob/decoder.go:155
					// _ = "end of CoverTab[84593]"
//line /usr/local/go/src/encoding/gob/decoder.go:155
				}
//line /usr/local/go/src/encoding/gob/decoder.go:155
				// _ = "end of CoverTab[84589]"
//line /usr/local/go/src/encoding/gob/decoder.go:155
				_go_fuzz_dep_.CoverTab[84590]++
										break
//line /usr/local/go/src/encoding/gob/decoder.go:156
				// _ = "end of CoverTab[84590]"
			} else {
//line /usr/local/go/src/encoding/gob/decoder.go:157
				_go_fuzz_dep_.CoverTab[84594]++
//line /usr/local/go/src/encoding/gob/decoder.go:157
				// _ = "end of CoverTab[84594]"
//line /usr/local/go/src/encoding/gob/decoder.go:157
			}
//line /usr/local/go/src/encoding/gob/decoder.go:157
			// _ = "end of CoverTab[84588]"
		} else {
//line /usr/local/go/src/encoding/gob/decoder.go:158
			_go_fuzz_dep_.CoverTab[84595]++
//line /usr/local/go/src/encoding/gob/decoder.go:158
			// _ = "end of CoverTab[84595]"
//line /usr/local/go/src/encoding/gob/decoder.go:158
		}
//line /usr/local/go/src/encoding/gob/decoder.go:158
		// _ = "end of CoverTab[84583]"
//line /usr/local/go/src/encoding/gob/decoder.go:158
		_go_fuzz_dep_.CoverTab[84584]++

								id := typeId(dec.nextInt())
								if id >= 0 {
//line /usr/local/go/src/encoding/gob/decoder.go:161
			_go_fuzz_dep_.CoverTab[84596]++

									return id
//line /usr/local/go/src/encoding/gob/decoder.go:163
			// _ = "end of CoverTab[84596]"
		} else {
//line /usr/local/go/src/encoding/gob/decoder.go:164
			_go_fuzz_dep_.CoverTab[84597]++
//line /usr/local/go/src/encoding/gob/decoder.go:164
			// _ = "end of CoverTab[84597]"
//line /usr/local/go/src/encoding/gob/decoder.go:164
		}
//line /usr/local/go/src/encoding/gob/decoder.go:164
		// _ = "end of CoverTab[84584]"
//line /usr/local/go/src/encoding/gob/decoder.go:164
		_go_fuzz_dep_.CoverTab[84585]++

								dec.recvType(-id)
								if dec.err != nil {
//line /usr/local/go/src/encoding/gob/decoder.go:167
			_go_fuzz_dep_.CoverTab[84598]++
									break
//line /usr/local/go/src/encoding/gob/decoder.go:168
			// _ = "end of CoverTab[84598]"
		} else {
//line /usr/local/go/src/encoding/gob/decoder.go:169
			_go_fuzz_dep_.CoverTab[84599]++
//line /usr/local/go/src/encoding/gob/decoder.go:169
			// _ = "end of CoverTab[84599]"
//line /usr/local/go/src/encoding/gob/decoder.go:169
		}
//line /usr/local/go/src/encoding/gob/decoder.go:169
		// _ = "end of CoverTab[84585]"
//line /usr/local/go/src/encoding/gob/decoder.go:169
		_go_fuzz_dep_.CoverTab[84586]++

//line /usr/local/go/src/encoding/gob/decoder.go:174
		if dec.buf.Len() > 0 {
//line /usr/local/go/src/encoding/gob/decoder.go:174
			_go_fuzz_dep_.CoverTab[84600]++
									if !isInterface {
//line /usr/local/go/src/encoding/gob/decoder.go:175
				_go_fuzz_dep_.CoverTab[84602]++
										dec.err = errors.New("extra data in buffer")
										break
//line /usr/local/go/src/encoding/gob/decoder.go:177
				// _ = "end of CoverTab[84602]"
			} else {
//line /usr/local/go/src/encoding/gob/decoder.go:178
				_go_fuzz_dep_.CoverTab[84603]++
//line /usr/local/go/src/encoding/gob/decoder.go:178
				// _ = "end of CoverTab[84603]"
//line /usr/local/go/src/encoding/gob/decoder.go:178
			}
//line /usr/local/go/src/encoding/gob/decoder.go:178
			// _ = "end of CoverTab[84600]"
//line /usr/local/go/src/encoding/gob/decoder.go:178
			_go_fuzz_dep_.CoverTab[84601]++
									dec.nextUint()
//line /usr/local/go/src/encoding/gob/decoder.go:179
			// _ = "end of CoverTab[84601]"
		} else {
//line /usr/local/go/src/encoding/gob/decoder.go:180
			_go_fuzz_dep_.CoverTab[84604]++
//line /usr/local/go/src/encoding/gob/decoder.go:180
			// _ = "end of CoverTab[84604]"
//line /usr/local/go/src/encoding/gob/decoder.go:180
		}
//line /usr/local/go/src/encoding/gob/decoder.go:180
		// _ = "end of CoverTab[84586]"
//line /usr/local/go/src/encoding/gob/decoder.go:180
		_go_fuzz_dep_.CoverTab[84587]++
								firstMessage = false
//line /usr/local/go/src/encoding/gob/decoder.go:181
		// _ = "end of CoverTab[84587]"
	}
//line /usr/local/go/src/encoding/gob/decoder.go:182
	// _ = "end of CoverTab[84581]"
//line /usr/local/go/src/encoding/gob/decoder.go:182
	_go_fuzz_dep_.CoverTab[84582]++
							return -1
//line /usr/local/go/src/encoding/gob/decoder.go:183
	// _ = "end of CoverTab[84582]"
}

// Decode reads the next value from the input stream and stores
//line /usr/local/go/src/encoding/gob/decoder.go:186
// it in the data represented by the empty interface value.
//line /usr/local/go/src/encoding/gob/decoder.go:186
// If e is nil, the value will be discarded. Otherwise,
//line /usr/local/go/src/encoding/gob/decoder.go:186
// the value underlying e must be a pointer to the
//line /usr/local/go/src/encoding/gob/decoder.go:186
// correct type for the next data item received.
//line /usr/local/go/src/encoding/gob/decoder.go:186
// If the input is at EOF, Decode returns io.EOF and
//line /usr/local/go/src/encoding/gob/decoder.go:186
// does not modify e.
//line /usr/local/go/src/encoding/gob/decoder.go:193
func (dec *Decoder) Decode(e any) error {
//line /usr/local/go/src/encoding/gob/decoder.go:193
	_go_fuzz_dep_.CoverTab[84605]++
							if e == nil {
//line /usr/local/go/src/encoding/gob/decoder.go:194
		_go_fuzz_dep_.CoverTab[84608]++
								return dec.DecodeValue(reflect.Value{})
//line /usr/local/go/src/encoding/gob/decoder.go:195
		// _ = "end of CoverTab[84608]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:196
		_go_fuzz_dep_.CoverTab[84609]++
//line /usr/local/go/src/encoding/gob/decoder.go:196
		// _ = "end of CoverTab[84609]"
//line /usr/local/go/src/encoding/gob/decoder.go:196
	}
//line /usr/local/go/src/encoding/gob/decoder.go:196
	// _ = "end of CoverTab[84605]"
//line /usr/local/go/src/encoding/gob/decoder.go:196
	_go_fuzz_dep_.CoverTab[84606]++
							value := reflect.ValueOf(e)

//line /usr/local/go/src/encoding/gob/decoder.go:200
	if value.Type().Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/gob/decoder.go:200
		_go_fuzz_dep_.CoverTab[84610]++
								dec.err = errors.New("gob: attempt to decode into a non-pointer")
								return dec.err
//line /usr/local/go/src/encoding/gob/decoder.go:202
		// _ = "end of CoverTab[84610]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:203
		_go_fuzz_dep_.CoverTab[84611]++
//line /usr/local/go/src/encoding/gob/decoder.go:203
		// _ = "end of CoverTab[84611]"
//line /usr/local/go/src/encoding/gob/decoder.go:203
	}
//line /usr/local/go/src/encoding/gob/decoder.go:203
	// _ = "end of CoverTab[84606]"
//line /usr/local/go/src/encoding/gob/decoder.go:203
	_go_fuzz_dep_.CoverTab[84607]++
							return dec.DecodeValue(value)
//line /usr/local/go/src/encoding/gob/decoder.go:204
	// _ = "end of CoverTab[84607]"
}

// DecodeValue reads the next value from the input stream.
//line /usr/local/go/src/encoding/gob/decoder.go:207
// If v is the zero reflect.Value (v.Kind() == Invalid), DecodeValue discards the value.
//line /usr/local/go/src/encoding/gob/decoder.go:207
// Otherwise, it stores the value into v. In that case, v must represent
//line /usr/local/go/src/encoding/gob/decoder.go:207
// a non-nil pointer to data or be an assignable reflect.Value (v.CanSet())
//line /usr/local/go/src/encoding/gob/decoder.go:207
// If the input is at EOF, DecodeValue returns io.EOF and
//line /usr/local/go/src/encoding/gob/decoder.go:207
// does not modify v.
//line /usr/local/go/src/encoding/gob/decoder.go:213
func (dec *Decoder) DecodeValue(v reflect.Value) error {
//line /usr/local/go/src/encoding/gob/decoder.go:213
	_go_fuzz_dep_.CoverTab[84612]++
							if v.IsValid() {
//line /usr/local/go/src/encoding/gob/decoder.go:214
		_go_fuzz_dep_.CoverTab[84615]++
								if v.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/gob/decoder.go:215
			_go_fuzz_dep_.CoverTab[84616]++
//line /usr/local/go/src/encoding/gob/decoder.go:215
			return !v.IsNil()
//line /usr/local/go/src/encoding/gob/decoder.go:215
			// _ = "end of CoverTab[84616]"
//line /usr/local/go/src/encoding/gob/decoder.go:215
		}() {
//line /usr/local/go/src/encoding/gob/decoder.go:215
			_go_fuzz_dep_.CoverTab[84617]++
//line /usr/local/go/src/encoding/gob/decoder.go:215
			// _ = "end of CoverTab[84617]"

		} else {
//line /usr/local/go/src/encoding/gob/decoder.go:217
			_go_fuzz_dep_.CoverTab[84618]++
//line /usr/local/go/src/encoding/gob/decoder.go:217
			if !v.CanSet() {
//line /usr/local/go/src/encoding/gob/decoder.go:217
				_go_fuzz_dep_.CoverTab[84619]++
										return errors.New("gob: DecodeValue of unassignable value")
//line /usr/local/go/src/encoding/gob/decoder.go:218
				// _ = "end of CoverTab[84619]"
			} else {
//line /usr/local/go/src/encoding/gob/decoder.go:219
				_go_fuzz_dep_.CoverTab[84620]++
//line /usr/local/go/src/encoding/gob/decoder.go:219
				// _ = "end of CoverTab[84620]"
//line /usr/local/go/src/encoding/gob/decoder.go:219
			}
//line /usr/local/go/src/encoding/gob/decoder.go:219
			// _ = "end of CoverTab[84618]"
//line /usr/local/go/src/encoding/gob/decoder.go:219
		}
//line /usr/local/go/src/encoding/gob/decoder.go:219
		// _ = "end of CoverTab[84615]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:220
		_go_fuzz_dep_.CoverTab[84621]++
//line /usr/local/go/src/encoding/gob/decoder.go:220
		// _ = "end of CoverTab[84621]"
//line /usr/local/go/src/encoding/gob/decoder.go:220
	}
//line /usr/local/go/src/encoding/gob/decoder.go:220
	// _ = "end of CoverTab[84612]"
//line /usr/local/go/src/encoding/gob/decoder.go:220
	_go_fuzz_dep_.CoverTab[84613]++

							dec.mutex.Lock()
							defer dec.mutex.Unlock()

							dec.buf.Reset()
							dec.err = nil
							id := dec.decodeTypeSequence(false)
							if dec.err == nil {
//line /usr/local/go/src/encoding/gob/decoder.go:228
		_go_fuzz_dep_.CoverTab[84622]++
								dec.decodeValue(id, v)
//line /usr/local/go/src/encoding/gob/decoder.go:229
		// _ = "end of CoverTab[84622]"
	} else {
//line /usr/local/go/src/encoding/gob/decoder.go:230
		_go_fuzz_dep_.CoverTab[84623]++
//line /usr/local/go/src/encoding/gob/decoder.go:230
		// _ = "end of CoverTab[84623]"
//line /usr/local/go/src/encoding/gob/decoder.go:230
	}
//line /usr/local/go/src/encoding/gob/decoder.go:230
	// _ = "end of CoverTab[84613]"
//line /usr/local/go/src/encoding/gob/decoder.go:230
	_go_fuzz_dep_.CoverTab[84614]++
							return dec.err
//line /usr/local/go/src/encoding/gob/decoder.go:231
	// _ = "end of CoverTab[84614]"
}

// If debug.go is compiled into the program, debugFunc prints a human-readable
//line /usr/local/go/src/encoding/gob/decoder.go:234
// representation of the gob data read from r by calling that file's Debug function.
//line /usr/local/go/src/encoding/gob/decoder.go:234
// Otherwise it is nil.
//line /usr/local/go/src/encoding/gob/decoder.go:237
var debugFunc func(io.Reader)
//line /usr/local/go/src/encoding/gob/decoder.go:237
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/decoder.go:237
var _ = _go_fuzz_dep_.CoverTab
