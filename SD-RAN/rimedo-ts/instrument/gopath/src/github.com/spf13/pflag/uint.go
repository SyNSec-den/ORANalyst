//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:1
)

import "strconv"

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:8
	_go_fuzz_dep_.CoverTab[120630]++
										*p = val
										return (*uintValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:10
	// _ = "end of CoverTab[120630]"
}

func (i *uintValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:13
	_go_fuzz_dep_.CoverTab[120631]++
										v, err := strconv.ParseUint(s, 0, 64)
										*i = uintValue(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:16
	// _ = "end of CoverTab[120631]"
}

func (i *uintValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:19
	_go_fuzz_dep_.CoverTab[120632]++
										return "uint"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:20
	// _ = "end of CoverTab[120632]"
}

func (i *uintValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:23
	_go_fuzz_dep_.CoverTab[120633]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:23
	return strconv.FormatUint(uint64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:23
	// _ = "end of CoverTab[120633]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:23
}

func uintConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:25
	_go_fuzz_dep_.CoverTab[120634]++
										v, err := strconv.ParseUint(sval, 0, 0)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:27
		_go_fuzz_dep_.CoverTab[120636]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:28
		// _ = "end of CoverTab[120636]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:29
		_go_fuzz_dep_.CoverTab[120637]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:29
		// _ = "end of CoverTab[120637]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:29
	// _ = "end of CoverTab[120634]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:29
	_go_fuzz_dep_.CoverTab[120635]++
										return uint(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:30
	// _ = "end of CoverTab[120635]"
}

// GetUint return the uint value of a flag with the given name
func (f *FlagSet) GetUint(name string) (uint, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:34
	_go_fuzz_dep_.CoverTab[120638]++
										val, err := f.getFlagType(name, "uint", uintConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:36
		_go_fuzz_dep_.CoverTab[120640]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:37
		// _ = "end of CoverTab[120640]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:38
		_go_fuzz_dep_.CoverTab[120641]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:38
		// _ = "end of CoverTab[120641]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:38
	// _ = "end of CoverTab[120638]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:38
	_go_fuzz_dep_.CoverTab[120639]++
										return val.(uint), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:39
	// _ = "end of CoverTab[120639]"
}

// UintVar defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:42
// The argument p points to a uint variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:44
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:44
	_go_fuzz_dep_.CoverTab[120642]++
										f.VarP(newUintValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:45
	// _ = "end of CoverTab[120642]"
}

// UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:49
	_go_fuzz_dep_.CoverTab[120643]++
										f.VarP(newUintValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:50
	// _ = "end of CoverTab[120643]"
}

// UintVar defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:53
// The argument p points to a uint  variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:55
func UintVar(p *uint, name string, value uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:55
	_go_fuzz_dep_.CoverTab[120644]++
										CommandLine.VarP(newUintValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:56
	// _ = "end of CoverTab[120644]"
}

// UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.
func UintVarP(p *uint, name, shorthand string, value uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:60
	_go_fuzz_dep_.CoverTab[120645]++
										CommandLine.VarP(newUintValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:61
	// _ = "end of CoverTab[120645]"
}

// Uint defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:64
// The return value is the address of a uint  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:66
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:66
	_go_fuzz_dep_.CoverTab[120646]++
										p := new(uint)
										f.UintVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:69
	// _ = "end of CoverTab[120646]"
}

// UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:73
	_go_fuzz_dep_.CoverTab[120647]++
										p := new(uint)
										f.UintVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:76
	// _ = "end of CoverTab[120647]"
}

// Uint defines a uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:79
// The return value is the address of a uint  variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:81
func Uint(name string, value uint, usage string) *uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:81
	_go_fuzz_dep_.CoverTab[120648]++
										return CommandLine.UintP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:82
	// _ = "end of CoverTab[120648]"
}

// UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.
func UintP(name, shorthand string, value uint, usage string) *uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:86
	_go_fuzz_dep_.CoverTab[120649]++
										return CommandLine.UintP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:87
	// _ = "end of CoverTab[120649]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint.go:88
var _ = _go_fuzz_dep_.CoverTab
