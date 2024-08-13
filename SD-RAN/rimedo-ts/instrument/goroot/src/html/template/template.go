// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/template.go:5
package template

//line /usr/local/go/src/html/template/template.go:5
import (
//line /usr/local/go/src/html/template/template.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/template.go:5
)
//line /usr/local/go/src/html/template/template.go:5
import (
//line /usr/local/go/src/html/template/template.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/template.go:5
)

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
	"text/template"
	"text/template/parse"
)

// Template is a specialized Template from "text/template" that produces a safe
//line /usr/local/go/src/html/template/template.go:19
// HTML document fragment.
//line /usr/local/go/src/html/template/template.go:21
type Template struct {
	// Sticky error if escaping fails, or escapeOK if succeeded.
	escapeErr	error
	// We could embed the text/template field, but it's safer not to because
	// we need to keep our version of the name space and the underlying
	// template's in sync.
	text	*template.Template
	// The underlying template's parse tree, updated to be HTML-safe.
	Tree		*parse.Tree
	*nameSpace	// common to all associated templates
}

// escapeOK is a sentinel value used to indicate valid escaping.
var escapeOK = fmt.Errorf("template escaped correctly")

// nameSpace is the data structure shared by all templates in an association.
type nameSpace struct {
	mu	sync.Mutex
	set	map[string]*Template
	escaped	bool
	esc	escaper
}

// Templates returns a slice of the templates associated with t, including t
//line /usr/local/go/src/html/template/template.go:44
// itself.
//line /usr/local/go/src/html/template/template.go:46
func (t *Template) Templates() []*Template {
//line /usr/local/go/src/html/template/template.go:46
	_go_fuzz_dep_.CoverTab[31437]++
							ns := t.nameSpace
							ns.mu.Lock()
							defer ns.mu.Unlock()

							m := make([]*Template, 0, len(ns.set))
							for _, v := range ns.set {
//line /usr/local/go/src/html/template/template.go:52
		_go_fuzz_dep_.CoverTab[31439]++
								m = append(m, v)
//line /usr/local/go/src/html/template/template.go:53
		// _ = "end of CoverTab[31439]"
	}
//line /usr/local/go/src/html/template/template.go:54
	// _ = "end of CoverTab[31437]"
//line /usr/local/go/src/html/template/template.go:54
	_go_fuzz_dep_.CoverTab[31438]++
							return m
//line /usr/local/go/src/html/template/template.go:55
	// _ = "end of CoverTab[31438]"
}

// Option sets options for the template. Options are described by
//line /usr/local/go/src/html/template/template.go:58
// strings, either a simple string or "key=value". There can be at
//line /usr/local/go/src/html/template/template.go:58
// most one equals sign in an option string. If the option string
//line /usr/local/go/src/html/template/template.go:58
// is unrecognized or otherwise invalid, Option panics.
//line /usr/local/go/src/html/template/template.go:58
//
//line /usr/local/go/src/html/template/template.go:58
// Known options:
//line /usr/local/go/src/html/template/template.go:58
//
//line /usr/local/go/src/html/template/template.go:58
// missingkey: Control the behavior during execution if a map is
//line /usr/local/go/src/html/template/template.go:58
// indexed with a key that is not present in the map.
//line /usr/local/go/src/html/template/template.go:58
//
//line /usr/local/go/src/html/template/template.go:58
//	"missingkey=default" or "missingkey=invalid"
//line /usr/local/go/src/html/template/template.go:58
//		The default behavior: Do nothing and continue execution.
//line /usr/local/go/src/html/template/template.go:58
//		If printed, the result of the index operation is the string
//line /usr/local/go/src/html/template/template.go:58
//		"<no value>".
//line /usr/local/go/src/html/template/template.go:58
//	"missingkey=zero"
//line /usr/local/go/src/html/template/template.go:58
//		The operation returns the zero value for the map type's element.
//line /usr/local/go/src/html/template/template.go:58
//	"missingkey=error"
//line /usr/local/go/src/html/template/template.go:58
//		Execution stops immediately with an error.
//line /usr/local/go/src/html/template/template.go:76
func (t *Template) Option(opt ...string) *Template {
//line /usr/local/go/src/html/template/template.go:76
	_go_fuzz_dep_.CoverTab[31440]++
							t.text.Option(opt...)
							return t
//line /usr/local/go/src/html/template/template.go:78
	// _ = "end of CoverTab[31440]"
}

// checkCanParse checks whether it is OK to parse templates.
//line /usr/local/go/src/html/template/template.go:81
// If not, it returns an error.
//line /usr/local/go/src/html/template/template.go:83
func (t *Template) checkCanParse() error {
//line /usr/local/go/src/html/template/template.go:83
	_go_fuzz_dep_.CoverTab[31441]++
							if t == nil {
//line /usr/local/go/src/html/template/template.go:84
		_go_fuzz_dep_.CoverTab[31444]++
								return nil
//line /usr/local/go/src/html/template/template.go:85
		// _ = "end of CoverTab[31444]"
	} else {
//line /usr/local/go/src/html/template/template.go:86
		_go_fuzz_dep_.CoverTab[31445]++
//line /usr/local/go/src/html/template/template.go:86
		// _ = "end of CoverTab[31445]"
//line /usr/local/go/src/html/template/template.go:86
	}
//line /usr/local/go/src/html/template/template.go:86
	// _ = "end of CoverTab[31441]"
//line /usr/local/go/src/html/template/template.go:86
	_go_fuzz_dep_.CoverTab[31442]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							if t.nameSpace.escaped {
//line /usr/local/go/src/html/template/template.go:89
		_go_fuzz_dep_.CoverTab[31446]++
								return fmt.Errorf("html/template: cannot Parse after Execute")
//line /usr/local/go/src/html/template/template.go:90
		// _ = "end of CoverTab[31446]"
	} else {
//line /usr/local/go/src/html/template/template.go:91
		_go_fuzz_dep_.CoverTab[31447]++
//line /usr/local/go/src/html/template/template.go:91
		// _ = "end of CoverTab[31447]"
//line /usr/local/go/src/html/template/template.go:91
	}
//line /usr/local/go/src/html/template/template.go:91
	// _ = "end of CoverTab[31442]"
//line /usr/local/go/src/html/template/template.go:91
	_go_fuzz_dep_.CoverTab[31443]++
							return nil
//line /usr/local/go/src/html/template/template.go:92
	// _ = "end of CoverTab[31443]"
}

