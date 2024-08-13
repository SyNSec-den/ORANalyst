//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:1
// Package asn1tools provides tools for managing ASN1 marshaled data.
package asn1tools

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:2
)

import (
	"github.com/jcmturner/gofork/encoding/asn1"
)

// MarshalLengthBytes returns the ASN1 encoded bytes for the length 'l'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
// There are two forms: short (for lengths between 0 and 127), and long definite (for lengths between 0 and 2^1008 -1).
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
// Short form: One octet. Bit 8 has value "0" and bits 7-1 give the length.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:8
// Long form: Two to 127 octets. Bit 8 of first octet has value "1" and bits 7-1 give the number of additional length octets. Second and following octets give the length, base 256, most significant digit first.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:15
func MarshalLengthBytes(l int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:15
	_go_fuzz_dep_.CoverTab[83092]++
												if l <= 127 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:16
		_go_fuzz_dep_.CoverTab[83095]++
													return []byte{byte(l)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:17
		// _ = "end of CoverTab[83095]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:18
		_go_fuzz_dep_.CoverTab[83096]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:18
		// _ = "end of CoverTab[83096]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:18
	// _ = "end of CoverTab[83092]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:18
	_go_fuzz_dep_.CoverTab[83093]++
												var b []byte
												p := 1
												for i := 1; i < 127; {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:21
		_go_fuzz_dep_.CoverTab[83097]++
													b = append([]byte{byte((l % (p * 256)) / p)}, b...)
													p = p * 256
													l = l - l%p
													if l <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:25
			_go_fuzz_dep_.CoverTab[83098]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:26
			// _ = "end of CoverTab[83098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:27
			_go_fuzz_dep_.CoverTab[83099]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:27
			// _ = "end of CoverTab[83099]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:27
		// _ = "end of CoverTab[83097]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:28
	// _ = "end of CoverTab[83093]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:28
	_go_fuzz_dep_.CoverTab[83094]++
												return append([]byte{byte(128 + len(b))}, b...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:29
	// _ = "end of CoverTab[83094]"
}

// GetLengthFromASN returns the length of a slice of ASN1 encoded bytes from the ASN1 length header it contains.
func GetLengthFromASN(b []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:33
	_go_fuzz_dep_.CoverTab[83100]++
												if int(b[1]) <= 127 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:34
		_go_fuzz_dep_.CoverTab[83103]++
													return int(b[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:35
		// _ = "end of CoverTab[83103]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:36
		_go_fuzz_dep_.CoverTab[83104]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:36
		// _ = "end of CoverTab[83104]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:36
	// _ = "end of CoverTab[83100]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:36
	_go_fuzz_dep_.CoverTab[83101]++

												lb := b[2 : 2+int(b[1])-128]
												base := 1
												l := 0
												for i := len(lb) - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:41
		_go_fuzz_dep_.CoverTab[83105]++
													l += int(lb[i]) * base
													base = base * 256
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:43
		// _ = "end of CoverTab[83105]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:44
	// _ = "end of CoverTab[83101]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:44
	_go_fuzz_dep_.CoverTab[83102]++
												return l
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:45
	// _ = "end of CoverTab[83102]"
}

// GetNumberBytesInLengthHeader returns the number of bytes in the ASn1 header that indicate the length.
func GetNumberBytesInLengthHeader(b []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:49
	_go_fuzz_dep_.CoverTab[83106]++
												if int(b[1]) <= 127 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:50
		_go_fuzz_dep_.CoverTab[83108]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:51
		// _ = "end of CoverTab[83108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:52
		_go_fuzz_dep_.CoverTab[83109]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:52
		// _ = "end of CoverTab[83109]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:52
	// _ = "end of CoverTab[83106]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:52
	_go_fuzz_dep_.CoverTab[83107]++

												return 1 + int(b[1]) - 128
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:54
	// _ = "end of CoverTab[83107]"
}

// AddASNAppTag adds an ASN1 encoding application tag value to the raw bytes provided.
func AddASNAppTag(b []byte, tag int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:58
	_go_fuzz_dep_.CoverTab[83110]++
												r := asn1.RawValue{
		Class:		asn1.ClassApplication,
		IsCompound:	true,
		Tag:		tag,
		Bytes:		b,
	}
												ab, _ := asn1.Marshal(r)
												return ab
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:66
	// _ = "end of CoverTab[83110]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:67
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/asn1tools/tools.go:67
var _ = _go_fuzz_dep_.CoverTab
