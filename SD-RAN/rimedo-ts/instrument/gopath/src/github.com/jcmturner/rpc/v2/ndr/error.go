//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:1
)

import "fmt"

// Malformed implements the error interface for malformed NDR encoding errors.
type Malformed struct {
	EText string
}

// Error implements the error interface on the Malformed struct.
func (e Malformed) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:11
	_go_fuzz_dep_.CoverTab[87127]++
											return fmt.Sprintf("malformed NDR stream: %s", e.EText)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:12
	// _ = "end of CoverTab[87127]"
}

// Errorf formats an error message into a malformed NDR error.
func Errorf(format string, a ...interface{}) Malformed {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:16
	_go_fuzz_dep_.CoverTab[87128]++
											return Malformed{EText: fmt.Sprintf(format, a...)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:17
	// _ = "end of CoverTab[87128]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/error.go:18
var _ = _go_fuzz_dep_.CoverTab
