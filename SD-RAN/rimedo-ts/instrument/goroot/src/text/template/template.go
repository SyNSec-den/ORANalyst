// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/template/template.go:5
package template

//line /usr/local/go/src/text/template/template.go:5
import (
//line /usr/local/go/src/text/template/template.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/template.go:5
)
//line /usr/local/go/src/text/template/template.go:5
import (
//line /usr/local/go/src/text/template/template.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/template.go:5
)

import (
	"reflect"
	"sync"
	"text/template/parse"
)

// common holds the information shared by related templates.
type common struct {
	tmpl	map[string]*Template	// Map from name to defined templates.
	muTmpl	sync.RWMutex		// protects tmpl
	option	option
	// We use two maps, one for parsing and one for execution.
	// This separation makes the API cleaner since it doesn't
	// expose reflection to the client.
	muFuncs		sync.RWMutex	// protects parseFuncs and execFuncs
	parseFuncs	FuncMap
	execFuncs	map[string]reflect.Value
}

// Template is the representation of a parsed template. The *parse.Tree
//line /usr/local/go/src/text/template/template.go:26
// field is exported only for use by html/template and should be treated
//line /usr/local/go/src/text/template/template.go:26
// as unexported by all other clients.
//line /usr/local/go/src/text/template/template.go:29
type Template struct {
	name	string
	*parse.Tree
	*common
	leftDelim	string
	rightDelim	string
}

// New allocates a new, undefined template with the given name.
func New(name string) *Template {
//line /usr/local/go/src/text/template/template.go:38
	_go_fuzz_dep_.CoverTab[30580]++
							t := &Template{
		name: name,
	}
							t.init()
							return t
//line /usr/local/go/src/text/template/template.go:43
	// _ = "end of CoverTab[30580]"
}

// Name returns the name of the template.
func (t *Template) Name() string {
//line /usr/local/go/src/text/template/template.go:47
	_go_fuzz_dep_.CoverTab[30581]++
							return t.name
//line /usr/local/go/src/text/template/template.go:48
	// _ = "end of CoverTab[30581]"
}

// New allocates a new, undefined template associated with the given one and with the same
//line /usr/local/go/src/text/template/template.go:51
// delimiters. The association, which is transitive, allows one template to
//line /usr/local/go/src/text/template/template.go:51
// invoke another with a {{template}} action.
//line /usr/local/go/src/text/template/template.go:51
//
//line /usr/local/go/src/text/template/template.go:51
// Because associated templates share underlying data, template construction
//line /usr/local/go/src/text/template/template.go:51
// cannot be done safely in parallel. Once the templates are constructed, they
//line /usr/local/go/src/text/template/template.go:51
// can be executed in parallel.
//line /usr/local/go/src/text/template/template.go:58
func (t *Template) New(name string) *Template {
//line /usr/local/go/src/text/template/template.go:58
	_go_fuzz_dep_.CoverTab[30582]++
							t.init()
							nt := &Template{
		name:		name,
		common:		t.common,
		leftDelim:	t.leftDelim,
		rightDelim:	t.rightDelim,
	}
							return nt
//line /usr/local/go/src/text/template/template.go:66
	// _ = "end of CoverTab[30582]"
}

// init guarantees that t has a valid common structure.
func (t *Template) init() {
	if t.common == nil {
		c := new(common)
		c.tmpl = make(map[string]*Template)
		c.parseFuncs = make(FuncMap)
		c.execFuncs = make(map[string]reflect.Value)
		t.common = c
	}
}

