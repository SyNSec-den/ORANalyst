//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:6
import (
	"github.com/jcmturner/gofork/encoding/asn1"
)

// NewKrbFlags returns an ASN1 BitString struct of the right size for KrbFlags.
func NewKrbFlags() asn1.BitString {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:11
	_go_fuzz_dep_.CoverTab[86044]++
													f := asn1.BitString{}
													f.Bytes = make([]byte, 4)
													f.BitLength = len(f.Bytes) * 8
													return f
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:15
	// _ = "end of CoverTab[86044]"
}

// SetFlags sets the flags of an ASN1 BitString.
func SetFlags(f *asn1.BitString, j []int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:19
	_go_fuzz_dep_.CoverTab[86045]++
													for _, i := range j {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:20
		_go_fuzz_dep_.CoverTab[86046]++
														SetFlag(f, i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:21
		// _ = "end of CoverTab[86046]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:22
	// _ = "end of CoverTab[86045]"
}

// SetFlag sets a flag in an ASN1 BitString.
func SetFlag(f *asn1.BitString, i int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:26
	_go_fuzz_dep_.CoverTab[86047]++
													for l := len(f.Bytes); l < 4; l++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:27
		_go_fuzz_dep_.CoverTab[86049]++
														(*f).Bytes = append((*f).Bytes, byte(0))
														(*f).BitLength = len((*f).Bytes) * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:29
		// _ = "end of CoverTab[86049]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:30
	// _ = "end of CoverTab[86047]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:30
	_go_fuzz_dep_.CoverTab[86048]++

													b := i / 8

													p := uint(7 - (i - 8*b))
													(*f).Bytes[b] = (*f).Bytes[b] | (1 << p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:35
	// _ = "end of CoverTab[86048]"
}

// UnsetFlags unsets flags in an ASN1 BitString.
func UnsetFlags(f *asn1.BitString, j []int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:39
	_go_fuzz_dep_.CoverTab[86050]++
													for _, i := range j {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:40
		_go_fuzz_dep_.CoverTab[86051]++
														UnsetFlag(f, i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:41
		// _ = "end of CoverTab[86051]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:42
	// _ = "end of CoverTab[86050]"
}

// UnsetFlag unsets a flag in an ASN1 BitString.
func UnsetFlag(f *asn1.BitString, i int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:46
	_go_fuzz_dep_.CoverTab[86052]++
													for l := len(f.Bytes); l < 4; l++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:47
		_go_fuzz_dep_.CoverTab[86054]++
														(*f).Bytes = append((*f).Bytes, byte(0))
														(*f).BitLength = len((*f).Bytes) * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:49
		// _ = "end of CoverTab[86054]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:50
	// _ = "end of CoverTab[86052]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:50
	_go_fuzz_dep_.CoverTab[86053]++

													b := i / 8

													p := uint(7 - (i - 8*b))
													(*f).Bytes[b] = (*f).Bytes[b] &^ (1 << p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:55
	// _ = "end of CoverTab[86053]"
}

// IsFlagSet tests if a flag is set in the ASN1 BitString.
func IsFlagSet(f *asn1.BitString, i int) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:59
	_go_fuzz_dep_.CoverTab[86055]++

													b := i / 8

													p := uint(7 - (i - 8*b))
													if (*f).Bytes[b]&(1<<p) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:64
		_go_fuzz_dep_.CoverTab[86057]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:65
		// _ = "end of CoverTab[86057]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:66
		_go_fuzz_dep_.CoverTab[86058]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:66
		// _ = "end of CoverTab[86058]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:66
	// _ = "end of CoverTab[86055]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:66
	_go_fuzz_dep_.CoverTab[86056]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:67
	// _ = "end of CoverTab[86056]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:68
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/KerberosFlags.go:68
var _ = _go_fuzz_dep_.CoverTab
