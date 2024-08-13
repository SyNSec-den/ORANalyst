//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:1
// Package credentials provides credentials management for Kerberos 5 authentication.
package credentials

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:2
)

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/types"
)

const (
	// AttributeKeyADCredentials assigned number for AD credentials.
	AttributeKeyADCredentials = "gokrb5AttributeKeyADCredentials"
)

// Credentials struct for a user.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:21
// Contains either a keytab, password or both.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:21
// Keytabs are used over passwords if both are defined.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:24
type Credentials struct {
	username	string
	displayName	string
	realm		string
	cname		types.PrincipalName
	keytab		*keytab.Keytab
	password	string
	attributes	map[string]interface{}
	validUntil	time.Time
	authenticated	bool
	human		bool
	authTime	time.Time
	groupMembership	map[string]bool
	sessionID	string
}

// marshalCredentials is used to enable marshaling and unmarshaling of credentials
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:40
// without having exported fields on the Credentials struct
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:42
type marshalCredentials struct {
	Username	string
	DisplayName	string
	Realm		string
	CName		types.PrincipalName	`json:"-"`
	Keytab		bool
	Password	bool
	Attributes	map[string]interface{}	`json:"-"`
	ValidUntil	time.Time
	Authenticated	bool
	Human		bool
	AuthTime	time.Time
	GroupMembership	map[string]bool	`json:"-"`
	SessionID	string
}

// ADCredentials contains information obtained from the PAC.
type ADCredentials struct {
	EffectiveName		string
	FullName		string
	UserID			int
	PrimaryGroupID		int
	LogOnTime		time.Time
	LogOffTime		time.Time
	PasswordLastSet		time.Time
	GroupMembershipSIDs	[]string
	LogonDomainName		string
	LogonDomainID		string
	LogonServer		string
}

// New creates a new Credentials instance.
func New(username string, realm string) *Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:74
	_go_fuzz_dep_.CoverTab[86642]++
													uid, err := uuid.GenerateUUID()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:76
		_go_fuzz_dep_.CoverTab[86644]++
														uid = "00unique-sess-ions-uuid-unavailable0"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:77
		// _ = "end of CoverTab[86644]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:78
		_go_fuzz_dep_.CoverTab[86645]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:78
		// _ = "end of CoverTab[86645]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:78
	// _ = "end of CoverTab[86642]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:78
	_go_fuzz_dep_.CoverTab[86643]++
													return &Credentials{
		username:		username,
		displayName:		username,
		realm:			realm,
		cname:			types.NewPrincipalName(nametype.KRB_NT_PRINCIPAL, username),
		keytab:			keytab.New(),
		attributes:		make(map[string]interface{}),
		groupMembership:	make(map[string]bool),
		sessionID:		uid,
		human:			true,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:89
	// _ = "end of CoverTab[86643]"
}

// NewFromPrincipalName creates a new Credentials instance with the user details provides as a PrincipalName type.
func NewFromPrincipalName(cname types.PrincipalName, realm string) *Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:93
	_go_fuzz_dep_.CoverTab[86646]++
													c := New(cname.PrincipalNameString(), realm)
													c.cname = cname
													return c
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:96
	// _ = "end of CoverTab[86646]"
}

// WithKeytab sets the Keytab in the Credentials struct.
func (c *Credentials) WithKeytab(kt *keytab.Keytab) *Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:100
	_go_fuzz_dep_.CoverTab[86647]++
													c.keytab = kt
													c.password = ""
													return c
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:103
	// _ = "end of CoverTab[86647]"
}

// Keytab returns the credential's Keytab.
func (c *Credentials) Keytab() *keytab.Keytab {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:107
	_go_fuzz_dep_.CoverTab[86648]++
													return c.keytab
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:108
	// _ = "end of CoverTab[86648]"
}

// HasKeytab queries if the Credentials has a keytab defined.
func (c *Credentials) HasKeytab() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:112
	_go_fuzz_dep_.CoverTab[86649]++
													if c.keytab != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:113
		_go_fuzz_dep_.CoverTab[86651]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:113
		return len(c.keytab.Entries) > 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:113
		// _ = "end of CoverTab[86651]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:113
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:113
		_go_fuzz_dep_.CoverTab[86652]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:114
		// _ = "end of CoverTab[86652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:115
		_go_fuzz_dep_.CoverTab[86653]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:115
		// _ = "end of CoverTab[86653]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:115
	// _ = "end of CoverTab[86649]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:115
	_go_fuzz_dep_.CoverTab[86650]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:116
	// _ = "end of CoverTab[86650]"
}

