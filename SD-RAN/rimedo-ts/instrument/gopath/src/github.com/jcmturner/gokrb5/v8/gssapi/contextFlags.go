//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
package gssapi

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:1
)

import "github.com/jcmturner/gofork/encoding/asn1"

// GSS-API context flags assigned numbers.
const (
	ContextFlagDeleg	= 1
	ContextFlagMutual	= 2
	ContextFlagReplay	= 4
	ContextFlagSequence	= 8
	ContextFlagConf		= 16
	ContextFlagInteg	= 32
	ContextFlagAnon		= 64
)

// ContextFlags flags for GSSAPI
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:16
// DEPRECATED - do not use
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:18
type ContextFlags asn1.BitString

// NewContextFlags creates a new ContextFlags instance
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:20
// DEPRECATED - do not use
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:22
func NewContextFlags() ContextFlags {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:22
	_go_fuzz_dep_.CoverTab[88830]++
													var c ContextFlags
													c.BitLength = 32
													c.Bytes = make([]byte, 4)
													return c
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:26
	// _ = "end of CoverTab[88830]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/contextFlags.go:27
var _ = _go_fuzz_dep_.CoverTab
