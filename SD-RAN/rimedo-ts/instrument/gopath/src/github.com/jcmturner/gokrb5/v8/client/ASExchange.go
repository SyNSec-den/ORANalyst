//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:1
)

import (
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/patype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// ASExchange performs an AS exchange for the client to retrieve a TGT.
func (cl *Client) ASExchange(realm string, ASReq messages.ASReq, referral int) (messages.ASRep, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:15
	_go_fuzz_dep_.CoverTab[88236]++
													if ok, err := cl.IsConfigured(); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:16
		_go_fuzz_dep_.CoverTab[88243]++
														return messages.ASRep{}, krberror.Errorf(err, krberror.ConfigError, "AS Exchange cannot be performed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:17
		// _ = "end of CoverTab[88243]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:18
		_go_fuzz_dep_.CoverTab[88244]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:18
		// _ = "end of CoverTab[88244]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:18
	// _ = "end of CoverTab[88236]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:18
	_go_fuzz_dep_.CoverTab[88237]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:21
	err := setPAData(cl, nil, &ASReq)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:22
		_go_fuzz_dep_.CoverTab[88245]++
														return messages.ASRep{}, krberror.Errorf(err, krberror.KRBMsgError, "AS Exchange Error: issue with setting PAData on AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:23
		// _ = "end of CoverTab[88245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:24
		_go_fuzz_dep_.CoverTab[88246]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:24
		// _ = "end of CoverTab[88246]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:24
	// _ = "end of CoverTab[88237]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:24
	_go_fuzz_dep_.CoverTab[88238]++

													b, err := ASReq.Marshal()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:27
		_go_fuzz_dep_.CoverTab[88247]++
														return messages.ASRep{}, krberror.Errorf(err, krberror.EncodingError, "AS Exchange Error: failed marshaling AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:28
		// _ = "end of CoverTab[88247]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:29
		_go_fuzz_dep_.CoverTab[88248]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:29
		// _ = "end of CoverTab[88248]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:29
	// _ = "end of CoverTab[88238]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:29
	_go_fuzz_dep_.CoverTab[88239]++
													var ASRep messages.ASRep

													rb, err := cl.sendToKDC(b, realm)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:33
		_go_fuzz_dep_.CoverTab[88249]++
														if e, ok := err.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:34
			_go_fuzz_dep_.CoverTab[88250]++
															switch e.ErrorCode {
			case errorcode.KDC_ERR_PREAUTH_REQUIRED, errorcode.KDC_ERR_PREAUTH_FAILED:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:36
				_go_fuzz_dep_.CoverTab[88251]++

																cl.settings.assumePreAuthentication = true
																err = setPAData(cl, &e, &ASReq)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:40
					_go_fuzz_dep_.CoverTab[88257]++
																	return messages.ASRep{}, krberror.Errorf(err, krberror.KRBMsgError, "AS Exchange Error: failed setting AS_REQ PAData for pre-authentication required")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:41
					// _ = "end of CoverTab[88257]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:42
					_go_fuzz_dep_.CoverTab[88258]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:42
					// _ = "end of CoverTab[88258]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:42
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:42
				// _ = "end of CoverTab[88251]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:42
				_go_fuzz_dep_.CoverTab[88252]++
																b, err := ASReq.Marshal()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:44
					_go_fuzz_dep_.CoverTab[88259]++
																	return messages.ASRep{}, krberror.Errorf(err, krberror.EncodingError, "AS Exchange Error: failed marshaling AS_REQ with PAData")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:45
					// _ = "end of CoverTab[88259]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:46
					_go_fuzz_dep_.CoverTab[88260]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:46
					// _ = "end of CoverTab[88260]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:46
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:46
				// _ = "end of CoverTab[88252]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:46
				_go_fuzz_dep_.CoverTab[88253]++
																rb, err = cl.sendToKDC(b, realm)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:48
					_go_fuzz_dep_.CoverTab[88261]++
																	if _, ok := err.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:49
						_go_fuzz_dep_.CoverTab[88263]++
																		return messages.ASRep{}, krberror.Errorf(err, krberror.KDCError, "AS Exchange Error: kerberos error response from KDC")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:50
						// _ = "end of CoverTab[88263]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:51
						_go_fuzz_dep_.CoverTab[88264]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:51
						// _ = "end of CoverTab[88264]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:51
					}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:51
					// _ = "end of CoverTab[88261]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:51
					_go_fuzz_dep_.CoverTab[88262]++
																	return messages.ASRep{}, krberror.Errorf(err, krberror.NetworkingError, "AS Exchange Error: failed sending AS_REQ to KDC")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:52
					// _ = "end of CoverTab[88262]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:53
					_go_fuzz_dep_.CoverTab[88265]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:53
					// _ = "end of CoverTab[88265]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:53
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:53
				// _ = "end of CoverTab[88253]"
			case errorcode.KDC_ERR_WRONG_REALM:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:54
				_go_fuzz_dep_.CoverTab[88254]++

																if referral > 5 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:56
					_go_fuzz_dep_.CoverTab[88266]++
																	return messages.ASRep{}, krberror.Errorf(err, krberror.KRBMsgError, "maximum number of client referrals exceeded")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:57
					// _ = "end of CoverTab[88266]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:58
					_go_fuzz_dep_.CoverTab[88267]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:58
					// _ = "end of CoverTab[88267]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:58
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:58
				// _ = "end of CoverTab[88254]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:58
				_go_fuzz_dep_.CoverTab[88255]++
																referral++
																return cl.ASExchange(e.CRealm, ASReq, referral)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:60
				// _ = "end of CoverTab[88255]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:61
				_go_fuzz_dep_.CoverTab[88256]++
																return messages.ASRep{}, krberror.Errorf(err, krberror.KDCError, "AS Exchange Error: kerberos error response from KDC")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:62
				// _ = "end of CoverTab[88256]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:63
			// _ = "end of CoverTab[88250]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:64
			_go_fuzz_dep_.CoverTab[88268]++
															return messages.ASRep{}, krberror.Errorf(err, krberror.NetworkingError, "AS Exchange Error: failed sending AS_REQ to KDC")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:65
			// _ = "end of CoverTab[88268]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:66
		// _ = "end of CoverTab[88249]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:67
		_go_fuzz_dep_.CoverTab[88269]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:67
		// _ = "end of CoverTab[88269]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:67
	// _ = "end of CoverTab[88239]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:67
	_go_fuzz_dep_.CoverTab[88240]++
													err = ASRep.Unmarshal(rb)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:69
		_go_fuzz_dep_.CoverTab[88270]++
														return messages.ASRep{}, krberror.Errorf(err, krberror.EncodingError, "AS Exchange Error: failed to process the AS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:70
		// _ = "end of CoverTab[88270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:71
		_go_fuzz_dep_.CoverTab[88271]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:71
		// _ = "end of CoverTab[88271]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:71
	// _ = "end of CoverTab[88240]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:71
	_go_fuzz_dep_.CoverTab[88241]++
													if ok, err := ASRep.Verify(cl.Config, cl.Credentials, ASReq); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:72
		_go_fuzz_dep_.CoverTab[88272]++
														return messages.ASRep{}, krberror.Errorf(err, krberror.KRBMsgError, "AS Exchange Error: AS_REP is not valid or client password/keytab incorrect")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:73
		// _ = "end of CoverTab[88272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:74
		_go_fuzz_dep_.CoverTab[88273]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:74
		// _ = "end of CoverTab[88273]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:74
	// _ = "end of CoverTab[88241]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:74
	_go_fuzz_dep_.CoverTab[88242]++
													return ASRep, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:75
	// _ = "end of CoverTab[88242]"
}

// setPAData adds pre-authentication data to the AS_REQ.
func setPAData(cl *Client, krberr *messages.KRBError, ASReq *messages.ASReq) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:79
	_go_fuzz_dep_.CoverTab[88274]++
													if !cl.settings.DisablePAFXFAST() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:80
		_go_fuzz_dep_.CoverTab[88277]++
														pa := types.PAData{PADataType: patype.PA_REQ_ENC_PA_REP}
														ASReq.PAData = append(ASReq.PAData, pa)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:82
		// _ = "end of CoverTab[88277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:83
		_go_fuzz_dep_.CoverTab[88278]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:83
		// _ = "end of CoverTab[88278]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:83
	// _ = "end of CoverTab[88274]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:83
	_go_fuzz_dep_.CoverTab[88275]++
													if cl.settings.AssumePreAuthentication() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:84
		_go_fuzz_dep_.CoverTab[88279]++
		// Identify the etype to use to encrypt the PA Data
		var et etype.EType
		var err error
		var key types.EncryptionKey
		var kvno int
		if krberr == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:90
			_go_fuzz_dep_.CoverTab[88285]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:93
			etn := cl.settings.preAuthEType
			if etn == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:94
				_go_fuzz_dep_.CoverTab[88288]++
																etn = int32(cl.Config.LibDefaults.PreferredPreauthTypes[0])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:95
				// _ = "end of CoverTab[88288]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:96
				_go_fuzz_dep_.CoverTab[88289]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:96
				// _ = "end of CoverTab[88289]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:96
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:96
			// _ = "end of CoverTab[88285]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:96
			_go_fuzz_dep_.CoverTab[88286]++
															et, err = crypto.GetEtype(etn)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:98
				_go_fuzz_dep_.CoverTab[88290]++
																return krberror.Errorf(err, krberror.EncryptingError, "error getting etype for pre-auth encryption")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:99
				// _ = "end of CoverTab[88290]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:100
				_go_fuzz_dep_.CoverTab[88291]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:100
				// _ = "end of CoverTab[88291]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:100
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:100
			// _ = "end of CoverTab[88286]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:100
			_go_fuzz_dep_.CoverTab[88287]++
															key, kvno, err = cl.Key(et, 0, nil)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:102
				_go_fuzz_dep_.CoverTab[88292]++
																return krberror.Errorf(err, krberror.EncryptingError, "error getting key from credentials")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:103
				// _ = "end of CoverTab[88292]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:104
				_go_fuzz_dep_.CoverTab[88293]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:104
				// _ = "end of CoverTab[88293]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:104
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:104
			// _ = "end of CoverTab[88287]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:105
			_go_fuzz_dep_.CoverTab[88294]++

															et, err = preAuthEType(krberr)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:108
				_go_fuzz_dep_.CoverTab[88296]++
																return krberror.Errorf(err, krberror.EncryptingError, "error getting etype for pre-auth encryption")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:109
				// _ = "end of CoverTab[88296]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:110
				_go_fuzz_dep_.CoverTab[88297]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:110
				// _ = "end of CoverTab[88297]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:110
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:110
			// _ = "end of CoverTab[88294]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:110
			_go_fuzz_dep_.CoverTab[88295]++
															cl.settings.preAuthEType = et.GetETypeID()
															key, kvno, err = cl.Key(et, 0, krberr)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:113
				_go_fuzz_dep_.CoverTab[88298]++
																return krberror.Errorf(err, krberror.EncryptingError, "error getting key from credentials")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:114
				// _ = "end of CoverTab[88298]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:115
				_go_fuzz_dep_.CoverTab[88299]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:115
				// _ = "end of CoverTab[88299]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:115
			// _ = "end of CoverTab[88295]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:116
		// _ = "end of CoverTab[88279]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:116
		_go_fuzz_dep_.CoverTab[88280]++

														paTSb, err := types.GetPAEncTSEncAsnMarshalled()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:119
			_go_fuzz_dep_.CoverTab[88300]++
															return krberror.Errorf(err, krberror.KRBMsgError, "error creating PAEncTSEnc for Pre-Authentication")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:120
			// _ = "end of CoverTab[88300]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:121
			_go_fuzz_dep_.CoverTab[88301]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:121
			// _ = "end of CoverTab[88301]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:121
		// _ = "end of CoverTab[88280]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:121
		_go_fuzz_dep_.CoverTab[88281]++
														paEncTS, err := crypto.GetEncryptedData(paTSb, key, keyusage.AS_REQ_PA_ENC_TIMESTAMP, kvno)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:123
			_go_fuzz_dep_.CoverTab[88302]++
															return krberror.Errorf(err, krberror.EncryptingError, "error encrypting pre-authentication timestamp")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:124
			// _ = "end of CoverTab[88302]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:125
			_go_fuzz_dep_.CoverTab[88303]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:125
			// _ = "end of CoverTab[88303]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:125
		// _ = "end of CoverTab[88281]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:125
		_go_fuzz_dep_.CoverTab[88282]++
														pb, err := paEncTS.Marshal()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:127
			_go_fuzz_dep_.CoverTab[88304]++
															return krberror.Errorf(err, krberror.EncodingError, "error marshaling the PAEncTSEnc encrypted data")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:128
			// _ = "end of CoverTab[88304]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:129
			_go_fuzz_dep_.CoverTab[88305]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:129
			// _ = "end of CoverTab[88305]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:129
		// _ = "end of CoverTab[88282]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:129
		_go_fuzz_dep_.CoverTab[88283]++
														pa := types.PAData{
			PADataType:	patype.PA_ENC_TIMESTAMP,
			PADataValue:	pb,
		}

		for i, pa := range ASReq.PAData {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:135
			_go_fuzz_dep_.CoverTab[88306]++
															if pa.PADataType == patype.PA_ENC_TIMESTAMP {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:136
				_go_fuzz_dep_.CoverTab[88307]++
																ASReq.PAData[i] = ASReq.PAData[len(ASReq.PAData)-1]
																ASReq.PAData = ASReq.PAData[:len(ASReq.PAData)-1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:138
				// _ = "end of CoverTab[88307]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:139
				_go_fuzz_dep_.CoverTab[88308]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:139
				// _ = "end of CoverTab[88308]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:139
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:139
			// _ = "end of CoverTab[88306]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:140
		// _ = "end of CoverTab[88283]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:140
		_go_fuzz_dep_.CoverTab[88284]++
														ASReq.PAData = append(ASReq.PAData, pa)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:141
		// _ = "end of CoverTab[88284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:142
		_go_fuzz_dep_.CoverTab[88309]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:142
		// _ = "end of CoverTab[88309]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:142
	// _ = "end of CoverTab[88275]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:142
	_go_fuzz_dep_.CoverTab[88276]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:143
	// _ = "end of CoverTab[88276]"
}

// preAuthEType establishes what encryption type to use for pre-authentication from the KRBError returned from the KDC.
func preAuthEType(krberr *messages.KRBError) (etype etype.EType, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:147
	_go_fuzz_dep_.CoverTab[88310]++
	//RFC 4120 5.2.7.5 covers the preference order of ETYPE-INFO2 and ETYPE-INFO.
	var etypeID int32
	var pas types.PADataSequence
	e := pas.Unmarshal(krberr.EData)
	if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:152
		_go_fuzz_dep_.CoverTab[88314]++
														err = krberror.Errorf(e, krberror.EncodingError, "error unmashalling KRBError data")
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:154
		// _ = "end of CoverTab[88314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:155
		_go_fuzz_dep_.CoverTab[88315]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:155
		// _ = "end of CoverTab[88315]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:155
	// _ = "end of CoverTab[88310]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:155
	_go_fuzz_dep_.CoverTab[88311]++
Loop:
	for _, pa := range pas {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:157
		_go_fuzz_dep_.CoverTab[88316]++
														switch pa.PADataType {
		case patype.PA_ETYPE_INFO2:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:159
			_go_fuzz_dep_.CoverTab[88317]++
															info, e := pa.GetETypeInfo2()
															if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:161
				_go_fuzz_dep_.CoverTab[88322]++
																err = krberror.Errorf(e, krberror.EncodingError, "error unmashalling ETYPE-INFO2 data")
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:163
				// _ = "end of CoverTab[88322]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:164
				_go_fuzz_dep_.CoverTab[88323]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:164
				// _ = "end of CoverTab[88323]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:164
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:164
			// _ = "end of CoverTab[88317]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:164
			_go_fuzz_dep_.CoverTab[88318]++
															etypeID = info[0].EType
															break Loop
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:166
			// _ = "end of CoverTab[88318]"
		case patype.PA_ETYPE_INFO:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:167
			_go_fuzz_dep_.CoverTab[88319]++
															info, e := pa.GetETypeInfo()
															if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:169
				_go_fuzz_dep_.CoverTab[88324]++
																err = krberror.Errorf(e, krberror.EncodingError, "error unmashalling ETYPE-INFO data")
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:171
				// _ = "end of CoverTab[88324]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:172
				_go_fuzz_dep_.CoverTab[88325]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:172
				// _ = "end of CoverTab[88325]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:172
			// _ = "end of CoverTab[88319]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:172
			_go_fuzz_dep_.CoverTab[88320]++
															etypeID = info[0].EType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:173
			// _ = "end of CoverTab[88320]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:173
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:173
			_go_fuzz_dep_.CoverTab[88321]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:173
			// _ = "end of CoverTab[88321]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:174
		// _ = "end of CoverTab[88316]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:175
	// _ = "end of CoverTab[88311]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:175
	_go_fuzz_dep_.CoverTab[88312]++
													etype, e = crypto.GetEtype(etypeID)
													if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:177
		_go_fuzz_dep_.CoverTab[88326]++
														err = krberror.Errorf(e, krberror.EncryptingError, "error creating etype")
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:179
		// _ = "end of CoverTab[88326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:180
		_go_fuzz_dep_.CoverTab[88327]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:180
		// _ = "end of CoverTab[88327]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:180
	// _ = "end of CoverTab[88312]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:180
	_go_fuzz_dep_.CoverTab[88313]++
													return etype, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:181
	// _ = "end of CoverTab[88313]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:182
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/ASExchange.go:182
var _ = _go_fuzz_dep_.CoverTab
