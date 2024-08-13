//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
package credentials

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:1
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"strings"
	"time"
	"unsafe"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/types"
)

const (
	headerFieldTagKDCOffset = 1
)

// CCache is the file credentials cache as define here: https://web.mit.edu/kerberos/krb5-latest/doc/formats/ccache_file_format.html
type CCache struct {
	Version			uint8
	Header			header
	DefaultPrincipal	principal
	Credentials		[]*Credential
	Path			string
}

type header struct {
	length	uint16
	fields	[]headerField
}

type headerField struct {
	tag	uint16
	length	uint16
	value	[]byte
}

// Credential cache entry principal struct.
type principal struct {
	Realm		string
	PrincipalName	types.PrincipalName
}

// Credential holds a Kerberos client's ccache credential information.
type Credential struct {
	Client		principal
	Server		principal
	Key		types.EncryptionKey
	AuthTime	time.Time
	StartTime	time.Time
	EndTime		time.Time
	RenewTill	time.Time
	IsSKey		bool
	TicketFlags	asn1.BitString
	Addresses	[]types.HostAddress
	AuthData	[]types.AuthorizationDataEntry
	Ticket		[]byte
	SecondTicket	[]byte
}

// LoadCCache loads a credential cache file into a CCache type.
func LoadCCache(cpath string) (*CCache, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:64
	_go_fuzz_dep_.CoverTab[86542]++
													c := new(CCache)
													b, err := ioutil.ReadFile(cpath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:67
		_go_fuzz_dep_.CoverTab[86544]++
														return c, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:68
		// _ = "end of CoverTab[86544]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:69
		_go_fuzz_dep_.CoverTab[86545]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:69
		// _ = "end of CoverTab[86545]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:69
	// _ = "end of CoverTab[86542]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:69
	_go_fuzz_dep_.CoverTab[86543]++
													err = c.Unmarshal(b)
													return c, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:71
	// _ = "end of CoverTab[86543]"
}

// Unmarshal a byte slice of credential cache data into CCache type.
func (c *CCache) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:75
	_go_fuzz_dep_.CoverTab[86546]++
													p := 0

													if int8(b[p]) != 5 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:78
		_go_fuzz_dep_.CoverTab[86552]++
														return errors.New("Invalid credential cache data. First byte does not equal 5")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:79
		// _ = "end of CoverTab[86552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:80
		_go_fuzz_dep_.CoverTab[86553]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:80
		// _ = "end of CoverTab[86553]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:80
	// _ = "end of CoverTab[86546]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:80
	_go_fuzz_dep_.CoverTab[86547]++
													p++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:84
	c.Version = b[p]
	if c.Version < 1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:85
		_go_fuzz_dep_.CoverTab[86554]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:85
		return c.Version > 4
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:85
		// _ = "end of CoverTab[86554]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:85
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:85
		_go_fuzz_dep_.CoverTab[86555]++
														return errors.New("Invalid credential cache data. Keytab version is not within 1 to 4")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:86
		// _ = "end of CoverTab[86555]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:87
		_go_fuzz_dep_.CoverTab[86556]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:87
		// _ = "end of CoverTab[86556]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:87
	// _ = "end of CoverTab[86547]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:87
	_go_fuzz_dep_.CoverTab[86548]++
													p++
	//Version 1 or 2 of the file format uses native byte order for integer representations. Versions 3 & 4 always uses big-endian byte order
	var endian binary.ByteOrder
	endian = binary.BigEndian
	if (c.Version == 1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		_go_fuzz_dep_.CoverTab[86557]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		return c.Version == 2
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		// _ = "end of CoverTab[86557]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		_go_fuzz_dep_.CoverTab[86558]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		return isNativeEndianLittle()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		// _ = "end of CoverTab[86558]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:92
		_go_fuzz_dep_.CoverTab[86559]++
														endian = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:93
		// _ = "end of CoverTab[86559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:94
		_go_fuzz_dep_.CoverTab[86560]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:94
		// _ = "end of CoverTab[86560]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:94
	// _ = "end of CoverTab[86548]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:94
	_go_fuzz_dep_.CoverTab[86549]++
													if c.Version == 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:95
		_go_fuzz_dep_.CoverTab[86561]++
														err := parseHeader(b, &p, c, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:97
			_go_fuzz_dep_.CoverTab[86562]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:98
			// _ = "end of CoverTab[86562]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:99
			_go_fuzz_dep_.CoverTab[86563]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:99
			// _ = "end of CoverTab[86563]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:99
		// _ = "end of CoverTab[86561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:100
		_go_fuzz_dep_.CoverTab[86564]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:100
		// _ = "end of CoverTab[86564]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:100
	// _ = "end of CoverTab[86549]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:100
	_go_fuzz_dep_.CoverTab[86550]++
													c.DefaultPrincipal = parsePrincipal(b, &p, c, &endian)
													for p < len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:102
		_go_fuzz_dep_.CoverTab[86565]++
														cred, err := parseCredential(b, &p, c, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:104
			_go_fuzz_dep_.CoverTab[86567]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:105
			// _ = "end of CoverTab[86567]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:106
			_go_fuzz_dep_.CoverTab[86568]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:106
			// _ = "end of CoverTab[86568]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:106
		// _ = "end of CoverTab[86565]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:106
		_go_fuzz_dep_.CoverTab[86566]++
														c.Credentials = append(c.Credentials, cred)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:107
		// _ = "end of CoverTab[86566]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:108
	// _ = "end of CoverTab[86550]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:108
	_go_fuzz_dep_.CoverTab[86551]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:109
	// _ = "end of CoverTab[86551]"
}

func parseHeader(b []byte, p *int, c *CCache, e *binary.ByteOrder) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:112
	_go_fuzz_dep_.CoverTab[86569]++
													if c.Version != 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:113
		_go_fuzz_dep_.CoverTab[86572]++
														return errors.New("Credentials cache version is not 4 so there is no header to parse.")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:114
		// _ = "end of CoverTab[86572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:115
		_go_fuzz_dep_.CoverTab[86573]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:115
		// _ = "end of CoverTab[86573]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:115
	// _ = "end of CoverTab[86569]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:115
	_go_fuzz_dep_.CoverTab[86570]++
													h := header{}
													h.length = uint16(readInt16(b, p, e))
													for *p <= int(h.length) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:118
		_go_fuzz_dep_.CoverTab[86574]++
														f := headerField{}
														f.tag = uint16(readInt16(b, p, e))
														f.length = uint16(readInt16(b, p, e))
														f.value = b[*p : *p+int(f.length)]
														*p += int(f.length)
														if !f.valid() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:124
			_go_fuzz_dep_.CoverTab[86576]++
															return errors.New("Invalid credential cache header found")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:125
			// _ = "end of CoverTab[86576]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:126
			_go_fuzz_dep_.CoverTab[86577]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:126
			// _ = "end of CoverTab[86577]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:126
		// _ = "end of CoverTab[86574]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:126
		_go_fuzz_dep_.CoverTab[86575]++
														h.fields = append(h.fields, f)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:127
		// _ = "end of CoverTab[86575]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:128
	// _ = "end of CoverTab[86570]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:128
	_go_fuzz_dep_.CoverTab[86571]++
													c.Header = h
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:130
	// _ = "end of CoverTab[86571]"
}

// Parse the Keytab bytes of a principal into a Keytab entry's principal.
func parsePrincipal(b []byte, p *int, c *CCache, e *binary.ByteOrder) (princ principal) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:134
	_go_fuzz_dep_.CoverTab[86578]++
													if c.Version != 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:135
		_go_fuzz_dep_.CoverTab[86582]++

														princ.PrincipalName.NameType = readInt32(b, p, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:137
		// _ = "end of CoverTab[86582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:138
		_go_fuzz_dep_.CoverTab[86583]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:138
		// _ = "end of CoverTab[86583]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:138
	// _ = "end of CoverTab[86578]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:138
	_go_fuzz_dep_.CoverTab[86579]++
													nc := int(readInt32(b, p, e))
													if c.Version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:140
		_go_fuzz_dep_.CoverTab[86584]++

														nc--
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:142
		// _ = "end of CoverTab[86584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:143
		_go_fuzz_dep_.CoverTab[86585]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:143
		// _ = "end of CoverTab[86585]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:143
	// _ = "end of CoverTab[86579]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:143
	_go_fuzz_dep_.CoverTab[86580]++
													lenRealm := readInt32(b, p, e)
													princ.Realm = string(readBytes(b, p, int(lenRealm), e))
													for i := 0; i < nc; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:146
		_go_fuzz_dep_.CoverTab[86586]++
														l := readInt32(b, p, e)
														princ.PrincipalName.NameString = append(princ.PrincipalName.NameString, string(readBytes(b, p, int(l), e)))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:148
		// _ = "end of CoverTab[86586]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:149
	// _ = "end of CoverTab[86580]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:149
	_go_fuzz_dep_.CoverTab[86581]++
													return princ
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:150
	// _ = "end of CoverTab[86581]"
}

func parseCredential(b []byte, p *int, c *CCache, e *binary.ByteOrder) (cred *Credential, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:153
	_go_fuzz_dep_.CoverTab[86587]++
													cred = new(Credential)
													cred.Client = parsePrincipal(b, p, c, e)
													cred.Server = parsePrincipal(b, p, c, e)
													key := types.EncryptionKey{}
													key.KeyType = int32(readInt16(b, p, e))
													if c.Version == 3 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:159
		_go_fuzz_dep_.CoverTab[86592]++

														key.KeyType = int32(readInt16(b, p, e))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:161
		// _ = "end of CoverTab[86592]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:162
		_go_fuzz_dep_.CoverTab[86593]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:162
		// _ = "end of CoverTab[86593]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:162
	// _ = "end of CoverTab[86587]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:162
	_go_fuzz_dep_.CoverTab[86588]++
													key.KeyValue = readData(b, p, e)
													cred.Key = key
													cred.AuthTime = readTimestamp(b, p, e)
													cred.StartTime = readTimestamp(b, p, e)
													cred.EndTime = readTimestamp(b, p, e)
													cred.RenewTill = readTimestamp(b, p, e)
													if ik := readInt8(b, p, e); ik == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:169
		_go_fuzz_dep_.CoverTab[86594]++
														cred.IsSKey = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:170
		// _ = "end of CoverTab[86594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:171
		_go_fuzz_dep_.CoverTab[86595]++
														cred.IsSKey = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:172
		// _ = "end of CoverTab[86595]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:173
	// _ = "end of CoverTab[86588]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:173
	_go_fuzz_dep_.CoverTab[86589]++
													cred.TicketFlags = types.NewKrbFlags()
													cred.TicketFlags.Bytes = readBytes(b, p, 4, e)
													l := int(readInt32(b, p, e))
													cred.Addresses = make([]types.HostAddress, l, l)
													for i := range cred.Addresses {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:178
		_go_fuzz_dep_.CoverTab[86596]++
														cred.Addresses[i] = readAddress(b, p, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:179
		// _ = "end of CoverTab[86596]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:180
	// _ = "end of CoverTab[86589]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:180
	_go_fuzz_dep_.CoverTab[86590]++
													l = int(readInt32(b, p, e))
													cred.AuthData = make([]types.AuthorizationDataEntry, l, l)
													for i := range cred.AuthData {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:183
		_go_fuzz_dep_.CoverTab[86597]++
														cred.AuthData[i] = readAuthDataEntry(b, p, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:184
		// _ = "end of CoverTab[86597]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:185
	// _ = "end of CoverTab[86590]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:185
	_go_fuzz_dep_.CoverTab[86591]++
													cred.Ticket = readData(b, p, e)
													cred.SecondTicket = readData(b, p, e)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:188
	// _ = "end of CoverTab[86591]"
}

// GetClientPrincipalName returns a PrincipalName type for the client the credentials cache is for.
func (c *CCache) GetClientPrincipalName() types.PrincipalName {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:192
	_go_fuzz_dep_.CoverTab[86598]++
													return c.DefaultPrincipal.PrincipalName
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:193
	// _ = "end of CoverTab[86598]"
}

// GetClientRealm returns the reals of the client the credentials cache is for.
func (c *CCache) GetClientRealm() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:197
	_go_fuzz_dep_.CoverTab[86599]++
													return c.DefaultPrincipal.Realm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:198
	// _ = "end of CoverTab[86599]"
}

// GetClientCredentials returns a Credentials object representing the client of the credentials cache.
func (c *CCache) GetClientCredentials() *Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:202
	_go_fuzz_dep_.CoverTab[86600]++
													return &Credentials{
		username:	c.DefaultPrincipal.PrincipalName.PrincipalNameString(),
		realm:		c.GetClientRealm(),
		cname:		c.DefaultPrincipal.PrincipalName,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:207
	// _ = "end of CoverTab[86600]"
}

// Contains tests if the cache contains a credential for the provided server PrincipalName
func (c *CCache) Contains(p types.PrincipalName) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:211
	_go_fuzz_dep_.CoverTab[86601]++
													for _, cred := range c.Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:212
		_go_fuzz_dep_.CoverTab[86603]++
														if cred.Server.PrincipalName.Equal(p) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:213
			_go_fuzz_dep_.CoverTab[86604]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:214
			// _ = "end of CoverTab[86604]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:215
			_go_fuzz_dep_.CoverTab[86605]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:215
			// _ = "end of CoverTab[86605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:215
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:215
		// _ = "end of CoverTab[86603]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:216
	// _ = "end of CoverTab[86601]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:216
	_go_fuzz_dep_.CoverTab[86602]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:217
	// _ = "end of CoverTab[86602]"
}

// GetEntry returns a specific credential for the PrincipalName provided.
func (c *CCache) GetEntry(p types.PrincipalName) (*Credential, bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:221
	_go_fuzz_dep_.CoverTab[86606]++
													cred := new(Credential)
													var found bool
													for i := range c.Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:224
		_go_fuzz_dep_.CoverTab[86609]++
														if c.Credentials[i].Server.PrincipalName.Equal(p) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:225
			_go_fuzz_dep_.CoverTab[86610]++
															cred = c.Credentials[i]
															found = true
															break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:228
			// _ = "end of CoverTab[86610]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:229
			_go_fuzz_dep_.CoverTab[86611]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:229
			// _ = "end of CoverTab[86611]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:229
		// _ = "end of CoverTab[86609]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:230
	// _ = "end of CoverTab[86606]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:230
	_go_fuzz_dep_.CoverTab[86607]++
													if !found {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:231
		_go_fuzz_dep_.CoverTab[86612]++
														return cred, false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:232
		// _ = "end of CoverTab[86612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:233
		_go_fuzz_dep_.CoverTab[86613]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:233
		// _ = "end of CoverTab[86613]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:233
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:233
	// _ = "end of CoverTab[86607]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:233
	_go_fuzz_dep_.CoverTab[86608]++
													return cred, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:234
	// _ = "end of CoverTab[86608]"
}

// GetEntries filters out configuration entries an returns a slice of credentials.
func (c *CCache) GetEntries() []*Credential {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:238
	_go_fuzz_dep_.CoverTab[86614]++
													creds := make([]*Credential, 0)
													for _, cred := range c.Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:240
		_go_fuzz_dep_.CoverTab[86616]++

														if strings.HasPrefix(cred.Server.Realm, "X-CACHECONF") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:242
			_go_fuzz_dep_.CoverTab[86618]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:243
			// _ = "end of CoverTab[86618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:244
			_go_fuzz_dep_.CoverTab[86619]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:244
			// _ = "end of CoverTab[86619]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:244
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:244
		// _ = "end of CoverTab[86616]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:244
		_go_fuzz_dep_.CoverTab[86617]++
														creds = append(creds, cred)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:245
		// _ = "end of CoverTab[86617]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:246
	// _ = "end of CoverTab[86614]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:246
	_go_fuzz_dep_.CoverTab[86615]++
													return creds
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:247
	// _ = "end of CoverTab[86615]"
}

func (h *headerField) valid() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:250
	_go_fuzz_dep_.CoverTab[86620]++

													switch h.tag {
	case headerFieldTagKDCOffset:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:253
		_go_fuzz_dep_.CoverTab[86622]++
														if h.length != 8 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:254
			_go_fuzz_dep_.CoverTab[86625]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:254
			return len(h.value) != 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:254
			// _ = "end of CoverTab[86625]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:254
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:254
			_go_fuzz_dep_.CoverTab[86626]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:255
			// _ = "end of CoverTab[86626]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:256
			_go_fuzz_dep_.CoverTab[86627]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:256
			// _ = "end of CoverTab[86627]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:256
		// _ = "end of CoverTab[86622]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:256
		_go_fuzz_dep_.CoverTab[86623]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:257
		// _ = "end of CoverTab[86623]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:257
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:257
		_go_fuzz_dep_.CoverTab[86624]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:257
		// _ = "end of CoverTab[86624]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:258
	// _ = "end of CoverTab[86620]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:258
	_go_fuzz_dep_.CoverTab[86621]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:259
	// _ = "end of CoverTab[86621]"
}

func readData(b []byte, p *int, e *binary.ByteOrder) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:262
	_go_fuzz_dep_.CoverTab[86628]++
													l := readInt32(b, p, e)
													return readBytes(b, p, int(l), e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:264
	// _ = "end of CoverTab[86628]"
}

func readAddress(b []byte, p *int, e *binary.ByteOrder) types.HostAddress {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:267
	_go_fuzz_dep_.CoverTab[86629]++
													a := types.HostAddress{}
													a.AddrType = int32(readInt16(b, p, e))
													a.Address = readData(b, p, e)
													return a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:271
	// _ = "end of CoverTab[86629]"
}

func readAuthDataEntry(b []byte, p *int, e *binary.ByteOrder) types.AuthorizationDataEntry {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:274
	_go_fuzz_dep_.CoverTab[86630]++
													a := types.AuthorizationDataEntry{}
													a.ADType = int32(readInt16(b, p, e))
													a.ADData = readData(b, p, e)
													return a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:278
	// _ = "end of CoverTab[86630]"
}

// Read bytes representing a timestamp.
func readTimestamp(b []byte, p *int, e *binary.ByteOrder) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:282
	_go_fuzz_dep_.CoverTab[86631]++
													return time.Unix(int64(readInt32(b, p, e)), 0)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:283
	// _ = "end of CoverTab[86631]"
}

// Read bytes representing an eight bit integer.
func readInt8(b []byte, p *int, e *binary.ByteOrder) (i int8) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:287
	_go_fuzz_dep_.CoverTab[86632]++
													buf := bytes.NewBuffer(b[*p : *p+1])
													binary.Read(buf, *e, &i)
													*p++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:291
	// _ = "end of CoverTab[86632]"
}

// Read bytes representing a sixteen bit integer.
func readInt16(b []byte, p *int, e *binary.ByteOrder) (i int16) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:295
	_go_fuzz_dep_.CoverTab[86633]++
													buf := bytes.NewBuffer(b[*p : *p+2])
													binary.Read(buf, *e, &i)
													*p += 2
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:299
	// _ = "end of CoverTab[86633]"
}

// Read bytes representing a thirty two bit integer.
func readInt32(b []byte, p *int, e *binary.ByteOrder) (i int32) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:303
	_go_fuzz_dep_.CoverTab[86634]++
													buf := bytes.NewBuffer(b[*p : *p+4])
													binary.Read(buf, *e, &i)
													*p += 4
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:307
	// _ = "end of CoverTab[86634]"
}

func readBytes(b []byte, p *int, s int, e *binary.ByteOrder) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:310
	_go_fuzz_dep_.CoverTab[86635]++
													buf := bytes.NewBuffer(b[*p : *p+s])
													r := make([]byte, s)
													binary.Read(buf, *e, &r)
													*p += s
													return r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:315
	// _ = "end of CoverTab[86635]"
}

func isNativeEndianLittle() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:318
	_go_fuzz_dep_.CoverTab[86636]++
													var x = 0x012345678
													var p = unsafe.Pointer(&x)
													var bp = (*[4]byte)(p)

													var endian bool
													if 0x01 == bp[0] {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:324
		_go_fuzz_dep_.CoverTab[86638]++
														endian = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:325
		// _ = "end of CoverTab[86638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:326
		_go_fuzz_dep_.CoverTab[86639]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:326
		if (0x78 & 0xff) == (bp[0] & 0xff) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:326
			_go_fuzz_dep_.CoverTab[86640]++
															endian = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:327
			// _ = "end of CoverTab[86640]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:328
			_go_fuzz_dep_.CoverTab[86641]++

															endian = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:330
			// _ = "end of CoverTab[86641]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:331
		// _ = "end of CoverTab[86639]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:331
	// _ = "end of CoverTab[86636]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:331
	_go_fuzz_dep_.CoverTab[86637]++
													return endian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:332
	// _ = "end of CoverTab[86637]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:333
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/ccache.go:333
var _ = _go_fuzz_dep_.CoverTab
