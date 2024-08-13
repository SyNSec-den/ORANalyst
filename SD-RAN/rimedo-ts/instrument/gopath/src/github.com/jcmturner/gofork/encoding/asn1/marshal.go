// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
package asn1

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:5
)

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"time"
	"unicode/utf8"
)

// A forkableWriter is an in-memory buffer that can be
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
// 'forked' to create new forkableWriters that bracket the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
// original. After
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
//	pre, post := w.fork()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:18
// the overall sequence of bytes represented is logically w+pre+post.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:23
type forkableWriter struct {
	*bytes.Buffer
	pre, post	*forkableWriter
}

func newForkableWriter() *forkableWriter {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:28
	_go_fuzz_dep_.CoverTab[82767]++
													return &forkableWriter{new(bytes.Buffer), nil, nil}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:29
	// _ = "end of CoverTab[82767]"
}

func (f *forkableWriter) fork() (pre, post *forkableWriter) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:32
	_go_fuzz_dep_.CoverTab[82768]++
													if f.pre != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:33
		_go_fuzz_dep_.CoverTab[82770]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:33
		return f.post != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:33
		// _ = "end of CoverTab[82770]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:33
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:33
		_go_fuzz_dep_.CoverTab[82771]++
														panic("have already forked")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:34
		// _ = "end of CoverTab[82771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:35
		_go_fuzz_dep_.CoverTab[82772]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:35
		// _ = "end of CoverTab[82772]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:35
	// _ = "end of CoverTab[82768]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:35
	_go_fuzz_dep_.CoverTab[82769]++
													f.pre = newForkableWriter()
													f.post = newForkableWriter()
													return f.pre, f.post
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:38
	// _ = "end of CoverTab[82769]"
}

func (f *forkableWriter) Len() (l int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:41
	_go_fuzz_dep_.CoverTab[82773]++
													l += f.Buffer.Len()
													if f.pre != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:43
		_go_fuzz_dep_.CoverTab[82776]++
														l += f.pre.Len()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:44
		// _ = "end of CoverTab[82776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:45
		_go_fuzz_dep_.CoverTab[82777]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:45
		// _ = "end of CoverTab[82777]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:45
	// _ = "end of CoverTab[82773]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:45
	_go_fuzz_dep_.CoverTab[82774]++
													if f.post != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:46
		_go_fuzz_dep_.CoverTab[82778]++
														l += f.post.Len()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:47
		// _ = "end of CoverTab[82778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:48
		_go_fuzz_dep_.CoverTab[82779]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:48
		// _ = "end of CoverTab[82779]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:48
	// _ = "end of CoverTab[82774]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:48
	_go_fuzz_dep_.CoverTab[82775]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:49
	// _ = "end of CoverTab[82775]"
}

func (f *forkableWriter) writeTo(out io.Writer) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:52
	_go_fuzz_dep_.CoverTab[82780]++
													n, err = out.Write(f.Bytes())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:54
		_go_fuzz_dep_.CoverTab[82784]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:55
		// _ = "end of CoverTab[82784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:56
		_go_fuzz_dep_.CoverTab[82785]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:56
		// _ = "end of CoverTab[82785]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:56
	// _ = "end of CoverTab[82780]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:56
	_go_fuzz_dep_.CoverTab[82781]++

													var nn int

													if f.pre != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:60
		_go_fuzz_dep_.CoverTab[82786]++
														nn, err = f.pre.writeTo(out)
														n += nn
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:63
			_go_fuzz_dep_.CoverTab[82787]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:64
			// _ = "end of CoverTab[82787]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:65
			_go_fuzz_dep_.CoverTab[82788]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:65
			// _ = "end of CoverTab[82788]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:65
		// _ = "end of CoverTab[82786]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:66
		_go_fuzz_dep_.CoverTab[82789]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:66
		// _ = "end of CoverTab[82789]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:66
	// _ = "end of CoverTab[82781]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:66
	_go_fuzz_dep_.CoverTab[82782]++

													if f.post != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:68
		_go_fuzz_dep_.CoverTab[82790]++
														nn, err = f.post.writeTo(out)
														n += nn
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:70
		// _ = "end of CoverTab[82790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:71
		_go_fuzz_dep_.CoverTab[82791]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:71
		// _ = "end of CoverTab[82791]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:71
	// _ = "end of CoverTab[82782]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:71
	_go_fuzz_dep_.CoverTab[82783]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:72
	// _ = "end of CoverTab[82783]"
}

func marshalBase128Int(out *forkableWriter, n int64) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:75
	_go_fuzz_dep_.CoverTab[82792]++
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:76
		_go_fuzz_dep_.CoverTab[82796]++
														err = out.WriteByte(0)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:78
		// _ = "end of CoverTab[82796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:79
		_go_fuzz_dep_.CoverTab[82797]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:79
		// _ = "end of CoverTab[82797]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:79
	// _ = "end of CoverTab[82792]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:79
	_go_fuzz_dep_.CoverTab[82793]++

													l := 0
													for i := n; i > 0; i >>= 7 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:82
		_go_fuzz_dep_.CoverTab[82798]++
														l++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:83
		// _ = "end of CoverTab[82798]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:84
	// _ = "end of CoverTab[82793]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:84
	_go_fuzz_dep_.CoverTab[82794]++

													for i := l - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:86
		_go_fuzz_dep_.CoverTab[82799]++
														o := byte(n >> uint(i*7))
														o &= 0x7f
														if i != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:89
			_go_fuzz_dep_.CoverTab[82801]++
															o |= 0x80
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:90
			// _ = "end of CoverTab[82801]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:91
			_go_fuzz_dep_.CoverTab[82802]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:91
			// _ = "end of CoverTab[82802]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:91
		// _ = "end of CoverTab[82799]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:91
		_go_fuzz_dep_.CoverTab[82800]++
														err = out.WriteByte(o)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:93
			_go_fuzz_dep_.CoverTab[82803]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:94
			// _ = "end of CoverTab[82803]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:95
			_go_fuzz_dep_.CoverTab[82804]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:95
			// _ = "end of CoverTab[82804]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:95
		// _ = "end of CoverTab[82800]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:96
	// _ = "end of CoverTab[82794]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:96
	_go_fuzz_dep_.CoverTab[82795]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:98
	// _ = "end of CoverTab[82795]"
}

