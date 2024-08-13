//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
package hcl

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:1
)

import (
	"fmt"

	"github.com/hashicorp/hcl/hcl/ast"
	hclParser "github.com/hashicorp/hcl/hcl/parser"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

// ParseBytes accepts as input byte slice and returns ast tree.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:11
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:11
// Input can be either JSON or HCL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:14
func ParseBytes(in []byte) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:14
	_go_fuzz_dep_.CoverTab[122147]++
										return parse(in)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:15
	// _ = "end of CoverTab[122147]"
}

// ParseString accepts input as a string and returns ast tree.
func ParseString(input string) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:19
	_go_fuzz_dep_.CoverTab[122148]++
										return parse([]byte(input))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:20
	// _ = "end of CoverTab[122148]"
}

func parse(in []byte) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:23
	_go_fuzz_dep_.CoverTab[122149]++
										switch lexMode(in) {
	case lexModeHcl:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:25
		_go_fuzz_dep_.CoverTab[122151]++
											return hclParser.Parse(in)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:26
		// _ = "end of CoverTab[122151]"
	case lexModeJson:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:27
		_go_fuzz_dep_.CoverTab[122152]++
											return jsonParser.Parse(in)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:28
		// _ = "end of CoverTab[122152]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:28
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:28
		_go_fuzz_dep_.CoverTab[122153]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:28
		// _ = "end of CoverTab[122153]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:29
	// _ = "end of CoverTab[122149]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:29
	_go_fuzz_dep_.CoverTab[122150]++

										return nil, fmt.Errorf("unknown config format")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:31
	// _ = "end of CoverTab[122150]"
}

// Parse parses the given input and returns the root object.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:34
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:34
// The input format can be either HCL or JSON.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:37
func Parse(input string) (*ast.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:37
	_go_fuzz_dep_.CoverTab[122154]++
										return parse([]byte(input))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:38
	// _ = "end of CoverTab[122154]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:39
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/parse.go:39
var _ = _go_fuzz_dep_.CoverTab