// Clone returns a duplicate of the template, including all associated
//line /usr/local/go/src/text/template/template.go:80
// templates. The actual representation is not copied, but the name space of
//line /usr/local/go/src/text/template/template.go:80
// associated templates is, so further calls to Parse in the copy will add
//line /usr/local/go/src/text/template/template.go:80
// templates to the copy but not to the original. Clone can be used to prepare
//line /usr/local/go/src/text/template/template.go:80
// common templates and use them with variant definitions for other templates
//line /usr/local/go/src/text/template/template.go:80
// by adding the variants after the clone is made.
//line /usr/local/go/src/text/template/template.go:86
func (t *Template) Clone() (*Template, error) {
//line /usr/local/go/src/text/template/template.go:86
	_go_fuzz_dep_.CoverTab[30583]++
							nt := t.copy(nil)
							nt.init()
							if t.common == nil {
//line /usr/local/go/src/text/template/template.go:89
		_go_fuzz_dep_.CoverTab[30588]++
								return nt, nil
//line /usr/local/go/src/text/template/template.go:90
		// _ = "end of CoverTab[30588]"
	} else {
//line /usr/local/go/src/text/template/template.go:91
		_go_fuzz_dep_.CoverTab[30589]++
//line /usr/local/go/src/text/template/template.go:91
		// _ = "end of CoverTab[30589]"
//line /usr/local/go/src/text/template/template.go:91
	}
//line /usr/local/go/src/text/template/template.go:91
	// _ = "end of CoverTab[30583]"
//line /usr/local/go/src/text/template/template.go:91
	_go_fuzz_dep_.CoverTab[30584]++
							t.muTmpl.RLock()
							defer t.muTmpl.RUnlock()
							for k, v := range t.tmpl {
//line /usr/local/go/src/text/template/template.go:94
		_go_fuzz_dep_.CoverTab[30590]++
								if k == t.name {
//line /usr/local/go/src/text/template/template.go:95
			_go_fuzz_dep_.CoverTab[30592]++
									nt.tmpl[t.name] = nt
									continue
//line /usr/local/go/src/text/template/template.go:97
			// _ = "end of CoverTab[30592]"
		} else {
//line /usr/local/go/src/text/template/template.go:98
			_go_fuzz_dep_.CoverTab[30593]++
//line /usr/local/go/src/text/template/template.go:98
			// _ = "end of CoverTab[30593]"
//line /usr/local/go/src/text/template/template.go:98
		}
//line /usr/local/go/src/text/template/template.go:98
		// _ = "end of CoverTab[30590]"
//line /usr/local/go/src/text/template/template.go:98
		_go_fuzz_dep_.CoverTab[30591]++

								tmpl := v.copy(nt.common)
								nt.tmpl[k] = tmpl
//line /usr/local/go/src/text/template/template.go:101
		// _ = "end of CoverTab[30591]"
	}
//line /usr/local/go/src/text/template/template.go:102
	// _ = "end of CoverTab[30584]"
//line /usr/local/go/src/text/template/template.go:102
	_go_fuzz_dep_.CoverTab[30585]++
							t.muFuncs.RLock()
							defer t.muFuncs.RUnlock()
							for k, v := range t.parseFuncs {
//line /usr/local/go/src/text/template/template.go:105
		_go_fuzz_dep_.CoverTab[30594]++
								nt.parseFuncs[k] = v
//line /usr/local/go/src/text/template/template.go:106
		// _ = "end of CoverTab[30594]"
	}
//line /usr/local/go/src/text/template/template.go:107
	// _ = "end of CoverTab[30585]"
//line /usr/local/go/src/text/template/template.go:107
	_go_fuzz_dep_.CoverTab[30586]++
							for k, v := range t.execFuncs {
//line /usr/local/go/src/text/template/template.go:108
		_go_fuzz_dep_.CoverTab[30595]++
								nt.execFuncs[k] = v
//line /usr/local/go/src/text/template/template.go:109
		// _ = "end of CoverTab[30595]"
	}
//line /usr/local/go/src/text/template/template.go:110
	// _ = "end of CoverTab[30586]"
//line /usr/local/go/src/text/template/template.go:110
	_go_fuzz_dep_.CoverTab[30587]++
							return nt, nil
//line /usr/local/go/src/text/template/template.go:111
	// _ = "end of CoverTab[30587]"
}

// copy returns a shallow copy of t, with common set to the argument.
func (t *Template) copy(c *common) *Template {
//line /usr/local/go/src/text/template/template.go:115
	_go_fuzz_dep_.CoverTab[30596]++
							return &Template{
		name:		t.name,
		Tree:		t.Tree,
		common:		c,
		leftDelim:	t.leftDelim,
		rightDelim:	t.rightDelim,
	}
//line /usr/local/go/src/text/template/template.go:122
	// _ = "end of CoverTab[30596]"
}

// AddParseTree associates the argument parse tree with the template t, giving
//line /usr/local/go/src/text/template/template.go:125
// it the specified name. If the template has not been defined, this tree becomes
//line /usr/local/go/src/text/template/template.go:125
// its definition. If it has been defined and already has that name, the existing
//line /usr/local/go/src/text/template/template.go:125
// definition is replaced; otherwise a new template is created, defined, and returned.
//line /usr/local/go/src/text/template/template.go:129
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
//line /usr/local/go/src/text/template/template.go:129
	_go_fuzz_dep_.CoverTab[30597]++
							t.init()
							t.muTmpl.Lock()
							defer t.muTmpl.Unlock()
							nt := t
							if name != t.name {
//line /usr/local/go/src/text/template/template.go:134
		_go_fuzz_dep_.CoverTab[30600]++
								nt = t.New(name)
//line /usr/local/go/src/text/template/template.go:135
		// _ = "end of CoverTab[30600]"
	} else {
//line /usr/local/go/src/text/template/template.go:136
		_go_fuzz_dep_.CoverTab[30601]++
//line /usr/local/go/src/text/template/template.go:136
		// _ = "end of CoverTab[30601]"
//line /usr/local/go/src/text/template/template.go:136
	}
