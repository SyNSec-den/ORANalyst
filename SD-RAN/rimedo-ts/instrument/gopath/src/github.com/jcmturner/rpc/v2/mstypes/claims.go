//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:1
)

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jcmturner/rpc/v2/ndr"
	"golang.org/x/net/http2/hpack"
)

// Compression format assigned numbers. https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-xca/a8b7cb0a-92a6-4187-a23b-5e14273b96f8
const (
	CompressionFormatNone		uint16	= 0
	CompressionFormatLZNT1		uint16	= 2	// LZNT1 aka ntfs compression
	CompressionFormatXPress		uint16	= 3	// plain LZ77
	CompressionFormatXPressHuff	uint16	= 4	// LZ77+Huffman - The Huffman variant of the XPRESS compression format uses LZ77-style dictionary compression combined with Huffman coding.
)

// ClaimsSourceTypeAD https://msdn.microsoft.com/en-us/library/hh553809.aspx
const ClaimsSourceTypeAD uint16 = 1

// Claim Type assigned numbers
const (
	ClaimTypeIDInt64	uint16	= 1
	ClaimTypeIDUInt64	uint16	= 2
	ClaimTypeIDString	uint16	= 3
	ClaimsTypeIDBoolean	uint16	= 6
)

// ClaimsBlob implements https://msdn.microsoft.com/en-us/library/hh554119.aspx
type ClaimsBlob struct {
	Size		uint32
	EncodedBlob	EncodedBlob
}

// EncodedBlob are the bytes of the encoded Claims
type EncodedBlob []byte

// Size returns the size of the bytes of the encoded Claims
func (b EncodedBlob) Size(c interface{}) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:42
	_go_fuzz_dep_.CoverTab[87334]++
												cb := c.(ClaimsBlob)
												return int(cb.Size)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:44
	// _ = "end of CoverTab[87334]"
}

// ClaimsSetMetadata implements https://msdn.microsoft.com/en-us/library/hh554073.aspx
type ClaimsSetMetadata struct {
	ClaimsSetSize			uint32
	ClaimsSetBytes			[]byte	`ndr:"pointer,conformant"`
	CompressionFormat		uint16	// Enum see constants for options
	UncompressedClaimsSetSize	uint32
	ReservedType			uint16
	ReservedFieldSize		uint32
	ReservedField			[]byte	`ndr:"pointer,conformant"`
}