// escape escapes all associated templates.
func (t *Template) escape() error {
//line /usr/local/go/src/html/template/template.go:96
	_go_fuzz_dep_.CoverTab[31448]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							t.nameSpace.escaped = true
							if t.escapeErr == nil {
//line /usr/local/go/src/html/template/template.go:100
		_go_fuzz_dep_.CoverTab[31450]++
								if t.Tree == nil {
//line /usr/local/go/src/html/template/template.go:101
			_go_fuzz_dep_.CoverTab[31452]++
									return fmt.Errorf("template: %q is an incomplete or empty template", t.Name())
//line /usr/local/go/src/html/template/template.go:102
			// _ = "end of CoverTab[31452]"
		} else {
//line /usr/local/go/src/html/template/template.go:103
			_go_fuzz_dep_.CoverTab[31453]++
//line /usr/local/go/src/html/template/template.go:103
			// _ = "end of CoverTab[31453]"
//line /usr/local/go/src/html/template/template.go:103
		}
//line /usr/local/go/src/html/template/template.go:103
		// _ = "end of CoverTab[31450]"
//line /usr/local/go/src/html/template/template.go:103
		_go_fuzz_dep_.CoverTab[31451]++
								if err := escapeTemplate(t, t.text.Root, t.Name()); err != nil {
//line /usr/local/go/src/html/template/template.go:104
			_go_fuzz_dep_.CoverTab[31454]++
									return err
//line /usr/local/go/src/html/template/template.go:105
			// _ = "end of CoverTab[31454]"
		} else {
//line /usr/local/go/src/html/template/template.go:106
			_go_fuzz_dep_.CoverTab[31455]++
//line /usr/local/go/src/html/template/template.go:106
			// _ = "end of CoverTab[31455]"
//line /usr/local/go/src/html/template/template.go:106
		}
//line /usr/local/go/src/html/template/template.go:106
		// _ = "end of CoverTab[31451]"
	} else {
//line /usr/local/go/src/html/template/template.go:107
		_go_fuzz_dep_.CoverTab[31456]++
//line /usr/local/go/src/html/template/template.go:107
		if t.escapeErr != escapeOK {
//line /usr/local/go/src/html/template/template.go:107
			_go_fuzz_dep_.CoverTab[31457]++
									return t.escapeErr
//line /usr/local/go/src/html/template/template.go:108
			// _ = "end of CoverTab[31457]"
		} else {
//line /usr/local/go/src/html/template/template.go:109
			_go_fuzz_dep_.CoverTab[31458]++
//line /usr/local/go/src/html/template/template.go:109
			// _ = "end of CoverTab[31458]"
//line /usr/local/go/src/html/template/template.go:109
		}
//line /usr/local/go/src/html/template/template.go:109
		// _ = "end of CoverTab[31456]"
//line /usr/local/go/src/html/template/template.go:109
	}
//line /usr/local/go/src/html/template/template.go:109
	// _ = "end of CoverTab[31448]"
//line /usr/local/go/src/html/template/template.go:109
	_go_fuzz_dep_.CoverTab[31449]++
							return nil
//line /usr/local/go/src/html/template/template.go:110
	// _ = "end of CoverTab[31449]"
}

// Execute applies a parsed template to the specified data object,
//line /usr/local/go/src/html/template/template.go:113
// writing the output to wr.
//line /usr/local/go/src/html/template/template.go:113
// If an error occurs executing the template or writing its output,
//line /usr/local/go/src/html/template/template.go:113
// execution stops, but partial results may already have been written to
//line /usr/local/go/src/html/template/template.go:113
// the output writer.
//line /usr/local/go/src/html/template/template.go:113
// A template may be executed safely in parallel, although if parallel
//line /usr/local/go/src/html/template/template.go:113
// executions share a Writer the output may be interleaved.
//line /usr/local/go/src/html/template/template.go:120
func (t *Template) Execute(wr io.Writer, data any) error {
//line /usr/local/go/src/html/template/template.go:120
	_go_fuzz_dep_.CoverTab[31459]++
							if err := t.escape(); err != nil {
//line /usr/local/go/src/html/template/template.go:121
		_go_fuzz_dep_.CoverTab[31461]++
								return err
//line /usr/local/go/src/html/template/template.go:122
		// _ = "end of CoverTab[31461]"
	} else {
//line /usr/local/go/src/html/template/template.go:123
		_go_fuzz_dep_.CoverTab[31462]++
//line /usr/local/go/src/html/template/template.go:123
		// _ = "end of CoverTab[31462]"
//line /usr/local/go/src/html/template/template.go:123
	}
//line /usr/local/go/src/html/template/template.go:123
	// _ = "end of CoverTab[31459]"
//line /usr/local/go/src/html/template/template.go:123
	_go_fuzz_dep_.CoverTab[31460]++
							return t.text.Execute(wr, data)
//line /usr/local/go/src/html/template/template.go:124
	// _ = "end of CoverTab[31460]"
}

// ExecuteTemplate applies the template associated with t that has the given
//line /usr/local/go/src/html/template/template.go:127
// name to the specified data object and writes the output to wr.
//line /usr/local/go/src/html/template/template.go:127
// If an error occurs executing the template or writing its output,
//line /usr/local/go/src/html/template/template.go:127
// execution stops, but partial results may already have been written to
//line /usr/local/go/src/html/template/template.go:127
// the output writer.
//line /usr/local/go/src/html/template/template.go:127
// A template may be executed safely in parallel, although if parallel
//line /usr/local/go/src/html/template/template.go:127
// executions share a Writer the output may be interleaved.
//line /usr/local/go/src/html/template/template.go:134
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error {
//line /usr/local/go/src/html/template/template.go:134
	_go_fuzz_dep_.CoverTab[31463]++
							tmpl, err := t.lookupAndEscapeTemplate(name)
							if err != nil {
//line /usr/local/go/src/html/template/template.go:136
		_go_fuzz_dep_.CoverTab[31465]++
								return err
//line /usr/local/go/src/html/template/template.go:137
		// _ = "end of CoverTab[31465]"
	} else {
//line /usr/local/go/src/html/template/template.go:138
		_go_fuzz_dep_.CoverTab[31466]++
//line /usr/local/go/src/html/template/template.go:138
		// _ = "end of CoverTab[31466]"
//line /usr/local/go/src/html/template/template.go:138
	}
//line /usr/local/go/src/html/template/template.go:138
	// _ = "end of CoverTab[31463]"
//line /usr/local/go/src/html/template/template.go:138
	_go_fuzz_dep_.CoverTab[31464]++
							return tmpl.text.Execute(wr, data)
//line /usr/local/go/src/html/template/template.go:139
	// _ = "end of CoverTab[31464]"
}

