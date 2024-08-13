//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:1
)

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/messages"
)

// SendToKDC performs network actions to send data to the KDC.
func (cl *Client) sendToKDC(b []byte, realm string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:17
	_go_fuzz_dep_.CoverTab[88545]++
												var rb []byte
												if cl.Config.LibDefaults.UDPPreferenceLimit == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:19
		_go_fuzz_dep_.CoverTab[88549]++

													rb, errtcp := cl.sendKDCTCP(realm, b)
													if errtcp != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:22
			_go_fuzz_dep_.CoverTab[88551]++
														if e, ok := errtcp.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:23
				_go_fuzz_dep_.CoverTab[88553]++
															return rb, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:24
				// _ = "end of CoverTab[88553]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:25
				_go_fuzz_dep_.CoverTab[88554]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:25
				// _ = "end of CoverTab[88554]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:25
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:25
			// _ = "end of CoverTab[88551]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:25
			_go_fuzz_dep_.CoverTab[88552]++
														return rb, fmt.Errorf("communication error with KDC via TCP: %v", errtcp)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:26
			// _ = "end of CoverTab[88552]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:27
			_go_fuzz_dep_.CoverTab[88555]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:27
			// _ = "end of CoverTab[88555]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:27
		// _ = "end of CoverTab[88549]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:27
		_go_fuzz_dep_.CoverTab[88550]++
													return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:28
		// _ = "end of CoverTab[88550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:29
		_go_fuzz_dep_.CoverTab[88556]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:29
		// _ = "end of CoverTab[88556]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:29
	// _ = "end of CoverTab[88545]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:29
	_go_fuzz_dep_.CoverTab[88546]++
												if len(b) <= cl.Config.LibDefaults.UDPPreferenceLimit {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:30
		_go_fuzz_dep_.CoverTab[88557]++

													rb, errudp := cl.sendKDCUDP(realm, b)
													if errudp != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:33
			_go_fuzz_dep_.CoverTab[88559]++
														if e, ok := errudp.(messages.KRBError); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:34
				_go_fuzz_dep_.CoverTab[88562]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:34
				return e.ErrorCode != errorcode.KRB_ERR_RESPONSE_TOO_BIG
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:34
				// _ = "end of CoverTab[88562]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:34
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:34
				_go_fuzz_dep_.CoverTab[88563]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:37
				return rb, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:37
				// _ = "end of CoverTab[88563]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:38
				_go_fuzz_dep_.CoverTab[88564]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:38
				// _ = "end of CoverTab[88564]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:38
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:38
			// _ = "end of CoverTab[88559]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:38
			_go_fuzz_dep_.CoverTab[88560]++

														r, errtcp := cl.sendKDCTCP(realm, b)
														if errtcp != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:41
				_go_fuzz_dep_.CoverTab[88565]++
															if e, ok := errtcp.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:42
					_go_fuzz_dep_.CoverTab[88567]++

																return r, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:44
					// _ = "end of CoverTab[88567]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:45
					_go_fuzz_dep_.CoverTab[88568]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:45
					// _ = "end of CoverTab[88568]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:45
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:45
				// _ = "end of CoverTab[88565]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:45
				_go_fuzz_dep_.CoverTab[88566]++
															return r, fmt.Errorf("failed to communicate with KDC. Attempts made with UDP (%v) and then TCP (%v)", errudp, errtcp)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:46
				// _ = "end of CoverTab[88566]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:47
				_go_fuzz_dep_.CoverTab[88569]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:47
				// _ = "end of CoverTab[88569]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:47
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:47
			// _ = "end of CoverTab[88560]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:47
			_go_fuzz_dep_.CoverTab[88561]++
														rb = r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:48
			// _ = "end of CoverTab[88561]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:49
			_go_fuzz_dep_.CoverTab[88570]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:49
			// _ = "end of CoverTab[88570]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:49
		// _ = "end of CoverTab[88557]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:49
		_go_fuzz_dep_.CoverTab[88558]++
													return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:50
		// _ = "end of CoverTab[88558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:51
		_go_fuzz_dep_.CoverTab[88571]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:51
		// _ = "end of CoverTab[88571]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:51
	// _ = "end of CoverTab[88546]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:51
	_go_fuzz_dep_.CoverTab[88547]++

												rb, errtcp := cl.sendKDCTCP(realm, b)
												if errtcp != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:54
		_go_fuzz_dep_.CoverTab[88572]++
													if e, ok := errtcp.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:55
			_go_fuzz_dep_.CoverTab[88574]++

														return rb, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:57
			// _ = "end of CoverTab[88574]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:58
			_go_fuzz_dep_.CoverTab[88575]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:58
			// _ = "end of CoverTab[88575]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:58
		// _ = "end of CoverTab[88572]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:58
		_go_fuzz_dep_.CoverTab[88573]++
													rb, errudp := cl.sendKDCUDP(realm, b)
													if errudp != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:60
			_go_fuzz_dep_.CoverTab[88576]++
														if e, ok := errudp.(messages.KRBError); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:61
				_go_fuzz_dep_.CoverTab[88578]++

															return rb, e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:63
				// _ = "end of CoverTab[88578]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:64
				_go_fuzz_dep_.CoverTab[88579]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:64
				// _ = "end of CoverTab[88579]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:64
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:64
			// _ = "end of CoverTab[88576]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:64
			_go_fuzz_dep_.CoverTab[88577]++
														return rb, fmt.Errorf("failed to communicate with KDC. Attempts made with TCP (%v) and then UDP (%v)", errtcp, errudp)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:65
			// _ = "end of CoverTab[88577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:66
			_go_fuzz_dep_.CoverTab[88580]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:66
			// _ = "end of CoverTab[88580]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:66
		// _ = "end of CoverTab[88573]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:67
		_go_fuzz_dep_.CoverTab[88581]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:67
		// _ = "end of CoverTab[88581]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:67
	// _ = "end of CoverTab[88547]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:67
	_go_fuzz_dep_.CoverTab[88548]++
												return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:68
	// _ = "end of CoverTab[88548]"
}

// sendKDCUDP sends bytes to the KDC via UDP.
func (cl *Client) sendKDCUDP(realm string, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:72
	_go_fuzz_dep_.CoverTab[88582]++
												var r []byte
												_, kdcs, err := cl.Config.GetKDCs(realm, false)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:75
		_go_fuzz_dep_.CoverTab[88585]++
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:76
		// _ = "end of CoverTab[88585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:77
		_go_fuzz_dep_.CoverTab[88586]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:77
		// _ = "end of CoverTab[88586]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:77
	// _ = "end of CoverTab[88582]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:77
	_go_fuzz_dep_.CoverTab[88583]++
												r, err = dialSendUDP(kdcs, b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:79
		_go_fuzz_dep_.CoverTab[88587]++
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:80
		// _ = "end of CoverTab[88587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:81
		_go_fuzz_dep_.CoverTab[88588]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:81
		// _ = "end of CoverTab[88588]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:81
	// _ = "end of CoverTab[88583]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:81
	_go_fuzz_dep_.CoverTab[88584]++
												return checkForKRBError(r)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:82
	// _ = "end of CoverTab[88584]"
}

// dialSendUDP establishes a UDP connection to a KDC.
func dialSendUDP(kdcs map[int]string, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:86
	_go_fuzz_dep_.CoverTab[88589]++
												var errs []string
												for i := 1; i <= len(kdcs); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:88
		_go_fuzz_dep_.CoverTab[88591]++
													udpAddr, err := net.ResolveUDPAddr("udp", kdcs[i])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:90
			_go_fuzz_dep_.CoverTab[88596]++
														errs = append(errs, fmt.Sprintf("error resolving KDC address: %v", err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:92
			// _ = "end of CoverTab[88596]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:93
			_go_fuzz_dep_.CoverTab[88597]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:93
			// _ = "end of CoverTab[88597]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:93
		// _ = "end of CoverTab[88591]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:93
		_go_fuzz_dep_.CoverTab[88592]++

													conn, err := net.DialTimeout("udp", udpAddr.String(), 5*time.Second)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:96
			_go_fuzz_dep_.CoverTab[88598]++
														errs = append(errs, fmt.Sprintf("error setting dial timeout on connection to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:98
			// _ = "end of CoverTab[88598]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:99
			_go_fuzz_dep_.CoverTab[88599]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:99
			// _ = "end of CoverTab[88599]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:99
		// _ = "end of CoverTab[88592]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:99
		_go_fuzz_dep_.CoverTab[88593]++
													if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:100
			_go_fuzz_dep_.CoverTab[88600]++
														errs = append(errs, fmt.Sprintf("error setting deadline on connection to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:102
			// _ = "end of CoverTab[88600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:103
			_go_fuzz_dep_.CoverTab[88601]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:103
			// _ = "end of CoverTab[88601]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:103
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:103
		// _ = "end of CoverTab[88593]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:103
		_go_fuzz_dep_.CoverTab[88594]++

													rb, err := sendUDP(conn.(*net.UDPConn), b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:106
			_go_fuzz_dep_.CoverTab[88602]++
														errs = append(errs, fmt.Sprintf("error sneding to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:108
			// _ = "end of CoverTab[88602]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:109
			_go_fuzz_dep_.CoverTab[88603]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:109
			// _ = "end of CoverTab[88603]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:109
		// _ = "end of CoverTab[88594]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:109
		_go_fuzz_dep_.CoverTab[88595]++
													return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:110
		// _ = "end of CoverTab[88595]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:111
	// _ = "end of CoverTab[88589]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:111
	_go_fuzz_dep_.CoverTab[88590]++
												return nil, fmt.Errorf("error sending to a KDC: %s", strings.Join(errs, "; "))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:112
	// _ = "end of CoverTab[88590]"
}

// sendUDP sends bytes to connection over UDP.
func sendUDP(conn *net.UDPConn, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:116
	_go_fuzz_dep_.CoverTab[88604]++
												var r []byte
												defer conn.Close()
												_, err := conn.Write(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:120
		_go_fuzz_dep_.CoverTab[88608]++
													return r, fmt.Errorf("error sending to (%s): %v", conn.RemoteAddr().String(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:121
		// _ = "end of CoverTab[88608]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:122
		_go_fuzz_dep_.CoverTab[88609]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:122
		// _ = "end of CoverTab[88609]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:122
	// _ = "end of CoverTab[88604]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:122
	_go_fuzz_dep_.CoverTab[88605]++
												udpbuf := make([]byte, 4096)
												n, _, err := conn.ReadFrom(udpbuf)
												r = udpbuf[:n]
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:126
		_go_fuzz_dep_.CoverTab[88610]++
													return r, fmt.Errorf("sending over UDP failed to %s: %v", conn.RemoteAddr().String(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:127
		// _ = "end of CoverTab[88610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:128
		_go_fuzz_dep_.CoverTab[88611]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:128
		// _ = "end of CoverTab[88611]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:128
	// _ = "end of CoverTab[88605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:128
	_go_fuzz_dep_.CoverTab[88606]++
												if len(r) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:129
		_go_fuzz_dep_.CoverTab[88612]++
													return r, fmt.Errorf("no response data from %s", conn.RemoteAddr().String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:130
		// _ = "end of CoverTab[88612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:131
		_go_fuzz_dep_.CoverTab[88613]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:131
		// _ = "end of CoverTab[88613]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:131
	// _ = "end of CoverTab[88606]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:131
	_go_fuzz_dep_.CoverTab[88607]++
												return r, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:132
	// _ = "end of CoverTab[88607]"
}

// sendKDCTCP sends bytes to the KDC via TCP.
func (cl *Client) sendKDCTCP(realm string, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:136
	_go_fuzz_dep_.CoverTab[88614]++
												var r []byte
												_, kdcs, err := cl.Config.GetKDCs(realm, true)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:139
		_go_fuzz_dep_.CoverTab[88617]++
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:140
		// _ = "end of CoverTab[88617]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:141
		_go_fuzz_dep_.CoverTab[88618]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:141
		// _ = "end of CoverTab[88618]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:141
	// _ = "end of CoverTab[88614]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:141
	_go_fuzz_dep_.CoverTab[88615]++
												r, err = dialSendTCP(kdcs, b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:143
		_go_fuzz_dep_.CoverTab[88619]++
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:144
		// _ = "end of CoverTab[88619]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:145
		_go_fuzz_dep_.CoverTab[88620]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:145
		// _ = "end of CoverTab[88620]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:145
	// _ = "end of CoverTab[88615]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:145
	_go_fuzz_dep_.CoverTab[88616]++
												return checkForKRBError(r)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:146
	// _ = "end of CoverTab[88616]"
}

// dialKDCTCP establishes a TCP connection to a KDC.
func dialSendTCP(kdcs map[int]string, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:150
	_go_fuzz_dep_.CoverTab[88621]++
												var errs []string
												for i := 1; i <= len(kdcs); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:152
		_go_fuzz_dep_.CoverTab[88623]++
													tcpAddr, err := net.ResolveTCPAddr("tcp", kdcs[i])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:154
			_go_fuzz_dep_.CoverTab[88628]++
														errs = append(errs, fmt.Sprintf("error resolving KDC address: %v", err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:156
			// _ = "end of CoverTab[88628]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:157
			_go_fuzz_dep_.CoverTab[88629]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:157
			// _ = "end of CoverTab[88629]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:157
		// _ = "end of CoverTab[88623]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:157
		_go_fuzz_dep_.CoverTab[88624]++

													conn, err := net.DialTimeout("tcp", tcpAddr.String(), 5*time.Second)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:160
			_go_fuzz_dep_.CoverTab[88630]++
														errs = append(errs, fmt.Sprintf("error setting dial timeout on connection to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:162
			// _ = "end of CoverTab[88630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:163
			_go_fuzz_dep_.CoverTab[88631]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:163
			// _ = "end of CoverTab[88631]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:163
		// _ = "end of CoverTab[88624]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:163
		_go_fuzz_dep_.CoverTab[88625]++
													if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:164
			_go_fuzz_dep_.CoverTab[88632]++
														errs = append(errs, fmt.Sprintf("error setting deadline on connection to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:166
			// _ = "end of CoverTab[88632]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:167
			_go_fuzz_dep_.CoverTab[88633]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:167
			// _ = "end of CoverTab[88633]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:167
		// _ = "end of CoverTab[88625]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:167
		_go_fuzz_dep_.CoverTab[88626]++

													rb, err := sendTCP(conn.(*net.TCPConn), b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:170
			_go_fuzz_dep_.CoverTab[88634]++
														errs = append(errs, fmt.Sprintf("error sneding to %s: %v", kdcs[i], err))
														continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:172
			// _ = "end of CoverTab[88634]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:173
			_go_fuzz_dep_.CoverTab[88635]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:173
			// _ = "end of CoverTab[88635]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:173
		// _ = "end of CoverTab[88626]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:173
		_go_fuzz_dep_.CoverTab[88627]++
													return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:174
		// _ = "end of CoverTab[88627]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:175
	// _ = "end of CoverTab[88621]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:175
	_go_fuzz_dep_.CoverTab[88622]++
												return nil, errors.New("error in getting a TCP connection to any of the KDCs")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:176
	// _ = "end of CoverTab[88622]"
}

// sendTCP sends bytes to connection over TCP.
func sendTCP(conn *net.TCPConn, b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:180
	_go_fuzz_dep_.CoverTab[88636]++
												defer conn.Close()
												var r []byte

												hb := make([]byte, 4, 4)
												binary.BigEndian.PutUint32(hb, uint32(len(b)))
												b = append(hb, b...)

												_, err := conn.Write(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:189
		_go_fuzz_dep_.CoverTab[88641]++
													return r, fmt.Errorf("error sending to KDC (%s): %v", conn.RemoteAddr().String(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:190
		// _ = "end of CoverTab[88641]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:191
		_go_fuzz_dep_.CoverTab[88642]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:191
		// _ = "end of CoverTab[88642]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:191
	// _ = "end of CoverTab[88636]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:191
	_go_fuzz_dep_.CoverTab[88637]++

												sh := make([]byte, 4, 4)
												_, err = conn.Read(sh)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:195
		_go_fuzz_dep_.CoverTab[88643]++
													return r, fmt.Errorf("error reading response size header: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:196
		// _ = "end of CoverTab[88643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:197
		_go_fuzz_dep_.CoverTab[88644]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:197
		// _ = "end of CoverTab[88644]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:197
	// _ = "end of CoverTab[88637]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:197
	_go_fuzz_dep_.CoverTab[88638]++
												s := binary.BigEndian.Uint32(sh)

												rb := make([]byte, s, s)
												_, err = io.ReadFull(conn, rb)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:202
		_go_fuzz_dep_.CoverTab[88645]++
													return r, fmt.Errorf("error reading response: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:203
		// _ = "end of CoverTab[88645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:204
		_go_fuzz_dep_.CoverTab[88646]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:204
		// _ = "end of CoverTab[88646]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:204
	// _ = "end of CoverTab[88638]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:204
	_go_fuzz_dep_.CoverTab[88639]++
												if len(rb) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:205
		_go_fuzz_dep_.CoverTab[88647]++
													return r, fmt.Errorf("no response data from KDC %s", conn.RemoteAddr().String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:206
		// _ = "end of CoverTab[88647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:207
		_go_fuzz_dep_.CoverTab[88648]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:207
		// _ = "end of CoverTab[88648]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:207
	// _ = "end of CoverTab[88639]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:207
	_go_fuzz_dep_.CoverTab[88640]++
												return rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:208
	// _ = "end of CoverTab[88640]"
}

// checkForKRBError checks if the response bytes from the KDC are a KRBError.
func checkForKRBError(b []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:212
	_go_fuzz_dep_.CoverTab[88649]++
												var KRBErr messages.KRBError
												if err := KRBErr.Unmarshal(b); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:214
		_go_fuzz_dep_.CoverTab[88651]++
													return b, KRBErr
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:215
		// _ = "end of CoverTab[88651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:216
		_go_fuzz_dep_.CoverTab[88652]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:216
		// _ = "end of CoverTab[88652]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:216
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:216
	// _ = "end of CoverTab[88649]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:216
	_go_fuzz_dep_.CoverTab[88650]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:217
	// _ = "end of CoverTab[88650]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/network.go:218
var _ = _go_fuzz_dep_.CoverTab
