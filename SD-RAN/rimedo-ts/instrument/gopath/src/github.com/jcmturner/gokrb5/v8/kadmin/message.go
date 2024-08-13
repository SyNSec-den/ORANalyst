//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
package kadmin

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:1
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

const (
	verisonHex = "ff80"
)

// Request message for changing password.
type Request struct {
	APREQ	messages.APReq
	KRBPriv	messages.KRBPriv
}

// Reply message for a password change.
type Reply struct {
	MessageLength	int
	Version		int
	APREPLength	int
	APREP		messages.APRep
	KRBPriv		messages.KRBPriv
	KRBError	messages.KRBError
	IsKRBError	bool
	ResultCode	uint16
	Result		string
}

// Marshal a Request into a byte slice.
func (m *Request) Marshal() (b []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:38
	_go_fuzz_dep_.CoverTab[88184]++
												b = []byte{255, 128}
												ab, e := m.APREQ.Marshal()
												if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:41
		_go_fuzz_dep_.CoverTab[88189]++
													err = fmt.Errorf("error marshaling AP_REQ: %v", e)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:43
		// _ = "end of CoverTab[88189]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:44
		_go_fuzz_dep_.CoverTab[88190]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:44
		// _ = "end of CoverTab[88190]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:44
	// _ = "end of CoverTab[88184]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:44
	_go_fuzz_dep_.CoverTab[88185]++
												if len(ab) > math.MaxUint16 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:45
		_go_fuzz_dep_.CoverTab[88191]++
													err = errors.New("length of AP_REQ greater then max Uint16 size")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:47
		// _ = "end of CoverTab[88191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:48
		_go_fuzz_dep_.CoverTab[88192]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:48
		// _ = "end of CoverTab[88192]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:48
	// _ = "end of CoverTab[88185]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:48
	_go_fuzz_dep_.CoverTab[88186]++
												al := make([]byte, 2)
												binary.BigEndian.PutUint16(al, uint16(len(ab)))
												b = append(b, al...)
												b = append(b, ab...)
												pb, e := m.KRBPriv.Marshal()
												if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:54
		_go_fuzz_dep_.CoverTab[88193]++
													err = fmt.Errorf("error marshaling KRB_Priv: %v", e)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:56
		// _ = "end of CoverTab[88193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:57
		_go_fuzz_dep_.CoverTab[88194]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:57
		// _ = "end of CoverTab[88194]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:57
	// _ = "end of CoverTab[88186]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:57
	_go_fuzz_dep_.CoverTab[88187]++
												b = append(b, pb...)
												if len(b)+2 > math.MaxUint16 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:59
		_go_fuzz_dep_.CoverTab[88195]++
													err = errors.New("length of message greater then max Uint16 size")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:61
		// _ = "end of CoverTab[88195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:62
		_go_fuzz_dep_.CoverTab[88196]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:62
		// _ = "end of CoverTab[88196]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:62
	// _ = "end of CoverTab[88187]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:62
	_go_fuzz_dep_.CoverTab[88188]++
												ml := make([]byte, 2)
												binary.BigEndian.PutUint16(ml, uint16(len(b)+2))
												b = append(ml, b...)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:66
	// _ = "end of CoverTab[88188]"
}

// Unmarshal a byte slice into a Reply.
func (m *Reply) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:70
	_go_fuzz_dep_.CoverTab[88197]++
												m.MessageLength = int(binary.BigEndian.Uint16(b[0:2]))
												m.Version = int(binary.BigEndian.Uint16(b[2:4]))
												if m.Version != 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:73
		_go_fuzz_dep_.CoverTab[88200]++
													return fmt.Errorf("kadmin reply has incorrect protocol version number: %d", m.Version)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:74
		// _ = "end of CoverTab[88200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:75
		_go_fuzz_dep_.CoverTab[88201]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:75
		// _ = "end of CoverTab[88201]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:75
	// _ = "end of CoverTab[88197]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:75
	_go_fuzz_dep_.CoverTab[88198]++
												m.APREPLength = int(binary.BigEndian.Uint16(b[4:6]))
												if m.APREPLength != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:77
		_go_fuzz_dep_.CoverTab[88202]++
													err := m.APREP.Unmarshal(b[6 : 6+m.APREPLength])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:79
			_go_fuzz_dep_.CoverTab[88204]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:80
			// _ = "end of CoverTab[88204]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:81
			_go_fuzz_dep_.CoverTab[88205]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:81
			// _ = "end of CoverTab[88205]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:81
		// _ = "end of CoverTab[88202]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:81
		_go_fuzz_dep_.CoverTab[88203]++
													err = m.KRBPriv.Unmarshal(b[6+m.APREPLength : m.MessageLength])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:83
			_go_fuzz_dep_.CoverTab[88206]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:84
			// _ = "end of CoverTab[88206]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:85
			_go_fuzz_dep_.CoverTab[88207]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:85
			// _ = "end of CoverTab[88207]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:85
		// _ = "end of CoverTab[88203]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:86
		_go_fuzz_dep_.CoverTab[88208]++
													m.IsKRBError = true
													m.KRBError.Unmarshal(b[6:m.MessageLength])
													m.ResultCode, m.Result = parseResponse(m.KRBError.EData)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:89
		// _ = "end of CoverTab[88208]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:90
	// _ = "end of CoverTab[88198]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:90
	_go_fuzz_dep_.CoverTab[88199]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:91
	// _ = "end of CoverTab[88199]"
}

func parseResponse(b []byte) (c uint16, s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:94
	_go_fuzz_dep_.CoverTab[88209]++
												c = binary.BigEndian.Uint16(b[0:2])
												buf := bytes.NewBuffer(b[2:])
												m := make([]byte, len(b)-2)
												binary.Read(buf, binary.BigEndian, &m)
												s = string(m)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:100
	// _ = "end of CoverTab[88209]"
}

// Decrypt the encrypted part of the KRBError within the change password Reply.
func (m *Reply) Decrypt(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:104
	_go_fuzz_dep_.CoverTab[88210]++
												if m.IsKRBError {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:105
		_go_fuzz_dep_.CoverTab[88213]++
													return m.KRBError
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:106
		// _ = "end of CoverTab[88213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:107
		_go_fuzz_dep_.CoverTab[88214]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:107
		// _ = "end of CoverTab[88214]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:107
	// _ = "end of CoverTab[88210]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:107
	_go_fuzz_dep_.CoverTab[88211]++
												err := m.KRBPriv.DecryptEncPart(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:109
		_go_fuzz_dep_.CoverTab[88215]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:110
		// _ = "end of CoverTab[88215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:111
		_go_fuzz_dep_.CoverTab[88216]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:111
		// _ = "end of CoverTab[88216]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:111
	// _ = "end of CoverTab[88211]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:111
	_go_fuzz_dep_.CoverTab[88212]++
												m.ResultCode, m.Result = parseResponse(m.KRBPriv.DecryptedEncPart.UserData)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:113
	// _ = "end of CoverTab[88212]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/message.go:114
var _ = _go_fuzz_dep_.CoverTab
