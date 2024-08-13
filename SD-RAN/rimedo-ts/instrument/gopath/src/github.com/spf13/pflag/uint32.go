//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:1
)

import "strconv"

// -- uint32 value
type uint32Value uint32

func newUint32Value(val uint32, p *uint32) *uint32Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:8
	_go_fuzz_dep_.CoverTab[120670]++
										*p = val
										return (*uint32Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:10
	// _ = "end of CoverTab[120670]"
}

func (i *uint32Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:13
	_go_fuzz_dep_.CoverTab[120671]++
										v, err := strconv.ParseUint(s, 0, 32)
										*i = uint32Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:16
	// _ = "end of CoverTab[120671]"
}

func (i *uint32Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:19
	_go_fuzz_dep_.CoverTab[120672]++
										return "uint32"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:20
	// _ = "end of CoverTab[120672]"
}

func (i *uint32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:23
	_go_fuzz_dep_.CoverTab[120673]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:23
	return strconv.FormatUint(uint64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:23
	// _ = "end of CoverTab[120673]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:23
}

func uint32Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:25
	_go_fuzz_dep_.CoverTab[120674]++
										v, err := strconv.ParseUint(sval, 0, 32)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:27
		_go_fuzz_dep_.CoverTab[120676]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:28
		// _ = "end of CoverTab[120676]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:29
		_go_fuzz_dep_.CoverTab[120677]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:29
		// _ = "end of CoverTab[120677]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:29
	// _ = "end of CoverTab[120674]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:29
	_go_fuzz_dep_.CoverTab[120675]++
										return uint32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:30
	// _ = "end of CoverTab[120675]"
}

// GetUint32 return the uint32 value of a flag with the given name
func (f *FlagSet) GetUint32(name string) (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:34
	_go_fuzz_dep_.CoverTab[120678]++
										val, err := f.getFlagType(name, "uint32", uint32Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:36
		_go_fuzz_dep_.CoverTab[120680]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:37
		// _ = "end of CoverTab[120680]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:38
		_go_fuzz_dep_.CoverTab[120681]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:38
		// _ = "end of CoverTab[120681]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:38
	// _ = "end of CoverTab[120678]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:38
	_go_fuzz_dep_.CoverTab[120679]++
										return val.(uint32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:39
	// _ = "end of CoverTab[120679]"
}

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:42
// The argument p points to a uint32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:44
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:44
	_go_fuzz_dep_.CoverTab[120682]++
										f.VarP(newUint32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:45
	// _ = "end of CoverTab[120682]"
}

// Uint32VarP is like Uint32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:49
	_go_fuzz_dep_.CoverTab[120683]++
										f.VarP(newUint32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:50
	// _ = "end of CoverTab[120683]"
}

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:53
// The argument p points to a uint32  variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:55
func Uint32Var(p *uint32, name string, value uint32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:55
	_go_fuzz_dep_.CoverTab[120684]++
										CommandLine.VarP(newUint32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:56
	// _ = "end of CoverTab[120684]"
}

// Uint32VarP is like Uint32Var, but accepts a shorthand letter that can be used after a single dash.
func Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:60
	_go_fuzz_dep_.CoverTab[120685]++
										CommandLine.VarP(newUint32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:61
	// _ = "end of CoverTab[120685]"
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:64
// The return value is the address of a uint32  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:66
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:66
	_go_fuzz_dep_.CoverTab[120686]++
										p := new(uint32)
										f.Uint32VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:69
	// _ = "end of CoverTab[120686]"
}

// Uint32P is like Uint32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint32P(name, shorthand string, value uint32, usage string) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:73
	_go_fuzz_dep_.CoverTab[120687]++
										p := new(uint32)
										f.Uint32VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:76
	// _ = "end of CoverTab[120687]"
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:79
// The return value is the address of a uint32  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:81
func Uint32(name string, value uint32, usage string) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:81
	_go_fuzz_dep_.CoverTab[120688]++
										return CommandLine.Uint32P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:82
	// _ = "end of CoverTab[120688]"
}

// Uint32P is like Uint32, but accepts a shorthand letter that can be used after a single dash.
func Uint32P(name, shorthand string, value uint32, usage string) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:86
	_go_fuzz_dep_.CoverTab[120689]++
										return CommandLine.Uint32P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:87
	// _ = "end of CoverTab[120689]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint32.go:88
var _ = _go_fuzz_dep_.CoverTab
