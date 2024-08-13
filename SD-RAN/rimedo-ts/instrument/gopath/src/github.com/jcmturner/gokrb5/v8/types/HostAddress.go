//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:6
import (
	"bytes"
	"fmt"
	"net"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/iana/addrtype"
)

// HostAddresses implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.5
type HostAddresses []HostAddress

// HostAddress implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.5
type HostAddress struct {
	AddrType	int32	`asn1:"explicit,tag:0"`
	Address		[]byte	`asn1:"explicit,tag:1"`
}

// GetHostAddress returns a HostAddress struct from a string in the format <hostname>:<port>
func GetHostAddress(s string) (HostAddress, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:25
	_go_fuzz_dep_.CoverTab[85970]++
													var h HostAddress
													cAddr, _, err := net.SplitHostPort(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:28
		_go_fuzz_dep_.CoverTab[85973]++
														return h, fmt.Errorf("invalid format of client address: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:29
		// _ = "end of CoverTab[85973]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:30
		_go_fuzz_dep_.CoverTab[85974]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:30
		// _ = "end of CoverTab[85974]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:30
	// _ = "end of CoverTab[85970]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:30
	_go_fuzz_dep_.CoverTab[85971]++
													ip := net.ParseIP(cAddr)
													var ht int32
													if ip.To4() != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:33
		_go_fuzz_dep_.CoverTab[85975]++
														ht = addrtype.IPv4
														ip = ip.To4()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:35
		// _ = "end of CoverTab[85975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:36
		_go_fuzz_dep_.CoverTab[85976]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:36
		if ip.To16() != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:36
			_go_fuzz_dep_.CoverTab[85977]++
															ht = addrtype.IPv6
															ip = ip.To16()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:38
			// _ = "end of CoverTab[85977]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:39
			_go_fuzz_dep_.CoverTab[85978]++
															return h, fmt.Errorf("could not determine client's address types: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:40
			// _ = "end of CoverTab[85978]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:41
		// _ = "end of CoverTab[85976]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:41
	// _ = "end of CoverTab[85971]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:41
	_go_fuzz_dep_.CoverTab[85972]++
													h = HostAddress{
		AddrType:	ht,
		Address:	ip,
	}
													return h, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:46
	// _ = "end of CoverTab[85972]"
}

// GetAddress returns a string representation of the HostAddress.
func (h *HostAddress) GetAddress() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:50
	_go_fuzz_dep_.CoverTab[85979]++
													var b []byte
													_, err := asn1.Unmarshal(h.Address, &b)
													return string(b), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:53
	// _ = "end of CoverTab[85979]"
}

// LocalHostAddresses returns a HostAddresses struct for the local machines interface IP addresses.
func LocalHostAddresses() (ha HostAddresses, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:57
	_go_fuzz_dep_.CoverTab[85980]++
													ifs, err := net.Interfaces()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:59
		_go_fuzz_dep_.CoverTab[85983]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:60
		// _ = "end of CoverTab[85983]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:61
		_go_fuzz_dep_.CoverTab[85984]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:61
		// _ = "end of CoverTab[85984]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:61
	// _ = "end of CoverTab[85980]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:61
	_go_fuzz_dep_.CoverTab[85981]++
													for _, iface := range ifs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:62
		_go_fuzz_dep_.CoverTab[85985]++
														if iface.Flags&net.FlagLoopback != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:63
			_go_fuzz_dep_.CoverTab[85988]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:63
			return iface.Flags&net.FlagUp == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:63
			// _ = "end of CoverTab[85988]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:63
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:63
			_go_fuzz_dep_.CoverTab[85989]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:65
			// _ = "end of CoverTab[85989]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:66
			_go_fuzz_dep_.CoverTab[85990]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:66
			// _ = "end of CoverTab[85990]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:66
		// _ = "end of CoverTab[85985]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:66
		_go_fuzz_dep_.CoverTab[85986]++
														addrs, err := iface.Addrs()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:68
			_go_fuzz_dep_.CoverTab[85991]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:69
			// _ = "end of CoverTab[85991]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:70
			_go_fuzz_dep_.CoverTab[85992]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:70
			// _ = "end of CoverTab[85992]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:70
		// _ = "end of CoverTab[85986]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:70
		_go_fuzz_dep_.CoverTab[85987]++
														for _, addr := range addrs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:71
			_go_fuzz_dep_.CoverTab[85993]++
															var ip net.IP
															switch v := addr.(type) {
			case *net.IPNet:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:74
				_go_fuzz_dep_.CoverTab[85997]++
																ip = v.IP
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:75
				// _ = "end of CoverTab[85997]"
			case *net.IPAddr:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:76
				_go_fuzz_dep_.CoverTab[85998]++
																ip = v.IP
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:77
				// _ = "end of CoverTab[85998]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:78
			// _ = "end of CoverTab[85993]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:78
			_go_fuzz_dep_.CoverTab[85994]++
															var a HostAddress
															if ip.To16() == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:80
				_go_fuzz_dep_.CoverTab[85999]++

																continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:82
				// _ = "end of CoverTab[85999]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:83
				_go_fuzz_dep_.CoverTab[86000]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:83
				// _ = "end of CoverTab[86000]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:83
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:83
			// _ = "end of CoverTab[85994]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:83
			_go_fuzz_dep_.CoverTab[85995]++
															if ip.To4() != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:84
				_go_fuzz_dep_.CoverTab[86001]++

																a.AddrType = addrtype.IPv4
																a.Address = ip.To4()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:87
				// _ = "end of CoverTab[86001]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:88
				_go_fuzz_dep_.CoverTab[86002]++
																a.AddrType = addrtype.IPv6
																a.Address = ip.To16()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:90
				// _ = "end of CoverTab[86002]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:91
			// _ = "end of CoverTab[85995]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:91
			_go_fuzz_dep_.CoverTab[85996]++
															ha = append(ha, a)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:92
			// _ = "end of CoverTab[85996]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:93
		// _ = "end of CoverTab[85987]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:94
	// _ = "end of CoverTab[85981]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:94
	_go_fuzz_dep_.CoverTab[85982]++
													return ha, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:95
	// _ = "end of CoverTab[85982]"
}

// HostAddressesFromNetIPs returns a HostAddresses type from a slice of net.IP
func HostAddressesFromNetIPs(ips []net.IP) (ha HostAddresses) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:99
	_go_fuzz_dep_.CoverTab[86003]++
													for _, ip := range ips {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:100
		_go_fuzz_dep_.CoverTab[86005]++
														ha = append(ha, HostAddressFromNetIP(ip))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:101
		// _ = "end of CoverTab[86005]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:102
	// _ = "end of CoverTab[86003]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:102
	_go_fuzz_dep_.CoverTab[86004]++
													return ha
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:103
	// _ = "end of CoverTab[86004]"
}

// HostAddressFromNetIP returns a HostAddress type from a net.IP
func HostAddressFromNetIP(ip net.IP) HostAddress {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:107
	_go_fuzz_dep_.CoverTab[86006]++
													if ip.To4() != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:108
		_go_fuzz_dep_.CoverTab[86008]++

														return HostAddress{
			AddrType:	addrtype.IPv4,
			Address:	ip.To4(),
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:113
		// _ = "end of CoverTab[86008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:114
		_go_fuzz_dep_.CoverTab[86009]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:114
		// _ = "end of CoverTab[86009]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:114
	// _ = "end of CoverTab[86006]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:114
	_go_fuzz_dep_.CoverTab[86007]++
													return HostAddress{
		AddrType:	addrtype.IPv6,
		Address:	ip.To16(),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:118
	// _ = "end of CoverTab[86007]"
}

// HostAddressesEqual tests if two HostAddress slices are equal.
func HostAddressesEqual(h, a []HostAddress) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:122
	_go_fuzz_dep_.CoverTab[86010]++
													if len(h) != len(a) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:123
		_go_fuzz_dep_.CoverTab[86013]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:124
		// _ = "end of CoverTab[86013]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:125
		_go_fuzz_dep_.CoverTab[86014]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:125
		// _ = "end of CoverTab[86014]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:125
	// _ = "end of CoverTab[86010]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:125
	_go_fuzz_dep_.CoverTab[86011]++
													for _, e := range a {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:126
		_go_fuzz_dep_.CoverTab[86015]++
														var found bool
														for _, i := range h {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:128
			_go_fuzz_dep_.CoverTab[86017]++
															if e.Equal(i) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:129
				_go_fuzz_dep_.CoverTab[86018]++
																found = true
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:131
				// _ = "end of CoverTab[86018]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:132
				_go_fuzz_dep_.CoverTab[86019]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:132
				// _ = "end of CoverTab[86019]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:132
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:132
			// _ = "end of CoverTab[86017]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:133
		// _ = "end of CoverTab[86015]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:133
		_go_fuzz_dep_.CoverTab[86016]++
														if !found {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:134
			_go_fuzz_dep_.CoverTab[86020]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:135
			// _ = "end of CoverTab[86020]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:136
			_go_fuzz_dep_.CoverTab[86021]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:136
			// _ = "end of CoverTab[86021]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:136
		// _ = "end of CoverTab[86016]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:137
	// _ = "end of CoverTab[86011]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:137
	_go_fuzz_dep_.CoverTab[86012]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:138
	// _ = "end of CoverTab[86012]"
}

// HostAddressesContains tests if a HostAddress is contained in a HostAddress slice.
func HostAddressesContains(h []HostAddress, a HostAddress) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:142
	_go_fuzz_dep_.CoverTab[86022]++
													for _, e := range h {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:143
		_go_fuzz_dep_.CoverTab[86024]++
														if e.Equal(a) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:144
			_go_fuzz_dep_.CoverTab[86025]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:145
			// _ = "end of CoverTab[86025]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:146
			_go_fuzz_dep_.CoverTab[86026]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:146
			// _ = "end of CoverTab[86026]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:146
		// _ = "end of CoverTab[86024]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:147
	// _ = "end of CoverTab[86022]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:147
	_go_fuzz_dep_.CoverTab[86023]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:148
	// _ = "end of CoverTab[86023]"
}

// Equal tests if the HostAddress is equal to another HostAddress provided.
func (h *HostAddress) Equal(a HostAddress) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:152
	_go_fuzz_dep_.CoverTab[86027]++
													if h.AddrType != a.AddrType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:153
		_go_fuzz_dep_.CoverTab[86029]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:154
		// _ = "end of CoverTab[86029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:155
		_go_fuzz_dep_.CoverTab[86030]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:155
		// _ = "end of CoverTab[86030]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:155
	// _ = "end of CoverTab[86027]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:155
	_go_fuzz_dep_.CoverTab[86028]++
													return bytes.Equal(h.Address, a.Address)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:156
	// _ = "end of CoverTab[86028]"
}

// Contains tests if a HostAddress is contained within the HostAddresses struct.
func (h *HostAddresses) Contains(a HostAddress) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:160
	_go_fuzz_dep_.CoverTab[86031]++
													for _, e := range *h {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:161
		_go_fuzz_dep_.CoverTab[86033]++
														if e.Equal(a) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:162
			_go_fuzz_dep_.CoverTab[86034]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:163
			// _ = "end of CoverTab[86034]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:164
			_go_fuzz_dep_.CoverTab[86035]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:164
			// _ = "end of CoverTab[86035]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:164
		// _ = "end of CoverTab[86033]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:165
	// _ = "end of CoverTab[86031]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:165
	_go_fuzz_dep_.CoverTab[86032]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:166
	// _ = "end of CoverTab[86032]"
}

// Equal tests if a HostAddress slice is equal to the HostAddresses struct.
func (h *HostAddresses) Equal(a []HostAddress) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:170
	_go_fuzz_dep_.CoverTab[86036]++
													if len(*h) != len(a) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:171
		_go_fuzz_dep_.CoverTab[86039]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:172
		// _ = "end of CoverTab[86039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:173
		_go_fuzz_dep_.CoverTab[86040]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:173
		// _ = "end of CoverTab[86040]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:173
	// _ = "end of CoverTab[86036]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:173
	_go_fuzz_dep_.CoverTab[86037]++
													for _, e := range a {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:174
		_go_fuzz_dep_.CoverTab[86041]++
														if !h.Contains(e) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:175
			_go_fuzz_dep_.CoverTab[86042]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:176
			// _ = "end of CoverTab[86042]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:177
			_go_fuzz_dep_.CoverTab[86043]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:177
			// _ = "end of CoverTab[86043]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:177
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:177
		// _ = "end of CoverTab[86041]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:178
	// _ = "end of CoverTab[86037]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:178
	_go_fuzz_dep_.CoverTab[86038]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:179
	// _ = "end of CoverTab[86038]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:180
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/HostAddress.go:180
var _ = _go_fuzz_dep_.CoverTab
