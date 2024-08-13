// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/flag/flag.go:5
/*
Package flag implements command-line flag parsing.

# Usage

Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -n, stored in the pointer nFlag, with type *int:

	import "flag"
	var nFlag = flag.Int("n", 1234, "help message for flag n")

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

After parsing, the arguments following the flags are available as the
slice flag.Args() or individually as flag.Arg(i).
The arguments are indexed from 0 through flag.NArg()-1.

# Command line flag syntax

The following forms are permitted:

	-flag
	--flag   // double dashes are also permitted
	-flag=x
	-flag x  // non-boolean flags only

One or two dashes may be used; they are equivalent.
The last form is not permitted for boolean flags because the
meaning of the command

	cmd -x *

where * is a Unix shell wildcard, will change if there is a file
called 0, false, etc. You must use the -flag=false form to turn
off a boolean flag.

Flag parsing stops just before the first non-flag argument
("-" is a non-flag argument) or after the terminator "--".

Integer flags accept 1234, 0664, 0x1234 and may be negative.
Boolean flags may be:

	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by
top-level functions.  The FlagSet type allows one to define
independent sets of flags, such as to implement subcommands
in a command-line interface. The methods of FlagSet are
analogous to the top-level functions for the command-line
flag set.
*/
package flag

//line /usr/local/go/src/flag/flag.go:83
import (
//line /usr/local/go/src/flag/flag.go:83
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/flag/flag.go:83
)
//line /usr/local/go/src/flag/flag.go:83
import (
//line /usr/local/go/src/flag/flag.go:83
	_atomic_ "sync/atomic"
//line /usr/local/go/src/flag/flag.go:83
)

import (
	"encoding"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ErrHelp is the error returned if the -help or -h flag is invoked
//line /usr/local/go/src/flag/flag.go:98
// but no such flag is defined.
//line /usr/local/go/src/flag/flag.go:100
var ErrHelp = errors.New("flag: help requested")

// errParse is returned by Set if a flag's value fails to parse, such as with an invalid integer for Int.
//line /usr/local/go/src/flag/flag.go:102
// It then gets wrapped through failf to provide more information.
//line /usr/local/go/src/flag/flag.go:104
var errParse = errors.New("parse error")

// errRange is returned by Set if a flag's value is out of range.
//line /usr/local/go/src/flag/flag.go:106
// It then gets wrapped through failf to provide more information.
//line /usr/local/go/src/flag/flag.go:108
var errRange = errors.New("value out of range")

func numError(err error) error {
//line /usr/local/go/src/flag/flag.go:110
	_go_fuzz_dep_.CoverTab[115200]++
						ne, ok := err.(*strconv.NumError)
						if !ok {
//line /usr/local/go/src/flag/flag.go:112
		_go_fuzz_dep_.CoverTab[115204]++
							return err
//line /usr/local/go/src/flag/flag.go:113
		// _ = "end of CoverTab[115204]"
	} else {
//line /usr/local/go/src/flag/flag.go:114
		_go_fuzz_dep_.CoverTab[115205]++
//line /usr/local/go/src/flag/flag.go:114
		// _ = "end of CoverTab[115205]"
//line /usr/local/go/src/flag/flag.go:114
	}
//line /usr/local/go/src/flag/flag.go:114
	// _ = "end of CoverTab[115200]"
//line /usr/local/go/src/flag/flag.go:114
	_go_fuzz_dep_.CoverTab[115201]++
						if ne.Err == strconv.ErrSyntax {
//line /usr/local/go/src/flag/flag.go:115
		_go_fuzz_dep_.CoverTab[115206]++
							return errParse
//line /usr/local/go/src/flag/flag.go:116
		// _ = "end of CoverTab[115206]"
	} else {
//line /usr/local/go/src/flag/flag.go:117
		_go_fuzz_dep_.CoverTab[115207]++
//line /usr/local/go/src/flag/flag.go:117
		// _ = "end of CoverTab[115207]"
//line /usr/local/go/src/flag/flag.go:117
	}
//line /usr/local/go/src/flag/flag.go:117
	// _ = "end of CoverTab[115201]"
//line /usr/local/go/src/flag/flag.go:117
	_go_fuzz_dep_.CoverTab[115202]++
						if ne.Err == strconv.ErrRange {
//line /usr/local/go/src/flag/flag.go:118
		_go_fuzz_dep_.CoverTab[115208]++
							return errRange
//line /usr/local/go/src/flag/flag.go:119
		// _ = "end of CoverTab[115208]"
	} else {
//line /usr/local/go/src/flag/flag.go:120
		_go_fuzz_dep_.CoverTab[115209]++
//line /usr/local/go/src/flag/flag.go:120
		// _ = "end of CoverTab[115209]"
//line /usr/local/go/src/flag/flag.go:120
	}
//line /usr/local/go/src/flag/flag.go:120
	// _ = "end of CoverTab[115202]"
//line /usr/local/go/src/flag/flag.go:120
	_go_fuzz_dep_.CoverTab[115203]++
						return err
//line /usr/local/go/src/flag/flag.go:121
	// _ = "end of CoverTab[115203]"
}

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
//line /usr/local/go/src/flag/flag.go:127
	_go_fuzz_dep_.CoverTab[115210]++
						*p = val
						return (*boolValue)(p)
//line /usr/local/go/src/flag/flag.go:129
	// _ = "end of CoverTab[115210]"
}

func (b *boolValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:132
	_go_fuzz_dep_.CoverTab[115211]++
						v, err := strconv.ParseBool(s)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:134
		_go_fuzz_dep_.CoverTab[115213]++
							err = errParse
//line /usr/local/go/src/flag/flag.go:135
		// _ = "end of CoverTab[115213]"
	} else {
//line /usr/local/go/src/flag/flag.go:136
		_go_fuzz_dep_.CoverTab[115214]++
//line /usr/local/go/src/flag/flag.go:136
		// _ = "end of CoverTab[115214]"
//line /usr/local/go/src/flag/flag.go:136
	}
//line /usr/local/go/src/flag/flag.go:136
	// _ = "end of CoverTab[115211]"
//line /usr/local/go/src/flag/flag.go:136
	_go_fuzz_dep_.CoverTab[115212]++
						*b = boolValue(v)
						return err
//line /usr/local/go/src/flag/flag.go:138
	// _ = "end of CoverTab[115212]"
}

func (b *boolValue) Get() any {
//line /usr/local/go/src/flag/flag.go:141
	_go_fuzz_dep_.CoverTab[115215]++
//line /usr/local/go/src/flag/flag.go:141
	return bool(*b)
//line /usr/local/go/src/flag/flag.go:141
	// _ = "end of CoverTab[115215]"
//line /usr/local/go/src/flag/flag.go:141
}

func (b *boolValue) String() string {
//line /usr/local/go/src/flag/flag.go:143
	_go_fuzz_dep_.CoverTab[115216]++
//line /usr/local/go/src/flag/flag.go:143
	return strconv.FormatBool(bool(*b))
//line /usr/local/go/src/flag/flag.go:143
	// _ = "end of CoverTab[115216]"
//line /usr/local/go/src/flag/flag.go:143
}

func (b *boolValue) IsBoolFlag() bool {
//line /usr/local/go/src/flag/flag.go:145
	_go_fuzz_dep_.CoverTab[115217]++
//line /usr/local/go/src/flag/flag.go:145
	return true
//line /usr/local/go/src/flag/flag.go:145
	// _ = "end of CoverTab[115217]"
//line /usr/local/go/src/flag/flag.go:145
}

// optional interface to indicate boolean flags that can be
//line /usr/local/go/src/flag/flag.go:147
// supplied without "=value" text
//line /usr/local/go/src/flag/flag.go:149
type boolFlag interface {
	Value
	IsBoolFlag() bool
}

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
//line /usr/local/go/src/flag/flag.go:157
	_go_fuzz_dep_.CoverTab[115218]++
						*p = val
						return (*intValue)(p)
//line /usr/local/go/src/flag/flag.go:159
	// _ = "end of CoverTab[115218]"
}

func (i *intValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:162
	_go_fuzz_dep_.CoverTab[115219]++
						v, err := strconv.ParseInt(s, 0, strconv.IntSize)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:164
		_go_fuzz_dep_.CoverTab[115221]++
							err = numError(err)
//line /usr/local/go/src/flag/flag.go:165
		// _ = "end of CoverTab[115221]"
	} else {
//line /usr/local/go/src/flag/flag.go:166
		_go_fuzz_dep_.CoverTab[115222]++
//line /usr/local/go/src/flag/flag.go:166
		// _ = "end of CoverTab[115222]"
//line /usr/local/go/src/flag/flag.go:166
	}
//line /usr/local/go/src/flag/flag.go:166
	// _ = "end of CoverTab[115219]"
//line /usr/local/go/src/flag/flag.go:166
	_go_fuzz_dep_.CoverTab[115220]++
						*i = intValue(v)
						return err
//line /usr/local/go/src/flag/flag.go:168
	// _ = "end of CoverTab[115220]"
}

func (i *intValue) Get() any {
//line /usr/local/go/src/flag/flag.go:171
	_go_fuzz_dep_.CoverTab[115223]++
//line /usr/local/go/src/flag/flag.go:171
	return int(*i)
//line /usr/local/go/src/flag/flag.go:171
	// _ = "end of CoverTab[115223]"
//line /usr/local/go/src/flag/flag.go:171
}

func (i *intValue) String() string {
//line /usr/local/go/src/flag/flag.go:173
	_go_fuzz_dep_.CoverTab[115224]++
//line /usr/local/go/src/flag/flag.go:173
	return strconv.Itoa(int(*i))
//line /usr/local/go/src/flag/flag.go:173
	// _ = "end of CoverTab[115224]"
//line /usr/local/go/src/flag/flag.go:173
}

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
//line /usr/local/go/src/flag/flag.go:178
	_go_fuzz_dep_.CoverTab[115225]++
						*p = val
						return (*int64Value)(p)
//line /usr/local/go/src/flag/flag.go:180
	// _ = "end of CoverTab[115225]"
}

func (i *int64Value) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:183
	_go_fuzz_dep_.CoverTab[115226]++
						v, err := strconv.ParseInt(s, 0, 64)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:185
		_go_fuzz_dep_.CoverTab[115228]++
							err = numError(err)
//line /usr/local/go/src/flag/flag.go:186
		// _ = "end of CoverTab[115228]"
	} else {
//line /usr/local/go/src/flag/flag.go:187
		_go_fuzz_dep_.CoverTab[115229]++
//line /usr/local/go/src/flag/flag.go:187
		// _ = "end of CoverTab[115229]"
//line /usr/local/go/src/flag/flag.go:187
	}
//line /usr/local/go/src/flag/flag.go:187
	// _ = "end of CoverTab[115226]"
//line /usr/local/go/src/flag/flag.go:187
	_go_fuzz_dep_.CoverTab[115227]++
						*i = int64Value(v)
						return err
//line /usr/local/go/src/flag/flag.go:189
	// _ = "end of CoverTab[115227]"
}

func (i *int64Value) Get() any {
//line /usr/local/go/src/flag/flag.go:192
	_go_fuzz_dep_.CoverTab[115230]++
//line /usr/local/go/src/flag/flag.go:192
	return int64(*i)
//line /usr/local/go/src/flag/flag.go:192
	// _ = "end of CoverTab[115230]"
//line /usr/local/go/src/flag/flag.go:192
}

func (i *int64Value) String() string {
//line /usr/local/go/src/flag/flag.go:194
	_go_fuzz_dep_.CoverTab[115231]++
//line /usr/local/go/src/flag/flag.go:194
	return strconv.FormatInt(int64(*i), 10)
//line /usr/local/go/src/flag/flag.go:194
	// _ = "end of CoverTab[115231]"
//line /usr/local/go/src/flag/flag.go:194
}

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
//line /usr/local/go/src/flag/flag.go:199
	_go_fuzz_dep_.CoverTab[115232]++
						*p = val
						return (*uintValue)(p)
//line /usr/local/go/src/flag/flag.go:201
	// _ = "end of CoverTab[115232]"
}

func (i *uintValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:204
	_go_fuzz_dep_.CoverTab[115233]++
						v, err := strconv.ParseUint(s, 0, strconv.IntSize)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:206
		_go_fuzz_dep_.CoverTab[115235]++
							err = numError(err)
//line /usr/local/go/src/flag/flag.go:207
		// _ = "end of CoverTab[115235]"
	} else {
//line /usr/local/go/src/flag/flag.go:208
		_go_fuzz_dep_.CoverTab[115236]++
//line /usr/local/go/src/flag/flag.go:208
		// _ = "end of CoverTab[115236]"
//line /usr/local/go/src/flag/flag.go:208
	}
//line /usr/local/go/src/flag/flag.go:208
	// _ = "end of CoverTab[115233]"
//line /usr/local/go/src/flag/flag.go:208
	_go_fuzz_dep_.CoverTab[115234]++
						*i = uintValue(v)
						return err
//line /usr/local/go/src/flag/flag.go:210
	// _ = "end of CoverTab[115234]"
}

func (i *uintValue) Get() any {
//line /usr/local/go/src/flag/flag.go:213
	_go_fuzz_dep_.CoverTab[115237]++
//line /usr/local/go/src/flag/flag.go:213
	return uint(*i)
//line /usr/local/go/src/flag/flag.go:213
	// _ = "end of CoverTab[115237]"
//line /usr/local/go/src/flag/flag.go:213
}

