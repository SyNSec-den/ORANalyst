//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:1
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/klauspost/compress/huff0"
)

type dict struct {
	id	uint32

	litEnc			*huff0.Scratch
	llDec, ofDec, mlDec	sequenceDec
	//llEnc, ofEnc, mlEnc []*fseEncoder
	offsets	[3]int
	content	[]byte
}

var dictMagic = [4]byte{0x37, 0xa4, 0x30, 0xec}

// ID returns the dictionary id or 0 if d is nil.
func (d *dict) ID() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:26
	_go_fuzz_dep_.CoverTab[91954]++
												if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:27
		_go_fuzz_dep_.CoverTab[91956]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:28
		// _ = "end of CoverTab[91956]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:29
		_go_fuzz_dep_.CoverTab[91957]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:29
		// _ = "end of CoverTab[91957]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:29
	// _ = "end of CoverTab[91954]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:29
	_go_fuzz_dep_.CoverTab[91955]++
												return d.id
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:30
	// _ = "end of CoverTab[91955]"
}

// DictContentSize returns the dictionary content size or 0 if d is nil.
func (d *dict) DictContentSize() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:34
	_go_fuzz_dep_.CoverTab[91958]++
												if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:35
		_go_fuzz_dep_.CoverTab[91960]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:36
		// _ = "end of CoverTab[91960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:37
		_go_fuzz_dep_.CoverTab[91961]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:37
		// _ = "end of CoverTab[91961]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:37
	// _ = "end of CoverTab[91958]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:37
	_go_fuzz_dep_.CoverTab[91959]++
												return len(d.content)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:38
	// _ = "end of CoverTab[91959]"
}

