//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:1
// Package mstypes provides implemnations of some Microsoft data types [MS-DTYP] https://msdn.microsoft.com/en-us/library/cc230283.aspx
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:2
)

// LPWSTR implements https://msdn.microsoft.com/en-us/library/cc230355.aspx
type LPWSTR struct {
	Value string `ndr:"pointer,conformant,varying"`
}

// String returns the string representation of LPWSTR data type.
func (s *LPWSTR) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:10
	_go_fuzz_dep_.CoverTab[87354]++
												return s.Value
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:11
	// _ = "end of CoverTab[87354]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:12
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/common.go:12
var _ = _go_fuzz_dep_.CoverTab
