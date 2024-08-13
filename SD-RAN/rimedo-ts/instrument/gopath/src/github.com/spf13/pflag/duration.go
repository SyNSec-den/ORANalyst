//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:1
)

import (
	"time"
)

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:10
	_go_fuzz_dep_.CoverTab[119405]++
										*p = val
										return (*durationValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:12
	// _ = "end of CoverTab[119405]"
}

func (d *durationValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:15
	_go_fuzz_dep_.CoverTab[119406]++
										v, err := time.ParseDuration(s)
										*d = durationValue(v)
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:18
	// _ = "end of CoverTab[119406]"
}

func (d *durationValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:21
	_go_fuzz_dep_.CoverTab[119407]++
										return "duration"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:22
	// _ = "end of CoverTab[119407]"
}

func (d *durationValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:25
	_go_fuzz_dep_.CoverTab[119408]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:25
	return (*time.Duration)(d).String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:25
	// _ = "end of CoverTab[119408]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:25
}

func durationConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:27
	_go_fuzz_dep_.CoverTab[119409]++
										return time.ParseDuration(sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:28
	// _ = "end of CoverTab[119409]"
}

// GetDuration return the duration value of a flag with the given name
func (f *FlagSet) GetDuration(name string) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:32
	_go_fuzz_dep_.CoverTab[119410]++
										val, err := f.getFlagType(name, "duration", durationConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:34
		_go_fuzz_dep_.CoverTab[119412]++
											return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:35
		// _ = "end of CoverTab[119412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:36
		_go_fuzz_dep_.CoverTab[119413]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:36
		// _ = "end of CoverTab[119413]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:36
	// _ = "end of CoverTab[119410]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:36
	_go_fuzz_dep_.CoverTab[119411]++
										return val.(time.Duration), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:37
	// _ = "end of CoverTab[119411]"
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:40
// The argument p points to a time.Duration variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:42
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:42
	_go_fuzz_dep_.CoverTab[119414]++
										f.VarP(newDurationValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:43
	// _ = "end of CoverTab[119414]"
}

// DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:47
	_go_fuzz_dep_.CoverTab[119415]++
										f.VarP(newDurationValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:48
	// _ = "end of CoverTab[119415]"
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:51
// The argument p points to a time.Duration variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:53
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:53
	_go_fuzz_dep_.CoverTab[119416]++
										CommandLine.VarP(newDurationValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:54
	// _ = "end of CoverTab[119416]"
}

// DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.
func DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:58
	_go_fuzz_dep_.CoverTab[119417]++
										CommandLine.VarP(newDurationValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:59
	// _ = "end of CoverTab[119417]"
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:62
// The return value is the address of a time.Duration variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:64
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:64
	_go_fuzz_dep_.CoverTab[119418]++
										p := new(time.Duration)
										f.DurationVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:67
	// _ = "end of CoverTab[119418]"
}

// DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:71
	_go_fuzz_dep_.CoverTab[119419]++
										p := new(time.Duration)
										f.DurationVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:74
	// _ = "end of CoverTab[119419]"
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:77
// The return value is the address of a time.Duration variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:79
func Duration(name string, value time.Duration, usage string) *time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:79
	_go_fuzz_dep_.CoverTab[119420]++
										return CommandLine.DurationP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:80
	// _ = "end of CoverTab[119420]"
}

// DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.
func DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:84
	_go_fuzz_dep_.CoverTab[119421]++
										return CommandLine.DurationP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:85
	// _ = "end of CoverTab[119421]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration.go:86
var _ = _go_fuzz_dep_.CoverTab
