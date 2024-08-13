// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
package json

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:5
)

import (
	"bytes"
	"strconv"
)

// parseNumber reads the given []byte for a valid JSON number. If it is valid,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:12
// it returns the number of bytes.  Parsing logic follows the definition in
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:12
// https://tools.ietf.org/html/rfc7159#section-6, and is based off
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:12
// encoding/json.isValidNumber function.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:16
func parseNumber(input []byte) (int, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:16
	_go_fuzz_dep_.CoverTab[65585]++
															var n int

															s := input
															if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:20
		_go_fuzz_dep_.CoverTab[65592]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:21
		// _ = "end of CoverTab[65592]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:22
		_go_fuzz_dep_.CoverTab[65593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:22
		// _ = "end of CoverTab[65593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:22
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:22
	// _ = "end of CoverTab[65585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:22
	_go_fuzz_dep_.CoverTab[65586]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:25
	if s[0] == '-' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:25
		_go_fuzz_dep_.CoverTab[65594]++
																s = s[1:]
																n++
																if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:28
			_go_fuzz_dep_.CoverTab[65595]++
																	return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:29
			// _ = "end of CoverTab[65595]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:30
			_go_fuzz_dep_.CoverTab[65596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:30
			// _ = "end of CoverTab[65596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:30
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:30
		// _ = "end of CoverTab[65594]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:31
		_go_fuzz_dep_.CoverTab[65597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:31
		// _ = "end of CoverTab[65597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:31
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:31
	// _ = "end of CoverTab[65586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:31
	_go_fuzz_dep_.CoverTab[65587]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:34
	switch {
	case s[0] == '0':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:35
		_go_fuzz_dep_.CoverTab[65598]++
																s = s[1:]
																n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:37
		// _ = "end of CoverTab[65598]"

	case '1' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:39
		_go_fuzz_dep_.CoverTab[65601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:39
		return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:39
		// _ = "end of CoverTab[65601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:39
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:39
		_go_fuzz_dep_.CoverTab[65599]++
																s = s[1:]
																n++
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			_go_fuzz_dep_.CoverTab[65602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			// _ = "end of CoverTab[65602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			_go_fuzz_dep_.CoverTab[65603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			// _ = "end of CoverTab[65603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:42
			_go_fuzz_dep_.CoverTab[65604]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:44
			// _ = "end of CoverTab[65604]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:45
		// _ = "end of CoverTab[65599]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:47
		_go_fuzz_dep_.CoverTab[65600]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:48
		// _ = "end of CoverTab[65600]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:49
	// _ = "end of CoverTab[65587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:49
	_go_fuzz_dep_.CoverTab[65588]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		_go_fuzz_dep_.CoverTab[65605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		// _ = "end of CoverTab[65605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		_go_fuzz_dep_.CoverTab[65606]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		return '0' <= s[1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		// _ = "end of CoverTab[65606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		_go_fuzz_dep_.CoverTab[65607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		return s[1] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		// _ = "end of CoverTab[65607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:52
		_go_fuzz_dep_.CoverTab[65608]++
																s = s[2:]
																n += 2
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			_go_fuzz_dep_.CoverTab[65609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			// _ = "end of CoverTab[65609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			_go_fuzz_dep_.CoverTab[65610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			// _ = "end of CoverTab[65610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:55
			_go_fuzz_dep_.CoverTab[65611]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:57
			// _ = "end of CoverTab[65611]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:58
		// _ = "end of CoverTab[65608]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:59
		_go_fuzz_dep_.CoverTab[65612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:59
		// _ = "end of CoverTab[65612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:59
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:59
	// _ = "end of CoverTab[65588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:59
	_go_fuzz_dep_.CoverTab[65589]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
		_go_fuzz_dep_.CoverTab[65613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
		return (s[0] == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
			_go_fuzz_dep_.CoverTab[65614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
			return s[0] == 'E'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
			// _ = "end of CoverTab[65614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
		// _ = "end of CoverTab[65613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:63
		_go_fuzz_dep_.CoverTab[65615]++
																s = s[1:]
																n++
																if s[0] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:66
			_go_fuzz_dep_.CoverTab[65617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:66
			return s[0] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:66
			// _ = "end of CoverTab[65617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:66
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:66
			_go_fuzz_dep_.CoverTab[65618]++
																	s = s[1:]
																	n++
																	if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:69
				_go_fuzz_dep_.CoverTab[65619]++
																		return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:70
				// _ = "end of CoverTab[65619]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:71
				_go_fuzz_dep_.CoverTab[65620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:71
				// _ = "end of CoverTab[65620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:71
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:71
			// _ = "end of CoverTab[65618]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:72
			_go_fuzz_dep_.CoverTab[65621]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:72
			// _ = "end of CoverTab[65621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:72
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:72
		// _ = "end of CoverTab[65615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:72
		_go_fuzz_dep_.CoverTab[65616]++
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			_go_fuzz_dep_.CoverTab[65622]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			// _ = "end of CoverTab[65622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			_go_fuzz_dep_.CoverTab[65623]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			// _ = "end of CoverTab[65623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:73
			_go_fuzz_dep_.CoverTab[65624]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:75
			// _ = "end of CoverTab[65624]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:76
		// _ = "end of CoverTab[65616]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:77
		_go_fuzz_dep_.CoverTab[65625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:77
		// _ = "end of CoverTab[65625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:77
	// _ = "end of CoverTab[65589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:77
	_go_fuzz_dep_.CoverTab[65590]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
	if n < len(input) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
		_go_fuzz_dep_.CoverTab[65626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
		return isNotDelim(input[n])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
		// _ = "end of CoverTab[65626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:80
		_go_fuzz_dep_.CoverTab[65627]++
																return 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:81
		// _ = "end of CoverTab[65627]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:82
		_go_fuzz_dep_.CoverTab[65628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:82
		// _ = "end of CoverTab[65628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:82
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:82
	// _ = "end of CoverTab[65590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:82
	_go_fuzz_dep_.CoverTab[65591]++

															return n, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:84
	// _ = "end of CoverTab[65591]"
}

// numberParts is the result of parsing out a valid JSON number. It contains
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:87
// the parts of a number. The parts are used for integer conversion.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:89
type numberParts struct {
	neg	bool
	intp	[]byte
	frac	[]byte
	exp	[]byte
}

// parseNumber constructs numberParts from given []byte. The logic here is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:96
// similar to consumeNumber above with the difference of having to construct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:96
// numberParts. The slice fields in numberParts are subslices of the input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:99
func parseNumberParts(input []byte) (numberParts, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:99
	_go_fuzz_dep_.CoverTab[65629]++
															var neg bool
															var intp []byte
															var frac []byte
															var exp []byte

															s := input
															if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:106
		_go_fuzz_dep_.CoverTab[65635]++
																return numberParts{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:107
		// _ = "end of CoverTab[65635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:108
		_go_fuzz_dep_.CoverTab[65636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:108
		// _ = "end of CoverTab[65636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:108
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:108
	// _ = "end of CoverTab[65629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:108
	_go_fuzz_dep_.CoverTab[65630]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:111
	if s[0] == '-' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:111
		_go_fuzz_dep_.CoverTab[65637]++
																neg = true
																s = s[1:]
																if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:114
			_go_fuzz_dep_.CoverTab[65638]++
																	return numberParts{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:115
			// _ = "end of CoverTab[65638]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:116
			_go_fuzz_dep_.CoverTab[65639]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:116
			// _ = "end of CoverTab[65639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:116
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:116
		// _ = "end of CoverTab[65637]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:117
		_go_fuzz_dep_.CoverTab[65640]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:117
		// _ = "end of CoverTab[65640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:117
	// _ = "end of CoverTab[65630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:117
	_go_fuzz_dep_.CoverTab[65631]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:120
	switch {
	case s[0] == '0':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:121
		_go_fuzz_dep_.CoverTab[65641]++

																s = s[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:123
		// _ = "end of CoverTab[65641]"

	case '1' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:125
		_go_fuzz_dep_.CoverTab[65645]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:125
		return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:125
		// _ = "end of CoverTab[65645]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:125
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:125
		_go_fuzz_dep_.CoverTab[65642]++
																intp = s
																n := 1
																s = s[1:]
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			_go_fuzz_dep_.CoverTab[65646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			// _ = "end of CoverTab[65646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			_go_fuzz_dep_.CoverTab[65647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			// _ = "end of CoverTab[65647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:129
			_go_fuzz_dep_.CoverTab[65648]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:131
			// _ = "end of CoverTab[65648]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:132
		// _ = "end of CoverTab[65642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:132
		_go_fuzz_dep_.CoverTab[65643]++
																intp = intp[:n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:133
		// _ = "end of CoverTab[65643]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:135
		_go_fuzz_dep_.CoverTab[65644]++
																return numberParts{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:136
		// _ = "end of CoverTab[65644]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:137
	// _ = "end of CoverTab[65631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:137
	_go_fuzz_dep_.CoverTab[65632]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		_go_fuzz_dep_.CoverTab[65649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		// _ = "end of CoverTab[65649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		_go_fuzz_dep_.CoverTab[65650]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		return '0' <= s[1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		// _ = "end of CoverTab[65650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		_go_fuzz_dep_.CoverTab[65651]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		return s[1] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		// _ = "end of CoverTab[65651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:140
		_go_fuzz_dep_.CoverTab[65652]++
																frac = s[1:]
																n := 1
																s = s[2:]
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			_go_fuzz_dep_.CoverTab[65654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			// _ = "end of CoverTab[65654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			_go_fuzz_dep_.CoverTab[65655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			// _ = "end of CoverTab[65655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:144
			_go_fuzz_dep_.CoverTab[65656]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:146
			// _ = "end of CoverTab[65656]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:147
		// _ = "end of CoverTab[65652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:147
		_go_fuzz_dep_.CoverTab[65653]++
																frac = frac[:n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:148
		// _ = "end of CoverTab[65653]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:149
		_go_fuzz_dep_.CoverTab[65657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:149
		// _ = "end of CoverTab[65657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:149
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:149
	// _ = "end of CoverTab[65632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:149
	_go_fuzz_dep_.CoverTab[65633]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
		_go_fuzz_dep_.CoverTab[65658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
		return (s[0] == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
			_go_fuzz_dep_.CoverTab[65659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
			return s[0] == 'E'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
			// _ = "end of CoverTab[65659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
		// _ = "end of CoverTab[65658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:153
		_go_fuzz_dep_.CoverTab[65660]++
																s = s[1:]
																exp = s
																n := 0
																if s[0] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:157
			_go_fuzz_dep_.CoverTab[65663]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:157
			return s[0] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:157
			// _ = "end of CoverTab[65663]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:157
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:157
			_go_fuzz_dep_.CoverTab[65664]++
																	s = s[1:]
																	n++
																	if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:160
				_go_fuzz_dep_.CoverTab[65665]++
																		return numberParts{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:161
				// _ = "end of CoverTab[65665]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:162
				_go_fuzz_dep_.CoverTab[65666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:162
				// _ = "end of CoverTab[65666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:162
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:162
			// _ = "end of CoverTab[65664]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:163
			_go_fuzz_dep_.CoverTab[65667]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:163
			// _ = "end of CoverTab[65667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:163
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:163
		// _ = "end of CoverTab[65660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:163
		_go_fuzz_dep_.CoverTab[65661]++
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			_go_fuzz_dep_.CoverTab[65668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			// _ = "end of CoverTab[65668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			_go_fuzz_dep_.CoverTab[65669]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			// _ = "end of CoverTab[65669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:164
			_go_fuzz_dep_.CoverTab[65670]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:166
			// _ = "end of CoverTab[65670]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:167
		// _ = "end of CoverTab[65661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:167
		_go_fuzz_dep_.CoverTab[65662]++
																exp = exp[:n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:168
		// _ = "end of CoverTab[65662]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:169
		_go_fuzz_dep_.CoverTab[65671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:169
		// _ = "end of CoverTab[65671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:169
	// _ = "end of CoverTab[65633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:169
	_go_fuzz_dep_.CoverTab[65634]++

															return numberParts{
		neg:	neg,
		intp:	intp,
		frac:	bytes.TrimRight(frac, "0"),
		exp:	exp,
	}, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:176
	// _ = "end of CoverTab[65634]"
}

// normalizeToIntString returns an integer string in normal form without the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:179
// E-notation for given numberParts. It will return false if it is not an
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:179
// integer or if the exponent exceeds than max/min int value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:182
func normalizeToIntString(n numberParts) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:182
	_go_fuzz_dep_.CoverTab[65672]++
															intpSize := len(n.intp)
															fracSize := len(n.frac)

															if intpSize == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:186
		_go_fuzz_dep_.CoverTab[65677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:186
		return fracSize == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:186
		// _ = "end of CoverTab[65677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:186
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:186
		_go_fuzz_dep_.CoverTab[65678]++
																return "0", true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:187
		// _ = "end of CoverTab[65678]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:188
		_go_fuzz_dep_.CoverTab[65679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:188
		// _ = "end of CoverTab[65679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:188
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:188
	// _ = "end of CoverTab[65672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:188
	_go_fuzz_dep_.CoverTab[65673]++

															var exp int
															if len(n.exp) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:191
		_go_fuzz_dep_.CoverTab[65680]++
																i, err := strconv.ParseInt(string(n.exp), 10, 32)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:193
			_go_fuzz_dep_.CoverTab[65682]++
																	return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:194
			// _ = "end of CoverTab[65682]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:195
			_go_fuzz_dep_.CoverTab[65683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:195
			// _ = "end of CoverTab[65683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:195
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:195
		// _ = "end of CoverTab[65680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:195
		_go_fuzz_dep_.CoverTab[65681]++
																exp = int(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:196
		// _ = "end of CoverTab[65681]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:197
		_go_fuzz_dep_.CoverTab[65684]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:197
		// _ = "end of CoverTab[65684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:197
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:197
	// _ = "end of CoverTab[65673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:197
	_go_fuzz_dep_.CoverTab[65674]++

															var num []byte
															if exp >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:200
		_go_fuzz_dep_.CoverTab[65685]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:206
		if fracSize > exp {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:206
			_go_fuzz_dep_.CoverTab[65688]++
																	return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:207
			// _ = "end of CoverTab[65688]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:208
			_go_fuzz_dep_.CoverTab[65689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:208
			// _ = "end of CoverTab[65689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:208
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:208
		// _ = "end of CoverTab[65685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:208
		_go_fuzz_dep_.CoverTab[65686]++

		// Make sure resulting digits are within max value limit to avoid
		// unnecessarily constructing a large byte slice that may simply fail
		// later on.
		const maxDigits = 20	// Max uint64 value has 20 decimal digits.
		if intpSize+exp > maxDigits {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:214
			_go_fuzz_dep_.CoverTab[65690]++
																	return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:215
			// _ = "end of CoverTab[65690]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:216
			_go_fuzz_dep_.CoverTab[65691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:216
			// _ = "end of CoverTab[65691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:216
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:216
		// _ = "end of CoverTab[65686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:216
		_go_fuzz_dep_.CoverTab[65687]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:219
		num = n.intp[:len(n.intp):len(n.intp)]
		num = append(num, n.frac...)
		for i := 0; i < exp-fracSize; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:221
			_go_fuzz_dep_.CoverTab[65692]++
																	num = append(num, '0')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:222
			// _ = "end of CoverTab[65692]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:223
		// _ = "end of CoverTab[65687]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:224
		_go_fuzz_dep_.CoverTab[65693]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:228
		if fracSize > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:228
			_go_fuzz_dep_.CoverTab[65697]++
																	return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:229
			// _ = "end of CoverTab[65697]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:230
			_go_fuzz_dep_.CoverTab[65698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:230
			// _ = "end of CoverTab[65698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:230
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:230
		// _ = "end of CoverTab[65693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:230
		_go_fuzz_dep_.CoverTab[65694]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:234
		index := intpSize + exp
		if index < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:235
			_go_fuzz_dep_.CoverTab[65699]++
																	return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:236
			// _ = "end of CoverTab[65699]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:237
			_go_fuzz_dep_.CoverTab[65700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:237
			// _ = "end of CoverTab[65700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:237
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:237
		// _ = "end of CoverTab[65694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:237
		_go_fuzz_dep_.CoverTab[65695]++

																num = n.intp

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:242
		for i := index; i < intpSize; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:242
			_go_fuzz_dep_.CoverTab[65701]++
																	if num[i] != '0' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:243
				_go_fuzz_dep_.CoverTab[65702]++
																		return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:244
				// _ = "end of CoverTab[65702]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:245
				_go_fuzz_dep_.CoverTab[65703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:245
				// _ = "end of CoverTab[65703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:245
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:245
			// _ = "end of CoverTab[65701]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:246
		// _ = "end of CoverTab[65695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:246
		_go_fuzz_dep_.CoverTab[65696]++
																num = num[:index]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:247
		// _ = "end of CoverTab[65696]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:248
	// _ = "end of CoverTab[65674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:248
	_go_fuzz_dep_.CoverTab[65675]++

															if n.neg {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:250
		_go_fuzz_dep_.CoverTab[65704]++
																return "-" + string(num), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:251
		// _ = "end of CoverTab[65704]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:252
		_go_fuzz_dep_.CoverTab[65705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:252
		// _ = "end of CoverTab[65705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:252
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:252
	// _ = "end of CoverTab[65675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:252
	_go_fuzz_dep_.CoverTab[65676]++
															return string(num), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:253
	// _ = "end of CoverTab[65676]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:254
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_number.go:254
var _ = _go_fuzz_dep_.CoverTab
