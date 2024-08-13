//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:17
)

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ConfigState houses the configuration options used by spew to format and
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// display values.  There is a global instance, Config, that is used to control
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// all top-level Formatter and Dump functionality.  Each ConfigState instance
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// provides methods equivalent to the top-level functions.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// The zero value for ConfigState provides no indentation.  You would typically
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// want to set it to a space or a tab.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// Alternatively, you can use NewDefaultConfig to get a ConfigState instance
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// with default settings.  See the documentation of NewDefaultConfig for default
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:26
// values.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:37
type ConfigState struct {
	// Indent specifies the string to use for each indentation level.  The
	// global config instance that all top-level functions use set this to a
	// single space by default.  If you would like more indentation, you might
	// set this to a tab with "\t" or perhaps two spaces with "  ".
	Indent	string

	// MaxDepth controls the maximum number of levels to descend into nested
	// data structures.  The default, 0, means there is no limit.
	//
	// NOTE: Circular data structures are properly detected, so it is not
	// necessary to set this value unless you specifically want to limit deeply
	// nested data structures.
	MaxDepth	int

	// DisableMethods specifies whether or not error and Stringer interfaces are
	// invoked for types that implement them.
	DisableMethods	bool

	// DisablePointerMethods specifies whether or not to check for and invoke
	// error and Stringer interfaces on types which only accept a pointer
	// receiver when the current type is not a pointer.
	//
	// NOTE: This might be an unsafe action since calling one of these methods
	// with a pointer receiver could technically mutate the value, however,
	// in practice, types which choose to satisify an error or Stringer
	// interface with a pointer receiver should not be mutating their state
	// inside these interface methods.  As a result, this option relies on
	// access to the unsafe package, so it will not have any effect when
	// running in environments without access to the unsafe package such as
	// Google App Engine or with the "safe" build tag specified.
	DisablePointerMethods	bool

	// DisablePointerAddresses specifies whether to disable the printing of
	// pointer addresses. This is useful when diffing data structures in tests.
	DisablePointerAddresses	bool

	// DisableCapacities specifies whether to disable the printing of capacities
	// for arrays, slices, maps and channels. This is useful when diffing
	// data structures in tests.
	DisableCapacities	bool

	// ContinueOnMethod specifies whether or not recursion should continue once
	// a custom error or Stringer interface is invoked.  The default, false,
	// means it will print the results of invoking the custom error or Stringer
	// interface and return immediately instead of continuing to recurse into
	// the internals of the data type.
	//
	// NOTE: This flag does not have any effect if method invocation is disabled
	// via the DisableMethods or DisablePointerMethods options.
	ContinueOnMethod	bool

	// SortKeys specifies map keys should be sorted before being printed. Use
	// this to have a more deterministic, diffable output.  Note that only
	// native types (bool, int, uint, floats, uintptr and string) and types
	// that support the error or Stringer interfaces (if methods are
	// enabled) are supported, with other types sorted according to the
	// reflect.Value.String() output which guarantees display stability.
	SortKeys	bool

	// SpewKeys specifies that, as a last resort attempt, map keys should
	// be spewed to strings and sorted by those strings.  This is only
	// considered if SortKeys is true.
	SpewKeys	bool
}

// Config is the active configuration of the top-level functions.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:103
// The configuration can be changed by modifying the contents of spew.Config.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:105
var Config = ConfigState{Indent: " "}

