//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:1
)

import (
	"github.com/jcmturner/gofork/encoding/asn1"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:10
// AuthorizationData implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6
type AuthorizationData []AuthorizationDataEntry

// AuthorizationDataEntry implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6
type AuthorizationDataEntry struct {
	ADType	int32	`asn1:"explicit,tag:0"`
	ADData	[]byte	`asn1:"explicit,tag:1"`
}

// ADIfRelevant implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6.1
type ADIfRelevant AuthorizationData

// ADKDCIssued implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6.2
type ADKDCIssued struct {
	ADChecksum	Checksum		`asn1:"explicit,tag:0"`
	IRealm		string			`asn1:"optional,generalstring,explicit,tag:1"`
	Isname		PrincipalName		`asn1:"optional,explicit,tag:2"`
	Elements	AuthorizationData	`asn1:"explicit,tag:3"`
}

// ADAndOr implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6.3
type ADAndOr struct {
	ConditionCount	int32			`asn1:"explicit,tag:0"`
	Elements	AuthorizationData	`asn1:"explicit,tag:1"`
}

// ADMandatoryForKDC implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.6.4
type ADMandatoryForKDC AuthorizationData

// Unmarshal bytes into the ADKDCIssued.
func (a *ADKDCIssued) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:40
	_go_fuzz_dep_.CoverTab[85956]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:42
	// _ = "end of CoverTab[85956]"
}

// Unmarshal bytes into the AuthorizationData.
func (a *AuthorizationData) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:46
	_go_fuzz_dep_.CoverTab[85957]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:48
	// _ = "end of CoverTab[85957]"
}

// Unmarshal bytes into the AuthorizationDataEntry.
func (a *AuthorizationDataEntry) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:52
	_go_fuzz_dep_.CoverTab[85958]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:54
	// _ = "end of CoverTab[85958]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:55
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/AuthorizationData.go:55
var _ = _go_fuzz_dep_.CoverTab
