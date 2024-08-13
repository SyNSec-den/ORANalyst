//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
package gssapi

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:1
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

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:18
const (
	// HdrLen is the length of the Wrap Token's header
	HdrLen	= 16
	// FillerByte is a filler in the WrapToken structure
	FillerByte	byte	= 0xFF
)

// WrapToken represents a GSS API Wrap token, as defined in RFC 4121.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:25
// It contains the header fields, the payload and the checksum, and provides
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:25
// the logic for converting to/from bytes plus computing and verifying checksums
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:28
type WrapToken struct {
	// const GSS Token ID: 0x0504
	Flags	byte	// contains three flags: acceptor, sealed, acceptor subkey
	// const Filler: 0xFF
	EC		uint16	// checksum length. big-endian
	RRC		uint16	// right rotation count. big-endian
	SndSeqNum	uint64	// sender's sequence number. big-endian
	Payload		[]byte	// your data! :)
	CheckSum	[]byte	// authenticated checksum of { payload | header }
}

// Return the 2 bytes identifying a GSS API Wrap token
func getGssWrapTokenId() *[2]byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:40
	_go_fuzz_dep_.CoverTab[88869]++
												return &[2]byte{0x05, 0x04}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:41
	// _ = "end of CoverTab[88869]"
}

// Marshal the WrapToken into a byte slice.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:44
// The payload should have been set and the checksum computed, otherwise an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:46
func (wt *WrapToken) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:46
	_go_fuzz_dep_.CoverTab[88870]++
												if wt.CheckSum == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:47
		_go_fuzz_dep_.CoverTab[88873]++
													return nil, errors.New("checksum has not been set")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:48
		// _ = "end of CoverTab[88873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:49
		_go_fuzz_dep_.CoverTab[88874]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:49
		// _ = "end of CoverTab[88874]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:49
	// _ = "end of CoverTab[88870]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:49
	_go_fuzz_dep_.CoverTab[88871]++
												if wt.Payload == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:50
		_go_fuzz_dep_.CoverTab[88875]++
													return nil, errors.New("payload has not been set")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:51
		// _ = "end of CoverTab[88875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:52
		_go_fuzz_dep_.CoverTab[88876]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:52
		// _ = "end of CoverTab[88876]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:52
	// _ = "end of CoverTab[88871]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:52
	_go_fuzz_dep_.CoverTab[88872]++

												pldOffset := HdrLen
												chkSOffset := HdrLen + len(wt.Payload)

												bytes := make([]byte, chkSOffset+int(wt.EC))
												copy(bytes[0:], getGssWrapTokenId()[:])
												bytes[2] = wt.Flags
												bytes[3] = FillerByte
												binary.BigEndian.PutUint16(bytes[4:6], wt.EC)
												binary.BigEndian.PutUint16(bytes[6:8], wt.RRC)
												binary.BigEndian.PutUint64(bytes[8:16], wt.SndSeqNum)
												copy(bytes[pldOffset:], wt.Payload)
												copy(bytes[chkSOffset:], wt.CheckSum)
												return bytes, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:66
	// _ = "end of CoverTab[88872]"
}

// SetCheckSum uses the passed encryption key and key usage to compute the checksum over the payload and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:69
// the header, and sets the CheckSum field of this WrapToken.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:69
// If the payload has not been set or the checksum has already been set, an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:72
func (wt *WrapToken) SetCheckSum(key types.EncryptionKey, keyUsage uint32) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:72
	_go_fuzz_dep_.CoverTab[88877]++
												if wt.Payload == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:73
		_go_fuzz_dep_.CoverTab[88881]++
													return errors.New("payload has not been set")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:74
		// _ = "end of CoverTab[88881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:75
		_go_fuzz_dep_.CoverTab[88882]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:75
		// _ = "end of CoverTab[88882]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:75
	// _ = "end of CoverTab[88877]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:75
	_go_fuzz_dep_.CoverTab[88878]++
												if wt.CheckSum != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:76
		_go_fuzz_dep_.CoverTab[88883]++
													return errors.New("checksum has already been computed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:77
		// _ = "end of CoverTab[88883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:78
		_go_fuzz_dep_.CoverTab[88884]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:78
		// _ = "end of CoverTab[88884]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:78
	// _ = "end of CoverTab[88878]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:78
	_go_fuzz_dep_.CoverTab[88879]++
												chkSum, cErr := wt.computeCheckSum(key, keyUsage)
												if cErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:80
		_go_fuzz_dep_.CoverTab[88885]++
													return cErr
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:81
		// _ = "end of CoverTab[88885]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:82
		_go_fuzz_dep_.CoverTab[88886]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:82
		// _ = "end of CoverTab[88886]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:82
	// _ = "end of CoverTab[88879]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:82
	_go_fuzz_dep_.CoverTab[88880]++
												wt.CheckSum = chkSum
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:84
	// _ = "end of CoverTab[88880]"
}

