//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:1
)

import "strconv"

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:8
	_go_fuzz_dep_.CoverTab[120145]++
										*p = val
										return (*int64Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:10
	// _ = "end of CoverTab[120145]"
}

func (i *int64Value) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:13
	_go_fuzz_dep_.CoverTab[120146]++
										v, err := strconv.ParseInt(s, 0, 64)
										*i = int64Value(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:16
	// _ = "end of CoverTab[120146]"
}

func (i *int64Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:19
	_go_fuzz_dep_.CoverTab[120147]++
										return "int64"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:20
	// _ = "end of CoverTab[120147]"
}

func (i *int64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:23
	_go_fuzz_dep_.CoverTab[120148]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:23
	return strconv.FormatInt(int64(*i), 10)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:23
	// _ = "end of CoverTab[120148]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:23
}

func int64Conv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:25
	_go_fuzz_dep_.CoverTab[120149]++
										return strconv.ParseInt(sval, 0, 64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:26
	// _ = "end of CoverTab[120149]"
}

// GetInt64 return the int64 value of a flag with the given name
func (f *FlagSet) GetInt64(name string) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:30
	_go_fuzz_dep_.CoverTab[120150]++
										val, err := f.getFlagType(name, "int64", int64Conv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:32
		_go_fuzz_dep_.CoverTab[120152]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:33
		// _ = "end of CoverTab[120152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:34
		_go_fuzz_dep_.CoverTab[120153]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:34
		// _ = "end of CoverTab[120153]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:34
	// _ = "end of CoverTab[120150]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:34
	_go_fuzz_dep_.CoverTab[120151]++
										return val.(int64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:35
	// _ = "end of CoverTab[120151]"
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:38
// The argument p points to an int64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:40
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:40
	_go_fuzz_dep_.CoverTab[120154]++
										f.VarP(newInt64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:41
	// _ = "end of CoverTab[120154]"
}

// Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64VarP(p *int64, name, shorthand string, value int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:45
	_go_fuzz_dep_.CoverTab[120155]++
										f.VarP(newInt64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:46
	// _ = "end of CoverTab[120155]"
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:49
// The argument p points to an int64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:51
func Int64Var(p *int64, name string, value int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:51
	_go_fuzz_dep_.CoverTab[120156]++
										CommandLine.VarP(newInt64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:52
	// _ = "end of CoverTab[120156]"
}

// Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.
func Int64VarP(p *int64, name, shorthand string, value int64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:56
	_go_fuzz_dep_.CoverTab[120157]++
										CommandLine.VarP(newInt64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:57
	// _ = "end of CoverTab[120157]"
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:60
// The return value is the address of an int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:62
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:62
	_go_fuzz_dep_.CoverTab[120158]++
										p := new(int64)
										f.Int64VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:65
	// _ = "end of CoverTab[120158]"
}

// Int64P is like Int64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int64P(name, shorthand string, value int64, usage string) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:69
	_go_fuzz_dep_.CoverTab[120159]++
										p := new(int64)
										f.Int64VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:72
	// _ = "end of CoverTab[120159]"
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:75
// The return value is the address of an int64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:77
func Int64(name string, value int64, usage string) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:77
	_go_fuzz_dep_.CoverTab[120160]++
										return CommandLine.Int64P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:78
	// _ = "end of CoverTab[120160]"
}

// Int64P is like Int64, but accepts a shorthand letter that can be used after a single dash.
func Int64P(name, shorthand string, value int64, usage string) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:82
	_go_fuzz_dep_.CoverTab[120161]++
										return CommandLine.Int64P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:83
	// _ = "end of CoverTab[120161]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:84
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int64.go:84
var _ = _go_fuzz_dep_.CoverTab
