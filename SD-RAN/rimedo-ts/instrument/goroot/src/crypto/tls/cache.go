// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/cache.go:5
package tls

//line /usr/local/go/src/crypto/tls/cache.go:5
import (
//line /usr/local/go/src/crypto/tls/cache.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/cache.go:5
)
//line /usr/local/go/src/crypto/tls/cache.go:5
import (
//line /usr/local/go/src/crypto/tls/cache.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/cache.go:5
)

import (
	"crypto/x509"
	"runtime"
	"sync"
	"sync/atomic"
)

type cacheEntry struct {
	refs	atomic.Int64
	cert	*x509.Certificate
}

// certCache implements an intern table for reference counted x509.Certificates,
//line /usr/local/go/src/crypto/tls/cache.go:19
// implemented in a similar fashion to BoringSSL's CRYPTO_BUFFER_POOL. This
//line /usr/local/go/src/crypto/tls/cache.go:19
// allows for a single x509.Certificate to be kept in memory and referenced from
//line /usr/local/go/src/crypto/tls/cache.go:19
// multiple Conns. Returned references should not be mutated by callers. Certificates
//line /usr/local/go/src/crypto/tls/cache.go:19
// are still safe to use after they are removed from the cache.
//line /usr/local/go/src/crypto/tls/cache.go:19
//
//line /usr/local/go/src/crypto/tls/cache.go:19
// Certificates are returned wrapped in a activeCert struct that should be held by
//line /usr/local/go/src/crypto/tls/cache.go:19
// the caller. When references to the activeCert are freed, the number of references
//line /usr/local/go/src/crypto/tls/cache.go:19
// to the certificate in the cache is decremented. Once the number of references
//line /usr/local/go/src/crypto/tls/cache.go:19
// reaches zero, the entry is evicted from the cache.
//line /usr/local/go/src/crypto/tls/cache.go:19
//
//line /usr/local/go/src/crypto/tls/cache.go:19
// The main difference between this implementation and CRYPTO_BUFFER_POOL is that
//line /usr/local/go/src/crypto/tls/cache.go:19
// CRYPTO_BUFFER_POOL is a more  generic structure which supports blobs of data,
//line /usr/local/go/src/crypto/tls/cache.go:19
// rather than specific structures. Since we only care about x509.Certificates,
//line /usr/local/go/src/crypto/tls/cache.go:19
// certCache is implemented as a specific cache, rather than a generic one.
//line /usr/local/go/src/crypto/tls/cache.go:19
//
//line /usr/local/go/src/crypto/tls/cache.go:19
// See https://boringssl.googlesource.com/boringssl/+/master/include/openssl/pool.h
//line /usr/local/go/src/crypto/tls/cache.go:19
// and https://boringssl.googlesource.com/boringssl/+/master/crypto/pool/pool.c
//line /usr/local/go/src/crypto/tls/cache.go:19
// for the BoringSSL reference.
//line /usr/local/go/src/crypto/tls/cache.go:38
type certCache struct {
	sync.Map
}

var clientCertCache = new(certCache)

// activeCert is a handle to a certificate held in the cache. Once there are
//line /usr/local/go/src/crypto/tls/cache.go:44
// no alive activeCerts for a given certificate, the certificate is removed
//line /usr/local/go/src/crypto/tls/cache.go:44
// from the cache by a finalizer.
//line /usr/local/go/src/crypto/tls/cache.go:47
type activeCert struct {
	cert *x509.Certificate
}

// active increments the number of references to the entry, wraps the
//line /usr/local/go/src/crypto/tls/cache.go:51
// certificate in the entry in a activeCert, and sets the finalizer.
//line /usr/local/go/src/crypto/tls/cache.go:51
//
//line /usr/local/go/src/crypto/tls/cache.go:51
// Note that there is a race between active and the finalizer set on the
//line /usr/local/go/src/crypto/tls/cache.go:51
// returned activeCert, triggered if active is called after the ref count is
//line /usr/local/go/src/crypto/tls/cache.go:51
// decremented such that refs may be > 0 when evict is called. We consider this
//line /usr/local/go/src/crypto/tls/cache.go:51
// safe, since the caller holding an activeCert for an entry that is no longer
//line /usr/local/go/src/crypto/tls/cache.go:51
// in the cache is fine, with the only side effect being the memory overhead of
//line /usr/local/go/src/crypto/tls/cache.go:51
// there being more than one distinct reference to a certificate alive at once.
//line /usr/local/go/src/crypto/tls/cache.go:60
func (cc *certCache) active(e *cacheEntry) *activeCert {
//line /usr/local/go/src/crypto/tls/cache.go:60
	_go_fuzz_dep_.CoverTab[21123]++
							e.refs.Add(1)
							a := &activeCert{e.cert}
							runtime.SetFinalizer(a, func(_ *activeCert) {
//line /usr/local/go/src/crypto/tls/cache.go:63
		_go_fuzz_dep_.CoverTab[21125]++
								if e.refs.Add(-1) == 0 {
//line /usr/local/go/src/crypto/tls/cache.go:64
			_go_fuzz_dep_.CoverTab[21126]++
									cc.evict(e)
//line /usr/local/go/src/crypto/tls/cache.go:65
			// _ = "end of CoverTab[21126]"
		} else {
//line /usr/local/go/src/crypto/tls/cache.go:66
			_go_fuzz_dep_.CoverTab[21127]++
//line /usr/local/go/src/crypto/tls/cache.go:66
			// _ = "end of CoverTab[21127]"
//line /usr/local/go/src/crypto/tls/cache.go:66
		}
//line /usr/local/go/src/crypto/tls/cache.go:66
		// _ = "end of CoverTab[21125]"
	})
//line /usr/local/go/src/crypto/tls/cache.go:67
	// _ = "end of CoverTab[21123]"
//line /usr/local/go/src/crypto/tls/cache.go:67
	_go_fuzz_dep_.CoverTab[21124]++
							return a
//line /usr/local/go/src/crypto/tls/cache.go:68
	// _ = "end of CoverTab[21124]"
}

