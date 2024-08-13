//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// Package cachecontrol implements the logic for HTTP Caching
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// Deciding if an HTTP Response can be cached is often harder
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// and more bug prone than an actual cache storage backend.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// cachecontrol provides a simple interface to determine if
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// request and response pairs are cachable as defined under
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:18
// RFC 7234 http://tools.ietf.org/html/rfc7234
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
package cachecontrol

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
)

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/doc.go:25
var _ = _go_fuzz_dep_.CoverTab
