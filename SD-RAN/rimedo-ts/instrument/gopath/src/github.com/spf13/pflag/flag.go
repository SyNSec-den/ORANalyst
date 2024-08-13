// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:5
/*
Package pflag is a drop-in replacement for Go's flag package, implementing
POSIX/GNU-style --flags.

pflag is compatible with the GNU extensions to the POSIX recommendations
for command-line options. See
http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

Usage:

pflag is a drop-in replacement of Go's native flag package. If you import
pflag under the name "flag" then all code should continue to function
with no changes.

	import flag "github.com/spf13/pflag"

There is one exception to this: if you directly instantiate the Flag struct
there is one more field "Shorthand" that you will need to set.
Most code never instantiates this struct directly, and instead uses
functions such as String(), BoolVar(), and Var(), and is therefore
unaffected.

Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -flagname, stored in the pointer ip, with type *int.

	var ip = flag.Int("flagname", 1234, "help message for flagname")

If you like, you can bind the flag to a variable using the Var() functions.

	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}

Or you can create custom flags that satisfy the Value interface (with
pointer receivers) and couple them to flag parsing by

	flag.Var(&flagVal, "name", "help message for flagname")

For such flags, the default value is just the initial value of the variable.

After all flags are defined, call

	flag.Parse()

to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves,
they are all pointers; if you bind to variables, they're values.

	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

After parsing, the arguments after the flag are available as the
slice flag.Args() or individually as flag.Arg(i).
The arguments are indexed from 0 through flag.NArg()-1.

The pflag package also defines some new functions that are not in flag,
that give one-letter shorthands for flags. You can use these by appending
'P' to the name of any function that defines a flag.

	var ip = flag.IntP("flagname", "f", 1234, "help message")
	var flagvar bool
	func init() {
		flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
	}
	flag.VarP(&flagval, "varname", "v", "help message")

Shorthand letters can be used with single dashes on the command line.
Boolean shorthand flags can be combined with other shorthand flags.

Command line flag syntax:

	--flag    // boolean flags only
	--flag=x

Unlike the flag package, a single dash before an option means something
different than a double dash. Single dashes signify a series of shorthand
letters for flags. All but the last shorthand letter must be boolean flags.

	// boolean flags
	-f
	-abc
	// non-boolean flags
	-n 1234
	-Ifile
	// mixed
	-abcs "hello"
	-abcn1234

Flag parsing stops after the terminator "--". Unlike the flag package,
flags can be interspersed with arguments anywhere on the command line
before this terminator.

Integer flags accept 1234, 0664, 0x1234 and may be negative.
Boolean flags (in their long form) accept 1, 0, t, f, true, false,
TRUE, FALSE, True, False.
Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by
top-level functions.  The FlagSet type allows one to define
independent sets of flags, such as to implement subcommands
in a command-line interface. The methods of FlagSet are
analogous to the top-level functions for the command-line
flag set.
*/
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:99
)

import (
	"bytes"
	"errors"
	goflag "flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// ErrHelp is the error returned if the flag -help is invoked but no such flag is defined.
var ErrHelp = errors.New("pflag: help requested")

// ErrorHandling defines how to handle flag parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError	ErrorHandling	= iota
	// ExitOnError will call os.Exit(2) if an error is found when parsing
	ExitOnError
	// PanicOnError will panic() if an error is found when parsing flags
	PanicOnError
)

// ParseErrorsWhitelist defines the parsing errors that can be ignored
type ParseErrorsWhitelist struct {
	// UnknownFlags will ignore unknown flags errors and continue parsing rest of the flags
	UnknownFlags bool
}

// NormalizedName is a flag name that has been normalized according to rules
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:133
// for the FlagSet (e.g. making '-' and '_' equivalent).
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:135
type NormalizedName string

// A FlagSet represents a set of defined flags.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage	func()

	// SortFlags is used to indicate, if user wants to have sorted flags in
	// help/usage messages.
	SortFlags	bool

	// ParseErrorsWhitelist is used to configure a whitelist of errors
	ParseErrorsWhitelist	ParseErrorsWhitelist

	name			string
	parsed			bool
	actual			map[NormalizedName]*Flag
	orderedActual		[]*Flag
	sortedActual		[]*Flag
	formal			map[NormalizedName]*Flag
	orderedFormal		[]*Flag
	sortedFormal		[]*Flag
	shorthands		map[byte]*Flag
	args			[]string	// arguments after flags
	argsLenAtDash		int		// len(args) when a '--' was located when parsing, or -1 if no --
	errorHandling		ErrorHandling
	output			io.Writer	// nil means stderr; use out() accessor
	interspersed		bool		// allow interspersed option/non-option args
	normalizeNameFunc	func(f *FlagSet, name string) NormalizedName

	addedGoFlagSets	[]*goflag.FlagSet
}

// A Flag represents the state of a flag.
type Flag struct {
	Name			string			// name as it appears on command line
	Shorthand		string			// one-letter abbreviated flag
	Usage			string			// help message
	Value			Value			// value as set
	DefValue		string			// default value (as text); for usage message
	Changed			bool			// If the user set the value (or if left to default)
	NoOptDefVal		string			// default value (as text); if the flag is on the command line without any options
	Deprecated		string			// If this flag is deprecated, this string is the new or now thing to use
	Hidden			bool			// used by cobra.Command to allow flags to be hidden from help/usage text
	ShorthandDeprecated	string			// If the shorthand of this flag is deprecated, this string is the new or now thing to use
	Annotations		map[string][]string	// used by cobra.Command bash autocomple code
}

// Value is the interface to the dynamic value stored in a flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:185
// (The default value is represented as a string.)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:187
type Value interface {
	String() string
	Set(string) error
	Type() string
}

