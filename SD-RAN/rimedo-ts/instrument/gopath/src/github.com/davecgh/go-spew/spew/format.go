//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:17
)

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// supportedFlags is a list of all the character flags supported by fmt package.
const supportedFlags = "0-+# "

// formatState implements the fmt.Formatter interface and contains information
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:30
// about the state of a formatting operation.  The NewFormatter function can
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:30
// be used to get a new Formatter which can be used directly as arguments
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:30
// in standard fmt package printing calls.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:34
type formatState struct {
	value		interface{}
	fs		fmt.State
	depth		int
	pointers	map[uintptr]int
	ignoreNextType	bool
	cs		*ConfigState
}

// buildDefaultFormat recreates the original format string without precision
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:43
// and width information to pass in to fmt.Sprintf in the case of an
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:43
// unrecognized type.  Unless new types are added to the language, this
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:43
// function won't ever be called.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:47
func (f *formatState) buildDefaultFormat() (format string) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:47
	_go_fuzz_dep_.CoverTab[81820]++
											buf := bytes.NewBuffer(percentBytes)

											for _, flag := range supportedFlags {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:50
		_go_fuzz_dep_.CoverTab[81822]++
												if f.fs.Flag(int(flag)) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:51
			_go_fuzz_dep_.CoverTab[81823]++
													buf.WriteRune(flag)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:52
			// _ = "end of CoverTab[81823]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:53
			_go_fuzz_dep_.CoverTab[81824]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:53
			// _ = "end of CoverTab[81824]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:53
		// _ = "end of CoverTab[81822]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:54
	// _ = "end of CoverTab[81820]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:54
	_go_fuzz_dep_.CoverTab[81821]++

											buf.WriteRune('v')

											format = buf.String()
											return format
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:59
	// _ = "end of CoverTab[81821]"
}

