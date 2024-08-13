//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:1
// Package etype provides the Kerberos Encryption Type interface
package etype

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:2
)

import "hash"

// EType is the interface defining the Encryption Type.
type EType interface {
	GetETypeID() int32
	GetHashID() int32
	GetKeyByteSize() int
	GetKeySeedBitLength() int
	GetDefaultStringToKeyParams() string
	StringToKey(string, salt, s2kparams string) ([]byte, error)
	RandomToKey(b []byte) []byte
	GetHMACBitLength() int
	GetMessageBlockByteSize() int
	EncryptData(key, data []byte) ([]byte, []byte, error)
	EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error)
	DecryptData(key, data []byte) ([]byte, error)
	DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error)
	GetCypherBlockBitLength() int
	GetConfounderByteSize() int
	DeriveKey(protocolKey, usage []byte) ([]byte, error)
	DeriveRandom(protocolKey, usage []byte) ([]byte, error)
	VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool
	GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error)
	VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool
	GetHashFunc() func() hash.Hash
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/etype/etype.go:29
var _ = _go_fuzz_dep_.CoverTab
