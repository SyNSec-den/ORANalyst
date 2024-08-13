//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:1
// Package yaml implements YAML support for the Go language.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:1
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:1
// Source code and other details for the project are available at GitHub:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:1
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:1
//	https://github.com/go-yaml/yaml
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:7
)

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
)

// MapSlice encodes and decodes as a YAML map.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:18
// The order of keys is preserved when encoding and decoding.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:20
type MapSlice []MapItem

// MapItem is an item in a MapSlice.
type MapItem struct {
	Key, Value interface{}
}

// The Unmarshaler interface may be implemented by types to customize their
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:27
// behavior when being unmarshaled from a YAML document. The UnmarshalYAML
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:27
// method receives a function that may be called to unmarshal the original
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:27
// YAML value into a field or variable. It is safe to call the unmarshal
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:27
// function parameter more than once if necessary.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:32
type Unmarshaler interface {
	UnmarshalYAML(unmarshal func(interface{}) error) error
}

// The Marshaler interface may be implemented by types to customize their
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:36
// behavior when being marshaled into a YAML document. The returned value
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:36
// is marshaled in place of the original value implementing Marshaler.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:36
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:36
// If an error is returned by MarshalYAML, the marshaling procedure stops
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:36
// and returns with the provided error.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:42
type Marshaler interface {
	MarshalYAML() (interface{}, error)
}

