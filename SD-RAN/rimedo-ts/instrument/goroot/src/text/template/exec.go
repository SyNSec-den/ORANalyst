// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/template/exec.go:5
package template

//line /usr/local/go/src/text/template/exec.go:5
import (
//line /usr/local/go/src/text/template/exec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/exec.go:5
)
//line /usr/local/go/src/text/template/exec.go:5
import (
//line /usr/local/go/src/text/template/exec.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/exec.go:5
)

import (
	"errors"
	"fmt"
	"internal/fmtsort"
	"io"
	"reflect"
	"runtime"
	"strings"
	"text/template/parse"
)

// maxExecDepth specifies the maximum stack depth of templates within
//line /usr/local/go/src/text/template/exec.go:18
// templates. This limit is only practically reached by accidentally
//line /usr/local/go/src/text/template/exec.go:18
// recursive template invocations. This limit allows us to return
//line /usr/local/go/src/text/template/exec.go:18
// an error instead of triggering a stack overflow.
//line /usr/local/go/src/text/template/exec.go:22
var maxExecDepth = initMaxExecDepth()

func initMaxExecDepth() int {
//line /usr/local/go/src/text/template/exec.go:24
	_go_fuzz_dep_.CoverTab[29710]++
							if runtime.GOARCH == "wasm" {
//line /usr/local/go/src/text/template/exec.go:25
		_go_fuzz_dep_.CoverTab[29712]++
								return 1000
//line /usr/local/go/src/text/template/exec.go:26
		// _ = "end of CoverTab[29712]"
	} else {
//line /usr/local/go/src/text/template/exec.go:27
		_go_fuzz_dep_.CoverTab[29713]++
//line /usr/local/go/src/text/template/exec.go:27
		// _ = "end of CoverTab[29713]"
//line /usr/local/go/src/text/template/exec.go:27
	}
//line /usr/local/go/src/text/template/exec.go:27
	// _ = "end of CoverTab[29710]"
//line /usr/local/go/src/text/template/exec.go:27
	_go_fuzz_dep_.CoverTab[29711]++
							return 100000
//line /usr/local/go/src/text/template/exec.go:28
	// _ = "end of CoverTab[29711]"
}

// state represents the state of an execution. It's not part of the
//line /usr/local/go/src/text/template/exec.go:31
// template so that multiple executions of the same template
//line /usr/local/go/src/text/template/exec.go:31
// can execute in parallel.
//line /usr/local/go/src/text/template/exec.go:34
type state struct {
	tmpl	*Template
	wr	io.Writer
	node	parse.Node	// current node, for errors
	vars	[]variable	// push-down stack of variable values.
	depth	int		// the height of the stack of executing templates.
}

// variable holds the dynamic value of a variable such as $, $x etc.
type variable struct {
	name	string
	value	reflect.Value
}

// push pushes a new variable on the stack.
func (s *state) push(name string, value reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:49
	_go_fuzz_dep_.CoverTab[29714]++
							s.vars = append(s.vars, variable{name, value})
//line /usr/local/go/src/text/template/exec.go:50
	// _ = "end of CoverTab[29714]"
}

// mark returns the length of the variable stack.
func (s *state) mark() int {
//line /usr/local/go/src/text/template/exec.go:54
	_go_fuzz_dep_.CoverTab[29715]++
							return len(s.vars)
//line /usr/local/go/src/text/template/exec.go:55
	// _ = "end of CoverTab[29715]"
}

// pop pops the variable stack up to the mark.
func (s *state) pop(mark int) {
//line /usr/local/go/src/text/template/exec.go:59
	_go_fuzz_dep_.CoverTab[29716]++
							s.vars = s.vars[0:mark]
//line /usr/local/go/src/text/template/exec.go:60
	// _ = "end of CoverTab[29716]"
}

// setVar overwrites the last declared variable with the given name.
//line /usr/local/go/src/text/template/exec.go:63
// Used by variable assignments.
//line /usr/local/go/src/text/template/exec.go:65
func (s *state) setVar(name string, value reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:65
	_go_fuzz_dep_.CoverTab[29717]++
							for i := s.mark() - 1; i >= 0; i-- {
//line /usr/local/go/src/text/template/exec.go:66
		_go_fuzz_dep_.CoverTab[29719]++
								if s.vars[i].name == name {
//line /usr/local/go/src/text/template/exec.go:67
			_go_fuzz_dep_.CoverTab[29720]++
									s.vars[i].value = value
									return
//line /usr/local/go/src/text/template/exec.go:69
			// _ = "end of CoverTab[29720]"
		} else {
//line /usr/local/go/src/text/template/exec.go:70
			_go_fuzz_dep_.CoverTab[29721]++
//line /usr/local/go/src/text/template/exec.go:70
			// _ = "end of CoverTab[29721]"
//line /usr/local/go/src/text/template/exec.go:70
		}
//line /usr/local/go/src/text/template/exec.go:70
		// _ = "end of CoverTab[29719]"
	}
//line /usr/local/go/src/text/template/exec.go:71
	// _ = "end of CoverTab[29717]"
//line /usr/local/go/src/text/template/exec.go:71
	_go_fuzz_dep_.CoverTab[29718]++
							s.errorf("undefined variable: %s", name)
//line /usr/local/go/src/text/template/exec.go:72
	// _ = "end of CoverTab[29718]"
}

// setTopVar overwrites the top-nth variable on the stack. Used by range iterations.
func (s *state) setTopVar(n int, value reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:76
	_go_fuzz_dep_.CoverTab[29722]++
							s.vars[len(s.vars)-n].value = value
//line /usr/local/go/src/text/template/exec.go:77
	// _ = "end of CoverTab[29722]"
}

// varValue returns the value of the named variable.
func (s *state) varValue(name string) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:81
	_go_fuzz_dep_.CoverTab[29723]++
							for i := s.mark() - 1; i >= 0; i-- {
//line /usr/local/go/src/text/template/exec.go:82
		_go_fuzz_dep_.CoverTab[29725]++
								if s.vars[i].name == name {
//line /usr/local/go/src/text/template/exec.go:83
			_go_fuzz_dep_.CoverTab[29726]++
									return s.vars[i].value
//line /usr/local/go/src/text/template/exec.go:84
			// _ = "end of CoverTab[29726]"
		} else {
//line /usr/local/go/src/text/template/exec.go:85
			_go_fuzz_dep_.CoverTab[29727]++
//line /usr/local/go/src/text/template/exec.go:85
			// _ = "end of CoverTab[29727]"
//line /usr/local/go/src/text/template/exec.go:85
		}
//line /usr/local/go/src/text/template/exec.go:85
		// _ = "end of CoverTab[29725]"
	}
//line /usr/local/go/src/text/template/exec.go:86
	// _ = "end of CoverTab[29723]"
//line /usr/local/go/src/text/template/exec.go:86
	_go_fuzz_dep_.CoverTab[29724]++
							s.errorf("undefined variable: %s", name)
							return zero
//line /usr/local/go/src/text/template/exec.go:88
	// _ = "end of CoverTab[29724]"
}

var zero reflect.Value

type missingValType struct{}

var missingVal = reflect.ValueOf(missingValType{})

var missingValReflectType = reflect.TypeOf(missingValType{})

func isMissing(v reflect.Value) bool {
//line /usr/local/go/src/text/template/exec.go:99
	_go_fuzz_dep_.CoverTab[29728]++
							return v.IsValid() && func() bool {
//line /usr/local/go/src/text/template/exec.go:100
		_go_fuzz_dep_.CoverTab[29729]++
//line /usr/local/go/src/text/template/exec.go:100
		return v.Type() == missingValReflectType
//line /usr/local/go/src/text/template/exec.go:100
		// _ = "end of CoverTab[29729]"
//line /usr/local/go/src/text/template/exec.go:100
	}()
//line /usr/local/go/src/text/template/exec.go:100
	// _ = "end of CoverTab[29728]"
}

// at marks the state to be on node n, for error reporting.
func (s *state) at(node parse.Node) {
//line /usr/local/go/src/text/template/exec.go:104
	_go_fuzz_dep_.CoverTab[29730]++
							s.node = node
//line /usr/local/go/src/text/template/exec.go:105
	// _ = "end of CoverTab[29730]"
}

// doublePercent returns the string with %'s replaced by %%, if necessary,
//line /usr/local/go/src/text/template/exec.go:108
// so it can be used safely inside a Printf format string.
//line /usr/local/go/src/text/template/exec.go:110
func doublePercent(str string) string {
//line /usr/local/go/src/text/template/exec.go:110
	_go_fuzz_dep_.CoverTab[29731]++
							return strings.ReplaceAll(str, "%", "%%")
//line /usr/local/go/src/text/template/exec.go:111
	// _ = "end of CoverTab[29731]"
}

//line /usr/local/go/src/text/template/exec.go:118
// ExecError is the custom error type returned when Execute has an
//line /usr/local/go/src/text/template/exec.go:118
// error evaluating its template. (If a write error occurs, the actual
//line /usr/local/go/src/text/template/exec.go:118
// error is returned; it will not be of type ExecError.)
//line /usr/local/go/src/text/template/exec.go:121
type ExecError struct {
	Name	string	// Name of template.
	Err	error	// Pre-formatted error.
}

func (e ExecError) Error() string {
//line /usr/local/go/src/text/template/exec.go:126
	_go_fuzz_dep_.CoverTab[29732]++
							return e.Err.Error()
//line /usr/local/go/src/text/template/exec.go:127
	// _ = "end of CoverTab[29732]"
}

func (e ExecError) Unwrap() error {
//line /usr/local/go/src/text/template/exec.go:130
	_go_fuzz_dep_.CoverTab[29733]++
							return e.Err
//line /usr/local/go/src/text/template/exec.go:131
	// _ = "end of CoverTab[29733]"
}

// errorf records an ExecError and terminates processing.
func (s *state) errorf(format string, args ...any) {
//line /usr/local/go/src/text/template/exec.go:135
	_go_fuzz_dep_.CoverTab[29734]++
							name := doublePercent(s.tmpl.Name())
							if s.node == nil {
//line /usr/local/go/src/text/template/exec.go:137
		_go_fuzz_dep_.CoverTab[29736]++
								format = fmt.Sprintf("template: %s: %s", name, format)
//line /usr/local/go/src/text/template/exec.go:138
		// _ = "end of CoverTab[29736]"
	} else {
//line /usr/local/go/src/text/template/exec.go:139
		_go_fuzz_dep_.CoverTab[29737]++
								location, context := s.tmpl.ErrorContext(s.node)
								format = fmt.Sprintf("template: %s: executing %q at <%s>: %s", location, name, doublePercent(context), format)
//line /usr/local/go/src/text/template/exec.go:141
		// _ = "end of CoverTab[29737]"
	}
//line /usr/local/go/src/text/template/exec.go:142
	// _ = "end of CoverTab[29734]"
//line /usr/local/go/src/text/template/exec.go:142
	_go_fuzz_dep_.CoverTab[29735]++
							panic(ExecError{
		Name:	s.tmpl.Name(),
		Err:	fmt.Errorf(format, args...),
	})
//line /usr/local/go/src/text/template/exec.go:146
	// _ = "end of CoverTab[29735]"
}

// writeError is the wrapper type used internally when Execute has an
//line /usr/local/go/src/text/template/exec.go:149
// error writing to its output. We strip the wrapper in errRecover.
//line /usr/local/go/src/text/template/exec.go:149
// Note that this is not an implementation of error, so it cannot escape
//line /usr/local/go/src/text/template/exec.go:149
// from the package as an error value.
//line /usr/local/go/src/text/template/exec.go:153
type writeError struct {
	Err error	// Original error.
}

func (s *state) writeError(err error) {
//line /usr/local/go/src/text/template/exec.go:157
	_go_fuzz_dep_.CoverTab[29738]++
							panic(writeError{
		Err: err,
	})
//line /usr/local/go/src/text/template/exec.go:160
	// _ = "end of CoverTab[29738]"
}

// errRecover is the handler that turns panics into returns from the top
//line /usr/local/go/src/text/template/exec.go:163
// level of Parse.
//line /usr/local/go/src/text/template/exec.go:165
func errRecover(errp *error) {
//line /usr/local/go/src/text/template/exec.go:165
	_go_fuzz_dep_.CoverTab[29739]++
							e := recover()
							if e != nil {
//line /usr/local/go/src/text/template/exec.go:167
		_go_fuzz_dep_.CoverTab[29740]++
								switch err := e.(type) {
		case runtime.Error:
//line /usr/local/go/src/text/template/exec.go:169
			_go_fuzz_dep_.CoverTab[29741]++
									panic(e)
//line /usr/local/go/src/text/template/exec.go:170
			// _ = "end of CoverTab[29741]"
		case writeError:
//line /usr/local/go/src/text/template/exec.go:171
			_go_fuzz_dep_.CoverTab[29742]++
									*errp = err.Err
//line /usr/local/go/src/text/template/exec.go:172
			// _ = "end of CoverTab[29742]"
		case ExecError:
//line /usr/local/go/src/text/template/exec.go:173
			_go_fuzz_dep_.CoverTab[29743]++
									*errp = err
//line /usr/local/go/src/text/template/exec.go:174
			// _ = "end of CoverTab[29743]"
		default:
//line /usr/local/go/src/text/template/exec.go:175
			_go_fuzz_dep_.CoverTab[29744]++
									panic(e)
//line /usr/local/go/src/text/template/exec.go:176
			// _ = "end of CoverTab[29744]"
		}
//line /usr/local/go/src/text/template/exec.go:177
		// _ = "end of CoverTab[29740]"
	} else {
//line /usr/local/go/src/text/template/exec.go:178
		_go_fuzz_dep_.CoverTab[29745]++
//line /usr/local/go/src/text/template/exec.go:178
		// _ = "end of CoverTab[29745]"
//line /usr/local/go/src/text/template/exec.go:178
	}
//line /usr/local/go/src/text/template/exec.go:178
	// _ = "end of CoverTab[29739]"
}

// ExecuteTemplate applies the template associated with t that has the given name
//line /usr/local/go/src/text/template/exec.go:181
// to the specified data object and writes the output to wr.
//line /usr/local/go/src/text/template/exec.go:181
// If an error occurs executing the template or writing its output,
//line /usr/local/go/src/text/template/exec.go:181
// execution stops, but partial results may already have been written to
//line /usr/local/go/src/text/template/exec.go:181
// the output writer.
//line /usr/local/go/src/text/template/exec.go:181
// A template may be executed safely in parallel, although if parallel
//line /usr/local/go/src/text/template/exec.go:181
// executions share a Writer the output may be interleaved.
//line /usr/local/go/src/text/template/exec.go:188
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error {
//line /usr/local/go/src/text/template/exec.go:188
	_go_fuzz_dep_.CoverTab[29746]++
							tmpl := t.Lookup(name)
							if tmpl == nil {
//line /usr/local/go/src/text/template/exec.go:190
		_go_fuzz_dep_.CoverTab[29748]++
								return fmt.Errorf("template: no template %q associated with template %q", name, t.name)
//line /usr/local/go/src/text/template/exec.go:191
		// _ = "end of CoverTab[29748]"
	} else {
//line /usr/local/go/src/text/template/exec.go:192
		_go_fuzz_dep_.CoverTab[29749]++
//line /usr/local/go/src/text/template/exec.go:192
		// _ = "end of CoverTab[29749]"
//line /usr/local/go/src/text/template/exec.go:192
	}
//line /usr/local/go/src/text/template/exec.go:192
	// _ = "end of CoverTab[29746]"
//line /usr/local/go/src/text/template/exec.go:192
	_go_fuzz_dep_.CoverTab[29747]++
							return tmpl.Execute(wr, data)
//line /usr/local/go/src/text/template/exec.go:193
	// _ = "end of CoverTab[29747]"
}

// Execute applies a parsed template to the specified data object,
//line /usr/local/go/src/text/template/exec.go:196
// and writes the output to wr.
//line /usr/local/go/src/text/template/exec.go:196
// If an error occurs executing the template or writing its output,
//line /usr/local/go/src/text/template/exec.go:196
// execution stops, but partial results may already have been written to
//line /usr/local/go/src/text/template/exec.go:196
// the output writer.
//line /usr/local/go/src/text/template/exec.go:196
// A template may be executed safely in parallel, although if parallel
//line /usr/local/go/src/text/template/exec.go:196
// executions share a Writer the output may be interleaved.
//line /usr/local/go/src/text/template/exec.go:196
//
//line /usr/local/go/src/text/template/exec.go:196
// If data is a reflect.Value, the template applies to the concrete
//line /usr/local/go/src/text/template/exec.go:196
// value that the reflect.Value holds, as in fmt.Print.
//line /usr/local/go/src/text/template/exec.go:206
func (t *Template) Execute(wr io.Writer, data any) error {
//line /usr/local/go/src/text/template/exec.go:206
	_go_fuzz_dep_.CoverTab[29750]++
							return t.execute(wr, data)
//line /usr/local/go/src/text/template/exec.go:207
	// _ = "end of CoverTab[29750]"
}

func (t *Template) execute(wr io.Writer, data any) (err error) {
//line /usr/local/go/src/text/template/exec.go:210
	_go_fuzz_dep_.CoverTab[29751]++
							defer errRecover(&err)
							value, ok := data.(reflect.Value)
							if !ok {
//line /usr/local/go/src/text/template/exec.go:213
		_go_fuzz_dep_.CoverTab[29754]++
								value = reflect.ValueOf(data)
//line /usr/local/go/src/text/template/exec.go:214
		// _ = "end of CoverTab[29754]"
	} else {
//line /usr/local/go/src/text/template/exec.go:215
		_go_fuzz_dep_.CoverTab[29755]++
//line /usr/local/go/src/text/template/exec.go:215
		// _ = "end of CoverTab[29755]"
//line /usr/local/go/src/text/template/exec.go:215
	}
//line /usr/local/go/src/text/template/exec.go:215
	// _ = "end of CoverTab[29751]"
//line /usr/local/go/src/text/template/exec.go:215
	_go_fuzz_dep_.CoverTab[29752]++
							state := &state{
		tmpl:	t,
		wr:	wr,
		vars:	[]variable{{"$", value}},
	}
	if t.Tree == nil || func() bool {
//line /usr/local/go/src/text/template/exec.go:221
		_go_fuzz_dep_.CoverTab[29756]++
//line /usr/local/go/src/text/template/exec.go:221
		return t.Root == nil
//line /usr/local/go/src/text/template/exec.go:221
		// _ = "end of CoverTab[29756]"
//line /usr/local/go/src/text/template/exec.go:221
	}() {
//line /usr/local/go/src/text/template/exec.go:221
		_go_fuzz_dep_.CoverTab[29757]++
								state.errorf("%q is an incomplete or empty template", t.Name())
//line /usr/local/go/src/text/template/exec.go:222
		// _ = "end of CoverTab[29757]"
	} else {
//line /usr/local/go/src/text/template/exec.go:223
		_go_fuzz_dep_.CoverTab[29758]++
//line /usr/local/go/src/text/template/exec.go:223
		// _ = "end of CoverTab[29758]"
//line /usr/local/go/src/text/template/exec.go:223
	}
//line /usr/local/go/src/text/template/exec.go:223
	// _ = "end of CoverTab[29752]"
//line /usr/local/go/src/text/template/exec.go:223
	_go_fuzz_dep_.CoverTab[29753]++
							state.walk(value, t.Root)
							return
//line /usr/local/go/src/text/template/exec.go:225
	// _ = "end of CoverTab[29753]"
}

