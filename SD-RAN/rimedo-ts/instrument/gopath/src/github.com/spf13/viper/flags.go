//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
package viper

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:1
)

import "github.com/spf13/pflag"

// FlagValueSet is an interface that users can implement
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:5
// to bind a set of flags to viper.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:7
type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}

// FlagValue is an interface that users can implement
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:11
// to bind different flags to viper.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:13
type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}

// pflagValueSet is a wrapper around *pflag.ValueSet
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:20
// that implements FlagValueSet.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:22
type pflagValueSet struct {
	flags *pflag.FlagSet
}

// VisitAll iterates over all *pflag.Flag inside the *pflag.FlagSet.
func (p pflagValueSet) VisitAll(fn func(flag FlagValue)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:27
	_go_fuzz_dep_.CoverTab[129548]++
										p.flags.VisitAll(func(flag *pflag.Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:28
		_go_fuzz_dep_.CoverTab[129549]++
											fn(pflagValue{flag})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:29
		// _ = "end of CoverTab[129549]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:30
	// _ = "end of CoverTab[129548]"
}

// pflagValue is a wrapper aroung *pflag.flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:33
// that implements FlagValue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:35
type pflagValue struct {
	flag *pflag.Flag
}

// HasChanged returns whether the flag has changes or not.
func (p pflagValue) HasChanged() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:40
	_go_fuzz_dep_.CoverTab[129550]++
										return p.flag.Changed
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:41
	// _ = "end of CoverTab[129550]"
}

// Name returns the name of the flag.
func (p pflagValue) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:45
	_go_fuzz_dep_.CoverTab[129551]++
										return p.flag.Name
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:46
	// _ = "end of CoverTab[129551]"
}

// ValueString returns the value of the flag as a string.
func (p pflagValue) ValueString() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:50
	_go_fuzz_dep_.CoverTab[129552]++
										return p.flag.Value.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:51
	// _ = "end of CoverTab[129552]"
}

// ValueType returns the type of the flag as a string.
func (p pflagValue) ValueType() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:55
	_go_fuzz_dep_.CoverTab[129553]++
										return p.flag.Value.Type()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:56
	// _ = "end of CoverTab[129553]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/flags.go:57
var _ = _go_fuzz_dep_.CoverTab