func (i *uintValue) String() string {
//line /usr/local/go/src/flag/flag.go:215
	_go_fuzz_dep_.CoverTab[115238]++
//line /usr/local/go/src/flag/flag.go:215
	return strconv.FormatUint(uint64(*i), 10)
//line /usr/local/go/src/flag/flag.go:215
	// _ = "end of CoverTab[115238]"
//line /usr/local/go/src/flag/flag.go:215
}

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
//line /usr/local/go/src/flag/flag.go:220
	_go_fuzz_dep_.CoverTab[115239]++
						*p = val
						return (*uint64Value)(p)
//line /usr/local/go/src/flag/flag.go:222
	// _ = "end of CoverTab[115239]"
}

func (i *uint64Value) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:225
	_go_fuzz_dep_.CoverTab[115240]++
						v, err := strconv.ParseUint(s, 0, 64)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:227
		_go_fuzz_dep_.CoverTab[115242]++
							err = numError(err)
//line /usr/local/go/src/flag/flag.go:228
		// _ = "end of CoverTab[115242]"
	} else {
//line /usr/local/go/src/flag/flag.go:229
		_go_fuzz_dep_.CoverTab[115243]++
//line /usr/local/go/src/flag/flag.go:229
		// _ = "end of CoverTab[115243]"
//line /usr/local/go/src/flag/flag.go:229
	}
//line /usr/local/go/src/flag/flag.go:229
	// _ = "end of CoverTab[115240]"
//line /usr/local/go/src/flag/flag.go:229
	_go_fuzz_dep_.CoverTab[115241]++
						*i = uint64Value(v)
						return err
//line /usr/local/go/src/flag/flag.go:231
	// _ = "end of CoverTab[115241]"
}

func (i *uint64Value) Get() any {
//line /usr/local/go/src/flag/flag.go:234
	_go_fuzz_dep_.CoverTab[115244]++
//line /usr/local/go/src/flag/flag.go:234
	return uint64(*i)
//line /usr/local/go/src/flag/flag.go:234
	// _ = "end of CoverTab[115244]"
//line /usr/local/go/src/flag/flag.go:234
}

func (i *uint64Value) String() string {
//line /usr/local/go/src/flag/flag.go:236
	_go_fuzz_dep_.CoverTab[115245]++
//line /usr/local/go/src/flag/flag.go:236
	return strconv.FormatUint(uint64(*i), 10)
//line /usr/local/go/src/flag/flag.go:236
	// _ = "end of CoverTab[115245]"
//line /usr/local/go/src/flag/flag.go:236
}

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
//line /usr/local/go/src/flag/flag.go:241
	_go_fuzz_dep_.CoverTab[115246]++
						*p = val
						return (*stringValue)(p)
//line /usr/local/go/src/flag/flag.go:243
	// _ = "end of CoverTab[115246]"
}

func (s *stringValue) Set(val string) error {
//line /usr/local/go/src/flag/flag.go:246
	_go_fuzz_dep_.CoverTab[115247]++
						*s = stringValue(val)
						return nil
//line /usr/local/go/src/flag/flag.go:248
	// _ = "end of CoverTab[115247]"
}

func (s *stringValue) Get() any {
//line /usr/local/go/src/flag/flag.go:251
	_go_fuzz_dep_.CoverTab[115248]++
//line /usr/local/go/src/flag/flag.go:251
	return string(*s)
//line /usr/local/go/src/flag/flag.go:251
	// _ = "end of CoverTab[115248]"
//line /usr/local/go/src/flag/flag.go:251
}

func (s *stringValue) String() string {
//line /usr/local/go/src/flag/flag.go:253
	_go_fuzz_dep_.CoverTab[115249]++
//line /usr/local/go/src/flag/flag.go:253
	return string(*s)
//line /usr/local/go/src/flag/flag.go:253
	// _ = "end of CoverTab[115249]"
//line /usr/local/go/src/flag/flag.go:253
}

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
//line /usr/local/go/src/flag/flag.go:258
	_go_fuzz_dep_.CoverTab[115250]++
						*p = val
						return (*float64Value)(p)
//line /usr/local/go/src/flag/flag.go:260
	// _ = "end of CoverTab[115250]"
}

func (f *float64Value) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:263
	_go_fuzz_dep_.CoverTab[115251]++
						v, err := strconv.ParseFloat(s, 64)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:265
		_go_fuzz_dep_.CoverTab[115253]++
							err = numError(err)
//line /usr/local/go/src/flag/flag.go:266
		// _ = "end of CoverTab[115253]"
	} else {
//line /usr/local/go/src/flag/flag.go:267
		_go_fuzz_dep_.CoverTab[115254]++
//line /usr/local/go/src/flag/flag.go:267
		// _ = "end of CoverTab[115254]"
//line /usr/local/go/src/flag/flag.go:267
	}
//line /usr/local/go/src/flag/flag.go:267
	// _ = "end of CoverTab[115251]"
//line /usr/local/go/src/flag/flag.go:267
	_go_fuzz_dep_.CoverTab[115252]++
						*f = float64Value(v)
						return err
//line /usr/local/go/src/flag/flag.go:269
	// _ = "end of CoverTab[115252]"
}

func (f *float64Value) Get() any {
//line /usr/local/go/src/flag/flag.go:272
	_go_fuzz_dep_.CoverTab[115255]++
//line /usr/local/go/src/flag/flag.go:272
	return float64(*f)
//line /usr/local/go/src/flag/flag.go:272
	// _ = "end of CoverTab[115255]"
//line /usr/local/go/src/flag/flag.go:272
}

func (f *float64Value) String() string {
//line /usr/local/go/src/flag/flag.go:274
	_go_fuzz_dep_.CoverTab[115256]++
//line /usr/local/go/src/flag/flag.go:274
	return strconv.FormatFloat(float64(*f), 'g', -1, 64)
//line /usr/local/go/src/flag/flag.go:274
	// _ = "end of CoverTab[115256]"
//line /usr/local/go/src/flag/flag.go:274
}

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
//line /usr/local/go/src/flag/flag.go:279
	_go_fuzz_dep_.CoverTab[115257]++
						*p = val
						return (*durationValue)(p)
//line /usr/local/go/src/flag/flag.go:281
	// _ = "end of CoverTab[115257]"
}

func (d *durationValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:284
	_go_fuzz_dep_.CoverTab[115258]++
						v, err := time.ParseDuration(s)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:286
		_go_fuzz_dep_.CoverTab[115260]++
							err = errParse
//line /usr/local/go/src/flag/flag.go:287
		// _ = "end of CoverTab[115260]"
	} else {
//line /usr/local/go/src/flag/flag.go:288
		_go_fuzz_dep_.CoverTab[115261]++
//line /usr/local/go/src/flag/flag.go:288
		// _ = "end of CoverTab[115261]"
//line /usr/local/go/src/flag/flag.go:288
	}
//line /usr/local/go/src/flag/flag.go:288
	// _ = "end of CoverTab[115258]"
//line /usr/local/go/src/flag/flag.go:288
	_go_fuzz_dep_.CoverTab[115259]++
						*d = durationValue(v)
						return err
//line /usr/local/go/src/flag/flag.go:290
	// _ = "end of CoverTab[115259]"
}

func (d *durationValue) Get() any {
//line /usr/local/go/src/flag/flag.go:293
	_go_fuzz_dep_.CoverTab[115262]++
//line /usr/local/go/src/flag/flag.go:293
	return time.Duration(*d)
//line /usr/local/go/src/flag/flag.go:293
	// _ = "end of CoverTab[115262]"
//line /usr/local/go/src/flag/flag.go:293
}

func (d *durationValue) String() string {
//line /usr/local/go/src/flag/flag.go:295
	_go_fuzz_dep_.CoverTab[115263]++
//line /usr/local/go/src/flag/flag.go:295
	return (*time.Duration)(d).String()
//line /usr/local/go/src/flag/flag.go:295
	// _ = "end of CoverTab[115263]"
//line /usr/local/go/src/flag/flag.go:295
}

// -- encoding.TextUnmarshaler Value
type textValue struct{ p encoding.TextUnmarshaler }

func newTextValue(val encoding.TextMarshaler, p encoding.TextUnmarshaler) textValue {
//line /usr/local/go/src/flag/flag.go:300
	_go_fuzz_dep_.CoverTab[115264]++
						ptrVal := reflect.ValueOf(p)
						if ptrVal.Kind() != reflect.Ptr {
//line /usr/local/go/src/flag/flag.go:302
		_go_fuzz_dep_.CoverTab[115268]++
							panic("variable value type must be a pointer")
//line /usr/local/go/src/flag/flag.go:303
		// _ = "end of CoverTab[115268]"
	} else {
//line /usr/local/go/src/flag/flag.go:304
		_go_fuzz_dep_.CoverTab[115269]++
//line /usr/local/go/src/flag/flag.go:304
		// _ = "end of CoverTab[115269]"
//line /usr/local/go/src/flag/flag.go:304
	}
//line /usr/local/go/src/flag/flag.go:304
	// _ = "end of CoverTab[115264]"
//line /usr/local/go/src/flag/flag.go:304
	_go_fuzz_dep_.CoverTab[115265]++
						defVal := reflect.ValueOf(val)
						if defVal.Kind() == reflect.Ptr {
//line /usr/local/go/src/flag/flag.go:306
		_go_fuzz_dep_.CoverTab[115270]++
							defVal = defVal.Elem()
//line /usr/local/go/src/flag/flag.go:307
		// _ = "end of CoverTab[115270]"
	} else {
//line /usr/local/go/src/flag/flag.go:308
		_go_fuzz_dep_.CoverTab[115271]++
//line /usr/local/go/src/flag/flag.go:308
		// _ = "end of CoverTab[115271]"
//line /usr/local/go/src/flag/flag.go:308
	}
//line /usr/local/go/src/flag/flag.go:308
	// _ = "end of CoverTab[115265]"
//line /usr/local/go/src/flag/flag.go:308
	_go_fuzz_dep_.CoverTab[115266]++
						if defVal.Type() != ptrVal.Type().Elem() {
//line /usr/local/go/src/flag/flag.go:309
		_go_fuzz_dep_.CoverTab[115272]++
							panic(fmt.Sprintf("default type does not match variable type: %v != %v", defVal.Type(), ptrVal.Type().Elem()))
//line /usr/local/go/src/flag/flag.go:310
		// _ = "end of CoverTab[115272]"
	} else {
//line /usr/local/go/src/flag/flag.go:311
		_go_fuzz_dep_.CoverTab[115273]++
//line /usr/local/go/src/flag/flag.go:311
		// _ = "end of CoverTab[115273]"
//line /usr/local/go/src/flag/flag.go:311
	}
//line /usr/local/go/src/flag/flag.go:311
	// _ = "end of CoverTab[115266]"
//line /usr/local/go/src/flag/flag.go:311
	_go_fuzz_dep_.CoverTab[115267]++
						ptrVal.Elem().Set(defVal)
						return textValue{p}
//line /usr/local/go/src/flag/flag.go:313
	// _ = "end of CoverTab[115267]"
}

func (v textValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:316
	_go_fuzz_dep_.CoverTab[115274]++
						return v.p.UnmarshalText([]byte(s))
//line /usr/local/go/src/flag/flag.go:317
	// _ = "end of CoverTab[115274]"
}

func (v textValue) Get() interface{} {
//line /usr/local/go/src/flag/flag.go:320
	_go_fuzz_dep_.CoverTab[115275]++
						return v.p
//line /usr/local/go/src/flag/flag.go:321
	// _ = "end of CoverTab[115275]"
}

func (v textValue) String() string {
//line /usr/local/go/src/flag/flag.go:324
	_go_fuzz_dep_.CoverTab[115276]++
						if m, ok := v.p.(encoding.TextMarshaler); ok {
//line /usr/local/go/src/flag/flag.go:325
		_go_fuzz_dep_.CoverTab[115278]++
							if b, err := m.MarshalText(); err == nil {
//line /usr/local/go/src/flag/flag.go:326
			_go_fuzz_dep_.CoverTab[115279]++
								return string(b)
//line /usr/local/go/src/flag/flag.go:327
			// _ = "end of CoverTab[115279]"
		} else {
//line /usr/local/go/src/flag/flag.go:328
			_go_fuzz_dep_.CoverTab[115280]++
//line /usr/local/go/src/flag/flag.go:328
			// _ = "end of CoverTab[115280]"
//line /usr/local/go/src/flag/flag.go:328
		}
//line /usr/local/go/src/flag/flag.go:328
		// _ = "end of CoverTab[115278]"
	} else {
//line /usr/local/go/src/flag/flag.go:329
		_go_fuzz_dep_.CoverTab[115281]++
//line /usr/local/go/src/flag/flag.go:329
		// _ = "end of CoverTab[115281]"
//line /usr/local/go/src/flag/flag.go:329
	}
//line /usr/local/go/src/flag/flag.go:329
	// _ = "end of CoverTab[115276]"
//line /usr/local/go/src/flag/flag.go:329
	_go_fuzz_dep_.CoverTab[115277]++
						return ""
//line /usr/local/go/src/flag/flag.go:330
	// _ = "end of CoverTab[115277]"
}

// -- func Value
type funcValue func(string) error

func (f funcValue) Set(s string) error {
//line /usr/local/go/src/flag/flag.go:336
	_go_fuzz_dep_.CoverTab[115282]++
//line /usr/local/go/src/flag/flag.go:336
	return f(s)
//line /usr/local/go/src/flag/flag.go:336
	// _ = "end of CoverTab[115282]"
//line /usr/local/go/src/flag/flag.go:336
}

