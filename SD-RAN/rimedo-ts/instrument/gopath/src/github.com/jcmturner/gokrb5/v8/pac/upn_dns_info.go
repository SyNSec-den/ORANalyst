//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:1
)

import (
	"bytes"

	"github.com/jcmturner/rpc/v2/mstypes"
)

// UPNDNSInfo implements https://msdn.microsoft.com/en-us/library/dd240468.aspx
type UPNDNSInfo struct {
	UPNLength		uint16	// An unsigned 16-bit integer in little-endian format that specifies the length, in bytes, of the UPN field.
	UPNOffset		uint16	// An unsigned 16-bit integer in little-endian format that contains the offset to the beginning of the buffer, in bytes, from the beginning of the UPN_DNS_INFO structure.
	DNSDomainNameLength	uint16
	DNSDomainNameOffset	uint16
	Flags			uint32
	UPN			string
	DNSDomain		string
}

const (
	upnNoUPNAttr = 31	// The user account object does not have the userPrincipalName attribute ([MS-ADA3] section 2.349) set. A UPN constructed by concatenating the user name with the DNS domain name of the account domain is provided.
)

// Unmarshal bytes into the UPN_DNSInfo struct
func (k *UPNDNSInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:25
	_go_fuzz_dep_.CoverTab[87638]++

												r := mstypes.NewReader(bytes.NewReader(b))
												k.UPNLength, err = r.Uint16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:29
		_go_fuzz_dep_.CoverTab[87646]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:30
		// _ = "end of CoverTab[87646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:31
		_go_fuzz_dep_.CoverTab[87647]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:31
		// _ = "end of CoverTab[87647]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:31
	// _ = "end of CoverTab[87638]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:31
	_go_fuzz_dep_.CoverTab[87639]++
												k.UPNOffset, err = r.Uint16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:33
		_go_fuzz_dep_.CoverTab[87648]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:34
		// _ = "end of CoverTab[87648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:35
		_go_fuzz_dep_.CoverTab[87649]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:35
		// _ = "end of CoverTab[87649]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:35
	// _ = "end of CoverTab[87639]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:35
	_go_fuzz_dep_.CoverTab[87640]++
												k.DNSDomainNameLength, err = r.Uint16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:37
		_go_fuzz_dep_.CoverTab[87650]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:38
		// _ = "end of CoverTab[87650]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:39
		_go_fuzz_dep_.CoverTab[87651]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:39
		// _ = "end of CoverTab[87651]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:39
	// _ = "end of CoverTab[87640]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:39
	_go_fuzz_dep_.CoverTab[87641]++
												k.DNSDomainNameOffset, err = r.Uint16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:41
		_go_fuzz_dep_.CoverTab[87652]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:42
		// _ = "end of CoverTab[87652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:43
		_go_fuzz_dep_.CoverTab[87653]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:43
		// _ = "end of CoverTab[87653]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:43
	// _ = "end of CoverTab[87641]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:43
	_go_fuzz_dep_.CoverTab[87642]++
												k.Flags, err = r.Uint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:45
		_go_fuzz_dep_.CoverTab[87654]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:46
		// _ = "end of CoverTab[87654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:47
		_go_fuzz_dep_.CoverTab[87655]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:47
		// _ = "end of CoverTab[87655]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:47
	// _ = "end of CoverTab[87642]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:47
	_go_fuzz_dep_.CoverTab[87643]++
												ub := mstypes.NewReader(bytes.NewReader(b[k.UPNOffset : k.UPNOffset+k.UPNLength]))
												db := mstypes.NewReader(bytes.NewReader(b[k.DNSDomainNameOffset : k.DNSDomainNameOffset+k.DNSDomainNameLength]))

												u := make([]rune, k.UPNLength/2, k.UPNLength/2)
												for i := 0; i < len(u); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:52
		_go_fuzz_dep_.CoverTab[87656]++
													var r uint16
													r, err = ub.Uint16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:55
			_go_fuzz_dep_.CoverTab[87658]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:56
			// _ = "end of CoverTab[87658]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:57
			_go_fuzz_dep_.CoverTab[87659]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:57
			// _ = "end of CoverTab[87659]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:57
		// _ = "end of CoverTab[87656]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:57
		_go_fuzz_dep_.CoverTab[87657]++
													u[i] = rune(r)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:58
		// _ = "end of CoverTab[87657]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:59
	// _ = "end of CoverTab[87643]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:59
	_go_fuzz_dep_.CoverTab[87644]++
												k.UPN = string(u)
												d := make([]rune, k.DNSDomainNameLength/2, k.DNSDomainNameLength/2)
												for i := 0; i < len(d); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:62
		_go_fuzz_dep_.CoverTab[87660]++
													var r uint16
													r, err = db.Uint16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:65
			_go_fuzz_dep_.CoverTab[87662]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:66
			// _ = "end of CoverTab[87662]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:67
			_go_fuzz_dep_.CoverTab[87663]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:67
			// _ = "end of CoverTab[87663]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:67
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:67
		// _ = "end of CoverTab[87660]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:67
		_go_fuzz_dep_.CoverTab[87661]++
													d[i] = rune(r)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:68
		// _ = "end of CoverTab[87661]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:69
	// _ = "end of CoverTab[87644]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:69
	_go_fuzz_dep_.CoverTab[87645]++
												k.DNSDomain = string(d)

												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:72
	// _ = "end of CoverTab[87645]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/upn_dns_info.go:73
var _ = _go_fuzz_dep_.CoverTab
