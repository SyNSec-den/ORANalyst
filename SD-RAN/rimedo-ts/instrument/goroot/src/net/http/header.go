// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/header.go:5
package http

//line /usr/local/go/src/net/http/header.go:5
import (
//line /usr/local/go/src/net/http/header.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/header.go:5
)
//line /usr/local/go/src/net/http/header.go:5
import (
//line /usr/local/go/src/net/http/header.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/header.go:5
)

import (
	"io"
	"net/http/httptrace"
	"net/http/internal/ascii"
	"net/textproto"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http/httpguts"
)

// A Header represents the key-value pairs in an HTTP header.
//line /usr/local/go/src/net/http/header.go:20
//
//line /usr/local/go/src/net/http/header.go:20
// The keys should be in canonical form, as returned by
//line /usr/local/go/src/net/http/header.go:20
// CanonicalHeaderKey.
//line /usr/local/go/src/net/http/header.go:24
type Header map[string][]string

// Add adds the key, value pair to the header.
//line /usr/local/go/src/net/http/header.go:26
// It appends to any existing values associated with key.
//line /usr/local/go/src/net/http/header.go:26
// The key is case insensitive; it is canonicalized by
//line /usr/local/go/src/net/http/header.go:26
// CanonicalHeaderKey.
//line /usr/local/go/src/net/http/header.go:30
func (h Header) Add(key, value string) {
//line /usr/local/go/src/net/http/header.go:30
	_go_fuzz_dep_.CoverTab[41358]++
						textproto.MIMEHeader(h).Add(key, value)
//line /usr/local/go/src/net/http/header.go:31
	// _ = "end of CoverTab[41358]"
}

// Set sets the header entries associated with key to the
//line /usr/local/go/src/net/http/header.go:34
// single element value. It replaces any existing values
//line /usr/local/go/src/net/http/header.go:34
// associated with key. The key is case insensitive; it is
//line /usr/local/go/src/net/http/header.go:34
// canonicalized by textproto.CanonicalMIMEHeaderKey.
//line /usr/local/go/src/net/http/header.go:34
// To use non-canonical keys, assign to the map directly.
//line /usr/local/go/src/net/http/header.go:39
func (h Header) Set(key, value string) {
//line /usr/local/go/src/net/http/header.go:39
	_go_fuzz_dep_.CoverTab[41359]++
						textproto.MIMEHeader(h).Set(key, value)
//line /usr/local/go/src/net/http/header.go:40
	// _ = "end of CoverTab[41359]"
}

// Get gets the first value associated with the given key. If
//line /usr/local/go/src/net/http/header.go:43
// there are no values associated with the key, Get returns "".
//line /usr/local/go/src/net/http/header.go:43
// It is case insensitive; textproto.CanonicalMIMEHeaderKey is
//line /usr/local/go/src/net/http/header.go:43
// used to canonicalize the provided key. Get assumes that all
//line /usr/local/go/src/net/http/header.go:43
// keys are stored in canonical form. To use non-canonical keys,
//line /usr/local/go/src/net/http/header.go:43
// access the map directly.
//line /usr/local/go/src/net/http/header.go:49
func (h Header) Get(key string) string {
//line /usr/local/go/src/net/http/header.go:49
	_go_fuzz_dep_.CoverTab[41360]++
						return textproto.MIMEHeader(h).Get(key)
//line /usr/local/go/src/net/http/header.go:50
	// _ = "end of CoverTab[41360]"
}

// Values returns all values associated with the given key.
//line /usr/local/go/src/net/http/header.go:53
// It is case insensitive; textproto.CanonicalMIMEHeaderKey is
//line /usr/local/go/src/net/http/header.go:53
// used to canonicalize the provided key. To use non-canonical
//line /usr/local/go/src/net/http/header.go:53
// keys, access the map directly.
//line /usr/local/go/src/net/http/header.go:53
// The returned slice is not a copy.
//line /usr/local/go/src/net/http/header.go:58
func (h Header) Values(key string) []string {
//line /usr/local/go/src/net/http/header.go:58
	_go_fuzz_dep_.CoverTab[41361]++
						return textproto.MIMEHeader(h).Values(key)
//line /usr/local/go/src/net/http/header.go:59
	// _ = "end of CoverTab[41361]"
}

