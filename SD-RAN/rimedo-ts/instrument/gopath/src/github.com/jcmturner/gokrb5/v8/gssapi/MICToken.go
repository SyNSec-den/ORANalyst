//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
package gssapi

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:1
)

import (
	"bytes"
	"crypto/hmac"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/types"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:18
const (
	// MICTokenFlagSentByAcceptor - this flag indicates the sender is the context acceptor.  When not set, it indicates the sender is the context initiator
	MICTokenFlagSentByAcceptor	= 1 << iota
	// MICTokenFlagSealed - this flag indicates confidentiality is provided for.  It SHALL NOT be set in MIC tokens
	MICTokenFlagSealed
	// MICTokenFlagAcceptorSubkey - a subkey asserted by the context acceptor is used to protect the message
	MICTokenFlagAcceptorSubkey
)

const (
	micHdrLen = 16	// Length of the MIC Token's header
)

// MICToken represents a GSS API MIC token, as defined in RFC 4121.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:31
// It contains the header fields, the payload (this is not transmitted) and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:31
// the checksum, and provides the logic for converting to/from bytes plus
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:31
// computing and verifying checksums
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:35
type MICToken struct {
	// const GSS Token ID: 0x0404
	Flags	byte	// contains three flags: acceptor, sealed, acceptor subkey
	// const Filler: 0xFF 0xFF 0xFF 0xFF 0xFF
	SndSeqNum	uint64	// sender's sequence number. big-endian
	Payload		[]byte	// your data! :)
	Checksum	[]byte	// checksum of { payload | header }
}

// Return the 2 bytes identifying a GSS API MIC token
func getGSSMICTokenID() *[2]byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:45
	_go_fuzz_dep_.CoverTab[88780]++
												return &[2]byte{0x04, 0x04}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:46
	// _ = "end of CoverTab[88780]"
}

// Return the filler bytes used in header
func fillerBytes() *[5]byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:50
	_go_fuzz_dep_.CoverTab[88781]++
												return &[5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:51
	// _ = "end of CoverTab[88781]"
}

// Marshal the MICToken into a byte slice.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:54
// The payload should have been set and the checksum computed, otherwise an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:56
func (mt *MICToken) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:56
	_go_fuzz_dep_.CoverTab[88782]++
												if mt.Checksum == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:57
		_go_fuzz_dep_.CoverTab[88784]++
													return nil, errors.New("checksum has not been set")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:58
		// _ = "end of CoverTab[88784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:59
		_go_fuzz_dep_.CoverTab[88785]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:59
		// _ = "end of CoverTab[88785]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:59
	// _ = "end of CoverTab[88782]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:59
	_go_fuzz_dep_.CoverTab[88783]++

												bytes := make([]byte, micHdrLen+len(mt.Checksum))
												copy(bytes[0:micHdrLen], mt.getMICChecksumHeader()[:])
												copy(bytes[micHdrLen:], mt.Checksum)

												return bytes, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:65
	// _ = "end of CoverTab[88783]"
}

