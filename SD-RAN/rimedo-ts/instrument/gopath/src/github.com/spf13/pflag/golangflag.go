// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:5
)

import (
	goflag "flag"
	"reflect"
	"strings"
)

// flagValueWrapper implements pflag.Value around a flag.Value.  The main
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:13
// difference here is the addition of the Type method that returns a string
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:13
// name of the type.  As this is generally unknown, we approximate that with
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:13
// reflection.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:17
type flagValueWrapper struct {
	inner		goflag.Value
	flagType	string
}

// We are just copying the boolFlag interface out of goflag as that is what
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:22
// they use to decide if a flag should get "true" when no arg is given.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:24
type goBoolFlag interface {
	goflag.Value
	IsBoolFlag() bool
}

func wrapFlagValue(v goflag.Value) Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:29
	_go_fuzz_dep_.CoverTab[120004]++

											if pv, ok := v.(Value); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:31
		_go_fuzz_dep_.CoverTab[120007]++
												return pv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:32
		// _ = "end of CoverTab[120007]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:33
		_go_fuzz_dep_.CoverTab[120008]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:33
		// _ = "end of CoverTab[120008]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:33
	// _ = "end of CoverTab[120004]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:33
	_go_fuzz_dep_.CoverTab[120005]++

											pv := &flagValueWrapper{
		inner: v,
	}

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Interface || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:40
		_go_fuzz_dep_.CoverTab[120009]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:40
		return t.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:40
		// _ = "end of CoverTab[120009]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:40
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:40
		_go_fuzz_dep_.CoverTab[120010]++
												t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:41
		// _ = "end of CoverTab[120010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:42
		_go_fuzz_dep_.CoverTab[120011]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:42
		// _ = "end of CoverTab[120011]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:42
	// _ = "end of CoverTab[120005]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:42
	_go_fuzz_dep_.CoverTab[120006]++

											pv.flagType = strings.TrimSuffix(t.Name(), "Value")
											return pv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:45
	// _ = "end of CoverTab[120006]"
}

func (v *flagValueWrapper) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:48
	_go_fuzz_dep_.CoverTab[120012]++
											return v.inner.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:49
	// _ = "end of CoverTab[120012]"
}

func (v *flagValueWrapper) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:52
	_go_fuzz_dep_.CoverTab[120013]++
											return v.inner.Set(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:53
	// _ = "end of CoverTab[120013]"
}

func (v *flagValueWrapper) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:56
	_go_fuzz_dep_.CoverTab[120014]++
											return v.flagType
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:57
	// _ = "end of CoverTab[120014]"
}

// PFlagFromGoFlag will return a *pflag.Flag given a *flag.Flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:60
// If the *flag.Flag.Name was a single character (ex: `v`) it will be accessiblei
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:60
// with both `-v` and `--v` in flags. If the golang flag was more than a single
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:60
// character (ex: `verbose`) it will only be accessible via `--verbose`
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:64
func PFlagFromGoFlag(goflag *goflag.Flag) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:64
	_go_fuzz_dep_.CoverTab[120015]++

											flag := &Flag{
												Name:	goflag.Name,
												Usage:	goflag.Usage,
												Value:	wrapFlagValue(goflag.Value),

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:72
		DefValue:	goflag.Value.String(),
	}

	if len(flag.Name) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:75
		_go_fuzz_dep_.CoverTab[120018]++
												flag.Shorthand = flag.Name
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:76
		// _ = "end of CoverTab[120018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:77
		_go_fuzz_dep_.CoverTab[120019]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:77
		// _ = "end of CoverTab[120019]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:77
	// _ = "end of CoverTab[120015]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:77
	_go_fuzz_dep_.CoverTab[120016]++
											if fv, ok := goflag.Value.(goBoolFlag); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:78
		_go_fuzz_dep_.CoverTab[120020]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:78
		return fv.IsBoolFlag()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:78
		// _ = "end of CoverTab[120020]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:78
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:78
		_go_fuzz_dep_.CoverTab[120021]++
												flag.NoOptDefVal = "true"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:79
		// _ = "end of CoverTab[120021]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:80
		_go_fuzz_dep_.CoverTab[120022]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:80
		// _ = "end of CoverTab[120022]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:80
	// _ = "end of CoverTab[120016]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:80
	_go_fuzz_dep_.CoverTab[120017]++
											return flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:81
	// _ = "end of CoverTab[120017]"
}

// AddGoFlag will add the given *flag.Flag to the pflag.FlagSet
func (f *FlagSet) AddGoFlag(goflag *goflag.Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:85
	_go_fuzz_dep_.CoverTab[120023]++
											if f.Lookup(goflag.Name) != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:86
		_go_fuzz_dep_.CoverTab[120025]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:87
		// _ = "end of CoverTab[120025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:88
		_go_fuzz_dep_.CoverTab[120026]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:88
		// _ = "end of CoverTab[120026]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:88
	// _ = "end of CoverTab[120023]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:88
	_go_fuzz_dep_.CoverTab[120024]++
											newflag := PFlagFromGoFlag(goflag)
											f.AddFlag(newflag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:90
	// _ = "end of CoverTab[120024]"
}

// AddGoFlagSet will add the given *flag.FlagSet to the pflag.FlagSet
func (f *FlagSet) AddGoFlagSet(newSet *goflag.FlagSet) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:94
	_go_fuzz_dep_.CoverTab[120027]++
											if newSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:95
		_go_fuzz_dep_.CoverTab[120031]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:96
		// _ = "end of CoverTab[120031]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:97
		_go_fuzz_dep_.CoverTab[120032]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:97
		// _ = "end of CoverTab[120032]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:97
	// _ = "end of CoverTab[120027]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:97
	_go_fuzz_dep_.CoverTab[120028]++
											newSet.VisitAll(func(goflag *goflag.Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:98
		_go_fuzz_dep_.CoverTab[120033]++
												f.AddGoFlag(goflag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:99
		// _ = "end of CoverTab[120033]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:100
	// _ = "end of CoverTab[120028]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:100
	_go_fuzz_dep_.CoverTab[120029]++
											if f.addedGoFlagSets == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:101
		_go_fuzz_dep_.CoverTab[120034]++
												f.addedGoFlagSets = make([]*goflag.FlagSet, 0)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:102
		// _ = "end of CoverTab[120034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:103
		_go_fuzz_dep_.CoverTab[120035]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:103
		// _ = "end of CoverTab[120035]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:103
	// _ = "end of CoverTab[120029]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:103
	_go_fuzz_dep_.CoverTab[120030]++
											f.addedGoFlagSets = append(f.addedGoFlagSets, newSet)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:104
	// _ = "end of CoverTab[120030]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:105
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/golangflag.go:105
var _ = _go_fuzz_dep_.CoverTab
