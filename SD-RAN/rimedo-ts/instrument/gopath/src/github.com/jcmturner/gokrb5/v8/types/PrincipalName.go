//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:1
)

import (
	"strings"

	"github.com/jcmturner/gokrb5/v8/iana/nametype"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:12
// PrincipalName implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.2
type PrincipalName struct {
	NameType	int32		`asn1:"explicit,tag:0"`
	NameString	[]string	`asn1:"generalstring,explicit,tag:1"`
}

// NewPrincipalName creates a new PrincipalName from the name type int32 and name string provided.
func NewPrincipalName(ntype int32, spn string) PrincipalName {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:19
	_go_fuzz_dep_.CoverTab[86085]++
													return PrincipalName{
		NameType:	ntype,
		NameString:	strings.Split(spn, "/"),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:23
	// _ = "end of CoverTab[86085]"
}

// GetSalt returns a salt derived from the PrincipalName.
func (pn PrincipalName) GetSalt(realm string) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:27
	_go_fuzz_dep_.CoverTab[86086]++
													var sb []byte
													sb = append(sb, realm...)
													for _, n := range pn.NameString {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:30
		_go_fuzz_dep_.CoverTab[86088]++
														sb = append(sb, n...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:31
		// _ = "end of CoverTab[86088]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:32
	// _ = "end of CoverTab[86086]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:32
	_go_fuzz_dep_.CoverTab[86087]++
													return string(sb)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:33
	// _ = "end of CoverTab[86087]"
}

// Equal tests if the PrincipalName is equal to the one provided.
func (pn PrincipalName) Equal(n PrincipalName) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:37
	_go_fuzz_dep_.CoverTab[86089]++
													if len(pn.NameString) != len(n.NameString) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:38
		_go_fuzz_dep_.CoverTab[86092]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:39
		// _ = "end of CoverTab[86092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:40
		_go_fuzz_dep_.CoverTab[86093]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:40
		// _ = "end of CoverTab[86093]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:40
	// _ = "end of CoverTab[86089]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:40
	_go_fuzz_dep_.CoverTab[86090]++

													for i, s := range pn.NameString {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:42
		_go_fuzz_dep_.CoverTab[86094]++
														if n.NameString[i] != s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:43
			_go_fuzz_dep_.CoverTab[86095]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:44
			// _ = "end of CoverTab[86095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:45
			_go_fuzz_dep_.CoverTab[86096]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:45
			// _ = "end of CoverTab[86096]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:45
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:45
		// _ = "end of CoverTab[86094]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:46
	// _ = "end of CoverTab[86090]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:46
	_go_fuzz_dep_.CoverTab[86091]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:47
	// _ = "end of CoverTab[86091]"
}

// PrincipalNameString returns the PrincipalName in string form.
func (pn PrincipalName) PrincipalNameString() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:51
	_go_fuzz_dep_.CoverTab[86097]++
													return strings.Join(pn.NameString, "/")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:52
	// _ = "end of CoverTab[86097]"
}

// ParseSPNString will parse a string in the format <service>/<name>@<realm>
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:55
// a PrincipalName type will be returned with the name type set to KRB_NT_PRINCIPAL(1)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:55
// and the realm will be returned as a string. If the "@<realm>" suffix
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:55
// is not included in the SPN then the value of realm string returned will be ""
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:59
func ParseSPNString(spn string) (pn PrincipalName, realm string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:59
	_go_fuzz_dep_.CoverTab[86098]++
													if strings.Contains(spn, "@") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:60
		_go_fuzz_dep_.CoverTab[86100]++
														s := strings.Split(spn, "@")
														realm = s[len(s)-1]
														spn = strings.TrimSuffix(spn, "@"+realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:63
		// _ = "end of CoverTab[86100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:64
		_go_fuzz_dep_.CoverTab[86101]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:64
		// _ = "end of CoverTab[86101]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:64
	// _ = "end of CoverTab[86098]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:64
	_go_fuzz_dep_.CoverTab[86099]++
													pn = NewPrincipalName(nametype.KRB_NT_PRINCIPAL, spn)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:66
	// _ = "end of CoverTab[86099]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:67
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PrincipalName.go:67
var _ = _go_fuzz_dep_.CoverTab
