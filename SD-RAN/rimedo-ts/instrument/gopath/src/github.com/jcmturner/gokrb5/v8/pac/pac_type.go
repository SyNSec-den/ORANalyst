//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:1
)

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/types"
	"github.com/jcmturner/rpc/v2/mstypes"
)

const (
	infoTypeKerbValidationInfo	uint32	= 1
	infoTypeCredentials		uint32	= 2
	infoTypePACServerSignatureData	uint32	= 6
	infoTypePACKDCSignatureData	uint32	= 7
	infoTypePACClientInfo		uint32	= 10
	infoTypeS4UDelegationInfo	uint32	= 11
	infoTypeUPNDNSInfo		uint32	= 12
	infoTypePACClientClaimsInfo	uint32	= 13
	infoTypePACDeviceInfo		uint32	= 14
	infoTypePACDeviceClaimsInfo	uint32	= 15
)

// PACType implements: https://msdn.microsoft.com/en-us/library/cc237950.aspx
type PACType struct {
	CBuffers		uint32
	Version			uint32
	Buffers			[]InfoBuffer
	Data			[]byte
	KerbValidationInfo	*KerbValidationInfo
	CredentialsInfo		*CredentialsInfo
	ServerChecksum		*SignatureData
	KDCChecksum		*SignatureData
	ClientInfo		*ClientInfo
	S4UDelegationInfo	*S4UDelegationInfo
	UPNDNSInfo		*UPNDNSInfo
	ClientClaimsInfo	*ClientClaimsInfo
	DeviceInfo		*DeviceInfo
	DeviceClaimsInfo	*DeviceClaimsInfo
	ZeroSigData		[]byte
}

// InfoBuffer implements the PAC Info Buffer: https://msdn.microsoft.com/en-us/library/cc237954.aspx
type InfoBuffer struct {
	ULType		uint32	// A 32-bit unsigned integer in little-endian format that describes the type of data present in the buffer contained at Offset.
	CBBufferSize	uint32	// A 32-bit unsigned integer in little-endian format that contains the size, in bytes, of the buffer in the PAC located at Offset.
	Offset		uint64	// A 64-bit unsigned integer in little-endian format that contains the offset to the beginning of the buffer, in bytes, from the beginning of the PACTYPE structure. The data offset MUST be a multiple of eight. The following sections specify the format of each type of element.
}

// Unmarshal bytes into the PACType struct
func (pac *PACType) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:55
	_go_fuzz_dep_.CoverTab[87481]++
												pac.Data = b
												zb := make([]byte, len(b), len(b))
												copy(zb, b)
												pac.ZeroSigData = zb
												r := mstypes.NewReader(bytes.NewReader(b))
												pac.CBuffers, err = r.Uint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:62
		_go_fuzz_dep_.CoverTab[87485]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:63
		// _ = "end of CoverTab[87485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:64
		_go_fuzz_dep_.CoverTab[87486]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:64
		// _ = "end of CoverTab[87486]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:64
	// _ = "end of CoverTab[87481]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:64
	_go_fuzz_dep_.CoverTab[87482]++
												pac.Version, err = r.Uint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:66
		_go_fuzz_dep_.CoverTab[87487]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:67
		// _ = "end of CoverTab[87487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:68
		_go_fuzz_dep_.CoverTab[87488]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:68
		// _ = "end of CoverTab[87488]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:68
	// _ = "end of CoverTab[87482]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:68
	_go_fuzz_dep_.CoverTab[87483]++
												buf := make([]InfoBuffer, pac.CBuffers, pac.CBuffers)
												for i := range buf {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:70
		_go_fuzz_dep_.CoverTab[87489]++
													buf[i].ULType, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:72
			_go_fuzz_dep_.CoverTab[87492]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:73
			// _ = "end of CoverTab[87492]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:74
			_go_fuzz_dep_.CoverTab[87493]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:74
			// _ = "end of CoverTab[87493]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:74
		// _ = "end of CoverTab[87489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:74
		_go_fuzz_dep_.CoverTab[87490]++
													buf[i].CBBufferSize, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:76
			_go_fuzz_dep_.CoverTab[87494]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:77
			// _ = "end of CoverTab[87494]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:78
			_go_fuzz_dep_.CoverTab[87495]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:78
			// _ = "end of CoverTab[87495]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:78
		// _ = "end of CoverTab[87490]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:78
		_go_fuzz_dep_.CoverTab[87491]++
													buf[i].Offset, err = r.Uint64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:80
			_go_fuzz_dep_.CoverTab[87496]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:81
			// _ = "end of CoverTab[87496]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:82
			_go_fuzz_dep_.CoverTab[87497]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:82
			// _ = "end of CoverTab[87497]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:82
		// _ = "end of CoverTab[87491]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:83
	// _ = "end of CoverTab[87483]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:83
	_go_fuzz_dep_.CoverTab[87484]++
												pac.Buffers = buf
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:85
	// _ = "end of CoverTab[87484]"
}

