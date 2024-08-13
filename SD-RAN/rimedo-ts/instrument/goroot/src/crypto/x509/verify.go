// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/verify.go:5
package x509

//line /usr/local/go/src/crypto/x509/verify.go:5
import (
//line /usr/local/go/src/crypto/x509/verify.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/verify.go:5
)
//line /usr/local/go/src/crypto/x509/verify.go:5
import (
//line /usr/local/go/src/crypto/x509/verify.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/verify.go:5
)

import (
	"bytes"
	"crypto"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

type InvalidReason int

const (
	// NotAuthorizedToSign results when a certificate is signed by another
	// which isn't marked as a CA certificate.
	NotAuthorizedToSign	InvalidReason	= iota
	// Expired results when a certificate has expired, based on the time
	// given in the VerifyOptions.
	Expired
	// CANotAuthorizedForThisName results when an intermediate or root
	// certificate has a name constraint which doesn't permit a DNS or
	// other name (including IP address) in the leaf certificate.
	CANotAuthorizedForThisName
	// TooManyIntermediates results when a path length constraint is
	// violated.
	TooManyIntermediates
	// IncompatibleUsage results when the certificate's key usage indicates
	// that it may only be used for a different purpose.
	IncompatibleUsage
	// NameMismatch results when the subject name of a parent certificate
	// does not match the issuer name in the child.
	NameMismatch
	// NameConstraintsWithoutSANs is a legacy error and is no longer returned.
	NameConstraintsWithoutSANs
	// UnconstrainedName results when a CA certificate contains permitted
	// name constraints, but leaf certificate contains a name of an
	// unsupported or unconstrained type.
	UnconstrainedName
	// TooManyConstraints results when the number of comparison operations
	// needed to check a certificate exceeds the limit set by
	// VerifyOptions.MaxConstraintComparisions. This limit exists to
	// prevent pathological certificates can consuming excessive amounts of
	// CPU time to verify.
	TooManyConstraints
	// CANotAuthorizedForExtKeyUsage results when an intermediate or root
	// certificate does not permit a requested extended key usage.
	CANotAuthorizedForExtKeyUsage
)

// CertificateInvalidError results when an odd error occurs. Users of this
//line /usr/local/go/src/crypto/x509/verify.go:61
// library probably want to handle all these errors uniformly.
//line /usr/local/go/src/crypto/x509/verify.go:63
type CertificateInvalidError struct {
	Cert	*Certificate
	Reason	InvalidReason
	Detail	string
}

func (e CertificateInvalidError) Error() string {
//line /usr/local/go/src/crypto/x509/verify.go:69
	_go_fuzz_dep_.CoverTab[19349]++
							switch e.Reason {
	case NotAuthorizedToSign:
//line /usr/local/go/src/crypto/x509/verify.go:71
		_go_fuzz_dep_.CoverTab[19351]++
								return "x509: certificate is not authorized to sign other certificates"
//line /usr/local/go/src/crypto/x509/verify.go:72
		// _ = "end of CoverTab[19351]"
	case Expired:
//line /usr/local/go/src/crypto/x509/verify.go:73
		_go_fuzz_dep_.CoverTab[19352]++
								return "x509: certificate has expired or is not yet valid: " + e.Detail
//line /usr/local/go/src/crypto/x509/verify.go:74
		// _ = "end of CoverTab[19352]"
	case CANotAuthorizedForThisName:
//line /usr/local/go/src/crypto/x509/verify.go:75
		_go_fuzz_dep_.CoverTab[19353]++
								return "x509: a root or intermediate certificate is not authorized to sign for this name: " + e.Detail
//line /usr/local/go/src/crypto/x509/verify.go:76
		// _ = "end of CoverTab[19353]"
	case CANotAuthorizedForExtKeyUsage:
//line /usr/local/go/src/crypto/x509/verify.go:77
		_go_fuzz_dep_.CoverTab[19354]++
								return "x509: a root or intermediate certificate is not authorized for an extended key usage: " + e.Detail
//line /usr/local/go/src/crypto/x509/verify.go:78
		// _ = "end of CoverTab[19354]"
	case TooManyIntermediates:
//line /usr/local/go/src/crypto/x509/verify.go:79
		_go_fuzz_dep_.CoverTab[19355]++
								return "x509: too many intermediates for path length constraint"
//line /usr/local/go/src/crypto/x509/verify.go:80
		// _ = "end of CoverTab[19355]"
	case IncompatibleUsage:
//line /usr/local/go/src/crypto/x509/verify.go:81
		_go_fuzz_dep_.CoverTab[19356]++
								return "x509: certificate specifies an incompatible key usage"
//line /usr/local/go/src/crypto/x509/verify.go:82
		// _ = "end of CoverTab[19356]"
	case NameMismatch:
//line /usr/local/go/src/crypto/x509/verify.go:83
		_go_fuzz_dep_.CoverTab[19357]++
								return "x509: issuer name does not match subject from issuing certificate"
//line /usr/local/go/src/crypto/x509/verify.go:84
		// _ = "end of CoverTab[19357]"
	case NameConstraintsWithoutSANs:
//line /usr/local/go/src/crypto/x509/verify.go:85
		_go_fuzz_dep_.CoverTab[19358]++
								return "x509: issuer has name constraints but leaf doesn't have a SAN extension"
//line /usr/local/go/src/crypto/x509/verify.go:86
		// _ = "end of CoverTab[19358]"
	case UnconstrainedName:
//line /usr/local/go/src/crypto/x509/verify.go:87
		_go_fuzz_dep_.CoverTab[19359]++
								return "x509: issuer has name constraints but leaf contains unknown or unconstrained name: " + e.Detail
//line /usr/local/go/src/crypto/x509/verify.go:88
		// _ = "end of CoverTab[19359]"
//line /usr/local/go/src/crypto/x509/verify.go:88
	default:
//line /usr/local/go/src/crypto/x509/verify.go:88
		_go_fuzz_dep_.CoverTab[19360]++
//line /usr/local/go/src/crypto/x509/verify.go:88
		// _ = "end of CoverTab[19360]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:89
	// _ = "end of CoverTab[19349]"
//line /usr/local/go/src/crypto/x509/verify.go:89
	_go_fuzz_dep_.CoverTab[19350]++
							return "x509: unknown error"
//line /usr/local/go/src/crypto/x509/verify.go:90
	// _ = "end of CoverTab[19350]"
}

// HostnameError results when the set of authorized names doesn't match the
//line /usr/local/go/src/crypto/x509/verify.go:93
// requested name.
//line /usr/local/go/src/crypto/x509/verify.go:95
type HostnameError struct {
	Certificate	*Certificate
	Host		string
}

func (h HostnameError) Error() string {
//line /usr/local/go/src/crypto/x509/verify.go:100
	_go_fuzz_dep_.CoverTab[19361]++
							c := h.Certificate

							if !c.hasSANExtension() && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:103
		_go_fuzz_dep_.CoverTab[19365]++
//line /usr/local/go/src/crypto/x509/verify.go:103
		return matchHostnames(c.Subject.CommonName, h.Host)
//line /usr/local/go/src/crypto/x509/verify.go:103
		// _ = "end of CoverTab[19365]"
//line /usr/local/go/src/crypto/x509/verify.go:103
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:103
		_go_fuzz_dep_.CoverTab[19366]++
								return "x509: certificate relies on legacy Common Name field, use SANs instead"
//line /usr/local/go/src/crypto/x509/verify.go:104
		// _ = "end of CoverTab[19366]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:105
		_go_fuzz_dep_.CoverTab[19367]++
//line /usr/local/go/src/crypto/x509/verify.go:105
		// _ = "end of CoverTab[19367]"
//line /usr/local/go/src/crypto/x509/verify.go:105
	}
//line /usr/local/go/src/crypto/x509/verify.go:105
	// _ = "end of CoverTab[19361]"
//line /usr/local/go/src/crypto/x509/verify.go:105
	_go_fuzz_dep_.CoverTab[19362]++

							var valid string
							if ip := net.ParseIP(h.Host); ip != nil {
//line /usr/local/go/src/crypto/x509/verify.go:108
		_go_fuzz_dep_.CoverTab[19368]++

								if len(c.IPAddresses) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:110
			_go_fuzz_dep_.CoverTab[19370]++
									return "x509: cannot validate certificate for " + h.Host + " because it doesn't contain any IP SANs"
//line /usr/local/go/src/crypto/x509/verify.go:111
			// _ = "end of CoverTab[19370]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:112
			_go_fuzz_dep_.CoverTab[19371]++
//line /usr/local/go/src/crypto/x509/verify.go:112
			// _ = "end of CoverTab[19371]"
//line /usr/local/go/src/crypto/x509/verify.go:112
		}
//line /usr/local/go/src/crypto/x509/verify.go:112
		// _ = "end of CoverTab[19368]"
//line /usr/local/go/src/crypto/x509/verify.go:112
		_go_fuzz_dep_.CoverTab[19369]++
								for _, san := range c.IPAddresses {
//line /usr/local/go/src/crypto/x509/verify.go:113
			_go_fuzz_dep_.CoverTab[19372]++
									if len(valid) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:114
				_go_fuzz_dep_.CoverTab[19374]++
										valid += ", "
//line /usr/local/go/src/crypto/x509/verify.go:115
				// _ = "end of CoverTab[19374]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:116
				_go_fuzz_dep_.CoverTab[19375]++
//line /usr/local/go/src/crypto/x509/verify.go:116
				// _ = "end of CoverTab[19375]"
//line /usr/local/go/src/crypto/x509/verify.go:116
			}
//line /usr/local/go/src/crypto/x509/verify.go:116
			// _ = "end of CoverTab[19372]"
//line /usr/local/go/src/crypto/x509/verify.go:116
			_go_fuzz_dep_.CoverTab[19373]++
									valid += san.String()
//line /usr/local/go/src/crypto/x509/verify.go:117
			// _ = "end of CoverTab[19373]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:118
		// _ = "end of CoverTab[19369]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:119
		_go_fuzz_dep_.CoverTab[19376]++
								valid = strings.Join(c.DNSNames, ", ")
//line /usr/local/go/src/crypto/x509/verify.go:120
		// _ = "end of CoverTab[19376]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:121
	// _ = "end of CoverTab[19362]"
//line /usr/local/go/src/crypto/x509/verify.go:121
	_go_fuzz_dep_.CoverTab[19363]++

							if len(valid) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:123
		_go_fuzz_dep_.CoverTab[19377]++
								return "x509: certificate is not valid for any names, but wanted to match " + h.Host
//line /usr/local/go/src/crypto/x509/verify.go:124
		// _ = "end of CoverTab[19377]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:125
		_go_fuzz_dep_.CoverTab[19378]++
//line /usr/local/go/src/crypto/x509/verify.go:125
		// _ = "end of CoverTab[19378]"
//line /usr/local/go/src/crypto/x509/verify.go:125
	}
//line /usr/local/go/src/crypto/x509/verify.go:125
	// _ = "end of CoverTab[19363]"
//line /usr/local/go/src/crypto/x509/verify.go:125
	_go_fuzz_dep_.CoverTab[19364]++
							return "x509: certificate is valid for " + valid + ", not " + h.Host
//line /usr/local/go/src/crypto/x509/verify.go:126
	// _ = "end of CoverTab[19364]"
}

// UnknownAuthorityError results when the certificate issuer is unknown
type UnknownAuthorityError struct {
	Cert	*Certificate
	// hintErr contains an error that may be helpful in determining why an
	// authority wasn't found.
	hintErr	error
	// hintCert contains a possible authority certificate that was rejected
	// because of the error in hintErr.
	hintCert	*Certificate
}

func (e UnknownAuthorityError) Error() string {
//line /usr/local/go/src/crypto/x509/verify.go:140
	_go_fuzz_dep_.CoverTab[19379]++
							s := "x509: certificate signed by unknown authority"
							if e.hintErr != nil {
//line /usr/local/go/src/crypto/x509/verify.go:142
		_go_fuzz_dep_.CoverTab[19381]++
								certName := e.hintCert.Subject.CommonName
								if len(certName) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:144
			_go_fuzz_dep_.CoverTab[19383]++
									if len(e.hintCert.Subject.Organization) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:145
				_go_fuzz_dep_.CoverTab[19384]++
										certName = e.hintCert.Subject.Organization[0]
//line /usr/local/go/src/crypto/x509/verify.go:146
				// _ = "end of CoverTab[19384]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:147
				_go_fuzz_dep_.CoverTab[19385]++
										certName = "serial:" + e.hintCert.SerialNumber.String()
//line /usr/local/go/src/crypto/x509/verify.go:148
				// _ = "end of CoverTab[19385]"
			}
//line /usr/local/go/src/crypto/x509/verify.go:149
			// _ = "end of CoverTab[19383]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:150
			_go_fuzz_dep_.CoverTab[19386]++
//line /usr/local/go/src/crypto/x509/verify.go:150
			// _ = "end of CoverTab[19386]"
//line /usr/local/go/src/crypto/x509/verify.go:150
		}
//line /usr/local/go/src/crypto/x509/verify.go:150
		// _ = "end of CoverTab[19381]"
//line /usr/local/go/src/crypto/x509/verify.go:150
		_go_fuzz_dep_.CoverTab[19382]++
								s += fmt.Sprintf(" (possibly because of %q while trying to verify candidate authority certificate %q)", e.hintErr, certName)
//line /usr/local/go/src/crypto/x509/verify.go:151
		// _ = "end of CoverTab[19382]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:152
		_go_fuzz_dep_.CoverTab[19387]++
//line /usr/local/go/src/crypto/x509/verify.go:152
		// _ = "end of CoverTab[19387]"
//line /usr/local/go/src/crypto/x509/verify.go:152
	}
//line /usr/local/go/src/crypto/x509/verify.go:152
	// _ = "end of CoverTab[19379]"
//line /usr/local/go/src/crypto/x509/verify.go:152
	_go_fuzz_dep_.CoverTab[19380]++
							return s
//line /usr/local/go/src/crypto/x509/verify.go:153
	// _ = "end of CoverTab[19380]"
}

// SystemRootsError results when we fail to load the system root certificates.
type SystemRootsError struct {
	Err error
}

func (se SystemRootsError) Error() string {
//line /usr/local/go/src/crypto/x509/verify.go:161
	_go_fuzz_dep_.CoverTab[19388]++
							msg := "x509: failed to load system roots and no roots provided"
							if se.Err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:163
		_go_fuzz_dep_.CoverTab[19390]++
								return msg + "; " + se.Err.Error()
//line /usr/local/go/src/crypto/x509/verify.go:164
		// _ = "end of CoverTab[19390]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:165
		_go_fuzz_dep_.CoverTab[19391]++
//line /usr/local/go/src/crypto/x509/verify.go:165
		// _ = "end of CoverTab[19391]"
//line /usr/local/go/src/crypto/x509/verify.go:165
	}
//line /usr/local/go/src/crypto/x509/verify.go:165
	// _ = "end of CoverTab[19388]"
//line /usr/local/go/src/crypto/x509/verify.go:165
	_go_fuzz_dep_.CoverTab[19389]++
							return msg
//line /usr/local/go/src/crypto/x509/verify.go:166
	// _ = "end of CoverTab[19389]"
}

func (se SystemRootsError) Unwrap() error {
//line /usr/local/go/src/crypto/x509/verify.go:169
	_go_fuzz_dep_.CoverTab[19392]++
//line /usr/local/go/src/crypto/x509/verify.go:169
	return se.Err
//line /usr/local/go/src/crypto/x509/verify.go:169
	// _ = "end of CoverTab[19392]"
//line /usr/local/go/src/crypto/x509/verify.go:169
}

// errNotParsed is returned when a certificate without ASN.1 contents is
//line /usr/local/go/src/crypto/x509/verify.go:171
// verified. Platform-specific verification needs the ASN.1 contents.
//line /usr/local/go/src/crypto/x509/verify.go:173
var errNotParsed = errors.New("x509: missing ASN.1 contents; use ParseCertificate")

// VerifyOptions contains parameters for Certificate.Verify.
type VerifyOptions struct {
	// DNSName, if set, is checked against the leaf certificate with
	// Certificate.VerifyHostname or the platform verifier.
	DNSName	string

	// Intermediates is an optional pool of certificates that are not trust
	// anchors, but can be used to form a chain from the leaf certificate to a
	// root certificate.
	Intermediates	*CertPool
	// Roots is the set of trusted root certificates the leaf certificate needs
	// to chain up to. If nil, the system roots or the platform verifier are used.
	Roots	*CertPool

	// CurrentTime is used to check the validity of all certificates in the
	// chain. If zero, the current time is used.
	CurrentTime	time.Time

	// KeyUsages specifies which Extended Key Usage values are acceptable. A
	// chain is accepted if it allows any of the listed values. An empty list
	// means ExtKeyUsageServerAuth. To accept any key usage, include ExtKeyUsageAny.
	KeyUsages	[]ExtKeyUsage

	// MaxConstraintComparisions is the maximum number of comparisons to
	// perform when checking a given certificate's name constraints. If
	// zero, a sensible default is used. This limit prevents pathological
	// certificates from consuming excessive amounts of CPU time when
	// validating. It does not apply to the platform verifier.
	MaxConstraintComparisions	int
}

const (
	leafCertificate	= iota
	intermediateCertificate
	rootCertificate
)

// rfc2821Mailbox represents a “mailbox” (which is an email address to most
//line /usr/local/go/src/crypto/x509/verify.go:212
// people) by breaking it into the “local” (i.e. before the '@') and “domain”
//line /usr/local/go/src/crypto/x509/verify.go:212
// parts.
//line /usr/local/go/src/crypto/x509/verify.go:215
type rfc2821Mailbox struct {
	local, domain string
}

// parseRFC2821Mailbox parses an email address into local and domain parts,
//line /usr/local/go/src/crypto/x509/verify.go:219
// based on the ABNF for a “Mailbox” from RFC 2821. According to RFC 5280,
//line /usr/local/go/src/crypto/x509/verify.go:219
// Section 4.2.1.6 that's correct for an rfc822Name from a certificate: “The
//line /usr/local/go/src/crypto/x509/verify.go:219
// format of an rfc822Name is a "Mailbox" as defined in RFC 2821, Section 4.1.2”.
//line /usr/local/go/src/crypto/x509/verify.go:223
func parseRFC2821Mailbox(in string) (mailbox rfc2821Mailbox, ok bool) {
//line /usr/local/go/src/crypto/x509/verify.go:223
	_go_fuzz_dep_.CoverTab[19393]++
							if len(in) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:224
		_go_fuzz_dep_.CoverTab[19398]++
								return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:225
		// _ = "end of CoverTab[19398]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:226
		_go_fuzz_dep_.CoverTab[19399]++
//line /usr/local/go/src/crypto/x509/verify.go:226
		// _ = "end of CoverTab[19399]"
//line /usr/local/go/src/crypto/x509/verify.go:226
	}
//line /usr/local/go/src/crypto/x509/verify.go:226
	// _ = "end of CoverTab[19393]"
//line /usr/local/go/src/crypto/x509/verify.go:226
	_go_fuzz_dep_.CoverTab[19394]++

							localPartBytes := make([]byte, 0, len(in)/2)

							if in[0] == '"' {
//line /usr/local/go/src/crypto/x509/verify.go:230
		_go_fuzz_dep_.CoverTab[19400]++

//line /usr/local/go/src/crypto/x509/verify.go:241
		in = in[1:]
	QuotedString:
		for {
//line /usr/local/go/src/crypto/x509/verify.go:243
			_go_fuzz_dep_.CoverTab[19401]++
									if len(in) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:244
				_go_fuzz_dep_.CoverTab[19403]++
										return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:245
				// _ = "end of CoverTab[19403]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:246
				_go_fuzz_dep_.CoverTab[19404]++
//line /usr/local/go/src/crypto/x509/verify.go:246
				// _ = "end of CoverTab[19404]"
//line /usr/local/go/src/crypto/x509/verify.go:246
			}
//line /usr/local/go/src/crypto/x509/verify.go:246
			// _ = "end of CoverTab[19401]"
//line /usr/local/go/src/crypto/x509/verify.go:246
			_go_fuzz_dep_.CoverTab[19402]++
									c := in[0]
									in = in[1:]

									switch {
			case c == '"':
//line /usr/local/go/src/crypto/x509/verify.go:251
				_go_fuzz_dep_.CoverTab[19405]++
										break QuotedString
//line /usr/local/go/src/crypto/x509/verify.go:252
				// _ = "end of CoverTab[19405]"

			case c == '\\':
//line /usr/local/go/src/crypto/x509/verify.go:254
				_go_fuzz_dep_.CoverTab[19406]++

										if len(in) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:256
					_go_fuzz_dep_.CoverTab[19410]++
											return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:257
					// _ = "end of CoverTab[19410]"
				} else {
//line /usr/local/go/src/crypto/x509/verify.go:258
					_go_fuzz_dep_.CoverTab[19411]++
//line /usr/local/go/src/crypto/x509/verify.go:258
					// _ = "end of CoverTab[19411]"
//line /usr/local/go/src/crypto/x509/verify.go:258
				}
//line /usr/local/go/src/crypto/x509/verify.go:258
				// _ = "end of CoverTab[19406]"
//line /usr/local/go/src/crypto/x509/verify.go:258
				_go_fuzz_dep_.CoverTab[19407]++
										if in[0] == 11 || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:259
					_go_fuzz_dep_.CoverTab[19412]++
//line /usr/local/go/src/crypto/x509/verify.go:259
					return in[0] == 12
											// _ = "end of CoverTab[19412]"
//line /usr/local/go/src/crypto/x509/verify.go:260
				}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:260
					_go_fuzz_dep_.CoverTab[19413]++
