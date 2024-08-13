//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:1
)

import (
	"github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// TGSREQGenerateAndExchange generates the TGS_REQ and performs a TGS exchange to retrieve a ticket to the specified SPN.
func (cl *Client) TGSREQGenerateAndExchange(spn types.PrincipalName, kdcRealm string, tgt messages.Ticket, sessionKey types.EncryptionKey, renewal bool) (tgsReq messages.TGSReq, tgsRep messages.TGSRep, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:12
	_go_fuzz_dep_.CoverTab[88328]++
													tgsReq, err = messages.NewTGSReq(cl.Credentials.CName(), kdcRealm, cl.Config, tgt, sessionKey, spn, renewal)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:14
		_go_fuzz_dep_.CoverTab[88330]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.KRBMsgError, "TGS Exchange Error: failed to generate a new TGS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:15
		// _ = "end of CoverTab[88330]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:16
		_go_fuzz_dep_.CoverTab[88331]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:16
		// _ = "end of CoverTab[88331]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:16
	// _ = "end of CoverTab[88328]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:16
	_go_fuzz_dep_.CoverTab[88329]++
													return cl.TGSExchange(tgsReq, kdcRealm, tgsRep.Ticket, sessionKey, 0)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:17
	// _ = "end of CoverTab[88329]"
}

// TGSExchange exchanges the provided TGS_REQ with the KDC to retrieve a TGS_REP.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:20
// Referrals are automatically handled.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:20
// The client's cache is updated with the ticket received.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:23
func (cl *Client) TGSExchange(tgsReq messages.TGSReq, kdcRealm string, tgt messages.Ticket, sessionKey types.EncryptionKey, referral int) (messages.TGSReq, messages.TGSRep, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:23
	_go_fuzz_dep_.CoverTab[88332]++
													var tgsRep messages.TGSRep
													b, err := tgsReq.Marshal()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:26
		_go_fuzz_dep_.CoverTab[88339]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.EncodingError, "TGS Exchange Error: failed to marshal TGS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:27
		// _ = "end of CoverTab[88339]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:28
		_go_fuzz_dep_.CoverTab[88340]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:28
		// _ = "end of CoverTab[88340]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:28
	// _ = "end of CoverTab[88332]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:28
	_go_fuzz_dep_.CoverTab[88333]++
													r, err := cl.sendToKDC(b, kdcRealm)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:30
		_go_fuzz_dep_.CoverTab[88341]++
														if _, ok := err.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:31
			_go_fuzz_dep_.CoverTab[88343]++
															return tgsReq, tgsRep, krberror.Errorf(err, krberror.KDCError, "TGS Exchange Error: kerberos error response from KDC when requesting for %s", tgsReq.ReqBody.SName.PrincipalNameString())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:32
			// _ = "end of CoverTab[88343]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:33
			_go_fuzz_dep_.CoverTab[88344]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:33
			// _ = "end of CoverTab[88344]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:33
		// _ = "end of CoverTab[88341]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:33
		_go_fuzz_dep_.CoverTab[88342]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.NetworkingError, "TGS Exchange Error: issue sending TGS_REQ to KDC")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:34
		// _ = "end of CoverTab[88342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:35
		_go_fuzz_dep_.CoverTab[88345]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:35
		// _ = "end of CoverTab[88345]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:35
	// _ = "end of CoverTab[88333]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:35
	_go_fuzz_dep_.CoverTab[88334]++
													err = tgsRep.Unmarshal(r)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:37
		_go_fuzz_dep_.CoverTab[88346]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.EncodingError, "TGS Exchange Error: failed to process the TGS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:38
		// _ = "end of CoverTab[88346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:39
		_go_fuzz_dep_.CoverTab[88347]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:39
		// _ = "end of CoverTab[88347]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:39
	// _ = "end of CoverTab[88334]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:39
	_go_fuzz_dep_.CoverTab[88335]++
													err = tgsRep.DecryptEncPart(sessionKey)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:41
		_go_fuzz_dep_.CoverTab[88348]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.EncodingError, "TGS Exchange Error: failed to process the TGS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:42
		// _ = "end of CoverTab[88348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:43
		_go_fuzz_dep_.CoverTab[88349]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:43
		// _ = "end of CoverTab[88349]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:43
	// _ = "end of CoverTab[88335]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:43
	_go_fuzz_dep_.CoverTab[88336]++
													if ok, err := tgsRep.Verify(cl.Config, tgsReq); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:44
		_go_fuzz_dep_.CoverTab[88350]++
														return tgsReq, tgsRep, krberror.Errorf(err, krberror.EncodingError, "TGS Exchange Error: TGS_REP is not valid")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:45
		// _ = "end of CoverTab[88350]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:46
		_go_fuzz_dep_.CoverTab[88351]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:46
		// _ = "end of CoverTab[88351]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:46
	// _ = "end of CoverTab[88336]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:46
	_go_fuzz_dep_.CoverTab[88337]++

													if tgsRep.Ticket.SName.NameString[0] == "krbtgt" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:48
		_go_fuzz_dep_.CoverTab[88352]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:48
		return !tgsRep.Ticket.SName.Equal(tgsReq.ReqBody.SName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:48
		// _ = "end of CoverTab[88352]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:48
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:48
		_go_fuzz_dep_.CoverTab[88353]++
														if referral > 5 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:49
			_go_fuzz_dep_.CoverTab[88357]++
															return tgsReq, tgsRep, krberror.Errorf(err, krberror.KRBMsgError, "TGS Exchange Error: maximum number of referrals exceeded")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:50
			// _ = "end of CoverTab[88357]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:51
			_go_fuzz_dep_.CoverTab[88358]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:51
			// _ = "end of CoverTab[88358]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:51
		// _ = "end of CoverTab[88353]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:51
		_go_fuzz_dep_.CoverTab[88354]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:54
		cl.addSession(tgsRep.Ticket, tgsRep.DecryptedEncPart)
		realm := tgsRep.Ticket.SName.NameString[len(tgsRep.Ticket.SName.NameString)-1]
		referral++
		if types.IsFlagSet(&tgsReq.ReqBody.KDCOptions, flags.EncTktInSkey) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:57
			_go_fuzz_dep_.CoverTab[88359]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:57
			return len(tgsReq.ReqBody.AdditionalTickets) > 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:57
			// _ = "end of CoverTab[88359]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:57
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:57
			_go_fuzz_dep_.CoverTab[88360]++
															tgsReq, err = messages.NewUser2UserTGSReq(cl.Credentials.CName(), kdcRealm, cl.Config, tgt, sessionKey, tgsReq.ReqBody.SName, tgsReq.Renewal, tgsReq.ReqBody.AdditionalTickets[0])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:59
				_go_fuzz_dep_.CoverTab[88361]++
																return tgsReq, tgsRep, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:60
				// _ = "end of CoverTab[88361]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:61
				_go_fuzz_dep_.CoverTab[88362]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:61
				// _ = "end of CoverTab[88362]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:61
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:61
			// _ = "end of CoverTab[88360]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:62
			_go_fuzz_dep_.CoverTab[88363]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:62
			// _ = "end of CoverTab[88363]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:62
		// _ = "end of CoverTab[88354]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:62
		_go_fuzz_dep_.CoverTab[88355]++
														tgsReq, err = messages.NewTGSReq(cl.Credentials.CName(), realm, cl.Config, tgsRep.Ticket, tgsRep.DecryptedEncPart.Key, tgsReq.ReqBody.SName, tgsReq.Renewal)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:64
			_go_fuzz_dep_.CoverTab[88364]++
															return tgsReq, tgsRep, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:65
			// _ = "end of CoverTab[88364]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:66
			_go_fuzz_dep_.CoverTab[88365]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:66
			// _ = "end of CoverTab[88365]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:66
		// _ = "end of CoverTab[88355]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:66
		_go_fuzz_dep_.CoverTab[88356]++
														return cl.TGSExchange(tgsReq, realm, tgsRep.Ticket, tgsRep.DecryptedEncPart.Key, referral)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:67
		// _ = "end of CoverTab[88356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:68
		_go_fuzz_dep_.CoverTab[88366]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:68
		// _ = "end of CoverTab[88366]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:68
	// _ = "end of CoverTab[88337]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:68
	_go_fuzz_dep_.CoverTab[88338]++
													cl.cache.addEntry(
		tgsRep.Ticket,
		tgsRep.DecryptedEncPart.AuthTime,
		tgsRep.DecryptedEncPart.StartTime,
		tgsRep.DecryptedEncPart.EndTime,
		tgsRep.DecryptedEncPart.RenewTill,
		tgsRep.DecryptedEncPart.Key,
	)
													cl.Log("ticket added to cache for %s (EndTime: %v)", tgsRep.Ticket.SName.PrincipalNameString(), tgsRep.DecryptedEncPart.EndTime)
													return tgsReq, tgsRep, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:78
	// _ = "end of CoverTab[88338]"
}

// GetServiceTicket makes a request to get a service ticket for the SPN specified
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:81
// SPN format: <SERVICE>/<FQDN> Eg. HTTP/www.example.com
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:81
// The ticket will be added to the client's ticket cache
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:84
func (cl *Client) GetServiceTicket(spn string) (messages.Ticket, types.EncryptionKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:84
	_go_fuzz_dep_.CoverTab[88367]++
													var tkt messages.Ticket
													var skey types.EncryptionKey
													if tkt, skey, ok := cl.GetCachedTicket(spn); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:87
		_go_fuzz_dep_.CoverTab[88371]++

														return tkt, skey, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:89
		// _ = "end of CoverTab[88371]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:90
		_go_fuzz_dep_.CoverTab[88372]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:90
		// _ = "end of CoverTab[88372]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:90
	// _ = "end of CoverTab[88367]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:90
	_go_fuzz_dep_.CoverTab[88368]++
													princ := types.NewPrincipalName(nametype.KRB_NT_PRINCIPAL, spn)
													realm := cl.Config.ResolveRealm(princ.NameString[len(princ.NameString)-1])

													tgt, skey, err := cl.sessionTGT(realm)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:95
		_go_fuzz_dep_.CoverTab[88373]++
														return tkt, skey, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:96
		// _ = "end of CoverTab[88373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:97
		_go_fuzz_dep_.CoverTab[88374]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:97
		// _ = "end of CoverTab[88374]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:97
	// _ = "end of CoverTab[88368]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:97
	_go_fuzz_dep_.CoverTab[88369]++
													_, tgsRep, err := cl.TGSREQGenerateAndExchange(princ, realm, tgt, skey, false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:99
		_go_fuzz_dep_.CoverTab[88375]++
														return tkt, skey, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:100
		// _ = "end of CoverTab[88375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:101
		_go_fuzz_dep_.CoverTab[88376]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:101
		// _ = "end of CoverTab[88376]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:101
	// _ = "end of CoverTab[88369]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:101
	_go_fuzz_dep_.CoverTab[88370]++
													return tgsRep.Ticket, tgsRep.DecryptedEncPart.Key, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:102
	// _ = "end of CoverTab[88370]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:103
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/TGSExchange.go:103
var _ = _go_fuzz_dep_.CoverTab
