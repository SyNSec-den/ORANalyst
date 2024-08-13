// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/cert_pool.go:5
package x509

//line /usr/local/go/src/crypto/x509/cert_pool.go:5
import (
//line /usr/local/go/src/crypto/x509/cert_pool.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/cert_pool.go:5
)
//line /usr/local/go/src/crypto/x509/cert_pool.go:5
import (
//line /usr/local/go/src/crypto/x509/cert_pool.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/cert_pool.go:5
)

import (
	"bytes"
	"crypto/sha256"
	"encoding/pem"
	"sync"
)

type sum224 [sha256.Size224]byte

// CertPool is a set of certificates.
type CertPool struct {
	byName	map[string][]int	// cert.RawSubject => index into lazyCerts

	// lazyCerts contains funcs that return a certificate,
	// lazily parsing/decompressing it as needed.
	lazyCerts	[]lazyCert

	// haveSum maps from sum224(cert.Raw) to true. It's used only
	// for AddCert duplicate detection, to avoid CertPool.contains
	// calls in the AddCert path (because the contains method can
	// call getCert and otherwise negate savings from lazy getCert
	// funcs).
	haveSum	map[sum224]bool

	// systemPool indicates whether this is a special pool derived from the
	// system roots. If it includes additional roots, it requires doing two
	// verifications, one using the roots provided by the caller, and one using
	// the system platform verifier.
	systemPool	bool
}

// lazyCert is minimal metadata about a Cert and a func to retrieve it
//line /usr/local/go/src/crypto/x509/cert_pool.go:38
// in its normal expanded *Certificate form.
//line /usr/local/go/src/crypto/x509/cert_pool.go:40
type lazyCert struct {
	// rawSubject is the Certificate.RawSubject value.
	// It's the same as the CertPool.byName key, but in []byte
	// form to make CertPool.Subjects (as used by crypto/tls) do
	// fewer allocations.
	rawSubject	[]byte

	// getCert returns the certificate.
	//
	// It is not meant to do network operations or anything else
	// where a failure is likely; the func is meant to lazily
	// parse/decompress data that is already known to be good. The
	// error in the signature primarily is meant for use in the
	// case where a cert file existed on local disk when the program
	// started up is deleted later before it's read.
	getCert	func() (*Certificate, error)
}

// NewCertPool returns a new, empty CertPool.
func NewCertPool() *CertPool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:59
	_go_fuzz_dep_.CoverTab[18314]++
							return &CertPool{
		byName:		make(map[string][]int),
		haveSum:	make(map[sum224]bool),
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:63
	// _ = "end of CoverTab[18314]"
}

// len returns the number of certs in the set.
//line /usr/local/go/src/crypto/x509/cert_pool.go:66
// A nil set is a valid empty set.
//line /usr/local/go/src/crypto/x509/cert_pool.go:68
func (s *CertPool) len() int {
//line /usr/local/go/src/crypto/x509/cert_pool.go:68
	_go_fuzz_dep_.CoverTab[18315]++
							if s == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:69
		_go_fuzz_dep_.CoverTab[18317]++
								return 0
//line /usr/local/go/src/crypto/x509/cert_pool.go:70
		// _ = "end of CoverTab[18317]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:71
		_go_fuzz_dep_.CoverTab[18318]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:71
		// _ = "end of CoverTab[18318]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:71
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:71
	// _ = "end of CoverTab[18315]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:71
	_go_fuzz_dep_.CoverTab[18316]++
							return len(s.lazyCerts)
//line /usr/local/go/src/crypto/x509/cert_pool.go:72
	// _ = "end of CoverTab[18316]"
}

// cert returns cert index n in s.
func (s *CertPool) cert(n int) (*Certificate, error) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:76
	_go_fuzz_dep_.CoverTab[18319]++
							return s.lazyCerts[n].getCert()
//line /usr/local/go/src/crypto/x509/cert_pool.go:77
	// _ = "end of CoverTab[18319]"
}

