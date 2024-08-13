// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/error.go:5
package template

//line /usr/local/go/src/html/template/error.go:5
import (
//line /usr/local/go/src/html/template/error.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/error.go:5
)
//line /usr/local/go/src/html/template/error.go:5
import (
//line /usr/local/go/src/html/template/error.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/error.go:5
)

import (
	"fmt"
	"text/template/parse"
)

// Error describes a problem encountered during template Escaping.
type Error struct {
	// ErrorCode describes the kind of error.
	ErrorCode	ErrorCode
	// Node is the node that caused the problem, if known.
	// If not nil, it overrides Name and Line.
	Node	parse.Node
	// Name is the name of the template in which the error was encountered.
	Name	string
	// Line is the line number of the error in the template source or 0.
	Line	int
	// Description is a human-readable description of the problem.
	Description	string
}

// ErrorCode is a code for a kind of error.
type ErrorCode int

// We define codes for each error that manifests while escaping templates, but
//line /usr/local/go/src/html/template/error.go:30
// escaped templates may also fail at runtime.
//line /usr/local/go/src/html/template/error.go:30
//
//line /usr/local/go/src/html/template/error.go:30
// Output: "ZgotmplZ"
//line /usr/local/go/src/html/template/error.go:30
// Example:
//line /usr/local/go/src/html/template/error.go:30
//
//line /usr/local/go/src/html/template/error.go:30
//	<img src="{{.X}}">
//line /usr/local/go/src/html/template/error.go:30
//	where {{.X}} evaluates to `javascript:...`
//line /usr/local/go/src/html/template/error.go:30
//
//line /usr/local/go/src/html/template/error.go:30
// Discussion:
//line /usr/local/go/src/html/template/error.go:30
//
//line /usr/local/go/src/html/template/error.go:30
//	"ZgotmplZ" is a special value that indicates that unsafe content reached a
//line /usr/local/go/src/html/template/error.go:30
//	CSS or URL context at runtime. The output of the example will be
//line /usr/local/go/src/html/template/error.go:30
//	  <img src="#ZgotmplZ">
//line /usr/local/go/src/html/template/error.go:30
//	If the data comes from a trusted source, use content types to exempt it
//line /usr/local/go/src/html/template/error.go:30
//	from filtering: URL(`javascript:...`).
//line /usr/local/go/src/html/template/error.go:46
const (
	// OK indicates the lack of an error.
	OK	ErrorCode	= iota

	// ErrAmbigContext: "... appears in an ambiguous context within a URL"
	// Example:
	//   <a href="
	//      {{if .C}}
	//        /path/
	//      {{else}}
	//        /search?q=
	//      {{end}}
	//      {{.X}}
	//   ">
	// Discussion:
	//   {{.X}} is in an ambiguous URL context since, depending on {{.C}},
	//  it may be either a URL suffix or a query parameter.
	//   Moving {{.X}} into the condition removes the ambiguity:
	//   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}">
	ErrAmbigContext

	// ErrBadHTML: "expected space, attr name, or end of tag, but got ...",
	//   "... in unquoted attr", "... in attribute name"
	// Example:
	//   <a href = /search?q=foo>
	//   <href=foo>
	//   <form na<e=...>
	//   <option selected<
	// Discussion:
	//   This is often due to a typo in an HTML element, but some runes
	//   are banned in tag names, attribute names, and unquoted attribute
	//   values because they can tickle parser ambiguities.
	//   Quoting all attributes is the best policy.
	ErrBadHTML

	// ErrBranchEnd: "{{if}} branches end in different contexts"
	// Example:
	//   {{if .C}}<a href="{{end}}{{.X}}
	// Discussion:
	//   Package html/template statically examines each path through an
	//   {{if}}, {{range}}, or {{with}} to escape any following pipelines.
	//   The example is ambiguous since {{.X}} might be an HTML text node,
	//   or a URL prefix in an HTML attribute. The context of {{.X}} is
	//   used to figure out how to escape it, but that context depends on
	//   the run-time value of {{.C}} which is not statically known.
	//
	//   The problem is usually something like missing quotes or angle
	//   brackets, or can be avoided by refactoring to put the two contexts
	//   into different branches of an if, range or with. If the problem
	//   is in a {{range}} over a collection that should never be empty,
	//   adding a dummy {{else}} can help.
	ErrBranchEnd

	// ErrEndContext: "... ends in a non-text context: ..."
	// Examples:
	//   <div
	//   <div title="no close quote>
	//   <script>f()
	// Discussion:
	//   Executed templates should produce a DocumentFragment of HTML.
	//   Templates that end without closing tags will trigger this error.
	//   Templates that should not be used in an HTML context or that
	//   produce incomplete Fragments should not be executed directly.
	//
	//   {{define "main"}} <script>{{template "helper"}}</script> {{end}}
	//   {{define "helper"}} document.write(' <div title=" ') {{end}}
	//
	//   "helper" does not produce a valid document fragment, so should
	//   not be Executed directly.
	ErrEndContext

	// ErrNoSuchTemplate: "no such template ..."
	// Examples:
	//   {{define "main"}}<div {{template "attrs"}}>{{end}}
	//   {{define "attrs"}}href="{{.URL}}"{{end}}
	// Discussion:
	//   Package html/template looks through template calls to compute the
	//   context.
	//   Here the {{.URL}} in "attrs" must be treated as a URL when called
	//   from "main", but you will get this error if "attrs" is not defined
	//   when "main" is parsed.
	ErrNoSuchTemplate

	// ErrOutputContext: "cannot compute output context for template ..."
	// Examples:
	//   {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
	// Discussion:
	//   A recursive template does not end in the same context in which it
	//   starts, and a reliable output context cannot be computed.
	//   Look for typos in the named template.
	//   If the template should not be called in the named start context,
	//   look for calls to that template in unexpected contexts.
	//   Maybe refactor recursive templates to not be recursive.
	ErrOutputContext

	// ErrPartialCharset: "unfinished JS regexp charset in ..."
	// Example:
	//     <script>var pattern = /foo[{{.Chars}}]/</script>
	// Discussion:
	//   Package html/template does not support interpolation into regular
	//   expression literal character sets.
	ErrPartialCharset

	// ErrPartialEscape: "unfinished escape sequence in ..."
	// Example:
	//   <script>alert("\{{.X}}")</script>
	// Discussion:
	//   Package html/template does not support actions following a
	//   backslash.
	//   This is usually an error and there are better solutions; for
	//   example
	//     <script>alert("{{.X}}")</script>
	//   should work, and if {{.X}} is a partial escape sequence such as
	//   "xA0", mark the whole sequence as safe content: JSStr(`\xA0`)
	ErrPartialEscape

	// ErrRangeLoopReentry: "on range loop re-entry: ..."
	// Example:
	//   <script>var x = [{{range .}}'{{.}},{{end}}]</script>
	// Discussion:
	//   If an iteration through a range would cause it to end in a
	//   different context than an earlier pass, there is no single context.
	//   In the example, there is missing a quote, so it is not clear
	//   whether {{.}} is meant to be inside a JS string or in a JS value
	//   context. The second iteration would produce something like
	//
	//     <script>var x = ['firstValue,'secondValue]</script>
	ErrRangeLoopReentry

	// ErrSlashAmbig: '/' could start a division or regexp.
	// Example:
	//   <script>
	//     {{if .C}}var x = 1{{end}}
	//     /-{{.N}}/i.test(x) ? doThis : doThat();
	//   </script>
	// Discussion:
	//   The example above could produce `var x = 1/-2/i.test(s)...`
	//   in which the first '/' is a mathematical division operator or it
	//   could produce `/-2/i.test(s)` in which the first '/' starts a
	//   regexp literal.
	//   Look for missing semicolons inside branches, and maybe add
	//   parentheses to make it clear which interpretation you intend.
	ErrSlashAmbig

	// ErrPredefinedEscaper: "predefined escaper ... disallowed in template"
	// Example:
	//   <div class={{. | html}}>Hello<div>
	// Discussion:
	//   Package html/template already contextually escapes all pipelines to
	//   produce HTML output safe against code injection. Manually escaping
	//   pipeline output using the predefined escapers "html" or "urlquery" is
	//   unnecessary, and may affect the correctness or safety of the escaped
	//   pipeline output in Go 1.8 and earlier.
	//
	//   In most cases, such as the given example, this error can be resolved by
	//   simply removing the predefined escaper from the pipeline and letting the
	//   contextual autoescaper handle the escaping of the pipeline. In other
	//   instances, where the predefined escaper occurs in the middle of a
	//   pipeline where subsequent commands expect escaped input, e.g.
	//     {{.X | html | makeALink}}
	//   where makeALink does
	//     return `<a href="`+input+`">link</a>`
	//   consider refactoring the surrounding template to make use of the
	//   contextual autoescaper, i.e.
	//     <a href="{{.X}}">link</a>
	//
	//   To ease migration to Go 1.9 and beyond, "html" and "urlquery" will
	//   continue to be allowed as the last command in a pipeline. However, if the
	//   pipeline occurs in an unquoted attribute value context, "html" is
	//   disallowed. Avoid using "html" and "urlquery" entirely in new templates.
	ErrPredefinedEscaper

	// errJSTmplLit: "... appears in a JS template literal"
	// Example:
	//     <script>var tmpl = `{{.Interp}`</script>
	// Discussion:
	//   Package html/template does not support actions inside of JS template
	//   literals.
	//
	// TODO(rolandshoemaker): we cannot add this as an exported error in a minor
	// release, since it is backwards incompatible with the other minor
	// releases. As such we need to leave it unexported, and then we'll add it
	// in the next major release.
	errJSTmplLit
)