// WithPassword sets the password in the Credentials struct.
func (c *Credentials) WithPassword(password string) *Credentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:120
	_go_fuzz_dep_.CoverTab[86654]++
													c.password = password
													c.keytab = keytab.New()
													return c
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:123
	// _ = "end of CoverTab[86654]"
}

// Password returns the credential's password.
func (c *Credentials) Password() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:127
	_go_fuzz_dep_.CoverTab[86655]++
													return c.password
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:128
	// _ = "end of CoverTab[86655]"
}

// HasPassword queries if the Credentials has a password defined.
func (c *Credentials) HasPassword() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:132
	_go_fuzz_dep_.CoverTab[86656]++
													if c.password != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:133
		_go_fuzz_dep_.CoverTab[86658]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:134
		// _ = "end of CoverTab[86658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:135
		_go_fuzz_dep_.CoverTab[86659]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:135
		// _ = "end of CoverTab[86659]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:135
	// _ = "end of CoverTab[86656]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:135
	_go_fuzz_dep_.CoverTab[86657]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:136
	// _ = "end of CoverTab[86657]"
}

// SetValidUntil sets the expiry time of the credentials
func (c *Credentials) SetValidUntil(t time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:140
	_go_fuzz_dep_.CoverTab[86660]++
													c.validUntil = t
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:141
	// _ = "end of CoverTab[86660]"
}

// SetADCredentials adds ADCredentials attributes to the credentials
func (c *Credentials) SetADCredentials(a ADCredentials) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:145
	_go_fuzz_dep_.CoverTab[86661]++
													c.SetAttribute(AttributeKeyADCredentials, a)
													if a.FullName != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:147
		_go_fuzz_dep_.CoverTab[86664]++
														c.SetDisplayName(a.FullName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:148
		// _ = "end of CoverTab[86664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:149
		_go_fuzz_dep_.CoverTab[86665]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:149
		// _ = "end of CoverTab[86665]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:149
	// _ = "end of CoverTab[86661]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:149
	_go_fuzz_dep_.CoverTab[86662]++
													if a.EffectiveName != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:150
		_go_fuzz_dep_.CoverTab[86666]++
														c.SetUserName(a.EffectiveName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:151
		// _ = "end of CoverTab[86666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:152
		_go_fuzz_dep_.CoverTab[86667]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:152
		// _ = "end of CoverTab[86667]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:152
	// _ = "end of CoverTab[86662]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:152
	_go_fuzz_dep_.CoverTab[86663]++
													for i := range a.GroupMembershipSIDs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:153
		_go_fuzz_dep_.CoverTab[86668]++
														c.AddAuthzAttribute(a.GroupMembershipSIDs[i])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:154
		// _ = "end of CoverTab[86668]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:155
	// _ = "end of CoverTab[86663]"
}

// GetADCredentials returns ADCredentials attributes sorted in the credential
func (c *Credentials) GetADCredentials() ADCredentials {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:159
	_go_fuzz_dep_.CoverTab[86669]++
													if a, ok := c.attributes[AttributeKeyADCredentials].(ADCredentials); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:160
		_go_fuzz_dep_.CoverTab[86671]++
														return a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:161
		// _ = "end of CoverTab[86671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:162
		_go_fuzz_dep_.CoverTab[86672]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:162
		// _ = "end of CoverTab[86672]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:162
	// _ = "end of CoverTab[86669]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:162
	_go_fuzz_dep_.CoverTab[86670]++
													return ADCredentials{}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:163
	// _ = "end of CoverTab[86670]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:168
// UserName returns the credential's username.
func (c *Credentials) UserName() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:169
	_go_fuzz_dep_.CoverTab[86673]++
													return c.username
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:170
	// _ = "end of CoverTab[86673]"
}

// SetUserName sets the username value on the credential.
func (c *Credentials) SetUserName(s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:174
	_go_fuzz_dep_.CoverTab[86674]++
													c.username = s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:175
	// _ = "end of CoverTab[86674]"
}

// CName returns the credential's client principal name.
func (c *Credentials) CName() types.PrincipalName {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:179
	_go_fuzz_dep_.CoverTab[86675]++
													return c.cname
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:180
	// _ = "end of CoverTab[86675]"
}

// SetCName sets the client principal name on the credential.
func (c *Credentials) SetCName(pn types.PrincipalName) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:184
	_go_fuzz_dep_.CoverTab[86676]++
													c.cname = pn
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:185
	// _ = "end of CoverTab[86676]"
}