// Errorf is a wrapper for fmt.Errorf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
// the formatted string as a value that satisfies error.  See NewFormatter
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
// for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:107
//	fmt.Errorf(format, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:115
func (c *ConfigState) Errorf(format string, a ...interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:115
	_go_fuzz_dep_.CoverTab[81637]++
											return fmt.Errorf(format, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:116
	// _ = "end of CoverTab[81637]"
}

// Fprint is a wrapper for fmt.Fprint that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
// the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:119
//	fmt.Fprint(w, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:127
func (c *ConfigState) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:127
	_go_fuzz_dep_.CoverTab[81638]++
											return fmt.Fprint(w, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:128
	// _ = "end of CoverTab[81638]"
}

// Fprintf is a wrapper for fmt.Fprintf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
// the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:131
//	fmt.Fprintf(w, format, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:139
func (c *ConfigState) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:139
	_go_fuzz_dep_.CoverTab[81639]++
											return fmt.Fprintf(w, format, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:140
	// _ = "end of CoverTab[81639]"
}

// Fprintln is a wrapper for fmt.Fprintln that treats each argument as if it
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
// passed with a Formatter interface returned by c.NewFormatter.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:143
//	fmt.Fprintln(w, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:150
func (c *ConfigState) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:150
	_go_fuzz_dep_.CoverTab[81640]++
											return fmt.Fprintln(w, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:151
	// _ = "end of CoverTab[81640]"
}

// Print is a wrapper for fmt.Print that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
// the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:154
//	fmt.Print(c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:162
func (c *ConfigState) Print(a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:162
	_go_fuzz_dep_.CoverTab[81641]++
											return fmt.Print(c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:163
	// _ = "end of CoverTab[81641]"
}

// Printf is a wrapper for fmt.Printf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
// the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:166
//	fmt.Printf(format, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:174
func (c *ConfigState) Printf(format string, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:174
	_go_fuzz_dep_.CoverTab[81642]++
											return fmt.Printf(format, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:175
	// _ = "end of CoverTab[81642]"
}

// Println is a wrapper for fmt.Println that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
// the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:178
//	fmt.Println(c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:186
func (c *ConfigState) Println(a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:186
	_go_fuzz_dep_.CoverTab[81643]++
											return fmt.Println(c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:187
	// _ = "end of CoverTab[81643]"
}

// Sprint is a wrapper for fmt.Sprint that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
// the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:190
//	fmt.Sprint(c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:197
func (c *ConfigState) Sprint(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:197
	_go_fuzz_dep_.CoverTab[81644]++
											return fmt.Sprint(c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:198
	// _ = "end of CoverTab[81644]"
}

// Sprintf is a wrapper for fmt.Sprintf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
// passed with a Formatter interface returned by c.NewFormatter.  It returns
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
// the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:201
//	fmt.Sprintf(format, c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:208
func (c *ConfigState) Sprintf(format string, a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:208
	_go_fuzz_dep_.CoverTab[81645]++
											return fmt.Sprintf(format, c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:209
	// _ = "end of CoverTab[81645]"
}

// Sprintln is a wrapper for fmt.Sprintln that treats each argument as if it
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
// were passed with a Formatter interface returned by c.NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
// returns the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:212
//	fmt.Sprintln(c.NewFormatter(a), c.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:219
func (c *ConfigState) Sprintln(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:219
	_go_fuzz_dep_.CoverTab[81646]++
											return fmt.Sprintln(c.convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:220
	// _ = "end of CoverTab[81646]"
}

/*
NewFormatter returns a custom formatter that satisfies the fmt.Formatter
interface.  As a result, it integrates cleanly with standard fmt package
printing functions.  The formatter is useful for inline printing of smaller data
types similar to the standard %v format specifier.

The custom formatter only responds to the %v (most compact), %+v (adds pointer
addresses), %#v (adds types), and %#+v (adds types and pointer addresses) verb
combinations.  Any other verbs such as %x and %q will be sent to the the
standard fmt package for formatting.  In addition, the custom formatter ignores
the width and precision arguments (however they will still work on the format
specifiers not handled by the custom formatter).

Typically this function shouldn't be called directly.  It is much easier to make
use of the custom formatter by calling one of the convenience functions such as
c.Printf, c.Println, or c.Printf.
*/
func (c *ConfigState) NewFormatter(v interface{}) fmt.Formatter {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:240
	_go_fuzz_dep_.CoverTab[81647]++
											return newFormatter(c, v)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:241
	// _ = "end of CoverTab[81647]"
}

// Fdump formats and displays the passed arguments to io.Writer w.  It formats
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:244
// exactly the same as Dump.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:246
func (c *ConfigState) Fdump(w io.Writer, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:246
	_go_fuzz_dep_.CoverTab[81648]++
											fdump(c, w, a...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:247
	// _ = "end of CoverTab[81648]"
}

/*
Dump displays the passed parameters to standard out with newlines, customizable
indentation, and additional debug information such as complete types and all
pointer addresses used to indirect to the final value.  It provides the
following features over the built-in printing facilities provided by the fmt
package:

  - Pointers are dereferenced and followed
  - Circular data structures are detected and handled properly
  - Custom Stringer/error interfaces are optionally invoked, including
    on unexported types
  - Custom types which only implement the Stringer/error interfaces via
    a pointer receiver are optionally invoked when passing non-pointer
    variables
  - Byte arrays and slices are dumped like the hexdump -C command which
    includes offsets, byte values in hex, and ASCII output

The configuration options are controlled by modifying the public members
of c.  See ConfigState for options documentation.

See Fdump if you would prefer dumping to an arbitrary io.Writer or Sdump to
get the formatted result as a string.
*/
func (c *ConfigState) Dump(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:273
	_go_fuzz_dep_.CoverTab[81649]++
											fdump(c, os.Stdout, a...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:274
	// _ = "end of CoverTab[81649]"
}

// Sdump returns a string with the passed arguments formatted exactly the same
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:277
// as Dump.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:279
func (c *ConfigState) Sdump(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:279
	_go_fuzz_dep_.CoverTab[81650]++
											var buf bytes.Buffer
											fdump(c, &buf, a...)
											return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:282
	// _ = "end of CoverTab[81650]"
}

// convertArgs accepts a slice of arguments and returns a slice of the same
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:285
// length with each argument converted to a spew Formatter interface using
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:285
// the ConfigState associated with s.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:288
func (c *ConfigState) convertArgs(args []interface{}) (formatters []interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:288
	_go_fuzz_dep_.CoverTab[81651]++
											formatters = make([]interface{}, len(args))
											for index, arg := range args {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:290
		_go_fuzz_dep_.CoverTab[81653]++
												formatters[index] = newFormatter(c, arg)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:291
		// _ = "end of CoverTab[81653]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:292
	// _ = "end of CoverTab[81651]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:292
	_go_fuzz_dep_.CoverTab[81652]++
											return formatters
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:293
	// _ = "end of CoverTab[81652]"
}

// NewDefaultConfig returns a ConfigState with the following default settings.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	Indent: " "
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	MaxDepth: 0
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	DisableMethods: false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	DisablePointerMethods: false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	ContinueOnMethod: false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:296
//	SortKeys: false
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:304
func NewDefaultConfig() *ConfigState {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:304
	_go_fuzz_dep_.CoverTab[81654]++
											return &ConfigState{Indent: " "}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:305
	// _ = "end of CoverTab[81654]"
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:306
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/config.go:306
var _ = _go_fuzz_dep_.CoverTab
