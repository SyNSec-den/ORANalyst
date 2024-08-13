//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:1
)

import "strconv"

// -- int32 Value
type int32Value int32

func newInt32Value(val int32, p *int32) *int32Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:8
	_go_fuzz_dep_.CoverTab[120073]++
										*p = val
										return (*int32Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:10
	// _ = "end of CoverTab[120073]"
}

func (i *int32Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:13
	_go_fuzz_dep_.CoverTab[120074]++
										v, err := strconv.ParseInt(s, 0, 32)
										*i = int32Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:16
	// _ = "end of CoverTab[120074]"
}

func (i *int32Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:19
	_go_fuzz_dep_.CoverTab[120075]++
										return "int32"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:20
	// _ = "end of CoverTab[120075]"
}

func (i *int32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:23
	_go_fuzz_dep_.CoverTab[120076]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:23
	return strconv.FormatInt(int64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:23
	// _ = "end of CoverTab[120076]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:23
}

func int32Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:25
	_go_fuzz_dep_.CoverTab[120077]++
										v, err := strconv.ParseInt(sval, 0, 32)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:27
		_go_fuzz_dep_.CoverTab[120079]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:28
		// _ = "end of CoverTab[120079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:29
		_go_fuzz_dep_.CoverTab[120080]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:29
		// _ = "end of CoverTab[120080]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:29
	// _ = "end of CoverTab[120077]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:29
	_go_fuzz_dep_.CoverTab[120078]++
										return int32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:30
	// _ = "end of CoverTab[120078]"
}

// GetInt32 return the int32 value of a flag with the given name
func (f *FlagSet) GetInt32(name string) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:34
	_go_fuzz_dep_.CoverTab[120081]++
										val, err := f.getFlagType(name, "int32", int32Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:36
		_go_fuzz_dep_.CoverTab[120083]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:37
		// _ = "end of CoverTab[120083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:38
		_go_fuzz_dep_.CoverTab[120084]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:38
		// _ = "end of CoverTab[120084]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:38
	// _ = "end of CoverTab[120081]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:38
	_go_fuzz_dep_.CoverTab[120082]++
										return val.(int32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:39
	// _ = "end of CoverTab[120082]"
}

// Int32Var defines an int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:42
// The argument p points to an int32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:44
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:44
	_go_fuzz_dep_.CoverTab[120085]++
										f.VarP(newInt32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:45
	// _ = "end of CoverTab[120085]"
}

// Int32VarP is like Int32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32VarP(p *int32, name, shorthand string, value int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:49
	_go_fuzz_dep_.CoverTab[120086]++
										f.VarP(newInt32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:50
	// _ = "end of CoverTab[120086]"
}

// Int32Var defines an int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:53
// The argument p points to an int32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:55
func Int32Var(p *int32, name string, value int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:55
	_go_fuzz_dep_.CoverTab[120087]++
										CommandLine.VarP(newInt32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:56
	// _ = "end of CoverTab[120087]"
}

// Int32VarP is like Int32Var, but accepts a shorthand letter that can be used after a single dash.
func Int32VarP(p *int32, name, shorthand string, value int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:60
	_go_fuzz_dep_.CoverTab[120088]++
										CommandLine.VarP(newInt32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:61
	// _ = "end of CoverTab[120088]"
}

// Int32 defines an int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:64
// The return value is the address of an int32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:66
func (f *FlagSet) Int32(name string, value int32, usage string) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:66
	_go_fuzz_dep_.CoverTab[120089]++
										p := new(int32)
										f.Int32VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:69
	// _ = "end of CoverTab[120089]"
}

// Int32P is like Int32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32P(name, shorthand string, value int32, usage string) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:73
	_go_fuzz_dep_.CoverTab[120090]++
										p := new(int32)
										f.Int32VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:76
	// _ = "end of CoverTab[120090]"
}

// Int32 defines an int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:79
// The return value is the address of an int32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:81
func Int32(name string, value int32, usage string) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:81
	_go_fuzz_dep_.CoverTab[120091]++
										return CommandLine.Int32P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:82
	// _ = "end of CoverTab[120091]"
}

// Int32P is like Int32, but accepts a shorthand letter that can be used after a single dash.
func Int32P(name, shorthand string, value int32, usage string) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:86
	_go_fuzz_dep_.CoverTab[120092]++
										return CommandLine.Int32P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:87
	// _ = "end of CoverTab[120092]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32.go:88
var _ = _go_fuzz_dep_.CoverTab