// get is like Get, but key must already be in CanonicalHeaderKey form.
func (h Header) get(key string) string {
//line /usr/local/go/src/net/http/header.go:63
	_go_fuzz_dep_.CoverTab[41362]++
						if v := h[key]; len(v) > 0 {
//line /usr/local/go/src/net/http/header.go:64
		_go_fuzz_dep_.CoverTab[41364]++
							return v[0]
//line /usr/local/go/src/net/http/header.go:65
		// _ = "end of CoverTab[41364]"
	} else {
//line /usr/local/go/src/net/http/header.go:66
		_go_fuzz_dep_.CoverTab[41365]++
//line /usr/local/go/src/net/http/header.go:66
		// _ = "end of CoverTab[41365]"
//line /usr/local/go/src/net/http/header.go:66
	}
//line /usr/local/go/src/net/http/header.go:66
	// _ = "end of CoverTab[41362]"
//line /usr/local/go/src/net/http/header.go:66
	_go_fuzz_dep_.CoverTab[41363]++
						return ""
//line /usr/local/go/src/net/http/header.go:67
	// _ = "end of CoverTab[41363]"
}

// has reports whether h has the provided key defined, even if it's
//line /usr/local/go/src/net/http/header.go:70
// set to 0-length slice.
//line /usr/local/go/src/net/http/header.go:72
func (h Header) has(key string) bool {
//line /usr/local/go/src/net/http/header.go:72
	_go_fuzz_dep_.CoverTab[41366]++
						_, ok := h[key]
						return ok
//line /usr/local/go/src/net/http/header.go:74
	// _ = "end of CoverTab[41366]"
}

// Del deletes the values associated with key.
//line /usr/local/go/src/net/http/header.go:77
// The key is case insensitive; it is canonicalized by
//line /usr/local/go/src/net/http/header.go:77
// CanonicalHeaderKey.
//line /usr/local/go/src/net/http/header.go:80
func (h Header) Del(key string) {
//line /usr/local/go/src/net/http/header.go:80
	_go_fuzz_dep_.CoverTab[41367]++
						textproto.MIMEHeader(h).Del(key)
//line /usr/local/go/src/net/http/header.go:81
	// _ = "end of CoverTab[41367]"
}

// Write writes a header in wire format.
func (h Header) Write(w io.Writer) error {
//line /usr/local/go/src/net/http/header.go:85
	_go_fuzz_dep_.CoverTab[41368]++
						return h.write(w, nil)
//line /usr/local/go/src/net/http/header.go:86
	// _ = "end of CoverTab[41368]"
}

func (h Header) write(w io.Writer, trace *httptrace.ClientTrace) error {
//line /usr/local/go/src/net/http/header.go:89
	_go_fuzz_dep_.CoverTab[41369]++
						return h.writeSubset(w, nil, trace)
//line /usr/local/go/src/net/http/header.go:90
	// _ = "end of CoverTab[41369]"
}