// Clone returns a copy of s.
func (s *CertPool) Clone() *CertPool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:81
	_go_fuzz_dep_.CoverTab[18320]++
							p := &CertPool{
		byName:		make(map[string][]int, len(s.byName)),
		lazyCerts:	make([]lazyCert, len(s.lazyCerts)),
		haveSum:	make(map[sum224]bool, len(s.haveSum)),
		systemPool:	s.systemPool,
	}
	for k, v := range s.byName {
//line /usr/local/go/src/crypto/x509/cert_pool.go:88
		_go_fuzz_dep_.CoverTab[18323]++
								indexes := make([]int, len(v))
								copy(indexes, v)
								p.byName[k] = indexes
//line /usr/local/go/src/crypto/x509/cert_pool.go:91
		// _ = "end of CoverTab[18323]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:92
	// _ = "end of CoverTab[18320]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:92
	_go_fuzz_dep_.CoverTab[18321]++
							for k := range s.haveSum {
//line /usr/local/go/src/crypto/x509/cert_pool.go:93
		_go_fuzz_dep_.CoverTab[18324]++
								p.haveSum[k] = true
//line /usr/local/go/src/crypto/x509/cert_pool.go:94
		// _ = "end of CoverTab[18324]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:95
	// _ = "end of CoverTab[18321]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:95
	_go_fuzz_dep_.CoverTab[18322]++
							copy(p.lazyCerts, s.lazyCerts)
							return p
//line /usr/local/go/src/crypto/x509/cert_pool.go:97
	// _ = "end of CoverTab[18322]"
}

// SystemCertPool returns a copy of the system cert pool.
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// On Unix systems other than macOS the environment variables SSL_CERT_FILE and
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// SSL_CERT_DIR can be used to override the system default locations for the SSL
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// certificate file and SSL certificate files directory, respectively. The
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// latter can be a colon-separated list.
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// Any mutations to the returned pool are not written to disk and do not affect
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// any other pool returned by SystemCertPool.
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:100
// New changes in the system cert pool might not be reflected in subsequent calls.
//line /usr/local/go/src/crypto/x509/cert_pool.go:111
func SystemCertPool() (*CertPool, error) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:111
	_go_fuzz_dep_.CoverTab[18325]++
							if sysRoots := systemRootsPool(); sysRoots != nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:112
		_go_fuzz_dep_.CoverTab[18327]++
								return sysRoots.Clone(), nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:113
		// _ = "end of CoverTab[18327]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:114
		_go_fuzz_dep_.CoverTab[18328]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:114
		// _ = "end of CoverTab[18328]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:114
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:114
	// _ = "end of CoverTab[18325]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:114
	_go_fuzz_dep_.CoverTab[18326]++

							return loadSystemRoots()
//line /usr/local/go/src/crypto/x509/cert_pool.go:116
	// _ = "end of CoverTab[18326]"
}

