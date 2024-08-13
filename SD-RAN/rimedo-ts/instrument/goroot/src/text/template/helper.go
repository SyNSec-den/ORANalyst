// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helper functions to make constructing templates easier.

//line /usr/local/go/src/text/template/helper.go:7
package template

//line /usr/local/go/src/text/template/helper.go:7
import (
//line /usr/local/go/src/text/template/helper.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/helper.go:7
)
//line /usr/local/go/src/text/template/helper.go:7
import (
//line /usr/local/go/src/text/template/helper.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/helper.go:7
)

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

//line /usr/local/go/src/text/template/helper.go:19
// Must is a helper that wraps a call to a function returning (*Template, error)
//line /usr/local/go/src/text/template/helper.go:19
// and panics if the error is non-nil. It is intended for use in variable
//line /usr/local/go/src/text/template/helper.go:19
// initializations such as
//line /usr/local/go/src/text/template/helper.go:19
//
//line /usr/local/go/src/text/template/helper.go:19
//	var t = template.Must(template.New("name").Parse("text"))
//line /usr/local/go/src/text/template/helper.go:24
func Must(t *Template, err error) *Template {
//line /usr/local/go/src/text/template/helper.go:24
	_go_fuzz_dep_.CoverTab[30518]++
							if err != nil {
//line /usr/local/go/src/text/template/helper.go:25
		_go_fuzz_dep_.CoverTab[30520]++
								panic(err)
//line /usr/local/go/src/text/template/helper.go:26
		// _ = "end of CoverTab[30520]"
	} else {
//line /usr/local/go/src/text/template/helper.go:27
		_go_fuzz_dep_.CoverTab[30521]++
//line /usr/local/go/src/text/template/helper.go:27
		// _ = "end of CoverTab[30521]"
//line /usr/local/go/src/text/template/helper.go:27
	}
//line /usr/local/go/src/text/template/helper.go:27
	// _ = "end of CoverTab[30518]"
//line /usr/local/go/src/text/template/helper.go:27
	_go_fuzz_dep_.CoverTab[30519]++
							return t
//line /usr/local/go/src/text/template/helper.go:28
	// _ = "end of CoverTab[30519]"
}

// ParseFiles creates a new Template and parses the template definitions from
//line /usr/local/go/src/text/template/helper.go:31
// the named files. The returned template's name will have the base name and
//line /usr/local/go/src/text/template/helper.go:31
// parsed contents of the first file. There must be at least one file.
//line /usr/local/go/src/text/template/helper.go:31
// If an error occurs, parsing stops and the returned *Template is nil.
//line /usr/local/go/src/text/template/helper.go:31
//
//line /usr/local/go/src/text/template/helper.go:31
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/text/template/helper.go:31
// the last one mentioned will be the one that results.
//line /usr/local/go/src/text/template/helper.go:31
// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
//line /usr/local/go/src/text/template/helper.go:31
// named "foo", while "a/foo" is unavailable.
//line /usr/local/go/src/text/template/helper.go:40
func ParseFiles(filenames ...string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:40
	_go_fuzz_dep_.CoverTab[30522]++
							return parseFiles(nil, readFileOS, filenames...)
//line /usr/local/go/src/text/template/helper.go:41
	// _ = "end of CoverTab[30522]"
}

// ParseFiles parses the named files and associates the resulting templates with
//line /usr/local/go/src/text/template/helper.go:44
// t. If an error occurs, parsing stops and the returned template is nil;
//line /usr/local/go/src/text/template/helper.go:44
// otherwise it is t. There must be at least one file.
//line /usr/local/go/src/text/template/helper.go:44
// Since the templates created by ParseFiles are named by the base
//line /usr/local/go/src/text/template/helper.go:44
// names of the argument files, t should usually have the name of one
//line /usr/local/go/src/text/template/helper.go:44
// of the (base) names of the files. If it does not, depending on t's
//line /usr/local/go/src/text/template/helper.go:44
// contents before calling ParseFiles, t.Execute may fail. In that
//line /usr/local/go/src/text/template/helper.go:44
// case use t.ExecuteTemplate to execute a valid template.
//line /usr/local/go/src/text/template/helper.go:44
//
//line /usr/local/go/src/text/template/helper.go:44
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/text/template/helper.go:44
// the last one mentioned will be the one that results.
//line /usr/local/go/src/text/template/helper.go:55
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:55
	_go_fuzz_dep_.CoverTab[30523]++
							t.init()
							return parseFiles(t, readFileOS, filenames...)
//line /usr/local/go/src/text/template/helper.go:57
	// _ = "end of CoverTab[30523]"
}