// evict removes a cacheEntry from the cache.
func (cc *certCache) evict(e *cacheEntry) {
//line /usr/local/go/src/crypto/tls/cache.go:72
	_go_fuzz_dep_.CoverTab[21128]++
							cc.Delete(string(e.cert.Raw))
//line /usr/local/go/src/crypto/tls/cache.go:73
	// _ = "end of CoverTab[21128]"
}

// newCert returns a x509.Certificate parsed from der. If there is already a copy
//line /usr/local/go/src/crypto/tls/cache.go:76
// of the certificate in the cache, a reference to the existing certificate will
//line /usr/local/go/src/crypto/tls/cache.go:76
// be returned. Otherwise, a fresh certificate will be added to the cache, and
//line /usr/local/go/src/crypto/tls/cache.go:76
// the reference returned. The returned reference should not be mutated.
//line /usr/local/go/src/crypto/tls/cache.go:80
func (cc *certCache) newCert(der []byte) (*activeCert, error) {
//line /usr/local/go/src/crypto/tls/cache.go:80
	_go_fuzz_dep_.CoverTab[21129]++
							if entry, ok := cc.Load(string(der)); ok {
//line /usr/local/go/src/crypto/tls/cache.go:81
		_go_fuzz_dep_.CoverTab[21133]++
								return cc.active(entry.(*cacheEntry)), nil
//line /usr/local/go/src/crypto/tls/cache.go:82
		// _ = "end of CoverTab[21133]"
	} else {
//line /usr/local/go/src/crypto/tls/cache.go:83
		_go_fuzz_dep_.CoverTab[21134]++
//line /usr/local/go/src/crypto/tls/cache.go:83
		// _ = "end of CoverTab[21134]"
//line /usr/local/go/src/crypto/tls/cache.go:83
	}
//line /usr/local/go/src/crypto/tls/cache.go:83
	// _ = "end of CoverTab[21129]"
//line /usr/local/go/src/crypto/tls/cache.go:83
	_go_fuzz_dep_.CoverTab[21130]++

							cert, err := x509.ParseCertificate(der)
							if err != nil {
//line /usr/local/go/src/crypto/tls/cache.go:86
		_go_fuzz_dep_.CoverTab[21135]++
								return nil, err
//line /usr/local/go/src/crypto/tls/cache.go:87
		// _ = "end of CoverTab[21135]"
	} else {
//line /usr/local/go/src/crypto/tls/cache.go:88
		_go_fuzz_dep_.CoverTab[21136]++
//line /usr/local/go/src/crypto/tls/cache.go:88
		// _ = "end of CoverTab[21136]"
//line /usr/local/go/src/crypto/tls/cache.go:88
	}
//line /usr/local/go/src/crypto/tls/cache.go:88
	// _ = "end of CoverTab[21130]"
//line /usr/local/go/src/crypto/tls/cache.go:88
	_go_fuzz_dep_.CoverTab[21131]++

							entry := &cacheEntry{cert: cert}
							if entry, loaded := cc.LoadOrStore(string(der), entry); loaded {
//line /usr/local/go/src/crypto/tls/cache.go:91
		_go_fuzz_dep_.CoverTab[21137]++
								return cc.active(entry.(*cacheEntry)), nil
//line /usr/local/go/src/crypto/tls/cache.go:92
		// _ = "end of CoverTab[21137]"
	} else {
//line /usr/local/go/src/crypto/tls/cache.go:93
		_go_fuzz_dep_.CoverTab[21138]++
//line /usr/local/go/src/crypto/tls/cache.go:93
		// _ = "end of CoverTab[21138]"
//line /usr/local/go/src/crypto/tls/cache.go:93
	}
//line /usr/local/go/src/crypto/tls/cache.go:93
	// _ = "end of CoverTab[21131]"
//line /usr/local/go/src/crypto/tls/cache.go:93
	_go_fuzz_dep_.CoverTab[21132]++
							return cc.active(entry), nil
//line /usr/local/go/src/crypto/tls/cache.go:94
	// _ = "end of CoverTab[21132]"
}

//line /usr/local/go/src/crypto/tls/cache.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/cache.go:95
var _ = _go_fuzz_dep_.CoverTab