// findPotentialParents returns the indexes of certificates in s which might
//line /usr/local/go/src/crypto/x509/cert_pool.go:119
// have signed cert.
//line /usr/local/go/src/crypto/x509/cert_pool.go:121
func (s *CertPool) findPotentialParents(cert *Certificate) []*Certificate {
//line /usr/local/go/src/crypto/x509/cert_pool.go:121
	_go_fuzz_dep_.CoverTab[18329]++
							if s == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:122
		_go_fuzz_dep_.CoverTab[18333]++
								return nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:123
		// _ = "end of CoverTab[18333]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:124
		_go_fuzz_dep_.CoverTab[18334]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:124
		// _ = "end of CoverTab[18334]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:124
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:124
	// _ = "end of CoverTab[18329]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:124
	_go_fuzz_dep_.CoverTab[18330]++

	// consider all candidates where cert.Issuer matches cert.Subject.
	// when picking possible candidates the list is built in the order
	// of match plausibility as to save cycles in buildChains:
	//   AKID and SKID match
	//   AKID present, SKID missing / AKID missing, SKID present
	//   AKID and SKID don't match
	var matchingKeyID, oneKeyID, mismatchKeyID []*Certificate
	for _, c := range s.byName[string(cert.RawIssuer)] {
//line /usr/local/go/src/crypto/x509/cert_pool.go:133
		_go_fuzz_dep_.CoverTab[18335]++
								candidate, err := s.cert(c)
								if err != nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:135
			_go_fuzz_dep_.CoverTab[18337]++
									continue
//line /usr/local/go/src/crypto/x509/cert_pool.go:136
			// _ = "end of CoverTab[18337]"
		} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:137
			_go_fuzz_dep_.CoverTab[18338]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:137
			// _ = "end of CoverTab[18338]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:137
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:137
		// _ = "end of CoverTab[18335]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:137
		_go_fuzz_dep_.CoverTab[18336]++
								kidMatch := bytes.Equal(candidate.SubjectKeyId, cert.AuthorityKeyId)
								switch {
		case kidMatch:
//line /usr/local/go/src/crypto/x509/cert_pool.go:140
			_go_fuzz_dep_.CoverTab[18339]++
									matchingKeyID = append(matchingKeyID, candidate)
//line /usr/local/go/src/crypto/x509/cert_pool.go:141
			// _ = "end of CoverTab[18339]"
		case (len(candidate.SubjectKeyId) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
			_go_fuzz_dep_.CoverTab[18342]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
			return len(cert.AuthorityKeyId) > 0
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
			// _ = "end of CoverTab[18342]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
		}()) || func() bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
			_go_fuzz_dep_.CoverTab[18343]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:142
			return (len(candidate.SubjectKeyId) > 0 && func() bool {
										_go_fuzz_dep_.CoverTab[18344]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
				return len(cert.AuthorityKeyId) == 0
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
				// _ = "end of CoverTab[18344]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
			}())
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
			// _ = "end of CoverTab[18343]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
		}():
//line /usr/local/go/src/crypto/x509/cert_pool.go:143
			_go_fuzz_dep_.CoverTab[18340]++
									oneKeyID = append(oneKeyID, candidate)
//line /usr/local/go/src/crypto/x509/cert_pool.go:144
			// _ = "end of CoverTab[18340]"
		default:
//line /usr/local/go/src/crypto/x509/cert_pool.go:145
			_go_fuzz_dep_.CoverTab[18341]++
									mismatchKeyID = append(mismatchKeyID, candidate)
//line /usr/local/go/src/crypto/x509/cert_pool.go:146
			// _ = "end of CoverTab[18341]"
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:147
		// _ = "end of CoverTab[18336]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:148
	// _ = "end of CoverTab[18330]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:148
	_go_fuzz_dep_.CoverTab[18331]++

							found := len(matchingKeyID) + len(oneKeyID) + len(mismatchKeyID)
							if found == 0 {
//line /usr/local/go/src/crypto/x509/cert_pool.go:151
		_go_fuzz_dep_.CoverTab[18345]++
								return nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:152
		// _ = "end of CoverTab[18345]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:153
		_go_fuzz_dep_.CoverTab[18346]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:153
		// _ = "end of CoverTab[18346]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:153
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:153
	// _ = "end of CoverTab[18331]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:153
	_go_fuzz_dep_.CoverTab[18332]++
							candidates := make([]*Certificate, 0, found)
							candidates = append(candidates, matchingKeyID...)
							candidates = append(candidates, oneKeyID...)
							candidates = append(candidates, mismatchKeyID...)
							return candidates
//line /usr/local/go/src/crypto/x509/cert_pool.go:158
	// _ = "end of CoverTab[18332]"
}

func (s *CertPool) contains(cert *Certificate) bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:161
	_go_fuzz_dep_.CoverTab[18347]++
							if s == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:162
		_go_fuzz_dep_.CoverTab[18349]++
								return false
//line /usr/local/go/src/crypto/x509/cert_pool.go:163
		// _ = "end of CoverTab[18349]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:164
		_go_fuzz_dep_.CoverTab[18350]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:164
		// _ = "end of CoverTab[18350]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:164
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:164
	// _ = "end of CoverTab[18347]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:164
	_go_fuzz_dep_.CoverTab[18348]++
							return s.haveSum[sha256.Sum224(cert.Raw)]
//line /usr/local/go/src/crypto/x509/cert_pool.go:165
	// _ = "end of CoverTab[18348]"
}