// Clone returns a copy of h or nil if h is nil.
func (h Header) Clone() Header {
//line /usr/local/go/src/net/http/header.go:94
	_go_fuzz_dep_.CoverTab[41370]++
						if h == nil {
//line /usr/local/go/src/net/http/header.go:95
		_go_fuzz_dep_.CoverTab[41374]++
							return nil
//line /usr/local/go/src/net/http/header.go:96
		// _ = "end of CoverTab[41374]"
	} else {
//line /usr/local/go/src/net/http/header.go:97
		_go_fuzz_dep_.CoverTab[41375]++
//line /usr/local/go/src/net/http/header.go:97
		// _ = "end of CoverTab[41375]"
//line /usr/local/go/src/net/http/header.go:97
	}
//line /usr/local/go/src/net/http/header.go:97
		// _ = "end of CoverTab[41370]"
//line /usr/local/go/src/net/http/header.go:97
		_go_fuzz_dep_.CoverTab[41371]++

//line /usr/local/go/src/net/http/header.go:100
	nv := 0
	for _, vv := range h {
//line /usr/local/go/src/net/http/header.go:101
		_go_fuzz_dep_.CoverTab[41376]++
								nv += len(vv)
//line /usr/local/go/src/net/http/header.go:102
		// _ = "end of CoverTab[41376]"
	}
//line /usr/local/go/src/net/http/header.go:103
	// _ = "end of CoverTab[41371]"
//line /usr/local/go/src/net/http/header.go:103
	_go_fuzz_dep_.CoverTab[41372]++
							sv := make([]string, nv)
							h2 := make(Header, len(h))
							for k, vv := range h {
//line /usr/local/go/src/net/http/header.go:106
		_go_fuzz_dep_.CoverTab[41377]++
								if vv == nil {
//line /usr/local/go/src/net/http/header.go:107
			_go_fuzz_dep_.CoverTab[41379]++

//line /usr/local/go/src/net/http/header.go:110
			h2[k] = nil
									continue
//line /usr/local/go/src/net/http/header.go:111
			// _ = "end of CoverTab[41379]"
		} else {
//line /usr/local/go/src/net/http/header.go:112
			_go_fuzz_dep_.CoverTab[41380]++
//line /usr/local/go/src/net/http/header.go:112
			// _ = "end of CoverTab[41380]"
//line /usr/local/go/src/net/http/header.go:112
		}
//line /usr/local/go/src/net/http/header.go:112
		// _ = "end of CoverTab[41377]"
//line /usr/local/go/src/net/http/header.go:112
		_go_fuzz_dep_.CoverTab[41378]++
								n := copy(sv, vv)
								h2[k] = sv[:n:n]
								sv = sv[n:]
//line /usr/local/go/src/net/http/header.go:115
		// _ = "end of CoverTab[41378]"
	}
//line /usr/local/go/src/net/http/header.go:116
	// _ = "end of CoverTab[41372]"
//line /usr/local/go/src/net/http/header.go:116
	_go_fuzz_dep_.CoverTab[41373]++
							return h2
//line /usr/local/go/src/net/http/header.go:117
	// _ = "end of CoverTab[41373]"
}

var timeFormats = []string{
	TimeFormat,
	time.RFC850,
	time.ANSIC,
}

// ParseTime parses a time header (such as the Date: header),
//line /usr/local/go/src/net/http/header.go:126
// trying each of the three formats allowed by HTTP/1.1:
//line /usr/local/go/src/net/http/header.go:126
// TimeFormat, time.RFC850, and time.ANSIC.
//line /usr/local/go/src/net/http/header.go:129
func ParseTime(text string) (t time.Time, err error) {
//line /usr/local/go/src/net/http/header.go:129
	_go_fuzz_dep_.CoverTab[41381]++
							for _, layout := range timeFormats {
//line /usr/local/go/src/net/http/header.go:130
		_go_fuzz_dep_.CoverTab[41383]++
								t, err = time.Parse(layout, text)
								if err == nil {
//line /usr/local/go/src/net/http/header.go:132
			_go_fuzz_dep_.CoverTab[41384]++
									return
//line /usr/local/go/src/net/http/header.go:133
			// _ = "end of CoverTab[41384]"
		} else {
//line /usr/local/go/src/net/http/header.go:134
			_go_fuzz_dep_.CoverTab[41385]++
//line /usr/local/go/src/net/http/header.go:134
			// _ = "end of CoverTab[41385]"
//line /usr/local/go/src/net/http/header.go:134
		}
//line /usr/local/go/src/net/http/header.go:134
		// _ = "end of CoverTab[41383]"
	}
//line /usr/local/go/src/net/http/header.go:135
	// _ = "end of CoverTab[41381]"
//line /usr/local/go/src/net/http/header.go:135
	_go_fuzz_dep_.CoverTab[41382]++
							return
//line /usr/local/go/src/net/http/header.go:136
	// _ = "end of CoverTab[41382]"
}

var headerNewlineToSpace = strings.NewReplacer("\n", " ", "\r", " ")

// stringWriter implements WriteString on a Writer.
type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (n int, err error) {
//line /usr/local/go/src/net/http/header.go:146
	_go_fuzz_dep_.CoverTab[41386]++
							return w.w.Write([]byte(s))
//line /usr/local/go/src/net/http/header.go:147
	// _ = "end of CoverTab[41386]"
}

type keyValues struct {
	key	string
	values	[]string
}

