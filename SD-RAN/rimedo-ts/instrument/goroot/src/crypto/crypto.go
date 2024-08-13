// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/crypto.go:5
// Package crypto collects common cryptographic constants.
package crypto

//line /usr/local/go/src/crypto/crypto.go:6
import (
//line /usr/local/go/src/crypto/crypto.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/crypto.go:6
)
//line /usr/local/go/src/crypto/crypto.go:6
import (
//line /usr/local/go/src/crypto/crypto.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/crypto.go:6
)

import (
	"hash"
	"io"
	"strconv"
)

// Hash identifies a cryptographic hash function that is implemented in another
//line /usr/local/go/src/crypto/crypto.go:14
// package.
//line /usr/local/go/src/crypto/crypto.go:16
type Hash uint

// HashFunc simply returns the value of h so that Hash implements SignerOpts.
func (h Hash) HashFunc() Hash {
//line /usr/local/go/src/crypto/crypto.go:19
	_go_fuzz_dep_.CoverTab[1095]++
						return h
//line /usr/local/go/src/crypto/crypto.go:20
	// _ = "end of CoverTab[1095]"
}

func (h Hash) String() string {
//line /usr/local/go/src/crypto/crypto.go:23
	_go_fuzz_dep_.CoverTab[1096]++
						switch h {
	case MD4:
//line /usr/local/go/src/crypto/crypto.go:25
		_go_fuzz_dep_.CoverTab[1097]++
							return "MD4"
//line /usr/local/go/src/crypto/crypto.go:26
		// _ = "end of CoverTab[1097]"
	case MD5:
//line /usr/local/go/src/crypto/crypto.go:27
		_go_fuzz_dep_.CoverTab[1098]++
							return "MD5"
//line /usr/local/go/src/crypto/crypto.go:28
		// _ = "end of CoverTab[1098]"
	case SHA1:
//line /usr/local/go/src/crypto/crypto.go:29
		_go_fuzz_dep_.CoverTab[1099]++
							return "SHA-1"
//line /usr/local/go/src/crypto/crypto.go:30
		// _ = "end of CoverTab[1099]"
	case SHA224:
//line /usr/local/go/src/crypto/crypto.go:31
		_go_fuzz_dep_.CoverTab[1100]++
							return "SHA-224"
//line /usr/local/go/src/crypto/crypto.go:32
		// _ = "end of CoverTab[1100]"
	case SHA256:
//line /usr/local/go/src/crypto/crypto.go:33
		_go_fuzz_dep_.CoverTab[1101]++
							return "SHA-256"
//line /usr/local/go/src/crypto/crypto.go:34
		// _ = "end of CoverTab[1101]"
	case SHA384:
//line /usr/local/go/src/crypto/crypto.go:35
		_go_fuzz_dep_.CoverTab[1102]++
							return "SHA-384"
//line /usr/local/go/src/crypto/crypto.go:36
		// _ = "end of CoverTab[1102]"
	case SHA512:
//line /usr/local/go/src/crypto/crypto.go:37
		_go_fuzz_dep_.CoverTab[1103]++
							return "SHA-512"
//line /usr/local/go/src/crypto/crypto.go:38
		// _ = "end of CoverTab[1103]"
	case MD5SHA1:
//line /usr/local/go/src/crypto/crypto.go:39
		_go_fuzz_dep_.CoverTab[1104]++
							return "MD5+SHA1"
//line /usr/local/go/src/crypto/crypto.go:40
		// _ = "end of CoverTab[1104]"
	case RIPEMD160:
//line /usr/local/go/src/crypto/crypto.go:41
		_go_fuzz_dep_.CoverTab[1105]++
							return "RIPEMD-160"
//line /usr/local/go/src/crypto/crypto.go:42
		// _ = "end of CoverTab[1105]"
	case SHA3_224:
//line /usr/local/go/src/crypto/crypto.go:43
		_go_fuzz_dep_.CoverTab[1106]++
							return "SHA3-224"
//line /usr/local/go/src/crypto/crypto.go:44
		// _ = "end of CoverTab[1106]"
	case SHA3_256:
//line /usr/local/go/src/crypto/crypto.go:45
		_go_fuzz_dep_.CoverTab[1107]++
							return "SHA3-256"
//line /usr/local/go/src/crypto/crypto.go:46
		// _ = "end of CoverTab[1107]"
	case SHA3_384:
//line /usr/local/go/src/crypto/crypto.go:47
		_go_fuzz_dep_.CoverTab[1108]++
							return "SHA3-384"
//line /usr/local/go/src/crypto/crypto.go:48
		// _ = "end of CoverTab[1108]"
	case SHA3_512:
//line /usr/local/go/src/crypto/crypto.go:49
		_go_fuzz_dep_.CoverTab[1109]++
							return "SHA3-512"
//line /usr/local/go/src/crypto/crypto.go:50
		// _ = "end of CoverTab[1109]"
	case SHA512_224:
//line /usr/local/go/src/crypto/crypto.go:51
		_go_fuzz_dep_.CoverTab[1110]++
							return "SHA-512/224"
//line /usr/local/go/src/crypto/crypto.go:52
		// _ = "end of CoverTab[1110]"
	case SHA512_256:
//line /usr/local/go/src/crypto/crypto.go:53
		_go_fuzz_dep_.CoverTab[1111]++
							return "SHA-512/256"
//line /usr/local/go/src/crypto/crypto.go:54
		// _ = "end of CoverTab[1111]"
	case BLAKE2s_256:
//line /usr/local/go/src/crypto/crypto.go:55
		_go_fuzz_dep_.CoverTab[1112]++
							return "BLAKE2s-256"
//line /usr/local/go/src/crypto/crypto.go:56
		// _ = "end of CoverTab[1112]"
	case BLAKE2b_256:
//line /usr/local/go/src/crypto/crypto.go:57
		_go_fuzz_dep_.CoverTab[1113]++
							return "BLAKE2b-256"
//line /usr/local/go/src/crypto/crypto.go:58
		// _ = "end of CoverTab[1113]"
	case BLAKE2b_384:
//line /usr/local/go/src/crypto/crypto.go:59
		_go_fuzz_dep_.CoverTab[1114]++
							return "BLAKE2b-384"
//line /usr/local/go/src/crypto/crypto.go:60
		// _ = "end of CoverTab[1114]"
	case BLAKE2b_512:
//line /usr/local/go/src/crypto/crypto.go:61
		_go_fuzz_dep_.CoverTab[1115]++
							return "BLAKE2b-512"
//line /usr/local/go/src/crypto/crypto.go:62
		// _ = "end of CoverTab[1115]"
	default:
//line /usr/local/go/src/crypto/crypto.go:63
		_go_fuzz_dep_.CoverTab[1116]++
							return "unknown hash value " + strconv.Itoa(int(h))
//line /usr/local/go/src/crypto/crypto.go:64
		// _ = "end of CoverTab[1116]"
	}
//line /usr/local/go/src/crypto/crypto.go:65
	// _ = "end of CoverTab[1096]"
}

