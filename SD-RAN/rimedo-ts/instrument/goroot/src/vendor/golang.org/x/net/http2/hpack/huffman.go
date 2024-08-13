// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
package hpack

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:5
)

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:15
		_go_fuzz_dep_.CoverTab[35328]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:15
		return new(bytes.Buffer)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:15
		// _ = "end of CoverTab[35328]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:15
	},
}

// HuffmanDecode decodes the string in v and writes the expanded
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:18
// result to w, returning the number of bytes written to w and the
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:18
// Write call's return value. At most one Write call is made.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:21
func HuffmanDecode(w io.Writer, v []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:21
	_go_fuzz_dep_.CoverTab[35329]++
										buf := bufPool.Get().(*bytes.Buffer)
										buf.Reset()
										defer bufPool.Put(buf)
										if err := huffmanDecode(buf, 0, v); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:25
		_go_fuzz_dep_.CoverTab[35331]++
											return 0, err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:26
		// _ = "end of CoverTab[35331]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:27
		_go_fuzz_dep_.CoverTab[35332]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:27
		// _ = "end of CoverTab[35332]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:27
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:27
	// _ = "end of CoverTab[35329]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:27
	_go_fuzz_dep_.CoverTab[35330]++
										return w.Write(buf.Bytes())
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:28
	// _ = "end of CoverTab[35330]"
}

// HuffmanDecodeToString decodes the string in v.
func HuffmanDecodeToString(v []byte) (string, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:32
	_go_fuzz_dep_.CoverTab[35333]++
										buf := bufPool.Get().(*bytes.Buffer)
										buf.Reset()
										defer bufPool.Put(buf)
										if err := huffmanDecode(buf, 0, v); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:36
		_go_fuzz_dep_.CoverTab[35335]++
											return "", err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:37
		// _ = "end of CoverTab[35335]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:38
		_go_fuzz_dep_.CoverTab[35336]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:38
		// _ = "end of CoverTab[35336]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:38
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:38
	// _ = "end of CoverTab[35333]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:38
	_go_fuzz_dep_.CoverTab[35334]++
										return buf.String(), nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:39
	// _ = "end of CoverTab[35334]"
}

// ErrInvalidHuffman is returned for errors found decoding
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:42
// Huffman-encoded strings.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:44
var ErrInvalidHuffman = errors.New("hpack: invalid Huffman-encoded data")