// A headerSorter implements sort.Interface by sorting a []keyValues
//line /usr/local/go/src/net/http/header.go:155
// by key. It's used as a pointer, so it can fit in a sort.Interface
//line /usr/local/go/src/net/http/header.go:155
// interface value without allocation.
//line /usr/local/go/src/net/http/header.go:158
type headerSorter struct {
	kvs []keyValues
}

func (s *headerSorter) Len() int {
//line /usr/local/go/src/net/http/header.go:162
	_go_fuzz_dep_.CoverTab[41387]++
//line /usr/local/go/src/net/http/header.go:162
	return len(s.kvs)
//line /usr/local/go/src/net/http/header.go:162
	// _ = "end of CoverTab[41387]"
//line /usr/local/go/src/net/http/header.go:162
}
func (s *headerSorter) Swap(i, j int) {
//line /usr/local/go/src/net/http/header.go:163
	_go_fuzz_dep_.CoverTab[41388]++
//line /usr/local/go/src/net/http/header.go:163
	s.kvs[i], s.kvs[j] = s.kvs[j], s.kvs[i]
//line /usr/local/go/src/net/http/header.go:163
	// _ = "end of CoverTab[41388]"
//line /usr/local/go/src/net/http/header.go:163
}
func (s *headerSorter) Less(i, j int) bool {
//line /usr/local/go/src/net/http/header.go:164
	_go_fuzz_dep_.CoverTab[41389]++
//line /usr/local/go/src/net/http/header.go:164
	return s.kvs[i].key < s.kvs[j].key
//line /usr/local/go/src/net/http/header.go:164
	// _ = "end of CoverTab[41389]"
//line /usr/local/go/src/net/http/header.go:164
}