const (
	MD4		Hash	= 1 + iota	// import golang.org/x/crypto/md4
	MD5					// import crypto/md5
	SHA1					// import crypto/sha1
	SHA224					// import crypto/sha256
	SHA256					// import crypto/sha256
	SHA384					// import crypto/sha512
	SHA512					// import crypto/sha512
	MD5SHA1					// no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160				// import golang.org/x/crypto/ripemd160
	SHA3_224				// import golang.org/x/crypto/sha3
	SHA3_256				// import golang.org/x/crypto/sha3
	SHA3_384				// import golang.org/x/crypto/sha3
	SHA3_512				// import golang.org/x/crypto/sha3
	SHA512_224				// import crypto/sha512
	SHA512_256				// import crypto/sha512
	BLAKE2s_256				// import golang.org/x/crypto/blake2s
	BLAKE2b_256				// import golang.org/x/crypto/blake2b
	BLAKE2b_384				// import golang.org/x/crypto/blake2b
	BLAKE2b_512				// import golang.org/x/crypto/blake2b
	maxHash
)

var digestSizes = []uint8{
	MD4:		16,
	MD5:		16,
	SHA1:		20,
	SHA224:		28,
	SHA256:		32,
	SHA384:		48,
	SHA512:		64,
	SHA512_224:	28,
	SHA512_256:	32,
	SHA3_224:	28,
	SHA3_256:	32,
	SHA3_384:	48,
	SHA3_512:	64,
	MD5SHA1:	36,
	RIPEMD160:	20,
	BLAKE2s_256:	32,
	BLAKE2b_256:	32,
	BLAKE2b_384:	48,
	BLAKE2b_512:	64,
}