//line /usr/local/go/src/crypto/x509/verify.go:260
					return (1 <= in[0] && func() bool {
												_go_fuzz_dep_.CoverTab[19414]++
//line /usr/local/go/src/crypto/x509/verify.go:261
						return in[0] <= 9
//line /usr/local/go/src/crypto/x509/verify.go:261
						// _ = "end of CoverTab[19414]"
//line /usr/local/go/src/crypto/x509/verify.go:261
					}())
//line /usr/local/go/src/crypto/x509/verify.go:261
					// _ = "end of CoverTab[19413]"
//line /usr/local/go/src/crypto/x509/verify.go:261
				}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:261
					_go_fuzz_dep_.CoverTab[19415]++
//line /usr/local/go/src/crypto/x509/verify.go:261
					return (14 <= in[0] && func() bool {
												_go_fuzz_dep_.CoverTab[19416]++
//line /usr/local/go/src/crypto/x509/verify.go:262
						return in[0] <= 127
//line /usr/local/go/src/crypto/x509/verify.go:262
						// _ = "end of CoverTab[19416]"
//line /usr/local/go/src/crypto/x509/verify.go:262
					}())
//line /usr/local/go/src/crypto/x509/verify.go:262
					// _ = "end of CoverTab[19415]"
//line /usr/local/go/src/crypto/x509/verify.go:262
				}() {
//line /usr/local/go/src/crypto/x509/verify.go:262
					_go_fuzz_dep_.CoverTab[19417]++
											localPartBytes = append(localPartBytes, in[0])
											in = in[1:]
//line /usr/local/go/src/crypto/x509/verify.go:264
					// _ = "end of CoverTab[19417]"
				} else {
//line /usr/local/go/src/crypto/x509/verify.go:265
					_go_fuzz_dep_.CoverTab[19418]++
											return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:266
					// _ = "end of CoverTab[19418]"
				}
//line /usr/local/go/src/crypto/x509/verify.go:267
				// _ = "end of CoverTab[19407]"

			case c == 11 || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:269
				_go_fuzz_dep_.CoverTab[19419]++
//line /usr/local/go/src/crypto/x509/verify.go:269
				return c == 12
										// _ = "end of CoverTab[19419]"
//line /usr/local/go/src/crypto/x509/verify.go:270
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:270
				_go_fuzz_dep_.CoverTab[19420]++
//line /usr/local/go/src/crypto/x509/verify.go:270
				return c == 32
//line /usr/local/go/src/crypto/x509/verify.go:276
				// _ = "end of CoverTab[19420]"
//line /usr/local/go/src/crypto/x509/verify.go:276
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:276
				_go_fuzz_dep_.CoverTab[19421]++
//line /usr/local/go/src/crypto/x509/verify.go:276
				return c == 33
										// _ = "end of CoverTab[19421]"
//line /usr/local/go/src/crypto/x509/verify.go:277
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:277
				_go_fuzz_dep_.CoverTab[19422]++
//line /usr/local/go/src/crypto/x509/verify.go:277
				return c == 127
										// _ = "end of CoverTab[19422]"
//line /usr/local/go/src/crypto/x509/verify.go:278
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:278
				_go_fuzz_dep_.CoverTab[19423]++
//line /usr/local/go/src/crypto/x509/verify.go:278
				return (1 <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19424]++
//line /usr/local/go/src/crypto/x509/verify.go:279
					return c <= 8
//line /usr/local/go/src/crypto/x509/verify.go:279
					// _ = "end of CoverTab[19424]"
//line /usr/local/go/src/crypto/x509/verify.go:279
				}())
//line /usr/local/go/src/crypto/x509/verify.go:279
				// _ = "end of CoverTab[19423]"
//line /usr/local/go/src/crypto/x509/verify.go:279
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:279
				_go_fuzz_dep_.CoverTab[19425]++
//line /usr/local/go/src/crypto/x509/verify.go:279
				return (14 <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19426]++
//line /usr/local/go/src/crypto/x509/verify.go:280
					return c <= 31
//line /usr/local/go/src/crypto/x509/verify.go:280
					// _ = "end of CoverTab[19426]"
//line /usr/local/go/src/crypto/x509/verify.go:280
				}())
//line /usr/local/go/src/crypto/x509/verify.go:280
				// _ = "end of CoverTab[19425]"
//line /usr/local/go/src/crypto/x509/verify.go:280
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:280
				_go_fuzz_dep_.CoverTab[19427]++
//line /usr/local/go/src/crypto/x509/verify.go:280
				return (35 <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19428]++
//line /usr/local/go/src/crypto/x509/verify.go:281
					return c <= 91
//line /usr/local/go/src/crypto/x509/verify.go:281
					// _ = "end of CoverTab[19428]"
//line /usr/local/go/src/crypto/x509/verify.go:281
				}())
//line /usr/local/go/src/crypto/x509/verify.go:281
				// _ = "end of CoverTab[19427]"
//line /usr/local/go/src/crypto/x509/verify.go:281
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:281
				_go_fuzz_dep_.CoverTab[19429]++
//line /usr/local/go/src/crypto/x509/verify.go:281
				return (93 <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19430]++
//line /usr/local/go/src/crypto/x509/verify.go:282
					return c <= 126
//line /usr/local/go/src/crypto/x509/verify.go:282
					// _ = "end of CoverTab[19430]"
//line /usr/local/go/src/crypto/x509/verify.go:282
				}())
//line /usr/local/go/src/crypto/x509/verify.go:282
				// _ = "end of CoverTab[19429]"
//line /usr/local/go/src/crypto/x509/verify.go:282
			}():
//line /usr/local/go/src/crypto/x509/verify.go:282
				_go_fuzz_dep_.CoverTab[19408]++

										localPartBytes = append(localPartBytes, c)
//line /usr/local/go/src/crypto/x509/verify.go:284
				// _ = "end of CoverTab[19408]"

			default:
//line /usr/local/go/src/crypto/x509/verify.go:286
				_go_fuzz_dep_.CoverTab[19409]++
										return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:287
				// _ = "end of CoverTab[19409]"
			}
//line /usr/local/go/src/crypto/x509/verify.go:288
			// _ = "end of CoverTab[19402]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:289
		// _ = "end of CoverTab[19400]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:290
		_go_fuzz_dep_.CoverTab[19431]++

	NextChar:
		for len(in) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:293
			_go_fuzz_dep_.CoverTab[19434]++

									c := in[0]

									switch {
			case c == '\\':
//line /usr/local/go/src/crypto/x509/verify.go:298
				_go_fuzz_dep_.CoverTab[19435]++

//line /usr/local/go/src/crypto/x509/verify.go:304
				in = in[1:]
				if len(in) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:305
					_go_fuzz_dep_.CoverTab[19439]++
											return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:306
					// _ = "end of CoverTab[19439]"
				} else {
//line /usr/local/go/src/crypto/x509/verify.go:307
					_go_fuzz_dep_.CoverTab[19440]++
//line /usr/local/go/src/crypto/x509/verify.go:307
					// _ = "end of CoverTab[19440]"
//line /usr/local/go/src/crypto/x509/verify.go:307
				}
//line /usr/local/go/src/crypto/x509/verify.go:307
				// _ = "end of CoverTab[19435]"
//line /usr/local/go/src/crypto/x509/verify.go:307
				_go_fuzz_dep_.CoverTab[19436]++
										fallthrough
//line /usr/local/go/src/crypto/x509/verify.go:308
				// _ = "end of CoverTab[19436]"

			case ('0' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:310
				_go_fuzz_dep_.CoverTab[19441]++
//line /usr/local/go/src/crypto/x509/verify.go:310
				return c <= '9'
//line /usr/local/go/src/crypto/x509/verify.go:310
				// _ = "end of CoverTab[19441]"
//line /usr/local/go/src/crypto/x509/verify.go:310
			}()) || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:310
				_go_fuzz_dep_.CoverTab[19442]++
//line /usr/local/go/src/crypto/x509/verify.go:310
				return ('a' <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19443]++
//line /usr/local/go/src/crypto/x509/verify.go:311
					return c <= 'z'
//line /usr/local/go/src/crypto/x509/verify.go:311
					// _ = "end of CoverTab[19443]"
//line /usr/local/go/src/crypto/x509/verify.go:311
				}())
//line /usr/local/go/src/crypto/x509/verify.go:311
				// _ = "end of CoverTab[19442]"
//line /usr/local/go/src/crypto/x509/verify.go:311
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:311
				_go_fuzz_dep_.CoverTab[19444]++
//line /usr/local/go/src/crypto/x509/verify.go:311
				return ('A' <= c && func() bool {
											_go_fuzz_dep_.CoverTab[19445]++
//line /usr/local/go/src/crypto/x509/verify.go:312
					return c <= 'Z'
//line /usr/local/go/src/crypto/x509/verify.go:312
					// _ = "end of CoverTab[19445]"
//line /usr/local/go/src/crypto/x509/verify.go:312
				}())
//line /usr/local/go/src/crypto/x509/verify.go:312
				// _ = "end of CoverTab[19444]"
//line /usr/local/go/src/crypto/x509/verify.go:312
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:312
				_go_fuzz_dep_.CoverTab[19446]++
//line /usr/local/go/src/crypto/x509/verify.go:312
				return c == '!'
										// _ = "end of CoverTab[19446]"
//line /usr/local/go/src/crypto/x509/verify.go:313
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:313
				_go_fuzz_dep_.CoverTab[19447]++
//line /usr/local/go/src/crypto/x509/verify.go:313
				return c == '#'
//line /usr/local/go/src/crypto/x509/verify.go:313
				// _ = "end of CoverTab[19447]"
//line /usr/local/go/src/crypto/x509/verify.go:313
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:313
				_go_fuzz_dep_.CoverTab[19448]++
//line /usr/local/go/src/crypto/x509/verify.go:313
				return c == '$'
//line /usr/local/go/src/crypto/x509/verify.go:313
				// _ = "end of CoverTab[19448]"
//line /usr/local/go/src/crypto/x509/verify.go:313
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:313
				_go_fuzz_dep_.CoverTab[19449]++
//line /usr/local/go/src/crypto/x509/verify.go:313
				return c == '%'
//line /usr/local/go/src/crypto/x509/verify.go:313
				// _ = "end of CoverTab[19449]"
//line /usr/local/go/src/crypto/x509/verify.go:313
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:313
				_go_fuzz_dep_.CoverTab[19450]++
//line /usr/local/go/src/crypto/x509/verify.go:313
				return c == '&'
										// _ = "end of CoverTab[19450]"
//line /usr/local/go/src/crypto/x509/verify.go:314
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:314
				_go_fuzz_dep_.CoverTab[19451]++
//line /usr/local/go/src/crypto/x509/verify.go:314
				return c == '\''
//line /usr/local/go/src/crypto/x509/verify.go:314
				// _ = "end of CoverTab[19451]"
//line /usr/local/go/src/crypto/x509/verify.go:314
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:314
				_go_fuzz_dep_.CoverTab[19452]++
//line /usr/local/go/src/crypto/x509/verify.go:314
				return c == '*'
//line /usr/local/go/src/crypto/x509/verify.go:314
				// _ = "end of CoverTab[19452]"
//line /usr/local/go/src/crypto/x509/verify.go:314
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:314
				_go_fuzz_dep_.CoverTab[19453]++
//line /usr/local/go/src/crypto/x509/verify.go:314
				return c == '+'
//line /usr/local/go/src/crypto/x509/verify.go:314
				// _ = "end of CoverTab[19453]"
//line /usr/local/go/src/crypto/x509/verify.go:314
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:314
				_go_fuzz_dep_.CoverTab[19454]++
//line /usr/local/go/src/crypto/x509/verify.go:314
				return c == '-'
										// _ = "end of CoverTab[19454]"
//line /usr/local/go/src/crypto/x509/verify.go:315
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:315
				_go_fuzz_dep_.CoverTab[19455]++
//line /usr/local/go/src/crypto/x509/verify.go:315
				return c == '/'
//line /usr/local/go/src/crypto/x509/verify.go:315
				// _ = "end of CoverTab[19455]"
//line /usr/local/go/src/crypto/x509/verify.go:315
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:315
				_go_fuzz_dep_.CoverTab[19456]++
//line /usr/local/go/src/crypto/x509/verify.go:315
				return c == '='
//line /usr/local/go/src/crypto/x509/verify.go:315
				// _ = "end of CoverTab[19456]"
//line /usr/local/go/src/crypto/x509/verify.go:315
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:315
				_go_fuzz_dep_.CoverTab[19457]++
//line /usr/local/go/src/crypto/x509/verify.go:315
				return c == '?'
//line /usr/local/go/src/crypto/x509/verify.go:315
				// _ = "end of CoverTab[19457]"
//line /usr/local/go/src/crypto/x509/verify.go:315
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:315
				_go_fuzz_dep_.CoverTab[19458]++
//line /usr/local/go/src/crypto/x509/verify.go:315
				return c == '^'
										// _ = "end of CoverTab[19458]"
//line /usr/local/go/src/crypto/x509/verify.go:316
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:316
				_go_fuzz_dep_.CoverTab[19459]++
//line /usr/local/go/src/crypto/x509/verify.go:316
				return c == '_'
//line /usr/local/go/src/crypto/x509/verify.go:316
				// _ = "end of CoverTab[19459]"
//line /usr/local/go/src/crypto/x509/verify.go:316
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:316
				_go_fuzz_dep_.CoverTab[19460]++
//line /usr/local/go/src/crypto/x509/verify.go:316
				return c == '`'
//line /usr/local/go/src/crypto/x509/verify.go:316
				// _ = "end of CoverTab[19460]"
//line /usr/local/go/src/crypto/x509/verify.go:316
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:316
				_go_fuzz_dep_.CoverTab[19461]++
//line /usr/local/go/src/crypto/x509/verify.go:316
				return c == '{'
//line /usr/local/go/src/crypto/x509/verify.go:316
				// _ = "end of CoverTab[19461]"
//line /usr/local/go/src/crypto/x509/verify.go:316
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:316
				_go_fuzz_dep_.CoverTab[19462]++
//line /usr/local/go/src/crypto/x509/verify.go:316
				return c == '|'
										// _ = "end of CoverTab[19462]"
//line /usr/local/go/src/crypto/x509/verify.go:317
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:317
				_go_fuzz_dep_.CoverTab[19463]++
//line /usr/local/go/src/crypto/x509/verify.go:317
				return c == '}'
//line /usr/local/go/src/crypto/x509/verify.go:317
				// _ = "end of CoverTab[19463]"
//line /usr/local/go/src/crypto/x509/verify.go:317
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:317
				_go_fuzz_dep_.CoverTab[19464]++
//line /usr/local/go/src/crypto/x509/verify.go:317
				return c == '~'
//line /usr/local/go/src/crypto/x509/verify.go:317
				// _ = "end of CoverTab[19464]"
//line /usr/local/go/src/crypto/x509/verify.go:317
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:317
				_go_fuzz_dep_.CoverTab[19465]++
//line /usr/local/go/src/crypto/x509/verify.go:317
				return c == '.'
//line /usr/local/go/src/crypto/x509/verify.go:317
				// _ = "end of CoverTab[19465]"
//line /usr/local/go/src/crypto/x509/verify.go:317
			}():
//line /usr/local/go/src/crypto/x509/verify.go:317
				_go_fuzz_dep_.CoverTab[19437]++
										localPartBytes = append(localPartBytes, in[0])
										in = in[1:]
//line /usr/local/go/src/crypto/x509/verify.go:319
				// _ = "end of CoverTab[19437]"

			default:
//line /usr/local/go/src/crypto/x509/verify.go:321
				_go_fuzz_dep_.CoverTab[19438]++
										break NextChar
//line /usr/local/go/src/crypto/x509/verify.go:322
				// _ = "end of CoverTab[19438]"
			}
//line /usr/local/go/src/crypto/x509/verify.go:323
			// _ = "end of CoverTab[19434]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:324
		// _ = "end of CoverTab[19431]"
//line /usr/local/go/src/crypto/x509/verify.go:324
		_go_fuzz_dep_.CoverTab[19432]++

								if len(localPartBytes) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:326
			_go_fuzz_dep_.CoverTab[19466]++
									return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:327
			// _ = "end of CoverTab[19466]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:328
			_go_fuzz_dep_.CoverTab[19467]++
//line /usr/local/go/src/crypto/x509/verify.go:328
			// _ = "end of CoverTab[19467]"
//line /usr/local/go/src/crypto/x509/verify.go:328
		}
//line /usr/local/go/src/crypto/x509/verify.go:328
		// _ = "end of CoverTab[19432]"
//line /usr/local/go/src/crypto/x509/verify.go:328
		_go_fuzz_dep_.CoverTab[19433]++

//line /usr/local/go/src/crypto/x509/verify.go:334
		twoDots := []byte{'.', '.'}
		if localPartBytes[0] == '.' || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:335
			_go_fuzz_dep_.CoverTab[19468]++
//line /usr/local/go/src/crypto/x509/verify.go:335
			return localPartBytes[len(localPartBytes)-1] == '.'
									// _ = "end of CoverTab[19468]"
//line /usr/local/go/src/crypto/x509/verify.go:336
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:336
			_go_fuzz_dep_.CoverTab[19469]++
//line /usr/local/go/src/crypto/x509/verify.go:336
			return bytes.Contains(localPartBytes, twoDots)
									// _ = "end of CoverTab[19469]"
//line /usr/local/go/src/crypto/x509/verify.go:337
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:337
			_go_fuzz_dep_.CoverTab[19470]++
									return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:338
			// _ = "end of CoverTab[19470]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:339
			_go_fuzz_dep_.CoverTab[19471]++
//line /usr/local/go/src/crypto/x509/verify.go:339
			// _ = "end of CoverTab[19471]"
//line /usr/local/go/src/crypto/x509/verify.go:339
		}
//line /usr/local/go/src/crypto/x509/verify.go:339
		// _ = "end of CoverTab[19433]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:340
	// _ = "end of CoverTab[19394]"
//line /usr/local/go/src/crypto/x509/verify.go:340
	_go_fuzz_dep_.CoverTab[19395]++

							if len(in) == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:342
		_go_fuzz_dep_.CoverTab[19472]++
//line /usr/local/go/src/crypto/x509/verify.go:342
		return in[0] != '@'
//line /usr/local/go/src/crypto/x509/verify.go:342
		// _ = "end of CoverTab[19472]"
//line /usr/local/go/src/crypto/x509/verify.go:342
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:342
		_go_fuzz_dep_.CoverTab[19473]++
								return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:343
		// _ = "end of CoverTab[19473]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:344
		_go_fuzz_dep_.CoverTab[19474]++
//line /usr/local/go/src/crypto/x509/verify.go:344
		// _ = "end of CoverTab[19474]"
//line /usr/local/go/src/crypto/x509/verify.go:344
	}
//line /usr/local/go/src/crypto/x509/verify.go:344
	// _ = "end of CoverTab[19395]"
//line /usr/local/go/src/crypto/x509/verify.go:344
	_go_fuzz_dep_.CoverTab[19396]++
							in = in[1:]

//line /usr/local/go/src/crypto/x509/verify.go:350
	if _, ok := domainToReverseLabels(in); !ok {
//line /usr/local/go/src/crypto/x509/verify.go:350
		_go_fuzz_dep_.CoverTab[19475]++
								return mailbox, false
//line /usr/local/go/src/crypto/x509/verify.go:351
		// _ = "end of CoverTab[19475]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:352
		_go_fuzz_dep_.CoverTab[19476]++
//line /usr/local/go/src/crypto/x509/verify.go:352
		// _ = "end of CoverTab[19476]"
//line /usr/local/go/src/crypto/x509/verify.go:352
	}
//line /usr/local/go/src/crypto/x509/verify.go:352
	// _ = "end of CoverTab[19396]"
//line /usr/local/go/src/crypto/x509/verify.go:352
	_go_fuzz_dep_.CoverTab[19397]++

							mailbox.local = string(localPartBytes)
							mailbox.domain = in
							return mailbox, true
//line /usr/local/go/src/crypto/x509/verify.go:356
	// _ = "end of CoverTab[19397]"
}