// Load a dictionary as described in
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:41
// https://github.com/facebook/zstd/blob/master/doc/zstd_compression_format.md#dictionary-format
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:43
func loadDict(b []byte) (*dict, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:43
	_go_fuzz_dep_.CoverTab[91962]++

												if len(b) <= 8+(3*4) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:45
		_go_fuzz_dep_.CoverTab[91974]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:46
		// _ = "end of CoverTab[91974]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:47
		_go_fuzz_dep_.CoverTab[91975]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:47
		// _ = "end of CoverTab[91975]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:47
	// _ = "end of CoverTab[91962]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:47
	_go_fuzz_dep_.CoverTab[91963]++
												d := dict{
		llDec:	sequenceDec{fse: &fseDecoder{}},
		ofDec:	sequenceDec{fse: &fseDecoder{}},
		mlDec:	sequenceDec{fse: &fseDecoder{}},
	}
	if !bytes.Equal(b[:4], dictMagic[:]) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:53
		_go_fuzz_dep_.CoverTab[91976]++
													return nil, ErrMagicMismatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:54
		// _ = "end of CoverTab[91976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:55
		_go_fuzz_dep_.CoverTab[91977]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:55
		// _ = "end of CoverTab[91977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:55
	// _ = "end of CoverTab[91963]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:55
	_go_fuzz_dep_.CoverTab[91964]++
												d.id = binary.LittleEndian.Uint32(b[4:8])
												if d.id == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:57
		_go_fuzz_dep_.CoverTab[91978]++
													return nil, errors.New("dictionaries cannot have ID 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:58
		// _ = "end of CoverTab[91978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:59
		_go_fuzz_dep_.CoverTab[91979]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:59
		// _ = "end of CoverTab[91979]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:59
	// _ = "end of CoverTab[91964]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:59
	_go_fuzz_dep_.CoverTab[91965]++

	// Read literal table
	var err error
	d.litEnc, b, err = huff0.ReadTable(b[8:], nil)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:64
		_go_fuzz_dep_.CoverTab[91980]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:65
		// _ = "end of CoverTab[91980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:66
		_go_fuzz_dep_.CoverTab[91981]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:66
		// _ = "end of CoverTab[91981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:66
	// _ = "end of CoverTab[91965]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:66
	_go_fuzz_dep_.CoverTab[91966]++
												d.litEnc.Reuse = huff0.ReusePolicyMust

												br := byteReader{
		b:	b,
		off:	0,
	}
	readDec := func(i tableIndex, dec *fseDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:73
		_go_fuzz_dep_.CoverTab[91982]++
													if err := dec.readNCount(&br, uint16(maxTableSymbol[i])); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:74
			_go_fuzz_dep_.CoverTab[91987]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:75
			// _ = "end of CoverTab[91987]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:76
			_go_fuzz_dep_.CoverTab[91988]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:76
			// _ = "end of CoverTab[91988]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:76
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:76
		// _ = "end of CoverTab[91982]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:76
		_go_fuzz_dep_.CoverTab[91983]++
													if br.overread() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:77
			_go_fuzz_dep_.CoverTab[91989]++
														return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:78
			// _ = "end of CoverTab[91989]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:79
			_go_fuzz_dep_.CoverTab[91990]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:79
			// _ = "end of CoverTab[91990]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:79
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:79
		// _ = "end of CoverTab[91983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:79
		_go_fuzz_dep_.CoverTab[91984]++
													err = dec.transform(symbolTableX[i])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:81
			_go_fuzz_dep_.CoverTab[91991]++
														println("Transform table error:", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:83
			// _ = "end of CoverTab[91991]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:84
			_go_fuzz_dep_.CoverTab[91992]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:84
			// _ = "end of CoverTab[91992]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:84
		// _ = "end of CoverTab[91984]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:84
		_go_fuzz_dep_.CoverTab[91985]++
													if debugDecoder || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:85
			_go_fuzz_dep_.CoverTab[91993]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:85
			return debugEncoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:85
			// _ = "end of CoverTab[91993]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:85
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:85
			_go_fuzz_dep_.CoverTab[91994]++
														println("Read table ok", "symbolLen:", dec.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:86
			// _ = "end of CoverTab[91994]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:87
			_go_fuzz_dep_.CoverTab[91995]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:87
			// _ = "end of CoverTab[91995]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:87
		// _ = "end of CoverTab[91985]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:87
		_go_fuzz_dep_.CoverTab[91986]++

													dec.preDefined = true
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:90
		// _ = "end of CoverTab[91986]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:91
	// _ = "end of CoverTab[91966]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:91
	_go_fuzz_dep_.CoverTab[91967]++

												if err := readDec(tableOffsets, d.ofDec.fse); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:93
		_go_fuzz_dep_.CoverTab[91996]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:94
		// _ = "end of CoverTab[91996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:95
		_go_fuzz_dep_.CoverTab[91997]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:95
		// _ = "end of CoverTab[91997]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:95
	// _ = "end of CoverTab[91967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:95
	_go_fuzz_dep_.CoverTab[91968]++
												if err := readDec(tableMatchLengths, d.mlDec.fse); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:96
		_go_fuzz_dep_.CoverTab[91998]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:97
		// _ = "end of CoverTab[91998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:98
		_go_fuzz_dep_.CoverTab[91999]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:98
		// _ = "end of CoverTab[91999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:98
	// _ = "end of CoverTab[91968]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:98
	_go_fuzz_dep_.CoverTab[91969]++
												if err := readDec(tableLiteralLengths, d.llDec.fse); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:99
		_go_fuzz_dep_.CoverTab[92000]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:100
		// _ = "end of CoverTab[92000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:101
		_go_fuzz_dep_.CoverTab[92001]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:101
		// _ = "end of CoverTab[92001]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:101
	// _ = "end of CoverTab[91969]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:101
	_go_fuzz_dep_.CoverTab[91970]++
												if br.remain() < 12 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:102
		_go_fuzz_dep_.CoverTab[92002]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:103
		// _ = "end of CoverTab[92002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:104
		_go_fuzz_dep_.CoverTab[92003]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:104
		// _ = "end of CoverTab[92003]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:104
	// _ = "end of CoverTab[91970]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:104
	_go_fuzz_dep_.CoverTab[91971]++

												d.offsets[0] = int(br.Uint32())
												br.advance(4)
												d.offsets[1] = int(br.Uint32())
												br.advance(4)
												d.offsets[2] = int(br.Uint32())
												br.advance(4)
												if d.offsets[0] <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		_go_fuzz_dep_.CoverTab[92004]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		return d.offsets[1] <= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		// _ = "end of CoverTab[92004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		_go_fuzz_dep_.CoverTab[92005]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		return d.offsets[2] <= 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		// _ = "end of CoverTab[92005]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:112
		_go_fuzz_dep_.CoverTab[92006]++
													return nil, errors.New("invalid offset in dictionary")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:113
		// _ = "end of CoverTab[92006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:114
		_go_fuzz_dep_.CoverTab[92007]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:114
		// _ = "end of CoverTab[92007]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:114
	// _ = "end of CoverTab[91971]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:114
	_go_fuzz_dep_.CoverTab[91972]++
												d.content = make([]byte, br.remain())
												copy(d.content, br.unread())
												if d.offsets[0] > len(d.content) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		_go_fuzz_dep_.CoverTab[92008]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		return d.offsets[1] > len(d.content)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		// _ = "end of CoverTab[92008]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		_go_fuzz_dep_.CoverTab[92009]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		return d.offsets[2] > len(d.content)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		// _ = "end of CoverTab[92009]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:117
		_go_fuzz_dep_.CoverTab[92010]++
													return nil, fmt.Errorf("initial offset bigger than dictionary content size %d, offsets: %v", len(d.content), d.offsets)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:118
		// _ = "end of CoverTab[92010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:119
		_go_fuzz_dep_.CoverTab[92011]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:119
		// _ = "end of CoverTab[92011]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:119
	// _ = "end of CoverTab[91972]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:119
	_go_fuzz_dep_.CoverTab[91973]++

												return &d, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:121
	// _ = "end of CoverTab[91973]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/dict.go:122
var _ = _go_fuzz_dep_.CoverTab
