//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
package multipart

//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
import (
//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
)
//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
import (
//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/multipart/readmimeheader.go:4
)

import (
	"net/textproto"
	_ "unsafe"
)

//line /usr/local/go/src/mime/multipart/readmimeheader.go:13
//go:linkname readMIMEHeader net/textproto.readMIMEHeader
func readMIMEHeader(r *textproto.Reader, maxMemory, maxHeaders int64) (textproto.MIMEHeader, error)

//line /usr/local/go/src/mime/multipart/readmimeheader.go:14
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/multipart/readmimeheader.go:14
var _ = _go_fuzz_dep_.CoverTab
