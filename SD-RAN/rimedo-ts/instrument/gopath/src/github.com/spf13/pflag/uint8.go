//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:1
)

import "strconv"

// -- uint8 Value
type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:8
	_go_fuzz_dep_.CoverTab[120710]++
										*p = val
										return (*uint8Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:10
	// _ = "end of CoverTab[120710]"
}

func (i *uint8Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:13
	_go_fuzz_dep_.CoverTab[120711]++
										v, err := strconv.ParseUint(s, 0, 8)
										*i = uint8Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:16
	// _ = "end of CoverTab[120711]"
}

func (i *uint8Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:19
	_go_fuzz_dep_.CoverTab[120712]++
										return "uint8"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:20
	// _ = "end of CoverTab[120712]"
}

func (i *uint8Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:23
	_go_fuzz_dep_.CoverTab[120713]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:23
	return strconv.FormatUint(uint64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:23
	// _ = "end of CoverTab[120713]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:23
}

func uint8Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:25
	_go_fuzz_dep_.CoverTab[120714]++
										v, err := strconv.ParseUint(sval, 0, 8)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:27
		_go_fuzz_dep_.CoverTab[120716]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:28
		// _ = "end of CoverTab[120716]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:29
		_go_fuzz_dep_.CoverTab[120717]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:29
		// _ = "end of CoverTab[120717]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:29
	// _ = "end of CoverTab[120714]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:29
	_go_fuzz_dep_.CoverTab[120715]++
										return uint8(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:30
	// _ = "end of CoverTab[120715]"
}

// GetUint8 return the uint8 value of a flag with the given name
func (f *FlagSet) GetUint8(name string) (uint8, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:34
	_go_fuzz_dep_.CoverTab[120718]++
										val, err := f.getFlagType(name, "uint8", uint8Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:36
		_go_fuzz_dep_.CoverTab[120720]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:37
		// _ = "end of CoverTab[120720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:38
		_go_fuzz_dep_.CoverTab[120721]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:38
		// _ = "end of CoverTab[120721]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:38
	// _ = "end of CoverTab[120718]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:38
	_go_fuzz_dep_.CoverTab[120719]++
										return val.(uint8), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:39
	// _ = "end of CoverTab[120719]"
}

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:42
// The argument p points to a uint8 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:44
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:44
	_go_fuzz_dep_.CoverTab[120722]++
										f.VarP(newUint8Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:45
	// _ = "end of CoverTab[120722]"
}

// Uint8VarP is like Uint8Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:49
	_go_fuzz_dep_.CoverTab[120723]++
										f.VarP(newUint8Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:50
	// _ = "end of CoverTab[120723]"
}

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:53
// The argument p points to a uint8 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:55
func Uint8Var(p *uint8, name string, value uint8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:55
	_go_fuzz_dep_.CoverTab[120724]++
										CommandLine.VarP(newUint8Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:56
	// _ = "end of CoverTab[120724]"
}

// Uint8VarP is like Uint8Var, but accepts a shorthand letter that can be used after a single dash.
func Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:60
	_go_fuzz_dep_.CoverTab[120725]++
										CommandLine.VarP(newUint8Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:61
	// _ = "end of CoverTab[120725]"
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:64
// The return value is the address of a uint8 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:66
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:66
	_go_fuzz_dep_.CoverTab[120726]++
										p := new(uint8)
										f.Uint8VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:69
	// _ = "end of CoverTab[120726]"
}

// Uint8P is like Uint8, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint8P(name, shorthand string, value uint8, usage string) *uint8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:73
	_go_fuzz_dep_.CoverTab[120727]++
										p := new(uint8)
										f.Uint8VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:76
	// _ = "end of CoverTab[120727]"
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:79
// The return value is the address of a uint8 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:81
func Uint8(name string, value uint8, usage string) *uint8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:81
	_go_fuzz_dep_.CoverTab[120728]++
										return CommandLine.Uint8P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:82
	// _ = "end of CoverTab[120728]"
}

// Uint8P is like Uint8, but accepts a shorthand letter that can be used after a single dash.
func Uint8P(name, shorthand string, value uint8, usage string) *uint8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:86
	_go_fuzz_dep_.CoverTab[120729]++
										return CommandLine.Uint8P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:87
	// _ = "end of CoverTab[120729]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint8.go:88
var _ = _go_fuzz_dep_.CoverTab