// DefinedTemplates returns a string listing the defined templates,
//line /usr/local/go/src/text/template/exec.go:228
// prefixed by the string "; defined templates are: ". If there are none,
//line /usr/local/go/src/text/template/exec.go:228
// it returns the empty string. For generating an error message here
//line /usr/local/go/src/text/template/exec.go:228
// and in html/template.
//line /usr/local/go/src/text/template/exec.go:232
func (t *Template) DefinedTemplates() string {
//line /usr/local/go/src/text/template/exec.go:232
	_go_fuzz_dep_.CoverTab[29759]++
							if t.common == nil {
//line /usr/local/go/src/text/template/exec.go:233
		_go_fuzz_dep_.CoverTab[29762]++
								return ""
//line /usr/local/go/src/text/template/exec.go:234
		// _ = "end of CoverTab[29762]"
	} else {
//line /usr/local/go/src/text/template/exec.go:235
		_go_fuzz_dep_.CoverTab[29763]++
//line /usr/local/go/src/text/template/exec.go:235
		// _ = "end of CoverTab[29763]"
//line /usr/local/go/src/text/template/exec.go:235
	}
//line /usr/local/go/src/text/template/exec.go:235
	// _ = "end of CoverTab[29759]"
//line /usr/local/go/src/text/template/exec.go:235
	_go_fuzz_dep_.CoverTab[29760]++
							var b strings.Builder
							t.muTmpl.RLock()
							defer t.muTmpl.RUnlock()
							for name, tmpl := range t.tmpl {
//line /usr/local/go/src/text/template/exec.go:239
		_go_fuzz_dep_.CoverTab[29764]++
								if tmpl.Tree == nil || func() bool {
//line /usr/local/go/src/text/template/exec.go:240
			_go_fuzz_dep_.CoverTab[29767]++
//line /usr/local/go/src/text/template/exec.go:240
			return tmpl.Root == nil
//line /usr/local/go/src/text/template/exec.go:240
			// _ = "end of CoverTab[29767]"
//line /usr/local/go/src/text/template/exec.go:240
		}() {
//line /usr/local/go/src/text/template/exec.go:240
			_go_fuzz_dep_.CoverTab[29768]++
									continue
//line /usr/local/go/src/text/template/exec.go:241
			// _ = "end of CoverTab[29768]"
		} else {
//line /usr/local/go/src/text/template/exec.go:242
			_go_fuzz_dep_.CoverTab[29769]++
//line /usr/local/go/src/text/template/exec.go:242
			// _ = "end of CoverTab[29769]"
//line /usr/local/go/src/text/template/exec.go:242
		}
//line /usr/local/go/src/text/template/exec.go:242
		// _ = "end of CoverTab[29764]"
//line /usr/local/go/src/text/template/exec.go:242
		_go_fuzz_dep_.CoverTab[29765]++
								if b.Len() == 0 {
//line /usr/local/go/src/text/template/exec.go:243
			_go_fuzz_dep_.CoverTab[29770]++
									b.WriteString("; defined templates are: ")
//line /usr/local/go/src/text/template/exec.go:244
			// _ = "end of CoverTab[29770]"
		} else {
//line /usr/local/go/src/text/template/exec.go:245
			_go_fuzz_dep_.CoverTab[29771]++
									b.WriteString(", ")
//line /usr/local/go/src/text/template/exec.go:246
			// _ = "end of CoverTab[29771]"
		}
//line /usr/local/go/src/text/template/exec.go:247
		// _ = "end of CoverTab[29765]"
//line /usr/local/go/src/text/template/exec.go:247
		_go_fuzz_dep_.CoverTab[29766]++
								fmt.Fprintf(&b, "%q", name)
//line /usr/local/go/src/text/template/exec.go:248
		// _ = "end of CoverTab[29766]"
	}
//line /usr/local/go/src/text/template/exec.go:249
	// _ = "end of CoverTab[29760]"
//line /usr/local/go/src/text/template/exec.go:249
	_go_fuzz_dep_.CoverTab[29761]++
							return b.String()
//line /usr/local/go/src/text/template/exec.go:250
	// _ = "end of CoverTab[29761]"
}

// Sentinel errors for use with panic to signal early exits from range loops.
var (
	walkBreak	= errors.New("break")
	walkContinue	= errors.New("continue")
)

// Walk functions step through the major pieces of the template structure,
//line /usr/local/go/src/text/template/exec.go:259
// generating output as they go.
//line /usr/local/go/src/text/template/exec.go:261
func (s *state) walk(dot reflect.Value, node parse.Node) {
//line /usr/local/go/src/text/template/exec.go:261
	_go_fuzz_dep_.CoverTab[29772]++
							s.at(node)
							switch node := node.(type) {
	case *parse.ActionNode:
//line /usr/local/go/src/text/template/exec.go:264
		_go_fuzz_dep_.CoverTab[29773]++

//line /usr/local/go/src/text/template/exec.go:267
		val := s.evalPipeline(dot, node.Pipe)
		if len(node.Pipe.Decl) == 0 {
//line /usr/local/go/src/text/template/exec.go:268
			_go_fuzz_dep_.CoverTab[29784]++
									s.printValue(node, val)
//line /usr/local/go/src/text/template/exec.go:269
			// _ = "end of CoverTab[29784]"
		} else {
//line /usr/local/go/src/text/template/exec.go:270
			_go_fuzz_dep_.CoverTab[29785]++
//line /usr/local/go/src/text/template/exec.go:270
			// _ = "end of CoverTab[29785]"
//line /usr/local/go/src/text/template/exec.go:270
		}
//line /usr/local/go/src/text/template/exec.go:270
		// _ = "end of CoverTab[29773]"
	case *parse.BreakNode:
//line /usr/local/go/src/text/template/exec.go:271
		_go_fuzz_dep_.CoverTab[29774]++
								panic(walkBreak)
//line /usr/local/go/src/text/template/exec.go:272
		// _ = "end of CoverTab[29774]"
	case *parse.CommentNode:
//line /usr/local/go/src/text/template/exec.go:273
		_go_fuzz_dep_.CoverTab[29775]++
//line /usr/local/go/src/text/template/exec.go:273
		// _ = "end of CoverTab[29775]"
	case *parse.ContinueNode:
//line /usr/local/go/src/text/template/exec.go:274
		_go_fuzz_dep_.CoverTab[29776]++
								panic(walkContinue)
//line /usr/local/go/src/text/template/exec.go:275
		// _ = "end of CoverTab[29776]"
	case *parse.IfNode:
//line /usr/local/go/src/text/template/exec.go:276
		_go_fuzz_dep_.CoverTab[29777]++
								s.walkIfOrWith(parse.NodeIf, dot, node.Pipe, node.List, node.ElseList)
//line /usr/local/go/src/text/template/exec.go:277
		// _ = "end of CoverTab[29777]"
	case *parse.ListNode:
//line /usr/local/go/src/text/template/exec.go:278
		_go_fuzz_dep_.CoverTab[29778]++
								for _, node := range node.Nodes {
//line /usr/local/go/src/text/template/exec.go:279
			_go_fuzz_dep_.CoverTab[29786]++
									s.walk(dot, node)
//line /usr/local/go/src/text/template/exec.go:280
			// _ = "end of CoverTab[29786]"
		}
//line /usr/local/go/src/text/template/exec.go:281
		// _ = "end of CoverTab[29778]"
	case *parse.RangeNode:
//line /usr/local/go/src/text/template/exec.go:282
		_go_fuzz_dep_.CoverTab[29779]++
								s.walkRange(dot, node)
//line /usr/local/go/src/text/template/exec.go:283
		// _ = "end of CoverTab[29779]"
	case *parse.TemplateNode:
//line /usr/local/go/src/text/template/exec.go:284
		_go_fuzz_dep_.CoverTab[29780]++
								s.walkTemplate(dot, node)
//line /usr/local/go/src/text/template/exec.go:285
		// _ = "end of CoverTab[29780]"
	case *parse.TextNode:
//line /usr/local/go/src/text/template/exec.go:286
		_go_fuzz_dep_.CoverTab[29781]++
								if _, err := s.wr.Write(node.Text); err != nil {
//line /usr/local/go/src/text/template/exec.go:287
			_go_fuzz_dep_.CoverTab[29787]++
									s.writeError(err)
//line /usr/local/go/src/text/template/exec.go:288
			// _ = "end of CoverTab[29787]"
		} else {
//line /usr/local/go/src/text/template/exec.go:289
			_go_fuzz_dep_.CoverTab[29788]++
//line /usr/local/go/src/text/template/exec.go:289
			// _ = "end of CoverTab[29788]"
//line /usr/local/go/src/text/template/exec.go:289
		}
//line /usr/local/go/src/text/template/exec.go:289
		// _ = "end of CoverTab[29781]"
	case *parse.WithNode:
//line /usr/local/go/src/text/template/exec.go:290
		_go_fuzz_dep_.CoverTab[29782]++
								s.walkIfOrWith(parse.NodeWith, dot, node.Pipe, node.List, node.ElseList)
//line /usr/local/go/src/text/template/exec.go:291
		// _ = "end of CoverTab[29782]"
	default:
//line /usr/local/go/src/text/template/exec.go:292
		_go_fuzz_dep_.CoverTab[29783]++
								s.errorf("unknown node: %s", node)
//line /usr/local/go/src/text/template/exec.go:293
		// _ = "end of CoverTab[29783]"
	}
//line /usr/local/go/src/text/template/exec.go:294
	// _ = "end of CoverTab[29772]"
}

// walkIfOrWith walks an 'if' or 'with' node. The two control structures
//line /usr/local/go/src/text/template/exec.go:297
// are identical in behavior except that 'with' sets dot.
//line /usr/local/go/src/text/template/exec.go:299
func (s *state) walkIfOrWith(typ parse.NodeType, dot reflect.Value, pipe *parse.PipeNode, list, elseList *parse.ListNode) {
//line /usr/local/go/src/text/template/exec.go:299
	_go_fuzz_dep_.CoverTab[29789]++
							defer s.pop(s.mark())
							val := s.evalPipeline(dot, pipe)
							truth, ok := isTrue(indirectInterface(val))
							if !ok {
//line /usr/local/go/src/text/template/exec.go:303
		_go_fuzz_dep_.CoverTab[29791]++
								s.errorf("if/with can't use %v", val)
//line /usr/local/go/src/text/template/exec.go:304
		// _ = "end of CoverTab[29791]"
	} else {
//line /usr/local/go/src/text/template/exec.go:305
		_go_fuzz_dep_.CoverTab[29792]++
//line /usr/local/go/src/text/template/exec.go:305
		// _ = "end of CoverTab[29792]"
//line /usr/local/go/src/text/template/exec.go:305
	}
//line /usr/local/go/src/text/template/exec.go:305
	// _ = "end of CoverTab[29789]"
//line /usr/local/go/src/text/template/exec.go:305
	_go_fuzz_dep_.CoverTab[29790]++
							if truth {
//line /usr/local/go/src/text/template/exec.go:306
		_go_fuzz_dep_.CoverTab[29793]++
								if typ == parse.NodeWith {
//line /usr/local/go/src/text/template/exec.go:307
			_go_fuzz_dep_.CoverTab[29794]++
									s.walk(val, list)
//line /usr/local/go/src/text/template/exec.go:308
			// _ = "end of CoverTab[29794]"
		} else {
//line /usr/local/go/src/text/template/exec.go:309
			_go_fuzz_dep_.CoverTab[29795]++
									s.walk(dot, list)
//line /usr/local/go/src/text/template/exec.go:310
			// _ = "end of CoverTab[29795]"
		}
//line /usr/local/go/src/text/template/exec.go:311
		// _ = "end of CoverTab[29793]"
	} else {
//line /usr/local/go/src/text/template/exec.go:312
		_go_fuzz_dep_.CoverTab[29796]++
//line /usr/local/go/src/text/template/exec.go:312
		if elseList != nil {
//line /usr/local/go/src/text/template/exec.go:312
			_go_fuzz_dep_.CoverTab[29797]++
									s.walk(dot, elseList)
//line /usr/local/go/src/text/template/exec.go:313
			// _ = "end of CoverTab[29797]"
		} else {
//line /usr/local/go/src/text/template/exec.go:314
			_go_fuzz_dep_.CoverTab[29798]++
//line /usr/local/go/src/text/template/exec.go:314
			// _ = "end of CoverTab[29798]"
//line /usr/local/go/src/text/template/exec.go:314
		}
//line /usr/local/go/src/text/template/exec.go:314
		// _ = "end of CoverTab[29796]"
//line /usr/local/go/src/text/template/exec.go:314
	}
//line /usr/local/go/src/text/template/exec.go:314
	// _ = "end of CoverTab[29790]"
}

// IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
//line /usr/local/go/src/text/template/exec.go:317
// and whether the value has a meaningful truth value. This is the definition of
//line /usr/local/go/src/text/template/exec.go:317
// truth used by if and other such actions.
//line /usr/local/go/src/text/template/exec.go:320
func IsTrue(val any) (truth, ok bool) {
//line /usr/local/go/src/text/template/exec.go:320
	_go_fuzz_dep_.CoverTab[29799]++
							return isTrue(reflect.ValueOf(val))
//line /usr/local/go/src/text/template/exec.go:321
	// _ = "end of CoverTab[29799]"
}

func isTrue(val reflect.Value) (truth, ok bool) {
//line /usr/local/go/src/text/template/exec.go:324
	_go_fuzz_dep_.CoverTab[29800]++
							if !val.IsValid() {
//line /usr/local/go/src/text/template/exec.go:325
		_go_fuzz_dep_.CoverTab[29803]++

								return false, true
//line /usr/local/go/src/text/template/exec.go:327
		// _ = "end of CoverTab[29803]"
	} else {
//line /usr/local/go/src/text/template/exec.go:328
		_go_fuzz_dep_.CoverTab[29804]++
//line /usr/local/go/src/text/template/exec.go:328
		// _ = "end of CoverTab[29804]"
//line /usr/local/go/src/text/template/exec.go:328
	}
//line /usr/local/go/src/text/template/exec.go:328
	// _ = "end of CoverTab[29800]"
//line /usr/local/go/src/text/template/exec.go:328
	_go_fuzz_dep_.CoverTab[29801]++
							switch val.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /usr/local/go/src/text/template/exec.go:330
		_go_fuzz_dep_.CoverTab[29805]++
								truth = val.Len() > 0
//line /usr/local/go/src/text/template/exec.go:331
		// _ = "end of CoverTab[29805]"
	case reflect.Bool:
//line /usr/local/go/src/text/template/exec.go:332
		_go_fuzz_dep_.CoverTab[29806]++
								truth = val.Bool()
//line /usr/local/go/src/text/template/exec.go:333
		// _ = "end of CoverTab[29806]"
	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/text/template/exec.go:334
		_go_fuzz_dep_.CoverTab[29807]++
								truth = val.Complex() != 0
//line /usr/local/go/src/text/template/exec.go:335
		// _ = "end of CoverTab[29807]"
	case reflect.Chan, reflect.Func, reflect.Pointer, reflect.Interface:
//line /usr/local/go/src/text/template/exec.go:336
		_go_fuzz_dep_.CoverTab[29808]++
								truth = !val.IsNil()
//line /usr/local/go/src/text/template/exec.go:337
		// _ = "end of CoverTab[29808]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/text/template/exec.go:338
		_go_fuzz_dep_.CoverTab[29809]++
								truth = val.Int() != 0
//line /usr/local/go/src/text/template/exec.go:339
		// _ = "end of CoverTab[29809]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/text/template/exec.go:340
		_go_fuzz_dep_.CoverTab[29810]++
								truth = val.Float() != 0
//line /usr/local/go/src/text/template/exec.go:341
		// _ = "end of CoverTab[29810]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/text/template/exec.go:342
		_go_fuzz_dep_.CoverTab[29811]++
								truth = val.Uint() != 0
//line /usr/local/go/src/text/template/exec.go:343
		// _ = "end of CoverTab[29811]"
	case reflect.Struct:
//line /usr/local/go/src/text/template/exec.go:344
		_go_fuzz_dep_.CoverTab[29812]++
								truth = true
//line /usr/local/go/src/text/template/exec.go:345
		// _ = "end of CoverTab[29812]"
	default:
//line /usr/local/go/src/text/template/exec.go:346
		_go_fuzz_dep_.CoverTab[29813]++
								return
//line /usr/local/go/src/text/template/exec.go:347
		// _ = "end of CoverTab[29813]"
	}
//line /usr/local/go/src/text/template/exec.go:348
	// _ = "end of CoverTab[29801]"
//line /usr/local/go/src/text/template/exec.go:348
	_go_fuzz_dep_.CoverTab[29802]++
							return truth, true
//line /usr/local/go/src/text/template/exec.go:349
	// _ = "end of CoverTab[29802]"
}

