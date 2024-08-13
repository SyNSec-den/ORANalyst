//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:1
// Package keytab implements Kerberos keytabs: https://web.mit.edu/kerberos/krb5-devel/doc/formats/keytab_file_format.html.
package keytab

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:2
)

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
	"unsafe"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/types"
)

const (
	keytabFirstByte byte = 05
)

// Keytab struct.
type Keytab struct {
	version	uint8
	Entries	[]entry
}

// Keytab entry struct.
type entry struct {
	Principal	principal
	Timestamp	time.Time
	KVNO8		uint8
	Key		types.EncryptionKey
	KVNO		uint32
}

func (e entry) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:39
	_go_fuzz_dep_.CoverTab[86328]++
												return fmt.Sprintf("% 4d %s %-56s %2d %-64x",
		e.KVNO8,
		e.Timestamp.Format("02/01/06 15:04:05"),
		e.Principal.String(),
		e.Key.KeyType,
		e.Key.KeyValue,
	)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:46
	// _ = "end of CoverTab[86328]"
}

// Keytab entry principal struct.
type principal struct {
	NumComponents	int16	`json:"-"`
	Realm		string
	Components	[]string
	NameType	int32
}

func (p principal) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:57
	_go_fuzz_dep_.CoverTab[86329]++
												return fmt.Sprintf("%s@%s", strings.Join(p.Components, "/"), p.Realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:58
	// _ = "end of CoverTab[86329]"
}