// AddCert adds a certificate to a pool.
func (s *CertPool) AddCert(cert *Certificate) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:169
	_go_fuzz_dep_.CoverTab[18351]++
							if cert == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:170
		_go_fuzz_dep_.CoverTab[18353]++
								panic("adding nil Certificate to CertPool")
//line /usr/local/go/src/crypto/x509/cert_pool.go:171
		// _ = "end of CoverTab[18353]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:172
		_go_fuzz_dep_.CoverTab[18354]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:172
		// _ = "end of CoverTab[18354]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:172
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:172
	// _ = "end of CoverTab[18351]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:172
	_go_fuzz_dep_.CoverTab[18352]++
							s.addCertFunc(sha256.Sum224(cert.Raw), string(cert.RawSubject), func() (*Certificate, error) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:173
		_go_fuzz_dep_.CoverTab[18355]++
								return cert, nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:174
		// _ = "end of CoverTab[18355]"
	})
//line /usr/local/go/src/crypto/x509/cert_pool.go:175
	// _ = "end of CoverTab[18352]"
}

// addCertFunc adds metadata about a certificate to a pool, along with
//line /usr/local/go/src/crypto/x509/cert_pool.go:178
// a func to fetch that certificate later when needed.
//line /usr/local/go/src/crypto/x509/cert_pool.go:178
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:178
// The rawSubject is Certificate.RawSubject and must be non-empty.
//line /usr/local/go/src/crypto/x509/cert_pool.go:178
// The getCert func may be called 0 or more times.
//line /usr/local/go/src/crypto/x509/cert_pool.go:183
func (s *CertPool) addCertFunc(rawSum224 sum224, rawSubject string, getCert func() (*Certificate, error)) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:183
	_go_fuzz_dep_.CoverTab[18356]++
							if getCert == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:184
		_go_fuzz_dep_.CoverTab[18359]++
								panic("getCert can't be nil")
//line /usr/local/go/src/crypto/x509/cert_pool.go:185
		// _ = "end of CoverTab[18359]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:186
		_go_fuzz_dep_.CoverTab[18360]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:186
		// _ = "end of CoverTab[18360]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:186
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:186
	// _ = "end of CoverTab[18356]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:186
	_go_fuzz_dep_.CoverTab[18357]++

//line /usr/local/go/src/crypto/x509/cert_pool.go:189
	if s.haveSum[rawSum224] {
//line /usr/local/go/src/crypto/x509/cert_pool.go:189
		_go_fuzz_dep_.CoverTab[18361]++
								return
//line /usr/local/go/src/crypto/x509/cert_pool.go:190
		// _ = "end of CoverTab[18361]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:191
		_go_fuzz_dep_.CoverTab[18362]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:191
		// _ = "end of CoverTab[18362]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:191
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:191
	// _ = "end of CoverTab[18357]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:191
	_go_fuzz_dep_.CoverTab[18358]++

							s.haveSum[rawSum224] = true
							s.lazyCerts = append(s.lazyCerts, lazyCert{
		rawSubject:	[]byte(rawSubject),
		getCert:	getCert,
	})
							s.byName[rawSubject] = append(s.byName[rawSubject], len(s.lazyCerts)-1)
//line /usr/local/go/src/crypto/x509/cert_pool.go:198
	// _ = "end of CoverTab[18358]"
}

// AppendCertsFromPEM attempts to parse a series of PEM encoded certificates.
//line /usr/local/go/src/crypto/x509/cert_pool.go:201
// It appends any certificates found to s and reports whether any certificates
//line /usr/local/go/src/crypto/x509/cert_pool.go:201
// were successfully parsed.
//line /usr/local/go/src/crypto/x509/cert_pool.go:201
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:201
// On many Linux systems, /etc/ssl/cert.pem will contain the system wide set
//line /usr/local/go/src/crypto/x509/cert_pool.go:201
// of root CAs in a format suitable for this function.
//line /usr/local/go/src/crypto/x509/cert_pool.go:207
func (s *CertPool) AppendCertsFromPEM(pemCerts []byte) (ok bool) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:207
	_go_fuzz_dep_.CoverTab[18363]++
							for len(pemCerts) > 0 {
//line /usr/local/go/src/crypto/x509/cert_pool.go:208
		_go_fuzz_dep_.CoverTab[18365]++
								var block *pem.Block
								block, pemCerts = pem.Decode(pemCerts)
								if block == nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:211
			_go_fuzz_dep_.CoverTab[18370]++
									break
//line /usr/local/go/src/crypto/x509/cert_pool.go:212
			// _ = "end of CoverTab[18370]"
		} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:213
			_go_fuzz_dep_.CoverTab[18371]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:213
			// _ = "end of CoverTab[18371]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:213
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:213
		// _ = "end of CoverTab[18365]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:213
		_go_fuzz_dep_.CoverTab[18366]++
								if block.Type != "CERTIFICATE" || func() bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:214
			_go_fuzz_dep_.CoverTab[18372]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:214
			return len(block.Headers) != 0