// Unmarshal decodes the first document found within the in byte slice
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// and assigns decoded values into the out value.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// Maps and pointers (to a struct, string, int, etc) are accepted as out
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// values. If an internal pointer within a struct is not initialized,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// the yaml package will initialize it if necessary for unmarshalling
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// the provided data. The out parameter must not be nil.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// The type of the decoded values should be compatible with the respective
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// values in out. If one or more values cannot be decoded due to a type
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// mismatches, decoding continues partially until the end of the YAML
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// content, and a *yaml.TypeError is returned with details for all
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// missed values.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// Struct fields are only unmarshalled if they are exported (have an
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// upper case first letter), and are unmarshalled using the field name
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// lowercased as the default key. Custom keys may be defined via the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// "yaml" name in the field tag: the content preceding the first comma
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// is used as the key, and the following comma-separated options are
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// used to tweak the marshalling process (see Marshal).
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// Conflicting names result in a runtime error.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// For example:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	type T struct {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	    F int `yaml:"a,omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	    B int
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	var t T
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//	yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// See the documentation of Marshal for the format of tags and a list of
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:46
// supported tag options.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:80
func Unmarshal(in []byte, out interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:80
	_go_fuzz_dep_.CoverTab[127851]++
									return unmarshal(in, out, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:81
	// _ = "end of CoverTab[127851]"
}

// UnmarshalStrict is like Unmarshal except that any fields that are found
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:84
// in the data that do not have corresponding struct members, or mapping
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:84
// keys that are duplicates, will result in
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:84
// an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:88
func UnmarshalStrict(in []byte, out interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:88
	_go_fuzz_dep_.CoverTab[127852]++
									return unmarshal(in, out, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:89
	// _ = "end of CoverTab[127852]"
}

// A Decoder reads and decodes YAML values from an input stream.
type Decoder struct {
	strict	bool
	parser	*parser
}

// NewDecoder returns a new decoder that reads from r.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:98
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:98
// The decoder introduces its own buffering and may read
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:98
// data from r beyond the YAML values requested.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:102
func NewDecoder(r io.Reader) *Decoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:102
	_go_fuzz_dep_.CoverTab[127853]++
									return &Decoder{
		parser: newParserFromReader(r),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:105
	// _ = "end of CoverTab[127853]"
}

// SetStrict sets whether strict decoding behaviour is enabled when
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:108
// decoding items in the data (see UnmarshalStrict). By default, decoding is not strict.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:110
func (dec *Decoder) SetStrict(strict bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:110
	_go_fuzz_dep_.CoverTab[127854]++
									dec.strict = strict
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:111
	// _ = "end of CoverTab[127854]"
}

// Decode reads the next YAML-encoded value from its input
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:114
// and stores it in the value pointed to by v.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:114
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:114
// See the documentation for Unmarshal for details about the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:114
// conversion of YAML into a Go value.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:119
func (dec *Decoder) Decode(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:119
	_go_fuzz_dep_.CoverTab[127855]++
									d := newDecoder(dec.strict)
									defer handleErr(&err)
									node := dec.parser.parse()
									if node == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:123
		_go_fuzz_dep_.CoverTab[127859]++
										return io.EOF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:124
		// _ = "end of CoverTab[127859]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:125
		_go_fuzz_dep_.CoverTab[127860]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:125
		// _ = "end of CoverTab[127860]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:125
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:125
	// _ = "end of CoverTab[127855]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:125
	_go_fuzz_dep_.CoverTab[127856]++
									out := reflect.ValueOf(v)
									if out.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:127
		_go_fuzz_dep_.CoverTab[127861]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:127
		return !out.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:127
		// _ = "end of CoverTab[127861]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:127
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:127
		_go_fuzz_dep_.CoverTab[127862]++
										out = out.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:128
		// _ = "end of CoverTab[127862]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:129
		_go_fuzz_dep_.CoverTab[127863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:129
		// _ = "end of CoverTab[127863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:129
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:129
	// _ = "end of CoverTab[127856]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:129
	_go_fuzz_dep_.CoverTab[127857]++
									d.unmarshal(node, out)
									if len(d.terrors) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:131
		_go_fuzz_dep_.CoverTab[127864]++
										return &TypeError{d.terrors}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:132
		// _ = "end of CoverTab[127864]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:133
		_go_fuzz_dep_.CoverTab[127865]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:133
		// _ = "end of CoverTab[127865]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:133
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:133
	// _ = "end of CoverTab[127857]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:133
	_go_fuzz_dep_.CoverTab[127858]++
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:134
	// _ = "end of CoverTab[127858]"
}

func unmarshal(in []byte, out interface{}, strict bool) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:137
	_go_fuzz_dep_.CoverTab[127866]++
									defer handleErr(&err)
									d := newDecoder(strict)
									p := newParser(in)
									defer p.destroy()
									node := p.parse()
									if node != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:143
		_go_fuzz_dep_.CoverTab[127869]++
										v := reflect.ValueOf(out)
										if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:145
			_go_fuzz_dep_.CoverTab[127871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:145
			return !v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:145
			// _ = "end of CoverTab[127871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:145
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:145
			_go_fuzz_dep_.CoverTab[127872]++
											v = v.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:146
			// _ = "end of CoverTab[127872]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:147
			_go_fuzz_dep_.CoverTab[127873]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:147
			// _ = "end of CoverTab[127873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:147
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:147
		// _ = "end of CoverTab[127869]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:147
		_go_fuzz_dep_.CoverTab[127870]++
										d.unmarshal(node, v)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:148
		// _ = "end of CoverTab[127870]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:149
		_go_fuzz_dep_.CoverTab[127874]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:149
		// _ = "end of CoverTab[127874]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:149
	// _ = "end of CoverTab[127866]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:149
	_go_fuzz_dep_.CoverTab[127867]++
									if len(d.terrors) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:150
		_go_fuzz_dep_.CoverTab[127875]++
										return &TypeError{d.terrors}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:151
		// _ = "end of CoverTab[127875]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:152
		_go_fuzz_dep_.CoverTab[127876]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:152
		// _ = "end of CoverTab[127876]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:152
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:152
	// _ = "end of CoverTab[127867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:152
	_go_fuzz_dep_.CoverTab[127868]++
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:153
	// _ = "end of CoverTab[127868]"
}

// Marshal serializes the value provided into a YAML document. The structure
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// of the generated document will reflect the structure of the value itself.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// Maps and pointers (to struct, string, int, etc) are accepted as the in value.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// Struct fields are only marshalled if they are exported (have an upper case
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// first letter), and are marshalled using the field name lowercased as the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// default key. Custom keys may be defined via the "yaml" name in the field
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// tag: the content preceding the first comma is used as the key, and the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// following comma-separated options are used to tweak the marshalling process.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// Conflicting names result in a runtime error.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// The field tag format accepted is:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	`(...) yaml:"[<key>][,<flag1>[,<flag2>]]" (...)`
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// The following flags are currently supported:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	omitempty    Only include the field if it's not set to the zero
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             value for the type or to empty slices or maps.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             Zero valued structs will be omitted if all their public
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             fields are zero, unless they implement an IsZero
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             method (see the IsZeroer interface type), in which
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             case the field will be excluded if IsZero returns true.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	flow         Marshal using a flow style (useful for structs,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             sequences and maps).
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	inline       Inline the field, which must be a struct or a map,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             causing all of its fields or keys to be processed as if
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             they were part of the outer struct. For maps, keys must
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	             not conflict with the yaml keys of other struct fields.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// In addition, if the key is "-", the field is ignored.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
// For example:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	type T struct {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	    F int `yaml:"a,omitempty"`
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	    B int
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:156
//	yaml.Marshal(&T{F: 1}} // Returns "a: 1\nb: 0\n"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:199
func Marshal(in interface{}) (out []byte, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:199
	_go_fuzz_dep_.CoverTab[127877]++
									defer handleErr(&err)
									e := newEncoder()
									defer e.destroy()
									e.marshalDoc("", reflect.ValueOf(in))
									e.finish()
									out = e.out
									return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:206
	// _ = "end of CoverTab[127877]"
}

// An Encoder writes YAML values to an output stream.
type Encoder struct {
	encoder *encoder
}

// NewEncoder returns a new encoder that writes to w.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:214
// The Encoder should be closed after use to flush all data
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:214
// to w.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:217
func NewEncoder(w io.Writer) *Encoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:217
	_go_fuzz_dep_.CoverTab[127878]++
									return &Encoder{
		encoder: newEncoderWithWriter(w),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:220
	// _ = "end of CoverTab[127878]"
}

// Encode writes the YAML encoding of v to the stream.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
// If multiple items are encoded to the stream, the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
// second and subsequent document will be preceded
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
// with a "---" document separator, but the first will not.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
// See the documentation for Marshal for details about the conversion of Go
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:223
// values to YAML.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:230
func (e *Encoder) Encode(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:230
	_go_fuzz_dep_.CoverTab[127879]++
									defer handleErr(&err)
									e.encoder.marshalDoc("", reflect.ValueOf(v))
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:233
	// _ = "end of CoverTab[127879]"
}

// Close closes the encoder by writing any remaining data.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:236
// It does not write a stream terminating string "...".
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:238
func (e *Encoder) Close() (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:238
	_go_fuzz_dep_.CoverTab[127880]++
									defer handleErr(&err)
									e.encoder.finish()
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:241
	// _ = "end of CoverTab[127880]"
}

func handleErr(err *error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:244
	_go_fuzz_dep_.CoverTab[127881]++
									if v := recover(); v != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:245
		_go_fuzz_dep_.CoverTab[127882]++
										if e, ok := v.(yamlError); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:246
			_go_fuzz_dep_.CoverTab[127883]++
											*err = e.err
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:247
			// _ = "end of CoverTab[127883]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:248
			_go_fuzz_dep_.CoverTab[127884]++
											panic(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:249
			// _ = "end of CoverTab[127884]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:250
		// _ = "end of CoverTab[127882]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:251
		_go_fuzz_dep_.CoverTab[127885]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:251
		// _ = "end of CoverTab[127885]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:251
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:251
	// _ = "end of CoverTab[127881]"
}

type yamlError struct {
	err error
}

func fail(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:258
	_go_fuzz_dep_.CoverTab[127886]++
									panic(yamlError{err})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:259
	// _ = "end of CoverTab[127886]"
}

func failf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:262
	_go_fuzz_dep_.CoverTab[127887]++
									panic(yamlError{fmt.Errorf("yaml: "+format, args...)})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:263
	// _ = "end of CoverTab[127887]"
}

// A TypeError is returned by Unmarshal when one or more fields in
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:266
// the YAML document cannot be properly decoded into the requested
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:266
// types. When this error is returned, the value is still
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:266
// unmarshaled partially.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:270
type TypeError struct {
	Errors []string
}

func (e *TypeError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:274
	_go_fuzz_dep_.CoverTab[127888]++
									return fmt.Sprintf("yaml: unmarshal errors:\n  %s", strings.Join(e.Errors, "\n  "))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:275
	// _ = "end of CoverTab[127888]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:283
// structInfo holds details for the serialization of fields of
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:283
// a given struct.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:285
type structInfo struct {
	FieldsMap	map[string]fieldInfo
	FieldsList	[]fieldInfo

	// InlineMap is the number of the field in the struct that
	// contains an ,inline map, or -1 if there's none.
	InlineMap	int
}

type fieldInfo struct {
	Key		string
	Num		int
	OmitEmpty	bool
	Flow		bool
	// Id holds the unique field identifier, so we can cheaply
	// check for field duplicates without maintaining an extra map.
	Id	int

	// Inline holds the field index if the field is part of an inlined struct.
	Inline	[]int
}

var structMap = make(map[reflect.Type]*structInfo)
var fieldMapMutex sync.RWMutex

func getStructInfo(st reflect.Type) (*structInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:310
	_go_fuzz_dep_.CoverTab[127889]++
									fieldMapMutex.RLock()
									sinfo, found := structMap[st]
									fieldMapMutex.RUnlock()
									if found {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:314
		_go_fuzz_dep_.CoverTab[127892]++
										return sinfo, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:315
		// _ = "end of CoverTab[127892]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:316
		_go_fuzz_dep_.CoverTab[127893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:316
		// _ = "end of CoverTab[127893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:316
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:316
	// _ = "end of CoverTab[127889]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:316
	_go_fuzz_dep_.CoverTab[127890]++

									n := st.NumField()
									fieldsMap := make(map[string]fieldInfo)
									fieldsList := make([]fieldInfo, 0, n)
									inlineMap := -1
									for i := 0; i != n; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:322
		_go_fuzz_dep_.CoverTab[127894]++
										field := st.Field(i)
										if field.PkgPath != "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:324
			_go_fuzz_dep_.CoverTab[127902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:324
			return !field.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:324
			// _ = "end of CoverTab[127902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:324
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:324
			_go_fuzz_dep_.CoverTab[127903]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:325
			// _ = "end of CoverTab[127903]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:326
			_go_fuzz_dep_.CoverTab[127904]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:326
			// _ = "end of CoverTab[127904]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:326
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:326
		// _ = "end of CoverTab[127894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:326
		_go_fuzz_dep_.CoverTab[127895]++

										info := fieldInfo{Num: i}

										tag := field.Tag.Get("yaml")
										if tag == "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:331
			_go_fuzz_dep_.CoverTab[127905]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:331
			return strings.Index(string(field.Tag), ":") < 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:331
			// _ = "end of CoverTab[127905]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:331
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:331
			_go_fuzz_dep_.CoverTab[127906]++
											tag = string(field.Tag)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:332
			// _ = "end of CoverTab[127906]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:333
			_go_fuzz_dep_.CoverTab[127907]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:333
			// _ = "end of CoverTab[127907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:333
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:333
		// _ = "end of CoverTab[127895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:333
		_go_fuzz_dep_.CoverTab[127896]++
										if tag == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:334
			_go_fuzz_dep_.CoverTab[127908]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:335
			// _ = "end of CoverTab[127908]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:336
			_go_fuzz_dep_.CoverTab[127909]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:336
			// _ = "end of CoverTab[127909]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:336
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:336
		// _ = "end of CoverTab[127896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:336
		_go_fuzz_dep_.CoverTab[127897]++

										inline := false
										fields := strings.Split(tag, ",")
										if len(fields) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:340
			_go_fuzz_dep_.CoverTab[127910]++
											for _, flag := range fields[1:] {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:341
				_go_fuzz_dep_.CoverTab[127912]++
												switch flag {
				case "omitempty":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:343
					_go_fuzz_dep_.CoverTab[127913]++
													info.OmitEmpty = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:344
					// _ = "end of CoverTab[127913]"
				case "flow":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:345
					_go_fuzz_dep_.CoverTab[127914]++
													info.Flow = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:346
					// _ = "end of CoverTab[127914]"
				case "inline":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:347
					_go_fuzz_dep_.CoverTab[127915]++
													inline = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:348
					// _ = "end of CoverTab[127915]"
				default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:349
					_go_fuzz_dep_.CoverTab[127916]++
													return nil, errors.New(fmt.Sprintf("Unsupported flag %q in tag %q of type %s", flag, tag, st))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:350
					// _ = "end of CoverTab[127916]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:351
				// _ = "end of CoverTab[127912]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:352
			// _ = "end of CoverTab[127910]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:352
			_go_fuzz_dep_.CoverTab[127911]++
											tag = fields[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:353
			// _ = "end of CoverTab[127911]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:354
			_go_fuzz_dep_.CoverTab[127917]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:354
			// _ = "end of CoverTab[127917]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:354
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:354
		// _ = "end of CoverTab[127897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:354
		_go_fuzz_dep_.CoverTab[127898]++

										if inline {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:356
			_go_fuzz_dep_.CoverTab[127918]++
											switch field.Type.Kind() {
			case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:358
				_go_fuzz_dep_.CoverTab[127920]++
												if inlineMap >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:359
					_go_fuzz_dep_.CoverTab[127926]++
													return nil, errors.New("Multiple ,inline maps in struct " + st.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:360
					// _ = "end of CoverTab[127926]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:361
					_go_fuzz_dep_.CoverTab[127927]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:361
					// _ = "end of CoverTab[127927]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:361
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:361
				// _ = "end of CoverTab[127920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:361
				_go_fuzz_dep_.CoverTab[127921]++
												if field.Type.Key() != reflect.TypeOf("") {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:362
					_go_fuzz_dep_.CoverTab[127928]++
													return nil, errors.New("Option ,inline needs a map with string keys in struct " + st.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:363
					// _ = "end of CoverTab[127928]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:364
					_go_fuzz_dep_.CoverTab[127929]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:364
					// _ = "end of CoverTab[127929]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:364
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:364
				// _ = "end of CoverTab[127921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:364
				_go_fuzz_dep_.CoverTab[127922]++
												inlineMap = info.Num
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:365
				// _ = "end of CoverTab[127922]"
			case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:366
				_go_fuzz_dep_.CoverTab[127923]++
												sinfo, err := getStructInfo(field.Type)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:368
					_go_fuzz_dep_.CoverTab[127930]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:369
					// _ = "end of CoverTab[127930]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:370
					_go_fuzz_dep_.CoverTab[127931]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:370
					// _ = "end of CoverTab[127931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:370
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:370
				// _ = "end of CoverTab[127923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:370
				_go_fuzz_dep_.CoverTab[127924]++
												for _, finfo := range sinfo.FieldsList {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:371
					_go_fuzz_dep_.CoverTab[127932]++
													if _, found := fieldsMap[finfo.Key]; found {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:372
						_go_fuzz_dep_.CoverTab[127935]++
														msg := "Duplicated key '" + finfo.Key + "' in struct " + st.String()
														return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:374
						// _ = "end of CoverTab[127935]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:375
						_go_fuzz_dep_.CoverTab[127936]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:375
						// _ = "end of CoverTab[127936]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:375
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:375
					// _ = "end of CoverTab[127932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:375
					_go_fuzz_dep_.CoverTab[127933]++
													if finfo.Inline == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:376
						_go_fuzz_dep_.CoverTab[127937]++
														finfo.Inline = []int{i, finfo.Num}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:377
						// _ = "end of CoverTab[127937]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:378
						_go_fuzz_dep_.CoverTab[127938]++
														finfo.Inline = append([]int{i}, finfo.Inline...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:379
						// _ = "end of CoverTab[127938]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:380
					// _ = "end of CoverTab[127933]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:380
					_go_fuzz_dep_.CoverTab[127934]++
													finfo.Id = len(fieldsList)
													fieldsMap[finfo.Key] = finfo
													fieldsList = append(fieldsList, finfo)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:383
					// _ = "end of CoverTab[127934]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:384
				// _ = "end of CoverTab[127924]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:385
				_go_fuzz_dep_.CoverTab[127925]++

												return nil, errors.New("Option ,inline needs a struct value field")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:387
				// _ = "end of CoverTab[127925]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:388
			// _ = "end of CoverTab[127918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:388
			_go_fuzz_dep_.CoverTab[127919]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:389
			// _ = "end of CoverTab[127919]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:390
			_go_fuzz_dep_.CoverTab[127939]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:390
			// _ = "end of CoverTab[127939]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:390
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:390
		// _ = "end of CoverTab[127898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:390
		_go_fuzz_dep_.CoverTab[127899]++

										if tag != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:392
			_go_fuzz_dep_.CoverTab[127940]++
											info.Key = tag
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:393
			// _ = "end of CoverTab[127940]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:394
			_go_fuzz_dep_.CoverTab[127941]++
											info.Key = strings.ToLower(field.Name)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:395
			// _ = "end of CoverTab[127941]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:396
		// _ = "end of CoverTab[127899]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:396
		_go_fuzz_dep_.CoverTab[127900]++

										if _, found = fieldsMap[info.Key]; found {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:398
			_go_fuzz_dep_.CoverTab[127942]++
											msg := "Duplicated key '" + info.Key + "' in struct " + st.String()
											return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:400
			// _ = "end of CoverTab[127942]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:401
			_go_fuzz_dep_.CoverTab[127943]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:401
			// _ = "end of CoverTab[127943]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:401
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:401
		// _ = "end of CoverTab[127900]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:401
		_go_fuzz_dep_.CoverTab[127901]++

										info.Id = len(fieldsList)
										fieldsList = append(fieldsList, info)
										fieldsMap[info.Key] = info
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:405
		// _ = "end of CoverTab[127901]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:406
	// _ = "end of CoverTab[127890]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:406
	_go_fuzz_dep_.CoverTab[127891]++

									sinfo = &structInfo{
		FieldsMap:	fieldsMap,
		FieldsList:	fieldsList,
		InlineMap:	inlineMap,
	}

									fieldMapMutex.Lock()
									structMap[st] = sinfo
									fieldMapMutex.Unlock()
									return sinfo, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:417
	// _ = "end of CoverTab[127891]"
}

// IsZeroer is used to check whether an object is zero to
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:420
// determine whether it should be omitted when marshaling
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:420
// with the omitempty flag. One notable implementation
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:420
// is time.Time.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:424
type IsZeroer interface {
	IsZero() bool
}

func isZero(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:428
	_go_fuzz_dep_.CoverTab[127944]++
									kind := v.Kind()
									if z, ok := v.Interface().(IsZeroer); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:430
		_go_fuzz_dep_.CoverTab[127947]++
										if (kind == reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			_go_fuzz_dep_.CoverTab[127949]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			return kind == reflect.Interface
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			// _ = "end of CoverTab[127949]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			_go_fuzz_dep_.CoverTab[127950]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			// _ = "end of CoverTab[127950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:431
			_go_fuzz_dep_.CoverTab[127951]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:432
			// _ = "end of CoverTab[127951]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:433
			_go_fuzz_dep_.CoverTab[127952]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:433
			// _ = "end of CoverTab[127952]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:433
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:433
		// _ = "end of CoverTab[127947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:433
		_go_fuzz_dep_.CoverTab[127948]++
										return z.IsZero()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:434
		// _ = "end of CoverTab[127948]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:435
		_go_fuzz_dep_.CoverTab[127953]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:435
		// _ = "end of CoverTab[127953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:435
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:435
	// _ = "end of CoverTab[127944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:435
	_go_fuzz_dep_.CoverTab[127945]++
									switch kind {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:437
		_go_fuzz_dep_.CoverTab[127954]++
										return len(v.String()) == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:438
		// _ = "end of CoverTab[127954]"
	case reflect.Interface, reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:439
		_go_fuzz_dep_.CoverTab[127955]++
										return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:440
		// _ = "end of CoverTab[127955]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:441
		_go_fuzz_dep_.CoverTab[127956]++
										return v.Len() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:442
		// _ = "end of CoverTab[127956]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:443
		_go_fuzz_dep_.CoverTab[127957]++
										return v.Len() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:444
		// _ = "end of CoverTab[127957]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:445
		_go_fuzz_dep_.CoverTab[127958]++
										return v.Int() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:446
		// _ = "end of CoverTab[127958]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:447
		_go_fuzz_dep_.CoverTab[127959]++
										return v.Float() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:448
		// _ = "end of CoverTab[127959]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:449
		_go_fuzz_dep_.CoverTab[127960]++
										return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:450
		// _ = "end of CoverTab[127960]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:451
		_go_fuzz_dep_.CoverTab[127961]++
										return !v.Bool()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:452
		// _ = "end of CoverTab[127961]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:453
		_go_fuzz_dep_.CoverTab[127962]++
										vt := v.Type()
										for i := v.NumField() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:455
			_go_fuzz_dep_.CoverTab[127965]++
											if vt.Field(i).PkgPath != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:456
				_go_fuzz_dep_.CoverTab[127967]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:457
				// _ = "end of CoverTab[127967]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:458
				_go_fuzz_dep_.CoverTab[127968]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:458
				// _ = "end of CoverTab[127968]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:458
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:458
			// _ = "end of CoverTab[127965]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:458
			_go_fuzz_dep_.CoverTab[127966]++
											if !isZero(v.Field(i)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:459
				_go_fuzz_dep_.CoverTab[127969]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:460
				// _ = "end of CoverTab[127969]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:461
				_go_fuzz_dep_.CoverTab[127970]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:461
				// _ = "end of CoverTab[127970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:461
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:461
			// _ = "end of CoverTab[127966]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:462
		// _ = "end of CoverTab[127962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:462
		_go_fuzz_dep_.CoverTab[127963]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:463
		// _ = "end of CoverTab[127963]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:463
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:463
		_go_fuzz_dep_.CoverTab[127964]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:463
		// _ = "end of CoverTab[127964]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:464
	// _ = "end of CoverTab[127945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:464
	_go_fuzz_dep_.CoverTab[127946]++
									return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:465
	// _ = "end of CoverTab[127946]"
}

// FutureLineWrap globally disables line wrapping when encoding long strings.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// This is a temporary and thus deprecated method introduced to faciliate
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// migration towards v3, which offers more control of line lengths on
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// individual encodings, and has a default matching the behavior introduced
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// by this function.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// The default formatting of v2 was erroneously changed in v2.3.0 and reverted
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:468
// in v2.4.0, at which point this function was introduced to help migration.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:476
func FutureLineWrap() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:476
	_go_fuzz_dep_.CoverTab[127971]++
									disableLineWrapping = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:477
	// _ = "end of CoverTab[127971]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:478
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yaml.go:478
var _ = _go_fuzz_dep_.CoverTab