var headerSorterPool = sync.Pool{
	New: func() any { _go_fuzz_dep_.CoverTab[41390]++; return new(headerSorter); // _ = "end of CoverTab[41390]" },
}

// sortedKeyValues returns h's keys sorted in the returned kvs
//line /usr/local/go/src/net/http/header.go:170
// slice. The headerSorter used to sort is also returned, for possible
//line /usr/local/go/src/net/http/header.go:170
// return to headerSorterCache.
//line /usr/local/go/src/net/http/header.go:173
func (h Header) sortedKeyValues(exclude map[string]bool) (kvs []keyValues, hs *headerSorter) {
//line /usr/local/go/src/net/http/header.go:173
	_go_fuzz_dep_.CoverTab[41391]++
							hs = headerSorterPool.Get().(*headerSorter)
							if cap(hs.kvs) < len(h) {
//line /usr/local/go/src/net/http/header.go:175
		_go_fuzz_dep_.CoverTab[41394]++
								hs.kvs = make([]keyValues, 0, len(h))
//line /usr/local/go/src/net/http/header.go:176
		// _ = "end of CoverTab[41394]"
	} else {
//line /usr/local/go/src/net/http/header.go:177
		_go_fuzz_dep_.CoverTab[41395]++
//line /usr/local/go/src/net/http/header.go:177
		// _ = "end of CoverTab[41395]"
//line /usr/local/go/src/net/http/header.go:177
	}
//line /usr/local/go/src/net/http/header.go:177
	// _ = "end of CoverTab[41391]"
//line /usr/local/go/src/net/http/header.go:177
	_go_fuzz_dep_.CoverTab[41392]++
							kvs = hs.kvs[:0]
							for k, vv := range h {
//line /usr/local/go/src/net/http/header.go:179
		_go_fuzz_dep_.CoverTab[41396]++
								if !exclude[k] {
//line /usr/local/go/src/net/http/header.go:180
			_go_fuzz_dep_.CoverTab[41397]++
									kvs = append(kvs, keyValues{k, vv})
//line /usr/local/go/src/net/http/header.go:181
			// _ = "end of CoverTab[41397]"
		} else {
//line /usr/local/go/src/net/http/header.go:182
			_go_fuzz_dep_.CoverTab[41398]++
//line /usr/local/go/src/net/http/header.go:182
			// _ = "end of CoverTab[41398]"
//line /usr/local/go/src/net/http/header.go:182
		}
//line /usr/local/go/src/net/http/header.go:182
		// _ = "end of CoverTab[41396]"
	}
//line /usr/local/go/src/net/http/header.go:183
	// _ = "end of CoverTab[41392]"
//line /usr/local/go/src/net/http/header.go:183
	_go_fuzz_dep_.CoverTab[41393]++
							hs.kvs = kvs
							sort.Sort(hs)
							return kvs, hs
//line /usr/local/go/src/net/http/header.go:186
	// _ = "end of CoverTab[41393]"
}

// WriteSubset writes a header in wire format.
//line /usr/local/go/src/net/http/header.go:189
// If exclude is not nil, keys where exclude[key] == true are not written.
//line /usr/local/go/src/net/http/header.go:189
// Keys are not canonicalized before checking the exclude map.
//line /usr/local/go/src/net/http/header.go:192
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error {
//line /usr/local/go/src/net/http/header.go:192
	_go_fuzz_dep_.CoverTab[41399]++
							return h.writeSubset(w, exclude, nil)
//line /usr/local/go/src/net/http/header.go:193
	// _ = "end of CoverTab[41399]"
}

func (h Header) writeSubset(w io.Writer, exclude map[string]bool, trace *httptrace.ClientTrace) error {
//line /usr/local/go/src/net/http/header.go:196
	_go_fuzz_dep_.CoverTab[41400]++
							ws, ok := w.(io.StringWriter)
							if !ok {
//line /usr/local/go/src/net/http/header.go:198
		_go_fuzz_dep_.CoverTab[41403]++
								ws = stringWriter{w}
//line /usr/local/go/src/net/http/header.go:199
		// _ = "end of CoverTab[41403]"
	} else {
//line /usr/local/go/src/net/http/header.go:200
		_go_fuzz_dep_.CoverTab[41404]++
//line /usr/local/go/src/net/http/header.go:200
		// _ = "end of CoverTab[41404]"
//line /usr/local/go/src/net/http/header.go:200
	}
//line /usr/local/go/src/net/http/header.go:200
	// _ = "end of CoverTab[41400]"
//line /usr/local/go/src/net/http/header.go:200
	_go_fuzz_dep_.CoverTab[41401]++
							kvs, sorter := h.sortedKeyValues(exclude)
							var formattedVals []string
							for _, kv := range kvs {
//line /usr/local/go/src/net/http/header.go:203
		_go_fuzz_dep_.CoverTab[41405]++
								if !httpguts.ValidHeaderFieldName(kv.key) {
//line /usr/local/go/src/net/http/header.go:204
			_go_fuzz_dep_.CoverTab[41408]++

//line /usr/local/go/src/net/http/header.go:209
			continue
//line /usr/local/go/src/net/http/header.go:209
			// _ = "end of CoverTab[41408]"
		} else {
//line /usr/local/go/src/net/http/header.go:210
			_go_fuzz_dep_.CoverTab[41409]++
//line /usr/local/go/src/net/http/header.go:210
			// _ = "end of CoverTab[41409]"
//line /usr/local/go/src/net/http/header.go:210
		}
//line /usr/local/go/src/net/http/header.go:210
		// _ = "end of CoverTab[41405]"
//line /usr/local/go/src/net/http/header.go:210
		_go_fuzz_dep_.CoverTab[41406]++
								for _, v := range kv.values {
//line /usr/local/go/src/net/http/header.go:211
			_go_fuzz_dep_.CoverTab[41410]++
									v = headerNewlineToSpace.Replace(v)
									v = textproto.TrimString(v)
									for _, s := range []string{kv.key, ": ", v, "\r\n"} {
//line /usr/local/go/src/net/http/header.go:214
				_go_fuzz_dep_.CoverTab[41412]++
										if _, err := ws.WriteString(s); err != nil {
//line /usr/local/go/src/net/http/header.go:215
					_go_fuzz_dep_.CoverTab[41413]++
											headerSorterPool.Put(sorter)
											return err
//line /usr/local/go/src/net/http/header.go:217
					// _ = "end of CoverTab[41413]"
				} else {
//line /usr/local/go/src/net/http/header.go:218
					_go_fuzz_dep_.CoverTab[41414]++
//line /usr/local/go/src/net/http/header.go:218
					// _ = "end of CoverTab[41414]"
//line /usr/local/go/src/net/http/header.go:218
				}
//line /usr/local/go/src/net/http/header.go:218
				// _ = "end of CoverTab[41412]"
			}
//line /usr/local/go/src/net/http/header.go:219
			// _ = "end of CoverTab[41410]"
//line /usr/local/go/src/net/http/header.go:219
			_go_fuzz_dep_.CoverTab[41411]++
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/header.go:220
				_go_fuzz_dep_.CoverTab[41415]++
//line /usr/local/go/src/net/http/header.go:220
				return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/header.go:220
				// _ = "end of CoverTab[41415]"
//line /usr/local/go/src/net/http/header.go:220
			}() {
//line /usr/local/go/src/net/http/header.go:220
				_go_fuzz_dep_.CoverTab[41416]++
										formattedVals = append(formattedVals, v)
//line /usr/local/go/src/net/http/header.go:221
				// _ = "end of CoverTab[41416]"
			} else {
//line /usr/local/go/src/net/http/header.go:222
				_go_fuzz_dep_.CoverTab[41417]++
//line /usr/local/go/src/net/http/header.go:222
				// _ = "end of CoverTab[41417]"
//line /usr/local/go/src/net/http/header.go:222
			}
//line /usr/local/go/src/net/http/header.go:222
			// _ = "end of CoverTab[41411]"
		}
//line /usr/local/go/src/net/http/header.go:223
		// _ = "end of CoverTab[41406]"
//line /usr/local/go/src/net/http/header.go:223
		_go_fuzz_dep_.CoverTab[41407]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/header.go:224
			_go_fuzz_dep_.CoverTab[41418]++
//line /usr/local/go/src/net/http/header.go:224
			return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/header.go:224
			// _ = "end of CoverTab[41418]"
//line /usr/local/go/src/net/http/header.go:224
		}() {
//line /usr/local/go/src/net/http/header.go:224
			_go_fuzz_dep_.CoverTab[41419]++
									trace.WroteHeaderField(kv.key, formattedVals)
									formattedVals = nil
//line /usr/local/go/src/net/http/header.go:226
			// _ = "end of CoverTab[41419]"
		} else {
//line /usr/local/go/src/net/http/header.go:227
			_go_fuzz_dep_.CoverTab[41420]++
//line /usr/local/go/src/net/http/header.go:227
			// _ = "end of CoverTab[41420]"
//line /usr/local/go/src/net/http/header.go:227
		}
//line /usr/local/go/src/net/http/header.go:227
		// _ = "end of CoverTab[41407]"
	}
