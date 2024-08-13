//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// Package toml is a TOML parser and manipulation library.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// This version supports the specification as described in
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// https://github.com/toml-lang/toml/blob/master/versions/en/toml-v0.5.0.md
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// # Marshaling
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// Go-toml can marshal and unmarshal TOML documents from and to data
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// structures.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// # TOML document as a tree
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// Go-toml can operate on a TOML document as a tree. Use one of the Load*
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// functions to parse TOML data and obtain a Tree instance, then one of its
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// methods to manipulate the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// # JSONPath-like queries
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// The package github.com/pelletier/go-toml/query implements a system
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// similar to JSONPath to quickly retrieve elements of a TOML document using a
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:1
// single expression. See the package documentation for more information.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
)

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/doc.go:23
var _ = _go_fuzz_dep_.CoverTab
