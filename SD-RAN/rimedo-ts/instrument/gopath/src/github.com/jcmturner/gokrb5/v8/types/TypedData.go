//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:1
)

import "github.com/jcmturner/gofork/encoding/asn1"

// TypedData implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.9.1
type TypedData struct {
	DataType	int32	`asn1:"explicit,tag:0"`
	DataValue	[]byte	`asn1:"optional,explicit,tag:1"`
}

// TypedDataSequence implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.9.1
type TypedDataSequence []TypedData

// Unmarshal bytes into the TypedDataSequence.
func (a *TypedDataSequence) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:15
	_go_fuzz_dep_.CoverTab[86102]++
												_, err := asn1.Unmarshal(b, a)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:17
	// _ = "end of CoverTab[86102]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/TypedData.go:18
var _ = _go_fuzz_dep_.CoverTab
