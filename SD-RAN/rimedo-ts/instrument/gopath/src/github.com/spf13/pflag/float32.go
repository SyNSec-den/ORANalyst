//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:1
)

import "strconv"

// -- float32 Value
type float32Value float32

func newFloat32Value(val float32, p *float32) *float32Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:8
	_go_fuzz_dep_.CoverTab[119868]++
										*p = val
										return (*float32Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:10
	// _ = "end of CoverTab[119868]"
}

func (f *float32Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:13
	_go_fuzz_dep_.CoverTab[119869]++
										v, err := strconv.ParseFloat(s, 32)
										*f = float32Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:16
	// _ = "end of CoverTab[119869]"
}

func (f *float32Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:19
	_go_fuzz_dep_.CoverTab[119870]++
										return "float32"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:20
	// _ = "end of CoverTab[119870]"
}

func (f *float32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:23
	_go_fuzz_dep_.CoverTab[119871]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:23
	return strconv.FormatFloat(float64(*f), 'g', -1, 32)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:23
	// _ = "end of CoverTab[119871]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:23
}

func float32Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:25
	_go_fuzz_dep_.CoverTab[119872]++
										v, err := strconv.ParseFloat(sval, 32)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:27
		_go_fuzz_dep_.CoverTab[119874]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:28
		// _ = "end of CoverTab[119874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:29
		_go_fuzz_dep_.CoverTab[119875]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:29
		// _ = "end of CoverTab[119875]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:29
	// _ = "end of CoverTab[119872]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:29
	_go_fuzz_dep_.CoverTab[119873]++
										return float32(v), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:30
	// _ = "end of CoverTab[119873]"
}

// GetFloat32 return the float32 value of a flag with the given name
func (f *FlagSet) GetFloat32(name string) (float32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:34
	_go_fuzz_dep_.CoverTab[119876]++
										val, err := f.getFlagType(name, "float32", float32Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:36
		_go_fuzz_dep_.CoverTab[119878]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:37
		// _ = "end of CoverTab[119878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:38
		_go_fuzz_dep_.CoverTab[119879]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:38
		// _ = "end of CoverTab[119879]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:38
	// _ = "end of CoverTab[119876]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:38
	_go_fuzz_dep_.CoverTab[119877]++
										return val.(float32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:39
	// _ = "end of CoverTab[119877]"
}

// Float32Var defines a float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:42
// The argument p points to a float32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:44
func (f *FlagSet) Float32Var(p *float32, name string, value float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:44
	_go_fuzz_dep_.CoverTab[119880]++
										f.VarP(newFloat32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:45
	// _ = "end of CoverTab[119880]"
}

// Float32VarP is like Float32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32VarP(p *float32, name, shorthand string, value float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:49
	_go_fuzz_dep_.CoverTab[119881]++
										f.VarP(newFloat32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:50
	// _ = "end of CoverTab[119881]"
}

// Float32Var defines a float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:53
// The argument p points to a float32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:55
func Float32Var(p *float32, name string, value float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:55
	_go_fuzz_dep_.CoverTab[119882]++
										CommandLine.VarP(newFloat32Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:56
	// _ = "end of CoverTab[119882]"
}

// Float32VarP is like Float32Var, but accepts a shorthand letter that can be used after a single dash.
func Float32VarP(p *float32, name, shorthand string, value float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:60
	_go_fuzz_dep_.CoverTab[119883]++
										CommandLine.VarP(newFloat32Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:61
	// _ = "end of CoverTab[119883]"
}

// Float32 defines a float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:64
// The return value is the address of a float32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:66
func (f *FlagSet) Float32(name string, value float32, usage string) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:66
	_go_fuzz_dep_.CoverTab[119884]++
										p := new(float32)
										f.Float32VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:69
	// _ = "end of CoverTab[119884]"
}

// Float32P is like Float32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32P(name, shorthand string, value float32, usage string) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:73
	_go_fuzz_dep_.CoverTab[119885]++
										p := new(float32)
										f.Float32VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:76
	// _ = "end of CoverTab[119885]"
}

// Float32 defines a float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:79
// The return value is the address of a float32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:81
func Float32(name string, value float32, usage string) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:81
	_go_fuzz_dep_.CoverTab[119886]++
										return CommandLine.Float32P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:82
	// _ = "end of CoverTab[119886]"
}

// Float32P is like Float32, but accepts a shorthand letter that can be used after a single dash.
func Float32P(name, shorthand string, value float32, usage string) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:86
	_go_fuzz_dep_.CoverTab[119887]++
										return CommandLine.Float32P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:87
	// _ = "end of CoverTab[119887]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32.go:88
var _ = _go_fuzz_dep_.CoverTab
