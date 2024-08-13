// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// Package text implements the text format for protocol buffers.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// This package has no semantic understanding for protocol buffers and is only
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// a parser and composer for the format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// There is no formal specification for the protobuf text format, as such the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// C++ implementation (see google::protobuf::TextFormat) is the reference
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// implementation of the text format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// This package is neither a superset nor a subset of the C++ implementation.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// This implementation permits a more liberal grammar in some cases to be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// backwards compatible with the historical Go implementation.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// Future parsings unique to Go should not be added.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// Some grammars allowed by the C++ implementation are deliberately
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// not implemented here because they are considered a bug by the protobuf team
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// and should not be replicated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// The Go implementation should implement a sufficient amount of the C++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// grammar such that the default text serialization by C++ can be parsed by Go.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// However, just because the C++ parser accepts some input does not mean that
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// the Go implementation should as well.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
// The text format is almost a superset of JSON except:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//   - message keys are not quoted strings, but identifiers
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:5
//   - the top-level value must be a message without the delimiters
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/doc.go:29
var _ = _go_fuzz_dep_.CoverTab
