// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
package hpack

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:5
)

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:15
		_go_fuzz_dep_.CoverTab[72110]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:15
		return new(bytes.Buffer)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:15
		// _ = "end of CoverTab[72110]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:15
	},
}

// HuffmanDecode decodes the string in v and writes the expanded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:18
// result to w, returning the number of bytes written to w and the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:18
// Write call's return value. At most one Write call is made.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:21
func HuffmanDecode(w io.Writer, v []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:21
	_go_fuzz_dep_.CoverTab[72111]++
											buf := bufPool.Get().(*bytes.Buffer)
											buf.Reset()
											defer bufPool.Put(buf)
											if err := huffmanDecode(buf, 0, v); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:25
		_go_fuzz_dep_.CoverTab[72113]++
												return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:26
		// _ = "end of CoverTab[72113]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:27
		_go_fuzz_dep_.CoverTab[72114]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:27
		// _ = "end of CoverTab[72114]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:27
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:27
	// _ = "end of CoverTab[72111]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:27
	_go_fuzz_dep_.CoverTab[72112]++
											return w.Write(buf.Bytes())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:28
	// _ = "end of CoverTab[72112]"
}

// HuffmanDecodeToString decodes the string in v.
func HuffmanDecodeToString(v []byte) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:32
	_go_fuzz_dep_.CoverTab[72115]++
											buf := bufPool.Get().(*bytes.Buffer)
											buf.Reset()
											defer bufPool.Put(buf)
											if err := huffmanDecode(buf, 0, v); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:36
		_go_fuzz_dep_.CoverTab[72117]++
												return "", err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:37
		// _ = "end of CoverTab[72117]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:38
		_go_fuzz_dep_.CoverTab[72118]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:38
		// _ = "end of CoverTab[72118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:38
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:38
	// _ = "end of CoverTab[72115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:38
	_go_fuzz_dep_.CoverTab[72116]++
											return buf.String(), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:39
	// _ = "end of CoverTab[72116]"
}

// ErrInvalidHuffman is returned for errors found decoding
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:42
// Huffman-encoded strings.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:44
var ErrInvalidHuffman = errors.New("hpack: invalid Huffman-encoded data")

