// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:5
// Package asn1 implements parsing of DER-encoded ASN.1 data structures,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:5
// as defined in ITU-T Rec X.690.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:5
// See also “A Layman's Guide to a Subset of ASN.1, BER, and DER,”
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:5
// http://luca.ntop.org/Teaching/Appunti/asn1.html.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
package asn1

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:10
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:22
import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

// A StructuralError suggests that the ASN.1 data is valid, but the Go type
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:32
// which is receiving it doesn't match.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:34
type StructuralError struct {
	Msg string
}

func (e StructuralError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:38
	_go_fuzz_dep_.CoverTab[82293]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:38
	return "asn1: structure error: " + e.Msg
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:38
	// _ = "end of CoverTab[82293]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:38
}

// A SyntaxError suggests that the ASN.1 data is invalid.
type SyntaxError struct {
	Msg string
}

func (e SyntaxError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:45
	_go_fuzz_dep_.CoverTab[82294]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:45
	return "asn1: syntax error: " + e.Msg
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:45
	// _ = "end of CoverTab[82294]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:45
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:51
func parseBool(bytes []byte) (ret bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:51
	_go_fuzz_dep_.CoverTab[82295]++
												if len(bytes) != 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:52
		_go_fuzz_dep_.CoverTab[82298]++
													err = SyntaxError{"invalid boolean"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:54
		// _ = "end of CoverTab[82298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:55
		_go_fuzz_dep_.CoverTab[82299]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:55
		// _ = "end of CoverTab[82299]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:55
	// _ = "end of CoverTab[82295]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:55
	_go_fuzz_dep_.CoverTab[82296]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:60
	switch bytes[0] {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:61
		_go_fuzz_dep_.CoverTab[82300]++
													ret = false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:62
		// _ = "end of CoverTab[82300]"
	case 0xff:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:63
		_go_fuzz_dep_.CoverTab[82301]++
													ret = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:64
		// _ = "end of CoverTab[82301]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:65
		_go_fuzz_dep_.CoverTab[82302]++
													err = SyntaxError{"invalid boolean"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:66
		// _ = "end of CoverTab[82302]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:67
	// _ = "end of CoverTab[82296]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:67
	_go_fuzz_dep_.CoverTab[82297]++

												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:69
	// _ = "end of CoverTab[82297]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:74
// checkInteger returns nil if the given bytes are a valid DER-encoded
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:74
// INTEGER and an error otherwise.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:76
func checkInteger(bytes []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:76
	_go_fuzz_dep_.CoverTab[82303]++
												if len(bytes) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:77
		_go_fuzz_dep_.CoverTab[82307]++
													return StructuralError{"empty integer"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:78
		// _ = "end of CoverTab[82307]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:79
		_go_fuzz_dep_.CoverTab[82308]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:79
		// _ = "end of CoverTab[82308]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:79
	// _ = "end of CoverTab[82303]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:79
	_go_fuzz_dep_.CoverTab[82304]++
												if len(bytes) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:80
		_go_fuzz_dep_.CoverTab[82309]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:81
		// _ = "end of CoverTab[82309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:82
		_go_fuzz_dep_.CoverTab[82310]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:82
		// _ = "end of CoverTab[82310]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:82
	// _ = "end of CoverTab[82304]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:82
	_go_fuzz_dep_.CoverTab[82305]++
												if (bytes[0] == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		_go_fuzz_dep_.CoverTab[82311]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		return bytes[1]&0x80 == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		// _ = "end of CoverTab[82311]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		_go_fuzz_dep_.CoverTab[82312]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		return (bytes[0] == 0xff && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
			_go_fuzz_dep_.CoverTab[82313]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
			return bytes[1]&0x80 == 0x80
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
			// _ = "end of CoverTab[82313]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		// _ = "end of CoverTab[82312]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:83
		_go_fuzz_dep_.CoverTab[82314]++
													return StructuralError{"integer not minimally-encoded"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:84
		// _ = "end of CoverTab[82314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:85
		_go_fuzz_dep_.CoverTab[82315]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:85
		// _ = "end of CoverTab[82315]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:85
	// _ = "end of CoverTab[82305]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:85
	_go_fuzz_dep_.CoverTab[82306]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:86
	// _ = "end of CoverTab[82306]"
}

// parseInt64 treats the given bytes as a big-endian, signed integer and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:89
// returns the result.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:91
func parseInt64(bytes []byte) (ret int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:91
	_go_fuzz_dep_.CoverTab[82316]++
												err = checkInteger(bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:93
		_go_fuzz_dep_.CoverTab[82320]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:94
		// _ = "end of CoverTab[82320]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:95
		_go_fuzz_dep_.CoverTab[82321]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:95
		// _ = "end of CoverTab[82321]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:95
	// _ = "end of CoverTab[82316]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:95
	_go_fuzz_dep_.CoverTab[82317]++
												if len(bytes) > 8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:96
		_go_fuzz_dep_.CoverTab[82322]++

													err = StructuralError{"integer too large"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:99
		// _ = "end of CoverTab[82322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:100
		_go_fuzz_dep_.CoverTab[82323]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:100
		// _ = "end of CoverTab[82323]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:100
	// _ = "end of CoverTab[82317]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:100
	_go_fuzz_dep_.CoverTab[82318]++
												for bytesRead := 0; bytesRead < len(bytes); bytesRead++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:101
		_go_fuzz_dep_.CoverTab[82324]++
													ret <<= 8
													ret |= int64(bytes[bytesRead])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:103
		// _ = "end of CoverTab[82324]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:104
	// _ = "end of CoverTab[82318]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:104
	_go_fuzz_dep_.CoverTab[82319]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:107
	ret <<= 64 - uint8(len(bytes))*8
												ret >>= 64 - uint8(len(bytes))*8
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:109
	// _ = "end of CoverTab[82319]"
}

// parseInt treats the given bytes as a big-endian, signed integer and returns
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:112
// the result.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:114
func parseInt32(bytes []byte) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:114
	_go_fuzz_dep_.CoverTab[82325]++
												if err := checkInteger(bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:115
		_go_fuzz_dep_.CoverTab[82329]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:116
		// _ = "end of CoverTab[82329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:117
		_go_fuzz_dep_.CoverTab[82330]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:117
		// _ = "end of CoverTab[82330]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:117
	// _ = "end of CoverTab[82325]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:117
	_go_fuzz_dep_.CoverTab[82326]++
												ret64, err := parseInt64(bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:119
		_go_fuzz_dep_.CoverTab[82331]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:120
		// _ = "end of CoverTab[82331]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:121
		_go_fuzz_dep_.CoverTab[82332]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:121
		// _ = "end of CoverTab[82332]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:121
	// _ = "end of CoverTab[82326]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:121
	_go_fuzz_dep_.CoverTab[82327]++
												if ret64 != int64(int32(ret64)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:122
		_go_fuzz_dep_.CoverTab[82333]++
													return 0, StructuralError{"integer too large"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:123
		// _ = "end of CoverTab[82333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:124
		_go_fuzz_dep_.CoverTab[82334]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:124
		// _ = "end of CoverTab[82334]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:124
	// _ = "end of CoverTab[82327]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:124
	_go_fuzz_dep_.CoverTab[82328]++
												return int32(ret64), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:125
	// _ = "end of CoverTab[82328]"
}

var bigOne = big.NewInt(1)

// parseBigInt treats the given bytes as a big-endian, signed integer and returns
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:130
// the result.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:132
func parseBigInt(bytes []byte) (*big.Int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:132
	_go_fuzz_dep_.CoverTab[82335]++
												if err := checkInteger(bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:133
		_go_fuzz_dep_.CoverTab[82338]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:134
		// _ = "end of CoverTab[82338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:135
		_go_fuzz_dep_.CoverTab[82339]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:135
		// _ = "end of CoverTab[82339]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:135
	// _ = "end of CoverTab[82335]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:135
	_go_fuzz_dep_.CoverTab[82336]++
												ret := new(big.Int)
												if len(bytes) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:137
		_go_fuzz_dep_.CoverTab[82340]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:137
		return bytes[0]&0x80 == 0x80
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:137
		// _ = "end of CoverTab[82340]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:137
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:137
		_go_fuzz_dep_.CoverTab[82341]++

													notBytes := make([]byte, len(bytes))
													for i := range notBytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:140
			_go_fuzz_dep_.CoverTab[82343]++
														notBytes[i] = ^bytes[i]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:141
			// _ = "end of CoverTab[82343]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:142
		// _ = "end of CoverTab[82341]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:142
		_go_fuzz_dep_.CoverTab[82342]++
													ret.SetBytes(notBytes)
													ret.Add(ret, bigOne)
													ret.Neg(ret)
													return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:146
		// _ = "end of CoverTab[82342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:147
		_go_fuzz_dep_.CoverTab[82344]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:147
		// _ = "end of CoverTab[82344]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:147
	// _ = "end of CoverTab[82336]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:147
	_go_fuzz_dep_.CoverTab[82337]++
												ret.SetBytes(bytes)
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:149
	// _ = "end of CoverTab[82337]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:154
// BitString is the structure to use when you want an ASN.1 BIT STRING type. A
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:154
// bit string is padded up to the nearest byte in memory and the number of
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:154
// valid bits is recorded. Padding bits will be zero.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:157
type BitString struct {
	Bytes		[]byte	// bits packed into bytes.
	BitLength	int	// length in bits.
}

// At returns the bit at the given index. If the index is out of range it
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:162
// returns false.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:164
func (b BitString) At(i int) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:164
	_go_fuzz_dep_.CoverTab[82345]++
												if i < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:165
		_go_fuzz_dep_.CoverTab[82347]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:165
		return i >= b.BitLength
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:165
		// _ = "end of CoverTab[82347]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:165
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:165
		_go_fuzz_dep_.CoverTab[82348]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:166
		// _ = "end of CoverTab[82348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:167
		_go_fuzz_dep_.CoverTab[82349]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:167
		// _ = "end of CoverTab[82349]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:167
	// _ = "end of CoverTab[82345]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:167
	_go_fuzz_dep_.CoverTab[82346]++
												x := i / 8
												y := 7 - uint(i%8)
												return int(b.Bytes[x]>>y) & 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:170
	// _ = "end of CoverTab[82346]"
}

// RightAlign returns a slice where the padding bits are at the beginning. The
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:173
// slice may share memory with the BitString.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:175
func (b BitString) RightAlign() []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:175
	_go_fuzz_dep_.CoverTab[82350]++
												shift := uint(8 - (b.BitLength % 8))
												if shift == 8 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:177
		_go_fuzz_dep_.CoverTab[82353]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:177
		return len(b.Bytes) == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:177
		// _ = "end of CoverTab[82353]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:177
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:177
		_go_fuzz_dep_.CoverTab[82354]++
													return b.Bytes
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:178
		// _ = "end of CoverTab[82354]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:179
		_go_fuzz_dep_.CoverTab[82355]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:179
		// _ = "end of CoverTab[82355]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:179
	// _ = "end of CoverTab[82350]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:179
	_go_fuzz_dep_.CoverTab[82351]++

												a := make([]byte, len(b.Bytes))
												a[0] = b.Bytes[0] >> shift
												for i := 1; i < len(b.Bytes); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:183
		_go_fuzz_dep_.CoverTab[82356]++
													a[i] = b.Bytes[i-1] << (8 - shift)
													a[i] |= b.Bytes[i] >> shift
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:185
		// _ = "end of CoverTab[82356]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:186
	// _ = "end of CoverTab[82351]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:186
	_go_fuzz_dep_.CoverTab[82352]++

												return a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:188
	// _ = "end of CoverTab[82352]"
}

// parseBitString parses an ASN.1 bit string from the given byte slice and returns it.
func parseBitString(bytes []byte) (ret BitString, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:192
	_go_fuzz_dep_.CoverTab[82357]++
												if len(bytes) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:193
		_go_fuzz_dep_.CoverTab[82360]++
													err = SyntaxError{"zero length BIT STRING"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:195
		// _ = "end of CoverTab[82360]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:196
		_go_fuzz_dep_.CoverTab[82361]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:196
		// _ = "end of CoverTab[82361]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:196
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:196
	// _ = "end of CoverTab[82357]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:196
	_go_fuzz_dep_.CoverTab[82358]++
												paddingBits := int(bytes[0])
												if paddingBits > 7 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:198
		_go_fuzz_dep_.CoverTab[82362]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:198
		return len(bytes) == 1 && func() bool {
														_go_fuzz_dep_.CoverTab[82363]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
			return paddingBits > 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
			// _ = "end of CoverTab[82363]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
		}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
		// _ = "end of CoverTab[82362]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
		_go_fuzz_dep_.CoverTab[82364]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:199
		return bytes[len(bytes)-1]&((1<<bytes[0])-1) != 0
													// _ = "end of CoverTab[82364]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:200
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:200
		_go_fuzz_dep_.CoverTab[82365]++
													err = SyntaxError{"invalid padding bits in BIT STRING"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:202
		// _ = "end of CoverTab[82365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:203
		_go_fuzz_dep_.CoverTab[82366]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:203
		// _ = "end of CoverTab[82366]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:203
	// _ = "end of CoverTab[82358]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:203
	_go_fuzz_dep_.CoverTab[82359]++
												ret.BitLength = (len(bytes)-1)*8 - paddingBits
												ret.Bytes = bytes[1:]
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:206
	// _ = "end of CoverTab[82359]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:211
// An ObjectIdentifier represents an ASN.1 OBJECT IDENTIFIER.
type ObjectIdentifier []int

// Equal reports whether oi and other represent the same identifier.
func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:215
	_go_fuzz_dep_.CoverTab[82367]++
												if len(oi) != len(other) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:216
		_go_fuzz_dep_.CoverTab[82370]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:217
		// _ = "end of CoverTab[82370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:218
		_go_fuzz_dep_.CoverTab[82371]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:218
		// _ = "end of CoverTab[82371]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:218
	// _ = "end of CoverTab[82367]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:218
	_go_fuzz_dep_.CoverTab[82368]++
												for i := 0; i < len(oi); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:219
		_go_fuzz_dep_.CoverTab[82372]++
													if oi[i] != other[i] {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:220
			_go_fuzz_dep_.CoverTab[82373]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:221
			// _ = "end of CoverTab[82373]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:222
			_go_fuzz_dep_.CoverTab[82374]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:222
			// _ = "end of CoverTab[82374]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:222
		// _ = "end of CoverTab[82372]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:223
	// _ = "end of CoverTab[82368]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:223
	_go_fuzz_dep_.CoverTab[82369]++

												return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:225
	// _ = "end of CoverTab[82369]"
}

func (oi ObjectIdentifier) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:228
	_go_fuzz_dep_.CoverTab[82375]++
												var s string

												for i, v := range oi {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:231
		_go_fuzz_dep_.CoverTab[82377]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:232
			_go_fuzz_dep_.CoverTab[82379]++
														s += "."
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:233
			// _ = "end of CoverTab[82379]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:234
			_go_fuzz_dep_.CoverTab[82380]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:234
			// _ = "end of CoverTab[82380]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:234
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:234
		// _ = "end of CoverTab[82377]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:234
		_go_fuzz_dep_.CoverTab[82378]++
													s += strconv.Itoa(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:235
		// _ = "end of CoverTab[82378]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:236
	// _ = "end of CoverTab[82375]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:236
	_go_fuzz_dep_.CoverTab[82376]++

												return s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:238
	// _ = "end of CoverTab[82376]"
}

// parseObjectIdentifier parses an OBJECT IDENTIFIER from the given bytes and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:241
// returns it. An object identifier is a sequence of variable length integers
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:241
// that are assigned in a hierarchy.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:244
func parseObjectIdentifier(bytes []byte) (s []int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:244
	_go_fuzz_dep_.CoverTab[82381]++
												if len(bytes) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:245
		_go_fuzz_dep_.CoverTab[82386]++
													err = SyntaxError{"zero length OBJECT IDENTIFIER"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:247
		// _ = "end of CoverTab[82386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:248
		_go_fuzz_dep_.CoverTab[82387]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:248
		// _ = "end of CoverTab[82387]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:248
	// _ = "end of CoverTab[82381]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:248
	_go_fuzz_dep_.CoverTab[82382]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:252
	s = make([]int, len(bytes)+1)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:258
	v, offset, err := parseBase128Int(bytes, 0)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:259
		_go_fuzz_dep_.CoverTab[82388]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:260
		// _ = "end of CoverTab[82388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:261
		_go_fuzz_dep_.CoverTab[82389]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:261
		// _ = "end of CoverTab[82389]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:261
	// _ = "end of CoverTab[82382]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:261
	_go_fuzz_dep_.CoverTab[82383]++
												if v < 80 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:262
		_go_fuzz_dep_.CoverTab[82390]++
													s[0] = v / 40
													s[1] = v % 40
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:264
		// _ = "end of CoverTab[82390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:265
		_go_fuzz_dep_.CoverTab[82391]++
													s[0] = 2
													s[1] = v - 80
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:267
		// _ = "end of CoverTab[82391]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:268
	// _ = "end of CoverTab[82383]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:268
	_go_fuzz_dep_.CoverTab[82384]++

												i := 2
												for ; offset < len(bytes); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:271
		_go_fuzz_dep_.CoverTab[82392]++
													v, offset, err = parseBase128Int(bytes, offset)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:273
			_go_fuzz_dep_.CoverTab[82394]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:274
			// _ = "end of CoverTab[82394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:275
			_go_fuzz_dep_.CoverTab[82395]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:275
			// _ = "end of CoverTab[82395]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:275
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:275
		// _ = "end of CoverTab[82392]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:275
		_go_fuzz_dep_.CoverTab[82393]++
													s[i] = v
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:276
		// _ = "end of CoverTab[82393]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:277
	// _ = "end of CoverTab[82384]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:277
	_go_fuzz_dep_.CoverTab[82385]++
												s = s[0:i]
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:279
	// _ = "end of CoverTab[82385]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:284
// An Enumerated is represented as a plain int.
type Enumerated int

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:289
// A Flag accepts any data and is set to true if present.
type Flag bool

// parseBase128Int parses a base-128 encoded int from the given offset in the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:292
// given byte slice. It returns the value and the new offset.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:294
func parseBase128Int(bytes []byte, initOffset int) (ret, offset int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:294
	_go_fuzz_dep_.CoverTab[82396]++
												offset = initOffset
												for shifted := 0; offset < len(bytes); shifted++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:296
		_go_fuzz_dep_.CoverTab[82398]++
													if shifted == 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:297
			_go_fuzz_dep_.CoverTab[82400]++
														err = StructuralError{"base 128 integer too large"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:299
			// _ = "end of CoverTab[82400]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:300
			_go_fuzz_dep_.CoverTab[82401]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:300
			// _ = "end of CoverTab[82401]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:300
		// _ = "end of CoverTab[82398]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:300
		_go_fuzz_dep_.CoverTab[82399]++
													ret <<= 7
													b := bytes[offset]
													ret |= int(b & 0x7f)
													offset++
													if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:305
			_go_fuzz_dep_.CoverTab[82402]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:306
			// _ = "end of CoverTab[82402]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:307
			_go_fuzz_dep_.CoverTab[82403]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:307
			// _ = "end of CoverTab[82403]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:307
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:307
		// _ = "end of CoverTab[82399]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:308
	// _ = "end of CoverTab[82396]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:308
	_go_fuzz_dep_.CoverTab[82397]++
												err = SyntaxError{"truncated base 128 integer"}
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:310
	// _ = "end of CoverTab[82397]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:315
func parseUTCTime(bytes []byte) (ret time.Time, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:315
	_go_fuzz_dep_.CoverTab[82404]++
												s := string(bytes)

												formatStr := "0601021504Z0700"
												ret, err = time.Parse(formatStr, s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:320
		_go_fuzz_dep_.CoverTab[82409]++
													formatStr = "060102150405Z0700"
													ret, err = time.Parse(formatStr, s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:322
		// _ = "end of CoverTab[82409]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:323
		_go_fuzz_dep_.CoverTab[82410]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:323
		// _ = "end of CoverTab[82410]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:323
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:323
	// _ = "end of CoverTab[82404]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:323
	_go_fuzz_dep_.CoverTab[82405]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:324
		_go_fuzz_dep_.CoverTab[82411]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:325
		// _ = "end of CoverTab[82411]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:326
		_go_fuzz_dep_.CoverTab[82412]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:326
		// _ = "end of CoverTab[82412]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:326
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:326
	// _ = "end of CoverTab[82405]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:326
	_go_fuzz_dep_.CoverTab[82406]++

												if serialized := ret.Format(formatStr); serialized != s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:328
		_go_fuzz_dep_.CoverTab[82413]++
													err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:330
		// _ = "end of CoverTab[82413]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:331
		_go_fuzz_dep_.CoverTab[82414]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:331
		// _ = "end of CoverTab[82414]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:331
	// _ = "end of CoverTab[82406]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:331
	_go_fuzz_dep_.CoverTab[82407]++

												if ret.Year() >= 2050 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:333
		_go_fuzz_dep_.CoverTab[82415]++

													ret = ret.AddDate(-100, 0, 0)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:335
		// _ = "end of CoverTab[82415]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:336
		_go_fuzz_dep_.CoverTab[82416]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:336
		// _ = "end of CoverTab[82416]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:336
	// _ = "end of CoverTab[82407]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:336
	_go_fuzz_dep_.CoverTab[82408]++

												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:338
	// _ = "end of CoverTab[82408]"
}

// parseGeneralizedTime parses the GeneralizedTime from the given byte slice
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:341
// and returns the resulting time.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:343
func parseGeneralizedTime(bytes []byte) (ret time.Time, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:343
	_go_fuzz_dep_.CoverTab[82417]++
												const formatStr = "20060102150405Z0700"
												s := string(bytes)

												if ret, err = time.Parse(formatStr, s); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:347
		_go_fuzz_dep_.CoverTab[82420]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:348
		// _ = "end of CoverTab[82420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:349
		_go_fuzz_dep_.CoverTab[82421]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:349
		// _ = "end of CoverTab[82421]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:349
	// _ = "end of CoverTab[82417]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:349
	_go_fuzz_dep_.CoverTab[82418]++

												if serialized := ret.Format(formatStr); serialized != s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:351
		_go_fuzz_dep_.CoverTab[82422]++
													err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:352
		// _ = "end of CoverTab[82422]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:353
		_go_fuzz_dep_.CoverTab[82423]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:353
		// _ = "end of CoverTab[82423]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:353
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:353
	// _ = "end of CoverTab[82418]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:353
	_go_fuzz_dep_.CoverTab[82419]++

												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:355
	// _ = "end of CoverTab[82419]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:360
// parsePrintableString parses a ASN.1 PrintableString from the given byte
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:360
// array and returns it.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:362
func parsePrintableString(bytes []byte) (ret string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:362
	_go_fuzz_dep_.CoverTab[82424]++
												for _, b := range bytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:363
		_go_fuzz_dep_.CoverTab[82426]++
													if !isPrintable(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:364
			_go_fuzz_dep_.CoverTab[82427]++
														err = SyntaxError{"PrintableString contains invalid character"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:366
			// _ = "end of CoverTab[82427]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:367
			_go_fuzz_dep_.CoverTab[82428]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:367
			// _ = "end of CoverTab[82428]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:367
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:367
		// _ = "end of CoverTab[82426]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:368
	// _ = "end of CoverTab[82424]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:368
	_go_fuzz_dep_.CoverTab[82425]++
												ret = string(bytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:370
	// _ = "end of CoverTab[82425]"
}

// isPrintable reports whether the given b is in the ASN.1 PrintableString set.
func isPrintable(b byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:374
	_go_fuzz_dep_.CoverTab[82429]++
												return 'a' <= b && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
		_go_fuzz_dep_.CoverTab[82430]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
		return b <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
		// _ = "end of CoverTab[82430]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
		_go_fuzz_dep_.CoverTab[82431]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:375
		return 'A' <= b && func() bool {
														_go_fuzz_dep_.CoverTab[82432]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
			return b <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
			// _ = "end of CoverTab[82432]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
		}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
		// _ = "end of CoverTab[82431]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
		_go_fuzz_dep_.CoverTab[82433]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:376
		return '0' <= b && func() bool {
														_go_fuzz_dep_.CoverTab[82434]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
			return b <= '9'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
			// _ = "end of CoverTab[82434]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
		}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
		// _ = "end of CoverTab[82433]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
		_go_fuzz_dep_.CoverTab[82435]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:377
		return '\'' <= b && func() bool {
														_go_fuzz_dep_.CoverTab[82436]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
			return b <= ')'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
			// _ = "end of CoverTab[82436]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
		}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
		// _ = "end of CoverTab[82435]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
		_go_fuzz_dep_.CoverTab[82437]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:378
		return '+' <= b && func() bool {
														_go_fuzz_dep_.CoverTab[82438]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
			return b <= '/'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
			// _ = "end of CoverTab[82438]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
		}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
		// _ = "end of CoverTab[82437]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
		_go_fuzz_dep_.CoverTab[82439]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:379
		return b == ' '
													// _ = "end of CoverTab[82439]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:380
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:380
		_go_fuzz_dep_.CoverTab[82440]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:380
		return b == ':'
													// _ = "end of CoverTab[82440]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:381
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:381
		_go_fuzz_dep_.CoverTab[82441]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:381
		return b == '='
													// _ = "end of CoverTab[82441]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:382
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:382
		_go_fuzz_dep_.CoverTab[82442]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:382
		return b == '?'
													// _ = "end of CoverTab[82442]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:383
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:383
		_go_fuzz_dep_.CoverTab[82443]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:383
		return b == '*'
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:387
		// _ = "end of CoverTab[82443]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:387
	}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:387
	// _ = "end of CoverTab[82429]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:392
// parseIA5String parses a ASN.1 IA5String (ASCII string) from the given
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:392
// byte slice and returns it.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:394
func parseIA5String(bytes []byte) (ret string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:394
	_go_fuzz_dep_.CoverTab[82444]++
												for _, b := range bytes {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:395
		_go_fuzz_dep_.CoverTab[82446]++
													if b >= utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:396
			_go_fuzz_dep_.CoverTab[82447]++
														err = SyntaxError{"IA5String contains invalid character"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:398
			// _ = "end of CoverTab[82447]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:399
			_go_fuzz_dep_.CoverTab[82448]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:399
			// _ = "end of CoverTab[82448]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:399
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:399
		// _ = "end of CoverTab[82446]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:400
	// _ = "end of CoverTab[82444]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:400
	_go_fuzz_dep_.CoverTab[82445]++
												ret = string(bytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:402
	// _ = "end of CoverTab[82445]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:407
// parseT61String parses a ASN.1 T61String (8-bit clean string) from the given
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:407
// byte slice and returns it.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:409
func parseT61String(bytes []byte) (ret string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:409
	_go_fuzz_dep_.CoverTab[82449]++
												return string(bytes), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:410
	// _ = "end of CoverTab[82449]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:415
// parseUTF8String parses a ASN.1 UTF8String (raw UTF-8) from the given byte
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:415
// array and returns it.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:417
func parseUTF8String(bytes []byte) (ret string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:417
	_go_fuzz_dep_.CoverTab[82450]++
												if !utf8.Valid(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:418
		_go_fuzz_dep_.CoverTab[82452]++
													return "", errors.New("asn1: invalid UTF-8 string")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:419
		// _ = "end of CoverTab[82452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:420
		_go_fuzz_dep_.CoverTab[82453]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:420
		// _ = "end of CoverTab[82453]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:420
	// _ = "end of CoverTab[82450]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:420
	_go_fuzz_dep_.CoverTab[82451]++
												return string(bytes), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:421
	// _ = "end of CoverTab[82451]"
}

// A RawValue represents an undecoded ASN.1 object.
type RawValue struct {
	Class, Tag	int
	IsCompound	bool
	Bytes		[]byte
	FullBytes	[]byte	// includes the tag and length
}

// RawContent is used to signal that the undecoded, DER data needs to be
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:432
// preserved for a struct. To use it, the first field of the struct must have
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:432
// this type. It's an error for any of the other fields to have this type.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:435
type RawContent []byte

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:439
// parseTagAndLength parses an ASN.1 tag and length pair from the given offset
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:439
// into a byte slice. It returns the parsed data and the new offset. SET and
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:439
// SET OF (tag 17) are mapped to SEQUENCE and SEQUENCE OF (tag 16) since we
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:439
// don't distinguish between ordered and unordered objects in this code.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:443
func parseTagAndLength(bytes []byte, initOffset int) (ret tagAndLength, offset int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:443
	_go_fuzz_dep_.CoverTab[82454]++
												offset = initOffset

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:447
	if offset >= len(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:447
		_go_fuzz_dep_.CoverTab[82459]++
													err = errors.New("asn1: internal error in parseTagAndLength")
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:449
		// _ = "end of CoverTab[82459]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:450
		_go_fuzz_dep_.CoverTab[82460]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:450
		// _ = "end of CoverTab[82460]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:450
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:450
	// _ = "end of CoverTab[82454]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:450
	_go_fuzz_dep_.CoverTab[82455]++
												b := bytes[offset]
												offset++
												ret.class = int(b >> 6)
												ret.isCompound = b&0x20 == 0x20
												ret.tag = int(b & 0x1f)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:459
	if ret.tag == 0x1f {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:459
		_go_fuzz_dep_.CoverTab[82461]++
													ret.tag, offset, err = parseBase128Int(bytes, offset)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:461
			_go_fuzz_dep_.CoverTab[82463]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:462
			// _ = "end of CoverTab[82463]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:463
			_go_fuzz_dep_.CoverTab[82464]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:463
			// _ = "end of CoverTab[82464]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:463
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:463
		// _ = "end of CoverTab[82461]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:463
		_go_fuzz_dep_.CoverTab[82462]++

													if ret.tag < 0x1f {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:465
			_go_fuzz_dep_.CoverTab[82465]++
														err = SyntaxError{"non-minimal tag"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:467
			// _ = "end of CoverTab[82465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:468
			_go_fuzz_dep_.CoverTab[82466]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:468
			// _ = "end of CoverTab[82466]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:468
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:468
		// _ = "end of CoverTab[82462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:469
		_go_fuzz_dep_.CoverTab[82467]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:469
		// _ = "end of CoverTab[82467]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:469
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:469
	// _ = "end of CoverTab[82455]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:469
	_go_fuzz_dep_.CoverTab[82456]++
												if offset >= len(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:470
		_go_fuzz_dep_.CoverTab[82468]++
													err = SyntaxError{"truncated tag or length"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:472
		// _ = "end of CoverTab[82468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:473
		_go_fuzz_dep_.CoverTab[82469]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:473
		// _ = "end of CoverTab[82469]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:473
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:473
	// _ = "end of CoverTab[82456]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:473
	_go_fuzz_dep_.CoverTab[82457]++
												b = bytes[offset]
												offset++
												if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:476
		_go_fuzz_dep_.CoverTab[82470]++

													ret.length = int(b & 0x7f)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:478
		// _ = "end of CoverTab[82470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:479
		_go_fuzz_dep_.CoverTab[82471]++

													numBytes := int(b & 0x7f)
													if numBytes == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:482
			_go_fuzz_dep_.CoverTab[82474]++
														err = SyntaxError{"indefinite length found (not DER)"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:484
			// _ = "end of CoverTab[82474]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:485
			_go_fuzz_dep_.CoverTab[82475]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:485
			// _ = "end of CoverTab[82475]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:485
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:485
		// _ = "end of CoverTab[82471]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:485
		_go_fuzz_dep_.CoverTab[82472]++
													ret.length = 0
													for i := 0; i < numBytes; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:487
			_go_fuzz_dep_.CoverTab[82476]++
														if offset >= len(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:488
				_go_fuzz_dep_.CoverTab[82479]++
															err = SyntaxError{"truncated tag or length"}
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:490
				// _ = "end of CoverTab[82479]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:491
				_go_fuzz_dep_.CoverTab[82480]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:491
				// _ = "end of CoverTab[82480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:491
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:491
			// _ = "end of CoverTab[82476]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:491
			_go_fuzz_dep_.CoverTab[82477]++
														b = bytes[offset]
														offset++
														if ret.length >= 1<<23 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:494
				_go_fuzz_dep_.CoverTab[82481]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:497
				err = StructuralError{"length too large"}
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:498
				// _ = "end of CoverTab[82481]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:499
				_go_fuzz_dep_.CoverTab[82482]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:499
				// _ = "end of CoverTab[82482]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:499
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:499
			// _ = "end of CoverTab[82477]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:499
			_go_fuzz_dep_.CoverTab[82478]++
														ret.length <<= 8
														ret.length |= int(b)
														if ret.length == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:502
				_go_fuzz_dep_.CoverTab[82483]++

															err = StructuralError{"superfluous leading zeros in length"}
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:505
				// _ = "end of CoverTab[82483]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:506
				_go_fuzz_dep_.CoverTab[82484]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:506
				// _ = "end of CoverTab[82484]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:506
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:506
			// _ = "end of CoverTab[82478]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:507
		// _ = "end of CoverTab[82472]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:507
		_go_fuzz_dep_.CoverTab[82473]++

													if ret.length < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:509
			_go_fuzz_dep_.CoverTab[82485]++
														err = StructuralError{"non-minimal length"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:511
			// _ = "end of CoverTab[82485]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:512
			_go_fuzz_dep_.CoverTab[82486]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:512
			// _ = "end of CoverTab[82486]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:512
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:512
		// _ = "end of CoverTab[82473]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:513
	// _ = "end of CoverTab[82457]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:513
	_go_fuzz_dep_.CoverTab[82458]++

												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:515
	// _ = "end of CoverTab[82458]"
}

// parseSequenceOf is used for SEQUENCE OF and SET OF values. It tries to parse
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:518
// a number of ASN.1 values from the given byte slice and returns them as a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:518
// slice of Go values of the given type.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:521
func parseSequenceOf(bytes []byte, sliceType reflect.Type, elemType reflect.Type) (ret reflect.Value, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:521
	_go_fuzz_dep_.CoverTab[82487]++
												expectedTag, compoundType, ok := getUniversalType(elemType)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:523
		_go_fuzz_dep_.CoverTab[82491]++
													err = StructuralError{"unknown Go type for slice"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:525
		// _ = "end of CoverTab[82491]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:526
		_go_fuzz_dep_.CoverTab[82492]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:526
		// _ = "end of CoverTab[82492]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:526
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:526
	// _ = "end of CoverTab[82487]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:526
	_go_fuzz_dep_.CoverTab[82488]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:530
	numElements := 0
	for offset := 0; offset < len(bytes); {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:531
		_go_fuzz_dep_.CoverTab[82493]++
													var t tagAndLength
													t, offset, err = parseTagAndLength(bytes, offset)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:534
			_go_fuzz_dep_.CoverTab[82498]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:535
			// _ = "end of CoverTab[82498]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:536
			_go_fuzz_dep_.CoverTab[82499]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:536
			// _ = "end of CoverTab[82499]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:536
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:536
		// _ = "end of CoverTab[82493]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:536
		_go_fuzz_dep_.CoverTab[82494]++
													switch t.tag {
		case TagIA5String, TagGeneralString, TagT61String, TagUTF8String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:538
			_go_fuzz_dep_.CoverTab[82500]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:542
			t.tag = TagPrintableString
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:542
			// _ = "end of CoverTab[82500]"
		case TagGeneralizedTime, TagUTCTime:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:543
			_go_fuzz_dep_.CoverTab[82501]++

														t.tag = TagUTCTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:545
			// _ = "end of CoverTab[82501]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:545
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:545
			_go_fuzz_dep_.CoverTab[82502]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:545
			// _ = "end of CoverTab[82502]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:546
		// _ = "end of CoverTab[82494]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:546
		_go_fuzz_dep_.CoverTab[82495]++

													if t.class != ClassUniversal || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			_go_fuzz_dep_.CoverTab[82503]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			return t.isCompound != compoundType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			// _ = "end of CoverTab[82503]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			_go_fuzz_dep_.CoverTab[82504]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			return t.tag != expectedTag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			// _ = "end of CoverTab[82504]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:548
			_go_fuzz_dep_.CoverTab[82505]++
														err = StructuralError{"sequence tag mismatch"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:550
			// _ = "end of CoverTab[82505]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:551
			_go_fuzz_dep_.CoverTab[82506]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:551
			// _ = "end of CoverTab[82506]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:551
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:551
		// _ = "end of CoverTab[82495]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:551
		_go_fuzz_dep_.CoverTab[82496]++
													if invalidLength(offset, t.length, len(bytes)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:552
			_go_fuzz_dep_.CoverTab[82507]++
														err = SyntaxError{"truncated sequence"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:554
			// _ = "end of CoverTab[82507]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:555
			_go_fuzz_dep_.CoverTab[82508]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:555
			// _ = "end of CoverTab[82508]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:555
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:555
		// _ = "end of CoverTab[82496]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:555
		_go_fuzz_dep_.CoverTab[82497]++
													offset += t.length
													numElements++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:557
		// _ = "end of CoverTab[82497]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:558
	// _ = "end of CoverTab[82488]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:558
	_go_fuzz_dep_.CoverTab[82489]++
												ret = reflect.MakeSlice(sliceType, numElements, numElements)
												params := fieldParameters{}
												offset := 0
												for i := 0; i < numElements; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:562
		_go_fuzz_dep_.CoverTab[82509]++
													offset, err = parseField(ret.Index(i), bytes, offset, params)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:564
			_go_fuzz_dep_.CoverTab[82510]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:565
			// _ = "end of CoverTab[82510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:566
			_go_fuzz_dep_.CoverTab[82511]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:566
			// _ = "end of CoverTab[82511]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:566
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:566
		// _ = "end of CoverTab[82509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:567
	// _ = "end of CoverTab[82489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:567
	_go_fuzz_dep_.CoverTab[82490]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:568
	// _ = "end of CoverTab[82490]"
}

var (
	bitStringType		= reflect.TypeOf(BitString{})
	objectIdentifierType	= reflect.TypeOf(ObjectIdentifier{})
	enumeratedType		= reflect.TypeOf(Enumerated(0))
	flagType		= reflect.TypeOf(Flag(false))
	timeType		= reflect.TypeOf(time.Time{})
	rawValueType		= reflect.TypeOf(RawValue{})
	rawContentsType		= reflect.TypeOf(RawContent(nil))
	bigIntType		= reflect.TypeOf(new(big.Int))
)

// invalidLength returns true iff offset + length > sliceLength, or if the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:582
// addition would overflow.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:584
func invalidLength(offset, length, sliceLength int) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:584
	_go_fuzz_dep_.CoverTab[82512]++
												return offset+length < offset || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:585
		_go_fuzz_dep_.CoverTab[82513]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:585
		return offset+length > sliceLength
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:585
		// _ = "end of CoverTab[82513]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:585
	}()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:585
	// _ = "end of CoverTab[82512]"
}

// parseField is the main parsing function. Given a byte slice and an offset
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:588
// into the array, it will try to parse a suitable ASN.1 value out and store it
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:588
// in the given Value.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:591
func parseField(v reflect.Value, bytes []byte, initOffset int, params fieldParameters) (offset int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:591
	_go_fuzz_dep_.CoverTab[82514]++
												offset = initOffset
												fieldType := v.Type()

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:596
	if offset == len(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:596
		_go_fuzz_dep_.CoverTab[82530]++
													if !setDefaultValue(v, params) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:597
			_go_fuzz_dep_.CoverTab[82532]++
														err = SyntaxError{"sequence truncated"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:598
			// _ = "end of CoverTab[82532]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:599
			_go_fuzz_dep_.CoverTab[82533]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:599
			// _ = "end of CoverTab[82533]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:599
		// _ = "end of CoverTab[82530]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:599
		_go_fuzz_dep_.CoverTab[82531]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:600
		// _ = "end of CoverTab[82531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:601
		_go_fuzz_dep_.CoverTab[82534]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:601
		// _ = "end of CoverTab[82534]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:601
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:601
	// _ = "end of CoverTab[82514]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:601
	_go_fuzz_dep_.CoverTab[82515]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:604
	if fieldType == rawValueType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:604
		_go_fuzz_dep_.CoverTab[82535]++
													var t tagAndLength
													t, offset, err = parseTagAndLength(bytes, offset)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:607
			_go_fuzz_dep_.CoverTab[82538]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:608
			// _ = "end of CoverTab[82538]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:609
			_go_fuzz_dep_.CoverTab[82539]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:609
			// _ = "end of CoverTab[82539]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:609
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:609
		// _ = "end of CoverTab[82535]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:609
		_go_fuzz_dep_.CoverTab[82536]++
													if invalidLength(offset, t.length, len(bytes)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:610
			_go_fuzz_dep_.CoverTab[82540]++
														err = SyntaxError{"data truncated"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:612
			// _ = "end of CoverTab[82540]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:613
			_go_fuzz_dep_.CoverTab[82541]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:613
			// _ = "end of CoverTab[82541]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:613
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:613
		// _ = "end of CoverTab[82536]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:613
		_go_fuzz_dep_.CoverTab[82537]++
													result := RawValue{t.class, t.tag, t.isCompound, bytes[offset : offset+t.length], bytes[initOffset : offset+t.length]}
													offset += t.length
													v.Set(reflect.ValueOf(result))
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:617
		// _ = "end of CoverTab[82537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:618
		_go_fuzz_dep_.CoverTab[82542]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:618
		// _ = "end of CoverTab[82542]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:618
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:618
	// _ = "end of CoverTab[82515]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:618
	_go_fuzz_dep_.CoverTab[82516]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
	if ifaceType := fieldType; ifaceType.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
		_go_fuzz_dep_.CoverTab[82543]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
		return ifaceType.NumMethod() == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
		// _ = "end of CoverTab[82543]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:621
		_go_fuzz_dep_.CoverTab[82544]++
													var t tagAndLength
													t, offset, err = parseTagAndLength(bytes, offset)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:624
			_go_fuzz_dep_.CoverTab[82550]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:625
			// _ = "end of CoverTab[82550]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:626
			_go_fuzz_dep_.CoverTab[82551]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:626
			// _ = "end of CoverTab[82551]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:626
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:626
		// _ = "end of CoverTab[82544]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:626
		_go_fuzz_dep_.CoverTab[82545]++
													if invalidLength(offset, t.length, len(bytes)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:627
			_go_fuzz_dep_.CoverTab[82552]++
														err = SyntaxError{"data truncated"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:629
			// _ = "end of CoverTab[82552]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:630
			_go_fuzz_dep_.CoverTab[82553]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:630
			// _ = "end of CoverTab[82553]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:630
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:630
		// _ = "end of CoverTab[82545]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:630
		_go_fuzz_dep_.CoverTab[82546]++
													var result interface{}
													if !t.isCompound && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:632
			_go_fuzz_dep_.CoverTab[82554]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:632
			return t.class == ClassUniversal
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:632
			// _ = "end of CoverTab[82554]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:632
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:632
			_go_fuzz_dep_.CoverTab[82555]++
														innerBytes := bytes[offset : offset+t.length]
														switch t.tag {
			case TagPrintableString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:635
				_go_fuzz_dep_.CoverTab[82556]++
															result, err = parsePrintableString(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:636
				// _ = "end of CoverTab[82556]"
			case TagIA5String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:637
				_go_fuzz_dep_.CoverTab[82557]++
															result, err = parseIA5String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:638
				// _ = "end of CoverTab[82557]"

			case TagGeneralString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:640
				_go_fuzz_dep_.CoverTab[82558]++
															result, err = parseIA5String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:641
				// _ = "end of CoverTab[82558]"
			case TagT61String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:642
				_go_fuzz_dep_.CoverTab[82559]++
															result, err = parseT61String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:643
				// _ = "end of CoverTab[82559]"
			case TagUTF8String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:644
				_go_fuzz_dep_.CoverTab[82560]++
															result, err = parseUTF8String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:645
				// _ = "end of CoverTab[82560]"
			case TagInteger:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:646
				_go_fuzz_dep_.CoverTab[82561]++
															result, err = parseInt64(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:647
				// _ = "end of CoverTab[82561]"
			case TagBitString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:648
				_go_fuzz_dep_.CoverTab[82562]++
															result, err = parseBitString(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:649
				// _ = "end of CoverTab[82562]"
			case TagOID:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:650
				_go_fuzz_dep_.CoverTab[82563]++
															result, err = parseObjectIdentifier(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:651
				// _ = "end of CoverTab[82563]"
			case TagUTCTime:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:652
				_go_fuzz_dep_.CoverTab[82564]++
															result, err = parseUTCTime(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:653
				// _ = "end of CoverTab[82564]"
			case TagGeneralizedTime:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:654
				_go_fuzz_dep_.CoverTab[82565]++
															result, err = parseGeneralizedTime(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:655
				// _ = "end of CoverTab[82565]"
			case TagOctetString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:656
				_go_fuzz_dep_.CoverTab[82566]++
															result = innerBytes
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:657
				// _ = "end of CoverTab[82566]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:658
				_go_fuzz_dep_.CoverTab[82567]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:658
				// _ = "end of CoverTab[82567]"

			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:660
			// _ = "end of CoverTab[82555]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:661
			_go_fuzz_dep_.CoverTab[82568]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:661
			// _ = "end of CoverTab[82568]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:661
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:661
		// _ = "end of CoverTab[82546]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:661
		_go_fuzz_dep_.CoverTab[82547]++
													offset += t.length
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:663
			_go_fuzz_dep_.CoverTab[82569]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:664
			// _ = "end of CoverTab[82569]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:665
			_go_fuzz_dep_.CoverTab[82570]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:665
			// _ = "end of CoverTab[82570]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:665
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:665
		// _ = "end of CoverTab[82547]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:665
		_go_fuzz_dep_.CoverTab[82548]++
													if result != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:666
			_go_fuzz_dep_.CoverTab[82571]++
														v.Set(reflect.ValueOf(result))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:667
			// _ = "end of CoverTab[82571]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:668
			_go_fuzz_dep_.CoverTab[82572]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:668
			// _ = "end of CoverTab[82572]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:668
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:668
		// _ = "end of CoverTab[82548]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:668
		_go_fuzz_dep_.CoverTab[82549]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:669
		// _ = "end of CoverTab[82549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:670
		_go_fuzz_dep_.CoverTab[82573]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:670
		// _ = "end of CoverTab[82573]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:670
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:670
	// _ = "end of CoverTab[82516]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:670
	_go_fuzz_dep_.CoverTab[82517]++
												universalTag, compoundType, ok1 := getUniversalType(fieldType)
												if !ok1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:672
		_go_fuzz_dep_.CoverTab[82574]++
													err = StructuralError{fmt.Sprintf("unknown Go type: %v", fieldType)}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:674
		// _ = "end of CoverTab[82574]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:675
		_go_fuzz_dep_.CoverTab[82575]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:675
		// _ = "end of CoverTab[82575]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:675
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:675
	// _ = "end of CoverTab[82517]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:675
	_go_fuzz_dep_.CoverTab[82518]++

												t, offset, err := parseTagAndLength(bytes, offset)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:678
		_go_fuzz_dep_.CoverTab[82576]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:679
		// _ = "end of CoverTab[82576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:680
		_go_fuzz_dep_.CoverTab[82577]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:680
		// _ = "end of CoverTab[82577]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:680
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:680
	// _ = "end of CoverTab[82518]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:680
	_go_fuzz_dep_.CoverTab[82519]++
												if params.explicit {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:681
		_go_fuzz_dep_.CoverTab[82578]++
													expectedClass := ClassContextSpecific
													if params.application {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:683
			_go_fuzz_dep_.CoverTab[82581]++
														expectedClass = ClassApplication
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:684
			// _ = "end of CoverTab[82581]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:685
			_go_fuzz_dep_.CoverTab[82582]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:685
			// _ = "end of CoverTab[82582]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:685
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:685
		// _ = "end of CoverTab[82578]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:685
		_go_fuzz_dep_.CoverTab[82579]++
													if offset == len(bytes) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:686
			_go_fuzz_dep_.CoverTab[82583]++
														err = StructuralError{"explicit tag has no child"}
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:688
			// _ = "end of CoverTab[82583]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:689
			_go_fuzz_dep_.CoverTab[82584]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:689
			// _ = "end of CoverTab[82584]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:689
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:689
		// _ = "end of CoverTab[82579]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:689
		_go_fuzz_dep_.CoverTab[82580]++
													if t.class == expectedClass && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			_go_fuzz_dep_.CoverTab[82585]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			return t.tag == *params.tag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			// _ = "end of CoverTab[82585]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			_go_fuzz_dep_.CoverTab[82586]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			return (t.length == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
				_go_fuzz_dep_.CoverTab[82587]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
				return t.isCompound
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
				// _ = "end of CoverTab[82587]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			}())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			// _ = "end of CoverTab[82586]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:690
			_go_fuzz_dep_.CoverTab[82588]++
														if t.length > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:691
				_go_fuzz_dep_.CoverTab[82589]++
															t, offset, err = parseTagAndLength(bytes, offset)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:693
					_go_fuzz_dep_.CoverTab[82590]++
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:694
					// _ = "end of CoverTab[82590]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:695
					_go_fuzz_dep_.CoverTab[82591]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:695
					// _ = "end of CoverTab[82591]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:695
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:695
				// _ = "end of CoverTab[82589]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:696
				_go_fuzz_dep_.CoverTab[82592]++
															if fieldType != flagType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:697
					_go_fuzz_dep_.CoverTab[82594]++
																err = StructuralError{"zero length explicit tag was not an asn1.Flag"}
																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:699
					// _ = "end of CoverTab[82594]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:700
					_go_fuzz_dep_.CoverTab[82595]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:700
					// _ = "end of CoverTab[82595]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:700
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:700
				// _ = "end of CoverTab[82592]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:700
				_go_fuzz_dep_.CoverTab[82593]++
															v.SetBool(true)
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:702
				// _ = "end of CoverTab[82593]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:703
			// _ = "end of CoverTab[82588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:704
			_go_fuzz_dep_.CoverTab[82596]++

														ok := setDefaultValue(v, params)
														if ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:707
				_go_fuzz_dep_.CoverTab[82598]++
															offset = initOffset
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:708
				// _ = "end of CoverTab[82598]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:709
				_go_fuzz_dep_.CoverTab[82599]++
															err = StructuralError{"explicitly tagged member didn't match"}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:710
				// _ = "end of CoverTab[82599]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:711
			// _ = "end of CoverTab[82596]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:711
			_go_fuzz_dep_.CoverTab[82597]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:712
			// _ = "end of CoverTab[82597]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:713
		// _ = "end of CoverTab[82580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:714
		_go_fuzz_dep_.CoverTab[82600]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:714
		// _ = "end of CoverTab[82600]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:714
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:714
	// _ = "end of CoverTab[82519]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:714
	_go_fuzz_dep_.CoverTab[82520]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:720
	if universalTag == TagPrintableString {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:720
		_go_fuzz_dep_.CoverTab[82601]++
													if t.class == ClassUniversal {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:721
			_go_fuzz_dep_.CoverTab[82602]++
														switch t.tag {
			case TagIA5String, TagGeneralString, TagT61String, TagUTF8String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:723
				_go_fuzz_dep_.CoverTab[82603]++
															universalTag = t.tag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:724
				// _ = "end of CoverTab[82603]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:724
			default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:724
				_go_fuzz_dep_.CoverTab[82604]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:724
				// _ = "end of CoverTab[82604]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:725
			// _ = "end of CoverTab[82602]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:726
			_go_fuzz_dep_.CoverTab[82605]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:726
			if params.stringType != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:726
				_go_fuzz_dep_.CoverTab[82606]++
															universalTag = params.stringType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:727
				// _ = "end of CoverTab[82606]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
				_go_fuzz_dep_.CoverTab[82607]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
				// _ = "end of CoverTab[82607]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
			// _ = "end of CoverTab[82605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:728
		// _ = "end of CoverTab[82601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:729
		_go_fuzz_dep_.CoverTab[82608]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:729
		// _ = "end of CoverTab[82608]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:729
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:729
	// _ = "end of CoverTab[82520]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:729
	_go_fuzz_dep_.CoverTab[82521]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
	if universalTag == TagUTCTime && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		_go_fuzz_dep_.CoverTab[82609]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		return t.tag == TagGeneralizedTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		// _ = "end of CoverTab[82609]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		_go_fuzz_dep_.CoverTab[82610]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		return t.class == ClassUniversal
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		// _ = "end of CoverTab[82610]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:733
		_go_fuzz_dep_.CoverTab[82611]++
													universalTag = TagGeneralizedTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:734
		// _ = "end of CoverTab[82611]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:735
		_go_fuzz_dep_.CoverTab[82612]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:735
		// _ = "end of CoverTab[82612]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:735
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:735
	// _ = "end of CoverTab[82521]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:735
	_go_fuzz_dep_.CoverTab[82522]++

												if params.set {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:737
		_go_fuzz_dep_.CoverTab[82613]++
													universalTag = TagSet
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:738
		// _ = "end of CoverTab[82613]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:739
		_go_fuzz_dep_.CoverTab[82614]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:739
		// _ = "end of CoverTab[82614]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:739
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:739
	// _ = "end of CoverTab[82522]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:739
	_go_fuzz_dep_.CoverTab[82523]++

												expectedClass := ClassUniversal
												expectedTag := universalTag

												if !params.explicit && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:744
		_go_fuzz_dep_.CoverTab[82615]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:744
		return params.tag != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:744
		// _ = "end of CoverTab[82615]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:744
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:744
		_go_fuzz_dep_.CoverTab[82616]++
													expectedClass = ClassContextSpecific
													expectedTag = *params.tag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:746
		// _ = "end of CoverTab[82616]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:747
		_go_fuzz_dep_.CoverTab[82617]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:747
		// _ = "end of CoverTab[82617]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:747
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:747
	// _ = "end of CoverTab[82523]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:747
	_go_fuzz_dep_.CoverTab[82524]++

												if !params.explicit && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		_go_fuzz_dep_.CoverTab[82618]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		return params.application
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		// _ = "end of CoverTab[82618]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		_go_fuzz_dep_.CoverTab[82619]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		return params.tag != nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		// _ = "end of CoverTab[82619]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:749
		_go_fuzz_dep_.CoverTab[82620]++
													expectedClass = ClassApplication
													expectedTag = *params.tag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:751
		// _ = "end of CoverTab[82620]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:752
		_go_fuzz_dep_.CoverTab[82621]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:752
		// _ = "end of CoverTab[82621]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:752
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:752
	// _ = "end of CoverTab[82524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:752
	_go_fuzz_dep_.CoverTab[82525]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
	if t.class != expectedClass || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		_go_fuzz_dep_.CoverTab[82622]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		return t.tag != expectedTag
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		// _ = "end of CoverTab[82622]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		_go_fuzz_dep_.CoverTab[82623]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		return t.isCompound != compoundType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		// _ = "end of CoverTab[82623]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:755
		_go_fuzz_dep_.CoverTab[82624]++

													ok := setDefaultValue(v, params)
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:758
			_go_fuzz_dep_.CoverTab[82626]++
														offset = initOffset
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:759
			// _ = "end of CoverTab[82626]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:760
			_go_fuzz_dep_.CoverTab[82627]++
														err = StructuralError{fmt.Sprintf("tags don't match (%d vs %+v) %+v %s @%d", expectedTag, t, params, fieldType.Name(), offset)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:761
			// _ = "end of CoverTab[82627]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:762
		// _ = "end of CoverTab[82624]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:762
		_go_fuzz_dep_.CoverTab[82625]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:763
		// _ = "end of CoverTab[82625]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:764
		_go_fuzz_dep_.CoverTab[82628]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:764
		// _ = "end of CoverTab[82628]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:764
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:764
	// _ = "end of CoverTab[82525]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:764
	_go_fuzz_dep_.CoverTab[82526]++
												if invalidLength(offset, t.length, len(bytes)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:765
		_go_fuzz_dep_.CoverTab[82629]++
													err = SyntaxError{"data truncated"}
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:767
		// _ = "end of CoverTab[82629]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:768
		_go_fuzz_dep_.CoverTab[82630]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:768
		// _ = "end of CoverTab[82630]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:768
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:768
	// _ = "end of CoverTab[82526]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:768
	_go_fuzz_dep_.CoverTab[82527]++
												innerBytes := bytes[offset : offset+t.length]
												offset += t.length

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:773
	switch fieldType {
	case objectIdentifierType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:774
		_go_fuzz_dep_.CoverTab[82631]++
													newSlice, err1 := parseObjectIdentifier(innerBytes)
													v.Set(reflect.MakeSlice(v.Type(), len(newSlice), len(newSlice)))
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:777
			_go_fuzz_dep_.CoverTab[82644]++
														reflect.Copy(v, reflect.ValueOf(newSlice))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:778
			// _ = "end of CoverTab[82644]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:779
			_go_fuzz_dep_.CoverTab[82645]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:779
			// _ = "end of CoverTab[82645]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:779
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:779
		// _ = "end of CoverTab[82631]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:779
		_go_fuzz_dep_.CoverTab[82632]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:781
		// _ = "end of CoverTab[82632]"
	case bitStringType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:782
		_go_fuzz_dep_.CoverTab[82633]++
													bs, err1 := parseBitString(innerBytes)
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:784
			_go_fuzz_dep_.CoverTab[82646]++
														v.Set(reflect.ValueOf(bs))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:785
			// _ = "end of CoverTab[82646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:786
			_go_fuzz_dep_.CoverTab[82647]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:786
			// _ = "end of CoverTab[82647]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:786
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:786
		// _ = "end of CoverTab[82633]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:786
		_go_fuzz_dep_.CoverTab[82634]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:788
		// _ = "end of CoverTab[82634]"
	case timeType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:789
		_go_fuzz_dep_.CoverTab[82635]++
													var time time.Time
													var err1 error
													if universalTag == TagUTCTime {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:792
			_go_fuzz_dep_.CoverTab[82648]++
														time, err1 = parseUTCTime(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:793
			// _ = "end of CoverTab[82648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:794
			_go_fuzz_dep_.CoverTab[82649]++
														time, err1 = parseGeneralizedTime(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:795
			// _ = "end of CoverTab[82649]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:796
		// _ = "end of CoverTab[82635]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:796
		_go_fuzz_dep_.CoverTab[82636]++
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:797
			_go_fuzz_dep_.CoverTab[82650]++
														v.Set(reflect.ValueOf(time))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:798
			// _ = "end of CoverTab[82650]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:799
			_go_fuzz_dep_.CoverTab[82651]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:799
			// _ = "end of CoverTab[82651]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:799
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:799
		// _ = "end of CoverTab[82636]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:799
		_go_fuzz_dep_.CoverTab[82637]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:801
		// _ = "end of CoverTab[82637]"
	case enumeratedType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:802
		_go_fuzz_dep_.CoverTab[82638]++
													parsedInt, err1 := parseInt32(innerBytes)
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:804
			_go_fuzz_dep_.CoverTab[82652]++
														v.SetInt(int64(parsedInt))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:805
			// _ = "end of CoverTab[82652]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:806
			_go_fuzz_dep_.CoverTab[82653]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:806
			// _ = "end of CoverTab[82653]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:806
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:806
		// _ = "end of CoverTab[82638]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:806
		_go_fuzz_dep_.CoverTab[82639]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:808
		// _ = "end of CoverTab[82639]"
	case flagType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:809
		_go_fuzz_dep_.CoverTab[82640]++
													v.SetBool(true)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:811
		// _ = "end of CoverTab[82640]"
	case bigIntType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:812
		_go_fuzz_dep_.CoverTab[82641]++
													parsedInt, err1 := parseBigInt(innerBytes)
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:814
			_go_fuzz_dep_.CoverTab[82654]++
														v.Set(reflect.ValueOf(parsedInt))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:815
			// _ = "end of CoverTab[82654]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:816
			_go_fuzz_dep_.CoverTab[82655]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:816
			// _ = "end of CoverTab[82655]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:816
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:816
		// _ = "end of CoverTab[82641]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:816
		_go_fuzz_dep_.CoverTab[82642]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:818
		// _ = "end of CoverTab[82642]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:818
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:818
		_go_fuzz_dep_.CoverTab[82643]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:818
		// _ = "end of CoverTab[82643]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:819
	// _ = "end of CoverTab[82527]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:819
	_go_fuzz_dep_.CoverTab[82528]++
												switch val := v; val.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:821
		_go_fuzz_dep_.CoverTab[82656]++
													parsedBool, err1 := parseBool(innerBytes)
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:823
			_go_fuzz_dep_.CoverTab[82670]++
														val.SetBool(parsedBool)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:824
			// _ = "end of CoverTab[82670]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:825
			_go_fuzz_dep_.CoverTab[82671]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:825
			// _ = "end of CoverTab[82671]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:825
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:825
		// _ = "end of CoverTab[82656]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:825
		_go_fuzz_dep_.CoverTab[82657]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:827
		// _ = "end of CoverTab[82657]"
	case reflect.Int, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:828
		_go_fuzz_dep_.CoverTab[82658]++
													if val.Type().Size() == 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:829
			_go_fuzz_dep_.CoverTab[82672]++
														parsedInt, err1 := parseInt32(innerBytes)
														if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:831
				_go_fuzz_dep_.CoverTab[82674]++
															val.SetInt(int64(parsedInt))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:832
				// _ = "end of CoverTab[82674]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:833
				_go_fuzz_dep_.CoverTab[82675]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:833
				// _ = "end of CoverTab[82675]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:833
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:833
			// _ = "end of CoverTab[82672]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:833
			_go_fuzz_dep_.CoverTab[82673]++
														err = err1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:834
			// _ = "end of CoverTab[82673]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:835
			_go_fuzz_dep_.CoverTab[82676]++
														parsedInt, err1 := parseInt64(innerBytes)
														if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:837
				_go_fuzz_dep_.CoverTab[82678]++
															val.SetInt(parsedInt)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:838
				// _ = "end of CoverTab[82678]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:839
				_go_fuzz_dep_.CoverTab[82679]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:839
				// _ = "end of CoverTab[82679]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:839
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:839
			// _ = "end of CoverTab[82676]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:839
			_go_fuzz_dep_.CoverTab[82677]++
														err = err1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:840
			// _ = "end of CoverTab[82677]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:841
		// _ = "end of CoverTab[82658]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:841
		_go_fuzz_dep_.CoverTab[82659]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:842
		// _ = "end of CoverTab[82659]"

	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:844
		_go_fuzz_dep_.CoverTab[82660]++
													structType := fieldType

													if structType.NumField() > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:847
			_go_fuzz_dep_.CoverTab[82680]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:847
			return structType.Field(0).Type == rawContentsType
														// _ = "end of CoverTab[82680]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:848
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:848
			_go_fuzz_dep_.CoverTab[82681]++
														bytes := bytes[initOffset:offset]
														val.Field(0).Set(reflect.ValueOf(RawContent(bytes)))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:850
			// _ = "end of CoverTab[82681]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:851
			_go_fuzz_dep_.CoverTab[82682]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:851
			// _ = "end of CoverTab[82682]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:851
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:851
		// _ = "end of CoverTab[82660]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:851
		_go_fuzz_dep_.CoverTab[82661]++

													innerOffset := 0
													for i := 0; i < structType.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:854
			_go_fuzz_dep_.CoverTab[82683]++
														field := structType.Field(i)
														if i == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:856
				_go_fuzz_dep_.CoverTab[82685]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:856
				return field.Type == rawContentsType
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:856
				// _ = "end of CoverTab[82685]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:856
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:856
				_go_fuzz_dep_.CoverTab[82686]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:857
				// _ = "end of CoverTab[82686]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:858
				_go_fuzz_dep_.CoverTab[82687]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:858
				// _ = "end of CoverTab[82687]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:858
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:858
			// _ = "end of CoverTab[82683]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:858
			_go_fuzz_dep_.CoverTab[82684]++
														innerOffset, err = parseField(val.Field(i), innerBytes, innerOffset, parseFieldParameters(field.Tag.Get("asn1")))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:860
				_go_fuzz_dep_.CoverTab[82688]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:861
				// _ = "end of CoverTab[82688]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:862
				_go_fuzz_dep_.CoverTab[82689]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:862
				// _ = "end of CoverTab[82689]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:862
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:862
			// _ = "end of CoverTab[82684]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:863
		// _ = "end of CoverTab[82661]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:863
		_go_fuzz_dep_.CoverTab[82662]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:867
		return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:867
		// _ = "end of CoverTab[82662]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:868
		_go_fuzz_dep_.CoverTab[82663]++
													sliceType := fieldType
													if sliceType.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:870
			_go_fuzz_dep_.CoverTab[82690]++
														val.Set(reflect.MakeSlice(sliceType, len(innerBytes), len(innerBytes)))
														reflect.Copy(val, reflect.ValueOf(innerBytes))
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:873
			// _ = "end of CoverTab[82690]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:874
			_go_fuzz_dep_.CoverTab[82691]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:874
			// _ = "end of CoverTab[82691]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:874
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:874
		// _ = "end of CoverTab[82663]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:874
		_go_fuzz_dep_.CoverTab[82664]++
													newSlice, err1 := parseSequenceOf(innerBytes, sliceType, sliceType.Elem())
													if err1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:876
			_go_fuzz_dep_.CoverTab[82692]++
														val.Set(newSlice)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:877
			// _ = "end of CoverTab[82692]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:878
			_go_fuzz_dep_.CoverTab[82693]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:878
			// _ = "end of CoverTab[82693]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:878
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:878
		// _ = "end of CoverTab[82664]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:878
		_go_fuzz_dep_.CoverTab[82665]++
													err = err1
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:880
		// _ = "end of CoverTab[82665]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:881
		_go_fuzz_dep_.CoverTab[82666]++
													var v string
													switch universalTag {
		case TagPrintableString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:884
			_go_fuzz_dep_.CoverTab[82694]++
														v, err = parsePrintableString(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:885
			// _ = "end of CoverTab[82694]"
		case TagIA5String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:886
			_go_fuzz_dep_.CoverTab[82695]++
														v, err = parseIA5String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:887
			// _ = "end of CoverTab[82695]"
		case TagT61String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:888
			_go_fuzz_dep_.CoverTab[82696]++
														v, err = parseT61String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:889
			// _ = "end of CoverTab[82696]"
		case TagUTF8String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:890
			_go_fuzz_dep_.CoverTab[82697]++
														v, err = parseUTF8String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:891
			// _ = "end of CoverTab[82697]"
		case TagGeneralString:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:892
			_go_fuzz_dep_.CoverTab[82698]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:897
			v, err = parseT61String(innerBytes)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:897
			// _ = "end of CoverTab[82698]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:898
			_go_fuzz_dep_.CoverTab[82699]++
														err = SyntaxError{fmt.Sprintf("internal error: unknown string type %d", universalTag)}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:899
			// _ = "end of CoverTab[82699]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:900
		// _ = "end of CoverTab[82666]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:900
		_go_fuzz_dep_.CoverTab[82667]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:901
			_go_fuzz_dep_.CoverTab[82700]++
														val.SetString(v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:902
			// _ = "end of CoverTab[82700]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:903
			_go_fuzz_dep_.CoverTab[82701]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:903
			// _ = "end of CoverTab[82701]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:903
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:903
		// _ = "end of CoverTab[82667]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:903
		_go_fuzz_dep_.CoverTab[82668]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:904
		// _ = "end of CoverTab[82668]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:904
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:904
		_go_fuzz_dep_.CoverTab[82669]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:904
		// _ = "end of CoverTab[82669]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:905
	// _ = "end of CoverTab[82528]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:905
	_go_fuzz_dep_.CoverTab[82529]++
												err = StructuralError{"unsupported: " + v.Type().String()}
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:907
	// _ = "end of CoverTab[82529]"
}

// canHaveDefaultValue reports whether k is a Kind that we will set a default
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:910
// value for. (A signed integer, essentially.)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:912
func canHaveDefaultValue(k reflect.Kind) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:912
	_go_fuzz_dep_.CoverTab[82702]++
												switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:914
		_go_fuzz_dep_.CoverTab[82704]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:915
		// _ = "end of CoverTab[82704]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:915
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:915
		_go_fuzz_dep_.CoverTab[82705]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:915
		// _ = "end of CoverTab[82705]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:916
	// _ = "end of CoverTab[82702]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:916
	_go_fuzz_dep_.CoverTab[82703]++

												return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:918
	// _ = "end of CoverTab[82703]"
}

// setDefaultValue is used to install a default value, from a tag string, into
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:921
// a Value. It is successful if the field was optional, even if a default value
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:921
// wasn't provided or it failed to install it into the Value.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:924
func setDefaultValue(v reflect.Value, params fieldParameters) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:924
	_go_fuzz_dep_.CoverTab[82706]++
												if !params.optional {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:925
		_go_fuzz_dep_.CoverTab[82710]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:926
		// _ = "end of CoverTab[82710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:927
		_go_fuzz_dep_.CoverTab[82711]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:927
		// _ = "end of CoverTab[82711]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:927
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:927
	// _ = "end of CoverTab[82706]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:927
	_go_fuzz_dep_.CoverTab[82707]++
												ok = true
												if params.defaultValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:929
		_go_fuzz_dep_.CoverTab[82712]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:930
		// _ = "end of CoverTab[82712]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:931
		_go_fuzz_dep_.CoverTab[82713]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:931
		// _ = "end of CoverTab[82713]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:931
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:931
	// _ = "end of CoverTab[82707]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:931
	_go_fuzz_dep_.CoverTab[82708]++
												if canHaveDefaultValue(v.Kind()) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:932
		_go_fuzz_dep_.CoverTab[82714]++
													v.SetInt(*params.defaultValue)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:933
		// _ = "end of CoverTab[82714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:934
		_go_fuzz_dep_.CoverTab[82715]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:934
		// _ = "end of CoverTab[82715]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:934
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:934
	// _ = "end of CoverTab[82708]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:934
	_go_fuzz_dep_.CoverTab[82709]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:935
	// _ = "end of CoverTab[82709]"
}

// Unmarshal parses the DER-encoded ASN.1 data structure b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// and uses the reflect package to fill in an arbitrary value pointed at by val.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// Because Unmarshal uses the reflect package, the structs
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// being written to must use upper case field names.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 INTEGER can be written to an int, int32, int64,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// or *big.Int (from the math/big package).
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// If the encoded value does not fit in the Go type,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// Unmarshal returns a parse error.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 BIT STRING can be written to a BitString.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 OCTET STRING can be written to a []byte.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 OBJECT IDENTIFIER can be written to an
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// ObjectIdentifier.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 ENUMERATED can be written to an Enumerated.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a time.Time.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 PrintableString or IA5String can be written to a string.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// Any of the above ASN.1 values can be written to an interface{}.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// The value stored in the interface has the corresponding Go type.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// For integers, that type is int64.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 SEQUENCE OF x or SET OF x can be written
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// to a slice if an x can be written to the slice's element type.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// An ASN.1 SEQUENCE or SET can be written to a struct
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// if each of the elements in the sequence can be
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// written to the corresponding element in the struct.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// The following tags on struct fields have special meaning to Unmarshal:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	application	specifies that a APPLICATION tag is used
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	default:x	sets the default value for optional integer fields
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	explicit	specifies that an additional, explicit tag wraps the implicit one
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	optional	marks the field as ASN.1 OPTIONAL
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	set		causes a SET, rather than a SEQUENCE type to be expected
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//	tag:x		specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// If the type of the first field of a structure is RawContent then the raw
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// ASN1 contents of the struct will be stored in it.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// If the type name of a slice element ends with "SET" then it's treated as if
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// the "set" tag was set on it. This can be used with nested slices where a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// struct tag cannot be given.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// Other ASN.1 types are not supported; if it encounters them,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:938
// Unmarshal returns a parse error.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:990
func Unmarshal(b []byte, val interface{}) (rest []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:990
	_go_fuzz_dep_.CoverTab[82716]++
												return UnmarshalWithParams(b, val, "")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:991
	// _ = "end of CoverTab[82716]"
}

// UnmarshalWithParams allows field parameters to be specified for the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:994
// top-level element. The form of the params is the same as the field tags.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:996
func UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:996
	_go_fuzz_dep_.CoverTab[82717]++
												v := reflect.ValueOf(val).Elem()
												offset, err := parseField(v, b, 0, parseFieldParameters(params))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:999
			_go_fuzz_dep_.CoverTab[82719]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1000
		// _ = "end of CoverTab[82719]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1001
		_go_fuzz_dep_.CoverTab[82720]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1001
		// _ = "end of CoverTab[82720]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1001
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1001
	// _ = "end of CoverTab[82717]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1001
	_go_fuzz_dep_.CoverTab[82718]++
													return b[offset:], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1002
	// _ = "end of CoverTab[82718]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1003
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/asn1.go:1003
var _ = _go_fuzz_dep_.CoverTab
