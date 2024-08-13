//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// Package hcl decodes HCL into usable Go structures.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// hcl input can come in either pure HCL format or JSON format.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// It can be parsed into an AST, and then decoded into a structure,
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// or it can be decoded directly from a string into a structure.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// If you choose to parse HCL into a raw AST, the benefit is that you
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// can write custom visitor implementations to implement custom
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// semantic checks. By default, HCL does not perform any semantic
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:1
// checks.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
package hcl

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl.go:11
var _ = _go_fuzz_dep_.CoverTab