// lookupAndEscapeTemplate guarantees that the template with the given name
//line /usr/local/go/src/html/template/template.go:142
// is escaped, or returns an error if it cannot be. It returns the named
//line /usr/local/go/src/html/template/template.go:142
// template.
//line /usr/local/go/src/html/template/template.go:145
func (t *Template) lookupAndEscapeTemplate(name string) (tmpl *Template, err error) {
//line /usr/local/go/src/html/template/template.go:145
	_go_fuzz_dep_.CoverTab[31467]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							t.nameSpace.escaped = true
							tmpl = t.set[name]
							if tmpl == nil {
//line /usr/local/go/src/html/template/template.go:150
		_go_fuzz_dep_.CoverTab[31473]++
								return nil, fmt.Errorf("html/template: %q is undefined", name)
//line /usr/local/go/src/html/template/template.go:151
		// _ = "end of CoverTab[31473]"
	} else {
//line /usr/local/go/src/html/template/template.go:152
		_go_fuzz_dep_.CoverTab[31474]++
//line /usr/local/go/src/html/template/template.go:152
		// _ = "end of CoverTab[31474]"
//line /usr/local/go/src/html/template/template.go:152
	}
//line /usr/local/go/src/html/template/template.go:152
	// _ = "end of CoverTab[31467]"
//line /usr/local/go/src/html/template/template.go:152
	_go_fuzz_dep_.CoverTab[31468]++
							if tmpl.escapeErr != nil && func() bool {
//line /usr/local/go/src/html/template/template.go:153
		_go_fuzz_dep_.CoverTab[31475]++
//line /usr/local/go/src/html/template/template.go:153
		return tmpl.escapeErr != escapeOK
//line /usr/local/go/src/html/template/template.go:153
		// _ = "end of CoverTab[31475]"
//line /usr/local/go/src/html/template/template.go:153
	}() {
//line /usr/local/go/src/html/template/template.go:153
		_go_fuzz_dep_.CoverTab[31476]++
								return nil, tmpl.escapeErr
//line /usr/local/go/src/html/template/template.go:154
		// _ = "end of CoverTab[31476]"
	} else {
//line /usr/local/go/src/html/template/template.go:155
		_go_fuzz_dep_.CoverTab[31477]++
//line /usr/local/go/src/html/template/template.go:155
		// _ = "end of CoverTab[31477]"
//line /usr/local/go/src/html/template/template.go:155
	}
//line /usr/local/go/src/html/template/template.go:155
	// _ = "end of CoverTab[31468]"
//line /usr/local/go/src/html/template/template.go:155
	_go_fuzz_dep_.CoverTab[31469]++
							if tmpl.text.Tree == nil || func() bool {
//line /usr/local/go/src/html/template/template.go:156
		_go_fuzz_dep_.CoverTab[31478]++
//line /usr/local/go/src/html/template/template.go:156
		return tmpl.text.Root == nil
//line /usr/local/go/src/html/template/template.go:156
		// _ = "end of CoverTab[31478]"
//line /usr/local/go/src/html/template/template.go:156
	}() {
//line /usr/local/go/src/html/template/template.go:156
		_go_fuzz_dep_.CoverTab[31479]++
								return nil, fmt.Errorf("html/template: %q is an incomplete template", name)
//line /usr/local/go/src/html/template/template.go:157
		// _ = "end of CoverTab[31479]"
	} else {
//line /usr/local/go/src/html/template/template.go:158
		_go_fuzz_dep_.CoverTab[31480]++
//line /usr/local/go/src/html/template/template.go:158
		// _ = "end of CoverTab[31480]"
//line /usr/local/go/src/html/template/template.go:158
	}
//line /usr/local/go/src/html/template/template.go:158
	// _ = "end of CoverTab[31469]"
//line /usr/local/go/src/html/template/template.go:158
	_go_fuzz_dep_.CoverTab[31470]++
							if t.text.Lookup(name) == nil {
//line /usr/local/go/src/html/template/template.go:159
		_go_fuzz_dep_.CoverTab[31481]++
								panic("html/template internal error: template escaping out of sync")
//line /usr/local/go/src/html/template/template.go:160
		// _ = "end of CoverTab[31481]"
	} else {
//line /usr/local/go/src/html/template/template.go:161
		_go_fuzz_dep_.CoverTab[31482]++
//line /usr/local/go/src/html/template/template.go:161
		// _ = "end of CoverTab[31482]"
//line /usr/local/go/src/html/template/template.go:161
	}
//line /usr/local/go/src/html/template/template.go:161
	// _ = "end of CoverTab[31470]"
//line /usr/local/go/src/html/template/template.go:161
	_go_fuzz_dep_.CoverTab[31471]++
							if tmpl.escapeErr == nil {
//line /usr/local/go/src/html/template/template.go:162
		_go_fuzz_dep_.CoverTab[31483]++
								err = escapeTemplate(tmpl, tmpl.text.Root, name)
//line /usr/local/go/src/html/template/template.go:163
		// _ = "end of CoverTab[31483]"
	} else {
//line /usr/local/go/src/html/template/template.go:164
		_go_fuzz_dep_.CoverTab[31484]++
//line /usr/local/go/src/html/template/template.go:164
		// _ = "end of CoverTab[31484]"
//line /usr/local/go/src/html/template/template.go:164
	}
//line /usr/local/go/src/html/template/template.go:164
	// _ = "end of CoverTab[31471]"
//line /usr/local/go/src/html/template/template.go:164
	_go_fuzz_dep_.CoverTab[31472]++
							return tmpl, err
//line /usr/local/go/src/html/template/template.go:165
	// _ = "end of CoverTab[31472]"
}

// DefinedTemplates returns a string listing the defined templates,
//line /usr/local/go/src/html/template/template.go:168
// prefixed by the string "; defined templates are: ". If there are none,
//line /usr/local/go/src/html/template/template.go:168
// it returns the empty string. Used to generate an error message.
//line /usr/local/go/src/html/template/template.go:171
func (t *Template) DefinedTemplates() string {
//line /usr/local/go/src/html/template/template.go:171
	_go_fuzz_dep_.CoverTab[31485]++
							return t.text.DefinedTemplates()
//line /usr/local/go/src/html/template/template.go:172
	// _ = "end of CoverTab[31485]"
}

