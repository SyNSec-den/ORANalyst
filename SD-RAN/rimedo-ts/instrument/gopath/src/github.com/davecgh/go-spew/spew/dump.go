//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:17
)

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	// uint8Type is a reflect.Type representing a uint8.  It is used to
	// convert cgo types to uint8 slices for hexdumping.
	uint8Type	= reflect.TypeOf(uint8(0))

	// cCharRE is a regular expression that matches a cgo char.
	// It is used to detect character arrays to hexdump them.
	cCharRE	= regexp.MustCompile(`^.*\._Ctype_char$`)

	// cUnsignedCharRE is a regular expression that matches a cgo unsigned
	// char.  It is used to detect unsigned character arrays to hexdump
	// them.
	cUnsignedCharRE	= regexp.MustCompile(`^.*\._Ctype_unsignedchar$`)

	// cUint8tCharRE is a regular expression that matches a cgo uint8_t.
	// It is used to detect uint8_t arrays to hexdump them.
	cUint8tCharRE	= regexp.MustCompile(`^.*\._Ctype_uint8_t$`)
)

// dumpState contains information about the state of a dump operation.
type dumpState struct {
	w			io.Writer
	depth			int
	pointers		map[uintptr]int
	ignoreNextType		bool
	ignoreNextIndent	bool
	cs			*ConfigState
}

// indent performs indentation according to the depth level and cs.Indent
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:60
// option.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:62
func (d *dumpState) indent() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:62
	_go_fuzz_dep_.CoverTab[81655]++
											if d.ignoreNextIndent {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:63
		_go_fuzz_dep_.CoverTab[81657]++
												d.ignoreNextIndent = false
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:65
		// _ = "end of CoverTab[81657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:66
		_go_fuzz_dep_.CoverTab[81658]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:66
		// _ = "end of CoverTab[81658]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:66
	// _ = "end of CoverTab[81655]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:66
	_go_fuzz_dep_.CoverTab[81656]++
											d.w.Write(bytes.Repeat([]byte(d.cs.Indent), d.depth))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:67
	// _ = "end of CoverTab[81656]"
}

// unpackValue returns values inside of non-nil interfaces when possible.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:70
// This is useful for data types like structs, arrays, slices, and maps which
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:70
// can contain varying types packed inside an interface.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:73
func (d *dumpState) unpackValue(v reflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:73
	_go_fuzz_dep_.CoverTab[81659]++
											if v.Kind() == reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:74
		_go_fuzz_dep_.CoverTab[81661]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:74
		return !v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:74
		// _ = "end of CoverTab[81661]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:74
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:74
		_go_fuzz_dep_.CoverTab[81662]++
												v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:75
		// _ = "end of CoverTab[81662]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:76
		_go_fuzz_dep_.CoverTab[81663]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:76
		// _ = "end of CoverTab[81663]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:76
	// _ = "end of CoverTab[81659]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:76
	_go_fuzz_dep_.CoverTab[81660]++
											return v
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:77
	// _ = "end of CoverTab[81660]"
}

