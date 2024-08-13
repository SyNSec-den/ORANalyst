// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/textproto/header.go:5
package textproto

//line /usr/local/go/src/net/textproto/header.go:5
import (
//line /usr/local/go/src/net/textproto/header.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/textproto/header.go:5
)
//line /usr/local/go/src/net/textproto/header.go:5
import (
//line /usr/local/go/src/net/textproto/header.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/textproto/header.go:5
)

// A MIMEHeader represents a MIME-style header mapping
//line /usr/local/go/src/net/textproto/header.go:7
// keys to sets of values.
//line /usr/local/go/src/net/textproto/header.go:9
type MIMEHeader map[string][]string

// Add adds the key, value pair to the header.
//line /usr/local/go/src/net/textproto/header.go:11
// It appends to any existing values associated with key.
//line /usr/local/go/src/net/textproto/header.go:13
func (h MIMEHeader) Add(key, value string) {
//line /usr/local/go/src/net/textproto/header.go:13
	_go_fuzz_dep_.CoverTab[34503]++
							key = CanonicalMIMEHeaderKey(key)
							h[key] = append(h[key], value)
//line /usr/local/go/src/net/textproto/header.go:15
	// _ = "end of CoverTab[34503]"
}

// Set sets the header entries associated with key to
//line /usr/local/go/src/net/textproto/header.go:18
// the single element value. It replaces any existing
//line /usr/local/go/src/net/textproto/header.go:18
// values associated with key.
//line /usr/local/go/src/net/textproto/header.go:21
func (h MIMEHeader) Set(key, value string) {
//line /usr/local/go/src/net/textproto/header.go:21
	_go_fuzz_dep_.CoverTab[34504]++
							h[CanonicalMIMEHeaderKey(key)] = []string{value}
//line /usr/local/go/src/net/textproto/header.go:22
	// _ = "end of CoverTab[34504]"
}

// Get gets the first value associated with the given key.
//line /usr/local/go/src/net/textproto/header.go:25
// It is case insensitive; CanonicalMIMEHeaderKey is used
//line /usr/local/go/src/net/textproto/header.go:25
// to canonicalize the provided key.
//line /usr/local/go/src/net/textproto/header.go:25
// If there are no values associated with the key, Get returns "".
//line /usr/local/go/src/net/textproto/header.go:25
// To use non-canonical keys, access the map directly.
//line /usr/local/go/src/net/textproto/header.go:30
func (h MIMEHeader) Get(key string) string {
//line /usr/local/go/src/net/textproto/header.go:30
	_go_fuzz_dep_.CoverTab[34505]++
							if h == nil {
//line /usr/local/go/src/net/textproto/header.go:31
		_go_fuzz_dep_.CoverTab[34508]++
								return ""
//line /usr/local/go/src/net/textproto/header.go:32
		// _ = "end of CoverTab[34508]"
	} else {
//line /usr/local/go/src/net/textproto/header.go:33
		_go_fuzz_dep_.CoverTab[34509]++
//line /usr/local/go/src/net/textproto/header.go:33
		// _ = "end of CoverTab[34509]"
//line /usr/local/go/src/net/textproto/header.go:33
	}
//line /usr/local/go/src/net/textproto/header.go:33
	// _ = "end of CoverTab[34505]"
//line /usr/local/go/src/net/textproto/header.go:33
	_go_fuzz_dep_.CoverTab[34506]++
							v := h[CanonicalMIMEHeaderKey(key)]
							if len(v) == 0 {
//line /usr/local/go/src/net/textproto/header.go:35
		_go_fuzz_dep_.CoverTab[34510]++
								return ""
//line /usr/local/go/src/net/textproto/header.go:36
		// _ = "end of CoverTab[34510]"
	} else {
//line /usr/local/go/src/net/textproto/header.go:37
		_go_fuzz_dep_.CoverTab[34511]++
//line /usr/local/go/src/net/textproto/header.go:37
		// _ = "end of CoverTab[34511]"
//line /usr/local/go/src/net/textproto/header.go:37
	}
//line /usr/local/go/src/net/textproto/header.go:37
	// _ = "end of CoverTab[34506]"
//line /usr/local/go/src/net/textproto/header.go:37
	_go_fuzz_dep_.CoverTab[34507]++
							return v[0]
//line /usr/local/go/src/net/textproto/header.go:38
	// _ = "end of CoverTab[34507]"
}

// Values returns all values associated with the given key.
//line /usr/local/go/src/net/textproto/header.go:41
// It is case insensitive; CanonicalMIMEHeaderKey is
//line /usr/local/go/src/net/textproto/header.go:41
// used to canonicalize the provided key. To use non-canonical
//line /usr/local/go/src/net/textproto/header.go:41
// keys, access the map directly.
//line /usr/local/go/src/net/textproto/header.go:41
// The returned slice is not a copy.
//line /usr/local/go/src/net/textproto/header.go:46
func (h MIMEHeader) Values(key string) []string {
//line /usr/local/go/src/net/textproto/header.go:46
	_go_fuzz_dep_.CoverTab[34512]++
							if h == nil {
//line /usr/local/go/src/net/textproto/header.go:47
		_go_fuzz_dep_.CoverTab[34514]++
								return nil
//line /usr/local/go/src/net/textproto/header.go:48
		// _ = "end of CoverTab[34514]"
	} else {
//line /usr/local/go/src/net/textproto/header.go:49
		_go_fuzz_dep_.CoverTab[34515]++
//line /usr/local/go/src/net/textproto/header.go:49
		// _ = "end of CoverTab[34515]"
//line /usr/local/go/src/net/textproto/header.go:49
	}
//line /usr/local/go/src/net/textproto/header.go:49
	// _ = "end of CoverTab[34512]"
//line /usr/local/go/src/net/textproto/header.go:49
	_go_fuzz_dep_.CoverTab[34513]++
							return h[CanonicalMIMEHeaderKey(key)]
//line /usr/local/go/src/net/textproto/header.go:50
	// _ = "end of CoverTab[34513]"
}

// Del deletes the values associated with key.
func (h MIMEHeader) Del(key string) {
//line /usr/local/go/src/net/textproto/header.go:54
	_go_fuzz_dep_.CoverTab[34516]++
							delete(h, CanonicalMIMEHeaderKey(key))
//line /usr/local/go/src/net/textproto/header.go:55
	// _ = "end of CoverTab[34516]"
}

//line /usr/local/go/src/net/textproto/header.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/textproto/header.go:56
var _ = _go_fuzz_dep_.CoverTab