// Parse parses text as a template body for t.
//line /usr/local/go/src/html/template/template.go:175
// Named template definitions ({{define ...}} or {{block ...}} statements) in text
//line /usr/local/go/src/html/template/template.go:175
// define additional templates associated with t and are removed from the
//line /usr/local/go/src/html/template/template.go:175
// definition of t itself.
//line /usr/local/go/src/html/template/template.go:175
//
//line /usr/local/go/src/html/template/template.go:175
// Templates can be redefined in successive calls to Parse,
//line /usr/local/go/src/html/template/template.go:175
// before the first use of Execute on t or any associated template.
//line /usr/local/go/src/html/template/template.go:175
// A template definition with a body containing only white space and comments
//line /usr/local/go/src/html/template/template.go:175
// is considered empty and will not replace an existing template's body.
//line /usr/local/go/src/html/template/template.go:175
// This allows using Parse to add new named template definitions without
//line /usr/local/go/src/html/template/template.go:175
// overwriting the main template body.
//line /usr/local/go/src/html/template/template.go:186
func (t *Template) Parse(text string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:186
	_go_fuzz_dep_.CoverTab[31486]++
							if err := t.checkCanParse(); err != nil {
//line /usr/local/go/src/html/template/template.go:187
		_go_fuzz_dep_.CoverTab[31490]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:188
		// _ = "end of CoverTab[31490]"
	} else {
//line /usr/local/go/src/html/template/template.go:189
		_go_fuzz_dep_.CoverTab[31491]++
//line /usr/local/go/src/html/template/template.go:189
		// _ = "end of CoverTab[31491]"
//line /usr/local/go/src/html/template/template.go:189
	}
//line /usr/local/go/src/html/template/template.go:189
	// _ = "end of CoverTab[31486]"
//line /usr/local/go/src/html/template/template.go:189
	_go_fuzz_dep_.CoverTab[31487]++

							ret, err := t.text.Parse(text)
							if err != nil {
//line /usr/local/go/src/html/template/template.go:192
		_go_fuzz_dep_.CoverTab[31492]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:193
		// _ = "end of CoverTab[31492]"
	} else {
//line /usr/local/go/src/html/template/template.go:194
		_go_fuzz_dep_.CoverTab[31493]++
//line /usr/local/go/src/html/template/template.go:194
		// _ = "end of CoverTab[31493]"
//line /usr/local/go/src/html/template/template.go:194
	}
//line /usr/local/go/src/html/template/template.go:194
	// _ = "end of CoverTab[31487]"
//line /usr/local/go/src/html/template/template.go:194
	_go_fuzz_dep_.CoverTab[31488]++

//line /usr/local/go/src/html/template/template.go:199
	t.nameSpace.mu.Lock()
	defer t.nameSpace.mu.Unlock()
	for _, v := range ret.Templates() {
//line /usr/local/go/src/html/template/template.go:201
		_go_fuzz_dep_.CoverTab[31494]++
								name := v.Name()
								tmpl := t.set[name]
								if tmpl == nil {
//line /usr/local/go/src/html/template/template.go:204
			_go_fuzz_dep_.CoverTab[31496]++
									tmpl = t.new(name)
//line /usr/local/go/src/html/template/template.go:205
			// _ = "end of CoverTab[31496]"
		} else {
//line /usr/local/go/src/html/template/template.go:206
			_go_fuzz_dep_.CoverTab[31497]++
//line /usr/local/go/src/html/template/template.go:206
			// _ = "end of CoverTab[31497]"
//line /usr/local/go/src/html/template/template.go:206
		}
//line /usr/local/go/src/html/template/template.go:206
		// _ = "end of CoverTab[31494]"
//line /usr/local/go/src/html/template/template.go:206
		_go_fuzz_dep_.CoverTab[31495]++
								tmpl.text = v
								tmpl.Tree = v.Tree
//line /usr/local/go/src/html/template/template.go:208
		// _ = "end of CoverTab[31495]"
	}
//line /usr/local/go/src/html/template/template.go:209
	// _ = "end of CoverTab[31488]"
//line /usr/local/go/src/html/template/template.go:209
	_go_fuzz_dep_.CoverTab[31489]++
							return t, nil
//line /usr/local/go/src/html/template/template.go:210
	// _ = "end of CoverTab[31489]"
}

// AddParseTree creates a new template with the name and parse tree
//line /usr/local/go/src/html/template/template.go:213
// and associates it with t.
//line /usr/local/go/src/html/template/template.go:213
//
//line /usr/local/go/src/html/template/template.go:213
// It returns an error if t or any associated template has already been executed.
//line /usr/local/go/src/html/template/template.go:217
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:217
	_go_fuzz_dep_.CoverTab[31498]++
							if err := t.checkCanParse(); err != nil {
//line /usr/local/go/src/html/template/template.go:218
		_go_fuzz_dep_.CoverTab[31501]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:219
		// _ = "end of CoverTab[31501]"
	} else {
//line /usr/local/go/src/html/template/template.go:220
		_go_fuzz_dep_.CoverTab[31502]++
//line /usr/local/go/src/html/template/template.go:220
		// _ = "end of CoverTab[31502]"
//line /usr/local/go/src/html/template/template.go:220
	}
//line /usr/local/go/src/html/template/template.go:220
	// _ = "end of CoverTab[31498]"
//line /usr/local/go/src/html/template/template.go:220
	_go_fuzz_dep_.CoverTab[31499]++

							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							text, err := t.text.AddParseTree(name, tree)
							if err != nil {
//line /usr/local/go/src/html/template/template.go:225
		_go_fuzz_dep_.CoverTab[31503]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:226
		// _ = "end of CoverTab[31503]"
	} else {
//line /usr/local/go/src/html/template/template.go:227
		_go_fuzz_dep_.CoverTab[31504]++
//line /usr/local/go/src/html/template/template.go:227
		// _ = "end of CoverTab[31504]"
//line /usr/local/go/src/html/template/template.go:227
	}
//line /usr/local/go/src/html/template/template.go:227
	// _ = "end of CoverTab[31499]"
//line /usr/local/go/src/html/template/template.go:227
	_go_fuzz_dep_.CoverTab[31500]++
							ret := &Template{
		nil,
		text,
		text.Tree,
		t.nameSpace,
	}
							t.set[name] = ret
							return ret, nil
//line /usr/local/go/src/html/template/template.go:235
	// _ = "end of CoverTab[31500]"
}

