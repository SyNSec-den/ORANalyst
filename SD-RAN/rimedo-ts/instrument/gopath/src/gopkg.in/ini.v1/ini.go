// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:15
// Package ini provides INI file read and write functionality in Go.
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:16
)

import (
	"os"
	"regexp"
	"runtime"
	"strings"
)

const (
	// DefaultSection is the name of default section. You can use this constant or the string literal.
	// In most of cases, an empty string is all you need to access the section.
	DefaultSection	= "DEFAULT"

	// Maximum allowed depth when recursively substituing variable names.
	depthValues	= 99
)

var (
	// LineBreak is the delimiter to determine or compose a new line.
	// This variable will be changed to "\r\n" automatically on Windows at package init time.
	LineBreak	= "\n"

	// Variable regexp pattern: %(variable)s
	varPattern	= regexp.MustCompile(`%\(([^)]+)\)s`)

	// DefaultHeader explicitly writes default section header.
	DefaultHeader	= false

	// PrettySection indicates whether to put a line between sections.
	PrettySection	= true
	// PrettyFormat indicates whether to align "=" sign with spaces to produce pretty output
	// or reduce all possible spaces for compact format.
	PrettyFormat	= true
	// PrettyEqual places spaces around "=" sign even when PrettyFormat is false.
	PrettyEqual	= false
	// DefaultFormatLeft places custom spaces on the left when PrettyFormat and PrettyEqual are both disabled.
	DefaultFormatLeft	= ""
	// DefaultFormatRight places custom spaces on the right when PrettyFormat and PrettyEqual are both disabled.
	DefaultFormatRight	= ""
)

var inTest = len(os.Args) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:58
	_go_fuzz_dep_.CoverTab[128501]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:58
	return strings.HasSuffix(strings.TrimSuffix(os.Args[0], ".exe"), ".test")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:58
	// _ = "end of CoverTab[128501]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:58
}()

func init() {
	if runtime.GOOS == "windows" && !inTest {
		LineBreak = "\r\n"
	}
}

// LoadOptions contains all customized options used for load data source(s).
type LoadOptions struct {
	// Loose indicates whether the parser should ignore nonexistent files or return error.
	Loose	bool
	// Insensitive indicates whether the parser forces all section and key names to lowercase.
	Insensitive	bool
	// InsensitiveSections indicates whether the parser forces all section to lowercase.
	InsensitiveSections	bool
	// InsensitiveKeys indicates whether the parser forces all key names to lowercase.
	InsensitiveKeys	bool
	// IgnoreContinuation indicates whether to ignore continuation lines while parsing.
	IgnoreContinuation	bool
	// IgnoreInlineComment indicates whether to ignore comments at the end of value and treat it as part of value.
	IgnoreInlineComment	bool
	// SkipUnrecognizableLines indicates whether to skip unrecognizable lines that do not conform to key/value pairs.
	SkipUnrecognizableLines	bool
	// ShortCircuit indicates whether to ignore other configuration sources after loaded the first available configuration source.
	ShortCircuit	bool
	// AllowBooleanKeys indicates whether to allow boolean type keys or treat as value is missing.
	// This type of keys are mostly used in my.cnf.
	AllowBooleanKeys	bool
	// AllowShadows indicates whether to keep track of keys with same name under same section.
	AllowShadows	bool
	// AllowNestedValues indicates whether to allow AWS-like nested values.
	// Docs: http://docs.aws.amazon.com/cli/latest/topic/config-vars.html#nested-values
	AllowNestedValues	bool
	// AllowPythonMultilineValues indicates whether to allow Python-like multi-line values.
	// Docs: https://docs.python.org/3/library/configparser.html#supported-ini-file-structure
	// Relevant quote:  Values can also span multiple lines, as long as they are indented deeper
	// than the first line of the value.
	AllowPythonMultilineValues	bool
	// SpaceBeforeInlineComment indicates whether to allow comment symbols (\# and \;) inside value.
	// Docs: https://docs.python.org/2/library/configparser.html
	// Quote: Comments may appear on their own in an otherwise empty line, or may be entered in lines holding values or section names.
	// In the latter case, they need to be preceded by a whitespace character to be recognized as a comment.
	SpaceBeforeInlineComment	bool
	// UnescapeValueDoubleQuotes indicates whether to unescape double quotes inside value to regular format
	// when value is surrounded by double quotes, e.g. key="a \"value\"" => key=a "value"
	UnescapeValueDoubleQuotes	bool
	// UnescapeValueCommentSymbols indicates to unescape comment symbols (\# and \;) inside value to regular format
	// when value is NOT surrounded by any quotes.
	// Note: UNSTABLE, behavior might change to only unescape inside double quotes but may noy necessary at all.
	UnescapeValueCommentSymbols	bool
	// UnparseableSections stores a list of blocks that are allowed with raw content which do not otherwise
	// conform to key/value pairs. Specify the names of those blocks here.
	UnparseableSections	[]string
	// KeyValueDelimiters is the sequence of delimiters that are used to separate key and value. By default, it is "=:".
	KeyValueDelimiters	string
	// KeyValueDelimiterOnWrite is the delimiter that are used to separate key and value output. By default, it is "=".
	KeyValueDelimiterOnWrite	string
	// ChildSectionDelimiter is the delimiter that is used to separate child sections. By default, it is ".".
	ChildSectionDelimiter	string
	// PreserveSurroundedQuote indicates whether to preserve surrounded quote (single and double quotes).
	PreserveSurroundedQuote	bool
	// DebugFunc is called to collect debug information (currently only useful to debug parsing Python-style multiline values).
	DebugFunc	DebugFunc
	// ReaderBufferSize is the buffer size of the reader in bytes.
	ReaderBufferSize	int
	// AllowNonUniqueSections indicates whether to allow sections with the same name multiple times.
	AllowNonUniqueSections	bool
	// AllowDuplicateShadowValues indicates whether values for shadowed keys should be deduplicated.
	AllowDuplicateShadowValues	bool
}

