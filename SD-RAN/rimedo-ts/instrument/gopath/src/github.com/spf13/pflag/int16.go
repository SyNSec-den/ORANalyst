//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:1
)

import "strconv"

// -- int16 Value
type int16Value int16

func newInt16Value(val int16, p *int16) *int16Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:8
	_go_fuzz_dep_.CoverTab[120053]++
										*p = val
										return (*int16Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:10
	// _ = "end of CoverTab[120053]"
}

func (i *int16Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:13
	_go_fuzz_dep_.CoverTab[120054]++
										v, err := strconv.ParseInt(s, 0, 16)
										*i = int16Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:16
	// _ = "end of CoverTab[120054]"
}

func (i *int16Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:19
	_go_fuzz_dep_.CoverTab[120055]++
										return "int16"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:20
	// _ = "end of CoverTab[120055]"
}

func (i *int16Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:23
	_go_fuzz_dep_.CoverTab[120056]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:23
	return strconv.FormatInt(int64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:23
	// _ = "end of CoverTab[120056]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:23
}

func int16Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:25
	_go_fuzz_dep_.CoverTab[120057]++
										v, err := strconv.ParseInt(sval, 0, 16)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:27
		_go_fuzz_dep_.CoverTab[120059]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:28
		// _ = "end of CoverTab[120059]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:29
		_go_fuzz_dep_.CoverTab[120060]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:29
		// _ = "end of CoverTab[120060]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:29
	// _ = "end of CoverTab[120057]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:29
	_go_fuzz_dep_.CoverTab[120058]++
										return int16(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:30
	// _ = "end of CoverTab[120058]"
}

// GetInt16 returns the int16 value of a flag with the given name
func (f *FlagSet) GetInt16(name string) (int16, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:34
	_go_fuzz_dep_.CoverTab[120061]++
										val, err := f.getFlagType(name, "int16", int16Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:36
		_go_fuzz_dep_.CoverTab[120063]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:37
		// _ = "end of CoverTab[120063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:38
		_go_fuzz_dep_.CoverTab[120064]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:38
		// _ = "end of CoverTab[120064]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:38
	// _ = "end of CoverTab[120061]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:38
	_go_fuzz_dep_.CoverTab[120062]++
										return val.(int16), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:39
	// _ = "end of CoverTab[120062]"
}

// Int16Var defines an int16 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:42
// The argument p points to an int16 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:44
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:44
	_go_fuzz_dep_.CoverTab[120065]++
										f.VarP(newInt16Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:45
	// _ = "end of CoverTab[120065]"
}

// Int16VarP is like Int16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:49
	_go_fuzz_dep_.CoverTab[120066]++
										f.VarP(newInt16Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:50
	// _ = "end of CoverTab[120066]"
}

// Int16Var defines an int16 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:53
// The argument p points to an int16 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:55
func Int16Var(p *int16, name string, value int16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:55
	_go_fuzz_dep_.CoverTab[120067]++
										CommandLine.VarP(newInt16Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:56
	// _ = "end of CoverTab[120067]"
}

// Int16VarP is like Int16Var, but accepts a shorthand letter that can be used after a single dash.
func Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:60
	_go_fuzz_dep_.CoverTab[120068]++
										CommandLine.VarP(newInt16Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:61
	// _ = "end of CoverTab[120068]"
}

// Int16 defines an int16 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:64
// The return value is the address of an int16 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:66
func (f *FlagSet) Int16(name string, value int16, usage string) *int16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:66
	_go_fuzz_dep_.CoverTab[120069]++
										p := new(int16)
										f.Int16VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:69
	// _ = "end of CoverTab[120069]"
}

// Int16P is like Int16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:73
	_go_fuzz_dep_.CoverTab[120070]++
										p := new(int16)
										f.Int16VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:76
	// _ = "end of CoverTab[120070]"
}

// Int16 defines an int16 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:79
// The return value is the address of an int16 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:81
func Int16(name string, value int16, usage string) *int16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:81
	_go_fuzz_dep_.CoverTab[120071]++
										return CommandLine.Int16P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:82
	// _ = "end of CoverTab[120071]"
}

// Int16P is like Int16, but accepts a shorthand letter that can be used after a single dash.
func Int16P(name, shorthand string, value int16, usage string) *int16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:86
	_go_fuzz_dep_.CoverTab[120072]++
										return CommandLine.Int16P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:87
	// _ = "end of CoverTab[120072]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int16.go:88
var _ = _go_fuzz_dep_.CoverTab