// huffmanDecode decodes v to buf.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:46
// If maxLen is greater than 0, attempts to write more to buf than
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:46
// maxLen bytes will return ErrStringLength.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:49
func huffmanDecode(buf *bytes.Buffer, maxLen int, v []byte) error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:49
	_go_fuzz_dep_.CoverTab[35337]++
										rootHuffmanNode := getRootHuffmanNode()
										n := rootHuffmanNode

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:55
	cur, cbits, sbits := uint(0), uint8(0), uint8(0)
	for _, b := range v {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:56
		_go_fuzz_dep_.CoverTab[35342]++
											cur = cur<<8 | uint(b)
											cbits += 8
											sbits += 8
											for cbits >= 8 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:60
			_go_fuzz_dep_.CoverTab[35343]++
												idx := byte(cur >> (cbits - 8))
												n = n.children[idx]
												if n == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:63
				_go_fuzz_dep_.CoverTab[35345]++
													return ErrInvalidHuffman
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:64
				// _ = "end of CoverTab[35345]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:65
				_go_fuzz_dep_.CoverTab[35346]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:65
				// _ = "end of CoverTab[35346]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:65
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:65
			// _ = "end of CoverTab[35343]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:65
			_go_fuzz_dep_.CoverTab[35344]++
												if n.children == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:66
				_go_fuzz_dep_.CoverTab[35347]++
													if maxLen != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:67
					_go_fuzz_dep_.CoverTab[35349]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:67
					return buf.Len() == maxLen
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:67
					// _ = "end of CoverTab[35349]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:67
				}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:67
					_go_fuzz_dep_.CoverTab[35350]++
														return ErrStringLength
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:68
					// _ = "end of CoverTab[35350]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:69
					_go_fuzz_dep_.CoverTab[35351]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:69
					// _ = "end of CoverTab[35351]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:69
				}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:69
				// _ = "end of CoverTab[35347]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:69
				_go_fuzz_dep_.CoverTab[35348]++
													buf.WriteByte(n.sym)
													cbits -= n.codeLen
													n = rootHuffmanNode
													sbits = cbits
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:73
				// _ = "end of CoverTab[35348]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:74
				_go_fuzz_dep_.CoverTab[35352]++
													cbits -= 8
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:75
				// _ = "end of CoverTab[35352]"
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:76
			// _ = "end of CoverTab[35344]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:77
		// _ = "end of CoverTab[35342]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:78
	// _ = "end of CoverTab[35337]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:78
	_go_fuzz_dep_.CoverTab[35338]++
										for cbits > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:79
		_go_fuzz_dep_.CoverTab[35353]++
											n = n.children[byte(cur<<(8-cbits))]
											if n == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:81
			_go_fuzz_dep_.CoverTab[35357]++
												return ErrInvalidHuffman
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:82
			// _ = "end of CoverTab[35357]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:83
			_go_fuzz_dep_.CoverTab[35358]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:83
			// _ = "end of CoverTab[35358]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:83
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:83
		// _ = "end of CoverTab[35353]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:83
		_go_fuzz_dep_.CoverTab[35354]++
											if n.children != nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:84
			_go_fuzz_dep_.CoverTab[35359]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:84
			return n.codeLen > cbits
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:84
			// _ = "end of CoverTab[35359]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:84
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:84
			_go_fuzz_dep_.CoverTab[35360]++
												break
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:85
			// _ = "end of CoverTab[35360]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:86
			_go_fuzz_dep_.CoverTab[35361]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:86
			// _ = "end of CoverTab[35361]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:86
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:86
		// _ = "end of CoverTab[35354]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:86
		_go_fuzz_dep_.CoverTab[35355]++
											if maxLen != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:87
			_go_fuzz_dep_.CoverTab[35362]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:87
			return buf.Len() == maxLen
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:87
			// _ = "end of CoverTab[35362]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:87
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:87
			_go_fuzz_dep_.CoverTab[35363]++
												return ErrStringLength
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:88
			// _ = "end of CoverTab[35363]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:89
			_go_fuzz_dep_.CoverTab[35364]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:89
			// _ = "end of CoverTab[35364]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:89
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:89
		// _ = "end of CoverTab[35355]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:89
		_go_fuzz_dep_.CoverTab[35356]++
											buf.WriteByte(n.sym)
											cbits -= n.codeLen
											n = rootHuffmanNode
											sbits = cbits
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:93
		// _ = "end of CoverTab[35356]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:94
	// _ = "end of CoverTab[35338]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:94
	_go_fuzz_dep_.CoverTab[35339]++
										if sbits > 7 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:95
		_go_fuzz_dep_.CoverTab[35365]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:98
		return ErrInvalidHuffman
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:98
		// _ = "end of CoverTab[35365]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:99
		_go_fuzz_dep_.CoverTab[35366]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:99
		// _ = "end of CoverTab[35366]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:99
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:99
	// _ = "end of CoverTab[35339]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:99
	_go_fuzz_dep_.CoverTab[35340]++
										if mask := uint(1<<cbits - 1); cur&mask != mask {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:100
		_go_fuzz_dep_.CoverTab[35367]++

											return ErrInvalidHuffman
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:102
		// _ = "end of CoverTab[35367]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:103
		_go_fuzz_dep_.CoverTab[35368]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:103
		// _ = "end of CoverTab[35368]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:103
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:103
	// _ = "end of CoverTab[35340]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:103
	_go_fuzz_dep_.CoverTab[35341]++

										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:105
	// _ = "end of CoverTab[35341]"
}

// incomparable is a zero-width, non-comparable type. Adding it to a struct
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:108
// makes that struct also non-comparable, and generally doesn't add
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:108
// any size (as long as it's first).
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:111
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
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:124
	_go_fuzz_dep_.CoverTab[35369]++
										return &node{children: new([256]*node)}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:125
	// _ = "end of CoverTab[35369]"
}

var (
	buildRootOnce		sync.Once
	lazyRootHuffmanNode	*node
)

func getRootHuffmanNode() *node {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:133
	_go_fuzz_dep_.CoverTab[35370]++
										buildRootOnce.Do(buildRootHuffmanNode)
										return lazyRootHuffmanNode
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:135
	// _ = "end of CoverTab[35370]"
}

func buildRootHuffmanNode() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:138
	_go_fuzz_dep_.CoverTab[35371]++
										if len(huffmanCodes) != 256 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:139
		_go_fuzz_dep_.CoverTab[35373]++
											panic("unexpected size")
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:140
		// _ = "end of CoverTab[35373]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:141
		_go_fuzz_dep_.CoverTab[35374]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:141
		// _ = "end of CoverTab[35374]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:141
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:141
	// _ = "end of CoverTab[35371]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:141
	_go_fuzz_dep_.CoverTab[35372]++
										lazyRootHuffmanNode = newInternalNode()

										leaves := new([256]node)

										for sym, code := range huffmanCodes {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:146
		_go_fuzz_dep_.CoverTab[35375]++
											codeLen := huffmanCodeLen[sym]

											cur := lazyRootHuffmanNode
											for codeLen > 8 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:150
			_go_fuzz_dep_.CoverTab[35377]++
												codeLen -= 8
												i := uint8(code >> codeLen)
												if cur.children[i] == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:153
				_go_fuzz_dep_.CoverTab[35379]++
													cur.children[i] = newInternalNode()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:154
				// _ = "end of CoverTab[35379]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:155
				_go_fuzz_dep_.CoverTab[35380]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:155
				// _ = "end of CoverTab[35380]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:155
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:155
			// _ = "end of CoverTab[35377]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:155
			_go_fuzz_dep_.CoverTab[35378]++
												cur = cur.children[i]
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:156
			// _ = "end of CoverTab[35378]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:157
		// _ = "end of CoverTab[35375]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:157
		_go_fuzz_dep_.CoverTab[35376]++
											shift := 8 - codeLen
											start, end := int(uint8(code<<shift)), int(1<<shift)

											leaves[sym].sym = byte(sym)
											leaves[sym].codeLen = codeLen
											for i := start; i < start+end; i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:163
			_go_fuzz_dep_.CoverTab[35381]++
												cur.children[i] = &leaves[sym]
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:164
			// _ = "end of CoverTab[35381]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:165
		// _ = "end of CoverTab[35376]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:166
	// _ = "end of CoverTab[35372]"
}

