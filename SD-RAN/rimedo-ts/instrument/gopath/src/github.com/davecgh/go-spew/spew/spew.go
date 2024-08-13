//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
package spew

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:17
)

import (
	"fmt"
	"io"
)

// Errorf is a wrapper for fmt.Errorf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
// returns the formatted string as a value that satisfies error.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:24
//	fmt.Errorf(format, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:32
func Errorf(format string, a ...interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:32
	_go_fuzz_dep_.CoverTab[81976]++
											return fmt.Errorf(format, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:33
	// _ = "end of CoverTab[81976]"
}

// Fprint is a wrapper for fmt.Fprint that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
// returns the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:36
//	fmt.Fprint(w, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:44
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:44
	_go_fuzz_dep_.CoverTab[81977]++
											return fmt.Fprint(w, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:45
	// _ = "end of CoverTab[81977]"
}

// Fprintf is a wrapper for fmt.Fprintf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
// returns the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:48
//	fmt.Fprintf(w, format, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:56
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:56
	_go_fuzz_dep_.CoverTab[81978]++
											return fmt.Fprintf(w, format, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:57
	// _ = "end of CoverTab[81978]"
}

// Fprintln is a wrapper for fmt.Fprintln that treats each argument as if it
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
// passed with a default Formatter interface returned by NewFormatter.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:60
//	fmt.Fprintln(w, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:67
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:67
	_go_fuzz_dep_.CoverTab[81979]++
											return fmt.Fprintln(w, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:68
	// _ = "end of CoverTab[81979]"
}

// Print is a wrapper for fmt.Print that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
// returns the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:71
//	fmt.Print(spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:79
func Print(a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:79
	_go_fuzz_dep_.CoverTab[81980]++
											return fmt.Print(convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:80
	// _ = "end of CoverTab[81980]"
}

// Printf is a wrapper for fmt.Printf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
// returns the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:83
//	fmt.Printf(format, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:91
func Printf(format string, a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:91
	_go_fuzz_dep_.CoverTab[81981]++
											return fmt.Printf(format, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:92
	// _ = "end of CoverTab[81981]"
}

// Println is a wrapper for fmt.Println that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
// returns the number of bytes written and any write error encountered.  See
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
// NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:95
//	fmt.Println(spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:103
func Println(a ...interface{}) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:103
	_go_fuzz_dep_.CoverTab[81982]++
											return fmt.Println(convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:104
	// _ = "end of CoverTab[81982]"
}

// Sprint is a wrapper for fmt.Sprint that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
// returns the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:107
//	fmt.Sprint(spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:114
func Sprint(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:114
	_go_fuzz_dep_.CoverTab[81983]++
											return fmt.Sprint(convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:115
	// _ = "end of CoverTab[81983]"
}

// Sprintf is a wrapper for fmt.Sprintf that treats each argument as if it were
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
// passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
// returns the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:118
//	fmt.Sprintf(format, spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:125
func Sprintf(format string, a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:125
	_go_fuzz_dep_.CoverTab[81984]++
											return fmt.Sprintf(format, convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:126
	// _ = "end of CoverTab[81984]"
}

// Sprintln is a wrapper for fmt.Sprintln that treats each argument as if it
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
// were passed with a default Formatter interface returned by NewFormatter.  It
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
// returns the resulting string.  See NewFormatter for formatting details.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
// This function is shorthand for the following syntax:
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
//
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:129
//	fmt.Sprintln(spew.NewFormatter(a), spew.NewFormatter(b))
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:136
func Sprintln(a ...interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:136
	_go_fuzz_dep_.CoverTab[81985]++
											return fmt.Sprintln(convertArgs(a)...)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:137
	// _ = "end of CoverTab[81985]"
}

// convertArgs accepts a slice of arguments and returns a slice of the same
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:140
// length with each argument converted to a default spew Formatter interface.
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:142
func convertArgs(args []interface{}) (formatters []interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:142
	_go_fuzz_dep_.CoverTab[81986]++
											formatters = make([]interface{}, len(args))
											for index, arg := range args {
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:144
		_go_fuzz_dep_.CoverTab[81988]++
												formatters[index] = NewFormatter(arg)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:145
		// _ = "end of CoverTab[81988]"
	}
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:146
	// _ = "end of CoverTab[81986]"
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:146
	_go_fuzz_dep_.CoverTab[81987]++
											return formatters
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:147
	// _ = "end of CoverTab[81987]"
}

//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:148
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/davecgh/go-spew@v1.1.1/spew/spew.go:148
var _ = _go_fuzz_dep_.CoverTab