// Domain returns the credential's domain.
func (c *Credentials) Domain() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:189
	_go_fuzz_dep_.CoverTab[86677]++
													return c.realm
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:190
	// _ = "end of CoverTab[86677]"
}

// SetDomain sets the domain value on the credential.
func (c *Credentials) SetDomain(s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:194
	_go_fuzz_dep_.CoverTab[86678]++
													c.realm = s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:195
	// _ = "end of CoverTab[86678]"
}

// Realm returns the credential's realm. Same as the domain.
func (c *Credentials) Realm() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:199
	_go_fuzz_dep_.CoverTab[86679]++
													return c.Domain()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:200
	// _ = "end of CoverTab[86679]"
}

// SetRealm sets the realm value on the credential. Same as the domain
func (c *Credentials) SetRealm(s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:204
	_go_fuzz_dep_.CoverTab[86680]++
													c.SetDomain(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:205
	// _ = "end of CoverTab[86680]"
}

// DisplayName returns the credential's display name.
func (c *Credentials) DisplayName() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:209
	_go_fuzz_dep_.CoverTab[86681]++
													return c.displayName
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:210
	// _ = "end of CoverTab[86681]"
}

// SetDisplayName sets the display name value on the credential.
func (c *Credentials) SetDisplayName(s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:214
	_go_fuzz_dep_.CoverTab[86682]++
													c.displayName = s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:215
	// _ = "end of CoverTab[86682]"
}

// Human returns if the  credential represents a human or not.
func (c *Credentials) Human() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:219
	_go_fuzz_dep_.CoverTab[86683]++
													return c.human
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:220
	// _ = "end of CoverTab[86683]"
}

// SetHuman sets the credential as human.
func (c *Credentials) SetHuman(b bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:224
	_go_fuzz_dep_.CoverTab[86684]++
													c.human = b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:225
	// _ = "end of CoverTab[86684]"
}

// AuthTime returns the time the credential was authenticated.
func (c *Credentials) AuthTime() time.Time {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:229
	_go_fuzz_dep_.CoverTab[86685]++
													return c.authTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:230
	// _ = "end of CoverTab[86685]"
}

// SetAuthTime sets the time the credential was authenticated.
func (c *Credentials) SetAuthTime(t time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:234
	_go_fuzz_dep_.CoverTab[86686]++
													c.authTime = t
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:235
	// _ = "end of CoverTab[86686]"
}

// AuthzAttributes returns the credentials authorizing attributes.
func (c *Credentials) AuthzAttributes() []string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:239
	_go_fuzz_dep_.CoverTab[86687]++
													s := make([]string, len(c.groupMembership))
													i := 0
													for a := range c.groupMembership {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:242
		_go_fuzz_dep_.CoverTab[86689]++
														s[i] = a
														i++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:244
		// _ = "end of CoverTab[86689]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:245
	// _ = "end of CoverTab[86687]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:245
	_go_fuzz_dep_.CoverTab[86688]++
													return s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:246
	// _ = "end of CoverTab[86688]"
}

// Authenticated indicates if the credential has been successfully authenticated or not.
func (c *Credentials) Authenticated() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:250
	_go_fuzz_dep_.CoverTab[86690]++
													return c.authenticated
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:251
	// _ = "end of CoverTab[86690]"
}

// SetAuthenticated sets the credential as having been successfully authenticated.
func (c *Credentials) SetAuthenticated(b bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:255
	_go_fuzz_dep_.CoverTab[86691]++
													c.authenticated = b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:256
	// _ = "end of CoverTab[86691]"
}

// AddAuthzAttribute adds an authorization attribute to the credential.
func (c *Credentials) AddAuthzAttribute(a string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:260
	_go_fuzz_dep_.CoverTab[86692]++
													c.groupMembership[a] = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:261
	// _ = "end of CoverTab[86692]"
}

