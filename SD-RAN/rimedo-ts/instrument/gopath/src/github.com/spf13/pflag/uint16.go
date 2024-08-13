//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:1
)

import "strconv"

// -- uint16 value
type uint16Value uint16

func newUint16Value(val uint16, p *uint16) *uint16Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:8
	_go_fuzz_dep_.CoverTab[120650]++
										*p = val
										return (*uint16Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:10
	// _ = "end of CoverTab[120650]"
}

func (i *uint16Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:13
	_go_fuzz_dep_.CoverTab[120651]++
										v, err := strconv.ParseUint(s, 0, 16)
										*i = uint16Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:16
	// _ = "end of CoverTab[120651]"
}

func (i *uint16Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:19
	_go_fuzz_dep_.CoverTab[120652]++
										return "uint16"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:20
	// _ = "end of CoverTab[120652]"
}

func (i *uint16Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:23
	_go_fuzz_dep_.CoverTab[120653]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:23
	return strconv.FormatUint(uint64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:23
	// _ = "end of CoverTab[120653]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:23
}

func uint16Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:25
	_go_fuzz_dep_.CoverTab[120654]++
										v, err := strconv.ParseUint(sval, 0, 16)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:27
		_go_fuzz_dep_.CoverTab[120656]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:28
		// _ = "end of CoverTab[120656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:29
		_go_fuzz_dep_.CoverTab[120657]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:29
		// _ = "end of CoverTab[120657]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:29
	// _ = "end of CoverTab[120654]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:29
	_go_fuzz_dep_.CoverTab[120655]++
										return uint16(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:30
	// _ = "end of CoverTab[120655]"
}

// GetUint16 return the uint16 value of a flag with the given name
func (f *FlagSet) GetUint16(name string) (uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:34
	_go_fuzz_dep_.CoverTab[120658]++
										val, err := f.getFlagType(name, "uint16", uint16Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:36
		_go_fuzz_dep_.CoverTab[120660]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:37
		// _ = "end of CoverTab[120660]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:38
		_go_fuzz_dep_.CoverTab[120661]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:38
		// _ = "end of CoverTab[120661]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:38
	// _ = "end of CoverTab[120658]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:38
	_go_fuzz_dep_.CoverTab[120659]++
										return val.(uint16), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:39
	// _ = "end of CoverTab[120659]"
}

// Uint16Var defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:42
// The argument p points to a uint variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:44
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:44
	_go_fuzz_dep_.CoverTab[120662]++
										f.VarP(newUint16Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:45
	// _ = "end of CoverTab[120662]"
}

// Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:49
	_go_fuzz_dep_.CoverTab[120663]++
										f.VarP(newUint16Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:50
	// _ = "end of CoverTab[120663]"
}

// Uint16Var defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:53
// The argument p points to a uint  variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:55
func Uint16Var(p *uint16, name string, value uint16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:55
	_go_fuzz_dep_.CoverTab[120664]++
										CommandLine.VarP(newUint16Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:56
	// _ = "end of CoverTab[120664]"
}

// Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.
func Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:60
	_go_fuzz_dep_.CoverTab[120665]++
										CommandLine.VarP(newUint16Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:61
	// _ = "end of CoverTab[120665]"
}

// Uint16 defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:64
// The return value is the address of a uint  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:66
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:66
	_go_fuzz_dep_.CoverTab[120666]++
										p := new(uint16)
										f.Uint16VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:69
	// _ = "end of CoverTab[120666]"
}

// Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:73
	_go_fuzz_dep_.CoverTab[120667]++
										p := new(uint16)
										f.Uint16VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:76
	// _ = "end of CoverTab[120667]"
}

// Uint16 defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:79
// The return value is the address of a uint  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:81
func Uint16(name string, value uint16, usage string) *uint16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:81
	_go_fuzz_dep_.CoverTab[120668]++
										return CommandLine.Uint16P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:82
	// _ = "end of CoverTab[120668]"
}

// Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.
func Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:86
	_go_fuzz_dep_.CoverTab[120669]++
										return CommandLine.Uint16P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:87
	// _ = "end of CoverTab[120669]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint16.go:88
var _ = _go_fuzz_dep_.CoverTab
