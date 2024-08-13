// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the code to handle template options.

//line /usr/local/go/src/text/template/option.go:7
package template

//line /usr/local/go/src/text/template/option.go:7
import (
//line /usr/local/go/src/text/template/option.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/option.go:7
)
//line /usr/local/go/src/text/template/option.go:7
import (
//line /usr/local/go/src/text/template/option.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/option.go:7
)

import "strings"

// missingKeyAction defines how to respond to indexing a map with a key that is not present.
type missingKeyAction int

const (
	mapInvalid	missingKeyAction	= iota	// Return an invalid reflect.Value.
	mapZeroValue					// Return the zero value for the map element.
	mapError					// Error out
)

type option struct {
	missingKey missingKeyAction
}

// Option sets options for the template. Options are described by
//line /usr/local/go/src/text/template/option.go:24
// strings, either a simple string or "key=value". There can be at
//line /usr/local/go/src/text/template/option.go:24
// most one equals sign in an option string. If the option string
//line /usr/local/go/src/text/template/option.go:24
// is unrecognized or otherwise invalid, Option panics.
//line /usr/local/go/src/text/template/option.go:24
//
//line /usr/local/go/src/text/template/option.go:24
// Known options:
//line /usr/local/go/src/text/template/option.go:24
//
//line /usr/local/go/src/text/template/option.go:24
// missingkey: Control the behavior during execution if a map is
//line /usr/local/go/src/text/template/option.go:24
// indexed with a key that is not present in the map.
//line /usr/local/go/src/text/template/option.go:24
//
//line /usr/local/go/src/text/template/option.go:24
//	"missingkey=default" or "missingkey=invalid"
//line /usr/local/go/src/text/template/option.go:24
//		The default behavior: Do nothing and continue execution.
//line /usr/local/go/src/text/template/option.go:24
//		If printed, the result of the index operation is the string
//line /usr/local/go/src/text/template/option.go:24
//		"<no value>".
//line /usr/local/go/src/text/template/option.go:24
//	"missingkey=zero"
//line /usr/local/go/src/text/template/option.go:24
//		The operation returns the zero value for the map type's element.
//line /usr/local/go/src/text/template/option.go:24
//	"missingkey=error"
//line /usr/local/go/src/text/template/option.go:24
//		Execution stops immediately with an error.
//line /usr/local/go/src/text/template/option.go:42
func (t *Template) Option(opt ...string) *Template {
//line /usr/local/go/src/text/template/option.go:42
	_go_fuzz_dep_.CoverTab[30564]++
							t.init()
							for _, s := range opt {
//line /usr/local/go/src/text/template/option.go:44
		_go_fuzz_dep_.CoverTab[30566]++
								t.setOption(s)
//line /usr/local/go/src/text/template/option.go:45
		// _ = "end of CoverTab[30566]"
	}
//line /usr/local/go/src/text/template/option.go:46
	// _ = "end of CoverTab[30564]"
//line /usr/local/go/src/text/template/option.go:46
	_go_fuzz_dep_.CoverTab[30565]++
							return t
//line /usr/local/go/src/text/template/option.go:47
	// _ = "end of CoverTab[30565]"
}

func (t *Template) setOption(opt string) {
//line /usr/local/go/src/text/template/option.go:50
	_go_fuzz_dep_.CoverTab[30567]++
							if opt == "" {
//line /usr/local/go/src/text/template/option.go:51
		_go_fuzz_dep_.CoverTab[30570]++
								panic("empty option string")
//line /usr/local/go/src/text/template/option.go:52
		// _ = "end of CoverTab[30570]"
	} else {
//line /usr/local/go/src/text/template/option.go:53
		_go_fuzz_dep_.CoverTab[30571]++
//line /usr/local/go/src/text/template/option.go:53
		// _ = "end of CoverTab[30571]"
//line /usr/local/go/src/text/template/option.go:53
	}
//line /usr/local/go/src/text/template/option.go:53
	// _ = "end of CoverTab[30567]"
//line /usr/local/go/src/text/template/option.go:53
	_go_fuzz_dep_.CoverTab[30568]++

							if key, value, ok := strings.Cut(opt, "="); ok {
//line /usr/local/go/src/text/template/option.go:55
		_go_fuzz_dep_.CoverTab[30572]++
								switch key {
		case "missingkey":
//line /usr/local/go/src/text/template/option.go:57
			_go_fuzz_dep_.CoverTab[30573]++
									switch value {
			case "invalid", "default":
//line /usr/local/go/src/text/template/option.go:59
				_go_fuzz_dep_.CoverTab[30575]++
										t.option.missingKey = mapInvalid
										return
//line /usr/local/go/src/text/template/option.go:61
				// _ = "end of CoverTab[30575]"
			case "zero":
//line /usr/local/go/src/text/template/option.go:62
				_go_fuzz_dep_.CoverTab[30576]++
										t.option.missingKey = mapZeroValue
										return
//line /usr/local/go/src/text/template/option.go:64
				// _ = "end of CoverTab[30576]"
			case "error":
//line /usr/local/go/src/text/template/option.go:65
				_go_fuzz_dep_.CoverTab[30577]++
										t.option.missingKey = mapError
										return
//line /usr/local/go/src/text/template/option.go:67
				// _ = "end of CoverTab[30577]"
//line /usr/local/go/src/text/template/option.go:67
			default:
//line /usr/local/go/src/text/template/option.go:67
				_go_fuzz_dep_.CoverTab[30578]++
//line /usr/local/go/src/text/template/option.go:67
				// _ = "end of CoverTab[30578]"
			}
//line /usr/local/go/src/text/template/option.go:68
			// _ = "end of CoverTab[30573]"
//line /usr/local/go/src/text/template/option.go:68
		default:
//line /usr/local/go/src/text/template/option.go:68
			_go_fuzz_dep_.CoverTab[30574]++
//line /usr/local/go/src/text/template/option.go:68
			// _ = "end of CoverTab[30574]"
		}
//line /usr/local/go/src/text/template/option.go:69
		// _ = "end of CoverTab[30572]"
	} else {
//line /usr/local/go/src/text/template/option.go:70
		_go_fuzz_dep_.CoverTab[30579]++
//line /usr/local/go/src/text/template/option.go:70
		// _ = "end of CoverTab[30579]"
//line /usr/local/go/src/text/template/option.go:70
	}
//line /usr/local/go/src/text/template/option.go:70
	// _ = "end of CoverTab[30568]"
//line /usr/local/go/src/text/template/option.go:70
	_go_fuzz_dep_.CoverTab[30569]++
							panic("unrecognized option: " + opt)
//line /usr/local/go/src/text/template/option.go:71
	// _ = "end of CoverTab[30569]"
}

//line /usr/local/go/src/text/template/option.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/option.go:72
var _ = _go_fuzz_dep_.CoverTab