//line /usr/local/go/src/text/template/template.go:136
	// _ = "end of CoverTab[30597]"
//line /usr/local/go/src/text/template/template.go:136
	_go_fuzz_dep_.CoverTab[30598]++

							if t.associate(nt, tree) || func() bool {
//line /usr/local/go/src/text/template/template.go:138
		_go_fuzz_dep_.CoverTab[30602]++
//line /usr/local/go/src/text/template/template.go:138
		return nt.Tree == nil
//line /usr/local/go/src/text/template/template.go:138
		// _ = "end of CoverTab[30602]"
//line /usr/local/go/src/text/template/template.go:138
	}() {
//line /usr/local/go/src/text/template/template.go:138
		_go_fuzz_dep_.CoverTab[30603]++
								nt.Tree = tree
//line /usr/local/go/src/text/template/template.go:139
		// _ = "end of CoverTab[30603]"
	} else {
//line /usr/local/go/src/text/template/template.go:140
		_go_fuzz_dep_.CoverTab[30604]++
//line /usr/local/go/src/text/template/template.go:140
		// _ = "end of CoverTab[30604]"
//line /usr/local/go/src/text/template/template.go:140
	}
//line /usr/local/go/src/text/template/template.go:140
	// _ = "end of CoverTab[30598]"
//line /usr/local/go/src/text/template/template.go:140
	_go_fuzz_dep_.CoverTab[30599]++
							return nt, nil
//line /usr/local/go/src/text/template/template.go:141
	// _ = "end of CoverTab[30599]"
}

// Templates returns a slice of defined templates associated with t.
func (t *Template) Templates() []*Template {
//line /usr/local/go/src/text/template/template.go:145
	_go_fuzz_dep_.CoverTab[30605]++
							if t.common == nil {
//line /usr/local/go/src/text/template/template.go:146
		_go_fuzz_dep_.CoverTab[30608]++
								return nil
//line /usr/local/go/src/text/template/template.go:147
		// _ = "end of CoverTab[30608]"
	} else {
//line /usr/local/go/src/text/template/template.go:148
		_go_fuzz_dep_.CoverTab[30609]++
//line /usr/local/go/src/text/template/template.go:148
		// _ = "end of CoverTab[30609]"
//line /usr/local/go/src/text/template/template.go:148
	}
//line /usr/local/go/src/text/template/template.go:148
	// _ = "end of CoverTab[30605]"
//line /usr/local/go/src/text/template/template.go:148
	_go_fuzz_dep_.CoverTab[30606]++

							t.muTmpl.RLock()
							defer t.muTmpl.RUnlock()
							m := make([]*Template, 0, len(t.tmpl))
							for _, v := range t.tmpl {
//line /usr/local/go/src/text/template/template.go:153
		_go_fuzz_dep_.CoverTab[30610]++
								m = append(m, v)
//line /usr/local/go/src/text/template/template.go:154
		// _ = "end of CoverTab[30610]"
	}
//line /usr/local/go/src/text/template/template.go:155
	// _ = "end of CoverTab[30606]"
//line /usr/local/go/src/text/template/template.go:155
	_go_fuzz_dep_.CoverTab[30607]++
							return m
//line /usr/local/go/src/text/template/template.go:156
	// _ = "end of CoverTab[30607]"
}

// Delims sets the action delimiters to the specified strings, to be used in
//line /usr/local/go/src/text/template/template.go:159
// subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
//line /usr/local/go/src/text/template/template.go:159
// definitions will inherit the settings. An empty delimiter stands for the
//line /usr/local/go/src/text/template/template.go:159
// corresponding default: {{ or }}.
//line /usr/local/go/src/text/template/template.go:159
// The return value is the template, so calls can be chained.
//line /usr/local/go/src/text/template/template.go:164
func (t *Template) Delims(left, right string) *Template {
//line /usr/local/go/src/text/template/template.go:164
	_go_fuzz_dep_.CoverTab[30611]++
							t.init()
							t.leftDelim = left
							t.rightDelim = right
							return t
//line /usr/local/go/src/text/template/template.go:168
	// _ = "end of CoverTab[30611]"
}