func marshalInt64(out *forkableWriter, i int64) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:101
	_go_fuzz_dep_.CoverTab[82805]++
													n := int64Length(i)

													for ; n > 0; n-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:104
		_go_fuzz_dep_.CoverTab[82807]++
														err = out.WriteByte(byte(i >> uint((n-1)*8)))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:106
			_go_fuzz_dep_.CoverTab[82808]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:107
			// _ = "end of CoverTab[82808]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:108
			_go_fuzz_dep_.CoverTab[82809]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:108
			// _ = "end of CoverTab[82809]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:108
		// _ = "end of CoverTab[82807]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:109
	// _ = "end of CoverTab[82805]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:109
	_go_fuzz_dep_.CoverTab[82806]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:111
	// _ = "end of CoverTab[82806]"
}

func int64Length(i int64) (numBytes int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:114
	_go_fuzz_dep_.CoverTab[82810]++
													numBytes = 1

													for i > 127 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:117
		_go_fuzz_dep_.CoverTab[82813]++
														numBytes++
														i >>= 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:119
		// _ = "end of CoverTab[82813]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:120
	// _ = "end of CoverTab[82810]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:120
	_go_fuzz_dep_.CoverTab[82811]++

													for i < -128 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:122
		_go_fuzz_dep_.CoverTab[82814]++
														numBytes++
														i >>= 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:124
		// _ = "end of CoverTab[82814]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:125
	// _ = "end of CoverTab[82811]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:125
	_go_fuzz_dep_.CoverTab[82812]++

													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:127
	// _ = "end of CoverTab[82812]"
}

