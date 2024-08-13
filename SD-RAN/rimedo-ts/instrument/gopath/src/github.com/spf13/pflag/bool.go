//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:1
)

import "strconv"

// optional interface to indicate boolean flags that can be
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:5
// supplied without "=value" text
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:7
type boolFlag interface {
	Value
	IsBoolFlag() bool
}

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:15
	_go_fuzz_dep_.CoverTab[119266]++
										*p = val
										return (*boolValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:17
	// _ = "end of CoverTab[119266]"
}

func (b *boolValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:20
	_go_fuzz_dep_.CoverTab[119267]++
										v, err := strconv.ParseBool(s)
										*b = boolValue(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:23
	// _ = "end of CoverTab[119267]"
}

func (b *boolValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:26
	_go_fuzz_dep_.CoverTab[119268]++
										return "bool"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:27
	// _ = "end of CoverTab[119268]"
}

func (b *boolValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:30
	_go_fuzz_dep_.CoverTab[119269]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:30
	return strconv.FormatBool(bool(*b))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:30
	// _ = "end of CoverTab[119269]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:30
}

func (b *boolValue) IsBoolFlag() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:32
	_go_fuzz_dep_.CoverTab[119270]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:32
	return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:32
	// _ = "end of CoverTab[119270]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:32
}

func boolConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:34
	_go_fuzz_dep_.CoverTab[119271]++
										return strconv.ParseBool(sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:35
	// _ = "end of CoverTab[119271]"
}

// GetBool return the bool value of a flag with the given name
func (f *FlagSet) GetBool(name string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:39
	_go_fuzz_dep_.CoverTab[119272]++
										val, err := f.getFlagType(name, "bool", boolConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:41
		_go_fuzz_dep_.CoverTab[119274]++
											return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:42
		// _ = "end of CoverTab[119274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:43
		_go_fuzz_dep_.CoverTab[119275]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:43
		// _ = "end of CoverTab[119275]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:43
	// _ = "end of CoverTab[119272]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:43
	_go_fuzz_dep_.CoverTab[119273]++
										return val.(bool), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:44
	// _ = "end of CoverTab[119273]"
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:47
// The argument p points to a bool variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:49
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:49
	_go_fuzz_dep_.CoverTab[119276]++
										f.BoolVarP(p, name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:50
	// _ = "end of CoverTab[119276]"
}

// BoolVarP is like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:54
	_go_fuzz_dep_.CoverTab[119277]++
										flag := f.VarPF(newBoolValue(value, p), name, shorthand, usage)
										flag.NoOptDefVal = "true"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:56
	// _ = "end of CoverTab[119277]"
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:59
// The argument p points to a bool variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:61
func BoolVar(p *bool, name string, value bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:61
	_go_fuzz_dep_.CoverTab[119278]++
										BoolVarP(p, name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:62
	// _ = "end of CoverTab[119278]"
}

// BoolVarP is like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:66
	_go_fuzz_dep_.CoverTab[119279]++
										flag := CommandLine.VarPF(newBoolValue(value, p), name, shorthand, usage)
										flag.NoOptDefVal = "true"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:68
	// _ = "end of CoverTab[119279]"
}

// Bool defines a bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:71
// The return value is the address of a bool variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:73
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:73
	_go_fuzz_dep_.CoverTab[119280]++
										return f.BoolP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:74
	// _ = "end of CoverTab[119280]"
}

// BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:78
	_go_fuzz_dep_.CoverTab[119281]++
										p := new(bool)
										f.BoolVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:81
	// _ = "end of CoverTab[119281]"
}

// Bool defines a bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:84
// The return value is the address of a bool variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:86
func Bool(name string, value bool, usage string) *bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:86
	_go_fuzz_dep_.CoverTab[119282]++
										return BoolP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:87
	// _ = "end of CoverTab[119282]"
}

// BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
func BoolP(name, shorthand string, value bool, usage string) *bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:91
	_go_fuzz_dep_.CoverTab[119283]++
										b := CommandLine.BoolP(name, shorthand, value, usage)
										return b
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:93
	// _ = "end of CoverTab[119283]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool.go:94
var _ = _go_fuzz_dep_.CoverTab
