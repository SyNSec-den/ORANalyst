//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
package token

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:1
)

import "fmt"

// Pos describes an arbitrary source position
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:5
// including the file, line, and column location.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:5
// A Position is valid if the line number is > 0.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:8
type Pos struct {
	Filename	string	// filename, if any
	Offset		int	// offset, starting at 0
	Line		int	// line number, starting at 1
	Column		int	// column number, starting at 1 (character count)
}

// IsValid returns true if the position is valid.
func (p *Pos) IsValid() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:16
	_go_fuzz_dep_.CoverTab[121547]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:16
	return p.Line > 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:16
	// _ = "end of CoverTab[121547]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:16
}

// String returns a string in one of several forms:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:18
//	file:line:column    valid position with file name
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:18
//	line:column         valid position without file name
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:18
//	file                invalid position with file name
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:18
//	-                   invalid position without file name
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:24
func (p Pos) String() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:24
	_go_fuzz_dep_.CoverTab[121548]++
												s := p.Filename
												if p.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:26
		_go_fuzz_dep_.CoverTab[121551]++
													if s != "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:27
			_go_fuzz_dep_.CoverTab[121553]++
														s += ":"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:28
			// _ = "end of CoverTab[121553]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:29
			_go_fuzz_dep_.CoverTab[121554]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:29
			// _ = "end of CoverTab[121554]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:29
		// _ = "end of CoverTab[121551]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:29
		_go_fuzz_dep_.CoverTab[121552]++
													s += fmt.Sprintf("%d:%d", p.Line, p.Column)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:30
		// _ = "end of CoverTab[121552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:31
		_go_fuzz_dep_.CoverTab[121555]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:31
		// _ = "end of CoverTab[121555]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:31
	// _ = "end of CoverTab[121548]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:31
	_go_fuzz_dep_.CoverTab[121549]++
												if s == "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:32
		_go_fuzz_dep_.CoverTab[121556]++
													s = "-"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:33
		// _ = "end of CoverTab[121556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:34
		_go_fuzz_dep_.CoverTab[121557]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:34
		// _ = "end of CoverTab[121557]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:34
	// _ = "end of CoverTab[121549]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:34
	_go_fuzz_dep_.CoverTab[121550]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:35
	// _ = "end of CoverTab[121550]"
}

// Before reports whether the position p is before u.
func (p Pos) Before(u Pos) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:39
	_go_fuzz_dep_.CoverTab[121558]++
												return u.Offset > p.Offset || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:40
		_go_fuzz_dep_.CoverTab[121559]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:40
		return u.Line > p.Line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:40
		// _ = "end of CoverTab[121559]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:40
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:40
	// _ = "end of CoverTab[121558]"
}

// After reports whether the position p is after u.
func (p Pos) After(u Pos) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:44
	_go_fuzz_dep_.CoverTab[121560]++
												return u.Offset < p.Offset || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:45
		_go_fuzz_dep_.CoverTab[121561]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:45
		return u.Line < p.Line
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:45
		// _ = "end of CoverTab[121561]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:45
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:45
	// _ = "end of CoverTab[121560]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/token/position.go:46
var _ = _go_fuzz_dep_.CoverTab
