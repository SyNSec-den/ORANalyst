//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:1
)

import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

// KRBSafe implements RFC 4120 KRB_SAFE: https://tools.ietf.org/html/rfc4120#section-5.6.1.
type KRBSafe struct {
	PVNO		int		`asn1:"explicit,tag:0"`
	MsgType		int		`asn1:"explicit,tag:1"`
	SafeBody	KRBSafeBody	`asn1:"explicit,tag:2"`
	Cksum		types.Checksum	`asn1:"explicit,tag:3"`
}

// KRBSafeBody implements the KRB_SAFE_BODY of KRB_SAFE.
type KRBSafeBody struct {
	UserData	[]byte			`asn1:"explicit,tag:0"`
	Timestamp	time.Time		`asn1:"generalized,optional,explicit,tag:1"`
	Usec		int			`asn1:"optional,explicit,tag:2"`
	SequenceNumber	int64			`asn1:"optional,explicit,tag:3"`
	SAddress	types.HostAddress	`asn1:"explicit,tag:4"`
	RAddress	types.HostAddress	`asn1:"optional,explicit,tag:5"`
}

// Unmarshal bytes b into the KRBSafe struct.
func (s *KRBSafe) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:33
	_go_fuzz_dep_.CoverTab[88091]++
												_, err := asn1.UnmarshalWithParams(b, s, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.KRBSafe))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:35
		_go_fuzz_dep_.CoverTab[88094]++
													return processUnmarshalReplyError(b, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:36
		// _ = "end of CoverTab[88094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:37
		_go_fuzz_dep_.CoverTab[88095]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:37
		// _ = "end of CoverTab[88095]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:37
	// _ = "end of CoverTab[88091]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:37
	_go_fuzz_dep_.CoverTab[88092]++
												expectedMsgType := msgtype.KRB_SAFE
												if s.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:39
		_go_fuzz_dep_.CoverTab[88096]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a KRB_SAFE. Expected: %v; Actual: %v", expectedMsgType, s.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:40
		// _ = "end of CoverTab[88096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:41
		_go_fuzz_dep_.CoverTab[88097]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:41
		// _ = "end of CoverTab[88097]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:41
	// _ = "end of CoverTab[88092]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:41
	_go_fuzz_dep_.CoverTab[88093]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:42
	// _ = "end of CoverTab[88093]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBSafe.go:43
var _ = _go_fuzz_dep_.CoverTab