func (s *state) walkRange(dot reflect.Value, r *parse.RangeNode) {
//line /usr/local/go/src/text/template/exec.go:352
	_go_fuzz_dep_.CoverTab[29814]++
							s.at(r)
							defer func() {
//line /usr/local/go/src/text/template/exec.go:354
		_go_fuzz_dep_.CoverTab[29818]++
								if r := recover(); r != nil && func() bool {
//line /usr/local/go/src/text/template/exec.go:355
			_go_fuzz_dep_.CoverTab[29819]++
//line /usr/local/go/src/text/template/exec.go:355
			return r != walkBreak
//line /usr/local/go/src/text/template/exec.go:355
			// _ = "end of CoverTab[29819]"
//line /usr/local/go/src/text/template/exec.go:355
		}() {
//line /usr/local/go/src/text/template/exec.go:355
			_go_fuzz_dep_.CoverTab[29820]++
									panic(r)
//line /usr/local/go/src/text/template/exec.go:356
			// _ = "end of CoverTab[29820]"
		} else {
//line /usr/local/go/src/text/template/exec.go:357
			_go_fuzz_dep_.CoverTab[29821]++
//line /usr/local/go/src/text/template/exec.go:357
			// _ = "end of CoverTab[29821]"
//line /usr/local/go/src/text/template/exec.go:357
		}
//line /usr/local/go/src/text/template/exec.go:357
		// _ = "end of CoverTab[29818]"
	}()
//line /usr/local/go/src/text/template/exec.go:358
	// _ = "end of CoverTab[29814]"
//line /usr/local/go/src/text/template/exec.go:358
	_go_fuzz_dep_.CoverTab[29815]++
							defer s.pop(s.mark())
							val, _ := indirect(s.evalPipeline(dot, r.Pipe))

							mark := s.mark()
							oneIteration := func(index, elem reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:363
		_go_fuzz_dep_.CoverTab[29822]++
								if len(r.Pipe.Decl) > 0 {
//line /usr/local/go/src/text/template/exec.go:364
			_go_fuzz_dep_.CoverTab[29826]++
									if r.Pipe.IsAssign {
//line /usr/local/go/src/text/template/exec.go:365
				_go_fuzz_dep_.CoverTab[29827]++

//line /usr/local/go/src/text/template/exec.go:368
				if len(r.Pipe.Decl) > 1 {
//line /usr/local/go/src/text/template/exec.go:368
					_go_fuzz_dep_.CoverTab[29828]++
											s.setVar(r.Pipe.Decl[0].Ident[0], index)
//line /usr/local/go/src/text/template/exec.go:369
					// _ = "end of CoverTab[29828]"
				} else {
//line /usr/local/go/src/text/template/exec.go:370
					_go_fuzz_dep_.CoverTab[29829]++
											s.setVar(r.Pipe.Decl[0].Ident[0], elem)
//line /usr/local/go/src/text/template/exec.go:371
					// _ = "end of CoverTab[29829]"
				}
//line /usr/local/go/src/text/template/exec.go:372
				// _ = "end of CoverTab[29827]"
			} else {
//line /usr/local/go/src/text/template/exec.go:373
				_go_fuzz_dep_.CoverTab[29830]++

//line /usr/local/go/src/text/template/exec.go:376
				s.setTopVar(1, elem)
//line /usr/local/go/src/text/template/exec.go:376
				// _ = "end of CoverTab[29830]"
			}
//line /usr/local/go/src/text/template/exec.go:377
			// _ = "end of CoverTab[29826]"
		} else {
//line /usr/local/go/src/text/template/exec.go:378
			_go_fuzz_dep_.CoverTab[29831]++
//line /usr/local/go/src/text/template/exec.go:378
			// _ = "end of CoverTab[29831]"
//line /usr/local/go/src/text/template/exec.go:378
		}
//line /usr/local/go/src/text/template/exec.go:378
		// _ = "end of CoverTab[29822]"
//line /usr/local/go/src/text/template/exec.go:378
		_go_fuzz_dep_.CoverTab[29823]++
								if len(r.Pipe.Decl) > 1 {
//line /usr/local/go/src/text/template/exec.go:379
			_go_fuzz_dep_.CoverTab[29832]++
									if r.Pipe.IsAssign {
//line /usr/local/go/src/text/template/exec.go:380
				_go_fuzz_dep_.CoverTab[29833]++
										s.setVar(r.Pipe.Decl[1].Ident[0], elem)
//line /usr/local/go/src/text/template/exec.go:381
				// _ = "end of CoverTab[29833]"
			} else {
//line /usr/local/go/src/text/template/exec.go:382
				_go_fuzz_dep_.CoverTab[29834]++

//line /usr/local/go/src/text/template/exec.go:385
				s.setTopVar(2, index)
//line /usr/local/go/src/text/template/exec.go:385
				// _ = "end of CoverTab[29834]"
			}
//line /usr/local/go/src/text/template/exec.go:386
			// _ = "end of CoverTab[29832]"
		} else {
//line /usr/local/go/src/text/template/exec.go:387
			_go_fuzz_dep_.CoverTab[29835]++
//line /usr/local/go/src/text/template/exec.go:387
			// _ = "end of CoverTab[29835]"
//line /usr/local/go/src/text/template/exec.go:387
		}
//line /usr/local/go/src/text/template/exec.go:387
		// _ = "end of CoverTab[29823]"
//line /usr/local/go/src/text/template/exec.go:387
		_go_fuzz_dep_.CoverTab[29824]++
								defer s.pop(mark)
								defer func() {
//line /usr/local/go/src/text/template/exec.go:389
			_go_fuzz_dep_.CoverTab[29836]++

									if r := recover(); r != nil && func() bool {
//line /usr/local/go/src/text/template/exec.go:391
				_go_fuzz_dep_.CoverTab[29837]++
//line /usr/local/go/src/text/template/exec.go:391
				return r != walkContinue
//line /usr/local/go/src/text/template/exec.go:391
				// _ = "end of CoverTab[29837]"
//line /usr/local/go/src/text/template/exec.go:391
			}() {
//line /usr/local/go/src/text/template/exec.go:391
				_go_fuzz_dep_.CoverTab[29838]++
										panic(r)
//line /usr/local/go/src/text/template/exec.go:392
				// _ = "end of CoverTab[29838]"
			} else {
//line /usr/local/go/src/text/template/exec.go:393
				_go_fuzz_dep_.CoverTab[29839]++
//line /usr/local/go/src/text/template/exec.go:393
				// _ = "end of CoverTab[29839]"
//line /usr/local/go/src/text/template/exec.go:393
			}
//line /usr/local/go/src/text/template/exec.go:393
			// _ = "end of CoverTab[29836]"
		}()
//line /usr/local/go/src/text/template/exec.go:394
		// _ = "end of CoverTab[29824]"
//line /usr/local/go/src/text/template/exec.go:394
		_go_fuzz_dep_.CoverTab[29825]++
								s.walk(elem, r.List)
//line /usr/local/go/src/text/template/exec.go:395
		// _ = "end of CoverTab[29825]"
	}
//line /usr/local/go/src/text/template/exec.go:396
	// _ = "end of CoverTab[29815]"
//line /usr/local/go/src/text/template/exec.go:396
	_go_fuzz_dep_.CoverTab[29816]++
							switch val.Kind() {
	case reflect.Array, reflect.Slice:
//line /usr/local/go/src/text/template/exec.go:398
		_go_fuzz_dep_.CoverTab[29840]++
								if val.Len() == 0 {
//line /usr/local/go/src/text/template/exec.go:399
			_go_fuzz_dep_.CoverTab[29853]++
									break
//line /usr/local/go/src/text/template/exec.go:400
			// _ = "end of CoverTab[29853]"
		} else {
//line /usr/local/go/src/text/template/exec.go:401
			_go_fuzz_dep_.CoverTab[29854]++
//line /usr/local/go/src/text/template/exec.go:401
			// _ = "end of CoverTab[29854]"
//line /usr/local/go/src/text/template/exec.go:401
		}
//line /usr/local/go/src/text/template/exec.go:401
		// _ = "end of CoverTab[29840]"
//line /usr/local/go/src/text/template/exec.go:401
		_go_fuzz_dep_.CoverTab[29841]++
								for i := 0; i < val.Len(); i++ {
//line /usr/local/go/src/text/template/exec.go:402
			_go_fuzz_dep_.CoverTab[29855]++
									oneIteration(reflect.ValueOf(i), val.Index(i))
//line /usr/local/go/src/text/template/exec.go:403
			// _ = "end of CoverTab[29855]"
		}
//line /usr/local/go/src/text/template/exec.go:404
		// _ = "end of CoverTab[29841]"
//line /usr/local/go/src/text/template/exec.go:404
		_go_fuzz_dep_.CoverTab[29842]++
								return
//line /usr/local/go/src/text/template/exec.go:405
		// _ = "end of CoverTab[29842]"
	case reflect.Map:
//line /usr/local/go/src/text/template/exec.go:406
		_go_fuzz_dep_.CoverTab[29843]++
								if val.Len() == 0 {
//line /usr/local/go/src/text/template/exec.go:407
			_go_fuzz_dep_.CoverTab[29856]++
									break
//line /usr/local/go/src/text/template/exec.go:408
			// _ = "end of CoverTab[29856]"
		} else {
//line /usr/local/go/src/text/template/exec.go:409
			_go_fuzz_dep_.CoverTab[29857]++
//line /usr/local/go/src/text/template/exec.go:409
			// _ = "end of CoverTab[29857]"
//line /usr/local/go/src/text/template/exec.go:409
		}
//line /usr/local/go/src/text/template/exec.go:409
		// _ = "end of CoverTab[29843]"
//line /usr/local/go/src/text/template/exec.go:409
		_go_fuzz_dep_.CoverTab[29844]++
								om := fmtsort.Sort(val)
								for i, key := range om.Key {
//line /usr/local/go/src/text/template/exec.go:411
			_go_fuzz_dep_.CoverTab[29858]++
									oneIteration(key, om.Value[i])
//line /usr/local/go/src/text/template/exec.go:412
			// _ = "end of CoverTab[29858]"
		}
//line /usr/local/go/src/text/template/exec.go:413
		// _ = "end of CoverTab[29844]"
//line /usr/local/go/src/text/template/exec.go:413
		_go_fuzz_dep_.CoverTab[29845]++
								return
//line /usr/local/go/src/text/template/exec.go:414
		// _ = "end of CoverTab[29845]"
	case reflect.Chan:
//line /usr/local/go/src/text/template/exec.go:415
		_go_fuzz_dep_.CoverTab[29846]++
								if val.IsNil() {
//line /usr/local/go/src/text/template/exec.go:416
			_go_fuzz_dep_.CoverTab[29859]++
									break
//line /usr/local/go/src/text/template/exec.go:417
			// _ = "end of CoverTab[29859]"
		} else {
//line /usr/local/go/src/text/template/exec.go:418
			_go_fuzz_dep_.CoverTab[29860]++
//line /usr/local/go/src/text/template/exec.go:418
			// _ = "end of CoverTab[29860]"
//line /usr/local/go/src/text/template/exec.go:418
		}
//line /usr/local/go/src/text/template/exec.go:418
		// _ = "end of CoverTab[29846]"
//line /usr/local/go/src/text/template/exec.go:418
		_go_fuzz_dep_.CoverTab[29847]++
								if val.Type().ChanDir() == reflect.SendDir {
//line /usr/local/go/src/text/template/exec.go:419
			_go_fuzz_dep_.CoverTab[29861]++
									s.errorf("range over send-only channel %v", val)
									break
//line /usr/local/go/src/text/template/exec.go:421
			// _ = "end of CoverTab[29861]"
		} else {
//line /usr/local/go/src/text/template/exec.go:422
			_go_fuzz_dep_.CoverTab[29862]++
//line /usr/local/go/src/text/template/exec.go:422
			// _ = "end of CoverTab[29862]"
//line /usr/local/go/src/text/template/exec.go:422
		}
//line /usr/local/go/src/text/template/exec.go:422
		// _ = "end of CoverTab[29847]"
//line /usr/local/go/src/text/template/exec.go:422
		_go_fuzz_dep_.CoverTab[29848]++
								i := 0
								for ; ; i++ {
//line /usr/local/go/src/text/template/exec.go:424
			_go_fuzz_dep_.CoverTab[29863]++
									elem, ok := val.Recv()
									if !ok {
//line /usr/local/go/src/text/template/exec.go:426
				_go_fuzz_dep_.CoverTab[29865]++
										break
//line /usr/local/go/src/text/template/exec.go:427
				// _ = "end of CoverTab[29865]"
			} else {
//line /usr/local/go/src/text/template/exec.go:428
				_go_fuzz_dep_.CoverTab[29866]++
//line /usr/local/go/src/text/template/exec.go:428
				// _ = "end of CoverTab[29866]"
//line /usr/local/go/src/text/template/exec.go:428
			}
//line /usr/local/go/src/text/template/exec.go:428
			// _ = "end of CoverTab[29863]"
//line /usr/local/go/src/text/template/exec.go:428
			_go_fuzz_dep_.CoverTab[29864]++
									oneIteration(reflect.ValueOf(i), elem)
//line /usr/local/go/src/text/template/exec.go:429
			// _ = "end of CoverTab[29864]"
		}
//line /usr/local/go/src/text/template/exec.go:430
		// _ = "end of CoverTab[29848]"
//line /usr/local/go/src/text/template/exec.go:430
		_go_fuzz_dep_.CoverTab[29849]++
								if i == 0 {
//line /usr/local/go/src/text/template/exec.go:431
			_go_fuzz_dep_.CoverTab[29867]++
									break
//line /usr/local/go/src/text/template/exec.go:432
			// _ = "end of CoverTab[29867]"
		} else {
//line /usr/local/go/src/text/template/exec.go:433
			_go_fuzz_dep_.CoverTab[29868]++
//line /usr/local/go/src/text/template/exec.go:433
			// _ = "end of CoverTab[29868]"
//line /usr/local/go/src/text/template/exec.go:433
		}
//line /usr/local/go/src/text/template/exec.go:433
		// _ = "end of CoverTab[29849]"
//line /usr/local/go/src/text/template/exec.go:433
		_go_fuzz_dep_.CoverTab[29850]++
								return
//line /usr/local/go/src/text/template/exec.go:434
		// _ = "end of CoverTab[29850]"
	case reflect.Invalid:
//line /usr/local/go/src/text/template/exec.go:435
		_go_fuzz_dep_.CoverTab[29851]++
								break
//line /usr/local/go/src/text/template/exec.go:436
		// _ = "end of CoverTab[29851]"
	default:
//line /usr/local/go/src/text/template/exec.go:437
		_go_fuzz_dep_.CoverTab[29852]++
								s.errorf("range can't iterate over %v", val)
//line /usr/local/go/src/text/template/exec.go:438
		// _ = "end of CoverTab[29852]"
	}
//line /usr/local/go/src/text/template/exec.go:439
	// _ = "end of CoverTab[29816]"
//line /usr/local/go/src/text/template/exec.go:439
	_go_fuzz_dep_.CoverTab[29817]++
							if r.ElseList != nil {
//line /usr/local/go/src/text/template/exec.go:440
		_go_fuzz_dep_.CoverTab[29869]++
								s.walk(dot, r.ElseList)
//line /usr/local/go/src/text/template/exec.go:441
		// _ = "end of CoverTab[29869]"
	} else {
//line /usr/local/go/src/text/template/exec.go:442
		_go_fuzz_dep_.CoverTab[29870]++
//line /usr/local/go/src/text/template/exec.go:442
		// _ = "end of CoverTab[29870]"
//line /usr/local/go/src/text/template/exec.go:442
	}
//line /usr/local/go/src/text/template/exec.go:442
	// _ = "end of CoverTab[29817]"
}

func (s *state) walkTemplate(dot reflect.Value, t *parse.TemplateNode) {
//line /usr/local/go/src/text/template/exec.go:445
	_go_fuzz_dep_.CoverTab[29871]++
							s.at(t)
							tmpl := s.tmpl.Lookup(t.Name)
							if tmpl == nil {
//line /usr/local/go/src/text/template/exec.go:448
		_go_fuzz_dep_.CoverTab[29874]++
								s.errorf("template %q not defined", t.Name)
//line /usr/local/go/src/text/template/exec.go:449
		// _ = "end of CoverTab[29874]"
	} else {
//line /usr/local/go/src/text/template/exec.go:450
		_go_fuzz_dep_.CoverTab[29875]++
//line /usr/local/go/src/text/template/exec.go:450
		// _ = "end of CoverTab[29875]"
//line /usr/local/go/src/text/template/exec.go:450
	}
//line /usr/local/go/src/text/template/exec.go:450
	// _ = "end of CoverTab[29871]"
//line /usr/local/go/src/text/template/exec.go:450
	_go_fuzz_dep_.CoverTab[29872]++
							if s.depth == maxExecDepth {
//line /usr/local/go/src/text/template/exec.go:451
		_go_fuzz_dep_.CoverTab[29876]++
								s.errorf("exceeded maximum template depth (%v)", maxExecDepth)
//line /usr/local/go/src/text/template/exec.go:452
		// _ = "end of CoverTab[29876]"
	} else {
//line /usr/local/go/src/text/template/exec.go:453
		_go_fuzz_dep_.CoverTab[29877]++
//line /usr/local/go/src/text/template/exec.go:453
		// _ = "end of CoverTab[29877]"
//line /usr/local/go/src/text/template/exec.go:453
	}
//line /usr/local/go/src/text/template/exec.go:453
	// _ = "end of CoverTab[29872]"
//line /usr/local/go/src/text/template/exec.go:453
	_go_fuzz_dep_.CoverTab[29873]++

							dot = s.evalPipeline(dot, t.Pipe)
							newState := *s
							newState.depth++
							newState.tmpl = tmpl

							newState.vars = []variable{{"$", dot}}
							newState.walk(dot, tmpl.Root)
//line /usr/local/go/src/text/template/exec.go:461
	// _ = "end of CoverTab[29873]"
}

//line /usr/local/go/src/text/template/exec.go:468
// evalPipeline returns the value acquired by evaluating a pipeline. If the
//line /usr/local/go/src/text/template/exec.go:468
// pipeline has a variable declaration, the variable will be pushed on the
//line /usr/local/go/src/text/template/exec.go:468
// stack. Callers should therefore pop the stack after they are finished
//line /usr/local/go/src/text/template/exec.go:468
// executing commands depending on the pipeline value.
//line /usr/local/go/src/text/template/exec.go:472
func (s *state) evalPipeline(dot reflect.Value, pipe *parse.PipeNode) (value reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:472
	_go_fuzz_dep_.CoverTab[29878]++
							if pipe == nil {
//line /usr/local/go/src/text/template/exec.go:473
		_go_fuzz_dep_.CoverTab[29882]++
								return
//line /usr/local/go/src/text/template/exec.go:474
		// _ = "end of CoverTab[29882]"
	} else {
//line /usr/local/go/src/text/template/exec.go:475
		_go_fuzz_dep_.CoverTab[29883]++
//line /usr/local/go/src/text/template/exec.go:475
		// _ = "end of CoverTab[29883]"
//line /usr/local/go/src/text/template/exec.go:475
	}
//line /usr/local/go/src/text/template/exec.go:475
	// _ = "end of CoverTab[29878]"
//line /usr/local/go/src/text/template/exec.go:475
	_go_fuzz_dep_.CoverTab[29879]++
							s.at(pipe)
							value = missingVal
							for _, cmd := range pipe.Cmds {
//line /usr/local/go/src/text/template/exec.go:478
		_go_fuzz_dep_.CoverTab[29884]++
								value = s.evalCommand(dot, cmd, value)

								if value.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/text/template/exec.go:481
			_go_fuzz_dep_.CoverTab[29885]++
//line /usr/local/go/src/text/template/exec.go:481
			return value.Type().NumMethod() == 0
//line /usr/local/go/src/text/template/exec.go:481
			// _ = "end of CoverTab[29885]"
//line /usr/local/go/src/text/template/exec.go:481
		}() {
//line /usr/local/go/src/text/template/exec.go:481
			_go_fuzz_dep_.CoverTab[29886]++
									value = reflect.ValueOf(value.Interface())
//line /usr/local/go/src/text/template/exec.go:482
			// _ = "end of CoverTab[29886]"
		} else {
//line /usr/local/go/src/text/template/exec.go:483
			_go_fuzz_dep_.CoverTab[29887]++
//line /usr/local/go/src/text/template/exec.go:483
			// _ = "end of CoverTab[29887]"
//line /usr/local/go/src/text/template/exec.go:483
		}
//line /usr/local/go/src/text/template/exec.go:483
		// _ = "end of CoverTab[29884]"
	}
