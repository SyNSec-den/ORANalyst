//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:1
)

import "strconv"

// -- count Value
type countValue int

func newCountValue(val int, p *int) *countValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:8
	_go_fuzz_dep_.CoverTab[119382]++
										*p = val
										return (*countValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:10
	// _ = "end of CoverTab[119382]"
}

func (i *countValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:13
	_go_fuzz_dep_.CoverTab[119383]++

										if s == "+1" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:15
		_go_fuzz_dep_.CoverTab[119385]++
											*i = countValue(*i + 1)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:17
		// _ = "end of CoverTab[119385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:18
		_go_fuzz_dep_.CoverTab[119386]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:18
		// _ = "end of CoverTab[119386]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:18
	// _ = "end of CoverTab[119383]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:18
	_go_fuzz_dep_.CoverTab[119384]++
										v, err := strconv.ParseInt(s, 0, 0)
										*i = countValue(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:21
	// _ = "end of CoverTab[119384]"
}

func (i *countValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:24
	_go_fuzz_dep_.CoverTab[119387]++
										return "count"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:25
	// _ = "end of CoverTab[119387]"
}

func (i *countValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:28
	_go_fuzz_dep_.CoverTab[119388]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:28
	return strconv.Itoa(int(*i))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:28
	// _ = "end of CoverTab[119388]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:28
}

func countConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:30
	_go_fuzz_dep_.CoverTab[119389]++
										i, err := strconv.Atoi(sval)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:32
		_go_fuzz_dep_.CoverTab[119391]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:33
		// _ = "end of CoverTab[119391]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:34
		_go_fuzz_dep_.CoverTab[119392]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:34
		// _ = "end of CoverTab[119392]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:34
	// _ = "end of CoverTab[119389]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:34
	_go_fuzz_dep_.CoverTab[119390]++
										return i, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:35
	// _ = "end of CoverTab[119390]"
}

// GetCount return the int value of a flag with the given name
func (f *FlagSet) GetCount(name string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:39
	_go_fuzz_dep_.CoverTab[119393]++
										val, err := f.getFlagType(name, "count", countConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:41
		_go_fuzz_dep_.CoverTab[119395]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:42
		// _ = "end of CoverTab[119395]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:43
		_go_fuzz_dep_.CoverTab[119396]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:43
		// _ = "end of CoverTab[119396]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:43
	// _ = "end of CoverTab[119393]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:43
	_go_fuzz_dep_.CoverTab[119394]++
										return val.(int), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:44
	// _ = "end of CoverTab[119394]"
}

// CountVar defines a count flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:47
// The argument p points to an int variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:47
// A count flag will add 1 to its value every time it is found on the command line
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:50
func (f *FlagSet) CountVar(p *int, name string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:50
	_go_fuzz_dep_.CoverTab[119397]++
										f.CountVarP(p, name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:51
	// _ = "end of CoverTab[119397]"
}

// CountVarP is like CountVar only take a shorthand for the flag name.
func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:55
	_go_fuzz_dep_.CoverTab[119398]++
										flag := f.VarPF(newCountValue(0, p), name, shorthand, usage)
										flag.NoOptDefVal = "+1"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:57
	// _ = "end of CoverTab[119398]"
}

// CountVar like CountVar only the flag is placed on the CommandLine instead of a given flag set
func CountVar(p *int, name string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:61
	_go_fuzz_dep_.CoverTab[119399]++
										CommandLine.CountVar(p, name, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:62
	// _ = "end of CoverTab[119399]"
}

// CountVarP is like CountVar only take a shorthand for the flag name.
func CountVarP(p *int, name, shorthand string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:66
	_go_fuzz_dep_.CoverTab[119400]++
										CommandLine.CountVarP(p, name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:67
	// _ = "end of CoverTab[119400]"
}

// Count defines a count flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:70
// The return value is the address of an int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:70
// A count flag will add 1 to its value every time it is found on the command line
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:73
func (f *FlagSet) Count(name string, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:73
	_go_fuzz_dep_.CoverTab[119401]++
										p := new(int)
										f.CountVarP(p, name, "", usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:76
	// _ = "end of CoverTab[119401]"
}

// CountP is like Count only takes a shorthand for the flag name.
func (f *FlagSet) CountP(name, shorthand string, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:80
	_go_fuzz_dep_.CoverTab[119402]++
										p := new(int)
										f.CountVarP(p, name, shorthand, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:83
	// _ = "end of CoverTab[119402]"
}

// Count defines a count flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:86
// The return value is the address of an int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:86
// A count flag will add 1 to its value evey time it is found on the command line
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:89
func Count(name string, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:89
	_go_fuzz_dep_.CoverTab[119403]++
										return CommandLine.CountP(name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:90
	// _ = "end of CoverTab[119403]"
}

// CountP is like Count only takes a shorthand for the flag name.
func CountP(name, shorthand string, usage string) *int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:94
	_go_fuzz_dep_.CoverTab[119404]++
										return CommandLine.CountP(name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:95
	// _ = "end of CoverTab[119404]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:96
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/count.go:96
var _ = _go_fuzz_dep_.CoverTab
