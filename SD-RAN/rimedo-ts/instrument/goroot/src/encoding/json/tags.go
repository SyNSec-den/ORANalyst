// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/tags.go:5
package json

//line /usr/local/go/src/encoding/json/tags.go:5
import (
//line /usr/local/go/src/encoding/json/tags.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/tags.go:5
)
//line /usr/local/go/src/encoding/json/tags.go:5
import (
//line /usr/local/go/src/encoding/json/tags.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/tags.go:5
)

import (
	"strings"
)

// tagOptions is the string following a comma in a struct field's "json"
//line /usr/local/go/src/encoding/json/tags.go:11
// tag, or the empty string. It does not include the leading comma.
//line /usr/local/go/src/encoding/json/tags.go:13
type tagOptions string

// parseTag splits a struct field's json tag into its name and
//line /usr/local/go/src/encoding/json/tags.go:15
// comma-separated options.
//line /usr/local/go/src/encoding/json/tags.go:17
func parseTag(tag string) (string, tagOptions) {
//line /usr/local/go/src/encoding/json/tags.go:17
	_go_fuzz_dep_.CoverTab[28558]++
							tag, opt, _ := strings.Cut(tag, ",")
							return tag, tagOptions(opt)
//line /usr/local/go/src/encoding/json/tags.go:19
	// _ = "end of CoverTab[28558]"
}

// Contains reports whether a comma-separated list of options
//line /usr/local/go/src/encoding/json/tags.go:22
// contains a particular substr flag. substr must be surrounded by a
//line /usr/local/go/src/encoding/json/tags.go:22
// string boundary or commas.
//line /usr/local/go/src/encoding/json/tags.go:25
func (o tagOptions) Contains(optionName string) bool {
//line /usr/local/go/src/encoding/json/tags.go:25
	_go_fuzz_dep_.CoverTab[28559]++
							if len(o) == 0 {
//line /usr/local/go/src/encoding/json/tags.go:26
		_go_fuzz_dep_.CoverTab[28562]++
								return false
//line /usr/local/go/src/encoding/json/tags.go:27
		// _ = "end of CoverTab[28562]"
	} else {
//line /usr/local/go/src/encoding/json/tags.go:28
		_go_fuzz_dep_.CoverTab[28563]++
//line /usr/local/go/src/encoding/json/tags.go:28
		// _ = "end of CoverTab[28563]"
//line /usr/local/go/src/encoding/json/tags.go:28
	}
//line /usr/local/go/src/encoding/json/tags.go:28
	// _ = "end of CoverTab[28559]"
//line /usr/local/go/src/encoding/json/tags.go:28
	_go_fuzz_dep_.CoverTab[28560]++
							s := string(o)
							for s != "" {
//line /usr/local/go/src/encoding/json/tags.go:30
		_go_fuzz_dep_.CoverTab[28564]++
								var name string
								name, s, _ = strings.Cut(s, ",")
								if name == optionName {
//line /usr/local/go/src/encoding/json/tags.go:33
			_go_fuzz_dep_.CoverTab[28565]++
									return true
//line /usr/local/go/src/encoding/json/tags.go:34
			// _ = "end of CoverTab[28565]"
		} else {
//line /usr/local/go/src/encoding/json/tags.go:35
			_go_fuzz_dep_.CoverTab[28566]++
//line /usr/local/go/src/encoding/json/tags.go:35
			// _ = "end of CoverTab[28566]"
//line /usr/local/go/src/encoding/json/tags.go:35
		}
//line /usr/local/go/src/encoding/json/tags.go:35
		// _ = "end of CoverTab[28564]"
	}
//line /usr/local/go/src/encoding/json/tags.go:36
	// _ = "end of CoverTab[28560]"
//line /usr/local/go/src/encoding/json/tags.go:36
	_go_fuzz_dep_.CoverTab[28561]++
							return false
//line /usr/local/go/src/encoding/json/tags.go:37
	// _ = "end of CoverTab[28561]"
}

//line /usr/local/go/src/encoding/json/tags.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/tags.go:38
var _ = _go_fuzz_dep_.CoverTab