//line /usr/local/go/src/text/template/exec.go:484
	// _ = "end of CoverTab[29879]"
//line /usr/local/go/src/text/template/exec.go:484
	_go_fuzz_dep_.CoverTab[29880]++
							for _, variable := range pipe.Decl {
//line /usr/local/go/src/text/template/exec.go:485
		_go_fuzz_dep_.CoverTab[29888]++
								if pipe.IsAssign {
//line /usr/local/go/src/text/template/exec.go:486
			_go_fuzz_dep_.CoverTab[29889]++
									s.setVar(variable.Ident[0], value)
//line /usr/local/go/src/text/template/exec.go:487
			// _ = "end of CoverTab[29889]"
		} else {
//line /usr/local/go/src/text/template/exec.go:488
			_go_fuzz_dep_.CoverTab[29890]++
									s.push(variable.Ident[0], value)
//line /usr/local/go/src/text/template/exec.go:489
			// _ = "end of CoverTab[29890]"
		}
//line /usr/local/go/src/text/template/exec.go:490
		// _ = "end of CoverTab[29888]"
	}
//line /usr/local/go/src/text/template/exec.go:491
	// _ = "end of CoverTab[29880]"
//line /usr/local/go/src/text/template/exec.go:491
	_go_fuzz_dep_.CoverTab[29881]++
							return value
//line /usr/local/go/src/text/template/exec.go:492
	// _ = "end of CoverTab[29881]"
}

func (s *state) notAFunction(args []parse.Node, final reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:495
	_go_fuzz_dep_.CoverTab[29891]++
							if len(args) > 1 || func() bool {
//line /usr/local/go/src/text/template/exec.go:496
		_go_fuzz_dep_.CoverTab[29892]++
//line /usr/local/go/src/text/template/exec.go:496
		return !isMissing(final)
//line /usr/local/go/src/text/template/exec.go:496
		// _ = "end of CoverTab[29892]"
//line /usr/local/go/src/text/template/exec.go:496
	}() {
//line /usr/local/go/src/text/template/exec.go:496
		_go_fuzz_dep_.CoverTab[29893]++
								s.errorf("can't give argument to non-function %s", args[0])
//line /usr/local/go/src/text/template/exec.go:497
		// _ = "end of CoverTab[29893]"
	} else {
//line /usr/local/go/src/text/template/exec.go:498
		_go_fuzz_dep_.CoverTab[29894]++
//line /usr/local/go/src/text/template/exec.go:498
		// _ = "end of CoverTab[29894]"
//line /usr/local/go/src/text/template/exec.go:498
	}
//line /usr/local/go/src/text/template/exec.go:498
	// _ = "end of CoverTab[29891]"
}

func (s *state) evalCommand(dot reflect.Value, cmd *parse.CommandNode, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:501
	_go_fuzz_dep_.CoverTab[29895]++
							firstWord := cmd.Args[0]
							switch n := firstWord.(type) {
	case *parse.FieldNode:
//line /usr/local/go/src/text/template/exec.go:504
		_go_fuzz_dep_.CoverTab[29898]++
								return s.evalFieldNode(dot, n, cmd.Args, final)
//line /usr/local/go/src/text/template/exec.go:505
		// _ = "end of CoverTab[29898]"
	case *parse.ChainNode:
//line /usr/local/go/src/text/template/exec.go:506
		_go_fuzz_dep_.CoverTab[29899]++
								return s.evalChainNode(dot, n, cmd.Args, final)
//line /usr/local/go/src/text/template/exec.go:507
		// _ = "end of CoverTab[29899]"
	case *parse.IdentifierNode:
//line /usr/local/go/src/text/template/exec.go:508
		_go_fuzz_dep_.CoverTab[29900]++

								return s.evalFunction(dot, n, cmd, cmd.Args, final)
//line /usr/local/go/src/text/template/exec.go:510
		// _ = "end of CoverTab[29900]"
	case *parse.PipeNode:
//line /usr/local/go/src/text/template/exec.go:511
		_go_fuzz_dep_.CoverTab[29901]++

								s.notAFunction(cmd.Args, final)
								return s.evalPipeline(dot, n)
//line /usr/local/go/src/text/template/exec.go:514
		// _ = "end of CoverTab[29901]"
	case *parse.VariableNode:
//line /usr/local/go/src/text/template/exec.go:515
		_go_fuzz_dep_.CoverTab[29902]++
								return s.evalVariableNode(dot, n, cmd.Args, final)
//line /usr/local/go/src/text/template/exec.go:516
		// _ = "end of CoverTab[29902]"
	}
//line /usr/local/go/src/text/template/exec.go:517
	// _ = "end of CoverTab[29895]"
//line /usr/local/go/src/text/template/exec.go:517
	_go_fuzz_dep_.CoverTab[29896]++
							s.at(firstWord)
							s.notAFunction(cmd.Args, final)
							switch word := firstWord.(type) {
	case *parse.BoolNode:
//line /usr/local/go/src/text/template/exec.go:521
		_go_fuzz_dep_.CoverTab[29903]++
								return reflect.ValueOf(word.True)
//line /usr/local/go/src/text/template/exec.go:522
		// _ = "end of CoverTab[29903]"
	case *parse.DotNode:
//line /usr/local/go/src/text/template/exec.go:523
		_go_fuzz_dep_.CoverTab[29904]++
								return dot
//line /usr/local/go/src/text/template/exec.go:524
		// _ = "end of CoverTab[29904]"
	case *parse.NilNode:
//line /usr/local/go/src/text/template/exec.go:525
		_go_fuzz_dep_.CoverTab[29905]++
								s.errorf("nil is not a command")
//line /usr/local/go/src/text/template/exec.go:526
		// _ = "end of CoverTab[29905]"
	case *parse.NumberNode:
//line /usr/local/go/src/text/template/exec.go:527
		_go_fuzz_dep_.CoverTab[29906]++
								return s.idealConstant(word)
//line /usr/local/go/src/text/template/exec.go:528
		// _ = "end of CoverTab[29906]"
	case *parse.StringNode:
//line /usr/local/go/src/text/template/exec.go:529
		_go_fuzz_dep_.CoverTab[29907]++
								return reflect.ValueOf(word.Text)
//line /usr/local/go/src/text/template/exec.go:530
		// _ = "end of CoverTab[29907]"
	}
//line /usr/local/go/src/text/template/exec.go:531
	// _ = "end of CoverTab[29896]"
//line /usr/local/go/src/text/template/exec.go:531
	_go_fuzz_dep_.CoverTab[29897]++
							s.errorf("can't evaluate command %q", firstWord)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:533
	// _ = "end of CoverTab[29897]"
}

// idealConstant is called to return the value of a number in a context where
//line /usr/local/go/src/text/template/exec.go:536
// we don't know the type. In that case, the syntax of the number tells us
//line /usr/local/go/src/text/template/exec.go:536
// its type, and we use Go rules to resolve. Note there is no such thing as
//line /usr/local/go/src/text/template/exec.go:536
// a uint ideal constant in this situation - the value must be of int type.
//line /usr/local/go/src/text/template/exec.go:540
func (s *state) idealConstant(constant *parse.NumberNode) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:540
	_go_fuzz_dep_.CoverTab[29908]++

//line /usr/local/go/src/text/template/exec.go:544
	s.at(constant)
	switch {
	case constant.IsComplex:
//line /usr/local/go/src/text/template/exec.go:546
		_go_fuzz_dep_.CoverTab[29910]++
								return reflect.ValueOf(constant.Complex128)
//line /usr/local/go/src/text/template/exec.go:547
		// _ = "end of CoverTab[29910]"

	case constant.IsFloat && func() bool {
//line /usr/local/go/src/text/template/exec.go:549
		_go_fuzz_dep_.CoverTab[29916]++
//line /usr/local/go/src/text/template/exec.go:549
		return !isHexInt(constant.Text)
								// _ = "end of CoverTab[29916]"
//line /usr/local/go/src/text/template/exec.go:550
	}() && func() bool {
//line /usr/local/go/src/text/template/exec.go:550
		_go_fuzz_dep_.CoverTab[29917]++
//line /usr/local/go/src/text/template/exec.go:550
		return !isRuneInt(constant.Text)
//line /usr/local/go/src/text/template/exec.go:550
		// _ = "end of CoverTab[29917]"
//line /usr/local/go/src/text/template/exec.go:550
	}() && func() bool {
//line /usr/local/go/src/text/template/exec.go:550
		_go_fuzz_dep_.CoverTab[29918]++
//line /usr/local/go/src/text/template/exec.go:550
		return strings.ContainsAny(constant.Text, ".eEpP")
								// _ = "end of CoverTab[29918]"
//line /usr/local/go/src/text/template/exec.go:551
	}():
//line /usr/local/go/src/text/template/exec.go:551
		_go_fuzz_dep_.CoverTab[29911]++
								return reflect.ValueOf(constant.Float64)
//line /usr/local/go/src/text/template/exec.go:552
		// _ = "end of CoverTab[29911]"

	case constant.IsInt:
//line /usr/local/go/src/text/template/exec.go:554
		_go_fuzz_dep_.CoverTab[29912]++
								n := int(constant.Int64)
								if int64(n) != constant.Int64 {
//line /usr/local/go/src/text/template/exec.go:556
			_go_fuzz_dep_.CoverTab[29919]++
									s.errorf("%s overflows int", constant.Text)
//line /usr/local/go/src/text/template/exec.go:557
			// _ = "end of CoverTab[29919]"
		} else {
//line /usr/local/go/src/text/template/exec.go:558
			_go_fuzz_dep_.CoverTab[29920]++
//line /usr/local/go/src/text/template/exec.go:558
			// _ = "end of CoverTab[29920]"
//line /usr/local/go/src/text/template/exec.go:558
		}
//line /usr/local/go/src/text/template/exec.go:558
		// _ = "end of CoverTab[29912]"
//line /usr/local/go/src/text/template/exec.go:558
		_go_fuzz_dep_.CoverTab[29913]++
								return reflect.ValueOf(n)
//line /usr/local/go/src/text/template/exec.go:559
		// _ = "end of CoverTab[29913]"

	case constant.IsUint:
//line /usr/local/go/src/text/template/exec.go:561
		_go_fuzz_dep_.CoverTab[29914]++
								s.errorf("%s overflows int", constant.Text)
//line /usr/local/go/src/text/template/exec.go:562
		// _ = "end of CoverTab[29914]"
//line /usr/local/go/src/text/template/exec.go:562
	default:
//line /usr/local/go/src/text/template/exec.go:562
		_go_fuzz_dep_.CoverTab[29915]++
//line /usr/local/go/src/text/template/exec.go:562
		// _ = "end of CoverTab[29915]"
	}
//line /usr/local/go/src/text/template/exec.go:563
	// _ = "end of CoverTab[29908]"
//line /usr/local/go/src/text/template/exec.go:563
	_go_fuzz_dep_.CoverTab[29909]++
							return zero
//line /usr/local/go/src/text/template/exec.go:564
	// _ = "end of CoverTab[29909]"
}

func isRuneInt(s string) bool {
//line /usr/local/go/src/text/template/exec.go:567
	_go_fuzz_dep_.CoverTab[29921]++
							return len(s) > 0 && func() bool {
//line /usr/local/go/src/text/template/exec.go:568
		_go_fuzz_dep_.CoverTab[29922]++
//line /usr/local/go/src/text/template/exec.go:568
		return s[0] == '\''
//line /usr/local/go/src/text/template/exec.go:568
		// _ = "end of CoverTab[29922]"
//line /usr/local/go/src/text/template/exec.go:568
	}()
//line /usr/local/go/src/text/template/exec.go:568
	// _ = "end of CoverTab[29921]"
}

func isHexInt(s string) bool {
//line /usr/local/go/src/text/template/exec.go:571
	_go_fuzz_dep_.CoverTab[29923]++
							return len(s) > 2 && func() bool {
//line /usr/local/go/src/text/template/exec.go:572
		_go_fuzz_dep_.CoverTab[29924]++
//line /usr/local/go/src/text/template/exec.go:572
		return s[0] == '0'
//line /usr/local/go/src/text/template/exec.go:572
		// _ = "end of CoverTab[29924]"
//line /usr/local/go/src/text/template/exec.go:572
	}() && func() bool {
//line /usr/local/go/src/text/template/exec.go:572
		_go_fuzz_dep_.CoverTab[29925]++
//line /usr/local/go/src/text/template/exec.go:572
		return (s[1] == 'x' || func() bool {
//line /usr/local/go/src/text/template/exec.go:572
			_go_fuzz_dep_.CoverTab[29926]++
//line /usr/local/go/src/text/template/exec.go:572
			return s[1] == 'X'
//line /usr/local/go/src/text/template/exec.go:572
			// _ = "end of CoverTab[29926]"
//line /usr/local/go/src/text/template/exec.go:572
		}())
//line /usr/local/go/src/text/template/exec.go:572
		// _ = "end of CoverTab[29925]"
//line /usr/local/go/src/text/template/exec.go:572
	}() && func() bool {
//line /usr/local/go/src/text/template/exec.go:572
		_go_fuzz_dep_.CoverTab[29927]++
//line /usr/local/go/src/text/template/exec.go:572
		return !strings.ContainsAny(s, "pP")
//line /usr/local/go/src/text/template/exec.go:572
		// _ = "end of CoverTab[29927]"
//line /usr/local/go/src/text/template/exec.go:572
	}()
//line /usr/local/go/src/text/template/exec.go:572
	// _ = "end of CoverTab[29923]"
}

func (s *state) evalFieldNode(dot reflect.Value, field *parse.FieldNode, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:575
	_go_fuzz_dep_.CoverTab[29928]++
							s.at(field)
							return s.evalFieldChain(dot, dot, field, field.Ident, args, final)
//line /usr/local/go/src/text/template/exec.go:577
	// _ = "end of CoverTab[29928]"
}

func (s *state) evalChainNode(dot reflect.Value, chain *parse.ChainNode, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:580
	_go_fuzz_dep_.CoverTab[29929]++
							s.at(chain)
							if len(chain.Field) == 0 {
//line /usr/local/go/src/text/template/exec.go:582
		_go_fuzz_dep_.CoverTab[29932]++
								s.errorf("internal error: no fields in evalChainNode")
//line /usr/local/go/src/text/template/exec.go:583
		// _ = "end of CoverTab[29932]"
	} else {
//line /usr/local/go/src/text/template/exec.go:584
		_go_fuzz_dep_.CoverTab[29933]++
//line /usr/local/go/src/text/template/exec.go:584
		// _ = "end of CoverTab[29933]"
//line /usr/local/go/src/text/template/exec.go:584
	}
//line /usr/local/go/src/text/template/exec.go:584
	// _ = "end of CoverTab[29929]"
//line /usr/local/go/src/text/template/exec.go:584
	_go_fuzz_dep_.CoverTab[29930]++
							if chain.Node.Type() == parse.NodeNil {
//line /usr/local/go/src/text/template/exec.go:585
		_go_fuzz_dep_.CoverTab[29934]++
								s.errorf("indirection through explicit nil in %s", chain)
//line /usr/local/go/src/text/template/exec.go:586
		// _ = "end of CoverTab[29934]"
	} else {
//line /usr/local/go/src/text/template/exec.go:587
		_go_fuzz_dep_.CoverTab[29935]++
//line /usr/local/go/src/text/template/exec.go:587
		// _ = "end of CoverTab[29935]"
//line /usr/local/go/src/text/template/exec.go:587
	}
//line /usr/local/go/src/text/template/exec.go:587
	// _ = "end of CoverTab[29930]"
//line /usr/local/go/src/text/template/exec.go:587
	_go_fuzz_dep_.CoverTab[29931]++

							pipe := s.evalArg(dot, nil, chain.Node)
							return s.evalFieldChain(dot, pipe, chain, chain.Field, args, final)
//line /usr/local/go/src/text/template/exec.go:590
	// _ = "end of CoverTab[29931]"
}

func (s *state) evalVariableNode(dot reflect.Value, variable *parse.VariableNode, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:593
	_go_fuzz_dep_.CoverTab[29936]++

							s.at(variable)
							value := s.varValue(variable.Ident[0])
							if len(variable.Ident) == 1 {
//line /usr/local/go/src/text/template/exec.go:597
		_go_fuzz_dep_.CoverTab[29938]++
								s.notAFunction(args, final)
								return value
//line /usr/local/go/src/text/template/exec.go:599
		// _ = "end of CoverTab[29938]"
	} else {
//line /usr/local/go/src/text/template/exec.go:600
		_go_fuzz_dep_.CoverTab[29939]++
//line /usr/local/go/src/text/template/exec.go:600
		// _ = "end of CoverTab[29939]"
//line /usr/local/go/src/text/template/exec.go:600
	}
//line /usr/local/go/src/text/template/exec.go:600
	// _ = "end of CoverTab[29936]"
//line /usr/local/go/src/text/template/exec.go:600
	_go_fuzz_dep_.CoverTab[29937]++
							return s.evalFieldChain(dot, value, variable, variable.Ident[1:], args, final)
//line /usr/local/go/src/text/template/exec.go:601
	// _ = "end of CoverTab[29937]"
}