// domainToReverseLabels converts a textual domain name like foo.example.com to
//line /usr/local/go/src/crypto/x509/verify.go:359
// the list of labels in reverse order, e.g. ["com", "example", "foo"].
//line /usr/local/go/src/crypto/x509/verify.go:361
func domainToReverseLabels(domain string) (reverseLabels []string, ok bool) {
//line /usr/local/go/src/crypto/x509/verify.go:361
	_go_fuzz_dep_.CoverTab[19477]++
							for len(domain) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:362
		_go_fuzz_dep_.CoverTab[19481]++
								if i := strings.LastIndexByte(domain, '.'); i == -1 {
//line /usr/local/go/src/crypto/x509/verify.go:363
			_go_fuzz_dep_.CoverTab[19482]++
									reverseLabels = append(reverseLabels, domain)
									domain = ""
//line /usr/local/go/src/crypto/x509/verify.go:365
			// _ = "end of CoverTab[19482]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:366
			_go_fuzz_dep_.CoverTab[19483]++
									reverseLabels = append(reverseLabels, domain[i+1:])
									domain = domain[:i]
//line /usr/local/go/src/crypto/x509/verify.go:368
			// _ = "end of CoverTab[19483]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:369
		// _ = "end of CoverTab[19481]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:370
	// _ = "end of CoverTab[19477]"
//line /usr/local/go/src/crypto/x509/verify.go:370
	_go_fuzz_dep_.CoverTab[19478]++

							if len(reverseLabels) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:372
		_go_fuzz_dep_.CoverTab[19484]++
//line /usr/local/go/src/crypto/x509/verify.go:372
		return len(reverseLabels[0]) == 0
//line /usr/local/go/src/crypto/x509/verify.go:372
		// _ = "end of CoverTab[19484]"
//line /usr/local/go/src/crypto/x509/verify.go:372
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:372
		_go_fuzz_dep_.CoverTab[19485]++

								return nil, false
//line /usr/local/go/src/crypto/x509/verify.go:374
		// _ = "end of CoverTab[19485]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:375
		_go_fuzz_dep_.CoverTab[19486]++
//line /usr/local/go/src/crypto/x509/verify.go:375
		// _ = "end of CoverTab[19486]"
//line /usr/local/go/src/crypto/x509/verify.go:375
	}
//line /usr/local/go/src/crypto/x509/verify.go:375
	// _ = "end of CoverTab[19478]"
//line /usr/local/go/src/crypto/x509/verify.go:375
	_go_fuzz_dep_.CoverTab[19479]++

							for _, label := range reverseLabels {
//line /usr/local/go/src/crypto/x509/verify.go:377
		_go_fuzz_dep_.CoverTab[19487]++
								if len(label) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:378
			_go_fuzz_dep_.CoverTab[19489]++

									return nil, false
//line /usr/local/go/src/crypto/x509/verify.go:380
			// _ = "end of CoverTab[19489]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:381
			_go_fuzz_dep_.CoverTab[19490]++
//line /usr/local/go/src/crypto/x509/verify.go:381
			// _ = "end of CoverTab[19490]"
//line /usr/local/go/src/crypto/x509/verify.go:381
		}
//line /usr/local/go/src/crypto/x509/verify.go:381
		// _ = "end of CoverTab[19487]"
//line /usr/local/go/src/crypto/x509/verify.go:381
		_go_fuzz_dep_.CoverTab[19488]++

								for _, c := range label {
//line /usr/local/go/src/crypto/x509/verify.go:383
			_go_fuzz_dep_.CoverTab[19491]++
									if c < 33 || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:384
				_go_fuzz_dep_.CoverTab[19492]++
//line /usr/local/go/src/crypto/x509/verify.go:384
				return c > 126
//line /usr/local/go/src/crypto/x509/verify.go:384
				// _ = "end of CoverTab[19492]"
//line /usr/local/go/src/crypto/x509/verify.go:384
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:384
				_go_fuzz_dep_.CoverTab[19493]++

										return nil, false
//line /usr/local/go/src/crypto/x509/verify.go:386
				// _ = "end of CoverTab[19493]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:387
				_go_fuzz_dep_.CoverTab[19494]++
//line /usr/local/go/src/crypto/x509/verify.go:387
				// _ = "end of CoverTab[19494]"
//line /usr/local/go/src/crypto/x509/verify.go:387
			}
//line /usr/local/go/src/crypto/x509/verify.go:387
			// _ = "end of CoverTab[19491]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:388
		// _ = "end of CoverTab[19488]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:389
	// _ = "end of CoverTab[19479]"
//line /usr/local/go/src/crypto/x509/verify.go:389
	_go_fuzz_dep_.CoverTab[19480]++

							return reverseLabels, true
//line /usr/local/go/src/crypto/x509/verify.go:391
	// _ = "end of CoverTab[19480]"
}

func matchEmailConstraint(mailbox rfc2821Mailbox, constraint string) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:394
	_go_fuzz_dep_.CoverTab[19495]++

//line /usr/local/go/src/crypto/x509/verify.go:397
	if strings.Contains(constraint, "@") {
//line /usr/local/go/src/crypto/x509/verify.go:397
		_go_fuzz_dep_.CoverTab[19497]++
								constraintMailbox, ok := parseRFC2821Mailbox(constraint)
								if !ok {
//line /usr/local/go/src/crypto/x509/verify.go:399
			_go_fuzz_dep_.CoverTab[19499]++
									return false, fmt.Errorf("x509: internal error: cannot parse constraint %q", constraint)
//line /usr/local/go/src/crypto/x509/verify.go:400
			// _ = "end of CoverTab[19499]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:401
			_go_fuzz_dep_.CoverTab[19500]++
//line /usr/local/go/src/crypto/x509/verify.go:401
			// _ = "end of CoverTab[19500]"
//line /usr/local/go/src/crypto/x509/verify.go:401
		}
//line /usr/local/go/src/crypto/x509/verify.go:401
		// _ = "end of CoverTab[19497]"
//line /usr/local/go/src/crypto/x509/verify.go:401
		_go_fuzz_dep_.CoverTab[19498]++
								return mailbox.local == constraintMailbox.local && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:402
			_go_fuzz_dep_.CoverTab[19501]++
//line /usr/local/go/src/crypto/x509/verify.go:402
			return strings.EqualFold(mailbox.domain, constraintMailbox.domain)
//line /usr/local/go/src/crypto/x509/verify.go:402
			// _ = "end of CoverTab[19501]"
//line /usr/local/go/src/crypto/x509/verify.go:402
		}(), nil
//line /usr/local/go/src/crypto/x509/verify.go:402
		// _ = "end of CoverTab[19498]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:403
		_go_fuzz_dep_.CoverTab[19502]++
//line /usr/local/go/src/crypto/x509/verify.go:403
		// _ = "end of CoverTab[19502]"
//line /usr/local/go/src/crypto/x509/verify.go:403
	}
//line /usr/local/go/src/crypto/x509/verify.go:403
	// _ = "end of CoverTab[19495]"
//line /usr/local/go/src/crypto/x509/verify.go:403
	_go_fuzz_dep_.CoverTab[19496]++

//line /usr/local/go/src/crypto/x509/verify.go:407
	return matchDomainConstraint(mailbox.domain, constraint)
//line /usr/local/go/src/crypto/x509/verify.go:407
	// _ = "end of CoverTab[19496]"
}

func matchURIConstraint(uri *url.URL, constraint string) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:410
	_go_fuzz_dep_.CoverTab[19503]++

//line /usr/local/go/src/crypto/x509/verify.go:419
	host := uri.Host
	if len(host) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:420
		_go_fuzz_dep_.CoverTab[19507]++
								return false, fmt.Errorf("URI with empty host (%q) cannot be matched against constraints", uri.String())
//line /usr/local/go/src/crypto/x509/verify.go:421
		// _ = "end of CoverTab[19507]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:422
		_go_fuzz_dep_.CoverTab[19508]++
//line /usr/local/go/src/crypto/x509/verify.go:422
		// _ = "end of CoverTab[19508]"
//line /usr/local/go/src/crypto/x509/verify.go:422
	}
//line /usr/local/go/src/crypto/x509/verify.go:422
	// _ = "end of CoverTab[19503]"
//line /usr/local/go/src/crypto/x509/verify.go:422
	_go_fuzz_dep_.CoverTab[19504]++

							if strings.Contains(host, ":") && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:424
		_go_fuzz_dep_.CoverTab[19509]++
//line /usr/local/go/src/crypto/x509/verify.go:424
		return !strings.HasSuffix(host, "]")
//line /usr/local/go/src/crypto/x509/verify.go:424
		// _ = "end of CoverTab[19509]"
//line /usr/local/go/src/crypto/x509/verify.go:424
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:424
		_go_fuzz_dep_.CoverTab[19510]++
								var err error
								host, _, err = net.SplitHostPort(uri.Host)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:427
			_go_fuzz_dep_.CoverTab[19511]++
									return false, err
//line /usr/local/go/src/crypto/x509/verify.go:428
			// _ = "end of CoverTab[19511]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:429
			_go_fuzz_dep_.CoverTab[19512]++
//line /usr/local/go/src/crypto/x509/verify.go:429
			// _ = "end of CoverTab[19512]"
//line /usr/local/go/src/crypto/x509/verify.go:429
		}
//line /usr/local/go/src/crypto/x509/verify.go:429
		// _ = "end of CoverTab[19510]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:430
		_go_fuzz_dep_.CoverTab[19513]++
//line /usr/local/go/src/crypto/x509/verify.go:430
		// _ = "end of CoverTab[19513]"
//line /usr/local/go/src/crypto/x509/verify.go:430
	}
//line /usr/local/go/src/crypto/x509/verify.go:430
	// _ = "end of CoverTab[19504]"
//line /usr/local/go/src/crypto/x509/verify.go:430
	_go_fuzz_dep_.CoverTab[19505]++

							if strings.HasPrefix(host, "[") && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:432
		_go_fuzz_dep_.CoverTab[19514]++
//line /usr/local/go/src/crypto/x509/verify.go:432
		return strings.HasSuffix(host, "]")
//line /usr/local/go/src/crypto/x509/verify.go:432
		// _ = "end of CoverTab[19514]"
//line /usr/local/go/src/crypto/x509/verify.go:432
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:432
		_go_fuzz_dep_.CoverTab[19515]++
//line /usr/local/go/src/crypto/x509/verify.go:432
		return net.ParseIP(host) != nil
								// _ = "end of CoverTab[19515]"
//line /usr/local/go/src/crypto/x509/verify.go:433
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:433
		_go_fuzz_dep_.CoverTab[19516]++
								return false, fmt.Errorf("URI with IP (%q) cannot be matched against constraints", uri.String())
//line /usr/local/go/src/crypto/x509/verify.go:434
		// _ = "end of CoverTab[19516]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:435
		_go_fuzz_dep_.CoverTab[19517]++
//line /usr/local/go/src/crypto/x509/verify.go:435
		// _ = "end of CoverTab[19517]"
//line /usr/local/go/src/crypto/x509/verify.go:435
	}
//line /usr/local/go/src/crypto/x509/verify.go:435
	// _ = "end of CoverTab[19505]"
//line /usr/local/go/src/crypto/x509/verify.go:435
	_go_fuzz_dep_.CoverTab[19506]++

							return matchDomainConstraint(host, constraint)
//line /usr/local/go/src/crypto/x509/verify.go:437
	// _ = "end of CoverTab[19506]"
}

func matchIPConstraint(ip net.IP, constraint *net.IPNet) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:440
	_go_fuzz_dep_.CoverTab[19518]++
							if len(ip) != len(constraint.IP) {
//line /usr/local/go/src/crypto/x509/verify.go:441
		_go_fuzz_dep_.CoverTab[19521]++
								return false, nil
//line /usr/local/go/src/crypto/x509/verify.go:442
		// _ = "end of CoverTab[19521]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:443
		_go_fuzz_dep_.CoverTab[19522]++
//line /usr/local/go/src/crypto/x509/verify.go:443
		// _ = "end of CoverTab[19522]"
//line /usr/local/go/src/crypto/x509/verify.go:443
	}
//line /usr/local/go/src/crypto/x509/verify.go:443
	// _ = "end of CoverTab[19518]"
//line /usr/local/go/src/crypto/x509/verify.go:443
	_go_fuzz_dep_.CoverTab[19519]++

							for i := range ip {
//line /usr/local/go/src/crypto/x509/verify.go:445
		_go_fuzz_dep_.CoverTab[19523]++
								if mask := constraint.Mask[i]; ip[i]&mask != constraint.IP[i]&mask {
//line /usr/local/go/src/crypto/x509/verify.go:446
			_go_fuzz_dep_.CoverTab[19524]++
									return false, nil
//line /usr/local/go/src/crypto/x509/verify.go:447
			// _ = "end of CoverTab[19524]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:448
			_go_fuzz_dep_.CoverTab[19525]++
//line /usr/local/go/src/crypto/x509/verify.go:448
			// _ = "end of CoverTab[19525]"
//line /usr/local/go/src/crypto/x509/verify.go:448
		}
//line /usr/local/go/src/crypto/x509/verify.go:448
		// _ = "end of CoverTab[19523]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:449
	// _ = "end of CoverTab[19519]"
//line /usr/local/go/src/crypto/x509/verify.go:449
	_go_fuzz_dep_.CoverTab[19520]++

							return true, nil
//line /usr/local/go/src/crypto/x509/verify.go:451
	// _ = "end of CoverTab[19520]"
}

func matchDomainConstraint(domain, constraint string) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:454
	_go_fuzz_dep_.CoverTab[19526]++

//line /usr/local/go/src/crypto/x509/verify.go:457
	if len(constraint) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:457
		_go_fuzz_dep_.CoverTab[19533]++
								return true, nil
//line /usr/local/go/src/crypto/x509/verify.go:458
		// _ = "end of CoverTab[19533]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:459
		_go_fuzz_dep_.CoverTab[19534]++
//line /usr/local/go/src/crypto/x509/verify.go:459
		// _ = "end of CoverTab[19534]"
//line /usr/local/go/src/crypto/x509/verify.go:459
	}
//line /usr/local/go/src/crypto/x509/verify.go:459
	// _ = "end of CoverTab[19526]"
//line /usr/local/go/src/crypto/x509/verify.go:459
	_go_fuzz_dep_.CoverTab[19527]++

							domainLabels, ok := domainToReverseLabels(domain)
							if !ok {
//line /usr/local/go/src/crypto/x509/verify.go:462
		_go_fuzz_dep_.CoverTab[19535]++
								return false, fmt.Errorf("x509: internal error: cannot parse domain %q", domain)
//line /usr/local/go/src/crypto/x509/verify.go:463
		// _ = "end of CoverTab[19535]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:464
		_go_fuzz_dep_.CoverTab[19536]++
//line /usr/local/go/src/crypto/x509/verify.go:464
		// _ = "end of CoverTab[19536]"
//line /usr/local/go/src/crypto/x509/verify.go:464
	}
//line /usr/local/go/src/crypto/x509/verify.go:464
	// _ = "end of CoverTab[19527]"
//line /usr/local/go/src/crypto/x509/verify.go:464
	_go_fuzz_dep_.CoverTab[19528]++

//line /usr/local/go/src/crypto/x509/verify.go:471
	mustHaveSubdomains := false
	if constraint[0] == '.' {
//line /usr/local/go/src/crypto/x509/verify.go:472
		_go_fuzz_dep_.CoverTab[19537]++
								mustHaveSubdomains = true
								constraint = constraint[1:]
//line /usr/local/go/src/crypto/x509/verify.go:474
		// _ = "end of CoverTab[19537]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:475
		_go_fuzz_dep_.CoverTab[19538]++
//line /usr/local/go/src/crypto/x509/verify.go:475
		// _ = "end of CoverTab[19538]"
//line /usr/local/go/src/crypto/x509/verify.go:475
	}
//line /usr/local/go/src/crypto/x509/verify.go:475
	// _ = "end of CoverTab[19528]"
//line /usr/local/go/src/crypto/x509/verify.go:475
	_go_fuzz_dep_.CoverTab[19529]++

							constraintLabels, ok := domainToReverseLabels(constraint)
							if !ok {
//line /usr/local/go/src/crypto/x509/verify.go:478
		_go_fuzz_dep_.CoverTab[19539]++
								return false, fmt.Errorf("x509: internal error: cannot parse domain %q", constraint)
//line /usr/local/go/src/crypto/x509/verify.go:479
		// _ = "end of CoverTab[19539]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:480
		_go_fuzz_dep_.CoverTab[19540]++
//line /usr/local/go/src/crypto/x509/verify.go:480
		// _ = "end of CoverTab[19540]"
//line /usr/local/go/src/crypto/x509/verify.go:480
	}
//line /usr/local/go/src/crypto/x509/verify.go:480
	// _ = "end of CoverTab[19529]"
//line /usr/local/go/src/crypto/x509/verify.go:480
	_go_fuzz_dep_.CoverTab[19530]++

							if len(domainLabels) < len(constraintLabels) || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:482
		_go_fuzz_dep_.CoverTab[19541]++
//line /usr/local/go/src/crypto/x509/verify.go:482
		return (mustHaveSubdomains && func() bool {
									_go_fuzz_dep_.CoverTab[19542]++
//line /usr/local/go/src/crypto/x509/verify.go:483
			return len(domainLabels) == len(constraintLabels)
//line /usr/local/go/src/crypto/x509/verify.go:483
			// _ = "end of CoverTab[19542]"
//line /usr/local/go/src/crypto/x509/verify.go:483
		}())
//line /usr/local/go/src/crypto/x509/verify.go:483
		// _ = "end of CoverTab[19541]"
//line /usr/local/go/src/crypto/x509/verify.go:483
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:483
		_go_fuzz_dep_.CoverTab[19543]++
								return false, nil
//line /usr/local/go/src/crypto/x509/verify.go:484
		// _ = "end of CoverTab[19543]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:485
		_go_fuzz_dep_.CoverTab[19544]++
//line /usr/local/go/src/crypto/x509/verify.go:485
		// _ = "end of CoverTab[19544]"
//line /usr/local/go/src/crypto/x509/verify.go:485
	}
//line /usr/local/go/src/crypto/x509/verify.go:485
	// _ = "end of CoverTab[19530]"
//line /usr/local/go/src/crypto/x509/verify.go:485
	_go_fuzz_dep_.CoverTab[19531]++

							for i, constraintLabel := range constraintLabels {
//line /usr/local/go/src/crypto/x509/verify.go:487
		_go_fuzz_dep_.CoverTab[19545]++
								if !strings.EqualFold(constraintLabel, domainLabels[i]) {
//line /usr/local/go/src/crypto/x509/verify.go:488
			_go_fuzz_dep_.CoverTab[19546]++
									return false, nil
//line /usr/local/go/src/crypto/x509/verify.go:489
			// _ = "end of CoverTab[19546]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:490
			_go_fuzz_dep_.CoverTab[19547]++
//line /usr/local/go/src/crypto/x509/verify.go:490
			// _ = "end of CoverTab[19547]"
//line /usr/local/go/src/crypto/x509/verify.go:490
		}
//line /usr/local/go/src/crypto/x509/verify.go:490
		// _ = "end of CoverTab[19545]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:491
	// _ = "end of CoverTab[19531]"
//line /usr/local/go/src/crypto/x509/verify.go:491
	_go_fuzz_dep_.CoverTab[19532]++

							return true, nil
//line /usr/local/go/src/crypto/x509/verify.go:493
	// _ = "end of CoverTab[19532]"
}

