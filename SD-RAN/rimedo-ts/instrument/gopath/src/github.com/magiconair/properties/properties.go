// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:5
)

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:10
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const maxExpansionDepth = 64

// ErrorHandlerFunc defines the type of function which handles failures
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:26
// of the MustXXX() functions. An error handler function must exit
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:26
// the application after handling the error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:29
type ErrorHandlerFunc func(error)

// ErrorHandler is the function which handles failures of the MustXXX()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:31
// functions. The default is LogFatalHandler.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:33
var ErrorHandler ErrorHandlerFunc = LogFatalHandler

// LogHandlerFunc defines the function prototype for logging errors.
type LogHandlerFunc func(fmt string, args ...interface{})

// LogPrintf defines a log handler which uses log.Printf.
var LogPrintf LogHandlerFunc = log.Printf

// LogFatalHandler handles the error by logging a fatal error and exiting.
func LogFatalHandler(err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:42
	_go_fuzz_dep_.CoverTab[115854]++
												log.Fatal(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:43
	// _ = "end of CoverTab[115854]"
}

// PanicHandler handles the error by panicking.
func PanicHandler(err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:47
	_go_fuzz_dep_.CoverTab[115855]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:48
	// _ = "end of CoverTab[115855]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:53
// A Properties contains the key/value pairs from the properties input.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:53
// All values are stored in unexpanded form and are expanded at runtime
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:55
type Properties struct {
	// Pre-/Postfix for property expansion.
	Prefix	string
	Postfix	string

	// DisableExpansion controls the expansion of properties on Get()
	// and the check for circular references on Set(). When set to
	// true Properties behaves like a simple key/value store and does
	// not check for circular references on Get() or on Set().
	DisableExpansion	bool

	// Stores the key/value pairs
	m	map[string]string

	// Stores the comments per key.
	c	map[string][]string

	// Stores the keys in order of appearance.
	k	[]string

	// WriteSeparator specifies the separator of key and value while writing the properties.
	WriteSeparator	string
}

// NewProperties creates a new Properties struct with the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:79
// configuration for "${key}" expressions.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:81
func NewProperties() *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:81
	_go_fuzz_dep_.CoverTab[115856]++
												return &Properties{
		Prefix:		"${",
		Postfix:	"}",
		m:		map[string]string{},
		c:		map[string][]string{},
		k:		[]string{},
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:88
	// _ = "end of CoverTab[115856]"
}

// Load reads a buffer into the given Properties struct.
func (p *Properties) Load(buf []byte, enc Encoding) error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:92
	_go_fuzz_dep_.CoverTab[115857]++
												l := &Loader{Encoding: enc, DisableExpansion: p.DisableExpansion}
												newProperties, err := l.LoadBytes(buf)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:95
		_go_fuzz_dep_.CoverTab[115859]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:96
		// _ = "end of CoverTab[115859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:97
		_go_fuzz_dep_.CoverTab[115860]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:97
		// _ = "end of CoverTab[115860]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:97
	// _ = "end of CoverTab[115857]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:97
	_go_fuzz_dep_.CoverTab[115858]++
												p.Merge(newProperties)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:99
	// _ = "end of CoverTab[115858]"
}

// Get returns the expanded value for the given key if exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:102
// Otherwise, ok is false.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:104
func (p *Properties) Get(key string) (value string, ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:104
	_go_fuzz_dep_.CoverTab[115861]++
												v, ok := p.m[key]
												if p.DisableExpansion {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:106
		_go_fuzz_dep_.CoverTab[115865]++
													return v, ok
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:107
		// _ = "end of CoverTab[115865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:108
		_go_fuzz_dep_.CoverTab[115866]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:108
		// _ = "end of CoverTab[115866]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:108
	// _ = "end of CoverTab[115861]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:108
	_go_fuzz_dep_.CoverTab[115862]++
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:109
		_go_fuzz_dep_.CoverTab[115867]++
													return "", false
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:110
		// _ = "end of CoverTab[115867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:111
		_go_fuzz_dep_.CoverTab[115868]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:111
		// _ = "end of CoverTab[115868]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:111
	// _ = "end of CoverTab[115862]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:111
	_go_fuzz_dep_.CoverTab[115863]++

												expanded, err := p.expand(key, v)

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:118
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:118
		_go_fuzz_dep_.CoverTab[115869]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:119
		// _ = "end of CoverTab[115869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:120
		_go_fuzz_dep_.CoverTab[115870]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:120
		// _ = "end of CoverTab[115870]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:120
	// _ = "end of CoverTab[115863]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:120
	_go_fuzz_dep_.CoverTab[115864]++

												return expanded, true
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:122
	// _ = "end of CoverTab[115864]"
}

// MustGet returns the expanded value for the given key if exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:125
// Otherwise, it panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:127
func (p *Properties) MustGet(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:127
	_go_fuzz_dep_.CoverTab[115871]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:128
		_go_fuzz_dep_.CoverTab[115873]++
													return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:129
		// _ = "end of CoverTab[115873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:130
		_go_fuzz_dep_.CoverTab[115874]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:130
		// _ = "end of CoverTab[115874]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:130
	// _ = "end of CoverTab[115871]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:130
	_go_fuzz_dep_.CoverTab[115872]++
												ErrorHandler(invalidKeyError(key))
												panic("ErrorHandler should exit")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:132
	// _ = "end of CoverTab[115872]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:137
// ClearComments removes the comments for all keys.
func (p *Properties) ClearComments() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:138
	_go_fuzz_dep_.CoverTab[115875]++
												p.c = map[string][]string{}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:139
	// _ = "end of CoverTab[115875]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:144
// GetComment returns the last comment before the given key or an empty string.
func (p *Properties) GetComment(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:145
	_go_fuzz_dep_.CoverTab[115876]++
												comments, ok := p.c[key]
												if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:147
		_go_fuzz_dep_.CoverTab[115878]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:147
		return len(comments) == 0
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:147
		// _ = "end of CoverTab[115878]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:147
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:147
		_go_fuzz_dep_.CoverTab[115879]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:148
		// _ = "end of CoverTab[115879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:149
		_go_fuzz_dep_.CoverTab[115880]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:149
		// _ = "end of CoverTab[115880]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:149
	// _ = "end of CoverTab[115876]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:149
	_go_fuzz_dep_.CoverTab[115877]++
												return comments[len(comments)-1]
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:150
	// _ = "end of CoverTab[115877]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:155
// GetComments returns all comments that appeared before the given key or nil.
func (p *Properties) GetComments(key string) []string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:156
	_go_fuzz_dep_.CoverTab[115881]++
												if comments, ok := p.c[key]; ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:157
		_go_fuzz_dep_.CoverTab[115883]++
													return comments
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:158
		// _ = "end of CoverTab[115883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:159
		_go_fuzz_dep_.CoverTab[115884]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:159
		// _ = "end of CoverTab[115884]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:159
	// _ = "end of CoverTab[115881]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:159
	_go_fuzz_dep_.CoverTab[115882]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:160
	// _ = "end of CoverTab[115882]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:165
// SetComment sets the comment for the key.
func (p *Properties) SetComment(key, comment string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:166
	_go_fuzz_dep_.CoverTab[115885]++
												p.c[key] = []string{comment}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:167
	// _ = "end of CoverTab[115885]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:172
// SetComments sets the comments for the key. If the comments are nil then
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:172
// all comments for this key are deleted.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:174
func (p *Properties) SetComments(key string, comments []string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:174
	_go_fuzz_dep_.CoverTab[115886]++
												if comments == nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:175
		_go_fuzz_dep_.CoverTab[115888]++
													delete(p.c, key)
													return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:177
		// _ = "end of CoverTab[115888]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:178
		_go_fuzz_dep_.CoverTab[115889]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:178
		// _ = "end of CoverTab[115889]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:178
	// _ = "end of CoverTab[115886]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:178
	_go_fuzz_dep_.CoverTab[115887]++
												p.c[key] = comments
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:179
	// _ = "end of CoverTab[115887]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:184
// GetBool checks if the expanded value is one of '1', 'yes',
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:184
// 'true' or 'on' if the key exists. The comparison is case-insensitive.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:184
// If the key does not exist the default value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:187
func (p *Properties) GetBool(key string, def bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:187
	_go_fuzz_dep_.CoverTab[115890]++
												v, err := p.getBool(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:189
		_go_fuzz_dep_.CoverTab[115892]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:190
		// _ = "end of CoverTab[115892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:191
		_go_fuzz_dep_.CoverTab[115893]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:191
		// _ = "end of CoverTab[115893]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:191
	// _ = "end of CoverTab[115890]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:191
	_go_fuzz_dep_.CoverTab[115891]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:192
	// _ = "end of CoverTab[115891]"
}

// MustGetBool checks if the expanded value is one of '1', 'yes',
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:195
// 'true' or 'on' if the key exists. The comparison is case-insensitive.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:195
// If the key does not exist the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:198
func (p *Properties) MustGetBool(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:198
	_go_fuzz_dep_.CoverTab[115894]++
												v, err := p.getBool(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:200
		_go_fuzz_dep_.CoverTab[115896]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:201
		// _ = "end of CoverTab[115896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:202
		_go_fuzz_dep_.CoverTab[115897]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:202
		// _ = "end of CoverTab[115897]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:202
	// _ = "end of CoverTab[115894]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:202
	_go_fuzz_dep_.CoverTab[115895]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:203
	// _ = "end of CoverTab[115895]"
}

func (p *Properties) getBool(key string) (value bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:206
	_go_fuzz_dep_.CoverTab[115898]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:207
		_go_fuzz_dep_.CoverTab[115900]++
													return boolVal(v), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:208
		// _ = "end of CoverTab[115900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:209
		_go_fuzz_dep_.CoverTab[115901]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:209
		// _ = "end of CoverTab[115901]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:209
	// _ = "end of CoverTab[115898]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:209
	_go_fuzz_dep_.CoverTab[115899]++
												return false, invalidKeyError(key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:210
	// _ = "end of CoverTab[115899]"
}

func boolVal(v string) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:213
	_go_fuzz_dep_.CoverTab[115902]++
												v = strings.ToLower(v)
												return v == "1" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		_go_fuzz_dep_.CoverTab[115903]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		return v == "true"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		// _ = "end of CoverTab[115903]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		_go_fuzz_dep_.CoverTab[115904]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		return v == "yes"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		// _ = "end of CoverTab[115904]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		_go_fuzz_dep_.CoverTab[115905]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		return v == "on"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
		// _ = "end of CoverTab[115905]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:215
	// _ = "end of CoverTab[115902]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:220
// GetDuration parses the expanded value as an time.Duration (in ns) if the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:220
// key exists. If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:220
// value is returned. In almost all cases you want to use GetParsedDuration().
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:223
func (p *Properties) GetDuration(key string, def time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:223
	_go_fuzz_dep_.CoverTab[115906]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:225
		_go_fuzz_dep_.CoverTab[115908]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:226
		// _ = "end of CoverTab[115908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:227
		_go_fuzz_dep_.CoverTab[115909]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:227
		// _ = "end of CoverTab[115909]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:227
	// _ = "end of CoverTab[115906]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:227
	_go_fuzz_dep_.CoverTab[115907]++
												return time.Duration(v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:228
	// _ = "end of CoverTab[115907]"
}

// MustGetDuration parses the expanded value as an time.Duration (in ns) if
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:231
// the key exists. If key does not exist or the value cannot be parsed the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:231
// function panics. In almost all cases you want to use MustGetParsedDuration().
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:234
func (p *Properties) MustGetDuration(key string) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:234
	_go_fuzz_dep_.CoverTab[115910]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:236
		_go_fuzz_dep_.CoverTab[115912]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:237
		// _ = "end of CoverTab[115912]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:238
		_go_fuzz_dep_.CoverTab[115913]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:238
		// _ = "end of CoverTab[115913]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:238
	// _ = "end of CoverTab[115910]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:238
	_go_fuzz_dep_.CoverTab[115911]++
												return time.Duration(v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:239
	// _ = "end of CoverTab[115911]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:244
// GetParsedDuration parses the expanded value with time.ParseDuration() if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:244
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:244
// value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:247
func (p *Properties) GetParsedDuration(key string, def time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:247
	_go_fuzz_dep_.CoverTab[115914]++
												s, ok := p.Get(key)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:249
		_go_fuzz_dep_.CoverTab[115917]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:250
		// _ = "end of CoverTab[115917]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:251
		_go_fuzz_dep_.CoverTab[115918]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:251
		// _ = "end of CoverTab[115918]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:251
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:251
	// _ = "end of CoverTab[115914]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:251
	_go_fuzz_dep_.CoverTab[115915]++
												v, err := time.ParseDuration(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:253
		_go_fuzz_dep_.CoverTab[115919]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:254
		// _ = "end of CoverTab[115919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:255
		_go_fuzz_dep_.CoverTab[115920]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:255
		// _ = "end of CoverTab[115920]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:255
	// _ = "end of CoverTab[115915]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:255
	_go_fuzz_dep_.CoverTab[115916]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:256
	// _ = "end of CoverTab[115916]"
}

// MustGetParsedDuration parses the expanded value with time.ParseDuration() if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:259
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:261
func (p *Properties) MustGetParsedDuration(key string) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:261
	_go_fuzz_dep_.CoverTab[115921]++
												s, ok := p.Get(key)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:263
		_go_fuzz_dep_.CoverTab[115924]++
													ErrorHandler(invalidKeyError(key))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:264
		// _ = "end of CoverTab[115924]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:265
		_go_fuzz_dep_.CoverTab[115925]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:265
		// _ = "end of CoverTab[115925]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:265
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:265
	// _ = "end of CoverTab[115921]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:265
	_go_fuzz_dep_.CoverTab[115922]++
												v, err := time.ParseDuration(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:267
		_go_fuzz_dep_.CoverTab[115926]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:268
		// _ = "end of CoverTab[115926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:269
		_go_fuzz_dep_.CoverTab[115927]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:269
		// _ = "end of CoverTab[115927]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:269
	// _ = "end of CoverTab[115922]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:269
	_go_fuzz_dep_.CoverTab[115923]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:270
	// _ = "end of CoverTab[115923]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:275
// GetFloat64 parses the expanded value as a float64 if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:275
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:275
// value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:278
func (p *Properties) GetFloat64(key string, def float64) float64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:278
	_go_fuzz_dep_.CoverTab[115928]++
												v, err := p.getFloat64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:280
		_go_fuzz_dep_.CoverTab[115930]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:281
		// _ = "end of CoverTab[115930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:282
		_go_fuzz_dep_.CoverTab[115931]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:282
		// _ = "end of CoverTab[115931]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:282
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:282
	// _ = "end of CoverTab[115928]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:282
	_go_fuzz_dep_.CoverTab[115929]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:283
	// _ = "end of CoverTab[115929]"
}

// MustGetFloat64 parses the expanded value as a float64 if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:286
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:288
func (p *Properties) MustGetFloat64(key string) float64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:288
	_go_fuzz_dep_.CoverTab[115932]++
												v, err := p.getFloat64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:290
		_go_fuzz_dep_.CoverTab[115934]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:291
		// _ = "end of CoverTab[115934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:292
		_go_fuzz_dep_.CoverTab[115935]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:292
		// _ = "end of CoverTab[115935]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:292
	// _ = "end of CoverTab[115932]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:292
	_go_fuzz_dep_.CoverTab[115933]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:293
	// _ = "end of CoverTab[115933]"
}

func (p *Properties) getFloat64(key string) (value float64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:296
	_go_fuzz_dep_.CoverTab[115936]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:297
		_go_fuzz_dep_.CoverTab[115938]++
													value, err = strconv.ParseFloat(v, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:299
			_go_fuzz_dep_.CoverTab[115940]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:300
			// _ = "end of CoverTab[115940]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:301
			_go_fuzz_dep_.CoverTab[115941]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:301
			// _ = "end of CoverTab[115941]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:301
		// _ = "end of CoverTab[115938]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:301
		_go_fuzz_dep_.CoverTab[115939]++
													return value, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:302
		// _ = "end of CoverTab[115939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:303
		_go_fuzz_dep_.CoverTab[115942]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:303
		// _ = "end of CoverTab[115942]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:303
	// _ = "end of CoverTab[115936]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:303
	_go_fuzz_dep_.CoverTab[115937]++
												return 0, invalidKeyError(key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:304
	// _ = "end of CoverTab[115937]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:309
// GetInt parses the expanded value as an int if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:309
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:309
// value is returned. If the value does not fit into an int the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:309
// function panics with an out of range error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:313
func (p *Properties) GetInt(key string, def int) int {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:313
	_go_fuzz_dep_.CoverTab[115943]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:315
		_go_fuzz_dep_.CoverTab[115945]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:316
		// _ = "end of CoverTab[115945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:317
		_go_fuzz_dep_.CoverTab[115946]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:317
		// _ = "end of CoverTab[115946]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:317
	// _ = "end of CoverTab[115943]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:317
	_go_fuzz_dep_.CoverTab[115944]++
												return intRangeCheck(key, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:318
	// _ = "end of CoverTab[115944]"
}

// MustGetInt parses the expanded value as an int if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:321
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:321
// If the value does not fit into an int the function panics with
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:321
// an out of range error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:325
func (p *Properties) MustGetInt(key string) int {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:325
	_go_fuzz_dep_.CoverTab[115947]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:327
		_go_fuzz_dep_.CoverTab[115949]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:328
		// _ = "end of CoverTab[115949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:329
		_go_fuzz_dep_.CoverTab[115950]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:329
		// _ = "end of CoverTab[115950]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:329
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:329
	// _ = "end of CoverTab[115947]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:329
	_go_fuzz_dep_.CoverTab[115948]++
												return intRangeCheck(key, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:330
	// _ = "end of CoverTab[115948]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:335
// GetInt64 parses the expanded value as an int64 if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:335
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:335
// value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:338
func (p *Properties) GetInt64(key string, def int64) int64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:338
	_go_fuzz_dep_.CoverTab[115951]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:340
		_go_fuzz_dep_.CoverTab[115953]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:341
		// _ = "end of CoverTab[115953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:342
		_go_fuzz_dep_.CoverTab[115954]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:342
		// _ = "end of CoverTab[115954]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:342
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:342
	// _ = "end of CoverTab[115951]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:342
	_go_fuzz_dep_.CoverTab[115952]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:343
	// _ = "end of CoverTab[115952]"
}

// MustGetInt64 parses the expanded value as an int if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:346
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:348
func (p *Properties) MustGetInt64(key string) int64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:348
	_go_fuzz_dep_.CoverTab[115955]++
												v, err := p.getInt64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:350
		_go_fuzz_dep_.CoverTab[115957]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:351
		// _ = "end of CoverTab[115957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:352
		_go_fuzz_dep_.CoverTab[115958]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:352
		// _ = "end of CoverTab[115958]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:352
	// _ = "end of CoverTab[115955]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:352
	_go_fuzz_dep_.CoverTab[115956]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:353
	// _ = "end of CoverTab[115956]"
}

func (p *Properties) getInt64(key string) (value int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:356
	_go_fuzz_dep_.CoverTab[115959]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:357
		_go_fuzz_dep_.CoverTab[115961]++
													value, err = strconv.ParseInt(v, 10, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:359
			_go_fuzz_dep_.CoverTab[115963]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:360
			// _ = "end of CoverTab[115963]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:361
			_go_fuzz_dep_.CoverTab[115964]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:361
			// _ = "end of CoverTab[115964]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:361
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:361
		// _ = "end of CoverTab[115961]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:361
		_go_fuzz_dep_.CoverTab[115962]++
													return value, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:362
		// _ = "end of CoverTab[115962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:363
		_go_fuzz_dep_.CoverTab[115965]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:363
		// _ = "end of CoverTab[115965]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:363
	// _ = "end of CoverTab[115959]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:363
	_go_fuzz_dep_.CoverTab[115960]++
												return 0, invalidKeyError(key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:364
	// _ = "end of CoverTab[115960]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:369
// GetUint parses the expanded value as an uint if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:369
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:369
// value is returned. If the value does not fit into an int the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:369
// function panics with an out of range error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:373
func (p *Properties) GetUint(key string, def uint) uint {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:373
	_go_fuzz_dep_.CoverTab[115966]++
												v, err := p.getUint64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:375
		_go_fuzz_dep_.CoverTab[115968]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:376
		// _ = "end of CoverTab[115968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:377
		_go_fuzz_dep_.CoverTab[115969]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:377
		// _ = "end of CoverTab[115969]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:377
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:377
	// _ = "end of CoverTab[115966]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:377
	_go_fuzz_dep_.CoverTab[115967]++
												return uintRangeCheck(key, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:378
	// _ = "end of CoverTab[115967]"
}

// MustGetUint parses the expanded value as an int if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:381
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:381
// If the value does not fit into an int the function panics with
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:381
// an out of range error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:385
func (p *Properties) MustGetUint(key string) uint {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:385
	_go_fuzz_dep_.CoverTab[115970]++
												v, err := p.getUint64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:387
		_go_fuzz_dep_.CoverTab[115972]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:388
		// _ = "end of CoverTab[115972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:389
		_go_fuzz_dep_.CoverTab[115973]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:389
		// _ = "end of CoverTab[115973]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:389
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:389
	// _ = "end of CoverTab[115970]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:389
	_go_fuzz_dep_.CoverTab[115971]++
												return uintRangeCheck(key, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:390
	// _ = "end of CoverTab[115971]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:395
// GetUint64 parses the expanded value as an uint64 if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:395
// If key does not exist or the value cannot be parsed the default
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:395
// value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:398
func (p *Properties) GetUint64(key string, def uint64) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:398
	_go_fuzz_dep_.CoverTab[115974]++
												v, err := p.getUint64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:400
		_go_fuzz_dep_.CoverTab[115976]++
													return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:401
		// _ = "end of CoverTab[115976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:402
		_go_fuzz_dep_.CoverTab[115977]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:402
		// _ = "end of CoverTab[115977]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:402
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:402
	// _ = "end of CoverTab[115974]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:402
	_go_fuzz_dep_.CoverTab[115975]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:403
	// _ = "end of CoverTab[115975]"
}

// MustGetUint64 parses the expanded value as an int if the key exists.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:406
// If key does not exist or the value cannot be parsed the function panics.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:408
func (p *Properties) MustGetUint64(key string) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:408
	_go_fuzz_dep_.CoverTab[115978]++
												v, err := p.getUint64(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:410
		_go_fuzz_dep_.CoverTab[115980]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:411
		// _ = "end of CoverTab[115980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:412
		_go_fuzz_dep_.CoverTab[115981]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:412
		// _ = "end of CoverTab[115981]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:412
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:412
	// _ = "end of CoverTab[115978]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:412
	_go_fuzz_dep_.CoverTab[115979]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:413
	// _ = "end of CoverTab[115979]"
}

func (p *Properties) getUint64(key string) (value uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:416
	_go_fuzz_dep_.CoverTab[115982]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:417
		_go_fuzz_dep_.CoverTab[115984]++
													value, err = strconv.ParseUint(v, 10, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:419
			_go_fuzz_dep_.CoverTab[115986]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:420
			// _ = "end of CoverTab[115986]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:421
			_go_fuzz_dep_.CoverTab[115987]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:421
			// _ = "end of CoverTab[115987]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:421
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:421
		// _ = "end of CoverTab[115984]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:421
		_go_fuzz_dep_.CoverTab[115985]++
													return value, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:422
		// _ = "end of CoverTab[115985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:423
		_go_fuzz_dep_.CoverTab[115988]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:423
		// _ = "end of CoverTab[115988]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:423
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:423
	// _ = "end of CoverTab[115982]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:423
	_go_fuzz_dep_.CoverTab[115983]++
												return 0, invalidKeyError(key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:424
	// _ = "end of CoverTab[115983]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:429
// GetString returns the expanded value for the given key if exists or
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:429
// the default value otherwise.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:431
func (p *Properties) GetString(key, def string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:431
	_go_fuzz_dep_.CoverTab[115989]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:432
		_go_fuzz_dep_.CoverTab[115991]++
													return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:433
		// _ = "end of CoverTab[115991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:434
		_go_fuzz_dep_.CoverTab[115992]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:434
		// _ = "end of CoverTab[115992]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:434
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:434
	// _ = "end of CoverTab[115989]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:434
	_go_fuzz_dep_.CoverTab[115990]++
												return def
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:435
	// _ = "end of CoverTab[115990]"
}

// MustGetString returns the expanded value for the given key if exists or
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:438
// panics otherwise.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:440
func (p *Properties) MustGetString(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:440
	_go_fuzz_dep_.CoverTab[115993]++
												if v, ok := p.Get(key); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:441
		_go_fuzz_dep_.CoverTab[115995]++
													return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:442
		// _ = "end of CoverTab[115995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:443
		_go_fuzz_dep_.CoverTab[115996]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:443
		// _ = "end of CoverTab[115996]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:443
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:443
	// _ = "end of CoverTab[115993]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:443
	_go_fuzz_dep_.CoverTab[115994]++
												ErrorHandler(invalidKeyError(key))
												panic("ErrorHandler should exit")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:445
	// _ = "end of CoverTab[115994]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:450
// Filter returns a new properties object which contains all properties
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:450
// for which the key matches the pattern.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:452
func (p *Properties) Filter(pattern string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:452
	_go_fuzz_dep_.CoverTab[115997]++
												re, err := regexp.Compile(pattern)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:454
		_go_fuzz_dep_.CoverTab[115999]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:455
		// _ = "end of CoverTab[115999]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:456
		_go_fuzz_dep_.CoverTab[116000]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:456
		// _ = "end of CoverTab[116000]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:456
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:456
	// _ = "end of CoverTab[115997]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:456
	_go_fuzz_dep_.CoverTab[115998]++

												return p.FilterRegexp(re), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:458
	// _ = "end of CoverTab[115998]"
}

// FilterRegexp returns a new properties object which contains all properties
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:461
// for which the key matches the regular expression.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:463
func (p *Properties) FilterRegexp(re *regexp.Regexp) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:463
	_go_fuzz_dep_.CoverTab[116001]++
												pp := NewProperties()
												for _, k := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:465
		_go_fuzz_dep_.CoverTab[116003]++
													if re.MatchString(k) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:466
			_go_fuzz_dep_.CoverTab[116004]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:469
			pp.Set(k, p.m[k])
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:469
			// _ = "end of CoverTab[116004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:470
			_go_fuzz_dep_.CoverTab[116005]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:470
			// _ = "end of CoverTab[116005]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:470
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:470
		// _ = "end of CoverTab[116003]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:471
	// _ = "end of CoverTab[116001]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:471
	_go_fuzz_dep_.CoverTab[116002]++
												return pp
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:472
	// _ = "end of CoverTab[116002]"
}

// FilterPrefix returns a new properties object with a subset of all keys
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:475
// with the given prefix.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:477
func (p *Properties) FilterPrefix(prefix string) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:477
	_go_fuzz_dep_.CoverTab[116006]++
												pp := NewProperties()
												for _, k := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:479
		_go_fuzz_dep_.CoverTab[116008]++
													if strings.HasPrefix(k, prefix) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:480
			_go_fuzz_dep_.CoverTab[116009]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:483
			pp.Set(k, p.m[k])
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:483
			// _ = "end of CoverTab[116009]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:484
			_go_fuzz_dep_.CoverTab[116010]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:484
			// _ = "end of CoverTab[116010]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:484
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:484
		// _ = "end of CoverTab[116008]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:485
	// _ = "end of CoverTab[116006]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:485
	_go_fuzz_dep_.CoverTab[116007]++
												return pp
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:486
	// _ = "end of CoverTab[116007]"
}

// FilterStripPrefix returns a new properties object with a subset of all keys
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:489
// with the given prefix and the prefix removed from the keys.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:491
func (p *Properties) FilterStripPrefix(prefix string) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:491
	_go_fuzz_dep_.CoverTab[116011]++
												pp := NewProperties()
												n := len(prefix)
												for _, k := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:494
		_go_fuzz_dep_.CoverTab[116013]++
													if len(k) > len(prefix) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:495
			_go_fuzz_dep_.CoverTab[116014]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:495
			return strings.HasPrefix(k, prefix)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:495
			// _ = "end of CoverTab[116014]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:495
		}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:495
			_go_fuzz_dep_.CoverTab[116015]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:499
			pp.Set(k[n:], p.m[k])
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:499
			// _ = "end of CoverTab[116015]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:500
			_go_fuzz_dep_.CoverTab[116016]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:500
			// _ = "end of CoverTab[116016]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:500
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:500
		// _ = "end of CoverTab[116013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:501
	// _ = "end of CoverTab[116011]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:501
	_go_fuzz_dep_.CoverTab[116012]++
												return pp
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:502
	// _ = "end of CoverTab[116012]"
}

// Len returns the number of keys.
func (p *Properties) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:506
	_go_fuzz_dep_.CoverTab[116017]++
												return len(p.m)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:507
	// _ = "end of CoverTab[116017]"
}

// Keys returns all keys in the same order as in the input.
func (p *Properties) Keys() []string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:511
	_go_fuzz_dep_.CoverTab[116018]++
												keys := make([]string, len(p.k))
												copy(keys, p.k)
												return keys
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:514
	// _ = "end of CoverTab[116018]"
}

// Set sets the property key to the corresponding value.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:517
// If a value for key existed before then ok is true and prev
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:517
// contains the previous value. If the value contains a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:517
// circular reference or a malformed expression then
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:517
// an error is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:517
// An empty key is silently ignored.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:523
func (p *Properties) Set(key, value string) (prev string, ok bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:523
	_go_fuzz_dep_.CoverTab[116019]++
												if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:524
		_go_fuzz_dep_.CoverTab[116024]++
													return "", false, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:525
		// _ = "end of CoverTab[116024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:526
		_go_fuzz_dep_.CoverTab[116025]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:526
		// _ = "end of CoverTab[116025]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:526
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:526
	// _ = "end of CoverTab[116019]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:526
	_go_fuzz_dep_.CoverTab[116020]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:529
	if p.DisableExpansion {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:529
		_go_fuzz_dep_.CoverTab[116026]++
													prev, ok = p.Get(key)
													p.m[key] = value
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:532
			_go_fuzz_dep_.CoverTab[116028]++
														p.k = append(p.k, key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:533
			// _ = "end of CoverTab[116028]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:534
			_go_fuzz_dep_.CoverTab[116029]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:534
			// _ = "end of CoverTab[116029]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:534
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:534
		// _ = "end of CoverTab[116026]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:534
		_go_fuzz_dep_.CoverTab[116027]++
													return prev, ok, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:535
		// _ = "end of CoverTab[116027]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:536
		_go_fuzz_dep_.CoverTab[116030]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:536
		// _ = "end of CoverTab[116030]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:536
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:536
	// _ = "end of CoverTab[116020]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:536
	_go_fuzz_dep_.CoverTab[116021]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:542
	prev, ok = p.Get(key)
												p.m[key] = value

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:546
	_, err = p.expand(key, value)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:547
		_go_fuzz_dep_.CoverTab[116031]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:550
		if ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:550
			_go_fuzz_dep_.CoverTab[116033]++
														p.m[key] = prev
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:551
			// _ = "end of CoverTab[116033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:552
			_go_fuzz_dep_.CoverTab[116034]++
														delete(p.m, key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:553
			// _ = "end of CoverTab[116034]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:554
		// _ = "end of CoverTab[116031]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:554
		_go_fuzz_dep_.CoverTab[116032]++

													return "", false, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:556
		// _ = "end of CoverTab[116032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:557
		_go_fuzz_dep_.CoverTab[116035]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:557
		// _ = "end of CoverTab[116035]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:557
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:557
	// _ = "end of CoverTab[116021]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:557
	_go_fuzz_dep_.CoverTab[116022]++

												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:559
		_go_fuzz_dep_.CoverTab[116036]++
													p.k = append(p.k, key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:560
		// _ = "end of CoverTab[116036]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:561
		_go_fuzz_dep_.CoverTab[116037]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:561
		// _ = "end of CoverTab[116037]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:561
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:561
	// _ = "end of CoverTab[116022]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:561
	_go_fuzz_dep_.CoverTab[116023]++

												return prev, ok, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:563
	// _ = "end of CoverTab[116023]"
}

// SetValue sets property key to the default string value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:566
// as defined by fmt.Sprintf("%v").
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:568
func (p *Properties) SetValue(key string, value interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:568
	_go_fuzz_dep_.CoverTab[116038]++
												_, _, err := p.Set(key, fmt.Sprintf("%v", value))
												return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:570
	// _ = "end of CoverTab[116038]"
}

// MustSet sets the property key to the corresponding value.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:573
// If a value for key existed before then ok is true and prev
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:573
// contains the previous value. An empty key is silently ignored.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:576
func (p *Properties) MustSet(key, value string) (prev string, ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:576
	_go_fuzz_dep_.CoverTab[116039]++
												prev, ok, err := p.Set(key, value)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:578
		_go_fuzz_dep_.CoverTab[116041]++
													ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:579
		// _ = "end of CoverTab[116041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:580
		_go_fuzz_dep_.CoverTab[116042]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:580
		// _ = "end of CoverTab[116042]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:580
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:580
	// _ = "end of CoverTab[116039]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:580
	_go_fuzz_dep_.CoverTab[116040]++
												return prev, ok
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:581
	// _ = "end of CoverTab[116040]"
}

// String returns a string of all expanded 'key = value' pairs.
func (p *Properties) String() string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:585
	_go_fuzz_dep_.CoverTab[116043]++
												var s string
												for _, key := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:587
		_go_fuzz_dep_.CoverTab[116045]++
													value, _ := p.Get(key)
													s = fmt.Sprintf("%s%s = %s\n", s, key, value)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:589
		// _ = "end of CoverTab[116045]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:590
	// _ = "end of CoverTab[116043]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:590
	_go_fuzz_dep_.CoverTab[116044]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:591
	// _ = "end of CoverTab[116044]"
}

// Sort sorts the properties keys in alphabetical order.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:594
// This is helpfully before writing the properties.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:596
func (p *Properties) Sort() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:596
	_go_fuzz_dep_.CoverTab[116046]++
												sort.Strings(p.k)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:597
	// _ = "end of CoverTab[116046]"
}

// Write writes all unexpanded 'key = value' pairs to the given writer.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:600
// Write returns the number of bytes written and any write error encountered.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:602
func (p *Properties) Write(w io.Writer, enc Encoding) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:602
	_go_fuzz_dep_.CoverTab[116047]++
												return p.WriteComment(w, "", enc)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:603
	// _ = "end of CoverTab[116047]"
}

// WriteComment writes all unexpanced 'key = value' pairs to the given writer.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:606
// If prefix is not empty then comments are written with a blank line and the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:606
// given prefix. The prefix should be either "# " or "! " to be compatible with
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:606
// the properties file format. Otherwise, the properties parser will not be
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:606
// able to read the file back in. It returns the number of bytes written and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:606
// any write error encountered.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:612
func (p *Properties) WriteComment(w io.Writer, prefix string, enc Encoding) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:612
	_go_fuzz_dep_.CoverTab[116048]++
												var x int

												for _, key := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:615
		_go_fuzz_dep_.CoverTab[116050]++
													value := p.m[key]

													if prefix != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:618
			_go_fuzz_dep_.CoverTab[116054]++
														if comments, ok := p.c[key]; ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:619
				_go_fuzz_dep_.CoverTab[116055]++

															allEmpty := true
															for _, c := range comments {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:622
					_go_fuzz_dep_.CoverTab[116057]++
																if c != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:623
						_go_fuzz_dep_.CoverTab[116058]++
																	allEmpty = false
																	break
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:625
						// _ = "end of CoverTab[116058]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:626
						_go_fuzz_dep_.CoverTab[116059]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:626
						// _ = "end of CoverTab[116059]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:626
					}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:626
					// _ = "end of CoverTab[116057]"
				}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:627
				// _ = "end of CoverTab[116055]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:627
				_go_fuzz_dep_.CoverTab[116056]++

															if !allEmpty {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:629
					_go_fuzz_dep_.CoverTab[116060]++

																if len(comments) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:631
						_go_fuzz_dep_.CoverTab[116062]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:631
						return n > 0
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:631
						// _ = "end of CoverTab[116062]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:631
					}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:631
						_go_fuzz_dep_.CoverTab[116063]++
																	x, err = fmt.Fprintln(w)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:633
							_go_fuzz_dep_.CoverTab[116065]++
																		return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:634
							// _ = "end of CoverTab[116065]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:635
							_go_fuzz_dep_.CoverTab[116066]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:635
							// _ = "end of CoverTab[116066]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:635
						}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:635
						// _ = "end of CoverTab[116063]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:635
						_go_fuzz_dep_.CoverTab[116064]++
																	n += x
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:636
						// _ = "end of CoverTab[116064]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:637
						_go_fuzz_dep_.CoverTab[116067]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:637
						// _ = "end of CoverTab[116067]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:637
					}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:637
					// _ = "end of CoverTab[116060]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:637
					_go_fuzz_dep_.CoverTab[116061]++

																for _, c := range comments {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:639
						_go_fuzz_dep_.CoverTab[116068]++
																	x, err = fmt.Fprintf(w, "%s%s\n", prefix, c)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:641
							_go_fuzz_dep_.CoverTab[116070]++
																		return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:642
							// _ = "end of CoverTab[116070]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:643
							_go_fuzz_dep_.CoverTab[116071]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:643
							// _ = "end of CoverTab[116071]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:643
						}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:643
						// _ = "end of CoverTab[116068]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:643
						_go_fuzz_dep_.CoverTab[116069]++
																	n += x
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:644
						// _ = "end of CoverTab[116069]"
					}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:645
					// _ = "end of CoverTab[116061]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:646
					_go_fuzz_dep_.CoverTab[116072]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:646
					// _ = "end of CoverTab[116072]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:646
				}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:646
				// _ = "end of CoverTab[116056]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:647
				_go_fuzz_dep_.CoverTab[116073]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:647
				// _ = "end of CoverTab[116073]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:647
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:647
			// _ = "end of CoverTab[116054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:648
			_go_fuzz_dep_.CoverTab[116074]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:648
			// _ = "end of CoverTab[116074]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:648
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:648
		// _ = "end of CoverTab[116050]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:648
		_go_fuzz_dep_.CoverTab[116051]++
													sep := " = "
													if p.WriteSeparator != "" {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:650
			_go_fuzz_dep_.CoverTab[116075]++
														sep = p.WriteSeparator
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:651
			// _ = "end of CoverTab[116075]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:652
			_go_fuzz_dep_.CoverTab[116076]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:652
			// _ = "end of CoverTab[116076]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:652
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:652
		// _ = "end of CoverTab[116051]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:652
		_go_fuzz_dep_.CoverTab[116052]++
													x, err = fmt.Fprintf(w, "%s%s%s\n", encode(key, " :", enc), sep, encode(value, "", enc))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:654
			_go_fuzz_dep_.CoverTab[116077]++
														return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:655
			// _ = "end of CoverTab[116077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:656
			_go_fuzz_dep_.CoverTab[116078]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:656
			// _ = "end of CoverTab[116078]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:656
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:656
		// _ = "end of CoverTab[116052]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:656
		_go_fuzz_dep_.CoverTab[116053]++
													n += x
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:657
		// _ = "end of CoverTab[116053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:658
	// _ = "end of CoverTab[116048]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:658
	_go_fuzz_dep_.CoverTab[116049]++
												return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:659
	// _ = "end of CoverTab[116049]"
}

// Map returns a copy of the properties as a map.
func (p *Properties) Map() map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:663
	_go_fuzz_dep_.CoverTab[116079]++
												m := make(map[string]string)
												for k, v := range p.m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:665
		_go_fuzz_dep_.CoverTab[116081]++
													m[k] = v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:666
		// _ = "end of CoverTab[116081]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:667
	// _ = "end of CoverTab[116079]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:667
	_go_fuzz_dep_.CoverTab[116080]++
												return m
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:668
	// _ = "end of CoverTab[116080]"
}

// FilterFunc returns a copy of the properties which includes the values which passed all filters.
func (p *Properties) FilterFunc(filters ...func(k, v string) bool) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:672
	_go_fuzz_dep_.CoverTab[116082]++
												pp := NewProperties()
outer:
	for k, v := range p.m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:675
		_go_fuzz_dep_.CoverTab[116084]++
													for _, f := range filters {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:676
			_go_fuzz_dep_.CoverTab[116085]++
														if !f(k, v) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:677
				_go_fuzz_dep_.CoverTab[116087]++
															continue outer
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:678
				// _ = "end of CoverTab[116087]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:679
				_go_fuzz_dep_.CoverTab[116088]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:679
				// _ = "end of CoverTab[116088]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:679
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:679
			// _ = "end of CoverTab[116085]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:679
			_go_fuzz_dep_.CoverTab[116086]++
														pp.Set(k, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:680
			// _ = "end of CoverTab[116086]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:681
		// _ = "end of CoverTab[116084]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:682
	// _ = "end of CoverTab[116082]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:682
	_go_fuzz_dep_.CoverTab[116083]++
												return pp
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:683
	// _ = "end of CoverTab[116083]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:688
// Delete removes the key and its comments.
func (p *Properties) Delete(key string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:689
	_go_fuzz_dep_.CoverTab[116089]++
												delete(p.m, key)
												delete(p.c, key)
												newKeys := []string{}
												for _, k := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:693
		_go_fuzz_dep_.CoverTab[116091]++
													if k != key {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:694
			_go_fuzz_dep_.CoverTab[116092]++
														newKeys = append(newKeys, k)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:695
			// _ = "end of CoverTab[116092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:696
			_go_fuzz_dep_.CoverTab[116093]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:696
			// _ = "end of CoverTab[116093]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:696
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:696
		// _ = "end of CoverTab[116091]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:697
	// _ = "end of CoverTab[116089]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:697
	_go_fuzz_dep_.CoverTab[116090]++
												p.k = newKeys
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:698
	// _ = "end of CoverTab[116090]"
}

// Merge merges properties, comments and keys from other *Properties into p
func (p *Properties) Merge(other *Properties) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:702
	_go_fuzz_dep_.CoverTab[116094]++
												for k, v := range other.m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:703
		_go_fuzz_dep_.CoverTab[116097]++
													p.m[k] = v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:704
		// _ = "end of CoverTab[116097]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:705
	// _ = "end of CoverTab[116094]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:705
	_go_fuzz_dep_.CoverTab[116095]++
												for k, v := range other.c {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:706
		_go_fuzz_dep_.CoverTab[116098]++
													p.c[k] = v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:707
		// _ = "end of CoverTab[116098]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:708
	// _ = "end of CoverTab[116095]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:708
	_go_fuzz_dep_.CoverTab[116096]++

outer:
	for _, otherKey := range other.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:711
		_go_fuzz_dep_.CoverTab[116099]++
													for _, key := range p.k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:712
			_go_fuzz_dep_.CoverTab[116101]++
														if otherKey == key {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:713
				_go_fuzz_dep_.CoverTab[116102]++
															continue outer
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:714
				// _ = "end of CoverTab[116102]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:715
				_go_fuzz_dep_.CoverTab[116103]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:715
				// _ = "end of CoverTab[116103]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:715
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:715
			// _ = "end of CoverTab[116101]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:716
		// _ = "end of CoverTab[116099]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:716
		_go_fuzz_dep_.CoverTab[116100]++
													p.k = append(p.k, otherKey)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:717
		// _ = "end of CoverTab[116100]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:718
	// _ = "end of CoverTab[116096]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:723
// check expands all values and returns an error if a circular reference or
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:723
// a malformed expression was found.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:725
func (p *Properties) check() error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:725
	_go_fuzz_dep_.CoverTab[116104]++
												for key, value := range p.m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:726
		_go_fuzz_dep_.CoverTab[116106]++
													if _, err := p.expand(key, value); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:727
			_go_fuzz_dep_.CoverTab[116107]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:728
			// _ = "end of CoverTab[116107]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:729
			_go_fuzz_dep_.CoverTab[116108]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:729
			// _ = "end of CoverTab[116108]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:729
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:729
		// _ = "end of CoverTab[116106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:730
	// _ = "end of CoverTab[116104]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:730
	_go_fuzz_dep_.CoverTab[116105]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:731
	// _ = "end of CoverTab[116105]"
}

func (p *Properties) expand(key, input string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:734
	_go_fuzz_dep_.CoverTab[116109]++

												if p.Prefix == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:736
		_go_fuzz_dep_.CoverTab[116111]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:736
		return p.Postfix == ""
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:736
		// _ = "end of CoverTab[116111]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:736
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:736
		_go_fuzz_dep_.CoverTab[116112]++
													return input, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:737
		// _ = "end of CoverTab[116112]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:738
		_go_fuzz_dep_.CoverTab[116113]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:738
		// _ = "end of CoverTab[116113]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:738
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:738
	// _ = "end of CoverTab[116109]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:738
	_go_fuzz_dep_.CoverTab[116110]++

												return expand(input, []string{key}, p.Prefix, p.Postfix, p.m)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:740
	// _ = "end of CoverTab[116110]"
}

// expand recursively expands expressions of '(prefix)key(postfix)' to their corresponding values.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:743
// The function keeps track of the keys that were already expanded and stops if it
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:743
// detects a circular reference or a malformed expression of the form '(prefix)key'.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:746
func expand(s string, keys []string, prefix, postfix string, values map[string]string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:746
	_go_fuzz_dep_.CoverTab[116114]++
												if len(keys) > maxExpansionDepth {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:747
		_go_fuzz_dep_.CoverTab[116117]++
													return "", fmt.Errorf("expansion too deep")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:748
		// _ = "end of CoverTab[116117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:749
		_go_fuzz_dep_.CoverTab[116118]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:749
		// _ = "end of CoverTab[116118]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:749
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:749
	// _ = "end of CoverTab[116114]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:749
	_go_fuzz_dep_.CoverTab[116115]++

												for {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:751
		_go_fuzz_dep_.CoverTab[116119]++
													start := strings.Index(s, prefix)
													if start == -1 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:753
			_go_fuzz_dep_.CoverTab[116125]++
														return s, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:754
			// _ = "end of CoverTab[116125]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:755
			_go_fuzz_dep_.CoverTab[116126]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:755
			// _ = "end of CoverTab[116126]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:755
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:755
		// _ = "end of CoverTab[116119]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:755
		_go_fuzz_dep_.CoverTab[116120]++

													keyStart := start + len(prefix)
													keyLen := strings.Index(s[keyStart:], postfix)
													if keyLen == -1 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:759
			_go_fuzz_dep_.CoverTab[116127]++
														return "", fmt.Errorf("malformed expression")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:760
			// _ = "end of CoverTab[116127]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:761
			_go_fuzz_dep_.CoverTab[116128]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:761
			// _ = "end of CoverTab[116128]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:761
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:761
		// _ = "end of CoverTab[116120]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:761
		_go_fuzz_dep_.CoverTab[116121]++

													end := keyStart + keyLen + len(postfix) - 1
													key := s[keyStart : keyStart+keyLen]

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:768
		for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:768
			_go_fuzz_dep_.CoverTab[116129]++
														if key == k {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:769
				_go_fuzz_dep_.CoverTab[116130]++
															var b bytes.Buffer
															b.WriteString("circular reference in:\n")
															for _, k1 := range keys {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:772
					_go_fuzz_dep_.CoverTab[116132]++
																fmt.Fprintf(&b, "%s=%s\n", k1, values[k1])
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:773
					// _ = "end of CoverTab[116132]"
				}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:774
				// _ = "end of CoverTab[116130]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:774
				_go_fuzz_dep_.CoverTab[116131]++
															return "", fmt.Errorf(b.String())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:775
				// _ = "end of CoverTab[116131]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:776
				_go_fuzz_dep_.CoverTab[116133]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:776
				// _ = "end of CoverTab[116133]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:776
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:776
			// _ = "end of CoverTab[116129]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:777
		// _ = "end of CoverTab[116121]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:777
		_go_fuzz_dep_.CoverTab[116122]++

													val, ok := values[key]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:780
			_go_fuzz_dep_.CoverTab[116134]++
														val = os.Getenv(key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:781
			// _ = "end of CoverTab[116134]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:782
			_go_fuzz_dep_.CoverTab[116135]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:782
			// _ = "end of CoverTab[116135]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:782
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:782
		// _ = "end of CoverTab[116122]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:782
		_go_fuzz_dep_.CoverTab[116123]++
													new_val, err := expand(val, append(keys, key), prefix, postfix, values)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:784
			_go_fuzz_dep_.CoverTab[116136]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:785
			// _ = "end of CoverTab[116136]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:786
			_go_fuzz_dep_.CoverTab[116137]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:786
			// _ = "end of CoverTab[116137]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:786
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:786
		// _ = "end of CoverTab[116123]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:786
		_go_fuzz_dep_.CoverTab[116124]++
													s = s[:start] + new_val + s[end+1:]
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:787
		// _ = "end of CoverTab[116124]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:788
	// _ = "end of CoverTab[116115]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:788
	_go_fuzz_dep_.CoverTab[116116]++
												return s, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:789
	// _ = "end of CoverTab[116116]"
}

// encode encodes a UTF-8 string to ISO-8859-1 and escapes some characters.
func encode(s string, special string, enc Encoding) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:793
	_go_fuzz_dep_.CoverTab[116138]++
												switch enc {
	case UTF8:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:795
		_go_fuzz_dep_.CoverTab[116139]++
													return encodeUtf8(s, special)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:796
		// _ = "end of CoverTab[116139]"
	case ISO_8859_1:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:797
		_go_fuzz_dep_.CoverTab[116140]++
													return encodeIso(s, special)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:798
		// _ = "end of CoverTab[116140]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:799
		_go_fuzz_dep_.CoverTab[116141]++
													panic(fmt.Sprintf("unsupported encoding %v", enc))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:800
		// _ = "end of CoverTab[116141]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:801
	// _ = "end of CoverTab[116138]"
}

func encodeUtf8(s string, special string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:804
	_go_fuzz_dep_.CoverTab[116142]++
												v := ""
												for pos := 0; pos < len(s); {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:806
		_go_fuzz_dep_.CoverTab[116144]++
													r, w := utf8.DecodeRuneInString(s[pos:])
													pos += w
													v += escape(r, special)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:809
		// _ = "end of CoverTab[116144]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:810
	// _ = "end of CoverTab[116142]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:810
	_go_fuzz_dep_.CoverTab[116143]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:811
	// _ = "end of CoverTab[116143]"
}

func encodeIso(s string, special string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:814
	_go_fuzz_dep_.CoverTab[116145]++
												var r rune
												var w int
												var v string
												for pos := 0; pos < len(s); {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:818
		_go_fuzz_dep_.CoverTab[116147]++
													switch r, w = utf8.DecodeRuneInString(s[pos:]); {
		case r < 1<<8:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:820
			_go_fuzz_dep_.CoverTab[116149]++
														v += escape(r, special)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:821
			// _ = "end of CoverTab[116149]"
		case r < 1<<16:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:822
			_go_fuzz_dep_.CoverTab[116150]++
														v += fmt.Sprintf("\\u%04x", r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:823
			// _ = "end of CoverTab[116150]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:824
			_go_fuzz_dep_.CoverTab[116151]++
														v += "?"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:825
			// _ = "end of CoverTab[116151]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:826
		// _ = "end of CoverTab[116147]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:826
		_go_fuzz_dep_.CoverTab[116148]++
													pos += w
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:827
		// _ = "end of CoverTab[116148]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:828
	// _ = "end of CoverTab[116145]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:828
	_go_fuzz_dep_.CoverTab[116146]++
												return v
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:829
	// _ = "end of CoverTab[116146]"
}

func escape(r rune, special string) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:832
	_go_fuzz_dep_.CoverTab[116152]++
												switch r {
	case '\f':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:834
		_go_fuzz_dep_.CoverTab[116153]++
													return "\\f"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:835
		// _ = "end of CoverTab[116153]"
	case '\n':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:836
		_go_fuzz_dep_.CoverTab[116154]++
													return "\\n"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:837
		// _ = "end of CoverTab[116154]"
	case '\r':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:838
		_go_fuzz_dep_.CoverTab[116155]++
													return "\\r"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:839
		// _ = "end of CoverTab[116155]"
	case '\t':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:840
		_go_fuzz_dep_.CoverTab[116156]++
													return "\\t"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:841
		// _ = "end of CoverTab[116156]"
	case '\\':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:842
		_go_fuzz_dep_.CoverTab[116157]++
													return "\\\\"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:843
		// _ = "end of CoverTab[116157]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:844
		_go_fuzz_dep_.CoverTab[116158]++
													if strings.ContainsRune(special, r) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:845
			_go_fuzz_dep_.CoverTab[116160]++
														return "\\" + string(r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:846
			// _ = "end of CoverTab[116160]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:847
			_go_fuzz_dep_.CoverTab[116161]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:847
			// _ = "end of CoverTab[116161]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:847
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:847
		// _ = "end of CoverTab[116158]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:847
		_go_fuzz_dep_.CoverTab[116159]++
													return string(r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:848
		// _ = "end of CoverTab[116159]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:849
	// _ = "end of CoverTab[116152]"
}

func invalidKeyError(key string) error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:852
	_go_fuzz_dep_.CoverTab[116162]++
												return fmt.Errorf("unknown property: %s", key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:853
	// _ = "end of CoverTab[116162]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:854
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/properties.go:854
var _ = _go_fuzz_dep_.CoverTab