// Clone returns a duplicate of the template, including all associated
//line /usr/local/go/src/html/template/template.go:238
// templates. The actual representation is not copied, but the name space of
//line /usr/local/go/src/html/template/template.go:238
// associated templates is, so further calls to Parse in the copy will add
//line /usr/local/go/src/html/template/template.go:238
// templates to the copy but not to the original. Clone can be used to prepare
//line /usr/local/go/src/html/template/template.go:238
// common templates and use them with variant definitions for other templates
//line /usr/local/go/src/html/template/template.go:238
// by adding the variants after the clone is made.
//line /usr/local/go/src/html/template/template.go:238
//
//line /usr/local/go/src/html/template/template.go:238
// It returns an error if t has already been executed.
//line /usr/local/go/src/html/template/template.go:246
func (t *Template) Clone() (*Template, error) {
//line /usr/local/go/src/html/template/template.go:246
	_go_fuzz_dep_.CoverTab[31505]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							if t.escapeErr != nil {
//line /usr/local/go/src/html/template/template.go:249
		_go_fuzz_dep_.CoverTab[31509]++
								return nil, fmt.Errorf("html/template: cannot Clone %q after it has executed", t.Name())
//line /usr/local/go/src/html/template/template.go:250
		// _ = "end of CoverTab[31509]"
	} else {
//line /usr/local/go/src/html/template/template.go:251
		_go_fuzz_dep_.CoverTab[31510]++
//line /usr/local/go/src/html/template/template.go:251
		// _ = "end of CoverTab[31510]"
//line /usr/local/go/src/html/template/template.go:251
	}
//line /usr/local/go/src/html/template/template.go:251
	// _ = "end of CoverTab[31505]"
//line /usr/local/go/src/html/template/template.go:251
	_go_fuzz_dep_.CoverTab[31506]++
							textClone, err := t.text.Clone()
							if err != nil {
//line /usr/local/go/src/html/template/template.go:253
		_go_fuzz_dep_.CoverTab[31511]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:254
		// _ = "end of CoverTab[31511]"
	} else {
//line /usr/local/go/src/html/template/template.go:255
		_go_fuzz_dep_.CoverTab[31512]++
//line /usr/local/go/src/html/template/template.go:255
		// _ = "end of CoverTab[31512]"
//line /usr/local/go/src/html/template/template.go:255
	}
//line /usr/local/go/src/html/template/template.go:255
	// _ = "end of CoverTab[31506]"
//line /usr/local/go/src/html/template/template.go:255
	_go_fuzz_dep_.CoverTab[31507]++
							ns := &nameSpace{set: make(map[string]*Template)}
							ns.esc = makeEscaper(ns)
							ret := &Template{
		nil,
		textClone,
		textClone.Tree,
		ns,
	}
	ret.set[ret.Name()] = ret
	for _, x := range textClone.Templates() {
//line /usr/local/go/src/html/template/template.go:265
		_go_fuzz_dep_.CoverTab[31513]++
								name := x.Name()
								src := t.set[name]
								if src == nil || func() bool {
//line /usr/local/go/src/html/template/template.go:268
			_go_fuzz_dep_.CoverTab[31515]++
//line /usr/local/go/src/html/template/template.go:268
			return src.escapeErr != nil
//line /usr/local/go/src/html/template/template.go:268
			// _ = "end of CoverTab[31515]"
//line /usr/local/go/src/html/template/template.go:268
		}() {
//line /usr/local/go/src/html/template/template.go:268
			_go_fuzz_dep_.CoverTab[31516]++
									return nil, fmt.Errorf("html/template: cannot Clone %q after it has executed", t.Name())
//line /usr/local/go/src/html/template/template.go:269
			// _ = "end of CoverTab[31516]"
		} else {
//line /usr/local/go/src/html/template/template.go:270
			_go_fuzz_dep_.CoverTab[31517]++
//line /usr/local/go/src/html/template/template.go:270
			// _ = "end of CoverTab[31517]"
//line /usr/local/go/src/html/template/template.go:270
		}
//line /usr/local/go/src/html/template/template.go:270
		// _ = "end of CoverTab[31513]"
//line /usr/local/go/src/html/template/template.go:270
		_go_fuzz_dep_.CoverTab[31514]++
								x.Tree = x.Tree.Copy()
								ret.set[name] = &Template{
			nil,
			x,
			x.Tree,
			ret.nameSpace,
		}
//line /usr/local/go/src/html/template/template.go:277
		// _ = "end of CoverTab[31514]"
	}
//line /usr/local/go/src/html/template/template.go:278
	// _ = "end of CoverTab[31507]"
//line /usr/local/go/src/html/template/template.go:278
	_go_fuzz_dep_.CoverTab[31508]++

							return ret.set[ret.Name()], nil
//line /usr/local/go/src/html/template/template.go:280
	// _ = "end of CoverTab[31508]"
}

// New allocates a new HTML template with the given name.
func New(name string) *Template {
//line /usr/local/go/src/html/template/template.go:284
	_go_fuzz_dep_.CoverTab[31518]++
							ns := &nameSpace{set: make(map[string]*Template)}
							ns.esc = makeEscaper(ns)
							tmpl := &Template{
		nil,
		template.New(name),
		nil,
		ns,
	}
							tmpl.set[name] = tmpl
							return tmpl
//line /usr/local/go/src/html/template/template.go:294
	// _ = "end of CoverTab[31518]"
}

// New allocates a new HTML template associated with the given one
//line /usr/local/go/src/html/template/template.go:297
// and with the same delimiters. The association, which is transitive,
//line /usr/local/go/src/html/template/template.go:297
// allows one template to invoke another with a {{template}} action.
//line /usr/local/go/src/html/template/template.go:297
//
//line /usr/local/go/src/html/template/template.go:297
// If a template with the given name already exists, the new HTML template
//line /usr/local/go/src/html/template/template.go:297
// will replace it. The existing template will be reset and disassociated with
//line /usr/local/go/src/html/template/template.go:297
// t.
//line /usr/local/go/src/html/template/template.go:304
func (t *Template) New(name string) *Template {
//line /usr/local/go/src/html/template/template.go:304
	_go_fuzz_dep_.CoverTab[31519]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							return t.new(name)
//line /usr/local/go/src/html/template/template.go:307
	// _ = "end of CoverTab[31519]"
}

// new is the implementation of New, without the lock.
func (t *Template) new(name string) *Template {
//line /usr/local/go/src/html/template/template.go:311
	_go_fuzz_dep_.CoverTab[31520]++
							tmpl := &Template{
		nil,
		t.text.New(name),
		nil,
		t.nameSpace,
	}
	if existing, ok := tmpl.set[name]; ok {
//line /usr/local/go/src/html/template/template.go:318
		_go_fuzz_dep_.CoverTab[31522]++
								emptyTmpl := New(existing.Name())
								*existing = *emptyTmpl
//line /usr/local/go/src/html/template/template.go:320
		// _ = "end of CoverTab[31522]"
	} else {
//line /usr/local/go/src/html/template/template.go:321
		_go_fuzz_dep_.CoverTab[31523]++
//line /usr/local/go/src/html/template/template.go:321
		// _ = "end of CoverTab[31523]"
//line /usr/local/go/src/html/template/template.go:321
	}
//line /usr/local/go/src/html/template/template.go:321
	// _ = "end of CoverTab[31520]"
//line /usr/local/go/src/html/template/template.go:321
	_go_fuzz_dep_.CoverTab[31521]++
							tmpl.set[name] = tmpl
							return tmpl
//line /usr/local/go/src/html/template/template.go:323
	// _ = "end of CoverTab[31521]"
}

// Name returns the name of the template.
func (t *Template) Name() string {
//line /usr/local/go/src/html/template/template.go:327
	_go_fuzz_dep_.CoverTab[31524]++
							return t.text.Name()
//line /usr/local/go/src/html/template/template.go:328
	// _ = "end of CoverTab[31524]"
}

type FuncMap = template.FuncMap

// Funcs adds the elements of the argument map to the template's function map.
//line /usr/local/go/src/html/template/template.go:333
// It must be called before the template is parsed.
//line /usr/local/go/src/html/template/template.go:333
// It panics if a value in the map is not a function with appropriate return
//line /usr/local/go/src/html/template/template.go:333
// type. However, it is legal to overwrite elements of the map. The return
//line /usr/local/go/src/html/template/template.go:333
// value is the template, so calls can be chained.
//line /usr/local/go/src/html/template/template.go:338
func (t *Template) Funcs(funcMap FuncMap) *Template {
//line /usr/local/go/src/html/template/template.go:338
	_go_fuzz_dep_.CoverTab[31525]++
							t.text.Funcs(template.FuncMap(funcMap))
							return t
//line /usr/local/go/src/html/template/template.go:340
	// _ = "end of CoverTab[31525]"
}