// checkNameConstraints checks that c permits a child certificate to claim the
//line /usr/local/go/src/crypto/x509/verify.go:496
// given name, of type nameType. The argument parsedName contains the parsed
//line /usr/local/go/src/crypto/x509/verify.go:496
// form of name, suitable for passing to the match function. The total number
//line /usr/local/go/src/crypto/x509/verify.go:496
// of comparisons is tracked in the given count and should not exceed the given
//line /usr/local/go/src/crypto/x509/verify.go:496
// limit.
//line /usr/local/go/src/crypto/x509/verify.go:501
func (c *Certificate) checkNameConstraints(count *int,
	maxConstraintComparisons int,
	nameType string,
	name string,
	parsedName any,
	match func(parsedName, constraint any) (match bool, err error),
	permitted, excluded any) error {
//line /usr/local/go/src/crypto/x509/verify.go:507
	_go_fuzz_dep_.CoverTab[19548]++

							excludedValue := reflect.ValueOf(excluded)

							*count += excludedValue.Len()
							if *count > maxConstraintComparisons {
//line /usr/local/go/src/crypto/x509/verify.go:512
		_go_fuzz_dep_.CoverTab[19554]++
								return CertificateInvalidError{c, TooManyConstraints, ""}
//line /usr/local/go/src/crypto/x509/verify.go:513
		// _ = "end of CoverTab[19554]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:514
		_go_fuzz_dep_.CoverTab[19555]++
//line /usr/local/go/src/crypto/x509/verify.go:514
		// _ = "end of CoverTab[19555]"
//line /usr/local/go/src/crypto/x509/verify.go:514
	}
//line /usr/local/go/src/crypto/x509/verify.go:514
	// _ = "end of CoverTab[19548]"
//line /usr/local/go/src/crypto/x509/verify.go:514
	_go_fuzz_dep_.CoverTab[19549]++

							for i := 0; i < excludedValue.Len(); i++ {
//line /usr/local/go/src/crypto/x509/verify.go:516
		_go_fuzz_dep_.CoverTab[19556]++
								constraint := excludedValue.Index(i).Interface()
								match, err := match(parsedName, constraint)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:519
			_go_fuzz_dep_.CoverTab[19558]++
									return CertificateInvalidError{c, CANotAuthorizedForThisName, err.Error()}
//line /usr/local/go/src/crypto/x509/verify.go:520
			// _ = "end of CoverTab[19558]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:521
			_go_fuzz_dep_.CoverTab[19559]++
//line /usr/local/go/src/crypto/x509/verify.go:521
			// _ = "end of CoverTab[19559]"
//line /usr/local/go/src/crypto/x509/verify.go:521
		}
//line /usr/local/go/src/crypto/x509/verify.go:521
		// _ = "end of CoverTab[19556]"
//line /usr/local/go/src/crypto/x509/verify.go:521
		_go_fuzz_dep_.CoverTab[19557]++

								if match {
//line /usr/local/go/src/crypto/x509/verify.go:523
			_go_fuzz_dep_.CoverTab[19560]++
									return CertificateInvalidError{c, CANotAuthorizedForThisName, fmt.Sprintf("%s %q is excluded by constraint %q", nameType, name, constraint)}
//line /usr/local/go/src/crypto/x509/verify.go:524
			// _ = "end of CoverTab[19560]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:525
			_go_fuzz_dep_.CoverTab[19561]++
//line /usr/local/go/src/crypto/x509/verify.go:525
			// _ = "end of CoverTab[19561]"
//line /usr/local/go/src/crypto/x509/verify.go:525
		}
//line /usr/local/go/src/crypto/x509/verify.go:525
		// _ = "end of CoverTab[19557]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:526
	// _ = "end of CoverTab[19549]"
//line /usr/local/go/src/crypto/x509/verify.go:526
	_go_fuzz_dep_.CoverTab[19550]++

							permittedValue := reflect.ValueOf(permitted)

							*count += permittedValue.Len()
							if *count > maxConstraintComparisons {
//line /usr/local/go/src/crypto/x509/verify.go:531
		_go_fuzz_dep_.CoverTab[19562]++
								return CertificateInvalidError{c, TooManyConstraints, ""}
//line /usr/local/go/src/crypto/x509/verify.go:532
		// _ = "end of CoverTab[19562]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:533
		_go_fuzz_dep_.CoverTab[19563]++
//line /usr/local/go/src/crypto/x509/verify.go:533
		// _ = "end of CoverTab[19563]"
//line /usr/local/go/src/crypto/x509/verify.go:533
	}
//line /usr/local/go/src/crypto/x509/verify.go:533
	// _ = "end of CoverTab[19550]"
//line /usr/local/go/src/crypto/x509/verify.go:533
	_go_fuzz_dep_.CoverTab[19551]++

							ok := true
							for i := 0; i < permittedValue.Len(); i++ {
//line /usr/local/go/src/crypto/x509/verify.go:536
		_go_fuzz_dep_.CoverTab[19564]++
								constraint := permittedValue.Index(i).Interface()

								var err error
								if ok, err = match(parsedName, constraint); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:540
			_go_fuzz_dep_.CoverTab[19566]++
									return CertificateInvalidError{c, CANotAuthorizedForThisName, err.Error()}
//line /usr/local/go/src/crypto/x509/verify.go:541
			// _ = "end of CoverTab[19566]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:542
			_go_fuzz_dep_.CoverTab[19567]++
//line /usr/local/go/src/crypto/x509/verify.go:542
			// _ = "end of CoverTab[19567]"
//line /usr/local/go/src/crypto/x509/verify.go:542
		}
//line /usr/local/go/src/crypto/x509/verify.go:542
		// _ = "end of CoverTab[19564]"
//line /usr/local/go/src/crypto/x509/verify.go:542
		_go_fuzz_dep_.CoverTab[19565]++

								if ok {
//line /usr/local/go/src/crypto/x509/verify.go:544
			_go_fuzz_dep_.CoverTab[19568]++
									break
//line /usr/local/go/src/crypto/x509/verify.go:545
			// _ = "end of CoverTab[19568]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:546
			_go_fuzz_dep_.CoverTab[19569]++
//line /usr/local/go/src/crypto/x509/verify.go:546
			// _ = "end of CoverTab[19569]"
//line /usr/local/go/src/crypto/x509/verify.go:546
		}
//line /usr/local/go/src/crypto/x509/verify.go:546
		// _ = "end of CoverTab[19565]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:547
	// _ = "end of CoverTab[19551]"
//line /usr/local/go/src/crypto/x509/verify.go:547
	_go_fuzz_dep_.CoverTab[19552]++

							if !ok {
//line /usr/local/go/src/crypto/x509/verify.go:549
		_go_fuzz_dep_.CoverTab[19570]++
								return CertificateInvalidError{c, CANotAuthorizedForThisName, fmt.Sprintf("%s %q is not permitted by any constraint", nameType, name)}
//line /usr/local/go/src/crypto/x509/verify.go:550
		// _ = "end of CoverTab[19570]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:551
		_go_fuzz_dep_.CoverTab[19571]++
//line /usr/local/go/src/crypto/x509/verify.go:551
		// _ = "end of CoverTab[19571]"
//line /usr/local/go/src/crypto/x509/verify.go:551
	}
//line /usr/local/go/src/crypto/x509/verify.go:551
	// _ = "end of CoverTab[19552]"
//line /usr/local/go/src/crypto/x509/verify.go:551
	_go_fuzz_dep_.CoverTab[19553]++

							return nil
//line /usr/local/go/src/crypto/x509/verify.go:553
	// _ = "end of CoverTab[19553]"
}

// isValid performs validity checks on c given that it is a candidate to append
//line /usr/local/go/src/crypto/x509/verify.go:556
// to the chain in currentChain.
//line /usr/local/go/src/crypto/x509/verify.go:558
func (c *Certificate) isValid(certType int, currentChain []*Certificate, opts *VerifyOptions) error {
//line /usr/local/go/src/crypto/x509/verify.go:558
	_go_fuzz_dep_.CoverTab[19572]++
							if len(c.UnhandledCriticalExtensions) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:559
		_go_fuzz_dep_.CoverTab[19583]++
								return UnhandledCriticalExtension{}
//line /usr/local/go/src/crypto/x509/verify.go:560
		// _ = "end of CoverTab[19583]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:561
		_go_fuzz_dep_.CoverTab[19584]++
//line /usr/local/go/src/crypto/x509/verify.go:561
		// _ = "end of CoverTab[19584]"
//line /usr/local/go/src/crypto/x509/verify.go:561
	}
//line /usr/local/go/src/crypto/x509/verify.go:561
	// _ = "end of CoverTab[19572]"
//line /usr/local/go/src/crypto/x509/verify.go:561
	_go_fuzz_dep_.CoverTab[19573]++

							if len(currentChain) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:563
		_go_fuzz_dep_.CoverTab[19585]++
								child := currentChain[len(currentChain)-1]
								if !bytes.Equal(child.RawIssuer, c.RawSubject) {
//line /usr/local/go/src/crypto/x509/verify.go:565
			_go_fuzz_dep_.CoverTab[19586]++
									return CertificateInvalidError{c, NameMismatch, ""}
//line /usr/local/go/src/crypto/x509/verify.go:566
			// _ = "end of CoverTab[19586]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:567
			_go_fuzz_dep_.CoverTab[19587]++
//line /usr/local/go/src/crypto/x509/verify.go:567
			// _ = "end of CoverTab[19587]"
//line /usr/local/go/src/crypto/x509/verify.go:567
		}
//line /usr/local/go/src/crypto/x509/verify.go:567
		// _ = "end of CoverTab[19585]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:568
		_go_fuzz_dep_.CoverTab[19588]++
//line /usr/local/go/src/crypto/x509/verify.go:568
		// _ = "end of CoverTab[19588]"
//line /usr/local/go/src/crypto/x509/verify.go:568
	}
//line /usr/local/go/src/crypto/x509/verify.go:568
	// _ = "end of CoverTab[19573]"
//line /usr/local/go/src/crypto/x509/verify.go:568
	_go_fuzz_dep_.CoverTab[19574]++

							now := opts.CurrentTime
							if now.IsZero() {
//line /usr/local/go/src/crypto/x509/verify.go:571
		_go_fuzz_dep_.CoverTab[19589]++
								now = time.Now()
//line /usr/local/go/src/crypto/x509/verify.go:572
		// _ = "end of CoverTab[19589]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:573
		_go_fuzz_dep_.CoverTab[19590]++
//line /usr/local/go/src/crypto/x509/verify.go:573
		// _ = "end of CoverTab[19590]"
//line /usr/local/go/src/crypto/x509/verify.go:573
	}
//line /usr/local/go/src/crypto/x509/verify.go:573
	// _ = "end of CoverTab[19574]"
//line /usr/local/go/src/crypto/x509/verify.go:573
	_go_fuzz_dep_.CoverTab[19575]++
							if now.Before(c.NotBefore) {
//line /usr/local/go/src/crypto/x509/verify.go:574
		_go_fuzz_dep_.CoverTab[19591]++
								return CertificateInvalidError{
			Cert:	c,
			Reason:	Expired,
			Detail:	fmt.Sprintf("current time %s is before %s", now.Format(time.RFC3339), c.NotBefore.Format(time.RFC3339)),
		}
//line /usr/local/go/src/crypto/x509/verify.go:579
		// _ = "end of CoverTab[19591]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:580
		_go_fuzz_dep_.CoverTab[19592]++
//line /usr/local/go/src/crypto/x509/verify.go:580
		if now.After(c.NotAfter) {
//line /usr/local/go/src/crypto/x509/verify.go:580
			_go_fuzz_dep_.CoverTab[19593]++
									return CertificateInvalidError{
				Cert:	c,
				Reason:	Expired,
				Detail:	fmt.Sprintf("current time %s is after %s", now.Format(time.RFC3339), c.NotAfter.Format(time.RFC3339)),
			}
//line /usr/local/go/src/crypto/x509/verify.go:585
			// _ = "end of CoverTab[19593]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:586
			_go_fuzz_dep_.CoverTab[19594]++
//line /usr/local/go/src/crypto/x509/verify.go:586
			// _ = "end of CoverTab[19594]"
//line /usr/local/go/src/crypto/x509/verify.go:586
		}
//line /usr/local/go/src/crypto/x509/verify.go:586
		// _ = "end of CoverTab[19592]"
//line /usr/local/go/src/crypto/x509/verify.go:586
	}
//line /usr/local/go/src/crypto/x509/verify.go:586
	// _ = "end of CoverTab[19575]"
//line /usr/local/go/src/crypto/x509/verify.go:586
	_go_fuzz_dep_.CoverTab[19576]++

							maxConstraintComparisons := opts.MaxConstraintComparisions
							if maxConstraintComparisons == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:589
		_go_fuzz_dep_.CoverTab[19595]++
								maxConstraintComparisons = 250000
//line /usr/local/go/src/crypto/x509/verify.go:590
		// _ = "end of CoverTab[19595]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:591
		_go_fuzz_dep_.CoverTab[19596]++
//line /usr/local/go/src/crypto/x509/verify.go:591
		// _ = "end of CoverTab[19596]"
//line /usr/local/go/src/crypto/x509/verify.go:591
	}
//line /usr/local/go/src/crypto/x509/verify.go:591
	// _ = "end of CoverTab[19576]"
//line /usr/local/go/src/crypto/x509/verify.go:591
	_go_fuzz_dep_.CoverTab[19577]++
							comparisonCount := 0

							var leaf *Certificate
							if certType == intermediateCertificate || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:595
		_go_fuzz_dep_.CoverTab[19597]++
//line /usr/local/go/src/crypto/x509/verify.go:595
		return certType == rootCertificate
//line /usr/local/go/src/crypto/x509/verify.go:595
		// _ = "end of CoverTab[19597]"
//line /usr/local/go/src/crypto/x509/verify.go:595
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:595
		_go_fuzz_dep_.CoverTab[19598]++
								if len(currentChain) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:596
			_go_fuzz_dep_.CoverTab[19600]++
									return errors.New("x509: internal error: empty chain when appending CA cert")
//line /usr/local/go/src/crypto/x509/verify.go:597
			// _ = "end of CoverTab[19600]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:598
			_go_fuzz_dep_.CoverTab[19601]++
//line /usr/local/go/src/crypto/x509/verify.go:598
			// _ = "end of CoverTab[19601]"
//line /usr/local/go/src/crypto/x509/verify.go:598
		}
//line /usr/local/go/src/crypto/x509/verify.go:598
		// _ = "end of CoverTab[19598]"
//line /usr/local/go/src/crypto/x509/verify.go:598
		_go_fuzz_dep_.CoverTab[19599]++
								leaf = currentChain[0]
//line /usr/local/go/src/crypto/x509/verify.go:599
		// _ = "end of CoverTab[19599]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:600
		_go_fuzz_dep_.CoverTab[19602]++
//line /usr/local/go/src/crypto/x509/verify.go:600
		// _ = "end of CoverTab[19602]"
//line /usr/local/go/src/crypto/x509/verify.go:600
	}
//line /usr/local/go/src/crypto/x509/verify.go:600
	// _ = "end of CoverTab[19577]"
//line /usr/local/go/src/crypto/x509/verify.go:600
	_go_fuzz_dep_.CoverTab[19578]++

							if (certType == intermediateCertificate || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:602
		_go_fuzz_dep_.CoverTab[19603]++
//line /usr/local/go/src/crypto/x509/verify.go:602
		return certType == rootCertificate
//line /usr/local/go/src/crypto/x509/verify.go:602
		// _ = "end of CoverTab[19603]"
//line /usr/local/go/src/crypto/x509/verify.go:602
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:602
		_go_fuzz_dep_.CoverTab[19604]++
//line /usr/local/go/src/crypto/x509/verify.go:602
		return c.hasNameConstraints()
								// _ = "end of CoverTab[19604]"
//line /usr/local/go/src/crypto/x509/verify.go:603
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:603
		_go_fuzz_dep_.CoverTab[19605]++
								toCheck := []*Certificate{}
								if leaf.hasSANExtension() {
//line /usr/local/go/src/crypto/x509/verify.go:605
			_go_fuzz_dep_.CoverTab[19608]++
									toCheck = append(toCheck, leaf)
//line /usr/local/go/src/crypto/x509/verify.go:606
			// _ = "end of CoverTab[19608]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:607
			_go_fuzz_dep_.CoverTab[19609]++
//line /usr/local/go/src/crypto/x509/verify.go:607
			// _ = "end of CoverTab[19609]"
//line /usr/local/go/src/crypto/x509/verify.go:607
		}
//line /usr/local/go/src/crypto/x509/verify.go:607
		// _ = "end of CoverTab[19605]"
//line /usr/local/go/src/crypto/x509/verify.go:607
		_go_fuzz_dep_.CoverTab[19606]++
								if c.hasSANExtension() {
//line /usr/local/go/src/crypto/x509/verify.go:608
			_go_fuzz_dep_.CoverTab[19610]++
									toCheck = append(toCheck, c)
//line /usr/local/go/src/crypto/x509/verify.go:609
			// _ = "end of CoverTab[19610]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:610
			_go_fuzz_dep_.CoverTab[19611]++
//line /usr/local/go/src/crypto/x509/verify.go:610
			// _ = "end of CoverTab[19611]"
//line /usr/local/go/src/crypto/x509/verify.go:610
		}
//line /usr/local/go/src/crypto/x509/verify.go:610
		// _ = "end of CoverTab[19606]"
//line /usr/local/go/src/crypto/x509/verify.go:610
		_go_fuzz_dep_.CoverTab[19607]++
								for _, sanCert := range toCheck {
//line /usr/local/go/src/crypto/x509/verify.go:611
			_go_fuzz_dep_.CoverTab[19612]++
									err := forEachSAN(sanCert.getSANExtension(), func(tag int, data []byte) error {
//line /usr/local/go/src/crypto/x509/verify.go:612
				_go_fuzz_dep_.CoverTab[19614]++
										switch tag {
				case nameTypeEmail:
//line /usr/local/go/src/crypto/x509/verify.go:614
					_go_fuzz_dep_.CoverTab[19616]++
											name := string(data)
											mailbox, ok := parseRFC2821Mailbox(name)
											if !ok {
//line /usr/local/go/src/crypto/x509/verify.go:617
						_go_fuzz_dep_.CoverTab[19625]++
												return fmt.Errorf("x509: cannot parse rfc822Name %q", mailbox)
//line /usr/local/go/src/crypto/x509/verify.go:618
						// _ = "end of CoverTab[19625]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:619
						_go_fuzz_dep_.CoverTab[19626]++
//line /usr/local/go/src/crypto/x509/verify.go:619
						// _ = "end of CoverTab[19626]"
//line /usr/local/go/src/crypto/x509/verify.go:619
					}
//line /usr/local/go/src/crypto/x509/verify.go:619
					// _ = "end of CoverTab[19616]"
//line /usr/local/go/src/crypto/x509/verify.go:619
					_go_fuzz_dep_.CoverTab[19617]++

											if err := c.checkNameConstraints(&comparisonCount, maxConstraintComparisons, "email address", name, mailbox,
						func(parsedName, constraint any) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:622
							_go_fuzz_dep_.CoverTab[19627]++
													return matchEmailConstraint(parsedName.(rfc2821Mailbox), constraint.(string))
//line /usr/local/go/src/crypto/x509/verify.go:623
							// _ = "end of CoverTab[19627]"
						}, c.PermittedEmailAddresses, c.ExcludedEmailAddresses); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:624
						_go_fuzz_dep_.CoverTab[19628]++
												return err
//line /usr/local/go/src/crypto/x509/verify.go:625
						// _ = "end of CoverTab[19628]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:626
						_go_fuzz_dep_.CoverTab[19629]++
//line /usr/local/go/src/crypto/x509/verify.go:626
						// _ = "end of CoverTab[19629]"
//line /usr/local/go/src/crypto/x509/verify.go:626
					}
//line /usr/local/go/src/crypto/x509/verify.go:626
					// _ = "end of CoverTab[19617]"

				case nameTypeDNS:
//line /usr/local/go/src/crypto/x509/verify.go:628
					_go_fuzz_dep_.CoverTab[19618]++
											name := string(data)
											if _, ok := domainToReverseLabels(name); !ok {
//line /usr/local/go/src/crypto/x509/verify.go:630
						_go_fuzz_dep_.CoverTab[19630]++
												return fmt.Errorf("x509: cannot parse dnsName %q", name)
//line /usr/local/go/src/crypto/x509/verify.go:631
						// _ = "end of CoverTab[19630]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:632
						_go_fuzz_dep_.CoverTab[19631]++
//line /usr/local/go/src/crypto/x509/verify.go:632
						// _ = "end of CoverTab[19631]"
//line /usr/local/go/src/crypto/x509/verify.go:632
					}
//line /usr/local/go/src/crypto/x509/verify.go:632
					// _ = "end of CoverTab[19618]"
//line /usr/local/go/src/crypto/x509/verify.go:632
					_go_fuzz_dep_.CoverTab[19619]++

											if err := c.checkNameConstraints(&comparisonCount, maxConstraintComparisons, "DNS name", name, name,
						func(parsedName, constraint any) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:635
							_go_fuzz_dep_.CoverTab[19632]++
													return matchDomainConstraint(parsedName.(string), constraint.(string))
//line /usr/local/go/src/crypto/x509/verify.go:636
							// _ = "end of CoverTab[19632]"
						}, c.PermittedDNSDomains, c.ExcludedDNSDomains); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:637
						_go_fuzz_dep_.CoverTab[19633]++
												return err
//line /usr/local/go/src/crypto/x509/verify.go:638
						// _ = "end of CoverTab[19633]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:639
						_go_fuzz_dep_.CoverTab[19634]++
//line /usr/local/go/src/crypto/x509/verify.go:639
						// _ = "end of CoverTab[19634]"
//line /usr/local/go/src/crypto/x509/verify.go:639
					}
//line /usr/local/go/src/crypto/x509/verify.go:639
					// _ = "end of CoverTab[19619]"

				case nameTypeURI:
//line /usr/local/go/src/crypto/x509/verify.go:641
					_go_fuzz_dep_.CoverTab[19620]++
											name := string(data)
											uri, err := url.Parse(name)
											if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:644
						_go_fuzz_dep_.CoverTab[19635]++
												return fmt.Errorf("x509: internal error: URI SAN %q failed to parse", name)
//line /usr/local/go/src/crypto/x509/verify.go:645
						// _ = "end of CoverTab[19635]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:646
						_go_fuzz_dep_.CoverTab[19636]++
//line /usr/local/go/src/crypto/x509/verify.go:646
						// _ = "end of CoverTab[19636]"
//line /usr/local/go/src/crypto/x509/verify.go:646
					}
