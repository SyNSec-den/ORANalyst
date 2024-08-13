//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:1
)

import (
	"fmt"

	"github.com/jcmturner/gokrb5/v8/kadmin"
	"github.com/jcmturner/gokrb5/v8/messages"
)

// Kpasswd server response codes.
const (
	KRB5_KPASSWD_SUCCESS			= 0
	KRB5_KPASSWD_MALFORMED			= 1
	KRB5_KPASSWD_HARDERROR			= 2
	KRB5_KPASSWD_AUTHERROR			= 3
	KRB5_KPASSWD_SOFTERROR			= 4
	KRB5_KPASSWD_ACCESSDENIED		= 5
	KRB5_KPASSWD_BAD_VERSION		= 6
	KRB5_KPASSWD_INITIAL_FLAG_NEEDED	= 7
)

// ChangePasswd changes the password of the client to the value provided.
func (cl *Client) ChangePasswd(newPasswd string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:23
	_go_fuzz_dep_.CoverTab[88653]++
												ASReq, err := messages.NewASReqForChgPasswd(cl.Credentials.Domain(), cl.Config, cl.Credentials.CName())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:25
		_go_fuzz_dep_.CoverTab[88660]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:26
		// _ = "end of CoverTab[88660]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:27
		_go_fuzz_dep_.CoverTab[88661]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:27
		// _ = "end of CoverTab[88661]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:27
	// _ = "end of CoverTab[88653]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:27
	_go_fuzz_dep_.CoverTab[88654]++
												ASRep, err := cl.ASExchange(cl.Credentials.Domain(), ASReq, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:29
		_go_fuzz_dep_.CoverTab[88662]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:30
		// _ = "end of CoverTab[88662]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:31
		_go_fuzz_dep_.CoverTab[88663]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:31
		// _ = "end of CoverTab[88663]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:31
	// _ = "end of CoverTab[88654]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:31
	_go_fuzz_dep_.CoverTab[88655]++

												msg, key, err := kadmin.ChangePasswdMsg(cl.Credentials.CName(), cl.Credentials.Domain(), newPasswd, ASRep.Ticket, ASRep.DecryptedEncPart.Key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:34
		_go_fuzz_dep_.CoverTab[88664]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:35
		// _ = "end of CoverTab[88664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:36
		_go_fuzz_dep_.CoverTab[88665]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:36
		// _ = "end of CoverTab[88665]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:36
	// _ = "end of CoverTab[88655]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:36
	_go_fuzz_dep_.CoverTab[88656]++
												r, err := cl.sendToKPasswd(msg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:38
		_go_fuzz_dep_.CoverTab[88666]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:39
		// _ = "end of CoverTab[88666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:40
		_go_fuzz_dep_.CoverTab[88667]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:40
		// _ = "end of CoverTab[88667]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:40
	// _ = "end of CoverTab[88656]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:40
	_go_fuzz_dep_.CoverTab[88657]++
												err = r.Decrypt(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:42
		_go_fuzz_dep_.CoverTab[88668]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:43
		// _ = "end of CoverTab[88668]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:44
		_go_fuzz_dep_.CoverTab[88669]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:44
		// _ = "end of CoverTab[88669]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:44
	// _ = "end of CoverTab[88657]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:44
	_go_fuzz_dep_.CoverTab[88658]++
												if r.ResultCode != KRB5_KPASSWD_SUCCESS {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:45
		_go_fuzz_dep_.CoverTab[88670]++
													return false, fmt.Errorf("error response from kadmin: code: %d; result: %s; krberror: %v", r.ResultCode, r.Result, r.KRBError)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:46
		// _ = "end of CoverTab[88670]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:47
		_go_fuzz_dep_.CoverTab[88671]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:47
		// _ = "end of CoverTab[88671]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:47
	// _ = "end of CoverTab[88658]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:47
	_go_fuzz_dep_.CoverTab[88659]++
												cl.Credentials.WithPassword(newPasswd)
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:49
	// _ = "end of CoverTab[88659]"
}

func (cl *Client) sendToKPasswd(msg kadmin.Request) (r kadmin.Reply, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:52
	_go_fuzz_dep_.CoverTab[88672]++
												_, kps, err := cl.Config.GetKpasswdServers(cl.Credentials.Domain(), true)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:54
		_go_fuzz_dep_.CoverTab[88676]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:55
		// _ = "end of CoverTab[88676]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:56
		_go_fuzz_dep_.CoverTab[88677]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:56
		// _ = "end of CoverTab[88677]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:56
	// _ = "end of CoverTab[88672]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:56
	_go_fuzz_dep_.CoverTab[88673]++
												b, err := msg.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:58
		_go_fuzz_dep_.CoverTab[88678]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:59
		// _ = "end of CoverTab[88678]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:60
		_go_fuzz_dep_.CoverTab[88679]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:60
		// _ = "end of CoverTab[88679]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:60
	// _ = "end of CoverTab[88673]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:60
	_go_fuzz_dep_.CoverTab[88674]++
												var rb []byte
												if len(b) <= cl.Config.LibDefaults.UDPPreferenceLimit {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:62
		_go_fuzz_dep_.CoverTab[88680]++
													rb, err = dialSendUDP(kps, b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:64
			_go_fuzz_dep_.CoverTab[88681]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:65
			// _ = "end of CoverTab[88681]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:66
			_go_fuzz_dep_.CoverTab[88682]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:66
			// _ = "end of CoverTab[88682]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:66
		// _ = "end of CoverTab[88680]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:67
		_go_fuzz_dep_.CoverTab[88683]++
													rb, err = dialSendTCP(kps, b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:69
			_go_fuzz_dep_.CoverTab[88684]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:70
			// _ = "end of CoverTab[88684]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:71
			_go_fuzz_dep_.CoverTab[88685]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:71
			// _ = "end of CoverTab[88685]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:71
		// _ = "end of CoverTab[88683]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:72
	// _ = "end of CoverTab[88674]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:72
	_go_fuzz_dep_.CoverTab[88675]++
												err = r.Unmarshal(rb)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:74
	// _ = "end of CoverTab[88675]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:75
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/passwd.go:75
var _ = _go_fuzz_dep_.CoverTab