// SetChecksum uses the passed encryption key and key usage to compute the checksum over the payload and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:68
// the header, and sets the Checksum field of this MICToken.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:68
// If the payload has not been set or the checksum has already been set, an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:71
func (mt *MICToken) SetChecksum(key types.EncryptionKey, keyUsage uint32) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:71
	_go_fuzz_dep_.CoverTab[88786]++
												if mt.Checksum != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:72
		_go_fuzz_dep_.CoverTab[88789]++
													return errors.New("checksum has already been computed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:73
		// _ = "end of CoverTab[88789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:74
		_go_fuzz_dep_.CoverTab[88790]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:74
		// _ = "end of CoverTab[88790]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:74
	// _ = "end of CoverTab[88786]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:74
	_go_fuzz_dep_.CoverTab[88787]++
												checksum, err := mt.checksum(key, keyUsage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:76
		_go_fuzz_dep_.CoverTab[88791]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:77
		// _ = "end of CoverTab[88791]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:78
		_go_fuzz_dep_.CoverTab[88792]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:78
		// _ = "end of CoverTab[88792]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:78
	// _ = "end of CoverTab[88787]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:78
	_go_fuzz_dep_.CoverTab[88788]++
												mt.Checksum = checksum
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:80
	// _ = "end of CoverTab[88788]"
}

// Compute and return the checksum of this token, computed using the passed key and key usage.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:83
// Note: This will NOT update the struct's Checksum field.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:85
func (mt *MICToken) checksum(key types.EncryptionKey, keyUsage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:85
	_go_fuzz_dep_.CoverTab[88793]++
												if mt.Payload == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:86
		_go_fuzz_dep_.CoverTab[88796]++
													return nil, errors.New("cannot compute checksum with uninitialized payload")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:87
		// _ = "end of CoverTab[88796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:88
		_go_fuzz_dep_.CoverTab[88797]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:88
		// _ = "end of CoverTab[88797]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:88
	// _ = "end of CoverTab[88793]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:88
	_go_fuzz_dep_.CoverTab[88794]++
												d := make([]byte, micHdrLen+len(mt.Payload))
												copy(d[0:], mt.Payload)
												copy(d[len(mt.Payload):], mt.getMICChecksumHeader())

												encType, err := crypto.GetEtype(key.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:94
		_go_fuzz_dep_.CoverTab[88798]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:95
		// _ = "end of CoverTab[88798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:96
		_go_fuzz_dep_.CoverTab[88799]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:96
		// _ = "end of CoverTab[88799]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:96
	// _ = "end of CoverTab[88794]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:96
	_go_fuzz_dep_.CoverTab[88795]++
												return encType.GetChecksumHash(key.KeyValue, d, keyUsage)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:97
	// _ = "end of CoverTab[88795]"
}

// Build a header suitable for a checksum computation
func (mt *MICToken) getMICChecksumHeader() []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:101
	_go_fuzz_dep_.CoverTab[88800]++
												header := make([]byte, micHdrLen)
												copy(header[0:2], getGSSMICTokenID()[:])
												header[2] = mt.Flags
												copy(header[3:8], fillerBytes()[:])
												binary.BigEndian.PutUint64(header[8:16], mt.SndSeqNum)
												return header
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:107
	// _ = "end of CoverTab[88800]"
}

// Verify computes the token's checksum with the provided key and usage,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:110
// and compares it to the checksum present in the token.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:110
// In case of any failure, (false, err) is returned, with err an explanatory error.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:113
func (mt *MICToken) Verify(key types.EncryptionKey, keyUsage uint32) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:113
	_go_fuzz_dep_.CoverTab[88801]++
												computed, err := mt.checksum(key, keyUsage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:115
		_go_fuzz_dep_.CoverTab[88804]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:116
		// _ = "end of CoverTab[88804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:117
		_go_fuzz_dep_.CoverTab[88805]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:117
		// _ = "end of CoverTab[88805]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:117
	// _ = "end of CoverTab[88801]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:117
	_go_fuzz_dep_.CoverTab[88802]++
												if !hmac.Equal(computed, mt.Checksum) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:118
		_go_fuzz_dep_.CoverTab[88806]++
													return false, fmt.Errorf(
			"checksum mismatch. Computed: %s, Contained in token: %s",
			hex.EncodeToString(computed), hex.EncodeToString(mt.Checksum))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:121
		// _ = "end of CoverTab[88806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:122
		_go_fuzz_dep_.CoverTab[88807]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:122
		// _ = "end of CoverTab[88807]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:122
	// _ = "end of CoverTab[88802]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:122
	_go_fuzz_dep_.CoverTab[88803]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:123
	// _ = "end of CoverTab[88803]"
}

// Unmarshal bytes into the corresponding MICToken.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:126
// If expectFromAcceptor is true we expect the token to have been emitted by the gss acceptor,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:126
// and will check the according flag, returning an error if the token does not match the expectation.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:129
func (mt *MICToken) Unmarshal(b []byte, expectFromAcceptor bool) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:129
	_go_fuzz_dep_.CoverTab[88808]++
												if len(b) < micHdrLen {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:130
		_go_fuzz_dep_.CoverTab[88814]++
													return errors.New("bytes shorter than header length")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:131
		// _ = "end of CoverTab[88814]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:132
		_go_fuzz_dep_.CoverTab[88815]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:132
		// _ = "end of CoverTab[88815]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:132
	// _ = "end of CoverTab[88808]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:132
	_go_fuzz_dep_.CoverTab[88809]++
												if !bytes.Equal(getGSSMICTokenID()[:], b[0:2]) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:133
		_go_fuzz_dep_.CoverTab[88816]++
													return fmt.Errorf("wrong Token ID, Expected %s, was %s",
			hex.EncodeToString(getGSSMICTokenID()[:]),
			hex.EncodeToString(b[0:2]))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:136
		// _ = "end of CoverTab[88816]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:137
		_go_fuzz_dep_.CoverTab[88817]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:137
		// _ = "end of CoverTab[88817]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:137
	// _ = "end of CoverTab[88809]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:137
	_go_fuzz_dep_.CoverTab[88810]++
												flags := b[2]
												isFromAcceptor := flags&MICTokenFlagSentByAcceptor != 0
												if isFromAcceptor && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:140
		_go_fuzz_dep_.CoverTab[88818]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:140
		return !expectFromAcceptor
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:140
		// _ = "end of CoverTab[88818]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:140
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:140
		_go_fuzz_dep_.CoverTab[88819]++
													return errors.New("unexpected acceptor flag is set: not expecting a token from the acceptor")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:141
		// _ = "end of CoverTab[88819]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:142
		_go_fuzz_dep_.CoverTab[88820]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:142
		// _ = "end of CoverTab[88820]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:142
	// _ = "end of CoverTab[88810]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:142
	_go_fuzz_dep_.CoverTab[88811]++
												if !isFromAcceptor && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:143
		_go_fuzz_dep_.CoverTab[88821]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:143
		return expectFromAcceptor
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:143
		// _ = "end of CoverTab[88821]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:143
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:143
		_go_fuzz_dep_.CoverTab[88822]++
													return errors.New("unexpected acceptor flag is not set: expecting a token from the acceptor, not in the initiator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:144
		// _ = "end of CoverTab[88822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:145
		_go_fuzz_dep_.CoverTab[88823]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:145
		// _ = "end of CoverTab[88823]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:145
	// _ = "end of CoverTab[88811]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:145
	_go_fuzz_dep_.CoverTab[88812]++
												if !bytes.Equal(b[3:8], fillerBytes()[:]) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:146
		_go_fuzz_dep_.CoverTab[88824]++
													return fmt.Errorf("unexpected filler bytes: expecting %s, was %s",
			hex.EncodeToString(fillerBytes()[:]),
			hex.EncodeToString(b[3:8]))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:149
		// _ = "end of CoverTab[88824]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:150
		_go_fuzz_dep_.CoverTab[88825]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:150
		// _ = "end of CoverTab[88825]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:150
	// _ = "end of CoverTab[88812]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:150
	_go_fuzz_dep_.CoverTab[88813]++

												mt.Flags = flags
												mt.SndSeqNum = binary.BigEndian.Uint64(b[8:16])
												mt.Checksum = b[micHdrLen:]
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:155
	// _ = "end of CoverTab[88813]"
}

// NewInitiatorMICToken builds a new initiator token (acceptor flag will be set to 0) and computes the authenticated checksum.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:158
// Other flags are set to 0.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:158
// Note that in certain circumstances you may need to provide a sequence number that has been defined earlier.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:158
// This is currently not supported.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:162
func NewInitiatorMICToken(payload []byte, key types.EncryptionKey) (*MICToken, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:162
	_go_fuzz_dep_.CoverTab[88826]++
												token := MICToken{
		Flags:		0x00,
		SndSeqNum:	0,
		Payload:	payload,
	}

	if err := token.SetChecksum(key, keyusage.GSSAPI_INITIATOR_SIGN); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:169
		_go_fuzz_dep_.CoverTab[88828]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:170
		// _ = "end of CoverTab[88828]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:171
		_go_fuzz_dep_.CoverTab[88829]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:171
		// _ = "end of CoverTab[88829]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:171
	// _ = "end of CoverTab[88826]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:171
	_go_fuzz_dep_.CoverTab[88827]++

												return &token, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:173
	// _ = "end of CoverTab[88827]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:174
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/MICToken.go:174
var _ = _go_fuzz_dep_.CoverTab