//line /usr/local/go/src/crypto/x509/verify.go:646
					// _ = "end of CoverTab[19620]"
//line /usr/local/go/src/crypto/x509/verify.go:646
					_go_fuzz_dep_.CoverTab[19621]++

											if err := c.checkNameConstraints(&comparisonCount, maxConstraintComparisons, "URI", name, uri,
						func(parsedName, constraint any) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:649
							_go_fuzz_dep_.CoverTab[19637]++
													return matchURIConstraint(parsedName.(*url.URL), constraint.(string))
//line /usr/local/go/src/crypto/x509/verify.go:650
							// _ = "end of CoverTab[19637]"
						}, c.PermittedURIDomains, c.ExcludedURIDomains); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:651
						_go_fuzz_dep_.CoverTab[19638]++
												return err
//line /usr/local/go/src/crypto/x509/verify.go:652
						// _ = "end of CoverTab[19638]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:653
						_go_fuzz_dep_.CoverTab[19639]++
//line /usr/local/go/src/crypto/x509/verify.go:653
						// _ = "end of CoverTab[19639]"
//line /usr/local/go/src/crypto/x509/verify.go:653
					}
//line /usr/local/go/src/crypto/x509/verify.go:653
					// _ = "end of CoverTab[19621]"

				case nameTypeIP:
//line /usr/local/go/src/crypto/x509/verify.go:655
					_go_fuzz_dep_.CoverTab[19622]++
											ip := net.IP(data)
											if l := len(ip); l != net.IPv4len && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:657
						_go_fuzz_dep_.CoverTab[19640]++
//line /usr/local/go/src/crypto/x509/verify.go:657
						return l != net.IPv6len
//line /usr/local/go/src/crypto/x509/verify.go:657
						// _ = "end of CoverTab[19640]"
//line /usr/local/go/src/crypto/x509/verify.go:657
					}() {
//line /usr/local/go/src/crypto/x509/verify.go:657
						_go_fuzz_dep_.CoverTab[19641]++
												return fmt.Errorf("x509: internal error: IP SAN %x failed to parse", data)
//line /usr/local/go/src/crypto/x509/verify.go:658
						// _ = "end of CoverTab[19641]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:659
						_go_fuzz_dep_.CoverTab[19642]++
//line /usr/local/go/src/crypto/x509/verify.go:659
						// _ = "end of CoverTab[19642]"
//line /usr/local/go/src/crypto/x509/verify.go:659
					}
//line /usr/local/go/src/crypto/x509/verify.go:659
					// _ = "end of CoverTab[19622]"
//line /usr/local/go/src/crypto/x509/verify.go:659
					_go_fuzz_dep_.CoverTab[19623]++

											if err := c.checkNameConstraints(&comparisonCount, maxConstraintComparisons, "IP address", ip.String(), ip,
						func(parsedName, constraint any) (bool, error) {
//line /usr/local/go/src/crypto/x509/verify.go:662
							_go_fuzz_dep_.CoverTab[19643]++
													return matchIPConstraint(parsedName.(net.IP), constraint.(*net.IPNet))
//line /usr/local/go/src/crypto/x509/verify.go:663
							// _ = "end of CoverTab[19643]"
						}, c.PermittedIPRanges, c.ExcludedIPRanges); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:664
						_go_fuzz_dep_.CoverTab[19644]++
												return err
//line /usr/local/go/src/crypto/x509/verify.go:665
						// _ = "end of CoverTab[19644]"
					} else {
//line /usr/local/go/src/crypto/x509/verify.go:666
						_go_fuzz_dep_.CoverTab[19645]++
//line /usr/local/go/src/crypto/x509/verify.go:666
						// _ = "end of CoverTab[19645]"
//line /usr/local/go/src/crypto/x509/verify.go:666
					}
//line /usr/local/go/src/crypto/x509/verify.go:666
					// _ = "end of CoverTab[19623]"

				default:
//line /usr/local/go/src/crypto/x509/verify.go:668
					_go_fuzz_dep_.CoverTab[19624]++
//line /usr/local/go/src/crypto/x509/verify.go:668
					// _ = "end of CoverTab[19624]"

				}
//line /usr/local/go/src/crypto/x509/verify.go:670
				// _ = "end of CoverTab[19614]"
//line /usr/local/go/src/crypto/x509/verify.go:670
				_go_fuzz_dep_.CoverTab[19615]++

										return nil
//line /usr/local/go/src/crypto/x509/verify.go:672
				// _ = "end of CoverTab[19615]"
			})
//line /usr/local/go/src/crypto/x509/verify.go:673
			// _ = "end of CoverTab[19612]"
//line /usr/local/go/src/crypto/x509/verify.go:673
			_go_fuzz_dep_.CoverTab[19613]++

									if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:675
				_go_fuzz_dep_.CoverTab[19646]++
										return err
//line /usr/local/go/src/crypto/x509/verify.go:676
				// _ = "end of CoverTab[19646]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:677
				_go_fuzz_dep_.CoverTab[19647]++
//line /usr/local/go/src/crypto/x509/verify.go:677
				// _ = "end of CoverTab[19647]"
//line /usr/local/go/src/crypto/x509/verify.go:677
			}
//line /usr/local/go/src/crypto/x509/verify.go:677
			// _ = "end of CoverTab[19613]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:678
		// _ = "end of CoverTab[19607]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:679
		_go_fuzz_dep_.CoverTab[19648]++
//line /usr/local/go/src/crypto/x509/verify.go:679
		// _ = "end of CoverTab[19648]"
//line /usr/local/go/src/crypto/x509/verify.go:679
	}
//line /usr/local/go/src/crypto/x509/verify.go:679
	// _ = "end of CoverTab[19578]"
//line /usr/local/go/src/crypto/x509/verify.go:679
	_go_fuzz_dep_.CoverTab[19579]++

//line /usr/local/go/src/crypto/x509/verify.go:698
	if certType == intermediateCertificate && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:698
		_go_fuzz_dep_.CoverTab[19649]++
//line /usr/local/go/src/crypto/x509/verify.go:698
		return (!c.BasicConstraintsValid || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:698
			_go_fuzz_dep_.CoverTab[19650]++
//line /usr/local/go/src/crypto/x509/verify.go:698
			return !c.IsCA
//line /usr/local/go/src/crypto/x509/verify.go:698
			// _ = "end of CoverTab[19650]"
//line /usr/local/go/src/crypto/x509/verify.go:698
		}())
//line /usr/local/go/src/crypto/x509/verify.go:698
		// _ = "end of CoverTab[19649]"
//line /usr/local/go/src/crypto/x509/verify.go:698
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:698
		_go_fuzz_dep_.CoverTab[19651]++
								return CertificateInvalidError{c, NotAuthorizedToSign, ""}
//line /usr/local/go/src/crypto/x509/verify.go:699
		// _ = "end of CoverTab[19651]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:700
		_go_fuzz_dep_.CoverTab[19652]++
//line /usr/local/go/src/crypto/x509/verify.go:700
		// _ = "end of CoverTab[19652]"
//line /usr/local/go/src/crypto/x509/verify.go:700
	}
//line /usr/local/go/src/crypto/x509/verify.go:700
	// _ = "end of CoverTab[19579]"
//line /usr/local/go/src/crypto/x509/verify.go:700
	_go_fuzz_dep_.CoverTab[19580]++

							if c.BasicConstraintsValid && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:702
		_go_fuzz_dep_.CoverTab[19653]++
//line /usr/local/go/src/crypto/x509/verify.go:702
		return c.MaxPathLen >= 0
//line /usr/local/go/src/crypto/x509/verify.go:702
		// _ = "end of CoverTab[19653]"
//line /usr/local/go/src/crypto/x509/verify.go:702
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:702
		_go_fuzz_dep_.CoverTab[19654]++
								numIntermediates := len(currentChain) - 1
								if numIntermediates > c.MaxPathLen {
//line /usr/local/go/src/crypto/x509/verify.go:704
			_go_fuzz_dep_.CoverTab[19655]++
									return CertificateInvalidError{c, TooManyIntermediates, ""}
//line /usr/local/go/src/crypto/x509/verify.go:705
			// _ = "end of CoverTab[19655]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:706
			_go_fuzz_dep_.CoverTab[19656]++
//line /usr/local/go/src/crypto/x509/verify.go:706
			// _ = "end of CoverTab[19656]"
//line /usr/local/go/src/crypto/x509/verify.go:706
		}
//line /usr/local/go/src/crypto/x509/verify.go:706
		// _ = "end of CoverTab[19654]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:707
		_go_fuzz_dep_.CoverTab[19657]++
//line /usr/local/go/src/crypto/x509/verify.go:707
		// _ = "end of CoverTab[19657]"
//line /usr/local/go/src/crypto/x509/verify.go:707
	}
//line /usr/local/go/src/crypto/x509/verify.go:707
	// _ = "end of CoverTab[19580]"
//line /usr/local/go/src/crypto/x509/verify.go:707
	_go_fuzz_dep_.CoverTab[19581]++

							if !boringAllowCert(c) {
//line /usr/local/go/src/crypto/x509/verify.go:709
		_go_fuzz_dep_.CoverTab[19658]++

//line /usr/local/go/src/crypto/x509/verify.go:713
		return CertificateInvalidError{c, IncompatibleUsage, ""}
//line /usr/local/go/src/crypto/x509/verify.go:713
		// _ = "end of CoverTab[19658]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:714
		_go_fuzz_dep_.CoverTab[19659]++
//line /usr/local/go/src/crypto/x509/verify.go:714
		// _ = "end of CoverTab[19659]"
//line /usr/local/go/src/crypto/x509/verify.go:714
	}
//line /usr/local/go/src/crypto/x509/verify.go:714
	// _ = "end of CoverTab[19581]"
//line /usr/local/go/src/crypto/x509/verify.go:714
	_go_fuzz_dep_.CoverTab[19582]++

							return nil
//line /usr/local/go/src/crypto/x509/verify.go:716
	// _ = "end of CoverTab[19582]"
}

// Verify attempts to verify c by building one or more chains from c to a
//line /usr/local/go/src/crypto/x509/verify.go:719
// certificate in opts.Roots, using certificates in opts.Intermediates if
//line /usr/local/go/src/crypto/x509/verify.go:719
// needed. If successful, it returns one or more chains where the first
//line /usr/local/go/src/crypto/x509/verify.go:719
// element of the chain is c and the last element is from opts.Roots.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// If opts.Roots is nil, the platform verifier might be used, and
//line /usr/local/go/src/crypto/x509/verify.go:719
// verification details might differ from what is described below. If system
//line /usr/local/go/src/crypto/x509/verify.go:719
// roots are unavailable the returned error will be of type SystemRootsError.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// Name constraints in the intermediates will be applied to all names claimed
//line /usr/local/go/src/crypto/x509/verify.go:719
// in the chain, not just opts.DNSName. Thus it is invalid for a leaf to claim
//line /usr/local/go/src/crypto/x509/verify.go:719
// example.com if an intermediate doesn't permit it, even if example.com is not
//line /usr/local/go/src/crypto/x509/verify.go:719
// the name being validated. Note that DirectoryName constraints are not
//line /usr/local/go/src/crypto/x509/verify.go:719
// supported.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// Name constraint validation follows the rules from RFC 5280, with the
//line /usr/local/go/src/crypto/x509/verify.go:719
// addition that DNS name constraints may use the leading period format
//line /usr/local/go/src/crypto/x509/verify.go:719
// defined for emails and URIs. When a constraint has a leading period
//line /usr/local/go/src/crypto/x509/verify.go:719
// it indicates that at least one additional label must be prepended to
//line /usr/local/go/src/crypto/x509/verify.go:719
// the constrained name to be considered valid.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// Extended Key Usage values are enforced nested down a chain, so an intermediate
//line /usr/local/go/src/crypto/x509/verify.go:719
// or root that enumerates EKUs prevents a leaf from asserting an EKU not in that
//line /usr/local/go/src/crypto/x509/verify.go:719
// list. (While this is not specified, it is common practice in order to limit
//line /usr/local/go/src/crypto/x509/verify.go:719
// the types of certificates a CA can issue.)
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// Certificates that use SHA1WithRSA and ECDSAWithSHA1 signatures are not supported,
//line /usr/local/go/src/crypto/x509/verify.go:719
// and will not be used to build chains.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// Certificates other than c in the returned chains should not be modified.
//line /usr/local/go/src/crypto/x509/verify.go:719
//
//line /usr/local/go/src/crypto/x509/verify.go:719
// WARNING: this function doesn't do any revocation checking.
//line /usr/local/go/src/crypto/x509/verify.go:751
func (c *Certificate) Verify(opts VerifyOptions) (chains [][]*Certificate, err error) {
//line /usr/local/go/src/crypto/x509/verify.go:751
	_go_fuzz_dep_.CoverTab[19660]++

//line /usr/local/go/src/crypto/x509/verify.go:754
	if len(c.Raw) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:754
		_go_fuzz_dep_.CoverTab[19672]++
								return nil, errNotParsed
//line /usr/local/go/src/crypto/x509/verify.go:755
		// _ = "end of CoverTab[19672]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:756
		_go_fuzz_dep_.CoverTab[19673]++
//line /usr/local/go/src/crypto/x509/verify.go:756
		// _ = "end of CoverTab[19673]"
//line /usr/local/go/src/crypto/x509/verify.go:756
	}
//line /usr/local/go/src/crypto/x509/verify.go:756
	// _ = "end of CoverTab[19660]"
//line /usr/local/go/src/crypto/x509/verify.go:756
	_go_fuzz_dep_.CoverTab[19661]++
							for i := 0; i < opts.Intermediates.len(); i++ {
//line /usr/local/go/src/crypto/x509/verify.go:757
		_go_fuzz_dep_.CoverTab[19674]++
								c, err := opts.Intermediates.cert(i)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:759
			_go_fuzz_dep_.CoverTab[19676]++
									return nil, fmt.Errorf("crypto/x509: error fetching intermediate: %w", err)
//line /usr/local/go/src/crypto/x509/verify.go:760
			// _ = "end of CoverTab[19676]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:761
			_go_fuzz_dep_.CoverTab[19677]++
//line /usr/local/go/src/crypto/x509/verify.go:761
			// _ = "end of CoverTab[19677]"
//line /usr/local/go/src/crypto/x509/verify.go:761
		}
//line /usr/local/go/src/crypto/x509/verify.go:761
		// _ = "end of CoverTab[19674]"
//line /usr/local/go/src/crypto/x509/verify.go:761
		_go_fuzz_dep_.CoverTab[19675]++
								if len(c.Raw) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:762
			_go_fuzz_dep_.CoverTab[19678]++
									return nil, errNotParsed
//line /usr/local/go/src/crypto/x509/verify.go:763
			// _ = "end of CoverTab[19678]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:764
			_go_fuzz_dep_.CoverTab[19679]++
//line /usr/local/go/src/crypto/x509/verify.go:764
			// _ = "end of CoverTab[19679]"
//line /usr/local/go/src/crypto/x509/verify.go:764
		}
//line /usr/local/go/src/crypto/x509/verify.go:764
		// _ = "end of CoverTab[19675]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:765
	// _ = "end of CoverTab[19661]"
//line /usr/local/go/src/crypto/x509/verify.go:765
	_go_fuzz_dep_.CoverTab[19662]++

//line /usr/local/go/src/crypto/x509/verify.go:768
	if runtime.GOOS == "windows" || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:768
		_go_fuzz_dep_.CoverTab[19680]++
//line /usr/local/go/src/crypto/x509/verify.go:768
		return runtime.GOOS == "darwin"
//line /usr/local/go/src/crypto/x509/verify.go:768
		// _ = "end of CoverTab[19680]"
//line /usr/local/go/src/crypto/x509/verify.go:768
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:768
		_go_fuzz_dep_.CoverTab[19681]++
//line /usr/local/go/src/crypto/x509/verify.go:768
		return runtime.GOOS == "ios"
//line /usr/local/go/src/crypto/x509/verify.go:768
		// _ = "end of CoverTab[19681]"
//line /usr/local/go/src/crypto/x509/verify.go:768
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:768
		_go_fuzz_dep_.CoverTab[19682]++

//line /usr/local/go/src/crypto/x509/verify.go:771
		systemPool := systemRootsPool()
		if opts.Roots == nil && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:772
			_go_fuzz_dep_.CoverTab[19684]++
//line /usr/local/go/src/crypto/x509/verify.go:772
			return (systemPool == nil || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:772
				_go_fuzz_dep_.CoverTab[19685]++
//line /usr/local/go/src/crypto/x509/verify.go:772
				return systemPool.systemPool
//line /usr/local/go/src/crypto/x509/verify.go:772
				// _ = "end of CoverTab[19685]"
//line /usr/local/go/src/crypto/x509/verify.go:772
			}())
//line /usr/local/go/src/crypto/x509/verify.go:772
			// _ = "end of CoverTab[19684]"
//line /usr/local/go/src/crypto/x509/verify.go:772
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:772
			_go_fuzz_dep_.CoverTab[19686]++
									return c.systemVerify(&opts)
//line /usr/local/go/src/crypto/x509/verify.go:773
			// _ = "end of CoverTab[19686]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:774
			_go_fuzz_dep_.CoverTab[19687]++
//line /usr/local/go/src/crypto/x509/verify.go:774
			// _ = "end of CoverTab[19687]"
//line /usr/local/go/src/crypto/x509/verify.go:774
		}
//line /usr/local/go/src/crypto/x509/verify.go:774
		// _ = "end of CoverTab[19682]"
//line /usr/local/go/src/crypto/x509/verify.go:774
		_go_fuzz_dep_.CoverTab[19683]++
								if opts.Roots != nil && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:775
			_go_fuzz_dep_.CoverTab[19688]++
//line /usr/local/go/src/crypto/x509/verify.go:775
			return opts.Roots.systemPool
//line /usr/local/go/src/crypto/x509/verify.go:775
			// _ = "end of CoverTab[19688]"
//line /usr/local/go/src/crypto/x509/verify.go:775
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:775
			_go_fuzz_dep_.CoverTab[19689]++
									platformChains, err := c.systemVerify(&opts)

//line /usr/local/go/src/crypto/x509/verify.go:780
			if err == nil || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:780
				_go_fuzz_dep_.CoverTab[19690]++
//line /usr/local/go/src/crypto/x509/verify.go:780
				return opts.Roots.len() == 0
//line /usr/local/go/src/crypto/x509/verify.go:780
				// _ = "end of CoverTab[19690]"
//line /usr/local/go/src/crypto/x509/verify.go:780
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:780
				_go_fuzz_dep_.CoverTab[19691]++
										return platformChains, err
//line /usr/local/go/src/crypto/x509/verify.go:781
				// _ = "end of CoverTab[19691]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:782
				_go_fuzz_dep_.CoverTab[19692]++
//line /usr/local/go/src/crypto/x509/verify.go:782
				// _ = "end of CoverTab[19692]"
//line /usr/local/go/src/crypto/x509/verify.go:782
			}
//line /usr/local/go/src/crypto/x509/verify.go:782
			// _ = "end of CoverTab[19689]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:783
			_go_fuzz_dep_.CoverTab[19693]++
//line /usr/local/go/src/crypto/x509/verify.go:783
			// _ = "end of CoverTab[19693]"
//line /usr/local/go/src/crypto/x509/verify.go:783
		}
//line /usr/local/go/src/crypto/x509/verify.go:783
		// _ = "end of CoverTab[19683]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:784
		_go_fuzz_dep_.CoverTab[19694]++
//line /usr/local/go/src/crypto/x509/verify.go:784
		// _ = "end of CoverTab[19694]"
//line /usr/local/go/src/crypto/x509/verify.go:784
	}
//line /usr/local/go/src/crypto/x509/verify.go:784
	// _ = "end of CoverTab[19662]"
//line /usr/local/go/src/crypto/x509/verify.go:784
	_go_fuzz_dep_.CoverTab[19663]++

							if opts.Roots == nil {
//line /usr/local/go/src/crypto/x509/verify.go:786
		_go_fuzz_dep_.CoverTab[19695]++
								opts.Roots = systemRootsPool()
								if opts.Roots == nil {
//line /usr/local/go/src/crypto/x509/verify.go:788
			_go_fuzz_dep_.CoverTab[19696]++
									return nil, SystemRootsError{systemRootsErr}
//line /usr/local/go/src/crypto/x509/verify.go:789
			// _ = "end of CoverTab[19696]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:790
			_go_fuzz_dep_.CoverTab[19697]++
//line /usr/local/go/src/crypto/x509/verify.go:790
			// _ = "end of CoverTab[19697]"
//line /usr/local/go/src/crypto/x509/verify.go:790
		}
//line /usr/local/go/src/crypto/x509/verify.go:790
		// _ = "end of CoverTab[19695]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:791
		_go_fuzz_dep_.CoverTab[19698]++
//line /usr/local/go/src/crypto/x509/verify.go:791
		// _ = "end of CoverTab[19698]"
//line /usr/local/go/src/crypto/x509/verify.go:791
	}
