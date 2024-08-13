//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:1
)

import "strconv"

// -- int8 Value
type int8Value int8

func newInt8Value(val int8, p *int8) *int8Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:8
	_go_fuzz_dep_.CoverTab[120209]++
										*p = val
										return (*int8Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:10
	// _ = "end of CoverTab[120209]"
}

func (i *int8Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:13
	_go_fuzz_dep_.CoverTab[120210]++
										v, err := strconv.ParseInt(s, 0, 8)
										*i = int8Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:16
	// _ = "end of CoverTab[120210]"
}

func (i *int8Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:19
	_go_fuzz_dep_.CoverTab[120211]++
										return "int8"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:20
	// _ = "end of CoverTab[120211]"
}

func (i *int8Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:23
	_go_fuzz_dep_.CoverTab[120212]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:23
	return strconv.FormatInt(int64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:23
	// _ = "end of CoverTab[120212]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:23
}

func int8Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:25
	_go_fuzz_dep_.CoverTab[120213]++
										v, err := strconv.ParseInt(sval, 0, 8)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:27
		_go_fuzz_dep_.CoverTab[120215]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:28
		// _ = "end of CoverTab[120215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:29
		_go_fuzz_dep_.CoverTab[120216]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:29
		// _ = "end of CoverTab[120216]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:29
	// _ = "end of CoverTab[120213]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:29
	_go_fuzz_dep_.CoverTab[120214]++
										return int8(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:30
	// _ = "end of CoverTab[120214]"
}

// GetInt8 return the int8 value of a flag with the given name
func (f *FlagSet) GetInt8(name string) (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:34
	_go_fuzz_dep_.CoverTab[120217]++
										val, err := f.getFlagType(name, "int8", int8Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:36
		_go_fuzz_dep_.CoverTab[120219]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:37
		// _ = "end of CoverTab[120219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:38
		_go_fuzz_dep_.CoverTab[120220]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:38
		// _ = "end of CoverTab[120220]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:38
	// _ = "end of CoverTab[120217]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:38
	_go_fuzz_dep_.CoverTab[120218]++
										return val.(int8), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:39
	// _ = "end of CoverTab[120218]"
}

// Int8Var defines an int8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:42
// The argument p points to an int8 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:44
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:44
	_go_fuzz_dep_.CoverTab[120221]++
										f.VarP(newInt8Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:45
	// _ = "end of CoverTab[120221]"
}

// Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:49
	_go_fuzz_dep_.CoverTab[120222]++
										f.VarP(newInt8Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:50
	// _ = "end of CoverTab[120222]"
}

// Int8Var defines an int8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:53
// The argument p points to an int8 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:55
func Int8Var(p *int8, name string, value int8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:55
	_go_fuzz_dep_.CoverTab[120223]++
										CommandLine.VarP(newInt8Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:56
	// _ = "end of CoverTab[120223]"
}

// Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.
func Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:60
	_go_fuzz_dep_.CoverTab[120224]++
										CommandLine.VarP(newInt8Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:61
	// _ = "end of CoverTab[120224]"
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:64
// The return value is the address of an int8 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:66
func (f *FlagSet) Int8(name string, value int8, usage string) *int8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:66
	_go_fuzz_dep_.CoverTab[120225]++
										p := new(int8)
										f.Int8VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:69
	// _ = "end of CoverTab[120225]"
}

// Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:73
	_go_fuzz_dep_.CoverTab[120226]++
										p := new(int8)
										f.Int8VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:76
	// _ = "end of CoverTab[120226]"
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:79
// The return value is the address of an int8 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:81
func Int8(name string, value int8, usage string) *int8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:81
	_go_fuzz_dep_.CoverTab[120227]++
										return CommandLine.Int8P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:82
	// _ = "end of CoverTab[120227]"
}

// Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.
func Int8P(name, shorthand string, value int8, usage string) *int8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:86
	_go_fuzz_dep_.CoverTab[120228]++
										return CommandLine.Int8P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:87
	// _ = "end of CoverTab[120228]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int8.go:88
var _ = _go_fuzz_dep_.CoverTab