// RemoveAuthzAttribute removes an authorization attribute from the credential.
func (c *Credentials) RemoveAuthzAttribute(a string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:265
	_go_fuzz_dep_.CoverTab[86693]++
													if _, ok := c.groupMembership[a]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:266
		_go_fuzz_dep_.CoverTab[86695]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:267
		// _ = "end of CoverTab[86695]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:268
		_go_fuzz_dep_.CoverTab[86696]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:268
		// _ = "end of CoverTab[86696]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:268
	// _ = "end of CoverTab[86693]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:268
	_go_fuzz_dep_.CoverTab[86694]++
													delete(c.groupMembership, a)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:269
	// _ = "end of CoverTab[86694]"
}

// EnableAuthzAttribute toggles an authorization attribute to an enabled state on the credential.
func (c *Credentials) EnableAuthzAttribute(a string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:273
	_go_fuzz_dep_.CoverTab[86697]++
													if enabled, ok := c.groupMembership[a]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:274
		_go_fuzz_dep_.CoverTab[86698]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:274
		return !enabled
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:274
		// _ = "end of CoverTab[86698]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:274
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:274
		_go_fuzz_dep_.CoverTab[86699]++
														c.groupMembership[a] = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:275
		// _ = "end of CoverTab[86699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:276
		_go_fuzz_dep_.CoverTab[86700]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:276
		// _ = "end of CoverTab[86700]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:276
	// _ = "end of CoverTab[86697]"
}

// DisableAuthzAttribute toggles an authorization attribute to a disabled state on the credential.
func (c *Credentials) DisableAuthzAttribute(a string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:280
	_go_fuzz_dep_.CoverTab[86701]++
													if enabled, ok := c.groupMembership[a]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:281
		_go_fuzz_dep_.CoverTab[86702]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:281
		return enabled
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:281
		// _ = "end of CoverTab[86702]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:281
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:281
		_go_fuzz_dep_.CoverTab[86703]++
														c.groupMembership[a] = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:282
		// _ = "end of CoverTab[86703]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:283
		_go_fuzz_dep_.CoverTab[86704]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:283
		// _ = "end of CoverTab[86704]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:283
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:283
	// _ = "end of CoverTab[86701]"
}

// Authorized indicates if the credential has the specified authorizing attribute.
func (c *Credentials) Authorized(a string) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:287
	_go_fuzz_dep_.CoverTab[86705]++
													if enabled, ok := c.groupMembership[a]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:288
		_go_fuzz_dep_.CoverTab[86707]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:288
		return enabled
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:288
		// _ = "end of CoverTab[86707]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:288
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:288
		_go_fuzz_dep_.CoverTab[86708]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:289
		// _ = "end of CoverTab[86708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:290
		_go_fuzz_dep_.CoverTab[86709]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:290
		// _ = "end of CoverTab[86709]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:290
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:290
	// _ = "end of CoverTab[86705]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:290
	_go_fuzz_dep_.CoverTab[86706]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:291
	// _ = "end of CoverTab[86706]"
}

// SessionID returns the credential's session ID.
func (c *Credentials) SessionID() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:295
	_go_fuzz_dep_.CoverTab[86710]++
													return c.sessionID
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:296
	// _ = "end of CoverTab[86710]"
}

// Expired indicates if the credential has expired.
func (c *Credentials) Expired() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:300
	_go_fuzz_dep_.CoverTab[86711]++
													if !c.validUntil.IsZero() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:301
		_go_fuzz_dep_.CoverTab[86713]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:301
		return time.Now().UTC().After(c.validUntil)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:301
		// _ = "end of CoverTab[86713]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:301
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:301
		_go_fuzz_dep_.CoverTab[86714]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:302
		// _ = "end of CoverTab[86714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:303
		_go_fuzz_dep_.CoverTab[86715]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:303
		// _ = "end of CoverTab[86715]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:303
	// _ = "end of CoverTab[86711]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:303
	_go_fuzz_dep_.CoverTab[86712]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:304
	// _ = "end of CoverTab[86712]"
}

// ValidUntil returns the credential's valid until date
func (c *Credentials) ValidUntil() time.Time {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:308
	_go_fuzz_dep_.CoverTab[86716]++
													return c.validUntil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:309
	// _ = "end of CoverTab[86716]"
}

// Attributes returns the Credentials' attributes map.
func (c *Credentials) Attributes() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:313
	_go_fuzz_dep_.CoverTab[86717]++
													return c.attributes
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:314
	// _ = "end of CoverTab[86717]"
}