//line /usr/local/go/src/crypto/x509/verify.go:791
	// _ = "end of CoverTab[19663]"
//line /usr/local/go/src/crypto/x509/verify.go:791
	_go_fuzz_dep_.CoverTab[19664]++

							err = c.isValid(leafCertificate, nil, &opts)
							if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:794
		_go_fuzz_dep_.CoverTab[19699]++
								return
//line /usr/local/go/src/crypto/x509/verify.go:795
		// _ = "end of CoverTab[19699]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:796
		_go_fuzz_dep_.CoverTab[19700]++
//line /usr/local/go/src/crypto/x509/verify.go:796
		// _ = "end of CoverTab[19700]"
//line /usr/local/go/src/crypto/x509/verify.go:796
	}
//line /usr/local/go/src/crypto/x509/verify.go:796
	// _ = "end of CoverTab[19664]"
//line /usr/local/go/src/crypto/x509/verify.go:796
	_go_fuzz_dep_.CoverTab[19665]++

							if len(opts.DNSName) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:798
		_go_fuzz_dep_.CoverTab[19701]++
								err = c.VerifyHostname(opts.DNSName)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:800
			_go_fuzz_dep_.CoverTab[19702]++
									return
//line /usr/local/go/src/crypto/x509/verify.go:801
			// _ = "end of CoverTab[19702]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:802
			_go_fuzz_dep_.CoverTab[19703]++
//line /usr/local/go/src/crypto/x509/verify.go:802
			// _ = "end of CoverTab[19703]"
//line /usr/local/go/src/crypto/x509/verify.go:802
		}
//line /usr/local/go/src/crypto/x509/verify.go:802
		// _ = "end of CoverTab[19701]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:803
		_go_fuzz_dep_.CoverTab[19704]++
//line /usr/local/go/src/crypto/x509/verify.go:803
		// _ = "end of CoverTab[19704]"
//line /usr/local/go/src/crypto/x509/verify.go:803
	}
//line /usr/local/go/src/crypto/x509/verify.go:803
	// _ = "end of CoverTab[19665]"
//line /usr/local/go/src/crypto/x509/verify.go:803
	_go_fuzz_dep_.CoverTab[19666]++

							var candidateChains [][]*Certificate
							if opts.Roots.contains(c) {
//line /usr/local/go/src/crypto/x509/verify.go:806
		_go_fuzz_dep_.CoverTab[19705]++
								candidateChains = [][]*Certificate{{c}}
//line /usr/local/go/src/crypto/x509/verify.go:807
		// _ = "end of CoverTab[19705]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:808
		_go_fuzz_dep_.CoverTab[19706]++
								candidateChains, err = c.buildChains([]*Certificate{c}, nil, &opts)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:810
			_go_fuzz_dep_.CoverTab[19707]++
									return nil, err
//line /usr/local/go/src/crypto/x509/verify.go:811
			// _ = "end of CoverTab[19707]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:812
			_go_fuzz_dep_.CoverTab[19708]++
//line /usr/local/go/src/crypto/x509/verify.go:812
			// _ = "end of CoverTab[19708]"
//line /usr/local/go/src/crypto/x509/verify.go:812
		}
//line /usr/local/go/src/crypto/x509/verify.go:812
		// _ = "end of CoverTab[19706]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:813
	// _ = "end of CoverTab[19666]"
//line /usr/local/go/src/crypto/x509/verify.go:813
	_go_fuzz_dep_.CoverTab[19667]++

							if len(opts.KeyUsages) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:815
		_go_fuzz_dep_.CoverTab[19709]++
								opts.KeyUsages = []ExtKeyUsage{ExtKeyUsageServerAuth}
//line /usr/local/go/src/crypto/x509/verify.go:816
		// _ = "end of CoverTab[19709]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:817
		_go_fuzz_dep_.CoverTab[19710]++
//line /usr/local/go/src/crypto/x509/verify.go:817
		// _ = "end of CoverTab[19710]"
//line /usr/local/go/src/crypto/x509/verify.go:817
	}
//line /usr/local/go/src/crypto/x509/verify.go:817
	// _ = "end of CoverTab[19667]"
//line /usr/local/go/src/crypto/x509/verify.go:817
	_go_fuzz_dep_.CoverTab[19668]++

							for _, eku := range opts.KeyUsages {
//line /usr/local/go/src/crypto/x509/verify.go:819
		_go_fuzz_dep_.CoverTab[19711]++
								if eku == ExtKeyUsageAny {
//line /usr/local/go/src/crypto/x509/verify.go:820
			_go_fuzz_dep_.CoverTab[19712]++

//line /usr/local/go/src/crypto/x509/verify.go:823
			return candidateChains, nil
//line /usr/local/go/src/crypto/x509/verify.go:823
			// _ = "end of CoverTab[19712]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:824
			_go_fuzz_dep_.CoverTab[19713]++
//line /usr/local/go/src/crypto/x509/verify.go:824
			// _ = "end of CoverTab[19713]"
//line /usr/local/go/src/crypto/x509/verify.go:824
		}
//line /usr/local/go/src/crypto/x509/verify.go:824
		// _ = "end of CoverTab[19711]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:825
	// _ = "end of CoverTab[19668]"
//line /usr/local/go/src/crypto/x509/verify.go:825
	_go_fuzz_dep_.CoverTab[19669]++

							chains = make([][]*Certificate, 0, len(candidateChains))
							for _, candidate := range candidateChains {
//line /usr/local/go/src/crypto/x509/verify.go:828
		_go_fuzz_dep_.CoverTab[19714]++
								if checkChainForKeyUsage(candidate, opts.KeyUsages) {
//line /usr/local/go/src/crypto/x509/verify.go:829
			_go_fuzz_dep_.CoverTab[19715]++
									chains = append(chains, candidate)
//line /usr/local/go/src/crypto/x509/verify.go:830
			// _ = "end of CoverTab[19715]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:831
			_go_fuzz_dep_.CoverTab[19716]++
//line /usr/local/go/src/crypto/x509/verify.go:831
			// _ = "end of CoverTab[19716]"
//line /usr/local/go/src/crypto/x509/verify.go:831
		}
//line /usr/local/go/src/crypto/x509/verify.go:831
		// _ = "end of CoverTab[19714]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:832
	// _ = "end of CoverTab[19669]"
//line /usr/local/go/src/crypto/x509/verify.go:832
	_go_fuzz_dep_.CoverTab[19670]++

							if len(chains) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:834
		_go_fuzz_dep_.CoverTab[19717]++
								return nil, CertificateInvalidError{c, IncompatibleUsage, ""}
//line /usr/local/go/src/crypto/x509/verify.go:835
		// _ = "end of CoverTab[19717]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:836
		_go_fuzz_dep_.CoverTab[19718]++
//line /usr/local/go/src/crypto/x509/verify.go:836
		// _ = "end of CoverTab[19718]"
//line /usr/local/go/src/crypto/x509/verify.go:836
	}
//line /usr/local/go/src/crypto/x509/verify.go:836
	// _ = "end of CoverTab[19670]"
//line /usr/local/go/src/crypto/x509/verify.go:836
	_go_fuzz_dep_.CoverTab[19671]++

							return chains, nil
//line /usr/local/go/src/crypto/x509/verify.go:838
	// _ = "end of CoverTab[19671]"
}

func appendToFreshChain(chain []*Certificate, cert *Certificate) []*Certificate {
//line /usr/local/go/src/crypto/x509/verify.go:841
	_go_fuzz_dep_.CoverTab[19719]++
							n := make([]*Certificate, len(chain)+1)
							copy(n, chain)
							n[len(chain)] = cert
							return n
//line /usr/local/go/src/crypto/x509/verify.go:845
	// _ = "end of CoverTab[19719]"
}

