//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:1
)

import (
	"bytes"
	"encoding/binary"
	"math"
)

// Byte sizes of primitive types
const (
	SizeBool	= 1
	SizeChar	= 1
	SizeUint8	= 1
	SizeUint16	= 2
	SizeUint32	= 4
	SizeUint64	= 8
	SizeEnum	= 2
	SizeSingle	= 4
	SizeDouble	= 8
	SizePtr		= 4
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:54
// readBool reads a byte representing a boolean.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:54
// NDR represents a Boolean as one octet.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:54
// It represents a value of FALSE as a zero octet, an octet in which every bit is reset.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:54
// It represents a value of TRUE as a non-zero octet, an octet in which one or more bits are set.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:58
func (dec *Decoder) readBool() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:58
	_go_fuzz_dep_.CoverTab[87183]++
												i, err := dec.readUint8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:60
		_go_fuzz_dep_.CoverTab[87186]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:61
		// _ = "end of CoverTab[87186]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:62
		_go_fuzz_dep_.CoverTab[87187]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:62
		// _ = "end of CoverTab[87187]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:62
	// _ = "end of CoverTab[87183]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:62
	_go_fuzz_dep_.CoverTab[87184]++
												if i != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:63
		_go_fuzz_dep_.CoverTab[87188]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:64
		// _ = "end of CoverTab[87188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:65
		_go_fuzz_dep_.CoverTab[87189]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:65
		// _ = "end of CoverTab[87189]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:65
	// _ = "end of CoverTab[87184]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:65
	_go_fuzz_dep_.CoverTab[87185]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:66
	// _ = "end of CoverTab[87185]"
}

// readChar reads bytes representing a 8bit ASCII integer cast to a rune.
func (dec *Decoder) readChar() (rune, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:70
	_go_fuzz_dep_.CoverTab[87190]++
												var r rune
												a, err := dec.readUint8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:73
		_go_fuzz_dep_.CoverTab[87192]++
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:74
		// _ = "end of CoverTab[87192]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:75
		_go_fuzz_dep_.CoverTab[87193]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:75
		// _ = "end of CoverTab[87193]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:75
	// _ = "end of CoverTab[87190]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:75
	_go_fuzz_dep_.CoverTab[87191]++
												return rune(a), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:76
	// _ = "end of CoverTab[87191]"
}