// SliceValue is a secondary interface to all flags which hold a list
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:193
// of values.  This allows full control over the value of list flags,
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:193
// and avoids complicated marshalling and unmarshalling to csv.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:196
type SliceValue interface {
	// Append adds the specified value to the end of the flag value list.
	Append(string) error
	// Replace will fully overwrite any data currently in the flag value list.
	Replace([]string) error
	// GetSlice returns the flag value list as an array of strings.
	GetSlice() []string
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[NormalizedName]*Flag) []*Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:206
	_go_fuzz_dep_.CoverTab[119469]++
										list := make(sort.StringSlice, len(flags))
										i := 0
										for k := range flags {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:209
		_go_fuzz_dep_.CoverTab[119472]++
											list[i] = string(k)
											i++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:211
		// _ = "end of CoverTab[119472]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:212
	// _ = "end of CoverTab[119469]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:212
	_go_fuzz_dep_.CoverTab[119470]++
										list.Sort()
										result := make([]*Flag, len(list))
										for i, name := range list {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:215
		_go_fuzz_dep_.CoverTab[119473]++
											result[i] = flags[NormalizedName(name)]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:216
		// _ = "end of CoverTab[119473]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:217
	// _ = "end of CoverTab[119470]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:217
	_go_fuzz_dep_.CoverTab[119471]++
										return result
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:218
	// _ = "end of CoverTab[119471]"
}

// SetNormalizeFunc allows you to add a function which can translate flag names.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:221
// Flags added to the FlagSet will be translated and then when anything tries to
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:221
// look up the flag that will also be translated. So it would be possible to create
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:221
// a flag named "getURL" and have it translated to "geturl".  A user could then pass
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:221
// "--getUrl" which may also be translated to "geturl" and everything will work.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:226
func (f *FlagSet) SetNormalizeFunc(n func(f *FlagSet, name string) NormalizedName) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:226
	_go_fuzz_dep_.CoverTab[119474]++
										f.normalizeNameFunc = n
										f.sortedFormal = f.sortedFormal[:0]
										for fname, flag := range f.formal {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:229
		_go_fuzz_dep_.CoverTab[119475]++
											nname := f.normalizeFlagName(flag.Name)
											if fname == nname {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:231
			_go_fuzz_dep_.CoverTab[119477]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:232
			// _ = "end of CoverTab[119477]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:233
			_go_fuzz_dep_.CoverTab[119478]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:233
			// _ = "end of CoverTab[119478]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:233
		// _ = "end of CoverTab[119475]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:233
		_go_fuzz_dep_.CoverTab[119476]++
											flag.Name = string(nname)
											delete(f.formal, fname)
											f.formal[nname] = flag
											if _, set := f.actual[fname]; set {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:237
			_go_fuzz_dep_.CoverTab[119479]++
												delete(f.actual, fname)
												f.actual[nname] = flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:239
			// _ = "end of CoverTab[119479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:240
			_go_fuzz_dep_.CoverTab[119480]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:240
			// _ = "end of CoverTab[119480]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:240
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:240
		// _ = "end of CoverTab[119476]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:241
	// _ = "end of CoverTab[119474]"
}

// GetNormalizeFunc returns the previously set NormalizeFunc of a function which
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:244
// does no translation, if not set previously.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:246
func (f *FlagSet) GetNormalizeFunc() func(f *FlagSet, name string) NormalizedName {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:246
	_go_fuzz_dep_.CoverTab[119481]++
										if f.normalizeNameFunc != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:247
		_go_fuzz_dep_.CoverTab[119483]++
											return f.normalizeNameFunc
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:248
		// _ = "end of CoverTab[119483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:249
		_go_fuzz_dep_.CoverTab[119484]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:249
		// _ = "end of CoverTab[119484]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:249
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:249
	// _ = "end of CoverTab[119481]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:249
	_go_fuzz_dep_.CoverTab[119482]++
										return func(f *FlagSet, name string) NormalizedName {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:250
		_go_fuzz_dep_.CoverTab[119485]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:250
		return NormalizedName(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:250
		// _ = "end of CoverTab[119485]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:250
	// _ = "end of CoverTab[119482]"
}

func (f *FlagSet) normalizeFlagName(name string) NormalizedName {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:253
	_go_fuzz_dep_.CoverTab[119486]++
										n := f.GetNormalizeFunc()
										return n(f, name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:255
	// _ = "end of CoverTab[119486]"
}

func (f *FlagSet) out() io.Writer {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:258
	_go_fuzz_dep_.CoverTab[119487]++
										if f.output == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:259
		_go_fuzz_dep_.CoverTab[119489]++
											return os.Stderr
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:260
		// _ = "end of CoverTab[119489]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:261
		_go_fuzz_dep_.CoverTab[119490]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:261
		// _ = "end of CoverTab[119490]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:261
	// _ = "end of CoverTab[119487]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:261
	_go_fuzz_dep_.CoverTab[119488]++
										return f.output
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:262
	// _ = "end of CoverTab[119488]"
}

// SetOutput sets the destination for usage and error messages.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:265
// If output is nil, os.Stderr is used.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:267
func (f *FlagSet) SetOutput(output io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:267
	_go_fuzz_dep_.CoverTab[119491]++
										f.output = output
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:268
	// _ = "end of CoverTab[119491]"
}

// VisitAll visits the flags in lexicographical order or
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:271
// in primordial order if f.SortFlags is false, calling fn for each.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:271
// It visits all flags, even those not set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:274
func (f *FlagSet) VisitAll(fn func(*Flag)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:274
	_go_fuzz_dep_.CoverTab[119492]++
										if len(f.formal) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:275
		_go_fuzz_dep_.CoverTab[119495]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:276
		// _ = "end of CoverTab[119495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:277
		_go_fuzz_dep_.CoverTab[119496]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:277
		// _ = "end of CoverTab[119496]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:277
	// _ = "end of CoverTab[119492]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:277
	_go_fuzz_dep_.CoverTab[119493]++

										var flags []*Flag
										if f.SortFlags {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:280
		_go_fuzz_dep_.CoverTab[119497]++
											if len(f.formal) != len(f.sortedFormal) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:281
			_go_fuzz_dep_.CoverTab[119499]++
												f.sortedFormal = sortFlags(f.formal)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:282
			// _ = "end of CoverTab[119499]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:283
			_go_fuzz_dep_.CoverTab[119500]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:283
			// _ = "end of CoverTab[119500]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:283
		// _ = "end of CoverTab[119497]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:283
		_go_fuzz_dep_.CoverTab[119498]++
											flags = f.sortedFormal
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:284
		// _ = "end of CoverTab[119498]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:285
		_go_fuzz_dep_.CoverTab[119501]++
											flags = f.orderedFormal
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:286
		// _ = "end of CoverTab[119501]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:287
	// _ = "end of CoverTab[119493]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:287
	_go_fuzz_dep_.CoverTab[119494]++

										for _, flag := range flags {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:289
		_go_fuzz_dep_.CoverTab[119502]++
											fn(flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:290
		// _ = "end of CoverTab[119502]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:291
	// _ = "end of CoverTab[119494]"
}

// HasFlags returns a bool to indicate if the FlagSet has any flags defined.
func (f *FlagSet) HasFlags() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:295
	_go_fuzz_dep_.CoverTab[119503]++
										return len(f.formal) > 0
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:296
	// _ = "end of CoverTab[119503]"
}

// HasAvailableFlags returns a bool to indicate if the FlagSet has any flags
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:299
// that are not hidden.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:301
func (f *FlagSet) HasAvailableFlags() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:301
	_go_fuzz_dep_.CoverTab[119504]++
										for _, flag := range f.formal {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:302
		_go_fuzz_dep_.CoverTab[119506]++
											if !flag.Hidden {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:303
			_go_fuzz_dep_.CoverTab[119507]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:304
			// _ = "end of CoverTab[119507]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:305
			_go_fuzz_dep_.CoverTab[119508]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:305
			// _ = "end of CoverTab[119508]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:305
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:305
		// _ = "end of CoverTab[119506]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:306
	// _ = "end of CoverTab[119504]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:306
	_go_fuzz_dep_.CoverTab[119505]++
										return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:307
	// _ = "end of CoverTab[119505]"
}

// VisitAll visits the command-line flags in lexicographical order or
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:310
// in primordial order if f.SortFlags is false, calling fn for each.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:310
// It visits all flags, even those not set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:313
func VisitAll(fn func(*Flag)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:313
	_go_fuzz_dep_.CoverTab[119509]++
										CommandLine.VisitAll(fn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:314
	// _ = "end of CoverTab[119509]"
}

// Visit visits the flags in lexicographical order or
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:317
// in primordial order if f.SortFlags is false, calling fn for each.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:317
// It visits only those flags that have been set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:320
func (f *FlagSet) Visit(fn func(*Flag)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:320
	_go_fuzz_dep_.CoverTab[119510]++
										if len(f.actual) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:321
		_go_fuzz_dep_.CoverTab[119513]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:322
		// _ = "end of CoverTab[119513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:323
		_go_fuzz_dep_.CoverTab[119514]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:323
		// _ = "end of CoverTab[119514]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:323
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:323
	// _ = "end of CoverTab[119510]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:323
	_go_fuzz_dep_.CoverTab[119511]++

										var flags []*Flag
										if f.SortFlags {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:326
		_go_fuzz_dep_.CoverTab[119515]++
											if len(f.actual) != len(f.sortedActual) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:327
			_go_fuzz_dep_.CoverTab[119517]++
												f.sortedActual = sortFlags(f.actual)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:328
			// _ = "end of CoverTab[119517]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:329
			_go_fuzz_dep_.CoverTab[119518]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:329
			// _ = "end of CoverTab[119518]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:329
		// _ = "end of CoverTab[119515]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:329
		_go_fuzz_dep_.CoverTab[119516]++
											flags = f.sortedActual
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:330
		// _ = "end of CoverTab[119516]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:331
		_go_fuzz_dep_.CoverTab[119519]++
											flags = f.orderedActual
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:332
		// _ = "end of CoverTab[119519]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:333
	// _ = "end of CoverTab[119511]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:333
	_go_fuzz_dep_.CoverTab[119512]++

										for _, flag := range flags {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:335
		_go_fuzz_dep_.CoverTab[119520]++
											fn(flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:336
		// _ = "end of CoverTab[119520]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:337
	// _ = "end of CoverTab[119512]"
}

// Visit visits the command-line flags in lexicographical order or
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:340
// in primordial order if f.SortFlags is false, calling fn for each.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:340
// It visits only those flags that have been set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:343
func Visit(fn func(*Flag)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:343
	_go_fuzz_dep_.CoverTab[119521]++
										CommandLine.Visit(fn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:344
	// _ = "end of CoverTab[119521]"
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:348
	_go_fuzz_dep_.CoverTab[119522]++
										return f.lookup(f.normalizeFlagName(name))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:349
	// _ = "end of CoverTab[119522]"
}

// ShorthandLookup returns the Flag structure of the short handed flag,
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:352
// returning nil if none exists.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:352
// It panics, if len(name) > 1.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:355
func (f *FlagSet) ShorthandLookup(name string) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:355
	_go_fuzz_dep_.CoverTab[119523]++
										if name == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:356
		_go_fuzz_dep_.CoverTab[119526]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:357
		// _ = "end of CoverTab[119526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:358
		_go_fuzz_dep_.CoverTab[119527]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:358
		// _ = "end of CoverTab[119527]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:358
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:358
	// _ = "end of CoverTab[119523]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:358
	_go_fuzz_dep_.CoverTab[119524]++
										if len(name) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:359
		_go_fuzz_dep_.CoverTab[119528]++
											msg := fmt.Sprintf("can not look up shorthand which is more than one ASCII character: %q", name)
											fmt.Fprintf(f.out(), msg)
											panic(msg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:362
		// _ = "end of CoverTab[119528]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:363
		_go_fuzz_dep_.CoverTab[119529]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:363
		// _ = "end of CoverTab[119529]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:363
	// _ = "end of CoverTab[119524]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:363
	_go_fuzz_dep_.CoverTab[119525]++
										c := name[0]
										return f.shorthands[c]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:365
	// _ = "end of CoverTab[119525]"
}

// lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) lookup(name NormalizedName) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:369
	_go_fuzz_dep_.CoverTab[119530]++
										return f.formal[name]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:370
	// _ = "end of CoverTab[119530]"
}

// func to return a given type for a given flag name
func (f *FlagSet) getFlagType(name string, ftype string, convFunc func(sval string) (interface{}, error)) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:374
	_go_fuzz_dep_.CoverTab[119531]++
										flag := f.Lookup(name)
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:376
		_go_fuzz_dep_.CoverTab[119535]++
											err := fmt.Errorf("flag accessed but not defined: %s", name)
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:378
		// _ = "end of CoverTab[119535]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:379
		_go_fuzz_dep_.CoverTab[119536]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:379
		// _ = "end of CoverTab[119536]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:379
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:379
	// _ = "end of CoverTab[119531]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:379
	_go_fuzz_dep_.CoverTab[119532]++

										if flag.Value.Type() != ftype {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:381
		_go_fuzz_dep_.CoverTab[119537]++
											err := fmt.Errorf("trying to get %s value of flag of type %s", ftype, flag.Value.Type())
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:383
		// _ = "end of CoverTab[119537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:384
		_go_fuzz_dep_.CoverTab[119538]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:384
		// _ = "end of CoverTab[119538]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:384
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:384
	// _ = "end of CoverTab[119532]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:384
	_go_fuzz_dep_.CoverTab[119533]++

										sval := flag.Value.String()
										result, err := convFunc(sval)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:388
		_go_fuzz_dep_.CoverTab[119539]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:389
		// _ = "end of CoverTab[119539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:390
		_go_fuzz_dep_.CoverTab[119540]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:390
		// _ = "end of CoverTab[119540]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:390
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:390
	// _ = "end of CoverTab[119533]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:390
	_go_fuzz_dep_.CoverTab[119534]++
										return result, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:391
	// _ = "end of CoverTab[119534]"
}

// ArgsLenAtDash will return the length of f.Args at the moment when a -- was
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:394
// found during arg parsing. This allows your program to know which args were
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:394
// before the -- and which came after.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:397
func (f *FlagSet) ArgsLenAtDash() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:397
	_go_fuzz_dep_.CoverTab[119541]++
										return f.argsLenAtDash
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:398
	// _ = "end of CoverTab[119541]"
}

// MarkDeprecated indicated that a flag is deprecated in your program. It will
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:401
// continue to function but will not show up in help or usage messages. Using
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:401
// this flag will also print the given usageMessage.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:404
func (f *FlagSet) MarkDeprecated(name string, usageMessage string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:404
	_go_fuzz_dep_.CoverTab[119542]++
										flag := f.Lookup(name)
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:406
		_go_fuzz_dep_.CoverTab[119545]++
											return fmt.Errorf("flag %q does not exist", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:407
		// _ = "end of CoverTab[119545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:408
		_go_fuzz_dep_.CoverTab[119546]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:408
		// _ = "end of CoverTab[119546]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:408
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:408
	// _ = "end of CoverTab[119542]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:408
	_go_fuzz_dep_.CoverTab[119543]++
										if usageMessage == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:409
		_go_fuzz_dep_.CoverTab[119547]++
											return fmt.Errorf("deprecated message for flag %q must be set", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:410
		// _ = "end of CoverTab[119547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:411
		_go_fuzz_dep_.CoverTab[119548]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:411
		// _ = "end of CoverTab[119548]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:411
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:411
	// _ = "end of CoverTab[119543]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:411
	_go_fuzz_dep_.CoverTab[119544]++
										flag.Deprecated = usageMessage
										flag.Hidden = true
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:414
	// _ = "end of CoverTab[119544]"
}

// MarkShorthandDeprecated will mark the shorthand of a flag deprecated in your
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:417
// program. It will continue to function but will not show up in help or usage
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:417
// messages. Using this flag will also print the given usageMessage.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:420
func (f *FlagSet) MarkShorthandDeprecated(name string, usageMessage string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:420
	_go_fuzz_dep_.CoverTab[119549]++
										flag := f.Lookup(name)
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:422
		_go_fuzz_dep_.CoverTab[119552]++
											return fmt.Errorf("flag %q does not exist", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:423
		// _ = "end of CoverTab[119552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:424
		_go_fuzz_dep_.CoverTab[119553]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:424
		// _ = "end of CoverTab[119553]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:424
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:424
	// _ = "end of CoverTab[119549]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:424
	_go_fuzz_dep_.CoverTab[119550]++
										if usageMessage == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:425
		_go_fuzz_dep_.CoverTab[119554]++
											return fmt.Errorf("deprecated message for flag %q must be set", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:426
		// _ = "end of CoverTab[119554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:427
		_go_fuzz_dep_.CoverTab[119555]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:427
		// _ = "end of CoverTab[119555]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:427
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:427
	// _ = "end of CoverTab[119550]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:427
	_go_fuzz_dep_.CoverTab[119551]++
										flag.ShorthandDeprecated = usageMessage
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:429
	// _ = "end of CoverTab[119551]"
}

// MarkHidden sets a flag to 'hidden' in your program. It will continue to
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:432
// function but will not show up in help or usage messages.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:434
func (f *FlagSet) MarkHidden(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:434
	_go_fuzz_dep_.CoverTab[119556]++
										flag := f.Lookup(name)
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:436
		_go_fuzz_dep_.CoverTab[119558]++
											return fmt.Errorf("flag %q does not exist", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:437
		// _ = "end of CoverTab[119558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:438
		_go_fuzz_dep_.CoverTab[119559]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:438
		// _ = "end of CoverTab[119559]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:438
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:438
	// _ = "end of CoverTab[119556]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:438
	_go_fuzz_dep_.CoverTab[119557]++
										flag.Hidden = true
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:440
	// _ = "end of CoverTab[119557]"
}

// Lookup returns the Flag structure of the named command-line flag,
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:443
// returning nil if none exists.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:445
func Lookup(name string) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:445
	_go_fuzz_dep_.CoverTab[119560]++
										return CommandLine.Lookup(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:446
	// _ = "end of CoverTab[119560]"
}

// ShorthandLookup returns the Flag structure of the short handed flag,
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:449
// returning nil if none exists.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:451
func ShorthandLookup(name string) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:451
	_go_fuzz_dep_.CoverTab[119561]++
										return CommandLine.ShorthandLookup(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:452
	// _ = "end of CoverTab[119561]"
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:456
	_go_fuzz_dep_.CoverTab[119562]++
										normalName := f.normalizeFlagName(name)
										flag, ok := f.formal[normalName]
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:459
		_go_fuzz_dep_.CoverTab[119567]++
											return fmt.Errorf("no such flag -%v", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:460
		// _ = "end of CoverTab[119567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:461
		_go_fuzz_dep_.CoverTab[119568]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:461
		// _ = "end of CoverTab[119568]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:461
	// _ = "end of CoverTab[119562]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:461
	_go_fuzz_dep_.CoverTab[119563]++

										err := flag.Value.Set(value)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:464
		_go_fuzz_dep_.CoverTab[119569]++
											var flagName string
											if flag.Shorthand != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:466
			_go_fuzz_dep_.CoverTab[119571]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:466
			return flag.ShorthandDeprecated == ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:466
			// _ = "end of CoverTab[119571]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:466
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:466
			_go_fuzz_dep_.CoverTab[119572]++
												flagName = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:467
			// _ = "end of CoverTab[119572]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:468
			_go_fuzz_dep_.CoverTab[119573]++
												flagName = fmt.Sprintf("--%s", flag.Name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:469
			// _ = "end of CoverTab[119573]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:470
		// _ = "end of CoverTab[119569]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:470
		_go_fuzz_dep_.CoverTab[119570]++
											return fmt.Errorf("invalid argument %q for %q flag: %v", value, flagName, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:471
		// _ = "end of CoverTab[119570]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:472
		_go_fuzz_dep_.CoverTab[119574]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:472
		// _ = "end of CoverTab[119574]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:472
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:472
	// _ = "end of CoverTab[119563]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:472
	_go_fuzz_dep_.CoverTab[119564]++

										if !flag.Changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:474
		_go_fuzz_dep_.CoverTab[119575]++
											if f.actual == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:475
			_go_fuzz_dep_.CoverTab[119577]++
												f.actual = make(map[NormalizedName]*Flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:476
			// _ = "end of CoverTab[119577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:477
			_go_fuzz_dep_.CoverTab[119578]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:477
			// _ = "end of CoverTab[119578]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:477
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:477
		// _ = "end of CoverTab[119575]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:477
		_go_fuzz_dep_.CoverTab[119576]++
											f.actual[normalName] = flag
											f.orderedActual = append(f.orderedActual, flag)

											flag.Changed = true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:481
		// _ = "end of CoverTab[119576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:482
		_go_fuzz_dep_.CoverTab[119579]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:482
		// _ = "end of CoverTab[119579]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:482
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:482
	// _ = "end of CoverTab[119564]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:482
	_go_fuzz_dep_.CoverTab[119565]++

										if flag.Deprecated != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:484
		_go_fuzz_dep_.CoverTab[119580]++
											fmt.Fprintf(f.out(), "Flag --%s has been deprecated, %s\n", flag.Name, flag.Deprecated)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:485
		// _ = "end of CoverTab[119580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:486
		_go_fuzz_dep_.CoverTab[119581]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:486
		// _ = "end of CoverTab[119581]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:486
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:486
	// _ = "end of CoverTab[119565]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:486
	_go_fuzz_dep_.CoverTab[119566]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:487
	// _ = "end of CoverTab[119566]"
}

// SetAnnotation allows one to set arbitrary annotations on a flag in the FlagSet.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:490
// This is sometimes used by spf13/cobra programs which want to generate additional
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:490
// bash completion information.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:493
func (f *FlagSet) SetAnnotation(name, key string, values []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:493
	_go_fuzz_dep_.CoverTab[119582]++
										normalName := f.normalizeFlagName(name)
										flag, ok := f.formal[normalName]
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:496
		_go_fuzz_dep_.CoverTab[119585]++
											return fmt.Errorf("no such flag -%v", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:497
		// _ = "end of CoverTab[119585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:498
		_go_fuzz_dep_.CoverTab[119586]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:498
		// _ = "end of CoverTab[119586]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:498
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:498
	// _ = "end of CoverTab[119582]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:498
	_go_fuzz_dep_.CoverTab[119583]++
										if flag.Annotations == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:499
		_go_fuzz_dep_.CoverTab[119587]++
											flag.Annotations = map[string][]string{}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:500
		// _ = "end of CoverTab[119587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:501
		_go_fuzz_dep_.CoverTab[119588]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:501
		// _ = "end of CoverTab[119588]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:501
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:501
	// _ = "end of CoverTab[119583]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:501
	_go_fuzz_dep_.CoverTab[119584]++
										flag.Annotations[key] = values
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:503
	// _ = "end of CoverTab[119584]"
}

// Changed returns true if the flag was explicitly set during Parse() and false
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:506
// otherwise
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:508
func (f *FlagSet) Changed(name string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:508
	_go_fuzz_dep_.CoverTab[119589]++
										flag := f.Lookup(name)

										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:511
		_go_fuzz_dep_.CoverTab[119591]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:512
		// _ = "end of CoverTab[119591]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:513
		_go_fuzz_dep_.CoverTab[119592]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:513
		// _ = "end of CoverTab[119592]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:513
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:513
	// _ = "end of CoverTab[119589]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:513
	_go_fuzz_dep_.CoverTab[119590]++
										return flag.Changed
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:514
	// _ = "end of CoverTab[119590]"
}

// Set sets the value of the named command-line flag.
func Set(name, value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:518
	_go_fuzz_dep_.CoverTab[119593]++
										return CommandLine.Set(name, value)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:519
	// _ = "end of CoverTab[119593]"
}

// PrintDefaults prints, to standard error unless configured
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:522
// otherwise, the default values of all defined flags in the set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:524
func (f *FlagSet) PrintDefaults() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:524
	_go_fuzz_dep_.CoverTab[119594]++
										usages := f.FlagUsages()
										fmt.Fprint(f.out(), usages)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:526
	// _ = "end of CoverTab[119594]"
}

// defaultIsZeroValue returns true if the default value for this flag represents
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:529
// a zero value.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:531
func (f *Flag) defaultIsZeroValue() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:531
	_go_fuzz_dep_.CoverTab[119595]++
										switch f.Value.(type) {
	case boolFlag:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:533
		_go_fuzz_dep_.CoverTab[119596]++
											return f.DefValue == "false"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:534
		// _ = "end of CoverTab[119596]"
	case *durationValue:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:535
		_go_fuzz_dep_.CoverTab[119597]++

											return f.DefValue == "0" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:537
			_go_fuzz_dep_.CoverTab[119604]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:537
			return f.DefValue == "0s"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:537
			// _ = "end of CoverTab[119604]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:537
		}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:537
		// _ = "end of CoverTab[119597]"
	case *intValue, *int8Value, *int32Value, *int64Value, *uintValue, *uint8Value, *uint16Value, *uint32Value, *uint64Value, *countValue, *float32Value, *float64Value:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:538
		_go_fuzz_dep_.CoverTab[119598]++
											return f.DefValue == "0"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:539
		// _ = "end of CoverTab[119598]"
	case *stringValue:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:540
		_go_fuzz_dep_.CoverTab[119599]++
											return f.DefValue == ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:541
		// _ = "end of CoverTab[119599]"
	case *ipValue, *ipMaskValue, *ipNetValue:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:542
		_go_fuzz_dep_.CoverTab[119600]++
											return f.DefValue == "<nil>"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:543
		// _ = "end of CoverTab[119600]"
	case *intSliceValue, *stringSliceValue, *stringArrayValue:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:544
		_go_fuzz_dep_.CoverTab[119601]++
											return f.DefValue == "[]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:545
		// _ = "end of CoverTab[119601]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:546
		_go_fuzz_dep_.CoverTab[119602]++
											switch f.Value.String() {
		case "false":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:548
			_go_fuzz_dep_.CoverTab[119605]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:549
			// _ = "end of CoverTab[119605]"
		case "<nil>":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:550
			_go_fuzz_dep_.CoverTab[119606]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:551
			// _ = "end of CoverTab[119606]"
		case "":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:552
			_go_fuzz_dep_.CoverTab[119607]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:553
			// _ = "end of CoverTab[119607]"
		case "0":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:554
			_go_fuzz_dep_.CoverTab[119608]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:555
			// _ = "end of CoverTab[119608]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:555
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:555
			_go_fuzz_dep_.CoverTab[119609]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:555
			// _ = "end of CoverTab[119609]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:556
		// _ = "end of CoverTab[119602]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:556
		_go_fuzz_dep_.CoverTab[119603]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:557
		// _ = "end of CoverTab[119603]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:558
	// _ = "end of CoverTab[119595]"
}

// UnquoteUsage extracts a back-quoted name from the usage
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:561
// string for a flag and returns it and the un-quoted usage.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:561
// Given "a `name` to show" it returns ("name", "a name to show").
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:561
// If there are no back quotes, the name is an educated guess of the
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:561
// type of the flag's value, or the empty string if the flag is boolean.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:566
func UnquoteUsage(flag *Flag) (name string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:566
	_go_fuzz_dep_.CoverTab[119610]++

										usage = flag.Usage
										for i := 0; i < len(usage); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:569
		_go_fuzz_dep_.CoverTab[119613]++
											if usage[i] == '`' {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:570
			_go_fuzz_dep_.CoverTab[119614]++
												for j := i + 1; j < len(usage); j++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:571
				_go_fuzz_dep_.CoverTab[119616]++
													if usage[j] == '`' {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:572
					_go_fuzz_dep_.CoverTab[119617]++
														name = usage[i+1 : j]
														usage = usage[:i] + name + usage[j+1:]
														return name, usage
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:575
					// _ = "end of CoverTab[119617]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:576
					_go_fuzz_dep_.CoverTab[119618]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:576
					// _ = "end of CoverTab[119618]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:576
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:576
				// _ = "end of CoverTab[119616]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:577
			// _ = "end of CoverTab[119614]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:577
			_go_fuzz_dep_.CoverTab[119615]++
												break
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:578
			// _ = "end of CoverTab[119615]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:579
			_go_fuzz_dep_.CoverTab[119619]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:579
			// _ = "end of CoverTab[119619]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:579
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:579
		// _ = "end of CoverTab[119613]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:580
	// _ = "end of CoverTab[119610]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:580
	_go_fuzz_dep_.CoverTab[119611]++

										name = flag.Value.Type()
										switch name {
	case "bool":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:584
		_go_fuzz_dep_.CoverTab[119620]++
											name = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:585
		// _ = "end of CoverTab[119620]"
	case "float64":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:586
		_go_fuzz_dep_.CoverTab[119621]++
											name = "float"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:587
		// _ = "end of CoverTab[119621]"
	case "int64":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:588
		_go_fuzz_dep_.CoverTab[119622]++
											name = "int"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:589
		// _ = "end of CoverTab[119622]"
	case "uint64":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:590
		_go_fuzz_dep_.CoverTab[119623]++
											name = "uint"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:591
		// _ = "end of CoverTab[119623]"
	case "stringSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:592
		_go_fuzz_dep_.CoverTab[119624]++
											name = "strings"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:593
		// _ = "end of CoverTab[119624]"
	case "intSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:594
		_go_fuzz_dep_.CoverTab[119625]++
											name = "ints"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:595
		// _ = "end of CoverTab[119625]"
	case "uintSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:596
		_go_fuzz_dep_.CoverTab[119626]++
											name = "uints"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:597
		// _ = "end of CoverTab[119626]"
	case "boolSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:598
		_go_fuzz_dep_.CoverTab[119627]++
											name = "bools"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:599
		// _ = "end of CoverTab[119627]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:599
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:599
		_go_fuzz_dep_.CoverTab[119628]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:599
		// _ = "end of CoverTab[119628]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:600
	// _ = "end of CoverTab[119611]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:600
	_go_fuzz_dep_.CoverTab[119612]++

										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:602
	// _ = "end of CoverTab[119612]"
}

// Splits the string `s` on whitespace into an initial substring up to
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:605
// `i` runes in length and the remainder. Will go `slop` over `i` if
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:605
// that encompasses the entire string (which allows the caller to
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:605
// avoid short orphan words on the final line).
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:609
func wrapN(i, slop int, s string) (string, string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:609
	_go_fuzz_dep_.CoverTab[119629]++
										if i+slop > len(s) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:610
		_go_fuzz_dep_.CoverTab[119633]++
											return s, ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:611
		// _ = "end of CoverTab[119633]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:612
		_go_fuzz_dep_.CoverTab[119634]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:612
		// _ = "end of CoverTab[119634]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:612
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:612
	// _ = "end of CoverTab[119629]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:612
	_go_fuzz_dep_.CoverTab[119630]++

										w := strings.LastIndexAny(s[:i], " \t\n")
										if w <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:615
		_go_fuzz_dep_.CoverTab[119635]++
											return s, ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:616
		// _ = "end of CoverTab[119635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:617
		_go_fuzz_dep_.CoverTab[119636]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:617
		// _ = "end of CoverTab[119636]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:617
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:617
	// _ = "end of CoverTab[119630]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:617
	_go_fuzz_dep_.CoverTab[119631]++
										nlPos := strings.LastIndex(s[:i], "\n")
										if nlPos > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:619
		_go_fuzz_dep_.CoverTab[119637]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:619
		return nlPos < w
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:619
		// _ = "end of CoverTab[119637]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:619
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:619
		_go_fuzz_dep_.CoverTab[119638]++
											return s[:nlPos], s[nlPos+1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:620
		// _ = "end of CoverTab[119638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:621
		_go_fuzz_dep_.CoverTab[119639]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:621
		// _ = "end of CoverTab[119639]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:621
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:621
	// _ = "end of CoverTab[119631]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:621
	_go_fuzz_dep_.CoverTab[119632]++
										return s[:w], s[w+1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:622
	// _ = "end of CoverTab[119632]"
}

// Wraps the string `s` to a maximum width `w` with leading indent
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:625
// `i`. The first line is not indented (this is assumed to be done by
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:625
// caller). Pass `w` == 0 to do no wrapping
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:628
func wrap(i, w int, s string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:628
	_go_fuzz_dep_.CoverTab[119640]++
										if w == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:629
		_go_fuzz_dep_.CoverTab[119645]++
											return strings.Replace(s, "\n", "\n"+strings.Repeat(" ", i), -1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:630
		// _ = "end of CoverTab[119645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:631
		_go_fuzz_dep_.CoverTab[119646]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:631
		// _ = "end of CoverTab[119646]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:631
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:631
	// _ = "end of CoverTab[119640]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:631
	_go_fuzz_dep_.CoverTab[119641]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:635
	wrap := w - i

										var r, l string

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:641
	if wrap < 24 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:641
		_go_fuzz_dep_.CoverTab[119647]++
											i = 16
											wrap = w - i
											r += "\n" + strings.Repeat(" ", i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:644
		// _ = "end of CoverTab[119647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:645
		_go_fuzz_dep_.CoverTab[119648]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:645
		// _ = "end of CoverTab[119648]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:645
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:645
	// _ = "end of CoverTab[119641]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:645
	_go_fuzz_dep_.CoverTab[119642]++

										if wrap < 24 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:647
		_go_fuzz_dep_.CoverTab[119649]++
											return strings.Replace(s, "\n", r, -1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:648
		// _ = "end of CoverTab[119649]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:649
		_go_fuzz_dep_.CoverTab[119650]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:649
		// _ = "end of CoverTab[119650]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:649
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:649
	// _ = "end of CoverTab[119642]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:649
	_go_fuzz_dep_.CoverTab[119643]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:654
	slop := 5
										wrap = wrap - slop

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:659
	l, s = wrapN(wrap, slop, s)
										r = r + strings.Replace(l, "\n", "\n"+strings.Repeat(" ", i), -1)

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:663
	for s != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:663
		_go_fuzz_dep_.CoverTab[119651]++
											var t string

											t, s = wrapN(wrap, slop, s)
											r = r + "\n" + strings.Repeat(" ", i) + strings.Replace(t, "\n", "\n"+strings.Repeat(" ", i), -1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:667
		// _ = "end of CoverTab[119651]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:668
	// _ = "end of CoverTab[119643]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:668
	_go_fuzz_dep_.CoverTab[119644]++

										return r
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:670
	// _ = "end of CoverTab[119644]"

}

// FlagUsagesWrapped returns a string containing the usage information
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:674
// for all flags in the FlagSet. Wrapped to `cols` columns (0 for no
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:674
// wrapping)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:677
func (f *FlagSet) FlagUsagesWrapped(cols int) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:677
	_go_fuzz_dep_.CoverTab[119652]++
										buf := new(bytes.Buffer)

										lines := make([]string, 0, len(f.formal))

										maxlen := 0
										f.VisitAll(func(flag *Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:683
		_go_fuzz_dep_.CoverTab[119655]++
											if flag.Hidden {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:684
			_go_fuzz_dep_.CoverTab[119663]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:685
			// _ = "end of CoverTab[119663]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:686
			_go_fuzz_dep_.CoverTab[119664]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:686
			// _ = "end of CoverTab[119664]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:686
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:686
		// _ = "end of CoverTab[119655]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:686
		_go_fuzz_dep_.CoverTab[119656]++

											line := ""
											if flag.Shorthand != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:689
			_go_fuzz_dep_.CoverTab[119665]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:689
			return flag.ShorthandDeprecated == ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:689
			// _ = "end of CoverTab[119665]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:689
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:689
			_go_fuzz_dep_.CoverTab[119666]++
												line = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:690
			// _ = "end of CoverTab[119666]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:691
			_go_fuzz_dep_.CoverTab[119667]++
												line = fmt.Sprintf("      --%s", flag.Name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:692
			// _ = "end of CoverTab[119667]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:693
		// _ = "end of CoverTab[119656]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:693
		_go_fuzz_dep_.CoverTab[119657]++

											varname, usage := UnquoteUsage(flag)
											if varname != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:696
			_go_fuzz_dep_.CoverTab[119668]++
												line += " " + varname
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:697
			// _ = "end of CoverTab[119668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:698
			_go_fuzz_dep_.CoverTab[119669]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:698
			// _ = "end of CoverTab[119669]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:698
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:698
		// _ = "end of CoverTab[119657]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:698
		_go_fuzz_dep_.CoverTab[119658]++
											if flag.NoOptDefVal != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:699
			_go_fuzz_dep_.CoverTab[119670]++
												switch flag.Value.Type() {
			case "string":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:701
				_go_fuzz_dep_.CoverTab[119671]++
													line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:702
				// _ = "end of CoverTab[119671]"
			case "bool":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:703
				_go_fuzz_dep_.CoverTab[119672]++
													if flag.NoOptDefVal != "true" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:704
					_go_fuzz_dep_.CoverTab[119675]++
														line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:705
					// _ = "end of CoverTab[119675]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:706
					_go_fuzz_dep_.CoverTab[119676]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:706
					// _ = "end of CoverTab[119676]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:706
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:706
				// _ = "end of CoverTab[119672]"
			case "count":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:707
				_go_fuzz_dep_.CoverTab[119673]++
													if flag.NoOptDefVal != "+1" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:708
					_go_fuzz_dep_.CoverTab[119677]++
														line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:709
					// _ = "end of CoverTab[119677]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:710
					_go_fuzz_dep_.CoverTab[119678]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:710
					// _ = "end of CoverTab[119678]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:710
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:710
				// _ = "end of CoverTab[119673]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:711
				_go_fuzz_dep_.CoverTab[119674]++
													line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:712
				// _ = "end of CoverTab[119674]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:713
			// _ = "end of CoverTab[119670]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:714
			_go_fuzz_dep_.CoverTab[119679]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:714
			// _ = "end of CoverTab[119679]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:714
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:714
		// _ = "end of CoverTab[119658]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:714
		_go_fuzz_dep_.CoverTab[119659]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:718
		line += "\x00"
		if len(line) > maxlen {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:719
			_go_fuzz_dep_.CoverTab[119680]++
												maxlen = len(line)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:720
			// _ = "end of CoverTab[119680]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:721
			_go_fuzz_dep_.CoverTab[119681]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:721
			// _ = "end of CoverTab[119681]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:721
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:721
		// _ = "end of CoverTab[119659]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:721
		_go_fuzz_dep_.CoverTab[119660]++

											line += usage
											if !flag.defaultIsZeroValue() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:724
			_go_fuzz_dep_.CoverTab[119682]++
												if flag.Value.Type() == "string" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:725
				_go_fuzz_dep_.CoverTab[119683]++
													line += fmt.Sprintf(" (default %q)", flag.DefValue)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:726
				// _ = "end of CoverTab[119683]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:727
				_go_fuzz_dep_.CoverTab[119684]++
													line += fmt.Sprintf(" (default %s)", flag.DefValue)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:728
				// _ = "end of CoverTab[119684]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:729
			// _ = "end of CoverTab[119682]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:730
			_go_fuzz_dep_.CoverTab[119685]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:730
			// _ = "end of CoverTab[119685]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:730
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:730
		// _ = "end of CoverTab[119660]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:730
		_go_fuzz_dep_.CoverTab[119661]++
											if len(flag.Deprecated) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:731
			_go_fuzz_dep_.CoverTab[119686]++
												line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:732
			// _ = "end of CoverTab[119686]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:733
			_go_fuzz_dep_.CoverTab[119687]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:733
			// _ = "end of CoverTab[119687]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:733
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:733
		// _ = "end of CoverTab[119661]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:733
		_go_fuzz_dep_.CoverTab[119662]++

											lines = append(lines, line)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:735
		// _ = "end of CoverTab[119662]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:736
	// _ = "end of CoverTab[119652]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:736
	_go_fuzz_dep_.CoverTab[119653]++

										for _, line := range lines {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:738
		_go_fuzz_dep_.CoverTab[119688]++
											sidx := strings.Index(line, "\x00")
											spacing := strings.Repeat(" ", maxlen-sidx)

											fmt.Fprintln(buf, line[:sidx], spacing, wrap(maxlen+2, cols, line[sidx+1:]))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:742
		// _ = "end of CoverTab[119688]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:743
	// _ = "end of CoverTab[119653]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:743
	_go_fuzz_dep_.CoverTab[119654]++

										return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:745
	// _ = "end of CoverTab[119654]"
}

// FlagUsages returns a string containing the usage information for all flags in
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:748
// the FlagSet
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:750
func (f *FlagSet) FlagUsages() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:750
	_go_fuzz_dep_.CoverTab[119689]++
										return f.FlagUsagesWrapped(0)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:751
	// _ = "end of CoverTab[119689]"
}

// PrintDefaults prints to standard error the default values of all defined command-line flags.
func PrintDefaults() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:755
	_go_fuzz_dep_.CoverTab[119690]++
										CommandLine.PrintDefaults()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:756
	// _ = "end of CoverTab[119690]"
}

// defaultUsage is the default function to print a usage message.
func defaultUsage(f *FlagSet) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:760
	_go_fuzz_dep_.CoverTab[119691]++
										fmt.Fprintf(f.out(), "Usage of %s:\n", f.name)
										f.PrintDefaults()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:762
	// _ = "end of CoverTab[119691]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:769
// Usage prints to standard error a usage message documenting all defined command-line flags.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:769
// The function is a variable that may be changed to point to a custom function.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:769
// By default it prints a simple header and calls PrintDefaults; for details about the
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:769
// format of the output and how to control it, see the documentation for PrintDefaults.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:773
var Usage = func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:773
	_go_fuzz_dep_.CoverTab[119692]++
										fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
										PrintDefaults()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:775
	// _ = "end of CoverTab[119692]"
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:779
	_go_fuzz_dep_.CoverTab[119693]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:779
	return len(f.actual)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:779
	// _ = "end of CoverTab[119693]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:779
}

// NFlag returns the number of command-line flags that have been set.
func NFlag() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:782
	_go_fuzz_dep_.CoverTab[119694]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:782
	return len(CommandLine.actual)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:782
	// _ = "end of CoverTab[119694]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:782
}

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:784
// after flags have been processed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:786
func (f *FlagSet) Arg(i int) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:786
	_go_fuzz_dep_.CoverTab[119695]++
										if i < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:787
		_go_fuzz_dep_.CoverTab[119697]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:787
		return i >= len(f.args)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:787
		// _ = "end of CoverTab[119697]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:787
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:787
		_go_fuzz_dep_.CoverTab[119698]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:788
		// _ = "end of CoverTab[119698]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:789
		_go_fuzz_dep_.CoverTab[119699]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:789
		// _ = "end of CoverTab[119699]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:789
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:789
	// _ = "end of CoverTab[119695]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:789
	_go_fuzz_dep_.CoverTab[119696]++
										return f.args[i]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:790
	// _ = "end of CoverTab[119696]"
}

// Arg returns the i'th command-line argument.  Arg(0) is the first remaining argument
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:793
// after flags have been processed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:795
func Arg(i int) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:795
	_go_fuzz_dep_.CoverTab[119700]++
										return CommandLine.Arg(i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:796
	// _ = "end of CoverTab[119700]"
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:800
	_go_fuzz_dep_.CoverTab[119701]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:800
	return len(f.args)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:800
	// _ = "end of CoverTab[119701]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:800
}

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:803
	_go_fuzz_dep_.CoverTab[119702]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:803
	return len(CommandLine.args)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:803
	// _ = "end of CoverTab[119702]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:803
}

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:806
	_go_fuzz_dep_.CoverTab[119703]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:806
	return f.args
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:806
	// _ = "end of CoverTab[119703]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:806
}

// Args returns the non-flag command-line arguments.
func Args() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:809
	_go_fuzz_dep_.CoverTab[119704]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:809
	return CommandLine.args
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:809
	// _ = "end of CoverTab[119704]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:809
}

// Var defines a flag with the specified name and usage string. The type and
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:811
// value of the flag are represented by the first argument, of type Value, which
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:811
// typically holds a user-defined implementation of Value. For instance, the
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:811
// caller could create a flag that turns a comma-separated string into a slice
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:811
// of strings by giving the slice the methods of Value; in particular, Set would
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:811
// decompose the comma-separated string into the slice.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:817
func (f *FlagSet) Var(value Value, name string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:817
	_go_fuzz_dep_.CoverTab[119705]++
										f.VarP(value, name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:818
	// _ = "end of CoverTab[119705]"
}

// VarPF is like VarP, but returns the flag created
func (f *FlagSet) VarPF(value Value, name, shorthand, usage string) *Flag {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:822
	_go_fuzz_dep_.CoverTab[119706]++

										flag := &Flag{
		Name:		name,
		Shorthand:	shorthand,
		Usage:		usage,
		Value:		value,
		DefValue:	value.String(),
	}
										f.AddFlag(flag)
										return flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:832
	// _ = "end of CoverTab[119706]"
}

// VarP is like Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) VarP(value Value, name, shorthand, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:836
	_go_fuzz_dep_.CoverTab[119707]++
										f.VarPF(value, name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:837
	// _ = "end of CoverTab[119707]"
}

// AddFlag will add the flag to the FlagSet
func (f *FlagSet) AddFlag(flag *Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:841
	_go_fuzz_dep_.CoverTab[119708]++
										normalizedFlagName := f.normalizeFlagName(flag.Name)

										_, alreadyThere := f.formal[normalizedFlagName]
										if alreadyThere {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:845
		_go_fuzz_dep_.CoverTab[119715]++
											msg := fmt.Sprintf("%s flag redefined: %s", f.name, flag.Name)
											fmt.Fprintln(f.out(), msg)
											panic(msg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:848
		// _ = "end of CoverTab[119715]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:849
		_go_fuzz_dep_.CoverTab[119716]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:849
		// _ = "end of CoverTab[119716]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:849
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:849
	// _ = "end of CoverTab[119708]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:849
	_go_fuzz_dep_.CoverTab[119709]++
										if f.formal == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:850
		_go_fuzz_dep_.CoverTab[119717]++
											f.formal = make(map[NormalizedName]*Flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:851
		// _ = "end of CoverTab[119717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:852
		_go_fuzz_dep_.CoverTab[119718]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:852
		// _ = "end of CoverTab[119718]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:852
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:852
	// _ = "end of CoverTab[119709]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:852
	_go_fuzz_dep_.CoverTab[119710]++

										flag.Name = string(normalizedFlagName)
										f.formal[normalizedFlagName] = flag
										f.orderedFormal = append(f.orderedFormal, flag)

										if flag.Shorthand == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:858
		_go_fuzz_dep_.CoverTab[119719]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:859
		// _ = "end of CoverTab[119719]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:860
		_go_fuzz_dep_.CoverTab[119720]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:860
		// _ = "end of CoverTab[119720]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:860
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:860
	// _ = "end of CoverTab[119710]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:860
	_go_fuzz_dep_.CoverTab[119711]++
										if len(flag.Shorthand) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:861
		_go_fuzz_dep_.CoverTab[119721]++
											msg := fmt.Sprintf("%q shorthand is more than one ASCII character", flag.Shorthand)
											fmt.Fprintf(f.out(), msg)
											panic(msg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:864
		// _ = "end of CoverTab[119721]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:865
		_go_fuzz_dep_.CoverTab[119722]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:865
		// _ = "end of CoverTab[119722]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:865
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:865
	// _ = "end of CoverTab[119711]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:865
	_go_fuzz_dep_.CoverTab[119712]++
										if f.shorthands == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:866
		_go_fuzz_dep_.CoverTab[119723]++
											f.shorthands = make(map[byte]*Flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:867
		// _ = "end of CoverTab[119723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:868
		_go_fuzz_dep_.CoverTab[119724]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:868
		// _ = "end of CoverTab[119724]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:868
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:868
	// _ = "end of CoverTab[119712]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:868
	_go_fuzz_dep_.CoverTab[119713]++
										c := flag.Shorthand[0]
										used, alreadyThere := f.shorthands[c]
										if alreadyThere {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:871
		_go_fuzz_dep_.CoverTab[119725]++
											msg := fmt.Sprintf("unable to redefine %q shorthand in %q flagset: it's already used for %q flag", c, f.name, used.Name)
											fmt.Fprintf(f.out(), msg)
											panic(msg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:874
		// _ = "end of CoverTab[119725]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:875
		_go_fuzz_dep_.CoverTab[119726]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:875
		// _ = "end of CoverTab[119726]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:875
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:875
	// _ = "end of CoverTab[119713]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:875
	_go_fuzz_dep_.CoverTab[119714]++
										f.shorthands[c] = flag
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:876
	// _ = "end of CoverTab[119714]"
}

// AddFlagSet adds one FlagSet to another. If a flag is already present in f
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:879
// the flag from newSet will be ignored.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:881
func (f *FlagSet) AddFlagSet(newSet *FlagSet) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:881
	_go_fuzz_dep_.CoverTab[119727]++
										if newSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:882
		_go_fuzz_dep_.CoverTab[119729]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:883
		// _ = "end of CoverTab[119729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:884
		_go_fuzz_dep_.CoverTab[119730]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:884
		// _ = "end of CoverTab[119730]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:884
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:884
	// _ = "end of CoverTab[119727]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:884
	_go_fuzz_dep_.CoverTab[119728]++
										newSet.VisitAll(func(flag *Flag) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:885
		_go_fuzz_dep_.CoverTab[119731]++
											if f.Lookup(flag.Name) == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:886
			_go_fuzz_dep_.CoverTab[119732]++
												f.AddFlag(flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:887
			// _ = "end of CoverTab[119732]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:888
			_go_fuzz_dep_.CoverTab[119733]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:888
			// _ = "end of CoverTab[119733]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:888
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:888
		// _ = "end of CoverTab[119731]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:889
	// _ = "end of CoverTab[119728]"
}

// Var defines a flag with the specified name and usage string. The type and
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:892
// value of the flag are represented by the first argument, of type Value, which
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:892
// typically holds a user-defined implementation of Value. For instance, the
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:892
// caller could create a flag that turns a comma-separated string into a slice
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:892
// of strings by giving the slice the methods of Value; in particular, Set would
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:892
// decompose the comma-separated string into the slice.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:898
func Var(value Value, name string, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:898
	_go_fuzz_dep_.CoverTab[119734]++
										CommandLine.VarP(value, name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:899
	// _ = "end of CoverTab[119734]"
}

// VarP is like Var, but accepts a shorthand letter that can be used after a single dash.
func VarP(value Value, name, shorthand, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:903
	_go_fuzz_dep_.CoverTab[119735]++
										CommandLine.VarP(value, name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:904
	// _ = "end of CoverTab[119735]"
}

// failf prints to standard error a formatted error and usage message and
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:907
// returns the error.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:909
func (f *FlagSet) failf(format string, a ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:909
	_go_fuzz_dep_.CoverTab[119736]++
										err := fmt.Errorf(format, a...)
										if f.errorHandling != ContinueOnError {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:911
		_go_fuzz_dep_.CoverTab[119738]++
											fmt.Fprintln(f.out(), err)
											f.usage()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:913
		// _ = "end of CoverTab[119738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:914
		_go_fuzz_dep_.CoverTab[119739]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:914
		// _ = "end of CoverTab[119739]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:914
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:914
	// _ = "end of CoverTab[119736]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:914
	_go_fuzz_dep_.CoverTab[119737]++
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:915
	// _ = "end of CoverTab[119737]"
}

// usage calls the Usage method for the flag set, or the usage function if
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:918
// the flag set is CommandLine.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:920
func (f *FlagSet) usage() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:920
	_go_fuzz_dep_.CoverTab[119740]++
										if f == CommandLine {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:921
		_go_fuzz_dep_.CoverTab[119741]++
											Usage()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:922
		// _ = "end of CoverTab[119741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:923
		_go_fuzz_dep_.CoverTab[119742]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:923
		if f.Usage == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:923
			_go_fuzz_dep_.CoverTab[119743]++
												defaultUsage(f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:924
			// _ = "end of CoverTab[119743]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:925
			_go_fuzz_dep_.CoverTab[119744]++
												f.Usage()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:926
			// _ = "end of CoverTab[119744]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:927
		// _ = "end of CoverTab[119742]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:927
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:927
	// _ = "end of CoverTab[119740]"
}

// --unknown (args will be empty)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:930
// --unknown --next-flag ... (args will be --next-flag ...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:930
// --unknown arg ... (args will be arg ...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:933
func stripUnknownFlagValue(args []string) []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:933
	_go_fuzz_dep_.CoverTab[119745]++
										if len(args) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:934
		_go_fuzz_dep_.CoverTab[119749]++

											return args
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:936
		// _ = "end of CoverTab[119749]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:937
		_go_fuzz_dep_.CoverTab[119750]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:937
		// _ = "end of CoverTab[119750]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:937
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:937
	// _ = "end of CoverTab[119745]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:937
	_go_fuzz_dep_.CoverTab[119746]++

										first := args[0]
										if len(first) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:940
		_go_fuzz_dep_.CoverTab[119751]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:940
		return first[0] == '-'
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:940
		// _ = "end of CoverTab[119751]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:940
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:940
		_go_fuzz_dep_.CoverTab[119752]++

											return args
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:942
		// _ = "end of CoverTab[119752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:943
		_go_fuzz_dep_.CoverTab[119753]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:943
		// _ = "end of CoverTab[119753]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:943
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:943
	// _ = "end of CoverTab[119746]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:943
	_go_fuzz_dep_.CoverTab[119747]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:946
	if len(args) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:946
		_go_fuzz_dep_.CoverTab[119754]++
											return args[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:947
		// _ = "end of CoverTab[119754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:948
		_go_fuzz_dep_.CoverTab[119755]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:948
		// _ = "end of CoverTab[119755]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:948
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:948
	// _ = "end of CoverTab[119747]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:948
	_go_fuzz_dep_.CoverTab[119748]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:949
	// _ = "end of CoverTab[119748]"
}

func (f *FlagSet) parseLongArg(s string, args []string, fn parseFunc) (a []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:952
	_go_fuzz_dep_.CoverTab[119756]++
										a = args
										name := s[2:]
										if len(name) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		_go_fuzz_dep_.CoverTab[119761]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		return name[0] == '-'
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		// _ = "end of CoverTab[119761]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		_go_fuzz_dep_.CoverTab[119762]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		return name[0] == '='
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		// _ = "end of CoverTab[119762]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:955
		_go_fuzz_dep_.CoverTab[119763]++
											err = f.failf("bad flag syntax: %s", s)
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:957
		// _ = "end of CoverTab[119763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:958
		_go_fuzz_dep_.CoverTab[119764]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:958
		// _ = "end of CoverTab[119764]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:958
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:958
	// _ = "end of CoverTab[119756]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:958
	_go_fuzz_dep_.CoverTab[119757]++

										split := strings.SplitN(name, "=", 2)
										name = split[0]
										flag, exists := f.formal[f.normalizeFlagName(name)]

										if !exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:964
		_go_fuzz_dep_.CoverTab[119765]++
											switch {
		case name == "help":
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:966
			_go_fuzz_dep_.CoverTab[119766]++
												f.usage()
												return a, ErrHelp
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:968
			// _ = "end of CoverTab[119766]"
		case f.ParseErrorsWhitelist.UnknownFlags:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:969
			_go_fuzz_dep_.CoverTab[119767]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:972
			if len(split) >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:972
				_go_fuzz_dep_.CoverTab[119770]++
													return a, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:973
				// _ = "end of CoverTab[119770]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:974
				_go_fuzz_dep_.CoverTab[119771]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:974
				// _ = "end of CoverTab[119771]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:974
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:974
			// _ = "end of CoverTab[119767]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:974
			_go_fuzz_dep_.CoverTab[119768]++

												return stripUnknownFlagValue(a), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:976
			// _ = "end of CoverTab[119768]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:977
			_go_fuzz_dep_.CoverTab[119769]++
												err = f.failf("unknown flag: --%s", name)
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:979
			// _ = "end of CoverTab[119769]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:980
		// _ = "end of CoverTab[119765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:981
		_go_fuzz_dep_.CoverTab[119772]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:981
		// _ = "end of CoverTab[119772]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:981
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:981
	// _ = "end of CoverTab[119757]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:981
	_go_fuzz_dep_.CoverTab[119758]++

										var value string
										if len(split) == 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:984
		_go_fuzz_dep_.CoverTab[119773]++

											value = split[1]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:986
		// _ = "end of CoverTab[119773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:987
		_go_fuzz_dep_.CoverTab[119774]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:987
		if flag.NoOptDefVal != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:987
			_go_fuzz_dep_.CoverTab[119775]++

												value = flag.NoOptDefVal
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:989
			// _ = "end of CoverTab[119775]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:990
			_go_fuzz_dep_.CoverTab[119776]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:990
			if len(a) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:990
				_go_fuzz_dep_.CoverTab[119777]++

													value = a[0]
													a = a[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:993
				// _ = "end of CoverTab[119777]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:994
				_go_fuzz_dep_.CoverTab[119778]++

													err = f.failf("flag needs an argument: %s", s)
													return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:997
				// _ = "end of CoverTab[119778]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
			// _ = "end of CoverTab[119776]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
		// _ = "end of CoverTab[119774]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
	// _ = "end of CoverTab[119758]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:998
	_go_fuzz_dep_.CoverTab[119759]++

										err = fn(flag, value)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1001
		_go_fuzz_dep_.CoverTab[119779]++
											f.failf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1002
		// _ = "end of CoverTab[119779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1003
		_go_fuzz_dep_.CoverTab[119780]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1003
		// _ = "end of CoverTab[119780]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1003
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1003
	// _ = "end of CoverTab[119759]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1003
	_go_fuzz_dep_.CoverTab[119760]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1004
	// _ = "end of CoverTab[119760]"
}

func (f *FlagSet) parseSingleShortArg(shorthands string, args []string, fn parseFunc) (outShorts string, outArgs []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1007
	_go_fuzz_dep_.CoverTab[119781]++
										outArgs = args

										if strings.HasPrefix(shorthands, "test.") {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1010
		_go_fuzz_dep_.CoverTab[119787]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1011
		// _ = "end of CoverTab[119787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1012
		_go_fuzz_dep_.CoverTab[119788]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1012
		// _ = "end of CoverTab[119788]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1012
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1012
	// _ = "end of CoverTab[119781]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1012
	_go_fuzz_dep_.CoverTab[119782]++

										outShorts = shorthands[1:]
										c := shorthands[0]

										flag, exists := f.shorthands[c]
										if !exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1018
		_go_fuzz_dep_.CoverTab[119789]++
											switch {
		case c == 'h':
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1020
			_go_fuzz_dep_.CoverTab[119790]++
												f.usage()
												err = ErrHelp
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1023
			// _ = "end of CoverTab[119790]"
		case f.ParseErrorsWhitelist.UnknownFlags:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1024
			_go_fuzz_dep_.CoverTab[119791]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
			if len(shorthands) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
				_go_fuzz_dep_.CoverTab[119794]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
				return shorthands[1] == '='
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
				// _ = "end of CoverTab[119794]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
			}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1027
				_go_fuzz_dep_.CoverTab[119795]++
													outShorts = ""
													return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1029
				// _ = "end of CoverTab[119795]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1030
				_go_fuzz_dep_.CoverTab[119796]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1030
				// _ = "end of CoverTab[119796]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1030
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1030
			// _ = "end of CoverTab[119791]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1030
			_go_fuzz_dep_.CoverTab[119792]++

												outArgs = stripUnknownFlagValue(outArgs)
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1033
			// _ = "end of CoverTab[119792]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1034
			_go_fuzz_dep_.CoverTab[119793]++
												err = f.failf("unknown shorthand flag: %q in -%s", c, shorthands)
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1036
			// _ = "end of CoverTab[119793]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1037
		// _ = "end of CoverTab[119789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1038
		_go_fuzz_dep_.CoverTab[119797]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1038
		// _ = "end of CoverTab[119797]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1038
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1038
	// _ = "end of CoverTab[119782]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1038
	_go_fuzz_dep_.CoverTab[119783]++

										var value string
										if len(shorthands) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1041
		_go_fuzz_dep_.CoverTab[119798]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1041
		return shorthands[1] == '='
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1041
		// _ = "end of CoverTab[119798]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1041
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1041
		_go_fuzz_dep_.CoverTab[119799]++

											value = shorthands[2:]
											outShorts = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1044
		// _ = "end of CoverTab[119799]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1045
		_go_fuzz_dep_.CoverTab[119800]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1045
		if flag.NoOptDefVal != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1045
			_go_fuzz_dep_.CoverTab[119801]++

												value = flag.NoOptDefVal
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1047
			// _ = "end of CoverTab[119801]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1048
			_go_fuzz_dep_.CoverTab[119802]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1048
			if len(shorthands) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1048
				_go_fuzz_dep_.CoverTab[119803]++

													value = shorthands[1:]
													outShorts = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1051
				// _ = "end of CoverTab[119803]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1052
				_go_fuzz_dep_.CoverTab[119804]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1052
				if len(args) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1052
					_go_fuzz_dep_.CoverTab[119805]++

														value = args[0]
														outArgs = args[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1055
					// _ = "end of CoverTab[119805]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1056
					_go_fuzz_dep_.CoverTab[119806]++

														err = f.failf("flag needs an argument: %q in -%s", c, shorthands)
														return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1059
					// _ = "end of CoverTab[119806]"
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
				// _ = "end of CoverTab[119804]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
			// _ = "end of CoverTab[119802]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
		// _ = "end of CoverTab[119800]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
	// _ = "end of CoverTab[119783]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1060
	_go_fuzz_dep_.CoverTab[119784]++

										if flag.ShorthandDeprecated != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1062
		_go_fuzz_dep_.CoverTab[119807]++
											fmt.Fprintf(f.out(), "Flag shorthand -%s has been deprecated, %s\n", flag.Shorthand, flag.ShorthandDeprecated)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1063
		// _ = "end of CoverTab[119807]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1064
		_go_fuzz_dep_.CoverTab[119808]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1064
		// _ = "end of CoverTab[119808]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1064
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1064
	// _ = "end of CoverTab[119784]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1064
	_go_fuzz_dep_.CoverTab[119785]++

										err = fn(flag, value)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1067
		_go_fuzz_dep_.CoverTab[119809]++
											f.failf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1068
		// _ = "end of CoverTab[119809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1069
		_go_fuzz_dep_.CoverTab[119810]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1069
		// _ = "end of CoverTab[119810]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1069
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1069
	// _ = "end of CoverTab[119785]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1069
	_go_fuzz_dep_.CoverTab[119786]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1070
	// _ = "end of CoverTab[119786]"
}

func (f *FlagSet) parseShortArg(s string, args []string, fn parseFunc) (a []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1073
	_go_fuzz_dep_.CoverTab[119811]++
										a = args
										shorthands := s[1:]

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1078
	for len(shorthands) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1078
		_go_fuzz_dep_.CoverTab[119813]++
											shorthands, a, err = f.parseSingleShortArg(shorthands, args, fn)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1080
			_go_fuzz_dep_.CoverTab[119814]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1081
			// _ = "end of CoverTab[119814]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1082
			_go_fuzz_dep_.CoverTab[119815]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1082
			// _ = "end of CoverTab[119815]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1082
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1082
		// _ = "end of CoverTab[119813]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1083
	// _ = "end of CoverTab[119811]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1083
	_go_fuzz_dep_.CoverTab[119812]++

										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1085
	// _ = "end of CoverTab[119812]"
}

func (f *FlagSet) parseArgs(args []string, fn parseFunc) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1088
	_go_fuzz_dep_.CoverTab[119816]++
										for len(args) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1089
		_go_fuzz_dep_.CoverTab[119818]++
											s := args[0]
											args = args[1:]
											if len(s) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			_go_fuzz_dep_.CoverTab[119821]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			return s[0] != '-'
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			// _ = "end of CoverTab[119821]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			_go_fuzz_dep_.CoverTab[119822]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			return len(s) == 1
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			// _ = "end of CoverTab[119822]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1092
			_go_fuzz_dep_.CoverTab[119823]++
												if !f.interspersed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1093
				_go_fuzz_dep_.CoverTab[119825]++
													f.args = append(f.args, s)
													f.args = append(f.args, args...)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1096
				// _ = "end of CoverTab[119825]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1097
				_go_fuzz_dep_.CoverTab[119826]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1097
				// _ = "end of CoverTab[119826]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1097
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1097
			// _ = "end of CoverTab[119823]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1097
			_go_fuzz_dep_.CoverTab[119824]++
												f.args = append(f.args, s)
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1099
			// _ = "end of CoverTab[119824]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1100
			_go_fuzz_dep_.CoverTab[119827]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1100
			// _ = "end of CoverTab[119827]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1100
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1100
		// _ = "end of CoverTab[119818]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1100
		_go_fuzz_dep_.CoverTab[119819]++

											if s[1] == '-' {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1102
			_go_fuzz_dep_.CoverTab[119828]++
												if len(s) == 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1103
				_go_fuzz_dep_.CoverTab[119830]++
													f.argsLenAtDash = len(f.args)
													f.args = append(f.args, args...)
													break
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1106
				// _ = "end of CoverTab[119830]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1107
				_go_fuzz_dep_.CoverTab[119831]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1107
				// _ = "end of CoverTab[119831]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1107
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1107
			// _ = "end of CoverTab[119828]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1107
			_go_fuzz_dep_.CoverTab[119829]++
												args, err = f.parseLongArg(s, args, fn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1108
			// _ = "end of CoverTab[119829]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1109
			_go_fuzz_dep_.CoverTab[119832]++
												args, err = f.parseShortArg(s, args, fn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1110
			// _ = "end of CoverTab[119832]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1111
		// _ = "end of CoverTab[119819]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1111
		_go_fuzz_dep_.CoverTab[119820]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1112
			_go_fuzz_dep_.CoverTab[119833]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1113
			// _ = "end of CoverTab[119833]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1114
			_go_fuzz_dep_.CoverTab[119834]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1114
			// _ = "end of CoverTab[119834]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1114
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1114
		// _ = "end of CoverTab[119820]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1115
	// _ = "end of CoverTab[119816]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1115
	_go_fuzz_dep_.CoverTab[119817]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1116
	// _ = "end of CoverTab[119817]"
}

// Parse parses flag definitions from the argument list, which should not
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1119
// include the command name.  Must be called after all flags in the FlagSet
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1119
// are defined and before flags are accessed by the program.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1119
// The return value will be ErrHelp if -help was set but not defined.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1123
func (f *FlagSet) Parse(arguments []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1123
	_go_fuzz_dep_.CoverTab[119835]++
										if f.addedGoFlagSets != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1124
		_go_fuzz_dep_.CoverTab[119840]++
											for _, goFlagSet := range f.addedGoFlagSets {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1125
			_go_fuzz_dep_.CoverTab[119841]++
												goFlagSet.Parse(nil)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1126
			// _ = "end of CoverTab[119841]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1127
		// _ = "end of CoverTab[119840]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1128
		_go_fuzz_dep_.CoverTab[119842]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1128
		// _ = "end of CoverTab[119842]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1128
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1128
	// _ = "end of CoverTab[119835]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1128
	_go_fuzz_dep_.CoverTab[119836]++
										f.parsed = true

										if len(arguments) < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1131
		_go_fuzz_dep_.CoverTab[119843]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1132
		// _ = "end of CoverTab[119843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1133
		_go_fuzz_dep_.CoverTab[119844]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1133
		// _ = "end of CoverTab[119844]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1133
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1133
	// _ = "end of CoverTab[119836]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1133
	_go_fuzz_dep_.CoverTab[119837]++

										f.args = make([]string, 0, len(arguments))

										set := func(flag *Flag, value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1137
		_go_fuzz_dep_.CoverTab[119845]++
											return f.Set(flag.Name, value)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1138
		// _ = "end of CoverTab[119845]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1139
	// _ = "end of CoverTab[119837]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1139
	_go_fuzz_dep_.CoverTab[119838]++

										err := f.parseArgs(arguments, set)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1142
		_go_fuzz_dep_.CoverTab[119846]++
											switch f.errorHandling {
		case ContinueOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1144
			_go_fuzz_dep_.CoverTab[119847]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1145
			// _ = "end of CoverTab[119847]"
		case ExitOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1146
			_go_fuzz_dep_.CoverTab[119848]++
												fmt.Println(err)
												os.Exit(2)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1148
			// _ = "end of CoverTab[119848]"
		case PanicOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1149
			_go_fuzz_dep_.CoverTab[119849]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1150
			// _ = "end of CoverTab[119849]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1150
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1150
			_go_fuzz_dep_.CoverTab[119850]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1150
			// _ = "end of CoverTab[119850]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1151
		// _ = "end of CoverTab[119846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1152
		_go_fuzz_dep_.CoverTab[119851]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1152
		// _ = "end of CoverTab[119851]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1152
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1152
	// _ = "end of CoverTab[119838]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1152
	_go_fuzz_dep_.CoverTab[119839]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1153
	// _ = "end of CoverTab[119839]"
}

type parseFunc func(flag *Flag, value string) error

// ParseAll parses flag definitions from the argument list, which should not
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1158
// include the command name. The arguments for fn are flag and value. Must be
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1158
// called after all flags in the FlagSet are defined and before flags are
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1158
// accessed by the program. The return value will be ErrHelp if -help was set
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1158
// but not defined.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1163
func (f *FlagSet) ParseAll(arguments []string, fn func(flag *Flag, value string) error) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1163
	_go_fuzz_dep_.CoverTab[119852]++
										f.parsed = true
										f.args = make([]string, 0, len(arguments))

										err := f.parseArgs(arguments, fn)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1168
		_go_fuzz_dep_.CoverTab[119854]++
											switch f.errorHandling {
		case ContinueOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1170
			_go_fuzz_dep_.CoverTab[119855]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1171
			// _ = "end of CoverTab[119855]"
		case ExitOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1172
			_go_fuzz_dep_.CoverTab[119856]++
												os.Exit(2)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1173
			// _ = "end of CoverTab[119856]"
		case PanicOnError:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1174
			_go_fuzz_dep_.CoverTab[119857]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1175
			// _ = "end of CoverTab[119857]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1175
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1175
			_go_fuzz_dep_.CoverTab[119858]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1175
			// _ = "end of CoverTab[119858]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1176
		// _ = "end of CoverTab[119854]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1177
		_go_fuzz_dep_.CoverTab[119859]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1177
		// _ = "end of CoverTab[119859]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1177
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1177
	// _ = "end of CoverTab[119852]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1177
	_go_fuzz_dep_.CoverTab[119853]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1178
	// _ = "end of CoverTab[119853]"
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1182
	_go_fuzz_dep_.CoverTab[119860]++
										return f.parsed
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1183
	// _ = "end of CoverTab[119860]"
}

// Parse parses the command-line flags from os.Args[1:].  Must be called
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1186
// after all flags are defined and before flags are accessed by the program.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1188
func Parse() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1188
	_go_fuzz_dep_.CoverTab[119861]++

										CommandLine.Parse(os.Args[1:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1190
	// _ = "end of CoverTab[119861]"
}

// ParseAll parses the command-line flags from os.Args[1:] and called fn for each.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1193
// The arguments for fn are flag and value. Must be called after all flags are
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1193
// defined and before flags are accessed by the program.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1196
func ParseAll(fn func(flag *Flag, value string) error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1196
	_go_fuzz_dep_.CoverTab[119862]++

										CommandLine.ParseAll(os.Args[1:], fn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1198
	// _ = "end of CoverTab[119862]"
}

// SetInterspersed sets whether to support interspersed option/non-option arguments.
func SetInterspersed(interspersed bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1202
	_go_fuzz_dep_.CoverTab[119863]++
										CommandLine.SetInterspersed(interspersed)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1203
	// _ = "end of CoverTab[119863]"
}

// Parsed returns true if the command-line flags have been parsed.
func Parsed() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1207
	_go_fuzz_dep_.CoverTab[119864]++
										return CommandLine.Parsed()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1208
	// _ = "end of CoverTab[119864]"
}

// CommandLine is the default set of command-line flags, parsed from os.Args.
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

// NewFlagSet returns a new, empty flag set with the specified name,
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1214
// error handling property and SortFlags set to true.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1216
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1216
	_go_fuzz_dep_.CoverTab[119865]++
										f := &FlagSet{
		name:		name,
		errorHandling:	errorHandling,
		argsLenAtDash:	-1,
		interspersed:	true,
		SortFlags:	true,
	}
										return f
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1224
	// _ = "end of CoverTab[119865]"
}

// SetInterspersed sets whether to support interspersed option/non-option arguments.
func (f *FlagSet) SetInterspersed(interspersed bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1228
	_go_fuzz_dep_.CoverTab[119866]++
										f.interspersed = interspersed
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1229
	// _ = "end of CoverTab[119866]"
}

// Init sets the name and error handling property for a flag set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1232
// By default, the zero FlagSet uses an empty name and the
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1232
// ContinueOnError error handling policy.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1235
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1235
	_go_fuzz_dep_.CoverTab[119867]++
										f.name = name
										f.errorHandling = errorHandling
										f.argsLenAtDash = -1
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1238
	// _ = "end of CoverTab[119867]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1239
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/flag.go:1239
var _ = _go_fuzz_dep_.CoverTab