// alreadyInChain checks whether a candidate certificate is present in a chain.
//line /usr/local/go/src/crypto/x509/verify.go:848
// Rather than doing a direct byte for byte equivalency check, we check if the
//line /usr/local/go/src/crypto/x509/verify.go:848
// subject, public key, and SAN, if present, are equal. This prevents loops that
//line /usr/local/go/src/crypto/x509/verify.go:848
// are created by mutual cross-signatures, or other cross-signature bridge
//line /usr/local/go/src/crypto/x509/verify.go:848
// oddities.
//line /usr/local/go/src/crypto/x509/verify.go:853
func alreadyInChain(candidate *Certificate, chain []*Certificate) bool {
//line /usr/local/go/src/crypto/x509/verify.go:853
	_go_fuzz_dep_.CoverTab[19720]++
							type pubKeyEqual interface {
		Equal(crypto.PublicKey) bool
	}

	var candidateSAN *pkix.Extension
	for _, ext := range candidate.Extensions {
//line /usr/local/go/src/crypto/x509/verify.go:859
		_go_fuzz_dep_.CoverTab[19723]++
								if ext.Id.Equal(oidExtensionSubjectAltName) {
//line /usr/local/go/src/crypto/x509/verify.go:860
			_go_fuzz_dep_.CoverTab[19724]++
									candidateSAN = &ext
									break
//line /usr/local/go/src/crypto/x509/verify.go:862
			// _ = "end of CoverTab[19724]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:863
			_go_fuzz_dep_.CoverTab[19725]++
//line /usr/local/go/src/crypto/x509/verify.go:863
			// _ = "end of CoverTab[19725]"
//line /usr/local/go/src/crypto/x509/verify.go:863
		}
//line /usr/local/go/src/crypto/x509/verify.go:863
		// _ = "end of CoverTab[19723]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:864
	// _ = "end of CoverTab[19720]"
//line /usr/local/go/src/crypto/x509/verify.go:864
	_go_fuzz_dep_.CoverTab[19721]++

							for _, cert := range chain {
//line /usr/local/go/src/crypto/x509/verify.go:866
		_go_fuzz_dep_.CoverTab[19726]++
								if !bytes.Equal(candidate.RawSubject, cert.RawSubject) {
//line /usr/local/go/src/crypto/x509/verify.go:867
			_go_fuzz_dep_.CoverTab[19731]++
									continue
//line /usr/local/go/src/crypto/x509/verify.go:868
			// _ = "end of CoverTab[19731]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:869
			_go_fuzz_dep_.CoverTab[19732]++
//line /usr/local/go/src/crypto/x509/verify.go:869
			// _ = "end of CoverTab[19732]"
//line /usr/local/go/src/crypto/x509/verify.go:869
		}
//line /usr/local/go/src/crypto/x509/verify.go:869
		// _ = "end of CoverTab[19726]"
//line /usr/local/go/src/crypto/x509/verify.go:869
		_go_fuzz_dep_.CoverTab[19727]++
								if !candidate.PublicKey.(pubKeyEqual).Equal(cert.PublicKey) {
//line /usr/local/go/src/crypto/x509/verify.go:870
			_go_fuzz_dep_.CoverTab[19733]++
									continue
//line /usr/local/go/src/crypto/x509/verify.go:871
			// _ = "end of CoverTab[19733]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:872
			_go_fuzz_dep_.CoverTab[19734]++
//line /usr/local/go/src/crypto/x509/verify.go:872
			// _ = "end of CoverTab[19734]"
//line /usr/local/go/src/crypto/x509/verify.go:872
		}
//line /usr/local/go/src/crypto/x509/verify.go:872
		// _ = "end of CoverTab[19727]"
//line /usr/local/go/src/crypto/x509/verify.go:872
		_go_fuzz_dep_.CoverTab[19728]++
								var certSAN *pkix.Extension
								for _, ext := range cert.Extensions {
//line /usr/local/go/src/crypto/x509/verify.go:874
			_go_fuzz_dep_.CoverTab[19735]++
									if ext.Id.Equal(oidExtensionSubjectAltName) {
//line /usr/local/go/src/crypto/x509/verify.go:875
				_go_fuzz_dep_.CoverTab[19736]++
										certSAN = &ext
										break
//line /usr/local/go/src/crypto/x509/verify.go:877
				// _ = "end of CoverTab[19736]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:878
				_go_fuzz_dep_.CoverTab[19737]++
//line /usr/local/go/src/crypto/x509/verify.go:878
				// _ = "end of CoverTab[19737]"
//line /usr/local/go/src/crypto/x509/verify.go:878
			}
//line /usr/local/go/src/crypto/x509/verify.go:878
			// _ = "end of CoverTab[19735]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:879
		// _ = "end of CoverTab[19728]"
//line /usr/local/go/src/crypto/x509/verify.go:879
		_go_fuzz_dep_.CoverTab[19729]++
								if candidateSAN == nil && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:880
			_go_fuzz_dep_.CoverTab[19738]++
//line /usr/local/go/src/crypto/x509/verify.go:880
			return certSAN == nil
//line /usr/local/go/src/crypto/x509/verify.go:880
			// _ = "end of CoverTab[19738]"
//line /usr/local/go/src/crypto/x509/verify.go:880
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:880
			_go_fuzz_dep_.CoverTab[19739]++
									return true
//line /usr/local/go/src/crypto/x509/verify.go:881
			// _ = "end of CoverTab[19739]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:882
			_go_fuzz_dep_.CoverTab[19740]++
//line /usr/local/go/src/crypto/x509/verify.go:882
			if candidateSAN == nil || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:882
				_go_fuzz_dep_.CoverTab[19741]++
//line /usr/local/go/src/crypto/x509/verify.go:882
				return certSAN == nil
//line /usr/local/go/src/crypto/x509/verify.go:882
				// _ = "end of CoverTab[19741]"
//line /usr/local/go/src/crypto/x509/verify.go:882
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:882
				_go_fuzz_dep_.CoverTab[19742]++
										return false
//line /usr/local/go/src/crypto/x509/verify.go:883
				// _ = "end of CoverTab[19742]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:884
				_go_fuzz_dep_.CoverTab[19743]++
//line /usr/local/go/src/crypto/x509/verify.go:884
				// _ = "end of CoverTab[19743]"
//line /usr/local/go/src/crypto/x509/verify.go:884
			}
//line /usr/local/go/src/crypto/x509/verify.go:884
			// _ = "end of CoverTab[19740]"
//line /usr/local/go/src/crypto/x509/verify.go:884
		}
//line /usr/local/go/src/crypto/x509/verify.go:884
		// _ = "end of CoverTab[19729]"
//line /usr/local/go/src/crypto/x509/verify.go:884
		_go_fuzz_dep_.CoverTab[19730]++
								if bytes.Equal(candidateSAN.Value, certSAN.Value) {
//line /usr/local/go/src/crypto/x509/verify.go:885
			_go_fuzz_dep_.CoverTab[19744]++
									return true
//line /usr/local/go/src/crypto/x509/verify.go:886
			// _ = "end of CoverTab[19744]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:887
			_go_fuzz_dep_.CoverTab[19745]++
//line /usr/local/go/src/crypto/x509/verify.go:887
			// _ = "end of CoverTab[19745]"
//line /usr/local/go/src/crypto/x509/verify.go:887
		}
//line /usr/local/go/src/crypto/x509/verify.go:887
		// _ = "end of CoverTab[19730]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:888
	// _ = "end of CoverTab[19721]"
//line /usr/local/go/src/crypto/x509/verify.go:888
	_go_fuzz_dep_.CoverTab[19722]++
							return false
//line /usr/local/go/src/crypto/x509/verify.go:889
	// _ = "end of CoverTab[19722]"
}

// maxChainSignatureChecks is the maximum number of CheckSignatureFrom calls
//line /usr/local/go/src/crypto/x509/verify.go:892
// that an invocation of buildChains will (transitively) make. Most chains are
//line /usr/local/go/src/crypto/x509/verify.go:892
// less than 15 certificates long, so this leaves space for multiple chains and
//line /usr/local/go/src/crypto/x509/verify.go:892
// for failed checks due to different intermediates having the same Subject.
//line /usr/local/go/src/crypto/x509/verify.go:896
const maxChainSignatureChecks = 100

func (c *Certificate) buildChains(currentChain []*Certificate, sigChecks *int, opts *VerifyOptions) (chains [][]*Certificate, err error) {
//line /usr/local/go/src/crypto/x509/verify.go:898
	_go_fuzz_dep_.CoverTab[19746]++
							var (
		hintErr		error
		hintCert	*Certificate
	)

	considerCandidate := func(certType int, candidate *Certificate) {
//line /usr/local/go/src/crypto/x509/verify.go:904
		_go_fuzz_dep_.CoverTab[19752]++
								if alreadyInChain(candidate, currentChain) {
//line /usr/local/go/src/crypto/x509/verify.go:905
			_go_fuzz_dep_.CoverTab[19758]++
									return
//line /usr/local/go/src/crypto/x509/verify.go:906
			// _ = "end of CoverTab[19758]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:907
			_go_fuzz_dep_.CoverTab[19759]++
//line /usr/local/go/src/crypto/x509/verify.go:907
			// _ = "end of CoverTab[19759]"
//line /usr/local/go/src/crypto/x509/verify.go:907
		}
//line /usr/local/go/src/crypto/x509/verify.go:907
		// _ = "end of CoverTab[19752]"
//line /usr/local/go/src/crypto/x509/verify.go:907
		_go_fuzz_dep_.CoverTab[19753]++

								if sigChecks == nil {
//line /usr/local/go/src/crypto/x509/verify.go:909
			_go_fuzz_dep_.CoverTab[19760]++
									sigChecks = new(int)
//line /usr/local/go/src/crypto/x509/verify.go:910
			// _ = "end of CoverTab[19760]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:911
			_go_fuzz_dep_.CoverTab[19761]++
//line /usr/local/go/src/crypto/x509/verify.go:911
			// _ = "end of CoverTab[19761]"
//line /usr/local/go/src/crypto/x509/verify.go:911
		}
//line /usr/local/go/src/crypto/x509/verify.go:911
		// _ = "end of CoverTab[19753]"
//line /usr/local/go/src/crypto/x509/verify.go:911
		_go_fuzz_dep_.CoverTab[19754]++
								*sigChecks++
								if *sigChecks > maxChainSignatureChecks {
//line /usr/local/go/src/crypto/x509/verify.go:913
			_go_fuzz_dep_.CoverTab[19762]++
									err = errors.New("x509: signature check attempts limit reached while verifying certificate chain")
									return
//line /usr/local/go/src/crypto/x509/verify.go:915
			// _ = "end of CoverTab[19762]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:916
			_go_fuzz_dep_.CoverTab[19763]++
//line /usr/local/go/src/crypto/x509/verify.go:916
			// _ = "end of CoverTab[19763]"
//line /usr/local/go/src/crypto/x509/verify.go:916
		}
//line /usr/local/go/src/crypto/x509/verify.go:916
		// _ = "end of CoverTab[19754]"
//line /usr/local/go/src/crypto/x509/verify.go:916
		_go_fuzz_dep_.CoverTab[19755]++

								if err := c.CheckSignatureFrom(candidate); err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:918
			_go_fuzz_dep_.CoverTab[19764]++
									if hintErr == nil {
//line /usr/local/go/src/crypto/x509/verify.go:919
				_go_fuzz_dep_.CoverTab[19766]++
										hintErr = err
										hintCert = candidate
//line /usr/local/go/src/crypto/x509/verify.go:921
				// _ = "end of CoverTab[19766]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:922
				_go_fuzz_dep_.CoverTab[19767]++
//line /usr/local/go/src/crypto/x509/verify.go:922
				// _ = "end of CoverTab[19767]"
//line /usr/local/go/src/crypto/x509/verify.go:922
			}
//line /usr/local/go/src/crypto/x509/verify.go:922
			// _ = "end of CoverTab[19764]"
//line /usr/local/go/src/crypto/x509/verify.go:922
			_go_fuzz_dep_.CoverTab[19765]++
									return
//line /usr/local/go/src/crypto/x509/verify.go:923
			// _ = "end of CoverTab[19765]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:924
			_go_fuzz_dep_.CoverTab[19768]++
//line /usr/local/go/src/crypto/x509/verify.go:924
			// _ = "end of CoverTab[19768]"
//line /usr/local/go/src/crypto/x509/verify.go:924
		}
//line /usr/local/go/src/crypto/x509/verify.go:924
		// _ = "end of CoverTab[19755]"
//line /usr/local/go/src/crypto/x509/verify.go:924
		_go_fuzz_dep_.CoverTab[19756]++

								err = candidate.isValid(certType, currentChain, opts)
								if err != nil {
//line /usr/local/go/src/crypto/x509/verify.go:927
			_go_fuzz_dep_.CoverTab[19769]++
									if hintErr == nil {
//line /usr/local/go/src/crypto/x509/verify.go:928
				_go_fuzz_dep_.CoverTab[19771]++
										hintErr = err
										hintCert = candidate
//line /usr/local/go/src/crypto/x509/verify.go:930
				// _ = "end of CoverTab[19771]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:931
				_go_fuzz_dep_.CoverTab[19772]++
//line /usr/local/go/src/crypto/x509/verify.go:931
				// _ = "end of CoverTab[19772]"
//line /usr/local/go/src/crypto/x509/verify.go:931
			}
//line /usr/local/go/src/crypto/x509/verify.go:931
			// _ = "end of CoverTab[19769]"
//line /usr/local/go/src/crypto/x509/verify.go:931
			_go_fuzz_dep_.CoverTab[19770]++
									return
//line /usr/local/go/src/crypto/x509/verify.go:932
			// _ = "end of CoverTab[19770]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:933
			_go_fuzz_dep_.CoverTab[19773]++
//line /usr/local/go/src/crypto/x509/verify.go:933
			// _ = "end of CoverTab[19773]"
//line /usr/local/go/src/crypto/x509/verify.go:933
		}
//line /usr/local/go/src/crypto/x509/verify.go:933
		// _ = "end of CoverTab[19756]"
//line /usr/local/go/src/crypto/x509/verify.go:933
		_go_fuzz_dep_.CoverTab[19757]++

								switch certType {
		case rootCertificate:
//line /usr/local/go/src/crypto/x509/verify.go:936
			_go_fuzz_dep_.CoverTab[19774]++
									chains = append(chains, appendToFreshChain(currentChain, candidate))
//line /usr/local/go/src/crypto/x509/verify.go:937
			// _ = "end of CoverTab[19774]"
		case intermediateCertificate:
//line /usr/local/go/src/crypto/x509/verify.go:938
			_go_fuzz_dep_.CoverTab[19775]++
									var childChains [][]*Certificate
									childChains, err = candidate.buildChains(appendToFreshChain(currentChain, candidate), sigChecks, opts)
									chains = append(chains, childChains...)
//line /usr/local/go/src/crypto/x509/verify.go:941
			// _ = "end of CoverTab[19775]"
//line /usr/local/go/src/crypto/x509/verify.go:941
		default:
//line /usr/local/go/src/crypto/x509/verify.go:941
			_go_fuzz_dep_.CoverTab[19776]++
//line /usr/local/go/src/crypto/x509/verify.go:941
			// _ = "end of CoverTab[19776]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:942
		// _ = "end of CoverTab[19757]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:943
	// _ = "end of CoverTab[19746]"
//line /usr/local/go/src/crypto/x509/verify.go:943
	_go_fuzz_dep_.CoverTab[19747]++

							for _, root := range opts.Roots.findPotentialParents(c) {
//line /usr/local/go/src/crypto/x509/verify.go:945
		_go_fuzz_dep_.CoverTab[19777]++
								considerCandidate(rootCertificate, root)
//line /usr/local/go/src/crypto/x509/verify.go:946
		// _ = "end of CoverTab[19777]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:947
	// _ = "end of CoverTab[19747]"
//line /usr/local/go/src/crypto/x509/verify.go:947
	_go_fuzz_dep_.CoverTab[19748]++
							for _, intermediate := range opts.Intermediates.findPotentialParents(c) {
//line /usr/local/go/src/crypto/x509/verify.go:948
		_go_fuzz_dep_.CoverTab[19778]++
								considerCandidate(intermediateCertificate, intermediate)
//line /usr/local/go/src/crypto/x509/verify.go:949
		// _ = "end of CoverTab[19778]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:950
	// _ = "end of CoverTab[19748]"
//line /usr/local/go/src/crypto/x509/verify.go:950
	_go_fuzz_dep_.CoverTab[19749]++

							if len(chains) > 0 {
//line /usr/local/go/src/crypto/x509/verify.go:952
		_go_fuzz_dep_.CoverTab[19779]++
								err = nil
//line /usr/local/go/src/crypto/x509/verify.go:953
		// _ = "end of CoverTab[19779]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:954
		_go_fuzz_dep_.CoverTab[19780]++
//line /usr/local/go/src/crypto/x509/verify.go:954
		// _ = "end of CoverTab[19780]"
//line /usr/local/go/src/crypto/x509/verify.go:954
	}
//line /usr/local/go/src/crypto/x509/verify.go:954
	// _ = "end of CoverTab[19749]"
//line /usr/local/go/src/crypto/x509/verify.go:954
	_go_fuzz_dep_.CoverTab[19750]++
							if len(chains) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:955
		_go_fuzz_dep_.CoverTab[19781]++
//line /usr/local/go/src/crypto/x509/verify.go:955
		return err == nil
//line /usr/local/go/src/crypto/x509/verify.go:955
		// _ = "end of CoverTab[19781]"
//line /usr/local/go/src/crypto/x509/verify.go:955
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:955
		_go_fuzz_dep_.CoverTab[19782]++
								err = UnknownAuthorityError{c, hintErr, hintCert}
//line /usr/local/go/src/crypto/x509/verify.go:956
		// _ = "end of CoverTab[19782]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:957
		_go_fuzz_dep_.CoverTab[19783]++
//line /usr/local/go/src/crypto/x509/verify.go:957
		// _ = "end of CoverTab[19783]"
//line /usr/local/go/src/crypto/x509/verify.go:957
	}
//line /usr/local/go/src/crypto/x509/verify.go:957
	// _ = "end of CoverTab[19750]"
//line /usr/local/go/src/crypto/x509/verify.go:957
	_go_fuzz_dep_.CoverTab[19751]++

							return
//line /usr/local/go/src/crypto/x509/verify.go:959
	// _ = "end of CoverTab[19751]"
}

func validHostnamePattern(host string) bool {
//line /usr/local/go/src/crypto/x509/verify.go:962
	_go_fuzz_dep_.CoverTab[19784]++
//line /usr/local/go/src/crypto/x509/verify.go:962
	return validHostname(host, true)
//line /usr/local/go/src/crypto/x509/verify.go:962
	// _ = "end of CoverTab[19784]"
//line /usr/local/go/src/crypto/x509/verify.go:962
}
func validHostnameInput(host string) bool {
//line /usr/local/go/src/crypto/x509/verify.go:963
	_go_fuzz_dep_.CoverTab[19785]++
//line /usr/local/go/src/crypto/x509/verify.go:963
	return validHostname(host, false)
//line /usr/local/go/src/crypto/x509/verify.go:963
	// _ = "end of CoverTab[19785]"
//line /usr/local/go/src/crypto/x509/verify.go:963
}

// validHostname reports whether host is a valid hostname that can be matched or
//line /usr/local/go/src/crypto/x509/verify.go:965
// matched against according to RFC 6125 2.2, with some leniency to accommodate
//line /usr/local/go/src/crypto/x509/verify.go:965
// legacy values.
//line /usr/local/go/src/crypto/x509/verify.go:968
func validHostname(host string, isPattern bool) bool {
//line /usr/local/go/src/crypto/x509/verify.go:968
	_go_fuzz_dep_.CoverTab[19786]++
							if !isPattern {
//line /usr/local/go/src/crypto/x509/verify.go:969
		_go_fuzz_dep_.CoverTab[19790]++
								host = strings.TrimSuffix(host, ".")
//line /usr/local/go/src/crypto/x509/verify.go:970
		// _ = "end of CoverTab[19790]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:971
		_go_fuzz_dep_.CoverTab[19791]++
//line /usr/local/go/src/crypto/x509/verify.go:971
		// _ = "end of CoverTab[19791]"
//line /usr/local/go/src/crypto/x509/verify.go:971
	}
//line /usr/local/go/src/crypto/x509/verify.go:971
	// _ = "end of CoverTab[19786]"
//line /usr/local/go/src/crypto/x509/verify.go:971
	_go_fuzz_dep_.CoverTab[19787]++
							if len(host) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:972
		_go_fuzz_dep_.CoverTab[19792]++
								return false
//line /usr/local/go/src/crypto/x509/verify.go:973
		// _ = "end of CoverTab[19792]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:974
		_go_fuzz_dep_.CoverTab[19793]++
//line /usr/local/go/src/crypto/x509/verify.go:974
		// _ = "end of CoverTab[19793]"
//line /usr/local/go/src/crypto/x509/verify.go:974
	}
//line /usr/local/go/src/crypto/x509/verify.go:974
	// _ = "end of CoverTab[19787]"
//line /usr/local/go/src/crypto/x509/verify.go:974
	_go_fuzz_dep_.CoverTab[19788]++

							for i, part := range strings.Split(host, ".") {
//line /usr/local/go/src/crypto/x509/verify.go:976
		_go_fuzz_dep_.CoverTab[19794]++
								if part == "" {
//line /usr/local/go/src/crypto/x509/verify.go:977
			_go_fuzz_dep_.CoverTab[19797]++

									return false
//line /usr/local/go/src/crypto/x509/verify.go:979
			// _ = "end of CoverTab[19797]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:980
			_go_fuzz_dep_.CoverTab[19798]++
//line /usr/local/go/src/crypto/x509/verify.go:980
			// _ = "end of CoverTab[19798]"
//line /usr/local/go/src/crypto/x509/verify.go:980
		}
//line /usr/local/go/src/crypto/x509/verify.go:980
		// _ = "end of CoverTab[19794]"
//line /usr/local/go/src/crypto/x509/verify.go:980
		_go_fuzz_dep_.CoverTab[19795]++
								if isPattern && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:981
			_go_fuzz_dep_.CoverTab[19799]++
//line /usr/local/go/src/crypto/x509/verify.go:981
			return i == 0
//line /usr/local/go/src/crypto/x509/verify.go:981
			// _ = "end of CoverTab[19799]"
//line /usr/local/go/src/crypto/x509/verify.go:981
		}() && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:981
			_go_fuzz_dep_.CoverTab[19800]++
//line /usr/local/go/src/crypto/x509/verify.go:981
			return part == "*"
//line /usr/local/go/src/crypto/x509/verify.go:981
			// _ = "end of CoverTab[19800]"
//line /usr/local/go/src/crypto/x509/verify.go:981
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:981
			_go_fuzz_dep_.CoverTab[19801]++

//line /usr/local/go/src/crypto/x509/verify.go:985
			continue
//line /usr/local/go/src/crypto/x509/verify.go:985
			// _ = "end of CoverTab[19801]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:986
			_go_fuzz_dep_.CoverTab[19802]++
//line /usr/local/go/src/crypto/x509/verify.go:986
			// _ = "end of CoverTab[19802]"
//line /usr/local/go/src/crypto/x509/verify.go:986
		}
//line /usr/local/go/src/crypto/x509/verify.go:986
		// _ = "end of CoverTab[19795]"
//line /usr/local/go/src/crypto/x509/verify.go:986
		_go_fuzz_dep_.CoverTab[19796]++
								for j, c := range part {
//line /usr/local/go/src/crypto/x509/verify.go:987
			_go_fuzz_dep_.CoverTab[19803]++
									if 'a' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:988
				_go_fuzz_dep_.CoverTab[19809]++
//line /usr/local/go/src/crypto/x509/verify.go:988
				return c <= 'z'
//line /usr/local/go/src/crypto/x509/verify.go:988
				// _ = "end of CoverTab[19809]"
//line /usr/local/go/src/crypto/x509/verify.go:988
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:988
				_go_fuzz_dep_.CoverTab[19810]++
										continue
//line /usr/local/go/src/crypto/x509/verify.go:989
				// _ = "end of CoverTab[19810]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:990
				_go_fuzz_dep_.CoverTab[19811]++
//line /usr/local/go/src/crypto/x509/verify.go:990
				// _ = "end of CoverTab[19811]"
//line /usr/local/go/src/crypto/x509/verify.go:990
			}
//line /usr/local/go/src/crypto/x509/verify.go:990
			// _ = "end of CoverTab[19803]"
//line /usr/local/go/src/crypto/x509/verify.go:990
			_go_fuzz_dep_.CoverTab[19804]++
									if '0' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:991
				_go_fuzz_dep_.CoverTab[19812]++
//line /usr/local/go/src/crypto/x509/verify.go:991
				return c <= '9'
//line /usr/local/go/src/crypto/x509/verify.go:991
				// _ = "end of CoverTab[19812]"
//line /usr/local/go/src/crypto/x509/verify.go:991
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:991
				_go_fuzz_dep_.CoverTab[19813]++
										continue
//line /usr/local/go/src/crypto/x509/verify.go:992
				// _ = "end of CoverTab[19813]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:993
				_go_fuzz_dep_.CoverTab[19814]++
//line /usr/local/go/src/crypto/x509/verify.go:993
				// _ = "end of CoverTab[19814]"
//line /usr/local/go/src/crypto/x509/verify.go:993
			}
//line /usr/local/go/src/crypto/x509/verify.go:993
			// _ = "end of CoverTab[19804]"
//line /usr/local/go/src/crypto/x509/verify.go:993
			_go_fuzz_dep_.CoverTab[19805]++
									if 'A' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:994
				_go_fuzz_dep_.CoverTab[19815]++
//line /usr/local/go/src/crypto/x509/verify.go:994
				return c <= 'Z'
//line /usr/local/go/src/crypto/x509/verify.go:994
				// _ = "end of CoverTab[19815]"
//line /usr/local/go/src/crypto/x509/verify.go:994
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:994
				_go_fuzz_dep_.CoverTab[19816]++
										continue
//line /usr/local/go/src/crypto/x509/verify.go:995
				// _ = "end of CoverTab[19816]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:996
				_go_fuzz_dep_.CoverTab[19817]++
//line /usr/local/go/src/crypto/x509/verify.go:996
				// _ = "end of CoverTab[19817]"
//line /usr/local/go/src/crypto/x509/verify.go:996
			}
//line /usr/local/go/src/crypto/x509/verify.go:996
			// _ = "end of CoverTab[19805]"
//line /usr/local/go/src/crypto/x509/verify.go:996
			_go_fuzz_dep_.CoverTab[19806]++
									if c == '-' && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:997
				_go_fuzz_dep_.CoverTab[19818]++
//line /usr/local/go/src/crypto/x509/verify.go:997
				return j != 0
//line /usr/local/go/src/crypto/x509/verify.go:997
				// _ = "end of CoverTab[19818]"
//line /usr/local/go/src/crypto/x509/verify.go:997
			}() {
//line /usr/local/go/src/crypto/x509/verify.go:997
				_go_fuzz_dep_.CoverTab[19819]++
										continue
//line /usr/local/go/src/crypto/x509/verify.go:998
				// _ = "end of CoverTab[19819]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:999
				_go_fuzz_dep_.CoverTab[19820]++
//line /usr/local/go/src/crypto/x509/verify.go:999
				// _ = "end of CoverTab[19820]"
//line /usr/local/go/src/crypto/x509/verify.go:999
			}
//line /usr/local/go/src/crypto/x509/verify.go:999
			// _ = "end of CoverTab[19806]"
//line /usr/local/go/src/crypto/x509/verify.go:999
			_go_fuzz_dep_.CoverTab[19807]++
									if c == '_' {
//line /usr/local/go/src/crypto/x509/verify.go:1000
				_go_fuzz_dep_.CoverTab[19821]++

//line /usr/local/go/src/crypto/x509/verify.go:1003
				continue
//line /usr/local/go/src/crypto/x509/verify.go:1003
				// _ = "end of CoverTab[19821]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1004
				_go_fuzz_dep_.CoverTab[19822]++
//line /usr/local/go/src/crypto/x509/verify.go:1004
				// _ = "end of CoverTab[19822]"
//line /usr/local/go/src/crypto/x509/verify.go:1004
			}
//line /usr/local/go/src/crypto/x509/verify.go:1004
			// _ = "end of CoverTab[19807]"
//line /usr/local/go/src/crypto/x509/verify.go:1004
			_go_fuzz_dep_.CoverTab[19808]++
									return false
//line /usr/local/go/src/crypto/x509/verify.go:1005
			// _ = "end of CoverTab[19808]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:1006
		// _ = "end of CoverTab[19796]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1007
	// _ = "end of CoverTab[19788]"
//line /usr/local/go/src/crypto/x509/verify.go:1007
	_go_fuzz_dep_.CoverTab[19789]++

							return true
//line /usr/local/go/src/crypto/x509/verify.go:1009
	// _ = "end of CoverTab[19789]"
}

func matchExactly(hostA, hostB string) bool {
//line /usr/local/go/src/crypto/x509/verify.go:1012
	_go_fuzz_dep_.CoverTab[19823]++
							if hostA == "" || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1013
		_go_fuzz_dep_.CoverTab[19825]++
//line /usr/local/go/src/crypto/x509/verify.go:1013
		return hostA == "."
//line /usr/local/go/src/crypto/x509/verify.go:1013
		// _ = "end of CoverTab[19825]"
//line /usr/local/go/src/crypto/x509/verify.go:1013
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1013
		_go_fuzz_dep_.CoverTab[19826]++
//line /usr/local/go/src/crypto/x509/verify.go:1013
		return hostB == ""
//line /usr/local/go/src/crypto/x509/verify.go:1013
		// _ = "end of CoverTab[19826]"
//line /usr/local/go/src/crypto/x509/verify.go:1013
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1013
		_go_fuzz_dep_.CoverTab[19827]++
//line /usr/local/go/src/crypto/x509/verify.go:1013
		return hostB == "."
//line /usr/local/go/src/crypto/x509/verify.go:1013
		// _ = "end of CoverTab[19827]"
//line /usr/local/go/src/crypto/x509/verify.go:1013
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:1013
		_go_fuzz_dep_.CoverTab[19828]++
								return false
//line /usr/local/go/src/crypto/x509/verify.go:1014
		// _ = "end of CoverTab[19828]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1015
		_go_fuzz_dep_.CoverTab[19829]++
//line /usr/local/go/src/crypto/x509/verify.go:1015
		// _ = "end of CoverTab[19829]"
//line /usr/local/go/src/crypto/x509/verify.go:1015
	}
//line /usr/local/go/src/crypto/x509/verify.go:1015
	// _ = "end of CoverTab[19823]"
//line /usr/local/go/src/crypto/x509/verify.go:1015
	_go_fuzz_dep_.CoverTab[19824]++
							return toLowerCaseASCII(hostA) == toLowerCaseASCII(hostB)
//line /usr/local/go/src/crypto/x509/verify.go:1016
	// _ = "end of CoverTab[19824]"
}

func matchHostnames(pattern, host string) bool {
//line /usr/local/go/src/crypto/x509/verify.go:1019
	_go_fuzz_dep_.CoverTab[19830]++
							pattern = toLowerCaseASCII(pattern)
							host = toLowerCaseASCII(strings.TrimSuffix(host, "."))

							if len(pattern) == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1023
		_go_fuzz_dep_.CoverTab[19834]++
//line /usr/local/go/src/crypto/x509/verify.go:1023
		return len(host) == 0
//line /usr/local/go/src/crypto/x509/verify.go:1023
		// _ = "end of CoverTab[19834]"
//line /usr/local/go/src/crypto/x509/verify.go:1023
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:1023
		_go_fuzz_dep_.CoverTab[19835]++
								return false
//line /usr/local/go/src/crypto/x509/verify.go:1024
		// _ = "end of CoverTab[19835]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1025
		_go_fuzz_dep_.CoverTab[19836]++
//line /usr/local/go/src/crypto/x509/verify.go:1025
		// _ = "end of CoverTab[19836]"
//line /usr/local/go/src/crypto/x509/verify.go:1025
	}
//line /usr/local/go/src/crypto/x509/verify.go:1025
	// _ = "end of CoverTab[19830]"
//line /usr/local/go/src/crypto/x509/verify.go:1025
	_go_fuzz_dep_.CoverTab[19831]++

							patternParts := strings.Split(pattern, ".")
							hostParts := strings.Split(host, ".")

							if len(patternParts) != len(hostParts) {
//line /usr/local/go/src/crypto/x509/verify.go:1030
		_go_fuzz_dep_.CoverTab[19837]++
								return false
//line /usr/local/go/src/crypto/x509/verify.go:1031
		// _ = "end of CoverTab[19837]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1032
		_go_fuzz_dep_.CoverTab[19838]++
//line /usr/local/go/src/crypto/x509/verify.go:1032
		// _ = "end of CoverTab[19838]"
//line /usr/local/go/src/crypto/x509/verify.go:1032
	}
//line /usr/local/go/src/crypto/x509/verify.go:1032
	// _ = "end of CoverTab[19831]"
//line /usr/local/go/src/crypto/x509/verify.go:1032
	_go_fuzz_dep_.CoverTab[19832]++

							for i, patternPart := range patternParts {
//line /usr/local/go/src/crypto/x509/verify.go:1034
		_go_fuzz_dep_.CoverTab[19839]++
								if i == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1035
			_go_fuzz_dep_.CoverTab[19841]++
//line /usr/local/go/src/crypto/x509/verify.go:1035
			return patternPart == "*"
//line /usr/local/go/src/crypto/x509/verify.go:1035
			// _ = "end of CoverTab[19841]"
//line /usr/local/go/src/crypto/x509/verify.go:1035
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:1035
			_go_fuzz_dep_.CoverTab[19842]++
									continue
//line /usr/local/go/src/crypto/x509/verify.go:1036
			// _ = "end of CoverTab[19842]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1037
			_go_fuzz_dep_.CoverTab[19843]++
//line /usr/local/go/src/crypto/x509/verify.go:1037
			// _ = "end of CoverTab[19843]"
//line /usr/local/go/src/crypto/x509/verify.go:1037
		}