// parseFiles is the helper for the method and function. If the argument
//line /usr/local/go/src/text/template/helper.go:60
// template is nil, it is created from the first file.
//line /usr/local/go/src/text/template/helper.go:62
func parseFiles(t *Template, readFile func(string) (string, []byte, error), filenames ...string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:62
	_go_fuzz_dep_.CoverTab[30524]++
							if len(filenames) == 0 {
//line /usr/local/go/src/text/template/helper.go:63
		_go_fuzz_dep_.CoverTab[30527]++

								return nil, fmt.Errorf("template: no files named in call to ParseFiles")
//line /usr/local/go/src/text/template/helper.go:65
		// _ = "end of CoverTab[30527]"
	} else {
//line /usr/local/go/src/text/template/helper.go:66
		_go_fuzz_dep_.CoverTab[30528]++
//line /usr/local/go/src/text/template/helper.go:66
		// _ = "end of CoverTab[30528]"
//line /usr/local/go/src/text/template/helper.go:66
	}
//line /usr/local/go/src/text/template/helper.go:66
	// _ = "end of CoverTab[30524]"
//line /usr/local/go/src/text/template/helper.go:66
	_go_fuzz_dep_.CoverTab[30525]++
							for _, filename := range filenames {
//line /usr/local/go/src/text/template/helper.go:67
		_go_fuzz_dep_.CoverTab[30529]++
								name, b, err := readFile(filename)
								if err != nil {
//line /usr/local/go/src/text/template/helper.go:69
			_go_fuzz_dep_.CoverTab[30533]++
									return nil, err
//line /usr/local/go/src/text/template/helper.go:70
			// _ = "end of CoverTab[30533]"
		} else {
//line /usr/local/go/src/text/template/helper.go:71
			_go_fuzz_dep_.CoverTab[30534]++
//line /usr/local/go/src/text/template/helper.go:71
			// _ = "end of CoverTab[30534]"
//line /usr/local/go/src/text/template/helper.go:71
		}
//line /usr/local/go/src/text/template/helper.go:71
		// _ = "end of CoverTab[30529]"
//line /usr/local/go/src/text/template/helper.go:71
		_go_fuzz_dep_.CoverTab[30530]++
								s := string(b)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *Template
		if t == nil {
//line /usr/local/go/src/text/template/helper.go:80
			_go_fuzz_dep_.CoverTab[30535]++
									t = New(name)
//line /usr/local/go/src/text/template/helper.go:81
			// _ = "end of CoverTab[30535]"
		} else {
//line /usr/local/go/src/text/template/helper.go:82
			_go_fuzz_dep_.CoverTab[30536]++
//line /usr/local/go/src/text/template/helper.go:82
			// _ = "end of CoverTab[30536]"
//line /usr/local/go/src/text/template/helper.go:82
		}
//line /usr/local/go/src/text/template/helper.go:82
		// _ = "end of CoverTab[30530]"
//line /usr/local/go/src/text/template/helper.go:82
		_go_fuzz_dep_.CoverTab[30531]++
								if name == t.Name() {
//line /usr/local/go/src/text/template/helper.go:83
			_go_fuzz_dep_.CoverTab[30537]++
									tmpl = t
//line /usr/local/go/src/text/template/helper.go:84
			// _ = "end of CoverTab[30537]"
		} else {
//line /usr/local/go/src/text/template/helper.go:85
			_go_fuzz_dep_.CoverTab[30538]++
									tmpl = t.New(name)
//line /usr/local/go/src/text/template/helper.go:86
			// _ = "end of CoverTab[30538]"
		}
//line /usr/local/go/src/text/template/helper.go:87
		// _ = "end of CoverTab[30531]"
//line /usr/local/go/src/text/template/helper.go:87
		_go_fuzz_dep_.CoverTab[30532]++
								_, err = tmpl.Parse(s)
								if err != nil {
//line /usr/local/go/src/text/template/helper.go:89
			_go_fuzz_dep_.CoverTab[30539]++
									return nil, err
//line /usr/local/go/src/text/template/helper.go:90
			// _ = "end of CoverTab[30539]"
		} else {
//line /usr/local/go/src/text/template/helper.go:91
			_go_fuzz_dep_.CoverTab[30540]++
//line /usr/local/go/src/text/template/helper.go:91
			// _ = "end of CoverTab[30540]"
//line /usr/local/go/src/text/template/helper.go:91
		}
//line /usr/local/go/src/text/template/helper.go:91
		// _ = "end of CoverTab[30532]"
	}
//line /usr/local/go/src/text/template/helper.go:92
	// _ = "end of CoverTab[30525]"
//line /usr/local/go/src/text/template/helper.go:92
	_go_fuzz_dep_.CoverTab[30526]++
							return t, nil
//line /usr/local/go/src/text/template/helper.go:93
	// _ = "end of CoverTab[30526]"
}