// ClaimsSet reads the ClaimsSet type from the NDR encoded ClaimsSetBytes in the ClaimsSetMetadata
func (m *ClaimsSetMetadata) ClaimsSet() (c ClaimsSet, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:59
	_go_fuzz_dep_.CoverTab[87335]++
												if len(m.ClaimsSetBytes) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:60
		_go_fuzz_dep_.CoverTab[87338]++
													err = errors.New("no bytes available for ClaimsSet")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:62
		// _ = "end of CoverTab[87338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:63
		_go_fuzz_dep_.CoverTab[87339]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:63
		// _ = "end of CoverTab[87339]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:63
	// _ = "end of CoverTab[87335]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:63
	_go_fuzz_dep_.CoverTab[87336]++

												switch m.CompressionFormat {
	case CompressionFormatLZNT1:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:66
		_go_fuzz_dep_.CoverTab[87340]++
													s := hex.EncodeToString(m.ClaimsSetBytes)
													err = fmt.Errorf("ClaimsSet compressed, format LZNT1 not currently supported: %s", s)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:69
		// _ = "end of CoverTab[87340]"
	case CompressionFormatXPress:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:70
		_go_fuzz_dep_.CoverTab[87341]++
													s := hex.EncodeToString(m.ClaimsSetBytes)
													err = fmt.Errorf("ClaimsSet compressed, format XPress not currently supported: %s", s)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:73
		// _ = "end of CoverTab[87341]"
	case CompressionFormatXPressHuff:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:74
		_go_fuzz_dep_.CoverTab[87342]++
													var b []byte
													buff := bytes.NewBuffer(b)
													_, e := hpack.HuffmanDecode(buff, m.ClaimsSetBytes)
													if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:78
			_go_fuzz_dep_.CoverTab[87345]++
														err = fmt.Errorf("error deflating: %v", e)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:80
			// _ = "end of CoverTab[87345]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:81
			_go_fuzz_dep_.CoverTab[87346]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:81
			// _ = "end of CoverTab[87346]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:81
		// _ = "end of CoverTab[87342]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:81
		_go_fuzz_dep_.CoverTab[87343]++
													m.ClaimsSetBytes = buff.Bytes()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:82
		// _ = "end of CoverTab[87343]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:82
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:82
		_go_fuzz_dep_.CoverTab[87344]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:82
		// _ = "end of CoverTab[87344]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:83
	// _ = "end of CoverTab[87336]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:83
	_go_fuzz_dep_.CoverTab[87337]++
												dec := ndr.NewDecoder(bytes.NewReader(m.ClaimsSetBytes))
												err = dec.Decode(&c)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:86
	// _ = "end of CoverTab[87337]"
}

// ClaimsSet implements https://msdn.microsoft.com/en-us/library/hh554122.aspx
type ClaimsSet struct {
	ClaimsArrayCount	uint32
	ClaimsArrays		[]ClaimsArray	`ndr:"pointer,conformant"`
	ReservedType		uint16
	ReservedFieldSize	uint32
	ReservedField		[]byte	`ndr:"pointer,conformant"`
}

// ClaimsArray implements https://msdn.microsoft.com/en-us/library/hh536458.aspx
type ClaimsArray struct {
	ClaimsSourceType	uint16
	ClaimsCount		uint32
	ClaimEntries		[]ClaimEntry	`ndr:"pointer,conformant"`
}

// ClaimEntry is a NDR union that implements https://msdn.microsoft.com/en-us/library/hh536374.aspx
type ClaimEntry struct {
	ID		string			`ndr:"pointer,conformant,varying"`
	Type		uint16			`ndr:"unionTag"`
	TypeInt64	ClaimTypeInt64		`ndr:"unionField"`
	TypeUInt64	ClaimTypeUInt64		`ndr:"unionField"`
	TypeString	ClaimTypeString		`ndr:"unionField"`
	TypeBool	ClaimTypeBoolean	`ndr:"unionField"`
}

// SwitchFunc is the ClaimEntry union field selection function
func (u ClaimEntry) SwitchFunc(_ interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:116
	_go_fuzz_dep_.CoverTab[87347]++
												switch u.Type {
	case ClaimTypeIDInt64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:118
		_go_fuzz_dep_.CoverTab[87349]++
													return "TypeInt64"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:119
		// _ = "end of CoverTab[87349]"
	case ClaimTypeIDUInt64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:120
		_go_fuzz_dep_.CoverTab[87350]++
													return "TypeUInt64"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:121
		// _ = "end of CoverTab[87350]"
	case ClaimTypeIDString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:122
		_go_fuzz_dep_.CoverTab[87351]++
													return "TypeString"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:123
		// _ = "end of CoverTab[87351]"
	case ClaimsTypeIDBoolean:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:124
		_go_fuzz_dep_.CoverTab[87352]++
													return "TypeBool"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:125
		// _ = "end of CoverTab[87352]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:125
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:125
		_go_fuzz_dep_.CoverTab[87353]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:125
		// _ = "end of CoverTab[87353]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:126
	// _ = "end of CoverTab[87347]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:126
	_go_fuzz_dep_.CoverTab[87348]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:127
	// _ = "end of CoverTab[87348]"
}

// ClaimTypeInt64 is a claim of type int64
type ClaimTypeInt64 struct {
	ValueCount	uint32
	Value		[]int64	`ndr:"pointer,conformant"`
}

// ClaimTypeUInt64 is a claim of type uint64
type ClaimTypeUInt64 struct {
	ValueCount	uint32
	Value		[]uint64	`ndr:"pointer,conformant"`
}

// ClaimTypeString is a claim of type string
type ClaimTypeString struct {
	ValueCount	uint32
	Value		[]LPWSTR	`ndr:"pointer,conformant"`
}

// ClaimTypeBoolean is a claim of type bool
type ClaimTypeBoolean struct {
	ValueCount	uint32
	Value		[]bool	`ndr:"pointer,conformant"`
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:152
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/claims.go:152
var _ = _go_fuzz_dep_.CoverTab