// Size returns the length, in bytes, of a digest resulting from the given hash
//line /usr/local/go/src/crypto/crypto.go:113
// function. It doesn't require that the hash function in question be linked
//line /usr/local/go/src/crypto/crypto.go:113
// into the program.
//line /usr/local/go/src/crypto/crypto.go:116
func (h Hash) Size() int {
//line /usr/local/go/src/crypto/crypto.go:116
	_go_fuzz_dep_.CoverTab[1117]++
						if h > 0 && func() bool {
//line /usr/local/go/src/crypto/crypto.go:117
		_go_fuzz_dep_.CoverTab[1119]++
//line /usr/local/go/src/crypto/crypto.go:117
		return h < maxHash
//line /usr/local/go/src/crypto/crypto.go:117
		// _ = "end of CoverTab[1119]"
//line /usr/local/go/src/crypto/crypto.go:117
	}() {
//line /usr/local/go/src/crypto/crypto.go:117
		_go_fuzz_dep_.CoverTab[1120]++
							return int(digestSizes[h])
//line /usr/local/go/src/crypto/crypto.go:118
		// _ = "end of CoverTab[1120]"
	} else {
//line /usr/local/go/src/crypto/crypto.go:119
		_go_fuzz_dep_.CoverTab[1121]++
//line /usr/local/go/src/crypto/crypto.go:119
		// _ = "end of CoverTab[1121]"
//line /usr/local/go/src/crypto/crypto.go:119
	}
//line /usr/local/go/src/crypto/crypto.go:119
	// _ = "end of CoverTab[1117]"
//line /usr/local/go/src/crypto/crypto.go:119
	_go_fuzz_dep_.CoverTab[1118]++
						panic("crypto: Size of unknown hash function")
//line /usr/local/go/src/crypto/crypto.go:120
	// _ = "end of CoverTab[1118]"
}

var hashes = make([]func() hash.Hash, maxHash)

// New returns a new hash.Hash calculating the given hash function. New panics
//line /usr/local/go/src/crypto/crypto.go:125
// if the hash function is not linked into the binary.
//line /usr/local/go/src/crypto/crypto.go:127
func (h Hash) New() hash.Hash {
//line /usr/local/go/src/crypto/crypto.go:127
	_go_fuzz_dep_.CoverTab[1122]++
						if h > 0 && func() bool {
//line /usr/local/go/src/crypto/crypto.go:128
		_go_fuzz_dep_.CoverTab[1124]++
//line /usr/local/go/src/crypto/crypto.go:128
		return h < maxHash
//line /usr/local/go/src/crypto/crypto.go:128
		// _ = "end of CoverTab[1124]"
//line /usr/local/go/src/crypto/crypto.go:128
	}() {
//line /usr/local/go/src/crypto/crypto.go:128
		_go_fuzz_dep_.CoverTab[1125]++
							f := hashes[h]
							if f != nil {
//line /usr/local/go/src/crypto/crypto.go:130
			_go_fuzz_dep_.CoverTab[1126]++
								return f()
//line /usr/local/go/src/crypto/crypto.go:131
			// _ = "end of CoverTab[1126]"
		} else {
//line /usr/local/go/src/crypto/crypto.go:132
			_go_fuzz_dep_.CoverTab[1127]++
//line /usr/local/go/src/crypto/crypto.go:132
			// _ = "end of CoverTab[1127]"
//line /usr/local/go/src/crypto/crypto.go:132
		}
//line /usr/local/go/src/crypto/crypto.go:132
		// _ = "end of CoverTab[1125]"
	} else {
//line /usr/local/go/src/crypto/crypto.go:133
		_go_fuzz_dep_.CoverTab[1128]++
//line /usr/local/go/src/crypto/crypto.go:133
		// _ = "end of CoverTab[1128]"
//line /usr/local/go/src/crypto/crypto.go:133
	}
//line /usr/local/go/src/crypto/crypto.go:133
	// _ = "end of CoverTab[1122]"
//line /usr/local/go/src/crypto/crypto.go:133
	_go_fuzz_dep_.CoverTab[1123]++
						panic("crypto: requested hash function #" + strconv.Itoa(int(h)) + " is unavailable")
//line /usr/local/go/src/crypto/crypto.go:134
	// _ = "end of CoverTab[1123]"
}

// Available reports whether the given hash function is linked into the binary.
func (h Hash) Available() bool {
//line /usr/local/go/src/crypto/crypto.go:138
	_go_fuzz_dep_.CoverTab[1129]++
						return h < maxHash && func() bool {
//line /usr/local/go/src/crypto/crypto.go:139
		_go_fuzz_dep_.CoverTab[1130]++
//line /usr/local/go/src/crypto/crypto.go:139
		return hashes[h] != nil
//line /usr/local/go/src/crypto/crypto.go:139
		// _ = "end of CoverTab[1130]"
//line /usr/local/go/src/crypto/crypto.go:139
	}()
//line /usr/local/go/src/crypto/crypto.go:139
	// _ = "end of CoverTab[1129]"
}

