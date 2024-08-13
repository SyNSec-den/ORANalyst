//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1
)

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	tagFieldName	= "toml"
	tagFieldComment	= "comment"
	tagCommented	= "commented"
	tagMultiline	= "multiline"
	tagLiteral	= "literal"
	tagDefault	= "default"
)

type tomlOpts struct {
	name		string
	nameFromTag	bool
	comment		string
	commented	bool
	multiline	bool
	literal		bool
	include		bool
	omitempty	bool
	defaultValue	string
}

type encOpts struct {
	quoteMapKeys		bool
	arraysOneElementPerLine	bool
}

var encOptsDefaults = encOpts{
	quoteMapKeys: false,
}

type annotation struct {
	tag		string
	comment		string
	commented	string
	multiline	string
	literal		string
	defaultValue	string
}

var annotationDefault = annotation{
	tag:		tagFieldName,
	comment:	tagFieldComment,
	commented:	tagCommented,
	multiline:	tagMultiline,
	literal:	tagLiteral,
	defaultValue:	tagDefault,
}

type MarshalOrder int

// Orders the Encoder can write the fields to the output stream.
const (
	// Sort fields alphabetically.
	OrderAlphabetical	MarshalOrder	= iota + 1
	// Preserve the order the fields are encountered. For example, the order of fields in
	// a struct.
	OrderPreserve
)

var timeType = reflect.TypeOf(time.Time{})
var marshalerType = reflect.TypeOf(new(Marshaler)).Elem()
var unmarshalerType = reflect.TypeOf(new(Unmarshaler)).Elem()
var textMarshalerType = reflect.TypeOf(new(encoding.TextMarshaler)).Elem()
var textUnmarshalerType = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
var localDateType = reflect.TypeOf(LocalDate{})
var localTimeType = reflect.TypeOf(LocalTime{})
var localDateTimeType = reflect.TypeOf(LocalDateTime{})
var mapStringInterfaceType = reflect.TypeOf(map[string]interface{}{})

// Check if the given marshal type maps to a Tree primitive
func isPrimitive(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:86
	_go_fuzz_dep_.CoverTab[123090]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:88
		_go_fuzz_dep_.CoverTab[123091]++
												return isPrimitive(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:89
		// _ = "end of CoverTab[123091]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:90
		_go_fuzz_dep_.CoverTab[123092]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:91
		// _ = "end of CoverTab[123092]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:92
		_go_fuzz_dep_.CoverTab[123093]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:93
		// _ = "end of CoverTab[123093]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:94
		_go_fuzz_dep_.CoverTab[123094]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:95
		// _ = "end of CoverTab[123094]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:96
		_go_fuzz_dep_.CoverTab[123095]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:97
		// _ = "end of CoverTab[123095]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:98
		_go_fuzz_dep_.CoverTab[123096]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:99
		// _ = "end of CoverTab[123096]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:100
		_go_fuzz_dep_.CoverTab[123097]++
												return isTimeType(mtype)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:101
		// _ = "end of CoverTab[123097]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:102
		_go_fuzz_dep_.CoverTab[123098]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:103
		// _ = "end of CoverTab[123098]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:104
	// _ = "end of CoverTab[123090]"
}

func isTimeType(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:107
	_go_fuzz_dep_.CoverTab[123099]++
											return mtype == timeType || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		_go_fuzz_dep_.CoverTab[123100]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		return mtype == localDateType
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		// _ = "end of CoverTab[123100]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		_go_fuzz_dep_.CoverTab[123101]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		return mtype == localDateTimeType
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		// _ = "end of CoverTab[123101]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		_go_fuzz_dep_.CoverTab[123102]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		return mtype == localTimeType
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
		// _ = "end of CoverTab[123102]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:108
	// _ = "end of CoverTab[123099]"
}

// Check if the given marshal type maps to a Tree slice or array
func isTreeSequence(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:112
	_go_fuzz_dep_.CoverTab[123103]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:114
		_go_fuzz_dep_.CoverTab[123104]++
												return isTreeSequence(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:115
		// _ = "end of CoverTab[123104]"
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:116
		_go_fuzz_dep_.CoverTab[123105]++
												return isTree(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:117
		// _ = "end of CoverTab[123105]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:118
		_go_fuzz_dep_.CoverTab[123106]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:119
		// _ = "end of CoverTab[123106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:120
	// _ = "end of CoverTab[123103]"
}

// Check if the given marshal type maps to a slice or array of a custom marshaler type
func isCustomMarshalerSequence(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:124
	_go_fuzz_dep_.CoverTab[123107]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:126
		_go_fuzz_dep_.CoverTab[123108]++
												return isCustomMarshalerSequence(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:127
		// _ = "end of CoverTab[123108]"
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:128
		_go_fuzz_dep_.CoverTab[123109]++
												return isCustomMarshaler(mtype.Elem()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:129
			_go_fuzz_dep_.CoverTab[123111]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:129
			return isCustomMarshaler(reflect.New(mtype.Elem()).Type())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:129
			// _ = "end of CoverTab[123111]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:129
		}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:129
		// _ = "end of CoverTab[123109]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:130
		_go_fuzz_dep_.CoverTab[123110]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:131
		// _ = "end of CoverTab[123110]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:132
	// _ = "end of CoverTab[123107]"
}

// Check if the given marshal type maps to a slice or array of a text marshaler type
func isTextMarshalerSequence(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:136
	_go_fuzz_dep_.CoverTab[123112]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:138
		_go_fuzz_dep_.CoverTab[123113]++
												return isTextMarshalerSequence(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:139
		// _ = "end of CoverTab[123113]"
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:140
		_go_fuzz_dep_.CoverTab[123114]++
												return isTextMarshaler(mtype.Elem()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:141
			_go_fuzz_dep_.CoverTab[123116]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:141
			return isTextMarshaler(reflect.New(mtype.Elem()).Type())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:141
			// _ = "end of CoverTab[123116]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:141
		}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:141
		// _ = "end of CoverTab[123114]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:142
		_go_fuzz_dep_.CoverTab[123115]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:143
		// _ = "end of CoverTab[123115]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:144
	// _ = "end of CoverTab[123112]"
}

// Check if the given marshal type maps to a non-Tree slice or array
func isOtherSequence(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:148
	_go_fuzz_dep_.CoverTab[123117]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:150
		_go_fuzz_dep_.CoverTab[123118]++
												return isOtherSequence(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:151
		// _ = "end of CoverTab[123118]"
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:152
		_go_fuzz_dep_.CoverTab[123119]++
												return !isTreeSequence(mtype)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:153
		// _ = "end of CoverTab[123119]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:154
		_go_fuzz_dep_.CoverTab[123120]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:155
		// _ = "end of CoverTab[123120]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:156
	// _ = "end of CoverTab[123117]"
}

// Check if the given marshal type maps to a Tree
func isTree(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:160
	_go_fuzz_dep_.CoverTab[123121]++
											switch mtype.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:162
		_go_fuzz_dep_.CoverTab[123122]++
												return isTree(mtype.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:163
		// _ = "end of CoverTab[123122]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:164
		_go_fuzz_dep_.CoverTab[123123]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:165
		// _ = "end of CoverTab[123123]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:166
		_go_fuzz_dep_.CoverTab[123124]++
												return !isPrimitive(mtype)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:167
		// _ = "end of CoverTab[123124]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:168
		_go_fuzz_dep_.CoverTab[123125]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:169
		// _ = "end of CoverTab[123125]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:170
	// _ = "end of CoverTab[123121]"
}

func isCustomMarshaler(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:173
	_go_fuzz_dep_.CoverTab[123126]++
											return mtype.Implements(marshalerType)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:174
	// _ = "end of CoverTab[123126]"
}

func callCustomMarshaler(mval reflect.Value) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:177
	_go_fuzz_dep_.CoverTab[123127]++
											return mval.Interface().(Marshaler).MarshalTOML()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:178
	// _ = "end of CoverTab[123127]"
}

func isTextMarshaler(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:181
	_go_fuzz_dep_.CoverTab[123128]++
											return mtype.Implements(textMarshalerType) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:182
		_go_fuzz_dep_.CoverTab[123129]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:182
		return !isTimeType(mtype)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:182
		// _ = "end of CoverTab[123129]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:182
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:182
	// _ = "end of CoverTab[123128]"
}

func callTextMarshaler(mval reflect.Value) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:185
	_go_fuzz_dep_.CoverTab[123130]++
											return mval.Interface().(encoding.TextMarshaler).MarshalText()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:186
	// _ = "end of CoverTab[123130]"
}

func isCustomUnmarshaler(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:189
	_go_fuzz_dep_.CoverTab[123131]++
											return mtype.Implements(unmarshalerType)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:190
	// _ = "end of CoverTab[123131]"
}

func callCustomUnmarshaler(mval reflect.Value, tval interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:193
	_go_fuzz_dep_.CoverTab[123132]++
											return mval.Interface().(Unmarshaler).UnmarshalTOML(tval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:194
	// _ = "end of CoverTab[123132]"
}

func isTextUnmarshaler(mtype reflect.Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:197
	_go_fuzz_dep_.CoverTab[123133]++
											return mtype.Implements(textUnmarshalerType)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:198
	// _ = "end of CoverTab[123133]"
}

func callTextUnmarshaler(mval reflect.Value, text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:201
	_go_fuzz_dep_.CoverTab[123134]++
											return mval.Interface().(encoding.TextUnmarshaler).UnmarshalText(text)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:202
	// _ = "end of CoverTab[123134]"
}

// Marshaler is the interface implemented by types that
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:205
// can marshal themselves into valid TOML.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:207
type Marshaler interface {
	MarshalTOML() ([]byte, error)
}

// Unmarshaler is the interface implemented by types that
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:211
// can unmarshal a TOML description of themselves.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:213
type Unmarshaler interface {
	UnmarshalTOML(interface{}) error
}

/*
Marshal returns the TOML encoding of v.  Behavior is similar to the Go json
encoder, except that there is no concept of a Marshaler interface or MarshalTOML
function for sub-structs, and currently only definite types can be marshaled
(i.e. no `interface{}`).

The following struct annotations are supported:

	toml:"Field"      Overrides the field's name to output.
	omitempty         When set, empty values and groups are not emitted.
	comment:"comment" Emits a # comment on the same line. This supports new lines.
	commented:"true"  Emits the value as commented.

Note that pointers are automatically assigned the "omitempty" option, as TOML
explicitly does not handle null values (saying instead the label should be
dropped).

Tree structural types and corresponding marshal types:

	*Tree                            (*)struct, (*)map[string]interface{}
	[]*Tree                          (*)[](*)struct, (*)[](*)map[string]interface{}
	[]interface{} (as interface{})   (*)[]primitive, (*)[]([]interface{})
	interface{}                      (*)primitive

Tree primitive types and corresponding marshal types:

	uint64     uint, uint8-uint64, pointers to same
	int64      int, int8-uint64, pointers to same
	float64    float32, float64, pointers to same
	string     string, pointers to same
	bool       bool, pointers to same
	time.LocalTime  time.LocalTime{}, pointers to same

For additional flexibility, use the Encoder API.
*/
func Marshal(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:252
	_go_fuzz_dep_.CoverTab[123135]++
											return NewEncoder(nil).marshal(v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:253
	// _ = "end of CoverTab[123135]"
}

// Encoder writes TOML values to an output stream.
type Encoder struct {
	w	io.Writer
	encOpts
	annotation
	line		int
	col		int
	order		MarshalOrder
	promoteAnon	bool
	compactComments	bool
	indentation	string
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:270
	_go_fuzz_dep_.CoverTab[123136]++
											return &Encoder{
		w:		w,
		encOpts:	encOptsDefaults,
		annotation:	annotationDefault,
		line:		0,
		col:		1,
		order:		OrderAlphabetical,
		indentation:	"  ",
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:279
	// _ = "end of CoverTab[123136]"
}

// Encode writes the TOML encoding of v to the stream.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:282
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:282
// See the documentation for Marshal for details.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:285
func (e *Encoder) Encode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:285
	_go_fuzz_dep_.CoverTab[123137]++
											b, err := e.marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:287
		_go_fuzz_dep_.CoverTab[123140]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:288
		// _ = "end of CoverTab[123140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:289
		_go_fuzz_dep_.CoverTab[123141]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:289
		// _ = "end of CoverTab[123141]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:289
	// _ = "end of CoverTab[123137]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:289
	_go_fuzz_dep_.CoverTab[123138]++
											if _, err := e.w.Write(b); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:290
		_go_fuzz_dep_.CoverTab[123142]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:291
		// _ = "end of CoverTab[123142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:292
		_go_fuzz_dep_.CoverTab[123143]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:292
		// _ = "end of CoverTab[123143]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:292
	// _ = "end of CoverTab[123138]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:292
	_go_fuzz_dep_.CoverTab[123139]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:293
	// _ = "end of CoverTab[123139]"
}

// QuoteMapKeys sets up the encoder to encode
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:296
// maps with string type keys with quoted TOML keys.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:296
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:296
// This relieves the character limitations on map keys.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:300
func (e *Encoder) QuoteMapKeys(v bool) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:300
	_go_fuzz_dep_.CoverTab[123144]++
											e.quoteMapKeys = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:302
	// _ = "end of CoverTab[123144]"
}

// ArraysWithOneElementPerLine sets up the encoder to encode arrays
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
// with more than one element on multiple lines instead of one.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
// For example:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	A = [1,2,3]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
// Becomes
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	A = [
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	  1,
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	  2,
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	  3,
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:305
//	]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:319
func (e *Encoder) ArraysWithOneElementPerLine(v bool) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:319
	_go_fuzz_dep_.CoverTab[123145]++
											e.arraysOneElementPerLine = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:321
	// _ = "end of CoverTab[123145]"
}

// Order allows to change in which order fields will be written to the output stream.
func (e *Encoder) Order(ord MarshalOrder) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:325
	_go_fuzz_dep_.CoverTab[123146]++
											e.order = ord
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:327
	// _ = "end of CoverTab[123146]"
}

// Indentation allows to change indentation when marshalling.
func (e *Encoder) Indentation(indent string) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:331
	_go_fuzz_dep_.CoverTab[123147]++
											e.indentation = indent
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:333
	// _ = "end of CoverTab[123147]"
}

// SetTagName allows changing default tag "toml"
func (e *Encoder) SetTagName(v string) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:337
	_go_fuzz_dep_.CoverTab[123148]++
											e.tag = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:339
	// _ = "end of CoverTab[123148]"
}

// SetTagComment allows changing default tag "comment"
func (e *Encoder) SetTagComment(v string) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:343
	_go_fuzz_dep_.CoverTab[123149]++
											e.comment = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:345
	// _ = "end of CoverTab[123149]"
}

// SetTagCommented allows changing default tag "commented"
func (e *Encoder) SetTagCommented(v string) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:349
	_go_fuzz_dep_.CoverTab[123150]++
											e.commented = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:351
	// _ = "end of CoverTab[123150]"
}

// SetTagMultiline allows changing default tag "multiline"
func (e *Encoder) SetTagMultiline(v string) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:355
	_go_fuzz_dep_.CoverTab[123151]++
											e.multiline = v
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:357
	// _ = "end of CoverTab[123151]"
}