// ParseGlob creates a new Template and parses the template definitions from
//line /usr/local/go/src/text/template/helper.go:96
// the files identified by the pattern. The files are matched according to the
//line /usr/local/go/src/text/template/helper.go:96
// semantics of filepath.Match, and the pattern must match at least one file.
//line /usr/local/go/src/text/template/helper.go:96
// The returned template will have the (base) name and (parsed) contents of the
//line /usr/local/go/src/text/template/helper.go:96
// first file matched by the pattern. ParseGlob is equivalent to calling
//line /usr/local/go/src/text/template/helper.go:96
// ParseFiles with the list of files matched by the pattern.
//line /usr/local/go/src/text/template/helper.go:96
//
//line /usr/local/go/src/text/template/helper.go:96
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/text/template/helper.go:96
// the last one mentioned will be the one that results.
//line /usr/local/go/src/text/template/helper.go:105
func ParseGlob(pattern string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:105
	_go_fuzz_dep_.CoverTab[30541]++
							return parseGlob(nil, pattern)
//line /usr/local/go/src/text/template/helper.go:106
	// _ = "end of CoverTab[30541]"
}

// ParseGlob parses the template definitions in the files identified by the
//line /usr/local/go/src/text/template/helper.go:109
// pattern and associates the resulting templates with t. The files are matched
//line /usr/local/go/src/text/template/helper.go:109
// according to the semantics of filepath.Match, and the pattern must match at
//line /usr/local/go/src/text/template/helper.go:109
// least one file. ParseGlob is equivalent to calling t.ParseFiles with the
//line /usr/local/go/src/text/template/helper.go:109
// list of files matched by the pattern.
//line /usr/local/go/src/text/template/helper.go:109
//
//line /usr/local/go/src/text/template/helper.go:109
// When parsing multiple files with the same name in different directories,
//line /usr/local/go/src/text/template/helper.go:109
// the last one mentioned will be the one that results.
//line /usr/local/go/src/text/template/helper.go:117
func (t *Template) ParseGlob(pattern string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:117
	_go_fuzz_dep_.CoverTab[30542]++
							t.init()
							return parseGlob(t, pattern)
//line /usr/local/go/src/text/template/helper.go:119
	// _ = "end of CoverTab[30542]"
}