// Delims sets the action delimiters to the specified strings, to be used in
//line /usr/local/go/src/html/template/template.go:343
// subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
//line /usr/local/go/src/html/template/template.go:343
// definitions will inherit the settings. An empty delimiter stands for the
//line /usr/local/go/src/html/template/template.go:343
// corresponding default: {{ or }}.
//line /usr/local/go/src/html/template/template.go:343
// The return value is the template, so calls can be chained.
//line /usr/local/go/src/html/template/template.go:348
func (t *Template) Delims(left, right string) *Template {
//line /usr/local/go/src/html/template/template.go:348
	_go_fuzz_dep_.CoverTab[31526]++
							t.text.Delims(left, right)
							return t
//line /usr/local/go/src/html/template/template.go:350
	// _ = "end of CoverTab[31526]"
}

// Lookup returns the template with the given name that is associated with t,
//line /usr/local/go/src/html/template/template.go:353
// or nil if there is no such template.
//line /usr/local/go/src/html/template/template.go:355
func (t *Template) Lookup(name string) *Template {
//line /usr/local/go/src/html/template/template.go:355
	_go_fuzz_dep_.CoverTab[31527]++
							t.nameSpace.mu.Lock()
							defer t.nameSpace.mu.Unlock()
							return t.set[name]
//line /usr/local/go/src/html/template/template.go:358
	// _ = "end of CoverTab[31527]"
}

// Must is a helper that wraps a call to a function returning (*Template, error)
//line /usr/local/go/src/html/template/template.go:361
// and panics if the error is non-nil. It is intended for use in variable initializations
//line /usr/local/go/src/html/template/template.go:361
// such as
//line /usr/local/go/src/html/template/template.go:361
//
//line /usr/local/go/src/html/template/template.go:361
//	var t = template.Must(template.New("name").Parse("html"))
//line /usr/local/go/src/html/template/template.go:366
func Must(t *Template, err error) *Template {
//line /usr/local/go/src/html/template/template.go:366
	_go_fuzz_dep_.CoverTab[31528]++
							if err != nil {
//line /usr/local/go/src/html/template/template.go:367
		_go_fuzz_dep_.CoverTab[31530]++
								panic(err)
//line /usr/local/go/src/html/template/template.go:368
		// _ = "end of CoverTab[31530]"
	} else {
//line /usr/local/go/src/html/template/template.go:369
		_go_fuzz_dep_.CoverTab[31531]++
//line /usr/local/go/src/html/template/template.go:369
		// _ = "end of CoverTab[31531]"
//line /usr/local/go/src/html/template/template.go:369
	}
//line /usr/local/go/src/html/template/template.go:369
	// _ = "end of CoverTab[31528]"
//line /usr/local/go/src/html/template/template.go:369
	_go_fuzz_dep_.CoverTab[31529]++
							return t
//line /usr/local/go/src/html/template/template.go:370
	// _ = "end of CoverTab[31529]"
}

// ParseFiles creates a new Template and parses the template definitions from
//line /usr/local/go/src/html/template/template.go:373
// the named files. The returned template's name will have the (base) name and
//line /usr/local/go/src/html/template/template.go:373
// (parsed) contents of the first file. There must be at least one file.
//line /usr/local/go/src/html/template/template.go:373
// If an error occurs, parsing stops and the returned *Template is nil.
//line /usr/local/go/src/html/template/template.go:373
//
//line /usr/local/go/src/html/template/template.go:373
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/html/template/template.go:373
// the last one mentioned will be the one that results.
//line /usr/local/go/src/html/template/template.go:373
// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
//line /usr/local/go/src/html/template/template.go:373
// named "foo", while "a/foo" is unavailable.
//line /usr/local/go/src/html/template/template.go:382
func ParseFiles(filenames ...string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:382
	_go_fuzz_dep_.CoverTab[31532]++
							return parseFiles(nil, readFileOS, filenames...)
//line /usr/local/go/src/html/template/template.go:383
	// _ = "end of CoverTab[31532]"
}

// ParseFiles parses the named files and associates the resulting templates with
//line /usr/local/go/src/html/template/template.go:386
// t. If an error occurs, parsing stops and the returned template is nil;
//line /usr/local/go/src/html/template/template.go:386
// otherwise it is t. There must be at least one file.
//line /usr/local/go/src/html/template/template.go:386
//
//line /usr/local/go/src/html/template/template.go:386
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/html/template/template.go:386
// the last one mentioned will be the one that results.
//line /usr/local/go/src/html/template/template.go:386
//
//line /usr/local/go/src/html/template/template.go:386
// ParseFiles returns an error if t or any associated template has already been executed.
//line /usr/local/go/src/html/template/template.go:394
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:394
	_go_fuzz_dep_.CoverTab[31533]++
							return parseFiles(t, readFileOS, filenames...)
//line /usr/local/go/src/html/template/template.go:395
	// _ = "end of CoverTab[31533]"
}

