//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:1
// Package asnAppTag provides ASN1 application tag numbers.
package asnAppTag

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:2
)

// ASN1 application tag numbers.
const (
	Ticket		= 1
	Authenticator	= 2
	EncTicketPart	= 3
	ASREQ		= 10
	TGSREQ		= 12
	ASREP		= 11
	TGSREP		= 13
	APREQ		= 14
	APREP		= 15
	KRBSafe		= 20
	KRBPriv		= 21
	KRBCred		= 22
	EncASRepPart	= 25
	EncTGSRepPart	= 26
	EncAPRepPart	= 27
	EncKrbPrivPart	= 28
	EncKrbCredPart	= 29
	KRBError	= 30
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/asnAppTag/constants.go:24
var _ = _go_fuzz_dep_.CoverTab
