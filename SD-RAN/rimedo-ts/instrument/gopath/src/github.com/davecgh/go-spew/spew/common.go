//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:17
)

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
)

// Some constants in the form of bytes to avoid string overhead.  This mirrors
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:28
// the technique used in the fmt package.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:30
var (
	panicBytes		= []byte("(PANIC=")
	plusBytes		= []byte("+")
	iBytes			= []byte("i")
	trueBytes		= []byte("true")
	falseBytes		= []byte("false")
	interfaceBytes		= []byte("(interface {})")
	commaNewlineBytes	= []byte(",\n")
	newlineBytes		= []byte("\n")
	openBraceBytes		= []byte("{")
	openBraceNewlineBytes	= []byte("{\n")
	closeBraceBytes		= []byte("}")
	asteriskBytes		= []byte("*")
	colonBytes		= []byte(":")
	colonSpaceBytes		= []byte(": ")
	openParenBytes		= []byte("(")
	closeParenBytes		= []byte(")")
	spaceBytes		= []byte(" ")
	pointerChainBytes	= []byte("->")
	nilAngleBytes		= []byte("<nil>")
	maxNewlineBytes		= []byte("<max depth reached>\n")
	maxShortBytes		= []byte("<max>")
	circularBytes		= []byte("<already shown>")
	circularShortBytes	= []byte("<shown>")
	invalidAngleBytes	= []byte("<invalid>")
	openBracketBytes	= []byte("[")
	closeBracketBytes	= []byte("]")
	percentBytes		= []byte("%")
	precisionBytes		= []byte(".")
	openAngleBytes		= []byte("<")
	closeAngleBytes		= []byte(">")
	openMapBytes		= []byte("map[")
	closeMapBytes		= []byte("]")
	lenEqualsBytes		= []byte("len=")
	capEqualsBytes		= []byte("cap=")
)

// hexDigits is used to map a decimal value to a hex digit.
var hexDigits = "0123456789abcdef"

