// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Package bisect can be used by compilers and other programs
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// to serve as a target for the bisect debugging tool.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// See [golang.org/x/tools/cmd/bisect] for details about using the tool.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// To be a bisect target, allowing bisect to help determine which of a set of independent
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// changes provokes a failure, a program needs to:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//  1. Define a way to accept a change pattern on its command line or in its environment.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     The most common mechanism is a command-line flag.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     The pattern can be passed to [New] to create a [Matcher], the compiled form of a pattern.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//  2. Assign each change a unique ID. One possibility is to use a sequence number,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     but the most common mechanism is to hash some kind of identifying information
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     like the file and line number where the change might be applied.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     [Hash] hashes its arguments to compute an ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//  3. Enable each change that the pattern says should be enabled.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     The [Matcher.ShouldEnable] method answers this question for a given change ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//  4. Print a report identifying each change that the pattern says should be printed.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     The [Matcher.ShouldPrint] method answers this question for a given change ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     The report consists of one more lines on standard error or standard output
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     that contain a “match marker”. [Marker] returns the match marker for a given ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     When bisect reports a change as causing the failure, it identifies the change
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     by printing the report lines with the match marker removed.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// # Example Usage
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// A program starts by defining how it receives the pattern. In this example, we will assume a flag.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// The next step is to compile the pattern:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	m, err := bisect.New(patternFlag)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	if err != nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		log.Fatal(err)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Then, each time a potential change is considered, the program computes
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// a change ID by hashing identifying information (source file and line, in this case)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// and then calls m.ShouldPrint and m.ShouldEnable to decide whether to
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// print and enable the change, respectively. The two can return different values
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// depending on whether bisect is trying to find a minimal set of changes to
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// disable or to enable to provoke the failure.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// It is usually helpful to write a helper function that accepts the identifying information
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// and then takes care of hashing, printing, and reporting whether the identified change
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// should be enabled. For example, a helper for changes identified by a file and line number
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// would be:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	func ShouldEnable(file string, line int) {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		h := bisect.Hash(file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		if m.ShouldPrint(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			fmt.Fprintf(os.Stderr, "%v %s:%d\n", bisect.Marker(h), file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		return m.ShouldEnable(h)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Finally, note that New returns a nil Matcher when there is no pattern,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// meaning that the target is not running under bisect at all,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// so all changes should be enabled and none should be printed.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// In that common case, the computation of the hash can be avoided entirely
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// by checking for m == nil first:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	func ShouldEnable(file string, line int) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		if m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			return false
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		h := bisect.Hash(file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		if m.ShouldPrint(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			fmt.Fprintf(os.Stderr, "%v %s:%d\n", bisect.Marker(h), file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		return m.ShouldEnable(h)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// When the identifying information is expensive to format, this code can call
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// [Matcher.MarkerOnly] to find out whether short report lines containing only the
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// marker are permitted for a given run. (Bisect permits such lines when it is
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// still exploring the space of possible changes and will not be showing the
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// output to the user.) If so, the client can choose to print only the marker:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	func ShouldEnable(file string, line int) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		if m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			return false
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		h := bisect.Hash(file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		if m.ShouldPrint(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			if m.MarkerOnly() {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//				bisect.PrintMarker(os.Stderr)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//				fmt.Fprintf(os.Stderr, "%v %s:%d\n", bisect.Marker(h), file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//			}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//		return m.ShouldEnable(h)
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//	}
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// This specific helper – deciding whether to enable a change identified by
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// file and line number and printing about the change when necessary – is
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// provided by the [Matcher.FileLine] method.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Another common usage is deciding whether to make a change in a function
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// based on the caller's stack, to identify the specific calling contexts that the
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// change breaks. The [Matcher.Stack] method takes care of obtaining the stack,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// printing it when necessary, and reporting whether to enable the change
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// based on that stack.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// # Pattern Syntax
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Patterns are generated by the bisect tool and interpreted by [New].
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Users should not have to understand the patterns except when
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// debugging a target's bisect support or debugging the bisect tool itself.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// The pattern syntax selecting a change is a sequence of bit strings
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// separated by + and - operators. Each bit string denotes the set of
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// changes with IDs ending in those bits, + is set addition, - is set subtraction,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// and the expression is evaluated in the usual left-to-right order.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// The special binary number “y” denotes the set of all changes,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// standing in for the empty bit string.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// In the expression, all the + operators must appear before all the - operators.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// A leading + adds to an empty set. A leading - subtracts from the set of all
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// possible suffixes.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// For example:
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//   - “01+10” and “+01+10” both denote the set of changes
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     with IDs ending with the bits 01 or 10.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//   - “01+10-1001” denotes the set of changes with IDs
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     ending with the bits 01 or 10, but excluding those ending in 1001.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//   - “-01-1000” and “y-01-1000 both denote the set of all changes
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     with IDs not ending in 01 nor 1000.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//   - “0+1-01+001” is not a valid pattern, because all the + operators do not
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//     appear before all the - operators.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// In the syntaxes described so far, the pattern specifies the changes to
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// enable and report. If a pattern is prefixed by a “!”, the meaning
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// changes: the pattern specifies the changes to DISABLE and report. This
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// mode of operation is needed when a program passes with all changes
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// enabled but fails with no changes enabled. In this case, bisect
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// searches for minimal sets of changes to disable.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Put another way, the leading “!” inverts the result from [Matcher.ShouldEnable]
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// but does not invert the result from [Matcher.ShouldPrint].
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// As a convenience for manual debugging, “n” is an alias for “!y”,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// meaning to disable and report all changes.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// Finally, a leading “v” in the pattern indicates that the reports will be shown
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// to the user of bisect to describe the changes involved in a failure.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// At the API level, the leading “v” causes [Matcher.Visible] to return true.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// See the next section for details.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// # Match Reports
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// The target program must enable only those changed matched
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// by the pattern, and it must print a match report for each such change.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// A match report consists of one or more lines of text that will be
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// printed by the bisect tool to describe a change implicated in causing
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// a failure. Each line in the report for a given change must contain a
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// match marker with that change ID, as returned by [Marker].
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// The markers are elided when displaying the lines to the user.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// A match marker has the form “[bisect-match 0x1234]” where
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// 0x1234 is the change ID in hexadecimal.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// An alternate form is “[bisect-match 010101]”, giving the change ID in binary.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
//
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// When [Matcher.Visible] returns false, the match reports are only
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// being processed by bisect to learn the set of enabled changes,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// not shown to the user, meaning that each report can be a match
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// marker on a line by itself, eliding the usual textual description.
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// When the textual description is expensive to compute,
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// checking [Matcher.Visible] can help the avoid that expense
//line /snap/go/10455/src/internal/bisect/bisect.go:5
// in most runs.
//line /snap/go/10455/src/internal/bisect/bisect.go:177
package bisect

//line /snap/go/10455/src/internal/bisect/bisect.go:177
import (
//line /snap/go/10455/src/internal/bisect/bisect.go:177
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/internal/bisect/bisect.go:177
)
//line /snap/go/10455/src/internal/bisect/bisect.go:177
import (
//line /snap/go/10455/src/internal/bisect/bisect.go:177
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/internal/bisect/bisect.go:177
)

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// New creates and returns a new Matcher implementing the given pattern.
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// The pattern syntax is defined in the package doc comment.
//line /snap/go/10455/src/internal/bisect/bisect.go:186
//
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// In addition to the pattern syntax syntax, New("") returns nil, nil.
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// The nil *Matcher is valid for use: it returns true from ShouldEnable
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// and false from ShouldPrint for all changes. Callers can avoid calling
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// [Hash], [Matcher.ShouldEnable], and [Matcher.ShouldPrint] entirely
//line /snap/go/10455/src/internal/bisect/bisect.go:186
// when they recognize the nil Matcher.
//line /snap/go/10455/src/internal/bisect/bisect.go:194
func New(pattern string) (*Matcher, error) {
//line /snap/go/10455/src/internal/bisect/bisect.go:194
	_go_fuzz_dep_.CoverTab[3528]++
								if pattern == "" {
//line /snap/go/10455/src/internal/bisect/bisect.go:195
		_go_fuzz_dep_.CoverTab[526747]++
//line /snap/go/10455/src/internal/bisect/bisect.go:195
		_go_fuzz_dep_.CoverTab[3535]++
									return nil, nil
//line /snap/go/10455/src/internal/bisect/bisect.go:196
		// _ = "end of CoverTab[3535]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:197
		_go_fuzz_dep_.CoverTab[526748]++
//line /snap/go/10455/src/internal/bisect/bisect.go:197
		_go_fuzz_dep_.CoverTab[3536]++
//line /snap/go/10455/src/internal/bisect/bisect.go:197
		// _ = "end of CoverTab[3536]"
//line /snap/go/10455/src/internal/bisect/bisect.go:197
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:197
	// _ = "end of CoverTab[3528]"
//line /snap/go/10455/src/internal/bisect/bisect.go:197
	_go_fuzz_dep_.CoverTab[3529]++

								m := new(Matcher)

								p := pattern

//line /snap/go/10455/src/internal/bisect/bisect.go:204
	if len(p) > 0 && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:204
		_go_fuzz_dep_.CoverTab[3537]++
//line /snap/go/10455/src/internal/bisect/bisect.go:204
		return p[0] == 'q'
//line /snap/go/10455/src/internal/bisect/bisect.go:204
		// _ = "end of CoverTab[3537]"
//line /snap/go/10455/src/internal/bisect/bisect.go:204
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:204
		_go_fuzz_dep_.CoverTab[526749]++
//line /snap/go/10455/src/internal/bisect/bisect.go:204
		_go_fuzz_dep_.CoverTab[3538]++
									m.quiet = true
									p = p[1:]
									if p == "" {
//line /snap/go/10455/src/internal/bisect/bisect.go:207
			_go_fuzz_dep_.CoverTab[526751]++
//line /snap/go/10455/src/internal/bisect/bisect.go:207
			_go_fuzz_dep_.CoverTab[3539]++
										return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:208
			// _ = "end of CoverTab[3539]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:209
			_go_fuzz_dep_.CoverTab[526752]++
//line /snap/go/10455/src/internal/bisect/bisect.go:209
			_go_fuzz_dep_.CoverTab[3540]++
//line /snap/go/10455/src/internal/bisect/bisect.go:209
			// _ = "end of CoverTab[3540]"
//line /snap/go/10455/src/internal/bisect/bisect.go:209
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:209
		// _ = "end of CoverTab[3538]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:210
		_go_fuzz_dep_.CoverTab[526750]++
//line /snap/go/10455/src/internal/bisect/bisect.go:210
		_go_fuzz_dep_.CoverTab[3541]++
//line /snap/go/10455/src/internal/bisect/bisect.go:210
		// _ = "end of CoverTab[3541]"
//line /snap/go/10455/src/internal/bisect/bisect.go:210
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:210
	// _ = "end of CoverTab[3529]"
//line /snap/go/10455/src/internal/bisect/bisect.go:210
	_go_fuzz_dep_.CoverTab[3530]++
//line /snap/go/10455/src/internal/bisect/bisect.go:210
	_go_fuzz_dep_.CoverTab[786601] = 0

								for len(p) > 0 && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		_go_fuzz_dep_.CoverTab[3542]++
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		return p[0] == 'v'
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		// _ = "end of CoverTab[3542]"
//line /snap/go/10455/src/internal/bisect/bisect.go:212
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		if _go_fuzz_dep_.CoverTab[786601] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:212
			_go_fuzz_dep_.CoverTab[526864]++
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:212
			_go_fuzz_dep_.CoverTab[526865]++
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		_go_fuzz_dep_.CoverTab[786601] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:212
		_go_fuzz_dep_.CoverTab[3543]++
									m.verbose = true
									m.quiet = false
									p = p[1:]
									if p == "" {
//line /snap/go/10455/src/internal/bisect/bisect.go:216
			_go_fuzz_dep_.CoverTab[526753]++
//line /snap/go/10455/src/internal/bisect/bisect.go:216
			_go_fuzz_dep_.CoverTab[3544]++
										return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:217
			// _ = "end of CoverTab[3544]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:218
			_go_fuzz_dep_.CoverTab[526754]++
//line /snap/go/10455/src/internal/bisect/bisect.go:218
			_go_fuzz_dep_.CoverTab[3545]++
//line /snap/go/10455/src/internal/bisect/bisect.go:218
			// _ = "end of CoverTab[3545]"
//line /snap/go/10455/src/internal/bisect/bisect.go:218
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:218
		// _ = "end of CoverTab[3543]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:219
	if _go_fuzz_dep_.CoverTab[786601] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:219
		_go_fuzz_dep_.CoverTab[526866]++
//line /snap/go/10455/src/internal/bisect/bisect.go:219
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:219
		_go_fuzz_dep_.CoverTab[526867]++
//line /snap/go/10455/src/internal/bisect/bisect.go:219
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:219
	// _ = "end of CoverTab[3530]"
//line /snap/go/10455/src/internal/bisect/bisect.go:219
	_go_fuzz_dep_.CoverTab[3531]++

//line /snap/go/10455/src/internal/bisect/bisect.go:223
	m.enable = true
//line /snap/go/10455/src/internal/bisect/bisect.go:223
	_go_fuzz_dep_.CoverTab[786602] = 0
								for len(p) > 0 && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		_go_fuzz_dep_.CoverTab[3546]++
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		return p[0] == '!'
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		// _ = "end of CoverTab[3546]"
//line /snap/go/10455/src/internal/bisect/bisect.go:224
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		if _go_fuzz_dep_.CoverTab[786602] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:224
			_go_fuzz_dep_.CoverTab[526868]++
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:224
			_go_fuzz_dep_.CoverTab[526869]++
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		_go_fuzz_dep_.CoverTab[786602] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:224
		_go_fuzz_dep_.CoverTab[3547]++
									m.enable = !m.enable
									p = p[1:]
									if p == "" {
//line /snap/go/10455/src/internal/bisect/bisect.go:227
			_go_fuzz_dep_.CoverTab[526755]++
//line /snap/go/10455/src/internal/bisect/bisect.go:227
			_go_fuzz_dep_.CoverTab[3548]++
										return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:228
			// _ = "end of CoverTab[3548]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:229
			_go_fuzz_dep_.CoverTab[526756]++
//line /snap/go/10455/src/internal/bisect/bisect.go:229
			_go_fuzz_dep_.CoverTab[3549]++
//line /snap/go/10455/src/internal/bisect/bisect.go:229
			// _ = "end of CoverTab[3549]"
//line /snap/go/10455/src/internal/bisect/bisect.go:229
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:229
		// _ = "end of CoverTab[3547]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:230
	if _go_fuzz_dep_.CoverTab[786602] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:230
		_go_fuzz_dep_.CoverTab[526870]++
//line /snap/go/10455/src/internal/bisect/bisect.go:230
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:230
		_go_fuzz_dep_.CoverTab[526871]++
//line /snap/go/10455/src/internal/bisect/bisect.go:230
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:230
	// _ = "end of CoverTab[3531]"
//line /snap/go/10455/src/internal/bisect/bisect.go:230
	_go_fuzz_dep_.CoverTab[3532]++

								if p == "n" {
//line /snap/go/10455/src/internal/bisect/bisect.go:232
		_go_fuzz_dep_.CoverTab[526757]++
//line /snap/go/10455/src/internal/bisect/bisect.go:232
		_go_fuzz_dep_.CoverTab[3550]++

									m.enable = !m.enable
									p = "y"
//line /snap/go/10455/src/internal/bisect/bisect.go:235
		// _ = "end of CoverTab[3550]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:236
		_go_fuzz_dep_.CoverTab[526758]++
//line /snap/go/10455/src/internal/bisect/bisect.go:236
		_go_fuzz_dep_.CoverTab[3551]++
//line /snap/go/10455/src/internal/bisect/bisect.go:236
		// _ = "end of CoverTab[3551]"
//line /snap/go/10455/src/internal/bisect/bisect.go:236
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:236
	// _ = "end of CoverTab[3532]"
//line /snap/go/10455/src/internal/bisect/bisect.go:236
	_go_fuzz_dep_.CoverTab[3533]++

//line /snap/go/10455/src/internal/bisect/bisect.go:239
	result := true
								bits := uint64(0)
								start := 0
								wid := 1
//line /snap/go/10455/src/internal/bisect/bisect.go:242
	_go_fuzz_dep_.CoverTab[786603] = 0
								for i := 0; i <= len(p); i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:243
		if _go_fuzz_dep_.CoverTab[786603] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:243
			_go_fuzz_dep_.CoverTab[526872]++
//line /snap/go/10455/src/internal/bisect/bisect.go:243
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:243
			_go_fuzz_dep_.CoverTab[526873]++
//line /snap/go/10455/src/internal/bisect/bisect.go:243
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:243
		_go_fuzz_dep_.CoverTab[786603] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:243
		_go_fuzz_dep_.CoverTab[3552]++

									c := byte('-')
									if i < len(p) {
//line /snap/go/10455/src/internal/bisect/bisect.go:246
			_go_fuzz_dep_.CoverTab[526759]++
//line /snap/go/10455/src/internal/bisect/bisect.go:246
			_go_fuzz_dep_.CoverTab[3555]++
										c = p[i]
//line /snap/go/10455/src/internal/bisect/bisect.go:247
			// _ = "end of CoverTab[3555]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:248
			_go_fuzz_dep_.CoverTab[526760]++
//line /snap/go/10455/src/internal/bisect/bisect.go:248
			_go_fuzz_dep_.CoverTab[3556]++
//line /snap/go/10455/src/internal/bisect/bisect.go:248
			// _ = "end of CoverTab[3556]"
//line /snap/go/10455/src/internal/bisect/bisect.go:248
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:248
		// _ = "end of CoverTab[3552]"
//line /snap/go/10455/src/internal/bisect/bisect.go:248
		_go_fuzz_dep_.CoverTab[3553]++
									if i == start && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			_go_fuzz_dep_.CoverTab[3557]++
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			return wid == 1
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			// _ = "end of CoverTab[3557]"
//line /snap/go/10455/src/internal/bisect/bisect.go:249
		}() && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			_go_fuzz_dep_.CoverTab[3558]++
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			return c == 'x'
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			// _ = "end of CoverTab[3558]"
//line /snap/go/10455/src/internal/bisect/bisect.go:249
		}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			_go_fuzz_dep_.CoverTab[526761]++
//line /snap/go/10455/src/internal/bisect/bisect.go:249
			_go_fuzz_dep_.CoverTab[3559]++
										start = i + 1
										wid = 4
										continue
//line /snap/go/10455/src/internal/bisect/bisect.go:252
			// _ = "end of CoverTab[3559]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:253
			_go_fuzz_dep_.CoverTab[526762]++
//line /snap/go/10455/src/internal/bisect/bisect.go:253
			_go_fuzz_dep_.CoverTab[3560]++
//line /snap/go/10455/src/internal/bisect/bisect.go:253
			// _ = "end of CoverTab[3560]"
//line /snap/go/10455/src/internal/bisect/bisect.go:253
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:253
		// _ = "end of CoverTab[3553]"
//line /snap/go/10455/src/internal/bisect/bisect.go:253
		_go_fuzz_dep_.CoverTab[3554]++
									switch c {
		default:
//line /snap/go/10455/src/internal/bisect/bisect.go:255
			_go_fuzz_dep_.CoverTab[526763]++
//line /snap/go/10455/src/internal/bisect/bisect.go:255
			_go_fuzz_dep_.CoverTab[3561]++
										return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:256
			// _ = "end of CoverTab[3561]"
		case '2', '3', '4', '5', '6', '7', '8', '9':
//line /snap/go/10455/src/internal/bisect/bisect.go:257
			_go_fuzz_dep_.CoverTab[526764]++
//line /snap/go/10455/src/internal/bisect/bisect.go:257
			_go_fuzz_dep_.CoverTab[3562]++
										if wid != 4 {
//line /snap/go/10455/src/internal/bisect/bisect.go:258
				_go_fuzz_dep_.CoverTab[526769]++
//line /snap/go/10455/src/internal/bisect/bisect.go:258
				_go_fuzz_dep_.CoverTab[3572]++
											return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:259
				// _ = "end of CoverTab[3572]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:260
				_go_fuzz_dep_.CoverTab[526770]++
//line /snap/go/10455/src/internal/bisect/bisect.go:260
				_go_fuzz_dep_.CoverTab[3573]++
//line /snap/go/10455/src/internal/bisect/bisect.go:260
				// _ = "end of CoverTab[3573]"
//line /snap/go/10455/src/internal/bisect/bisect.go:260
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:260
			// _ = "end of CoverTab[3562]"
//line /snap/go/10455/src/internal/bisect/bisect.go:260
			_go_fuzz_dep_.CoverTab[3563]++
										fallthrough
//line /snap/go/10455/src/internal/bisect/bisect.go:261
			// _ = "end of CoverTab[3563]"
		case '0', '1':
//line /snap/go/10455/src/internal/bisect/bisect.go:262
			_go_fuzz_dep_.CoverTab[526765]++
//line /snap/go/10455/src/internal/bisect/bisect.go:262
			_go_fuzz_dep_.CoverTab[3564]++
										bits <<= wid
										bits |= uint64(c - '0')
//line /snap/go/10455/src/internal/bisect/bisect.go:264
			// _ = "end of CoverTab[3564]"
		case 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
//line /snap/go/10455/src/internal/bisect/bisect.go:265
			_go_fuzz_dep_.CoverTab[526766]++
//line /snap/go/10455/src/internal/bisect/bisect.go:265
			_go_fuzz_dep_.CoverTab[3565]++
										if wid != 4 {
//line /snap/go/10455/src/internal/bisect/bisect.go:266
				_go_fuzz_dep_.CoverTab[526771]++
//line /snap/go/10455/src/internal/bisect/bisect.go:266
				_go_fuzz_dep_.CoverTab[3574]++
											return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:267
				// _ = "end of CoverTab[3574]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:268
				_go_fuzz_dep_.CoverTab[526772]++
//line /snap/go/10455/src/internal/bisect/bisect.go:268
				_go_fuzz_dep_.CoverTab[3575]++
//line /snap/go/10455/src/internal/bisect/bisect.go:268
				// _ = "end of CoverTab[3575]"
//line /snap/go/10455/src/internal/bisect/bisect.go:268
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:268
			// _ = "end of CoverTab[3565]"
//line /snap/go/10455/src/internal/bisect/bisect.go:268
			_go_fuzz_dep_.CoverTab[3566]++
										bits <<= 4
										bits |= uint64(c&^0x20 - 'A' + 10)
//line /snap/go/10455/src/internal/bisect/bisect.go:270
			// _ = "end of CoverTab[3566]"
		case 'y':
//line /snap/go/10455/src/internal/bisect/bisect.go:271
			_go_fuzz_dep_.CoverTab[526767]++
//line /snap/go/10455/src/internal/bisect/bisect.go:271
			_go_fuzz_dep_.CoverTab[3567]++
										if i+1 < len(p) && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				_go_fuzz_dep_.CoverTab[3576]++
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				return (p[i+1] == '0' || func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:272
					_go_fuzz_dep_.CoverTab[3577]++
//line /snap/go/10455/src/internal/bisect/bisect.go:272
					return p[i+1] == '1'
//line /snap/go/10455/src/internal/bisect/bisect.go:272
					// _ = "end of CoverTab[3577]"
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				}())
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				// _ = "end of CoverTab[3576]"
//line /snap/go/10455/src/internal/bisect/bisect.go:272
			}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				_go_fuzz_dep_.CoverTab[526773]++
//line /snap/go/10455/src/internal/bisect/bisect.go:272
				_go_fuzz_dep_.CoverTab[3578]++
											return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:273
				// _ = "end of CoverTab[3578]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:274
				_go_fuzz_dep_.CoverTab[526774]++
//line /snap/go/10455/src/internal/bisect/bisect.go:274
				_go_fuzz_dep_.CoverTab[3579]++
//line /snap/go/10455/src/internal/bisect/bisect.go:274
				// _ = "end of CoverTab[3579]"
//line /snap/go/10455/src/internal/bisect/bisect.go:274
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:274
			// _ = "end of CoverTab[3567]"
//line /snap/go/10455/src/internal/bisect/bisect.go:274
			_go_fuzz_dep_.CoverTab[3568]++
										bits = 0
//line /snap/go/10455/src/internal/bisect/bisect.go:275
			// _ = "end of CoverTab[3568]"
		case '+', '-':
//line /snap/go/10455/src/internal/bisect/bisect.go:276
			_go_fuzz_dep_.CoverTab[526768]++
//line /snap/go/10455/src/internal/bisect/bisect.go:276
			_go_fuzz_dep_.CoverTab[3569]++
										if c == '+' && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:277
				_go_fuzz_dep_.CoverTab[3580]++
//line /snap/go/10455/src/internal/bisect/bisect.go:277
				return result == false
//line /snap/go/10455/src/internal/bisect/bisect.go:277
				// _ = "end of CoverTab[3580]"
//line /snap/go/10455/src/internal/bisect/bisect.go:277
			}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:277
				_go_fuzz_dep_.CoverTab[526775]++
//line /snap/go/10455/src/internal/bisect/bisect.go:277
				_go_fuzz_dep_.CoverTab[3581]++

											return nil, &parseError{"invalid pattern syntax (+ after -): " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:279
				// _ = "end of CoverTab[3581]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:280
				_go_fuzz_dep_.CoverTab[526776]++
//line /snap/go/10455/src/internal/bisect/bisect.go:280
				_go_fuzz_dep_.CoverTab[3582]++
//line /snap/go/10455/src/internal/bisect/bisect.go:280
				// _ = "end of CoverTab[3582]"
//line /snap/go/10455/src/internal/bisect/bisect.go:280
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:280
			// _ = "end of CoverTab[3569]"
//line /snap/go/10455/src/internal/bisect/bisect.go:280
			_go_fuzz_dep_.CoverTab[3570]++
										if i > 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:281
				_go_fuzz_dep_.CoverTab[526777]++
//line /snap/go/10455/src/internal/bisect/bisect.go:281
				_go_fuzz_dep_.CoverTab[3583]++
											n := (i - start) * wid
											if n > 64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:283
					_go_fuzz_dep_.CoverTab[526779]++
//line /snap/go/10455/src/internal/bisect/bisect.go:283
					_go_fuzz_dep_.CoverTab[3587]++
												return nil, &parseError{"pattern bits too long: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:284
					// _ = "end of CoverTab[3587]"
				} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:285
					_go_fuzz_dep_.CoverTab[526780]++
//line /snap/go/10455/src/internal/bisect/bisect.go:285
					_go_fuzz_dep_.CoverTab[3588]++
//line /snap/go/10455/src/internal/bisect/bisect.go:285
					// _ = "end of CoverTab[3588]"
//line /snap/go/10455/src/internal/bisect/bisect.go:285
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:285
				// _ = "end of CoverTab[3583]"
//line /snap/go/10455/src/internal/bisect/bisect.go:285
				_go_fuzz_dep_.CoverTab[3584]++
											if n <= 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:286
					_go_fuzz_dep_.CoverTab[526781]++
//line /snap/go/10455/src/internal/bisect/bisect.go:286
					_go_fuzz_dep_.CoverTab[3589]++
												return nil, &parseError{"invalid pattern syntax: " + pattern}
//line /snap/go/10455/src/internal/bisect/bisect.go:287
					// _ = "end of CoverTab[3589]"
				} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:288
					_go_fuzz_dep_.CoverTab[526782]++
//line /snap/go/10455/src/internal/bisect/bisect.go:288
					_go_fuzz_dep_.CoverTab[3590]++
//line /snap/go/10455/src/internal/bisect/bisect.go:288
					// _ = "end of CoverTab[3590]"
//line /snap/go/10455/src/internal/bisect/bisect.go:288
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:288
				// _ = "end of CoverTab[3584]"
//line /snap/go/10455/src/internal/bisect/bisect.go:288
				_go_fuzz_dep_.CoverTab[3585]++
											if p[start] == 'y' {
//line /snap/go/10455/src/internal/bisect/bisect.go:289
					_go_fuzz_dep_.CoverTab[526783]++
//line /snap/go/10455/src/internal/bisect/bisect.go:289
					_go_fuzz_dep_.CoverTab[3591]++
												n = 0
//line /snap/go/10455/src/internal/bisect/bisect.go:290
					// _ = "end of CoverTab[3591]"
				} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:291
					_go_fuzz_dep_.CoverTab[526784]++
//line /snap/go/10455/src/internal/bisect/bisect.go:291
					_go_fuzz_dep_.CoverTab[3592]++
//line /snap/go/10455/src/internal/bisect/bisect.go:291
					// _ = "end of CoverTab[3592]"
//line /snap/go/10455/src/internal/bisect/bisect.go:291
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:291
				// _ = "end of CoverTab[3585]"
//line /snap/go/10455/src/internal/bisect/bisect.go:291
				_go_fuzz_dep_.CoverTab[3586]++
											mask := uint64(1)<<n - 1
											m.list = append(m.list, cond{mask, bits, result})
//line /snap/go/10455/src/internal/bisect/bisect.go:293
				// _ = "end of CoverTab[3586]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:294
				_go_fuzz_dep_.CoverTab[526778]++
//line /snap/go/10455/src/internal/bisect/bisect.go:294
				_go_fuzz_dep_.CoverTab[3593]++
//line /snap/go/10455/src/internal/bisect/bisect.go:294
				if c == '-' {
//line /snap/go/10455/src/internal/bisect/bisect.go:294
					_go_fuzz_dep_.CoverTab[526785]++
//line /snap/go/10455/src/internal/bisect/bisect.go:294
					_go_fuzz_dep_.CoverTab[3594]++

												m.list = append(m.list, cond{0, 0, true})
//line /snap/go/10455/src/internal/bisect/bisect.go:296
					// _ = "end of CoverTab[3594]"
				} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:297
					_go_fuzz_dep_.CoverTab[526786]++
//line /snap/go/10455/src/internal/bisect/bisect.go:297
					_go_fuzz_dep_.CoverTab[3595]++
//line /snap/go/10455/src/internal/bisect/bisect.go:297
					// _ = "end of CoverTab[3595]"
//line /snap/go/10455/src/internal/bisect/bisect.go:297
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:297
				// _ = "end of CoverTab[3593]"
//line /snap/go/10455/src/internal/bisect/bisect.go:297
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:297
			// _ = "end of CoverTab[3570]"
//line /snap/go/10455/src/internal/bisect/bisect.go:297
			_go_fuzz_dep_.CoverTab[3571]++
										bits = 0
										result = c == '+'
										start = i + 1
										wid = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:301
			// _ = "end of CoverTab[3571]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:302
		// _ = "end of CoverTab[3554]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:303
	if _go_fuzz_dep_.CoverTab[786603] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:303
		_go_fuzz_dep_.CoverTab[526874]++
//line /snap/go/10455/src/internal/bisect/bisect.go:303
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:303
		_go_fuzz_dep_.CoverTab[526875]++
//line /snap/go/10455/src/internal/bisect/bisect.go:303
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:303
	// _ = "end of CoverTab[3533]"
//line /snap/go/10455/src/internal/bisect/bisect.go:303
	_go_fuzz_dep_.CoverTab[3534]++
								return m, nil
//line /snap/go/10455/src/internal/bisect/bisect.go:304
	// _ = "end of CoverTab[3534]"
}

// A Matcher is the parsed, compiled form of a PATTERN string.
//line /snap/go/10455/src/internal/bisect/bisect.go:307
// The nil *Matcher is valid: it has all changes enabled but none reported.
//line /snap/go/10455/src/internal/bisect/bisect.go:309
type Matcher struct {
	verbose	bool	// annotate reporting with human-helpful information
	quiet	bool	// disables all reporting.  reset if verbose is true. use case is -d=fmahash=qn
	enable	bool	// when true, list is for “enable and report” (when false, “disable and report”)
	list	[]cond	// conditions; later ones win over earlier ones
	dedup	atomicPointerDedup
}

// atomicPointerDedup is an atomic.Pointer[dedup],
//line /snap/go/10455/src/internal/bisect/bisect.go:317
// but we are avoiding using Go 1.19's atomic.Pointer
//line /snap/go/10455/src/internal/bisect/bisect.go:317
// until the bootstrap toolchain can be relied upon to have it.
//line /snap/go/10455/src/internal/bisect/bisect.go:320
type atomicPointerDedup struct {
	p unsafe.Pointer
}

func (p *atomicPointerDedup) Load() *dedup {
//line /snap/go/10455/src/internal/bisect/bisect.go:324
	_go_fuzz_dep_.CoverTab[3596]++
								return (*dedup)(atomic.LoadPointer(&p.p))
//line /snap/go/10455/src/internal/bisect/bisect.go:325
	// _ = "end of CoverTab[3596]"
}

func (p *atomicPointerDedup) CompareAndSwap(old, new *dedup) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:328
	_go_fuzz_dep_.CoverTab[3597]++
								return atomic.CompareAndSwapPointer(&p.p, unsafe.Pointer(old), unsafe.Pointer(new))
//line /snap/go/10455/src/internal/bisect/bisect.go:329
	// _ = "end of CoverTab[3597]"
}

// A cond is a single condition in the matcher.
//line /snap/go/10455/src/internal/bisect/bisect.go:332
// Given an input id, if id&mask == bits, return the result.
//line /snap/go/10455/src/internal/bisect/bisect.go:334
type cond struct {
	mask	uint64
	bits	uint64
	result	bool
}

// MarkerOnly reports whether it is okay to print only the marker for
//line /snap/go/10455/src/internal/bisect/bisect.go:340
// a given change, omitting the identifying information.
//line /snap/go/10455/src/internal/bisect/bisect.go:340
// MarkerOnly returns true when bisect is using the printed reports
//line /snap/go/10455/src/internal/bisect/bisect.go:340
// only for an intermediate search step, not for showing to users.
//line /snap/go/10455/src/internal/bisect/bisect.go:344
func (m *Matcher) MarkerOnly() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:344
	_go_fuzz_dep_.CoverTab[3598]++
								return !m.verbose
//line /snap/go/10455/src/internal/bisect/bisect.go:345
	// _ = "end of CoverTab[3598]"
}

// ShouldEnable reports whether the change with the given id should be enabled.
func (m *Matcher) ShouldEnable(id uint64) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:349
	_go_fuzz_dep_.CoverTab[3599]++
								if m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:350
		_go_fuzz_dep_.CoverTab[526787]++
//line /snap/go/10455/src/internal/bisect/bisect.go:350
		_go_fuzz_dep_.CoverTab[3601]++
									return true
//line /snap/go/10455/src/internal/bisect/bisect.go:351
		// _ = "end of CoverTab[3601]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:352
		_go_fuzz_dep_.CoverTab[526788]++
//line /snap/go/10455/src/internal/bisect/bisect.go:352
		_go_fuzz_dep_.CoverTab[3602]++
//line /snap/go/10455/src/internal/bisect/bisect.go:352
		// _ = "end of CoverTab[3602]"
//line /snap/go/10455/src/internal/bisect/bisect.go:352
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:352
	// _ = "end of CoverTab[3599]"
//line /snap/go/10455/src/internal/bisect/bisect.go:352
	_go_fuzz_dep_.CoverTab[3600]++
								return m.matchResult(id) == m.enable
//line /snap/go/10455/src/internal/bisect/bisect.go:353
	// _ = "end of CoverTab[3600]"
}

// ShouldPrint reports whether to print identifying information about the change with the given id.
func (m *Matcher) ShouldPrint(id uint64) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:357
	_go_fuzz_dep_.CoverTab[3603]++
								if m == nil || func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:358
		_go_fuzz_dep_.CoverTab[3605]++
//line /snap/go/10455/src/internal/bisect/bisect.go:358
		return m.quiet
//line /snap/go/10455/src/internal/bisect/bisect.go:358
		// _ = "end of CoverTab[3605]"
//line /snap/go/10455/src/internal/bisect/bisect.go:358
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:358
		_go_fuzz_dep_.CoverTab[526789]++
//line /snap/go/10455/src/internal/bisect/bisect.go:358
		_go_fuzz_dep_.CoverTab[3606]++
									return false
//line /snap/go/10455/src/internal/bisect/bisect.go:359
		// _ = "end of CoverTab[3606]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:360
		_go_fuzz_dep_.CoverTab[526790]++
//line /snap/go/10455/src/internal/bisect/bisect.go:360
		_go_fuzz_dep_.CoverTab[3607]++
//line /snap/go/10455/src/internal/bisect/bisect.go:360
		// _ = "end of CoverTab[3607]"
//line /snap/go/10455/src/internal/bisect/bisect.go:360
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:360
	// _ = "end of CoverTab[3603]"
//line /snap/go/10455/src/internal/bisect/bisect.go:360
	_go_fuzz_dep_.CoverTab[3604]++
								return m.matchResult(id)
//line /snap/go/10455/src/internal/bisect/bisect.go:361
	// _ = "end of CoverTab[3604]"
}

// matchResult returns the result from the first condition that matches id.
func (m *Matcher) matchResult(id uint64) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:365
	_go_fuzz_dep_.CoverTab[3608]++
//line /snap/go/10455/src/internal/bisect/bisect.go:365
	_go_fuzz_dep_.CoverTab[786604] = 0
								for i := len(m.list) - 1; i >= 0; i-- {
//line /snap/go/10455/src/internal/bisect/bisect.go:366
		if _go_fuzz_dep_.CoverTab[786604] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:366
			_go_fuzz_dep_.CoverTab[526876]++
//line /snap/go/10455/src/internal/bisect/bisect.go:366
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:366
			_go_fuzz_dep_.CoverTab[526877]++
//line /snap/go/10455/src/internal/bisect/bisect.go:366
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:366
		_go_fuzz_dep_.CoverTab[786604] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:366
		_go_fuzz_dep_.CoverTab[3610]++
									c := &m.list[i]
									if id&c.mask == c.bits {
//line /snap/go/10455/src/internal/bisect/bisect.go:368
			_go_fuzz_dep_.CoverTab[526791]++
//line /snap/go/10455/src/internal/bisect/bisect.go:368
			_go_fuzz_dep_.CoverTab[3611]++
										return c.result
//line /snap/go/10455/src/internal/bisect/bisect.go:369
			// _ = "end of CoverTab[3611]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:370
			_go_fuzz_dep_.CoverTab[526792]++
//line /snap/go/10455/src/internal/bisect/bisect.go:370
			_go_fuzz_dep_.CoverTab[3612]++
//line /snap/go/10455/src/internal/bisect/bisect.go:370
			// _ = "end of CoverTab[3612]"
//line /snap/go/10455/src/internal/bisect/bisect.go:370
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:370
		// _ = "end of CoverTab[3610]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:371
	if _go_fuzz_dep_.CoverTab[786604] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:371
		_go_fuzz_dep_.CoverTab[526878]++
//line /snap/go/10455/src/internal/bisect/bisect.go:371
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:371
		_go_fuzz_dep_.CoverTab[526879]++
//line /snap/go/10455/src/internal/bisect/bisect.go:371
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:371
	// _ = "end of CoverTab[3608]"
//line /snap/go/10455/src/internal/bisect/bisect.go:371
	_go_fuzz_dep_.CoverTab[3609]++
								return false
//line /snap/go/10455/src/internal/bisect/bisect.go:372
	// _ = "end of CoverTab[3609]"
}

// FileLine reports whether the change identified by file and line should be enabled.
//line /snap/go/10455/src/internal/bisect/bisect.go:375
// If the change should be printed, FileLine prints a one-line report to w.
//line /snap/go/10455/src/internal/bisect/bisect.go:377
func (m *Matcher) FileLine(w Writer, file string, line int) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:377
	_go_fuzz_dep_.CoverTab[3613]++
								if m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:378
		_go_fuzz_dep_.CoverTab[526793]++
//line /snap/go/10455/src/internal/bisect/bisect.go:378
		_go_fuzz_dep_.CoverTab[3615]++
									return true
//line /snap/go/10455/src/internal/bisect/bisect.go:379
		// _ = "end of CoverTab[3615]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:380
		_go_fuzz_dep_.CoverTab[526794]++
//line /snap/go/10455/src/internal/bisect/bisect.go:380
		_go_fuzz_dep_.CoverTab[3616]++
//line /snap/go/10455/src/internal/bisect/bisect.go:380
		// _ = "end of CoverTab[3616]"
//line /snap/go/10455/src/internal/bisect/bisect.go:380
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:380
	// _ = "end of CoverTab[3613]"
//line /snap/go/10455/src/internal/bisect/bisect.go:380
	_go_fuzz_dep_.CoverTab[3614]++
								return m.fileLine(w, file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:381
	// _ = "end of CoverTab[3614]"
}

// fileLine does the real work for FileLine.
//line /snap/go/10455/src/internal/bisect/bisect.go:384
// This lets FileLine's body handle m == nil and potentially be inlined.
//line /snap/go/10455/src/internal/bisect/bisect.go:386
func (m *Matcher) fileLine(w Writer, file string, line int) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:386
	_go_fuzz_dep_.CoverTab[3617]++
								h := Hash(file, line)
								if m.ShouldPrint(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:388
		_go_fuzz_dep_.CoverTab[526795]++
//line /snap/go/10455/src/internal/bisect/bisect.go:388
		_go_fuzz_dep_.CoverTab[3619]++
									if m.MarkerOnly() {
//line /snap/go/10455/src/internal/bisect/bisect.go:389
			_go_fuzz_dep_.CoverTab[526797]++
//line /snap/go/10455/src/internal/bisect/bisect.go:389
			_go_fuzz_dep_.CoverTab[3620]++
										PrintMarker(w, h)
//line /snap/go/10455/src/internal/bisect/bisect.go:390
			// _ = "end of CoverTab[3620]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:391
			_go_fuzz_dep_.CoverTab[526798]++
//line /snap/go/10455/src/internal/bisect/bisect.go:391
			_go_fuzz_dep_.CoverTab[3621]++
										printFileLine(w, h, file, line)
//line /snap/go/10455/src/internal/bisect/bisect.go:392
			// _ = "end of CoverTab[3621]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:393
		// _ = "end of CoverTab[3619]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:394
		_go_fuzz_dep_.CoverTab[526796]++
//line /snap/go/10455/src/internal/bisect/bisect.go:394
		_go_fuzz_dep_.CoverTab[3622]++
//line /snap/go/10455/src/internal/bisect/bisect.go:394
		// _ = "end of CoverTab[3622]"
//line /snap/go/10455/src/internal/bisect/bisect.go:394
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:394
	// _ = "end of CoverTab[3617]"
//line /snap/go/10455/src/internal/bisect/bisect.go:394
	_go_fuzz_dep_.CoverTab[3618]++
								return m.ShouldEnable(h)
//line /snap/go/10455/src/internal/bisect/bisect.go:395
	// _ = "end of CoverTab[3618]"
}

// printFileLine prints a non-marker-only report for file:line to w.
func printFileLine(w Writer, h uint64, file string, line int) error {
//line /snap/go/10455/src/internal/bisect/bisect.go:399
	_go_fuzz_dep_.CoverTab[3623]++
								const markerLen = 40	// overestimate
								b := make([]byte, 0, markerLen+len(file)+24)
								b = AppendMarker(b, h)
								b = appendFileLine(b, file, line)
								b = append(b, '\n')
								_, err := w.Write(b)
								return err
//line /snap/go/10455/src/internal/bisect/bisect.go:406
	// _ = "end of CoverTab[3623]"
}

// appendFileLine appends file:line to dst, returning the extended slice.
func appendFileLine(dst []byte, file string, line int) []byte {
//line /snap/go/10455/src/internal/bisect/bisect.go:410
	_go_fuzz_dep_.CoverTab[3624]++
								dst = append(dst, file...)
								dst = append(dst, ':')
								u := uint(line)
								if line < 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:414
		_go_fuzz_dep_.CoverTab[526799]++
//line /snap/go/10455/src/internal/bisect/bisect.go:414
		_go_fuzz_dep_.CoverTab[3627]++
									dst = append(dst, '-')
									u = -u
//line /snap/go/10455/src/internal/bisect/bisect.go:416
		// _ = "end of CoverTab[3627]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:417
		_go_fuzz_dep_.CoverTab[526800]++
//line /snap/go/10455/src/internal/bisect/bisect.go:417
		_go_fuzz_dep_.CoverTab[3628]++
//line /snap/go/10455/src/internal/bisect/bisect.go:417
		// _ = "end of CoverTab[3628]"
//line /snap/go/10455/src/internal/bisect/bisect.go:417
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:417
	// _ = "end of CoverTab[3624]"
//line /snap/go/10455/src/internal/bisect/bisect.go:417
	_go_fuzz_dep_.CoverTab[3625]++
								var buf [24]byte
								i := len(buf)
//line /snap/go/10455/src/internal/bisect/bisect.go:419
	_go_fuzz_dep_.CoverTab[786605] = 0
								for i == len(buf) || func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		_go_fuzz_dep_.CoverTab[3629]++
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		return u > 0
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		// _ = "end of CoverTab[3629]"
//line /snap/go/10455/src/internal/bisect/bisect.go:420
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		if _go_fuzz_dep_.CoverTab[786605] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:420
			_go_fuzz_dep_.CoverTab[526880]++
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:420
			_go_fuzz_dep_.CoverTab[526881]++
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		_go_fuzz_dep_.CoverTab[786605] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:420
		_go_fuzz_dep_.CoverTab[3630]++
									i--
									buf[i] = '0' + byte(u%10)
									u /= 10
//line /snap/go/10455/src/internal/bisect/bisect.go:423
		// _ = "end of CoverTab[3630]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:424
	if _go_fuzz_dep_.CoverTab[786605] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:424
		_go_fuzz_dep_.CoverTab[526882]++
//line /snap/go/10455/src/internal/bisect/bisect.go:424
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:424
		_go_fuzz_dep_.CoverTab[526883]++
//line /snap/go/10455/src/internal/bisect/bisect.go:424
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:424
	// _ = "end of CoverTab[3625]"
//line /snap/go/10455/src/internal/bisect/bisect.go:424
	_go_fuzz_dep_.CoverTab[3626]++
								dst = append(dst, buf[i:]...)
								return dst
//line /snap/go/10455/src/internal/bisect/bisect.go:426
	// _ = "end of CoverTab[3626]"
}

// MatchStack assigns the current call stack a change ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:429
// If the stack should be printed, MatchStack prints it.
//line /snap/go/10455/src/internal/bisect/bisect.go:429
// Then MatchStack reports whether a change at the current call stack should be enabled.
//line /snap/go/10455/src/internal/bisect/bisect.go:432
func (m *Matcher) Stack(w Writer) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:432
	_go_fuzz_dep_.CoverTab[3631]++
								if m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:433
		_go_fuzz_dep_.CoverTab[526801]++
//line /snap/go/10455/src/internal/bisect/bisect.go:433
		_go_fuzz_dep_.CoverTab[3633]++
									return true
//line /snap/go/10455/src/internal/bisect/bisect.go:434
		// _ = "end of CoverTab[3633]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:435
		_go_fuzz_dep_.CoverTab[526802]++
//line /snap/go/10455/src/internal/bisect/bisect.go:435
		_go_fuzz_dep_.CoverTab[3634]++
//line /snap/go/10455/src/internal/bisect/bisect.go:435
		// _ = "end of CoverTab[3634]"
//line /snap/go/10455/src/internal/bisect/bisect.go:435
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:435
	// _ = "end of CoverTab[3631]"
//line /snap/go/10455/src/internal/bisect/bisect.go:435
	_go_fuzz_dep_.CoverTab[3632]++
								return m.stack(w)
//line /snap/go/10455/src/internal/bisect/bisect.go:436
	// _ = "end of CoverTab[3632]"
}

// stack does the real work for Stack.
//line /snap/go/10455/src/internal/bisect/bisect.go:439
// This lets stack's body handle m == nil and potentially be inlined.
//line /snap/go/10455/src/internal/bisect/bisect.go:441
func (m *Matcher) stack(w Writer) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:441
	_go_fuzz_dep_.CoverTab[3635]++
								const maxStack = 16
								var stk [maxStack]uintptr
								n := runtime.Callers(2, stk[:])

								if n <= 1 {
//line /snap/go/10455/src/internal/bisect/bisect.go:446
		_go_fuzz_dep_.CoverTab[526803]++
//line /snap/go/10455/src/internal/bisect/bisect.go:446
		_go_fuzz_dep_.CoverTab[3639]++
									return false
//line /snap/go/10455/src/internal/bisect/bisect.go:447
		// _ = "end of CoverTab[3639]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:448
		_go_fuzz_dep_.CoverTab[526804]++
//line /snap/go/10455/src/internal/bisect/bisect.go:448
		_go_fuzz_dep_.CoverTab[3640]++
//line /snap/go/10455/src/internal/bisect/bisect.go:448
		// _ = "end of CoverTab[3640]"
//line /snap/go/10455/src/internal/bisect/bisect.go:448
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:448
	// _ = "end of CoverTab[3635]"
//line /snap/go/10455/src/internal/bisect/bisect.go:448
	_go_fuzz_dep_.CoverTab[3636]++

								base := stk[0]
//line /snap/go/10455/src/internal/bisect/bisect.go:450
	_go_fuzz_dep_.CoverTab[786606] = 0

								for i := range stk[:n] {
//line /snap/go/10455/src/internal/bisect/bisect.go:452
		if _go_fuzz_dep_.CoverTab[786606] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:452
			_go_fuzz_dep_.CoverTab[526884]++
//line /snap/go/10455/src/internal/bisect/bisect.go:452
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:452
			_go_fuzz_dep_.CoverTab[526885]++
//line /snap/go/10455/src/internal/bisect/bisect.go:452
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:452
		_go_fuzz_dep_.CoverTab[786606] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:452
		_go_fuzz_dep_.CoverTab[3641]++
									stk[i] -= base
//line /snap/go/10455/src/internal/bisect/bisect.go:453
		// _ = "end of CoverTab[3641]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:454
	if _go_fuzz_dep_.CoverTab[786606] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:454
		_go_fuzz_dep_.CoverTab[526886]++
//line /snap/go/10455/src/internal/bisect/bisect.go:454
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:454
		_go_fuzz_dep_.CoverTab[526887]++
//line /snap/go/10455/src/internal/bisect/bisect.go:454
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:454
	// _ = "end of CoverTab[3636]"
//line /snap/go/10455/src/internal/bisect/bisect.go:454
	_go_fuzz_dep_.CoverTab[3637]++

								h := Hash(stk[:n])
								if m.ShouldPrint(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:457
		_go_fuzz_dep_.CoverTab[526805]++
//line /snap/go/10455/src/internal/bisect/bisect.go:457
		_go_fuzz_dep_.CoverTab[3642]++
									var d *dedup
//line /snap/go/10455/src/internal/bisect/bisect.go:458
		_go_fuzz_dep_.CoverTab[786607] = 0
									for {
//line /snap/go/10455/src/internal/bisect/bisect.go:459
			if _go_fuzz_dep_.CoverTab[786607] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:459
				_go_fuzz_dep_.CoverTab[526888]++
//line /snap/go/10455/src/internal/bisect/bisect.go:459
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:459
				_go_fuzz_dep_.CoverTab[526889]++
//line /snap/go/10455/src/internal/bisect/bisect.go:459
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:459
			_go_fuzz_dep_.CoverTab[786607] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:459
			_go_fuzz_dep_.CoverTab[3644]++
										d = m.dedup.Load()
										if d != nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:461
				_go_fuzz_dep_.CoverTab[526807]++
//line /snap/go/10455/src/internal/bisect/bisect.go:461
				_go_fuzz_dep_.CoverTab[3646]++
											break
//line /snap/go/10455/src/internal/bisect/bisect.go:462
				// _ = "end of CoverTab[3646]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:463
				_go_fuzz_dep_.CoverTab[526808]++
//line /snap/go/10455/src/internal/bisect/bisect.go:463
				_go_fuzz_dep_.CoverTab[3647]++
//line /snap/go/10455/src/internal/bisect/bisect.go:463
				// _ = "end of CoverTab[3647]"
//line /snap/go/10455/src/internal/bisect/bisect.go:463
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:463
			// _ = "end of CoverTab[3644]"
//line /snap/go/10455/src/internal/bisect/bisect.go:463
			_go_fuzz_dep_.CoverTab[3645]++
										d = new(dedup)
										if m.dedup.CompareAndSwap(nil, d) {
//line /snap/go/10455/src/internal/bisect/bisect.go:465
				_go_fuzz_dep_.CoverTab[526809]++
//line /snap/go/10455/src/internal/bisect/bisect.go:465
				_go_fuzz_dep_.CoverTab[3648]++
											break
//line /snap/go/10455/src/internal/bisect/bisect.go:466
				// _ = "end of CoverTab[3648]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:467
				_go_fuzz_dep_.CoverTab[526810]++
//line /snap/go/10455/src/internal/bisect/bisect.go:467
				_go_fuzz_dep_.CoverTab[3649]++
//line /snap/go/10455/src/internal/bisect/bisect.go:467
				// _ = "end of CoverTab[3649]"
//line /snap/go/10455/src/internal/bisect/bisect.go:467
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:467
			// _ = "end of CoverTab[3645]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:468
		// _ = "end of CoverTab[3642]"
//line /snap/go/10455/src/internal/bisect/bisect.go:468
		_go_fuzz_dep_.CoverTab[3643]++

									if m.MarkerOnly() {
//line /snap/go/10455/src/internal/bisect/bisect.go:470
			_go_fuzz_dep_.CoverTab[526811]++
//line /snap/go/10455/src/internal/bisect/bisect.go:470
			_go_fuzz_dep_.CoverTab[3650]++
										if !d.seenLossy(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:471
				_go_fuzz_dep_.CoverTab[526813]++
//line /snap/go/10455/src/internal/bisect/bisect.go:471
				_go_fuzz_dep_.CoverTab[3651]++
											PrintMarker(w, h)
//line /snap/go/10455/src/internal/bisect/bisect.go:472
				// _ = "end of CoverTab[3651]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:473
				_go_fuzz_dep_.CoverTab[526814]++
//line /snap/go/10455/src/internal/bisect/bisect.go:473
				_go_fuzz_dep_.CoverTab[3652]++
//line /snap/go/10455/src/internal/bisect/bisect.go:473
				// _ = "end of CoverTab[3652]"
//line /snap/go/10455/src/internal/bisect/bisect.go:473
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:473
			// _ = "end of CoverTab[3650]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:474
			_go_fuzz_dep_.CoverTab[526812]++
//line /snap/go/10455/src/internal/bisect/bisect.go:474
			_go_fuzz_dep_.CoverTab[3653]++
										if !d.seen(h) {
//line /snap/go/10455/src/internal/bisect/bisect.go:475
				_go_fuzz_dep_.CoverTab[526815]++
//line /snap/go/10455/src/internal/bisect/bisect.go:475
				_go_fuzz_dep_.CoverTab[3654]++
//line /snap/go/10455/src/internal/bisect/bisect.go:475
				_go_fuzz_dep_.CoverTab[786608] = 0

											for i := range stk[:n] {
//line /snap/go/10455/src/internal/bisect/bisect.go:477
					if _go_fuzz_dep_.CoverTab[786608] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:477
						_go_fuzz_dep_.CoverTab[526892]++
//line /snap/go/10455/src/internal/bisect/bisect.go:477
					} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:477
						_go_fuzz_dep_.CoverTab[526893]++
//line /snap/go/10455/src/internal/bisect/bisect.go:477
					}
//line /snap/go/10455/src/internal/bisect/bisect.go:477
					_go_fuzz_dep_.CoverTab[786608] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:477
					_go_fuzz_dep_.CoverTab[3656]++
												stk[i] += base
//line /snap/go/10455/src/internal/bisect/bisect.go:478
					// _ = "end of CoverTab[3656]"
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:479
				if _go_fuzz_dep_.CoverTab[786608] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:479
					_go_fuzz_dep_.CoverTab[526894]++
//line /snap/go/10455/src/internal/bisect/bisect.go:479
				} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:479
					_go_fuzz_dep_.CoverTab[526895]++
//line /snap/go/10455/src/internal/bisect/bisect.go:479
				}
//line /snap/go/10455/src/internal/bisect/bisect.go:479
				// _ = "end of CoverTab[3654]"
//line /snap/go/10455/src/internal/bisect/bisect.go:479
				_go_fuzz_dep_.CoverTab[3655]++
											printStack(w, h, stk[1:n])
//line /snap/go/10455/src/internal/bisect/bisect.go:480
				// _ = "end of CoverTab[3655]"
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:481
				_go_fuzz_dep_.CoverTab[526816]++
//line /snap/go/10455/src/internal/bisect/bisect.go:481
				_go_fuzz_dep_.CoverTab[3657]++
//line /snap/go/10455/src/internal/bisect/bisect.go:481
				// _ = "end of CoverTab[3657]"
//line /snap/go/10455/src/internal/bisect/bisect.go:481
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:481
			// _ = "end of CoverTab[3653]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:482
		// _ = "end of CoverTab[3643]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:483
		_go_fuzz_dep_.CoverTab[526806]++
//line /snap/go/10455/src/internal/bisect/bisect.go:483
		_go_fuzz_dep_.CoverTab[3658]++
//line /snap/go/10455/src/internal/bisect/bisect.go:483
		// _ = "end of CoverTab[3658]"
//line /snap/go/10455/src/internal/bisect/bisect.go:483
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:483
	// _ = "end of CoverTab[3637]"
//line /snap/go/10455/src/internal/bisect/bisect.go:483
	_go_fuzz_dep_.CoverTab[3638]++
								return m.ShouldEnable(h)
//line /snap/go/10455/src/internal/bisect/bisect.go:484
	// _ = "end of CoverTab[3638]"

}

// Writer is the same interface as io.Writer.
//line /snap/go/10455/src/internal/bisect/bisect.go:488
// It is duplicated here to avoid importing io.
//line /snap/go/10455/src/internal/bisect/bisect.go:490
type Writer interface {
	Write([]byte) (int, error)
}

// PrintMarker prints to w a one-line report containing only the marker for h.
//line /snap/go/10455/src/internal/bisect/bisect.go:494
// It is appropriate to use when [Matcher.ShouldPrint] and [Matcher.MarkerOnly] both return true.
//line /snap/go/10455/src/internal/bisect/bisect.go:496
func PrintMarker(w Writer, h uint64) error {
//line /snap/go/10455/src/internal/bisect/bisect.go:496
	_go_fuzz_dep_.CoverTab[3659]++
								var buf [50]byte
								b := AppendMarker(buf[:], h)
								b = append(b, '\n')
								_, err := w.Write(b)
								return err
//line /snap/go/10455/src/internal/bisect/bisect.go:501
	// _ = "end of CoverTab[3659]"
}

// printStack prints to w a multi-line report containing a formatting of the call stack stk,
//line /snap/go/10455/src/internal/bisect/bisect.go:504
// with each line preceded by the marker for h.
//line /snap/go/10455/src/internal/bisect/bisect.go:506
func printStack(w Writer, h uint64, stk []uintptr) error {
//line /snap/go/10455/src/internal/bisect/bisect.go:506
	_go_fuzz_dep_.CoverTab[3660]++
								buf := make([]byte, 0, 2048)

								var prefixBuf [100]byte
								prefix := AppendMarker(prefixBuf[:0], h)

								frames := runtime.CallersFrames(stk)
//line /snap/go/10455/src/internal/bisect/bisect.go:512
	_go_fuzz_dep_.CoverTab[786609] = 0
								for {
//line /snap/go/10455/src/internal/bisect/bisect.go:513
		if _go_fuzz_dep_.CoverTab[786609] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:513
			_go_fuzz_dep_.CoverTab[526896]++
//line /snap/go/10455/src/internal/bisect/bisect.go:513
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:513
			_go_fuzz_dep_.CoverTab[526897]++
//line /snap/go/10455/src/internal/bisect/bisect.go:513
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:513
		_go_fuzz_dep_.CoverTab[786609] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:513
		_go_fuzz_dep_.CoverTab[3662]++
									f, more := frames.Next()
									buf = append(buf, prefix...)
									buf = append(buf, f.Func.Name()...)
									buf = append(buf, "()\n"...)
									buf = append(buf, prefix...)
									buf = append(buf, '\t')
									buf = appendFileLine(buf, f.File, f.Line)
									buf = append(buf, '\n')
									if !more {
//line /snap/go/10455/src/internal/bisect/bisect.go:522
			_go_fuzz_dep_.CoverTab[526817]++
//line /snap/go/10455/src/internal/bisect/bisect.go:522
			_go_fuzz_dep_.CoverTab[3663]++
										break
//line /snap/go/10455/src/internal/bisect/bisect.go:523
			// _ = "end of CoverTab[3663]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:524
			_go_fuzz_dep_.CoverTab[526818]++
//line /snap/go/10455/src/internal/bisect/bisect.go:524
			_go_fuzz_dep_.CoverTab[3664]++
//line /snap/go/10455/src/internal/bisect/bisect.go:524
			// _ = "end of CoverTab[3664]"
//line /snap/go/10455/src/internal/bisect/bisect.go:524
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:524
		// _ = "end of CoverTab[3662]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:525
	// _ = "end of CoverTab[3660]"
//line /snap/go/10455/src/internal/bisect/bisect.go:525
	_go_fuzz_dep_.CoverTab[3661]++
								buf = append(buf, prefix...)
								buf = append(buf, '\n')
								_, err := w.Write(buf)
								return err
//line /snap/go/10455/src/internal/bisect/bisect.go:529
	// _ = "end of CoverTab[3661]"
}

// Marker returns the match marker text to use on any line reporting details
//line /snap/go/10455/src/internal/bisect/bisect.go:532
// about a match of the given ID.
//line /snap/go/10455/src/internal/bisect/bisect.go:532
// It always returns the hexadecimal format.
//line /snap/go/10455/src/internal/bisect/bisect.go:535
func Marker(id uint64) string {
//line /snap/go/10455/src/internal/bisect/bisect.go:535
	_go_fuzz_dep_.CoverTab[3665]++
								return string(AppendMarker(nil, id))
//line /snap/go/10455/src/internal/bisect/bisect.go:536
	// _ = "end of CoverTab[3665]"
}

// AppendMarker is like [Marker] but appends the marker to dst.
func AppendMarker(dst []byte, id uint64) []byte {
//line /snap/go/10455/src/internal/bisect/bisect.go:540
	_go_fuzz_dep_.CoverTab[3666]++
								const prefix = "[bisect-match 0x"
								var buf [len(prefix) + 16 + 1]byte
								copy(buf[:], prefix)
//line /snap/go/10455/src/internal/bisect/bisect.go:543
	_go_fuzz_dep_.CoverTab[786610] = 0
								for i := 0; i < 16; i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:544
		if _go_fuzz_dep_.CoverTab[786610] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:544
			_go_fuzz_dep_.CoverTab[526900]++
//line /snap/go/10455/src/internal/bisect/bisect.go:544
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:544
			_go_fuzz_dep_.CoverTab[526901]++
//line /snap/go/10455/src/internal/bisect/bisect.go:544
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:544
		_go_fuzz_dep_.CoverTab[786610] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:544
		_go_fuzz_dep_.CoverTab[3668]++
									buf[len(prefix)+i] = "0123456789abcdef"[id>>60]
									id <<= 4
//line /snap/go/10455/src/internal/bisect/bisect.go:546
		// _ = "end of CoverTab[3668]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:547
	if _go_fuzz_dep_.CoverTab[786610] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:547
		_go_fuzz_dep_.CoverTab[526902]++
//line /snap/go/10455/src/internal/bisect/bisect.go:547
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:547
		_go_fuzz_dep_.CoverTab[526903]++
//line /snap/go/10455/src/internal/bisect/bisect.go:547
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:547
	// _ = "end of CoverTab[3666]"
//line /snap/go/10455/src/internal/bisect/bisect.go:547
	_go_fuzz_dep_.CoverTab[3667]++
								buf[len(prefix)+16] = ']'
								return append(dst, buf[:]...)
//line /snap/go/10455/src/internal/bisect/bisect.go:549
	// _ = "end of CoverTab[3667]"
}

// CutMarker finds the first match marker in line and removes it,
//line /snap/go/10455/src/internal/bisect/bisect.go:552
// returning the shortened line (with the marker removed),
//line /snap/go/10455/src/internal/bisect/bisect.go:552
// the ID from the match marker,
//line /snap/go/10455/src/internal/bisect/bisect.go:552
// and whether a marker was found at all.
//line /snap/go/10455/src/internal/bisect/bisect.go:552
// If there is no marker, CutMarker returns line, 0, false.
//line /snap/go/10455/src/internal/bisect/bisect.go:557
func CutMarker(line string) (short string, id uint64, ok bool) {
//line /snap/go/10455/src/internal/bisect/bisect.go:557
	_go_fuzz_dep_.CoverTab[3669]++

								prefix := "[bisect-match "
								i := 0
//line /snap/go/10455/src/internal/bisect/bisect.go:560
	_go_fuzz_dep_.CoverTab[786611] = 0
								for ; ; i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:561
		if _go_fuzz_dep_.CoverTab[786611] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:561
			_go_fuzz_dep_.CoverTab[526904]++
//line /snap/go/10455/src/internal/bisect/bisect.go:561
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:561
			_go_fuzz_dep_.CoverTab[526905]++
//line /snap/go/10455/src/internal/bisect/bisect.go:561
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:561
		_go_fuzz_dep_.CoverTab[786611] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:561
		_go_fuzz_dep_.CoverTab[3675]++
									if i >= len(line)-len(prefix) {
//line /snap/go/10455/src/internal/bisect/bisect.go:562
			_go_fuzz_dep_.CoverTab[526819]++
//line /snap/go/10455/src/internal/bisect/bisect.go:562
			_go_fuzz_dep_.CoverTab[3677]++
										return line, 0, false
//line /snap/go/10455/src/internal/bisect/bisect.go:563
			// _ = "end of CoverTab[3677]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:564
			_go_fuzz_dep_.CoverTab[526820]++
//line /snap/go/10455/src/internal/bisect/bisect.go:564
			_go_fuzz_dep_.CoverTab[3678]++
//line /snap/go/10455/src/internal/bisect/bisect.go:564
			// _ = "end of CoverTab[3678]"
//line /snap/go/10455/src/internal/bisect/bisect.go:564
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:564
		// _ = "end of CoverTab[3675]"
//line /snap/go/10455/src/internal/bisect/bisect.go:564
		_go_fuzz_dep_.CoverTab[3676]++
									if line[i] == '[' && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:565
			_go_fuzz_dep_.CoverTab[3679]++
//line /snap/go/10455/src/internal/bisect/bisect.go:565
			return line[i:i+len(prefix)] == prefix
//line /snap/go/10455/src/internal/bisect/bisect.go:565
			// _ = "end of CoverTab[3679]"
//line /snap/go/10455/src/internal/bisect/bisect.go:565
		}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:565
			_go_fuzz_dep_.CoverTab[526821]++
//line /snap/go/10455/src/internal/bisect/bisect.go:565
			_go_fuzz_dep_.CoverTab[3680]++
										break
//line /snap/go/10455/src/internal/bisect/bisect.go:566
			// _ = "end of CoverTab[3680]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:567
			_go_fuzz_dep_.CoverTab[526822]++
//line /snap/go/10455/src/internal/bisect/bisect.go:567
			_go_fuzz_dep_.CoverTab[3681]++
//line /snap/go/10455/src/internal/bisect/bisect.go:567
			// _ = "end of CoverTab[3681]"
//line /snap/go/10455/src/internal/bisect/bisect.go:567
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:567
		// _ = "end of CoverTab[3676]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:568
	// _ = "end of CoverTab[3669]"
//line /snap/go/10455/src/internal/bisect/bisect.go:568
	_go_fuzz_dep_.CoverTab[3670]++

//line /snap/go/10455/src/internal/bisect/bisect.go:571
	j := i + len(prefix)
//line /snap/go/10455/src/internal/bisect/bisect.go:571
	_go_fuzz_dep_.CoverTab[786612] = 0
								for j < len(line) && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		_go_fuzz_dep_.CoverTab[3682]++
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		return line[j] != ']'
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		// _ = "end of CoverTab[3682]"
//line /snap/go/10455/src/internal/bisect/bisect.go:572
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		if _go_fuzz_dep_.CoverTab[786612] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:572
			_go_fuzz_dep_.CoverTab[526908]++
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:572
			_go_fuzz_dep_.CoverTab[526909]++
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		_go_fuzz_dep_.CoverTab[786612] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:572
		_go_fuzz_dep_.CoverTab[3683]++
									j++
//line /snap/go/10455/src/internal/bisect/bisect.go:573
		// _ = "end of CoverTab[3683]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:574
	if _go_fuzz_dep_.CoverTab[786612] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:574
		_go_fuzz_dep_.CoverTab[526910]++
//line /snap/go/10455/src/internal/bisect/bisect.go:574
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:574
		_go_fuzz_dep_.CoverTab[526911]++
//line /snap/go/10455/src/internal/bisect/bisect.go:574
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:574
	// _ = "end of CoverTab[3670]"
//line /snap/go/10455/src/internal/bisect/bisect.go:574
	_go_fuzz_dep_.CoverTab[3671]++
								if j >= len(line) {
//line /snap/go/10455/src/internal/bisect/bisect.go:575
		_go_fuzz_dep_.CoverTab[526823]++
//line /snap/go/10455/src/internal/bisect/bisect.go:575
		_go_fuzz_dep_.CoverTab[3684]++
									return line, 0, false
//line /snap/go/10455/src/internal/bisect/bisect.go:576
		// _ = "end of CoverTab[3684]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:577
		_go_fuzz_dep_.CoverTab[526824]++
//line /snap/go/10455/src/internal/bisect/bisect.go:577
		_go_fuzz_dep_.CoverTab[3685]++
//line /snap/go/10455/src/internal/bisect/bisect.go:577
		// _ = "end of CoverTab[3685]"
//line /snap/go/10455/src/internal/bisect/bisect.go:577
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:577
	// _ = "end of CoverTab[3671]"
//line /snap/go/10455/src/internal/bisect/bisect.go:577
	_go_fuzz_dep_.CoverTab[3672]++

//line /snap/go/10455/src/internal/bisect/bisect.go:580
	idstr := line[i+len(prefix) : j]
	if len(idstr) >= 3 && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:581
		_go_fuzz_dep_.CoverTab[3686]++
//line /snap/go/10455/src/internal/bisect/bisect.go:581
		return idstr[:2] == "0x"
//line /snap/go/10455/src/internal/bisect/bisect.go:581
		// _ = "end of CoverTab[3686]"
//line /snap/go/10455/src/internal/bisect/bisect.go:581
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:581
		_go_fuzz_dep_.CoverTab[526825]++
//line /snap/go/10455/src/internal/bisect/bisect.go:581
		_go_fuzz_dep_.CoverTab[3687]++

									if len(idstr) > 2+16 {
//line /snap/go/10455/src/internal/bisect/bisect.go:583
			_go_fuzz_dep_.CoverTab[526827]++
//line /snap/go/10455/src/internal/bisect/bisect.go:583
			_go_fuzz_dep_.CoverTab[3689]++
										return line, 0, false
//line /snap/go/10455/src/internal/bisect/bisect.go:584
			// _ = "end of CoverTab[3689]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:585
			_go_fuzz_dep_.CoverTab[526828]++
//line /snap/go/10455/src/internal/bisect/bisect.go:585
			_go_fuzz_dep_.CoverTab[3690]++
//line /snap/go/10455/src/internal/bisect/bisect.go:585
			// _ = "end of CoverTab[3690]"
//line /snap/go/10455/src/internal/bisect/bisect.go:585
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:585
		// _ = "end of CoverTab[3687]"
//line /snap/go/10455/src/internal/bisect/bisect.go:585
		_go_fuzz_dep_.CoverTab[3688]++
//line /snap/go/10455/src/internal/bisect/bisect.go:585
		_go_fuzz_dep_.CoverTab[786613] = 0
									for i := 2; i < len(idstr); i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:586
			if _go_fuzz_dep_.CoverTab[786613] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:586
				_go_fuzz_dep_.CoverTab[526912]++
//line /snap/go/10455/src/internal/bisect/bisect.go:586
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:586
				_go_fuzz_dep_.CoverTab[526913]++
//line /snap/go/10455/src/internal/bisect/bisect.go:586
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:586
			_go_fuzz_dep_.CoverTab[786613] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:586
			_go_fuzz_dep_.CoverTab[3691]++
										id <<= 4
										switch c := idstr[i]; {
			case '0' <= c && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:589
				_go_fuzz_dep_.CoverTab[3696]++
//line /snap/go/10455/src/internal/bisect/bisect.go:589
				return c <= '9'
//line /snap/go/10455/src/internal/bisect/bisect.go:589
				// _ = "end of CoverTab[3696]"
//line /snap/go/10455/src/internal/bisect/bisect.go:589
			}():
//line /snap/go/10455/src/internal/bisect/bisect.go:589
				_go_fuzz_dep_.CoverTab[526829]++
//line /snap/go/10455/src/internal/bisect/bisect.go:589
				_go_fuzz_dep_.CoverTab[3692]++
											id |= uint64(c - '0')
//line /snap/go/10455/src/internal/bisect/bisect.go:590
				// _ = "end of CoverTab[3692]"
			case 'a' <= c && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:591
				_go_fuzz_dep_.CoverTab[3697]++
//line /snap/go/10455/src/internal/bisect/bisect.go:591
				return c <= 'f'
//line /snap/go/10455/src/internal/bisect/bisect.go:591
				// _ = "end of CoverTab[3697]"
//line /snap/go/10455/src/internal/bisect/bisect.go:591
			}():
//line /snap/go/10455/src/internal/bisect/bisect.go:591
				_go_fuzz_dep_.CoverTab[526830]++
//line /snap/go/10455/src/internal/bisect/bisect.go:591
				_go_fuzz_dep_.CoverTab[3693]++
											id |= uint64(c - 'a' + 10)
//line /snap/go/10455/src/internal/bisect/bisect.go:592
				// _ = "end of CoverTab[3693]"
			case 'A' <= c && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:593
				_go_fuzz_dep_.CoverTab[3698]++
//line /snap/go/10455/src/internal/bisect/bisect.go:593
				return c <= 'F'
//line /snap/go/10455/src/internal/bisect/bisect.go:593
				// _ = "end of CoverTab[3698]"
//line /snap/go/10455/src/internal/bisect/bisect.go:593
			}():
//line /snap/go/10455/src/internal/bisect/bisect.go:593
				_go_fuzz_dep_.CoverTab[526831]++
//line /snap/go/10455/src/internal/bisect/bisect.go:593
				_go_fuzz_dep_.CoverTab[3694]++
											id |= uint64(c - 'A' + 10)
//line /snap/go/10455/src/internal/bisect/bisect.go:594
				// _ = "end of CoverTab[3694]"
//line /snap/go/10455/src/internal/bisect/bisect.go:594
			default:
//line /snap/go/10455/src/internal/bisect/bisect.go:594
				_go_fuzz_dep_.CoverTab[526832]++
//line /snap/go/10455/src/internal/bisect/bisect.go:594
				_go_fuzz_dep_.CoverTab[3695]++
//line /snap/go/10455/src/internal/bisect/bisect.go:594
				// _ = "end of CoverTab[3695]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:595
			// _ = "end of CoverTab[3691]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:596
		if _go_fuzz_dep_.CoverTab[786613] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:596
			_go_fuzz_dep_.CoverTab[526914]++
//line /snap/go/10455/src/internal/bisect/bisect.go:596
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:596
			_go_fuzz_dep_.CoverTab[526915]++
//line /snap/go/10455/src/internal/bisect/bisect.go:596
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:596
		// _ = "end of CoverTab[3688]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:597
		_go_fuzz_dep_.CoverTab[526826]++
//line /snap/go/10455/src/internal/bisect/bisect.go:597
		_go_fuzz_dep_.CoverTab[3699]++
									if idstr == "" || func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:598
			_go_fuzz_dep_.CoverTab[3701]++
//line /snap/go/10455/src/internal/bisect/bisect.go:598
			return len(idstr) > 64
//line /snap/go/10455/src/internal/bisect/bisect.go:598
			// _ = "end of CoverTab[3701]"
//line /snap/go/10455/src/internal/bisect/bisect.go:598
		}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:598
			_go_fuzz_dep_.CoverTab[526833]++
//line /snap/go/10455/src/internal/bisect/bisect.go:598
			_go_fuzz_dep_.CoverTab[3702]++
										return line, 0, false
//line /snap/go/10455/src/internal/bisect/bisect.go:599
			// _ = "end of CoverTab[3702]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:600
			_go_fuzz_dep_.CoverTab[526834]++
//line /snap/go/10455/src/internal/bisect/bisect.go:600
			_go_fuzz_dep_.CoverTab[3703]++
//line /snap/go/10455/src/internal/bisect/bisect.go:600
			// _ = "end of CoverTab[3703]"
//line /snap/go/10455/src/internal/bisect/bisect.go:600
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:600
		// _ = "end of CoverTab[3699]"
//line /snap/go/10455/src/internal/bisect/bisect.go:600
		_go_fuzz_dep_.CoverTab[3700]++
//line /snap/go/10455/src/internal/bisect/bisect.go:600
		_go_fuzz_dep_.CoverTab[786614] = 0

									for i := 0; i < len(idstr); i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:602
			if _go_fuzz_dep_.CoverTab[786614] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:602
				_go_fuzz_dep_.CoverTab[526916]++
//line /snap/go/10455/src/internal/bisect/bisect.go:602
			} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:602
				_go_fuzz_dep_.CoverTab[526917]++
//line /snap/go/10455/src/internal/bisect/bisect.go:602
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:602
			_go_fuzz_dep_.CoverTab[786614] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:602
			_go_fuzz_dep_.CoverTab[3704]++
										id <<= 1
										switch c := idstr[i]; c {
			default:
//line /snap/go/10455/src/internal/bisect/bisect.go:605
				_go_fuzz_dep_.CoverTab[526835]++
//line /snap/go/10455/src/internal/bisect/bisect.go:605
				_go_fuzz_dep_.CoverTab[3705]++
											return line, 0, false
//line /snap/go/10455/src/internal/bisect/bisect.go:606
				// _ = "end of CoverTab[3705]"
			case '0', '1':
//line /snap/go/10455/src/internal/bisect/bisect.go:607
				_go_fuzz_dep_.CoverTab[526836]++
//line /snap/go/10455/src/internal/bisect/bisect.go:607
				_go_fuzz_dep_.CoverTab[3706]++
											id |= uint64(c - '0')
//line /snap/go/10455/src/internal/bisect/bisect.go:608
				// _ = "end of CoverTab[3706]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:609
			// _ = "end of CoverTab[3704]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:610
		if _go_fuzz_dep_.CoverTab[786614] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:610
			_go_fuzz_dep_.CoverTab[526918]++
//line /snap/go/10455/src/internal/bisect/bisect.go:610
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:610
			_go_fuzz_dep_.CoverTab[526919]++
//line /snap/go/10455/src/internal/bisect/bisect.go:610
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:610
		// _ = "end of CoverTab[3700]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:611
	// _ = "end of CoverTab[3672]"
//line /snap/go/10455/src/internal/bisect/bisect.go:611
	_go_fuzz_dep_.CoverTab[3673]++

//line /snap/go/10455/src/internal/bisect/bisect.go:616
	j++
	if i > 0 && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:617
		_go_fuzz_dep_.CoverTab[3707]++
//line /snap/go/10455/src/internal/bisect/bisect.go:617
		return line[i-1] == ' '
//line /snap/go/10455/src/internal/bisect/bisect.go:617
		// _ = "end of CoverTab[3707]"
//line /snap/go/10455/src/internal/bisect/bisect.go:617
	}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:617
		_go_fuzz_dep_.CoverTab[526837]++
//line /snap/go/10455/src/internal/bisect/bisect.go:617
		_go_fuzz_dep_.CoverTab[3708]++
									i--
//line /snap/go/10455/src/internal/bisect/bisect.go:618
		// _ = "end of CoverTab[3708]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:619
		_go_fuzz_dep_.CoverTab[526838]++
//line /snap/go/10455/src/internal/bisect/bisect.go:619
		_go_fuzz_dep_.CoverTab[3709]++
//line /snap/go/10455/src/internal/bisect/bisect.go:619
		if j < len(line) && func() bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:619
			_go_fuzz_dep_.CoverTab[3710]++
//line /snap/go/10455/src/internal/bisect/bisect.go:619
			return line[j] == ' '
//line /snap/go/10455/src/internal/bisect/bisect.go:619
			// _ = "end of CoverTab[3710]"
//line /snap/go/10455/src/internal/bisect/bisect.go:619
		}() {
//line /snap/go/10455/src/internal/bisect/bisect.go:619
			_go_fuzz_dep_.CoverTab[526839]++
//line /snap/go/10455/src/internal/bisect/bisect.go:619
			_go_fuzz_dep_.CoverTab[3711]++
										j++
//line /snap/go/10455/src/internal/bisect/bisect.go:620
			// _ = "end of CoverTab[3711]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:621
			_go_fuzz_dep_.CoverTab[526840]++
//line /snap/go/10455/src/internal/bisect/bisect.go:621
			_go_fuzz_dep_.CoverTab[3712]++
//line /snap/go/10455/src/internal/bisect/bisect.go:621
			// _ = "end of CoverTab[3712]"
//line /snap/go/10455/src/internal/bisect/bisect.go:621
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:621
		// _ = "end of CoverTab[3709]"
//line /snap/go/10455/src/internal/bisect/bisect.go:621
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:621
	// _ = "end of CoverTab[3673]"
//line /snap/go/10455/src/internal/bisect/bisect.go:621
	_go_fuzz_dep_.CoverTab[3674]++
								short = line[:i] + line[j:]
								return short, id, true
//line /snap/go/10455/src/internal/bisect/bisect.go:623
	// _ = "end of CoverTab[3674]"
}

// Hash computes a hash of the data arguments,
//line /snap/go/10455/src/internal/bisect/bisect.go:626
// each of which must be of type string, byte, int, uint, int32, uint32, int64, uint64, uintptr, or a slice of one of those types.
//line /snap/go/10455/src/internal/bisect/bisect.go:628
func Hash(data ...any) uint64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:628
	_go_fuzz_dep_.CoverTab[3713]++
								h := offset64
//line /snap/go/10455/src/internal/bisect/bisect.go:629
	_go_fuzz_dep_.CoverTab[786615] = 0
								for _, v := range data {
//line /snap/go/10455/src/internal/bisect/bisect.go:630
		if _go_fuzz_dep_.CoverTab[786615] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:630
			_go_fuzz_dep_.CoverTab[526920]++
//line /snap/go/10455/src/internal/bisect/bisect.go:630
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:630
			_go_fuzz_dep_.CoverTab[526921]++
//line /snap/go/10455/src/internal/bisect/bisect.go:630
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:630
		_go_fuzz_dep_.CoverTab[786615] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:630
		_go_fuzz_dep_.CoverTab[3715]++
									switch v := v.(type) {
		default:
//line /snap/go/10455/src/internal/bisect/bisect.go:632
			_go_fuzz_dep_.CoverTab[526841]++
//line /snap/go/10455/src/internal/bisect/bisect.go:632
			_go_fuzz_dep_.CoverTab[3716]++

//line /snap/go/10455/src/internal/bisect/bisect.go:638
			panic("bisect.Hash: unexpected argument type")
//line /snap/go/10455/src/internal/bisect/bisect.go:638
			// _ = "end of CoverTab[3716]"
		case string:
//line /snap/go/10455/src/internal/bisect/bisect.go:639
			_go_fuzz_dep_.CoverTab[526842]++
//line /snap/go/10455/src/internal/bisect/bisect.go:639
			_go_fuzz_dep_.CoverTab[3717]++
										h = fnvString(h, v)
//line /snap/go/10455/src/internal/bisect/bisect.go:640
			// _ = "end of CoverTab[3717]"
		case byte:
//line /snap/go/10455/src/internal/bisect/bisect.go:641
			_go_fuzz_dep_.CoverTab[526843]++
//line /snap/go/10455/src/internal/bisect/bisect.go:641
			_go_fuzz_dep_.CoverTab[3718]++
										h = fnv(h, v)
//line /snap/go/10455/src/internal/bisect/bisect.go:642
			// _ = "end of CoverTab[3718]"
		case int:
//line /snap/go/10455/src/internal/bisect/bisect.go:643
			_go_fuzz_dep_.CoverTab[526844]++
//line /snap/go/10455/src/internal/bisect/bisect.go:643
			_go_fuzz_dep_.CoverTab[3719]++
										h = fnvUint64(h, uint64(v))
//line /snap/go/10455/src/internal/bisect/bisect.go:644
			// _ = "end of CoverTab[3719]"
		case uint:
//line /snap/go/10455/src/internal/bisect/bisect.go:645
			_go_fuzz_dep_.CoverTab[526845]++
//line /snap/go/10455/src/internal/bisect/bisect.go:645
			_go_fuzz_dep_.CoverTab[3720]++
										h = fnvUint64(h, uint64(v))
//line /snap/go/10455/src/internal/bisect/bisect.go:646
			// _ = "end of CoverTab[3720]"
		case int32:
//line /snap/go/10455/src/internal/bisect/bisect.go:647
			_go_fuzz_dep_.CoverTab[526846]++
//line /snap/go/10455/src/internal/bisect/bisect.go:647
			_go_fuzz_dep_.CoverTab[3721]++
										h = fnvUint32(h, uint32(v))
//line /snap/go/10455/src/internal/bisect/bisect.go:648
			// _ = "end of CoverTab[3721]"
		case uint32:
//line /snap/go/10455/src/internal/bisect/bisect.go:649
			_go_fuzz_dep_.CoverTab[526847]++
//line /snap/go/10455/src/internal/bisect/bisect.go:649
			_go_fuzz_dep_.CoverTab[3722]++
										h = fnvUint32(h, v)
//line /snap/go/10455/src/internal/bisect/bisect.go:650
			// _ = "end of CoverTab[3722]"
		case int64:
//line /snap/go/10455/src/internal/bisect/bisect.go:651
			_go_fuzz_dep_.CoverTab[526848]++
//line /snap/go/10455/src/internal/bisect/bisect.go:651
			_go_fuzz_dep_.CoverTab[3723]++
										h = fnvUint64(h, uint64(v))
//line /snap/go/10455/src/internal/bisect/bisect.go:652
			// _ = "end of CoverTab[3723]"
		case uint64:
//line /snap/go/10455/src/internal/bisect/bisect.go:653
			_go_fuzz_dep_.CoverTab[526849]++
//line /snap/go/10455/src/internal/bisect/bisect.go:653
			_go_fuzz_dep_.CoverTab[3724]++
										h = fnvUint64(h, v)
//line /snap/go/10455/src/internal/bisect/bisect.go:654
			// _ = "end of CoverTab[3724]"
		case uintptr:
//line /snap/go/10455/src/internal/bisect/bisect.go:655
			_go_fuzz_dep_.CoverTab[526850]++
//line /snap/go/10455/src/internal/bisect/bisect.go:655
			_go_fuzz_dep_.CoverTab[3725]++
										h = fnvUint64(h, uint64(v))
//line /snap/go/10455/src/internal/bisect/bisect.go:656
			// _ = "end of CoverTab[3725]"
		case []string:
//line /snap/go/10455/src/internal/bisect/bisect.go:657
			_go_fuzz_dep_.CoverTab[526851]++
//line /snap/go/10455/src/internal/bisect/bisect.go:657
			_go_fuzz_dep_.CoverTab[3726]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:658
				_go_fuzz_dep_.CoverTab[3735]++
											h = fnvString(h, x)
//line /snap/go/10455/src/internal/bisect/bisect.go:659
				// _ = "end of CoverTab[3735]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:660
			// _ = "end of CoverTab[3726]"
		case []byte:
//line /snap/go/10455/src/internal/bisect/bisect.go:661
			_go_fuzz_dep_.CoverTab[526852]++
//line /snap/go/10455/src/internal/bisect/bisect.go:661
			_go_fuzz_dep_.CoverTab[3727]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:662
				_go_fuzz_dep_.CoverTab[3736]++
											h = fnv(h, x)
//line /snap/go/10455/src/internal/bisect/bisect.go:663
				// _ = "end of CoverTab[3736]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:664
			// _ = "end of CoverTab[3727]"
		case []int:
//line /snap/go/10455/src/internal/bisect/bisect.go:665
			_go_fuzz_dep_.CoverTab[526853]++
//line /snap/go/10455/src/internal/bisect/bisect.go:665
			_go_fuzz_dep_.CoverTab[3728]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:666
				_go_fuzz_dep_.CoverTab[3737]++
											h = fnvUint64(h, uint64(x))
//line /snap/go/10455/src/internal/bisect/bisect.go:667
				// _ = "end of CoverTab[3737]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:668
			// _ = "end of CoverTab[3728]"
		case []uint:
//line /snap/go/10455/src/internal/bisect/bisect.go:669
			_go_fuzz_dep_.CoverTab[526854]++
//line /snap/go/10455/src/internal/bisect/bisect.go:669
			_go_fuzz_dep_.CoverTab[3729]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:670
				_go_fuzz_dep_.CoverTab[3738]++
											h = fnvUint64(h, uint64(x))
//line /snap/go/10455/src/internal/bisect/bisect.go:671
				// _ = "end of CoverTab[3738]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:672
			// _ = "end of CoverTab[3729]"
		case []int32:
//line /snap/go/10455/src/internal/bisect/bisect.go:673
			_go_fuzz_dep_.CoverTab[526855]++
//line /snap/go/10455/src/internal/bisect/bisect.go:673
			_go_fuzz_dep_.CoverTab[3730]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:674
				_go_fuzz_dep_.CoverTab[3739]++
											h = fnvUint32(h, uint32(x))
//line /snap/go/10455/src/internal/bisect/bisect.go:675
				// _ = "end of CoverTab[3739]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:676
			// _ = "end of CoverTab[3730]"
		case []uint32:
//line /snap/go/10455/src/internal/bisect/bisect.go:677
			_go_fuzz_dep_.CoverTab[526856]++
//line /snap/go/10455/src/internal/bisect/bisect.go:677
			_go_fuzz_dep_.CoverTab[3731]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:678
				_go_fuzz_dep_.CoverTab[3740]++
											h = fnvUint32(h, x)
//line /snap/go/10455/src/internal/bisect/bisect.go:679
				// _ = "end of CoverTab[3740]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:680
			// _ = "end of CoverTab[3731]"
		case []int64:
//line /snap/go/10455/src/internal/bisect/bisect.go:681
			_go_fuzz_dep_.CoverTab[526857]++
//line /snap/go/10455/src/internal/bisect/bisect.go:681
			_go_fuzz_dep_.CoverTab[3732]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:682
				_go_fuzz_dep_.CoverTab[3741]++
											h = fnvUint64(h, uint64(x))
//line /snap/go/10455/src/internal/bisect/bisect.go:683
				// _ = "end of CoverTab[3741]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:684
			// _ = "end of CoverTab[3732]"
		case []uint64:
//line /snap/go/10455/src/internal/bisect/bisect.go:685
			_go_fuzz_dep_.CoverTab[526858]++
//line /snap/go/10455/src/internal/bisect/bisect.go:685
			_go_fuzz_dep_.CoverTab[3733]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:686
				_go_fuzz_dep_.CoverTab[3742]++
											h = fnvUint64(h, x)
//line /snap/go/10455/src/internal/bisect/bisect.go:687
				// _ = "end of CoverTab[3742]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:688
			// _ = "end of CoverTab[3733]"
		case []uintptr:
//line /snap/go/10455/src/internal/bisect/bisect.go:689
			_go_fuzz_dep_.CoverTab[526859]++
//line /snap/go/10455/src/internal/bisect/bisect.go:689
			_go_fuzz_dep_.CoverTab[3734]++
										for _, x := range v {
//line /snap/go/10455/src/internal/bisect/bisect.go:690
				_go_fuzz_dep_.CoverTab[3743]++
											h = fnvUint64(h, uint64(x))
//line /snap/go/10455/src/internal/bisect/bisect.go:691
				// _ = "end of CoverTab[3743]"
			}
//line /snap/go/10455/src/internal/bisect/bisect.go:692
			// _ = "end of CoverTab[3734]"
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:693
		// _ = "end of CoverTab[3715]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:694
	if _go_fuzz_dep_.CoverTab[786615] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:694
		_go_fuzz_dep_.CoverTab[526922]++
//line /snap/go/10455/src/internal/bisect/bisect.go:694
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:694
		_go_fuzz_dep_.CoverTab[526923]++
//line /snap/go/10455/src/internal/bisect/bisect.go:694
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:694
	// _ = "end of CoverTab[3713]"
//line /snap/go/10455/src/internal/bisect/bisect.go:694
	_go_fuzz_dep_.CoverTab[3714]++
								return h
//line /snap/go/10455/src/internal/bisect/bisect.go:695
	// _ = "end of CoverTab[3714]"
}

//line /snap/go/10455/src/internal/bisect/bisect.go:700
// parseError is a trivial error implementation,
//line /snap/go/10455/src/internal/bisect/bisect.go:700
// defined here to avoid importing errors.
//line /snap/go/10455/src/internal/bisect/bisect.go:702
type parseError struct{ text string }

func (e *parseError) Error() string {
//line /snap/go/10455/src/internal/bisect/bisect.go:704
	_go_fuzz_dep_.CoverTab[3744]++
//line /snap/go/10455/src/internal/bisect/bisect.go:704
	return e.text
//line /snap/go/10455/src/internal/bisect/bisect.go:704
	// _ = "end of CoverTab[3744]"
//line /snap/go/10455/src/internal/bisect/bisect.go:704
}

//line /snap/go/10455/src/internal/bisect/bisect.go:710
const (
	offset64	uint64	= 14695981039346656037
	prime64		uint64	= 1099511628211
)

func fnv(h uint64, x byte) uint64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:715
	_go_fuzz_dep_.CoverTab[3745]++
								h ^= uint64(x)
								h *= prime64
								return h
//line /snap/go/10455/src/internal/bisect/bisect.go:718
	// _ = "end of CoverTab[3745]"
}

func fnvString(h uint64, x string) uint64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:721
	_go_fuzz_dep_.CoverTab[3746]++
//line /snap/go/10455/src/internal/bisect/bisect.go:721
	_go_fuzz_dep_.CoverTab[786616] = 0
								for i := 0; i < len(x); i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:722
		if _go_fuzz_dep_.CoverTab[786616] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:722
			_go_fuzz_dep_.CoverTab[526924]++
//line /snap/go/10455/src/internal/bisect/bisect.go:722
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:722
			_go_fuzz_dep_.CoverTab[526925]++
//line /snap/go/10455/src/internal/bisect/bisect.go:722
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:722
		_go_fuzz_dep_.CoverTab[786616] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:722
		_go_fuzz_dep_.CoverTab[3748]++
									h ^= uint64(x[i])
									h *= prime64
//line /snap/go/10455/src/internal/bisect/bisect.go:724
		// _ = "end of CoverTab[3748]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:725
	if _go_fuzz_dep_.CoverTab[786616] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:725
		_go_fuzz_dep_.CoverTab[526926]++
//line /snap/go/10455/src/internal/bisect/bisect.go:725
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:725
		_go_fuzz_dep_.CoverTab[526927]++
//line /snap/go/10455/src/internal/bisect/bisect.go:725
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:725
	// _ = "end of CoverTab[3746]"
//line /snap/go/10455/src/internal/bisect/bisect.go:725
	_go_fuzz_dep_.CoverTab[3747]++
								return h
//line /snap/go/10455/src/internal/bisect/bisect.go:726
	// _ = "end of CoverTab[3747]"
}

func fnvUint64(h uint64, x uint64) uint64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:729
	_go_fuzz_dep_.CoverTab[3749]++
//line /snap/go/10455/src/internal/bisect/bisect.go:729
	_go_fuzz_dep_.CoverTab[786617] = 0
								for i := 0; i < 8; i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:730
		if _go_fuzz_dep_.CoverTab[786617] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:730
			_go_fuzz_dep_.CoverTab[526928]++
//line /snap/go/10455/src/internal/bisect/bisect.go:730
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:730
			_go_fuzz_dep_.CoverTab[526929]++
//line /snap/go/10455/src/internal/bisect/bisect.go:730
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:730
		_go_fuzz_dep_.CoverTab[786617] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:730
		_go_fuzz_dep_.CoverTab[3751]++
									h ^= uint64(x & 0xFF)
									x >>= 8
									h *= prime64
//line /snap/go/10455/src/internal/bisect/bisect.go:733
		// _ = "end of CoverTab[3751]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:734
	if _go_fuzz_dep_.CoverTab[786617] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:734
		_go_fuzz_dep_.CoverTab[526930]++
//line /snap/go/10455/src/internal/bisect/bisect.go:734
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:734
		_go_fuzz_dep_.CoverTab[526931]++
//line /snap/go/10455/src/internal/bisect/bisect.go:734
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:734
	// _ = "end of CoverTab[3749]"
//line /snap/go/10455/src/internal/bisect/bisect.go:734
	_go_fuzz_dep_.CoverTab[3750]++
								return h
//line /snap/go/10455/src/internal/bisect/bisect.go:735
	// _ = "end of CoverTab[3750]"
}

func fnvUint32(h uint64, x uint32) uint64 {
//line /snap/go/10455/src/internal/bisect/bisect.go:738
	_go_fuzz_dep_.CoverTab[3752]++
//line /snap/go/10455/src/internal/bisect/bisect.go:738
	_go_fuzz_dep_.CoverTab[786618] = 0
								for i := 0; i < 4; i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:739
		if _go_fuzz_dep_.CoverTab[786618] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:739
			_go_fuzz_dep_.CoverTab[526932]++
//line /snap/go/10455/src/internal/bisect/bisect.go:739
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:739
			_go_fuzz_dep_.CoverTab[526933]++
//line /snap/go/10455/src/internal/bisect/bisect.go:739
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:739
		_go_fuzz_dep_.CoverTab[786618] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:739
		_go_fuzz_dep_.CoverTab[3754]++
									h ^= uint64(x & 0xFF)
									x >>= 8
									h *= prime64
//line /snap/go/10455/src/internal/bisect/bisect.go:742
		// _ = "end of CoverTab[3754]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:743
	if _go_fuzz_dep_.CoverTab[786618] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:743
		_go_fuzz_dep_.CoverTab[526934]++
//line /snap/go/10455/src/internal/bisect/bisect.go:743
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:743
		_go_fuzz_dep_.CoverTab[526935]++
//line /snap/go/10455/src/internal/bisect/bisect.go:743
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:743
	// _ = "end of CoverTab[3752]"
//line /snap/go/10455/src/internal/bisect/bisect.go:743
	_go_fuzz_dep_.CoverTab[3753]++
								return h
//line /snap/go/10455/src/internal/bisect/bisect.go:744
	// _ = "end of CoverTab[3753]"
}

// A dedup is a deduplicator for call stacks, so that we only print
//line /snap/go/10455/src/internal/bisect/bisect.go:747
// a report for new call stacks, not for call stacks we've already
//line /snap/go/10455/src/internal/bisect/bisect.go:747
// reported.
//line /snap/go/10455/src/internal/bisect/bisect.go:747
//
//line /snap/go/10455/src/internal/bisect/bisect.go:747
// It has two modes: an approximate but lock-free mode that
//line /snap/go/10455/src/internal/bisect/bisect.go:747
// may still emit some duplicates, and a precise mode that uses
//line /snap/go/10455/src/internal/bisect/bisect.go:747
// a lock and never emits duplicates.
//line /snap/go/10455/src/internal/bisect/bisect.go:754
type dedup struct {
	// 128-entry 4-way, lossy cache for seenLossy
	recent	[128][4]uint64

	// complete history for seen
	mu	sync.Mutex
	m	map[uint64]bool
}

// seen records that h has now been seen and reports whether it was seen before.
//line /snap/go/10455/src/internal/bisect/bisect.go:763
// When seen returns false, the caller is expected to print a report for h.
//line /snap/go/10455/src/internal/bisect/bisect.go:765
func (d *dedup) seen(h uint64) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:765
	_go_fuzz_dep_.CoverTab[3755]++
								d.mu.Lock()
								if d.m == nil {
//line /snap/go/10455/src/internal/bisect/bisect.go:767
		_go_fuzz_dep_.CoverTab[526860]++
//line /snap/go/10455/src/internal/bisect/bisect.go:767
		_go_fuzz_dep_.CoverTab[3757]++
									d.m = make(map[uint64]bool)
//line /snap/go/10455/src/internal/bisect/bisect.go:768
		// _ = "end of CoverTab[3757]"
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:769
		_go_fuzz_dep_.CoverTab[526861]++
//line /snap/go/10455/src/internal/bisect/bisect.go:769
		_go_fuzz_dep_.CoverTab[3758]++
//line /snap/go/10455/src/internal/bisect/bisect.go:769
		// _ = "end of CoverTab[3758]"
//line /snap/go/10455/src/internal/bisect/bisect.go:769
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:769
	// _ = "end of CoverTab[3755]"
//line /snap/go/10455/src/internal/bisect/bisect.go:769
	_go_fuzz_dep_.CoverTab[3756]++
								seen := d.m[h]
								d.m[h] = true
								d.mu.Unlock()
								return seen
//line /snap/go/10455/src/internal/bisect/bisect.go:773
	// _ = "end of CoverTab[3756]"
}

// seenLossy is a variant of seen that avoids a lock by using a cache of recently seen hashes.
//line /snap/go/10455/src/internal/bisect/bisect.go:776
// Each cache entry is N-way set-associative: h can appear in any of the slots.
//line /snap/go/10455/src/internal/bisect/bisect.go:776
// If h does not appear in any of them, then it is inserted into a random slot,
//line /snap/go/10455/src/internal/bisect/bisect.go:776
// overwriting whatever was there before.
//line /snap/go/10455/src/internal/bisect/bisect.go:780
func (d *dedup) seenLossy(h uint64) bool {
//line /snap/go/10455/src/internal/bisect/bisect.go:780
	_go_fuzz_dep_.CoverTab[3759]++
								cache := &d.recent[uint(h)%uint(len(d.recent))]
//line /snap/go/10455/src/internal/bisect/bisect.go:781
	_go_fuzz_dep_.CoverTab[786619] = 0
								for i := 0; i < len(cache); i++ {
//line /snap/go/10455/src/internal/bisect/bisect.go:782
		if _go_fuzz_dep_.CoverTab[786619] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:782
			_go_fuzz_dep_.CoverTab[526936]++
//line /snap/go/10455/src/internal/bisect/bisect.go:782
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:782
			_go_fuzz_dep_.CoverTab[526937]++
//line /snap/go/10455/src/internal/bisect/bisect.go:782
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:782
		_go_fuzz_dep_.CoverTab[786619] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:782
		_go_fuzz_dep_.CoverTab[3762]++
									if atomic.LoadUint64(&cache[i]) == h {
//line /snap/go/10455/src/internal/bisect/bisect.go:783
			_go_fuzz_dep_.CoverTab[526862]++
//line /snap/go/10455/src/internal/bisect/bisect.go:783
			_go_fuzz_dep_.CoverTab[3763]++
										return true
//line /snap/go/10455/src/internal/bisect/bisect.go:784
			// _ = "end of CoverTab[3763]"
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:785
			_go_fuzz_dep_.CoverTab[526863]++
//line /snap/go/10455/src/internal/bisect/bisect.go:785
			_go_fuzz_dep_.CoverTab[3764]++
//line /snap/go/10455/src/internal/bisect/bisect.go:785
			// _ = "end of CoverTab[3764]"
//line /snap/go/10455/src/internal/bisect/bisect.go:785
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:785
		// _ = "end of CoverTab[3762]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:786
	if _go_fuzz_dep_.CoverTab[786619] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:786
		_go_fuzz_dep_.CoverTab[526938]++
//line /snap/go/10455/src/internal/bisect/bisect.go:786
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:786
		_go_fuzz_dep_.CoverTab[526939]++
//line /snap/go/10455/src/internal/bisect/bisect.go:786
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:786
	// _ = "end of CoverTab[3759]"
//line /snap/go/10455/src/internal/bisect/bisect.go:786
	_go_fuzz_dep_.CoverTab[3760]++

//line /snap/go/10455/src/internal/bisect/bisect.go:789
	ch := offset64
//line /snap/go/10455/src/internal/bisect/bisect.go:789
	_go_fuzz_dep_.CoverTab[786620] = 0
								for _, x := range cache {
//line /snap/go/10455/src/internal/bisect/bisect.go:790
		if _go_fuzz_dep_.CoverTab[786620] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:790
			_go_fuzz_dep_.CoverTab[526940]++
//line /snap/go/10455/src/internal/bisect/bisect.go:790
		} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:790
			_go_fuzz_dep_.CoverTab[526941]++
//line /snap/go/10455/src/internal/bisect/bisect.go:790
		}
//line /snap/go/10455/src/internal/bisect/bisect.go:790
		_go_fuzz_dep_.CoverTab[786620] = 1
//line /snap/go/10455/src/internal/bisect/bisect.go:790
		_go_fuzz_dep_.CoverTab[3765]++
									ch = fnvUint64(ch, x)
//line /snap/go/10455/src/internal/bisect/bisect.go:791
		// _ = "end of CoverTab[3765]"
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:792
	if _go_fuzz_dep_.CoverTab[786620] == 0 {
//line /snap/go/10455/src/internal/bisect/bisect.go:792
		_go_fuzz_dep_.CoverTab[526942]++
//line /snap/go/10455/src/internal/bisect/bisect.go:792
	} else {
//line /snap/go/10455/src/internal/bisect/bisect.go:792
		_go_fuzz_dep_.CoverTab[526943]++
//line /snap/go/10455/src/internal/bisect/bisect.go:792
	}
//line /snap/go/10455/src/internal/bisect/bisect.go:792
	// _ = "end of CoverTab[3760]"
//line /snap/go/10455/src/internal/bisect/bisect.go:792
	_go_fuzz_dep_.CoverTab[3761]++
								atomic.StoreUint64(&cache[uint(ch)%uint(len(cache))], h)
								return false
//line /snap/go/10455/src/internal/bisect/bisect.go:794
	// _ = "end of CoverTab[3761]"
}

//line /snap/go/10455/src/internal/bisect/bisect.go:795
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/internal/bisect/bisect.go:795
var _ = _go_fuzz_dep_.CoverTab