// RegisterHash registers a function that returns a new instance of the given
//line /usr/local/go/src/crypto/crypto.go:142
// hash function. This is intended to be called from the init function in
//line /usr/local/go/src/crypto/crypto.go:142
// packages that implement hash functions.
//line /usr/local/go/src/crypto/crypto.go:145
func RegisterHash(h Hash, f func() hash.Hash) {
//line /usr/local/go/src/crypto/crypto.go:145
	_go_fuzz_dep_.CoverTab[1131]++
						if h >= maxHash {
//line /usr/local/go/src/crypto/crypto.go:146
		_go_fuzz_dep_.CoverTab[1133]++
							panic("crypto: RegisterHash of unknown hash function")
//line /usr/local/go/src/crypto/crypto.go:147
		// _ = "end of CoverTab[1133]"
	} else {
//line /usr/local/go/src/crypto/crypto.go:148
		_go_fuzz_dep_.CoverTab[1134]++
//line /usr/local/go/src/crypto/crypto.go:148
		// _ = "end of CoverTab[1134]"
//line /usr/local/go/src/crypto/crypto.go:148
	}
//line /usr/local/go/src/crypto/crypto.go:148
	// _ = "end of CoverTab[1131]"
//line /usr/local/go/src/crypto/crypto.go:148
	_go_fuzz_dep_.CoverTab[1132]++
						hashes[h] = f
//line /usr/local/go/src/crypto/crypto.go:149
	// _ = "end of CoverTab[1132]"
}

// PublicKey represents a public key using an unspecified algorithm.
//line /usr/local/go/src/crypto/crypto.go:152
//
//line /usr/local/go/src/crypto/crypto.go:152
// Although this type is an empty interface for backwards compatibility reasons,
//line /usr/local/go/src/crypto/crypto.go:152
// all public key types in the standard library implement the following interface
//line /usr/local/go/src/crypto/crypto.go:152
//
//line /usr/local/go/src/crypto/crypto.go:152
//	interface{
//line /usr/local/go/src/crypto/crypto.go:152
//	    Equal(x crypto.PublicKey) bool
//line /usr/local/go/src/crypto/crypto.go:152
//	}
//line /usr/local/go/src/crypto/crypto.go:152
//
//line /usr/local/go/src/crypto/crypto.go:152
// which can be used for increased type safety within applications.
//line /usr/local/go/src/crypto/crypto.go:162
type PublicKey any

// PrivateKey represents a private key using an unspecified algorithm.
//line /usr/local/go/src/crypto/crypto.go:164
//
//line /usr/local/go/src/crypto/crypto.go:164
// Although this type is an empty interface for backwards compatibility reasons,
//line /usr/local/go/src/crypto/crypto.go:164
// all private key types in the standard library implement the following interface
//line /usr/local/go/src/crypto/crypto.go:164
//
//line /usr/local/go/src/crypto/crypto.go:164
//	interface{
//line /usr/local/go/src/crypto/crypto.go:164
//	    Public() crypto.PublicKey
//line /usr/local/go/src/crypto/crypto.go:164
//	    Equal(x crypto.PrivateKey) bool
//line /usr/local/go/src/crypto/crypto.go:164
//	}
//line /usr/local/go/src/crypto/crypto.go:164
//
//line /usr/local/go/src/crypto/crypto.go:164
// as well as purpose-specific interfaces such as Signer and Decrypter, which
//line /usr/local/go/src/crypto/crypto.go:164
// can be used for increased type safety within applications.
//line /usr/local/go/src/crypto/crypto.go:176
type PrivateKey any

// Signer is an interface for an opaque private key that can be used for
//line /usr/local/go/src/crypto/crypto.go:178
// signing operations. For example, an RSA key kept in a hardware module.
//line /usr/local/go/src/crypto/crypto.go:180
type Signer interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Sign signs digest with the private key, possibly using entropy from
	// rand. For an RSA key, the resulting signature should be either a
	// PKCS #1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
	// key, it should be a DER-serialised, ASN.1 signature structure.
	//
	// Hash implements the SignerOpts interface and, in most cases, one can
	// simply pass in the hash function used as opts. Sign may also attempt
	// to type assert opts to other types in order to obtain algorithm
	// specific values. See the documentation in each package for details.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest) and the hash function (as opts) to Sign.
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}

// SignerOpts contains options for signing with a Signer.
type SignerOpts interface {
	// HashFunc returns an identifier for the hash function used to produce
	// the message passed to Signer.Sign, or else zero to indicate that no
	// hashing was done.
	HashFunc() Hash
}

// Decrypter is an interface for an opaque private key that can be used for
//line /usr/local/go/src/crypto/crypto.go:209
// asymmetric decryption operations. An example would be an RSA key
//line /usr/local/go/src/crypto/crypto.go:209
// kept in a hardware module.
//line /usr/local/go/src/crypto/crypto.go:212
type Decrypter interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Decrypt decrypts msg. The opts argument should be appropriate for
	// the primitive used. See the documentation in each implementation for
	// details.
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts any

//line /usr/local/go/src/crypto/crypto.go:223
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/crypto.go:223
var _ = _go_fuzz_dep_.CoverTab
