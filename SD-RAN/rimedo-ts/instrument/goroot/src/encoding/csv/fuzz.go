// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gofuzz

//line /usr/local/go/src/encoding/csv/fuzz.go:7
package csv

//line /usr/local/go/src/encoding/csv/fuzz.go:7
import (
//line /usr/local/go/src/encoding/csv/fuzz.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/csv/fuzz.go:7
)
//line /usr/local/go/src/encoding/csv/fuzz.go:7
import (
//line /usr/local/go/src/encoding/csv/fuzz.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/csv/fuzz.go:7
)

import (
	"bytes"
	"fmt"
	"reflect"
)

func Fuzz(data []byte) int {
//line /usr/local/go/src/encoding/csv/fuzz.go:15
	_go_fuzz_dep_.CoverTab[114760]++
							score := 0
							buf := new(bytes.Buffer)

							for _, tt := range []Reader{
		{},
		{Comma: ';'},
		{Comma: '\t'},
		{LazyQuotes: true},
		{TrimLeadingSpace: true},
		{Comment: '#'},
		{Comment: ';'},
	} {
//line /usr/local/go/src/encoding/csv/fuzz.go:27
		_go_fuzz_dep_.CoverTab[114762]++
								r := NewReader(bytes.NewReader(data))
								r.Comma = tt.Comma
								r.Comment = tt.Comment
								r.LazyQuotes = tt.LazyQuotes
								r.TrimLeadingSpace = tt.TrimLeadingSpace

								records, err := r.ReadAll()
								if err != nil {
//line /usr/local/go/src/encoding/csv/fuzz.go:35
			_go_fuzz_dep_.CoverTab[114766]++
									continue
//line /usr/local/go/src/encoding/csv/fuzz.go:36
			// _ = "end of CoverTab[114766]"
		} else {
//line /usr/local/go/src/encoding/csv/fuzz.go:37
			_go_fuzz_dep_.CoverTab[114767]++
//line /usr/local/go/src/encoding/csv/fuzz.go:37
			// _ = "end of CoverTab[114767]"
//line /usr/local/go/src/encoding/csv/fuzz.go:37
		}
//line /usr/local/go/src/encoding/csv/fuzz.go:37
		// _ = "end of CoverTab[114762]"
//line /usr/local/go/src/encoding/csv/fuzz.go:37
		_go_fuzz_dep_.CoverTab[114763]++
								score = 1

								buf.Reset()
								w := NewWriter(buf)
								w.Comma = tt.Comma
								err = w.WriteAll(records)
								if err != nil {
//line /usr/local/go/src/encoding/csv/fuzz.go:44
			_go_fuzz_dep_.CoverTab[114768]++
									fmt.Printf("writer  = %#v\n", w)
									fmt.Printf("records = %v\n", records)
									panic(err)
//line /usr/local/go/src/encoding/csv/fuzz.go:47
			// _ = "end of CoverTab[114768]"
		} else {
//line /usr/local/go/src/encoding/csv/fuzz.go:48
			_go_fuzz_dep_.CoverTab[114769]++
//line /usr/local/go/src/encoding/csv/fuzz.go:48
			// _ = "end of CoverTab[114769]"
//line /usr/local/go/src/encoding/csv/fuzz.go:48
		}
//line /usr/local/go/src/encoding/csv/fuzz.go:48
		// _ = "end of CoverTab[114763]"
//line /usr/local/go/src/encoding/csv/fuzz.go:48
		_go_fuzz_dep_.CoverTab[114764]++

								r = NewReader(buf)
								r.Comma = tt.Comma
								r.Comment = tt.Comment
								r.LazyQuotes = tt.LazyQuotes
								r.TrimLeadingSpace = tt.TrimLeadingSpace
								result, err := r.ReadAll()
								if err != nil {
//line /usr/local/go/src/encoding/csv/fuzz.go:56
			_go_fuzz_dep_.CoverTab[114770]++
									fmt.Printf("reader  = %#v\n", r)
									fmt.Printf("records = %v\n", records)
									panic(err)
//line /usr/local/go/src/encoding/csv/fuzz.go:59
			// _ = "end of CoverTab[114770]"
		} else {
//line /usr/local/go/src/encoding/csv/fuzz.go:60
			_go_fuzz_dep_.CoverTab[114771]++
//line /usr/local/go/src/encoding/csv/fuzz.go:60
			// _ = "end of CoverTab[114771]"
//line /usr/local/go/src/encoding/csv/fuzz.go:60
		}
//line /usr/local/go/src/encoding/csv/fuzz.go:60
		// _ = "end of CoverTab[114764]"
//line /usr/local/go/src/encoding/csv/fuzz.go:60
		_go_fuzz_dep_.CoverTab[114765]++

								if !reflect.DeepEqual(records, result) {
//line /usr/local/go/src/encoding/csv/fuzz.go:62
			_go_fuzz_dep_.CoverTab[114772]++
									fmt.Println("records = \n", records)
									fmt.Println("result  = \n", records)
									panic("not equal")
//line /usr/local/go/src/encoding/csv/fuzz.go:65
			// _ = "end of CoverTab[114772]"
		} else {
//line /usr/local/go/src/encoding/csv/fuzz.go:66
			_go_fuzz_dep_.CoverTab[114773]++
//line /usr/local/go/src/encoding/csv/fuzz.go:66
			// _ = "end of CoverTab[114773]"
//line /usr/local/go/src/encoding/csv/fuzz.go:66
		}
//line /usr/local/go/src/encoding/csv/fuzz.go:66
		// _ = "end of CoverTab[114765]"
	}
//line /usr/local/go/src/encoding/csv/fuzz.go:67
	// _ = "end of CoverTab[114760]"
//line /usr/local/go/src/encoding/csv/fuzz.go:67
	_go_fuzz_dep_.CoverTab[114761]++

							return score
//line /usr/local/go/src/encoding/csv/fuzz.go:69
	// _ = "end of CoverTab[114761]"
}

//line /usr/local/go/src/encoding/csv/fuzz.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/csv/fuzz.go:70
var _ = _go_fuzz_dep_.CoverTab