func marshalBigInt(out *forkableWriter, n *big.Int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:130
	_go_fuzz_dep_.CoverTab[82815]++
													if n.Sign() < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:131
		_go_fuzz_dep_.CoverTab[82817]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:136
		nMinus1 := new(big.Int).Neg(n)
		nMinus1.Sub(nMinus1, bigOne)
		bytes := nMinus1.Bytes()
		for i := range bytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:139
			_go_fuzz_dep_.CoverTab[82820]++
															bytes[i] ^= 0xff
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:140
			// _ = "end of CoverTab[82820]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:141
		// _ = "end of CoverTab[82817]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:141
		_go_fuzz_dep_.CoverTab[82818]++
														if len(bytes) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:142
			_go_fuzz_dep_.CoverTab[82821]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:142
			return bytes[0]&0x80 == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:142
			// _ = "end of CoverTab[82821]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:142
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:142
			_go_fuzz_dep_.CoverTab[82822]++
															err = out.WriteByte(0xff)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:144
				_go_fuzz_dep_.CoverTab[82823]++
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:145
				// _ = "end of CoverTab[82823]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:146
				_go_fuzz_dep_.CoverTab[82824]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:146
				// _ = "end of CoverTab[82824]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:146
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:146
			// _ = "end of CoverTab[82822]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:147
			_go_fuzz_dep_.CoverTab[82825]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:147
			// _ = "end of CoverTab[82825]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:147
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:147
		// _ = "end of CoverTab[82818]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:147
		_go_fuzz_dep_.CoverTab[82819]++
														_, err = out.Write(bytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:148
		// _ = "end of CoverTab[82819]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:149
		_go_fuzz_dep_.CoverTab[82826]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:149
		if n.Sign() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:149
			_go_fuzz_dep_.CoverTab[82827]++

															err = out.WriteByte(0x00)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:151
			// _ = "end of CoverTab[82827]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:152
			_go_fuzz_dep_.CoverTab[82828]++
															bytes := n.Bytes()
															if len(bytes) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:154
				_go_fuzz_dep_.CoverTab[82830]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:154
				return bytes[0]&0x80 != 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:154
				// _ = "end of CoverTab[82830]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:154
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:154
				_go_fuzz_dep_.CoverTab[82831]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:157
				err = out.WriteByte(0)
				if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:158
					_go_fuzz_dep_.CoverTab[82832]++
																	return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:159
					// _ = "end of CoverTab[82832]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:160
					_go_fuzz_dep_.CoverTab[82833]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:160
					// _ = "end of CoverTab[82833]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:160
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:160
				// _ = "end of CoverTab[82831]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:161
				_go_fuzz_dep_.CoverTab[82834]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:161
				// _ = "end of CoverTab[82834]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:161
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:161
			// _ = "end of CoverTab[82828]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:161
			_go_fuzz_dep_.CoverTab[82829]++
															_, err = out.Write(bytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:162
			// _ = "end of CoverTab[82829]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:163
		// _ = "end of CoverTab[82826]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:163
	// _ = "end of CoverTab[82815]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:163
	_go_fuzz_dep_.CoverTab[82816]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:164
	// _ = "end of CoverTab[82816]"
}

func marshalLength(out *forkableWriter, i int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:167
	_go_fuzz_dep_.CoverTab[82835]++
													n := lengthLength(i)

													for ; n > 0; n-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:170
		_go_fuzz_dep_.CoverTab[82837]++
														err = out.WriteByte(byte(i >> uint((n-1)*8)))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:172
			_go_fuzz_dep_.CoverTab[82838]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:173
			// _ = "end of CoverTab[82838]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:174
			_go_fuzz_dep_.CoverTab[82839]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:174
			// _ = "end of CoverTab[82839]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:174
		// _ = "end of CoverTab[82837]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:175
	// _ = "end of CoverTab[82835]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:175
	_go_fuzz_dep_.CoverTab[82836]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:177
	// _ = "end of CoverTab[82836]"
}

func lengthLength(i int) (numBytes int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:180
	_go_fuzz_dep_.CoverTab[82840]++
													numBytes = 1
													for i > 255 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:182
		_go_fuzz_dep_.CoverTab[82842]++
														numBytes++
														i >>= 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:184
		// _ = "end of CoverTab[82842]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:185
	// _ = "end of CoverTab[82840]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:185
	_go_fuzz_dep_.CoverTab[82841]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:186
	// _ = "end of CoverTab[82841]"
}

func marshalTagAndLength(out *forkableWriter, t tagAndLength) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:189
	_go_fuzz_dep_.CoverTab[82843]++
													b := uint8(t.class) << 6
													if t.isCompound {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:191
		_go_fuzz_dep_.CoverTab[82847]++
														b |= 0x20
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:192
		// _ = "end of CoverTab[82847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:193
		_go_fuzz_dep_.CoverTab[82848]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:193
		// _ = "end of CoverTab[82848]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:193
	// _ = "end of CoverTab[82843]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:193
	_go_fuzz_dep_.CoverTab[82844]++
													if t.tag >= 31 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:194
		_go_fuzz_dep_.CoverTab[82849]++
														b |= 0x1f
														err = out.WriteByte(b)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:197
			_go_fuzz_dep_.CoverTab[82851]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:198
			// _ = "end of CoverTab[82851]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:199
			_go_fuzz_dep_.CoverTab[82852]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:199
			// _ = "end of CoverTab[82852]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:199
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:199
		// _ = "end of CoverTab[82849]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:199
		_go_fuzz_dep_.CoverTab[82850]++
														err = marshalBase128Int(out, int64(t.tag))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:201
			_go_fuzz_dep_.CoverTab[82853]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:202
			// _ = "end of CoverTab[82853]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:203
			_go_fuzz_dep_.CoverTab[82854]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:203
			// _ = "end of CoverTab[82854]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:203
		// _ = "end of CoverTab[82850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:204
		_go_fuzz_dep_.CoverTab[82855]++
														b |= uint8(t.tag)
														err = out.WriteByte(b)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:207
			_go_fuzz_dep_.CoverTab[82856]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:208
			// _ = "end of CoverTab[82856]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:209
			_go_fuzz_dep_.CoverTab[82857]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:209
			// _ = "end of CoverTab[82857]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:209
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:209
		// _ = "end of CoverTab[82855]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:210
	// _ = "end of CoverTab[82844]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:210
	_go_fuzz_dep_.CoverTab[82845]++

													if t.length >= 128 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:212
		_go_fuzz_dep_.CoverTab[82858]++
														l := lengthLength(t.length)
														err = out.WriteByte(0x80 | byte(l))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:215
			_go_fuzz_dep_.CoverTab[82860]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:216
			// _ = "end of CoverTab[82860]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:217
			_go_fuzz_dep_.CoverTab[82861]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:217
			// _ = "end of CoverTab[82861]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:217
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:217
		// _ = "end of CoverTab[82858]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:217
		_go_fuzz_dep_.CoverTab[82859]++
														err = marshalLength(out, t.length)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:219
			_go_fuzz_dep_.CoverTab[82862]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:220
			// _ = "end of CoverTab[82862]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:221
			_go_fuzz_dep_.CoverTab[82863]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:221
			// _ = "end of CoverTab[82863]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:221
		// _ = "end of CoverTab[82859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:222
		_go_fuzz_dep_.CoverTab[82864]++
														err = out.WriteByte(byte(t.length))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:224
			_go_fuzz_dep_.CoverTab[82865]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:225
			// _ = "end of CoverTab[82865]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:226
			_go_fuzz_dep_.CoverTab[82866]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:226
			// _ = "end of CoverTab[82866]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:226
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:226
		// _ = "end of CoverTab[82864]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:227
	// _ = "end of CoverTab[82845]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:227
	_go_fuzz_dep_.CoverTab[82846]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:229
	// _ = "end of CoverTab[82846]"
}

func marshalBitString(out *forkableWriter, b BitString) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:232
	_go_fuzz_dep_.CoverTab[82867]++
													paddingBits := byte((8 - b.BitLength%8) % 8)
													err = out.WriteByte(paddingBits)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:235
		_go_fuzz_dep_.CoverTab[82869]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:236
		// _ = "end of CoverTab[82869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:237
		_go_fuzz_dep_.CoverTab[82870]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:237
		// _ = "end of CoverTab[82870]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:237
	// _ = "end of CoverTab[82867]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:237
	_go_fuzz_dep_.CoverTab[82868]++
													_, err = out.Write(b.Bytes)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:239
	// _ = "end of CoverTab[82868]"
}

func marshalObjectIdentifier(out *forkableWriter, oid []int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:242
	_go_fuzz_dep_.CoverTab[82871]++
													if len(oid) < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		_go_fuzz_dep_.CoverTab[82875]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		return oid[0] > 2
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		// _ = "end of CoverTab[82875]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		_go_fuzz_dep_.CoverTab[82876]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		return (oid[0] < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
			_go_fuzz_dep_.CoverTab[82877]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
			return oid[1] >= 40
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
			// _ = "end of CoverTab[82877]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		// _ = "end of CoverTab[82876]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:243
		_go_fuzz_dep_.CoverTab[82878]++
														return StructuralError{"invalid object identifier"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:244
		// _ = "end of CoverTab[82878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:245
		_go_fuzz_dep_.CoverTab[82879]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:245
		// _ = "end of CoverTab[82879]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:245
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:245
	// _ = "end of CoverTab[82871]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:245
	_go_fuzz_dep_.CoverTab[82872]++

													err = marshalBase128Int(out, int64(oid[0]*40+oid[1]))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:248
		_go_fuzz_dep_.CoverTab[82880]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:249
		// _ = "end of CoverTab[82880]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:250
		_go_fuzz_dep_.CoverTab[82881]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:250
		// _ = "end of CoverTab[82881]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:250
	// _ = "end of CoverTab[82872]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:250
	_go_fuzz_dep_.CoverTab[82873]++
													for i := 2; i < len(oid); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:251
		_go_fuzz_dep_.CoverTab[82882]++
														err = marshalBase128Int(out, int64(oid[i]))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:253
			_go_fuzz_dep_.CoverTab[82883]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:254
			// _ = "end of CoverTab[82883]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:255
			_go_fuzz_dep_.CoverTab[82884]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:255
			// _ = "end of CoverTab[82884]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:255
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:255
		// _ = "end of CoverTab[82882]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:256
	// _ = "end of CoverTab[82873]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:256
	_go_fuzz_dep_.CoverTab[82874]++

													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:258
	// _ = "end of CoverTab[82874]"
}

func marshalPrintableString(out *forkableWriter, s string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:261
	_go_fuzz_dep_.CoverTab[82885]++
													b := []byte(s)
													for _, c := range b {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:263
		_go_fuzz_dep_.CoverTab[82887]++
														if !isPrintable(c) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:264
			_go_fuzz_dep_.CoverTab[82888]++
															return StructuralError{"PrintableString contains invalid character"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:265
			// _ = "end of CoverTab[82888]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:266
			_go_fuzz_dep_.CoverTab[82889]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:266
			// _ = "end of CoverTab[82889]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:266
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:266
		// _ = "end of CoverTab[82887]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:267
	// _ = "end of CoverTab[82885]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:267
	_go_fuzz_dep_.CoverTab[82886]++

													_, err = out.Write(b)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:270
	// _ = "end of CoverTab[82886]"
}

func marshalIA5String(out *forkableWriter, s string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:273
	_go_fuzz_dep_.CoverTab[82890]++
													b := []byte(s)
													for _, c := range b {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:275
		_go_fuzz_dep_.CoverTab[82892]++
														if c > 127 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:276
			_go_fuzz_dep_.CoverTab[82893]++
															return StructuralError{"IA5String contains invalid character"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:277
			// _ = "end of CoverTab[82893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:278
			_go_fuzz_dep_.CoverTab[82894]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:278
			// _ = "end of CoverTab[82894]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:278
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:278
		// _ = "end of CoverTab[82892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:279
	// _ = "end of CoverTab[82890]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:279
	_go_fuzz_dep_.CoverTab[82891]++

													_, err = out.Write(b)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:282
	// _ = "end of CoverTab[82891]"
}

func marshalUTF8String(out *forkableWriter, s string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:285
	_go_fuzz_dep_.CoverTab[82895]++
													_, err = out.Write([]byte(s))
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:287
	// _ = "end of CoverTab[82895]"
}

func marshalTwoDigits(out *forkableWriter, v int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:290
	_go_fuzz_dep_.CoverTab[82896]++
													err = out.WriteByte(byte('0' + (v/10)%10))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:292
		_go_fuzz_dep_.CoverTab[82898]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:293
		// _ = "end of CoverTab[82898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:294
		_go_fuzz_dep_.CoverTab[82899]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:294
		// _ = "end of CoverTab[82899]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:294
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:294
	// _ = "end of CoverTab[82896]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:294
	_go_fuzz_dep_.CoverTab[82897]++
													return out.WriteByte(byte('0' + v%10))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:295
	// _ = "end of CoverTab[82897]"
}

func marshalFourDigits(out *forkableWriter, v int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:298
	_go_fuzz_dep_.CoverTab[82900]++
													var bytes [4]byte
													for i := range bytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:300
		_go_fuzz_dep_.CoverTab[82902]++
														bytes[3-i] = '0' + byte(v%10)
														v /= 10
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:302
		// _ = "end of CoverTab[82902]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:303
	// _ = "end of CoverTab[82900]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:303
	_go_fuzz_dep_.CoverTab[82901]++
													_, err = out.Write(bytes[:])
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:305
	// _ = "end of CoverTab[82901]"
}

func outsideUTCRange(t time.Time) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:308
	_go_fuzz_dep_.CoverTab[82903]++
													year := t.Year()
													return year < 1950 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:310
		_go_fuzz_dep_.CoverTab[82904]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:310
		return year >= 2050
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:310
		// _ = "end of CoverTab[82904]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:310
	}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:310
	// _ = "end of CoverTab[82903]"
}

func marshalUTCTime(out *forkableWriter, t time.Time) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:313
	_go_fuzz_dep_.CoverTab[82905]++
													year := t.Year()

													switch {
	case 1950 <= year && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:317
		_go_fuzz_dep_.CoverTab[82911]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:317
		return year < 2000
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:317
		// _ = "end of CoverTab[82911]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:317
	}():
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:317
		_go_fuzz_dep_.CoverTab[82908]++
														err = marshalTwoDigits(out, year-1900)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:318
		// _ = "end of CoverTab[82908]"
	case 2000 <= year && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:319
		_go_fuzz_dep_.CoverTab[82912]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:319
		return year < 2050
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:319
		// _ = "end of CoverTab[82912]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:319
	}():
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:319
		_go_fuzz_dep_.CoverTab[82909]++
														err = marshalTwoDigits(out, year-2000)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:320
		// _ = "end of CoverTab[82909]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:321
		_go_fuzz_dep_.CoverTab[82910]++
														return StructuralError{"cannot represent time as UTCTime"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:322
		// _ = "end of CoverTab[82910]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:323
	// _ = "end of CoverTab[82905]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:323
	_go_fuzz_dep_.CoverTab[82906]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:324
		_go_fuzz_dep_.CoverTab[82913]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:325
		// _ = "end of CoverTab[82913]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:326
		_go_fuzz_dep_.CoverTab[82914]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:326
		// _ = "end of CoverTab[82914]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:326
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:326
	// _ = "end of CoverTab[82906]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:326
	_go_fuzz_dep_.CoverTab[82907]++

													return marshalTimeCommon(out, t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:328
	// _ = "end of CoverTab[82907]"
}

func marshalGeneralizedTime(out *forkableWriter, t time.Time) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:331
	_go_fuzz_dep_.CoverTab[82915]++
													year := t.Year()
													if year < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:333
		_go_fuzz_dep_.CoverTab[82918]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:333
		return year > 9999
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:333
		// _ = "end of CoverTab[82918]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:333
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:333
		_go_fuzz_dep_.CoverTab[82919]++
														return StructuralError{"cannot represent time as GeneralizedTime"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:334
		// _ = "end of CoverTab[82919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:335
		_go_fuzz_dep_.CoverTab[82920]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:335
		// _ = "end of CoverTab[82920]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:335
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:335
	// _ = "end of CoverTab[82915]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:335
	_go_fuzz_dep_.CoverTab[82916]++
													if err = marshalFourDigits(out, year); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:336
		_go_fuzz_dep_.CoverTab[82921]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:337
		// _ = "end of CoverTab[82921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:338
		_go_fuzz_dep_.CoverTab[82922]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:338
		// _ = "end of CoverTab[82922]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:338
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:338
	// _ = "end of CoverTab[82916]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:338
	_go_fuzz_dep_.CoverTab[82917]++

													return marshalTimeCommon(out, t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:340
	// _ = "end of CoverTab[82917]"
}

func marshalTimeCommon(out *forkableWriter, t time.Time) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:343
	_go_fuzz_dep_.CoverTab[82923]++
													_, month, day := t.Date()

													err = marshalTwoDigits(out, int(month))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:347
		_go_fuzz_dep_.CoverTab[82933]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:348
		// _ = "end of CoverTab[82933]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:349
		_go_fuzz_dep_.CoverTab[82934]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:349
		// _ = "end of CoverTab[82934]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:349
	// _ = "end of CoverTab[82923]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:349
	_go_fuzz_dep_.CoverTab[82924]++

													err = marshalTwoDigits(out, day)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:352
		_go_fuzz_dep_.CoverTab[82935]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:353
		// _ = "end of CoverTab[82935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:354
		_go_fuzz_dep_.CoverTab[82936]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:354
		// _ = "end of CoverTab[82936]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:354
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:354
	// _ = "end of CoverTab[82924]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:354
	_go_fuzz_dep_.CoverTab[82925]++

													hour, min, sec := t.Clock()

													err = marshalTwoDigits(out, hour)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:359
		_go_fuzz_dep_.CoverTab[82937]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:360
		// _ = "end of CoverTab[82937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:361
		_go_fuzz_dep_.CoverTab[82938]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:361
		// _ = "end of CoverTab[82938]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:361
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:361
	// _ = "end of CoverTab[82925]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:361
	_go_fuzz_dep_.CoverTab[82926]++

													err = marshalTwoDigits(out, min)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:364
		_go_fuzz_dep_.CoverTab[82939]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:365
		// _ = "end of CoverTab[82939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:366
		_go_fuzz_dep_.CoverTab[82940]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:366
		// _ = "end of CoverTab[82940]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:366
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:366
	// _ = "end of CoverTab[82926]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:366
	_go_fuzz_dep_.CoverTab[82927]++

													err = marshalTwoDigits(out, sec)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:369
		_go_fuzz_dep_.CoverTab[82941]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:370
		// _ = "end of CoverTab[82941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:371
		_go_fuzz_dep_.CoverTab[82942]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:371
		// _ = "end of CoverTab[82942]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:371
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:371
	// _ = "end of CoverTab[82927]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:371
	_go_fuzz_dep_.CoverTab[82928]++

													_, offset := t.Zone()

													switch {
	case offset/60 == 0:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:376
		_go_fuzz_dep_.CoverTab[82943]++
														err = out.WriteByte('Z')
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:378
		// _ = "end of CoverTab[82943]"
	case offset > 0:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:379
		_go_fuzz_dep_.CoverTab[82944]++
														err = out.WriteByte('+')
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:380
		// _ = "end of CoverTab[82944]"
	case offset < 0:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:381
		_go_fuzz_dep_.CoverTab[82945]++
														err = out.WriteByte('-')
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:382
		// _ = "end of CoverTab[82945]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:382
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:382
		_go_fuzz_dep_.CoverTab[82946]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:382
		// _ = "end of CoverTab[82946]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:383
	// _ = "end of CoverTab[82928]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:383
	_go_fuzz_dep_.CoverTab[82929]++

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:385
		_go_fuzz_dep_.CoverTab[82947]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:386
		// _ = "end of CoverTab[82947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:387
		_go_fuzz_dep_.CoverTab[82948]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:387
		// _ = "end of CoverTab[82948]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:387
	// _ = "end of CoverTab[82929]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:387
	_go_fuzz_dep_.CoverTab[82930]++

													offsetMinutes := offset / 60
													if offsetMinutes < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:390
		_go_fuzz_dep_.CoverTab[82949]++
														offsetMinutes = -offsetMinutes
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:391
		// _ = "end of CoverTab[82949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:392
		_go_fuzz_dep_.CoverTab[82950]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:392
		// _ = "end of CoverTab[82950]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:392
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:392
	// _ = "end of CoverTab[82930]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:392
	_go_fuzz_dep_.CoverTab[82931]++

													err = marshalTwoDigits(out, offsetMinutes/60)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:395
		_go_fuzz_dep_.CoverTab[82951]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:396
		// _ = "end of CoverTab[82951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:397
		_go_fuzz_dep_.CoverTab[82952]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:397
		// _ = "end of CoverTab[82952]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:397
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:397
	// _ = "end of CoverTab[82931]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:397
	_go_fuzz_dep_.CoverTab[82932]++

													err = marshalTwoDigits(out, offsetMinutes%60)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:400
	// _ = "end of CoverTab[82932]"
}

func stripTagAndLength(in []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:403
	_go_fuzz_dep_.CoverTab[82953]++
													_, offset, err := parseTagAndLength(in, 0)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:405
		_go_fuzz_dep_.CoverTab[82955]++
														return in
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:406
		// _ = "end of CoverTab[82955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:407
		_go_fuzz_dep_.CoverTab[82956]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:407
		// _ = "end of CoverTab[82956]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:407
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:407
	// _ = "end of CoverTab[82953]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:407
	_go_fuzz_dep_.CoverTab[82954]++
													return in[offset:]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:408
	// _ = "end of CoverTab[82954]"
}

func marshalBody(out *forkableWriter, value reflect.Value, params fieldParameters) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:411
	_go_fuzz_dep_.CoverTab[82957]++
													switch value.Type() {
	case flagType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:413
		_go_fuzz_dep_.CoverTab[82960]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:414
		// _ = "end of CoverTab[82960]"
	case timeType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:415
		_go_fuzz_dep_.CoverTab[82961]++
														t := value.Interface().(time.Time)
														if params.timeType == TagGeneralizedTime || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:417
			_go_fuzz_dep_.CoverTab[82966]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:417
			return outsideUTCRange(t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:417
			// _ = "end of CoverTab[82966]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:417
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:417
			_go_fuzz_dep_.CoverTab[82967]++
															return marshalGeneralizedTime(out, t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:418
			// _ = "end of CoverTab[82967]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:419
			_go_fuzz_dep_.CoverTab[82968]++
															return marshalUTCTime(out, t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:420
			// _ = "end of CoverTab[82968]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:421
		// _ = "end of CoverTab[82961]"
	case bitStringType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:422
		_go_fuzz_dep_.CoverTab[82962]++
														return marshalBitString(out, value.Interface().(BitString))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:423
		// _ = "end of CoverTab[82962]"
	case objectIdentifierType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:424
		_go_fuzz_dep_.CoverTab[82963]++
														return marshalObjectIdentifier(out, value.Interface().(ObjectIdentifier))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:425
		// _ = "end of CoverTab[82963]"
	case bigIntType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:426
		_go_fuzz_dep_.CoverTab[82964]++
														return marshalBigInt(out, value.Interface().(*big.Int))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:427
		// _ = "end of CoverTab[82964]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:427
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:427
		_go_fuzz_dep_.CoverTab[82965]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:427
		// _ = "end of CoverTab[82965]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:428
	// _ = "end of CoverTab[82957]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:428
	_go_fuzz_dep_.CoverTab[82958]++

													switch v := value; v.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:431
		_go_fuzz_dep_.CoverTab[82969]++
														if v.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:432
			_go_fuzz_dep_.CoverTab[82979]++
															return out.WriteByte(255)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:433
			// _ = "end of CoverTab[82979]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:434
			_go_fuzz_dep_.CoverTab[82980]++
															return out.WriteByte(0)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:435
			// _ = "end of CoverTab[82980]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:436
		// _ = "end of CoverTab[82969]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:437
		_go_fuzz_dep_.CoverTab[82970]++
														return marshalInt64(out, v.Int())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:438
		// _ = "end of CoverTab[82970]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:439
		_go_fuzz_dep_.CoverTab[82971]++
														t := v.Type()

														startingField := 0

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
		if t.NumField() > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
			_go_fuzz_dep_.CoverTab[82981]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
			return t.Field(0).Type == rawContentsType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
			// _ = "end of CoverTab[82981]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:446
			_go_fuzz_dep_.CoverTab[82982]++
															s := v.Field(0)
															if s.Len() > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:448
				_go_fuzz_dep_.CoverTab[82983]++
																bytes := make([]byte, s.Len())
																for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:450
					_go_fuzz_dep_.CoverTab[82985]++
																	bytes[i] = uint8(s.Index(i).Uint())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:451
					// _ = "end of CoverTab[82985]"
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:452
				// _ = "end of CoverTab[82983]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:452
				_go_fuzz_dep_.CoverTab[82984]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:457
				_, err = out.Write(stripTagAndLength(bytes))
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:458
				// _ = "end of CoverTab[82984]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:459
				_go_fuzz_dep_.CoverTab[82986]++
																startingField = 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:460
				// _ = "end of CoverTab[82986]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:461
			// _ = "end of CoverTab[82982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:462
			_go_fuzz_dep_.CoverTab[82987]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:462
			// _ = "end of CoverTab[82987]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:462
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:462
		// _ = "end of CoverTab[82971]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:462
		_go_fuzz_dep_.CoverTab[82972]++

														for i := startingField; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:464
			_go_fuzz_dep_.CoverTab[82988]++
															var pre *forkableWriter
															pre, out = out.fork()
															err = marshalField(pre, v.Field(i), parseFieldParameters(t.Field(i).Tag.Get("asn1")))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:468
				_go_fuzz_dep_.CoverTab[82989]++
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:469
				// _ = "end of CoverTab[82989]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:470
				_go_fuzz_dep_.CoverTab[82990]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:470
				// _ = "end of CoverTab[82990]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:470
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:470
			// _ = "end of CoverTab[82988]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:471
		// _ = "end of CoverTab[82972]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:471
		_go_fuzz_dep_.CoverTab[82973]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:472
		// _ = "end of CoverTab[82973]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:473
		_go_fuzz_dep_.CoverTab[82974]++
														sliceType := v.Type()
														if sliceType.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:475
			_go_fuzz_dep_.CoverTab[82991]++
															bytes := make([]byte, v.Len())
															for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:477
				_go_fuzz_dep_.CoverTab[82993]++
																bytes[i] = uint8(v.Index(i).Uint())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:478
				// _ = "end of CoverTab[82993]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:479
			// _ = "end of CoverTab[82991]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:479
			_go_fuzz_dep_.CoverTab[82992]++
															_, err = out.Write(bytes)
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:481
			// _ = "end of CoverTab[82992]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:482
			_go_fuzz_dep_.CoverTab[82994]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:482
			// _ = "end of CoverTab[82994]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:482
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:482
		// _ = "end of CoverTab[82974]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:482
		_go_fuzz_dep_.CoverTab[82975]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:486
		params.explicit = false
		params.tag = nil
		for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:488
			_go_fuzz_dep_.CoverTab[82995]++
															var pre *forkableWriter
															pre, out = out.fork()
															err = marshalField(pre, v.Index(i), params)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:492
				_go_fuzz_dep_.CoverTab[82996]++
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:493
				// _ = "end of CoverTab[82996]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:494
				_go_fuzz_dep_.CoverTab[82997]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:494
				// _ = "end of CoverTab[82997]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:494
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:494
			// _ = "end of CoverTab[82995]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:495
		// _ = "end of CoverTab[82975]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:495
		_go_fuzz_dep_.CoverTab[82976]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:496
		// _ = "end of CoverTab[82976]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:497
		_go_fuzz_dep_.CoverTab[82977]++
														switch params.stringType {
		case TagIA5String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:499
			_go_fuzz_dep_.CoverTab[82998]++
															return marshalIA5String(out, v.String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:500
			// _ = "end of CoverTab[82998]"
		case TagPrintableString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:501
			_go_fuzz_dep_.CoverTab[82999]++
															return marshalPrintableString(out, v.String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:502
			// _ = "end of CoverTab[82999]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:503
			_go_fuzz_dep_.CoverTab[83000]++
															return marshalUTF8String(out, v.String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:504
			// _ = "end of CoverTab[83000]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:505
		// _ = "end of CoverTab[82977]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:505
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:505
		_go_fuzz_dep_.CoverTab[82978]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:505
		// _ = "end of CoverTab[82978]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:506
	// _ = "end of CoverTab[82958]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:506
	_go_fuzz_dep_.CoverTab[82959]++

													return StructuralError{"unknown Go type"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:508
	// _ = "end of CoverTab[82959]"
}

func marshalField(out *forkableWriter, v reflect.Value, params fieldParameters) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:511
	_go_fuzz_dep_.CoverTab[83001]++
													if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:512
		_go_fuzz_dep_.CoverTab[83018]++
														return fmt.Errorf("asn1: cannot marshal nil value")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:513
		// _ = "end of CoverTab[83018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:514
		_go_fuzz_dep_.CoverTab[83019]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:514
		// _ = "end of CoverTab[83019]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:514
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:514
	// _ = "end of CoverTab[83001]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:514
	_go_fuzz_dep_.CoverTab[83002]++

													if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:516
		_go_fuzz_dep_.CoverTab[83020]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:516
		return v.Type().NumMethod() == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:516
		// _ = "end of CoverTab[83020]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:516
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:516
		_go_fuzz_dep_.CoverTab[83021]++
														return marshalField(out, v.Elem(), params)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:517
		// _ = "end of CoverTab[83021]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:518
		_go_fuzz_dep_.CoverTab[83022]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:518
		// _ = "end of CoverTab[83022]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:518
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:518
	// _ = "end of CoverTab[83002]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:518
	_go_fuzz_dep_.CoverTab[83003]++

													if v.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		_go_fuzz_dep_.CoverTab[83023]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		return v.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		// _ = "end of CoverTab[83023]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		_go_fuzz_dep_.CoverTab[83024]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		return params.omitEmpty
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		// _ = "end of CoverTab[83024]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:520
		_go_fuzz_dep_.CoverTab[83025]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:521
		// _ = "end of CoverTab[83025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:522
		_go_fuzz_dep_.CoverTab[83026]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:522
		// _ = "end of CoverTab[83026]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:522
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:522
	// _ = "end of CoverTab[83003]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:522
	_go_fuzz_dep_.CoverTab[83004]++

													if params.optional && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		_go_fuzz_dep_.CoverTab[83027]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		return params.defaultValue != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		// _ = "end of CoverTab[83027]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		_go_fuzz_dep_.CoverTab[83028]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		return canHaveDefaultValue(v.Kind())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		// _ = "end of CoverTab[83028]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:524
		_go_fuzz_dep_.CoverTab[83029]++
														defaultValue := reflect.New(v.Type()).Elem()
														defaultValue.SetInt(*params.defaultValue)

														if reflect.DeepEqual(v.Interface(), defaultValue.Interface()) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:528
			_go_fuzz_dep_.CoverTab[83030]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:529
			// _ = "end of CoverTab[83030]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:530
			_go_fuzz_dep_.CoverTab[83031]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:530
			// _ = "end of CoverTab[83031]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:530
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:530
		// _ = "end of CoverTab[83029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:531
		_go_fuzz_dep_.CoverTab[83032]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:531
		// _ = "end of CoverTab[83032]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:531
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:531
	// _ = "end of CoverTab[83004]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:531
	_go_fuzz_dep_.CoverTab[83005]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
	if params.optional && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
		_go_fuzz_dep_.CoverTab[83033]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
		return params.defaultValue == nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
		// _ = "end of CoverTab[83033]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:536
		_go_fuzz_dep_.CoverTab[83034]++
														if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:537
			_go_fuzz_dep_.CoverTab[83035]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:538
			// _ = "end of CoverTab[83035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:539
			_go_fuzz_dep_.CoverTab[83036]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:539
			// _ = "end of CoverTab[83036]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:539
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:539
		// _ = "end of CoverTab[83034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:540
		_go_fuzz_dep_.CoverTab[83037]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:540
		// _ = "end of CoverTab[83037]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:540
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:540
	// _ = "end of CoverTab[83005]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:540
	_go_fuzz_dep_.CoverTab[83006]++

													if v.Type() == rawValueType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:542
		_go_fuzz_dep_.CoverTab[83038]++
														rv := v.Interface().(RawValue)
														if len(rv.FullBytes) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:544
			_go_fuzz_dep_.CoverTab[83040]++
															_, err = out.Write(rv.FullBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:545
			// _ = "end of CoverTab[83040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:546
			_go_fuzz_dep_.CoverTab[83041]++
															err = marshalTagAndLength(out, tagAndLength{rv.Class, rv.Tag, len(rv.Bytes), rv.IsCompound})
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:548
				_go_fuzz_dep_.CoverTab[83043]++
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:549
				// _ = "end of CoverTab[83043]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:550
				_go_fuzz_dep_.CoverTab[83044]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:550
				// _ = "end of CoverTab[83044]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:550
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:550
			// _ = "end of CoverTab[83041]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:550
			_go_fuzz_dep_.CoverTab[83042]++
															_, err = out.Write(rv.Bytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:551
			// _ = "end of CoverTab[83042]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:552
		// _ = "end of CoverTab[83038]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:552
		_go_fuzz_dep_.CoverTab[83039]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:553
		// _ = "end of CoverTab[83039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:554
		_go_fuzz_dep_.CoverTab[83045]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:554
		// _ = "end of CoverTab[83045]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:554
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:554
	// _ = "end of CoverTab[83006]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:554
	_go_fuzz_dep_.CoverTab[83007]++

													tag, isCompound, ok := getUniversalType(v.Type())
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:557
		_go_fuzz_dep_.CoverTab[83046]++
														err = StructuralError{fmt.Sprintf("unknown Go type: %v", v.Type())}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:559
		// _ = "end of CoverTab[83046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:560
		_go_fuzz_dep_.CoverTab[83047]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:560
		// _ = "end of CoverTab[83047]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:560
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:560
	// _ = "end of CoverTab[83007]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:560
	_go_fuzz_dep_.CoverTab[83008]++
													class := ClassUniversal

													if params.timeType != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:563
		_go_fuzz_dep_.CoverTab[83048]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:563
		return tag != TagUTCTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:563
		// _ = "end of CoverTab[83048]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:563
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:563
		_go_fuzz_dep_.CoverTab[83049]++
														return StructuralError{"explicit time type given to non-time member"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:564
		// _ = "end of CoverTab[83049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:565
		_go_fuzz_dep_.CoverTab[83050]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:565
		// _ = "end of CoverTab[83050]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:565
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:565
	// _ = "end of CoverTab[83008]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:565
	_go_fuzz_dep_.CoverTab[83009]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
	if params.stringType != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
		_go_fuzz_dep_.CoverTab[83051]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
		return !(tag == TagPrintableString || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
			_go_fuzz_dep_.CoverTab[83052]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
			return (v.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				_go_fuzz_dep_.CoverTab[83053]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				return tag == 16
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				// _ = "end of CoverTab[83053]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				_go_fuzz_dep_.CoverTab[83054]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				return v.Type().Elem().Kind() == reflect.String
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
				// _ = "end of CoverTab[83054]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
			}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
			// _ = "end of CoverTab[83052]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
		}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
		// _ = "end of CoverTab[83051]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:568
		_go_fuzz_dep_.CoverTab[83055]++
														return StructuralError{"explicit string type given to non-string member"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:569
		// _ = "end of CoverTab[83055]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:570
		_go_fuzz_dep_.CoverTab[83056]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:570
		// _ = "end of CoverTab[83056]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:570
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:570
	// _ = "end of CoverTab[83009]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:570
	_go_fuzz_dep_.CoverTab[83010]++

													switch tag {
	case TagPrintableString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:573
		_go_fuzz_dep_.CoverTab[83057]++
														if params.stringType == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:574
			_go_fuzz_dep_.CoverTab[83060]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:578
			for _, r := range v.String() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:578
				_go_fuzz_dep_.CoverTab[83061]++
																if r >= utf8.RuneSelf || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:579
					_go_fuzz_dep_.CoverTab[83062]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:579
					return !isPrintable(byte(r))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:579
					// _ = "end of CoverTab[83062]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:579
				}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:579
					_go_fuzz_dep_.CoverTab[83063]++
																	if !utf8.ValidString(v.String()) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:580
						_go_fuzz_dep_.CoverTab[83065]++
																		return errors.New("asn1: string not valid UTF-8")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:581
						// _ = "end of CoverTab[83065]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:582
						_go_fuzz_dep_.CoverTab[83066]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:582
						// _ = "end of CoverTab[83066]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:582
					}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:582
					// _ = "end of CoverTab[83063]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:582
					_go_fuzz_dep_.CoverTab[83064]++
																	tag = TagUTF8String
																	break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:584
					// _ = "end of CoverTab[83064]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:585
					_go_fuzz_dep_.CoverTab[83067]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:585
					// _ = "end of CoverTab[83067]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:585
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:585
				// _ = "end of CoverTab[83061]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:586
			// _ = "end of CoverTab[83060]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:587
			_go_fuzz_dep_.CoverTab[83068]++
															tag = params.stringType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:588
			// _ = "end of CoverTab[83068]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:589
		// _ = "end of CoverTab[83057]"
	case TagUTCTime:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:590
		_go_fuzz_dep_.CoverTab[83058]++
														if params.timeType == TagGeneralizedTime || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:591
			_go_fuzz_dep_.CoverTab[83069]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:591
			return outsideUTCRange(v.Interface().(time.Time))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:591
			// _ = "end of CoverTab[83069]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:591
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:591
			_go_fuzz_dep_.CoverTab[83070]++
															tag = TagGeneralizedTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:592
			// _ = "end of CoverTab[83070]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
			_go_fuzz_dep_.CoverTab[83071]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
			// _ = "end of CoverTab[83071]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
		// _ = "end of CoverTab[83058]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
		_go_fuzz_dep_.CoverTab[83059]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:593
		// _ = "end of CoverTab[83059]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:594
	// _ = "end of CoverTab[83010]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:594
	_go_fuzz_dep_.CoverTab[83011]++

													if params.set {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:596
		_go_fuzz_dep_.CoverTab[83072]++
														if tag != TagSequence {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:597
			_go_fuzz_dep_.CoverTab[83074]++
															return StructuralError{"non sequence tagged as set"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:598
			// _ = "end of CoverTab[83074]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:599
			_go_fuzz_dep_.CoverTab[83075]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:599
			// _ = "end of CoverTab[83075]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:599
		// _ = "end of CoverTab[83072]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:599
		_go_fuzz_dep_.CoverTab[83073]++
														tag = TagSet
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:600
		// _ = "end of CoverTab[83073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:601
		_go_fuzz_dep_.CoverTab[83076]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:601
		// _ = "end of CoverTab[83076]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:601
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:601
	// _ = "end of CoverTab[83011]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:601
	_go_fuzz_dep_.CoverTab[83012]++

													tags, body := out.fork()

													err = marshalBody(body, v, params)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:606
		_go_fuzz_dep_.CoverTab[83077]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:607
		// _ = "end of CoverTab[83077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:608
		_go_fuzz_dep_.CoverTab[83078]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:608
		// _ = "end of CoverTab[83078]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:608
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:608
	// _ = "end of CoverTab[83012]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:608
	_go_fuzz_dep_.CoverTab[83013]++

													bodyLen := body.Len()

													var explicitTag *forkableWriter
													if params.explicit {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:613
		_go_fuzz_dep_.CoverTab[83079]++
														explicitTag, tags = tags.fork()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:614
		// _ = "end of CoverTab[83079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:615
		_go_fuzz_dep_.CoverTab[83080]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:615
		// _ = "end of CoverTab[83080]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:615
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:615
	// _ = "end of CoverTab[83013]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:615
	_go_fuzz_dep_.CoverTab[83014]++

													if !params.explicit && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:617
		_go_fuzz_dep_.CoverTab[83081]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:617
		return params.tag != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:617
		// _ = "end of CoverTab[83081]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:617
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:617
		_go_fuzz_dep_.CoverTab[83082]++

														tag = *params.tag
														class = ClassContextSpecific
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:620
		// _ = "end of CoverTab[83082]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:621
		_go_fuzz_dep_.CoverTab[83083]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:621
		// _ = "end of CoverTab[83083]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:621
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:621
	// _ = "end of CoverTab[83014]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:621
	_go_fuzz_dep_.CoverTab[83015]++

													err = marshalTagAndLength(tags, tagAndLength{class, tag, bodyLen, isCompound})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:624
		_go_fuzz_dep_.CoverTab[83084]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:625
		// _ = "end of CoverTab[83084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:626
		_go_fuzz_dep_.CoverTab[83085]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:626
		// _ = "end of CoverTab[83085]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:626
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:626
	// _ = "end of CoverTab[83015]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:626
	_go_fuzz_dep_.CoverTab[83016]++

													if params.explicit {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:628
		_go_fuzz_dep_.CoverTab[83086]++
														err = marshalTagAndLength(explicitTag, tagAndLength{
			class:		ClassContextSpecific,
			tag:		*params.tag,
			length:		bodyLen + tags.Len(),
			isCompound:	true,
		})
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:634
		// _ = "end of CoverTab[83086]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:635
		_go_fuzz_dep_.CoverTab[83087]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:635
		// _ = "end of CoverTab[83087]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:635
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:635
	// _ = "end of CoverTab[83016]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:635
	_go_fuzz_dep_.CoverTab[83017]++

													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:637
	// _ = "end of CoverTab[83017]"
}

// Marshal returns the ASN.1 encoding of val.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
// In addition to the struct tags recognised by Unmarshal, the following can be
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
// used:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//	ia5:		causes strings to be marshaled as ASN.1, IA5 strings
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//	omitempty:	causes empty slices to be skipped
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//	printable:	causes strings to be marshaled as ASN.1, PrintableString strings.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:640
//	utf8:		causes strings to be marshaled as ASN.1, UTF8 strings
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:649
func Marshal(val interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:649
	_go_fuzz_dep_.CoverTab[83088]++
													var out bytes.Buffer
													v := reflect.ValueOf(val)
													f := newForkableWriter()
													err := marshalField(f, v, fieldParameters{})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:654
		_go_fuzz_dep_.CoverTab[83090]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:655
		// _ = "end of CoverTab[83090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:656
		_go_fuzz_dep_.CoverTab[83091]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:656
		// _ = "end of CoverTab[83091]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:656
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:656
	// _ = "end of CoverTab[83088]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:656
	_go_fuzz_dep_.CoverTab[83089]++
													_, err = f.writeTo(&out)
													return out.Bytes(), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:658
	// _ = "end of CoverTab[83089]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:659
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/marshal.go:659
var _ = _go_fuzz_dep_.CoverTab
