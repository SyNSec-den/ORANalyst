//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:1
// Package kadmin provides Kerberos administration capabilities.
package kadmin

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:2
)

import (
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// ChangePasswdMsg generate a change password request and also return the key needed to decrypt the reply.
func ChangePasswdMsg(cname types.PrincipalName, realm, password string, tkt messages.Ticket, sessionKey types.EncryptionKey) (r Request, k types.EncryptionKey, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:12
	_go_fuzz_dep_.CoverTab[88217]++

												chgpasswd := ChangePasswdData{
		NewPasswd:	[]byte(password),
		TargName:	cname,
		TargRealm:	realm,
	}
	chpwdb, err := chgpasswd.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:20
		_go_fuzz_dep_.CoverTab[88224]++
													err = krberror.Errorf(err, krberror.KRBMsgError, "error marshaling change passwd data")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:22
		// _ = "end of CoverTab[88224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:23
		_go_fuzz_dep_.CoverTab[88225]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:23
		// _ = "end of CoverTab[88225]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:23
	// _ = "end of CoverTab[88217]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:23
	_go_fuzz_dep_.CoverTab[88218]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:26
	auth, err := types.NewAuthenticator(realm, cname)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:27
		_go_fuzz_dep_.CoverTab[88226]++
													err = krberror.Errorf(err, krberror.KRBMsgError, "error generating new authenticator")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:29
		// _ = "end of CoverTab[88226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:30
		_go_fuzz_dep_.CoverTab[88227]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:30
		// _ = "end of CoverTab[88227]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:30
	// _ = "end of CoverTab[88218]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:30
	_go_fuzz_dep_.CoverTab[88219]++
												etype, err := crypto.GetEtype(sessionKey.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:32
		_go_fuzz_dep_.CoverTab[88228]++
													err = krberror.Errorf(err, krberror.KRBMsgError, "error generating subkey etype")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:34
		// _ = "end of CoverTab[88228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:35
		_go_fuzz_dep_.CoverTab[88229]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:35
		// _ = "end of CoverTab[88229]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:35
	// _ = "end of CoverTab[88219]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:35
	_go_fuzz_dep_.CoverTab[88220]++
												err = auth.GenerateSeqNumberAndSubKey(etype.GetETypeID(), etype.GetKeyByteSize())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:37
		_go_fuzz_dep_.CoverTab[88230]++
													err = krberror.Errorf(err, krberror.KRBMsgError, "error generating subkey")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:39
		// _ = "end of CoverTab[88230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:40
		_go_fuzz_dep_.CoverTab[88231]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:40
		// _ = "end of CoverTab[88231]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:40
	// _ = "end of CoverTab[88220]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:40
	_go_fuzz_dep_.CoverTab[88221]++
												k = auth.SubKey

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:44
	APreq, err := messages.NewAPReq(tkt, sessionKey, auth)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:45
		_go_fuzz_dep_.CoverTab[88232]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:46
		// _ = "end of CoverTab[88232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:47
		_go_fuzz_dep_.CoverTab[88233]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:47
		// _ = "end of CoverTab[88233]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:47
	// _ = "end of CoverTab[88221]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:47
	_go_fuzz_dep_.CoverTab[88222]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:50
	kp := messages.EncKrbPrivPart{
		UserData:	chpwdb,
		Timestamp:	auth.CTime,
		Usec:		auth.Cusec,
		SequenceNumber:	auth.SeqNumber,
	}
	kpriv := messages.NewKRBPriv(kp)
	err = kpriv.EncryptEncPart(k)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:58
		_go_fuzz_dep_.CoverTab[88234]++
													err = krberror.Errorf(err, krberror.EncryptingError, "error encrypting change passwd data")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:60
		// _ = "end of CoverTab[88234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:61
		_go_fuzz_dep_.CoverTab[88235]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:61
		// _ = "end of CoverTab[88235]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:61
	// _ = "end of CoverTab[88222]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:61
	_go_fuzz_dep_.CoverTab[88223]++

												r = Request{
		APREQ:		APreq,
		KRBPriv:	kpriv,
	}
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:67
	// _ = "end of CoverTab[88223]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:68
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/passwd.go:68
var _ = _go_fuzz_dep_.CoverTab