// Funcs adds the elements of the argument map to the template's function map.
//line /usr/local/go/src/text/template/template.go:171
// It must be called before the template is parsed.
//line /usr/local/go/src/text/template/template.go:171
// It panics if a value in the map is not a function with appropriate return
//line /usr/local/go/src/text/template/template.go:171
// type or if the name cannot be used syntactically as a function in a template.
//line /usr/local/go/src/text/template/template.go:171
// It is legal to overwrite elements of the map. The return value is the template,
//line /usr/local/go/src/text/template/template.go:171
// so calls can be chained.
//line /usr/local/go/src/text/template/template.go:177
func (t *Template) Funcs(funcMap FuncMap) *Template {
//line /usr/local/go/src/text/template/template.go:177
	_go_fuzz_dep_.CoverTab[30612]++
							t.init()
							t.muFuncs.Lock()
							defer t.muFuncs.Unlock()
							addValueFuncs(t.execFuncs, funcMap)
							addFuncs(t.parseFuncs, funcMap)
							return t
//line /usr/local/go/src/text/template/template.go:183
	// _ = "end of CoverTab[30612]"
}

// Lookup returns the template with the given name that is associated with t.
//line /usr/local/go/src/text/template/template.go:186
// It returns nil if there is no such template or the template has no definition.
//line /usr/local/go/src/text/template/template.go:188
func (t *Template) Lookup(name string) *Template {
//line /usr/local/go/src/text/template/template.go:188
	_go_fuzz_dep_.CoverTab[30613]++
							if t.common == nil {
//line /usr/local/go/src/text/template/template.go:189
		_go_fuzz_dep_.CoverTab[30615]++
								return nil
//line /usr/local/go/src/text/template/template.go:190
		// _ = "end of CoverTab[30615]"
	} else {
//line /usr/local/go/src/text/template/template.go:191
		_go_fuzz_dep_.CoverTab[30616]++
//line /usr/local/go/src/text/template/template.go:191
		// _ = "end of CoverTab[30616]"
//line /usr/local/go/src/text/template/template.go:191
	}
//line /usr/local/go/src/text/template/template.go:191
	// _ = "end of CoverTab[30613]"
//line /usr/local/go/src/text/template/template.go:191
	_go_fuzz_dep_.CoverTab[30614]++
							t.muTmpl.RLock()
							defer t.muTmpl.RUnlock()
							return t.tmpl[name]
//line /usr/local/go/src/text/template/template.go:194
	// _ = "end of CoverTab[30614]"
}

// Parse parses text as a template body for t.
//line /usr/local/go/src/text/template/template.go:197
// Named template definitions ({{define ...}} or {{block ...}} statements) in text
//line /usr/local/go/src/text/template/template.go:197
// define additional templates associated with t and are removed from the
//line /usr/local/go/src/text/template/template.go:197
// definition of t itself.
//line /usr/local/go/src/text/template/template.go:197
//
//line /usr/local/go/src/text/template/template.go:197
// Templates can be redefined in successive calls to Parse.
//line /usr/local/go/src/text/template/template.go:197
// A template definition with a body containing only white space and comments
//line /usr/local/go/src/text/template/template.go:197
// is considered empty and will not replace an existing template's body.
//line /usr/local/go/src/text/template/template.go:197
// This allows using Parse to add new named template definitions without
//line /usr/local/go/src/text/template/template.go:197
// overwriting the main template body.
//line /usr/local/go/src/text/template/template.go:207
func (t *Template) Parse(text string) (*Template, error) {
//line /usr/local/go/src/text/template/template.go:207
	_go_fuzz_dep_.CoverTab[30617]++
							t.init()
							t.muFuncs.RLock()
							trees, err := parse.Parse(t.name, text, t.leftDelim, t.rightDelim, t.parseFuncs, builtins())
							t.muFuncs.RUnlock()
							if err != nil {
//line /usr/local/go/src/text/template/template.go:212
		_go_fuzz_dep_.CoverTab[30620]++
								return nil, err
//line /usr/local/go/src/text/template/template.go:213
		// _ = "end of CoverTab[30620]"
	} else {
//line /usr/local/go/src/text/template/template.go:214
		_go_fuzz_dep_.CoverTab[30621]++
//line /usr/local/go/src/text/template/template.go:214
		// _ = "end of CoverTab[30621]"
//line /usr/local/go/src/text/template/template.go:214
	}
//line /usr/local/go/src/text/template/template.go:214
	// _ = "end of CoverTab[30617]"
//line /usr/local/go/src/text/template/template.go:214
	_go_fuzz_dep_.CoverTab[30618]++

							for name, tree := range trees {
//line /usr/local/go/src/text/template/template.go:216
		_go_fuzz_dep_.CoverTab[30622]++
								if _, err := t.AddParseTree(name, tree); err != nil {
//line /usr/local/go/src/text/template/template.go:217
			_go_fuzz_dep_.CoverTab[30623]++
									return nil, err
//line /usr/local/go/src/text/template/template.go:218
			// _ = "end of CoverTab[30623]"
		} else {
//line /usr/local/go/src/text/template/template.go:219
			_go_fuzz_dep_.CoverTab[30624]++
//line /usr/local/go/src/text/template/template.go:219
			// _ = "end of CoverTab[30624]"
//line /usr/local/go/src/text/template/template.go:219
		}
//line /usr/local/go/src/text/template/template.go:219
		// _ = "end of CoverTab[30622]"
	}
//line /usr/local/go/src/text/template/template.go:220
	// _ = "end of CoverTab[30618]"
//line /usr/local/go/src/text/template/template.go:220
	_go_fuzz_dep_.CoverTab[30619]++
							return t, nil
//line /usr/local/go/src/text/template/template.go:221
	// _ = "end of CoverTab[30619]"
}