//line /usr/local/go/src/net/http/header.go:228
	// _ = "end of CoverTab[41401]"
//line /usr/local/go/src/net/http/header.go:228
	_go_fuzz_dep_.CoverTab[41402]++
							headerSorterPool.Put(sorter)
							return nil
//line /usr/local/go/src/net/http/header.go:230
	// _ = "end of CoverTab[41402]"
}

// CanonicalHeaderKey returns the canonical format of the
//line /usr/local/go/src/net/http/header.go:233
// header key s. The canonicalization converts the first
//line /usr/local/go/src/net/http/header.go:233
// letter and any letter following a hyphen to upper case;
//line /usr/local/go/src/net/http/header.go:233
// the rest are converted to lowercase. For example, the
//line /usr/local/go/src/net/http/header.go:233
// canonical key for "accept-encoding" is "Accept-Encoding".
//line /usr/local/go/src/net/http/header.go:233
// If s contains a space or invalid header field bytes, it is
//line /usr/local/go/src/net/http/header.go:233
// returned without modifications.
//line /usr/local/go/src/net/http/header.go:240
func CanonicalHeaderKey(s string) string {
//line /usr/local/go/src/net/http/header.go:240
	_go_fuzz_dep_.CoverTab[41421]++
//line /usr/local/go/src/net/http/header.go:240
	return textproto.CanonicalMIMEHeaderKey(s)
//line /usr/local/go/src/net/http/header.go:240
	// _ = "end of CoverTab[41421]"
//line /usr/local/go/src/net/http/header.go:240
}