//line /usr/local/go/src/crypto/x509/cert_pool.go:214
			// _ = "end of CoverTab[18372]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:214
		}() {
//line /usr/local/go/src/crypto/x509/cert_pool.go:214
			_go_fuzz_dep_.CoverTab[18373]++
									continue
//line /usr/local/go/src/crypto/x509/cert_pool.go:215
			// _ = "end of CoverTab[18373]"
		} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:216
			_go_fuzz_dep_.CoverTab[18374]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:216
			// _ = "end of CoverTab[18374]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:216
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:216
		// _ = "end of CoverTab[18366]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:216
		_go_fuzz_dep_.CoverTab[18367]++

								certBytes := block.Bytes
								cert, err := ParseCertificate(certBytes)
								if err != nil {
//line /usr/local/go/src/crypto/x509/cert_pool.go:220
			_go_fuzz_dep_.CoverTab[18375]++
									continue
//line /usr/local/go/src/crypto/x509/cert_pool.go:221
			// _ = "end of CoverTab[18375]"
		} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:222
			_go_fuzz_dep_.CoverTab[18376]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:222
			// _ = "end of CoverTab[18376]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:222
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:222
		// _ = "end of CoverTab[18367]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:222
		_go_fuzz_dep_.CoverTab[18368]++
								var lazyCert struct {
			sync.Once
			v	*Certificate
		}
		s.addCertFunc(sha256.Sum224(cert.Raw), string(cert.RawSubject), func() (*Certificate, error) {
//line /usr/local/go/src/crypto/x509/cert_pool.go:227
			_go_fuzz_dep_.CoverTab[18377]++
									lazyCert.Do(func() {
//line /usr/local/go/src/crypto/x509/cert_pool.go:228
				_go_fuzz_dep_.CoverTab[18379]++

										lazyCert.v, _ = ParseCertificate(certBytes)
										certBytes = nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:231
				// _ = "end of CoverTab[18379]"
			})
//line /usr/local/go/src/crypto/x509/cert_pool.go:232
			// _ = "end of CoverTab[18377]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:232
			_go_fuzz_dep_.CoverTab[18378]++
									return lazyCert.v, nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:233
			// _ = "end of CoverTab[18378]"
		})
//line /usr/local/go/src/crypto/x509/cert_pool.go:234
		// _ = "end of CoverTab[18368]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:234
		_go_fuzz_dep_.CoverTab[18369]++
								ok = true
//line /usr/local/go/src/crypto/x509/cert_pool.go:235
		// _ = "end of CoverTab[18369]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:236
	// _ = "end of CoverTab[18363]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:236
	_go_fuzz_dep_.CoverTab[18364]++

							return ok
//line /usr/local/go/src/crypto/x509/cert_pool.go:238
	// _ = "end of CoverTab[18364]"
}

