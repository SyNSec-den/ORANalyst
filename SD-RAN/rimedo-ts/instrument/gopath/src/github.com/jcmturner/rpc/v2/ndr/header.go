//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:1
)

import (
	"encoding/binary"
	"fmt"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:26
const (
	protocolVersion		uint8	= 1
	commonHeaderBytes	uint16	= 8
	bigEndian			= 0
	littleEndian			= 1
	ascii			uint8	= 0
	ebcdic			uint8	= 1
	ieee			uint8	= 0
	vax			uint8	= 1
	cray			uint8	= 2
	ibm			uint8	= 3
)

// CommonHeader implements the NDR common header: https://msdn.microsoft.com/en-us/library/cc243889.aspx
type CommonHeader struct {
	Version			uint8
	Endianness		binary.ByteOrder
	CharacterEncoding	uint8
	FloatRepresentation	uint8
	HeaderLength		uint16
	Filler			[]byte
}

// PrivateHeader implements the NDR private header: https://msdn.microsoft.com/en-us/library/cc243919.aspx
type PrivateHeader struct {
	ObjectBufferLength	uint32
	Filler			[]byte
}

func (dec *Decoder) readCommonHeader() error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:55
	_go_fuzz_dep_.CoverTab[87129]++

											vb, err := dec.r.ReadByte()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:58
		_go_fuzz_dep_.CoverTab[87139]++
												return Malformed{EText: "could not read first byte of common header for version"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:59
		// _ = "end of CoverTab[87139]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:60
		_go_fuzz_dep_.CoverTab[87140]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:60
		// _ = "end of CoverTab[87140]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:60
	// _ = "end of CoverTab[87129]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:60
	_go_fuzz_dep_.CoverTab[87130]++
											dec.ch.Version = uint8(vb)
											if dec.ch.Version != protocolVersion {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:62
		_go_fuzz_dep_.CoverTab[87141]++
												return Malformed{EText: fmt.Sprintf("byte stream does not indicate a RPC Type serialization of version %v", protocolVersion)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:63
		// _ = "end of CoverTab[87141]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:64
		_go_fuzz_dep_.CoverTab[87142]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:64
		// _ = "end of CoverTab[87142]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:64
	// _ = "end of CoverTab[87130]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:64
	_go_fuzz_dep_.CoverTab[87131]++

											eb, err := dec.r.ReadByte()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:67
		_go_fuzz_dep_.CoverTab[87143]++
												return Malformed{EText: "could not read second byte of common header for endianness"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:68
		// _ = "end of CoverTab[87143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:69
		_go_fuzz_dep_.CoverTab[87144]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:69
		// _ = "end of CoverTab[87144]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:69
	// _ = "end of CoverTab[87131]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:69
	_go_fuzz_dep_.CoverTab[87132]++
											endian := int(eb >> 4 & 0xF)
											if endian != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:71
		_go_fuzz_dep_.CoverTab[87145]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:71
		return endian != 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:71
		// _ = "end of CoverTab[87145]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:71
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:71
		_go_fuzz_dep_.CoverTab[87146]++
												return Malformed{EText: "common header does not indicate a valid endianness"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:72
		// _ = "end of CoverTab[87146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:73
		_go_fuzz_dep_.CoverTab[87147]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:73
		// _ = "end of CoverTab[87147]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:73
	// _ = "end of CoverTab[87132]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:73
	_go_fuzz_dep_.CoverTab[87133]++
											dec.ch.CharacterEncoding = uint8(vb & 0xF)
											if dec.ch.CharacterEncoding != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:75
		_go_fuzz_dep_.CoverTab[87148]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:75
		return dec.ch.CharacterEncoding != 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:75
		// _ = "end of CoverTab[87148]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:75
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:75
		_go_fuzz_dep_.CoverTab[87149]++
												return Malformed{EText: "common header does not indicate a valid character encoding"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:76
		// _ = "end of CoverTab[87149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:77
		_go_fuzz_dep_.CoverTab[87150]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:77
		// _ = "end of CoverTab[87150]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:77
	// _ = "end of CoverTab[87133]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:77
	_go_fuzz_dep_.CoverTab[87134]++
											switch endian {
	case littleEndian:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:79
		_go_fuzz_dep_.CoverTab[87151]++
												dec.ch.Endianness = binary.LittleEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:80
		// _ = "end of CoverTab[87151]"
	case bigEndian:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:81
		_go_fuzz_dep_.CoverTab[87152]++
												dec.ch.Endianness = binary.BigEndian
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:82
		// _ = "end of CoverTab[87152]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:82
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:82
		_go_fuzz_dep_.CoverTab[87153]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:82
		// _ = "end of CoverTab[87153]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:83
	// _ = "end of CoverTab[87134]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:83
	_go_fuzz_dep_.CoverTab[87135]++

											lb, err := dec.readBytes(2)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:86
		_go_fuzz_dep_.CoverTab[87154]++
												return Malformed{EText: fmt.Sprintf("could not read common header length: %v", err)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:87
		// _ = "end of CoverTab[87154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:88
		_go_fuzz_dep_.CoverTab[87155]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:88
		// _ = "end of CoverTab[87155]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:88
	// _ = "end of CoverTab[87135]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:88
	_go_fuzz_dep_.CoverTab[87136]++
											dec.ch.HeaderLength = dec.ch.Endianness.Uint16(lb)
											if dec.ch.HeaderLength != commonHeaderBytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:90
		_go_fuzz_dep_.CoverTab[87156]++
												return Malformed{EText: "common header does not indicate a valid length"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:91
		// _ = "end of CoverTab[87156]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:92
		_go_fuzz_dep_.CoverTab[87157]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:92
		// _ = "end of CoverTab[87157]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:92
	// _ = "end of CoverTab[87136]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:92
	_go_fuzz_dep_.CoverTab[87137]++

											dec.ch.Filler, err = dec.readBytes(4)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:95
		_go_fuzz_dep_.CoverTab[87158]++
												return Malformed{EText: fmt.Sprintf("could not read common header filler: %v", err)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:96
		// _ = "end of CoverTab[87158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:97
		_go_fuzz_dep_.CoverTab[87159]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:97
		// _ = "end of CoverTab[87159]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:97
	// _ = "end of CoverTab[87137]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:97
	_go_fuzz_dep_.CoverTab[87138]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:98
	// _ = "end of CoverTab[87138]"
}

func (dec *Decoder) readPrivateHeader() error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:101
	_go_fuzz_dep_.CoverTab[87160]++

											err := binary.Read(dec.r, dec.ch.Endianness, &dec.ph.ObjectBufferLength)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:104
		_go_fuzz_dep_.CoverTab[87164]++
												return Malformed{EText: "could not read private header object buffer length"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:105
		// _ = "end of CoverTab[87164]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:106
		_go_fuzz_dep_.CoverTab[87165]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:106
		// _ = "end of CoverTab[87165]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:106
	// _ = "end of CoverTab[87160]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:106
	_go_fuzz_dep_.CoverTab[87161]++
											if dec.ph.ObjectBufferLength%8 != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:107
		_go_fuzz_dep_.CoverTab[87166]++
												return Malformed{EText: "object buffer length not a multiple of 8"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:108
		// _ = "end of CoverTab[87166]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:109
		_go_fuzz_dep_.CoverTab[87167]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:109
		// _ = "end of CoverTab[87167]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:109
	// _ = "end of CoverTab[87161]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:109
	_go_fuzz_dep_.CoverTab[87162]++

											dec.ph.Filler, err = dec.readBytes(4)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:112
		_go_fuzz_dep_.CoverTab[87168]++
												return Malformed{EText: fmt.Sprintf("could not read private header filler: %v", err)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:113
		// _ = "end of CoverTab[87168]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:114
		_go_fuzz_dep_.CoverTab[87169]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:114
		// _ = "end of CoverTab[87169]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:114
	// _ = "end of CoverTab[87162]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:114
	_go_fuzz_dep_.CoverTab[87163]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:115
	// _ = "end of CoverTab[87163]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:116
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/header.go:116
var _ = _go_fuzz_dep_.CoverTab