// AppendHuffmanString appends s, as encoded in Huffman codes, to dst
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:169
// and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:171
func AppendHuffmanString(dst []byte, s string) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:171
	_go_fuzz_dep_.CoverTab[35382]++
	// This relies on the maximum huffman code length being 30 (See tables.go huffmanCodeLen array)
	// So if a uint64 buffer has less than 32 valid bits can always accommodate another huffmanCode.
	var (
		x	uint64	// buffer
		n	uint	// number valid of bits present in x
	)
	for i := 0; i < len(s); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:178
		_go_fuzz_dep_.CoverTab[35386]++
											c := s[i]
											n += uint(huffmanCodeLen[c])
											x <<= huffmanCodeLen[c] % 64
											x |= uint64(huffmanCodes[c])
											if n >= 32 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:183
			_go_fuzz_dep_.CoverTab[35387]++
												n %= 32
												y := uint32(x >> n)
												dst = append(dst, byte(y>>24), byte(y>>16), byte(y>>8), byte(y))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:186
			// _ = "end of CoverTab[35387]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:187
			_go_fuzz_dep_.CoverTab[35388]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:187
			// _ = "end of CoverTab[35388]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:187
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:187
		// _ = "end of CoverTab[35386]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:188
	// _ = "end of CoverTab[35382]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:188
	_go_fuzz_dep_.CoverTab[35383]++

										if over := n % 8; over > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:190
		_go_fuzz_dep_.CoverTab[35389]++
											const (
			eosCode		= 0x3fffffff
			eosNBits	= 30
			eosPadByte	= eosCode >> (eosNBits - 8)
		)
											pad := 8 - over
											x = (x << pad) | (eosPadByte >> over)
											n += pad
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:198
		// _ = "end of CoverTab[35389]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:199
		_go_fuzz_dep_.CoverTab[35390]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:199
		// _ = "end of CoverTab[35390]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:199
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:199
	// _ = "end of CoverTab[35383]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:199
	_go_fuzz_dep_.CoverTab[35384]++

										switch n / 8 {
	case 0:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:202
		_go_fuzz_dep_.CoverTab[35391]++
											return dst
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:203
		// _ = "end of CoverTab[35391]"
	case 1:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:204
		_go_fuzz_dep_.CoverTab[35392]++
											return append(dst, byte(x))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:205
		// _ = "end of CoverTab[35392]"
	case 2:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:206
		_go_fuzz_dep_.CoverTab[35393]++
											y := uint16(x)
											return append(dst, byte(y>>8), byte(y))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:208
		// _ = "end of CoverTab[35393]"
	case 3:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:209
		_go_fuzz_dep_.CoverTab[35394]++
											y := uint16(x >> 8)
											return append(dst, byte(y>>8), byte(y), byte(x))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:211
		// _ = "end of CoverTab[35394]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:211
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:211
		_go_fuzz_dep_.CoverTab[35395]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:211
		// _ = "end of CoverTab[35395]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:212
	// _ = "end of CoverTab[35384]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:212
	_go_fuzz_dep_.CoverTab[35385]++

										y := uint32(x)
										return append(dst, byte(y>>24), byte(y>>16), byte(y>>8), byte(y))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:215
	// _ = "end of CoverTab[35385]"
}

// HuffmanEncodeLength returns the number of bytes required to encode
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:218
// s in Huffman codes. The result is round up to byte boundary.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:220
func HuffmanEncodeLength(s string) uint64 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:220
	_go_fuzz_dep_.CoverTab[35396]++
										n := uint64(0)
										for i := 0; i < len(s); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:222
		_go_fuzz_dep_.CoverTab[35398]++
											n += uint64(huffmanCodeLen[s[i]])
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:223
		// _ = "end of CoverTab[35398]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:224
	// _ = "end of CoverTab[35396]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:224
	_go_fuzz_dep_.CoverTab[35397]++
										return (n + 7) / 8
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:225
	// _ = "end of CoverTab[35397]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:226
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/huffman.go:226
var _ = _go_fuzz_dep_.CoverTab