// ComputeCheckSum computes and returns the checksum of this token, computed using the passed key and key usage.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:87
// Note: This will NOT update the struct's Checksum field.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:89
func (wt *WrapToken) computeCheckSum(key types.EncryptionKey, keyUsage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:89
	_go_fuzz_dep_.CoverTab[88887]++
												if wt.Payload == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:90
		_go_fuzz_dep_.CoverTab[88890]++
													return nil, errors.New("cannot compute checksum with uninitialized payload")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:91
		// _ = "end of CoverTab[88890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:92
		_go_fuzz_dep_.CoverTab[88891]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:92
		// _ = "end of CoverTab[88891]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:92
	// _ = "end of CoverTab[88887]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:92
	_go_fuzz_dep_.CoverTab[88888]++

												checksumMe := make([]byte, HdrLen+len(wt.Payload))
												copy(checksumMe[0:], wt.Payload)
												copy(checksumMe[len(wt.Payload):], getChecksumHeader(wt.Flags, wt.SndSeqNum))

												encType, err := crypto.GetEtype(key.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:99
			_go_fuzz_dep_.CoverTab[88892]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:100
		// _ = "end of CoverTab[88892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:101
		_go_fuzz_dep_.CoverTab[88893]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:101
		// _ = "end of CoverTab[88893]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:101
	// _ = "end of CoverTab[88888]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:101
	_go_fuzz_dep_.CoverTab[88889]++
													return encType.GetChecksumHash(key.KeyValue, checksumMe, keyUsage)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:102
	// _ = "end of CoverTab[88889]"
}

// Build a header suitable for a checksum computation
func getChecksumHeader(flags byte, senderSeqNum uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:106
	_go_fuzz_dep_.CoverTab[88894]++
													header := make([]byte, 16)
													copy(header[0:], []byte{0x05, 0x04, flags, 0xFF, 0x00, 0x00, 0x00, 0x00})
													binary.BigEndian.PutUint64(header[8:], senderSeqNum)
													return header
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:110
	// _ = "end of CoverTab[88894]"
}

// Verify computes the token's checksum with the provided key and usage,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:113
// and compares it to the checksum present in the token.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:113
// In case of any failure, (false, Err) is returned, with Err an explanatory error.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:116
func (wt *WrapToken) Verify(key types.EncryptionKey, keyUsage uint32) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:116
	_go_fuzz_dep_.CoverTab[88895]++
													computed, cErr := wt.computeCheckSum(key, keyUsage)
													if cErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:118
		_go_fuzz_dep_.CoverTab[88898]++
														return false, cErr
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:119
		// _ = "end of CoverTab[88898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:120
		_go_fuzz_dep_.CoverTab[88899]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:120
		// _ = "end of CoverTab[88899]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:120
	// _ = "end of CoverTab[88895]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:120
	_go_fuzz_dep_.CoverTab[88896]++
													if !hmac.Equal(computed, wt.CheckSum) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:121
		_go_fuzz_dep_.CoverTab[88900]++
														return false, fmt.Errorf(
			"checksum mismatch. Computed: %s, Contained in token: %s",
			hex.EncodeToString(computed), hex.EncodeToString(wt.CheckSum))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:124
		// _ = "end of CoverTab[88900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:125
		_go_fuzz_dep_.CoverTab[88901]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:125
		// _ = "end of CoverTab[88901]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:125
	// _ = "end of CoverTab[88896]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:125
	_go_fuzz_dep_.CoverTab[88897]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:126
	// _ = "end of CoverTab[88897]"
}