func (f funcValue) String() string {
//line /usr/local/go/src/flag/flag.go:338
	_go_fuzz_dep_.CoverTab[115283]++
//line /usr/local/go/src/flag/flag.go:338
	return ""
//line /usr/local/go/src/flag/flag.go:338
	// _ = "end of CoverTab[115283]"
//line /usr/local/go/src/flag/flag.go:338
}

// Value is the interface to the dynamic value stored in a flag.
//line /usr/local/go/src/flag/flag.go:340
// (The default value is represented as a string.)
//line /usr/local/go/src/flag/flag.go:340
//
//line /usr/local/go/src/flag/flag.go:340
// If a Value has an IsBoolFlag() bool method returning true,
//line /usr/local/go/src/flag/flag.go:340
// the command-line parser makes -name equivalent to -name=true
//line /usr/local/go/src/flag/flag.go:340
// rather than using the next command-line argument.
//line /usr/local/go/src/flag/flag.go:340
//
//line /usr/local/go/src/flag/flag.go:340
// Set is called once, in command line order, for each flag present.
//line /usr/local/go/src/flag/flag.go:340
// The flag package may call the String method with a zero-valued receiver,
//line /usr/local/go/src/flag/flag.go:340
// such as a nil pointer.
//line /usr/local/go/src/flag/flag.go:350
type Value interface {
	String() string
	Set(string) error
}

// Getter is an interface that allows the contents of a Value to be retrieved.
//line /usr/local/go/src/flag/flag.go:355
// It wraps the Value interface, rather than being part of it, because it
//line /usr/local/go/src/flag/flag.go:355
// appeared after Go 1 and its compatibility rules. All Value types provided
//line /usr/local/go/src/flag/flag.go:355
// by this package satisfy the Getter interface, except the type used by Func.
//line /usr/local/go/src/flag/flag.go:359
type Getter interface {
	Value
	Get() any
}

// ErrorHandling defines how FlagSet.Parse behaves if the parse fails.
type ErrorHandling int

// These constants cause FlagSet.Parse to behave as described if the parse fails.
const (
	ContinueOnError	ErrorHandling	= iota	// Return a descriptive error.
	ExitOnError				// Call os.Exit(2) or for -h/-help Exit(0).
	PanicOnError				// Call panic with a descriptive error.
)

// A FlagSet represents a set of defined flags. The zero value of a FlagSet
//line /usr/local/go/src/flag/flag.go:374
// has no name and has ContinueOnError error handling.
//line /usr/local/go/src/flag/flag.go:374
//
//line /usr/local/go/src/flag/flag.go:374
// Flag names must be unique within a FlagSet. An attempt to define a flag whose
//line /usr/local/go/src/flag/flag.go:374
// name is already in use will cause a panic.
//line /usr/local/go/src/flag/flag.go:379
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage	func()

	name		string
	parsed		bool
	actual		map[string]*Flag
	formal		map[string]*Flag
	args		[]string	// arguments after flags
	errorHandling	ErrorHandling
	output		io.Writer	// nil means stderr; use Output() accessor
}

// A Flag represents the state of a flag.
type Flag struct {
	Name		string	// name as it appears on command line
	Usage		string	// help message
	Value		Value	// value as set
	DefValue	string	// default value (as text); for usage message
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*Flag) []*Flag {
//line /usr/local/go/src/flag/flag.go:405
	_go_fuzz_dep_.CoverTab[115284]++
						result := make([]*Flag, len(flags))
						i := 0
						for _, f := range flags {
//line /usr/local/go/src/flag/flag.go:408
		_go_fuzz_dep_.CoverTab[115287]++
							result[i] = f
							i++
//line /usr/local/go/src/flag/flag.go:410
		// _ = "end of CoverTab[115287]"
	}
//line /usr/local/go/src/flag/flag.go:411
	// _ = "end of CoverTab[115284]"
//line /usr/local/go/src/flag/flag.go:411
	_go_fuzz_dep_.CoverTab[115285]++
						sort.Slice(result, func(i, j int) bool {
//line /usr/local/go/src/flag/flag.go:412
		_go_fuzz_dep_.CoverTab[115288]++
							return result[i].Name < result[j].Name
//line /usr/local/go/src/flag/flag.go:413
		// _ = "end of CoverTab[115288]"
	})
//line /usr/local/go/src/flag/flag.go:414
	// _ = "end of CoverTab[115285]"
//line /usr/local/go/src/flag/flag.go:414
	_go_fuzz_dep_.CoverTab[115286]++
						return result
//line /usr/local/go/src/flag/flag.go:415
	// _ = "end of CoverTab[115286]"
}

// Output returns the destination for usage and error messages. os.Stderr is returned if
//line /usr/local/go/src/flag/flag.go:418
// output was not set or was set to nil.
//line /usr/local/go/src/flag/flag.go:420
func (f *FlagSet) Output() io.Writer {
//line /usr/local/go/src/flag/flag.go:420
	_go_fuzz_dep_.CoverTab[115289]++
						if f.output == nil {
//line /usr/local/go/src/flag/flag.go:421
		_go_fuzz_dep_.CoverTab[115291]++
							return os.Stderr
//line /usr/local/go/src/flag/flag.go:422
		// _ = "end of CoverTab[115291]"
	} else {
//line /usr/local/go/src/flag/flag.go:423
		_go_fuzz_dep_.CoverTab[115292]++
//line /usr/local/go/src/flag/flag.go:423
		// _ = "end of CoverTab[115292]"
//line /usr/local/go/src/flag/flag.go:423
	}
//line /usr/local/go/src/flag/flag.go:423
	// _ = "end of CoverTab[115289]"
//line /usr/local/go/src/flag/flag.go:423
	_go_fuzz_dep_.CoverTab[115290]++
						return f.output
//line /usr/local/go/src/flag/flag.go:424
	// _ = "end of CoverTab[115290]"
}

// Name returns the name of the flag set.
func (f *FlagSet) Name() string {
//line /usr/local/go/src/flag/flag.go:428
	_go_fuzz_dep_.CoverTab[115293]++
						return f.name
//line /usr/local/go/src/flag/flag.go:429
	// _ = "end of CoverTab[115293]"
}

// ErrorHandling returns the error handling behavior of the flag set.
func (f *FlagSet) ErrorHandling() ErrorHandling {
//line /usr/local/go/src/flag/flag.go:433
	_go_fuzz_dep_.CoverTab[115294]++
						return f.errorHandling
//line /usr/local/go/src/flag/flag.go:434
	// _ = "end of CoverTab[115294]"
}