// catchPanic handles any panics that might occur during the handleMethods
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:70
// calls.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:72
func catchPanic(w io.Writer, v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:72
	_go_fuzz_dep_.CoverTab[81541]++
											if err := recover(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:73
		_go_fuzz_dep_.CoverTab[81542]++
												w.Write(panicBytes)
												fmt.Fprintf(w, "%v", err)
												w.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:76
		// _ = "end of CoverTab[81542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:77
		_go_fuzz_dep_.CoverTab[81543]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:77
		// _ = "end of CoverTab[81543]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:77
	// _ = "end of CoverTab[81541]"
}

// handleMethods attempts to call the Error and String methods on the underlying
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:80
// type the passed reflect.Value represents and outputes the result to Writer w.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:80
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:80
// It handles panics in any called methods by catching and displaying the error
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:80
// as the formatted value.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:85
func handleMethods(cs *ConfigState, w io.Writer, v reflect.Value) (handled bool) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:85
	_go_fuzz_dep_.CoverTab[81544]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:92
	if !v.CanInterface() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:92
		_go_fuzz_dep_.CoverTab[81549]++
												if UnsafeDisabled {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:93
			_go_fuzz_dep_.CoverTab[81551]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:94
			// _ = "end of CoverTab[81551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:95
			_go_fuzz_dep_.CoverTab[81552]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:95
			// _ = "end of CoverTab[81552]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:95
		// _ = "end of CoverTab[81549]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:95
		_go_fuzz_dep_.CoverTab[81550]++

												v = unsafeReflectValue(v)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:97
		// _ = "end of CoverTab[81550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:98
		_go_fuzz_dep_.CoverTab[81553]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:98
		// _ = "end of CoverTab[81553]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:98
	// _ = "end of CoverTab[81544]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:98
	_go_fuzz_dep_.CoverTab[81545]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
	if !cs.DisablePointerMethods && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		_go_fuzz_dep_.CoverTab[81554]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		return !UnsafeDisabled
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		// _ = "end of CoverTab[81554]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		_go_fuzz_dep_.CoverTab[81555]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		return !v.CanAddr()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		// _ = "end of CoverTab[81555]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:106
		_go_fuzz_dep_.CoverTab[81556]++
												v = unsafeReflectValue(v)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:107
		// _ = "end of CoverTab[81556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:108
		_go_fuzz_dep_.CoverTab[81557]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:108
		// _ = "end of CoverTab[81557]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:108
	// _ = "end of CoverTab[81545]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:108
	_go_fuzz_dep_.CoverTab[81546]++
											if v.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:109
		_go_fuzz_dep_.CoverTab[81558]++
												v = v.Addr()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:110
		// _ = "end of CoverTab[81558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:111
		_go_fuzz_dep_.CoverTab[81559]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:111
		// _ = "end of CoverTab[81559]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:111
	// _ = "end of CoverTab[81546]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:111
	_go_fuzz_dep_.CoverTab[81547]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:114
	switch iface := v.Interface().(type) {
	case error:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:115
		_go_fuzz_dep_.CoverTab[81560]++
												defer catchPanic(w, v)
												if cs.ContinueOnMethod {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:117
			_go_fuzz_dep_.CoverTab[81564]++
													w.Write(openParenBytes)
													w.Write([]byte(iface.Error()))
													w.Write(closeParenBytes)
													w.Write(spaceBytes)
													return false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:122
			// _ = "end of CoverTab[81564]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:123
			_go_fuzz_dep_.CoverTab[81565]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:123
			// _ = "end of CoverTab[81565]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:123
		// _ = "end of CoverTab[81560]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:123
		_go_fuzz_dep_.CoverTab[81561]++

												w.Write([]byte(iface.Error()))
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:126
		// _ = "end of CoverTab[81561]"

	case fmt.Stringer:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:128
		_go_fuzz_dep_.CoverTab[81562]++
												defer catchPanic(w, v)
												if cs.ContinueOnMethod {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:130
			_go_fuzz_dep_.CoverTab[81566]++
													w.Write(openParenBytes)
													w.Write([]byte(iface.String()))
													w.Write(closeParenBytes)
													w.Write(spaceBytes)
													return false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:135
			// _ = "end of CoverTab[81566]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:136
			_go_fuzz_dep_.CoverTab[81567]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:136
			// _ = "end of CoverTab[81567]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:136
		// _ = "end of CoverTab[81562]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:136
		_go_fuzz_dep_.CoverTab[81563]++
												w.Write([]byte(iface.String()))
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:138
		// _ = "end of CoverTab[81563]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:139
	// _ = "end of CoverTab[81547]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:139
	_go_fuzz_dep_.CoverTab[81548]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:140
	// _ = "end of CoverTab[81548]"
}

// printBool outputs a boolean value as true or false to Writer w.
func printBool(w io.Writer, val bool) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:144
	_go_fuzz_dep_.CoverTab[81568]++
											if val {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:145
		_go_fuzz_dep_.CoverTab[81569]++
												w.Write(trueBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:146
		// _ = "end of CoverTab[81569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:147
		_go_fuzz_dep_.CoverTab[81570]++
												w.Write(falseBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:148
		// _ = "end of CoverTab[81570]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:149
	// _ = "end of CoverTab[81568]"
}

// printInt outputs a signed integer value to Writer w.
func printInt(w io.Writer, val int64, base int) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:153
	_go_fuzz_dep_.CoverTab[81571]++
											w.Write([]byte(strconv.FormatInt(val, base)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:154
	// _ = "end of CoverTab[81571]"
}

// printUint outputs an unsigned integer value to Writer w.
func printUint(w io.Writer, val uint64, base int) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:158
	_go_fuzz_dep_.CoverTab[81572]++
											w.Write([]byte(strconv.FormatUint(val, base)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:159
	// _ = "end of CoverTab[81572]"
}

// printFloat outputs a floating point value using the specified precision,
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:162
// which is expected to be 32 or 64bit, to Writer w.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:164
func printFloat(w io.Writer, val float64, precision int) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:164
	_go_fuzz_dep_.CoverTab[81573]++
											w.Write([]byte(strconv.FormatFloat(val, 'g', -1, precision)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:165
	// _ = "end of CoverTab[81573]"
}

// printComplex outputs a complex value using the specified float precision
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:168
// for the real and imaginary parts to Writer w.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:170
func printComplex(w io.Writer, c complex128, floatPrecision int) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:170
	_go_fuzz_dep_.CoverTab[81574]++
											r := real(c)
											w.Write(openParenBytes)
											w.Write([]byte(strconv.FormatFloat(r, 'g', -1, floatPrecision)))
											i := imag(c)
											if i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:175
		_go_fuzz_dep_.CoverTab[81576]++
												w.Write(plusBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:176
		// _ = "end of CoverTab[81576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:177
		_go_fuzz_dep_.CoverTab[81577]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:177
		// _ = "end of CoverTab[81577]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:177
	// _ = "end of CoverTab[81574]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:177
	_go_fuzz_dep_.CoverTab[81575]++
											w.Write([]byte(strconv.FormatFloat(i, 'g', -1, floatPrecision)))
											w.Write(iBytes)
											w.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:180
	// _ = "end of CoverTab[81575]"
}

// printHexPtr outputs a uintptr formatted as hexadecimal with a leading '0x'
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:183
// prefix to Writer w.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:185
func printHexPtr(w io.Writer, p uintptr) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:185
	_go_fuzz_dep_.CoverTab[81578]++

											num := uint64(p)
											if num == 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:188
		_go_fuzz_dep_.CoverTab[81581]++
												w.Write(nilAngleBytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:190
		// _ = "end of CoverTab[81581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:191
		_go_fuzz_dep_.CoverTab[81582]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:191
		// _ = "end of CoverTab[81582]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:191
	// _ = "end of CoverTab[81578]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:191
	_go_fuzz_dep_.CoverTab[81579]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:194
	buf := make([]byte, 18)

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:197
	base := uint64(16)
	i := len(buf) - 1
	for num >= base {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:199
		_go_fuzz_dep_.CoverTab[81583]++
												buf[i] = hexDigits[num%base]
												num /= base
												i--
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:202
		// _ = "end of CoverTab[81583]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:203
	// _ = "end of CoverTab[81579]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:203
	_go_fuzz_dep_.CoverTab[81580]++
											buf[i] = hexDigits[num]

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:207
	i--
											buf[i] = 'x'
											i--
											buf[i] = '0'

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:213
	buf = buf[i:]
											w.Write(buf)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:214
	// _ = "end of CoverTab[81580]"
}

// valuesSorter implements sort.Interface to allow a slice of reflect.Value
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:217
// elements to be sorted.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:219
type valuesSorter struct {
	values	[]reflect.Value
	strings	[]string	// either nil or same len and values
	cs	*ConfigState
}

// newValuesSorter initializes a valuesSorter instance, which holds a set of
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:225
// surrogate keys on which the data should be sorted.  It uses flags in
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:225
// ConfigState to decide if and how to populate those surrogate keys.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:228
func newValuesSorter(values []reflect.Value, cs *ConfigState) sort.Interface {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:228
	_go_fuzz_dep_.CoverTab[81584]++
											vs := &valuesSorter{values: values, cs: cs}
											if canSortSimply(vs.values[0].Kind()) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:230
		_go_fuzz_dep_.CoverTab[81588]++
												return vs
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:231
		// _ = "end of CoverTab[81588]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:232
		_go_fuzz_dep_.CoverTab[81589]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:232
		// _ = "end of CoverTab[81589]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:232
	// _ = "end of CoverTab[81584]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:232
	_go_fuzz_dep_.CoverTab[81585]++
											if !cs.DisableMethods {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:233
		_go_fuzz_dep_.CoverTab[81590]++
												vs.strings = make([]string, len(values))
												for i := range vs.values {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:235
			_go_fuzz_dep_.CoverTab[81591]++
													b := bytes.Buffer{}
													if !handleMethods(cs, &b, vs.values[i]) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:237
				_go_fuzz_dep_.CoverTab[81593]++
														vs.strings = nil
														break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:239
				// _ = "end of CoverTab[81593]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:240
				_go_fuzz_dep_.CoverTab[81594]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:240
				// _ = "end of CoverTab[81594]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:240
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:240
			// _ = "end of CoverTab[81591]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:240
			_go_fuzz_dep_.CoverTab[81592]++
													vs.strings[i] = b.String()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:241
			// _ = "end of CoverTab[81592]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:242
		// _ = "end of CoverTab[81590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:243
		_go_fuzz_dep_.CoverTab[81595]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:243
		// _ = "end of CoverTab[81595]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:243
	// _ = "end of CoverTab[81585]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:243
	_go_fuzz_dep_.CoverTab[81586]++
											if vs.strings == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:244
		_go_fuzz_dep_.CoverTab[81596]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:244
		return cs.SpewKeys
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:244
		// _ = "end of CoverTab[81596]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:244
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:244
		_go_fuzz_dep_.CoverTab[81597]++
												vs.strings = make([]string, len(values))
												for i := range vs.values {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:246
			_go_fuzz_dep_.CoverTab[81598]++
													vs.strings[i] = Sprintf("%#v", vs.values[i].Interface())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:247
			// _ = "end of CoverTab[81598]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:248
		// _ = "end of CoverTab[81597]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:249
		_go_fuzz_dep_.CoverTab[81599]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:249
		// _ = "end of CoverTab[81599]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:249
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:249
	// _ = "end of CoverTab[81586]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:249
	_go_fuzz_dep_.CoverTab[81587]++
											return vs
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:250
	// _ = "end of CoverTab[81587]"
}

// canSortSimply tests whether a reflect.Kind is a primitive that can be sorted
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:253
// directly, or whether it should be considered for sorting by surrogate keys
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:253
// (if the ConfigState allows it).
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:256
func canSortSimply(kind reflect.Kind) bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:256
	_go_fuzz_dep_.CoverTab[81600]++

											switch kind {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:259
		_go_fuzz_dep_.CoverTab[81602]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:260
		// _ = "end of CoverTab[81602]"
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:261
		_go_fuzz_dep_.CoverTab[81603]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:262
		// _ = "end of CoverTab[81603]"
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:263
		_go_fuzz_dep_.CoverTab[81604]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:264
		// _ = "end of CoverTab[81604]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:265
		_go_fuzz_dep_.CoverTab[81605]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:266
		// _ = "end of CoverTab[81605]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:267
		_go_fuzz_dep_.CoverTab[81606]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:268
		// _ = "end of CoverTab[81606]"
	case reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:269
		_go_fuzz_dep_.CoverTab[81607]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:270
		// _ = "end of CoverTab[81607]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:271
		_go_fuzz_dep_.CoverTab[81608]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:272
		// _ = "end of CoverTab[81608]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:272
	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:272
		_go_fuzz_dep_.CoverTab[81609]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:272
		// _ = "end of CoverTab[81609]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:273
	// _ = "end of CoverTab[81600]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:273
	_go_fuzz_dep_.CoverTab[81601]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:274
	// _ = "end of CoverTab[81601]"
}

// Len returns the number of values in the slice.  It is part of the
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:277
// sort.Interface implementation.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:279
func (s *valuesSorter) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:279
	_go_fuzz_dep_.CoverTab[81610]++
											return len(s.values)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:280
	// _ = "end of CoverTab[81610]"
}

// Swap swaps the values at the passed indices.  It is part of the
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:283
// sort.Interface implementation.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:285
func (s *valuesSorter) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:285
	_go_fuzz_dep_.CoverTab[81611]++
											s.values[i], s.values[j] = s.values[j], s.values[i]
											if s.strings != nil {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:287
		_go_fuzz_dep_.CoverTab[81612]++
												s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:288
		// _ = "end of CoverTab[81612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:289
		_go_fuzz_dep_.CoverTab[81613]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:289
		// _ = "end of CoverTab[81613]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:289
	// _ = "end of CoverTab[81611]"
}

// valueSortLess returns whether the first value should sort before the second
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:292
// value.  It is used by valueSorter.Less as part of the sort.Interface
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:292
// implementation.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:295
func valueSortLess(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:295
	_go_fuzz_dep_.CoverTab[81614]++
											switch a.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:297
		_go_fuzz_dep_.CoverTab[81616]++
												return !a.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:298
			_go_fuzz_dep_.CoverTab[81624]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:298
			return b.Bool()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:298
			// _ = "end of CoverTab[81624]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:298
		}()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:298
		// _ = "end of CoverTab[81616]"
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:299
		_go_fuzz_dep_.CoverTab[81617]++
												return a.Int() < b.Int()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:300
		// _ = "end of CoverTab[81617]"
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:301
		_go_fuzz_dep_.CoverTab[81618]++
												return a.Uint() < b.Uint()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:302
		// _ = "end of CoverTab[81618]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:303
		_go_fuzz_dep_.CoverTab[81619]++
												return a.Float() < b.Float()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:304
		// _ = "end of CoverTab[81619]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:305
		_go_fuzz_dep_.CoverTab[81620]++
												return a.String() < b.String()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:306
		// _ = "end of CoverTab[81620]"
	case reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:307
		_go_fuzz_dep_.CoverTab[81621]++
												return a.Uint() < b.Uint()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:308
		// _ = "end of CoverTab[81621]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:309
		_go_fuzz_dep_.CoverTab[81622]++

												l := a.Len()
												for i := 0; i < l; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:312
			_go_fuzz_dep_.CoverTab[81625]++
													av := a.Index(i)
													bv := b.Index(i)
													if av.Interface() == bv.Interface() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:315
				_go_fuzz_dep_.CoverTab[81627]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:316
				// _ = "end of CoverTab[81627]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:317
				_go_fuzz_dep_.CoverTab[81628]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:317
				// _ = "end of CoverTab[81628]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:317
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:317
			// _ = "end of CoverTab[81625]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:317
			_go_fuzz_dep_.CoverTab[81626]++
													return valueSortLess(av, bv)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:318
			// _ = "end of CoverTab[81626]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:319
		// _ = "end of CoverTab[81622]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:319
	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:319
		_go_fuzz_dep_.CoverTab[81623]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:319
		// _ = "end of CoverTab[81623]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:320
	// _ = "end of CoverTab[81614]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:320
	_go_fuzz_dep_.CoverTab[81615]++
											return a.String() < b.String()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:321
	// _ = "end of CoverTab[81615]"
}

// Less returns whether the value at index i should sort before the
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:324
// value at index j.  It is part of the sort.Interface implementation.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:326
func (s *valuesSorter) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:326
	_go_fuzz_dep_.CoverTab[81629]++
											if s.strings == nil {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:327
		_go_fuzz_dep_.CoverTab[81631]++
												return valueSortLess(s.values[i], s.values[j])
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:328
		// _ = "end of CoverTab[81631]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:329
		_go_fuzz_dep_.CoverTab[81632]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:329
		// _ = "end of CoverTab[81632]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:329
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:329
	// _ = "end of CoverTab[81629]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:329
	_go_fuzz_dep_.CoverTab[81630]++
											return s.strings[i] < s.strings[j]
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:330
	// _ = "end of CoverTab[81630]"
}

// sortValues is a sort function that handles both native types and any type that
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:333
// can be converted to error or Stringer.  Other inputs are sorted according to
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:333
// their Value.String() value to ensure display stability.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:336
func sortValues(values []reflect.Value, cs *ConfigState) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:336
	_go_fuzz_dep_.CoverTab[81633]++
											if len(values) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:337
		_go_fuzz_dep_.CoverTab[81635]++
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:338
		// _ = "end of CoverTab[81635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:339
		_go_fuzz_dep_.CoverTab[81636]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:339
		// _ = "end of CoverTab[81636]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:339
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:339
	// _ = "end of CoverTab[81633]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:339
	_go_fuzz_dep_.CoverTab[81634]++
											sort.Sort(newValuesSorter(values, cs))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:340
	// _ = "end of CoverTab[81634]"
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:341
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/common.go:341
var _ = _go_fuzz_dep_.CoverTab