// dumpPtr handles formatting of pointers by indirecting them as necessary.
func (d *dumpState) dumpPtr(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:81
	_go_fuzz_dep_.CoverTab[81664]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:84
	for k, depth := range d.pointers {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:84
		_go_fuzz_dep_.CoverTab[81669]++
												if depth >= d.depth {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:85
			_go_fuzz_dep_.CoverTab[81670]++
													delete(d.pointers, k)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:86
			// _ = "end of CoverTab[81670]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:87
			_go_fuzz_dep_.CoverTab[81671]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:87
			// _ = "end of CoverTab[81671]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:87
		// _ = "end of CoverTab[81669]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:88
	// _ = "end of CoverTab[81664]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:88
	_go_fuzz_dep_.CoverTab[81665]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:91
	pointerChain := make([]uintptr, 0)

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:96
	nilFound := false
	cycleFound := false
	indirects := 0
	ve := v
	for ve.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:100
		_go_fuzz_dep_.CoverTab[81672]++
												if ve.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:101
			_go_fuzz_dep_.CoverTab[81675]++
													nilFound = true
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:103
			// _ = "end of CoverTab[81675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:104
			_go_fuzz_dep_.CoverTab[81676]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:104
			// _ = "end of CoverTab[81676]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:104
		// _ = "end of CoverTab[81672]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:104
		_go_fuzz_dep_.CoverTab[81673]++
												indirects++
												addr := ve.Pointer()
												pointerChain = append(pointerChain, addr)
												if pd, ok := d.pointers[addr]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:108
			_go_fuzz_dep_.CoverTab[81677]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:108
			return pd < d.depth
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:108
			// _ = "end of CoverTab[81677]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:108
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:108
			_go_fuzz_dep_.CoverTab[81678]++
													cycleFound = true
													indirects--
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:111
			// _ = "end of CoverTab[81678]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:112
			_go_fuzz_dep_.CoverTab[81679]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:112
			// _ = "end of CoverTab[81679]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:112
		// _ = "end of CoverTab[81673]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:112
		_go_fuzz_dep_.CoverTab[81674]++
												d.pointers[addr] = d.depth

												ve = ve.Elem()
												if ve.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:116
			_go_fuzz_dep_.CoverTab[81680]++
													if ve.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:117
				_go_fuzz_dep_.CoverTab[81682]++
														nilFound = true
														break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:119
				// _ = "end of CoverTab[81682]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:120
				_go_fuzz_dep_.CoverTab[81683]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:120
				// _ = "end of CoverTab[81683]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:120
			// _ = "end of CoverTab[81680]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:120
			_go_fuzz_dep_.CoverTab[81681]++
													ve = ve.Elem()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:121
			// _ = "end of CoverTab[81681]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:122
			_go_fuzz_dep_.CoverTab[81684]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:122
			// _ = "end of CoverTab[81684]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:122
		// _ = "end of CoverTab[81674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:123
	// _ = "end of CoverTab[81665]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:123
	_go_fuzz_dep_.CoverTab[81666]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:126
	d.w.Write(openParenBytes)
											d.w.Write(bytes.Repeat(asteriskBytes, indirects))
											d.w.Write([]byte(ve.Type().String()))
											d.w.Write(closeParenBytes)

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
	if !d.cs.DisablePointerAddresses && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
		_go_fuzz_dep_.CoverTab[81685]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
		return len(pointerChain) > 0
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
		// _ = "end of CoverTab[81685]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:132
		_go_fuzz_dep_.CoverTab[81686]++
												d.w.Write(openParenBytes)
												for i, addr := range pointerChain {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:134
			_go_fuzz_dep_.CoverTab[81688]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:135
				_go_fuzz_dep_.CoverTab[81690]++
														d.w.Write(pointerChainBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:136
				// _ = "end of CoverTab[81690]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:137
				_go_fuzz_dep_.CoverTab[81691]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:137
				// _ = "end of CoverTab[81691]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:137
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:137
			// _ = "end of CoverTab[81688]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:137
			_go_fuzz_dep_.CoverTab[81689]++
													printHexPtr(d.w, addr)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:138
			// _ = "end of CoverTab[81689]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:139
		// _ = "end of CoverTab[81686]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:139
		_go_fuzz_dep_.CoverTab[81687]++
												d.w.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:140
		// _ = "end of CoverTab[81687]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:141
		_go_fuzz_dep_.CoverTab[81692]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:141
		// _ = "end of CoverTab[81692]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:141
	// _ = "end of CoverTab[81666]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:141
	_go_fuzz_dep_.CoverTab[81667]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:144
	d.w.Write(openParenBytes)
	switch {
	case nilFound:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:146
		_go_fuzz_dep_.CoverTab[81693]++
												d.w.Write(nilAngleBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:147
		// _ = "end of CoverTab[81693]"

	case cycleFound:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:149
		_go_fuzz_dep_.CoverTab[81694]++
												d.w.Write(circularBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:150
		// _ = "end of CoverTab[81694]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:152
		_go_fuzz_dep_.CoverTab[81695]++
												d.ignoreNextType = true
												d.dump(ve)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:154
		// _ = "end of CoverTab[81695]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:155
	// _ = "end of CoverTab[81667]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:155
	_go_fuzz_dep_.CoverTab[81668]++
											d.w.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:156
	// _ = "end of CoverTab[81668]"
}

// dumpSlice handles formatting of arrays and slices.  Byte (uint8 under
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:159
// reflection) arrays and slices are dumped in hexdump -C fashion.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:161
func (d *dumpState) dumpSlice(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:161
	_go_fuzz_dep_.CoverTab[81696]++
	// Determine whether this type should be hex dumped or not.  Also,
	// for types which should be hexdumped, try to use the underlying data
	// first, then fall back to trying to convert them to a uint8 slice.
	var buf []uint8
	doConvert := false
	doHexDump := false
	numEntries := v.Len()
	if numEntries > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:169
		_go_fuzz_dep_.CoverTab[81699]++
												vt := v.Index(0).Type()
												vts := vt.String()
												switch {

		case cCharRE.MatchString(vts):
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:174
			_go_fuzz_dep_.CoverTab[81701]++
													fallthrough
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:175
			// _ = "end of CoverTab[81701]"
		case cUnsignedCharRE.MatchString(vts):
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:176
			_go_fuzz_dep_.CoverTab[81702]++
													fallthrough
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:177
			// _ = "end of CoverTab[81702]"
		case cUint8tCharRE.MatchString(vts):
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:178
			_go_fuzz_dep_.CoverTab[81703]++
													doConvert = true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:179
			// _ = "end of CoverTab[81703]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:183
		case vt.Kind() == reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:183
			_go_fuzz_dep_.CoverTab[81704]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:191
			vs := v
			if !vs.CanInterface() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:192
				_go_fuzz_dep_.CoverTab[81708]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:192
				return !vs.CanAddr()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:192
				// _ = "end of CoverTab[81708]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:192
			}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:192
				_go_fuzz_dep_.CoverTab[81709]++
														vs = unsafeReflectValue(vs)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:193
				// _ = "end of CoverTab[81709]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:194
				_go_fuzz_dep_.CoverTab[81710]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:194
				// _ = "end of CoverTab[81710]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:194
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:194
			// _ = "end of CoverTab[81704]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:194
			_go_fuzz_dep_.CoverTab[81705]++
													if !UnsafeDisabled {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:195
				_go_fuzz_dep_.CoverTab[81711]++
														vs = vs.Slice(0, numEntries)

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:200
				iface := vs.Interface()
				if slice, ok := iface.([]uint8); ok {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:201
					_go_fuzz_dep_.CoverTab[81712]++
															buf = slice
															doHexDump = true
															break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:204
					// _ = "end of CoverTab[81712]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:205
					_go_fuzz_dep_.CoverTab[81713]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:205
					// _ = "end of CoverTab[81713]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:205
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:205
				// _ = "end of CoverTab[81711]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:206
				_go_fuzz_dep_.CoverTab[81714]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:206
				// _ = "end of CoverTab[81714]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:206
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:206
			// _ = "end of CoverTab[81705]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:206
			_go_fuzz_dep_.CoverTab[81706]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:210
			doConvert = true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:210
			// _ = "end of CoverTab[81706]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:210
		default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:210
			_go_fuzz_dep_.CoverTab[81707]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:210
			// _ = "end of CoverTab[81707]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:211
		// _ = "end of CoverTab[81699]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:211
		_go_fuzz_dep_.CoverTab[81700]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
		if doConvert && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
			_go_fuzz_dep_.CoverTab[81715]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
			return vt.ConvertibleTo(uint8Type)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
			// _ = "end of CoverTab[81715]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:214
			_go_fuzz_dep_.CoverTab[81716]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:217
			buf = make([]uint8, numEntries)
			for i := 0; i < numEntries; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:218
				_go_fuzz_dep_.CoverTab[81718]++
														vv := v.Index(i)
														buf[i] = uint8(vv.Convert(uint8Type).Uint())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:220
				// _ = "end of CoverTab[81718]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:221
			// _ = "end of CoverTab[81716]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:221
			_go_fuzz_dep_.CoverTab[81717]++
													doHexDump = true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:222
			// _ = "end of CoverTab[81717]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:223
			_go_fuzz_dep_.CoverTab[81719]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:223
			// _ = "end of CoverTab[81719]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:223
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:223
		// _ = "end of CoverTab[81700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:224
		_go_fuzz_dep_.CoverTab[81720]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:224
		// _ = "end of CoverTab[81720]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:224
	// _ = "end of CoverTab[81696]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:224
	_go_fuzz_dep_.CoverTab[81697]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:227
	if doHexDump {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:227
		_go_fuzz_dep_.CoverTab[81721]++
												indent := strings.Repeat(d.cs.Indent, d.depth)
												str := indent + hex.Dump(buf)
												str = strings.Replace(str, "\n", "\n"+indent, -1)
												str = strings.TrimRight(str, d.cs.Indent)
												d.w.Write([]byte(str))
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:233
		// _ = "end of CoverTab[81721]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:234
		_go_fuzz_dep_.CoverTab[81722]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:234
		// _ = "end of CoverTab[81722]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:234
	// _ = "end of CoverTab[81697]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:234
	_go_fuzz_dep_.CoverTab[81698]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:237
	for i := 0; i < numEntries; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:237
		_go_fuzz_dep_.CoverTab[81723]++
												d.dump(d.unpackValue(v.Index(i)))
												if i < (numEntries - 1) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:239
			_go_fuzz_dep_.CoverTab[81724]++
													d.w.Write(commaNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:240
			// _ = "end of CoverTab[81724]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:241
			_go_fuzz_dep_.CoverTab[81725]++
													d.w.Write(newlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:242
			// _ = "end of CoverTab[81725]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:243
		// _ = "end of CoverTab[81723]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:244
	// _ = "end of CoverTab[81698]"
}

// dump is the main workhorse for dumping a value.  It uses the passed reflect
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:247
// value to figure out what kind of object we are dealing with and formats it
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:247
// appropriately.  It is a recursive function, however circular data structures
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:247
// are detected and handled properly.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:251
func (d *dumpState) dump(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:251
	_go_fuzz_dep_.CoverTab[81726]++

											kind := v.Kind()
											if kind == reflect.Invalid {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:254
		_go_fuzz_dep_.CoverTab[81733]++
												d.w.Write(invalidAngleBytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:256
		// _ = "end of CoverTab[81733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:257
		_go_fuzz_dep_.CoverTab[81734]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:257
		// _ = "end of CoverTab[81734]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:257
	// _ = "end of CoverTab[81726]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:257
	_go_fuzz_dep_.CoverTab[81727]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:260
	if kind == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:260
		_go_fuzz_dep_.CoverTab[81735]++
												d.indent()
												d.dumpPtr(v)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:263
		// _ = "end of CoverTab[81735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:264
		_go_fuzz_dep_.CoverTab[81736]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:264
		// _ = "end of CoverTab[81736]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:264
	// _ = "end of CoverTab[81727]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:264
	_go_fuzz_dep_.CoverTab[81728]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:267
	if !d.ignoreNextType {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:267
		_go_fuzz_dep_.CoverTab[81737]++
												d.indent()
												d.w.Write(openParenBytes)
												d.w.Write([]byte(v.Type().String()))
												d.w.Write(closeParenBytes)
												d.w.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:272
		// _ = "end of CoverTab[81737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:273
		_go_fuzz_dep_.CoverTab[81738]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:273
		// _ = "end of CoverTab[81738]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:273
	// _ = "end of CoverTab[81728]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:273
	_go_fuzz_dep_.CoverTab[81729]++
											d.ignoreNextType = false

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:278
	valueLen, valueCap := 0, 0
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:280
		_go_fuzz_dep_.CoverTab[81739]++
												valueLen, valueCap = v.Len(), v.Cap()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:281
		// _ = "end of CoverTab[81739]"
	case reflect.Map, reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:282
		_go_fuzz_dep_.CoverTab[81740]++
												valueLen = v.Len()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:283
		// _ = "end of CoverTab[81740]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:283
	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:283
		_go_fuzz_dep_.CoverTab[81741]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:283
		// _ = "end of CoverTab[81741]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:284
	// _ = "end of CoverTab[81729]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:284
	_go_fuzz_dep_.CoverTab[81730]++
											if valueLen != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
		_go_fuzz_dep_.CoverTab[81742]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
		return !d.cs.DisableCapacities && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
			_go_fuzz_dep_.CoverTab[81743]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
			return valueCap != 0
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
			// _ = "end of CoverTab[81743]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
		}()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
		// _ = "end of CoverTab[81742]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:285
		_go_fuzz_dep_.CoverTab[81744]++
												d.w.Write(openParenBytes)
												if valueLen != 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:287
			_go_fuzz_dep_.CoverTab[81747]++
													d.w.Write(lenEqualsBytes)
													printInt(d.w, int64(valueLen), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:289
			// _ = "end of CoverTab[81747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:290
			_go_fuzz_dep_.CoverTab[81748]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:290
			// _ = "end of CoverTab[81748]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:290
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:290
		// _ = "end of CoverTab[81744]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:290
		_go_fuzz_dep_.CoverTab[81745]++
												if !d.cs.DisableCapacities && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:291
			_go_fuzz_dep_.CoverTab[81749]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:291
			return valueCap != 0
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:291
			// _ = "end of CoverTab[81749]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:291
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:291
			_go_fuzz_dep_.CoverTab[81750]++
													if valueLen != 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:292
				_go_fuzz_dep_.CoverTab[81752]++
														d.w.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:293
				// _ = "end of CoverTab[81752]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:294
				_go_fuzz_dep_.CoverTab[81753]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:294
				// _ = "end of CoverTab[81753]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:294
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:294
			// _ = "end of CoverTab[81750]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:294
			_go_fuzz_dep_.CoverTab[81751]++
													d.w.Write(capEqualsBytes)
													printInt(d.w, int64(valueCap), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:296
			// _ = "end of CoverTab[81751]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:297
			_go_fuzz_dep_.CoverTab[81754]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:297
			// _ = "end of CoverTab[81754]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:297
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:297
		// _ = "end of CoverTab[81745]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:297
		_go_fuzz_dep_.CoverTab[81746]++
												d.w.Write(closeParenBytes)
												d.w.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:299
		// _ = "end of CoverTab[81746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:300
		_go_fuzz_dep_.CoverTab[81755]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:300
		// _ = "end of CoverTab[81755]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:300
	// _ = "end of CoverTab[81730]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:300
	_go_fuzz_dep_.CoverTab[81731]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:304
	if !d.cs.DisableMethods {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:304
		_go_fuzz_dep_.CoverTab[81756]++
												if (kind != reflect.Invalid) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:305
			_go_fuzz_dep_.CoverTab[81757]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:305
			return (kind != reflect.Interface)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:305
			// _ = "end of CoverTab[81757]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:305
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:305
			_go_fuzz_dep_.CoverTab[81758]++
													if handled := handleMethods(d.cs, d.w, v); handled {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:306
				_go_fuzz_dep_.CoverTab[81759]++
														return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:307
				// _ = "end of CoverTab[81759]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:308
				_go_fuzz_dep_.CoverTab[81760]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:308
				// _ = "end of CoverTab[81760]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:308
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:308
			// _ = "end of CoverTab[81758]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:309
			_go_fuzz_dep_.CoverTab[81761]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:309
			// _ = "end of CoverTab[81761]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:309
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:309
		// _ = "end of CoverTab[81756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:310
		_go_fuzz_dep_.CoverTab[81762]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:310
		// _ = "end of CoverTab[81762]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:310
	// _ = "end of CoverTab[81731]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:310
	_go_fuzz_dep_.CoverTab[81732]++

											switch kind {
	case reflect.Invalid:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:313
		_go_fuzz_dep_.CoverTab[81763]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:313
		// _ = "end of CoverTab[81763]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:317
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:317
		_go_fuzz_dep_.CoverTab[81764]++
												printBool(d.w, v.Bool())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:318
		// _ = "end of CoverTab[81764]"

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:320
		_go_fuzz_dep_.CoverTab[81765]++
												printInt(d.w, v.Int(), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:321
		// _ = "end of CoverTab[81765]"

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:323
		_go_fuzz_dep_.CoverTab[81766]++
												printUint(d.w, v.Uint(), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:324
		// _ = "end of CoverTab[81766]"

	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:326
		_go_fuzz_dep_.CoverTab[81767]++
												printFloat(d.w, v.Float(), 32)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:327
		// _ = "end of CoverTab[81767]"

	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:329
		_go_fuzz_dep_.CoverTab[81768]++
												printFloat(d.w, v.Float(), 64)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:330
		// _ = "end of CoverTab[81768]"

	case reflect.Complex64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:332
		_go_fuzz_dep_.CoverTab[81769]++
												printComplex(d.w, v.Complex(), 32)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:333
		// _ = "end of CoverTab[81769]"

	case reflect.Complex128:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:335
		_go_fuzz_dep_.CoverTab[81770]++
												printComplex(d.w, v.Complex(), 64)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:336
		// _ = "end of CoverTab[81770]"

	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:338
		_go_fuzz_dep_.CoverTab[81771]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:339
			_go_fuzz_dep_.CoverTab[81786]++
													d.w.Write(nilAngleBytes)
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:341
			// _ = "end of CoverTab[81786]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:342
			_go_fuzz_dep_.CoverTab[81787]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:342
			// _ = "end of CoverTab[81787]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:342
		// _ = "end of CoverTab[81771]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:342
		_go_fuzz_dep_.CoverTab[81772]++
												fallthrough
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:343
		// _ = "end of CoverTab[81772]"

	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:345
		_go_fuzz_dep_.CoverTab[81773]++
												d.w.Write(openBraceNewlineBytes)
												d.depth++
												if (d.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:348
			_go_fuzz_dep_.CoverTab[81788]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:348
			return (d.depth > d.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:348
			// _ = "end of CoverTab[81788]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:348
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:348
			_go_fuzz_dep_.CoverTab[81789]++
													d.indent()
													d.w.Write(maxNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:350
			// _ = "end of CoverTab[81789]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:351
			_go_fuzz_dep_.CoverTab[81790]++
													d.dumpSlice(v)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:352
			// _ = "end of CoverTab[81790]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:353
		// _ = "end of CoverTab[81773]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:353
		_go_fuzz_dep_.CoverTab[81774]++
												d.depth--
												d.indent()
												d.w.Write(closeBraceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:356
		// _ = "end of CoverTab[81774]"

	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:358
		_go_fuzz_dep_.CoverTab[81775]++
												d.w.Write([]byte(strconv.Quote(v.String())))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:359
		// _ = "end of CoverTab[81775]"

	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:361
		_go_fuzz_dep_.CoverTab[81776]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:364
		if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:364
			_go_fuzz_dep_.CoverTab[81791]++
													d.w.Write(nilAngleBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:365
			// _ = "end of CoverTab[81791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:366
			_go_fuzz_dep_.CoverTab[81792]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:366
			// _ = "end of CoverTab[81792]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:366
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:366
		// _ = "end of CoverTab[81776]"

	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:368
		_go_fuzz_dep_.CoverTab[81777]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:368
		// _ = "end of CoverTab[81777]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:372
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:372
		_go_fuzz_dep_.CoverTab[81778]++

												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:374
			_go_fuzz_dep_.CoverTab[81793]++
													d.w.Write(nilAngleBytes)
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:376
			// _ = "end of CoverTab[81793]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:377
			_go_fuzz_dep_.CoverTab[81794]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:377
			// _ = "end of CoverTab[81794]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:377
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:377
		// _ = "end of CoverTab[81778]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:377
		_go_fuzz_dep_.CoverTab[81779]++

												d.w.Write(openBraceNewlineBytes)
												d.depth++
												if (d.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:381
			_go_fuzz_dep_.CoverTab[81795]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:381
			return (d.depth > d.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:381
			// _ = "end of CoverTab[81795]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:381
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:381
			_go_fuzz_dep_.CoverTab[81796]++
													d.indent()
													d.w.Write(maxNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:383
			// _ = "end of CoverTab[81796]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:384
			_go_fuzz_dep_.CoverTab[81797]++
													numEntries := v.Len()
													keys := v.MapKeys()
													if d.cs.SortKeys {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:387
				_go_fuzz_dep_.CoverTab[81799]++
														sortValues(keys, d.cs)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:388
				// _ = "end of CoverTab[81799]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:389
				_go_fuzz_dep_.CoverTab[81800]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:389
				// _ = "end of CoverTab[81800]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:389
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:389
			// _ = "end of CoverTab[81797]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:389
			_go_fuzz_dep_.CoverTab[81798]++
													for i, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:390
				_go_fuzz_dep_.CoverTab[81801]++
														d.dump(d.unpackValue(key))
														d.w.Write(colonSpaceBytes)
														d.ignoreNextIndent = true
														d.dump(d.unpackValue(v.MapIndex(key)))
														if i < (numEntries - 1) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:395
					_go_fuzz_dep_.CoverTab[81802]++
															d.w.Write(commaNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:396
					// _ = "end of CoverTab[81802]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:397
					_go_fuzz_dep_.CoverTab[81803]++
															d.w.Write(newlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:398
					// _ = "end of CoverTab[81803]"
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:399
				// _ = "end of CoverTab[81801]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:400
			// _ = "end of CoverTab[81798]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:401
		// _ = "end of CoverTab[81779]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:401
		_go_fuzz_dep_.CoverTab[81780]++
												d.depth--
												d.indent()
												d.w.Write(closeBraceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:404
		// _ = "end of CoverTab[81780]"

	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:406
		_go_fuzz_dep_.CoverTab[81781]++
												d.w.Write(openBraceNewlineBytes)
												d.depth++
												if (d.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:409
			_go_fuzz_dep_.CoverTab[81804]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:409
			return (d.depth > d.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:409
			// _ = "end of CoverTab[81804]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:409
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:409
			_go_fuzz_dep_.CoverTab[81805]++
													d.indent()
													d.w.Write(maxNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:411
			// _ = "end of CoverTab[81805]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:412
			_go_fuzz_dep_.CoverTab[81806]++
													vt := v.Type()
													numFields := v.NumField()
													for i := 0; i < numFields; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:415
				_go_fuzz_dep_.CoverTab[81807]++
														d.indent()
														vtf := vt.Field(i)
														d.w.Write([]byte(vtf.Name))
														d.w.Write(colonSpaceBytes)
														d.ignoreNextIndent = true
														d.dump(d.unpackValue(v.Field(i)))
														if i < (numFields - 1) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:422
					_go_fuzz_dep_.CoverTab[81808]++
															d.w.Write(commaNewlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:423
					// _ = "end of CoverTab[81808]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:424
					_go_fuzz_dep_.CoverTab[81809]++
															d.w.Write(newlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:425
					// _ = "end of CoverTab[81809]"
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:426
				// _ = "end of CoverTab[81807]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:427
			// _ = "end of CoverTab[81806]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:428
		// _ = "end of CoverTab[81781]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:428
		_go_fuzz_dep_.CoverTab[81782]++
												d.depth--
												d.indent()
												d.w.Write(closeBraceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:431
		// _ = "end of CoverTab[81782]"

	case reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:433
		_go_fuzz_dep_.CoverTab[81783]++
												printHexPtr(d.w, uintptr(v.Uint()))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:434
		// _ = "end of CoverTab[81783]"

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:436
		_go_fuzz_dep_.CoverTab[81784]++
												printHexPtr(d.w, v.Pointer())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:437
		// _ = "end of CoverTab[81784]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:442
	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:442
		_go_fuzz_dep_.CoverTab[81785]++
												if v.CanInterface() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:443
			_go_fuzz_dep_.CoverTab[81810]++
													fmt.Fprintf(d.w, "%v", v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:444
			// _ = "end of CoverTab[81810]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:445
			_go_fuzz_dep_.CoverTab[81811]++
													fmt.Fprintf(d.w, "%v", v.String())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:446
			// _ = "end of CoverTab[81811]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:447
		// _ = "end of CoverTab[81785]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:448
	// _ = "end of CoverTab[81732]"
}

// fdump is a helper function to consolidate the logic from the various public
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:451
// methods which take varying writers and config states.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:453
func fdump(cs *ConfigState, w io.Writer, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:453
	_go_fuzz_dep_.CoverTab[81812]++
											for _, arg := range a {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:454
		_go_fuzz_dep_.CoverTab[81813]++
												if arg == nil {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:455
			_go_fuzz_dep_.CoverTab[81815]++
													w.Write(interfaceBytes)
													w.Write(spaceBytes)
													w.Write(nilAngleBytes)
													w.Write(newlineBytes)
													continue
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:460
			// _ = "end of CoverTab[81815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:461
			_go_fuzz_dep_.CoverTab[81816]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:461
			// _ = "end of CoverTab[81816]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:461
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:461
		// _ = "end of CoverTab[81813]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:461
		_go_fuzz_dep_.CoverTab[81814]++

												d := dumpState{w: w, cs: cs}
												d.pointers = make(map[uintptr]int)
												d.dump(reflect.ValueOf(arg))
												d.w.Write(newlineBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:466
		// _ = "end of CoverTab[81814]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:467
	// _ = "end of CoverTab[81812]"
}

// Fdump formats and displays the passed arguments to io.Writer w.  It formats
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:470
// exactly the same as Dump.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:472
func Fdump(w io.Writer, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:472
	_go_fuzz_dep_.CoverTab[81817]++
											fdump(&Config, w, a...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:473
	// _ = "end of CoverTab[81817]"
}

// Sdump returns a string with the passed arguments formatted exactly the same
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:476
// as Dump.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:478
func Sdump(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:478
	_go_fuzz_dep_.CoverTab[81818]++
											var buf bytes.Buffer
											fdump(&Config, &buf, a...)
											return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:481
	// _ = "end of CoverTab[81818]"
}

/*
Dump displays the passed parameters to standard out with newlines, customizable
indentation, and additional debug information such as complete types and all
pointer addresses used to indirect to the final value.  It provides the
following features over the built-in printing facilities provided by the fmt
package:

  - Pointers are dereferenced and followed
  - Circular data structures are detected and handled properly
  - Custom Stringer/error interfaces are optionally invoked, including
    on unexported types
  - Custom types which only implement the Stringer/error interfaces via
    a pointer receiver are optionally invoked when passing non-pointer
    variables
  - Byte arrays and slices are dumped like the hexdump -C command which
    includes offsets, byte values in hex, and ASCII output

The configuration options are controlled by an exported package global,
spew.Config.  See ConfigState for options documentation.

See Fdump if you would prefer dumping to an arbitrary io.Writer or Sdump to
get the formatted result as a string.
*/
func Dump(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:507
	_go_fuzz_dep_.CoverTab[81819]++
											fdump(&Config, os.Stdout, a...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:508
	// _ = "end of CoverTab[81819]"
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:509
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/dump.go:509
var _ = _go_fuzz_dep_.CoverTab