func (e *Error) Error() string {
//line /usr/local/go/src/html/template/error.go:232
	_go_fuzz_dep_.CoverTab[30853]++
							switch {
	case e.Node != nil:
//line /usr/local/go/src/html/template/error.go:234
		_go_fuzz_dep_.CoverTab[30855]++
								loc, _ := (*parse.Tree)(nil).ErrorContext(e.Node)
								return fmt.Sprintf("html/template:%s: %s", loc, e.Description)
//line /usr/local/go/src/html/template/error.go:236
		// _ = "end of CoverTab[30855]"
	case e.Line != 0:
//line /usr/local/go/src/html/template/error.go:237
		_go_fuzz_dep_.CoverTab[30856]++
								return fmt.Sprintf("html/template:%s:%d: %s", e.Name, e.Line, e.Description)
//line /usr/local/go/src/html/template/error.go:238
		// _ = "end of CoverTab[30856]"
	case e.Name != "":
//line /usr/local/go/src/html/template/error.go:239
		_go_fuzz_dep_.CoverTab[30857]++
								return fmt.Sprintf("html/template:%s: %s", e.Name, e.Description)
//line /usr/local/go/src/html/template/error.go:240
		// _ = "end of CoverTab[30857]"
//line /usr/local/go/src/html/template/error.go:240
	default:
//line /usr/local/go/src/html/template/error.go:240
		_go_fuzz_dep_.CoverTab[30858]++
//line /usr/local/go/src/html/template/error.go:240
		// _ = "end of CoverTab[30858]"
	}
//line /usr/local/go/src/html/template/error.go:241
	// _ = "end of CoverTab[30853]"
//line /usr/local/go/src/html/template/error.go:241
	_go_fuzz_dep_.CoverTab[30854]++
							return "html/template: " + e.Description
//line /usr/local/go/src/html/template/error.go:242
	// _ = "end of CoverTab[30854]"
}

// errorf creates an error given a format string f and args.
//line /usr/local/go/src/html/template/error.go:245
// The template Name still needs to be supplied.
//line /usr/local/go/src/html/template/error.go:247
func errorf(k ErrorCode, node parse.Node, line int, f string, args ...any) *Error {
//line /usr/local/go/src/html/template/error.go:247
	_go_fuzz_dep_.CoverTab[30859]++
							return &Error{k, node, "", line, fmt.Sprintf(f, args...)}
//line /usr/local/go/src/html/template/error.go:248
	// _ = "end of CoverTab[30859]"
}

//line /usr/local/go/src/html/template/error.go:249
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/error.go:249
var _ = _go_fuzz_dep_.CoverTab