// evalFieldChain evaluates .X.Y.Z possibly followed by arguments.
//line /usr/local/go/src/text/template/exec.go:604
// dot is the environment in which to evaluate arguments, while
//line /usr/local/go/src/text/template/exec.go:604
// receiver is the value being walked along the chain.
//line /usr/local/go/src/text/template/exec.go:607
func (s *state) evalFieldChain(dot, receiver reflect.Value, node parse.Node, ident []string, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:607
	_go_fuzz_dep_.CoverTab[29940]++
							n := len(ident)
							for i := 0; i < n-1; i++ {
//line /usr/local/go/src/text/template/exec.go:609
		_go_fuzz_dep_.CoverTab[29942]++
								receiver = s.evalField(dot, ident[i], node, nil, missingVal, receiver)
//line /usr/local/go/src/text/template/exec.go:610
		// _ = "end of CoverTab[29942]"
	}
//line /usr/local/go/src/text/template/exec.go:611
	// _ = "end of CoverTab[29940]"
//line /usr/local/go/src/text/template/exec.go:611
	_go_fuzz_dep_.CoverTab[29941]++

							return s.evalField(dot, ident[n-1], node, args, final, receiver)
//line /usr/local/go/src/text/template/exec.go:613
	// _ = "end of CoverTab[29941]"
}

func (s *state) evalFunction(dot reflect.Value, node *parse.IdentifierNode, cmd parse.Node, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:616
	_go_fuzz_dep_.CoverTab[29943]++
							s.at(node)
							name := node.Ident
							function, isBuiltin, ok := findFunction(name, s.tmpl)
							if !ok {
//line /usr/local/go/src/text/template/exec.go:620
		_go_fuzz_dep_.CoverTab[29945]++
								s.errorf("%q is not a defined function", name)
//line /usr/local/go/src/text/template/exec.go:621
		// _ = "end of CoverTab[29945]"
	} else {
//line /usr/local/go/src/text/template/exec.go:622
		_go_fuzz_dep_.CoverTab[29946]++
//line /usr/local/go/src/text/template/exec.go:622
		// _ = "end of CoverTab[29946]"
//line /usr/local/go/src/text/template/exec.go:622
	}
//line /usr/local/go/src/text/template/exec.go:622
	// _ = "end of CoverTab[29943]"
//line /usr/local/go/src/text/template/exec.go:622
	_go_fuzz_dep_.CoverTab[29944]++
							return s.evalCall(dot, function, isBuiltin, cmd, name, args, final)
//line /usr/local/go/src/text/template/exec.go:623
	// _ = "end of CoverTab[29944]"
}

// evalField evaluates an expression like (.Field) or (.Field arg1 arg2).
//line /usr/local/go/src/text/template/exec.go:626
// The 'final' argument represents the return value from the preceding
//line /usr/local/go/src/text/template/exec.go:626
// value of the pipeline, if any.
//line /usr/local/go/src/text/template/exec.go:629
func (s *state) evalField(dot reflect.Value, fieldName string, node parse.Node, args []parse.Node, final, receiver reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:629
	_go_fuzz_dep_.CoverTab[29947]++
							if !receiver.IsValid() {
//line /usr/local/go/src/text/template/exec.go:630
		_go_fuzz_dep_.CoverTab[29953]++
								if s.tmpl.option.missingKey == mapError {
//line /usr/local/go/src/text/template/exec.go:631
			_go_fuzz_dep_.CoverTab[29955]++
									s.errorf("nil data; no entry for key %q", fieldName)
//line /usr/local/go/src/text/template/exec.go:632
			// _ = "end of CoverTab[29955]"
		} else {
//line /usr/local/go/src/text/template/exec.go:633
			_go_fuzz_dep_.CoverTab[29956]++
//line /usr/local/go/src/text/template/exec.go:633
			// _ = "end of CoverTab[29956]"
//line /usr/local/go/src/text/template/exec.go:633
		}
//line /usr/local/go/src/text/template/exec.go:633
		// _ = "end of CoverTab[29953]"
//line /usr/local/go/src/text/template/exec.go:633
		_go_fuzz_dep_.CoverTab[29954]++
								return zero
//line /usr/local/go/src/text/template/exec.go:634
		// _ = "end of CoverTab[29954]"
	} else {
//line /usr/local/go/src/text/template/exec.go:635
		_go_fuzz_dep_.CoverTab[29957]++
//line /usr/local/go/src/text/template/exec.go:635
		// _ = "end of CoverTab[29957]"
//line /usr/local/go/src/text/template/exec.go:635
	}
//line /usr/local/go/src/text/template/exec.go:635
	// _ = "end of CoverTab[29947]"
//line /usr/local/go/src/text/template/exec.go:635
	_go_fuzz_dep_.CoverTab[29948]++
							typ := receiver.Type()
							receiver, isNil := indirect(receiver)
							if receiver.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/text/template/exec.go:638
		_go_fuzz_dep_.CoverTab[29958]++
//line /usr/local/go/src/text/template/exec.go:638
		return isNil
//line /usr/local/go/src/text/template/exec.go:638
		// _ = "end of CoverTab[29958]"
//line /usr/local/go/src/text/template/exec.go:638
	}() {
//line /usr/local/go/src/text/template/exec.go:638
		_go_fuzz_dep_.CoverTab[29959]++

//line /usr/local/go/src/text/template/exec.go:641
		s.errorf("nil pointer evaluating %s.%s", typ, fieldName)
								return zero
//line /usr/local/go/src/text/template/exec.go:642
		// _ = "end of CoverTab[29959]"
	} else {
//line /usr/local/go/src/text/template/exec.go:643
		_go_fuzz_dep_.CoverTab[29960]++
//line /usr/local/go/src/text/template/exec.go:643
		// _ = "end of CoverTab[29960]"
//line /usr/local/go/src/text/template/exec.go:643
	}
//line /usr/local/go/src/text/template/exec.go:643
	// _ = "end of CoverTab[29948]"
//line /usr/local/go/src/text/template/exec.go:643
	_go_fuzz_dep_.CoverTab[29949]++

//line /usr/local/go/src/text/template/exec.go:647
	ptr := receiver
	if ptr.Kind() != reflect.Interface && func() bool {
//line /usr/local/go/src/text/template/exec.go:648
		_go_fuzz_dep_.CoverTab[29961]++
//line /usr/local/go/src/text/template/exec.go:648
		return ptr.Kind() != reflect.Pointer
//line /usr/local/go/src/text/template/exec.go:648
		// _ = "end of CoverTab[29961]"
//line /usr/local/go/src/text/template/exec.go:648
	}() && func() bool {
//line /usr/local/go/src/text/template/exec.go:648
		_go_fuzz_dep_.CoverTab[29962]++
//line /usr/local/go/src/text/template/exec.go:648
		return ptr.CanAddr()
//line /usr/local/go/src/text/template/exec.go:648
		// _ = "end of CoverTab[29962]"
//line /usr/local/go/src/text/template/exec.go:648
	}() {
//line /usr/local/go/src/text/template/exec.go:648
		_go_fuzz_dep_.CoverTab[29963]++
								ptr = ptr.Addr()
//line /usr/local/go/src/text/template/exec.go:649
		// _ = "end of CoverTab[29963]"
	} else {
//line /usr/local/go/src/text/template/exec.go:650
		_go_fuzz_dep_.CoverTab[29964]++
//line /usr/local/go/src/text/template/exec.go:650
		// _ = "end of CoverTab[29964]"
//line /usr/local/go/src/text/template/exec.go:650
	}
//line /usr/local/go/src/text/template/exec.go:650
	// _ = "end of CoverTab[29949]"
//line /usr/local/go/src/text/template/exec.go:650
	_go_fuzz_dep_.CoverTab[29950]++
							if method := ptr.MethodByName(fieldName); method.IsValid() {
//line /usr/local/go/src/text/template/exec.go:651
		_go_fuzz_dep_.CoverTab[29965]++
								return s.evalCall(dot, method, false, node, fieldName, args, final)
//line /usr/local/go/src/text/template/exec.go:652
		// _ = "end of CoverTab[29965]"
	} else {
//line /usr/local/go/src/text/template/exec.go:653
		_go_fuzz_dep_.CoverTab[29966]++
//line /usr/local/go/src/text/template/exec.go:653
		// _ = "end of CoverTab[29966]"
//line /usr/local/go/src/text/template/exec.go:653
	}
//line /usr/local/go/src/text/template/exec.go:653
	// _ = "end of CoverTab[29950]"
//line /usr/local/go/src/text/template/exec.go:653
	_go_fuzz_dep_.CoverTab[29951]++
							hasArgs := len(args) > 1 || func() bool {
//line /usr/local/go/src/text/template/exec.go:654
		_go_fuzz_dep_.CoverTab[29967]++
//line /usr/local/go/src/text/template/exec.go:654
		return !isMissing(final)
//line /usr/local/go/src/text/template/exec.go:654
		// _ = "end of CoverTab[29967]"
//line /usr/local/go/src/text/template/exec.go:654
	}()

							switch receiver.Kind() {
	case reflect.Struct:
//line /usr/local/go/src/text/template/exec.go:657
		_go_fuzz_dep_.CoverTab[29968]++
								tField, ok := receiver.Type().FieldByName(fieldName)
								if ok {
//line /usr/local/go/src/text/template/exec.go:659
			_go_fuzz_dep_.CoverTab[29973]++
									field, err := receiver.FieldByIndexErr(tField.Index)
									if !tField.IsExported() {
//line /usr/local/go/src/text/template/exec.go:661
				_go_fuzz_dep_.CoverTab[29977]++
										s.errorf("%s is an unexported field of struct type %s", fieldName, typ)
//line /usr/local/go/src/text/template/exec.go:662
				// _ = "end of CoverTab[29977]"
			} else {
//line /usr/local/go/src/text/template/exec.go:663
				_go_fuzz_dep_.CoverTab[29978]++
//line /usr/local/go/src/text/template/exec.go:663
				// _ = "end of CoverTab[29978]"
//line /usr/local/go/src/text/template/exec.go:663
			}
//line /usr/local/go/src/text/template/exec.go:663
			// _ = "end of CoverTab[29973]"
//line /usr/local/go/src/text/template/exec.go:663
			_go_fuzz_dep_.CoverTab[29974]++
									if err != nil {
//line /usr/local/go/src/text/template/exec.go:664
				_go_fuzz_dep_.CoverTab[29979]++
										s.errorf("%v", err)
//line /usr/local/go/src/text/template/exec.go:665
				// _ = "end of CoverTab[29979]"
			} else {
//line /usr/local/go/src/text/template/exec.go:666
				_go_fuzz_dep_.CoverTab[29980]++
//line /usr/local/go/src/text/template/exec.go:666
				// _ = "end of CoverTab[29980]"
//line /usr/local/go/src/text/template/exec.go:666
			}
//line /usr/local/go/src/text/template/exec.go:666
			// _ = "end of CoverTab[29974]"
//line /usr/local/go/src/text/template/exec.go:666
			_go_fuzz_dep_.CoverTab[29975]++

									if hasArgs {
//line /usr/local/go/src/text/template/exec.go:668
				_go_fuzz_dep_.CoverTab[29981]++
										s.errorf("%s has arguments but cannot be invoked as function", fieldName)
//line /usr/local/go/src/text/template/exec.go:669
				// _ = "end of CoverTab[29981]"
			} else {
//line /usr/local/go/src/text/template/exec.go:670
				_go_fuzz_dep_.CoverTab[29982]++
//line /usr/local/go/src/text/template/exec.go:670
				// _ = "end of CoverTab[29982]"
//line /usr/local/go/src/text/template/exec.go:670
			}
//line /usr/local/go/src/text/template/exec.go:670
			// _ = "end of CoverTab[29975]"
//line /usr/local/go/src/text/template/exec.go:670
			_go_fuzz_dep_.CoverTab[29976]++
									return field
//line /usr/local/go/src/text/template/exec.go:671
			// _ = "end of CoverTab[29976]"
		} else {
//line /usr/local/go/src/text/template/exec.go:672
			_go_fuzz_dep_.CoverTab[29983]++
//line /usr/local/go/src/text/template/exec.go:672
			// _ = "end of CoverTab[29983]"
//line /usr/local/go/src/text/template/exec.go:672
		}
//line /usr/local/go/src/text/template/exec.go:672
		// _ = "end of CoverTab[29968]"
	case reflect.Map:
//line /usr/local/go/src/text/template/exec.go:673
		_go_fuzz_dep_.CoverTab[29969]++

								nameVal := reflect.ValueOf(fieldName)
								if nameVal.Type().AssignableTo(receiver.Type().Key()) {
//line /usr/local/go/src/text/template/exec.go:676
			_go_fuzz_dep_.CoverTab[29984]++
									if hasArgs {
//line /usr/local/go/src/text/template/exec.go:677
				_go_fuzz_dep_.CoverTab[29987]++
										s.errorf("%s is not a method but has arguments", fieldName)
//line /usr/local/go/src/text/template/exec.go:678
				// _ = "end of CoverTab[29987]"
			} else {
//line /usr/local/go/src/text/template/exec.go:679
				_go_fuzz_dep_.CoverTab[29988]++
//line /usr/local/go/src/text/template/exec.go:679
				// _ = "end of CoverTab[29988]"
//line /usr/local/go/src/text/template/exec.go:679
			}
//line /usr/local/go/src/text/template/exec.go:679
			// _ = "end of CoverTab[29984]"
//line /usr/local/go/src/text/template/exec.go:679
			_go_fuzz_dep_.CoverTab[29985]++
									result := receiver.MapIndex(nameVal)
									if !result.IsValid() {
//line /usr/local/go/src/text/template/exec.go:681
				_go_fuzz_dep_.CoverTab[29989]++
										switch s.tmpl.option.missingKey {
				case mapInvalid:
//line /usr/local/go/src/text/template/exec.go:683
					_go_fuzz_dep_.CoverTab[29990]++
//line /usr/local/go/src/text/template/exec.go:683
					// _ = "end of CoverTab[29990]"

				case mapZeroValue:
//line /usr/local/go/src/text/template/exec.go:685
					_go_fuzz_dep_.CoverTab[29991]++
											result = reflect.Zero(receiver.Type().Elem())
//line /usr/local/go/src/text/template/exec.go:686
					// _ = "end of CoverTab[29991]"
				case mapError:
//line /usr/local/go/src/text/template/exec.go:687
					_go_fuzz_dep_.CoverTab[29992]++
											s.errorf("map has no entry for key %q", fieldName)
//line /usr/local/go/src/text/template/exec.go:688
					// _ = "end of CoverTab[29992]"
//line /usr/local/go/src/text/template/exec.go:688
				default:
//line /usr/local/go/src/text/template/exec.go:688
					_go_fuzz_dep_.CoverTab[29993]++
//line /usr/local/go/src/text/template/exec.go:688
					// _ = "end of CoverTab[29993]"
				}
//line /usr/local/go/src/text/template/exec.go:689
				// _ = "end of CoverTab[29989]"
			} else {
//line /usr/local/go/src/text/template/exec.go:690
				_go_fuzz_dep_.CoverTab[29994]++
//line /usr/local/go/src/text/template/exec.go:690
				// _ = "end of CoverTab[29994]"
//line /usr/local/go/src/text/template/exec.go:690
			}
//line /usr/local/go/src/text/template/exec.go:690
			// _ = "end of CoverTab[29985]"
//line /usr/local/go/src/text/template/exec.go:690
			_go_fuzz_dep_.CoverTab[29986]++
									return result
//line /usr/local/go/src/text/template/exec.go:691
			// _ = "end of CoverTab[29986]"
		} else {
//line /usr/local/go/src/text/template/exec.go:692
			_go_fuzz_dep_.CoverTab[29995]++
//line /usr/local/go/src/text/template/exec.go:692
			// _ = "end of CoverTab[29995]"
//line /usr/local/go/src/text/template/exec.go:692
		}
//line /usr/local/go/src/text/template/exec.go:692
		// _ = "end of CoverTab[29969]"
	case reflect.Pointer:
//line /usr/local/go/src/text/template/exec.go:693
		_go_fuzz_dep_.CoverTab[29970]++
								etyp := receiver.Type().Elem()
								if etyp.Kind() == reflect.Struct {
//line /usr/local/go/src/text/template/exec.go:695
			_go_fuzz_dep_.CoverTab[29996]++
									if _, ok := etyp.FieldByName(fieldName); !ok {
//line /usr/local/go/src/text/template/exec.go:696
				_go_fuzz_dep_.CoverTab[29997]++

//line /usr/local/go/src/text/template/exec.go:699
				break
//line /usr/local/go/src/text/template/exec.go:699
				// _ = "end of CoverTab[29997]"
			} else {
//line /usr/local/go/src/text/template/exec.go:700
				_go_fuzz_dep_.CoverTab[29998]++
//line /usr/local/go/src/text/template/exec.go:700
				// _ = "end of CoverTab[29998]"
//line /usr/local/go/src/text/template/exec.go:700
			}
//line /usr/local/go/src/text/template/exec.go:700
			// _ = "end of CoverTab[29996]"
		} else {
//line /usr/local/go/src/text/template/exec.go:701
			_go_fuzz_dep_.CoverTab[29999]++
//line /usr/local/go/src/text/template/exec.go:701
			// _ = "end of CoverTab[29999]"
//line /usr/local/go/src/text/template/exec.go:701
		}
//line /usr/local/go/src/text/template/exec.go:701
		// _ = "end of CoverTab[29970]"
//line /usr/local/go/src/text/template/exec.go:701
		_go_fuzz_dep_.CoverTab[29971]++
								if isNil {
//line /usr/local/go/src/text/template/exec.go:702
			_go_fuzz_dep_.CoverTab[30000]++
									s.errorf("nil pointer evaluating %s.%s", typ, fieldName)
//line /usr/local/go/src/text/template/exec.go:703
			// _ = "end of CoverTab[30000]"
		} else {
//line /usr/local/go/src/text/template/exec.go:704
			_go_fuzz_dep_.CoverTab[30001]++
//line /usr/local/go/src/text/template/exec.go:704
			// _ = "end of CoverTab[30001]"
//line /usr/local/go/src/text/template/exec.go:704
		}
//line /usr/local/go/src/text/template/exec.go:704
		// _ = "end of CoverTab[29971]"
//line /usr/local/go/src/text/template/exec.go:704
	default:
//line /usr/local/go/src/text/template/exec.go:704
		_go_fuzz_dep_.CoverTab[29972]++
//line /usr/local/go/src/text/template/exec.go:704
		// _ = "end of CoverTab[29972]"
	}
//line /usr/local/go/src/text/template/exec.go:705
	// _ = "end of CoverTab[29951]"
//line /usr/local/go/src/text/template/exec.go:705
	_go_fuzz_dep_.CoverTab[29952]++
							s.errorf("can't evaluate field %s in type %s", fieldName, typ)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:707
	// _ = "end of CoverTab[29952]"
}

var (
	errorType		= reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType		= reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	reflectValueType	= reflect.TypeOf((*reflect.Value)(nil)).Elem()
)