// parseGlob is the implementation of the function and method ParseGlob.
func parseGlob(t *Template, pattern string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:123
	_go_fuzz_dep_.CoverTab[30543]++
							filenames, err := filepath.Glob(pattern)
							if err != nil {
//line /usr/local/go/src/text/template/helper.go:125
		_go_fuzz_dep_.CoverTab[30546]++
								return nil, err
//line /usr/local/go/src/text/template/helper.go:126
		// _ = "end of CoverTab[30546]"
	} else {
//line /usr/local/go/src/text/template/helper.go:127
		_go_fuzz_dep_.CoverTab[30547]++
//line /usr/local/go/src/text/template/helper.go:127
		// _ = "end of CoverTab[30547]"
//line /usr/local/go/src/text/template/helper.go:127
	}
//line /usr/local/go/src/text/template/helper.go:127
	// _ = "end of CoverTab[30543]"
//line /usr/local/go/src/text/template/helper.go:127
	_go_fuzz_dep_.CoverTab[30544]++
							if len(filenames) == 0 {
//line /usr/local/go/src/text/template/helper.go:128
		_go_fuzz_dep_.CoverTab[30548]++
								return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
//line /usr/local/go/src/text/template/helper.go:129
		// _ = "end of CoverTab[30548]"
	} else {
//line /usr/local/go/src/text/template/helper.go:130
		_go_fuzz_dep_.CoverTab[30549]++
//line /usr/local/go/src/text/template/helper.go:130
		// _ = "end of CoverTab[30549]"
//line /usr/local/go/src/text/template/helper.go:130
	}
//line /usr/local/go/src/text/template/helper.go:130
	// _ = "end of CoverTab[30544]"
//line /usr/local/go/src/text/template/helper.go:130
	_go_fuzz_dep_.CoverTab[30545]++
							return parseFiles(t, readFileOS, filenames...)
//line /usr/local/go/src/text/template/helper.go:131
	// _ = "end of CoverTab[30545]"
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys
//line /usr/local/go/src/text/template/helper.go:134
// instead of the host operating system's file system.
//line /usr/local/go/src/text/template/helper.go:134
// It accepts a list of glob patterns.
//line /usr/local/go/src/text/template/helper.go:134
// (Note that most file names serve as glob patterns matching only themselves.)
//line /usr/local/go/src/text/template/helper.go:138
func ParseFS(fsys fs.FS, patterns ...string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:138
	_go_fuzz_dep_.CoverTab[30550]++
							return parseFS(nil, fsys, patterns)
//line /usr/local/go/src/text/template/helper.go:139
	// _ = "end of CoverTab[30550]"
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys
//line /usr/local/go/src/text/template/helper.go:142
// instead of the host operating system's file system.
//line /usr/local/go/src/text/template/helper.go:142
// It accepts a list of glob patterns.
//line /usr/local/go/src/text/template/helper.go:142
// (Note that most file names serve as glob patterns matching only themselves.)
//line /usr/local/go/src/text/template/helper.go:146
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:146
	_go_fuzz_dep_.CoverTab[30551]++
							t.init()
							return parseFS(t, fsys, patterns)
//line /usr/local/go/src/text/template/helper.go:148
	// _ = "end of CoverTab[30551]"
}

func parseFS(t *Template, fsys fs.FS, patterns []string) (*Template, error) {
//line /usr/local/go/src/text/template/helper.go:151
	_go_fuzz_dep_.CoverTab[30552]++
							var filenames []string
							for _, pattern := range patterns {
//line /usr/local/go/src/text/template/helper.go:153
		_go_fuzz_dep_.CoverTab[30554]++
								list, err := fs.Glob(fsys, pattern)
								if err != nil {
//line /usr/local/go/src/text/template/helper.go:155
			_go_fuzz_dep_.CoverTab[30557]++
									return nil, err
//line /usr/local/go/src/text/template/helper.go:156
			// _ = "end of CoverTab[30557]"
		} else {
//line /usr/local/go/src/text/template/helper.go:157
			_go_fuzz_dep_.CoverTab[30558]++
//line /usr/local/go/src/text/template/helper.go:157
			// _ = "end of CoverTab[30558]"
//line /usr/local/go/src/text/template/helper.go:157
		}
//line /usr/local/go/src/text/template/helper.go:157
		// _ = "end of CoverTab[30554]"
//line /usr/local/go/src/text/template/helper.go:157
		_go_fuzz_dep_.CoverTab[30555]++
								if len(list) == 0 {
//line /usr/local/go/src/text/template/helper.go:158
			_go_fuzz_dep_.CoverTab[30559]++
									return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
//line /usr/local/go/src/text/template/helper.go:159
			// _ = "end of CoverTab[30559]"
		} else {
//line /usr/local/go/src/text/template/helper.go:160
			_go_fuzz_dep_.CoverTab[30560]++
//line /usr/local/go/src/text/template/helper.go:160
			// _ = "end of CoverTab[30560]"
//line /usr/local/go/src/text/template/helper.go:160
		}
//line /usr/local/go/src/text/template/helper.go:160
		// _ = "end of CoverTab[30555]"
//line /usr/local/go/src/text/template/helper.go:160
		_go_fuzz_dep_.CoverTab[30556]++
								filenames = append(filenames, list...)
//line /usr/local/go/src/text/template/helper.go:161
		// _ = "end of CoverTab[30556]"
	}
//line /usr/local/go/src/text/template/helper.go:162
	// _ = "end of CoverTab[30552]"
//line /usr/local/go/src/text/template/helper.go:162
	_go_fuzz_dep_.CoverTab[30553]++
							return parseFiles(t, readFileFS(fsys), filenames...)
//line /usr/local/go/src/text/template/helper.go:163
	// _ = "end of CoverTab[30553]"
}

func readFileOS(file string) (name string, b []byte, err error) {
//line /usr/local/go/src/text/template/helper.go:166
	_go_fuzz_dep_.CoverTab[30561]++
							name = filepath.Base(file)
							b, err = os.ReadFile(file)
							return
//line /usr/local/go/src/text/template/helper.go:169
	// _ = "end of CoverTab[30561]"
}

func readFileFS(fsys fs.FS) func(string) (string, []byte, error) {
//line /usr/local/go/src/text/template/helper.go:172
	_go_fuzz_dep_.CoverTab[30562]++
							return func(file string) (name string, b []byte, err error) {
//line /usr/local/go/src/text/template/helper.go:173
		_go_fuzz_dep_.CoverTab[30563]++
								name = path.Base(file)
								b, err = fs.ReadFile(fsys, file)
								return
//line /usr/local/go/src/text/template/helper.go:176
		// _ = "end of CoverTab[30563]"
	}
//line /usr/local/go/src/text/template/helper.go:177
	// _ = "end of CoverTab[30562]"
}

//line /usr/local/go/src/text/template/helper.go:178
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/helper.go:178
var _ = _go_fuzz_dep_.CoverTab
