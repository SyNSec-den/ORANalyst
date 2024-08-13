// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/cipher/io.go:5
package cipher

//line /usr/local/go/src/crypto/cipher/io.go:5
import (
//line /usr/local/go/src/crypto/cipher/io.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/io.go:5
)
//line /usr/local/go/src/crypto/cipher/io.go:5
import (
//line /usr/local/go/src/crypto/cipher/io.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/io.go:5
)

import "io"

//line /usr/local/go/src/crypto/cipher/io.go:12
// StreamReader wraps a Stream into an io.Reader. It calls XORKeyStream
//line /usr/local/go/src/crypto/cipher/io.go:12
// to process each slice of data which passes through.
//line /usr/local/go/src/crypto/cipher/io.go:14
type StreamReader struct {
	S	Stream
	R	io.Reader
}

func (r StreamReader) Read(dst []byte) (n int, err error) {
//line /usr/local/go/src/crypto/cipher/io.go:19
	_go_fuzz_dep_.CoverTab[1672]++
							n, err = r.R.Read(dst)
							r.S.XORKeyStream(dst[:n], dst[:n])
							return
//line /usr/local/go/src/crypto/cipher/io.go:22
	// _ = "end of CoverTab[1672]"
}

// StreamWriter wraps a Stream into an io.Writer. It calls XORKeyStream
//line /usr/local/go/src/crypto/cipher/io.go:25
// to process each slice of data which passes through. If any Write call
//line /usr/local/go/src/crypto/cipher/io.go:25
// returns short then the StreamWriter is out of sync and must be discarded.
//line /usr/local/go/src/crypto/cipher/io.go:25
// A StreamWriter has no internal buffering; Close does not need
//line /usr/local/go/src/crypto/cipher/io.go:25
// to be called to flush write data.
//line /usr/local/go/src/crypto/cipher/io.go:30
type StreamWriter struct {
	S	Stream
	W	io.Writer
	Err	error	// unused
}

func (w StreamWriter) Write(src []byte) (n int, err error) {
//line /usr/local/go/src/crypto/cipher/io.go:36
	_go_fuzz_dep_.CoverTab[1673]++
							c := make([]byte, len(src))
							w.S.XORKeyStream(c, src)
							n, err = w.W.Write(c)
							if n != len(src) && func() bool {
//line /usr/local/go/src/crypto/cipher/io.go:40
		_go_fuzz_dep_.CoverTab[1675]++
//line /usr/local/go/src/crypto/cipher/io.go:40
		return err == nil
//line /usr/local/go/src/crypto/cipher/io.go:40
		// _ = "end of CoverTab[1675]"
//line /usr/local/go/src/crypto/cipher/io.go:40
	}() {
//line /usr/local/go/src/crypto/cipher/io.go:40
		_go_fuzz_dep_.CoverTab[1676]++
								err = io.ErrShortWrite
//line /usr/local/go/src/crypto/cipher/io.go:41
		// _ = "end of CoverTab[1676]"
	} else {
//line /usr/local/go/src/crypto/cipher/io.go:42
		_go_fuzz_dep_.CoverTab[1677]++
//line /usr/local/go/src/crypto/cipher/io.go:42
		// _ = "end of CoverTab[1677]"
//line /usr/local/go/src/crypto/cipher/io.go:42
	}
//line /usr/local/go/src/crypto/cipher/io.go:42
	// _ = "end of CoverTab[1673]"
//line /usr/local/go/src/crypto/cipher/io.go:42
	_go_fuzz_dep_.CoverTab[1674]++
							return
//line /usr/local/go/src/crypto/cipher/io.go:43
	// _ = "end of CoverTab[1674]"
}

// Close closes the underlying Writer and returns its Close return value, if the Writer
//line /usr/local/go/src/crypto/cipher/io.go:46
// is also an io.Closer. Otherwise it returns nil.
//line /usr/local/go/src/crypto/cipher/io.go:48
func (w StreamWriter) Close() error {
//line /usr/local/go/src/crypto/cipher/io.go:48
	_go_fuzz_dep_.CoverTab[1678]++
							if c, ok := w.W.(io.Closer); ok {
//line /usr/local/go/src/crypto/cipher/io.go:49
		_go_fuzz_dep_.CoverTab[1680]++
								return c.Close()
//line /usr/local/go/src/crypto/cipher/io.go:50
		// _ = "end of CoverTab[1680]"
	} else {
//line /usr/local/go/src/crypto/cipher/io.go:51
		_go_fuzz_dep_.CoverTab[1681]++
//line /usr/local/go/src/crypto/cipher/io.go:51
		// _ = "end of CoverTab[1681]"
//line /usr/local/go/src/crypto/cipher/io.go:51
	}
//line /usr/local/go/src/crypto/cipher/io.go:51
	// _ = "end of CoverTab[1678]"
//line /usr/local/go/src/crypto/cipher/io.go:51
	_go_fuzz_dep_.CoverTab[1679]++
							return nil
//line /usr/local/go/src/crypto/cipher/io.go:52
	// _ = "end of CoverTab[1679]"
}

//line /usr/local/go/src/crypto/cipher/io.go:53
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/io.go:53
var _ = _go_fuzz_dep_.CoverTab