// parseFiles is the helper for the method and function. If the argument
//line /usr/local/go/src/html/template/template.go:398
// template is nil, it is created from the first file.
//line /usr/local/go/src/html/template/template.go:400
func parseFiles(t *Template, readFile func(string) (string, []byte, error), filenames ...string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:400
	_go_fuzz_dep_.CoverTab[31534]++
							if err := t.checkCanParse(); err != nil {
//line /usr/local/go/src/html/template/template.go:401
		_go_fuzz_dep_.CoverTab[31538]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:402
		// _ = "end of CoverTab[31538]"
	} else {
//line /usr/local/go/src/html/template/template.go:403
		_go_fuzz_dep_.CoverTab[31539]++
//line /usr/local/go/src/html/template/template.go:403
		// _ = "end of CoverTab[31539]"
//line /usr/local/go/src/html/template/template.go:403
	}
//line /usr/local/go/src/html/template/template.go:403
	// _ = "end of CoverTab[31534]"
//line /usr/local/go/src/html/template/template.go:403
	_go_fuzz_dep_.CoverTab[31535]++

							if len(filenames) == 0 {
//line /usr/local/go/src/html/template/template.go:405
		_go_fuzz_dep_.CoverTab[31540]++

								return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
//line /usr/local/go/src/html/template/template.go:407
		// _ = "end of CoverTab[31540]"
	} else {
//line /usr/local/go/src/html/template/template.go:408
		_go_fuzz_dep_.CoverTab[31541]++
//line /usr/local/go/src/html/template/template.go:408
		// _ = "end of CoverTab[31541]"
//line /usr/local/go/src/html/template/template.go:408
	}
//line /usr/local/go/src/html/template/template.go:408
	// _ = "end of CoverTab[31535]"
//line /usr/local/go/src/html/template/template.go:408
	_go_fuzz_dep_.CoverTab[31536]++
							for _, filename := range filenames {
//line /usr/local/go/src/html/template/template.go:409
		_go_fuzz_dep_.CoverTab[31542]++
								name, b, err := readFile(filename)
								if err != nil {
//line /usr/local/go/src/html/template/template.go:411
			_go_fuzz_dep_.CoverTab[31546]++
									return nil, err
//line /usr/local/go/src/html/template/template.go:412
			// _ = "end of CoverTab[31546]"
		} else {
//line /usr/local/go/src/html/template/template.go:413
			_go_fuzz_dep_.CoverTab[31547]++
//line /usr/local/go/src/html/template/template.go:413
			// _ = "end of CoverTab[31547]"
//line /usr/local/go/src/html/template/template.go:413
		}
//line /usr/local/go/src/html/template/template.go:413
		// _ = "end of CoverTab[31542]"
//line /usr/local/go/src/html/template/template.go:413
		_go_fuzz_dep_.CoverTab[31543]++
								s := string(b)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *Template
		if t == nil {
//line /usr/local/go/src/html/template/template.go:422
			_go_fuzz_dep_.CoverTab[31548]++
									t = New(name)
//line /usr/local/go/src/html/template/template.go:423
			// _ = "end of CoverTab[31548]"
		} else {
//line /usr/local/go/src/html/template/template.go:424
			_go_fuzz_dep_.CoverTab[31549]++
//line /usr/local/go/src/html/template/template.go:424
			// _ = "end of CoverTab[31549]"
//line /usr/local/go/src/html/template/template.go:424
		}
//line /usr/local/go/src/html/template/template.go:424
		// _ = "end of CoverTab[31543]"
//line /usr/local/go/src/html/template/template.go:424
		_go_fuzz_dep_.CoverTab[31544]++
								if name == t.Name() {
//line /usr/local/go/src/html/template/template.go:425
			_go_fuzz_dep_.CoverTab[31550]++
									tmpl = t
//line /usr/local/go/src/html/template/template.go:426
			// _ = "end of CoverTab[31550]"
		} else {
//line /usr/local/go/src/html/template/template.go:427
			_go_fuzz_dep_.CoverTab[31551]++
									tmpl = t.New(name)
//line /usr/local/go/src/html/template/template.go:428
			// _ = "end of CoverTab[31551]"
		}
//line /usr/local/go/src/html/template/template.go:429
		// _ = "end of CoverTab[31544]"
//line /usr/local/go/src/html/template/template.go:429
		_go_fuzz_dep_.CoverTab[31545]++
								_, err = tmpl.Parse(s)
								if err != nil {
//line /usr/local/go/src/html/template/template.go:431
			_go_fuzz_dep_.CoverTab[31552]++
									return nil, err
//line /usr/local/go/src/html/template/template.go:432
			// _ = "end of CoverTab[31552]"
		} else {
//line /usr/local/go/src/html/template/template.go:433
			_go_fuzz_dep_.CoverTab[31553]++
//line /usr/local/go/src/html/template/template.go:433
			// _ = "end of CoverTab[31553]"
//line /usr/local/go/src/html/template/template.go:433
		}
//line /usr/local/go/src/html/template/template.go:433
		// _ = "end of CoverTab[31545]"
	}
//line /usr/local/go/src/html/template/template.go:434
	// _ = "end of CoverTab[31536]"
//line /usr/local/go/src/html/template/template.go:434
	_go_fuzz_dep_.CoverTab[31537]++
							return t, nil
//line /usr/local/go/src/html/template/template.go:435
	// _ = "end of CoverTab[31537]"
}

// ParseGlob creates a new Template and parses the template definitions from
//line /usr/local/go/src/html/template/template.go:438
// the files identified by the pattern. The files are matched according to the
//line /usr/local/go/src/html/template/template.go:438
// semantics of filepath.Match, and the pattern must match at least one file.
//line /usr/local/go/src/html/template/template.go:438
// The returned template will have the (base) name and (parsed) contents of the
//line /usr/local/go/src/html/template/template.go:438
// first file matched by the pattern. ParseGlob is equivalent to calling
//line /usr/local/go/src/html/template/template.go:438
// ParseFiles with the list of files matched by the pattern.
//line /usr/local/go/src/html/template/template.go:438
//
//line /usr/local/go/src/html/template/template.go:438
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/html/template/template.go:438
// the last one mentioned will be the one that results.
//line /usr/local/go/src/html/template/template.go:447
func ParseGlob(pattern string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:447
	_go_fuzz_dep_.CoverTab[31554]++
							return parseGlob(nil, pattern)
//line /usr/local/go/src/html/template/template.go:448
	// _ = "end of CoverTab[31554]"
}

// ParseGlob parses the template definitions in the files identified by the
//line /usr/local/go/src/html/template/template.go:451
// pattern and associates the resulting templates with t. The files are matched
//line /usr/local/go/src/html/template/template.go:451
// according to the semantics of filepath.Match, and the pattern must match at
//line /usr/local/go/src/html/template/template.go:451
// least one file. ParseGlob is equivalent to calling t.ParseFiles with the
//line /usr/local/go/src/html/template/template.go:451
// list of files matched by the pattern.
//line /usr/local/go/src/html/template/template.go:451
//
//line /usr/local/go/src/html/template/template.go:451
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/html/template/template.go:451
// the last one mentioned will be the one that results.
//line /usr/local/go/src/html/template/template.go:451
//
//line /usr/local/go/src/html/template/template.go:451
// ParseGlob returns an error if t or any associated template has already been executed.
//line /usr/local/go/src/html/template/template.go:461
func (t *Template) ParseGlob(pattern string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:461
	_go_fuzz_dep_.CoverTab[31555]++
							return parseGlob(t, pattern)
//line /usr/local/go/src/html/template/template.go:462
	// _ = "end of CoverTab[31555]"
}

