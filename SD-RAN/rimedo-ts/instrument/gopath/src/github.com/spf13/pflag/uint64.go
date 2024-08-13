//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:1
)

import "strconv"

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:8
	_go_fuzz_dep_.CoverTab[120690]++
										*p = val
										return (*uint64Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:10
	// _ = "end of CoverTab[120690]"
}

func (i *uint64Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:13
	_go_fuzz_dep_.CoverTab[120691]++
										v, err := strconv.ParseUint(s, 0, 64)
										*i = uint64Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:16
	// _ = "end of CoverTab[120691]"
}

func (i *uint64Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:19
	_go_fuzz_dep_.CoverTab[120692]++
										return "uint64"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:20
	// _ = "end of CoverTab[120692]"
}

func (i *uint64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:23
	_go_fuzz_dep_.CoverTab[120693]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:23
	return strconv.FormatUint(uint64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:23
	// _ = "end of CoverTab[120693]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:23
}

func uint64Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:25
	_go_fuzz_dep_.CoverTab[120694]++
										v, err := strconv.ParseUint(sval, 0, 64)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:27
		_go_fuzz_dep_.CoverTab[120696]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:28
		// _ = "end of CoverTab[120696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:29
		_go_fuzz_dep_.CoverTab[120697]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:29
		// _ = "end of CoverTab[120697]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:29
	// _ = "end of CoverTab[120694]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:29
	_go_fuzz_dep_.CoverTab[120695]++
										return uint64(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:30
	// _ = "end of CoverTab[120695]"
}

// GetUint64 return the uint64 value of a flag with the given name
func (f *FlagSet) GetUint64(name string) (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:34
	_go_fuzz_dep_.CoverTab[120698]++
										val, err := f.getFlagType(name, "uint64", uint64Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:36
		_go_fuzz_dep_.CoverTab[120700]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:37
		// _ = "end of CoverTab[120700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:38
		_go_fuzz_dep_.CoverTab[120701]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:38
		// _ = "end of CoverTab[120701]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:38
	// _ = "end of CoverTab[120698]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:38
	_go_fuzz_dep_.CoverTab[120699]++
										return val.(uint64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:39
	// _ = "end of CoverTab[120699]"
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:42
// The argument p points to a uint64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:44
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:44
	_go_fuzz_dep_.CoverTab[120702]++
										f.VarP(newUint64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:45
	// _ = "end of CoverTab[120702]"
}

// Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:49
	_go_fuzz_dep_.CoverTab[120703]++
										f.VarP(newUint64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:50
	// _ = "end of CoverTab[120703]"
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:53
// The argument p points to a uint64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:55
func Uint64Var(p *uint64, name string, value uint64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:55
	_go_fuzz_dep_.CoverTab[120704]++
										CommandLine.VarP(newUint64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:56
	// _ = "end of CoverTab[120704]"
}

// Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.
func Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:60
	_go_fuzz_dep_.CoverTab[120705]++
										CommandLine.VarP(newUint64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:61
	// _ = "end of CoverTab[120705]"
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:64
// The return value is the address of a uint64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:66
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:66
	_go_fuzz_dep_.CoverTab[120706]++
										p := new(uint64)
										f.Uint64VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:69
	// _ = "end of CoverTab[120706]"
}

// Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:73
	_go_fuzz_dep_.CoverTab[120707]++
										p := new(uint64)
										f.Uint64VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:76
	// _ = "end of CoverTab[120707]"
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:79
// The return value is the address of a uint64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:81
func Uint64(name string, value uint64, usage string) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:81
	_go_fuzz_dep_.CoverTab[120708]++
										return CommandLine.Uint64P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:82
	// _ = "end of CoverTab[120708]"
}

// Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.
func Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:86
	_go_fuzz_dep_.CoverTab[120709]++
										return CommandLine.Uint64P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:87
	// _ = "end of CoverTab[120709]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint64.go:88
var _ = _go_fuzz_dep_.CoverTab