// Unmarshal bytes into the corresponding WrapToken.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:129
// If expectFromAcceptor is true, we expect the token to have been emitted by the gss acceptor,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:129
// and will check the according flag, returning an error if the token does not match the expectation.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:132
func (wt *WrapToken) Unmarshal(b []byte, expectFromAcceptor bool) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:132
	_go_fuzz_dep_.CoverTab[88902]++

													if len(b) < 16 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:134
		_go_fuzz_dep_.CoverTab[88909]++
														return errors.New("bytes shorter than header length")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:135
		// _ = "end of CoverTab[88909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:136
		_go_fuzz_dep_.CoverTab[88910]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:136
		// _ = "end of CoverTab[88910]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:136
	// _ = "end of CoverTab[88902]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:136
	_go_fuzz_dep_.CoverTab[88903]++

													if !bytes.Equal(getGssWrapTokenId()[:], b[0:2]) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:138
		_go_fuzz_dep_.CoverTab[88911]++
														return fmt.Errorf("wrong Token ID. Expected %s, was %s",
			hex.EncodeToString(getGssWrapTokenId()[:]),
			hex.EncodeToString(b[0:2]))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:141
		// _ = "end of CoverTab[88911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:142
		_go_fuzz_dep_.CoverTab[88912]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:142
		// _ = "end of CoverTab[88912]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:142
	// _ = "end of CoverTab[88903]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:142
	_go_fuzz_dep_.CoverTab[88904]++

													flags := b[2]
													isFromAcceptor := flags&0x01 == 1
													if isFromAcceptor && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:146
		_go_fuzz_dep_.CoverTab[88913]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:146
		return !expectFromAcceptor
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:146
		// _ = "end of CoverTab[88913]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:146
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:146
		_go_fuzz_dep_.CoverTab[88914]++
														return errors.New("unexpected acceptor flag is set: not expecting a token from the acceptor")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:147
		// _ = "end of CoverTab[88914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:148
		_go_fuzz_dep_.CoverTab[88915]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:148
		// _ = "end of CoverTab[88915]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:148
	// _ = "end of CoverTab[88904]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:148
	_go_fuzz_dep_.CoverTab[88905]++
													if !isFromAcceptor && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:149
		_go_fuzz_dep_.CoverTab[88916]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:149
		return expectFromAcceptor
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:149
		// _ = "end of CoverTab[88916]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:149
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:149
		_go_fuzz_dep_.CoverTab[88917]++
														return errors.New("expected acceptor flag is not set: expecting a token from the acceptor, not the initiator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:150
		// _ = "end of CoverTab[88917]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:151
		_go_fuzz_dep_.CoverTab[88918]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:151
		// _ = "end of CoverTab[88918]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:151
	// _ = "end of CoverTab[88905]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:151
	_go_fuzz_dep_.CoverTab[88906]++

													if b[3] != FillerByte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:153
		_go_fuzz_dep_.CoverTab[88919]++
														return fmt.Errorf("unexpected filler byte: expecting 0xFF, was %s ", hex.EncodeToString(b[3:4]))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:154
		// _ = "end of CoverTab[88919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:155
		_go_fuzz_dep_.CoverTab[88920]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:155
		// _ = "end of CoverTab[88920]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:155
	// _ = "end of CoverTab[88906]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:155
	_go_fuzz_dep_.CoverTab[88907]++
													checksumL := binary.BigEndian.Uint16(b[4:6])

													if int(checksumL) > len(b)-HdrLen {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:158
		_go_fuzz_dep_.CoverTab[88921]++
														return fmt.Errorf("inconsistent checksum length: %d bytes to parse, checksum length is %d", len(b), checksumL)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:159
		// _ = "end of CoverTab[88921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:160
		_go_fuzz_dep_.CoverTab[88922]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:160
		// _ = "end of CoverTab[88922]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:160
	// _ = "end of CoverTab[88907]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:160
	_go_fuzz_dep_.CoverTab[88908]++

													wt.Flags = flags
													wt.EC = checksumL
													wt.RRC = binary.BigEndian.Uint16(b[6:8])
													wt.SndSeqNum = binary.BigEndian.Uint64(b[8:16])
													wt.Payload = b[16 : len(b)-int(checksumL)]
													wt.CheckSum = b[len(b)-int(checksumL):]
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:168
	// _ = "end of CoverTab[88908]"
}

// NewInitiatorWrapToken builds a new initiator token (acceptor flag will be set to 0) and computes the authenticated checksum.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:171
// Other flags are set to 0, and the RRC and sequence number are initialized to 0.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:171
// Note that in certain circumstances you may need to provide a sequence number that has been defined earlier.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:171
// This is currently not supported.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:175
func NewInitiatorWrapToken(payload []byte, key types.EncryptionKey) (*WrapToken, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:175
	_go_fuzz_dep_.CoverTab[88923]++
													encType, err := crypto.GetEtype(key.KeyType)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:177
		_go_fuzz_dep_.CoverTab[88926]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:178
		// _ = "end of CoverTab[88926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:179
		_go_fuzz_dep_.CoverTab[88927]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:179
		// _ = "end of CoverTab[88927]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:179
	// _ = "end of CoverTab[88923]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:179
	_go_fuzz_dep_.CoverTab[88924]++

													token := WrapToken{
		Flags:	0x00,

		EC:		uint16(encType.GetHMACBitLength() / 8),
		RRC:		0,
		SndSeqNum:	0,
		Payload:	payload,
	}

	if err := token.SetCheckSum(key, keyusage.GSSAPI_INITIATOR_SEAL); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:190
		_go_fuzz_dep_.CoverTab[88928]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:191
		// _ = "end of CoverTab[88928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:192
		_go_fuzz_dep_.CoverTab[88929]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:192
		// _ = "end of CoverTab[88929]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:192
	// _ = "end of CoverTab[88924]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:192
	_go_fuzz_dep_.CoverTab[88925]++

													return &token, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:194
	// _ = "end of CoverTab[88925]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:195
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/wrapToken.go:195
var _ = _go_fuzz_dep_.CoverTab
