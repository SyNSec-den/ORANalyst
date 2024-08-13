//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
package parser

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:1
)

import (
	"fmt"

	"github.com/hashicorp/hcl/hcl/token"
)

// PosError is a parse error that contains a position.
type PosError struct {
	Pos	token.Pos
	Err	error
}

func (e *PosError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:15
	_go_fuzz_dep_.CoverTab[121380]++
												return fmt.Sprintf("At %s: %s", e.Pos, e.Err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:16
	// _ = "end of CoverTab[121380]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/parser/error.go:17
var _ = _go_fuzz_dep_.CoverTab
