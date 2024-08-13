//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
package config

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:1
)

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"

	"github.com/jcmturner/dnsutils/v2"
)

// GetKDCs returns the count of KDCs available and a map of KDC host names keyed on preference order.
func (c *Config) GetKDCs(realm string, tcp bool) (int, map[int]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:14
	_go_fuzz_dep_.CoverTab[83259]++
												if realm == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:15
		_go_fuzz_dep_.CoverTab[83268]++
													realm = c.LibDefaults.DefaultRealm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:16
		// _ = "end of CoverTab[83268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:17
		_go_fuzz_dep_.CoverTab[83269]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:17
		// _ = "end of CoverTab[83269]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:17
	// _ = "end of CoverTab[83259]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:17
	_go_fuzz_dep_.CoverTab[83260]++
												kdcs := make(map[int]string)
												var count int

	// Get the KDCs from the krb5.conf.
	var ks []string
	for _, r := range c.Realms {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:23
		_go_fuzz_dep_.CoverTab[83270]++
													if r.Realm != realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:24
			_go_fuzz_dep_.CoverTab[83272]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:25
			// _ = "end of CoverTab[83272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:26
			_go_fuzz_dep_.CoverTab[83273]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:26
			// _ = "end of CoverTab[83273]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:26
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:26
		// _ = "end of CoverTab[83270]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:26
		_go_fuzz_dep_.CoverTab[83271]++
													ks = r.KDC
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:27
		// _ = "end of CoverTab[83271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:28
	// _ = "end of CoverTab[83260]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:28
	_go_fuzz_dep_.CoverTab[83261]++
												count = len(ks)

												if count > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:31
		_go_fuzz_dep_.CoverTab[83274]++

													kdcs = randServOrder(ks)
													return count, kdcs, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:34
		// _ = "end of CoverTab[83274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:35
		_go_fuzz_dep_.CoverTab[83275]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:35
		// _ = "end of CoverTab[83275]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:35
	// _ = "end of CoverTab[83261]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:35
	_go_fuzz_dep_.CoverTab[83262]++

												if !c.LibDefaults.DNSLookupKDC {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:37
		_go_fuzz_dep_.CoverTab[83276]++
													return count, kdcs, fmt.Errorf("no KDCs defined in configuration for realm %s", realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:38
		// _ = "end of CoverTab[83276]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:39
		_go_fuzz_dep_.CoverTab[83277]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:39
		// _ = "end of CoverTab[83277]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:39
	// _ = "end of CoverTab[83262]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:39
	_go_fuzz_dep_.CoverTab[83263]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:42
	proto := "udp"
	if tcp {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:43
		_go_fuzz_dep_.CoverTab[83278]++
													proto = "tcp"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:44
		// _ = "end of CoverTab[83278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:45
		_go_fuzz_dep_.CoverTab[83279]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:45
		// _ = "end of CoverTab[83279]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:45
	// _ = "end of CoverTab[83263]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:45
	_go_fuzz_dep_.CoverTab[83264]++
												index, addrs, err := dnsutils.OrderedSRV("kerberos", proto, realm)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:47
		_go_fuzz_dep_.CoverTab[83280]++
													return count, kdcs, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:48
		// _ = "end of CoverTab[83280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:49
		_go_fuzz_dep_.CoverTab[83281]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:49
		// _ = "end of CoverTab[83281]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:49
	// _ = "end of CoverTab[83264]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:49
	_go_fuzz_dep_.CoverTab[83265]++
												if len(addrs) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:50
		_go_fuzz_dep_.CoverTab[83282]++
													return count, kdcs, fmt.Errorf("no KDC SRV records found for realm %s", realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:51
		// _ = "end of CoverTab[83282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:52
		_go_fuzz_dep_.CoverTab[83283]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:52
		// _ = "end of CoverTab[83283]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:52
	// _ = "end of CoverTab[83265]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:52
	_go_fuzz_dep_.CoverTab[83266]++
												count = index
												for k, v := range addrs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:54
		_go_fuzz_dep_.CoverTab[83284]++
													kdcs[k] = strings.TrimRight(v.Target, ".") + ":" + strconv.Itoa(int(v.Port))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:55
		// _ = "end of CoverTab[83284]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:56
	// _ = "end of CoverTab[83266]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:56
	_go_fuzz_dep_.CoverTab[83267]++
												return count, kdcs, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:57
	// _ = "end of CoverTab[83267]"
}

// GetKpasswdServers returns the count of kpasswd servers available and a map of kpasswd host names keyed on preference order.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:60
// https://web.mit.edu/kerberos/krb5-latest/doc/admin/conf_files/krb5_conf.html#realms - see kpasswd_server section
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:62
func (c *Config) GetKpasswdServers(realm string, tcp bool) (int, map[int]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:62
	_go_fuzz_dep_.CoverTab[83285]++
												kdcs := make(map[int]string)
												var count int

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:67
	if c.LibDefaults.DNSLookupKDC {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:67
		_go_fuzz_dep_.CoverTab[83287]++
													proto := "udp"
													if tcp {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:69
			_go_fuzz_dep_.CoverTab[83292]++
														proto = "tcp"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:70
			// _ = "end of CoverTab[83292]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:71
			_go_fuzz_dep_.CoverTab[83293]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:71
			// _ = "end of CoverTab[83293]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:71
		// _ = "end of CoverTab[83287]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:71
		_go_fuzz_dep_.CoverTab[83288]++
													c, addrs, err := dnsutils.OrderedSRV("kpasswd", proto, realm)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:73
			_go_fuzz_dep_.CoverTab[83294]++
														return count, kdcs, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:74
			// _ = "end of CoverTab[83294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:75
			_go_fuzz_dep_.CoverTab[83295]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:75
			// _ = "end of CoverTab[83295]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:75
		// _ = "end of CoverTab[83288]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:75
		_go_fuzz_dep_.CoverTab[83289]++
													if c < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:76
			_go_fuzz_dep_.CoverTab[83296]++
														c, addrs, err = dnsutils.OrderedSRV("kerberos-adm", proto, realm)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:78
				_go_fuzz_dep_.CoverTab[83297]++
															return count, kdcs, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:79
				// _ = "end of CoverTab[83297]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:80
				_go_fuzz_dep_.CoverTab[83298]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:80
				// _ = "end of CoverTab[83298]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:80
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:80
			// _ = "end of CoverTab[83296]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:81
			_go_fuzz_dep_.CoverTab[83299]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:81
			// _ = "end of CoverTab[83299]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:81
		// _ = "end of CoverTab[83289]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:81
		_go_fuzz_dep_.CoverTab[83290]++
													if len(addrs) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:82
			_go_fuzz_dep_.CoverTab[83300]++
														return count, kdcs, fmt.Errorf("no kpasswd or kadmin SRV records found for realm %s", realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:83
			// _ = "end of CoverTab[83300]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:84
			_go_fuzz_dep_.CoverTab[83301]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:84
			// _ = "end of CoverTab[83301]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:84
		// _ = "end of CoverTab[83290]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:84
		_go_fuzz_dep_.CoverTab[83291]++
													count = c
													for k, v := range addrs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:86
			_go_fuzz_dep_.CoverTab[83302]++
														kdcs[k] = strings.TrimRight(v.Target, ".") + ":" + strconv.Itoa(int(v.Port))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:87
			// _ = "end of CoverTab[83302]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:88
		// _ = "end of CoverTab[83291]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:89
		_go_fuzz_dep_.CoverTab[83303]++
		// Get the KDCs from the krb5.conf an order them randomly for preference.
		var ks []string
		var ka []string
		for _, r := range c.Realms {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:93
			_go_fuzz_dep_.CoverTab[83307]++
														if r.Realm == realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:94
				_go_fuzz_dep_.CoverTab[83308]++
															ks = r.KPasswdServer
															ka = r.AdminServer
															break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:97
				// _ = "end of CoverTab[83308]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:98
				_go_fuzz_dep_.CoverTab[83309]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:98
				// _ = "end of CoverTab[83309]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:98
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:98
			// _ = "end of CoverTab[83307]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:99
		// _ = "end of CoverTab[83303]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:99
		_go_fuzz_dep_.CoverTab[83304]++
													if len(ks) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:100
			_go_fuzz_dep_.CoverTab[83310]++
														for _, k := range ka {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:101
				_go_fuzz_dep_.CoverTab[83311]++
															h, _, err := net.SplitHostPort(k)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:103
					_go_fuzz_dep_.CoverTab[83313]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:104
					// _ = "end of CoverTab[83313]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:105
					_go_fuzz_dep_.CoverTab[83314]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:105
					// _ = "end of CoverTab[83314]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:105
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:105
				// _ = "end of CoverTab[83311]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:105
				_go_fuzz_dep_.CoverTab[83312]++
															ks = append(ks, h+":464")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:106
				// _ = "end of CoverTab[83312]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:107
			// _ = "end of CoverTab[83310]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:108
			_go_fuzz_dep_.CoverTab[83315]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:108
			// _ = "end of CoverTab[83315]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:108
		// _ = "end of CoverTab[83304]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:108
		_go_fuzz_dep_.CoverTab[83305]++
													count = len(ks)
													if count < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:110
			_go_fuzz_dep_.CoverTab[83316]++
														return count, kdcs, fmt.Errorf("no kpasswd or kadmin defined in configuration for realm %s", realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:111
			// _ = "end of CoverTab[83316]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:112
			_go_fuzz_dep_.CoverTab[83317]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:112
			// _ = "end of CoverTab[83317]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:112
		// _ = "end of CoverTab[83305]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:112
		_go_fuzz_dep_.CoverTab[83306]++
													kdcs = randServOrder(ks)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:113
		// _ = "end of CoverTab[83306]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:114
	// _ = "end of CoverTab[83285]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:114
	_go_fuzz_dep_.CoverTab[83286]++
												return count, kdcs, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:115
	// _ = "end of CoverTab[83286]"
}

func randServOrder(ks []string) map[int]string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:118
	_go_fuzz_dep_.CoverTab[83318]++
												kdcs := make(map[int]string)
												count := len(ks)
												i := 1
												if count > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:122
		_go_fuzz_dep_.CoverTab[83320]++
													l := len(ks)
													for l > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:124
			_go_fuzz_dep_.CoverTab[83321]++
														ri := rand.Intn(l)
														kdcs[i] = ks[ri]
														if l > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:127
				_go_fuzz_dep_.CoverTab[83323]++

															ks[len(ks)-1], ks[ri] = ks[ri], ks[len(ks)-1]
															ks = ks[:len(ks)-1]
															l = len(ks)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:131
				// _ = "end of CoverTab[83323]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:132
				_go_fuzz_dep_.CoverTab[83324]++
															l = 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:133
				// _ = "end of CoverTab[83324]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:134
			// _ = "end of CoverTab[83321]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:134
			_go_fuzz_dep_.CoverTab[83322]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:135
			// _ = "end of CoverTab[83322]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:136
		// _ = "end of CoverTab[83320]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:137
		_go_fuzz_dep_.CoverTab[83325]++
													kdcs[i] = ks[0]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:138
		// _ = "end of CoverTab[83325]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:139
	// _ = "end of CoverTab[83318]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:139
	_go_fuzz_dep_.CoverTab[83319]++
												return kdcs
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:140
	// _ = "end of CoverTab[83319]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/hosts.go:141
var _ = _go_fuzz_dep_.CoverTab