// PromoteAnonymous allows to change how anonymous struct fields are marshaled.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// Usually, they are marshaled as if the inner exported fields were fields in
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// the outer struct. However, if an anonymous struct field is given a name in
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// its TOML tag, it is treated like a regular struct field with that name.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// rather than being anonymous.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// In case anonymous promotion is enabled, all anonymous structs are promoted
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:360
// and treated like regular struct fields.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:368
func (e *Encoder) PromoteAnonymous(promote bool) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:368
	_go_fuzz_dep_.CoverTab[123152]++
											e.promoteAnon = promote
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:370
	// _ = "end of CoverTab[123152]"
}

// CompactComments removes the new line before each comment in the tree.
func (e *Encoder) CompactComments(cc bool) *Encoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:374
	_go_fuzz_dep_.CoverTab[123153]++
											e.compactComments = cc
											return e
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:376
	// _ = "end of CoverTab[123153]"
}

func (e *Encoder) marshal(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:379
	_go_fuzz_dep_.CoverTab[123154]++

											for _, char := range e.indentation {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:381
		_go_fuzz_dep_.CoverTab[123161]++
												if !isSpace(char) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:382
			_go_fuzz_dep_.CoverTab[123162]++
													return []byte{}, fmt.Errorf("invalid indentation: must only contains space or tab characters")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:383
			// _ = "end of CoverTab[123162]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:384
			_go_fuzz_dep_.CoverTab[123163]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:384
			// _ = "end of CoverTab[123163]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:384
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:384
		// _ = "end of CoverTab[123161]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:385
	// _ = "end of CoverTab[123154]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:385
	_go_fuzz_dep_.CoverTab[123155]++

											mtype := reflect.TypeOf(v)
											if mtype == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:388
		_go_fuzz_dep_.CoverTab[123164]++
												return []byte{}, errors.New("nil cannot be marshaled to TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:389
		// _ = "end of CoverTab[123164]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:390
		_go_fuzz_dep_.CoverTab[123165]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:390
		// _ = "end of CoverTab[123165]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:390
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:390
	// _ = "end of CoverTab[123155]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:390
	_go_fuzz_dep_.CoverTab[123156]++

											switch mtype.Kind() {
	case reflect.Struct, reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:393
		_go_fuzz_dep_.CoverTab[123166]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:393
		// _ = "end of CoverTab[123166]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:394
		_go_fuzz_dep_.CoverTab[123167]++
												if mtype.Elem().Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:395
			_go_fuzz_dep_.CoverTab[123170]++
													return []byte{}, errors.New("Only pointer to struct can be marshaled to TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:396
			// _ = "end of CoverTab[123170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:397
			_go_fuzz_dep_.CoverTab[123171]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:397
			// _ = "end of CoverTab[123171]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:397
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:397
		// _ = "end of CoverTab[123167]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:397
		_go_fuzz_dep_.CoverTab[123168]++
												if reflect.ValueOf(v).IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:398
			_go_fuzz_dep_.CoverTab[123172]++
													return []byte{}, errors.New("nil pointer cannot be marshaled to TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:399
			// _ = "end of CoverTab[123172]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:400
			_go_fuzz_dep_.CoverTab[123173]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:400
			// _ = "end of CoverTab[123173]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:400
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:400
		// _ = "end of CoverTab[123168]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:401
		_go_fuzz_dep_.CoverTab[123169]++
												return []byte{}, errors.New("Only a struct or map can be marshaled to TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:402
		// _ = "end of CoverTab[123169]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:403
	// _ = "end of CoverTab[123156]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:403
	_go_fuzz_dep_.CoverTab[123157]++

											sval := reflect.ValueOf(v)
											if isCustomMarshaler(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:406
		_go_fuzz_dep_.CoverTab[123174]++
												return callCustomMarshaler(sval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:407
		// _ = "end of CoverTab[123174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:408
		_go_fuzz_dep_.CoverTab[123175]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:408
		// _ = "end of CoverTab[123175]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:408
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:408
	// _ = "end of CoverTab[123157]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:408
	_go_fuzz_dep_.CoverTab[123158]++
											if isTextMarshaler(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:409
		_go_fuzz_dep_.CoverTab[123176]++
												return callTextMarshaler(sval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:410
		// _ = "end of CoverTab[123176]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:411
		_go_fuzz_dep_.CoverTab[123177]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:411
		// _ = "end of CoverTab[123177]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:411
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:411
	// _ = "end of CoverTab[123158]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:411
	_go_fuzz_dep_.CoverTab[123159]++
											t, err := e.valueToTree(mtype, sval)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:413
		_go_fuzz_dep_.CoverTab[123178]++
												return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:414
		// _ = "end of CoverTab[123178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:415
		_go_fuzz_dep_.CoverTab[123179]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:415
		// _ = "end of CoverTab[123179]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:415
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:415
	// _ = "end of CoverTab[123159]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:415
	_go_fuzz_dep_.CoverTab[123160]++

											var buf bytes.Buffer
											_, err = t.writeToOrdered(&buf, "", "", 0, e.arraysOneElementPerLine, e.order, e.indentation, e.compactComments, false)

											return buf.Bytes(), err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:420
	// _ = "end of CoverTab[123160]"
}

// Create next tree with a position based on Encoder.line
func (e *Encoder) nextTree() *Tree {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:424
	_go_fuzz_dep_.CoverTab[123180]++
											return newTreeWithPosition(Position{Line: e.line, Col: 1})
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:425
	// _ = "end of CoverTab[123180]"
}

// Convert given marshal struct or map value to toml tree
func (e *Encoder) valueToTree(mtype reflect.Type, mval reflect.Value) (*Tree, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:429
	_go_fuzz_dep_.CoverTab[123181]++
											if mtype.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:430
		_go_fuzz_dep_.CoverTab[123184]++
												return e.valueToTree(mtype.Elem(), mval.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:431
		// _ = "end of CoverTab[123184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:432
		_go_fuzz_dep_.CoverTab[123185]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:432
		// _ = "end of CoverTab[123185]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:432
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:432
	// _ = "end of CoverTab[123181]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:432
	_go_fuzz_dep_.CoverTab[123182]++
											tval := e.nextTree()
											switch mtype.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:435
		_go_fuzz_dep_.CoverTab[123186]++
												switch mval.Interface().(type) {
		case Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:437
			_go_fuzz_dep_.CoverTab[123190]++
													reflect.ValueOf(tval).Elem().Set(mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:438
			// _ = "end of CoverTab[123190]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:439
			_go_fuzz_dep_.CoverTab[123191]++
													for i := 0; i < mtype.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:440
				_go_fuzz_dep_.CoverTab[123192]++
														mtypef, mvalf := mtype.Field(i), mval.Field(i)
														opts := tomlOptions(mtypef, e.annotation)
														if opts.include && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					_go_fuzz_dep_.CoverTab[123193]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					return ((mtypef.Type.Kind() != reflect.Interface && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						_go_fuzz_dep_.CoverTab[123194]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						return !opts.omitempty
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						// _ = "end of CoverTab[123194]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						_go_fuzz_dep_.CoverTab[123195]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						return !isZero(mvalf)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
						// _ = "end of CoverTab[123195]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					// _ = "end of CoverTab[123193]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:443
					_go_fuzz_dep_.CoverTab[123196]++
															val, err := e.valueToToml(mtypef.Type, mvalf)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:445
						_go_fuzz_dep_.CoverTab[123198]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:446
						// _ = "end of CoverTab[123198]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:447
						_go_fuzz_dep_.CoverTab[123199]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:447
						// _ = "end of CoverTab[123199]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:447
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:447
					// _ = "end of CoverTab[123196]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:447
					_go_fuzz_dep_.CoverTab[123197]++
															if tree, ok := val.(*Tree); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						_go_fuzz_dep_.CoverTab[123200]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						return mtypef.Anonymous
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						// _ = "end of CoverTab[123200]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						_go_fuzz_dep_.CoverTab[123201]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						return !opts.nameFromTag
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						// _ = "end of CoverTab[123201]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
					}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						_go_fuzz_dep_.CoverTab[123202]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						return !e.promoteAnon
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						// _ = "end of CoverTab[123202]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
					}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:448
						_go_fuzz_dep_.CoverTab[123203]++
																e.appendTree(tval, tree)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:449
						// _ = "end of CoverTab[123203]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:450
						_go_fuzz_dep_.CoverTab[123204]++
																val = e.wrapTomlValue(val, tval)
																tval.SetPathWithOptions([]string{opts.name}, SetOptions{
							Comment:	opts.comment,
							Commented:	opts.commented,
							Multiline:	opts.multiline,
							Literal:	opts.literal,
						}, val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:457
						// _ = "end of CoverTab[123204]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:458
					// _ = "end of CoverTab[123197]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:459
					_go_fuzz_dep_.CoverTab[123205]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:459
					// _ = "end of CoverTab[123205]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:459
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:459
				// _ = "end of CoverTab[123192]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:460
			// _ = "end of CoverTab[123191]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:461
		// _ = "end of CoverTab[123186]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:462
		_go_fuzz_dep_.CoverTab[123187]++
												keys := mval.MapKeys()
												if e.order == OrderPreserve && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:464
			_go_fuzz_dep_.CoverTab[123206]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:464
			return len(keys) > 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:464
			// _ = "end of CoverTab[123206]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:464
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:464
			_go_fuzz_dep_.CoverTab[123207]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:469
			typ := keys[0].Type()
			kind := keys[0].Kind()
			if kind == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:471
				_go_fuzz_dep_.CoverTab[123208]++
														ikeys := make([]string, len(keys))
														for i := range keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:473
					_go_fuzz_dep_.CoverTab[123210]++
															ikeys[i] = keys[i].Interface().(string)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:474
					// _ = "end of CoverTab[123210]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:475
				// _ = "end of CoverTab[123208]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:475
				_go_fuzz_dep_.CoverTab[123209]++
														sort.Strings(ikeys)
														for i := range ikeys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:477
					_go_fuzz_dep_.CoverTab[123211]++
															keys[i] = reflect.ValueOf(ikeys[i]).Convert(typ)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:478
					// _ = "end of CoverTab[123211]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:479
				// _ = "end of CoverTab[123209]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:480
				_go_fuzz_dep_.CoverTab[123212]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:480
				// _ = "end of CoverTab[123212]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:480
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:480
			// _ = "end of CoverTab[123207]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:481
			_go_fuzz_dep_.CoverTab[123213]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:481
			// _ = "end of CoverTab[123213]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:481
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:481
		// _ = "end of CoverTab[123187]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:481
		_go_fuzz_dep_.CoverTab[123188]++
												for _, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:482
			_go_fuzz_dep_.CoverTab[123214]++
													mvalf := mval.MapIndex(key)
													if (mtype.Elem().Kind() == reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				_go_fuzz_dep_.CoverTab[123217]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				return mtype.Elem().Kind() == reflect.Interface
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				// _ = "end of CoverTab[123217]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				_go_fuzz_dep_.CoverTab[123218]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				return mvalf.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				// _ = "end of CoverTab[123218]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:484
				_go_fuzz_dep_.CoverTab[123219]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:485
				// _ = "end of CoverTab[123219]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:486
				_go_fuzz_dep_.CoverTab[123220]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:486
				// _ = "end of CoverTab[123220]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:486
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:486
			// _ = "end of CoverTab[123214]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:486
			_go_fuzz_dep_.CoverTab[123215]++
													val, err := e.valueToToml(mtype.Elem(), mvalf)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:488
				_go_fuzz_dep_.CoverTab[123221]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:489
				// _ = "end of CoverTab[123221]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:490
				_go_fuzz_dep_.CoverTab[123222]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:490
				// _ = "end of CoverTab[123222]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:490
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:490
			// _ = "end of CoverTab[123215]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:490
			_go_fuzz_dep_.CoverTab[123216]++
													val = e.wrapTomlValue(val, tval)
													if e.quoteMapKeys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:492
				_go_fuzz_dep_.CoverTab[123223]++
														keyStr, err := tomlValueStringRepresentation(key.String(), "", "", e.order, e.arraysOneElementPerLine)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:494
					_go_fuzz_dep_.CoverTab[123225]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:495
					// _ = "end of CoverTab[123225]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:496
					_go_fuzz_dep_.CoverTab[123226]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:496
					// _ = "end of CoverTab[123226]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:496
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:496
				// _ = "end of CoverTab[123223]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:496
				_go_fuzz_dep_.CoverTab[123224]++
														tval.SetPath([]string{keyStr}, val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:497
				// _ = "end of CoverTab[123224]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:498
				_go_fuzz_dep_.CoverTab[123227]++
														tval.SetPath([]string{key.String()}, val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:499
				// _ = "end of CoverTab[123227]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:500
			// _ = "end of CoverTab[123216]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:501
		// _ = "end of CoverTab[123188]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:501
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:501
		_go_fuzz_dep_.CoverTab[123189]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:501
		// _ = "end of CoverTab[123189]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:502
	// _ = "end of CoverTab[123182]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:502
	_go_fuzz_dep_.CoverTab[123183]++
											return tval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:503
	// _ = "end of CoverTab[123183]"
}

// Convert given marshal slice to slice of Toml trees
func (e *Encoder) valueToTreeSlice(mtype reflect.Type, mval reflect.Value) ([]*Tree, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:507
	_go_fuzz_dep_.CoverTab[123228]++
											tval := make([]*Tree, mval.Len(), mval.Len())
											for i := 0; i < mval.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:509
		_go_fuzz_dep_.CoverTab[123230]++
												val, err := e.valueToTree(mtype.Elem(), mval.Index(i))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:511
			_go_fuzz_dep_.CoverTab[123232]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:512
			// _ = "end of CoverTab[123232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:513
			_go_fuzz_dep_.CoverTab[123233]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:513
			// _ = "end of CoverTab[123233]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:513
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:513
		// _ = "end of CoverTab[123230]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:513
		_go_fuzz_dep_.CoverTab[123231]++
												tval[i] = val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:514
		// _ = "end of CoverTab[123231]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:515
	// _ = "end of CoverTab[123228]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:515
	_go_fuzz_dep_.CoverTab[123229]++
											return tval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:516
	// _ = "end of CoverTab[123229]"
}

// Convert given marshal slice to slice of toml values
func (e *Encoder) valueToOtherSlice(mtype reflect.Type, mval reflect.Value) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:520
	_go_fuzz_dep_.CoverTab[123234]++
											tval := make([]interface{}, mval.Len(), mval.Len())
											for i := 0; i < mval.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:522
		_go_fuzz_dep_.CoverTab[123236]++
												val, err := e.valueToToml(mtype.Elem(), mval.Index(i))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:524
			_go_fuzz_dep_.CoverTab[123238]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:525
			// _ = "end of CoverTab[123238]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:526
			_go_fuzz_dep_.CoverTab[123239]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:526
			// _ = "end of CoverTab[123239]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:526
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:526
		// _ = "end of CoverTab[123236]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:526
		_go_fuzz_dep_.CoverTab[123237]++
												tval[i] = val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:527
		// _ = "end of CoverTab[123237]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:528
	// _ = "end of CoverTab[123234]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:528
	_go_fuzz_dep_.CoverTab[123235]++
											return tval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:529
	// _ = "end of CoverTab[123235]"
}

// Convert given marshal value to toml value
func (e *Encoder) valueToToml(mtype reflect.Type, mval reflect.Value) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:533
	_go_fuzz_dep_.CoverTab[123240]++
											if mtype.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:534
		_go_fuzz_dep_.CoverTab[123243]++
												switch {
		case isCustomMarshaler(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:536
			_go_fuzz_dep_.CoverTab[123244]++
													return callCustomMarshaler(mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:537
			// _ = "end of CoverTab[123244]"
		case isTextMarshaler(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:538
			_go_fuzz_dep_.CoverTab[123245]++
													b, err := callTextMarshaler(mval)
													return string(b), err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:540
			// _ = "end of CoverTab[123245]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:541
			_go_fuzz_dep_.CoverTab[123246]++
													return e.valueToToml(mtype.Elem(), mval.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:542
			// _ = "end of CoverTab[123246]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:543
		// _ = "end of CoverTab[123243]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:544
		_go_fuzz_dep_.CoverTab[123247]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:544
		// _ = "end of CoverTab[123247]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:544
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:544
	// _ = "end of CoverTab[123240]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:544
	_go_fuzz_dep_.CoverTab[123241]++
											if mtype.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:545
		_go_fuzz_dep_.CoverTab[123248]++
												return e.valueToToml(mval.Elem().Type(), mval.Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:546
		// _ = "end of CoverTab[123248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:547
		_go_fuzz_dep_.CoverTab[123249]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:547
		// _ = "end of CoverTab[123249]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:547
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:547
	// _ = "end of CoverTab[123241]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:547
	_go_fuzz_dep_.CoverTab[123242]++
											switch {
	case isCustomMarshaler(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:549
		_go_fuzz_dep_.CoverTab[123250]++
												return callCustomMarshaler(mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:550
		// _ = "end of CoverTab[123250]"
	case isTextMarshaler(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:551
		_go_fuzz_dep_.CoverTab[123251]++
												b, err := callTextMarshaler(mval)
												return string(b), err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:553
		// _ = "end of CoverTab[123251]"
	case isTree(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:554
		_go_fuzz_dep_.CoverTab[123252]++
												return e.valueToTree(mtype, mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:555
		// _ = "end of CoverTab[123252]"
	case isOtherSequence(mtype), isCustomMarshalerSequence(mtype), isTextMarshalerSequence(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:556
		_go_fuzz_dep_.CoverTab[123253]++
												return e.valueToOtherSlice(mtype, mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:557
		// _ = "end of CoverTab[123253]"
	case isTreeSequence(mtype):
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:558
		_go_fuzz_dep_.CoverTab[123254]++
												return e.valueToTreeSlice(mtype, mval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:559
		// _ = "end of CoverTab[123254]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:560
		_go_fuzz_dep_.CoverTab[123255]++
												switch mtype.Kind() {
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:562
			_go_fuzz_dep_.CoverTab[123256]++
													return mval.Bool(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:563
			// _ = "end of CoverTab[123256]"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:564
			_go_fuzz_dep_.CoverTab[123257]++
													if mtype.Kind() == reflect.Int64 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:565
				_go_fuzz_dep_.CoverTab[123264]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:565
				return mtype == reflect.TypeOf(time.Duration(1))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:565
				// _ = "end of CoverTab[123264]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:565
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:565
				_go_fuzz_dep_.CoverTab[123265]++
														return fmt.Sprint(mval), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:566
				// _ = "end of CoverTab[123265]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:567
				_go_fuzz_dep_.CoverTab[123266]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:567
				// _ = "end of CoverTab[123266]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:567
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:567
			// _ = "end of CoverTab[123257]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:567
			_go_fuzz_dep_.CoverTab[123258]++
													return mval.Int(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:568
			// _ = "end of CoverTab[123258]"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:569
			_go_fuzz_dep_.CoverTab[123259]++
													return mval.Uint(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:570
			// _ = "end of CoverTab[123259]"
		case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:571
			_go_fuzz_dep_.CoverTab[123260]++
													return mval.Float(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:572
			// _ = "end of CoverTab[123260]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:573
			_go_fuzz_dep_.CoverTab[123261]++
													return mval.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:574
			// _ = "end of CoverTab[123261]"
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:575
			_go_fuzz_dep_.CoverTab[123262]++
													return mval.Interface(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:576
			// _ = "end of CoverTab[123262]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:577
			_go_fuzz_dep_.CoverTab[123263]++
													return nil, fmt.Errorf("Marshal can't handle %v(%v)", mtype, mtype.Kind())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:578
			// _ = "end of CoverTab[123263]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:579
		// _ = "end of CoverTab[123255]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:580
	// _ = "end of CoverTab[123242]"
}

func (e *Encoder) appendTree(t, o *Tree) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:583
	_go_fuzz_dep_.CoverTab[123267]++
											for key, value := range o.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:584
		_go_fuzz_dep_.CoverTab[123269]++
												if _, ok := t.values[key]; ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:585
			_go_fuzz_dep_.CoverTab[123272]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:586
			// _ = "end of CoverTab[123272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:587
			_go_fuzz_dep_.CoverTab[123273]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:587
			// _ = "end of CoverTab[123273]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:587
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:587
		// _ = "end of CoverTab[123269]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:587
		_go_fuzz_dep_.CoverTab[123270]++
												if tomlValue, ok := value.(*tomlValue); ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:588
			_go_fuzz_dep_.CoverTab[123274]++
													tomlValue.position.Col = t.position.Col
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:589
			// _ = "end of CoverTab[123274]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:590
			_go_fuzz_dep_.CoverTab[123275]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:590
			// _ = "end of CoverTab[123275]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:590
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:590
		// _ = "end of CoverTab[123270]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:590
		_go_fuzz_dep_.CoverTab[123271]++
												t.values[key] = value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:591
		// _ = "end of CoverTab[123271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:592
	// _ = "end of CoverTab[123267]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:592
	_go_fuzz_dep_.CoverTab[123268]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:593
	// _ = "end of CoverTab[123268]"
}

// Create a toml value with the current line number as the position line
func (e *Encoder) wrapTomlValue(val interface{}, parent *Tree) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:597
	_go_fuzz_dep_.CoverTab[123276]++
											_, isTree := val.(*Tree)
											_, isTreeS := val.([]*Tree)
											if isTree || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:600
		_go_fuzz_dep_.CoverTab[123278]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:600
		return isTreeS
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:600
		// _ = "end of CoverTab[123278]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:600
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:600
		_go_fuzz_dep_.CoverTab[123279]++
												e.line++
												return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:602
		// _ = "end of CoverTab[123279]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:603
		_go_fuzz_dep_.CoverTab[123280]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:603
		// _ = "end of CoverTab[123280]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:603
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:603
	// _ = "end of CoverTab[123276]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:603
	_go_fuzz_dep_.CoverTab[123277]++

											ret := &tomlValue{
		value:	val,
		position: Position{
			e.line,
			parent.position.Col,
		},
	}
											e.line++
											return ret
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:613
	// _ = "end of CoverTab[123277]"
}

// Unmarshal attempts to unmarshal the Tree into a Go struct pointed by v.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:616
// Neither Unmarshaler interfaces nor UnmarshalTOML functions are supported for
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:616
// sub-structs, and only definite types can be unmarshaled.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:619
func (t *Tree) Unmarshal(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:619
	_go_fuzz_dep_.CoverTab[123281]++
											d := Decoder{tval: t, tagName: tagFieldName}
											return d.unmarshal(v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:621
	// _ = "end of CoverTab[123281]"
}

// Marshal returns the TOML encoding of Tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:624
// See Marshal() documentation for types mapping table.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:626
func (t *Tree) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:626
	_go_fuzz_dep_.CoverTab[123282]++
											var buf bytes.Buffer
											_, err := t.WriteTo(&buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:629
		_go_fuzz_dep_.CoverTab[123284]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:630
		// _ = "end of CoverTab[123284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:631
		_go_fuzz_dep_.CoverTab[123285]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:631
		// _ = "end of CoverTab[123285]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:631
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:631
	// _ = "end of CoverTab[123282]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:631
	_go_fuzz_dep_.CoverTab[123283]++
											return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:632
	// _ = "end of CoverTab[123283]"
}

// Unmarshal parses the TOML-encoded data and stores the result in the value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// pointed to by v. Behavior is similar to the Go json encoder, except that there
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// is no concept of an Unmarshaler interface or UnmarshalTOML function for
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// sub-structs, and currently only definite types can be unmarshaled to (i.e. no
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// `interface{}`).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// The following struct annotations are supported:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//	toml:"Field" Overrides the field's name to map to.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//	default:"foo" Provides a default value.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// For default values, only fields of the following types are supported:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//   - string
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//   - bool
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//   - int
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//   - int64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//   - float64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:635
// See Marshal() documentation for types mapping table.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:654
func Unmarshal(data []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:654
	_go_fuzz_dep_.CoverTab[123286]++
											t, err := LoadReader(bytes.NewReader(data))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:656
		_go_fuzz_dep_.CoverTab[123288]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:657
		// _ = "end of CoverTab[123288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:658
		_go_fuzz_dep_.CoverTab[123289]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:658
		// _ = "end of CoverTab[123289]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:658
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:658
	// _ = "end of CoverTab[123286]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:658
	_go_fuzz_dep_.CoverTab[123287]++
											return t.Unmarshal(v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:659
	// _ = "end of CoverTab[123287]"
}

// Decoder reads and decodes TOML values from an input stream.
type Decoder struct {
	r	io.Reader
	tval	*Tree
	encOpts
	tagName	string
	strict	bool
	visitor	visitorState
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:673
	_go_fuzz_dep_.CoverTab[123290]++
											return &Decoder{
		r:		r,
		encOpts:	encOptsDefaults,
		tagName:	tagFieldName,
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:678
	// _ = "end of CoverTab[123290]"
}

// Decode reads a TOML-encoded value from it's input
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:681
// and unmarshals it in the value pointed at by v.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:681
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:681
// See the documentation for Marshal for details.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:685
func (d *Decoder) Decode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:685
	_go_fuzz_dep_.CoverTab[123291]++
											var err error
											d.tval, err = LoadReader(d.r)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:688
		_go_fuzz_dep_.CoverTab[123293]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:689
		// _ = "end of CoverTab[123293]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:690
		_go_fuzz_dep_.CoverTab[123294]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:690
		// _ = "end of CoverTab[123294]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:690
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:690
	// _ = "end of CoverTab[123291]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:690
	_go_fuzz_dep_.CoverTab[123292]++
											return d.unmarshal(v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:691
	// _ = "end of CoverTab[123292]"
}

// SetTagName allows changing default tag "toml"
func (d *Decoder) SetTagName(v string) *Decoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:695
	_go_fuzz_dep_.CoverTab[123295]++
											d.tagName = v
											return d
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:697
	// _ = "end of CoverTab[123295]"
}

// Strict allows changing to strict decoding. Any fields that are found in the
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:700
// input data and do not have a corresponding struct member cause an error.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:702
func (d *Decoder) Strict(strict bool) *Decoder {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:702
	_go_fuzz_dep_.CoverTab[123296]++
											d.strict = strict
											return d
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:704
	// _ = "end of CoverTab[123296]"
}

func (d *Decoder) unmarshal(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:707
	_go_fuzz_dep_.CoverTab[123297]++
											mtype := reflect.TypeOf(v)
											if mtype == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:709
		_go_fuzz_dep_.CoverTab[123305]++
												return errors.New("nil cannot be unmarshaled from TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:710
		// _ = "end of CoverTab[123305]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:711
		_go_fuzz_dep_.CoverTab[123306]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:711
		// _ = "end of CoverTab[123306]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:711
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:711
	// _ = "end of CoverTab[123297]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:711
	_go_fuzz_dep_.CoverTab[123298]++
											if mtype.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:712
		_go_fuzz_dep_.CoverTab[123307]++
												return errors.New("only a pointer to struct or map can be unmarshaled from TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:713
		// _ = "end of CoverTab[123307]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:714
		_go_fuzz_dep_.CoverTab[123308]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:714
		// _ = "end of CoverTab[123308]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:714
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:714
	// _ = "end of CoverTab[123298]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:714
	_go_fuzz_dep_.CoverTab[123299]++

											elem := mtype.Elem()

											switch elem.Kind() {
	case reflect.Struct, reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:719
		_go_fuzz_dep_.CoverTab[123309]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:719
		// _ = "end of CoverTab[123309]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:720
		_go_fuzz_dep_.CoverTab[123310]++
												elem = mapStringInterfaceType
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:721
		// _ = "end of CoverTab[123310]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:722
		_go_fuzz_dep_.CoverTab[123311]++
												return errors.New("only a pointer to struct or map can be unmarshaled from TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:723
		// _ = "end of CoverTab[123311]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:724
	// _ = "end of CoverTab[123299]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:724
	_go_fuzz_dep_.CoverTab[123300]++

											if reflect.ValueOf(v).IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:726
		_go_fuzz_dep_.CoverTab[123312]++
												return errors.New("nil pointer cannot be unmarshaled from TOML")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:727
		// _ = "end of CoverTab[123312]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:728
		_go_fuzz_dep_.CoverTab[123313]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:728
		// _ = "end of CoverTab[123313]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:728
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:728
	// _ = "end of CoverTab[123300]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:728
	_go_fuzz_dep_.CoverTab[123301]++

											vv := reflect.ValueOf(v).Elem()

											if d.strict {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:732
		_go_fuzz_dep_.CoverTab[123314]++
												d.visitor = newVisitorState(d.tval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:733
		// _ = "end of CoverTab[123314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:734
		_go_fuzz_dep_.CoverTab[123315]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:734
		// _ = "end of CoverTab[123315]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:734
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:734
	// _ = "end of CoverTab[123301]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:734
	_go_fuzz_dep_.CoverTab[123302]++

											sval, err := d.valueFromTree(elem, d.tval, &vv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:737
		_go_fuzz_dep_.CoverTab[123316]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:738
		// _ = "end of CoverTab[123316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:739
		_go_fuzz_dep_.CoverTab[123317]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:739
		// _ = "end of CoverTab[123317]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:739
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:739
	// _ = "end of CoverTab[123302]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:739
	_go_fuzz_dep_.CoverTab[123303]++
											if err := d.visitor.validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:740
		_go_fuzz_dep_.CoverTab[123318]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:741
		// _ = "end of CoverTab[123318]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:742
		_go_fuzz_dep_.CoverTab[123319]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:742
		// _ = "end of CoverTab[123319]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:742
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:742
	// _ = "end of CoverTab[123303]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:742
	_go_fuzz_dep_.CoverTab[123304]++
											reflect.ValueOf(v).Elem().Set(sval)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:744
	// _ = "end of CoverTab[123304]"
}

// Convert toml tree to marshal struct or map, using marshal type. When mval1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:747
// is non-nil, merge fields into the given value instead of allocating a new one.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:749
func (d *Decoder) valueFromTree(mtype reflect.Type, tval *Tree, mval1 *reflect.Value) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:749
	_go_fuzz_dep_.CoverTab[123320]++
											if mtype.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:750
		_go_fuzz_dep_.CoverTab[123324]++
												return d.unwrapPointer(mtype, tval, mval1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:751
		// _ = "end of CoverTab[123324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:752
		_go_fuzz_dep_.CoverTab[123325]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:752
		// _ = "end of CoverTab[123325]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:752
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:752
	// _ = "end of CoverTab[123320]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:752
	_go_fuzz_dep_.CoverTab[123321]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:755
	if mvalPtr := reflect.New(mtype); isCustomUnmarshaler(mvalPtr.Type()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:755
		_go_fuzz_dep_.CoverTab[123326]++
												d.visitor.visitAll()

												if tval == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:758
			_go_fuzz_dep_.CoverTab[123329]++
													return mvalPtr.Elem(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:759
			// _ = "end of CoverTab[123329]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:760
			_go_fuzz_dep_.CoverTab[123330]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:760
			// _ = "end of CoverTab[123330]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:760
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:760
		// _ = "end of CoverTab[123326]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:760
		_go_fuzz_dep_.CoverTab[123327]++

												if err := callCustomUnmarshaler(mvalPtr, tval.ToMap()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:762
			_go_fuzz_dep_.CoverTab[123331]++
													return reflect.ValueOf(nil), fmt.Errorf("unmarshal toml: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:763
			// _ = "end of CoverTab[123331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:764
			_go_fuzz_dep_.CoverTab[123332]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:764
			// _ = "end of CoverTab[123332]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:764
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:764
		// _ = "end of CoverTab[123327]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:764
		_go_fuzz_dep_.CoverTab[123328]++
												return mvalPtr.Elem(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:765
		// _ = "end of CoverTab[123328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:766
		_go_fuzz_dep_.CoverTab[123333]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:766
		// _ = "end of CoverTab[123333]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:766
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:766
	// _ = "end of CoverTab[123321]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:766
	_go_fuzz_dep_.CoverTab[123322]++

											var mval reflect.Value
											switch mtype.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:770
		_go_fuzz_dep_.CoverTab[123334]++
												if mval1 != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:771
			_go_fuzz_dep_.CoverTab[123338]++
													mval = *mval1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:772
			// _ = "end of CoverTab[123338]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:773
			_go_fuzz_dep_.CoverTab[123339]++
													mval = reflect.New(mtype).Elem()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:774
			// _ = "end of CoverTab[123339]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:775
		// _ = "end of CoverTab[123334]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:775
		_go_fuzz_dep_.CoverTab[123335]++

												switch mval.Interface().(type) {
		case Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:778
			_go_fuzz_dep_.CoverTab[123340]++
													mval.Set(reflect.ValueOf(tval).Elem())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:779
			// _ = "end of CoverTab[123340]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:780
			_go_fuzz_dep_.CoverTab[123341]++
													for i := 0; i < mtype.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:781
				_go_fuzz_dep_.CoverTab[123342]++
														mtypef := mtype.Field(i)
														an := annotation{tag: d.tagName}
														opts := tomlOptions(mtypef, an)
														if !opts.include {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:785
					_go_fuzz_dep_.CoverTab[123346]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:786
					// _ = "end of CoverTab[123346]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:787
					_go_fuzz_dep_.CoverTab[123347]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:787
					// _ = "end of CoverTab[123347]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:787
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:787
				// _ = "end of CoverTab[123342]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:787
				_go_fuzz_dep_.CoverTab[123343]++
														baseKey := opts.name
														keysToTry := []string{
					baseKey,
					strings.ToLower(baseKey),
					strings.ToTitle(baseKey),
					strings.ToLower(string(baseKey[0])) + baseKey[1:],
				}

				found := false
				if tval != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:797
					_go_fuzz_dep_.CoverTab[123348]++
															for _, key := range keysToTry {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:798
						_go_fuzz_dep_.CoverTab[123349]++
																exists := tval.HasPath([]string{key})
																if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:800
							_go_fuzz_dep_.CoverTab[123352]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:801
							// _ = "end of CoverTab[123352]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:802
							_go_fuzz_dep_.CoverTab[123353]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:802
							// _ = "end of CoverTab[123353]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:802
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:802
						// _ = "end of CoverTab[123349]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:802
						_go_fuzz_dep_.CoverTab[123350]++

																d.visitor.push(key)
																val := tval.GetPath([]string{key})
																fval := mval.Field(i)
																mvalf, err := d.valueFromToml(mtypef.Type, val, &fval)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:808
							_go_fuzz_dep_.CoverTab[123354]++
																	return mval, formatError(err, tval.GetPositionPath([]string{key}))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:809
							// _ = "end of CoverTab[123354]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:810
							_go_fuzz_dep_.CoverTab[123355]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:810
							// _ = "end of CoverTab[123355]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:810
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:810
						// _ = "end of CoverTab[123350]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:810
						_go_fuzz_dep_.CoverTab[123351]++
																mval.Field(i).Set(mvalf)
																found = true
																d.visitor.pop()
																break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:814
						// _ = "end of CoverTab[123351]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:815
					// _ = "end of CoverTab[123348]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:816
					_go_fuzz_dep_.CoverTab[123356]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:816
					// _ = "end of CoverTab[123356]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:816
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:816
				// _ = "end of CoverTab[123343]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:816
				_go_fuzz_dep_.CoverTab[123344]++

														if !found && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:818
					_go_fuzz_dep_.CoverTab[123357]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:818
					return opts.defaultValue != ""
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:818
					// _ = "end of CoverTab[123357]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:818
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:818
					_go_fuzz_dep_.CoverTab[123358]++
															mvalf := mval.Field(i)
															var val interface{}
															var err error
															switch mvalf.Kind() {
					case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:823
						_go_fuzz_dep_.CoverTab[123361]++
																val = opts.defaultValue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:824
						// _ = "end of CoverTab[123361]"
					case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:825
						_go_fuzz_dep_.CoverTab[123362]++
																val, err = strconv.ParseBool(opts.defaultValue)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:826
						// _ = "end of CoverTab[123362]"
					case reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:827
						_go_fuzz_dep_.CoverTab[123363]++
																val, err = strconv.ParseUint(opts.defaultValue, 10, 0)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:828
						// _ = "end of CoverTab[123363]"
					case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:829
						_go_fuzz_dep_.CoverTab[123364]++
																val, err = strconv.ParseUint(opts.defaultValue, 10, 8)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:830
						// _ = "end of CoverTab[123364]"
					case reflect.Uint16:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:831
						_go_fuzz_dep_.CoverTab[123365]++
																val, err = strconv.ParseUint(opts.defaultValue, 10, 16)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:832
						// _ = "end of CoverTab[123365]"
					case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:833
						_go_fuzz_dep_.CoverTab[123366]++
																val, err = strconv.ParseUint(opts.defaultValue, 10, 32)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:834
						// _ = "end of CoverTab[123366]"
					case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:835
						_go_fuzz_dep_.CoverTab[123367]++
																val, err = strconv.ParseUint(opts.defaultValue, 10, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:836
						// _ = "end of CoverTab[123367]"
					case reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:837
						_go_fuzz_dep_.CoverTab[123368]++
																val, err = strconv.ParseInt(opts.defaultValue, 10, 0)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:838
						// _ = "end of CoverTab[123368]"
					case reflect.Int8:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:839
						_go_fuzz_dep_.CoverTab[123369]++
																val, err = strconv.ParseInt(opts.defaultValue, 10, 8)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:840
						// _ = "end of CoverTab[123369]"
					case reflect.Int16:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:841
						_go_fuzz_dep_.CoverTab[123370]++
																val, err = strconv.ParseInt(opts.defaultValue, 10, 16)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:842
						// _ = "end of CoverTab[123370]"
					case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:843
						_go_fuzz_dep_.CoverTab[123371]++
																val, err = strconv.ParseInt(opts.defaultValue, 10, 32)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:844
						// _ = "end of CoverTab[123371]"
					case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:845
						_go_fuzz_dep_.CoverTab[123372]++
						// Check if the provided number has a non-numeric extension.
						var hasExtension bool
						if len(opts.defaultValue) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:848
							_go_fuzz_dep_.CoverTab[123377]++
																	lastChar := opts.defaultValue[len(opts.defaultValue)-1]
																	if lastChar < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:850
								_go_fuzz_dep_.CoverTab[123378]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:850
								return lastChar > '9'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:850
								// _ = "end of CoverTab[123378]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:850
							}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:850
								_go_fuzz_dep_.CoverTab[123379]++
																		hasExtension = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:851
								// _ = "end of CoverTab[123379]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:852
								_go_fuzz_dep_.CoverTab[123380]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:852
								// _ = "end of CoverTab[123380]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:852
							}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:852
							// _ = "end of CoverTab[123377]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:853
							_go_fuzz_dep_.CoverTab[123381]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:853
							// _ = "end of CoverTab[123381]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:853
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:853
						// _ = "end of CoverTab[123372]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:853
						_go_fuzz_dep_.CoverTab[123373]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
						if hasExtension && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
							_go_fuzz_dep_.CoverTab[123382]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
							return mvalf.Type().String() == "time.Duration"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
							// _ = "end of CoverTab[123382]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
						}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:856
							_go_fuzz_dep_.CoverTab[123383]++
																	val, err = time.ParseDuration(opts.defaultValue)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:857
							// _ = "end of CoverTab[123383]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:858
							_go_fuzz_dep_.CoverTab[123384]++
																	val, err = strconv.ParseInt(opts.defaultValue, 10, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:859
							// _ = "end of CoverTab[123384]"
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:860
						// _ = "end of CoverTab[123373]"
					case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:861
						_go_fuzz_dep_.CoverTab[123374]++
																val, err = strconv.ParseFloat(opts.defaultValue, 32)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:862
						// _ = "end of CoverTab[123374]"
					case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:863
						_go_fuzz_dep_.CoverTab[123375]++
																val, err = strconv.ParseFloat(opts.defaultValue, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:864
						// _ = "end of CoverTab[123375]"
					default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:865
						_go_fuzz_dep_.CoverTab[123376]++
																return mvalf, fmt.Errorf("unsupported field type for default option")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:866
						// _ = "end of CoverTab[123376]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:867
					// _ = "end of CoverTab[123358]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:867
					_go_fuzz_dep_.CoverTab[123359]++

															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:869
						_go_fuzz_dep_.CoverTab[123385]++
																return mvalf, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:870
						// _ = "end of CoverTab[123385]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:871
						_go_fuzz_dep_.CoverTab[123386]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:871
						// _ = "end of CoverTab[123386]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:871
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:871
					// _ = "end of CoverTab[123359]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:871
					_go_fuzz_dep_.CoverTab[123360]++
															mvalf.Set(reflect.ValueOf(val).Convert(mvalf.Type()))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:872
					// _ = "end of CoverTab[123360]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:873
					_go_fuzz_dep_.CoverTab[123387]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:873
					// _ = "end of CoverTab[123387]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:873
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:873
				// _ = "end of CoverTab[123344]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:873
				_go_fuzz_dep_.CoverTab[123345]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
				if !found && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					_go_fuzz_dep_.CoverTab[123388]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					return opts.defaultValue == ""
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					// _ = "end of CoverTab[123388]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					_go_fuzz_dep_.CoverTab[123389]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					return mtypef.Type.Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					// _ = "end of CoverTab[123389]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:876
					_go_fuzz_dep_.CoverTab[123390]++
															tmpTval := tval
															if !mtypef.Anonymous {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:878
						_go_fuzz_dep_.CoverTab[123393]++
																tmpTval = nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:879
						// _ = "end of CoverTab[123393]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:880
						_go_fuzz_dep_.CoverTab[123394]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:880
						// _ = "end of CoverTab[123394]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:880
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:880
					// _ = "end of CoverTab[123390]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:880
					_go_fuzz_dep_.CoverTab[123391]++
															fval := mval.Field(i)
															v, err := d.valueFromTree(mtypef.Type, tmpTval, &fval)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:883
						_go_fuzz_dep_.CoverTab[123395]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:884
						// _ = "end of CoverTab[123395]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:885
						_go_fuzz_dep_.CoverTab[123396]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:885
						// _ = "end of CoverTab[123396]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:885
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:885
					// _ = "end of CoverTab[123391]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:885
					_go_fuzz_dep_.CoverTab[123392]++
															mval.Field(i).Set(v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:886
					// _ = "end of CoverTab[123392]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:887
					_go_fuzz_dep_.CoverTab[123397]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:887
					// _ = "end of CoverTab[123397]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:887
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:887
				// _ = "end of CoverTab[123345]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:888
			// _ = "end of CoverTab[123341]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:889
		// _ = "end of CoverTab[123335]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:890
		_go_fuzz_dep_.CoverTab[123336]++
												mval = reflect.MakeMap(mtype)
												for _, key := range tval.Keys() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:892
			_go_fuzz_dep_.CoverTab[123398]++
													d.visitor.push(key)

													val := tval.GetPath([]string{key})
													mvalf, err := d.valueFromToml(mtype.Elem(), val, nil)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:897
				_go_fuzz_dep_.CoverTab[123400]++
														return mval, formatError(err, tval.GetPositionPath([]string{key}))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:898
				// _ = "end of CoverTab[123400]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:899
				_go_fuzz_dep_.CoverTab[123401]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:899
				// _ = "end of CoverTab[123401]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:899
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:899
			// _ = "end of CoverTab[123398]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:899
			_go_fuzz_dep_.CoverTab[123399]++
													mval.SetMapIndex(reflect.ValueOf(key).Convert(mtype.Key()), mvalf)
													d.visitor.pop()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:901
			// _ = "end of CoverTab[123399]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:902
		// _ = "end of CoverTab[123336]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:902
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:902
		_go_fuzz_dep_.CoverTab[123337]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:902
		// _ = "end of CoverTab[123337]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:903
	// _ = "end of CoverTab[123322]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:903
	_go_fuzz_dep_.CoverTab[123323]++
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:904
	// _ = "end of CoverTab[123323]"
}

// Convert toml value to marshal struct/map slice, using marshal type
func (d *Decoder) valueFromTreeSlice(mtype reflect.Type, tval []*Tree) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:908
	_go_fuzz_dep_.CoverTab[123402]++
											mval, err := makeSliceOrArray(mtype, len(tval))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:910
		_go_fuzz_dep_.CoverTab[123405]++
												return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:911
		// _ = "end of CoverTab[123405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:912
		_go_fuzz_dep_.CoverTab[123406]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:912
		// _ = "end of CoverTab[123406]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:912
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:912
	// _ = "end of CoverTab[123402]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:912
	_go_fuzz_dep_.CoverTab[123403]++

											for i := 0; i < len(tval); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:914
		_go_fuzz_dep_.CoverTab[123407]++
												d.visitor.push(strconv.Itoa(i))
												val, err := d.valueFromTree(mtype.Elem(), tval[i], nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:917
			_go_fuzz_dep_.CoverTab[123409]++
													return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:918
			// _ = "end of CoverTab[123409]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:919
			_go_fuzz_dep_.CoverTab[123410]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:919
			// _ = "end of CoverTab[123410]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:919
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:919
		// _ = "end of CoverTab[123407]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:919
		_go_fuzz_dep_.CoverTab[123408]++
												mval.Index(i).Set(val)
												d.visitor.pop()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:921
		// _ = "end of CoverTab[123408]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:922
	// _ = "end of CoverTab[123403]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:922
	_go_fuzz_dep_.CoverTab[123404]++
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:923
	// _ = "end of CoverTab[123404]"
}

// Convert toml value to marshal primitive slice, using marshal type
func (d *Decoder) valueFromOtherSlice(mtype reflect.Type, tval []interface{}) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:927
	_go_fuzz_dep_.CoverTab[123411]++
											mval, err := makeSliceOrArray(mtype, len(tval))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:929
		_go_fuzz_dep_.CoverTab[123414]++
												return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:930
		// _ = "end of CoverTab[123414]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:931
		_go_fuzz_dep_.CoverTab[123415]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:931
		// _ = "end of CoverTab[123415]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:931
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:931
	// _ = "end of CoverTab[123411]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:931
	_go_fuzz_dep_.CoverTab[123412]++

											for i := 0; i < len(tval); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:933
		_go_fuzz_dep_.CoverTab[123416]++
												val, err := d.valueFromToml(mtype.Elem(), tval[i], nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:935
			_go_fuzz_dep_.CoverTab[123418]++
													return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:936
			// _ = "end of CoverTab[123418]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:937
			_go_fuzz_dep_.CoverTab[123419]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:937
			// _ = "end of CoverTab[123419]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:937
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:937
		// _ = "end of CoverTab[123416]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:937
		_go_fuzz_dep_.CoverTab[123417]++
												mval.Index(i).Set(val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:938
		// _ = "end of CoverTab[123417]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:939
	// _ = "end of CoverTab[123412]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:939
	_go_fuzz_dep_.CoverTab[123413]++
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:940
	// _ = "end of CoverTab[123413]"
}

// Convert toml value to marshal primitive slice, using marshal type
func (d *Decoder) valueFromOtherSliceI(mtype reflect.Type, tval interface{}) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:944
	_go_fuzz_dep_.CoverTab[123420]++
											val := reflect.ValueOf(tval)
											length := val.Len()

											mval, err := makeSliceOrArray(mtype, length)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:949
		_go_fuzz_dep_.CoverTab[123423]++
												return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:950
		// _ = "end of CoverTab[123423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:951
		_go_fuzz_dep_.CoverTab[123424]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:951
		// _ = "end of CoverTab[123424]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:951
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:951
	// _ = "end of CoverTab[123420]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:951
	_go_fuzz_dep_.CoverTab[123421]++

											for i := 0; i < length; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:953
		_go_fuzz_dep_.CoverTab[123425]++
												val, err := d.valueFromToml(mtype.Elem(), val.Index(i).Interface(), nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:955
			_go_fuzz_dep_.CoverTab[123427]++
													return mval, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:956
			// _ = "end of CoverTab[123427]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:957
			_go_fuzz_dep_.CoverTab[123428]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:957
			// _ = "end of CoverTab[123428]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:957
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:957
		// _ = "end of CoverTab[123425]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:957
		_go_fuzz_dep_.CoverTab[123426]++
												mval.Index(i).Set(val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:958
		// _ = "end of CoverTab[123426]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:959
	// _ = "end of CoverTab[123421]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:959
	_go_fuzz_dep_.CoverTab[123422]++
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:960
	// _ = "end of CoverTab[123422]"
}

// Create a new slice or a new array with specified length
func makeSliceOrArray(mtype reflect.Type, tLength int) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:964
	_go_fuzz_dep_.CoverTab[123429]++
											var mval reflect.Value
											switch mtype.Kind() {
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:967
		_go_fuzz_dep_.CoverTab[123431]++
												mval = reflect.MakeSlice(mtype, tLength, tLength)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:968
		// _ = "end of CoverTab[123431]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:969
		_go_fuzz_dep_.CoverTab[123432]++
												mval = reflect.New(reflect.ArrayOf(mtype.Len(), mtype.Elem())).Elem()
												if tLength > mtype.Len() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:971
			_go_fuzz_dep_.CoverTab[123434]++
													return mval, fmt.Errorf("unmarshal: TOML array length (%v) exceeds destination array length (%v)", tLength, mtype.Len())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:972
			// _ = "end of CoverTab[123434]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
			_go_fuzz_dep_.CoverTab[123435]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
			// _ = "end of CoverTab[123435]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
		// _ = "end of CoverTab[123432]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
		_go_fuzz_dep_.CoverTab[123433]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:973
		// _ = "end of CoverTab[123433]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:974
	// _ = "end of CoverTab[123429]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:974
	_go_fuzz_dep_.CoverTab[123430]++
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:975
	// _ = "end of CoverTab[123430]"
}

// Convert toml value to marshal value, using marshal type. When mval1 is non-nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:978
// and the given type is a struct value, merge fields into it.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:980
func (d *Decoder) valueFromToml(mtype reflect.Type, tval interface{}, mval1 *reflect.Value) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:980
	_go_fuzz_dep_.CoverTab[123436]++
											if mtype.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:981
		_go_fuzz_dep_.CoverTab[123438]++
												return d.unwrapPointer(mtype, tval, mval1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:982
		// _ = "end of CoverTab[123438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:983
		_go_fuzz_dep_.CoverTab[123439]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:983
		// _ = "end of CoverTab[123439]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:983
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:983
	// _ = "end of CoverTab[123436]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:983
	_go_fuzz_dep_.CoverTab[123437]++

											switch t := tval.(type) {
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:986
		_go_fuzz_dep_.CoverTab[123440]++
												var mval11 *reflect.Value
												if mtype.Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:988
			_go_fuzz_dep_.CoverTab[123453]++
													mval11 = mval1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:989
			// _ = "end of CoverTab[123453]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:990
			_go_fuzz_dep_.CoverTab[123454]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:990
			// _ = "end of CoverTab[123454]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:990
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:990
		// _ = "end of CoverTab[123440]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:990
		_go_fuzz_dep_.CoverTab[123441]++

												if isTree(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:992
			_go_fuzz_dep_.CoverTab[123455]++
													return d.valueFromTree(mtype, t, mval11)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:993
			// _ = "end of CoverTab[123455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:994
			_go_fuzz_dep_.CoverTab[123456]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:994
			// _ = "end of CoverTab[123456]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:994
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:994
		// _ = "end of CoverTab[123441]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:994
		_go_fuzz_dep_.CoverTab[123442]++

												if mtype.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:996
			_go_fuzz_dep_.CoverTab[123457]++
													if mval1 == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:997
				_go_fuzz_dep_.CoverTab[123458]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:997
				return mval1.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:997
				// _ = "end of CoverTab[123458]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:997
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:997
				_go_fuzz_dep_.CoverTab[123459]++
														return d.valueFromTree(reflect.TypeOf(map[string]interface{}{}), t, nil)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:998
				// _ = "end of CoverTab[123459]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:999
				_go_fuzz_dep_.CoverTab[123460]++
														return d.valueFromToml(mval1.Elem().Type(), t, nil)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1000
				// _ = "end of CoverTab[123460]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1001
			// _ = "end of CoverTab[123457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1002
			_go_fuzz_dep_.CoverTab[123461]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1002
			// _ = "end of CoverTab[123461]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1002
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1002
		// _ = "end of CoverTab[123442]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1002
		_go_fuzz_dep_.CoverTab[123443]++

												return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to a tree", tval, tval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1004
		// _ = "end of CoverTab[123443]"
	case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1005
		_go_fuzz_dep_.CoverTab[123444]++
												if isTreeSequence(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1006
			_go_fuzz_dep_.CoverTab[123462]++
													return d.valueFromTreeSlice(mtype, t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1007
			// _ = "end of CoverTab[123462]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1008
			_go_fuzz_dep_.CoverTab[123463]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1008
			// _ = "end of CoverTab[123463]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1008
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1008
		// _ = "end of CoverTab[123444]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1008
		_go_fuzz_dep_.CoverTab[123445]++
												if mtype.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1009
			_go_fuzz_dep_.CoverTab[123464]++
													if mval1 == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1010
				_go_fuzz_dep_.CoverTab[123465]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1010
				return mval1.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1010
				// _ = "end of CoverTab[123465]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1010
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1010
				_go_fuzz_dep_.CoverTab[123466]++
														return d.valueFromTreeSlice(reflect.TypeOf([]map[string]interface{}{}), t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1011
				// _ = "end of CoverTab[123466]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1012
				_go_fuzz_dep_.CoverTab[123467]++
														ival := mval1.Elem()
														return d.valueFromToml(mval1.Elem().Type(), t, &ival)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1014
				// _ = "end of CoverTab[123467]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1015
			// _ = "end of CoverTab[123464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1016
			_go_fuzz_dep_.CoverTab[123468]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1016
			// _ = "end of CoverTab[123468]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1016
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1016
		// _ = "end of CoverTab[123445]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1016
		_go_fuzz_dep_.CoverTab[123446]++
												return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to trees", tval, tval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1017
		// _ = "end of CoverTab[123446]"
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1018
		_go_fuzz_dep_.CoverTab[123447]++
												d.visitor.visit()
												if isOtherSequence(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1020
			_go_fuzz_dep_.CoverTab[123469]++
													return d.valueFromOtherSlice(mtype, t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1021
			// _ = "end of CoverTab[123469]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1022
			_go_fuzz_dep_.CoverTab[123470]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1022
			// _ = "end of CoverTab[123470]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1022
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1022
		// _ = "end of CoverTab[123447]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1022
		_go_fuzz_dep_.CoverTab[123448]++
												if mtype.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1023
			_go_fuzz_dep_.CoverTab[123471]++
													if mval1 == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1024
				_go_fuzz_dep_.CoverTab[123472]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1024
				return mval1.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1024
				// _ = "end of CoverTab[123472]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1024
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1024
				_go_fuzz_dep_.CoverTab[123473]++
														return d.valueFromOtherSlice(reflect.TypeOf([]interface{}{}), t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1025
				// _ = "end of CoverTab[123473]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1026
				_go_fuzz_dep_.CoverTab[123474]++
														ival := mval1.Elem()
														return d.valueFromToml(mval1.Elem().Type(), t, &ival)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1028
				// _ = "end of CoverTab[123474]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1029
			// _ = "end of CoverTab[123471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1030
			_go_fuzz_dep_.CoverTab[123475]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1030
			// _ = "end of CoverTab[123475]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1030
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1030
		// _ = "end of CoverTab[123448]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1030
		_go_fuzz_dep_.CoverTab[123449]++
												return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to a slice", tval, tval)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1031
		// _ = "end of CoverTab[123449]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1032
		_go_fuzz_dep_.CoverTab[123450]++
												d.visitor.visit()
												mvalPtr := reflect.New(mtype)

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1037
		if isCustomUnmarshaler(mvalPtr.Type()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1037
			_go_fuzz_dep_.CoverTab[123476]++
													if err := callCustomUnmarshaler(mvalPtr, tval); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1038
				_go_fuzz_dep_.CoverTab[123478]++
														return reflect.ValueOf(nil), fmt.Errorf("unmarshal toml: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1039
				// _ = "end of CoverTab[123478]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1040
				_go_fuzz_dep_.CoverTab[123479]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1040
				// _ = "end of CoverTab[123479]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1040
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1040
			// _ = "end of CoverTab[123476]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1040
			_go_fuzz_dep_.CoverTab[123477]++
													return mvalPtr.Elem(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1041
			// _ = "end of CoverTab[123477]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1042
			_go_fuzz_dep_.CoverTab[123480]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1042
			// _ = "end of CoverTab[123480]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1042
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1042
		// _ = "end of CoverTab[123450]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1042
		_go_fuzz_dep_.CoverTab[123451]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
		if isTextUnmarshaler(mvalPtr.Type()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
			_go_fuzz_dep_.CoverTab[123481]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
			return !isTimeType(mtype)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
			// _ = "end of CoverTab[123481]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1045
			_go_fuzz_dep_.CoverTab[123482]++
													if err := d.unmarshalText(tval, mvalPtr); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1046
				_go_fuzz_dep_.CoverTab[123484]++
														return reflect.ValueOf(nil), fmt.Errorf("unmarshal text: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1047
				// _ = "end of CoverTab[123484]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1048
				_go_fuzz_dep_.CoverTab[123485]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1048
				// _ = "end of CoverTab[123485]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1048
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1048
			// _ = "end of CoverTab[123482]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1048
			_go_fuzz_dep_.CoverTab[123483]++
													return mvalPtr.Elem(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1049
			// _ = "end of CoverTab[123483]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1050
			_go_fuzz_dep_.CoverTab[123486]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1050
			// _ = "end of CoverTab[123486]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1050
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1050
		// _ = "end of CoverTab[123451]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1050
		_go_fuzz_dep_.CoverTab[123452]++

												switch mtype.Kind() {
		case reflect.Bool, reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1053
			_go_fuzz_dep_.CoverTab[123487]++
													val := reflect.ValueOf(tval)

													switch val.Type() {
			case localDateType:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1057
				_go_fuzz_dep_.CoverTab[123507]++
														localDate := val.Interface().(LocalDate)
														switch mtype {
				case timeType:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1060
					_go_fuzz_dep_.CoverTab[123510]++
															return reflect.ValueOf(time.Date(localDate.Year, localDate.Month, localDate.Day, 0, 0, 0, 0, time.Local)), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1061
					// _ = "end of CoverTab[123510]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1061
				default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1061
					_go_fuzz_dep_.CoverTab[123511]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1061
					// _ = "end of CoverTab[123511]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1062
				// _ = "end of CoverTab[123507]"
			case localDateTimeType:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1063
				_go_fuzz_dep_.CoverTab[123508]++
														localDateTime := val.Interface().(LocalDateTime)
														switch mtype {
				case timeType:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1066
					_go_fuzz_dep_.CoverTab[123512]++
															return reflect.ValueOf(time.Date(
						localDateTime.Date.Year,
						localDateTime.Date.Month,
						localDateTime.Date.Day,
						localDateTime.Time.Hour,
						localDateTime.Time.Minute,
						localDateTime.Time.Second,
						localDateTime.Time.Nanosecond,
						time.Local)), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1075
					// _ = "end of CoverTab[123512]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1075
				default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1075
					_go_fuzz_dep_.CoverTab[123513]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1075
					// _ = "end of CoverTab[123513]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1076
				// _ = "end of CoverTab[123508]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1076
			default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1076
				_go_fuzz_dep_.CoverTab[123509]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1076
				// _ = "end of CoverTab[123509]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1077
			// _ = "end of CoverTab[123487]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1077
			_go_fuzz_dep_.CoverTab[123488]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1080
			if !val.Type().ConvertibleTo(mtype) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1080
				_go_fuzz_dep_.CoverTab[123514]++
														return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1081
				// _ = "end of CoverTab[123514]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1082
				_go_fuzz_dep_.CoverTab[123515]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1082
				// _ = "end of CoverTab[123515]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1082
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1082
			// _ = "end of CoverTab[123488]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1082
			_go_fuzz_dep_.CoverTab[123489]++

													return val.Convert(mtype), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1084
			// _ = "end of CoverTab[123489]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1085
			_go_fuzz_dep_.CoverTab[123490]++
													val := reflect.ValueOf(tval)

													if !val.Type().ConvertibleTo(mtype) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1088
				_go_fuzz_dep_.CoverTab[123516]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1088
				return val.Kind() == reflect.Int64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1088
				// _ = "end of CoverTab[123516]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1088
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1088
				_go_fuzz_dep_.CoverTab[123517]++
														return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1089
				// _ = "end of CoverTab[123517]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1090
				_go_fuzz_dep_.CoverTab[123518]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1090
				// _ = "end of CoverTab[123518]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1090
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1090
			// _ = "end of CoverTab[123490]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1090
			_go_fuzz_dep_.CoverTab[123491]++

													return val.Convert(mtype), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1092
			// _ = "end of CoverTab[123491]"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1093
			_go_fuzz_dep_.CoverTab[123492]++
													val := reflect.ValueOf(tval)
													if mtype.Kind() == reflect.Int64 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				_go_fuzz_dep_.CoverTab[123519]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				return mtype == reflect.TypeOf(time.Duration(1))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				// _ = "end of CoverTab[123519]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				_go_fuzz_dep_.CoverTab[123520]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				return val.Kind() == reflect.String
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				// _ = "end of CoverTab[123520]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1095
				_go_fuzz_dep_.CoverTab[123521]++
														d, err := time.ParseDuration(val.String())
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1097
					_go_fuzz_dep_.CoverTab[123523]++
															return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v. %s", tval, tval, mtype.String(), err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1098
					// _ = "end of CoverTab[123523]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1099
					_go_fuzz_dep_.CoverTab[123524]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1099
					// _ = "end of CoverTab[123524]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1099
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1099
				// _ = "end of CoverTab[123521]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1099
				_go_fuzz_dep_.CoverTab[123522]++
														return reflect.ValueOf(d), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1100
				// _ = "end of CoverTab[123522]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1101
				_go_fuzz_dep_.CoverTab[123525]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1101
				// _ = "end of CoverTab[123525]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1101
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1101
			// _ = "end of CoverTab[123492]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1101
			_go_fuzz_dep_.CoverTab[123493]++
													if !val.Type().ConvertibleTo(mtype) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1102
				_go_fuzz_dep_.CoverTab[123526]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1102
				return val.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1102
				// _ = "end of CoverTab[123526]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1102
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1102
				_go_fuzz_dep_.CoverTab[123527]++
														return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1103
				// _ = "end of CoverTab[123527]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1104
				_go_fuzz_dep_.CoverTab[123528]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1104
				// _ = "end of CoverTab[123528]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1104
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1104
			// _ = "end of CoverTab[123493]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1104
			_go_fuzz_dep_.CoverTab[123494]++
													if reflect.Indirect(reflect.New(mtype)).OverflowInt(val.Convert(reflect.TypeOf(int64(0))).Int()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1105
				_go_fuzz_dep_.CoverTab[123529]++
														return reflect.ValueOf(nil), fmt.Errorf("%v(%T) would overflow %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1106
				// _ = "end of CoverTab[123529]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1107
				_go_fuzz_dep_.CoverTab[123530]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1107
				// _ = "end of CoverTab[123530]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1107
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1107
			// _ = "end of CoverTab[123494]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1107
			_go_fuzz_dep_.CoverTab[123495]++

													return val.Convert(mtype), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1109
			// _ = "end of CoverTab[123495]"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1110
			_go_fuzz_dep_.CoverTab[123496]++
													val := reflect.ValueOf(tval)
													if !val.Type().ConvertibleTo(mtype) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1112
				_go_fuzz_dep_.CoverTab[123531]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1112
				return val.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1112
				// _ = "end of CoverTab[123531]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1112
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1112
				_go_fuzz_dep_.CoverTab[123532]++
														return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1113
				// _ = "end of CoverTab[123532]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1114
				_go_fuzz_dep_.CoverTab[123533]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1114
				// _ = "end of CoverTab[123533]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1114
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1114
			// _ = "end of CoverTab[123496]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1114
			_go_fuzz_dep_.CoverTab[123497]++

													if val.Convert(reflect.TypeOf(int(1))).Int() < 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1116
				_go_fuzz_dep_.CoverTab[123534]++
														return reflect.ValueOf(nil), fmt.Errorf("%v(%T) is negative so does not fit in %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1117
				// _ = "end of CoverTab[123534]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1118
				_go_fuzz_dep_.CoverTab[123535]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1118
				// _ = "end of CoverTab[123535]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1118
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1118
			// _ = "end of CoverTab[123497]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1118
			_go_fuzz_dep_.CoverTab[123498]++
													if reflect.Indirect(reflect.New(mtype)).OverflowUint(val.Convert(reflect.TypeOf(uint64(0))).Uint()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1119
				_go_fuzz_dep_.CoverTab[123536]++
														return reflect.ValueOf(nil), fmt.Errorf("%v(%T) would overflow %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1120
				// _ = "end of CoverTab[123536]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1121
				_go_fuzz_dep_.CoverTab[123537]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1121
				// _ = "end of CoverTab[123537]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1121
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1121
			// _ = "end of CoverTab[123498]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1121
			_go_fuzz_dep_.CoverTab[123499]++

													return val.Convert(mtype), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1123
			// _ = "end of CoverTab[123499]"
		case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1124
			_go_fuzz_dep_.CoverTab[123500]++
													val := reflect.ValueOf(tval)
													if !val.Type().ConvertibleTo(mtype) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1126
				_go_fuzz_dep_.CoverTab[123538]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1126
				return val.Kind() == reflect.Int64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1126
				// _ = "end of CoverTab[123538]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1126
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1126
				_go_fuzz_dep_.CoverTab[123539]++
														return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1127
				// _ = "end of CoverTab[123539]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1128
				_go_fuzz_dep_.CoverTab[123540]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1128
				// _ = "end of CoverTab[123540]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1128
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1128
			// _ = "end of CoverTab[123500]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1128
			_go_fuzz_dep_.CoverTab[123501]++
													if reflect.Indirect(reflect.New(mtype)).OverflowFloat(val.Convert(reflect.TypeOf(float64(0))).Float()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1129
				_go_fuzz_dep_.CoverTab[123541]++
														return reflect.ValueOf(nil), fmt.Errorf("%v(%T) would overflow %v", tval, tval, mtype.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1130
				// _ = "end of CoverTab[123541]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1131
				_go_fuzz_dep_.CoverTab[123542]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1131
				// _ = "end of CoverTab[123542]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1131
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1131
			// _ = "end of CoverTab[123501]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1131
			_go_fuzz_dep_.CoverTab[123502]++

													return val.Convert(mtype), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1133
			// _ = "end of CoverTab[123502]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1134
			_go_fuzz_dep_.CoverTab[123503]++
													if mval1 == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1135
				_go_fuzz_dep_.CoverTab[123543]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1135
				return mval1.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1135
				// _ = "end of CoverTab[123543]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1135
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1135
				_go_fuzz_dep_.CoverTab[123544]++
														return reflect.ValueOf(tval), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1136
				// _ = "end of CoverTab[123544]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1137
				_go_fuzz_dep_.CoverTab[123545]++
														ival := mval1.Elem()
														return d.valueFromToml(mval1.Elem().Type(), t, &ival)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1139
				// _ = "end of CoverTab[123545]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1140
			// _ = "end of CoverTab[123503]"
		case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1141
			_go_fuzz_dep_.CoverTab[123504]++
													if isOtherSequence(mtype) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1142
				_go_fuzz_dep_.CoverTab[123546]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1142
				return isOtherSequence(reflect.TypeOf(t))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1142
				// _ = "end of CoverTab[123546]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1142
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1142
				_go_fuzz_dep_.CoverTab[123547]++
														return d.valueFromOtherSliceI(mtype, t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1143
				// _ = "end of CoverTab[123547]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1144
				_go_fuzz_dep_.CoverTab[123548]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1144
				// _ = "end of CoverTab[123548]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1144
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1144
			// _ = "end of CoverTab[123504]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1144
			_go_fuzz_dep_.CoverTab[123505]++
													return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v(%v)", tval, tval, mtype, mtype.Kind())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1145
			// _ = "end of CoverTab[123505]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1146
			_go_fuzz_dep_.CoverTab[123506]++
													return reflect.ValueOf(nil), fmt.Errorf("Can't convert %v(%T) to %v(%v)", tval, tval, mtype, mtype.Kind())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1147
			// _ = "end of CoverTab[123506]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1148
		// _ = "end of CoverTab[123452]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1149
	// _ = "end of CoverTab[123437]"
}

func (d *Decoder) unwrapPointer(mtype reflect.Type, tval interface{}, mval1 *reflect.Value) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1152
	_go_fuzz_dep_.CoverTab[123549]++
											var melem *reflect.Value

											if mval1 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		_go_fuzz_dep_.CoverTab[123552]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		return !mval1.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		// _ = "end of CoverTab[123552]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		_go_fuzz_dep_.CoverTab[123553]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		return (mtype.Elem().Kind() == reflect.Struct || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
			_go_fuzz_dep_.CoverTab[123554]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
			return mtype.Elem().Kind() == reflect.Interface
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
			// _ = "end of CoverTab[123554]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		// _ = "end of CoverTab[123553]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1155
		_go_fuzz_dep_.CoverTab[123555]++
												elem := mval1.Elem()
												melem = &elem
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1157
		// _ = "end of CoverTab[123555]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1158
		_go_fuzz_dep_.CoverTab[123556]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1158
		// _ = "end of CoverTab[123556]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1158
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1158
	// _ = "end of CoverTab[123549]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1158
	_go_fuzz_dep_.CoverTab[123550]++

											val, err := d.valueFromToml(mtype.Elem(), tval, melem)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1161
		_go_fuzz_dep_.CoverTab[123557]++
												return reflect.ValueOf(nil), err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1162
		// _ = "end of CoverTab[123557]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1163
		_go_fuzz_dep_.CoverTab[123558]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1163
		// _ = "end of CoverTab[123558]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1163
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1163
	// _ = "end of CoverTab[123550]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1163
	_go_fuzz_dep_.CoverTab[123551]++
											mval := reflect.New(mtype.Elem())
											mval.Elem().Set(val)
											return mval, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1166
	// _ = "end of CoverTab[123551]"
}

func (d *Decoder) unmarshalText(tval interface{}, mval reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1169
	_go_fuzz_dep_.CoverTab[123559]++
											var buf bytes.Buffer
											fmt.Fprint(&buf, tval)
											return callTextUnmarshaler(mval, buf.Bytes())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1172
	// _ = "end of CoverTab[123559]"
}

func tomlOptions(vf reflect.StructField, an annotation) tomlOpts {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1175
	_go_fuzz_dep_.CoverTab[123560]++
											tag := vf.Tag.Get(an.tag)
											parse := strings.Split(tag, ",")
											var comment string
											if c := vf.Tag.Get(an.comment); c != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1179
		_go_fuzz_dep_.CoverTab[123566]++
												comment = c
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1180
		// _ = "end of CoverTab[123566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1181
		_go_fuzz_dep_.CoverTab[123567]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1181
		// _ = "end of CoverTab[123567]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1181
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1181
	// _ = "end of CoverTab[123560]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1181
	_go_fuzz_dep_.CoverTab[123561]++
											commented, _ := strconv.ParseBool(vf.Tag.Get(an.commented))
											multiline, _ := strconv.ParseBool(vf.Tag.Get(an.multiline))
											literal, _ := strconv.ParseBool(vf.Tag.Get(an.literal))
											defaultValue := vf.Tag.Get(tagDefault)
											result := tomlOpts{
		name:		vf.Name,
		nameFromTag:	false,
		comment:	comment,
		commented:	commented,
		multiline:	multiline,
		literal:	literal,
		include:	true,
		omitempty:	false,
		defaultValue:	defaultValue,
	}
	if parse[0] != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1197
		_go_fuzz_dep_.CoverTab[123568]++
												if parse[0] == "-" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1198
			_go_fuzz_dep_.CoverTab[123569]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1198
			return len(parse) == 1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1198
			// _ = "end of CoverTab[123569]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1198
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1198
			_go_fuzz_dep_.CoverTab[123570]++
													result.include = false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1199
			// _ = "end of CoverTab[123570]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1200
			_go_fuzz_dep_.CoverTab[123571]++
													result.name = strings.Trim(parse[0], " ")
													result.nameFromTag = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1202
			// _ = "end of CoverTab[123571]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1203
		// _ = "end of CoverTab[123568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1204
		_go_fuzz_dep_.CoverTab[123572]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1204
		// _ = "end of CoverTab[123572]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1204
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1204
	// _ = "end of CoverTab[123561]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1204
	_go_fuzz_dep_.CoverTab[123562]++
											if vf.PkgPath != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1205
		_go_fuzz_dep_.CoverTab[123573]++
												result.include = false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1206
		// _ = "end of CoverTab[123573]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1207
		_go_fuzz_dep_.CoverTab[123574]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1207
		// _ = "end of CoverTab[123574]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1207
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1207
	// _ = "end of CoverTab[123562]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1207
	_go_fuzz_dep_.CoverTab[123563]++
											if len(parse) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1208
		_go_fuzz_dep_.CoverTab[123575]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1208
		return strings.Trim(parse[1], " ") == "omitempty"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1208
		// _ = "end of CoverTab[123575]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1208
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1208
		_go_fuzz_dep_.CoverTab[123576]++
												result.omitempty = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1209
		// _ = "end of CoverTab[123576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1210
		_go_fuzz_dep_.CoverTab[123577]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1210
		// _ = "end of CoverTab[123577]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1210
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1210
	// _ = "end of CoverTab[123563]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1210
	_go_fuzz_dep_.CoverTab[123564]++
											if vf.Type.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1211
		_go_fuzz_dep_.CoverTab[123578]++
												result.omitempty = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1212
		// _ = "end of CoverTab[123578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1213
		_go_fuzz_dep_.CoverTab[123579]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1213
		// _ = "end of CoverTab[123579]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1213
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1213
	// _ = "end of CoverTab[123564]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1213
	_go_fuzz_dep_.CoverTab[123565]++
											return result
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1214
	// _ = "end of CoverTab[123565]"
}

func isZero(val reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1217
	_go_fuzz_dep_.CoverTab[123580]++
											switch val.Type().Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1219
		_go_fuzz_dep_.CoverTab[123581]++
												return val.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1220
		// _ = "end of CoverTab[123581]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1221
		_go_fuzz_dep_.CoverTab[123582]++
												return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1222
		// _ = "end of CoverTab[123582]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1223
	// _ = "end of CoverTab[123580]"
}

func formatError(err error, pos Position) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1226
	_go_fuzz_dep_.CoverTab[123583]++
											if err.Error()[0] == '(' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1227
		_go_fuzz_dep_.CoverTab[123585]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1228
		// _ = "end of CoverTab[123585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1229
		_go_fuzz_dep_.CoverTab[123586]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1229
		// _ = "end of CoverTab[123586]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1229
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1229
	// _ = "end of CoverTab[123583]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1229
	_go_fuzz_dep_.CoverTab[123584]++
											return fmt.Errorf("%s: %s", pos, err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1230
	// _ = "end of CoverTab[123584]"
}

// visitorState keeps track of which keys were unmarshaled.
type visitorState struct {
	tree	*Tree
	path	[]string
	keys	map[string]struct{}
	active	bool
}

func newVisitorState(tree *Tree) visitorState {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1241
	_go_fuzz_dep_.CoverTab[123587]++
											path, result := []string{}, map[string]struct{}{}
											insertKeys(path, result, tree)
											return visitorState{
		tree:	tree,
		path:	path[:0],
		keys:	result,
		active:	true,
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1249
	// _ = "end of CoverTab[123587]"
}

func (s *visitorState) push(key string) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1252
	_go_fuzz_dep_.CoverTab[123588]++
											if s.active {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1253
		_go_fuzz_dep_.CoverTab[123589]++
												s.path = append(s.path, key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1254
		// _ = "end of CoverTab[123589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1255
		_go_fuzz_dep_.CoverTab[123590]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1255
		// _ = "end of CoverTab[123590]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1255
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1255
	// _ = "end of CoverTab[123588]"
}

func (s *visitorState) pop() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1258
	_go_fuzz_dep_.CoverTab[123591]++
											if s.active {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1259
		_go_fuzz_dep_.CoverTab[123592]++
												s.path = s.path[:len(s.path)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1260
		// _ = "end of CoverTab[123592]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1261
		_go_fuzz_dep_.CoverTab[123593]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1261
		// _ = "end of CoverTab[123593]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1261
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1261
	// _ = "end of CoverTab[123591]"
}

func (s *visitorState) visit() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1264
	_go_fuzz_dep_.CoverTab[123594]++
											if s.active {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1265
		_go_fuzz_dep_.CoverTab[123595]++
												delete(s.keys, strings.Join(s.path, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1266
		// _ = "end of CoverTab[123595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1267
		_go_fuzz_dep_.CoverTab[123596]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1267
		// _ = "end of CoverTab[123596]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1267
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1267
	// _ = "end of CoverTab[123594]"
}

func (s *visitorState) visitAll() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1270
	_go_fuzz_dep_.CoverTab[123597]++
											if s.active {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1271
		_go_fuzz_dep_.CoverTab[123598]++
												for k := range s.keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1272
			_go_fuzz_dep_.CoverTab[123599]++
													if strings.HasPrefix(k, strings.Join(s.path, ".")) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1273
				_go_fuzz_dep_.CoverTab[123600]++
														delete(s.keys, k)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1274
				// _ = "end of CoverTab[123600]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1275
				_go_fuzz_dep_.CoverTab[123601]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1275
				// _ = "end of CoverTab[123601]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1275
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1275
			// _ = "end of CoverTab[123599]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1276
		// _ = "end of CoverTab[123598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1277
		_go_fuzz_dep_.CoverTab[123602]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1277
		// _ = "end of CoverTab[123602]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1277
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1277
	// _ = "end of CoverTab[123597]"
}

func (s *visitorState) validate() error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1280
	_go_fuzz_dep_.CoverTab[123603]++
											if !s.active {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1281
		_go_fuzz_dep_.CoverTab[123607]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1282
		// _ = "end of CoverTab[123607]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1283
		_go_fuzz_dep_.CoverTab[123608]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1283
		// _ = "end of CoverTab[123608]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1283
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1283
	// _ = "end of CoverTab[123603]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1283
	_go_fuzz_dep_.CoverTab[123604]++
											undecoded := make([]string, 0, len(s.keys))
											for key := range s.keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1285
		_go_fuzz_dep_.CoverTab[123609]++
												undecoded = append(undecoded, key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1286
		// _ = "end of CoverTab[123609]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1287
	// _ = "end of CoverTab[123604]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1287
	_go_fuzz_dep_.CoverTab[123605]++
											sort.Strings(undecoded)
											if len(undecoded) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1289
		_go_fuzz_dep_.CoverTab[123610]++
												return fmt.Errorf("undecoded keys: %q", undecoded)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1290
		// _ = "end of CoverTab[123610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1291
		_go_fuzz_dep_.CoverTab[123611]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1291
		// _ = "end of CoverTab[123611]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1291
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1291
	// _ = "end of CoverTab[123605]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1291
	_go_fuzz_dep_.CoverTab[123606]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1292
	// _ = "end of CoverTab[123606]"
}

func insertKeys(path []string, m map[string]struct{}, tree *Tree) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1295
	_go_fuzz_dep_.CoverTab[123612]++
											for k, v := range tree.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1296
		_go_fuzz_dep_.CoverTab[123613]++
												switch node := v.(type) {
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1298
			_go_fuzz_dep_.CoverTab[123614]++
													for i, item := range node {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1299
				_go_fuzz_dep_.CoverTab[123617]++
														insertKeys(append(path, k, strconv.Itoa(i)), m, item)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1300
				// _ = "end of CoverTab[123617]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1301
			// _ = "end of CoverTab[123614]"
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1302
			_go_fuzz_dep_.CoverTab[123615]++
													insertKeys(append(path, k), m, node)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1303
			// _ = "end of CoverTab[123615]"
		case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1304
			_go_fuzz_dep_.CoverTab[123616]++
													m[strings.Join(append(path, k), ".")] = struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1305
			// _ = "end of CoverTab[123616]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1306
		// _ = "end of CoverTab[123613]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1307
	// _ = "end of CoverTab[123612]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1308
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/marshal.go:1308
var _ = _go_fuzz_dep_.CoverTab
