//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:1
// Package config implements KRB5 client and service configuration as described at https://web.mit.edu/kerberos/krb5-latest/doc/admin/conf_files/krb5_conf.html
package config

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:2
)

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

// Config represents the KRB5 configuration.
type Config struct {
	LibDefaults	LibDefaults
	Realms		[]Realm
	DomainRealm	DomainRealm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:31
}

// WeakETypeList is a list of encryption types that have been deemed weak.
const WeakETypeList = "des-cbc-crc des-cbc-md4 des-cbc-md5 des-cbc-raw des3-cbc-raw des-hmac-sha1 arcfour-hmac-exp rc4-hmac-exp arcfour-hmac-md5-exp des"

// New creates a new config struct instance.
func New() *Config {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:37
	_go_fuzz_dep_.CoverTab[83326]++
												d := make(DomainRealm)
												return &Config{
		LibDefaults:	newLibDefaults(),
		DomainRealm:	d,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:42
	// _ = "end of CoverTab[83326]"
}

// LibDefaults represents the [libdefaults] section of the configuration.
type LibDefaults struct {
	AllowWeakCrypto	bool	//default false
	// ap_req_checksum_type int //unlikely to support this
	Canonicalize	bool		//default false
	CCacheType	int		//default is 4. unlikely to implement older
	Clockskew	time.Duration	//max allowed skew in seconds, default 300
	//Default_ccache_name string // default /tmp/krb5cc_%{uid} //Not implementing as will hold in memory
	DefaultClientKeytabName	string	//default /usr/local/var/krb5/user/%{euid}/client.keytab
	DefaultKeytabName	string	//default /etc/krb5.keytab
	DefaultRealm		string
	DefaultTGSEnctypes	[]string	//default aes256-cts-hmac-sha1-96 aes128-cts-hmac-sha1-96 des3-cbc-sha1 arcfour-hmac-md5 camellia256-cts-cmac camellia128-cts-cmac des-cbc-crc des-cbc-md5 des-cbc-md4
	DefaultTktEnctypes	[]string	//default aes256-cts-hmac-sha1-96 aes128-cts-hmac-sha1-96 des3-cbc-sha1 arcfour-hmac-md5 camellia256-cts-cmac camellia128-cts-cmac des-cbc-crc des-cbc-md5 des-cbc-md4
	DefaultTGSEnctypeIDs	[]int32		//default aes256-cts-hmac-sha1-96 aes128-cts-hmac-sha1-96 des3-cbc-sha1 arcfour-hmac-md5 camellia256-cts-cmac camellia128-cts-cmac des-cbc-crc des-cbc-md5 des-cbc-md4
	DefaultTktEnctypeIDs	[]int32		//default aes256-cts-hmac-sha1-96 aes128-cts-hmac-sha1-96 des3-cbc-sha1 arcfour-hmac-md5 camellia256-cts-cmac camellia128-cts-cmac des-cbc-crc des-cbc-md5 des-cbc-md4
	DNSCanonicalizeHostname	bool		//default true
	DNSLookupKDC		bool		//default false
	DNSLookupRealm		bool
	ExtraAddresses		[]net.IP	//Not implementing yet
	Forwardable		bool		//default false
	IgnoreAcceptorHostname	bool		//default false
	K5LoginAuthoritative	bool		//default false
	K5LoginDirectory	string		//default user's home directory. Must be owned by the user or root
	KDCDefaultOptions	asn1.BitString	//default 0x00000010 (KDC_OPT_RENEWABLE_OK)
	KDCTimeSync		int		//default 1
	//kdc_req_checksum_type int //unlikely to implement as for very old KDCs
	NoAddresses		bool		//default true
	PermittedEnctypes	[]string	//default aes256-cts-hmac-sha1-96 aes128-cts-hmac-sha1-96 des3-cbc-sha1 arcfour-hmac-md5 camellia256-cts-cmac camellia128-cts-cmac des-cbc-crc des-cbc-md5 des-cbc-md4
	PermittedEnctypeIDs	[]int32
	//plugin_base_dir string //not supporting plugins
	PreferredPreauthTypes	[]int		//default “17, 16, 15, 14”, which forces libkrb5 to attempt to use PKINIT if it is supported
	Proxiable		bool		//default false
	RDNS			bool		//default true
	RealmTryDomains		int		//default -1
	RenewLifetime		time.Duration	//default 0
	SafeChecksumType	int		//default 8
	TicketLifetime		time.Duration	//default 1 day
	UDPPreferenceLimit	int		// 1 means to always use tcp. MIT krb5 has a default value of 1465, and it prevents user setting more than 32700.
	VerifyAPReqNofail	bool		//default false
}

// Create a new LibDefaults struct.
func newLibDefaults() LibDefaults {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:87
	_go_fuzz_dep_.CoverTab[83327]++
												uid := "0"
												var hdir string
												usr, _ := user.Current()
												if usr != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:91
		_go_fuzz_dep_.CoverTab[83329]++
													uid = usr.Uid
													hdir = usr.HomeDir
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:93
		// _ = "end of CoverTab[83329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:94
		_go_fuzz_dep_.CoverTab[83330]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:94
		// _ = "end of CoverTab[83330]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:94
	// _ = "end of CoverTab[83327]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:94
	_go_fuzz_dep_.CoverTab[83328]++
												opts := asn1.BitString{}
												opts.Bytes, _ = hex.DecodeString("00000010")
												opts.BitLength = len(opts.Bytes) * 8
												return LibDefaults{
		CCacheType:			4,
		Clockskew:			time.Duration(300) * time.Second,
		DefaultClientKeytabName:	fmt.Sprintf("/usr/local/var/krb5/user/%s/client.keytab", uid),
		DefaultKeytabName:		"/etc/krb5.keytab",
		DefaultTGSEnctypes:		[]string{"aes256-cts-hmac-sha1-96", "aes128-cts-hmac-sha1-96", "des3-cbc-sha1", "arcfour-hmac-md5", "camellia256-cts-cmac", "camellia128-cts-cmac", "des-cbc-crc", "des-cbc-md5", "des-cbc-md4"},
		DefaultTktEnctypes:		[]string{"aes256-cts-hmac-sha1-96", "aes128-cts-hmac-sha1-96", "des3-cbc-sha1", "arcfour-hmac-md5", "camellia256-cts-cmac", "camellia128-cts-cmac", "des-cbc-crc", "des-cbc-md5", "des-cbc-md4"},
		DNSCanonicalizeHostname:	true,
		K5LoginDirectory:		hdir,
		KDCDefaultOptions:		opts,
		KDCTimeSync:			1,
		NoAddresses:			true,
		PermittedEnctypes:		[]string{"aes256-cts-hmac-sha1-96", "aes128-cts-hmac-sha1-96", "des3-cbc-sha1", "arcfour-hmac-md5", "camellia256-cts-cmac", "camellia128-cts-cmac", "des-cbc-crc", "des-cbc-md5", "des-cbc-md4"},
		RDNS:				true,
		RealmTryDomains:		-1,
		SafeChecksumType:		8,
		TicketLifetime:			time.Duration(24) * time.Hour,
		UDPPreferenceLimit:		1465,
		PreferredPreauthTypes:		[]int{17, 16, 15, 14},
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:117
	// _ = "end of CoverTab[83328]"
}

// Parse the lines of the [libdefaults] section of the configuration into the LibDefaults struct.
func (l *LibDefaults) parseLines(lines []string) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:121
	_go_fuzz_dep_.CoverTab[83331]++
												for _, line := range lines {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:122
		_go_fuzz_dep_.CoverTab[83333]++

													if idx := strings.IndexAny(line, "#;"); idx != -1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:124
			_go_fuzz_dep_.CoverTab[83337]++
														line = line[:idx]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:125
			// _ = "end of CoverTab[83337]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:126
			_go_fuzz_dep_.CoverTab[83338]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:126
			// _ = "end of CoverTab[83338]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:126
		// _ = "end of CoverTab[83333]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:126
		_go_fuzz_dep_.CoverTab[83334]++
													line = strings.TrimSpace(line)
													if line == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:128
			_go_fuzz_dep_.CoverTab[83339]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:129
			// _ = "end of CoverTab[83339]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:130
			_go_fuzz_dep_.CoverTab[83340]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:130
			// _ = "end of CoverTab[83340]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:130
		// _ = "end of CoverTab[83334]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:130
		_go_fuzz_dep_.CoverTab[83335]++
													if !strings.Contains(line, "=") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:131
			_go_fuzz_dep_.CoverTab[83341]++
														return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:132
			// _ = "end of CoverTab[83341]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:133
			_go_fuzz_dep_.CoverTab[83342]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:133
			// _ = "end of CoverTab[83342]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:133
		// _ = "end of CoverTab[83335]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:133
		_go_fuzz_dep_.CoverTab[83336]++

													p := strings.Split(line, "=")
													key := strings.TrimSpace(strings.ToLower(p[0]))
													switch key {
		case "allow_weak_crypto":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:138
			_go_fuzz_dep_.CoverTab[83343]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:140
				_go_fuzz_dep_.CoverTab[83396]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:141
				// _ = "end of CoverTab[83396]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:142
				_go_fuzz_dep_.CoverTab[83397]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:142
				// _ = "end of CoverTab[83397]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:142
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:142
			// _ = "end of CoverTab[83343]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:142
			_go_fuzz_dep_.CoverTab[83344]++
														l.AllowWeakCrypto = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:143
			// _ = "end of CoverTab[83344]"
		case "canonicalize":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:144
			_go_fuzz_dep_.CoverTab[83345]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:146
				_go_fuzz_dep_.CoverTab[83398]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:147
				// _ = "end of CoverTab[83398]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:148
				_go_fuzz_dep_.CoverTab[83399]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:148
				// _ = "end of CoverTab[83399]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:148
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:148
			// _ = "end of CoverTab[83345]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:148
			_go_fuzz_dep_.CoverTab[83346]++
														l.Canonicalize = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:149
			// _ = "end of CoverTab[83346]"
		case "ccache_type":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:150
			_go_fuzz_dep_.CoverTab[83347]++
														p[1] = strings.TrimSpace(p[1])
														v, err := strconv.ParseUint(p[1], 10, 32)
														if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				_go_fuzz_dep_.CoverTab[83400]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				return v < 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				// _ = "end of CoverTab[83400]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				_go_fuzz_dep_.CoverTab[83401]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				return v > 4
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				// _ = "end of CoverTab[83401]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:153
				_go_fuzz_dep_.CoverTab[83402]++
															return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:154
				// _ = "end of CoverTab[83402]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:155
				_go_fuzz_dep_.CoverTab[83403]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:155
				// _ = "end of CoverTab[83403]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:155
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:155
			// _ = "end of CoverTab[83347]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:155
			_go_fuzz_dep_.CoverTab[83348]++
														l.CCacheType = int(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:156
			// _ = "end of CoverTab[83348]"
		case "clockskew":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:157
			_go_fuzz_dep_.CoverTab[83349]++
														d, err := parseDuration(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:159
				_go_fuzz_dep_.CoverTab[83404]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:160
				// _ = "end of CoverTab[83404]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:161
				_go_fuzz_dep_.CoverTab[83405]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:161
				// _ = "end of CoverTab[83405]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:161
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:161
			// _ = "end of CoverTab[83349]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:161
			_go_fuzz_dep_.CoverTab[83350]++
														l.Clockskew = d
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:162
			// _ = "end of CoverTab[83350]"
		case "default_client_keytab_name":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:163
			_go_fuzz_dep_.CoverTab[83351]++
														l.DefaultClientKeytabName = strings.TrimSpace(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:164
			// _ = "end of CoverTab[83351]"
		case "default_keytab_name":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:165
			_go_fuzz_dep_.CoverTab[83352]++
														l.DefaultKeytabName = strings.TrimSpace(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:166
			// _ = "end of CoverTab[83352]"
		case "default_realm":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:167
			_go_fuzz_dep_.CoverTab[83353]++
														l.DefaultRealm = strings.TrimSpace(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:168
			// _ = "end of CoverTab[83353]"
		case "default_tgs_enctypes":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:169
			_go_fuzz_dep_.CoverTab[83354]++
														l.DefaultTGSEnctypes = strings.Fields(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:170
			// _ = "end of CoverTab[83354]"
		case "default_tkt_enctypes":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:171
			_go_fuzz_dep_.CoverTab[83355]++
														l.DefaultTktEnctypes = strings.Fields(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:172
			// _ = "end of CoverTab[83355]"
		case "dns_canonicalize_hostname":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:173
			_go_fuzz_dep_.CoverTab[83356]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:175
				_go_fuzz_dep_.CoverTab[83406]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:176
				// _ = "end of CoverTab[83406]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:177
				_go_fuzz_dep_.CoverTab[83407]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:177
				// _ = "end of CoverTab[83407]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:177
			// _ = "end of CoverTab[83356]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:177
			_go_fuzz_dep_.CoverTab[83357]++
														l.DNSCanonicalizeHostname = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:178
			// _ = "end of CoverTab[83357]"
		case "dns_lookup_kdc":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:179
			_go_fuzz_dep_.CoverTab[83358]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:181
				_go_fuzz_dep_.CoverTab[83408]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:182
				// _ = "end of CoverTab[83408]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:183
				_go_fuzz_dep_.CoverTab[83409]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:183
				// _ = "end of CoverTab[83409]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:183
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:183
			// _ = "end of CoverTab[83358]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:183
			_go_fuzz_dep_.CoverTab[83359]++
														l.DNSLookupKDC = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:184
			// _ = "end of CoverTab[83359]"
		case "dns_lookup_realm":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:185
			_go_fuzz_dep_.CoverTab[83360]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:187
				_go_fuzz_dep_.CoverTab[83410]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:188
				// _ = "end of CoverTab[83410]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:189
				_go_fuzz_dep_.CoverTab[83411]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:189
				// _ = "end of CoverTab[83411]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:189
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:189
			// _ = "end of CoverTab[83360]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:189
			_go_fuzz_dep_.CoverTab[83361]++
														l.DNSLookupRealm = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:190
			// _ = "end of CoverTab[83361]"
		case "extra_addresses":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:191
			_go_fuzz_dep_.CoverTab[83362]++
														ipStr := strings.TrimSpace(p[1])
														for _, ip := range strings.Split(ipStr, ",") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:193
				_go_fuzz_dep_.CoverTab[83412]++
															if eip := net.ParseIP(ip); eip != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:194
					_go_fuzz_dep_.CoverTab[83413]++
																l.ExtraAddresses = append(l.ExtraAddresses, eip)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:195
					// _ = "end of CoverTab[83413]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:196
					_go_fuzz_dep_.CoverTab[83414]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:196
					// _ = "end of CoverTab[83414]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:196
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:196
				// _ = "end of CoverTab[83412]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:197
			// _ = "end of CoverTab[83362]"
		case "forwardable":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:198
			_go_fuzz_dep_.CoverTab[83363]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:200
				_go_fuzz_dep_.CoverTab[83415]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:201
				// _ = "end of CoverTab[83415]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:202
				_go_fuzz_dep_.CoverTab[83416]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:202
				// _ = "end of CoverTab[83416]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:202
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:202
			// _ = "end of CoverTab[83363]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:202
			_go_fuzz_dep_.CoverTab[83364]++
														l.Forwardable = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:203
			// _ = "end of CoverTab[83364]"
		case "ignore_acceptor_hostname":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:204
			_go_fuzz_dep_.CoverTab[83365]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:206
				_go_fuzz_dep_.CoverTab[83417]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:207
				// _ = "end of CoverTab[83417]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:208
				_go_fuzz_dep_.CoverTab[83418]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:208
				// _ = "end of CoverTab[83418]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:208
			// _ = "end of CoverTab[83365]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:208
			_go_fuzz_dep_.CoverTab[83366]++
														l.IgnoreAcceptorHostname = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:209
			// _ = "end of CoverTab[83366]"
		case "k5login_authoritative":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:210
			_go_fuzz_dep_.CoverTab[83367]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:212
				_go_fuzz_dep_.CoverTab[83419]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:213
				// _ = "end of CoverTab[83419]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:214
				_go_fuzz_dep_.CoverTab[83420]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:214
				// _ = "end of CoverTab[83420]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:214
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:214
			// _ = "end of CoverTab[83367]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:214
			_go_fuzz_dep_.CoverTab[83368]++
														l.K5LoginAuthoritative = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:215
			// _ = "end of CoverTab[83368]"
		case "k5login_directory":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:216
			_go_fuzz_dep_.CoverTab[83369]++
														l.K5LoginDirectory = strings.TrimSpace(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:217
			// _ = "end of CoverTab[83369]"
		case "kdc_default_options":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:218
			_go_fuzz_dep_.CoverTab[83370]++
														v := strings.TrimSpace(p[1])
														v = strings.Replace(v, "0x", "", -1)
														b, err := hex.DecodeString(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:222
				_go_fuzz_dep_.CoverTab[83421]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:223
				// _ = "end of CoverTab[83421]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:224
				_go_fuzz_dep_.CoverTab[83422]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:224
				// _ = "end of CoverTab[83422]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:224
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:224
			// _ = "end of CoverTab[83370]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:224
			_go_fuzz_dep_.CoverTab[83371]++
														l.KDCDefaultOptions.Bytes = b
														l.KDCDefaultOptions.BitLength = len(b) * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:226
			// _ = "end of CoverTab[83371]"
		case "kdc_timesync":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:227
			_go_fuzz_dep_.CoverTab[83372]++
														p[1] = strings.TrimSpace(p[1])
														v, err := strconv.ParseInt(p[1], 10, 32)
														if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:230
				_go_fuzz_dep_.CoverTab[83423]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:230
				return v < 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:230
				// _ = "end of CoverTab[83423]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:230
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:230
				_go_fuzz_dep_.CoverTab[83424]++
															return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:231
				// _ = "end of CoverTab[83424]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:232
				_go_fuzz_dep_.CoverTab[83425]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:232
				// _ = "end of CoverTab[83425]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:232
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:232
			// _ = "end of CoverTab[83372]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:232
			_go_fuzz_dep_.CoverTab[83373]++
														l.KDCTimeSync = int(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:233
			// _ = "end of CoverTab[83373]"
		case "noaddresses":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:234
			_go_fuzz_dep_.CoverTab[83374]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:236
				_go_fuzz_dep_.CoverTab[83426]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:237
				// _ = "end of CoverTab[83426]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:238
				_go_fuzz_dep_.CoverTab[83427]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:238
				// _ = "end of CoverTab[83427]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:238
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:238
			// _ = "end of CoverTab[83374]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:238
			_go_fuzz_dep_.CoverTab[83375]++
														l.NoAddresses = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:239
			// _ = "end of CoverTab[83375]"
		case "permitted_enctypes":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:240
			_go_fuzz_dep_.CoverTab[83376]++
														l.PermittedEnctypes = strings.Fields(p[1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:241
			// _ = "end of CoverTab[83376]"
		case "preferred_preauth_types":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:242
			_go_fuzz_dep_.CoverTab[83377]++
														p[1] = strings.TrimSpace(p[1])
														t := strings.Split(p[1], ",")
														var v []int
														for _, s := range t {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:246
				_go_fuzz_dep_.CoverTab[83428]++
															i, err := strconv.ParseInt(s, 10, 32)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:248
					_go_fuzz_dep_.CoverTab[83430]++
																return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:249
					// _ = "end of CoverTab[83430]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:250
					_go_fuzz_dep_.CoverTab[83431]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:250
					// _ = "end of CoverTab[83431]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:250
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:250
				// _ = "end of CoverTab[83428]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:250
				_go_fuzz_dep_.CoverTab[83429]++
															v = append(v, int(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:251
				// _ = "end of CoverTab[83429]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:252
			// _ = "end of CoverTab[83377]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:252
			_go_fuzz_dep_.CoverTab[83378]++
														l.PreferredPreauthTypes = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:253
			// _ = "end of CoverTab[83378]"
		case "proxiable":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:254
			_go_fuzz_dep_.CoverTab[83379]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:256
				_go_fuzz_dep_.CoverTab[83432]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:257
				// _ = "end of CoverTab[83432]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:258
				_go_fuzz_dep_.CoverTab[83433]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:258
				// _ = "end of CoverTab[83433]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:258
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:258
			// _ = "end of CoverTab[83379]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:258
			_go_fuzz_dep_.CoverTab[83380]++
														l.Proxiable = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:259
			// _ = "end of CoverTab[83380]"
		case "rdns":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:260
			_go_fuzz_dep_.CoverTab[83381]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:262
				_go_fuzz_dep_.CoverTab[83434]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:263
				// _ = "end of CoverTab[83434]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:264
				_go_fuzz_dep_.CoverTab[83435]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:264
				// _ = "end of CoverTab[83435]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:264
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:264
			// _ = "end of CoverTab[83381]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:264
			_go_fuzz_dep_.CoverTab[83382]++
														l.RDNS = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:265
			// _ = "end of CoverTab[83382]"
		case "realm_try_domains":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:266
			_go_fuzz_dep_.CoverTab[83383]++
														p[1] = strings.TrimSpace(p[1])
														v, err := strconv.ParseInt(p[1], 10, 32)
														if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:269
				_go_fuzz_dep_.CoverTab[83436]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:269
				return v < -1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:269
				// _ = "end of CoverTab[83436]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:269
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:269
				_go_fuzz_dep_.CoverTab[83437]++
															return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:270
				// _ = "end of CoverTab[83437]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:271
				_go_fuzz_dep_.CoverTab[83438]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:271
				// _ = "end of CoverTab[83438]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:271
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:271
			// _ = "end of CoverTab[83383]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:271
			_go_fuzz_dep_.CoverTab[83384]++
														l.RealmTryDomains = int(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:272
			// _ = "end of CoverTab[83384]"
		case "renew_lifetime":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:273
			_go_fuzz_dep_.CoverTab[83385]++
														d, err := parseDuration(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:275
				_go_fuzz_dep_.CoverTab[83439]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:276
				// _ = "end of CoverTab[83439]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:277
				_go_fuzz_dep_.CoverTab[83440]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:277
				// _ = "end of CoverTab[83440]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:277
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:277
			// _ = "end of CoverTab[83385]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:277
			_go_fuzz_dep_.CoverTab[83386]++
														l.RenewLifetime = d
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:278
			// _ = "end of CoverTab[83386]"
		case "safe_checksum_type":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:279
			_go_fuzz_dep_.CoverTab[83387]++
														p[1] = strings.TrimSpace(p[1])
														v, err := strconv.ParseInt(p[1], 10, 32)
														if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:282
				_go_fuzz_dep_.CoverTab[83441]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:282
				return v < 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:282
				// _ = "end of CoverTab[83441]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:282
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:282
				_go_fuzz_dep_.CoverTab[83442]++
															return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:283
				// _ = "end of CoverTab[83442]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:284
				_go_fuzz_dep_.CoverTab[83443]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:284
				// _ = "end of CoverTab[83443]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:284
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:284
			// _ = "end of CoverTab[83387]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:284
			_go_fuzz_dep_.CoverTab[83388]++
														l.SafeChecksumType = int(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:285
			// _ = "end of CoverTab[83388]"
		case "ticket_lifetime":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:286
			_go_fuzz_dep_.CoverTab[83389]++
														d, err := parseDuration(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:288
				_go_fuzz_dep_.CoverTab[83444]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:289
				// _ = "end of CoverTab[83444]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:290
				_go_fuzz_dep_.CoverTab[83445]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:290
				// _ = "end of CoverTab[83445]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:290
			// _ = "end of CoverTab[83389]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:290
			_go_fuzz_dep_.CoverTab[83390]++
														l.TicketLifetime = d
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:291
			// _ = "end of CoverTab[83390]"
		case "udp_preference_limit":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:292
			_go_fuzz_dep_.CoverTab[83391]++
														p[1] = strings.TrimSpace(p[1])
														v, err := strconv.ParseUint(p[1], 10, 32)
														if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:295
				_go_fuzz_dep_.CoverTab[83446]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:295
				return v > 32700
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:295
				// _ = "end of CoverTab[83446]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:295
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:295
				_go_fuzz_dep_.CoverTab[83447]++
															return InvalidErrorf("libdefaults section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:296
				// _ = "end of CoverTab[83447]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:297
				_go_fuzz_dep_.CoverTab[83448]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:297
				// _ = "end of CoverTab[83448]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:297
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:297
			// _ = "end of CoverTab[83391]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:297
			_go_fuzz_dep_.CoverTab[83392]++
														l.UDPPreferenceLimit = int(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:298
			// _ = "end of CoverTab[83392]"
		case "verify_ap_req_nofail":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:299
			_go_fuzz_dep_.CoverTab[83393]++
														v, err := parseBoolean(p[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:301
				_go_fuzz_dep_.CoverTab[83449]++
															return InvalidErrorf("libdefaults section line (%s): %v", line, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:302
				// _ = "end of CoverTab[83449]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:303
				_go_fuzz_dep_.CoverTab[83450]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:303
				// _ = "end of CoverTab[83450]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:303
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:303
			// _ = "end of CoverTab[83393]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:303
			_go_fuzz_dep_.CoverTab[83394]++
														l.VerifyAPReqNofail = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:304
			// _ = "end of CoverTab[83394]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:304
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:304
			_go_fuzz_dep_.CoverTab[83395]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:304
			// _ = "end of CoverTab[83395]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:305
		// _ = "end of CoverTab[83336]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:306
	// _ = "end of CoverTab[83331]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:306
	_go_fuzz_dep_.CoverTab[83332]++
												l.DefaultTGSEnctypeIDs = parseETypes(l.DefaultTGSEnctypes, l.AllowWeakCrypto)
												l.DefaultTktEnctypeIDs = parseETypes(l.DefaultTktEnctypes, l.AllowWeakCrypto)
												l.PermittedEnctypeIDs = parseETypes(l.PermittedEnctypes, l.AllowWeakCrypto)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:310
	// _ = "end of CoverTab[83332]"
}

// Realm represents an entry in the [realms] section of the configuration.
type Realm struct {
	Realm		string
	AdminServer	[]string
	//auth_to_local //Not implementing for now
	//auth_to_local_names //Not implementing for now
	DefaultDomain	string
	KDC		[]string
	KPasswdServer	[]string	//default admin_server:464
	MasterKDC	[]string
}

// Parse the lines of a [realms] entry into the Realm struct.
func (r *Realm) parseLines(name string, lines []string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:326
	_go_fuzz_dep_.CoverTab[83451]++
												r.Realm = name
												var adminServerFinal bool
												var KDCFinal bool
												var kpasswdServerFinal bool
												var masterKDCFinal bool
												var ignore bool
												var c int	// counts the depth of blocks within brackets { }
												for _, line := range lines {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:334
		_go_fuzz_dep_.CoverTab[83454]++
													if ignore && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			_go_fuzz_dep_.CoverTab[83462]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			return c > 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			// _ = "end of CoverTab[83462]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			_go_fuzz_dep_.CoverTab[83463]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			return !strings.Contains(line, "{")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			// _ = "end of CoverTab[83463]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			_go_fuzz_dep_.CoverTab[83464]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			return !strings.Contains(line, "}")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			// _ = "end of CoverTab[83464]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:335
			_go_fuzz_dep_.CoverTab[83465]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:336
			// _ = "end of CoverTab[83465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:337
			_go_fuzz_dep_.CoverTab[83466]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:337
			// _ = "end of CoverTab[83466]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:337
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:337
		// _ = "end of CoverTab[83454]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:337
		_go_fuzz_dep_.CoverTab[83455]++

													if idx := strings.IndexAny(line, "#;"); idx != -1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:339
			_go_fuzz_dep_.CoverTab[83467]++
														line = line[:idx]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:340
			// _ = "end of CoverTab[83467]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:341
			_go_fuzz_dep_.CoverTab[83468]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:341
			// _ = "end of CoverTab[83468]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:341
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:341
		// _ = "end of CoverTab[83455]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:341
		_go_fuzz_dep_.CoverTab[83456]++
													line = strings.TrimSpace(line)
													if line == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:343
			_go_fuzz_dep_.CoverTab[83469]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:344
			// _ = "end of CoverTab[83469]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:345
			_go_fuzz_dep_.CoverTab[83470]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:345
			// _ = "end of CoverTab[83470]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:345
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:345
		// _ = "end of CoverTab[83456]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:345
		_go_fuzz_dep_.CoverTab[83457]++
													if !strings.Contains(line, "=") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:346
			_go_fuzz_dep_.CoverTab[83471]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:346
			return !strings.Contains(line, "}")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:346
			// _ = "end of CoverTab[83471]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:346
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:346
			_go_fuzz_dep_.CoverTab[83472]++
														return InvalidErrorf("realms section line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:347
			// _ = "end of CoverTab[83472]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:348
			_go_fuzz_dep_.CoverTab[83473]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:348
			// _ = "end of CoverTab[83473]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:348
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:348
		// _ = "end of CoverTab[83457]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:348
		_go_fuzz_dep_.CoverTab[83458]++
													if strings.Contains(line, "v4_") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:349
			_go_fuzz_dep_.CoverTab[83474]++
														ignore = true
														err = UnsupportedDirective{"v4 configurations are not supported"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:351
			// _ = "end of CoverTab[83474]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:352
			_go_fuzz_dep_.CoverTab[83475]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:352
			// _ = "end of CoverTab[83475]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:352
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:352
		// _ = "end of CoverTab[83458]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:352
		_go_fuzz_dep_.CoverTab[83459]++
													if strings.Contains(line, "{") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:353
			_go_fuzz_dep_.CoverTab[83476]++
														c++
														if ignore {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:355
				_go_fuzz_dep_.CoverTab[83477]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:356
				// _ = "end of CoverTab[83477]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:357
				_go_fuzz_dep_.CoverTab[83478]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:357
				// _ = "end of CoverTab[83478]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:357
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:357
			// _ = "end of CoverTab[83476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:358
			_go_fuzz_dep_.CoverTab[83479]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:358
			// _ = "end of CoverTab[83479]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:358
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:358
		// _ = "end of CoverTab[83459]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:358
		_go_fuzz_dep_.CoverTab[83460]++
													if strings.Contains(line, "}") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:359
			_go_fuzz_dep_.CoverTab[83480]++
														c--
														if c < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:361
				_go_fuzz_dep_.CoverTab[83482]++
															return InvalidErrorf("unpaired curly brackets")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:362
				// _ = "end of CoverTab[83482]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:363
				_go_fuzz_dep_.CoverTab[83483]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:363
				// _ = "end of CoverTab[83483]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:363
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:363
			// _ = "end of CoverTab[83480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:363
			_go_fuzz_dep_.CoverTab[83481]++
														if ignore {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:364
				_go_fuzz_dep_.CoverTab[83484]++
															if c < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:365
					_go_fuzz_dep_.CoverTab[83486]++
																c = 0
																ignore = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:367
					// _ = "end of CoverTab[83486]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:368
					_go_fuzz_dep_.CoverTab[83487]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:368
					// _ = "end of CoverTab[83487]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:368
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:368
				// _ = "end of CoverTab[83484]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:368
				_go_fuzz_dep_.CoverTab[83485]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:369
				// _ = "end of CoverTab[83485]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:370
				_go_fuzz_dep_.CoverTab[83488]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:370
				// _ = "end of CoverTab[83488]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:370
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:370
			// _ = "end of CoverTab[83481]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:371
			_go_fuzz_dep_.CoverTab[83489]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:371
			// _ = "end of CoverTab[83489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:371
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:371
		// _ = "end of CoverTab[83460]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:371
		_go_fuzz_dep_.CoverTab[83461]++

													p := strings.Split(line, "=")
													key := strings.TrimSpace(strings.ToLower(p[0]))
													v := strings.TrimSpace(p[1])
													switch key {
		case "admin_server":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:377
			_go_fuzz_dep_.CoverTab[83490]++
														appendUntilFinal(&r.AdminServer, v, &adminServerFinal)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:378
			// _ = "end of CoverTab[83490]"
		case "default_domain":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:379
			_go_fuzz_dep_.CoverTab[83491]++
														r.DefaultDomain = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:380
			// _ = "end of CoverTab[83491]"
		case "kdc":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:381
			_go_fuzz_dep_.CoverTab[83492]++
														if !strings.Contains(v, ":") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:382
				_go_fuzz_dep_.CoverTab[83497]++

															if strings.HasSuffix(v, `*`) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:384
					_go_fuzz_dep_.CoverTab[83498]++
																v = strings.TrimSpace(strings.TrimSuffix(v, `*`)) + ":88*"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:385
					// _ = "end of CoverTab[83498]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:386
					_go_fuzz_dep_.CoverTab[83499]++
																v = strings.TrimSpace(v) + ":88"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:387
					// _ = "end of CoverTab[83499]"
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:388
				// _ = "end of CoverTab[83497]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:389
				_go_fuzz_dep_.CoverTab[83500]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:389
				// _ = "end of CoverTab[83500]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:389
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:389
			// _ = "end of CoverTab[83492]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:389
			_go_fuzz_dep_.CoverTab[83493]++
														appendUntilFinal(&r.KDC, v, &KDCFinal)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:390
			// _ = "end of CoverTab[83493]"
		case "kpasswd_server":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:391
			_go_fuzz_dep_.CoverTab[83494]++
														appendUntilFinal(&r.KPasswdServer, v, &kpasswdServerFinal)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:392
			// _ = "end of CoverTab[83494]"
		case "master_kdc":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:393
			_go_fuzz_dep_.CoverTab[83495]++
														appendUntilFinal(&r.MasterKDC, v, &masterKDCFinal)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:394
			// _ = "end of CoverTab[83495]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:394
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:394
			_go_fuzz_dep_.CoverTab[83496]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:394
			// _ = "end of CoverTab[83496]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:395
		// _ = "end of CoverTab[83461]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:396
	// _ = "end of CoverTab[83451]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:396
	_go_fuzz_dep_.CoverTab[83452]++

												if len(r.KPasswdServer) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:398
		_go_fuzz_dep_.CoverTab[83501]++
													for _, a := range r.AdminServer {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:399
			_go_fuzz_dep_.CoverTab[83502]++
														s := strings.Split(a, ":")
														r.KPasswdServer = append(r.KPasswdServer, s[0]+":464")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:401
			// _ = "end of CoverTab[83502]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:402
		// _ = "end of CoverTab[83501]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:403
		_go_fuzz_dep_.CoverTab[83503]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:403
		// _ = "end of CoverTab[83503]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:403
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:403
	// _ = "end of CoverTab[83452]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:403
	_go_fuzz_dep_.CoverTab[83453]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:404
	// _ = "end of CoverTab[83453]"
}

// Parse the lines of the [realms] section of the configuration into an slice of Realm structs.
func parseRealms(lines []string) (realms []Realm, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:408
	_go_fuzz_dep_.CoverTab[83504]++
												var name string
												var start int
												var c int
												for i, l := range lines {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:412
		_go_fuzz_dep_.CoverTab[83506]++

													if idx := strings.IndexAny(l, "#;"); idx != -1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:414
			_go_fuzz_dep_.CoverTab[83510]++
														l = l[:idx]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:415
			// _ = "end of CoverTab[83510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:416
			_go_fuzz_dep_.CoverTab[83511]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:416
			// _ = "end of CoverTab[83511]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:416
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:416
		// _ = "end of CoverTab[83506]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:416
		_go_fuzz_dep_.CoverTab[83507]++
													l = strings.TrimSpace(l)
													if l == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:418
			_go_fuzz_dep_.CoverTab[83512]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:419
			// _ = "end of CoverTab[83512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:420
			_go_fuzz_dep_.CoverTab[83513]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:420
			// _ = "end of CoverTab[83513]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:420
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:420
		// _ = "end of CoverTab[83507]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:420
		_go_fuzz_dep_.CoverTab[83508]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:424
		if strings.Contains(l, "{") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:424
			_go_fuzz_dep_.CoverTab[83514]++
														c++
														if !strings.Contains(l, "=") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:426
				_go_fuzz_dep_.CoverTab[83516]++
															return nil, fmt.Errorf("realm configuration line invalid: %s", l)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:427
				// _ = "end of CoverTab[83516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:428
				_go_fuzz_dep_.CoverTab[83517]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:428
				// _ = "end of CoverTab[83517]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:428
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:428
			// _ = "end of CoverTab[83514]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:428
			_go_fuzz_dep_.CoverTab[83515]++
														if c == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:429
				_go_fuzz_dep_.CoverTab[83518]++
															start = i
															p := strings.Split(l, "=")
															name = strings.TrimSpace(p[0])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:432
				// _ = "end of CoverTab[83518]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:433
				_go_fuzz_dep_.CoverTab[83519]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:433
				// _ = "end of CoverTab[83519]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:433
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:433
			// _ = "end of CoverTab[83515]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:434
			_go_fuzz_dep_.CoverTab[83520]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:434
			// _ = "end of CoverTab[83520]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:434
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:434
		// _ = "end of CoverTab[83508]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:434
		_go_fuzz_dep_.CoverTab[83509]++
													if strings.Contains(l, "}") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:435
			_go_fuzz_dep_.CoverTab[83521]++
														if c < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:436
				_go_fuzz_dep_.CoverTab[83523]++

															return nil, errors.New("invalid Realms section in configuration")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:438
				// _ = "end of CoverTab[83523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:439
				_go_fuzz_dep_.CoverTab[83524]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:439
				// _ = "end of CoverTab[83524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:439
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:439
			// _ = "end of CoverTab[83521]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:439
			_go_fuzz_dep_.CoverTab[83522]++
														c--
														if c == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:441
				_go_fuzz_dep_.CoverTab[83525]++
															var r Realm
															e := r.parseLines(name, lines[start+1:i])
															if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:444
					_go_fuzz_dep_.CoverTab[83527]++
																if _, ok := e.(UnsupportedDirective); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:445
						_go_fuzz_dep_.CoverTab[83529]++
																	err = e
																	return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:447
						// _ = "end of CoverTab[83529]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:448
						_go_fuzz_dep_.CoverTab[83530]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:448
						// _ = "end of CoverTab[83530]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:448
					}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:448
					// _ = "end of CoverTab[83527]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:448
					_go_fuzz_dep_.CoverTab[83528]++
																err = e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:449
					// _ = "end of CoverTab[83528]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:450
					_go_fuzz_dep_.CoverTab[83531]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:450
					// _ = "end of CoverTab[83531]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:450
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:450
				// _ = "end of CoverTab[83525]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:450
				_go_fuzz_dep_.CoverTab[83526]++
															realms = append(realms, r)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:451
				// _ = "end of CoverTab[83526]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:452
				_go_fuzz_dep_.CoverTab[83532]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:452
				// _ = "end of CoverTab[83532]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:452
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:452
			// _ = "end of CoverTab[83522]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:453
			_go_fuzz_dep_.CoverTab[83533]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:453
			// _ = "end of CoverTab[83533]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:453
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:453
		// _ = "end of CoverTab[83509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:454
	// _ = "end of CoverTab[83504]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:454
	_go_fuzz_dep_.CoverTab[83505]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:455
	// _ = "end of CoverTab[83505]"
}

// DomainRealm maps the domains to realms representing the [domain_realm] section of the configuration.
type DomainRealm map[string]string

// Parse the lines of the [domain_realm] section of the configuration and add to the mapping.
func (d *DomainRealm) parseLines(lines []string) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:462
	_go_fuzz_dep_.CoverTab[83534]++
												for _, line := range lines {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:463
		_go_fuzz_dep_.CoverTab[83536]++

													if idx := strings.IndexAny(line, "#;"); idx != -1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:465
			_go_fuzz_dep_.CoverTab[83540]++
														line = line[:idx]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:466
			// _ = "end of CoverTab[83540]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:467
			_go_fuzz_dep_.CoverTab[83541]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:467
			// _ = "end of CoverTab[83541]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:467
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:467
		// _ = "end of CoverTab[83536]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:467
		_go_fuzz_dep_.CoverTab[83537]++
													if strings.TrimSpace(line) == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:468
			_go_fuzz_dep_.CoverTab[83542]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:469
			// _ = "end of CoverTab[83542]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:470
			_go_fuzz_dep_.CoverTab[83543]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:470
			// _ = "end of CoverTab[83543]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:470
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:470
		// _ = "end of CoverTab[83537]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:470
		_go_fuzz_dep_.CoverTab[83538]++
													if !strings.Contains(line, "=") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:471
			_go_fuzz_dep_.CoverTab[83544]++
														return InvalidErrorf("realm line (%s)", line)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:472
			// _ = "end of CoverTab[83544]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:473
			_go_fuzz_dep_.CoverTab[83545]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:473
			// _ = "end of CoverTab[83545]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:473
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:473
		// _ = "end of CoverTab[83538]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:473
		_go_fuzz_dep_.CoverTab[83539]++
													p := strings.Split(line, "=")
													domain := strings.TrimSpace(strings.ToLower(p[0]))
													realm := strings.TrimSpace(p[1])
													d.addMapping(domain, realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:477
		// _ = "end of CoverTab[83539]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:478
	// _ = "end of CoverTab[83534]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:478
	_go_fuzz_dep_.CoverTab[83535]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:479
	// _ = "end of CoverTab[83535]"
}

// Add a domain to realm mapping.
func (d *DomainRealm) addMapping(domain, realm string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:483
	_go_fuzz_dep_.CoverTab[83546]++
												(*d)[domain] = realm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:484
	// _ = "end of CoverTab[83546]"
}

// Delete a domain to realm mapping.
func (d *DomainRealm) deleteMapping(domain, realm string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:488
	_go_fuzz_dep_.CoverTab[83547]++
												delete(*d, domain)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:489
	// _ = "end of CoverTab[83547]"
}

// ResolveRealm resolves the kerberos realm for the specified domain name from the domain to realm mapping.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:492
// The most specific mapping is returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:494
func (c *Config) ResolveRealm(domainName string) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:494
	_go_fuzz_dep_.CoverTab[83548]++
												domainName = strings.TrimSuffix(domainName, ".")

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:498
	if r, ok := c.DomainRealm[domainName]; ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:498
		_go_fuzz_dep_.CoverTab[83551]++
													return r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:499
		// _ = "end of CoverTab[83551]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:500
		_go_fuzz_dep_.CoverTab[83552]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:500
		// _ = "end of CoverTab[83552]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:500
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:500
	// _ = "end of CoverTab[83548]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:500
	_go_fuzz_dep_.CoverTab[83549]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:503
	periods := strings.Count(domainName, ".") + 1
	for i := 2; i <= periods; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:504
		_go_fuzz_dep_.CoverTab[83553]++
													z := strings.SplitN(domainName, ".", i)
													if r, ok := c.DomainRealm["."+z[len(z)-1]]; ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:506
			_go_fuzz_dep_.CoverTab[83554]++
														return r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:507
			// _ = "end of CoverTab[83554]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:508
			_go_fuzz_dep_.CoverTab[83555]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:508
			// _ = "end of CoverTab[83555]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:508
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:508
		// _ = "end of CoverTab[83553]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:509
	// _ = "end of CoverTab[83549]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:509
	_go_fuzz_dep_.CoverTab[83550]++
												return c.LibDefaults.DefaultRealm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:510
	// _ = "end of CoverTab[83550]"
}

// Load the KRB5 configuration from the specified file path.
func Load(cfgPath string) (*Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:514
	_go_fuzz_dep_.CoverTab[83556]++
												fh, err := os.Open(cfgPath)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:516
		_go_fuzz_dep_.CoverTab[83558]++
													return nil, errors.New("configuration file could not be opened: " + cfgPath + " " + err.Error())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:517
		// _ = "end of CoverTab[83558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:518
		_go_fuzz_dep_.CoverTab[83559]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:518
		// _ = "end of CoverTab[83559]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:518
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:518
	// _ = "end of CoverTab[83556]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:518
	_go_fuzz_dep_.CoverTab[83557]++
												defer fh.Close()
												scanner := bufio.NewScanner(fh)
												return NewFromScanner(scanner)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:521
	// _ = "end of CoverTab[83557]"
}

// NewFromString creates a new Config struct from a string.
func NewFromString(s string) (*Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:525
	_go_fuzz_dep_.CoverTab[83560]++
												reader := strings.NewReader(s)
												return NewFromReader(reader)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:527
	// _ = "end of CoverTab[83560]"
}

// NewFromReader creates a new Config struct from an io.Reader.
func NewFromReader(r io.Reader) (*Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:531
	_go_fuzz_dep_.CoverTab[83561]++
												scanner := bufio.NewScanner(r)
												return NewFromScanner(scanner)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:533
	// _ = "end of CoverTab[83561]"
}

// NewFromScanner creates a new Config struct from a bufio.Scanner.
func NewFromScanner(scanner *bufio.Scanner) (*Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:537
	_go_fuzz_dep_.CoverTab[83562]++
												c := New()
												var e error
												sections := make(map[int]string)
												var sectionLineNum []int
												var lines []string
												for scanner.Scan() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:543
		_go_fuzz_dep_.CoverTab[83565]++

													if matched, _ := regexp.MatchString(`^\s*(#|;|\n)`, scanner.Text()); matched {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:545
			_go_fuzz_dep_.CoverTab[83571]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:546
			// _ = "end of CoverTab[83571]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:547
			_go_fuzz_dep_.CoverTab[83572]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:547
			// _ = "end of CoverTab[83572]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:547
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:547
		// _ = "end of CoverTab[83565]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:547
		_go_fuzz_dep_.CoverTab[83566]++
													if matched, _ := regexp.MatchString(`^\s*\[libdefaults\]\s*`, scanner.Text()); matched {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:548
			_go_fuzz_dep_.CoverTab[83573]++
														sections[len(lines)] = "libdefaults"
														sectionLineNum = append(sectionLineNum, len(lines))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:551
			// _ = "end of CoverTab[83573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:552
			_go_fuzz_dep_.CoverTab[83574]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:552
			// _ = "end of CoverTab[83574]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:552
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:552
		// _ = "end of CoverTab[83566]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:552
		_go_fuzz_dep_.CoverTab[83567]++
													if matched, _ := regexp.MatchString(`^\s*\[realms\]\s*`, scanner.Text()); matched {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:553
			_go_fuzz_dep_.CoverTab[83575]++
														sections[len(lines)] = "realms"
														sectionLineNum = append(sectionLineNum, len(lines))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:556
			// _ = "end of CoverTab[83575]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:557
			_go_fuzz_dep_.CoverTab[83576]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:557
			// _ = "end of CoverTab[83576]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:557
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:557
		// _ = "end of CoverTab[83567]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:557
		_go_fuzz_dep_.CoverTab[83568]++
													if matched, _ := regexp.MatchString(`^\s*\[domain_realm\]\s*`, scanner.Text()); matched {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:558
			_go_fuzz_dep_.CoverTab[83577]++
														sections[len(lines)] = "domain_realm"
														sectionLineNum = append(sectionLineNum, len(lines))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:561
			// _ = "end of CoverTab[83577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:562
			_go_fuzz_dep_.CoverTab[83578]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:562
			// _ = "end of CoverTab[83578]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:562
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:562
		// _ = "end of CoverTab[83568]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:562
		_go_fuzz_dep_.CoverTab[83569]++
													if matched, _ := regexp.MatchString(`^\s*\[.*\]\s*`, scanner.Text()); matched {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:563
			_go_fuzz_dep_.CoverTab[83579]++
														sections[len(lines)] = "unknown_section"
														sectionLineNum = append(sectionLineNum, len(lines))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:566
			// _ = "end of CoverTab[83579]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:567
			_go_fuzz_dep_.CoverTab[83580]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:567
			// _ = "end of CoverTab[83580]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:567
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:567
		// _ = "end of CoverTab[83569]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:567
		_go_fuzz_dep_.CoverTab[83570]++
													lines = append(lines, scanner.Text())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:568
		// _ = "end of CoverTab[83570]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:569
	// _ = "end of CoverTab[83562]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:569
	_go_fuzz_dep_.CoverTab[83563]++
												for i, start := range sectionLineNum {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:570
		_go_fuzz_dep_.CoverTab[83581]++
													var end int
													if i+1 >= len(sectionLineNum) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:572
			_go_fuzz_dep_.CoverTab[83583]++
														end = len(lines)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:573
			// _ = "end of CoverTab[83583]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:574
			_go_fuzz_dep_.CoverTab[83584]++
														end = sectionLineNum[i+1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:575
			// _ = "end of CoverTab[83584]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:576
		// _ = "end of CoverTab[83581]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:576
		_go_fuzz_dep_.CoverTab[83582]++
													switch section := sections[start]; section {
		case "libdefaults":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:578
			_go_fuzz_dep_.CoverTab[83585]++
														err := c.LibDefaults.parseLines(lines[start:end])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:580
				_go_fuzz_dep_.CoverTab[83590]++
															if _, ok := err.(UnsupportedDirective); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:581
					_go_fuzz_dep_.CoverTab[83592]++
																return nil, fmt.Errorf("error processing libdefaults section: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:582
					// _ = "end of CoverTab[83592]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:583
					_go_fuzz_dep_.CoverTab[83593]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:583
					// _ = "end of CoverTab[83593]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:583
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:583
				// _ = "end of CoverTab[83590]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:583
				_go_fuzz_dep_.CoverTab[83591]++
															e = err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:584
				// _ = "end of CoverTab[83591]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:585
				_go_fuzz_dep_.CoverTab[83594]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:585
				// _ = "end of CoverTab[83594]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:585
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:585
			// _ = "end of CoverTab[83585]"
		case "realms":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:586
			_go_fuzz_dep_.CoverTab[83586]++
														realms, err := parseRealms(lines[start:end])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:588
				_go_fuzz_dep_.CoverTab[83595]++
															if _, ok := err.(UnsupportedDirective); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:589
					_go_fuzz_dep_.CoverTab[83597]++
																return nil, fmt.Errorf("error processing realms section: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:590
					// _ = "end of CoverTab[83597]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:591
					_go_fuzz_dep_.CoverTab[83598]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:591
					// _ = "end of CoverTab[83598]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:591
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:591
				// _ = "end of CoverTab[83595]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:591
				_go_fuzz_dep_.CoverTab[83596]++
															e = err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:592
				// _ = "end of CoverTab[83596]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:593
				_go_fuzz_dep_.CoverTab[83599]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:593
				// _ = "end of CoverTab[83599]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:593
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:593
			// _ = "end of CoverTab[83586]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:593
			_go_fuzz_dep_.CoverTab[83587]++
														c.Realms = realms
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:594
			// _ = "end of CoverTab[83587]"
		case "domain_realm":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:595
			_go_fuzz_dep_.CoverTab[83588]++
														err := c.DomainRealm.parseLines(lines[start:end])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:597
				_go_fuzz_dep_.CoverTab[83600]++
															if _, ok := err.(UnsupportedDirective); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:598
					_go_fuzz_dep_.CoverTab[83602]++
																return nil, fmt.Errorf("error processing domaain_realm section: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:599
					// _ = "end of CoverTab[83602]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:600
					_go_fuzz_dep_.CoverTab[83603]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:600
					// _ = "end of CoverTab[83603]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:600
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:600
				// _ = "end of CoverTab[83600]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:600
				_go_fuzz_dep_.CoverTab[83601]++
															e = err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:601
				// _ = "end of CoverTab[83601]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
				_go_fuzz_dep_.CoverTab[83604]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
				// _ = "end of CoverTab[83604]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
			// _ = "end of CoverTab[83588]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
			_go_fuzz_dep_.CoverTab[83589]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:602
			// _ = "end of CoverTab[83589]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:603
		// _ = "end of CoverTab[83582]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:604
	// _ = "end of CoverTab[83563]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:604
	_go_fuzz_dep_.CoverTab[83564]++
												return c, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:605
	// _ = "end of CoverTab[83564]"
}

// Parse a space delimited list of ETypes into a list of EType numbers optionally filtering out weak ETypes.
func parseETypes(s []string, w bool) []int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:609
	_go_fuzz_dep_.CoverTab[83605]++
												var eti []int32
												for _, et := range s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:611
		_go_fuzz_dep_.CoverTab[83607]++
													if !w {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:612
			_go_fuzz_dep_.CoverTab[83609]++
														var weak bool
														for _, wet := range strings.Fields(WeakETypeList) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:614
				_go_fuzz_dep_.CoverTab[83611]++
															if et == wet {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:615
					_go_fuzz_dep_.CoverTab[83612]++
																weak = true
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:617
					// _ = "end of CoverTab[83612]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:618
					_go_fuzz_dep_.CoverTab[83613]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:618
					// _ = "end of CoverTab[83613]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:618
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:618
				// _ = "end of CoverTab[83611]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:619
			// _ = "end of CoverTab[83609]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:619
			_go_fuzz_dep_.CoverTab[83610]++
														if weak {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:620
				_go_fuzz_dep_.CoverTab[83614]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:621
				// _ = "end of CoverTab[83614]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:622
				_go_fuzz_dep_.CoverTab[83615]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:622
				// _ = "end of CoverTab[83615]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:622
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:622
			// _ = "end of CoverTab[83610]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:623
			_go_fuzz_dep_.CoverTab[83616]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:623
			// _ = "end of CoverTab[83616]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:623
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:623
		// _ = "end of CoverTab[83607]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:623
		_go_fuzz_dep_.CoverTab[83608]++
													i := etypeID.EtypeSupported(et)
													if i != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:625
			_go_fuzz_dep_.CoverTab[83617]++
														eti = append(eti, i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:626
			// _ = "end of CoverTab[83617]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:627
			_go_fuzz_dep_.CoverTab[83618]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:627
			// _ = "end of CoverTab[83618]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:627
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:627
		// _ = "end of CoverTab[83608]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:628
	// _ = "end of CoverTab[83605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:628
	_go_fuzz_dep_.CoverTab[83606]++
												return eti
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:629
	// _ = "end of CoverTab[83606]"
}

// Parse a time duration string in the configuration to a golang time.Duration.
func parseDuration(s string) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:633
	_go_fuzz_dep_.CoverTab[83619]++
												s = strings.Replace(strings.TrimSpace(s), " ", "", -1)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:637
	if strings.Contains(s, "d") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:637
		_go_fuzz_dep_.CoverTab[83624]++
													ds := strings.SplitN(s, "d", 2)
													dn, err := strconv.ParseUint(ds[0], 10, 32)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:640
			_go_fuzz_dep_.CoverTab[83627]++
														return time.Duration(0), errors.New("invalid time duration")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:641
			// _ = "end of CoverTab[83627]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:642
			_go_fuzz_dep_.CoverTab[83628]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:642
			// _ = "end of CoverTab[83628]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:642
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:642
		// _ = "end of CoverTab[83624]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:642
		_go_fuzz_dep_.CoverTab[83625]++
													d := time.Duration(dn*24) * time.Hour
													if ds[1] != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:644
			_go_fuzz_dep_.CoverTab[83629]++
														dp, err := time.ParseDuration(ds[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:646
				_go_fuzz_dep_.CoverTab[83631]++
															return time.Duration(0), errors.New("invalid time duration")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:647
				// _ = "end of CoverTab[83631]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:648
				_go_fuzz_dep_.CoverTab[83632]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:648
				// _ = "end of CoverTab[83632]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:648
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:648
			// _ = "end of CoverTab[83629]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:648
			_go_fuzz_dep_.CoverTab[83630]++
														d = d + dp
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:649
			// _ = "end of CoverTab[83630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:650
			_go_fuzz_dep_.CoverTab[83633]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:650
			// _ = "end of CoverTab[83633]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:650
		// _ = "end of CoverTab[83625]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:650
		_go_fuzz_dep_.CoverTab[83626]++
													return d, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:651
		// _ = "end of CoverTab[83626]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:652
		_go_fuzz_dep_.CoverTab[83634]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:652
		// _ = "end of CoverTab[83634]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:652
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:652
	// _ = "end of CoverTab[83619]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:652
	_go_fuzz_dep_.CoverTab[83620]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:655
	d, err := time.ParseDuration(s)
	if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:656
		_go_fuzz_dep_.CoverTab[83635]++
													return d, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:657
		// _ = "end of CoverTab[83635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:658
		_go_fuzz_dep_.CoverTab[83636]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:658
		// _ = "end of CoverTab[83636]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:658
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:658
	// _ = "end of CoverTab[83620]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:658
	_go_fuzz_dep_.CoverTab[83621]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:661
	v, err := strconv.ParseUint(s, 10, 32)
	if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:662
		_go_fuzz_dep_.CoverTab[83637]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:662
		return v > 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:662
		// _ = "end of CoverTab[83637]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:662
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:662
		_go_fuzz_dep_.CoverTab[83638]++
													return time.Duration(v) * time.Second, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:663
		// _ = "end of CoverTab[83638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:664
		_go_fuzz_dep_.CoverTab[83639]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:664
		// _ = "end of CoverTab[83639]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:664
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:664
	// _ = "end of CoverTab[83621]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:664
	_go_fuzz_dep_.CoverTab[83622]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:667
	if strings.Contains(s, ":") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:667
		_go_fuzz_dep_.CoverTab[83640]++
													t := strings.Split(s, ":")
													if 2 > len(t) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:669
			_go_fuzz_dep_.CoverTab[83644]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:669
			return len(t) > 3
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:669
			// _ = "end of CoverTab[83644]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:669
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:669
			_go_fuzz_dep_.CoverTab[83645]++
														return time.Duration(0), errors.New("invalid time duration value")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:670
			// _ = "end of CoverTab[83645]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:671
			_go_fuzz_dep_.CoverTab[83646]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:671
			// _ = "end of CoverTab[83646]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:671
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:671
		// _ = "end of CoverTab[83640]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:671
		_go_fuzz_dep_.CoverTab[83641]++
													var i []int
													for _, n := range t {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:673
			_go_fuzz_dep_.CoverTab[83647]++
														j, err := strconv.ParseInt(n, 10, 16)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:675
				_go_fuzz_dep_.CoverTab[83649]++
															return time.Duration(0), errors.New("invalid time duration value")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:676
				// _ = "end of CoverTab[83649]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:677
				_go_fuzz_dep_.CoverTab[83650]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:677
				// _ = "end of CoverTab[83650]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:677
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:677
			// _ = "end of CoverTab[83647]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:677
			_go_fuzz_dep_.CoverTab[83648]++
														i = append(i, int(j))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:678
			// _ = "end of CoverTab[83648]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:679
		// _ = "end of CoverTab[83641]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:679
		_go_fuzz_dep_.CoverTab[83642]++
													d := time.Duration(i[0])*time.Hour + time.Duration(i[1])*time.Minute
													if len(i) == 3 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:681
			_go_fuzz_dep_.CoverTab[83651]++
														d = d + time.Duration(i[2])*time.Second
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:682
			// _ = "end of CoverTab[83651]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:683
			_go_fuzz_dep_.CoverTab[83652]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:683
			// _ = "end of CoverTab[83652]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:683
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:683
		// _ = "end of CoverTab[83642]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:683
		_go_fuzz_dep_.CoverTab[83643]++
													return d, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:684
		// _ = "end of CoverTab[83643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:685
		_go_fuzz_dep_.CoverTab[83653]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:685
		// _ = "end of CoverTab[83653]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:685
	// _ = "end of CoverTab[83622]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:685
	_go_fuzz_dep_.CoverTab[83623]++
												return time.Duration(0), errors.New("invalid time duration value")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:686
	// _ = "end of CoverTab[83623]"
}

// Parse possible boolean values to golang bool.
func parseBoolean(s string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:690
	_go_fuzz_dep_.CoverTab[83654]++
												s = strings.TrimSpace(s)
												v, err := strconv.ParseBool(s)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:693
		_go_fuzz_dep_.CoverTab[83657]++
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:694
		// _ = "end of CoverTab[83657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:695
		_go_fuzz_dep_.CoverTab[83658]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:695
		// _ = "end of CoverTab[83658]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:695
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:695
	// _ = "end of CoverTab[83654]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:695
	_go_fuzz_dep_.CoverTab[83655]++
												switch strings.ToLower(s) {
	case "yes":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:697
		_go_fuzz_dep_.CoverTab[83659]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:698
		// _ = "end of CoverTab[83659]"
	case "y":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:699
		_go_fuzz_dep_.CoverTab[83660]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:700
		// _ = "end of CoverTab[83660]"
	case "no":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:701
		_go_fuzz_dep_.CoverTab[83661]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:702
		// _ = "end of CoverTab[83661]"
	case "n":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:703
		_go_fuzz_dep_.CoverTab[83662]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:704
		// _ = "end of CoverTab[83662]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:704
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:704
		_go_fuzz_dep_.CoverTab[83663]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:704
		// _ = "end of CoverTab[83663]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:705
	// _ = "end of CoverTab[83655]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:705
	_go_fuzz_dep_.CoverTab[83656]++
												return false, errors.New("invalid boolean value")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:706
	// _ = "end of CoverTab[83656]"
}

// Parse array of strings but stop if an asterisk is placed at the end of a line.
func appendUntilFinal(s *[]string, value string, final *bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:710
	_go_fuzz_dep_.CoverTab[83664]++
												if *final {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:711
		_go_fuzz_dep_.CoverTab[83667]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:712
		// _ = "end of CoverTab[83667]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:713
		_go_fuzz_dep_.CoverTab[83668]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:713
		// _ = "end of CoverTab[83668]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:713
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:713
	// _ = "end of CoverTab[83664]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:713
	_go_fuzz_dep_.CoverTab[83665]++
												if last := len(value) - 1; last >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:714
		_go_fuzz_dep_.CoverTab[83669]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:714
		return value[last] == '*'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:714
		// _ = "end of CoverTab[83669]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:714
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:714
		_go_fuzz_dep_.CoverTab[83670]++
													*final = true
													value = value[:len(value)-1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:716
		// _ = "end of CoverTab[83670]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:717
		_go_fuzz_dep_.CoverTab[83671]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:717
		// _ = "end of CoverTab[83671]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:717
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:717
	// _ = "end of CoverTab[83665]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:717
	_go_fuzz_dep_.CoverTab[83666]++
												*s = append(*s, value)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:718
	// _ = "end of CoverTab[83666]"
}

// JSON return details of the config in a JSON format.
func (c *Config) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:722
	_go_fuzz_dep_.CoverTab[83672]++
												b, err := json.MarshalIndent(c, "", "  ")
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:724
		_go_fuzz_dep_.CoverTab[83674]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:725
		// _ = "end of CoverTab[83674]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:726
		_go_fuzz_dep_.CoverTab[83675]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:726
		// _ = "end of CoverTab[83675]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:726
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:726
	// _ = "end of CoverTab[83672]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:726
	_go_fuzz_dep_.CoverTab[83673]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:727
	// _ = "end of CoverTab[83673]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:728
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/krb5conf.go:728
var _ = _go_fuzz_dep_.CoverTab
