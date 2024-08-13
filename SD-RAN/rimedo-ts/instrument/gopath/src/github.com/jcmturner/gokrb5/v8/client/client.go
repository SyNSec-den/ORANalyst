//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:1
// Package client provides a client library and methods for Kerberos 5 authentication.
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:2
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Client side configuration and state.
type Client struct {
	Credentials	*credentials.Credentials
	Config		*config.Config
	settings	*Settings
	sessions	*sessions
	cache		*Cache
}

// NewWithPassword creates a new client from a password credential.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:33
// Set the realm to empty string to use the default realm from config.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:35
func NewWithPassword(username, realm, password string, krb5conf *config.Config, settings ...func(*Settings)) *Client {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:35
	_go_fuzz_dep_.CoverTab[88410]++
												creds := credentials.New(username, realm)
												return &Client{
		Credentials:	creds.WithPassword(password),
		Config:		krb5conf,
		settings:	NewSettings(settings...),
		sessions: &sessions{
			Entries: make(map[string]*session),
		},
		cache:	NewCache(),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:45
	// _ = "end of CoverTab[88410]"
}

// NewWithKeytab creates a new client from a keytab credential.
func NewWithKeytab(username, realm string, kt *keytab.Keytab, krb5conf *config.Config, settings ...func(*Settings)) *Client {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:49
	_go_fuzz_dep_.CoverTab[88411]++
												creds := credentials.New(username, realm)
												return &Client{
		Credentials:	creds.WithKeytab(kt),
		Config:		krb5conf,
		settings:	NewSettings(settings...),
		sessions: &sessions{
			Entries: make(map[string]*session),
		},
		cache:	NewCache(),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:59
	// _ = "end of CoverTab[88411]"
}

// NewFromCCache create a client from a populated client cache.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:62
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:62
// WARNING: A client created from CCache does not automatically renew TGTs and a failure will occur after the TGT expires.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:65
func NewFromCCache(c *credentials.CCache, krb5conf *config.Config, settings ...func(*Settings)) (*Client, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:65
	_go_fuzz_dep_.CoverTab[88412]++
												cl := &Client{
		Credentials:	c.GetClientCredentials(),
		Config:		krb5conf,
		settings:	NewSettings(settings...),
		sessions: &sessions{
			Entries: make(map[string]*session),
		},
		cache:	NewCache(),
	}
	spn := types.PrincipalName{
		NameType:	nametype.KRB_NT_SRV_INST,
		NameString:	[]string{"krbtgt", c.DefaultPrincipal.Realm},
	}
	cred, ok := c.GetEntry(spn)
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:80
		_go_fuzz_dep_.CoverTab[88416]++
													return cl, errors.New("TGT not found in CCache")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:81
		// _ = "end of CoverTab[88416]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:82
		_go_fuzz_dep_.CoverTab[88417]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:82
		// _ = "end of CoverTab[88417]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:82
	// _ = "end of CoverTab[88412]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:82
	_go_fuzz_dep_.CoverTab[88413]++
												var tgt messages.Ticket
												err := tgt.Unmarshal(cred.Ticket)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:85
		_go_fuzz_dep_.CoverTab[88418]++
													return cl, fmt.Errorf("TGT bytes in cache are not valid: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:86
		// _ = "end of CoverTab[88418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:87
		_go_fuzz_dep_.CoverTab[88419]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:87
		// _ = "end of CoverTab[88419]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:87
	// _ = "end of CoverTab[88413]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:87
	_go_fuzz_dep_.CoverTab[88414]++
												cl.sessions.Entries[c.DefaultPrincipal.Realm] = &session{
		realm:		c.DefaultPrincipal.Realm,
		authTime:	cred.AuthTime,
		endTime:	cred.EndTime,
		renewTill:	cred.RenewTill,
		tgt:		tgt,
		sessionKey:	cred.Key,
	}
	for _, cred := range c.GetEntries() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:96
		_go_fuzz_dep_.CoverTab[88420]++
													var tkt messages.Ticket
													err = tkt.Unmarshal(cred.Ticket)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:99
			_go_fuzz_dep_.CoverTab[88422]++
														return cl, fmt.Errorf("cache entry ticket bytes are not valid: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:100
			// _ = "end of CoverTab[88422]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:101
			_go_fuzz_dep_.CoverTab[88423]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:101
			// _ = "end of CoverTab[88423]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:101
		// _ = "end of CoverTab[88420]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:101
		_go_fuzz_dep_.CoverTab[88421]++
													cl.cache.addEntry(
			tkt,
			cred.AuthTime,
			cred.StartTime,
			cred.EndTime,
			cred.RenewTill,
			cred.Key,
		)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:109
		// _ = "end of CoverTab[88421]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:110
	// _ = "end of CoverTab[88414]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:110
	_go_fuzz_dep_.CoverTab[88415]++
												return cl, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:111
	// _ = "end of CoverTab[88415]"
}

// Key returns the client's encryption key for the specified encryption type and its kvno (kvno of zero will find latest).
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:114
// The key can be retrieved either from the keytab or generated from the client's password.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:114
// If the client has both a keytab and a password defined the keytab is favoured as the source for the key
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:114
// A KRBError can be passed in the event the KDC returns one of type KDC_ERR_PREAUTH_REQUIRED and is required to derive
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:114
// the key for pre-authentication from the client's password. If a KRBError is not available, pass nil to this argument.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:119
func (cl *Client) Key(etype etype.EType, kvno int, krberr *messages.KRBError) (types.EncryptionKey, int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:119
	_go_fuzz_dep_.CoverTab[88424]++
												if cl.Credentials.HasKeytab() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:120
		_go_fuzz_dep_.CoverTab[88426]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:120
		return etype != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:120
		// _ = "end of CoverTab[88426]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:120
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:120
		_go_fuzz_dep_.CoverTab[88427]++
													return cl.Credentials.Keytab().GetEncryptionKey(cl.Credentials.CName(), cl.Credentials.Domain(), kvno, etype.GetETypeID())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:121
		// _ = "end of CoverTab[88427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:122
		_go_fuzz_dep_.CoverTab[88428]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:122
		if cl.Credentials.HasPassword() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:122
			_go_fuzz_dep_.CoverTab[88429]++
														if krberr != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:123
				_go_fuzz_dep_.CoverTab[88431]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:123
				return krberr.ErrorCode == errorcode.KDC_ERR_PREAUTH_REQUIRED
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:123
				// _ = "end of CoverTab[88431]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:123
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:123
				_go_fuzz_dep_.CoverTab[88432]++
															var pas types.PADataSequence
															err := pas.Unmarshal(krberr.EData)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:126
					_go_fuzz_dep_.CoverTab[88434]++
																return types.EncryptionKey{}, 0, fmt.Errorf("could not get PAData from KRBError to generate key from password: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:127
					// _ = "end of CoverTab[88434]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:128
					_go_fuzz_dep_.CoverTab[88435]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:128
					// _ = "end of CoverTab[88435]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:128
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:128
				// _ = "end of CoverTab[88432]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:128
				_go_fuzz_dep_.CoverTab[88433]++
															key, _, err := crypto.GetKeyFromPassword(cl.Credentials.Password(), krberr.CName, krberr.CRealm, etype.GetETypeID(), pas)
															return key, 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:130
				// _ = "end of CoverTab[88433]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:131
				_go_fuzz_dep_.CoverTab[88436]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:131
				// _ = "end of CoverTab[88436]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:131
			// _ = "end of CoverTab[88429]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:131
			_go_fuzz_dep_.CoverTab[88430]++
														key, _, err := crypto.GetKeyFromPassword(cl.Credentials.Password(), cl.Credentials.CName(), cl.Credentials.Domain(), etype.GetETypeID(), types.PADataSequence{})
														return key, 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:133
			// _ = "end of CoverTab[88430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
			_go_fuzz_dep_.CoverTab[88437]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
			// _ = "end of CoverTab[88437]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
		// _ = "end of CoverTab[88428]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
	// _ = "end of CoverTab[88424]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:134
	_go_fuzz_dep_.CoverTab[88425]++
												return types.EncryptionKey{}, 0, errors.New("credential has neither keytab or password to generate key")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:135
	// _ = "end of CoverTab[88425]"
}

// IsConfigured indicates if the client has the values required set.
func (cl *Client) IsConfigured() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:139
	_go_fuzz_dep_.CoverTab[88438]++
												if cl.Credentials.UserName() == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:140
		_go_fuzz_dep_.CoverTab[88443]++
													return false, errors.New("client does not have a username")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:141
		// _ = "end of CoverTab[88443]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:142
		_go_fuzz_dep_.CoverTab[88444]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:142
		// _ = "end of CoverTab[88444]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:142
	// _ = "end of CoverTab[88438]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:142
	_go_fuzz_dep_.CoverTab[88439]++
												if cl.Credentials.Domain() == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:143
		_go_fuzz_dep_.CoverTab[88445]++
													return false, errors.New("client does not have a define realm")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:144
		// _ = "end of CoverTab[88445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:145
		_go_fuzz_dep_.CoverTab[88446]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:145
		// _ = "end of CoverTab[88446]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:145
	// _ = "end of CoverTab[88439]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:145
	_go_fuzz_dep_.CoverTab[88440]++

												if !cl.Credentials.HasPassword() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:147
		_go_fuzz_dep_.CoverTab[88447]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:147
		return !cl.Credentials.HasKeytab()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:147
		// _ = "end of CoverTab[88447]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:147
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:147
		_go_fuzz_dep_.CoverTab[88448]++
													authTime, _, _, _, err := cl.sessionTimes(cl.Credentials.Domain())
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:149
			_go_fuzz_dep_.CoverTab[88449]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:149
			return authTime.IsZero()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:149
			// _ = "end of CoverTab[88449]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:149
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:149
			_go_fuzz_dep_.CoverTab[88450]++
														return false, errors.New("client has neither a keytab nor a password set and no session")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:150
			// _ = "end of CoverTab[88450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:151
			_go_fuzz_dep_.CoverTab[88451]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:151
			// _ = "end of CoverTab[88451]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:151
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:151
		// _ = "end of CoverTab[88448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:152
		_go_fuzz_dep_.CoverTab[88452]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:152
		// _ = "end of CoverTab[88452]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:152
	// _ = "end of CoverTab[88440]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:152
	_go_fuzz_dep_.CoverTab[88441]++
												if !cl.Config.LibDefaults.DNSLookupKDC {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:153
		_go_fuzz_dep_.CoverTab[88453]++
													for _, r := range cl.Config.Realms {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:154
			_go_fuzz_dep_.CoverTab[88454]++
														if r.Realm == cl.Credentials.Domain() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:155
				_go_fuzz_dep_.CoverTab[88455]++
															if len(r.KDC) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:156
					_go_fuzz_dep_.CoverTab[88457]++
																return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:157
					// _ = "end of CoverTab[88457]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:158
					_go_fuzz_dep_.CoverTab[88458]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:158
					// _ = "end of CoverTab[88458]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:158
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:158
				// _ = "end of CoverTab[88455]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:158
				_go_fuzz_dep_.CoverTab[88456]++
															return false, errors.New("client krb5 config does not have any defined KDCs for the default realm")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:159
				// _ = "end of CoverTab[88456]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:160
				_go_fuzz_dep_.CoverTab[88459]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:160
				// _ = "end of CoverTab[88459]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:160
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:160
			// _ = "end of CoverTab[88454]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:161
		// _ = "end of CoverTab[88453]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:162
		_go_fuzz_dep_.CoverTab[88460]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:162
		// _ = "end of CoverTab[88460]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:162
	// _ = "end of CoverTab[88441]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:162
	_go_fuzz_dep_.CoverTab[88442]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:163
	// _ = "end of CoverTab[88442]"
}

// Login the client with the KDC via an AS exchange.
func (cl *Client) Login() error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:167
	_go_fuzz_dep_.CoverTab[88461]++
												if ok, err := cl.IsConfigured(); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:168
		_go_fuzz_dep_.CoverTab[88466]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:169
		// _ = "end of CoverTab[88466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:170
		_go_fuzz_dep_.CoverTab[88467]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:170
		// _ = "end of CoverTab[88467]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:170
	// _ = "end of CoverTab[88461]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:170
	_go_fuzz_dep_.CoverTab[88462]++
												if !cl.Credentials.HasPassword() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:171
		_go_fuzz_dep_.CoverTab[88468]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:171
		return !cl.Credentials.HasKeytab()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:171
		// _ = "end of CoverTab[88468]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:171
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:171
		_go_fuzz_dep_.CoverTab[88469]++
													_, endTime, _, _, err := cl.sessionTimes(cl.Credentials.Domain())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:173
			_go_fuzz_dep_.CoverTab[88472]++
														return krberror.Errorf(err, krberror.KRBMsgError, "no user credentials available and error getting any existing session")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:174
			// _ = "end of CoverTab[88472]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:175
			_go_fuzz_dep_.CoverTab[88473]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:175
			// _ = "end of CoverTab[88473]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:175
		// _ = "end of CoverTab[88469]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:175
		_go_fuzz_dep_.CoverTab[88470]++
													if time.Now().UTC().After(endTime) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:176
			_go_fuzz_dep_.CoverTab[88474]++
														return krberror.New(krberror.KRBMsgError, "cannot login, no user credentials available and no valid existing session")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:177
			// _ = "end of CoverTab[88474]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:178
			_go_fuzz_dep_.CoverTab[88475]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:178
			// _ = "end of CoverTab[88475]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:178
		// _ = "end of CoverTab[88470]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:178
		_go_fuzz_dep_.CoverTab[88471]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:180
		// _ = "end of CoverTab[88471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:181
		_go_fuzz_dep_.CoverTab[88476]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:181
		// _ = "end of CoverTab[88476]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:181
	// _ = "end of CoverTab[88462]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:181
	_go_fuzz_dep_.CoverTab[88463]++
												ASReq, err := messages.NewASReqForTGT(cl.Credentials.Domain(), cl.Config, cl.Credentials.CName())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:183
		_go_fuzz_dep_.CoverTab[88477]++
													return krberror.Errorf(err, krberror.KRBMsgError, "error generating new AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:184
		// _ = "end of CoverTab[88477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:185
		_go_fuzz_dep_.CoverTab[88478]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:185
		// _ = "end of CoverTab[88478]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:185
	// _ = "end of CoverTab[88463]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:185
	_go_fuzz_dep_.CoverTab[88464]++
												ASRep, err := cl.ASExchange(cl.Credentials.Domain(), ASReq, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:187
		_go_fuzz_dep_.CoverTab[88479]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:188
		// _ = "end of CoverTab[88479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:189
		_go_fuzz_dep_.CoverTab[88480]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:189
		// _ = "end of CoverTab[88480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:189
	// _ = "end of CoverTab[88464]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:189
	_go_fuzz_dep_.CoverTab[88465]++
												cl.addSession(ASRep.Ticket, ASRep.DecryptedEncPart)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:191
	// _ = "end of CoverTab[88465]"
}

// AffirmLogin will only perform an AS exchange with the KDC if the client does not already have a TGT.
func (cl *Client) AffirmLogin() error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:195
	_go_fuzz_dep_.CoverTab[88481]++
												_, endTime, _, _, err := cl.sessionTimes(cl.Credentials.Domain())
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:197
		_go_fuzz_dep_.CoverTab[88483]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:197
		return time.Now().UTC().After(endTime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:197
		// _ = "end of CoverTab[88483]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:197
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:197
		_go_fuzz_dep_.CoverTab[88484]++
													err := cl.Login()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:199
			_go_fuzz_dep_.CoverTab[88485]++
														return fmt.Errorf("could not get valid TGT for client's realm: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:200
			// _ = "end of CoverTab[88485]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:201
			_go_fuzz_dep_.CoverTab[88486]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:201
			// _ = "end of CoverTab[88486]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:201
		// _ = "end of CoverTab[88484]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:202
		_go_fuzz_dep_.CoverTab[88487]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:202
		// _ = "end of CoverTab[88487]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:202
	// _ = "end of CoverTab[88481]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:202
	_go_fuzz_dep_.CoverTab[88482]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:203
	// _ = "end of CoverTab[88482]"
}

// realmLogin obtains or renews a TGT and establishes a session for the realm specified.
func (cl *Client) realmLogin(realm string) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:207
	_go_fuzz_dep_.CoverTab[88488]++
												if realm == cl.Credentials.Domain() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:208
		_go_fuzz_dep_.CoverTab[88493]++
													return cl.Login()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:209
		// _ = "end of CoverTab[88493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:210
		_go_fuzz_dep_.CoverTab[88494]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:210
		// _ = "end of CoverTab[88494]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:210
	// _ = "end of CoverTab[88488]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:210
	_go_fuzz_dep_.CoverTab[88489]++
												_, endTime, _, _, err := cl.sessionTimes(cl.Credentials.Domain())
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:212
		_go_fuzz_dep_.CoverTab[88495]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:212
		return time.Now().UTC().After(endTime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:212
		// _ = "end of CoverTab[88495]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:212
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:212
		_go_fuzz_dep_.CoverTab[88496]++
													err := cl.Login()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:214
			_go_fuzz_dep_.CoverTab[88497]++
														return fmt.Errorf("could not get valid TGT for client's realm: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:215
			// _ = "end of CoverTab[88497]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:216
			_go_fuzz_dep_.CoverTab[88498]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:216
			// _ = "end of CoverTab[88498]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:216
		// _ = "end of CoverTab[88496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:217
		_go_fuzz_dep_.CoverTab[88499]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:217
		// _ = "end of CoverTab[88499]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:217
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:217
	// _ = "end of CoverTab[88489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:217
	_go_fuzz_dep_.CoverTab[88490]++
												tgt, skey, err := cl.sessionTGT(cl.Credentials.Domain())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:219
		_go_fuzz_dep_.CoverTab[88500]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:220
		// _ = "end of CoverTab[88500]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:221
		_go_fuzz_dep_.CoverTab[88501]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:221
		// _ = "end of CoverTab[88501]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:221
	// _ = "end of CoverTab[88490]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:221
	_go_fuzz_dep_.CoverTab[88491]++

												spn := types.PrincipalName{
		NameType:	nametype.KRB_NT_SRV_INST,
		NameString:	[]string{"krbtgt", realm},
	}

	_, tgsRep, err := cl.TGSREQGenerateAndExchange(spn, cl.Credentials.Domain(), tgt, skey, false)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:229
		_go_fuzz_dep_.CoverTab[88502]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:230
		// _ = "end of CoverTab[88502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:231
		_go_fuzz_dep_.CoverTab[88503]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:231
		// _ = "end of CoverTab[88503]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:231
	// _ = "end of CoverTab[88491]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:231
	_go_fuzz_dep_.CoverTab[88492]++
												cl.addSession(tgsRep.Ticket, tgsRep.DecryptedEncPart)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:234
	// _ = "end of CoverTab[88492]"
}

// Destroy stops the auto-renewal of all sessions and removes the sessions and cache entries from the client.
func (cl *Client) Destroy() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:238
	_go_fuzz_dep_.CoverTab[88504]++
												creds := credentials.New("", "")
												cl.sessions.destroy()
												cl.cache.clear()
												cl.Credentials = creds
												cl.Log("client destroyed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:243
	// _ = "end of CoverTab[88504]"
}

// Diagnostics runs a set of checks that the client is properly configured and writes details to the io.Writer provided.
func (cl *Client) Diagnostics(w io.Writer) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:247
	_go_fuzz_dep_.CoverTab[88505]++
												cl.Print(w)
												var errs []string
												if cl.Credentials.HasKeytab() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:250
		_go_fuzz_dep_.CoverTab[88512]++
													var loginRealmEncTypes []int32
													for _, e := range cl.Credentials.Keytab().Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:252
			_go_fuzz_dep_.CoverTab[88515]++
														if e.Principal.Realm == cl.Credentials.Realm() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:253
				_go_fuzz_dep_.CoverTab[88516]++
															loginRealmEncTypes = append(loginRealmEncTypes, e.Key.KeyType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:254
				// _ = "end of CoverTab[88516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:255
				_go_fuzz_dep_.CoverTab[88517]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:255
				// _ = "end of CoverTab[88517]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:255
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:255
			// _ = "end of CoverTab[88515]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:256
		// _ = "end of CoverTab[88512]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:256
		_go_fuzz_dep_.CoverTab[88513]++
													for _, et := range cl.Config.LibDefaults.DefaultTktEnctypeIDs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:257
			_go_fuzz_dep_.CoverTab[88518]++
														var etInKt bool
														for _, val := range loginRealmEncTypes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:259
				_go_fuzz_dep_.CoverTab[88520]++
															if val == et {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:260
					_go_fuzz_dep_.CoverTab[88521]++
																etInKt = true
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:262
					// _ = "end of CoverTab[88521]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:263
					_go_fuzz_dep_.CoverTab[88522]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:263
					// _ = "end of CoverTab[88522]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:263
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:263
				// _ = "end of CoverTab[88520]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:264
			// _ = "end of CoverTab[88518]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:264
			_go_fuzz_dep_.CoverTab[88519]++
														if !etInKt {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:265
				_go_fuzz_dep_.CoverTab[88523]++
															errs = append(errs, fmt.Sprintf("default_tkt_enctypes specifies %d but this enctype is not available in the client's keytab", et))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:266
				// _ = "end of CoverTab[88523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:267
				_go_fuzz_dep_.CoverTab[88524]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:267
				// _ = "end of CoverTab[88524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:267
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:267
			// _ = "end of CoverTab[88519]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:268
		// _ = "end of CoverTab[88513]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:268
		_go_fuzz_dep_.CoverTab[88514]++
													for _, et := range cl.Config.LibDefaults.PreferredPreauthTypes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:269
			_go_fuzz_dep_.CoverTab[88525]++
														var etInKt bool
														for _, val := range loginRealmEncTypes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:271
				_go_fuzz_dep_.CoverTab[88527]++
															if int(val) == et {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:272
					_go_fuzz_dep_.CoverTab[88528]++
																etInKt = true
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:274
					// _ = "end of CoverTab[88528]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:275
					_go_fuzz_dep_.CoverTab[88529]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:275
					// _ = "end of CoverTab[88529]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:275
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:275
				// _ = "end of CoverTab[88527]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:276
			// _ = "end of CoverTab[88525]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:276
			_go_fuzz_dep_.CoverTab[88526]++
														if !etInKt {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:277
				_go_fuzz_dep_.CoverTab[88530]++
															errs = append(errs, fmt.Sprintf("preferred_preauth_types specifies %d but this enctype is not available in the client's keytab", et))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:278
				// _ = "end of CoverTab[88530]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:279
				_go_fuzz_dep_.CoverTab[88531]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:279
				// _ = "end of CoverTab[88531]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:279
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:279
			// _ = "end of CoverTab[88526]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:280
		// _ = "end of CoverTab[88514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:281
		_go_fuzz_dep_.CoverTab[88532]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:281
		// _ = "end of CoverTab[88532]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:281
	// _ = "end of CoverTab[88505]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:281
	_go_fuzz_dep_.CoverTab[88506]++
												udpCnt, udpKDC, err := cl.Config.GetKDCs(cl.Credentials.Realm(), false)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:283
		_go_fuzz_dep_.CoverTab[88533]++
													errs = append(errs, fmt.Sprintf("error when resolving KDCs for UDP communication: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:284
		// _ = "end of CoverTab[88533]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:285
		_go_fuzz_dep_.CoverTab[88534]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:285
		// _ = "end of CoverTab[88534]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:285
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:285
	// _ = "end of CoverTab[88506]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:285
	_go_fuzz_dep_.CoverTab[88507]++
												if udpCnt < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:286
		_go_fuzz_dep_.CoverTab[88535]++
													errs = append(errs, "no KDCs resolved for communication via UDP.")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:287
		// _ = "end of CoverTab[88535]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:288
		_go_fuzz_dep_.CoverTab[88536]++
													b, _ := json.MarshalIndent(&udpKDC, "", "  ")
													fmt.Fprintf(w, "UDP KDCs: %s\n", string(b))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:290
		// _ = "end of CoverTab[88536]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:291
	// _ = "end of CoverTab[88507]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:291
	_go_fuzz_dep_.CoverTab[88508]++
												tcpCnt, tcpKDC, err := cl.Config.GetKDCs(cl.Credentials.Realm(), false)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:293
		_go_fuzz_dep_.CoverTab[88537]++
													errs = append(errs, fmt.Sprintf("error when resolving KDCs for TCP communication: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:294
		// _ = "end of CoverTab[88537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:295
		_go_fuzz_dep_.CoverTab[88538]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:295
		// _ = "end of CoverTab[88538]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:295
	// _ = "end of CoverTab[88508]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:295
	_go_fuzz_dep_.CoverTab[88509]++
												if tcpCnt < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:296
		_go_fuzz_dep_.CoverTab[88539]++
													errs = append(errs, "no KDCs resolved for communication via TCP.")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:297
		// _ = "end of CoverTab[88539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:298
		_go_fuzz_dep_.CoverTab[88540]++
													b, _ := json.MarshalIndent(&tcpKDC, "", "  ")
													fmt.Fprintf(w, "TCP KDCs: %s\n", string(b))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:300
		// _ = "end of CoverTab[88540]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:301
	// _ = "end of CoverTab[88509]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:301
	_go_fuzz_dep_.CoverTab[88510]++

												if errs == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:303
		_go_fuzz_dep_.CoverTab[88541]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:303
		return len(errs) < 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:303
		// _ = "end of CoverTab[88541]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:303
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:303
		_go_fuzz_dep_.CoverTab[88542]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:304
		// _ = "end of CoverTab[88542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:305
		_go_fuzz_dep_.CoverTab[88543]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:305
		// _ = "end of CoverTab[88543]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:305
	// _ = "end of CoverTab[88510]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:305
	_go_fuzz_dep_.CoverTab[88511]++
												err = fmt.Errorf(strings.Join(errs, "\n"))
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:307
	// _ = "end of CoverTab[88511]"
}

// Print writes the details of the client to the io.Writer provided.
func (cl *Client) Print(w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:311
	_go_fuzz_dep_.CoverTab[88544]++
												c, _ := cl.Credentials.JSON()
												fmt.Fprintf(w, "Credentials:\n%s\n", c)

												s, _ := cl.sessions.JSON()
												fmt.Fprintf(w, "TGT Sessions:\n%s\n", s)

												c, _ = cl.cache.JSON()
												fmt.Fprintf(w, "Service ticket cache:\n%s\n", c)

												s, _ = cl.settings.JSON()
												fmt.Fprintf(w, "Settings:\n%s\n", s)

												j, _ := cl.Config.JSON()
												fmt.Fprintf(w, "Krb5 config:\n%s\n", j)

												k, _ := cl.Credentials.Keytab().JSON()
												fmt.Fprintf(w, "Keytab:\n%s\n", k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:328
	// _ = "end of CoverTab[88544]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:329
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/client.go:329
var _ = _go_fuzz_dep_.CoverTab