// constructOrigFormat recreates the original format string including precision
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:62
// and width information to pass along to the standard fmt package.  This allows
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:62
// automatic deferral of all format strings this package doesn't support.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:65
func (f *formatState) constructOrigFormat(verb rune) (format string) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:65
	_go_fuzz_dep_.CoverTab[81825]++
											buf := bytes.NewBuffer(percentBytes)

											for _, flag := range supportedFlags {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:68
		_go_fuzz_dep_.CoverTab[81829]++
												if f.fs.Flag(int(flag)) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:69
			_go_fuzz_dep_.CoverTab[81830]++
													buf.WriteRune(flag)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:70
			// _ = "end of CoverTab[81830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:71
			_go_fuzz_dep_.CoverTab[81831]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:71
			// _ = "end of CoverTab[81831]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:71
		// _ = "end of CoverTab[81829]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:72
	// _ = "end of CoverTab[81825]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:72
	_go_fuzz_dep_.CoverTab[81826]++

											if width, ok := f.fs.Width(); ok {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:74
		_go_fuzz_dep_.CoverTab[81832]++
												buf.WriteString(strconv.Itoa(width))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:75
		// _ = "end of CoverTab[81832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:76
		_go_fuzz_dep_.CoverTab[81833]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:76
		// _ = "end of CoverTab[81833]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:76
	// _ = "end of CoverTab[81826]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:76
	_go_fuzz_dep_.CoverTab[81827]++

											if precision, ok := f.fs.Precision(); ok {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:78
		_go_fuzz_dep_.CoverTab[81834]++
												buf.Write(precisionBytes)
												buf.WriteString(strconv.Itoa(precision))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:80
		// _ = "end of CoverTab[81834]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:81
		_go_fuzz_dep_.CoverTab[81835]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:81
		// _ = "end of CoverTab[81835]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:81
	// _ = "end of CoverTab[81827]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:81
	_go_fuzz_dep_.CoverTab[81828]++

											buf.WriteRune(verb)

											format = buf.String()
											return format
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:86
	// _ = "end of CoverTab[81828]"
}

// unpackValue returns values inside of non-nil interfaces when possible and
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:89
// ensures that types for values which have been unpacked from an interface
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:89
// are displayed when the show types flag is also set.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:89
// This is useful for data types like structs, arrays, slices, and maps which
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:89
// can contain varying types packed inside an interface.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:94
func (f *formatState) unpackValue(v reflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:94
	_go_fuzz_dep_.CoverTab[81836]++
											if v.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:95
		_go_fuzz_dep_.CoverTab[81838]++
												f.ignoreNextType = false
												if !v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:97
			_go_fuzz_dep_.CoverTab[81839]++
													v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:98
			// _ = "end of CoverTab[81839]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:99
			_go_fuzz_dep_.CoverTab[81840]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:99
			// _ = "end of CoverTab[81840]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:99
		// _ = "end of CoverTab[81838]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:100
		_go_fuzz_dep_.CoverTab[81841]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:100
		// _ = "end of CoverTab[81841]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:100
	// _ = "end of CoverTab[81836]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:100
	_go_fuzz_dep_.CoverTab[81837]++
											return v
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:101
	// _ = "end of CoverTab[81837]"
}

// formatPtr handles formatting of pointers by indirecting them as necessary.
func (f *formatState) formatPtr(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:105
	_go_fuzz_dep_.CoverTab[81842]++

											showTypes := f.fs.Flag('#')
											if v.IsNil() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
		_go_fuzz_dep_.CoverTab[81848]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
		return (!showTypes || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
			_go_fuzz_dep_.CoverTab[81849]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
			return f.ignoreNextType
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
			// _ = "end of CoverTab[81849]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
		}())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
		// _ = "end of CoverTab[81848]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:108
		_go_fuzz_dep_.CoverTab[81850]++
												f.fs.Write(nilAngleBytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:110
		// _ = "end of CoverTab[81850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:111
		_go_fuzz_dep_.CoverTab[81851]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:111
		// _ = "end of CoverTab[81851]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:111
	// _ = "end of CoverTab[81842]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:111
	_go_fuzz_dep_.CoverTab[81843]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:115
	for k, depth := range f.pointers {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:115
		_go_fuzz_dep_.CoverTab[81852]++
												if depth >= f.depth {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:116
			_go_fuzz_dep_.CoverTab[81853]++
													delete(f.pointers, k)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:117
			// _ = "end of CoverTab[81853]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:118
			_go_fuzz_dep_.CoverTab[81854]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:118
			// _ = "end of CoverTab[81854]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:118
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:118
		// _ = "end of CoverTab[81852]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:119
	// _ = "end of CoverTab[81843]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:119
	_go_fuzz_dep_.CoverTab[81844]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:122
	pointerChain := make([]uintptr, 0)

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:127
	nilFound := false
	cycleFound := false
	indirects := 0
	ve := v
	for ve.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:131
		_go_fuzz_dep_.CoverTab[81855]++
												if ve.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:132
			_go_fuzz_dep_.CoverTab[81858]++
													nilFound = true
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:134
			// _ = "end of CoverTab[81858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:135
			_go_fuzz_dep_.CoverTab[81859]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:135
			// _ = "end of CoverTab[81859]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:135
		// _ = "end of CoverTab[81855]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:135
		_go_fuzz_dep_.CoverTab[81856]++
												indirects++
												addr := ve.Pointer()
												pointerChain = append(pointerChain, addr)
												if pd, ok := f.pointers[addr]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:139
			_go_fuzz_dep_.CoverTab[81860]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:139
			return pd < f.depth
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:139
			// _ = "end of CoverTab[81860]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:139
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:139
			_go_fuzz_dep_.CoverTab[81861]++
													cycleFound = true
													indirects--
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:142
			// _ = "end of CoverTab[81861]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:143
			_go_fuzz_dep_.CoverTab[81862]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:143
			// _ = "end of CoverTab[81862]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:143
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:143
		// _ = "end of CoverTab[81856]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:143
		_go_fuzz_dep_.CoverTab[81857]++
												f.pointers[addr] = f.depth

												ve = ve.Elem()
												if ve.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:147
			_go_fuzz_dep_.CoverTab[81863]++
													if ve.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:148
				_go_fuzz_dep_.CoverTab[81865]++
														nilFound = true
														break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:150
				// _ = "end of CoverTab[81865]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:151
				_go_fuzz_dep_.CoverTab[81866]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:151
				// _ = "end of CoverTab[81866]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:151
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:151
			// _ = "end of CoverTab[81863]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:151
			_go_fuzz_dep_.CoverTab[81864]++
													ve = ve.Elem()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:152
			// _ = "end of CoverTab[81864]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:153
			_go_fuzz_dep_.CoverTab[81867]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:153
			// _ = "end of CoverTab[81867]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:153
		// _ = "end of CoverTab[81857]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:154
	// _ = "end of CoverTab[81844]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:154
	_go_fuzz_dep_.CoverTab[81845]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
	if showTypes && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
		_go_fuzz_dep_.CoverTab[81868]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
		return !f.ignoreNextType
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
		// _ = "end of CoverTab[81868]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:157
		_go_fuzz_dep_.CoverTab[81869]++
												f.fs.Write(openParenBytes)
												f.fs.Write(bytes.Repeat(asteriskBytes, indirects))
												f.fs.Write([]byte(ve.Type().String()))
												f.fs.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:161
		// _ = "end of CoverTab[81869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:162
		_go_fuzz_dep_.CoverTab[81870]++
												if nilFound || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:163
			_go_fuzz_dep_.CoverTab[81872]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:163
			return cycleFound
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:163
			// _ = "end of CoverTab[81872]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:163
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:163
			_go_fuzz_dep_.CoverTab[81873]++
													indirects += strings.Count(ve.Type().String(), "*")
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:164
			// _ = "end of CoverTab[81873]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:165
			_go_fuzz_dep_.CoverTab[81874]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:165
			// _ = "end of CoverTab[81874]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:165
		// _ = "end of CoverTab[81870]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:165
		_go_fuzz_dep_.CoverTab[81871]++
												f.fs.Write(openAngleBytes)
												f.fs.Write([]byte(strings.Repeat("*", indirects)))
												f.fs.Write(closeAngleBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:168
		// _ = "end of CoverTab[81871]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:169
	// _ = "end of CoverTab[81845]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:169
	_go_fuzz_dep_.CoverTab[81846]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
	if f.fs.Flag('+') && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
		_go_fuzz_dep_.CoverTab[81875]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
		return (len(pointerChain) > 0)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
		// _ = "end of CoverTab[81875]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:172
		_go_fuzz_dep_.CoverTab[81876]++
												f.fs.Write(openParenBytes)
												for i, addr := range pointerChain {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:174
			_go_fuzz_dep_.CoverTab[81878]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:175
				_go_fuzz_dep_.CoverTab[81880]++
														f.fs.Write(pointerChainBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:176
				// _ = "end of CoverTab[81880]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:177
				_go_fuzz_dep_.CoverTab[81881]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:177
				// _ = "end of CoverTab[81881]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:177
			// _ = "end of CoverTab[81878]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:177
			_go_fuzz_dep_.CoverTab[81879]++
													printHexPtr(f.fs, addr)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:178
			// _ = "end of CoverTab[81879]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:179
		// _ = "end of CoverTab[81876]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:179
		_go_fuzz_dep_.CoverTab[81877]++
												f.fs.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:180
		// _ = "end of CoverTab[81877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:181
		_go_fuzz_dep_.CoverTab[81882]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:181
		// _ = "end of CoverTab[81882]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:181
	// _ = "end of CoverTab[81846]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:181
	_go_fuzz_dep_.CoverTab[81847]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:184
	switch {
	case nilFound:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:185
		_go_fuzz_dep_.CoverTab[81883]++
												f.fs.Write(nilAngleBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:186
		// _ = "end of CoverTab[81883]"

	case cycleFound:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:188
		_go_fuzz_dep_.CoverTab[81884]++
												f.fs.Write(circularShortBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:189
		// _ = "end of CoverTab[81884]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:191
		_go_fuzz_dep_.CoverTab[81885]++
												f.ignoreNextType = true
												f.format(ve)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:193
		// _ = "end of CoverTab[81885]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:194
	// _ = "end of CoverTab[81847]"
}

// format is the main workhorse for providing the Formatter interface.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:197
// uses the passed reflect value to figure out what kind of object we are
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:197
// dealing with and formats it appropriately.  It is a recursive function,
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:197
// however circular data structures are detected and handled properly.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:201
func (f *formatState) format(v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:201
	_go_fuzz_dep_.CoverTab[81886]++

											kind := v.Kind()
											if kind == reflect.Invalid {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:204
		_go_fuzz_dep_.CoverTab[81891]++
												f.fs.Write(invalidAngleBytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:206
		// _ = "end of CoverTab[81891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:207
		_go_fuzz_dep_.CoverTab[81892]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:207
		// _ = "end of CoverTab[81892]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:207
	// _ = "end of CoverTab[81886]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:207
	_go_fuzz_dep_.CoverTab[81887]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:210
	if kind == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:210
		_go_fuzz_dep_.CoverTab[81893]++
												f.formatPtr(v)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:212
		// _ = "end of CoverTab[81893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:213
		_go_fuzz_dep_.CoverTab[81894]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:213
		// _ = "end of CoverTab[81894]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:213
	// _ = "end of CoverTab[81887]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:213
	_go_fuzz_dep_.CoverTab[81888]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
	if !f.ignoreNextType && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
		_go_fuzz_dep_.CoverTab[81895]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
		return f.fs.Flag('#')
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
		// _ = "end of CoverTab[81895]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
	}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:216
		_go_fuzz_dep_.CoverTab[81896]++
												f.fs.Write(openParenBytes)
												f.fs.Write([]byte(v.Type().String()))
												f.fs.Write(closeParenBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:219
		// _ = "end of CoverTab[81896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:220
		_go_fuzz_dep_.CoverTab[81897]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:220
		// _ = "end of CoverTab[81897]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:220
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:220
	// _ = "end of CoverTab[81888]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:220
	_go_fuzz_dep_.CoverTab[81889]++
											f.ignoreNextType = false

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:225
	if !f.cs.DisableMethods {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:225
		_go_fuzz_dep_.CoverTab[81898]++
												if (kind != reflect.Invalid) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:226
			_go_fuzz_dep_.CoverTab[81899]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:226
			return (kind != reflect.Interface)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:226
			// _ = "end of CoverTab[81899]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:226
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:226
			_go_fuzz_dep_.CoverTab[81900]++
													if handled := handleMethods(f.cs, f.fs, v); handled {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:227
				_go_fuzz_dep_.CoverTab[81901]++
														return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:228
				// _ = "end of CoverTab[81901]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:229
				_go_fuzz_dep_.CoverTab[81902]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:229
				// _ = "end of CoverTab[81902]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:229
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:229
			// _ = "end of CoverTab[81900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:230
			_go_fuzz_dep_.CoverTab[81903]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:230
			// _ = "end of CoverTab[81903]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:230
		// _ = "end of CoverTab[81898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:231
		_go_fuzz_dep_.CoverTab[81904]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:231
		// _ = "end of CoverTab[81904]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:231
	// _ = "end of CoverTab[81889]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:231
	_go_fuzz_dep_.CoverTab[81890]++

											switch kind {
	case reflect.Invalid:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:234
		_go_fuzz_dep_.CoverTab[81905]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:234
		// _ = "end of CoverTab[81905]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:238
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:238
		_go_fuzz_dep_.CoverTab[81906]++
												printBool(f.fs, v.Bool())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:239
		// _ = "end of CoverTab[81906]"

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:241
		_go_fuzz_dep_.CoverTab[81907]++
												printInt(f.fs, v.Int(), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:242
		// _ = "end of CoverTab[81907]"

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:244
		_go_fuzz_dep_.CoverTab[81908]++
												printUint(f.fs, v.Uint(), 10)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:245
		// _ = "end of CoverTab[81908]"

	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:247
		_go_fuzz_dep_.CoverTab[81909]++
												printFloat(f.fs, v.Float(), 32)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:248
		// _ = "end of CoverTab[81909]"

	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:250
		_go_fuzz_dep_.CoverTab[81910]++
												printFloat(f.fs, v.Float(), 64)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:251
		// _ = "end of CoverTab[81910]"

	case reflect.Complex64:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:253
		_go_fuzz_dep_.CoverTab[81911]++
												printComplex(f.fs, v.Complex(), 32)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:254
		// _ = "end of CoverTab[81911]"

	case reflect.Complex128:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:256
		_go_fuzz_dep_.CoverTab[81912]++
												printComplex(f.fs, v.Complex(), 64)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:257
		// _ = "end of CoverTab[81912]"

	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:259
		_go_fuzz_dep_.CoverTab[81913]++
												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:260
			_go_fuzz_dep_.CoverTab[81928]++
													f.fs.Write(nilAngleBytes)
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:262
			// _ = "end of CoverTab[81928]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:263
			_go_fuzz_dep_.CoverTab[81929]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:263
			// _ = "end of CoverTab[81929]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:263
		// _ = "end of CoverTab[81913]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:263
		_go_fuzz_dep_.CoverTab[81914]++
												fallthrough
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:264
		// _ = "end of CoverTab[81914]"

	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:266
		_go_fuzz_dep_.CoverTab[81915]++
												f.fs.Write(openBracketBytes)
												f.depth++
												if (f.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:269
			_go_fuzz_dep_.CoverTab[81930]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:269
			return (f.depth > f.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:269
			// _ = "end of CoverTab[81930]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:269
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:269
			_go_fuzz_dep_.CoverTab[81931]++
													f.fs.Write(maxShortBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:270
			// _ = "end of CoverTab[81931]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:271
			_go_fuzz_dep_.CoverTab[81932]++
													numEntries := v.Len()
													for i := 0; i < numEntries; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:273
				_go_fuzz_dep_.CoverTab[81933]++
														if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:274
					_go_fuzz_dep_.CoverTab[81935]++
															f.fs.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:275
					// _ = "end of CoverTab[81935]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:276
					_go_fuzz_dep_.CoverTab[81936]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:276
					// _ = "end of CoverTab[81936]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:276
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:276
				// _ = "end of CoverTab[81933]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:276
				_go_fuzz_dep_.CoverTab[81934]++
														f.ignoreNextType = true
														f.format(f.unpackValue(v.Index(i)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:278
				// _ = "end of CoverTab[81934]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:279
			// _ = "end of CoverTab[81932]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:280
		// _ = "end of CoverTab[81915]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:280
		_go_fuzz_dep_.CoverTab[81916]++
												f.depth--
												f.fs.Write(closeBracketBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:282
		// _ = "end of CoverTab[81916]"

	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:284
		_go_fuzz_dep_.CoverTab[81917]++
												f.fs.Write([]byte(v.String()))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:285
		// _ = "end of CoverTab[81917]"

	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:287
		_go_fuzz_dep_.CoverTab[81918]++

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:290
		if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:290
			_go_fuzz_dep_.CoverTab[81937]++
													f.fs.Write(nilAngleBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:291
			// _ = "end of CoverTab[81937]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:292
			_go_fuzz_dep_.CoverTab[81938]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:292
			// _ = "end of CoverTab[81938]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:292
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:292
		// _ = "end of CoverTab[81918]"

	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:294
		_go_fuzz_dep_.CoverTab[81919]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:294
		// _ = "end of CoverTab[81919]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:298
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:298
		_go_fuzz_dep_.CoverTab[81920]++

												if v.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:300
			_go_fuzz_dep_.CoverTab[81939]++
													f.fs.Write(nilAngleBytes)
													break
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:302
			// _ = "end of CoverTab[81939]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:303
			_go_fuzz_dep_.CoverTab[81940]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:303
			// _ = "end of CoverTab[81940]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:303
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:303
		// _ = "end of CoverTab[81920]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:303
		_go_fuzz_dep_.CoverTab[81921]++

												f.fs.Write(openMapBytes)
												f.depth++
												if (f.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:307
			_go_fuzz_dep_.CoverTab[81941]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:307
			return (f.depth > f.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:307
			// _ = "end of CoverTab[81941]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:307
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:307
			_go_fuzz_dep_.CoverTab[81942]++
													f.fs.Write(maxShortBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:308
			// _ = "end of CoverTab[81942]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:309
			_go_fuzz_dep_.CoverTab[81943]++
													keys := v.MapKeys()
													if f.cs.SortKeys {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:311
				_go_fuzz_dep_.CoverTab[81945]++
														sortValues(keys, f.cs)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:312
				// _ = "end of CoverTab[81945]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:313
				_go_fuzz_dep_.CoverTab[81946]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:313
				// _ = "end of CoverTab[81946]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:313
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:313
			// _ = "end of CoverTab[81943]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:313
			_go_fuzz_dep_.CoverTab[81944]++
													for i, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:314
				_go_fuzz_dep_.CoverTab[81947]++
														if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:315
					_go_fuzz_dep_.CoverTab[81949]++
															f.fs.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:316
					// _ = "end of CoverTab[81949]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:317
					_go_fuzz_dep_.CoverTab[81950]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:317
					// _ = "end of CoverTab[81950]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:317
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:317
				// _ = "end of CoverTab[81947]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:317
				_go_fuzz_dep_.CoverTab[81948]++
														f.ignoreNextType = true
														f.format(f.unpackValue(key))
														f.fs.Write(colonBytes)
														f.ignoreNextType = true
														f.format(f.unpackValue(v.MapIndex(key)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:322
				// _ = "end of CoverTab[81948]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:323
			// _ = "end of CoverTab[81944]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:324
		// _ = "end of CoverTab[81921]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:324
		_go_fuzz_dep_.CoverTab[81922]++
												f.depth--
												f.fs.Write(closeMapBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:326
		// _ = "end of CoverTab[81922]"

	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:328
		_go_fuzz_dep_.CoverTab[81923]++
												numFields := v.NumField()
												f.fs.Write(openBraceBytes)
												f.depth++
												if (f.cs.MaxDepth != 0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:332
			_go_fuzz_dep_.CoverTab[81951]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:332
			return (f.depth > f.cs.MaxDepth)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:332
			// _ = "end of CoverTab[81951]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:332
		}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:332
			_go_fuzz_dep_.CoverTab[81952]++
													f.fs.Write(maxShortBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:333
			// _ = "end of CoverTab[81952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:334
			_go_fuzz_dep_.CoverTab[81953]++
													vt := v.Type()
													for i := 0; i < numFields; i++ {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:336
				_go_fuzz_dep_.CoverTab[81954]++
														if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:337
					_go_fuzz_dep_.CoverTab[81957]++
															f.fs.Write(spaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:338
					// _ = "end of CoverTab[81957]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:339
					_go_fuzz_dep_.CoverTab[81958]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:339
					// _ = "end of CoverTab[81958]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:339
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:339
				// _ = "end of CoverTab[81954]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:339
				_go_fuzz_dep_.CoverTab[81955]++
														vtf := vt.Field(i)
														if f.fs.Flag('+') || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:341
					_go_fuzz_dep_.CoverTab[81959]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:341
					return f.fs.Flag('#')
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:341
					// _ = "end of CoverTab[81959]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:341
				}() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:341
					_go_fuzz_dep_.CoverTab[81960]++
															f.fs.Write([]byte(vtf.Name))
															f.fs.Write(colonBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:343
					// _ = "end of CoverTab[81960]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:344
					_go_fuzz_dep_.CoverTab[81961]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:344
					// _ = "end of CoverTab[81961]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:344
				}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:344
				// _ = "end of CoverTab[81955]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:344
				_go_fuzz_dep_.CoverTab[81956]++
														f.format(f.unpackValue(v.Field(i)))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:345
				// _ = "end of CoverTab[81956]"
			}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:346
			// _ = "end of CoverTab[81953]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:347
		// _ = "end of CoverTab[81923]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:347
		_go_fuzz_dep_.CoverTab[81924]++
												f.depth--
												f.fs.Write(closeBraceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:349
		// _ = "end of CoverTab[81924]"

	case reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:351
		_go_fuzz_dep_.CoverTab[81925]++
												printHexPtr(f.fs, uintptr(v.Uint()))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:352
		// _ = "end of CoverTab[81925]"

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:354
		_go_fuzz_dep_.CoverTab[81926]++
												printHexPtr(f.fs, v.Pointer())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:355
		// _ = "end of CoverTab[81926]"

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:359
	default:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:359
		_go_fuzz_dep_.CoverTab[81927]++
												format := f.buildDefaultFormat()
												if v.CanInterface() {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:361
			_go_fuzz_dep_.CoverTab[81962]++
													fmt.Fprintf(f.fs, format, v.Interface())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:362
			// _ = "end of CoverTab[81962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:363
			_go_fuzz_dep_.CoverTab[81963]++
													fmt.Fprintf(f.fs, format, v.String())
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:364
			// _ = "end of CoverTab[81963]"
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:365
		// _ = "end of CoverTab[81927]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:366
	// _ = "end of CoverTab[81890]"
}

// Format satisfies the fmt.Formatter interface. See NewFormatter for usage
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:369
// details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:371
func (f *formatState) Format(fs fmt.State, verb rune) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:371
	_go_fuzz_dep_.CoverTab[81964]++
											f.fs = fs

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:375
	if verb != 'v' {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:375
		_go_fuzz_dep_.CoverTab[81967]++
												format := f.constructOrigFormat(verb)
												fmt.Fprintf(fs, format, f.value)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:378
		// _ = "end of CoverTab[81967]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:379
		_go_fuzz_dep_.CoverTab[81968]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:379
		// _ = "end of CoverTab[81968]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:379
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:379
	// _ = "end of CoverTab[81964]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:379
	_go_fuzz_dep_.CoverTab[81965]++

											if f.value == nil {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:381
		_go_fuzz_dep_.CoverTab[81969]++
												if fs.Flag('#') {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:382
			_go_fuzz_dep_.CoverTab[81971]++
													fs.Write(interfaceBytes)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:383
			// _ = "end of CoverTab[81971]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:384
			_go_fuzz_dep_.CoverTab[81972]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:384
			// _ = "end of CoverTab[81972]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:384
		}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:384
		// _ = "end of CoverTab[81969]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:384
		_go_fuzz_dep_.CoverTab[81970]++
												fs.Write(nilAngleBytes)
												return
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:386
		// _ = "end of CoverTab[81970]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:387
		_go_fuzz_dep_.CoverTab[81973]++
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:387
		// _ = "end of CoverTab[81973]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:387
	// _ = "end of CoverTab[81965]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:387
	_go_fuzz_dep_.CoverTab[81966]++

											f.format(reflect.ValueOf(f.value))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:389
	// _ = "end of CoverTab[81966]"
}

// newFormatter is a helper function to consolidate the logic from the various
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:392
// public methods which take varying config states.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:394
func newFormatter(cs *ConfigState, v interface{}) fmt.Formatter {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:394
	_go_fuzz_dep_.CoverTab[81974]++
											fs := &formatState{value: v, cs: cs}
											fs.pointers = make(map[uintptr]int)
											return fs
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:397
	// _ = "end of CoverTab[81974]"
}

/*
NewFormatter returns a custom formatter that satisfies the fmt.Formatter
interface.  As a result, it integrates cleanly with standard fmt package
printing functions.  The formatter is useful for inline printing of smaller data
types similar to the standard %v format specifier.

The custom formatter only responds to the %v (most compact), %+v (adds pointer
addresses), %#v (adds types), or %#+v (adds types and pointer addresses) verb
combinations.  Any other verbs such as %x and %q will be sent to the the
standard fmt package for formatting.  In addition, the custom formatter ignores
the width and precision arguments (however they will still work on the format
specifiers not handled by the custom formatter).

Typically this function shouldn't be called directly.  It is much easier to make
use of the custom formatter by calling one of the convenience functions such as
Printf, Println, or Fprintf.
*/
func NewFormatter(v interface{}) fmt.Formatter {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:417
	_go_fuzz_dep_.CoverTab[81975]++
											return newFormatter(&Config, v)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:418
	// _ = "end of CoverTab[81975]"
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:419
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/format.go:419
var _ = _go_fuzz_dep_.CoverTab