// hasToken reports whether token appears with v, ASCII
//line /usr/local/go/src/net/http/header.go:242
// case-insensitive, with space or comma boundaries.
//line /usr/local/go/src/net/http/header.go:242
// token must be all lowercase.
//line /usr/local/go/src/net/http/header.go:242
// v may contain mixed cased.
//line /usr/local/go/src/net/http/header.go:246
func hasToken(v, token string) bool {
//line /usr/local/go/src/net/http/header.go:246
	_go_fuzz_dep_.CoverTab[41422]++
							if len(token) > len(v) || func() bool {
//line /usr/local/go/src/net/http/header.go:247
		_go_fuzz_dep_.CoverTab[41426]++
//line /usr/local/go/src/net/http/header.go:247
		return token == ""
//line /usr/local/go/src/net/http/header.go:247
		// _ = "end of CoverTab[41426]"
//line /usr/local/go/src/net/http/header.go:247
	}() {
//line /usr/local/go/src/net/http/header.go:247
		_go_fuzz_dep_.CoverTab[41427]++
								return false
//line /usr/local/go/src/net/http/header.go:248
		// _ = "end of CoverTab[41427]"
	} else {
//line /usr/local/go/src/net/http/header.go:249
		_go_fuzz_dep_.CoverTab[41428]++
//line /usr/local/go/src/net/http/header.go:249
		// _ = "end of CoverTab[41428]"
//line /usr/local/go/src/net/http/header.go:249
	}
//line /usr/local/go/src/net/http/header.go:249
	// _ = "end of CoverTab[41422]"
//line /usr/local/go/src/net/http/header.go:249
	_go_fuzz_dep_.CoverTab[41423]++
							if v == token {
//line /usr/local/go/src/net/http/header.go:250
		_go_fuzz_dep_.CoverTab[41429]++
								return true
//line /usr/local/go/src/net/http/header.go:251
		// _ = "end of CoverTab[41429]"
	} else {
//line /usr/local/go/src/net/http/header.go:252
		_go_fuzz_dep_.CoverTab[41430]++
//line /usr/local/go/src/net/http/header.go:252
		// _ = "end of CoverTab[41430]"
//line /usr/local/go/src/net/http/header.go:252
	}
//line /usr/local/go/src/net/http/header.go:252
	// _ = "end of CoverTab[41423]"
//line /usr/local/go/src/net/http/header.go:252
	_go_fuzz_dep_.CoverTab[41424]++
							for sp := 0; sp <= len(v)-len(token); sp++ {
//line /usr/local/go/src/net/http/header.go:253
		_go_fuzz_dep_.CoverTab[41431]++

//line /usr/local/go/src/net/http/header.go:260
		if b := v[sp]; b != token[0] && func() bool {
//line /usr/local/go/src/net/http/header.go:260
			_go_fuzz_dep_.CoverTab[41435]++
//line /usr/local/go/src/net/http/header.go:260
			return b|0x20 != token[0]
//line /usr/local/go/src/net/http/header.go:260
			// _ = "end of CoverTab[41435]"
//line /usr/local/go/src/net/http/header.go:260
		}() {
//line /usr/local/go/src/net/http/header.go:260
			_go_fuzz_dep_.CoverTab[41436]++
									continue
//line /usr/local/go/src/net/http/header.go:261
			// _ = "end of CoverTab[41436]"
		} else {
//line /usr/local/go/src/net/http/header.go:262
			_go_fuzz_dep_.CoverTab[41437]++
//line /usr/local/go/src/net/http/header.go:262
			// _ = "end of CoverTab[41437]"
//line /usr/local/go/src/net/http/header.go:262
		}
//line /usr/local/go/src/net/http/header.go:262
		// _ = "end of CoverTab[41431]"
//line /usr/local/go/src/net/http/header.go:262
		_go_fuzz_dep_.CoverTab[41432]++

								if sp > 0 && func() bool {
//line /usr/local/go/src/net/http/header.go:264
			_go_fuzz_dep_.CoverTab[41438]++
//line /usr/local/go/src/net/http/header.go:264
			return !isTokenBoundary(v[sp-1])
//line /usr/local/go/src/net/http/header.go:264
			// _ = "end of CoverTab[41438]"
//line /usr/local/go/src/net/http/header.go:264
		}() {
//line /usr/local/go/src/net/http/header.go:264
			_go_fuzz_dep_.CoverTab[41439]++
									continue
//line /usr/local/go/src/net/http/header.go:265
			// _ = "end of CoverTab[41439]"
		} else {
//line /usr/local/go/src/net/http/header.go:266
			_go_fuzz_dep_.CoverTab[41440]++
//line /usr/local/go/src/net/http/header.go:266
			// _ = "end of CoverTab[41440]"
//line /usr/local/go/src/net/http/header.go:266
		}
//line /usr/local/go/src/net/http/header.go:266
		// _ = "end of CoverTab[41432]"
//line /usr/local/go/src/net/http/header.go:266
		_go_fuzz_dep_.CoverTab[41433]++

								if endPos := sp + len(token); endPos != len(v) && func() bool {
//line /usr/local/go/src/net/http/header.go:268
			_go_fuzz_dep_.CoverTab[41441]++
//line /usr/local/go/src/net/http/header.go:268
			return !isTokenBoundary(v[endPos])
//line /usr/local/go/src/net/http/header.go:268
			// _ = "end of CoverTab[41441]"
//line /usr/local/go/src/net/http/header.go:268
		}() {
//line /usr/local/go/src/net/http/header.go:268
			_go_fuzz_dep_.CoverTab[41442]++
									continue
//line /usr/local/go/src/net/http/header.go:269
			// _ = "end of CoverTab[41442]"
		} else {
//line /usr/local/go/src/net/http/header.go:270
			_go_fuzz_dep_.CoverTab[41443]++
//line /usr/local/go/src/net/http/header.go:270
			// _ = "end of CoverTab[41443]"
//line /usr/local/go/src/net/http/header.go:270
		}
//line /usr/local/go/src/net/http/header.go:270
		// _ = "end of CoverTab[41433]"
//line /usr/local/go/src/net/http/header.go:270
		_go_fuzz_dep_.CoverTab[41434]++
								if ascii.EqualFold(v[sp:sp+len(token)], token) {
//line /usr/local/go/src/net/http/header.go:271
			_go_fuzz_dep_.CoverTab[41444]++
									return true
//line /usr/local/go/src/net/http/header.go:272
			// _ = "end of CoverTab[41444]"
		} else {
//line /usr/local/go/src/net/http/header.go:273
			_go_fuzz_dep_.CoverTab[41445]++
//line /usr/local/go/src/net/http/header.go:273
			// _ = "end of CoverTab[41445]"
//line /usr/local/go/src/net/http/header.go:273
		}
//line /usr/local/go/src/net/http/header.go:273
		// _ = "end of CoverTab[41434]"
	}
//line /usr/local/go/src/net/http/header.go:274
	// _ = "end of CoverTab[41424]"
//line /usr/local/go/src/net/http/header.go:274
	_go_fuzz_dep_.CoverTab[41425]++
							return false
//line /usr/local/go/src/net/http/header.go:275
	// _ = "end of CoverTab[41425]"
}

func isTokenBoundary(b byte) bool {
//line /usr/local/go/src/net/http/header.go:278
	_go_fuzz_dep_.CoverTab[41446]++
							return b == ' ' || func() bool {
//line /usr/local/go/src/net/http/header.go:279
		_go_fuzz_dep_.CoverTab[41447]++
//line /usr/local/go/src/net/http/header.go:279
		return b == ','
//line /usr/local/go/src/net/http/header.go:279
		// _ = "end of CoverTab[41447]"
//line /usr/local/go/src/net/http/header.go:279
	}() || func() bool {
//line /usr/local/go/src/net/http/header.go:279
		_go_fuzz_dep_.CoverTab[41448]++
//line /usr/local/go/src/net/http/header.go:279
		return b == '\t'
//line /usr/local/go/src/net/http/header.go:279
		// _ = "end of CoverTab[41448]"
//line /usr/local/go/src/net/http/header.go:279
	}()
//line /usr/local/go/src/net/http/header.go:279
	// _ = "end of CoverTab[41446]"
}

//line /usr/local/go/src/net/http/header.go:280
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/header.go:280
var _ = _go_fuzz_dep_.CoverTab
