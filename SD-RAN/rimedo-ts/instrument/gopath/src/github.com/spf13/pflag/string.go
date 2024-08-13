//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:1
)

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:6
	_go_fuzz_dep_.CoverTab[120410]++
										*p = val
										return (*stringValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:8
	// _ = "end of CoverTab[120410]"
}

func (s *stringValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:11
	_go_fuzz_dep_.CoverTab[120411]++
										*s = stringValue(val)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:13
	// _ = "end of CoverTab[120411]"
}
func (s *stringValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:15
	_go_fuzz_dep_.CoverTab[120412]++
										return "string"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:16
	// _ = "end of CoverTab[120412]"
}

func (s *stringValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:19
	_go_fuzz_dep_.CoverTab[120413]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:19
	return string(*s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:19
	// _ = "end of CoverTab[120413]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:19
}

func stringConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:21
	_go_fuzz_dep_.CoverTab[120414]++
										return sval, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:22
	// _ = "end of CoverTab[120414]"
}

// GetString return the string value of a flag with the given name
func (f *FlagSet) GetString(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:26
	_go_fuzz_dep_.CoverTab[120415]++
										val, err := f.getFlagType(name, "string", stringConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:28
		_go_fuzz_dep_.CoverTab[120417]++
											return "", err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:29
		// _ = "end of CoverTab[120417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:30
		_go_fuzz_dep_.CoverTab[120418]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:30
		// _ = "end of CoverTab[120418]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:30
	// _ = "end of CoverTab[120415]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:30
	_go_fuzz_dep_.CoverTab[120416]++
										return val.(string), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:31
	// _ = "end of CoverTab[120416]"
}

// StringVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:34
// The argument p points to a string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:36
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:36
	_go_fuzz_dep_.CoverTab[120419]++
										f.VarP(newStringValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:37
	// _ = "end of CoverTab[120419]"
}

// StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:41
	_go_fuzz_dep_.CoverTab[120420]++
										f.VarP(newStringValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:42
	// _ = "end of CoverTab[120420]"
}

// StringVar defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:45
// The argument p points to a string variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:47
func StringVar(p *string, name string, value string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:47
	_go_fuzz_dep_.CoverTab[120421]++
										CommandLine.VarP(newStringValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:48
	// _ = "end of CoverTab[120421]"
}

// StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.
func StringVarP(p *string, name, shorthand string, value string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:52
	_go_fuzz_dep_.CoverTab[120422]++
										CommandLine.VarP(newStringValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:53
	// _ = "end of CoverTab[120422]"
}

// String defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:56
// The return value is the address of a string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:58
func (f *FlagSet) String(name string, value string, usage string) *string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:58
	_go_fuzz_dep_.CoverTab[120423]++
										p := new(string)
										f.StringVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:61
	// _ = "end of CoverTab[120423]"
}

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:65
	_go_fuzz_dep_.CoverTab[120424]++
										p := new(string)
										f.StringVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:68
	// _ = "end of CoverTab[120424]"
}

// String defines a string flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:71
// The return value is the address of a string variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:73
func String(name string, value string, usage string) *string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:73
	_go_fuzz_dep_.CoverTab[120425]++
										return CommandLine.StringP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:74
	// _ = "end of CoverTab[120425]"
}

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func StringP(name, shorthand string, value string, usage string) *string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:78
	_go_fuzz_dep_.CoverTab[120426]++
										return CommandLine.StringP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:79
	// _ = "end of CoverTab[120426]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/string.go:80
var _ = _go_fuzz_dep_.CoverTab