// evalCall executes a function or method call. If it's a method, fun already has the receiver bound, so
//line /usr/local/go/src/text/template/exec.go:716
// it looks just like a function call. The arg list, if non-nil, includes (in the manner of the shell), arg[0]
//line /usr/local/go/src/text/template/exec.go:716
// as the function itself.
//line /usr/local/go/src/text/template/exec.go:719
func (s *state) evalCall(dot, fun reflect.Value, isBuiltin bool, node parse.Node, name string, args []parse.Node, final reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:719
	_go_fuzz_dep_.CoverTab[30002]++
							if args != nil {
//line /usr/local/go/src/text/template/exec.go:720
		_go_fuzz_dep_.CoverTab[30013]++
								args = args[1:]
//line /usr/local/go/src/text/template/exec.go:721
		// _ = "end of CoverTab[30013]"
	} else {
//line /usr/local/go/src/text/template/exec.go:722
		_go_fuzz_dep_.CoverTab[30014]++
//line /usr/local/go/src/text/template/exec.go:722
		// _ = "end of CoverTab[30014]"
//line /usr/local/go/src/text/template/exec.go:722
	}
//line /usr/local/go/src/text/template/exec.go:722
	// _ = "end of CoverTab[30002]"
//line /usr/local/go/src/text/template/exec.go:722
	_go_fuzz_dep_.CoverTab[30003]++
							typ := fun.Type()
							numIn := len(args)
							if !isMissing(final) {
//line /usr/local/go/src/text/template/exec.go:725
		_go_fuzz_dep_.CoverTab[30015]++
								numIn++
//line /usr/local/go/src/text/template/exec.go:726
		// _ = "end of CoverTab[30015]"
	} else {
//line /usr/local/go/src/text/template/exec.go:727
		_go_fuzz_dep_.CoverTab[30016]++
//line /usr/local/go/src/text/template/exec.go:727
		// _ = "end of CoverTab[30016]"
//line /usr/local/go/src/text/template/exec.go:727
	}
//line /usr/local/go/src/text/template/exec.go:727
	// _ = "end of CoverTab[30003]"
//line /usr/local/go/src/text/template/exec.go:727
	_go_fuzz_dep_.CoverTab[30004]++
							numFixed := len(args)
							if typ.IsVariadic() {
//line /usr/local/go/src/text/template/exec.go:729
		_go_fuzz_dep_.CoverTab[30017]++
								numFixed = typ.NumIn() - 1
								if numIn < numFixed {
//line /usr/local/go/src/text/template/exec.go:731
			_go_fuzz_dep_.CoverTab[30018]++
									s.errorf("wrong number of args for %s: want at least %d got %d", name, typ.NumIn()-1, len(args))
//line /usr/local/go/src/text/template/exec.go:732
			// _ = "end of CoverTab[30018]"
		} else {
//line /usr/local/go/src/text/template/exec.go:733
			_go_fuzz_dep_.CoverTab[30019]++
//line /usr/local/go/src/text/template/exec.go:733
			// _ = "end of CoverTab[30019]"
//line /usr/local/go/src/text/template/exec.go:733
		}
//line /usr/local/go/src/text/template/exec.go:733
		// _ = "end of CoverTab[30017]"
	} else {
//line /usr/local/go/src/text/template/exec.go:734
		_go_fuzz_dep_.CoverTab[30020]++
//line /usr/local/go/src/text/template/exec.go:734
		if numIn != typ.NumIn() {
//line /usr/local/go/src/text/template/exec.go:734
			_go_fuzz_dep_.CoverTab[30021]++
									s.errorf("wrong number of args for %s: want %d got %d", name, typ.NumIn(), numIn)
//line /usr/local/go/src/text/template/exec.go:735
			// _ = "end of CoverTab[30021]"
		} else {
//line /usr/local/go/src/text/template/exec.go:736
			_go_fuzz_dep_.CoverTab[30022]++
//line /usr/local/go/src/text/template/exec.go:736
			// _ = "end of CoverTab[30022]"
//line /usr/local/go/src/text/template/exec.go:736
		}
//line /usr/local/go/src/text/template/exec.go:736
		// _ = "end of CoverTab[30020]"
//line /usr/local/go/src/text/template/exec.go:736
	}
//line /usr/local/go/src/text/template/exec.go:736
	// _ = "end of CoverTab[30004]"
//line /usr/local/go/src/text/template/exec.go:736
	_go_fuzz_dep_.CoverTab[30005]++
							if !goodFunc(typ) {
//line /usr/local/go/src/text/template/exec.go:737
		_go_fuzz_dep_.CoverTab[30023]++

								s.errorf("can't call method/function %q with %d results", name, typ.NumOut())
//line /usr/local/go/src/text/template/exec.go:739
		// _ = "end of CoverTab[30023]"
	} else {
//line /usr/local/go/src/text/template/exec.go:740
		_go_fuzz_dep_.CoverTab[30024]++
//line /usr/local/go/src/text/template/exec.go:740
		// _ = "end of CoverTab[30024]"
//line /usr/local/go/src/text/template/exec.go:740
	}
//line /usr/local/go/src/text/template/exec.go:740
	// _ = "end of CoverTab[30005]"
//line /usr/local/go/src/text/template/exec.go:740
	_go_fuzz_dep_.CoverTab[30006]++

							unwrap := func(v reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:742
		_go_fuzz_dep_.CoverTab[30025]++
								if v.Type() == reflectValueType {
//line /usr/local/go/src/text/template/exec.go:743
			_go_fuzz_dep_.CoverTab[30027]++
									v = v.Interface().(reflect.Value)
//line /usr/local/go/src/text/template/exec.go:744
			// _ = "end of CoverTab[30027]"
		} else {
//line /usr/local/go/src/text/template/exec.go:745
			_go_fuzz_dep_.CoverTab[30028]++
//line /usr/local/go/src/text/template/exec.go:745
			// _ = "end of CoverTab[30028]"
//line /usr/local/go/src/text/template/exec.go:745
		}
//line /usr/local/go/src/text/template/exec.go:745
		// _ = "end of CoverTab[30025]"
//line /usr/local/go/src/text/template/exec.go:745
		_go_fuzz_dep_.CoverTab[30026]++
								return v
//line /usr/local/go/src/text/template/exec.go:746
		// _ = "end of CoverTab[30026]"
	}
//line /usr/local/go/src/text/template/exec.go:747
	// _ = "end of CoverTab[30006]"
//line /usr/local/go/src/text/template/exec.go:747
	_go_fuzz_dep_.CoverTab[30007]++

//line /usr/local/go/src/text/template/exec.go:750
	if isBuiltin && func() bool {
//line /usr/local/go/src/text/template/exec.go:750
		_go_fuzz_dep_.CoverTab[30029]++
//line /usr/local/go/src/text/template/exec.go:750
		return (name == "and" || func() bool {
//line /usr/local/go/src/text/template/exec.go:750
			_go_fuzz_dep_.CoverTab[30030]++
//line /usr/local/go/src/text/template/exec.go:750
			return name == "or"
//line /usr/local/go/src/text/template/exec.go:750
			// _ = "end of CoverTab[30030]"
//line /usr/local/go/src/text/template/exec.go:750
		}())
//line /usr/local/go/src/text/template/exec.go:750
		// _ = "end of CoverTab[30029]"
//line /usr/local/go/src/text/template/exec.go:750
	}() {
//line /usr/local/go/src/text/template/exec.go:750
		_go_fuzz_dep_.CoverTab[30031]++
								argType := typ.In(0)
								var v reflect.Value
								for _, arg := range args {
//line /usr/local/go/src/text/template/exec.go:753
			_go_fuzz_dep_.CoverTab[30034]++
									v = s.evalArg(dot, argType, arg).Interface().(reflect.Value)
									if truth(v) == (name == "or") {
//line /usr/local/go/src/text/template/exec.go:755
				_go_fuzz_dep_.CoverTab[30035]++

//line /usr/local/go/src/text/template/exec.go:758
				return v
//line /usr/local/go/src/text/template/exec.go:758
				// _ = "end of CoverTab[30035]"
			} else {
//line /usr/local/go/src/text/template/exec.go:759
				_go_fuzz_dep_.CoverTab[30036]++
//line /usr/local/go/src/text/template/exec.go:759
				// _ = "end of CoverTab[30036]"
//line /usr/local/go/src/text/template/exec.go:759
			}
//line /usr/local/go/src/text/template/exec.go:759
			// _ = "end of CoverTab[30034]"
		}
//line /usr/local/go/src/text/template/exec.go:760
		// _ = "end of CoverTab[30031]"
//line /usr/local/go/src/text/template/exec.go:760
		_go_fuzz_dep_.CoverTab[30032]++
								if final != missingVal {
//line /usr/local/go/src/text/template/exec.go:761
			_go_fuzz_dep_.CoverTab[30037]++

//line /usr/local/go/src/text/template/exec.go:768
			v = unwrap(s.validateType(final, argType))
//line /usr/local/go/src/text/template/exec.go:768
			// _ = "end of CoverTab[30037]"
		} else {
//line /usr/local/go/src/text/template/exec.go:769
			_go_fuzz_dep_.CoverTab[30038]++
//line /usr/local/go/src/text/template/exec.go:769
			// _ = "end of CoverTab[30038]"
//line /usr/local/go/src/text/template/exec.go:769
		}
//line /usr/local/go/src/text/template/exec.go:769
		// _ = "end of CoverTab[30032]"
//line /usr/local/go/src/text/template/exec.go:769
		_go_fuzz_dep_.CoverTab[30033]++
								return v
//line /usr/local/go/src/text/template/exec.go:770
		// _ = "end of CoverTab[30033]"
	} else {
//line /usr/local/go/src/text/template/exec.go:771
		_go_fuzz_dep_.CoverTab[30039]++
//line /usr/local/go/src/text/template/exec.go:771
		// _ = "end of CoverTab[30039]"
//line /usr/local/go/src/text/template/exec.go:771
	}
//line /usr/local/go/src/text/template/exec.go:771
	// _ = "end of CoverTab[30007]"
//line /usr/local/go/src/text/template/exec.go:771
	_go_fuzz_dep_.CoverTab[30008]++

//line /usr/local/go/src/text/template/exec.go:774
	argv := make([]reflect.Value, numIn)

	i := 0
	for ; i < numFixed && func() bool {
//line /usr/local/go/src/text/template/exec.go:777
		_go_fuzz_dep_.CoverTab[30040]++
//line /usr/local/go/src/text/template/exec.go:777
		return i < len(args)
//line /usr/local/go/src/text/template/exec.go:777
		// _ = "end of CoverTab[30040]"
//line /usr/local/go/src/text/template/exec.go:777
	}(); i++ {
//line /usr/local/go/src/text/template/exec.go:777
		_go_fuzz_dep_.CoverTab[30041]++
								argv[i] = s.evalArg(dot, typ.In(i), args[i])
//line /usr/local/go/src/text/template/exec.go:778
		// _ = "end of CoverTab[30041]"
	}
//line /usr/local/go/src/text/template/exec.go:779
	// _ = "end of CoverTab[30008]"
//line /usr/local/go/src/text/template/exec.go:779
	_go_fuzz_dep_.CoverTab[30009]++

							if typ.IsVariadic() {
//line /usr/local/go/src/text/template/exec.go:781
		_go_fuzz_dep_.CoverTab[30042]++
								argType := typ.In(typ.NumIn() - 1).Elem()
								for ; i < len(args); i++ {
//line /usr/local/go/src/text/template/exec.go:783
			_go_fuzz_dep_.CoverTab[30043]++
									argv[i] = s.evalArg(dot, argType, args[i])
//line /usr/local/go/src/text/template/exec.go:784
			// _ = "end of CoverTab[30043]"
		}
//line /usr/local/go/src/text/template/exec.go:785
		// _ = "end of CoverTab[30042]"
	} else {
//line /usr/local/go/src/text/template/exec.go:786
		_go_fuzz_dep_.CoverTab[30044]++
//line /usr/local/go/src/text/template/exec.go:786
		// _ = "end of CoverTab[30044]"
//line /usr/local/go/src/text/template/exec.go:786
	}
//line /usr/local/go/src/text/template/exec.go:786
	// _ = "end of CoverTab[30009]"
//line /usr/local/go/src/text/template/exec.go:786
	_go_fuzz_dep_.CoverTab[30010]++

							if !isMissing(final) {
//line /usr/local/go/src/text/template/exec.go:788
		_go_fuzz_dep_.CoverTab[30045]++
								t := typ.In(typ.NumIn() - 1)
								if typ.IsVariadic() {
//line /usr/local/go/src/text/template/exec.go:790
			_go_fuzz_dep_.CoverTab[30047]++
									if numIn-1 < numFixed {
//line /usr/local/go/src/text/template/exec.go:791
				_go_fuzz_dep_.CoverTab[30048]++

//line /usr/local/go/src/text/template/exec.go:794
				t = typ.In(numIn - 1)
//line /usr/local/go/src/text/template/exec.go:794
				// _ = "end of CoverTab[30048]"
			} else {
//line /usr/local/go/src/text/template/exec.go:795
				_go_fuzz_dep_.CoverTab[30049]++

//line /usr/local/go/src/text/template/exec.go:798
				t = t.Elem()
//line /usr/local/go/src/text/template/exec.go:798
				// _ = "end of CoverTab[30049]"
			}
//line /usr/local/go/src/text/template/exec.go:799
			// _ = "end of CoverTab[30047]"
		} else {
//line /usr/local/go/src/text/template/exec.go:800
			_go_fuzz_dep_.CoverTab[30050]++
//line /usr/local/go/src/text/template/exec.go:800
			// _ = "end of CoverTab[30050]"
//line /usr/local/go/src/text/template/exec.go:800
		}
//line /usr/local/go/src/text/template/exec.go:800
		// _ = "end of CoverTab[30045]"
//line /usr/local/go/src/text/template/exec.go:800
		_go_fuzz_dep_.CoverTab[30046]++
								argv[i] = s.validateType(final, t)
//line /usr/local/go/src/text/template/exec.go:801
		// _ = "end of CoverTab[30046]"
	} else {
//line /usr/local/go/src/text/template/exec.go:802
		_go_fuzz_dep_.CoverTab[30051]++
//line /usr/local/go/src/text/template/exec.go:802
		// _ = "end of CoverTab[30051]"
//line /usr/local/go/src/text/template/exec.go:802
	}
//line /usr/local/go/src/text/template/exec.go:802
	// _ = "end of CoverTab[30010]"
//line /usr/local/go/src/text/template/exec.go:802
	_go_fuzz_dep_.CoverTab[30011]++
							v, err := safeCall(fun, argv)

//line /usr/local/go/src/text/template/exec.go:806
	if err != nil {
//line /usr/local/go/src/text/template/exec.go:806
		_go_fuzz_dep_.CoverTab[30052]++
								s.at(node)
								s.errorf("error calling %s: %w", name, err)
//line /usr/local/go/src/text/template/exec.go:808
		// _ = "end of CoverTab[30052]"
	} else {
//line /usr/local/go/src/text/template/exec.go:809
		_go_fuzz_dep_.CoverTab[30053]++
//line /usr/local/go/src/text/template/exec.go:809
		// _ = "end of CoverTab[30053]"
//line /usr/local/go/src/text/template/exec.go:809
	}
//line /usr/local/go/src/text/template/exec.go:809
	// _ = "end of CoverTab[30011]"
//line /usr/local/go/src/text/template/exec.go:809
	_go_fuzz_dep_.CoverTab[30012]++
							return unwrap(v)
//line /usr/local/go/src/text/template/exec.go:810
	// _ = "end of CoverTab[30012]"
}

// canBeNil reports whether an untyped nil can be assigned to the type. See reflect.Zero.
func canBeNil(typ reflect.Type) bool {
//line /usr/local/go/src/text/template/exec.go:814
	_go_fuzz_dep_.CoverTab[30054]++
							switch typ.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
//line /usr/local/go/src/text/template/exec.go:816
		_go_fuzz_dep_.CoverTab[30056]++
								return true
//line /usr/local/go/src/text/template/exec.go:817
		// _ = "end of CoverTab[30056]"
	case reflect.Struct:
//line /usr/local/go/src/text/template/exec.go:818
		_go_fuzz_dep_.CoverTab[30057]++
								return typ == reflectValueType
//line /usr/local/go/src/text/template/exec.go:819
		// _ = "end of CoverTab[30057]"
//line /usr/local/go/src/text/template/exec.go:819
	default:
//line /usr/local/go/src/text/template/exec.go:819
		_go_fuzz_dep_.CoverTab[30058]++
//line /usr/local/go/src/text/template/exec.go:819
		// _ = "end of CoverTab[30058]"
	}
//line /usr/local/go/src/text/template/exec.go:820
	// _ = "end of CoverTab[30054]"
//line /usr/local/go/src/text/template/exec.go:820
	_go_fuzz_dep_.CoverTab[30055]++
							return false
//line /usr/local/go/src/text/template/exec.go:821
	// _ = "end of CoverTab[30055]"
}

// validateType guarantees that the value is valid and assignable to the type.
func (s *state) validateType(value reflect.Value, typ reflect.Type) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:825
	_go_fuzz_dep_.CoverTab[30059]++
							if !value.IsValid() {
//line /usr/local/go/src/text/template/exec.go:826
		_go_fuzz_dep_.CoverTab[30063]++
								if typ == nil {
//line /usr/local/go/src/text/template/exec.go:827
			_go_fuzz_dep_.CoverTab[30066]++

									return reflect.ValueOf(nil)
//line /usr/local/go/src/text/template/exec.go:829
			// _ = "end of CoverTab[30066]"
		} else {
//line /usr/local/go/src/text/template/exec.go:830
			_go_fuzz_dep_.CoverTab[30067]++
//line /usr/local/go/src/text/template/exec.go:830
			// _ = "end of CoverTab[30067]"
//line /usr/local/go/src/text/template/exec.go:830
		}
//line /usr/local/go/src/text/template/exec.go:830
		// _ = "end of CoverTab[30063]"
//line /usr/local/go/src/text/template/exec.go:830
		_go_fuzz_dep_.CoverTab[30064]++
								if canBeNil(typ) {
//line /usr/local/go/src/text/template/exec.go:831
			_go_fuzz_dep_.CoverTab[30068]++

									return reflect.Zero(typ)
//line /usr/local/go/src/text/template/exec.go:833
			// _ = "end of CoverTab[30068]"
		} else {
//line /usr/local/go/src/text/template/exec.go:834
			_go_fuzz_dep_.CoverTab[30069]++
//line /usr/local/go/src/text/template/exec.go:834
			// _ = "end of CoverTab[30069]"
//line /usr/local/go/src/text/template/exec.go:834
		}
//line /usr/local/go/src/text/template/exec.go:834
		// _ = "end of CoverTab[30064]"
//line /usr/local/go/src/text/template/exec.go:834
		_go_fuzz_dep_.CoverTab[30065]++
								s.errorf("invalid value; expected %s", typ)
//line /usr/local/go/src/text/template/exec.go:835
		// _ = "end of CoverTab[30065]"
	} else {
//line /usr/local/go/src/text/template/exec.go:836
		_go_fuzz_dep_.CoverTab[30070]++
//line /usr/local/go/src/text/template/exec.go:836
		// _ = "end of CoverTab[30070]"
//line /usr/local/go/src/text/template/exec.go:836
	}