//line /usr/local/go/src/crypto/x509/verify.go:1037
		// _ = "end of CoverTab[19839]"
//line /usr/local/go/src/crypto/x509/verify.go:1037
		_go_fuzz_dep_.CoverTab[19840]++
								if patternPart != hostParts[i] {
//line /usr/local/go/src/crypto/x509/verify.go:1038
			_go_fuzz_dep_.CoverTab[19844]++
									return false
//line /usr/local/go/src/crypto/x509/verify.go:1039
			// _ = "end of CoverTab[19844]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1040
			_go_fuzz_dep_.CoverTab[19845]++
//line /usr/local/go/src/crypto/x509/verify.go:1040
			// _ = "end of CoverTab[19845]"
//line /usr/local/go/src/crypto/x509/verify.go:1040
		}
//line /usr/local/go/src/crypto/x509/verify.go:1040
		// _ = "end of CoverTab[19840]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1041
	// _ = "end of CoverTab[19832]"
//line /usr/local/go/src/crypto/x509/verify.go:1041
	_go_fuzz_dep_.CoverTab[19833]++

							return true
//line /usr/local/go/src/crypto/x509/verify.go:1043
	// _ = "end of CoverTab[19833]"
}

// toLowerCaseASCII returns a lower-case version of in. See RFC 6125 6.4.1. We use
//line /usr/local/go/src/crypto/x509/verify.go:1046
// an explicitly ASCII function to avoid any sharp corners resulting from
//line /usr/local/go/src/crypto/x509/verify.go:1046
// performing Unicode operations on DNS labels.
//line /usr/local/go/src/crypto/x509/verify.go:1049
func toLowerCaseASCII(in string) string {
//line /usr/local/go/src/crypto/x509/verify.go:1049
	_go_fuzz_dep_.CoverTab[19846]++

							isAlreadyLowerCase := true
							for _, c := range in {
//line /usr/local/go/src/crypto/x509/verify.go:1052
		_go_fuzz_dep_.CoverTab[19850]++
								if c == utf8.RuneError {
//line /usr/local/go/src/crypto/x509/verify.go:1053
			_go_fuzz_dep_.CoverTab[19852]++

//line /usr/local/go/src/crypto/x509/verify.go:1056
			isAlreadyLowerCase = false
									break
//line /usr/local/go/src/crypto/x509/verify.go:1057
			// _ = "end of CoverTab[19852]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1058
			_go_fuzz_dep_.CoverTab[19853]++
//line /usr/local/go/src/crypto/x509/verify.go:1058
			// _ = "end of CoverTab[19853]"
//line /usr/local/go/src/crypto/x509/verify.go:1058
		}
//line /usr/local/go/src/crypto/x509/verify.go:1058
		// _ = "end of CoverTab[19850]"
//line /usr/local/go/src/crypto/x509/verify.go:1058
		_go_fuzz_dep_.CoverTab[19851]++
								if 'A' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1059
			_go_fuzz_dep_.CoverTab[19854]++
//line /usr/local/go/src/crypto/x509/verify.go:1059
			return c <= 'Z'
//line /usr/local/go/src/crypto/x509/verify.go:1059
			// _ = "end of CoverTab[19854]"
//line /usr/local/go/src/crypto/x509/verify.go:1059
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:1059
			_go_fuzz_dep_.CoverTab[19855]++
									isAlreadyLowerCase = false
									break
//line /usr/local/go/src/crypto/x509/verify.go:1061
			// _ = "end of CoverTab[19855]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1062
			_go_fuzz_dep_.CoverTab[19856]++
//line /usr/local/go/src/crypto/x509/verify.go:1062
			// _ = "end of CoverTab[19856]"
//line /usr/local/go/src/crypto/x509/verify.go:1062
		}
//line /usr/local/go/src/crypto/x509/verify.go:1062
		// _ = "end of CoverTab[19851]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1063
	// _ = "end of CoverTab[19846]"
//line /usr/local/go/src/crypto/x509/verify.go:1063
	_go_fuzz_dep_.CoverTab[19847]++

							if isAlreadyLowerCase {
//line /usr/local/go/src/crypto/x509/verify.go:1065
		_go_fuzz_dep_.CoverTab[19857]++
								return in
//line /usr/local/go/src/crypto/x509/verify.go:1066
		// _ = "end of CoverTab[19857]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1067
		_go_fuzz_dep_.CoverTab[19858]++
//line /usr/local/go/src/crypto/x509/verify.go:1067
		// _ = "end of CoverTab[19858]"
//line /usr/local/go/src/crypto/x509/verify.go:1067
	}
//line /usr/local/go/src/crypto/x509/verify.go:1067
	// _ = "end of CoverTab[19847]"
//line /usr/local/go/src/crypto/x509/verify.go:1067
	_go_fuzz_dep_.CoverTab[19848]++

							out := []byte(in)
							for i, c := range out {
//line /usr/local/go/src/crypto/x509/verify.go:1070
		_go_fuzz_dep_.CoverTab[19859]++
								if 'A' <= c && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1071
			_go_fuzz_dep_.CoverTab[19860]++
//line /usr/local/go/src/crypto/x509/verify.go:1071
			return c <= 'Z'
//line /usr/local/go/src/crypto/x509/verify.go:1071
			// _ = "end of CoverTab[19860]"
//line /usr/local/go/src/crypto/x509/verify.go:1071
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:1071
			_go_fuzz_dep_.CoverTab[19861]++
									out[i] += 'a' - 'A'
//line /usr/local/go/src/crypto/x509/verify.go:1072
			// _ = "end of CoverTab[19861]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1073
			_go_fuzz_dep_.CoverTab[19862]++
//line /usr/local/go/src/crypto/x509/verify.go:1073
			// _ = "end of CoverTab[19862]"
//line /usr/local/go/src/crypto/x509/verify.go:1073
		}
//line /usr/local/go/src/crypto/x509/verify.go:1073
		// _ = "end of CoverTab[19859]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1074
	// _ = "end of CoverTab[19848]"
//line /usr/local/go/src/crypto/x509/verify.go:1074
	_go_fuzz_dep_.CoverTab[19849]++
							return string(out)
//line /usr/local/go/src/crypto/x509/verify.go:1075
	// _ = "end of CoverTab[19849]"
}

// VerifyHostname returns nil if c is a valid certificate for the named host.
//line /usr/local/go/src/crypto/x509/verify.go:1078
// Otherwise it returns an error describing the mismatch.
//line /usr/local/go/src/crypto/x509/verify.go:1078
//
//line /usr/local/go/src/crypto/x509/verify.go:1078
// IP addresses can be optionally enclosed in square brackets and are checked
//line /usr/local/go/src/crypto/x509/verify.go:1078
// against the IPAddresses field. Other names are checked case insensitively
//line /usr/local/go/src/crypto/x509/verify.go:1078
// against the DNSNames field. If the names are valid hostnames, the certificate
//line /usr/local/go/src/crypto/x509/verify.go:1078
// fields can have a wildcard as the left-most label.
//line /usr/local/go/src/crypto/x509/verify.go:1078
//
//line /usr/local/go/src/crypto/x509/verify.go:1078
// Note that the legacy Common Name field is ignored.
//line /usr/local/go/src/crypto/x509/verify.go:1087
func (c *Certificate) VerifyHostname(h string) error {
//line /usr/local/go/src/crypto/x509/verify.go:1087
	_go_fuzz_dep_.CoverTab[19863]++

							candidateIP := h
							if len(h) >= 3 && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1090
		_go_fuzz_dep_.CoverTab[19867]++
//line /usr/local/go/src/crypto/x509/verify.go:1090
		return h[0] == '['
//line /usr/local/go/src/crypto/x509/verify.go:1090
		// _ = "end of CoverTab[19867]"
//line /usr/local/go/src/crypto/x509/verify.go:1090
	}() && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1090
		_go_fuzz_dep_.CoverTab[19868]++
//line /usr/local/go/src/crypto/x509/verify.go:1090
		return h[len(h)-1] == ']'
//line /usr/local/go/src/crypto/x509/verify.go:1090
		// _ = "end of CoverTab[19868]"
//line /usr/local/go/src/crypto/x509/verify.go:1090
	}() {
//line /usr/local/go/src/crypto/x509/verify.go:1090
		_go_fuzz_dep_.CoverTab[19869]++
								candidateIP = h[1 : len(h)-1]
//line /usr/local/go/src/crypto/x509/verify.go:1091
		// _ = "end of CoverTab[19869]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1092
		_go_fuzz_dep_.CoverTab[19870]++
//line /usr/local/go/src/crypto/x509/verify.go:1092
		// _ = "end of CoverTab[19870]"
//line /usr/local/go/src/crypto/x509/verify.go:1092
	}
//line /usr/local/go/src/crypto/x509/verify.go:1092
	// _ = "end of CoverTab[19863]"
//line /usr/local/go/src/crypto/x509/verify.go:1092
	_go_fuzz_dep_.CoverTab[19864]++
							if ip := net.ParseIP(candidateIP); ip != nil {
//line /usr/local/go/src/crypto/x509/verify.go:1093
		_go_fuzz_dep_.CoverTab[19871]++

//line /usr/local/go/src/crypto/x509/verify.go:1096
		for _, candidate := range c.IPAddresses {
//line /usr/local/go/src/crypto/x509/verify.go:1096
			_go_fuzz_dep_.CoverTab[19873]++
									if ip.Equal(candidate) {
//line /usr/local/go/src/crypto/x509/verify.go:1097
				_go_fuzz_dep_.CoverTab[19874]++
										return nil
//line /usr/local/go/src/crypto/x509/verify.go:1098
				// _ = "end of CoverTab[19874]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1099
				_go_fuzz_dep_.CoverTab[19875]++
//line /usr/local/go/src/crypto/x509/verify.go:1099
				// _ = "end of CoverTab[19875]"
//line /usr/local/go/src/crypto/x509/verify.go:1099
			}
//line /usr/local/go/src/crypto/x509/verify.go:1099
			// _ = "end of CoverTab[19873]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:1100
		// _ = "end of CoverTab[19871]"
//line /usr/local/go/src/crypto/x509/verify.go:1100
		_go_fuzz_dep_.CoverTab[19872]++
								return HostnameError{c, candidateIP}
//line /usr/local/go/src/crypto/x509/verify.go:1101
		// _ = "end of CoverTab[19872]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1102
		_go_fuzz_dep_.CoverTab[19876]++
//line /usr/local/go/src/crypto/x509/verify.go:1102
		// _ = "end of CoverTab[19876]"
//line /usr/local/go/src/crypto/x509/verify.go:1102
	}
//line /usr/local/go/src/crypto/x509/verify.go:1102
	// _ = "end of CoverTab[19864]"
//line /usr/local/go/src/crypto/x509/verify.go:1102
	_go_fuzz_dep_.CoverTab[19865]++

							candidateName := toLowerCaseASCII(h)
							validCandidateName := validHostnameInput(candidateName)

							for _, match := range c.DNSNames {
//line /usr/local/go/src/crypto/x509/verify.go:1107
		_go_fuzz_dep_.CoverTab[19877]++

//line /usr/local/go/src/crypto/x509/verify.go:1113
		if validCandidateName && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1113
			_go_fuzz_dep_.CoverTab[19878]++
//line /usr/local/go/src/crypto/x509/verify.go:1113
			return validHostnamePattern(match)
//line /usr/local/go/src/crypto/x509/verify.go:1113
			// _ = "end of CoverTab[19878]"
//line /usr/local/go/src/crypto/x509/verify.go:1113
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:1113
			_go_fuzz_dep_.CoverTab[19879]++
									if matchHostnames(match, candidateName) {
//line /usr/local/go/src/crypto/x509/verify.go:1114
				_go_fuzz_dep_.CoverTab[19880]++
										return nil
//line /usr/local/go/src/crypto/x509/verify.go:1115
				// _ = "end of CoverTab[19880]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1116
				_go_fuzz_dep_.CoverTab[19881]++
//line /usr/local/go/src/crypto/x509/verify.go:1116
				// _ = "end of CoverTab[19881]"
//line /usr/local/go/src/crypto/x509/verify.go:1116
			}
//line /usr/local/go/src/crypto/x509/verify.go:1116
			// _ = "end of CoverTab[19879]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1117
			_go_fuzz_dep_.CoverTab[19882]++
									if matchExactly(match, candidateName) {
//line /usr/local/go/src/crypto/x509/verify.go:1118
				_go_fuzz_dep_.CoverTab[19883]++
										return nil
//line /usr/local/go/src/crypto/x509/verify.go:1119
				// _ = "end of CoverTab[19883]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1120
				_go_fuzz_dep_.CoverTab[19884]++
//line /usr/local/go/src/crypto/x509/verify.go:1120
				// _ = "end of CoverTab[19884]"
//line /usr/local/go/src/crypto/x509/verify.go:1120
			}
//line /usr/local/go/src/crypto/x509/verify.go:1120
			// _ = "end of CoverTab[19882]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:1121
		// _ = "end of CoverTab[19877]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1122
	// _ = "end of CoverTab[19865]"
//line /usr/local/go/src/crypto/x509/verify.go:1122
	_go_fuzz_dep_.CoverTab[19866]++

							return HostnameError{c, h}
//line /usr/local/go/src/crypto/x509/verify.go:1124
	// _ = "end of CoverTab[19866]"
}

func checkChainForKeyUsage(chain []*Certificate, keyUsages []ExtKeyUsage) bool {
//line /usr/local/go/src/crypto/x509/verify.go:1127
	_go_fuzz_dep_.CoverTab[19885]++
							usages := make([]ExtKeyUsage, len(keyUsages))
							copy(usages, keyUsages)

							if len(chain) == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:1131
		_go_fuzz_dep_.CoverTab[19888]++
								return false
//line /usr/local/go/src/crypto/x509/verify.go:1132
		// _ = "end of CoverTab[19888]"
	} else {
//line /usr/local/go/src/crypto/x509/verify.go:1133
		_go_fuzz_dep_.CoverTab[19889]++
//line /usr/local/go/src/crypto/x509/verify.go:1133
		// _ = "end of CoverTab[19889]"
//line /usr/local/go/src/crypto/x509/verify.go:1133
	}
//line /usr/local/go/src/crypto/x509/verify.go:1133
	// _ = "end of CoverTab[19885]"
//line /usr/local/go/src/crypto/x509/verify.go:1133
	_go_fuzz_dep_.CoverTab[19886]++

							usagesRemaining := len(usages)

//line /usr/local/go/src/crypto/x509/verify.go:1141
NextCert:
	for i := len(chain) - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/x509/verify.go:1142
		_go_fuzz_dep_.CoverTab[19890]++
								cert := chain[i]
								if len(cert.ExtKeyUsage) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/verify.go:1144
			_go_fuzz_dep_.CoverTab[19893]++
//line /usr/local/go/src/crypto/x509/verify.go:1144
			return len(cert.UnknownExtKeyUsage) == 0
//line /usr/local/go/src/crypto/x509/verify.go:1144
			// _ = "end of CoverTab[19893]"
//line /usr/local/go/src/crypto/x509/verify.go:1144
		}() {
//line /usr/local/go/src/crypto/x509/verify.go:1144
			_go_fuzz_dep_.CoverTab[19894]++

									continue
//line /usr/local/go/src/crypto/x509/verify.go:1146
			// _ = "end of CoverTab[19894]"
		} else {
//line /usr/local/go/src/crypto/x509/verify.go:1147
			_go_fuzz_dep_.CoverTab[19895]++
//line /usr/local/go/src/crypto/x509/verify.go:1147
			// _ = "end of CoverTab[19895]"
//line /usr/local/go/src/crypto/x509/verify.go:1147
		}
//line /usr/local/go/src/crypto/x509/verify.go:1147
		// _ = "end of CoverTab[19890]"
//line /usr/local/go/src/crypto/x509/verify.go:1147
		_go_fuzz_dep_.CoverTab[19891]++

								for _, usage := range cert.ExtKeyUsage {
//line /usr/local/go/src/crypto/x509/verify.go:1149
			_go_fuzz_dep_.CoverTab[19896]++
									if usage == ExtKeyUsageAny {
//line /usr/local/go/src/crypto/x509/verify.go:1150
				_go_fuzz_dep_.CoverTab[19897]++

										continue NextCert
//line /usr/local/go/src/crypto/x509/verify.go:1152
				// _ = "end of CoverTab[19897]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1153
				_go_fuzz_dep_.CoverTab[19898]++
//line /usr/local/go/src/crypto/x509/verify.go:1153
				// _ = "end of CoverTab[19898]"
//line /usr/local/go/src/crypto/x509/verify.go:1153
			}
//line /usr/local/go/src/crypto/x509/verify.go:1153
			// _ = "end of CoverTab[19896]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:1154
		// _ = "end of CoverTab[19891]"
//line /usr/local/go/src/crypto/x509/verify.go:1154
		_go_fuzz_dep_.CoverTab[19892]++

								const invalidUsage ExtKeyUsage = -1

	NextRequestedUsage:
		for i, requestedUsage := range usages {
//line /usr/local/go/src/crypto/x509/verify.go:1159
			_go_fuzz_dep_.CoverTab[19899]++
									if requestedUsage == invalidUsage {
//line /usr/local/go/src/crypto/x509/verify.go:1160
				_go_fuzz_dep_.CoverTab[19902]++
										continue
//line /usr/local/go/src/crypto/x509/verify.go:1161
				// _ = "end of CoverTab[19902]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1162
				_go_fuzz_dep_.CoverTab[19903]++
//line /usr/local/go/src/crypto/x509/verify.go:1162
				// _ = "end of CoverTab[19903]"
//line /usr/local/go/src/crypto/x509/verify.go:1162
			}
//line /usr/local/go/src/crypto/x509/verify.go:1162
			// _ = "end of CoverTab[19899]"
//line /usr/local/go/src/crypto/x509/verify.go:1162
			_go_fuzz_dep_.CoverTab[19900]++

									for _, usage := range cert.ExtKeyUsage {
//line /usr/local/go/src/crypto/x509/verify.go:1164
				_go_fuzz_dep_.CoverTab[19904]++
										if requestedUsage == usage {
//line /usr/local/go/src/crypto/x509/verify.go:1165
					_go_fuzz_dep_.CoverTab[19905]++
											continue NextRequestedUsage
//line /usr/local/go/src/crypto/x509/verify.go:1166
					// _ = "end of CoverTab[19905]"
				} else {
//line /usr/local/go/src/crypto/x509/verify.go:1167
					_go_fuzz_dep_.CoverTab[19906]++
//line /usr/local/go/src/crypto/x509/verify.go:1167
					// _ = "end of CoverTab[19906]"
//line /usr/local/go/src/crypto/x509/verify.go:1167
				}
//line /usr/local/go/src/crypto/x509/verify.go:1167
				// _ = "end of CoverTab[19904]"
			}
//line /usr/local/go/src/crypto/x509/verify.go:1168
			// _ = "end of CoverTab[19900]"
//line /usr/local/go/src/crypto/x509/verify.go:1168
			_go_fuzz_dep_.CoverTab[19901]++

									usages[i] = invalidUsage
									usagesRemaining--
									if usagesRemaining == 0 {
//line /usr/local/go/src/crypto/x509/verify.go:1172
				_go_fuzz_dep_.CoverTab[19907]++
										return false
//line /usr/local/go/src/crypto/x509/verify.go:1173
				// _ = "end of CoverTab[19907]"
			} else {
//line /usr/local/go/src/crypto/x509/verify.go:1174
				_go_fuzz_dep_.CoverTab[19908]++
//line /usr/local/go/src/crypto/x509/verify.go:1174
				// _ = "end of CoverTab[19908]"
//line /usr/local/go/src/crypto/x509/verify.go:1174
			}
//line /usr/local/go/src/crypto/x509/verify.go:1174
			// _ = "end of CoverTab[19901]"
		}
//line /usr/local/go/src/crypto/x509/verify.go:1175
		// _ = "end of CoverTab[19892]"
	}
//line /usr/local/go/src/crypto/x509/verify.go:1176
	// _ = "end of CoverTab[19886]"
//line /usr/local/go/src/crypto/x509/verify.go:1176
	_go_fuzz_dep_.CoverTab[19887]++

							return true
//line /usr/local/go/src/crypto/x509/verify.go:1178
	// _ = "end of CoverTab[19887]"
}

//line /usr/local/go/src/crypto/x509/verify.go:1179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/verify.go:1179
var _ = _go_fuzz_dep_.CoverTab