// parseGlob is the implementation of the function and method ParseGlob.
func parseGlob(t *Template, pattern string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:466
	_go_fuzz_dep_.CoverTab[31556]++
							if err := t.checkCanParse(); err != nil {
//line /usr/local/go/src/html/template/template.go:467
		_go_fuzz_dep_.CoverTab[31560]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:468
		// _ = "end of CoverTab[31560]"
	} else {
//line /usr/local/go/src/html/template/template.go:469
		_go_fuzz_dep_.CoverTab[31561]++
//line /usr/local/go/src/html/template/template.go:469
		// _ = "end of CoverTab[31561]"
//line /usr/local/go/src/html/template/template.go:469
	}
//line /usr/local/go/src/html/template/template.go:469
	// _ = "end of CoverTab[31556]"
//line /usr/local/go/src/html/template/template.go:469
	_go_fuzz_dep_.CoverTab[31557]++
							filenames, err := filepath.Glob(pattern)
							if err != nil {
//line /usr/local/go/src/html/template/template.go:471
		_go_fuzz_dep_.CoverTab[31562]++
								return nil, err
//line /usr/local/go/src/html/template/template.go:472
		// _ = "end of CoverTab[31562]"
	} else {
//line /usr/local/go/src/html/template/template.go:473
		_go_fuzz_dep_.CoverTab[31563]++
//line /usr/local/go/src/html/template/template.go:473
		// _ = "end of CoverTab[31563]"
//line /usr/local/go/src/html/template/template.go:473
	}
//line /usr/local/go/src/html/template/template.go:473
	// _ = "end of CoverTab[31557]"
//line /usr/local/go/src/html/template/template.go:473
	_go_fuzz_dep_.CoverTab[31558]++
							if len(filenames) == 0 {
//line /usr/local/go/src/html/template/template.go:474
		_go_fuzz_dep_.CoverTab[31564]++
								return nil, fmt.Errorf("html/template: pattern matches no files: %#q", pattern)
//line /usr/local/go/src/html/template/template.go:475
		// _ = "end of CoverTab[31564]"
	} else {
//line /usr/local/go/src/html/template/template.go:476
		_go_fuzz_dep_.CoverTab[31565]++
//line /usr/local/go/src/html/template/template.go:476
		// _ = "end of CoverTab[31565]"
//line /usr/local/go/src/html/template/template.go:476
	}
//line /usr/local/go/src/html/template/template.go:476
	// _ = "end of CoverTab[31558]"
//line /usr/local/go/src/html/template/template.go:476
	_go_fuzz_dep_.CoverTab[31559]++
							return parseFiles(t, readFileOS, filenames...)
//line /usr/local/go/src/html/template/template.go:477
	// _ = "end of CoverTab[31559]"
}

// IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
//line /usr/local/go/src/html/template/template.go:480
// and whether the value has a meaningful truth value. This is the definition of
//line /usr/local/go/src/html/template/template.go:480
// truth used by if and other such actions.
//line /usr/local/go/src/html/template/template.go:483
func IsTrue(val any) (truth, ok bool) {
//line /usr/local/go/src/html/template/template.go:483
	_go_fuzz_dep_.CoverTab[31566]++
							return template.IsTrue(val)
//line /usr/local/go/src/html/template/template.go:484
	// _ = "end of CoverTab[31566]"
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fs
//line /usr/local/go/src/html/template/template.go:487
// instead of the host operating system's file system.
//line /usr/local/go/src/html/template/template.go:487
// It accepts a list of glob patterns.
//line /usr/local/go/src/html/template/template.go:487
// (Note that most file names serve as glob patterns matching only themselves.)
//line /usr/local/go/src/html/template/template.go:491
func ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:491
	_go_fuzz_dep_.CoverTab[31567]++
							return parseFS(nil, fs, patterns)
//line /usr/local/go/src/html/template/template.go:492
	// _ = "end of CoverTab[31567]"
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fs
//line /usr/local/go/src/html/template/template.go:495
// instead of the host operating system's file system.
//line /usr/local/go/src/html/template/template.go:495
// It accepts a list of glob patterns.
//line /usr/local/go/src/html/template/template.go:495
// (Note that most file names serve as glob patterns matching only themselves.)
//line /usr/local/go/src/html/template/template.go:499
func (t *Template) ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:499
	_go_fuzz_dep_.CoverTab[31568]++
							return parseFS(t, fs, patterns)
//line /usr/local/go/src/html/template/template.go:500
	// _ = "end of CoverTab[31568]"
}

func parseFS(t *Template, fsys fs.FS, patterns []string) (*Template, error) {
//line /usr/local/go/src/html/template/template.go:503
	_go_fuzz_dep_.CoverTab[31569]++
							var filenames []string
							for _, pattern := range patterns {
//line /usr/local/go/src/html/template/template.go:505
		_go_fuzz_dep_.CoverTab[31571]++
								list, err := fs.Glob(fsys, pattern)
								if err != nil {
//line /usr/local/go/src/html/template/template.go:507
			_go_fuzz_dep_.CoverTab[31574]++
									return nil, err
//line /usr/local/go/src/html/template/template.go:508
			// _ = "end of CoverTab[31574]"
		} else {
//line /usr/local/go/src/html/template/template.go:509
			_go_fuzz_dep_.CoverTab[31575]++
//line /usr/local/go/src/html/template/template.go:509
			// _ = "end of CoverTab[31575]"
//line /usr/local/go/src/html/template/template.go:509
		}
//line /usr/local/go/src/html/template/template.go:509
		// _ = "end of CoverTab[31571]"
//line /usr/local/go/src/html/template/template.go:509
		_go_fuzz_dep_.CoverTab[31572]++
								if len(list) == 0 {
//line /usr/local/go/src/html/template/template.go:510
			_go_fuzz_dep_.CoverTab[31576]++
									return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
//line /usr/local/go/src/html/template/template.go:511
			// _ = "end of CoverTab[31576]"
		} else {
//line /usr/local/go/src/html/template/template.go:512
			_go_fuzz_dep_.CoverTab[31577]++
//line /usr/local/go/src/html/template/template.go:512
			// _ = "end of CoverTab[31577]"
//line /usr/local/go/src/html/template/template.go:512
		}
//line /usr/local/go/src/html/template/template.go:512
		// _ = "end of CoverTab[31572]"
//line /usr/local/go/src/html/template/template.go:512
		_go_fuzz_dep_.CoverTab[31573]++
								filenames = append(filenames, list...)
//line /usr/local/go/src/html/template/template.go:513
		// _ = "end of CoverTab[31573]"
	}
//line /usr/local/go/src/html/template/template.go:514
	// _ = "end of CoverTab[31569]"
//line /usr/local/go/src/html/template/template.go:514
	_go_fuzz_dep_.CoverTab[31570]++
							return parseFiles(t, readFileFS(fsys), filenames...)
//line /usr/local/go/src/html/template/template.go:515
	// _ = "end of CoverTab[31570]"
}

func readFileOS(file string) (name string, b []byte, err error) {
//line /usr/local/go/src/html/template/template.go:518
	_go_fuzz_dep_.CoverTab[31578]++
							name = filepath.Base(file)
							b, err = os.ReadFile(file)
							return
//line /usr/local/go/src/html/template/template.go:521
	// _ = "end of CoverTab[31578]"
}

func readFileFS(fsys fs.FS) func(string) (string, []byte, error) {
//line /usr/local/go/src/html/template/template.go:524
	_go_fuzz_dep_.CoverTab[31579]++
							return func(file string) (name string, b []byte, err error) {
//line /usr/local/go/src/html/template/template.go:525
		_go_fuzz_dep_.CoverTab[31580]++
								name = path.Base(file)
								b, err = fs.ReadFile(fsys, file)
								return
//line /usr/local/go/src/html/template/template.go:528
		// _ = "end of CoverTab[31580]"
	}
//line /usr/local/go/src/html/template/template.go:529
	// _ = "end of CoverTab[31579]"
}

//line /usr/local/go/src/html/template/template.go:530
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/template.go:530
var _ = _go_fuzz_dep_.CoverTab
