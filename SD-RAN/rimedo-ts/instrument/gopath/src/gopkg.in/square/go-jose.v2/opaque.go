//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:17
)

// OpaqueSigner is an interface that supports signing payloads with opaque
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:19
// private key(s). Private key operations preformed by implementors may, for
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:19
// example, occur in a hardware module. An OpaqueSigner may rotate signing keys
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:19
// transparently to the user of this interface.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:23
type OpaqueSigner interface {
	// Public returns the public key of the current signing key.
	Public() *JSONWebKey
	// Algs returns a list of supported signing algorithms.
	Algs() []SignatureAlgorithm
	// SignPayload signs a payload with the current signing key using the given
	// algorithm.
	SignPayload(payload []byte, alg SignatureAlgorithm) ([]byte, error)
}

type opaqueSigner struct {
	signer OpaqueSigner
}

func newOpaqueSigner(alg SignatureAlgorithm, signer OpaqueSigner) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:37
	_go_fuzz_dep_.CoverTab[189888]++
											var algSupported bool
											for _, salg := range signer.Algs() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:39
		_go_fuzz_dep_.CoverTab[189891]++
												if alg == salg {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:40
			_go_fuzz_dep_.CoverTab[189892]++
													algSupported = true
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:42
			// _ = "end of CoverTab[189892]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:43
			_go_fuzz_dep_.CoverTab[189893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:43
			// _ = "end of CoverTab[189893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:43
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:43
		// _ = "end of CoverTab[189891]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:44
	// _ = "end of CoverTab[189888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:44
	_go_fuzz_dep_.CoverTab[189889]++
											if !algSupported {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:45
		_go_fuzz_dep_.CoverTab[189894]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:46
		// _ = "end of CoverTab[189894]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:47
		_go_fuzz_dep_.CoverTab[189895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:47
		// _ = "end of CoverTab[189895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:47
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:47
	// _ = "end of CoverTab[189889]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:47
	_go_fuzz_dep_.CoverTab[189890]++

											return recipientSigInfo{
		sigAlg:		alg,
		publicKey:	signer.Public,
		signer: &opaqueSigner{
			signer: signer,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:55
	// _ = "end of CoverTab[189890]"
}

func (o *opaqueSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:58
	_go_fuzz_dep_.CoverTab[189896]++
											out, err := o.signer.SignPayload(payload, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:60
		_go_fuzz_dep_.CoverTab[189898]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:61
		// _ = "end of CoverTab[189898]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:62
		_go_fuzz_dep_.CoverTab[189899]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:62
		// _ = "end of CoverTab[189899]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:62
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:62
	// _ = "end of CoverTab[189896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:62
	_go_fuzz_dep_.CoverTab[189897]++

											return Signature{
		Signature:	out,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:67
	// _ = "end of CoverTab[189897]"
}

// OpaqueVerifier is an interface that supports verifying payloads with opaque
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:70
// public key(s). An OpaqueSigner may rotate signing keys transparently to the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:70
// user of this interface.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:73
type OpaqueVerifier interface {
	VerifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error
}

type opaqueVerifier struct {
	verifier OpaqueVerifier
}

func (o *opaqueVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:81
	_go_fuzz_dep_.CoverTab[189900]++
											return o.verifier.VerifyPayload(payload, signature, alg)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:82
	// _ = "end of CoverTab[189900]"
}

// OpaqueKeyEncrypter is an interface that supports encrypting keys with an opaque key.
type OpaqueKeyEncrypter interface {
	// KeyID returns the kid
	KeyID() string
	// Algs returns a list of supported key encryption algorithms.
	Algs() []KeyAlgorithm
	// encryptKey encrypts the CEK using the given algorithm.
	encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error)
}

type opaqueKeyEncrypter struct {
	encrypter OpaqueKeyEncrypter
}

func newOpaqueKeyEncrypter(alg KeyAlgorithm, encrypter OpaqueKeyEncrypter) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:99
	_go_fuzz_dep_.CoverTab[189901]++
											var algSupported bool
											for _, salg := range encrypter.Algs() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:101
		_go_fuzz_dep_.CoverTab[189904]++
												if alg == salg {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:102
			_go_fuzz_dep_.CoverTab[189905]++
													algSupported = true
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:104
			// _ = "end of CoverTab[189905]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:105
			_go_fuzz_dep_.CoverTab[189906]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:105
			// _ = "end of CoverTab[189906]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:105
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:105
		// _ = "end of CoverTab[189904]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:106
	// _ = "end of CoverTab[189901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:106
	_go_fuzz_dep_.CoverTab[189902]++
											if !algSupported {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:107
		_go_fuzz_dep_.CoverTab[189907]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:108
		// _ = "end of CoverTab[189907]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:109
		_go_fuzz_dep_.CoverTab[189908]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:109
		// _ = "end of CoverTab[189908]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:109
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:109
	// _ = "end of CoverTab[189902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:109
	_go_fuzz_dep_.CoverTab[189903]++

											return recipientKeyInfo{
		keyID:	encrypter.KeyID(),
		keyAlg:	alg,
		keyEncrypter: &opaqueKeyEncrypter{
			encrypter: encrypter,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:117
	// _ = "end of CoverTab[189903]"
}

func (oke *opaqueKeyEncrypter) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:120
	_go_fuzz_dep_.CoverTab[189909]++
											return oke.encrypter.encryptKey(cek, alg)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:121
	// _ = "end of CoverTab[189909]"
}

// OpaqueKeyDecrypter is an interface that supports decrypting keys with an opaque key.
type OpaqueKeyDecrypter interface {
	DecryptKey(encryptedKey []byte, header Header) ([]byte, error)
}

type opaqueKeyDecrypter struct {
	decrypter OpaqueKeyDecrypter
}

func (okd *opaqueKeyDecrypter) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:133
	_go_fuzz_dep_.CoverTab[189910]++
											mergedHeaders := rawHeader{}
											mergedHeaders.merge(&headers)
											mergedHeaders.merge(recipient.header)

											header, err := mergedHeaders.sanitized()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:139
		_go_fuzz_dep_.CoverTab[189912]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:140
		// _ = "end of CoverTab[189912]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:141
		_go_fuzz_dep_.CoverTab[189913]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:141
		// _ = "end of CoverTab[189913]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:141
	// _ = "end of CoverTab[189910]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:141
	_go_fuzz_dep_.CoverTab[189911]++

											return okd.decrypter.DecryptKey(recipient.encryptedKey, header)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:143
	// _ = "end of CoverTab[189911]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:144
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/opaque.go:144
var _ = _go_fuzz_dep_.CoverTab