// ProcessPACInfoBuffers processes the PAC Info Buffers.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:88
// https://msdn.microsoft.com/en-us/library/cc237954.aspx
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:90
func (pac *PACType) ProcessPACInfoBuffers(key types.EncryptionKey, l *log.Logger) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:90
	_go_fuzz_dep_.CoverTab[87498]++
												for _, buf := range pac.Buffers {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:91
		_go_fuzz_dep_.CoverTab[87501]++
													p := make([]byte, buf.CBBufferSize, buf.CBBufferSize)
													copy(p, pac.Data[int(buf.Offset):int(buf.Offset)+int(buf.CBBufferSize)])
													switch buf.ULType {
		case infoTypeKerbValidationInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:95
			_go_fuzz_dep_.CoverTab[87502]++
														if pac.KerbValidationInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:96
				_go_fuzz_dep_.CoverTab[87531]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:98
				// _ = "end of CoverTab[87531]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:99
				_go_fuzz_dep_.CoverTab[87532]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:99
				// _ = "end of CoverTab[87532]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:99
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:99
			// _ = "end of CoverTab[87502]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:99
			_go_fuzz_dep_.CoverTab[87503]++
														var k KerbValidationInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:102
				_go_fuzz_dep_.CoverTab[87533]++
															return fmt.Errorf("error processing KerbValidationInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:103
				// _ = "end of CoverTab[87533]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:104
				_go_fuzz_dep_.CoverTab[87534]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:104
				// _ = "end of CoverTab[87534]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:104
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:104
			// _ = "end of CoverTab[87503]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:104
			_go_fuzz_dep_.CoverTab[87504]++
														pac.KerbValidationInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:105
			// _ = "end of CoverTab[87504]"
		case infoTypeCredentials:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:106
			_go_fuzz_dep_.CoverTab[87505]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:110
			continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:110
			// _ = "end of CoverTab[87505]"

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:121
		case infoTypePACServerSignatureData:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:121
			_go_fuzz_dep_.CoverTab[87506]++
														if pac.ServerChecksum != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:122
				_go_fuzz_dep_.CoverTab[87535]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:124
				// _ = "end of CoverTab[87535]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:125
				_go_fuzz_dep_.CoverTab[87536]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:125
				// _ = "end of CoverTab[87536]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:125
			// _ = "end of CoverTab[87506]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:125
			_go_fuzz_dep_.CoverTab[87507]++
														var k SignatureData
														zb, err := k.Unmarshal(p)
														copy(pac.ZeroSigData[int(buf.Offset):int(buf.Offset)+int(buf.CBBufferSize)], zb)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:129
				_go_fuzz_dep_.CoverTab[87537]++
															return fmt.Errorf("error processing ServerChecksum: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:130
				// _ = "end of CoverTab[87537]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:131
				_go_fuzz_dep_.CoverTab[87538]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:131
				// _ = "end of CoverTab[87538]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:131
			// _ = "end of CoverTab[87507]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:131
			_go_fuzz_dep_.CoverTab[87508]++
														pac.ServerChecksum = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:132
			// _ = "end of CoverTab[87508]"
		case infoTypePACKDCSignatureData:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:133
			_go_fuzz_dep_.CoverTab[87509]++
														if pac.KDCChecksum != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:134
				_go_fuzz_dep_.CoverTab[87539]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:136
				// _ = "end of CoverTab[87539]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:137
				_go_fuzz_dep_.CoverTab[87540]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:137
				// _ = "end of CoverTab[87540]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:137
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:137
			// _ = "end of CoverTab[87509]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:137
			_go_fuzz_dep_.CoverTab[87510]++
														var k SignatureData
														zb, err := k.Unmarshal(p)
														copy(pac.ZeroSigData[int(buf.Offset):int(buf.Offset)+int(buf.CBBufferSize)], zb)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:141
				_go_fuzz_dep_.CoverTab[87541]++
															return fmt.Errorf("error processing KDCChecksum: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:142
				// _ = "end of CoverTab[87541]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:143
				_go_fuzz_dep_.CoverTab[87542]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:143
				// _ = "end of CoverTab[87542]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:143
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:143
			// _ = "end of CoverTab[87510]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:143
			_go_fuzz_dep_.CoverTab[87511]++
														pac.KDCChecksum = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:144
			// _ = "end of CoverTab[87511]"
		case infoTypePACClientInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:145
			_go_fuzz_dep_.CoverTab[87512]++
														if pac.ClientInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:146
				_go_fuzz_dep_.CoverTab[87543]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:148
				// _ = "end of CoverTab[87543]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:149
				_go_fuzz_dep_.CoverTab[87544]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:149
				// _ = "end of CoverTab[87544]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:149
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:149
			// _ = "end of CoverTab[87512]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:149
			_go_fuzz_dep_.CoverTab[87513]++
														var k ClientInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:152
				_go_fuzz_dep_.CoverTab[87545]++
															return fmt.Errorf("error processing ClientInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:153
				// _ = "end of CoverTab[87545]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:154
				_go_fuzz_dep_.CoverTab[87546]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:154
				// _ = "end of CoverTab[87546]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:154
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:154
			// _ = "end of CoverTab[87513]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:154
			_go_fuzz_dep_.CoverTab[87514]++
														pac.ClientInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:155
			// _ = "end of CoverTab[87514]"
		case infoTypeS4UDelegationInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:156
			_go_fuzz_dep_.CoverTab[87515]++
														if pac.S4UDelegationInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:157
				_go_fuzz_dep_.CoverTab[87547]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:159
				// _ = "end of CoverTab[87547]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:160
				_go_fuzz_dep_.CoverTab[87548]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:160
				// _ = "end of CoverTab[87548]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:160
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:160
			// _ = "end of CoverTab[87515]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:160
			_go_fuzz_dep_.CoverTab[87516]++
														var k S4UDelegationInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:163
				_go_fuzz_dep_.CoverTab[87549]++
															l.Printf("could not process S4U_DelegationInfo: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:165
				// _ = "end of CoverTab[87549]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:166
				_go_fuzz_dep_.CoverTab[87550]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:166
				// _ = "end of CoverTab[87550]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:166
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:166
			// _ = "end of CoverTab[87516]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:166
			_go_fuzz_dep_.CoverTab[87517]++
														pac.S4UDelegationInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:167
			// _ = "end of CoverTab[87517]"
		case infoTypeUPNDNSInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:168
			_go_fuzz_dep_.CoverTab[87518]++
														if pac.UPNDNSInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:169
				_go_fuzz_dep_.CoverTab[87551]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:171
				// _ = "end of CoverTab[87551]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:172
				_go_fuzz_dep_.CoverTab[87552]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:172
				// _ = "end of CoverTab[87552]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:172
			// _ = "end of CoverTab[87518]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:172
			_go_fuzz_dep_.CoverTab[87519]++
														var k UPNDNSInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:175
				_go_fuzz_dep_.CoverTab[87553]++
															l.Printf("could not process UPN_DNSInfo: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:177
				// _ = "end of CoverTab[87553]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:178
				_go_fuzz_dep_.CoverTab[87554]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:178
				// _ = "end of CoverTab[87554]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:178
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:178
			// _ = "end of CoverTab[87519]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:178
			_go_fuzz_dep_.CoverTab[87520]++
														pac.UPNDNSInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:179
			// _ = "end of CoverTab[87520]"
		case infoTypePACClientClaimsInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:180
			_go_fuzz_dep_.CoverTab[87521]++
														if pac.ClientClaimsInfo != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:181
				_go_fuzz_dep_.CoverTab[87555]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:181
				return len(p) < 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:181
				// _ = "end of CoverTab[87555]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:181
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:181
				_go_fuzz_dep_.CoverTab[87556]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:183
				// _ = "end of CoverTab[87556]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:184
				_go_fuzz_dep_.CoverTab[87557]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:184
				// _ = "end of CoverTab[87557]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:184
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:184
			// _ = "end of CoverTab[87521]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:184
			_go_fuzz_dep_.CoverTab[87522]++
														var k ClientClaimsInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:187
				_go_fuzz_dep_.CoverTab[87558]++
															l.Printf("could not process ClientClaimsInfo: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:189
				// _ = "end of CoverTab[87558]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:190
				_go_fuzz_dep_.CoverTab[87559]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:190
				// _ = "end of CoverTab[87559]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:190
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:190
			// _ = "end of CoverTab[87522]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:190
			_go_fuzz_dep_.CoverTab[87523]++
														pac.ClientClaimsInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:191
			// _ = "end of CoverTab[87523]"
		case infoTypePACDeviceInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:192
			_go_fuzz_dep_.CoverTab[87524]++
														if pac.DeviceInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:193
				_go_fuzz_dep_.CoverTab[87560]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:195
				// _ = "end of CoverTab[87560]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:196
				_go_fuzz_dep_.CoverTab[87561]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:196
				// _ = "end of CoverTab[87561]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:196
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:196
			// _ = "end of CoverTab[87524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:196
			_go_fuzz_dep_.CoverTab[87525]++
														var k DeviceInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:199
				_go_fuzz_dep_.CoverTab[87562]++
															l.Printf("could not process DeviceInfo: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:201
				// _ = "end of CoverTab[87562]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:202
				_go_fuzz_dep_.CoverTab[87563]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:202
				// _ = "end of CoverTab[87563]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:202
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:202
			// _ = "end of CoverTab[87525]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:202
			_go_fuzz_dep_.CoverTab[87526]++
														pac.DeviceInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:203
			// _ = "end of CoverTab[87526]"
		case infoTypePACDeviceClaimsInfo:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:204
			_go_fuzz_dep_.CoverTab[87527]++
														if pac.DeviceClaimsInfo != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:205
				_go_fuzz_dep_.CoverTab[87564]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:207
				// _ = "end of CoverTab[87564]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:208
				_go_fuzz_dep_.CoverTab[87565]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:208
				// _ = "end of CoverTab[87565]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:208
			// _ = "end of CoverTab[87527]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:208
			_go_fuzz_dep_.CoverTab[87528]++
														var k DeviceClaimsInfo
														err := k.Unmarshal(p)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:211
				_go_fuzz_dep_.CoverTab[87566]++
															l.Printf("could not process DeviceClaimsInfo: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:213
				// _ = "end of CoverTab[87566]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:214
				_go_fuzz_dep_.CoverTab[87567]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:214
				// _ = "end of CoverTab[87567]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:214
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:214
			// _ = "end of CoverTab[87528]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:214
			_go_fuzz_dep_.CoverTab[87529]++
														pac.DeviceClaimsInfo = &k
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:215
			// _ = "end of CoverTab[87529]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:215
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:215
			_go_fuzz_dep_.CoverTab[87530]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:215
			// _ = "end of CoverTab[87530]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:216
		// _ = "end of CoverTab[87501]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:217
	// _ = "end of CoverTab[87498]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:217
	_go_fuzz_dep_.CoverTab[87499]++

												if ok, err := pac.verify(key); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:219
		_go_fuzz_dep_.CoverTab[87568]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:220
		// _ = "end of CoverTab[87568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:221
		_go_fuzz_dep_.CoverTab[87569]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:221
		// _ = "end of CoverTab[87569]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:221
	// _ = "end of CoverTab[87499]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:221
	_go_fuzz_dep_.CoverTab[87500]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:223
	// _ = "end of CoverTab[87500]"
}

func (pac *PACType) verify(key types.EncryptionKey) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:226
	_go_fuzz_dep_.CoverTab[87570]++
												if pac.KerbValidationInfo == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:227
		_go_fuzz_dep_.CoverTab[87577]++
													return false, errors.New("PAC Info Buffers does not contain a KerbValidationInfo")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:228
		// _ = "end of CoverTab[87577]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:229
		_go_fuzz_dep_.CoverTab[87578]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:229
		// _ = "end of CoverTab[87578]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:229
	// _ = "end of CoverTab[87570]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:229
	_go_fuzz_dep_.CoverTab[87571]++
												if pac.ServerChecksum == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:230
		_go_fuzz_dep_.CoverTab[87579]++
													return false, errors.New("PAC Info Buffers does not contain a ServerChecksum")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:231
		// _ = "end of CoverTab[87579]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:232
		_go_fuzz_dep_.CoverTab[87580]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:232
		// _ = "end of CoverTab[87580]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:232
	// _ = "end of CoverTab[87571]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:232
	_go_fuzz_dep_.CoverTab[87572]++
												if pac.KDCChecksum == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:233
		_go_fuzz_dep_.CoverTab[87581]++
													return false, errors.New("PAC Info Buffers does not contain a KDCChecksum")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:234
		// _ = "end of CoverTab[87581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:235
		_go_fuzz_dep_.CoverTab[87582]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:235
		// _ = "end of CoverTab[87582]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:235
	// _ = "end of CoverTab[87572]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:235
	_go_fuzz_dep_.CoverTab[87573]++
												if pac.ClientInfo == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:236
		_go_fuzz_dep_.CoverTab[87583]++
													return false, errors.New("PAC Info Buffers does not contain a ClientInfo")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:237
		// _ = "end of CoverTab[87583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:238
		_go_fuzz_dep_.CoverTab[87584]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:238
		// _ = "end of CoverTab[87584]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:238
	// _ = "end of CoverTab[87573]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:238
	_go_fuzz_dep_.CoverTab[87574]++
												etype, err := crypto.GetChksumEtype(int32(pac.ServerChecksum.SignatureType))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:240
		_go_fuzz_dep_.CoverTab[87585]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:241
		// _ = "end of CoverTab[87585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:242
		_go_fuzz_dep_.CoverTab[87586]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:242
		// _ = "end of CoverTab[87586]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:242
	// _ = "end of CoverTab[87574]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:242
	_go_fuzz_dep_.CoverTab[87575]++
												if ok := etype.VerifyChecksum(key.KeyValue,
		pac.ZeroSigData,
		pac.ServerChecksum.Signature,
		keyusage.KERB_NON_KERB_CKSUM_SALT); !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:246
		_go_fuzz_dep_.CoverTab[87587]++
													return false, errors.New("PAC service checksum verification failed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:247
		// _ = "end of CoverTab[87587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:248
		_go_fuzz_dep_.CoverTab[87588]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:248
		// _ = "end of CoverTab[87588]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:248
	// _ = "end of CoverTab[87575]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:248
	_go_fuzz_dep_.CoverTab[87576]++

												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:250
	// _ = "end of CoverTab[87576]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:251
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/pac_type.go:251
var _ = _go_fuzz_dep_.CoverTab