// Subjects returns a list of the DER-encoded subjects of
//line /usr/local/go/src/crypto/x509/cert_pool.go:241
// all of the certificates in the pool.
//line /usr/local/go/src/crypto/x509/cert_pool.go:241
//
//line /usr/local/go/src/crypto/x509/cert_pool.go:241
// Deprecated: if s was returned by SystemCertPool, Subjects
//line /usr/local/go/src/crypto/x509/cert_pool.go:241
// will not include the system roots.
//line /usr/local/go/src/crypto/x509/cert_pool.go:246
func (s *CertPool) Subjects() [][]byte {
//line /usr/local/go/src/crypto/x509/cert_pool.go:246
	_go_fuzz_dep_.CoverTab[18380]++
							res := make([][]byte, s.len())
							for i, lc := range s.lazyCerts {
//line /usr/local/go/src/crypto/x509/cert_pool.go:248
		_go_fuzz_dep_.CoverTab[18382]++
								res[i] = lc.rawSubject
//line /usr/local/go/src/crypto/x509/cert_pool.go:249
		// _ = "end of CoverTab[18382]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:250
	// _ = "end of CoverTab[18380]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:250
	_go_fuzz_dep_.CoverTab[18381]++
							return res
//line /usr/local/go/src/crypto/x509/cert_pool.go:251
	// _ = "end of CoverTab[18381]"
}

// Equal reports whether s and other are equal.
func (s *CertPool) Equal(other *CertPool) bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:255
	_go_fuzz_dep_.CoverTab[18383]++
							if s == nil || func() bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:256
		_go_fuzz_dep_.CoverTab[18387]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:256
		return other == nil
//line /usr/local/go/src/crypto/x509/cert_pool.go:256
		// _ = "end of CoverTab[18387]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:256
	}() {
//line /usr/local/go/src/crypto/x509/cert_pool.go:256
		_go_fuzz_dep_.CoverTab[18388]++
								return s == other
//line /usr/local/go/src/crypto/x509/cert_pool.go:257
		// _ = "end of CoverTab[18388]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:258
		_go_fuzz_dep_.CoverTab[18389]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:258
		// _ = "end of CoverTab[18389]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:258
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:258
	// _ = "end of CoverTab[18383]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:258
	_go_fuzz_dep_.CoverTab[18384]++
							if s.systemPool != other.systemPool || func() bool {
//line /usr/local/go/src/crypto/x509/cert_pool.go:259
		_go_fuzz_dep_.CoverTab[18390]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:259
		return len(s.haveSum) != len(other.haveSum)
//line /usr/local/go/src/crypto/x509/cert_pool.go:259
		// _ = "end of CoverTab[18390]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:259
	}() {
//line /usr/local/go/src/crypto/x509/cert_pool.go:259
		_go_fuzz_dep_.CoverTab[18391]++
								return false
//line /usr/local/go/src/crypto/x509/cert_pool.go:260
		// _ = "end of CoverTab[18391]"
	} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:261
		_go_fuzz_dep_.CoverTab[18392]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:261
		// _ = "end of CoverTab[18392]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:261
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:261
	// _ = "end of CoverTab[18384]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:261
	_go_fuzz_dep_.CoverTab[18385]++
							for h := range s.haveSum {
//line /usr/local/go/src/crypto/x509/cert_pool.go:262
		_go_fuzz_dep_.CoverTab[18393]++
								if !other.haveSum[h] {
//line /usr/local/go/src/crypto/x509/cert_pool.go:263
			_go_fuzz_dep_.CoverTab[18394]++
									return false
//line /usr/local/go/src/crypto/x509/cert_pool.go:264
			// _ = "end of CoverTab[18394]"
		} else {
//line /usr/local/go/src/crypto/x509/cert_pool.go:265
			_go_fuzz_dep_.CoverTab[18395]++
//line /usr/local/go/src/crypto/x509/cert_pool.go:265
			// _ = "end of CoverTab[18395]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:265
		}
//line /usr/local/go/src/crypto/x509/cert_pool.go:265
		// _ = "end of CoverTab[18393]"
	}
//line /usr/local/go/src/crypto/x509/cert_pool.go:266
	// _ = "end of CoverTab[18385]"
//line /usr/local/go/src/crypto/x509/cert_pool.go:266
	_go_fuzz_dep_.CoverTab[18386]++
							return true
//line /usr/local/go/src/crypto/x509/cert_pool.go:267
	// _ = "end of CoverTab[18386]"
}

//line /usr/local/go/src/crypto/x509/cert_pool.go:268
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/cert_pool.go:268
var _ = _go_fuzz_dep_.CoverTab