// readUint8 reads bytes representing a 8bit unsigned integer.
func (dec *Decoder) readUint8() (uint8, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:80
	_go_fuzz_dep_.CoverTab[87194]++
												b, err := dec.r.ReadByte()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:82
		_go_fuzz_dep_.CoverTab[87196]++
													return uint8(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:83
		// _ = "end of CoverTab[87196]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:84
		_go_fuzz_dep_.CoverTab[87197]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:84
		// _ = "end of CoverTab[87197]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:84
	// _ = "end of CoverTab[87194]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:84
	_go_fuzz_dep_.CoverTab[87195]++
												return uint8(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:85
	// _ = "end of CoverTab[87195]"
}

// readUint16 reads bytes representing a 16bit unsigned integer.
func (dec *Decoder) readUint16() (uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:89
	_go_fuzz_dep_.CoverTab[87198]++
												dec.ensureAlignment(SizeUint16)
												b, err := dec.readBytes(SizeUint16)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:92
		_go_fuzz_dep_.CoverTab[87200]++
													return uint16(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:93
		// _ = "end of CoverTab[87200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:94
		_go_fuzz_dep_.CoverTab[87201]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:94
		// _ = "end of CoverTab[87201]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:94
	// _ = "end of CoverTab[87198]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:94
	_go_fuzz_dep_.CoverTab[87199]++
												return dec.ch.Endianness.Uint16(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:95
	// _ = "end of CoverTab[87199]"
}

// readUint32 reads bytes representing a 32bit unsigned integer.
func (dec *Decoder) readUint32() (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:99
	_go_fuzz_dep_.CoverTab[87202]++
												dec.ensureAlignment(SizeUint32)
												b, err := dec.readBytes(SizeUint32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:102
		_go_fuzz_dep_.CoverTab[87204]++
													return uint32(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:103
		// _ = "end of CoverTab[87204]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:104
		_go_fuzz_dep_.CoverTab[87205]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:104
		// _ = "end of CoverTab[87205]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:104
	// _ = "end of CoverTab[87202]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:104
	_go_fuzz_dep_.CoverTab[87203]++
												return dec.ch.Endianness.Uint32(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:105
	// _ = "end of CoverTab[87203]"
}

// readUint32 reads bytes representing a 32bit unsigned integer.
func (dec *Decoder) readUint64() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:109
	_go_fuzz_dep_.CoverTab[87206]++
												dec.ensureAlignment(SizeUint64)
												b, err := dec.readBytes(SizeUint64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:112
		_go_fuzz_dep_.CoverTab[87208]++
													return uint64(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:113
		// _ = "end of CoverTab[87208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:114
		_go_fuzz_dep_.CoverTab[87209]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:114
		// _ = "end of CoverTab[87209]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:114
	// _ = "end of CoverTab[87206]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:114
	_go_fuzz_dep_.CoverTab[87207]++
												return dec.ch.Endianness.Uint64(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:115
	// _ = "end of CoverTab[87207]"
}

func (dec *Decoder) readInt8() (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:118
	_go_fuzz_dep_.CoverTab[87210]++
												dec.ensureAlignment(SizeUint8)
												b, err := dec.readBytes(SizeUint8)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:121
		_go_fuzz_dep_.CoverTab[87213]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:122
		// _ = "end of CoverTab[87213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:123
		_go_fuzz_dep_.CoverTab[87214]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:123
		// _ = "end of CoverTab[87214]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:123
	// _ = "end of CoverTab[87210]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:123
	_go_fuzz_dep_.CoverTab[87211]++
												var i int8
												buf := bytes.NewReader(b)
												err = binary.Read(buf, dec.ch.Endianness, &i)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:127
		_go_fuzz_dep_.CoverTab[87215]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:128
		// _ = "end of CoverTab[87215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:129
		_go_fuzz_dep_.CoverTab[87216]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:129
		// _ = "end of CoverTab[87216]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:129
	// _ = "end of CoverTab[87211]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:129
	_go_fuzz_dep_.CoverTab[87212]++
												return i, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:130
	// _ = "end of CoverTab[87212]"
}

func (dec *Decoder) readInt16() (int16, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:133
	_go_fuzz_dep_.CoverTab[87217]++
												dec.ensureAlignment(SizeUint16)
												b, err := dec.readBytes(SizeUint16)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:136
		_go_fuzz_dep_.CoverTab[87220]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:137
		// _ = "end of CoverTab[87220]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:138
		_go_fuzz_dep_.CoverTab[87221]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:138
		// _ = "end of CoverTab[87221]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:138
	// _ = "end of CoverTab[87217]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:138
	_go_fuzz_dep_.CoverTab[87218]++
												var i int16
												buf := bytes.NewReader(b)
												err = binary.Read(buf, dec.ch.Endianness, &i)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:142
		_go_fuzz_dep_.CoverTab[87222]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:143
		// _ = "end of CoverTab[87222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:144
		_go_fuzz_dep_.CoverTab[87223]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:144
		// _ = "end of CoverTab[87223]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:144
	// _ = "end of CoverTab[87218]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:144
	_go_fuzz_dep_.CoverTab[87219]++
												return i, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:145
	// _ = "end of CoverTab[87219]"
}

func (dec *Decoder) readInt32() (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:148
	_go_fuzz_dep_.CoverTab[87224]++
												dec.ensureAlignment(SizeUint32)
												b, err := dec.readBytes(SizeUint32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:151
		_go_fuzz_dep_.CoverTab[87227]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:152
		// _ = "end of CoverTab[87227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:153
		_go_fuzz_dep_.CoverTab[87228]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:153
		// _ = "end of CoverTab[87228]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:153
	// _ = "end of CoverTab[87224]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:153
	_go_fuzz_dep_.CoverTab[87225]++
												var i int32
												buf := bytes.NewReader(b)
												err = binary.Read(buf, dec.ch.Endianness, &i)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:157
		_go_fuzz_dep_.CoverTab[87229]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:158
		// _ = "end of CoverTab[87229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:159
		_go_fuzz_dep_.CoverTab[87230]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:159
		// _ = "end of CoverTab[87230]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:159
	// _ = "end of CoverTab[87225]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:159
	_go_fuzz_dep_.CoverTab[87226]++
												return i, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:160
	// _ = "end of CoverTab[87226]"
}

func (dec *Decoder) readInt64() (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:163
	_go_fuzz_dep_.CoverTab[87231]++
												dec.ensureAlignment(SizeUint64)
												b, err := dec.readBytes(SizeUint64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:166
		_go_fuzz_dep_.CoverTab[87234]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:167
		// _ = "end of CoverTab[87234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:168
		_go_fuzz_dep_.CoverTab[87235]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:168
		// _ = "end of CoverTab[87235]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:168
	// _ = "end of CoverTab[87231]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:168
	_go_fuzz_dep_.CoverTab[87232]++
												var i int64
												buf := bytes.NewReader(b)
												err = binary.Read(buf, dec.ch.Endianness, &i)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:172
		_go_fuzz_dep_.CoverTab[87236]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:173
		// _ = "end of CoverTab[87236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:174
		_go_fuzz_dep_.CoverTab[87237]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:174
		// _ = "end of CoverTab[87237]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:174
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:174
	// _ = "end of CoverTab[87232]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:174
	_go_fuzz_dep_.CoverTab[87233]++
												return i, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:175
	// _ = "end of CoverTab[87233]"
}

// https://en.wikipedia.org/wiki/IEEE_754-1985
func (dec *Decoder) readFloat32() (f float32, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:179
	_go_fuzz_dep_.CoverTab[87238]++
												dec.ensureAlignment(SizeSingle)
												b, err := dec.readBytes(SizeSingle)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:182
		_go_fuzz_dep_.CoverTab[87240]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:183
		// _ = "end of CoverTab[87240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:184
		_go_fuzz_dep_.CoverTab[87241]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:184
		// _ = "end of CoverTab[87241]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:184
	// _ = "end of CoverTab[87238]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:184
	_go_fuzz_dep_.CoverTab[87239]++
												bits := dec.ch.Endianness.Uint32(b)
												f = math.Float32frombits(bits)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:187
	// _ = "end of CoverTab[87239]"
}

func (dec *Decoder) readFloat64() (f float64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:190
	_go_fuzz_dep_.CoverTab[87242]++
												dec.ensureAlignment(SizeDouble)
												b, err := dec.readBytes(SizeDouble)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:193
		_go_fuzz_dep_.CoverTab[87244]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:194
		// _ = "end of CoverTab[87244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:195
		_go_fuzz_dep_.CoverTab[87245]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:195
		// _ = "end of CoverTab[87245]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:195
	// _ = "end of CoverTab[87242]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:195
	_go_fuzz_dep_.CoverTab[87243]++
												bits := dec.ch.Endianness.Uint64(b)
												f = math.Float64frombits(bits)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:198
	// _ = "end of CoverTab[87243]"
}

// NDR enforces NDR alignment of primitive data; that is, any primitive of size n octets is aligned at a octet stream
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:201
// index that is a multiple of n. (In this version of NDR, n is one of {1, 2, 4, 8}.) An octet stream index indicates
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:201
// the number of an octet in an octet stream when octets are numbered, beginning with 0, from the first octet in the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:201
// stream. Where necessary, an alignment gap, consisting of octets of unspecified value, precedes the representation
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:201
// of a primitive. The gap is of the smallest size sufficient to align the primitive.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:206
func (dec *Decoder) ensureAlignment(n int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:206
	_go_fuzz_dep_.CoverTab[87246]++
												p := dec.size - dec.r.Buffered()
												if s := p % n; s != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:208
		_go_fuzz_dep_.CoverTab[87247]++
													dec.r.Discard(n - s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:209
		// _ = "end of CoverTab[87247]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:210
		_go_fuzz_dep_.CoverTab[87248]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:210
		// _ = "end of CoverTab[87248]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:210
	// _ = "end of CoverTab[87246]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:211
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/primitives.go:211
var _ = _go_fuzz_dep_.CoverTab