// SetAttribute sets the value of an attribute.
func (c *Credentials) SetAttribute(k string, v interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:318
	_go_fuzz_dep_.CoverTab[86718]++
													c.attributes[k] = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:319
	// _ = "end of CoverTab[86718]"
}

// SetAttributes replaces the attributes map with the one provided.
func (c *Credentials) SetAttributes(a map[string]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:323
	_go_fuzz_dep_.CoverTab[86719]++
													c.attributes = a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:324
	// _ = "end of CoverTab[86719]"
}

// RemoveAttribute deletes an attribute from the attribute map that has the key provided.
func (c *Credentials) RemoveAttribute(k string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:328
	_go_fuzz_dep_.CoverTab[86720]++
													delete(c.attributes, k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:329
	// _ = "end of CoverTab[86720]"
}

// Marshal the Credentials into a byte slice
func (c *Credentials) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:333
	_go_fuzz_dep_.CoverTab[86721]++
													gob.Register(map[string]interface{}{})
													gob.Register(ADCredentials{})
													buf := new(bytes.Buffer)
													enc := gob.NewEncoder(buf)
													mc := marshalCredentials{
		Username:		c.username,
		DisplayName:		c.displayName,
		Realm:			c.realm,
		CName:			c.cname,
		Keytab:			c.HasKeytab(),
		Password:		c.HasPassword(),
		Attributes:		c.attributes,
		ValidUntil:		c.validUntil,
		Authenticated:		c.authenticated,
		Human:			c.human,
		AuthTime:		c.authTime,
		GroupMembership:	c.groupMembership,
		SessionID:		c.sessionID,
	}
	err := enc.Encode(&mc)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:354
		_go_fuzz_dep_.CoverTab[86723]++
														return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:355
		// _ = "end of CoverTab[86723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:356
		_go_fuzz_dep_.CoverTab[86724]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:356
		// _ = "end of CoverTab[86724]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:356
	// _ = "end of CoverTab[86721]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:356
	_go_fuzz_dep_.CoverTab[86722]++
													return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:357
	// _ = "end of CoverTab[86722]"
}

// Unmarshal a byte slice into Credentials
func (c *Credentials) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:361
	_go_fuzz_dep_.CoverTab[86725]++
													gob.Register(map[string]interface{}{})
													gob.Register(ADCredentials{})
													mc := new(marshalCredentials)
													buf := bytes.NewBuffer(b)
													dec := gob.NewDecoder(buf)
													err := dec.Decode(mc)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:368
		_go_fuzz_dep_.CoverTab[86727]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:369
		// _ = "end of CoverTab[86727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:370
		_go_fuzz_dep_.CoverTab[86728]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:370
		// _ = "end of CoverTab[86728]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:370
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:370
	// _ = "end of CoverTab[86725]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:370
	_go_fuzz_dep_.CoverTab[86726]++
													c.username = mc.Username
													c.displayName = mc.DisplayName
													c.realm = mc.Realm
													c.cname = mc.CName
													c.attributes = mc.Attributes
													c.validUntil = mc.ValidUntil
													c.authenticated = mc.Authenticated
													c.human = mc.Human
													c.authTime = mc.AuthTime
													c.groupMembership = mc.GroupMembership
													c.sessionID = mc.SessionID
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:382
	// _ = "end of CoverTab[86726]"
}

// JSON return details of the Credentials in a JSON format.
func (c *Credentials) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:386
	_go_fuzz_dep_.CoverTab[86729]++
													mc := marshalCredentials{
		Username:	c.username,
		DisplayName:	c.displayName,
		Realm:		c.realm,
		CName:		c.cname,
		Keytab:		c.HasKeytab(),
		Password:	c.HasPassword(),
		ValidUntil:	c.validUntil,
		Authenticated:	c.authenticated,
		Human:		c.human,
		AuthTime:	c.authTime,
		SessionID:	c.sessionID,
	}
	b, err := json.MarshalIndent(mc, "", "  ")
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:401
		_go_fuzz_dep_.CoverTab[86731]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:402
		// _ = "end of CoverTab[86731]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:403
		_go_fuzz_dep_.CoverTab[86732]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:403
		// _ = "end of CoverTab[86732]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:403
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:403
	// _ = "end of CoverTab[86729]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:403
	_go_fuzz_dep_.CoverTab[86730]++
													return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:404
	// _ = "end of CoverTab[86730]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:405
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/credentials/credentials.go:405
var _ = _go_fuzz_dep_.CoverTab