// SetOutput sets the destination for usage and error messages.
//line /usr/local/go/src/flag/flag.go:437
// If output is nil, os.Stderr is used.
//line /usr/local/go/src/flag/flag.go:439
func (f *FlagSet) SetOutput(output io.Writer) {
//line /usr/local/go/src/flag/flag.go:439
	_go_fuzz_dep_.CoverTab[115295]++
						f.output = output
//line /usr/local/go/src/flag/flag.go:440
	// _ = "end of CoverTab[115295]"
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
//line /usr/local/go/src/flag/flag.go:443
// It visits all flags, even those not set.
//line /usr/local/go/src/flag/flag.go:445
func (f *FlagSet) VisitAll(fn func(*Flag)) {
//line /usr/local/go/src/flag/flag.go:445
	_go_fuzz_dep_.CoverTab[115296]++
						for _, flag := range sortFlags(f.formal) {
//line /usr/local/go/src/flag/flag.go:446
		_go_fuzz_dep_.CoverTab[115297]++
							fn(flag)
//line /usr/local/go/src/flag/flag.go:447
		// _ = "end of CoverTab[115297]"
	}
//line /usr/local/go/src/flag/flag.go:448
	// _ = "end of CoverTab[115296]"
}

// VisitAll visits the command-line flags in lexicographical order, calling
//line /usr/local/go/src/flag/flag.go:451
// fn for each. It visits all flags, even those not set.
//line /usr/local/go/src/flag/flag.go:453
func VisitAll(fn func(*Flag)) {
//line /usr/local/go/src/flag/flag.go:453
	_go_fuzz_dep_.CoverTab[115298]++
						CommandLine.VisitAll(fn)
//line /usr/local/go/src/flag/flag.go:454
	// _ = "end of CoverTab[115298]"
}

// Visit visits the flags in lexicographical order, calling fn for each.
//line /usr/local/go/src/flag/flag.go:457
// It visits only those flags that have been set.
//line /usr/local/go/src/flag/flag.go:459
func (f *FlagSet) Visit(fn func(*Flag)) {
//line /usr/local/go/src/flag/flag.go:459
	_go_fuzz_dep_.CoverTab[115299]++
						for _, flag := range sortFlags(f.actual) {
//line /usr/local/go/src/flag/flag.go:460
		_go_fuzz_dep_.CoverTab[115300]++
							fn(flag)
//line /usr/local/go/src/flag/flag.go:461
		// _ = "end of CoverTab[115300]"
	}
//line /usr/local/go/src/flag/flag.go:462
	// _ = "end of CoverTab[115299]"
}

// Visit visits the command-line flags in lexicographical order, calling fn
//line /usr/local/go/src/flag/flag.go:465
// for each. It visits only those flags that have been set.
//line /usr/local/go/src/flag/flag.go:467
func Visit(fn func(*Flag)) {
//line /usr/local/go/src/flag/flag.go:467
	_go_fuzz_dep_.CoverTab[115301]++
						CommandLine.Visit(fn)
//line /usr/local/go/src/flag/flag.go:468
	// _ = "end of CoverTab[115301]"
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
//line /usr/local/go/src/flag/flag.go:472
	_go_fuzz_dep_.CoverTab[115302]++
						return f.formal[name]
//line /usr/local/go/src/flag/flag.go:473
	// _ = "end of CoverTab[115302]"
}

// Lookup returns the Flag structure of the named command-line flag,
//line /usr/local/go/src/flag/flag.go:476
// returning nil if none exists.
//line /usr/local/go/src/flag/flag.go:478
func Lookup(name string) *Flag {
//line /usr/local/go/src/flag/flag.go:478
	_go_fuzz_dep_.CoverTab[115303]++
						return CommandLine.formal[name]
//line /usr/local/go/src/flag/flag.go:479
	// _ = "end of CoverTab[115303]"
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
//line /usr/local/go/src/flag/flag.go:483
	_go_fuzz_dep_.CoverTab[115304]++
						flag, ok := f.formal[name]
						if !ok {
//line /usr/local/go/src/flag/flag.go:485
		_go_fuzz_dep_.CoverTab[115308]++
							return fmt.Errorf("no such flag -%v", name)
//line /usr/local/go/src/flag/flag.go:486
		// _ = "end of CoverTab[115308]"
	} else {
//line /usr/local/go/src/flag/flag.go:487
		_go_fuzz_dep_.CoverTab[115309]++
//line /usr/local/go/src/flag/flag.go:487
		// _ = "end of CoverTab[115309]"
//line /usr/local/go/src/flag/flag.go:487
	}
//line /usr/local/go/src/flag/flag.go:487
	// _ = "end of CoverTab[115304]"
//line /usr/local/go/src/flag/flag.go:487
	_go_fuzz_dep_.CoverTab[115305]++
						err := flag.Value.Set(value)
						if err != nil {
//line /usr/local/go/src/flag/flag.go:489
		_go_fuzz_dep_.CoverTab[115310]++
							return err
//line /usr/local/go/src/flag/flag.go:490
		// _ = "end of CoverTab[115310]"
	} else {
//line /usr/local/go/src/flag/flag.go:491
		_go_fuzz_dep_.CoverTab[115311]++
//line /usr/local/go/src/flag/flag.go:491
		// _ = "end of CoverTab[115311]"
//line /usr/local/go/src/flag/flag.go:491
	}
//line /usr/local/go/src/flag/flag.go:491
	// _ = "end of CoverTab[115305]"
//line /usr/local/go/src/flag/flag.go:491
	_go_fuzz_dep_.CoverTab[115306]++
						if f.actual == nil {
//line /usr/local/go/src/flag/flag.go:492
		_go_fuzz_dep_.CoverTab[115312]++
							f.actual = make(map[string]*Flag)
//line /usr/local/go/src/flag/flag.go:493
		// _ = "end of CoverTab[115312]"
	} else {
//line /usr/local/go/src/flag/flag.go:494
		_go_fuzz_dep_.CoverTab[115313]++
//line /usr/local/go/src/flag/flag.go:494
		// _ = "end of CoverTab[115313]"
//line /usr/local/go/src/flag/flag.go:494
	}
//line /usr/local/go/src/flag/flag.go:494
	// _ = "end of CoverTab[115306]"
//line /usr/local/go/src/flag/flag.go:494
	_go_fuzz_dep_.CoverTab[115307]++
						f.actual[name] = flag
						return nil
//line /usr/local/go/src/flag/flag.go:496
	// _ = "end of CoverTab[115307]"
}

// Set sets the value of the named command-line flag.
func Set(name, value string) error {
//line /usr/local/go/src/flag/flag.go:500
	_go_fuzz_dep_.CoverTab[115314]++
						return CommandLine.Set(name, value)
//line /usr/local/go/src/flag/flag.go:501
	// _ = "end of CoverTab[115314]"
}

// isZeroValue determines whether the string represents the zero
//line /usr/local/go/src/flag/flag.go:504
// value for a flag.
//line /usr/local/go/src/flag/flag.go:506
func isZeroValue(flag *Flag, value string) (ok bool, err error) {
//line /usr/local/go/src/flag/flag.go:506
	_go_fuzz_dep_.CoverTab[115315]++

//line /usr/local/go/src/flag/flag.go:510
	typ := reflect.TypeOf(flag.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Pointer {
//line /usr/local/go/src/flag/flag.go:512
		_go_fuzz_dep_.CoverTab[115318]++
							z = reflect.New(typ.Elem())
//line /usr/local/go/src/flag/flag.go:513
		// _ = "end of CoverTab[115318]"
	} else {
//line /usr/local/go/src/flag/flag.go:514
		_go_fuzz_dep_.CoverTab[115319]++
							z = reflect.Zero(typ)
//line /usr/local/go/src/flag/flag.go:515
		// _ = "end of CoverTab[115319]"
	}
//line /usr/local/go/src/flag/flag.go:516
	// _ = "end of CoverTab[115315]"
//line /usr/local/go/src/flag/flag.go:516
	_go_fuzz_dep_.CoverTab[115316]++

//line /usr/local/go/src/flag/flag.go:520
	defer func() {
//line /usr/local/go/src/flag/flag.go:520
		_go_fuzz_dep_.CoverTab[115320]++
							if e := recover(); e != nil {
//line /usr/local/go/src/flag/flag.go:521
			_go_fuzz_dep_.CoverTab[115321]++
								if typ.Kind() == reflect.Pointer {
//line /usr/local/go/src/flag/flag.go:522
				_go_fuzz_dep_.CoverTab[115323]++
									typ = typ.Elem()
//line /usr/local/go/src/flag/flag.go:523
				// _ = "end of CoverTab[115323]"
			} else {
//line /usr/local/go/src/flag/flag.go:524
				_go_fuzz_dep_.CoverTab[115324]++
//line /usr/local/go/src/flag/flag.go:524
				// _ = "end of CoverTab[115324]"
//line /usr/local/go/src/flag/flag.go:524
			}
//line /usr/local/go/src/flag/flag.go:524
			// _ = "end of CoverTab[115321]"
//line /usr/local/go/src/flag/flag.go:524
			_go_fuzz_dep_.CoverTab[115322]++
								err = fmt.Errorf("panic calling String method on zero %v for flag %s: %v", typ, flag.Name, e)
//line /usr/local/go/src/flag/flag.go:525
			// _ = "end of CoverTab[115322]"
		} else {
//line /usr/local/go/src/flag/flag.go:526
			_go_fuzz_dep_.CoverTab[115325]++
//line /usr/local/go/src/flag/flag.go:526
			// _ = "end of CoverTab[115325]"
//line /usr/local/go/src/flag/flag.go:526
		}
//line /usr/local/go/src/flag/flag.go:526
		// _ = "end of CoverTab[115320]"
	}()
//line /usr/local/go/src/flag/flag.go:527
	// _ = "end of CoverTab[115316]"
//line /usr/local/go/src/flag/flag.go:527
	_go_fuzz_dep_.CoverTab[115317]++
						return value == z.Interface().(Value).String(), nil
//line /usr/local/go/src/flag/flag.go:528
	// _ = "end of CoverTab[115317]"
}

// UnquoteUsage extracts a back-quoted name from the usage
//line /usr/local/go/src/flag/flag.go:531
// string for a flag and returns it and the un-quoted usage.
//line /usr/local/go/src/flag/flag.go:531
// Given "a `name` to show" it returns ("name", "a name to show").
//line /usr/local/go/src/flag/flag.go:531
// If there are no back quotes, the name is an educated guess of the
//line /usr/local/go/src/flag/flag.go:531
// type of the flag's value, or the empty string if the flag is boolean.
//line /usr/local/go/src/flag/flag.go:536
func UnquoteUsage(flag *Flag) (name string, usage string) {
//line /usr/local/go/src/flag/flag.go:536
	_go_fuzz_dep_.CoverTab[115326]++

						usage = flag.Usage
						for i := 0; i < len(usage); i++ {
//line /usr/local/go/src/flag/flag.go:539
		_go_fuzz_dep_.CoverTab[115329]++
							if usage[i] == '`' {
//line /usr/local/go/src/flag/flag.go:540
			_go_fuzz_dep_.CoverTab[115330]++
								for j := i + 1; j < len(usage); j++ {
//line /usr/local/go/src/flag/flag.go:541
				_go_fuzz_dep_.CoverTab[115332]++
									if usage[j] == '`' {
//line /usr/local/go/src/flag/flag.go:542
					_go_fuzz_dep_.CoverTab[115333]++
										name = usage[i+1 : j]
										usage = usage[:i] + name + usage[j+1:]
										return name, usage
//line /usr/local/go/src/flag/flag.go:545
					// _ = "end of CoverTab[115333]"
				} else {
//line /usr/local/go/src/flag/flag.go:546
					_go_fuzz_dep_.CoverTab[115334]++
//line /usr/local/go/src/flag/flag.go:546
					// _ = "end of CoverTab[115334]"
//line /usr/local/go/src/flag/flag.go:546
				}
//line /usr/local/go/src/flag/flag.go:546
				// _ = "end of CoverTab[115332]"
			}
//line /usr/local/go/src/flag/flag.go:547
			// _ = "end of CoverTab[115330]"
//line /usr/local/go/src/flag/flag.go:547
			_go_fuzz_dep_.CoverTab[115331]++
								break
//line /usr/local/go/src/flag/flag.go:548
			// _ = "end of CoverTab[115331]"
		} else {
//line /usr/local/go/src/flag/flag.go:549
			_go_fuzz_dep_.CoverTab[115335]++
//line /usr/local/go/src/flag/flag.go:549
			// _ = "end of CoverTab[115335]"
//line /usr/local/go/src/flag/flag.go:549
		}
//line /usr/local/go/src/flag/flag.go:549
		// _ = "end of CoverTab[115329]"
	}
//line /usr/local/go/src/flag/flag.go:550
	// _ = "end of CoverTab[115326]"
//line /usr/local/go/src/flag/flag.go:550
	_go_fuzz_dep_.CoverTab[115327]++

						name = "value"
						switch fv := flag.Value.(type) {
	case boolFlag:
//line /usr/local/go/src/flag/flag.go:554
		_go_fuzz_dep_.CoverTab[115336]++
							if fv.IsBoolFlag() {
//line /usr/local/go/src/flag/flag.go:555
			_go_fuzz_dep_.CoverTab[115342]++
								name = ""
//line /usr/local/go/src/flag/flag.go:556
			// _ = "end of CoverTab[115342]"
		} else {
//line /usr/local/go/src/flag/flag.go:557
			_go_fuzz_dep_.CoverTab[115343]++
//line /usr/local/go/src/flag/flag.go:557
			// _ = "end of CoverTab[115343]"
//line /usr/local/go/src/flag/flag.go:557
		}
//line /usr/local/go/src/flag/flag.go:557
		// _ = "end of CoverTab[115336]"
	case *durationValue:
//line /usr/local/go/src/flag/flag.go:558
		_go_fuzz_dep_.CoverTab[115337]++
							name = "duration"
//line /usr/local/go/src/flag/flag.go:559
		// _ = "end of CoverTab[115337]"
	case *float64Value:
//line /usr/local/go/src/flag/flag.go:560
		_go_fuzz_dep_.CoverTab[115338]++
							name = "float"
//line /usr/local/go/src/flag/flag.go:561
		// _ = "end of CoverTab[115338]"
	case *intValue, *int64Value:
//line /usr/local/go/src/flag/flag.go:562
		_go_fuzz_dep_.CoverTab[115339]++
							name = "int"
//line /usr/local/go/src/flag/flag.go:563
		// _ = "end of CoverTab[115339]"
	case *stringValue:
//line /usr/local/go/src/flag/flag.go:564
		_go_fuzz_dep_.CoverTab[115340]++
							name = "string"
//line /usr/local/go/src/flag/flag.go:565
		// _ = "end of CoverTab[115340]"
	case *uintValue, *uint64Value:
//line /usr/local/go/src/flag/flag.go:566
		_go_fuzz_dep_.CoverTab[115341]++
							name = "uint"
//line /usr/local/go/src/flag/flag.go:567
		// _ = "end of CoverTab[115341]"
	}
//line /usr/local/go/src/flag/flag.go:568
	// _ = "end of CoverTab[115327]"
//line /usr/local/go/src/flag/flag.go:568
	_go_fuzz_dep_.CoverTab[115328]++
						return
//line /usr/local/go/src/flag/flag.go:569
	// _ = "end of CoverTab[115328]"
}

// PrintDefaults prints, to standard error unless configured otherwise, the
//line /usr/local/go/src/flag/flag.go:572
// default values of all defined command-line flags in the set. See the
//line /usr/local/go/src/flag/flag.go:572
// documentation for the global function PrintDefaults for more information.
//line /usr/local/go/src/flag/flag.go:575
func (f *FlagSet) PrintDefaults() {
//line /usr/local/go/src/flag/flag.go:575
	_go_fuzz_dep_.CoverTab[115344]++
						var isZeroValueErrs []error
						f.VisitAll(func(flag *Flag) {
//line /usr/local/go/src/flag/flag.go:577
		_go_fuzz_dep_.CoverTab[115346]++
							var b strings.Builder
							fmt.Fprintf(&b, "  -%s", flag.Name)
							name, usage := UnquoteUsage(flag)
							if len(name) > 0 {
//line /usr/local/go/src/flag/flag.go:581
			_go_fuzz_dep_.CoverTab[115350]++
								b.WriteString(" ")
								b.WriteString(name)
//line /usr/local/go/src/flag/flag.go:583
			// _ = "end of CoverTab[115350]"
		} else {
//line /usr/local/go/src/flag/flag.go:584
			_go_fuzz_dep_.CoverTab[115351]++
//line /usr/local/go/src/flag/flag.go:584
			// _ = "end of CoverTab[115351]"
//line /usr/local/go/src/flag/flag.go:584
		}
//line /usr/local/go/src/flag/flag.go:584
		// _ = "end of CoverTab[115346]"
//line /usr/local/go/src/flag/flag.go:584
		_go_fuzz_dep_.CoverTab[115347]++

//line /usr/local/go/src/flag/flag.go:587
		if b.Len() <= 4 {
//line /usr/local/go/src/flag/flag.go:587
			_go_fuzz_dep_.CoverTab[115352]++
								b.WriteString("\t")
//line /usr/local/go/src/flag/flag.go:588
			// _ = "end of CoverTab[115352]"
		} else {
//line /usr/local/go/src/flag/flag.go:589
			_go_fuzz_dep_.CoverTab[115353]++

//line /usr/local/go/src/flag/flag.go:592
			b.WriteString("\n    \t")
//line /usr/local/go/src/flag/flag.go:592
			// _ = "end of CoverTab[115353]"
		}
//line /usr/local/go/src/flag/flag.go:593
		// _ = "end of CoverTab[115347]"
//line /usr/local/go/src/flag/flag.go:593
		_go_fuzz_dep_.CoverTab[115348]++
							b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

//line /usr/local/go/src/flag/flag.go:598
		if isZero, err := isZeroValue(flag, flag.DefValue); err != nil {
//line /usr/local/go/src/flag/flag.go:598
			_go_fuzz_dep_.CoverTab[115354]++
								isZeroValueErrs = append(isZeroValueErrs, err)
//line /usr/local/go/src/flag/flag.go:599
			// _ = "end of CoverTab[115354]"
		} else {
//line /usr/local/go/src/flag/flag.go:600
			_go_fuzz_dep_.CoverTab[115355]++
//line /usr/local/go/src/flag/flag.go:600
			if !isZero {
//line /usr/local/go/src/flag/flag.go:600
				_go_fuzz_dep_.CoverTab[115356]++
									if _, ok := flag.Value.(*stringValue); ok {
//line /usr/local/go/src/flag/flag.go:601
					_go_fuzz_dep_.CoverTab[115357]++

										fmt.Fprintf(&b, " (default %q)", flag.DefValue)
//line /usr/local/go/src/flag/flag.go:603
					// _ = "end of CoverTab[115357]"
				} else {
//line /usr/local/go/src/flag/flag.go:604
					_go_fuzz_dep_.CoverTab[115358]++
										fmt.Fprintf(&b, " (default %v)", flag.DefValue)
//line /usr/local/go/src/flag/flag.go:605
					// _ = "end of CoverTab[115358]"
				}
//line /usr/local/go/src/flag/flag.go:606
				// _ = "end of CoverTab[115356]"
			} else {
//line /usr/local/go/src/flag/flag.go:607
				_go_fuzz_dep_.CoverTab[115359]++
//line /usr/local/go/src/flag/flag.go:607
				// _ = "end of CoverTab[115359]"
//line /usr/local/go/src/flag/flag.go:607
			}
//line /usr/local/go/src/flag/flag.go:607
			// _ = "end of CoverTab[115355]"
//line /usr/local/go/src/flag/flag.go:607
		}
//line /usr/local/go/src/flag/flag.go:607
		// _ = "end of CoverTab[115348]"
//line /usr/local/go/src/flag/flag.go:607
		_go_fuzz_dep_.CoverTab[115349]++
							fmt.Fprint(f.Output(), b.String(), "\n")
//line /usr/local/go/src/flag/flag.go:608
		// _ = "end of CoverTab[115349]"
	})
//line /usr/local/go/src/flag/flag.go:609
	// _ = "end of CoverTab[115344]"
//line /usr/local/go/src/flag/flag.go:609
	_go_fuzz_dep_.CoverTab[115345]++

//line /usr/local/go/src/flag/flag.go:613
	if errs := isZeroValueErrs; len(errs) > 0 {
//line /usr/local/go/src/flag/flag.go:613
		_go_fuzz_dep_.CoverTab[115360]++
							fmt.Fprintln(f.Output())
							for _, err := range errs {
//line /usr/local/go/src/flag/flag.go:615
			_go_fuzz_dep_.CoverTab[115361]++
								fmt.Fprintln(f.Output(), err)
//line /usr/local/go/src/flag/flag.go:616
			// _ = "end of CoverTab[115361]"
		}
//line /usr/local/go/src/flag/flag.go:617
		// _ = "end of CoverTab[115360]"
	} else {
//line /usr/local/go/src/flag/flag.go:618
		_go_fuzz_dep_.CoverTab[115362]++
//line /usr/local/go/src/flag/flag.go:618
		// _ = "end of CoverTab[115362]"
//line /usr/local/go/src/flag/flag.go:618
	}
//line /usr/local/go/src/flag/flag.go:618
	// _ = "end of CoverTab[115345]"
}

// PrintDefaults prints, to standard error unless configured otherwise,
//line /usr/local/go/src/flag/flag.go:621
// a usage message showing the default settings of all defined
//line /usr/local/go/src/flag/flag.go:621
// command-line flags.
//line /usr/local/go/src/flag/flag.go:621
// For an integer valued flag x, the default output has the form
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
//	-x int
//line /usr/local/go/src/flag/flag.go:621
//		usage-message-for-x (default 7)
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
// The usage message will appear on a separate line for anything but
//line /usr/local/go/src/flag/flag.go:621
// a bool flag with a one-byte name. For bool flags, the type is
//line /usr/local/go/src/flag/flag.go:621
// omitted and if the flag name is one byte the usage message appears
//line /usr/local/go/src/flag/flag.go:621
// on the same line. The parenthetical default is omitted if the
//line /usr/local/go/src/flag/flag.go:621
// default is the zero value for the type. The listed type, here int,
//line /usr/local/go/src/flag/flag.go:621
// can be changed by placing a back-quoted name in the flag's usage
//line /usr/local/go/src/flag/flag.go:621
// string; the first such item in the message is taken to be a parameter
//line /usr/local/go/src/flag/flag.go:621
// name to show in the message and the back quotes are stripped from
//line /usr/local/go/src/flag/flag.go:621
// the message when displayed. For instance, given
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
//	flag.String("I", "", "search `directory` for include files")
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
// the output will be
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
//	-I directory
//line /usr/local/go/src/flag/flag.go:621
//		search directory for include files.
//line /usr/local/go/src/flag/flag.go:621
//
//line /usr/local/go/src/flag/flag.go:621
// To change the destination for flag messages, call CommandLine.SetOutput.
//line /usr/local/go/src/flag/flag.go:647
func PrintDefaults() {
//line /usr/local/go/src/flag/flag.go:647
	_go_fuzz_dep_.CoverTab[115363]++
						CommandLine.PrintDefaults()
//line /usr/local/go/src/flag/flag.go:648
	// _ = "end of CoverTab[115363]"
}

// defaultUsage is the default function to print a usage message.
func (f *FlagSet) defaultUsage() {
//line /usr/local/go/src/flag/flag.go:652
	_go_fuzz_dep_.CoverTab[115364]++
						if f.name == "" {
//line /usr/local/go/src/flag/flag.go:653
		_go_fuzz_dep_.CoverTab[115366]++
							fmt.Fprintf(f.Output(), "Usage:\n")
//line /usr/local/go/src/flag/flag.go:654
		// _ = "end of CoverTab[115366]"
	} else {
//line /usr/local/go/src/flag/flag.go:655
		_go_fuzz_dep_.CoverTab[115367]++
							fmt.Fprintf(f.Output(), "Usage of %s:\n", f.name)
//line /usr/local/go/src/flag/flag.go:656
		// _ = "end of CoverTab[115367]"
	}
//line /usr/local/go/src/flag/flag.go:657
	// _ = "end of CoverTab[115364]"
//line /usr/local/go/src/flag/flag.go:657
	_go_fuzz_dep_.CoverTab[115365]++
						f.PrintDefaults()
//line /usr/local/go/src/flag/flag.go:658
	// _ = "end of CoverTab[115365]"
}

//line /usr/local/go/src/flag/flag.go:665
// Usage prints a usage message documenting all defined command-line flags
//line /usr/local/go/src/flag/flag.go:665
// to CommandLine's output, which by default is os.Stderr.
//line /usr/local/go/src/flag/flag.go:665
// It is called when an error occurs while parsing flags.
//line /usr/local/go/src/flag/flag.go:665
// The function is a variable that may be changed to point to a custom function.
//line /usr/local/go/src/flag/flag.go:665
// By default it prints a simple header and calls PrintDefaults; for details about the
//line /usr/local/go/src/flag/flag.go:665
// format of the output and how to control it, see the documentation for PrintDefaults.
//line /usr/local/go/src/flag/flag.go:665
// Custom usage functions may choose to exit the program; by default exiting
//line /usr/local/go/src/flag/flag.go:665
// happens anyway as the command line's error handling strategy is set to
//line /usr/local/go/src/flag/flag.go:665
// ExitOnError.
//line /usr/local/go/src/flag/flag.go:674
var Usage = func() {
//line /usr/local/go/src/flag/flag.go:674
	_go_fuzz_dep_.CoverTab[115368]++
						fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
						PrintDefaults()
//line /usr/local/go/src/flag/flag.go:676
	// _ = "end of CoverTab[115368]"
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int {
//line /usr/local/go/src/flag/flag.go:680
	_go_fuzz_dep_.CoverTab[115369]++
//line /usr/local/go/src/flag/flag.go:680
	return len(f.actual)
//line /usr/local/go/src/flag/flag.go:680
	// _ = "end of CoverTab[115369]"
//line /usr/local/go/src/flag/flag.go:680
}

// NFlag returns the number of command-line flags that have been set.
func NFlag() int {
//line /usr/local/go/src/flag/flag.go:683
	_go_fuzz_dep_.CoverTab[115370]++
//line /usr/local/go/src/flag/flag.go:683
	return len(CommandLine.actual)
//line /usr/local/go/src/flag/flag.go:683
	// _ = "end of CoverTab[115370]"
//line /usr/local/go/src/flag/flag.go:683
}

// Arg returns the i'th argument. Arg(0) is the first remaining argument
//line /usr/local/go/src/flag/flag.go:685
// after flags have been processed. Arg returns an empty string if the
//line /usr/local/go/src/flag/flag.go:685
// requested element does not exist.
//line /usr/local/go/src/flag/flag.go:688
func (f *FlagSet) Arg(i int) string {
//line /usr/local/go/src/flag/flag.go:688
	_go_fuzz_dep_.CoverTab[115371]++
						if i < 0 || func() bool {
//line /usr/local/go/src/flag/flag.go:689
		_go_fuzz_dep_.CoverTab[115373]++
//line /usr/local/go/src/flag/flag.go:689
		return i >= len(f.args)
//line /usr/local/go/src/flag/flag.go:689
		// _ = "end of CoverTab[115373]"
//line /usr/local/go/src/flag/flag.go:689
	}() {
//line /usr/local/go/src/flag/flag.go:689
		_go_fuzz_dep_.CoverTab[115374]++
							return ""
//line /usr/local/go/src/flag/flag.go:690
		// _ = "end of CoverTab[115374]"
	} else {
//line /usr/local/go/src/flag/flag.go:691
		_go_fuzz_dep_.CoverTab[115375]++
//line /usr/local/go/src/flag/flag.go:691
		// _ = "end of CoverTab[115375]"
//line /usr/local/go/src/flag/flag.go:691
	}
//line /usr/local/go/src/flag/flag.go:691
	// _ = "end of CoverTab[115371]"
//line /usr/local/go/src/flag/flag.go:691
	_go_fuzz_dep_.CoverTab[115372]++
						return f.args[i]
//line /usr/local/go/src/flag/flag.go:692
	// _ = "end of CoverTab[115372]"
}

// Arg returns the i'th command-line argument. Arg(0) is the first remaining argument
//line /usr/local/go/src/flag/flag.go:695
// after flags have been processed. Arg returns an empty string if the
//line /usr/local/go/src/flag/flag.go:695
// requested element does not exist.
//line /usr/local/go/src/flag/flag.go:698
func Arg(i int) string {
//line /usr/local/go/src/flag/flag.go:698
	_go_fuzz_dep_.CoverTab[115376]++
						return CommandLine.Arg(i)
//line /usr/local/go/src/flag/flag.go:699
	// _ = "end of CoverTab[115376]"
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int {
//line /usr/local/go/src/flag/flag.go:703
	_go_fuzz_dep_.CoverTab[115377]++
//line /usr/local/go/src/flag/flag.go:703
	return len(f.args)
//line /usr/local/go/src/flag/flag.go:703
	// _ = "end of CoverTab[115377]"
//line /usr/local/go/src/flag/flag.go:703
}

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int {
//line /usr/local/go/src/flag/flag.go:706
	_go_fuzz_dep_.CoverTab[115378]++
//line /usr/local/go/src/flag/flag.go:706
	return len(CommandLine.args)
//line /usr/local/go/src/flag/flag.go:706
	// _ = "end of CoverTab[115378]"
//line /usr/local/go/src/flag/flag.go:706
}

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string {
//line /usr/local/go/src/flag/flag.go:709
	_go_fuzz_dep_.CoverTab[115379]++
//line /usr/local/go/src/flag/flag.go:709
	return f.args
//line /usr/local/go/src/flag/flag.go:709
	// _ = "end of CoverTab[115379]"
//line /usr/local/go/src/flag/flag.go:709
}

// Args returns the non-flag command-line arguments.
func Args() []string {
//line /usr/local/go/src/flag/flag.go:712
	_go_fuzz_dep_.CoverTab[115380]++
//line /usr/local/go/src/flag/flag.go:712
	return CommandLine.args
//line /usr/local/go/src/flag/flag.go:712
	// _ = "end of CoverTab[115380]"
//line /usr/local/go/src/flag/flag.go:712
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:714
// The argument p points to a bool variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:716
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
//line /usr/local/go/src/flag/flag.go:716
	_go_fuzz_dep_.CoverTab[115381]++
						f.Var(newBoolValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:717
	// _ = "end of CoverTab[115381]"
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:720
// The argument p points to a bool variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:722
func BoolVar(p *bool, name string, value bool, usage string) {
//line /usr/local/go/src/flag/flag.go:722
	_go_fuzz_dep_.CoverTab[115382]++
						CommandLine.Var(newBoolValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:723
	// _ = "end of CoverTab[115382]"
}

// Bool defines a bool flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:726
// The return value is the address of a bool variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:728
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
//line /usr/local/go/src/flag/flag.go:728
	_go_fuzz_dep_.CoverTab[115383]++
						p := new(bool)
						f.BoolVar(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:731
	// _ = "end of CoverTab[115383]"
}

// Bool defines a bool flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:734
// The return value is the address of a bool variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:736
func Bool(name string, value bool, usage string) *bool {
//line /usr/local/go/src/flag/flag.go:736
	_go_fuzz_dep_.CoverTab[115384]++
						return CommandLine.Bool(name, value, usage)
//line /usr/local/go/src/flag/flag.go:737
	// _ = "end of CoverTab[115384]"
}

// IntVar defines an int flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:740
// The argument p points to an int variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:742
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
//line /usr/local/go/src/flag/flag.go:742
	_go_fuzz_dep_.CoverTab[115385]++
						f.Var(newIntValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:743
	// _ = "end of CoverTab[115385]"
}

// IntVar defines an int flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:746
// The argument p points to an int variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:748
func IntVar(p *int, name string, value int, usage string) {
//line /usr/local/go/src/flag/flag.go:748
	_go_fuzz_dep_.CoverTab[115386]++
						CommandLine.Var(newIntValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:749
	// _ = "end of CoverTab[115386]"
}

// Int defines an int flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:752
// The return value is the address of an int variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:754
func (f *FlagSet) Int(name string, value int, usage string) *int {
//line /usr/local/go/src/flag/flag.go:754
	_go_fuzz_dep_.CoverTab[115387]++
						p := new(int)
						f.IntVar(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:757
	// _ = "end of CoverTab[115387]"
}

// Int defines an int flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:760
// The return value is the address of an int variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:762
func Int(name string, value int, usage string) *int {
//line /usr/local/go/src/flag/flag.go:762
	_go_fuzz_dep_.CoverTab[115388]++
						return CommandLine.Int(name, value, usage)
//line /usr/local/go/src/flag/flag.go:763
	// _ = "end of CoverTab[115388]"
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:766
// The argument p points to an int64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:768
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
//line /usr/local/go/src/flag/flag.go:768
	_go_fuzz_dep_.CoverTab[115389]++
						f.Var(newInt64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:769
	// _ = "end of CoverTab[115389]"
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:772
// The argument p points to an int64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:774
func Int64Var(p *int64, name string, value int64, usage string) {
//line /usr/local/go/src/flag/flag.go:774
	_go_fuzz_dep_.CoverTab[115390]++
						CommandLine.Var(newInt64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:775
	// _ = "end of CoverTab[115390]"
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:778
// The return value is the address of an int64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:780
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
//line /usr/local/go/src/flag/flag.go:780
	_go_fuzz_dep_.CoverTab[115391]++
						p := new(int64)
						f.Int64Var(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:783
	// _ = "end of CoverTab[115391]"
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:786
// The return value is the address of an int64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:788
func Int64(name string, value int64, usage string) *int64 {
//line /usr/local/go/src/flag/flag.go:788
	_go_fuzz_dep_.CoverTab[115392]++
						return CommandLine.Int64(name, value, usage)
//line /usr/local/go/src/flag/flag.go:789
	// _ = "end of CoverTab[115392]"
}

// UintVar defines a uint flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:792
// The argument p points to a uint variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:794
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
//line /usr/local/go/src/flag/flag.go:794
	_go_fuzz_dep_.CoverTab[115393]++
						f.Var(newUintValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:795
	// _ = "end of CoverTab[115393]"
}

// UintVar defines a uint flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:798
// The argument p points to a uint variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:800
func UintVar(p *uint, name string, value uint, usage string) {
//line /usr/local/go/src/flag/flag.go:800
	_go_fuzz_dep_.CoverTab[115394]++
						CommandLine.Var(newUintValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:801
	// _ = "end of CoverTab[115394]"
}

// Uint defines a uint flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:804
// The return value is the address of a uint variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:806
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
//line /usr/local/go/src/flag/flag.go:806
	_go_fuzz_dep_.CoverTab[115395]++
						p := new(uint)
						f.UintVar(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:809
	// _ = "end of CoverTab[115395]"
}

// Uint defines a uint flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:812
// The return value is the address of a uint variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:814
func Uint(name string, value uint, usage string) *uint {
//line /usr/local/go/src/flag/flag.go:814
	_go_fuzz_dep_.CoverTab[115396]++
						return CommandLine.Uint(name, value, usage)
//line /usr/local/go/src/flag/flag.go:815
	// _ = "end of CoverTab[115396]"
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:818
// The argument p points to a uint64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:820
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
//line /usr/local/go/src/flag/flag.go:820
	_go_fuzz_dep_.CoverTab[115397]++
						f.Var(newUint64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:821
	// _ = "end of CoverTab[115397]"
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:824
// The argument p points to a uint64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:826
func Uint64Var(p *uint64, name string, value uint64, usage string) {
//line /usr/local/go/src/flag/flag.go:826
	_go_fuzz_dep_.CoverTab[115398]++
						CommandLine.Var(newUint64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:827
	// _ = "end of CoverTab[115398]"
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:830
// The return value is the address of a uint64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:832
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
//line /usr/local/go/src/flag/flag.go:832
	_go_fuzz_dep_.CoverTab[115399]++
						p := new(uint64)
						f.Uint64Var(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:835
	// _ = "end of CoverTab[115399]"
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:838
// The return value is the address of a uint64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:840
func Uint64(name string, value uint64, usage string) *uint64 {
//line /usr/local/go/src/flag/flag.go:840
	_go_fuzz_dep_.CoverTab[115400]++
						return CommandLine.Uint64(name, value, usage)
//line /usr/local/go/src/flag/flag.go:841
	// _ = "end of CoverTab[115400]"
}

// StringVar defines a string flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:844
// The argument p points to a string variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:846
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
//line /usr/local/go/src/flag/flag.go:846
	_go_fuzz_dep_.CoverTab[115401]++
						f.Var(newStringValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:847
	// _ = "end of CoverTab[115401]"
}

// StringVar defines a string flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:850
// The argument p points to a string variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:852
func StringVar(p *string, name string, value string, usage string) {
//line /usr/local/go/src/flag/flag.go:852
	_go_fuzz_dep_.CoverTab[115402]++
						CommandLine.Var(newStringValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:853
	// _ = "end of CoverTab[115402]"
}

// String defines a string flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:856
// The return value is the address of a string variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:858
func (f *FlagSet) String(name string, value string, usage string) *string {
//line /usr/local/go/src/flag/flag.go:858
	_go_fuzz_dep_.CoverTab[115403]++
						p := new(string)
						f.StringVar(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:861
	// _ = "end of CoverTab[115403]"
}

// String defines a string flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:864
// The return value is the address of a string variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:866
func String(name string, value string, usage string) *string {
//line /usr/local/go/src/flag/flag.go:866
	_go_fuzz_dep_.CoverTab[115404]++
						return CommandLine.String(name, value, usage)
//line /usr/local/go/src/flag/flag.go:867
	// _ = "end of CoverTab[115404]"
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:870
// The argument p points to a float64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:872
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
//line /usr/local/go/src/flag/flag.go:872
	_go_fuzz_dep_.CoverTab[115405]++
						f.Var(newFloat64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:873
	// _ = "end of CoverTab[115405]"
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:876
// The argument p points to a float64 variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:878
func Float64Var(p *float64, name string, value float64, usage string) {
//line /usr/local/go/src/flag/flag.go:878
	_go_fuzz_dep_.CoverTab[115406]++
						CommandLine.Var(newFloat64Value(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:879
	// _ = "end of CoverTab[115406]"
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:882
// The return value is the address of a float64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:884
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
//line /usr/local/go/src/flag/flag.go:884
	_go_fuzz_dep_.CoverTab[115407]++
						p := new(float64)
						f.Float64Var(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:887
	// _ = "end of CoverTab[115407]"
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:890
// The return value is the address of a float64 variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:892
func Float64(name string, value float64, usage string) *float64 {
//line /usr/local/go/src/flag/flag.go:892
	_go_fuzz_dep_.CoverTab[115408]++
						return CommandLine.Float64(name, value, usage)
//line /usr/local/go/src/flag/flag.go:893
	// _ = "end of CoverTab[115408]"
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:896
// The argument p points to a time.Duration variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:896
// The flag accepts a value acceptable to time.ParseDuration.
//line /usr/local/go/src/flag/flag.go:899
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
//line /usr/local/go/src/flag/flag.go:899
	_go_fuzz_dep_.CoverTab[115409]++
						f.Var(newDurationValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:900
	// _ = "end of CoverTab[115409]"
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:903
// The argument p points to a time.Duration variable in which to store the value of the flag.
//line /usr/local/go/src/flag/flag.go:903
// The flag accepts a value acceptable to time.ParseDuration.
//line /usr/local/go/src/flag/flag.go:906
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
//line /usr/local/go/src/flag/flag.go:906
	_go_fuzz_dep_.CoverTab[115410]++
						CommandLine.Var(newDurationValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:907
	// _ = "end of CoverTab[115410]"
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:910
// The return value is the address of a time.Duration variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:910
// The flag accepts a value acceptable to time.ParseDuration.
//line /usr/local/go/src/flag/flag.go:913
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
//line /usr/local/go/src/flag/flag.go:913
	_go_fuzz_dep_.CoverTab[115411]++
						p := new(time.Duration)
						f.DurationVar(p, name, value, usage)
						return p
//line /usr/local/go/src/flag/flag.go:916
	// _ = "end of CoverTab[115411]"
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:919
// The return value is the address of a time.Duration variable that stores the value of the flag.
//line /usr/local/go/src/flag/flag.go:919
// The flag accepts a value acceptable to time.ParseDuration.
//line /usr/local/go/src/flag/flag.go:922
func Duration(name string, value time.Duration, usage string) *time.Duration {
//line /usr/local/go/src/flag/flag.go:922
	_go_fuzz_dep_.CoverTab[115412]++
						return CommandLine.Duration(name, value, usage)
//line /usr/local/go/src/flag/flag.go:923
	// _ = "end of CoverTab[115412]"
}

// TextVar defines a flag with a specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:926
// The argument p must be a pointer to a variable that will hold the value
//line /usr/local/go/src/flag/flag.go:926
// of the flag, and p must implement encoding.TextUnmarshaler.
//line /usr/local/go/src/flag/flag.go:926
// If the flag is used, the flag value will be passed to p's UnmarshalText method.
//line /usr/local/go/src/flag/flag.go:926
// The type of the default value must be the same as the type of p.
//line /usr/local/go/src/flag/flag.go:931
func (f *FlagSet) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string) {
//line /usr/local/go/src/flag/flag.go:931
	_go_fuzz_dep_.CoverTab[115413]++
						f.Var(newTextValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:932
	// _ = "end of CoverTab[115413]"
}

// TextVar defines a flag with a specified name, default value, and usage string.
//line /usr/local/go/src/flag/flag.go:935
// The argument p must be a pointer to a variable that will hold the value
//line /usr/local/go/src/flag/flag.go:935
// of the flag, and p must implement encoding.TextUnmarshaler.
//line /usr/local/go/src/flag/flag.go:935
// If the flag is used, the flag value will be passed to p's UnmarshalText method.
//line /usr/local/go/src/flag/flag.go:935
// The type of the default value must be the same as the type of p.
//line /usr/local/go/src/flag/flag.go:940
func TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string) {
//line /usr/local/go/src/flag/flag.go:940
	_go_fuzz_dep_.CoverTab[115414]++
						CommandLine.Var(newTextValue(value, p), name, usage)
//line /usr/local/go/src/flag/flag.go:941
	// _ = "end of CoverTab[115414]"
}

// Func defines a flag with the specified name and usage string.
//line /usr/local/go/src/flag/flag.go:944
// Each time the flag is seen, fn is called with the value of the flag.
//line /usr/local/go/src/flag/flag.go:944
// If fn returns a non-nil error, it will be treated as a flag value parsing error.
//line /usr/local/go/src/flag/flag.go:947
func (f *FlagSet) Func(name, usage string, fn func(string) error) {
//line /usr/local/go/src/flag/flag.go:947
	_go_fuzz_dep_.CoverTab[115415]++
						f.Var(funcValue(fn), name, usage)
//line /usr/local/go/src/flag/flag.go:948
	// _ = "end of CoverTab[115415]"
}

// Func defines a flag with the specified name and usage string.
//line /usr/local/go/src/flag/flag.go:951
// Each time the flag is seen, fn is called with the value of the flag.
//line /usr/local/go/src/flag/flag.go:951
// If fn returns a non-nil error, it will be treated as a flag value parsing error.
//line /usr/local/go/src/flag/flag.go:954
func Func(name, usage string, fn func(string) error) {
//line /usr/local/go/src/flag/flag.go:954
	_go_fuzz_dep_.CoverTab[115416]++
						CommandLine.Func(name, usage, fn)
//line /usr/local/go/src/flag/flag.go:955
	// _ = "end of CoverTab[115416]"
}

// Var defines a flag with the specified name and usage string. The type and
//line /usr/local/go/src/flag/flag.go:958
// value of the flag are represented by the first argument, of type Value, which
//line /usr/local/go/src/flag/flag.go:958
// typically holds a user-defined implementation of Value. For instance, the
//line /usr/local/go/src/flag/flag.go:958
// caller could create a flag that turns a comma-separated string into a slice
//line /usr/local/go/src/flag/flag.go:958
// of strings by giving the slice the methods of Value; in particular, Set would
//line /usr/local/go/src/flag/flag.go:958
// decompose the comma-separated string into the slice.
//line /usr/local/go/src/flag/flag.go:964
func (f *FlagSet) Var(value Value, name string, usage string) {
//line /usr/local/go/src/flag/flag.go:964
	_go_fuzz_dep_.CoverTab[115417]++

						if strings.HasPrefix(name, "-") {
//line /usr/local/go/src/flag/flag.go:966
		_go_fuzz_dep_.CoverTab[115421]++
							panic(f.sprintf("flag %q begins with -", name))
//line /usr/local/go/src/flag/flag.go:967
		// _ = "end of CoverTab[115421]"
	} else {
//line /usr/local/go/src/flag/flag.go:968
		_go_fuzz_dep_.CoverTab[115422]++
//line /usr/local/go/src/flag/flag.go:968
		if strings.Contains(name, "=") {
//line /usr/local/go/src/flag/flag.go:968
			_go_fuzz_dep_.CoverTab[115423]++
								panic(f.sprintf("flag %q contains =", name))
//line /usr/local/go/src/flag/flag.go:969
			// _ = "end of CoverTab[115423]"
		} else {
//line /usr/local/go/src/flag/flag.go:970
			_go_fuzz_dep_.CoverTab[115424]++
//line /usr/local/go/src/flag/flag.go:970
			// _ = "end of CoverTab[115424]"
//line /usr/local/go/src/flag/flag.go:970
		}
//line /usr/local/go/src/flag/flag.go:970
		// _ = "end of CoverTab[115422]"
//line /usr/local/go/src/flag/flag.go:970
	}
//line /usr/local/go/src/flag/flag.go:970
	// _ = "end of CoverTab[115417]"
//line /usr/local/go/src/flag/flag.go:970
	_go_fuzz_dep_.CoverTab[115418]++

//line /usr/local/go/src/flag/flag.go:973
	flag := &Flag{name, usage, value, value.String()}
	_, alreadythere := f.formal[name]
	if alreadythere {
//line /usr/local/go/src/flag/flag.go:975
		_go_fuzz_dep_.CoverTab[115425]++
							var msg string
							if f.name == "" {
//line /usr/local/go/src/flag/flag.go:977
			_go_fuzz_dep_.CoverTab[115427]++
								msg = f.sprintf("flag redefined: %s", name)
//line /usr/local/go/src/flag/flag.go:978
			// _ = "end of CoverTab[115427]"
		} else {
//line /usr/local/go/src/flag/flag.go:979
			_go_fuzz_dep_.CoverTab[115428]++
								msg = f.sprintf("%s flag redefined: %s", f.name, name)
//line /usr/local/go/src/flag/flag.go:980
			// _ = "end of CoverTab[115428]"
		}
//line /usr/local/go/src/flag/flag.go:981
		// _ = "end of CoverTab[115425]"
//line /usr/local/go/src/flag/flag.go:981
		_go_fuzz_dep_.CoverTab[115426]++
							panic(msg)
//line /usr/local/go/src/flag/flag.go:982
		// _ = "end of CoverTab[115426]"
	} else {
//line /usr/local/go/src/flag/flag.go:983
		_go_fuzz_dep_.CoverTab[115429]++
//line /usr/local/go/src/flag/flag.go:983
		// _ = "end of CoverTab[115429]"
//line /usr/local/go/src/flag/flag.go:983
	}
//line /usr/local/go/src/flag/flag.go:983
	// _ = "end of CoverTab[115418]"
//line /usr/local/go/src/flag/flag.go:983
	_go_fuzz_dep_.CoverTab[115419]++
						if f.formal == nil {
//line /usr/local/go/src/flag/flag.go:984
		_go_fuzz_dep_.CoverTab[115430]++
							f.formal = make(map[string]*Flag)
//line /usr/local/go/src/flag/flag.go:985
		// _ = "end of CoverTab[115430]"
	} else {
//line /usr/local/go/src/flag/flag.go:986
		_go_fuzz_dep_.CoverTab[115431]++
//line /usr/local/go/src/flag/flag.go:986
		// _ = "end of CoverTab[115431]"
//line /usr/local/go/src/flag/flag.go:986
	}
//line /usr/local/go/src/flag/flag.go:986
	// _ = "end of CoverTab[115419]"
//line /usr/local/go/src/flag/flag.go:986
	_go_fuzz_dep_.CoverTab[115420]++
						f.formal[name] = flag
//line /usr/local/go/src/flag/flag.go:987
	// _ = "end of CoverTab[115420]"
}

// Var defines a flag with the specified name and usage string. The type and
//line /usr/local/go/src/flag/flag.go:990
// value of the flag are represented by the first argument, of type Value, which
//line /usr/local/go/src/flag/flag.go:990
// typically holds a user-defined implementation of Value. For instance, the
//line /usr/local/go/src/flag/flag.go:990
// caller could create a flag that turns a comma-separated string into a slice
//line /usr/local/go/src/flag/flag.go:990
// of strings by giving the slice the methods of Value; in particular, Set would
//line /usr/local/go/src/flag/flag.go:990
// decompose the comma-separated string into the slice.
//line /usr/local/go/src/flag/flag.go:996
func Var(value Value, name string, usage string) {
//line /usr/local/go/src/flag/flag.go:996
	_go_fuzz_dep_.CoverTab[115432]++
						CommandLine.Var(value, name, usage)
//line /usr/local/go/src/flag/flag.go:997
	// _ = "end of CoverTab[115432]"
}

// sprintf formats the message, prints it to output, and returns it.
func (f *FlagSet) sprintf(format string, a ...any) string {
//line /usr/local/go/src/flag/flag.go:1001
	_go_fuzz_dep_.CoverTab[115433]++
						msg := fmt.Sprintf(format, a...)
						fmt.Fprintln(f.Output(), msg)
						return msg
//line /usr/local/go/src/flag/flag.go:1004
	// _ = "end of CoverTab[115433]"
}

// failf prints to standard error a formatted error and usage message and
//line /usr/local/go/src/flag/flag.go:1007
// returns the error.
//line /usr/local/go/src/flag/flag.go:1009
func (f *FlagSet) failf(format string, a ...any) error {
//line /usr/local/go/src/flag/flag.go:1009
	_go_fuzz_dep_.CoverTab[115434]++
						msg := f.sprintf(format, a...)
						f.usage()
						return errors.New(msg)
//line /usr/local/go/src/flag/flag.go:1012
	// _ = "end of CoverTab[115434]"
}

// usage calls the Usage method for the flag set if one is specified,
//line /usr/local/go/src/flag/flag.go:1015
// or the appropriate default usage function otherwise.
//line /usr/local/go/src/flag/flag.go:1017
func (f *FlagSet) usage() {
//line /usr/local/go/src/flag/flag.go:1017
	_go_fuzz_dep_.CoverTab[115435]++
						if f.Usage == nil {
//line /usr/local/go/src/flag/flag.go:1018
		_go_fuzz_dep_.CoverTab[115436]++
							f.defaultUsage()
//line /usr/local/go/src/flag/flag.go:1019
		// _ = "end of CoverTab[115436]"
	} else {
//line /usr/local/go/src/flag/flag.go:1020
		_go_fuzz_dep_.CoverTab[115437]++
							f.Usage()
//line /usr/local/go/src/flag/flag.go:1021
		// _ = "end of CoverTab[115437]"
	}
//line /usr/local/go/src/flag/flag.go:1022
	// _ = "end of CoverTab[115435]"
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (bool, error) {
//line /usr/local/go/src/flag/flag.go:1026
	_go_fuzz_dep_.CoverTab[115438]++
						if len(f.args) == 0 {
//line /usr/local/go/src/flag/flag.go:1027
		_go_fuzz_dep_.CoverTab[115447]++
							return false, nil
//line /usr/local/go/src/flag/flag.go:1028
		// _ = "end of CoverTab[115447]"
	} else {
//line /usr/local/go/src/flag/flag.go:1029
		_go_fuzz_dep_.CoverTab[115448]++
//line /usr/local/go/src/flag/flag.go:1029
		// _ = "end of CoverTab[115448]"
//line /usr/local/go/src/flag/flag.go:1029
	}
//line /usr/local/go/src/flag/flag.go:1029
	// _ = "end of CoverTab[115438]"
//line /usr/local/go/src/flag/flag.go:1029
	_go_fuzz_dep_.CoverTab[115439]++
						s := f.args[0]
						if len(s) < 2 || func() bool {
//line /usr/local/go/src/flag/flag.go:1031
		_go_fuzz_dep_.CoverTab[115449]++
//line /usr/local/go/src/flag/flag.go:1031
		return s[0] != '-'
//line /usr/local/go/src/flag/flag.go:1031
		// _ = "end of CoverTab[115449]"
//line /usr/local/go/src/flag/flag.go:1031
	}() {
//line /usr/local/go/src/flag/flag.go:1031
		_go_fuzz_dep_.CoverTab[115450]++
							return false, nil
//line /usr/local/go/src/flag/flag.go:1032
		// _ = "end of CoverTab[115450]"
	} else {
//line /usr/local/go/src/flag/flag.go:1033
		_go_fuzz_dep_.CoverTab[115451]++
//line /usr/local/go/src/flag/flag.go:1033
		// _ = "end of CoverTab[115451]"
//line /usr/local/go/src/flag/flag.go:1033
	}
//line /usr/local/go/src/flag/flag.go:1033
	// _ = "end of CoverTab[115439]"
//line /usr/local/go/src/flag/flag.go:1033
	_go_fuzz_dep_.CoverTab[115440]++
						numMinuses := 1
						if s[1] == '-' {
//line /usr/local/go/src/flag/flag.go:1035
		_go_fuzz_dep_.CoverTab[115452]++
							numMinuses++
							if len(s) == 2 {
//line /usr/local/go/src/flag/flag.go:1037
			_go_fuzz_dep_.CoverTab[115453]++
								f.args = f.args[1:]
								return false, nil
//line /usr/local/go/src/flag/flag.go:1039
			// _ = "end of CoverTab[115453]"
		} else {
//line /usr/local/go/src/flag/flag.go:1040
			_go_fuzz_dep_.CoverTab[115454]++
//line /usr/local/go/src/flag/flag.go:1040
			// _ = "end of CoverTab[115454]"
//line /usr/local/go/src/flag/flag.go:1040
		}
//line /usr/local/go/src/flag/flag.go:1040
		// _ = "end of CoverTab[115452]"
	} else {
//line /usr/local/go/src/flag/flag.go:1041
		_go_fuzz_dep_.CoverTab[115455]++
//line /usr/local/go/src/flag/flag.go:1041
		// _ = "end of CoverTab[115455]"
//line /usr/local/go/src/flag/flag.go:1041
	}
//line /usr/local/go/src/flag/flag.go:1041
	// _ = "end of CoverTab[115440]"
//line /usr/local/go/src/flag/flag.go:1041
	_go_fuzz_dep_.CoverTab[115441]++
						name := s[numMinuses:]
						if len(name) == 0 || func() bool {
//line /usr/local/go/src/flag/flag.go:1043
		_go_fuzz_dep_.CoverTab[115456]++
//line /usr/local/go/src/flag/flag.go:1043
		return name[0] == '-'
//line /usr/local/go/src/flag/flag.go:1043
		// _ = "end of CoverTab[115456]"
//line /usr/local/go/src/flag/flag.go:1043
	}() || func() bool {
//line /usr/local/go/src/flag/flag.go:1043
		_go_fuzz_dep_.CoverTab[115457]++
//line /usr/local/go/src/flag/flag.go:1043
		return name[0] == '='
//line /usr/local/go/src/flag/flag.go:1043
		// _ = "end of CoverTab[115457]"
//line /usr/local/go/src/flag/flag.go:1043
	}() {
//line /usr/local/go/src/flag/flag.go:1043
		_go_fuzz_dep_.CoverTab[115458]++
							return false, f.failf("bad flag syntax: %s", s)
//line /usr/local/go/src/flag/flag.go:1044
		// _ = "end of CoverTab[115458]"
	} else {
//line /usr/local/go/src/flag/flag.go:1045
		_go_fuzz_dep_.CoverTab[115459]++
//line /usr/local/go/src/flag/flag.go:1045
		// _ = "end of CoverTab[115459]"
//line /usr/local/go/src/flag/flag.go:1045
	}
//line /usr/local/go/src/flag/flag.go:1045
	// _ = "end of CoverTab[115441]"
//line /usr/local/go/src/flag/flag.go:1045
	_go_fuzz_dep_.CoverTab[115442]++

//line /usr/local/go/src/flag/flag.go:1048
	f.args = f.args[1:]
	hasValue := false
	value := ""
	for i := 1; i < len(name); i++ {
//line /usr/local/go/src/flag/flag.go:1051
		_go_fuzz_dep_.CoverTab[115460]++
							if name[i] == '=' {
//line /usr/local/go/src/flag/flag.go:1052
			_go_fuzz_dep_.CoverTab[115461]++
								value = name[i+1:]
								hasValue = true
								name = name[0:i]
								break
//line /usr/local/go/src/flag/flag.go:1056
			// _ = "end of CoverTab[115461]"
		} else {
//line /usr/local/go/src/flag/flag.go:1057
			_go_fuzz_dep_.CoverTab[115462]++
//line /usr/local/go/src/flag/flag.go:1057
			// _ = "end of CoverTab[115462]"
//line /usr/local/go/src/flag/flag.go:1057
		}
//line /usr/local/go/src/flag/flag.go:1057
		// _ = "end of CoverTab[115460]"
	}
//line /usr/local/go/src/flag/flag.go:1058
	// _ = "end of CoverTab[115442]"
//line /usr/local/go/src/flag/flag.go:1058
	_go_fuzz_dep_.CoverTab[115443]++

						flag, ok := f.formal[name]
						if !ok {
//line /usr/local/go/src/flag/flag.go:1061
		_go_fuzz_dep_.CoverTab[115463]++
							if name == "help" || func() bool {
//line /usr/local/go/src/flag/flag.go:1062
			_go_fuzz_dep_.CoverTab[115465]++
//line /usr/local/go/src/flag/flag.go:1062
			return name == "h"
//line /usr/local/go/src/flag/flag.go:1062
			// _ = "end of CoverTab[115465]"
//line /usr/local/go/src/flag/flag.go:1062
		}() {
//line /usr/local/go/src/flag/flag.go:1062
			_go_fuzz_dep_.CoverTab[115466]++
								f.usage()
								return false, ErrHelp
//line /usr/local/go/src/flag/flag.go:1064
			// _ = "end of CoverTab[115466]"
		} else {
//line /usr/local/go/src/flag/flag.go:1065
			_go_fuzz_dep_.CoverTab[115467]++
//line /usr/local/go/src/flag/flag.go:1065
			// _ = "end of CoverTab[115467]"
//line /usr/local/go/src/flag/flag.go:1065
		}
//line /usr/local/go/src/flag/flag.go:1065
		// _ = "end of CoverTab[115463]"
//line /usr/local/go/src/flag/flag.go:1065
		_go_fuzz_dep_.CoverTab[115464]++
							return false, f.failf("flag provided but not defined: -%s", name)
//line /usr/local/go/src/flag/flag.go:1066
		// _ = "end of CoverTab[115464]"
	} else {
//line /usr/local/go/src/flag/flag.go:1067
		_go_fuzz_dep_.CoverTab[115468]++
//line /usr/local/go/src/flag/flag.go:1067
		// _ = "end of CoverTab[115468]"
//line /usr/local/go/src/flag/flag.go:1067
	}
//line /usr/local/go/src/flag/flag.go:1067
	// _ = "end of CoverTab[115443]"
//line /usr/local/go/src/flag/flag.go:1067
	_go_fuzz_dep_.CoverTab[115444]++

						if fv, ok := flag.Value.(boolFlag); ok && func() bool {
//line /usr/local/go/src/flag/flag.go:1069
		_go_fuzz_dep_.CoverTab[115469]++
//line /usr/local/go/src/flag/flag.go:1069
		return fv.IsBoolFlag()
//line /usr/local/go/src/flag/flag.go:1069
		// _ = "end of CoverTab[115469]"
//line /usr/local/go/src/flag/flag.go:1069
	}() {
//line /usr/local/go/src/flag/flag.go:1069
		_go_fuzz_dep_.CoverTab[115470]++
							if hasValue {
//line /usr/local/go/src/flag/flag.go:1070
			_go_fuzz_dep_.CoverTab[115471]++
								if err := fv.Set(value); err != nil {
//line /usr/local/go/src/flag/flag.go:1071
				_go_fuzz_dep_.CoverTab[115472]++
									return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
//line /usr/local/go/src/flag/flag.go:1072
				// _ = "end of CoverTab[115472]"
			} else {
//line /usr/local/go/src/flag/flag.go:1073
				_go_fuzz_dep_.CoverTab[115473]++
//line /usr/local/go/src/flag/flag.go:1073
				// _ = "end of CoverTab[115473]"
//line /usr/local/go/src/flag/flag.go:1073
			}
//line /usr/local/go/src/flag/flag.go:1073
			// _ = "end of CoverTab[115471]"
		} else {
//line /usr/local/go/src/flag/flag.go:1074
			_go_fuzz_dep_.CoverTab[115474]++
								if err := fv.Set("true"); err != nil {
//line /usr/local/go/src/flag/flag.go:1075
				_go_fuzz_dep_.CoverTab[115475]++
									return false, f.failf("invalid boolean flag %s: %v", name, err)
//line /usr/local/go/src/flag/flag.go:1076
				// _ = "end of CoverTab[115475]"
			} else {
//line /usr/local/go/src/flag/flag.go:1077
				_go_fuzz_dep_.CoverTab[115476]++
//line /usr/local/go/src/flag/flag.go:1077
				// _ = "end of CoverTab[115476]"
//line /usr/local/go/src/flag/flag.go:1077
			}
//line /usr/local/go/src/flag/flag.go:1077
			// _ = "end of CoverTab[115474]"
		}
//line /usr/local/go/src/flag/flag.go:1078
		// _ = "end of CoverTab[115470]"
	} else {
//line /usr/local/go/src/flag/flag.go:1079
		_go_fuzz_dep_.CoverTab[115477]++

							if !hasValue && func() bool {
//line /usr/local/go/src/flag/flag.go:1081
			_go_fuzz_dep_.CoverTab[115480]++
//line /usr/local/go/src/flag/flag.go:1081
			return len(f.args) > 0
//line /usr/local/go/src/flag/flag.go:1081
			// _ = "end of CoverTab[115480]"
//line /usr/local/go/src/flag/flag.go:1081
		}() {
//line /usr/local/go/src/flag/flag.go:1081
			_go_fuzz_dep_.CoverTab[115481]++

								hasValue = true
								value, f.args = f.args[0], f.args[1:]
//line /usr/local/go/src/flag/flag.go:1084
			// _ = "end of CoverTab[115481]"
		} else {
//line /usr/local/go/src/flag/flag.go:1085
			_go_fuzz_dep_.CoverTab[115482]++
//line /usr/local/go/src/flag/flag.go:1085
			// _ = "end of CoverTab[115482]"
//line /usr/local/go/src/flag/flag.go:1085
		}
//line /usr/local/go/src/flag/flag.go:1085
		// _ = "end of CoverTab[115477]"
//line /usr/local/go/src/flag/flag.go:1085
		_go_fuzz_dep_.CoverTab[115478]++
							if !hasValue {
//line /usr/local/go/src/flag/flag.go:1086
			_go_fuzz_dep_.CoverTab[115483]++
								return false, f.failf("flag needs an argument: -%s", name)
//line /usr/local/go/src/flag/flag.go:1087
			// _ = "end of CoverTab[115483]"
		} else {
//line /usr/local/go/src/flag/flag.go:1088
			_go_fuzz_dep_.CoverTab[115484]++
//line /usr/local/go/src/flag/flag.go:1088
			// _ = "end of CoverTab[115484]"
//line /usr/local/go/src/flag/flag.go:1088
		}
//line /usr/local/go/src/flag/flag.go:1088
		// _ = "end of CoverTab[115478]"
//line /usr/local/go/src/flag/flag.go:1088
		_go_fuzz_dep_.CoverTab[115479]++
							if err := flag.Value.Set(value); err != nil {
//line /usr/local/go/src/flag/flag.go:1089
			_go_fuzz_dep_.CoverTab[115485]++
								return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
//line /usr/local/go/src/flag/flag.go:1090
			// _ = "end of CoverTab[115485]"
		} else {
//line /usr/local/go/src/flag/flag.go:1091
			_go_fuzz_dep_.CoverTab[115486]++
//line /usr/local/go/src/flag/flag.go:1091
			// _ = "end of CoverTab[115486]"
//line /usr/local/go/src/flag/flag.go:1091
		}
//line /usr/local/go/src/flag/flag.go:1091
		// _ = "end of CoverTab[115479]"
	}
//line /usr/local/go/src/flag/flag.go:1092
	// _ = "end of CoverTab[115444]"
//line /usr/local/go/src/flag/flag.go:1092
	_go_fuzz_dep_.CoverTab[115445]++
						if f.actual == nil {
//line /usr/local/go/src/flag/flag.go:1093
		_go_fuzz_dep_.CoverTab[115487]++
							f.actual = make(map[string]*Flag)
//line /usr/local/go/src/flag/flag.go:1094
		// _ = "end of CoverTab[115487]"
	} else {
//line /usr/local/go/src/flag/flag.go:1095
		_go_fuzz_dep_.CoverTab[115488]++
//line /usr/local/go/src/flag/flag.go:1095
		// _ = "end of CoverTab[115488]"
//line /usr/local/go/src/flag/flag.go:1095
	}
//line /usr/local/go/src/flag/flag.go:1095
	// _ = "end of CoverTab[115445]"
//line /usr/local/go/src/flag/flag.go:1095
	_go_fuzz_dep_.CoverTab[115446]++
						f.actual[name] = flag
						return true, nil
//line /usr/local/go/src/flag/flag.go:1097
	// _ = "end of CoverTab[115446]"
}

// Parse parses flag definitions from the argument list, which should not
//line /usr/local/go/src/flag/flag.go:1100
// include the command name. Must be called after all flags in the FlagSet
//line /usr/local/go/src/flag/flag.go:1100
// are defined and before flags are accessed by the program.
//line /usr/local/go/src/flag/flag.go:1100
// The return value will be ErrHelp if -help or -h were set but not defined.
//line /usr/local/go/src/flag/flag.go:1104
func (f *FlagSet) Parse(arguments []string) error {
//line /usr/local/go/src/flag/flag.go:1104
	_go_fuzz_dep_.CoverTab[115489]++
						f.parsed = true
						f.args = arguments
						for {
//line /usr/local/go/src/flag/flag.go:1107
		_go_fuzz_dep_.CoverTab[115491]++
							seen, err := f.parseOne()
							if seen {
//line /usr/local/go/src/flag/flag.go:1109
			_go_fuzz_dep_.CoverTab[115494]++
								continue
//line /usr/local/go/src/flag/flag.go:1110
			// _ = "end of CoverTab[115494]"
		} else {
//line /usr/local/go/src/flag/flag.go:1111
			_go_fuzz_dep_.CoverTab[115495]++
//line /usr/local/go/src/flag/flag.go:1111
			// _ = "end of CoverTab[115495]"
//line /usr/local/go/src/flag/flag.go:1111
		}
//line /usr/local/go/src/flag/flag.go:1111
		// _ = "end of CoverTab[115491]"
//line /usr/local/go/src/flag/flag.go:1111
		_go_fuzz_dep_.CoverTab[115492]++
							if err == nil {
//line /usr/local/go/src/flag/flag.go:1112
			_go_fuzz_dep_.CoverTab[115496]++
								break
//line /usr/local/go/src/flag/flag.go:1113
			// _ = "end of CoverTab[115496]"
		} else {
//line /usr/local/go/src/flag/flag.go:1114
			_go_fuzz_dep_.CoverTab[115497]++
//line /usr/local/go/src/flag/flag.go:1114
			// _ = "end of CoverTab[115497]"
//line /usr/local/go/src/flag/flag.go:1114
		}
//line /usr/local/go/src/flag/flag.go:1114
		// _ = "end of CoverTab[115492]"
//line /usr/local/go/src/flag/flag.go:1114
		_go_fuzz_dep_.CoverTab[115493]++
							switch f.errorHandling {
		case ContinueOnError:
//line /usr/local/go/src/flag/flag.go:1116
			_go_fuzz_dep_.CoverTab[115498]++
								return err
//line /usr/local/go/src/flag/flag.go:1117
			// _ = "end of CoverTab[115498]"
		case ExitOnError:
//line /usr/local/go/src/flag/flag.go:1118
			_go_fuzz_dep_.CoverTab[115499]++
								if err == ErrHelp {
//line /usr/local/go/src/flag/flag.go:1119
				_go_fuzz_dep_.CoverTab[115503]++
									os.Exit(0)
//line /usr/local/go/src/flag/flag.go:1120
				// _ = "end of CoverTab[115503]"
			} else {
//line /usr/local/go/src/flag/flag.go:1121
				_go_fuzz_dep_.CoverTab[115504]++
//line /usr/local/go/src/flag/flag.go:1121
				// _ = "end of CoverTab[115504]"
//line /usr/local/go/src/flag/flag.go:1121
			}
//line /usr/local/go/src/flag/flag.go:1121
			// _ = "end of CoverTab[115499]"
//line /usr/local/go/src/flag/flag.go:1121
			_go_fuzz_dep_.CoverTab[115500]++
								os.Exit(2)
//line /usr/local/go/src/flag/flag.go:1122
			// _ = "end of CoverTab[115500]"
		case PanicOnError:
//line /usr/local/go/src/flag/flag.go:1123
			_go_fuzz_dep_.CoverTab[115501]++
								panic(err)
//line /usr/local/go/src/flag/flag.go:1124
			// _ = "end of CoverTab[115501]"
//line /usr/local/go/src/flag/flag.go:1124
		default:
//line /usr/local/go/src/flag/flag.go:1124
			_go_fuzz_dep_.CoverTab[115502]++
//line /usr/local/go/src/flag/flag.go:1124
			// _ = "end of CoverTab[115502]"
		}
//line /usr/local/go/src/flag/flag.go:1125
		// _ = "end of CoverTab[115493]"
	}
//line /usr/local/go/src/flag/flag.go:1126
	// _ = "end of CoverTab[115489]"
//line /usr/local/go/src/flag/flag.go:1126
	_go_fuzz_dep_.CoverTab[115490]++
						return nil
//line /usr/local/go/src/flag/flag.go:1127
	// _ = "end of CoverTab[115490]"
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
//line /usr/local/go/src/flag/flag.go:1131
	_go_fuzz_dep_.CoverTab[115505]++
						return f.parsed
//line /usr/local/go/src/flag/flag.go:1132
	// _ = "end of CoverTab[115505]"
}

// Parse parses the command-line flags from os.Args[1:]. Must be called
//line /usr/local/go/src/flag/flag.go:1135
// after all flags are defined and before flags are accessed by the program.
//line /usr/local/go/src/flag/flag.go:1137
func Parse() {
//line /usr/local/go/src/flag/flag.go:1137
	_go_fuzz_dep_.CoverTab[115506]++

						CommandLine.Parse(os.Args[1:])
//line /usr/local/go/src/flag/flag.go:1139
	// _ = "end of CoverTab[115506]"
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
//line /usr/local/go/src/flag/flag.go:1143
	_go_fuzz_dep_.CoverTab[115507]++
						return CommandLine.Parsed()
//line /usr/local/go/src/flag/flag.go:1144
	// _ = "end of CoverTab[115507]"
}

// CommandLine is the default set of command-line flags, parsed from os.Args.
//line /usr/local/go/src/flag/flag.go:1147
// The top-level functions such as BoolVar, Arg, and so on are wrappers for the
//line /usr/local/go/src/flag/flag.go:1147
// methods of CommandLine.
//line /usr/local/go/src/flag/flag.go:1150
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

func init() {

//line /usr/local/go/src/flag/flag.go:1157
	CommandLine.Usage = commandLineUsage
}

func commandLineUsage() {
//line /usr/local/go/src/flag/flag.go:1160
	_go_fuzz_dep_.CoverTab[115508]++
						Usage()
//line /usr/local/go/src/flag/flag.go:1161
	// _ = "end of CoverTab[115508]"
}

// NewFlagSet returns a new, empty flag set with the specified name and
//line /usr/local/go/src/flag/flag.go:1164
// error handling property. If the name is not empty, it will be printed
//line /usr/local/go/src/flag/flag.go:1164
// in the default usage message and in error messages.
//line /usr/local/go/src/flag/flag.go:1167
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
//line /usr/local/go/src/flag/flag.go:1167
	_go_fuzz_dep_.CoverTab[115509]++
						f := &FlagSet{
		name:		name,
		errorHandling:	errorHandling,
	}
						f.Usage = f.defaultUsage
						return f
//line /usr/local/go/src/flag/flag.go:1173
	// _ = "end of CoverTab[115509]"
}

// Init sets the name and error handling property for a flag set.
//line /usr/local/go/src/flag/flag.go:1176
// By default, the zero FlagSet uses an empty name and the
//line /usr/local/go/src/flag/flag.go:1176
// ContinueOnError error handling policy.
//line /usr/local/go/src/flag/flag.go:1179
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
//line /usr/local/go/src/flag/flag.go:1179
	_go_fuzz_dep_.CoverTab[115510]++
						f.name = name
						f.errorHandling = errorHandling
//line /usr/local/go/src/flag/flag.go:1181
	// _ = "end of CoverTab[115510]"
}

//line /usr/local/go/src/flag/flag.go:1182
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/flag/flag.go:1182
var _ = _go_fuzz_dep_.CoverTab