//line /usr/local/go/src/text/template/exec.go:836
	// _ = "end of CoverTab[30059]"
//line /usr/local/go/src/text/template/exec.go:836
	_go_fuzz_dep_.CoverTab[30060]++
							if typ == reflectValueType && func() bool {
//line /usr/local/go/src/text/template/exec.go:837
		_go_fuzz_dep_.CoverTab[30071]++
//line /usr/local/go/src/text/template/exec.go:837
		return value.Type() != typ
//line /usr/local/go/src/text/template/exec.go:837
		// _ = "end of CoverTab[30071]"
//line /usr/local/go/src/text/template/exec.go:837
	}() {
//line /usr/local/go/src/text/template/exec.go:837
		_go_fuzz_dep_.CoverTab[30072]++
								return reflect.ValueOf(value)
//line /usr/local/go/src/text/template/exec.go:838
		// _ = "end of CoverTab[30072]"
	} else {
//line /usr/local/go/src/text/template/exec.go:839
		_go_fuzz_dep_.CoverTab[30073]++
//line /usr/local/go/src/text/template/exec.go:839
		// _ = "end of CoverTab[30073]"
//line /usr/local/go/src/text/template/exec.go:839
	}
//line /usr/local/go/src/text/template/exec.go:839
	// _ = "end of CoverTab[30060]"
//line /usr/local/go/src/text/template/exec.go:839
	_go_fuzz_dep_.CoverTab[30061]++
							if typ != nil && func() bool {
//line /usr/local/go/src/text/template/exec.go:840
		_go_fuzz_dep_.CoverTab[30074]++
//line /usr/local/go/src/text/template/exec.go:840
		return !value.Type().AssignableTo(typ)
//line /usr/local/go/src/text/template/exec.go:840
		// _ = "end of CoverTab[30074]"
//line /usr/local/go/src/text/template/exec.go:840
	}() {
//line /usr/local/go/src/text/template/exec.go:840
		_go_fuzz_dep_.CoverTab[30075]++
								if value.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/text/template/exec.go:841
			_go_fuzz_dep_.CoverTab[30077]++
//line /usr/local/go/src/text/template/exec.go:841
			return !value.IsNil()
//line /usr/local/go/src/text/template/exec.go:841
			// _ = "end of CoverTab[30077]"
//line /usr/local/go/src/text/template/exec.go:841
		}() {
//line /usr/local/go/src/text/template/exec.go:841
			_go_fuzz_dep_.CoverTab[30078]++
									value = value.Elem()
									if value.Type().AssignableTo(typ) {
//line /usr/local/go/src/text/template/exec.go:843
				_go_fuzz_dep_.CoverTab[30079]++
										return value
//line /usr/local/go/src/text/template/exec.go:844
				// _ = "end of CoverTab[30079]"
			} else {
//line /usr/local/go/src/text/template/exec.go:845
				_go_fuzz_dep_.CoverTab[30080]++
//line /usr/local/go/src/text/template/exec.go:845
				// _ = "end of CoverTab[30080]"
//line /usr/local/go/src/text/template/exec.go:845
			}
//line /usr/local/go/src/text/template/exec.go:845
			// _ = "end of CoverTab[30078]"

		} else {
//line /usr/local/go/src/text/template/exec.go:847
			_go_fuzz_dep_.CoverTab[30081]++
//line /usr/local/go/src/text/template/exec.go:847
			// _ = "end of CoverTab[30081]"
//line /usr/local/go/src/text/template/exec.go:847
		}
//line /usr/local/go/src/text/template/exec.go:847
		// _ = "end of CoverTab[30075]"
//line /usr/local/go/src/text/template/exec.go:847
		_go_fuzz_dep_.CoverTab[30076]++

//line /usr/local/go/src/text/template/exec.go:852
		switch {
		case value.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/text/template/exec.go:853
			_go_fuzz_dep_.CoverTab[30085]++
//line /usr/local/go/src/text/template/exec.go:853
			return value.Type().Elem().AssignableTo(typ)
//line /usr/local/go/src/text/template/exec.go:853
			// _ = "end of CoverTab[30085]"
//line /usr/local/go/src/text/template/exec.go:853
		}():
//line /usr/local/go/src/text/template/exec.go:853
			_go_fuzz_dep_.CoverTab[30082]++
									value = value.Elem()
									if !value.IsValid() {
//line /usr/local/go/src/text/template/exec.go:855
				_go_fuzz_dep_.CoverTab[30086]++
										s.errorf("dereference of nil pointer of type %s", typ)
//line /usr/local/go/src/text/template/exec.go:856
				// _ = "end of CoverTab[30086]"
			} else {
//line /usr/local/go/src/text/template/exec.go:857
				_go_fuzz_dep_.CoverTab[30087]++
//line /usr/local/go/src/text/template/exec.go:857
				// _ = "end of CoverTab[30087]"
//line /usr/local/go/src/text/template/exec.go:857
			}
//line /usr/local/go/src/text/template/exec.go:857
			// _ = "end of CoverTab[30082]"
		case reflect.PointerTo(value.Type()).AssignableTo(typ) && func() bool {
//line /usr/local/go/src/text/template/exec.go:858
			_go_fuzz_dep_.CoverTab[30088]++
//line /usr/local/go/src/text/template/exec.go:858
			return value.CanAddr()
//line /usr/local/go/src/text/template/exec.go:858
			// _ = "end of CoverTab[30088]"
//line /usr/local/go/src/text/template/exec.go:858
		}():
//line /usr/local/go/src/text/template/exec.go:858
			_go_fuzz_dep_.CoverTab[30083]++
									value = value.Addr()
//line /usr/local/go/src/text/template/exec.go:859
			// _ = "end of CoverTab[30083]"
		default:
//line /usr/local/go/src/text/template/exec.go:860
			_go_fuzz_dep_.CoverTab[30084]++
									s.errorf("wrong type for value; expected %s; got %s", typ, value.Type())
//line /usr/local/go/src/text/template/exec.go:861
			// _ = "end of CoverTab[30084]"
		}
//line /usr/local/go/src/text/template/exec.go:862
		// _ = "end of CoverTab[30076]"
	} else {
//line /usr/local/go/src/text/template/exec.go:863
		_go_fuzz_dep_.CoverTab[30089]++
//line /usr/local/go/src/text/template/exec.go:863
		// _ = "end of CoverTab[30089]"
//line /usr/local/go/src/text/template/exec.go:863
	}
//line /usr/local/go/src/text/template/exec.go:863
	// _ = "end of CoverTab[30061]"
//line /usr/local/go/src/text/template/exec.go:863
	_go_fuzz_dep_.CoverTab[30062]++
							return value
//line /usr/local/go/src/text/template/exec.go:864
	// _ = "end of CoverTab[30062]"
}

func (s *state) evalArg(dot reflect.Value, typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:867
	_go_fuzz_dep_.CoverTab[30090]++
							s.at(n)
							switch arg := n.(type) {
	case *parse.DotNode:
//line /usr/local/go/src/text/template/exec.go:870
		_go_fuzz_dep_.CoverTab[30093]++
								return s.validateType(dot, typ)
//line /usr/local/go/src/text/template/exec.go:871
		// _ = "end of CoverTab[30093]"
	case *parse.NilNode:
//line /usr/local/go/src/text/template/exec.go:872
		_go_fuzz_dep_.CoverTab[30094]++
								if canBeNil(typ) {
//line /usr/local/go/src/text/template/exec.go:873
			_go_fuzz_dep_.CoverTab[30101]++
									return reflect.Zero(typ)
//line /usr/local/go/src/text/template/exec.go:874
			// _ = "end of CoverTab[30101]"
		} else {
//line /usr/local/go/src/text/template/exec.go:875
			_go_fuzz_dep_.CoverTab[30102]++
//line /usr/local/go/src/text/template/exec.go:875
			// _ = "end of CoverTab[30102]"
//line /usr/local/go/src/text/template/exec.go:875
		}
//line /usr/local/go/src/text/template/exec.go:875
		// _ = "end of CoverTab[30094]"
//line /usr/local/go/src/text/template/exec.go:875
		_go_fuzz_dep_.CoverTab[30095]++
								s.errorf("cannot assign nil to %s", typ)
//line /usr/local/go/src/text/template/exec.go:876
		// _ = "end of CoverTab[30095]"
	case *parse.FieldNode:
//line /usr/local/go/src/text/template/exec.go:877
		_go_fuzz_dep_.CoverTab[30096]++
								return s.validateType(s.evalFieldNode(dot, arg, []parse.Node{n}, missingVal), typ)
//line /usr/local/go/src/text/template/exec.go:878
		// _ = "end of CoverTab[30096]"
	case *parse.VariableNode:
//line /usr/local/go/src/text/template/exec.go:879
		_go_fuzz_dep_.CoverTab[30097]++
								return s.validateType(s.evalVariableNode(dot, arg, nil, missingVal), typ)
//line /usr/local/go/src/text/template/exec.go:880
		// _ = "end of CoverTab[30097]"
	case *parse.PipeNode:
//line /usr/local/go/src/text/template/exec.go:881
		_go_fuzz_dep_.CoverTab[30098]++
								return s.validateType(s.evalPipeline(dot, arg), typ)
//line /usr/local/go/src/text/template/exec.go:882
		// _ = "end of CoverTab[30098]"
	case *parse.IdentifierNode:
//line /usr/local/go/src/text/template/exec.go:883
		_go_fuzz_dep_.CoverTab[30099]++
								return s.validateType(s.evalFunction(dot, arg, arg, nil, missingVal), typ)
//line /usr/local/go/src/text/template/exec.go:884
		// _ = "end of CoverTab[30099]"
	case *parse.ChainNode:
//line /usr/local/go/src/text/template/exec.go:885
		_go_fuzz_dep_.CoverTab[30100]++
								return s.validateType(s.evalChainNode(dot, arg, nil, missingVal), typ)
//line /usr/local/go/src/text/template/exec.go:886
		// _ = "end of CoverTab[30100]"
	}
//line /usr/local/go/src/text/template/exec.go:887
	// _ = "end of CoverTab[30090]"
//line /usr/local/go/src/text/template/exec.go:887
	_go_fuzz_dep_.CoverTab[30091]++
							switch typ.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/text/template/exec.go:889
		_go_fuzz_dep_.CoverTab[30103]++
								return s.evalBool(typ, n)
//line /usr/local/go/src/text/template/exec.go:890
		// _ = "end of CoverTab[30103]"
	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/text/template/exec.go:891
		_go_fuzz_dep_.CoverTab[30104]++
								return s.evalComplex(typ, n)
//line /usr/local/go/src/text/template/exec.go:892
		// _ = "end of CoverTab[30104]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/text/template/exec.go:893
		_go_fuzz_dep_.CoverTab[30105]++
								return s.evalFloat(typ, n)
//line /usr/local/go/src/text/template/exec.go:894
		// _ = "end of CoverTab[30105]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/text/template/exec.go:895
		_go_fuzz_dep_.CoverTab[30106]++
								return s.evalInteger(typ, n)
//line /usr/local/go/src/text/template/exec.go:896
		// _ = "end of CoverTab[30106]"
	case reflect.Interface:
//line /usr/local/go/src/text/template/exec.go:897
		_go_fuzz_dep_.CoverTab[30107]++
								if typ.NumMethod() == 0 {
//line /usr/local/go/src/text/template/exec.go:898
			_go_fuzz_dep_.CoverTab[30112]++
									return s.evalEmptyInterface(dot, n)
//line /usr/local/go/src/text/template/exec.go:899
			// _ = "end of CoverTab[30112]"
		} else {
//line /usr/local/go/src/text/template/exec.go:900
			_go_fuzz_dep_.CoverTab[30113]++
//line /usr/local/go/src/text/template/exec.go:900
			// _ = "end of CoverTab[30113]"
//line /usr/local/go/src/text/template/exec.go:900
		}
//line /usr/local/go/src/text/template/exec.go:900
		// _ = "end of CoverTab[30107]"
	case reflect.Struct:
//line /usr/local/go/src/text/template/exec.go:901
		_go_fuzz_dep_.CoverTab[30108]++
								if typ == reflectValueType {
//line /usr/local/go/src/text/template/exec.go:902
			_go_fuzz_dep_.CoverTab[30114]++
									return reflect.ValueOf(s.evalEmptyInterface(dot, n))
//line /usr/local/go/src/text/template/exec.go:903
			// _ = "end of CoverTab[30114]"
		} else {
//line /usr/local/go/src/text/template/exec.go:904
			_go_fuzz_dep_.CoverTab[30115]++
//line /usr/local/go/src/text/template/exec.go:904
			// _ = "end of CoverTab[30115]"
//line /usr/local/go/src/text/template/exec.go:904
		}
//line /usr/local/go/src/text/template/exec.go:904
		// _ = "end of CoverTab[30108]"
	case reflect.String:
//line /usr/local/go/src/text/template/exec.go:905
		_go_fuzz_dep_.CoverTab[30109]++
								return s.evalString(typ, n)
//line /usr/local/go/src/text/template/exec.go:906
		// _ = "end of CoverTab[30109]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/text/template/exec.go:907
		_go_fuzz_dep_.CoverTab[30110]++
								return s.evalUnsignedInteger(typ, n)
//line /usr/local/go/src/text/template/exec.go:908
		// _ = "end of CoverTab[30110]"
//line /usr/local/go/src/text/template/exec.go:908
	default:
//line /usr/local/go/src/text/template/exec.go:908
		_go_fuzz_dep_.CoverTab[30111]++
//line /usr/local/go/src/text/template/exec.go:908
		// _ = "end of CoverTab[30111]"
	}
//line /usr/local/go/src/text/template/exec.go:909
	// _ = "end of CoverTab[30091]"
//line /usr/local/go/src/text/template/exec.go:909
	_go_fuzz_dep_.CoverTab[30092]++
							s.errorf("can't handle %s for arg of type %s", n, typ)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:911
	// _ = "end of CoverTab[30092]"
}

