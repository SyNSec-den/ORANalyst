//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:1
// Package addrtype provides Address type assigned numbers.
package addrtype

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:2
)

// Address type IDs.
const (
	IPv4		int32	= 2
	Directional	int32	= 3
	ChaosNet	int32	= 5
	XNS		int32	= 6
	ISO		int32	= 7
	DECNETPhaseIV	int32	= 12
	AppleTalkDDP	int32	= 16
	NetBios		int32	= 20
	IPv6		int32	= 24
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/addrtype/constants.go:15
var _ = _go_fuzz_dep_.CoverTab