// New creates new, empty Keytab type.
func New() *Keytab {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:62
	_go_fuzz_dep_.CoverTab[86330]++
												var e []entry
												return &Keytab{
		version:	2,
		Entries:	e,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:67
	// _ = "end of CoverTab[86330]"
}

// GetEncryptionKey returns the EncryptionKey from the Keytab for the newest entry with the required kvno, etype and matching principal.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:70
// If the kvno is zero then the latest kvno will be returned. The kvno is also returned for
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:72
func (kt *Keytab) GetEncryptionKey(princName types.PrincipalName, realm string, kvno int, etype int32) (types.EncryptionKey, int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:72
	_go_fuzz_dep_.CoverTab[86331]++
												var key types.EncryptionKey
												var t time.Time
												var kv int
												for _, k := range kt.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:76
		_go_fuzz_dep_.CoverTab[86334]++
													if k.Principal.Realm == realm && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
			_go_fuzz_dep_.CoverTab[86335]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
			return len(k.Principal.Components) == len(princName.NameString)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
			// _ = "end of CoverTab[86335]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
			_go_fuzz_dep_.CoverTab[86336]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:77
			return k.Key.KeyType == etype
														// _ = "end of CoverTab[86336]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:78
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:78
			_go_fuzz_dep_.CoverTab[86337]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:78
			return (k.KVNO == uint32(kvno) || func() bool {
															_go_fuzz_dep_.CoverTab[86338]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
				return kvno == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
				// _ = "end of CoverTab[86338]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
			}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
			// _ = "end of CoverTab[86337]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
			_go_fuzz_dep_.CoverTab[86339]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:79
			return k.Timestamp.After(t)
														// _ = "end of CoverTab[86339]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:80
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:80
			_go_fuzz_dep_.CoverTab[86340]++
														p := true
														for i, n := range k.Principal.Components {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:82
				_go_fuzz_dep_.CoverTab[86342]++
															if princName.NameString[i] != n {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:83
					_go_fuzz_dep_.CoverTab[86343]++
																p = false
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:85
					// _ = "end of CoverTab[86343]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:86
					_go_fuzz_dep_.CoverTab[86344]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:86
					// _ = "end of CoverTab[86344]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:86
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:86
				// _ = "end of CoverTab[86342]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:87
			// _ = "end of CoverTab[86340]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:87
			_go_fuzz_dep_.CoverTab[86341]++
														if p {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:88
				_go_fuzz_dep_.CoverTab[86345]++
															key = k.Key
															kv = int(k.KVNO)
															t = k.Timestamp
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:91
				// _ = "end of CoverTab[86345]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:92
				_go_fuzz_dep_.CoverTab[86346]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:92
				// _ = "end of CoverTab[86346]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:92
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:92
			// _ = "end of CoverTab[86341]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:93
			_go_fuzz_dep_.CoverTab[86347]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:93
			// _ = "end of CoverTab[86347]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:93
		// _ = "end of CoverTab[86334]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:94
	// _ = "end of CoverTab[86331]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:94
	_go_fuzz_dep_.CoverTab[86332]++
												if len(key.KeyValue) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:95
		_go_fuzz_dep_.CoverTab[86348]++
													return key, 0, fmt.Errorf("matching key not found in keytab. Looking for %v realm: %v kvno: %v etype: %v", princName.NameString, realm, kvno, etype)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:96
		// _ = "end of CoverTab[86348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:97
		_go_fuzz_dep_.CoverTab[86349]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:97
		// _ = "end of CoverTab[86349]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:97
	// _ = "end of CoverTab[86332]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:97
	_go_fuzz_dep_.CoverTab[86333]++
												return key, kv, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:98
	// _ = "end of CoverTab[86333]"
}

// Create a new Keytab entry.
func newEntry() entry {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:102
	_go_fuzz_dep_.CoverTab[86350]++
												var b []byte
												return entry{
		Principal:	newPrincipal(),
		Timestamp:	time.Time{},
		KVNO8:		0,
		Key: types.EncryptionKey{
			KeyType:	0,
			KeyValue:	b,
		},
		KVNO:	0,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:113
	// _ = "end of CoverTab[86350]"
}

func (kt Keytab) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:116
	_go_fuzz_dep_.CoverTab[86351]++
												var s string
												s = `KVNO Timestamp         Principal                                                ET Key
---- ----------------- -------------------------------------------------------- -- ----------------------------------------------------------------
`
	for _, entry := range kt.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:121
		_go_fuzz_dep_.CoverTab[86353]++
													s += entry.String() + "\n"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:122
		// _ = "end of CoverTab[86353]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:123
	// _ = "end of CoverTab[86351]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:123
	_go_fuzz_dep_.CoverTab[86352]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:124
	// _ = "end of CoverTab[86352]"
}

// AddEntry adds an entry to the keytab. The password should be provided in plain text and it will be converted using the defined enctype to be stored.
func (kt *Keytab) AddEntry(principalName, realm, password string, ts time.Time, KVNO uint8, encType int32) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:128
	_go_fuzz_dep_.CoverTab[86354]++

												princ, _ := types.ParseSPNString(principalName)
												key, _, err := crypto.GetKeyFromPassword(password, princ, realm, encType, types.PADataSequence{})
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:132
		_go_fuzz_dep_.CoverTab[86357]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:133
		// _ = "end of CoverTab[86357]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:134
		_go_fuzz_dep_.CoverTab[86358]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:134
		// _ = "end of CoverTab[86358]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:134
	// _ = "end of CoverTab[86354]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:134
	_go_fuzz_dep_.CoverTab[86355]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:137
	ktep := newPrincipal()
	ktep.NumComponents = int16(len(princ.NameString))
	if kt.version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:139
		_go_fuzz_dep_.CoverTab[86359]++
													ktep.NumComponents += 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:140
		// _ = "end of CoverTab[86359]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:141
		_go_fuzz_dep_.CoverTab[86360]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:141
		// _ = "end of CoverTab[86360]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:141
	// _ = "end of CoverTab[86355]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:141
	_go_fuzz_dep_.CoverTab[86356]++

												ktep.Realm = realm
												ktep.Components = princ.NameString
												ktep.NameType = princ.NameType

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:148
	e := newEntry()
												e.Principal = ktep
												e.Timestamp = ts
												e.KVNO8 = KVNO
												e.KVNO = uint32(KVNO)
												e.Key = key

												kt.Entries = append(kt.Entries, e)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:156
	// _ = "end of CoverTab[86356]"
}

// Create a new principal.
func newPrincipal() principal {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:160
	_go_fuzz_dep_.CoverTab[86361]++
												var c []string
												return principal{
		NumComponents:	0,
		Realm:		"",
		Components:	c,
		NameType:	0,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:167
	// _ = "end of CoverTab[86361]"
}

// Load a Keytab file into a Keytab type.
func Load(ktPath string) (*Keytab, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:171
	_go_fuzz_dep_.CoverTab[86362]++
												kt := new(Keytab)
												b, err := ioutil.ReadFile(ktPath)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:174
		_go_fuzz_dep_.CoverTab[86364]++
													return kt, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:175
		// _ = "end of CoverTab[86364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:176
		_go_fuzz_dep_.CoverTab[86365]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:176
		// _ = "end of CoverTab[86365]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:176
	// _ = "end of CoverTab[86362]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:176
	_go_fuzz_dep_.CoverTab[86363]++
												err = kt.Unmarshal(b)
												return kt, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:178
	// _ = "end of CoverTab[86363]"
}

// Marshal keytab into byte slice
func (kt *Keytab) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:182
	_go_fuzz_dep_.CoverTab[86366]++
												b := []byte{keytabFirstByte, kt.version}
												for _, e := range kt.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:184
		_go_fuzz_dep_.CoverTab[86368]++
													eb, err := e.marshal(int(kt.version))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:186
			_go_fuzz_dep_.CoverTab[86370]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:187
			// _ = "end of CoverTab[86370]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:188
			_go_fuzz_dep_.CoverTab[86371]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:188
			// _ = "end of CoverTab[86371]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:188
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:188
		// _ = "end of CoverTab[86368]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:188
		_go_fuzz_dep_.CoverTab[86369]++
													b = append(b, eb...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:189
		// _ = "end of CoverTab[86369]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:190
	// _ = "end of CoverTab[86366]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:190
	_go_fuzz_dep_.CoverTab[86367]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:191
	// _ = "end of CoverTab[86367]"
}

// Write the keytab bytes to io.Writer.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:194
// Returns the number of bytes written
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:196
func (kt *Keytab) Write(w io.Writer) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:196
	_go_fuzz_dep_.CoverTab[86372]++
												b, err := kt.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:198
		_go_fuzz_dep_.CoverTab[86374]++
													return 0, fmt.Errorf("error marshaling keytab: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:199
		// _ = "end of CoverTab[86374]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:200
		_go_fuzz_dep_.CoverTab[86375]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:200
		// _ = "end of CoverTab[86375]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:200
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:200
	// _ = "end of CoverTab[86372]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:200
	_go_fuzz_dep_.CoverTab[86373]++
												return w.Write(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:201
	// _ = "end of CoverTab[86373]"
}

// Unmarshal byte slice of Keytab data into Keytab type.
func (kt *Keytab) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:205
	_go_fuzz_dep_.CoverTab[86376]++
												if len(b) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:206
		_go_fuzz_dep_.CoverTab[86383]++
													return fmt.Errorf("byte array is less than 2 bytes: %d", len(b))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:207
		// _ = "end of CoverTab[86383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:208
		_go_fuzz_dep_.CoverTab[86384]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:208
		// _ = "end of CoverTab[86384]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:208
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:208
	// _ = "end of CoverTab[86376]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:208
	_go_fuzz_dep_.CoverTab[86377]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:211
	if b[0] != keytabFirstByte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:211
		_go_fuzz_dep_.CoverTab[86385]++
													return errors.New("invalid keytab data. First byte does not equal 5")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:212
		// _ = "end of CoverTab[86385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:213
		_go_fuzz_dep_.CoverTab[86386]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:213
		// _ = "end of CoverTab[86386]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:213
	// _ = "end of CoverTab[86377]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:213
	_go_fuzz_dep_.CoverTab[86378]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:216
	kt.version = b[1]
	if kt.version != 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:217
		_go_fuzz_dep_.CoverTab[86387]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:217
		return kt.version != 2
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:217
		// _ = "end of CoverTab[86387]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:217
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:217
		_go_fuzz_dep_.CoverTab[86388]++
													return errors.New("invalid keytab data. Keytab version is neither 1 nor 2")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:218
		// _ = "end of CoverTab[86388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:219
		_go_fuzz_dep_.CoverTab[86389]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:219
		// _ = "end of CoverTab[86389]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:219
	// _ = "end of CoverTab[86378]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:219
	_go_fuzz_dep_.CoverTab[86379]++
	//Version 1 of the file format uses native byte order for integer representations. Version 2 always uses big-endian byte order
	var endian binary.ByteOrder
	endian = binary.BigEndian
	if kt.version == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:223
		_go_fuzz_dep_.CoverTab[86390]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:223
		return isNativeEndianLittle()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:223
		// _ = "end of CoverTab[86390]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:223
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:223
		_go_fuzz_dep_.CoverTab[86391]++
													endian = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:224
		// _ = "end of CoverTab[86391]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:225
		_go_fuzz_dep_.CoverTab[86392]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:225
		// _ = "end of CoverTab[86392]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:225
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:225
	// _ = "end of CoverTab[86379]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:225
	_go_fuzz_dep_.CoverTab[86380]++

												n := 2
												l, err := readInt32(b, &n, &endian)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:229
		_go_fuzz_dep_.CoverTab[86393]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:230
		// _ = "end of CoverTab[86393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:231
		_go_fuzz_dep_.CoverTab[86394]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:231
		// _ = "end of CoverTab[86394]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:231
	// _ = "end of CoverTab[86380]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:231
	_go_fuzz_dep_.CoverTab[86381]++
												for l != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:232
		_go_fuzz_dep_.CoverTab[86395]++
													if l < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:233
			_go_fuzz_dep_.CoverTab[86398]++

														l = l * -1
														n = n + int(l)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:236
			// _ = "end of CoverTab[86398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:237
			_go_fuzz_dep_.CoverTab[86399]++
														if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:238
				_go_fuzz_dep_.CoverTab[86409]++
															return fmt.Errorf("%d can't be less than zero", n)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:239
				// _ = "end of CoverTab[86409]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:240
				_go_fuzz_dep_.CoverTab[86410]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:240
				// _ = "end of CoverTab[86410]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:240
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:240
			// _ = "end of CoverTab[86399]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:240
			_go_fuzz_dep_.CoverTab[86400]++
														if n+int(l) > len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:241
				_go_fuzz_dep_.CoverTab[86411]++
															return fmt.Errorf("%s's length is less than %d", b, n+int(l))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:242
				// _ = "end of CoverTab[86411]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:243
				_go_fuzz_dep_.CoverTab[86412]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:243
				// _ = "end of CoverTab[86412]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:243
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:243
			// _ = "end of CoverTab[86400]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:243
			_go_fuzz_dep_.CoverTab[86401]++
														eb := b[n : n+int(l)]
														n = n + int(l)
														ke := newEntry()
			// p keeps track as to where we are in the byte stream
			var p int
			var err error
			parsePrincipal(eb, &p, kt, &ke, &endian)
			ke.Timestamp, err = readTimestamp(eb, &p, &endian)
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:252
				_go_fuzz_dep_.CoverTab[86413]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:253
				// _ = "end of CoverTab[86413]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:254
				_go_fuzz_dep_.CoverTab[86414]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:254
				// _ = "end of CoverTab[86414]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:254
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:254
			// _ = "end of CoverTab[86401]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:254
			_go_fuzz_dep_.CoverTab[86402]++
														rei8, err := readInt8(eb, &p, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:256
				_go_fuzz_dep_.CoverTab[86415]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:257
				// _ = "end of CoverTab[86415]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:258
				_go_fuzz_dep_.CoverTab[86416]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:258
				// _ = "end of CoverTab[86416]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:258
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:258
			// _ = "end of CoverTab[86402]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:258
			_go_fuzz_dep_.CoverTab[86403]++
														ke.KVNO8 = uint8(rei8)
														rei16, err := readInt16(eb, &p, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:261
				_go_fuzz_dep_.CoverTab[86417]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:262
				// _ = "end of CoverTab[86417]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:263
				_go_fuzz_dep_.CoverTab[86418]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:263
				// _ = "end of CoverTab[86418]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:263
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:263
			// _ = "end of CoverTab[86403]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:263
			_go_fuzz_dep_.CoverTab[86404]++
														ke.Key.KeyType = int32(rei16)
														rei16, err = readInt16(eb, &p, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:266
				_go_fuzz_dep_.CoverTab[86419]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:267
				// _ = "end of CoverTab[86419]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:268
				_go_fuzz_dep_.CoverTab[86420]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:268
				// _ = "end of CoverTab[86420]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:268
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:268
			// _ = "end of CoverTab[86404]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:268
			_go_fuzz_dep_.CoverTab[86405]++
														kl := int(rei16)
														ke.Key.KeyValue, err = readBytes(eb, &p, kl, &endian)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:271
				_go_fuzz_dep_.CoverTab[86421]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:272
				// _ = "end of CoverTab[86421]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:273
				_go_fuzz_dep_.CoverTab[86422]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:273
				// _ = "end of CoverTab[86422]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:273
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:273
			// _ = "end of CoverTab[86405]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:273
			_go_fuzz_dep_.CoverTab[86406]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:277
			if len(eb)-p >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:277
				_go_fuzz_dep_.CoverTab[86423]++

															ri32, err := readInt32(eb, &p, &endian)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:280
					_go_fuzz_dep_.CoverTab[86425]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:281
					// _ = "end of CoverTab[86425]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:282
					_go_fuzz_dep_.CoverTab[86426]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:282
					// _ = "end of CoverTab[86426]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:282
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:282
				// _ = "end of CoverTab[86423]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:282
				_go_fuzz_dep_.CoverTab[86424]++
															ke.KVNO = uint32(ri32)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:283
				// _ = "end of CoverTab[86424]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:284
				_go_fuzz_dep_.CoverTab[86427]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:284
				// _ = "end of CoverTab[86427]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:284
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:284
			// _ = "end of CoverTab[86406]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:284
			_go_fuzz_dep_.CoverTab[86407]++
														if ke.KVNO == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:285
				_go_fuzz_dep_.CoverTab[86428]++

															ke.KVNO = uint32(ke.KVNO8)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:287
				// _ = "end of CoverTab[86428]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:288
				_go_fuzz_dep_.CoverTab[86429]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:288
				// _ = "end of CoverTab[86429]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:288
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:288
			// _ = "end of CoverTab[86407]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:288
			_go_fuzz_dep_.CoverTab[86408]++

														kt.Entries = append(kt.Entries, ke)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:290
			// _ = "end of CoverTab[86408]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:291
		// _ = "end of CoverTab[86395]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:291
		_go_fuzz_dep_.CoverTab[86396]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
		if n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			_go_fuzz_dep_.CoverTab[86430]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			return n > len(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			// _ = "end of CoverTab[86430]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			_go_fuzz_dep_.CoverTab[86431]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			return len(b[n:]) < 4
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			// _ = "end of CoverTab[86431]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:294
			_go_fuzz_dep_.CoverTab[86432]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:295
			// _ = "end of CoverTab[86432]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:296
			_go_fuzz_dep_.CoverTab[86433]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:296
			// _ = "end of CoverTab[86433]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:296
		// _ = "end of CoverTab[86396]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:296
		_go_fuzz_dep_.CoverTab[86397]++

													l, err = readInt32(b, &n, &endian)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:299
			_go_fuzz_dep_.CoverTab[86434]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:300
			// _ = "end of CoverTab[86434]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:301
			_go_fuzz_dep_.CoverTab[86435]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:301
			// _ = "end of CoverTab[86435]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:301
		// _ = "end of CoverTab[86397]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:302
	// _ = "end of CoverTab[86381]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:302
	_go_fuzz_dep_.CoverTab[86382]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:303
	// _ = "end of CoverTab[86382]"
}

func (e entry) marshal(v int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:306
	_go_fuzz_dep_.CoverTab[86436]++
												var b []byte
												pb, err := e.Principal.marshal(v)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:309
		_go_fuzz_dep_.CoverTab[86440]++
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:310
		// _ = "end of CoverTab[86440]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:311
		_go_fuzz_dep_.CoverTab[86441]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:311
		// _ = "end of CoverTab[86441]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:311
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:311
	// _ = "end of CoverTab[86436]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:311
	_go_fuzz_dep_.CoverTab[86437]++
												b = append(b, pb...)

												var endian binary.ByteOrder
												endian = binary.BigEndian
												if v == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:316
		_go_fuzz_dep_.CoverTab[86442]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:316
		return isNativeEndianLittle()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:316
		// _ = "end of CoverTab[86442]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:316
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:316
		_go_fuzz_dep_.CoverTab[86443]++
													endian = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:317
		// _ = "end of CoverTab[86443]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:318
		_go_fuzz_dep_.CoverTab[86444]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:318
		// _ = "end of CoverTab[86444]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:318
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:318
	// _ = "end of CoverTab[86437]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:318
	_go_fuzz_dep_.CoverTab[86438]++

												t := make([]byte, 9)
												endian.PutUint32(t[0:4], uint32(e.Timestamp.Unix()))
												t[4] = e.KVNO8
												endian.PutUint16(t[5:7], uint16(e.Key.KeyType))
												endian.PutUint16(t[7:9], uint16(len(e.Key.KeyValue)))
												b = append(b, t...)

												buf := new(bytes.Buffer)
												err = binary.Write(buf, endian, e.Key.KeyValue)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:329
		_go_fuzz_dep_.CoverTab[86445]++
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:330
		// _ = "end of CoverTab[86445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:331
		_go_fuzz_dep_.CoverTab[86446]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:331
		// _ = "end of CoverTab[86446]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:331
	// _ = "end of CoverTab[86438]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:331
	_go_fuzz_dep_.CoverTab[86439]++
												b = append(b, buf.Bytes()...)

												t = make([]byte, 4)
												endian.PutUint32(t, e.KVNO)
												b = append(b, t...)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:339
	t = make([]byte, 4)
												endian.PutUint32(t, uint32(len(b)))
												b = append(t, b...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:342
	// _ = "end of CoverTab[86439]"
}

// Parse the Keytab bytes of a principal into a Keytab entry's principal.
func parsePrincipal(b []byte, p *int, kt *Keytab, ke *entry, e *binary.ByteOrder) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:346
	_go_fuzz_dep_.CoverTab[86447]++
												var err error
												ke.Principal.NumComponents, err = readInt16(b, p, e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:349
		_go_fuzz_dep_.CoverTab[86454]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:350
		// _ = "end of CoverTab[86454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:351
		_go_fuzz_dep_.CoverTab[86455]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:351
		// _ = "end of CoverTab[86455]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:351
	// _ = "end of CoverTab[86447]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:351
	_go_fuzz_dep_.CoverTab[86448]++
												if kt.version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:352
		_go_fuzz_dep_.CoverTab[86456]++

													ke.Principal.NumComponents--
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:354
		// _ = "end of CoverTab[86456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:355
		_go_fuzz_dep_.CoverTab[86457]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:355
		// _ = "end of CoverTab[86457]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:355
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:355
	// _ = "end of CoverTab[86448]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:355
	_go_fuzz_dep_.CoverTab[86449]++
												lenRealm, err := readInt16(b, p, e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:357
		_go_fuzz_dep_.CoverTab[86458]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:358
		// _ = "end of CoverTab[86458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:359
		_go_fuzz_dep_.CoverTab[86459]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:359
		// _ = "end of CoverTab[86459]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:359
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:359
	// _ = "end of CoverTab[86449]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:359
	_go_fuzz_dep_.CoverTab[86450]++
												realmB, err := readBytes(b, p, int(lenRealm), e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:361
		_go_fuzz_dep_.CoverTab[86460]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:362
		// _ = "end of CoverTab[86460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:363
		_go_fuzz_dep_.CoverTab[86461]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:363
		// _ = "end of CoverTab[86461]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:363
	// _ = "end of CoverTab[86450]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:363
	_go_fuzz_dep_.CoverTab[86451]++
												ke.Principal.Realm = string(realmB)
												for i := 0; i < int(ke.Principal.NumComponents); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:365
		_go_fuzz_dep_.CoverTab[86462]++
													l, err := readInt16(b, p, e)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:367
			_go_fuzz_dep_.CoverTab[86465]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:368
			// _ = "end of CoverTab[86465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:369
			_go_fuzz_dep_.CoverTab[86466]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:369
			// _ = "end of CoverTab[86466]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:369
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:369
		// _ = "end of CoverTab[86462]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:369
		_go_fuzz_dep_.CoverTab[86463]++
													compB, err := readBytes(b, p, int(l), e)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:371
			_go_fuzz_dep_.CoverTab[86467]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:372
			// _ = "end of CoverTab[86467]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:373
			_go_fuzz_dep_.CoverTab[86468]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:373
			// _ = "end of CoverTab[86468]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:373
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:373
		// _ = "end of CoverTab[86463]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:373
		_go_fuzz_dep_.CoverTab[86464]++
													ke.Principal.Components = append(ke.Principal.Components, string(compB))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:374
		// _ = "end of CoverTab[86464]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:375
	// _ = "end of CoverTab[86451]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:375
	_go_fuzz_dep_.CoverTab[86452]++
												if kt.version != 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:376
		_go_fuzz_dep_.CoverTab[86469]++

													ke.Principal.NameType, err = readInt32(b, p, e)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:379
			_go_fuzz_dep_.CoverTab[86470]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:380
			// _ = "end of CoverTab[86470]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:381
			_go_fuzz_dep_.CoverTab[86471]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:381
			// _ = "end of CoverTab[86471]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:381
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:381
		// _ = "end of CoverTab[86469]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:382
		_go_fuzz_dep_.CoverTab[86472]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:382
		// _ = "end of CoverTab[86472]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:382
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:382
	// _ = "end of CoverTab[86452]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:382
	_go_fuzz_dep_.CoverTab[86453]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:383
	// _ = "end of CoverTab[86453]"
}

func (p principal) marshal(v int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:386
	_go_fuzz_dep_.CoverTab[86473]++

												b := make([]byte, 2)
												var endian binary.ByteOrder
												endian = binary.BigEndian
												if v == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:391
		_go_fuzz_dep_.CoverTab[86478]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:391
		return isNativeEndianLittle()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:391
		// _ = "end of CoverTab[86478]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:391
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:391
		_go_fuzz_dep_.CoverTab[86479]++
													endian = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:392
		// _ = "end of CoverTab[86479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:393
		_go_fuzz_dep_.CoverTab[86480]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:393
		// _ = "end of CoverTab[86480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:393
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:393
	// _ = "end of CoverTab[86473]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:393
	_go_fuzz_dep_.CoverTab[86474]++
												endian.PutUint16(b[0:], uint16(p.NumComponents))
												realm, err := marshalString(p.Realm, v)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:396
		_go_fuzz_dep_.CoverTab[86481]++
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:397
		// _ = "end of CoverTab[86481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:398
		_go_fuzz_dep_.CoverTab[86482]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:398
		// _ = "end of CoverTab[86482]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:398
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:398
	// _ = "end of CoverTab[86474]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:398
	_go_fuzz_dep_.CoverTab[86475]++
												b = append(b, realm...)
												for _, c := range p.Components {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:400
		_go_fuzz_dep_.CoverTab[86483]++
													cb, err := marshalString(c, v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:402
			_go_fuzz_dep_.CoverTab[86485]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:403
			// _ = "end of CoverTab[86485]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:404
			_go_fuzz_dep_.CoverTab[86486]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:404
			// _ = "end of CoverTab[86486]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:404
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:404
		// _ = "end of CoverTab[86483]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:404
		_go_fuzz_dep_.CoverTab[86484]++
													b = append(b, cb...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:405
		// _ = "end of CoverTab[86484]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:406
	// _ = "end of CoverTab[86475]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:406
	_go_fuzz_dep_.CoverTab[86476]++
												if v != 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:407
		_go_fuzz_dep_.CoverTab[86487]++
													t := make([]byte, 4)
													endian.PutUint32(t, uint32(p.NameType))
													b = append(b, t...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:410
		// _ = "end of CoverTab[86487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:411
		_go_fuzz_dep_.CoverTab[86488]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:411
		// _ = "end of CoverTab[86488]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:411
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:411
	// _ = "end of CoverTab[86476]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:411
	_go_fuzz_dep_.CoverTab[86477]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:412
	// _ = "end of CoverTab[86477]"
}

func marshalString(s string, v int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:415
	_go_fuzz_dep_.CoverTab[86489]++
												sb := []byte(s)
												b := make([]byte, 2)
												var endian binary.ByteOrder
												endian = binary.BigEndian
												if v == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:420
		_go_fuzz_dep_.CoverTab[86492]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:420
		return isNativeEndianLittle()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:420
		// _ = "end of CoverTab[86492]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:420
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:420
		_go_fuzz_dep_.CoverTab[86493]++
													endian = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:421
		// _ = "end of CoverTab[86493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:422
		_go_fuzz_dep_.CoverTab[86494]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:422
		// _ = "end of CoverTab[86494]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:422
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:422
	// _ = "end of CoverTab[86489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:422
	_go_fuzz_dep_.CoverTab[86490]++
												endian.PutUint16(b[0:], uint16(len(sb)))
												buf := new(bytes.Buffer)
												err := binary.Write(buf, endian, sb)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:426
		_go_fuzz_dep_.CoverTab[86495]++
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:427
		// _ = "end of CoverTab[86495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:428
		_go_fuzz_dep_.CoverTab[86496]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:428
		// _ = "end of CoverTab[86496]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:428
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:428
	// _ = "end of CoverTab[86490]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:428
	_go_fuzz_dep_.CoverTab[86491]++
												b = append(b, buf.Bytes()...)
												return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:430
	// _ = "end of CoverTab[86491]"
}

// Read bytes representing a timestamp.
func readTimestamp(b []byte, p *int, e *binary.ByteOrder) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:434
	_go_fuzz_dep_.CoverTab[86497]++
												i32, err := readInt32(b, p, e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:436
		_go_fuzz_dep_.CoverTab[86499]++
													return time.Time{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:437
		// _ = "end of CoverTab[86499]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:438
		_go_fuzz_dep_.CoverTab[86500]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:438
		// _ = "end of CoverTab[86500]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:438
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:438
	// _ = "end of CoverTab[86497]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:438
	_go_fuzz_dep_.CoverTab[86498]++
												return time.Unix(int64(i32), 0), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:439
	// _ = "end of CoverTab[86498]"
}

// Read bytes representing an eight bit integer.
func readInt8(b []byte, p *int, e *binary.ByteOrder) (i int8, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:443
	_go_fuzz_dep_.CoverTab[86501]++
												if *p < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:444
		_go_fuzz_dep_.CoverTab[86504]++
													return 0, fmt.Errorf("%d cannot be less than zero", *p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:445
		// _ = "end of CoverTab[86504]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:446
		_go_fuzz_dep_.CoverTab[86505]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:446
		// _ = "end of CoverTab[86505]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:446
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:446
	// _ = "end of CoverTab[86501]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:446
	_go_fuzz_dep_.CoverTab[86502]++

												if (*p + 1) > len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:448
		_go_fuzz_dep_.CoverTab[86506]++
													return 0, fmt.Errorf("%s's length is less than %d", b, *p+1)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:449
		// _ = "end of CoverTab[86506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:450
		_go_fuzz_dep_.CoverTab[86507]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:450
		// _ = "end of CoverTab[86507]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:450
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:450
	// _ = "end of CoverTab[86502]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:450
	_go_fuzz_dep_.CoverTab[86503]++
												buf := bytes.NewBuffer(b[*p : *p+1])
												binary.Read(buf, *e, &i)
												*p++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:454
	// _ = "end of CoverTab[86503]"
}

// Read bytes representing a sixteen bit integer.
func readInt16(b []byte, p *int, e *binary.ByteOrder) (i int16, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:458
	_go_fuzz_dep_.CoverTab[86508]++
												if *p < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:459
		_go_fuzz_dep_.CoverTab[86511]++
													return 0, fmt.Errorf("%d cannot be less than zero", *p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:460
		// _ = "end of CoverTab[86511]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:461
		_go_fuzz_dep_.CoverTab[86512]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:461
		// _ = "end of CoverTab[86512]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:461
	// _ = "end of CoverTab[86508]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:461
	_go_fuzz_dep_.CoverTab[86509]++

												if (*p + 2) > len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:463
		_go_fuzz_dep_.CoverTab[86513]++
													return 0, fmt.Errorf("%s's length is less than %d", b, *p+2)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:464
		// _ = "end of CoverTab[86513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:465
		_go_fuzz_dep_.CoverTab[86514]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:465
		// _ = "end of CoverTab[86514]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:465
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:465
	// _ = "end of CoverTab[86509]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:465
	_go_fuzz_dep_.CoverTab[86510]++

												buf := bytes.NewBuffer(b[*p : *p+2])
												binary.Read(buf, *e, &i)
												*p += 2
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:470
	// _ = "end of CoverTab[86510]"
}

// Read bytes representing a thirty two bit integer.
func readInt32(b []byte, p *int, e *binary.ByteOrder) (i int32, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:474
	_go_fuzz_dep_.CoverTab[86515]++
												if *p < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:475
		_go_fuzz_dep_.CoverTab[86518]++
													return 0, fmt.Errorf("%d cannot be less than zero", *p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:476
		// _ = "end of CoverTab[86518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:477
		_go_fuzz_dep_.CoverTab[86519]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:477
		// _ = "end of CoverTab[86519]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:477
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:477
	// _ = "end of CoverTab[86515]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:477
	_go_fuzz_dep_.CoverTab[86516]++

												if (*p + 4) > len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:479
		_go_fuzz_dep_.CoverTab[86520]++
													return 0, fmt.Errorf("%s's length is less than %d", b, *p+4)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:480
		// _ = "end of CoverTab[86520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:481
		_go_fuzz_dep_.CoverTab[86521]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:481
		// _ = "end of CoverTab[86521]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:481
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:481
	// _ = "end of CoverTab[86516]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:481
	_go_fuzz_dep_.CoverTab[86517]++

												buf := bytes.NewBuffer(b[*p : *p+4])
												binary.Read(buf, *e, &i)
												*p += 4
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:486
	// _ = "end of CoverTab[86517]"
}

func readBytes(b []byte, p *int, s int, e *binary.ByteOrder) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:489
	_go_fuzz_dep_.CoverTab[86522]++
												if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:490
		_go_fuzz_dep_.CoverTab[86526]++
													return nil, fmt.Errorf("%d cannot be less than zero", s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:491
		// _ = "end of CoverTab[86526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:492
		_go_fuzz_dep_.CoverTab[86527]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:492
		// _ = "end of CoverTab[86527]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:492
	// _ = "end of CoverTab[86522]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:492
	_go_fuzz_dep_.CoverTab[86523]++
												i := *p + s
												if i > len(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:494
		_go_fuzz_dep_.CoverTab[86528]++
													return nil, fmt.Errorf("%s's length is greater than %d", b, i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:495
		// _ = "end of CoverTab[86528]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:496
		_go_fuzz_dep_.CoverTab[86529]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:496
		// _ = "end of CoverTab[86529]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:496
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:496
	// _ = "end of CoverTab[86523]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:496
	_go_fuzz_dep_.CoverTab[86524]++
												buf := bytes.NewBuffer(b[*p:i])
												r := make([]byte, s)
												if err := binary.Read(buf, *e, &r); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:499
		_go_fuzz_dep_.CoverTab[86530]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:500
		// _ = "end of CoverTab[86530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:501
		_go_fuzz_dep_.CoverTab[86531]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:501
		// _ = "end of CoverTab[86531]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:501
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:501
	// _ = "end of CoverTab[86524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:501
	_go_fuzz_dep_.CoverTab[86525]++
												*p += s
												return r, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:503
	// _ = "end of CoverTab[86525]"
}

func isNativeEndianLittle() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:506
	_go_fuzz_dep_.CoverTab[86532]++
												var x = 0x012345678
												var p = unsafe.Pointer(&x)
												var bp = (*[4]byte)(p)

												var endian bool
												if 0x01 == bp[0] {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:512
		_go_fuzz_dep_.CoverTab[86534]++
													endian = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:513
		// _ = "end of CoverTab[86534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:514
		_go_fuzz_dep_.CoverTab[86535]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:514
		if (0x78 & 0xff) == (bp[0] & 0xff) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:514
			_go_fuzz_dep_.CoverTab[86536]++
														endian = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:515
			// _ = "end of CoverTab[86536]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:516
			_go_fuzz_dep_.CoverTab[86537]++

														endian = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:518
			// _ = "end of CoverTab[86537]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:519
		// _ = "end of CoverTab[86535]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:519
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:519
	// _ = "end of CoverTab[86532]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:519
	_go_fuzz_dep_.CoverTab[86533]++
												return endian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:520
	// _ = "end of CoverTab[86533]"
}

// JSON return information about the keys held in the keytab in a JSON format.
func (kt *Keytab) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:524
	_go_fuzz_dep_.CoverTab[86538]++
												b, err := json.MarshalIndent(kt, "", "  ")
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:526
		_go_fuzz_dep_.CoverTab[86540]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:527
		// _ = "end of CoverTab[86540]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:528
		_go_fuzz_dep_.CoverTab[86541]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:528
		// _ = "end of CoverTab[86541]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:528
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:528
	// _ = "end of CoverTab[86538]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:528
	_go_fuzz_dep_.CoverTab[86539]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:529
	// _ = "end of CoverTab[86539]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:530
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/keytab/keytab.go:530
var _ = _go_fuzz_dep_.CoverTab