func (s *state) evalBool(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:914
	_go_fuzz_dep_.CoverTab[30116]++
							s.at(n)
							if n, ok := n.(*parse.BoolNode); ok {
//line /usr/local/go/src/text/template/exec.go:916
		_go_fuzz_dep_.CoverTab[30118]++
								value := reflect.New(typ).Elem()
								value.SetBool(n.True)
								return value
//line /usr/local/go/src/text/template/exec.go:919
		// _ = "end of CoverTab[30118]"
	} else {
//line /usr/local/go/src/text/template/exec.go:920
		_go_fuzz_dep_.CoverTab[30119]++
//line /usr/local/go/src/text/template/exec.go:920
		// _ = "end of CoverTab[30119]"
//line /usr/local/go/src/text/template/exec.go:920
	}
//line /usr/local/go/src/text/template/exec.go:920
	// _ = "end of CoverTab[30116]"
//line /usr/local/go/src/text/template/exec.go:920
	_go_fuzz_dep_.CoverTab[30117]++
							s.errorf("expected bool; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:922
	// _ = "end of CoverTab[30117]"
}

func (s *state) evalString(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:925
	_go_fuzz_dep_.CoverTab[30120]++
							s.at(n)
							if n, ok := n.(*parse.StringNode); ok {
//line /usr/local/go/src/text/template/exec.go:927
		_go_fuzz_dep_.CoverTab[30122]++
								value := reflect.New(typ).Elem()
								value.SetString(n.Text)
								return value
//line /usr/local/go/src/text/template/exec.go:930
		// _ = "end of CoverTab[30122]"
	} else {
//line /usr/local/go/src/text/template/exec.go:931
		_go_fuzz_dep_.CoverTab[30123]++
//line /usr/local/go/src/text/template/exec.go:931
		// _ = "end of CoverTab[30123]"
//line /usr/local/go/src/text/template/exec.go:931
	}
//line /usr/local/go/src/text/template/exec.go:931
	// _ = "end of CoverTab[30120]"
//line /usr/local/go/src/text/template/exec.go:931
	_go_fuzz_dep_.CoverTab[30121]++
							s.errorf("expected string; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:933
	// _ = "end of CoverTab[30121]"
}

func (s *state) evalInteger(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:936
	_go_fuzz_dep_.CoverTab[30124]++
							s.at(n)
							if n, ok := n.(*parse.NumberNode); ok && func() bool {
//line /usr/local/go/src/text/template/exec.go:938
		_go_fuzz_dep_.CoverTab[30126]++
//line /usr/local/go/src/text/template/exec.go:938
		return n.IsInt
//line /usr/local/go/src/text/template/exec.go:938
		// _ = "end of CoverTab[30126]"
//line /usr/local/go/src/text/template/exec.go:938
	}() {
//line /usr/local/go/src/text/template/exec.go:938
		_go_fuzz_dep_.CoverTab[30127]++
								value := reflect.New(typ).Elem()
								value.SetInt(n.Int64)
								return value
//line /usr/local/go/src/text/template/exec.go:941
		// _ = "end of CoverTab[30127]"
	} else {
//line /usr/local/go/src/text/template/exec.go:942
		_go_fuzz_dep_.CoverTab[30128]++
//line /usr/local/go/src/text/template/exec.go:942
		// _ = "end of CoverTab[30128]"
//line /usr/local/go/src/text/template/exec.go:942
	}
//line /usr/local/go/src/text/template/exec.go:942
	// _ = "end of CoverTab[30124]"
//line /usr/local/go/src/text/template/exec.go:942
	_go_fuzz_dep_.CoverTab[30125]++
							s.errorf("expected integer; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:944
	// _ = "end of CoverTab[30125]"
}

func (s *state) evalUnsignedInteger(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:947
	_go_fuzz_dep_.CoverTab[30129]++
							s.at(n)
							if n, ok := n.(*parse.NumberNode); ok && func() bool {
//line /usr/local/go/src/text/template/exec.go:949
		_go_fuzz_dep_.CoverTab[30131]++
//line /usr/local/go/src/text/template/exec.go:949
		return n.IsUint
//line /usr/local/go/src/text/template/exec.go:949
		// _ = "end of CoverTab[30131]"
//line /usr/local/go/src/text/template/exec.go:949
	}() {
//line /usr/local/go/src/text/template/exec.go:949
		_go_fuzz_dep_.CoverTab[30132]++
								value := reflect.New(typ).Elem()
								value.SetUint(n.Uint64)
								return value
//line /usr/local/go/src/text/template/exec.go:952
		// _ = "end of CoverTab[30132]"
	} else {
//line /usr/local/go/src/text/template/exec.go:953
		_go_fuzz_dep_.CoverTab[30133]++
//line /usr/local/go/src/text/template/exec.go:953
		// _ = "end of CoverTab[30133]"
//line /usr/local/go/src/text/template/exec.go:953
	}
//line /usr/local/go/src/text/template/exec.go:953
	// _ = "end of CoverTab[30129]"
//line /usr/local/go/src/text/template/exec.go:953
	_go_fuzz_dep_.CoverTab[30130]++
							s.errorf("expected unsigned integer; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:955
	// _ = "end of CoverTab[30130]"
}

func (s *state) evalFloat(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:958
	_go_fuzz_dep_.CoverTab[30134]++
							s.at(n)
							if n, ok := n.(*parse.NumberNode); ok && func() bool {
//line /usr/local/go/src/text/template/exec.go:960
		_go_fuzz_dep_.CoverTab[30136]++
//line /usr/local/go/src/text/template/exec.go:960
		return n.IsFloat
//line /usr/local/go/src/text/template/exec.go:960
		// _ = "end of CoverTab[30136]"
//line /usr/local/go/src/text/template/exec.go:960
	}() {
//line /usr/local/go/src/text/template/exec.go:960
		_go_fuzz_dep_.CoverTab[30137]++
								value := reflect.New(typ).Elem()
								value.SetFloat(n.Float64)
								return value
//line /usr/local/go/src/text/template/exec.go:963
		// _ = "end of CoverTab[30137]"
	} else {
//line /usr/local/go/src/text/template/exec.go:964
		_go_fuzz_dep_.CoverTab[30138]++
//line /usr/local/go/src/text/template/exec.go:964
		// _ = "end of CoverTab[30138]"
//line /usr/local/go/src/text/template/exec.go:964
	}
//line /usr/local/go/src/text/template/exec.go:964
	// _ = "end of CoverTab[30134]"
//line /usr/local/go/src/text/template/exec.go:964
	_go_fuzz_dep_.CoverTab[30135]++
							s.errorf("expected float; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:966
	// _ = "end of CoverTab[30135]"
}

func (s *state) evalComplex(typ reflect.Type, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:969
	_go_fuzz_dep_.CoverTab[30139]++
							if n, ok := n.(*parse.NumberNode); ok && func() bool {
//line /usr/local/go/src/text/template/exec.go:970
		_go_fuzz_dep_.CoverTab[30141]++
//line /usr/local/go/src/text/template/exec.go:970
		return n.IsComplex
//line /usr/local/go/src/text/template/exec.go:970
		// _ = "end of CoverTab[30141]"
//line /usr/local/go/src/text/template/exec.go:970
	}() {
//line /usr/local/go/src/text/template/exec.go:970
		_go_fuzz_dep_.CoverTab[30142]++
								value := reflect.New(typ).Elem()
								value.SetComplex(n.Complex128)
								return value
//line /usr/local/go/src/text/template/exec.go:973
		// _ = "end of CoverTab[30142]"
	} else {
//line /usr/local/go/src/text/template/exec.go:974
		_go_fuzz_dep_.CoverTab[30143]++
//line /usr/local/go/src/text/template/exec.go:974
		// _ = "end of CoverTab[30143]"
//line /usr/local/go/src/text/template/exec.go:974
	}
//line /usr/local/go/src/text/template/exec.go:974
	// _ = "end of CoverTab[30139]"
//line /usr/local/go/src/text/template/exec.go:974
	_go_fuzz_dep_.CoverTab[30140]++
							s.errorf("expected complex; found %s", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:976
	// _ = "end of CoverTab[30140]"
}

func (s *state) evalEmptyInterface(dot reflect.Value, n parse.Node) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:979
	_go_fuzz_dep_.CoverTab[30144]++
							s.at(n)
							switch n := n.(type) {
	case *parse.BoolNode:
//line /usr/local/go/src/text/template/exec.go:982
		_go_fuzz_dep_.CoverTab[30146]++
								return reflect.ValueOf(n.True)
//line /usr/local/go/src/text/template/exec.go:983
		// _ = "end of CoverTab[30146]"
	case *parse.DotNode:
//line /usr/local/go/src/text/template/exec.go:984
		_go_fuzz_dep_.CoverTab[30147]++
								return dot
//line /usr/local/go/src/text/template/exec.go:985
		// _ = "end of CoverTab[30147]"
	case *parse.FieldNode:
//line /usr/local/go/src/text/template/exec.go:986
		_go_fuzz_dep_.CoverTab[30148]++
								return s.evalFieldNode(dot, n, nil, missingVal)
//line /usr/local/go/src/text/template/exec.go:987
		// _ = "end of CoverTab[30148]"
	case *parse.IdentifierNode:
//line /usr/local/go/src/text/template/exec.go:988
		_go_fuzz_dep_.CoverTab[30149]++
								return s.evalFunction(dot, n, n, nil, missingVal)
//line /usr/local/go/src/text/template/exec.go:989
		// _ = "end of CoverTab[30149]"
	case *parse.NilNode:
//line /usr/local/go/src/text/template/exec.go:990
		_go_fuzz_dep_.CoverTab[30150]++

								s.errorf("evalEmptyInterface: nil (can't happen)")
//line /usr/local/go/src/text/template/exec.go:992
		// _ = "end of CoverTab[30150]"
	case *parse.NumberNode:
//line /usr/local/go/src/text/template/exec.go:993
		_go_fuzz_dep_.CoverTab[30151]++
								return s.idealConstant(n)
//line /usr/local/go/src/text/template/exec.go:994
		// _ = "end of CoverTab[30151]"
	case *parse.StringNode:
//line /usr/local/go/src/text/template/exec.go:995
		_go_fuzz_dep_.CoverTab[30152]++
								return reflect.ValueOf(n.Text)
//line /usr/local/go/src/text/template/exec.go:996
		// _ = "end of CoverTab[30152]"
	case *parse.VariableNode:
//line /usr/local/go/src/text/template/exec.go:997
		_go_fuzz_dep_.CoverTab[30153]++
								return s.evalVariableNode(dot, n, nil, missingVal)
//line /usr/local/go/src/text/template/exec.go:998
		// _ = "end of CoverTab[30153]"
	case *parse.PipeNode:
//line /usr/local/go/src/text/template/exec.go:999
		_go_fuzz_dep_.CoverTab[30154]++
								return s.evalPipeline(dot, n)
//line /usr/local/go/src/text/template/exec.go:1000
		// _ = "end of CoverTab[30154]"
	}
//line /usr/local/go/src/text/template/exec.go:1001
	// _ = "end of CoverTab[30144]"
//line /usr/local/go/src/text/template/exec.go:1001
	_go_fuzz_dep_.CoverTab[30145]++
							s.errorf("can't handle assignment of %s to empty interface argument", n)
							panic("not reached")
//line /usr/local/go/src/text/template/exec.go:1003
	// _ = "end of CoverTab[30145]"
}

// indirect returns the item at the end of indirection, and a bool to indicate
//line /usr/local/go/src/text/template/exec.go:1006
// if it's nil. If the returned bool is true, the returned value's kind will be
//line /usr/local/go/src/text/template/exec.go:1006
// either a pointer or interface.
//line /usr/local/go/src/text/template/exec.go:1009
func indirect(v reflect.Value) (rv reflect.Value, isNil bool) {
//line /usr/local/go/src/text/template/exec.go:1009
	_go_fuzz_dep_.CoverTab[30155]++
							for ; v.Kind() == reflect.Pointer || func() bool {
//line /usr/local/go/src/text/template/exec.go:1010
		_go_fuzz_dep_.CoverTab[30157]++
//line /usr/local/go/src/text/template/exec.go:1010
		return v.Kind() == reflect.Interface
//line /usr/local/go/src/text/template/exec.go:1010
		// _ = "end of CoverTab[30157]"
//line /usr/local/go/src/text/template/exec.go:1010
	}(); v = v.Elem() {
//line /usr/local/go/src/text/template/exec.go:1010
		_go_fuzz_dep_.CoverTab[30158]++
								if v.IsNil() {
//line /usr/local/go/src/text/template/exec.go:1011
			_go_fuzz_dep_.CoverTab[30159]++
									return v, true
//line /usr/local/go/src/text/template/exec.go:1012
			// _ = "end of CoverTab[30159]"
		} else {
//line /usr/local/go/src/text/template/exec.go:1013
			_go_fuzz_dep_.CoverTab[30160]++
//line /usr/local/go/src/text/template/exec.go:1013
			// _ = "end of CoverTab[30160]"
//line /usr/local/go/src/text/template/exec.go:1013
		}
//line /usr/local/go/src/text/template/exec.go:1013
		// _ = "end of CoverTab[30158]"
	}
//line /usr/local/go/src/text/template/exec.go:1014
	// _ = "end of CoverTab[30155]"
//line /usr/local/go/src/text/template/exec.go:1014
	_go_fuzz_dep_.CoverTab[30156]++
							return v, false
//line /usr/local/go/src/text/template/exec.go:1015
	// _ = "end of CoverTab[30156]"
}

// indirectInterface returns the concrete value in an interface value,
//line /usr/local/go/src/text/template/exec.go:1018
// or else the zero reflect.Value.
//line /usr/local/go/src/text/template/exec.go:1018
// That is, if v represents the interface value x, the result is the same as reflect.ValueOf(x):
//line /usr/local/go/src/text/template/exec.go:1018
// the fact that x was an interface value is forgotten.
//line /usr/local/go/src/text/template/exec.go:1022
func indirectInterface(v reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/exec.go:1022
	_go_fuzz_dep_.CoverTab[30161]++
							if v.Kind() != reflect.Interface {
//line /usr/local/go/src/text/template/exec.go:1023
		_go_fuzz_dep_.CoverTab[30164]++
								return v
//line /usr/local/go/src/text/template/exec.go:1024
		// _ = "end of CoverTab[30164]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1025
		_go_fuzz_dep_.CoverTab[30165]++
//line /usr/local/go/src/text/template/exec.go:1025
		// _ = "end of CoverTab[30165]"
//line /usr/local/go/src/text/template/exec.go:1025
	}
//line /usr/local/go/src/text/template/exec.go:1025
	// _ = "end of CoverTab[30161]"
//line /usr/local/go/src/text/template/exec.go:1025
	_go_fuzz_dep_.CoverTab[30162]++
							if v.IsNil() {
//line /usr/local/go/src/text/template/exec.go:1026
		_go_fuzz_dep_.CoverTab[30166]++
								return reflect.Value{}
//line /usr/local/go/src/text/template/exec.go:1027
		// _ = "end of CoverTab[30166]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1028
		_go_fuzz_dep_.CoverTab[30167]++
//line /usr/local/go/src/text/template/exec.go:1028
		// _ = "end of CoverTab[30167]"
//line /usr/local/go/src/text/template/exec.go:1028
	}
//line /usr/local/go/src/text/template/exec.go:1028
	// _ = "end of CoverTab[30162]"
//line /usr/local/go/src/text/template/exec.go:1028
	_go_fuzz_dep_.CoverTab[30163]++
							return v.Elem()
//line /usr/local/go/src/text/template/exec.go:1029
	// _ = "end of CoverTab[30163]"
}

// printValue writes the textual representation of the value to the output of
//line /usr/local/go/src/text/template/exec.go:1032
// the template.
//line /usr/local/go/src/text/template/exec.go:1034
func (s *state) printValue(n parse.Node, v reflect.Value) {
//line /usr/local/go/src/text/template/exec.go:1034
	_go_fuzz_dep_.CoverTab[30168]++
							s.at(n)
							iface, ok := printableValue(v)
							if !ok {
//line /usr/local/go/src/text/template/exec.go:1037
		_go_fuzz_dep_.CoverTab[30170]++
								s.errorf("can't print %s of type %s", n, v.Type())
//line /usr/local/go/src/text/template/exec.go:1038
		// _ = "end of CoverTab[30170]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1039
		_go_fuzz_dep_.CoverTab[30171]++
//line /usr/local/go/src/text/template/exec.go:1039
		// _ = "end of CoverTab[30171]"
//line /usr/local/go/src/text/template/exec.go:1039
	}
//line /usr/local/go/src/text/template/exec.go:1039
	// _ = "end of CoverTab[30168]"
//line /usr/local/go/src/text/template/exec.go:1039
	_go_fuzz_dep_.CoverTab[30169]++
							_, err := fmt.Fprint(s.wr, iface)
							if err != nil {
//line /usr/local/go/src/text/template/exec.go:1041
		_go_fuzz_dep_.CoverTab[30172]++
								s.writeError(err)
//line /usr/local/go/src/text/template/exec.go:1042
		// _ = "end of CoverTab[30172]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1043
		_go_fuzz_dep_.CoverTab[30173]++
//line /usr/local/go/src/text/template/exec.go:1043
		// _ = "end of CoverTab[30173]"
//line /usr/local/go/src/text/template/exec.go:1043
	}
//line /usr/local/go/src/text/template/exec.go:1043
	// _ = "end of CoverTab[30169]"
}

// printableValue returns the, possibly indirected, interface value inside v that
//line /usr/local/go/src/text/template/exec.go:1046
// is best for a call to formatted printer.
//line /usr/local/go/src/text/template/exec.go:1048
func printableValue(v reflect.Value) (any, bool) {
//line /usr/local/go/src/text/template/exec.go:1048
	_go_fuzz_dep_.CoverTab[30174]++
							if v.Kind() == reflect.Pointer {
//line /usr/local/go/src/text/template/exec.go:1049
		_go_fuzz_dep_.CoverTab[30178]++
								v, _ = indirect(v)
//line /usr/local/go/src/text/template/exec.go:1050
		// _ = "end of CoverTab[30178]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1051
		_go_fuzz_dep_.CoverTab[30179]++
//line /usr/local/go/src/text/template/exec.go:1051
		// _ = "end of CoverTab[30179]"
//line /usr/local/go/src/text/template/exec.go:1051
	}
//line /usr/local/go/src/text/template/exec.go:1051
	// _ = "end of CoverTab[30174]"
//line /usr/local/go/src/text/template/exec.go:1051
	_go_fuzz_dep_.CoverTab[30175]++
							if !v.IsValid() {
//line /usr/local/go/src/text/template/exec.go:1052
		_go_fuzz_dep_.CoverTab[30180]++
								return "<no value>", true
//line /usr/local/go/src/text/template/exec.go:1053
		// _ = "end of CoverTab[30180]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1054
		_go_fuzz_dep_.CoverTab[30181]++
//line /usr/local/go/src/text/template/exec.go:1054
		// _ = "end of CoverTab[30181]"
//line /usr/local/go/src/text/template/exec.go:1054
	}
//line /usr/local/go/src/text/template/exec.go:1054
	// _ = "end of CoverTab[30175]"
//line /usr/local/go/src/text/template/exec.go:1054
	_go_fuzz_dep_.CoverTab[30176]++

							if !v.Type().Implements(errorType) && func() bool {
//line /usr/local/go/src/text/template/exec.go:1056
		_go_fuzz_dep_.CoverTab[30182]++
//line /usr/local/go/src/text/template/exec.go:1056
		return !v.Type().Implements(fmtStringerType)
//line /usr/local/go/src/text/template/exec.go:1056
		// _ = "end of CoverTab[30182]"
//line /usr/local/go/src/text/template/exec.go:1056
	}() {
//line /usr/local/go/src/text/template/exec.go:1056
		_go_fuzz_dep_.CoverTab[30183]++
								if v.CanAddr() && func() bool {
//line /usr/local/go/src/text/template/exec.go:1057
			_go_fuzz_dep_.CoverTab[30184]++
//line /usr/local/go/src/text/template/exec.go:1057
			return (reflect.PointerTo(v.Type()).Implements(errorType) || func() bool {
//line /usr/local/go/src/text/template/exec.go:1057
				_go_fuzz_dep_.CoverTab[30185]++
//line /usr/local/go/src/text/template/exec.go:1057
				return reflect.PointerTo(v.Type()).Implements(fmtStringerType)
//line /usr/local/go/src/text/template/exec.go:1057
				// _ = "end of CoverTab[30185]"
//line /usr/local/go/src/text/template/exec.go:1057
			}())
//line /usr/local/go/src/text/template/exec.go:1057
			// _ = "end of CoverTab[30184]"
//line /usr/local/go/src/text/template/exec.go:1057
		}() {
//line /usr/local/go/src/text/template/exec.go:1057
			_go_fuzz_dep_.CoverTab[30186]++
									v = v.Addr()
//line /usr/local/go/src/text/template/exec.go:1058
			// _ = "end of CoverTab[30186]"
		} else {
//line /usr/local/go/src/text/template/exec.go:1059
			_go_fuzz_dep_.CoverTab[30187]++
									switch v.Kind() {
			case reflect.Chan, reflect.Func:
//line /usr/local/go/src/text/template/exec.go:1061
				_go_fuzz_dep_.CoverTab[30188]++
										return nil, false
//line /usr/local/go/src/text/template/exec.go:1062
				// _ = "end of CoverTab[30188]"
//line /usr/local/go/src/text/template/exec.go:1062
			default:
//line /usr/local/go/src/text/template/exec.go:1062
				_go_fuzz_dep_.CoverTab[30189]++
//line /usr/local/go/src/text/template/exec.go:1062
				// _ = "end of CoverTab[30189]"
			}
//line /usr/local/go/src/text/template/exec.go:1063
			// _ = "end of CoverTab[30187]"
		}
//line /usr/local/go/src/text/template/exec.go:1064
		// _ = "end of CoverTab[30183]"
	} else {
//line /usr/local/go/src/text/template/exec.go:1065
		_go_fuzz_dep_.CoverTab[30190]++
//line /usr/local/go/src/text/template/exec.go:1065
		// _ = "end of CoverTab[30190]"
//line /usr/local/go/src/text/template/exec.go:1065
	}
//line /usr/local/go/src/text/template/exec.go:1065
	// _ = "end of CoverTab[30176]"
//line /usr/local/go/src/text/template/exec.go:1065
	_go_fuzz_dep_.CoverTab[30177]++
							return v.Interface(), true
//line /usr/local/go/src/text/template/exec.go:1066
	// _ = "end of CoverTab[30177]"
}

//line /usr/local/go/src/text/template/exec.go:1067
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/exec.go:1067
var _ = _go_fuzz_dep_.CoverTab
