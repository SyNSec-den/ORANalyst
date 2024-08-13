//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
package rfc4757

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:1
)

import "encoding/binary"

// UsageToMSMsgType converts Kerberos key usage numbers to Microsoft message type encoded as a little-endian four byte slice.
func UsageToMSMsgType(usage uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:6
	_go_fuzz_dep_.CoverTab[85837]++

													switch usage {
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:9
		_go_fuzz_dep_.CoverTab[85839]++
														usage = 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:10
		// _ = "end of CoverTab[85839]"
	case 9:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:11
		_go_fuzz_dep_.CoverTab[85840]++
														usage = 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:12
		// _ = "end of CoverTab[85840]"
	case 23:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:13
		_go_fuzz_dep_.CoverTab[85841]++
														usage = 13
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:14
		// _ = "end of CoverTab[85841]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:14
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:14
		_go_fuzz_dep_.CoverTab[85842]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:14
		// _ = "end of CoverTab[85842]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:15
	// _ = "end of CoverTab[85837]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:15
	_go_fuzz_dep_.CoverTab[85838]++

													tb := make([]byte, 4)
													binary.PutUvarint(tb, uint64(usage))
													return tb
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:19
	// _ = "end of CoverTab[85838]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/msgtype.go:20
var _ = _go_fuzz_dep_.CoverTab