// associate installs the new template into the group of templates associated
//line /usr/local/go/src/text/template/template.go:224
// with t. The two are already known to share the common structure.
//line /usr/local/go/src/text/template/template.go:224
// The boolean return value reports whether to store this tree as t.Tree.
//line /usr/local/go/src/text/template/template.go:227
func (t *Template) associate(new *Template, tree *parse.Tree) bool {
//line /usr/local/go/src/text/template/template.go:227
	_go_fuzz_dep_.CoverTab[30625]++
							if new.common != t.common {
//line /usr/local/go/src/text/template/template.go:228
		_go_fuzz_dep_.CoverTab[30628]++
								panic("internal error: associate not common")
//line /usr/local/go/src/text/template/template.go:229
		// _ = "end of CoverTab[30628]"
	} else {
//line /usr/local/go/src/text/template/template.go:230
		_go_fuzz_dep_.CoverTab[30629]++
//line /usr/local/go/src/text/template/template.go:230
		// _ = "end of CoverTab[30629]"
//line /usr/local/go/src/text/template/template.go:230
	}
//line /usr/local/go/src/text/template/template.go:230
	// _ = "end of CoverTab[30625]"
//line /usr/local/go/src/text/template/template.go:230
	_go_fuzz_dep_.CoverTab[30626]++
							if old := t.tmpl[new.name]; old != nil && func() bool {
//line /usr/local/go/src/text/template/template.go:231
		_go_fuzz_dep_.CoverTab[30630]++
//line /usr/local/go/src/text/template/template.go:231
		return parse.IsEmptyTree(tree.Root)
//line /usr/local/go/src/text/template/template.go:231
		// _ = "end of CoverTab[30630]"
//line /usr/local/go/src/text/template/template.go:231
	}() && func() bool {
//line /usr/local/go/src/text/template/template.go:231
		_go_fuzz_dep_.CoverTab[30631]++
//line /usr/local/go/src/text/template/template.go:231
		return old.Tree != nil
//line /usr/local/go/src/text/template/template.go:231
		// _ = "end of CoverTab[30631]"
//line /usr/local/go/src/text/template/template.go:231
	}() {
//line /usr/local/go/src/text/template/template.go:231
		_go_fuzz_dep_.CoverTab[30632]++

//line /usr/local/go/src/text/template/template.go:234
		return false
//line /usr/local/go/src/text/template/template.go:234
		// _ = "end of CoverTab[30632]"
	} else {
//line /usr/local/go/src/text/template/template.go:235
		_go_fuzz_dep_.CoverTab[30633]++
//line /usr/local/go/src/text/template/template.go:235
		// _ = "end of CoverTab[30633]"
//line /usr/local/go/src/text/template/template.go:235
	}
//line /usr/local/go/src/text/template/template.go:235
	// _ = "end of CoverTab[30626]"
//line /usr/local/go/src/text/template/template.go:235
	_go_fuzz_dep_.CoverTab[30627]++
							t.tmpl[new.name] = new
							return true
//line /usr/local/go/src/text/template/template.go:237
	// _ = "end of CoverTab[30627]"
}

//line /usr/local/go/src/text/template/template.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/template.go:238
var _ = _go_fuzz_dep_.CoverTab