// DebugFunc is the type of function called to log parse events.
type DebugFunc func(message string)

// LoadSources allows caller to apply customized options for loading from data source(s).
func LoadSources(opts LoadOptions, source interface{}, others ...interface{}) (_ *File, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:134
	_go_fuzz_dep_.CoverTab[128502]++
									sources := make([]dataSource, len(others)+1)
									sources[0], err = parseDataSource(source)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:137
		_go_fuzz_dep_.CoverTab[128506]++
										return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:138
		// _ = "end of CoverTab[128506]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:139
		_go_fuzz_dep_.CoverTab[128507]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:139
		// _ = "end of CoverTab[128507]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:139
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:139
	// _ = "end of CoverTab[128502]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:139
	_go_fuzz_dep_.CoverTab[128503]++
									for i := range others {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:140
		_go_fuzz_dep_.CoverTab[128508]++
										sources[i+1], err = parseDataSource(others[i])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:142
			_go_fuzz_dep_.CoverTab[128509]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:143
			// _ = "end of CoverTab[128509]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:144
			_go_fuzz_dep_.CoverTab[128510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:144
			// _ = "end of CoverTab[128510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:144
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:144
		// _ = "end of CoverTab[128508]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:145
	// _ = "end of CoverTab[128503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:145
	_go_fuzz_dep_.CoverTab[128504]++
									f := newFile(sources, opts)
									if err = f.Reload(); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:147
		_go_fuzz_dep_.CoverTab[128511]++
										return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:148
		// _ = "end of CoverTab[128511]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:149
		_go_fuzz_dep_.CoverTab[128512]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:149
		// _ = "end of CoverTab[128512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:149
	// _ = "end of CoverTab[128504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:149
	_go_fuzz_dep_.CoverTab[128505]++
									return f, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:150
	// _ = "end of CoverTab[128505]"
}

// Load loads and parses from INI data sources.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:153
// Arguments can be mixed of file name with string type, or raw data in []byte.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:153
// It will return error if list contains nonexistent files.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:156
func Load(source interface{}, others ...interface{}) (*File, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:156
	_go_fuzz_dep_.CoverTab[128513]++
									return LoadSources(LoadOptions{}, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:157
	// _ = "end of CoverTab[128513]"
}

// LooseLoad has exactly same functionality as Load function
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:160
// except it ignores nonexistent files instead of returning error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:162
func LooseLoad(source interface{}, others ...interface{}) (*File, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:162
	_go_fuzz_dep_.CoverTab[128514]++
									return LoadSources(LoadOptions{Loose: true}, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:163
	// _ = "end of CoverTab[128514]"
}

// InsensitiveLoad has exactly same functionality as Load function
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:166
// except it forces all section and key names to be lowercased.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:168
func InsensitiveLoad(source interface{}, others ...interface{}) (*File, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:168
	_go_fuzz_dep_.CoverTab[128515]++
									return LoadSources(LoadOptions{Insensitive: true}, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:169
	// _ = "end of CoverTab[128515]"
}

// ShadowLoad has exactly same functionality as Load function
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:172
// except it allows have shadow keys.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:174
func ShadowLoad(source interface{}, others ...interface{}) (*File, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:174
	_go_fuzz_dep_.CoverTab[128516]++
									return LoadSources(LoadOptions{AllowShadows: true}, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:175
	// _ = "end of CoverTab[128516]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/ini.go:176
var _ = _go_fuzz_dep_.CoverTab