// huffmanDecode decodes v to buf.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:46
// If maxLen is greater than 0, attempts to write more to buf than
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:46
// maxLen bytes will return ErrStringLength.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:49
func huffmanDecode(buf *bytes.Buffer, maxLen int, v []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:49
	_go_fuzz_dep_.CoverTab[72119]++
											rootHuffmanNode := getRootHuffmanNode()
											n := rootHuffmanNode

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:55
	cur, cbits, sbits := uint(0), uint8(0), uint8(0)
	for _, b := range v {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:56
		_go_fuzz_dep_.CoverTab[72124]++
												cur = cur<<8 | uint(b)
												cbits += 8
												sbits += 8
												for cbits >= 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:60
			_go_fuzz_dep_.CoverTab[72125]++
													idx := byte(cur >> (cbits - 8))
													n = n.children[idx]
													if n == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:63
				_go_fuzz_dep_.CoverTab[72127]++
														return ErrInvalidHuffman
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:64
				// _ = "end of CoverTab[72127]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:65
				_go_fuzz_dep_.CoverTab[72128]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:65
				// _ = "end of CoverTab[72128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:65
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:65
			// _ = "end of CoverTab[72125]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:65
			_go_fuzz_dep_.CoverTab[72126]++
													if n.children == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:66
				_go_fuzz_dep_.CoverTab[72129]++
														if maxLen != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:67
					_go_fuzz_dep_.CoverTab[72131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:67
					return buf.Len() == maxLen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:67
					// _ = "end of CoverTab[72131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:67
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:67
					_go_fuzz_dep_.CoverTab[72132]++
															return ErrStringLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:68
					// _ = "end of CoverTab[72132]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:69
					_go_fuzz_dep_.CoverTab[72133]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:69
					// _ = "end of CoverTab[72133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:69
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:69
				// _ = "end of CoverTab[72129]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:69
				_go_fuzz_dep_.CoverTab[72130]++
														buf.WriteByte(n.sym)
														cbits -= n.codeLen
														n = rootHuffmanNode
														sbits = cbits
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:73
				// _ = "end of CoverTab[72130]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:74
				_go_fuzz_dep_.CoverTab[72134]++
														cbits -= 8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:75
				// _ = "end of CoverTab[72134]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:76
			// _ = "end of CoverTab[72126]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:77
		// _ = "end of CoverTab[72124]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:78
	// _ = "end of CoverTab[72119]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:78
	_go_fuzz_dep_.CoverTab[72120]++
											for cbits > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:79
		_go_fuzz_dep_.CoverTab[72135]++
												n = n.children[byte(cur<<(8-cbits))]
												if n == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:81
			_go_fuzz_dep_.CoverTab[72139]++
													return ErrInvalidHuffman
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:82
			// _ = "end of CoverTab[72139]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:83
			_go_fuzz_dep_.CoverTab[72140]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:83
			// _ = "end of CoverTab[72140]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:83
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:83
		// _ = "end of CoverTab[72135]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:83
		_go_fuzz_dep_.CoverTab[72136]++
												if n.children != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:84
			_go_fuzz_dep_.CoverTab[72141]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:84
			return n.codeLen > cbits
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:84
			// _ = "end of CoverTab[72141]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:84
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:84
			_go_fuzz_dep_.CoverTab[72142]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:85
			// _ = "end of CoverTab[72142]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:86
			_go_fuzz_dep_.CoverTab[72143]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:86
			// _ = "end of CoverTab[72143]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:86
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:86
		// _ = "end of CoverTab[72136]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:86
		_go_fuzz_dep_.CoverTab[72137]++
												if maxLen != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:87
			_go_fuzz_dep_.CoverTab[72144]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:87
			return buf.Len() == maxLen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:87
			// _ = "end of CoverTab[72144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:87
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:87
			_go_fuzz_dep_.CoverTab[72145]++
													return ErrStringLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:88
			// _ = "end of CoverTab[72145]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:89
			_go_fuzz_dep_.CoverTab[72146]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:89
			// _ = "end of CoverTab[72146]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:89
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:89
		// _ = "end of CoverTab[72137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:89
		_go_fuzz_dep_.CoverTab[72138]++
												buf.WriteByte(n.sym)
												cbits -= n.codeLen
												n = rootHuffmanNode
												sbits = cbits
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:93
		// _ = "end of CoverTab[72138]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:94
	// _ = "end of CoverTab[72120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:94
	_go_fuzz_dep_.CoverTab[72121]++
											if sbits > 7 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:95
		_go_fuzz_dep_.CoverTab[72147]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:98
		return ErrInvalidHuffman
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:98
		// _ = "end of CoverTab[72147]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:99
		_go_fuzz_dep_.CoverTab[72148]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:99
		// _ = "end of CoverTab[72148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:99
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:99
	// _ = "end of CoverTab[72121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:99
	_go_fuzz_dep_.CoverTab[72122]++
											if mask := uint(1<<cbits - 1); cur&mask != mask {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:100
		_go_fuzz_dep_.CoverTab[72149]++

												return ErrInvalidHuffman
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:102
		// _ = "end of CoverTab[72149]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:103
		_go_fuzz_dep_.CoverTab[72150]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:103
		// _ = "end of CoverTab[72150]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:103
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:103
	// _ = "end of CoverTab[72122]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:103
	_go_fuzz_dep_.CoverTab[72123]++

											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:105
	// _ = "end of CoverTab[72123]"
}

// incomparable is a zero-width, non-comparable type. Adding it to a struct
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:108
// makes that struct also non-comparable, and generally doesn't add
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:108
// any size (as long as it's first).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:111
type incomparable [0]func()

type node struct {
	_	incomparable

	// children is non-nil for internal nodes
	children	*[256]*node

	// The following are only valid if children is nil:
	codeLen	uint8	// number of bits that led to the output of sym
	sym	byte	// output symbol
}

func newInternalNode() *node {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:124
	_go_fuzz_dep_.CoverTab[72151]++
											return &node{children: new([256]*node)}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:125
	// _ = "end of CoverTab[72151]"
}

var (
	buildRootOnce		sync.Once
	lazyRootHuffmanNode	*node
)

func getRootHuffmanNode() *node {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:133
	_go_fuzz_dep_.CoverTab[72152]++
											buildRootOnce.Do(buildRootHuffmanNode)
											return lazyRootHuffmanNode
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:135
	// _ = "end of CoverTab[72152]"
}

func buildRootHuffmanNode() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:138
	_go_fuzz_dep_.CoverTab[72153]++
											if len(huffmanCodes) != 256 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:139
		_go_fuzz_dep_.CoverTab[72155]++
												panic("unexpected size")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:140
		// _ = "end of CoverTab[72155]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:141
		_go_fuzz_dep_.CoverTab[72156]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:141
		// _ = "end of CoverTab[72156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:141
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:141
	// _ = "end of CoverTab[72153]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:141
	_go_fuzz_dep_.CoverTab[72154]++
											lazyRootHuffmanNode = newInternalNode()

											leaves := new([256]node)

											for sym, code := range huffmanCodes {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:146
		_go_fuzz_dep_.CoverTab[72157]++
												codeLen := huffmanCodeLen[sym]

												cur := lazyRootHuffmanNode
												for codeLen > 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:150
			_go_fuzz_dep_.CoverTab[72159]++
													codeLen -= 8
													i := uint8(code >> codeLen)
													if cur.children[i] == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:153
				_go_fuzz_dep_.CoverTab[72161]++
														cur.children[i] = newInternalNode()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:154
				// _ = "end of CoverTab[72161]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:155
				_go_fuzz_dep_.CoverTab[72162]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:155
				// _ = "end of CoverTab[72162]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:155
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:155
			// _ = "end of CoverTab[72159]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:155
			_go_fuzz_dep_.CoverTab[72160]++
													cur = cur.children[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:156
			// _ = "end of CoverTab[72160]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:157
		// _ = "end of CoverTab[72157]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:157
		_go_fuzz_dep_.CoverTab[72158]++
												shift := 8 - codeLen
												start, end := int(uint8(code<<shift)), int(1<<shift)

												leaves[sym].sym = byte(sym)
												leaves[sym].codeLen = codeLen
												for i := start; i < start+end; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:163
			_go_fuzz_dep_.CoverTab[72163]++
													cur.children[i] = &leaves[sym]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:164
			// _ = "end of CoverTab[72163]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:165
		// _ = "end of CoverTab[72158]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:166
	// _ = "end of CoverTab[72154]"
}

// AppendHuffmanString appends s, as encoded in Huffman codes, to dst
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:169
// and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:171
func AppendHuffmanString(dst []byte, s string) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:171
	_go_fuzz_dep_.CoverTab[72164]++
	// This relies on the maximum huffman code length being 30 (See tables.go huffmanCodeLen array)
	// So if a uint64 buffer has less than 32 valid bits can always accommodate another huffmanCode.
	var (
		x	uint64	// buffer
		n	uint	// number valid of bits present in x
	)
	for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:178
		_go_fuzz_dep_.CoverTab[72168]++
												c := s[i]
												n += uint(huffmanCodeLen[c])
												x <<= huffmanCodeLen[c] % 64
												x |= uint64(huffmanCodes[c])
												if n >= 32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:183
			_go_fuzz_dep_.CoverTab[72169]++
													n %= 32
													y := uint32(x >> n)
													dst = append(dst, byte(y>>24), byte(y>>16), byte(y>>8), byte(y))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:186
			// _ = "end of CoverTab[72169]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:187
			_go_fuzz_dep_.CoverTab[72170]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:187
			// _ = "end of CoverTab[72170]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:187
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:187
		// _ = "end of CoverTab[72168]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:188
	// _ = "end of CoverTab[72164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:188
	_go_fuzz_dep_.CoverTab[72165]++

											if over := n % 8; over > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:190
		_go_fuzz_dep_.CoverTab[72171]++
												const (
			eosCode		= 0x3fffffff
			eosNBits	= 30
			eosPadByte	= eosCode >> (eosNBits - 8)
		)
												pad := 8 - over
												x = (x << pad) | (eosPadByte >> over)
												n += pad
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:198
		// _ = "end of CoverTab[72171]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:199
		_go_fuzz_dep_.CoverTab[72172]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:199
		// _ = "end of CoverTab[72172]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:199
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:199
	// _ = "end of CoverTab[72165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:199
	_go_fuzz_dep_.CoverTab[72166]++

											switch n / 8 {
	case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:202
		_go_fuzz_dep_.CoverTab[72173]++
												return dst
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:203
		// _ = "end of CoverTab[72173]"
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:204
		_go_fuzz_dep_.CoverTab[72174]++
												return append(dst, byte(x))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:205
		// _ = "end of CoverTab[72174]"
	case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:206
		_go_fuzz_dep_.CoverTab[72175]++
												y := uint16(x)
												return append(dst, byte(y>>8), byte(y))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:208
		// _ = "end of CoverTab[72175]"
	case 3:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:209
		_go_fuzz_dep_.CoverTab[72176]++
												y := uint16(x >> 8)
												return append(dst, byte(y>>8), byte(y), byte(x))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:211
		// _ = "end of CoverTab[72176]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:211
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:211
		_go_fuzz_dep_.CoverTab[72177]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:211
		// _ = "end of CoverTab[72177]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:212
	// _ = "end of CoverTab[72166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:212
	_go_fuzz_dep_.CoverTab[72167]++

											y := uint32(x)
											return append(dst, byte(y>>24), byte(y>>16), byte(y>>8), byte(y))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:215
	// _ = "end of CoverTab[72167]"
}

// HuffmanEncodeLength returns the number of bytes required to encode
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:218
// s in Huffman codes. The result is round up to byte boundary.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:220
func HuffmanEncodeLength(s string) uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:220
	_go_fuzz_dep_.CoverTab[72178]++
											n := uint64(0)
											for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:222
		_go_fuzz_dep_.CoverTab[72180]++
												n += uint64(huffmanCodeLen[s[i]])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:223
		// _ = "end of CoverTab[72180]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:224
	// _ = "end of CoverTab[72178]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:224
	_go_fuzz_dep_.CoverTab[72179]++
											return (n + 7) / 8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:225
	// _ = "end of CoverTab[72179]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:226
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/huffman.go:226
var _ = _go_fuzz_dep_.CoverTab
