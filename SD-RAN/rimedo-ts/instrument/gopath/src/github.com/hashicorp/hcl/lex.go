//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
package hcl

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:1
)

import (
	"unicode"
	"unicode/utf8"
)

type lexModeValue byte

const (
	lexModeUnknown	lexModeValue	= iota
	lexModeHcl
	lexModeJson
)

// lexMode returns whether we're going to be parsing in JSON
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:16
// mode or HCL mode.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:18
func lexMode(v []byte) lexModeValue {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:18
	_go_fuzz_dep_.CoverTab[122138]++
										var (
		r	rune
		w	int
		offset	int
	)

	for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:25
		_go_fuzz_dep_.CoverTab[122140]++
											r, w = utf8.DecodeRune(v[offset:])
											offset += w
											if unicode.IsSpace(r) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:28
			_go_fuzz_dep_.CoverTab[122143]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:29
			// _ = "end of CoverTab[122143]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:30
			_go_fuzz_dep_.CoverTab[122144]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:30
			// _ = "end of CoverTab[122144]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:30
		// _ = "end of CoverTab[122140]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:30
		_go_fuzz_dep_.CoverTab[122141]++
											if r == '{' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:31
			_go_fuzz_dep_.CoverTab[122145]++
												return lexModeJson
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:32
			// _ = "end of CoverTab[122145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:33
			_go_fuzz_dep_.CoverTab[122146]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:33
			// _ = "end of CoverTab[122146]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:33
		// _ = "end of CoverTab[122141]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:33
		_go_fuzz_dep_.CoverTab[122142]++
											break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:34
		// _ = "end of CoverTab[122142]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:35
	// _ = "end of CoverTab[122138]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:35
	_go_fuzz_dep_.CoverTab[122139]++

										return lexModeHcl
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:37
	// _ = "end of CoverTab[122139]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/lex.go:38
var _ = _go_fuzz_dep_.CoverTab
