// +build !appengine
// +build gc
// +build !noasm

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:5
)

//go:noescape
func decodeBlock(dst, src []byte) int

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:8
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/decode_amd64.go:8
var _ = _go_fuzz_dep_.CoverTab
