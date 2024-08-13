//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:1
)

import "strconv"

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:8
	_go_fuzz_dep_.CoverTab[120036]++
										*p = val
										return (*intValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:10
	// _ = "end of CoverTab[120036]"
}

func (i *intValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:13
	_go_fuzz_dep_.CoverTab[120037]++
										v, err := strconv.ParseInt(s, 0, 64)
										*i = intValue(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:16
	// _ = "end of CoverTab[120037]"
}

func (i *intValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:19
	_go_fuzz_dep_.CoverTab[120038]++
										return "int"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:20
	// _ = "end of CoverTab[120038]"
}

func (i *intValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:23
	_go_fuzz_dep_.CoverTab[120039]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:23
	return strconv.Itoa(int(*i))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:23
	// _ = "end of CoverTab[120039]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:23
}

func intConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:25
	_go_fuzz_dep_.CoverTab[120040]++
										return strconv.Atoi(sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:26
	// _ = "end of CoverTab[120040]"
}

// GetInt return the int value of a flag with the given name
func (f *FlagSet) GetInt(name string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:30
	_go_fuzz_dep_.CoverTab[120041]++
										val, err := f.getFlagType(name, "int", intConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:32
		_go_fuzz_dep_.CoverTab[120043]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:33
		// _ = "end of CoverTab[120043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:34
		_go_fuzz_dep_.CoverTab[120044]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:34
		// _ = "end of CoverTab[120044]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:34
	// _ = "end of CoverTab[120041]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:34
	_go_fuzz_dep_.CoverTab[120042]++
										return val.(int), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:35
	// _ = "end of CoverTab[120042]"
}

// IntVar defines an int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:38
// The argument p points to an int variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:40
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:40
	_go_fuzz_dep_.CoverTab[120045]++
										f.VarP(newIntValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:41
	// _ = "end of CoverTab[120045]"
}

// IntVarP is like IntVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntVarP(p *int, name, shorthand string, value int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:45
	_go_fuzz_dep_.CoverTab[120046]++
										f.VarP(newIntValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:46
	// _ = "end of CoverTab[120046]"
}

// IntVar defines an int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:49
// The argument p points to an int variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:51
func IntVar(p *int, name string, value int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:51
	_go_fuzz_dep_.CoverTab[120047]++
										CommandLine.VarP(newIntValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:52
	// _ = "end of CoverTab[120047]"
}

// IntVarP is like IntVar, but accepts a shorthand letter that can be used after a single dash.
func IntVarP(p *int, name, shorthand string, value int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:56
	_go_fuzz_dep_.CoverTab[120048]++
										CommandLine.VarP(newIntValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:57
	// _ = "end of CoverTab[120048]"
}

// Int defines an int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:60
// The return value is the address of an int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:62
func (f *FlagSet) Int(name string, value int, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:62
	_go_fuzz_dep_.CoverTab[120049]++
										p := new(int)
										f.IntVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:65
	// _ = "end of CoverTab[120049]"
}

// IntP is like Int, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntP(name, shorthand string, value int, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:69
	_go_fuzz_dep_.CoverTab[120050]++
										p := new(int)
										f.IntVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:72
	// _ = "end of CoverTab[120050]"
}

// Int defines an int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:75
// The return value is the address of an int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:77
func Int(name string, value int, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:77
	_go_fuzz_dep_.CoverTab[120051]++
										return CommandLine.IntP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:78
	// _ = "end of CoverTab[120051]"
}

// IntP is like Int, but accepts a shorthand letter that can be used after a single dash.
func IntP(name, shorthand string, value int, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:82
	_go_fuzz_dep_.CoverTab[120052]++
										return CommandLine.IntP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:83
	// _ = "end of CoverTab[120052]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:84
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int.go:84
var _ = _go_fuzz_dep_.CoverTab
