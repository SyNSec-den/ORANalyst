// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/alert.go:5
package tls

//line /usr/local/go/src/crypto/tls/alert.go:5
import (
//line /usr/local/go/src/crypto/tls/alert.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/alert.go:5
)
//line /usr/local/go/src/crypto/tls/alert.go:5
import (
//line /usr/local/go/src/crypto/tls/alert.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/alert.go:5
)

import "strconv"

type alert uint8

const (
	// alert level
	alertLevelWarning	= 1
	alertLevelError		= 2
)

const (
	alertCloseNotify			alert	= 0
	alertUnexpectedMessage			alert	= 10
	alertBadRecordMAC			alert	= 20
	alertDecryptionFailed			alert	= 21
	alertRecordOverflow			alert	= 22
	alertDecompressionFailure		alert	= 30
	alertHandshakeFailure			alert	= 40
	alertBadCertificate			alert	= 42
	alertUnsupportedCertificate		alert	= 43
	alertCertificateRevoked			alert	= 44
	alertCertificateExpired			alert	= 45
	alertCertificateUnknown			alert	= 46
	alertIllegalParameter			alert	= 47
	alertUnknownCA				alert	= 48
	alertAccessDenied			alert	= 49
	alertDecodeError			alert	= 50
	alertDecryptError			alert	= 51
	alertExportRestriction			alert	= 60
	alertProtocolVersion			alert	= 70
	alertInsufficientSecurity		alert	= 71
	alertInternalError			alert	= 80
	alertInappropriateFallback		alert	= 86
	alertUserCanceled			alert	= 90
	alertNoRenegotiation			alert	= 100
	alertMissingExtension			alert	= 109
	alertUnsupportedExtension		alert	= 110
	alertCertificateUnobtainable		alert	= 111
	alertUnrecognizedName			alert	= 112
	alertBadCertificateStatusResponse	alert	= 113
	alertBadCertificateHashValue		alert	= 114
	alertUnknownPSKIdentity			alert	= 115
	alertCertificateRequired		alert	= 116
	alertNoApplicationProtocol		alert	= 120
)

var alertText = map[alert]string{
	alertCloseNotify:			"close notify",
	alertUnexpectedMessage:			"unexpected message",
	alertBadRecordMAC:			"bad record MAC",
	alertDecryptionFailed:			"decryption failed",
	alertRecordOverflow:			"record overflow",
	alertDecompressionFailure:		"decompression failure",
	alertHandshakeFailure:			"handshake failure",
	alertBadCertificate:			"bad certificate",
	alertUnsupportedCertificate:		"unsupported certificate",
	alertCertificateRevoked:		"revoked certificate",
	alertCertificateExpired:		"expired certificate",
	alertCertificateUnknown:		"unknown certificate",
	alertIllegalParameter:			"illegal parameter",
	alertUnknownCA:				"unknown certificate authority",
	alertAccessDenied:			"access denied",
	alertDecodeError:			"error decoding message",
	alertDecryptError:			"error decrypting message",
	alertExportRestriction:			"export restriction",
	alertProtocolVersion:			"protocol version not supported",
	alertInsufficientSecurity:		"insufficient security level",
	alertInternalError:			"internal error",
	alertInappropriateFallback:		"inappropriate fallback",
	alertUserCanceled:			"user canceled",
	alertNoRenegotiation:			"no renegotiation",
	alertMissingExtension:			"missing extension",
	alertUnsupportedExtension:		"unsupported extension",
	alertCertificateUnobtainable:		"certificate unobtainable",
	alertUnrecognizedName:			"unrecognized name",
	alertBadCertificateStatusResponse:	"bad certificate status response",
	alertBadCertificateHashValue:		"bad certificate hash value",
	alertUnknownPSKIdentity:		"unknown PSK identity",
	alertCertificateRequired:		"certificate required",
	alertNoApplicationProtocol:		"no application protocol",
}

func (e alert) String() string {
//line /usr/local/go/src/crypto/tls/alert.go:89
	_go_fuzz_dep_.CoverTab[21006]++
							s, ok := alertText[e]
							if ok {
//line /usr/local/go/src/crypto/tls/alert.go:91
		_go_fuzz_dep_.CoverTab[21008]++
								return "tls: " + s
//line /usr/local/go/src/crypto/tls/alert.go:92
		// _ = "end of CoverTab[21008]"
	} else {
//line /usr/local/go/src/crypto/tls/alert.go:93
		_go_fuzz_dep_.CoverTab[21009]++
//line /usr/local/go/src/crypto/tls/alert.go:93
		// _ = "end of CoverTab[21009]"
//line /usr/local/go/src/crypto/tls/alert.go:93
	}
//line /usr/local/go/src/crypto/tls/alert.go:93
	// _ = "end of CoverTab[21006]"
//line /usr/local/go/src/crypto/tls/alert.go:93
	_go_fuzz_dep_.CoverTab[21007]++
							return "tls: alert(" + strconv.Itoa(int(e)) + ")"
//line /usr/local/go/src/crypto/tls/alert.go:94
	// _ = "end of CoverTab[21007]"
}

func (e alert) Error() string {
//line /usr/local/go/src/crypto/tls/alert.go:97
	_go_fuzz_dep_.CoverTab[21010]++
							return e.String()
//line /usr/local/go/src/crypto/tls/alert.go:98
	// _ = "end of CoverTab[21010]"
}

//line /usr/local/go/src/crypto/tls/alert.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/alert.go:99
var _ = _go_fuzz_dep_.CoverTab
