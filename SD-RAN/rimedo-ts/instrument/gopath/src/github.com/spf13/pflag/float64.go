//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:1
)

import "strconv"

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:8
	_go_fuzz_dep_.CoverTab[119940]++
										*p = val
										return (*float64Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:10
	// _ = "end of CoverTab[119940]"
}

func (f *float64Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:13
	_go_fuzz_dep_.CoverTab[119941]++
										v, err := strconv.ParseFloat(s, 64)
										*f = float64Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:16
	// _ = "end of CoverTab[119941]"
}

func (f *float64Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:19
	_go_fuzz_dep_.CoverTab[119942]++
										return "float64"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:20
	// _ = "end of CoverTab[119942]"
}

func (f *float64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:23
	_go_fuzz_dep_.CoverTab[119943]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:23
	return strconv.FormatFloat(float64(*f), 'g', -1, 64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:23
	// _ = "end of CoverTab[119943]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:23
}

func float64Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:25
	_go_fuzz_dep_.CoverTab[119944]++
										return strconv.ParseFloat(sval, 64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:26
	// _ = "end of CoverTab[119944]"
}

// GetFloat64 return the float64 value of a flag with the given name
func (f *FlagSet) GetFloat64(name string) (float64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:30
	_go_fuzz_dep_.CoverTab[119945]++
										val, err := f.getFlagType(name, "float64", float64Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:32
		_go_fuzz_dep_.CoverTab[119947]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:33
		// _ = "end of CoverTab[119947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:34
		_go_fuzz_dep_.CoverTab[119948]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:34
		// _ = "end of CoverTab[119948]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:34
	// _ = "end of CoverTab[119945]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:34
	_go_fuzz_dep_.CoverTab[119946]++
										return val.(float64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:35
	// _ = "end of CoverTab[119946]"
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:38
// The argument p points to a float64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:40
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:40
	_go_fuzz_dep_.CoverTab[119949]++
										f.VarP(newFloat64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:41
	// _ = "end of CoverTab[119949]"
}

// Float64VarP is like Float64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64VarP(p *float64, name, shorthand string, value float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:45
	_go_fuzz_dep_.CoverTab[119950]++
										f.VarP(newFloat64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:46
	// _ = "end of CoverTab[119950]"
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:49
// The argument p points to a float64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:51
func Float64Var(p *float64, name string, value float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:51
	_go_fuzz_dep_.CoverTab[119951]++
										CommandLine.VarP(newFloat64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:52
	// _ = "end of CoverTab[119951]"
}

// Float64VarP is like Float64Var, but accepts a shorthand letter that can be used after a single dash.
func Float64VarP(p *float64, name, shorthand string, value float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:56
	_go_fuzz_dep_.CoverTab[119952]++
										CommandLine.VarP(newFloat64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:57
	// _ = "end of CoverTab[119952]"
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:60
// The return value is the address of a float64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:62
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:62
	_go_fuzz_dep_.CoverTab[119953]++
										p := new(float64)
										f.Float64VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:65
	// _ = "end of CoverTab[119953]"
}

// Float64P is like Float64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64P(name, shorthand string, value float64, usage string) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:69
	_go_fuzz_dep_.CoverTab[119954]++
										p := new(float64)
										f.Float64VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:72
	// _ = "end of CoverTab[119954]"
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:75
// The return value is the address of a float64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:77
func Float64(name string, value float64, usage string) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:77
	_go_fuzz_dep_.CoverTab[119955]++
										return CommandLine.Float64P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:78
	// _ = "end of CoverTab[119955]"
}

// Float64P is like Float64, but accepts a shorthand letter that can be used after a single dash.
func Float64P(name, shorthand string, value float64, usage string) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:82
	_go_fuzz_dep_.CoverTab[119956]++
										return CommandLine.Float64P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:83
	// _ = "end of CoverTab[119956]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:84
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64.go:84
var _ = _go_fuzz_dep_.CoverTab
