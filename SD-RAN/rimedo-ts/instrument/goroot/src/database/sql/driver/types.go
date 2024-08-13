// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/database/sql/driver/types.go:5
package driver

//line /usr/local/go/src/database/sql/driver/types.go:5
import (
//line /usr/local/go/src/database/sql/driver/types.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/database/sql/driver/types.go:5
)
//line /usr/local/go/src/database/sql/driver/types.go:5
import (
//line /usr/local/go/src/database/sql/driver/types.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/database/sql/driver/types.go:5
)

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ValueConverter is the interface providing the ConvertValue method.
//line /usr/local/go/src/database/sql/driver/types.go:14
//
//line /usr/local/go/src/database/sql/driver/types.go:14
// Various implementations of ValueConverter are provided by the
//line /usr/local/go/src/database/sql/driver/types.go:14
// driver package to provide consistent implementations of conversions
//line /usr/local/go/src/database/sql/driver/types.go:14
// between drivers. The ValueConverters have several uses:
//line /usr/local/go/src/database/sql/driver/types.go:14
//
//line /usr/local/go/src/database/sql/driver/types.go:14
//   - converting from the Value types as provided by the sql package
//line /usr/local/go/src/database/sql/driver/types.go:14
//     into a database table's specific column type and making sure it
//line /usr/local/go/src/database/sql/driver/types.go:14
//     fits, such as making sure a particular int64 fits in a
//line /usr/local/go/src/database/sql/driver/types.go:14
//     table's uint16 column.
//line /usr/local/go/src/database/sql/driver/types.go:14
//
//line /usr/local/go/src/database/sql/driver/types.go:14
//   - converting a value as given from the database into one of the
//line /usr/local/go/src/database/sql/driver/types.go:14
//     driver Value types.
//line /usr/local/go/src/database/sql/driver/types.go:14
//
//line /usr/local/go/src/database/sql/driver/types.go:14
//   - by the sql package, for converting from a driver's Value type
//line /usr/local/go/src/database/sql/driver/types.go:14
//     to a user's type in a scan.
//line /usr/local/go/src/database/sql/driver/types.go:30
type ValueConverter interface {
	// ConvertValue converts a value to a driver Value.
	ConvertValue(v any) (Value, error)
}

// Valuer is the interface providing the Value method.
//line /usr/local/go/src/database/sql/driver/types.go:35
//
//line /usr/local/go/src/database/sql/driver/types.go:35
// Types implementing Valuer interface are able to convert
//line /usr/local/go/src/database/sql/driver/types.go:35
// themselves to a driver Value.
//line /usr/local/go/src/database/sql/driver/types.go:39
type Valuer interface {
	// Value returns a driver Value.
	// Value must not panic.
	Value() (Value, error)
}

// Bool is a ValueConverter that converts input values to bools.
//line /usr/local/go/src/database/sql/driver/types.go:45
//
//line /usr/local/go/src/database/sql/driver/types.go:45
// The conversion rules are:
//line /usr/local/go/src/database/sql/driver/types.go:45
//   - booleans are returned unchanged
//line /usr/local/go/src/database/sql/driver/types.go:45
//   - for integer types,
//line /usr/local/go/src/database/sql/driver/types.go:45
//     1 is true
//line /usr/local/go/src/database/sql/driver/types.go:45
//     0 is false,
//line /usr/local/go/src/database/sql/driver/types.go:45
//     other integers are an error
//line /usr/local/go/src/database/sql/driver/types.go:45
//   - for strings and []byte, same rules as strconv.ParseBool
//line /usr/local/go/src/database/sql/driver/types.go:45
//   - all other types are an error
//line /usr/local/go/src/database/sql/driver/types.go:55
var Bool boolType

type boolType struct{}

var _ ValueConverter = boolType{}

func (boolType) String() string {
//line /usr/local/go/src/database/sql/driver/types.go:61
	_go_fuzz_dep_.CoverTab[179205]++
//line /usr/local/go/src/database/sql/driver/types.go:61
	return "Bool"
//line /usr/local/go/src/database/sql/driver/types.go:61
	// _ = "end of CoverTab[179205]"
//line /usr/local/go/src/database/sql/driver/types.go:61
}

func (boolType) ConvertValue(src any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:63
	_go_fuzz_dep_.CoverTab[179206]++
								switch s := src.(type) {
	case bool:
//line /usr/local/go/src/database/sql/driver/types.go:65
		_go_fuzz_dep_.CoverTab[179209]++
									return s, nil
//line /usr/local/go/src/database/sql/driver/types.go:66
		// _ = "end of CoverTab[179209]"
	case string:
//line /usr/local/go/src/database/sql/driver/types.go:67
		_go_fuzz_dep_.CoverTab[179210]++
									b, err := strconv.ParseBool(s)
									if err != nil {
//line /usr/local/go/src/database/sql/driver/types.go:69
			_go_fuzz_dep_.CoverTab[179214]++
										return nil, fmt.Errorf("sql/driver: couldn't convert %q into type bool", s)
//line /usr/local/go/src/database/sql/driver/types.go:70
			// _ = "end of CoverTab[179214]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:71
			_go_fuzz_dep_.CoverTab[179215]++
//line /usr/local/go/src/database/sql/driver/types.go:71
			// _ = "end of CoverTab[179215]"
//line /usr/local/go/src/database/sql/driver/types.go:71
		}
//line /usr/local/go/src/database/sql/driver/types.go:71
		// _ = "end of CoverTab[179210]"
//line /usr/local/go/src/database/sql/driver/types.go:71
		_go_fuzz_dep_.CoverTab[179211]++
									return b, nil
//line /usr/local/go/src/database/sql/driver/types.go:72
		// _ = "end of CoverTab[179211]"
	case []byte:
//line /usr/local/go/src/database/sql/driver/types.go:73
		_go_fuzz_dep_.CoverTab[179212]++
									b, err := strconv.ParseBool(string(s))
									if err != nil {
//line /usr/local/go/src/database/sql/driver/types.go:75
			_go_fuzz_dep_.CoverTab[179216]++
										return nil, fmt.Errorf("sql/driver: couldn't convert %q into type bool", s)
//line /usr/local/go/src/database/sql/driver/types.go:76
			// _ = "end of CoverTab[179216]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:77
			_go_fuzz_dep_.CoverTab[179217]++
//line /usr/local/go/src/database/sql/driver/types.go:77
			// _ = "end of CoverTab[179217]"
//line /usr/local/go/src/database/sql/driver/types.go:77
		}
//line /usr/local/go/src/database/sql/driver/types.go:77
		// _ = "end of CoverTab[179212]"
//line /usr/local/go/src/database/sql/driver/types.go:77
		_go_fuzz_dep_.CoverTab[179213]++
									return b, nil
//line /usr/local/go/src/database/sql/driver/types.go:78
		// _ = "end of CoverTab[179213]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:79
	// _ = "end of CoverTab[179206]"
//line /usr/local/go/src/database/sql/driver/types.go:79
	_go_fuzz_dep_.CoverTab[179207]++

								sv := reflect.ValueOf(src)
								switch sv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/database/sql/driver/types.go:83
		_go_fuzz_dep_.CoverTab[179218]++
									iv := sv.Int()
									if iv == 1 || func() bool {
//line /usr/local/go/src/database/sql/driver/types.go:85
			_go_fuzz_dep_.CoverTab[179223]++
//line /usr/local/go/src/database/sql/driver/types.go:85
			return iv == 0
//line /usr/local/go/src/database/sql/driver/types.go:85
			// _ = "end of CoverTab[179223]"
//line /usr/local/go/src/database/sql/driver/types.go:85
		}() {
//line /usr/local/go/src/database/sql/driver/types.go:85
			_go_fuzz_dep_.CoverTab[179224]++
										return iv == 1, nil
//line /usr/local/go/src/database/sql/driver/types.go:86
			// _ = "end of CoverTab[179224]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:87
			_go_fuzz_dep_.CoverTab[179225]++
//line /usr/local/go/src/database/sql/driver/types.go:87
			// _ = "end of CoverTab[179225]"
//line /usr/local/go/src/database/sql/driver/types.go:87
		}
//line /usr/local/go/src/database/sql/driver/types.go:87
		// _ = "end of CoverTab[179218]"
//line /usr/local/go/src/database/sql/driver/types.go:87
		_go_fuzz_dep_.CoverTab[179219]++
									return nil, fmt.Errorf("sql/driver: couldn't convert %d into type bool", iv)
//line /usr/local/go/src/database/sql/driver/types.go:88
		// _ = "end of CoverTab[179219]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /usr/local/go/src/database/sql/driver/types.go:89
		_go_fuzz_dep_.CoverTab[179220]++
									uv := sv.Uint()
									if uv == 1 || func() bool {
//line /usr/local/go/src/database/sql/driver/types.go:91
			_go_fuzz_dep_.CoverTab[179226]++
//line /usr/local/go/src/database/sql/driver/types.go:91
			return uv == 0
//line /usr/local/go/src/database/sql/driver/types.go:91
			// _ = "end of CoverTab[179226]"
//line /usr/local/go/src/database/sql/driver/types.go:91
		}() {
//line /usr/local/go/src/database/sql/driver/types.go:91
			_go_fuzz_dep_.CoverTab[179227]++
										return uv == 1, nil
//line /usr/local/go/src/database/sql/driver/types.go:92
			// _ = "end of CoverTab[179227]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:93
			_go_fuzz_dep_.CoverTab[179228]++
//line /usr/local/go/src/database/sql/driver/types.go:93
			// _ = "end of CoverTab[179228]"
//line /usr/local/go/src/database/sql/driver/types.go:93
		}
//line /usr/local/go/src/database/sql/driver/types.go:93
		// _ = "end of CoverTab[179220]"
//line /usr/local/go/src/database/sql/driver/types.go:93
		_go_fuzz_dep_.CoverTab[179221]++
									return nil, fmt.Errorf("sql/driver: couldn't convert %d into type bool", uv)
//line /usr/local/go/src/database/sql/driver/types.go:94
		// _ = "end of CoverTab[179221]"
//line /usr/local/go/src/database/sql/driver/types.go:94
	default:
//line /usr/local/go/src/database/sql/driver/types.go:94
		_go_fuzz_dep_.CoverTab[179222]++
//line /usr/local/go/src/database/sql/driver/types.go:94
		// _ = "end of CoverTab[179222]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:95
	// _ = "end of CoverTab[179207]"
//line /usr/local/go/src/database/sql/driver/types.go:95
	_go_fuzz_dep_.CoverTab[179208]++

								return nil, fmt.Errorf("sql/driver: couldn't convert %v (%T) into type bool", src, src)
//line /usr/local/go/src/database/sql/driver/types.go:97
	// _ = "end of CoverTab[179208]"
}

// Int32 is a ValueConverter that converts input values to int64,
//line /usr/local/go/src/database/sql/driver/types.go:100
// respecting the limits of an int32 value.
//line /usr/local/go/src/database/sql/driver/types.go:102
var Int32 int32Type

type int32Type struct{}

var _ ValueConverter = int32Type{}

func (int32Type) ConvertValue(v any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:108
	_go_fuzz_dep_.CoverTab[179229]++
								rv := reflect.ValueOf(v)
								switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/database/sql/driver/types.go:111
		_go_fuzz_dep_.CoverTab[179231]++
									i64 := rv.Int()
									if i64 > (1<<31)-1 || func() bool {
//line /usr/local/go/src/database/sql/driver/types.go:113
			_go_fuzz_dep_.CoverTab[179238]++
//line /usr/local/go/src/database/sql/driver/types.go:113
			return i64 < -(1 << 31)
//line /usr/local/go/src/database/sql/driver/types.go:113
			// _ = "end of CoverTab[179238]"
//line /usr/local/go/src/database/sql/driver/types.go:113
		}() {
//line /usr/local/go/src/database/sql/driver/types.go:113
			_go_fuzz_dep_.CoverTab[179239]++
										return nil, fmt.Errorf("sql/driver: value %d overflows int32", v)
//line /usr/local/go/src/database/sql/driver/types.go:114
			// _ = "end of CoverTab[179239]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:115
			_go_fuzz_dep_.CoverTab[179240]++
//line /usr/local/go/src/database/sql/driver/types.go:115
			// _ = "end of CoverTab[179240]"
//line /usr/local/go/src/database/sql/driver/types.go:115
		}
//line /usr/local/go/src/database/sql/driver/types.go:115
		// _ = "end of CoverTab[179231]"
//line /usr/local/go/src/database/sql/driver/types.go:115
		_go_fuzz_dep_.CoverTab[179232]++
									return i64, nil
//line /usr/local/go/src/database/sql/driver/types.go:116
		// _ = "end of CoverTab[179232]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /usr/local/go/src/database/sql/driver/types.go:117
		_go_fuzz_dep_.CoverTab[179233]++
									u64 := rv.Uint()
									if u64 > (1<<31)-1 {
//line /usr/local/go/src/database/sql/driver/types.go:119
			_go_fuzz_dep_.CoverTab[179241]++
										return nil, fmt.Errorf("sql/driver: value %d overflows int32", v)
//line /usr/local/go/src/database/sql/driver/types.go:120
			// _ = "end of CoverTab[179241]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:121
			_go_fuzz_dep_.CoverTab[179242]++
//line /usr/local/go/src/database/sql/driver/types.go:121
			// _ = "end of CoverTab[179242]"
//line /usr/local/go/src/database/sql/driver/types.go:121
		}
//line /usr/local/go/src/database/sql/driver/types.go:121
		// _ = "end of CoverTab[179233]"
//line /usr/local/go/src/database/sql/driver/types.go:121
		_go_fuzz_dep_.CoverTab[179234]++
									return int64(u64), nil
//line /usr/local/go/src/database/sql/driver/types.go:122
		// _ = "end of CoverTab[179234]"
	case reflect.String:
//line /usr/local/go/src/database/sql/driver/types.go:123
		_go_fuzz_dep_.CoverTab[179235]++
									i, err := strconv.Atoi(rv.String())
									if err != nil {
//line /usr/local/go/src/database/sql/driver/types.go:125
			_go_fuzz_dep_.CoverTab[179243]++
										return nil, fmt.Errorf("sql/driver: value %q can't be converted to int32", v)
//line /usr/local/go/src/database/sql/driver/types.go:126
			// _ = "end of CoverTab[179243]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:127
			_go_fuzz_dep_.CoverTab[179244]++
//line /usr/local/go/src/database/sql/driver/types.go:127
			// _ = "end of CoverTab[179244]"
//line /usr/local/go/src/database/sql/driver/types.go:127
		}
//line /usr/local/go/src/database/sql/driver/types.go:127
		// _ = "end of CoverTab[179235]"
//line /usr/local/go/src/database/sql/driver/types.go:127
		_go_fuzz_dep_.CoverTab[179236]++
									return int64(i), nil
//line /usr/local/go/src/database/sql/driver/types.go:128
		// _ = "end of CoverTab[179236]"
//line /usr/local/go/src/database/sql/driver/types.go:128
	default:
//line /usr/local/go/src/database/sql/driver/types.go:128
		_go_fuzz_dep_.CoverTab[179237]++
//line /usr/local/go/src/database/sql/driver/types.go:128
		// _ = "end of CoverTab[179237]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:129
	// _ = "end of CoverTab[179229]"
//line /usr/local/go/src/database/sql/driver/types.go:129
	_go_fuzz_dep_.CoverTab[179230]++
								return nil, fmt.Errorf("sql/driver: unsupported value %v (type %T) converting to int32", v, v)
//line /usr/local/go/src/database/sql/driver/types.go:130
	// _ = "end of CoverTab[179230]"
}

// String is a ValueConverter that converts its input to a string.
//line /usr/local/go/src/database/sql/driver/types.go:133
// If the value is already a string or []byte, it's unchanged.
//line /usr/local/go/src/database/sql/driver/types.go:133
// If the value is of another type, conversion to string is done
//line /usr/local/go/src/database/sql/driver/types.go:133
// with fmt.Sprintf("%v", v).
//line /usr/local/go/src/database/sql/driver/types.go:137
var String stringType

type stringType struct{}

func (stringType) ConvertValue(v any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:141
	_go_fuzz_dep_.CoverTab[179245]++
								switch v.(type) {
	case string, []byte:
//line /usr/local/go/src/database/sql/driver/types.go:143
		_go_fuzz_dep_.CoverTab[179247]++
									return v, nil
//line /usr/local/go/src/database/sql/driver/types.go:144
		// _ = "end of CoverTab[179247]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:145
	// _ = "end of CoverTab[179245]"
//line /usr/local/go/src/database/sql/driver/types.go:145
	_go_fuzz_dep_.CoverTab[179246]++
								return fmt.Sprintf("%v", v), nil
//line /usr/local/go/src/database/sql/driver/types.go:146
	// _ = "end of CoverTab[179246]"
}

// Null is a type that implements ValueConverter by allowing nil
//line /usr/local/go/src/database/sql/driver/types.go:149
// values but otherwise delegating to another ValueConverter.
//line /usr/local/go/src/database/sql/driver/types.go:151
type Null struct {
	Converter ValueConverter
}

func (n Null) ConvertValue(v any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:155
	_go_fuzz_dep_.CoverTab[179248]++
								if v == nil {
//line /usr/local/go/src/database/sql/driver/types.go:156
		_go_fuzz_dep_.CoverTab[179250]++
									return nil, nil
//line /usr/local/go/src/database/sql/driver/types.go:157
		// _ = "end of CoverTab[179250]"
	} else {
//line /usr/local/go/src/database/sql/driver/types.go:158
		_go_fuzz_dep_.CoverTab[179251]++
//line /usr/local/go/src/database/sql/driver/types.go:158
		// _ = "end of CoverTab[179251]"
//line /usr/local/go/src/database/sql/driver/types.go:158
	}
//line /usr/local/go/src/database/sql/driver/types.go:158
	// _ = "end of CoverTab[179248]"
//line /usr/local/go/src/database/sql/driver/types.go:158
	_go_fuzz_dep_.CoverTab[179249]++
								return n.Converter.ConvertValue(v)
//line /usr/local/go/src/database/sql/driver/types.go:159
	// _ = "end of CoverTab[179249]"
}

// NotNull is a type that implements ValueConverter by disallowing nil
//line /usr/local/go/src/database/sql/driver/types.go:162
// values but otherwise delegating to another ValueConverter.
//line /usr/local/go/src/database/sql/driver/types.go:164
type NotNull struct {
	Converter ValueConverter
}

func (n NotNull) ConvertValue(v any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:168
	_go_fuzz_dep_.CoverTab[179252]++
								if v == nil {
//line /usr/local/go/src/database/sql/driver/types.go:169
		_go_fuzz_dep_.CoverTab[179254]++
									return nil, fmt.Errorf("nil value not allowed")
//line /usr/local/go/src/database/sql/driver/types.go:170
		// _ = "end of CoverTab[179254]"
	} else {
//line /usr/local/go/src/database/sql/driver/types.go:171
		_go_fuzz_dep_.CoverTab[179255]++
//line /usr/local/go/src/database/sql/driver/types.go:171
		// _ = "end of CoverTab[179255]"
//line /usr/local/go/src/database/sql/driver/types.go:171
	}
//line /usr/local/go/src/database/sql/driver/types.go:171
	// _ = "end of CoverTab[179252]"
//line /usr/local/go/src/database/sql/driver/types.go:171
	_go_fuzz_dep_.CoverTab[179253]++
								return n.Converter.ConvertValue(v)
//line /usr/local/go/src/database/sql/driver/types.go:172
	// _ = "end of CoverTab[179253]"
}

// IsValue reports whether v is a valid Value parameter type.
func IsValue(v any) bool {
//line /usr/local/go/src/database/sql/driver/types.go:176
	_go_fuzz_dep_.CoverTab[179256]++
								if v == nil {
//line /usr/local/go/src/database/sql/driver/types.go:177
		_go_fuzz_dep_.CoverTab[179259]++
									return true
//line /usr/local/go/src/database/sql/driver/types.go:178
		// _ = "end of CoverTab[179259]"
	} else {
//line /usr/local/go/src/database/sql/driver/types.go:179
		_go_fuzz_dep_.CoverTab[179260]++
//line /usr/local/go/src/database/sql/driver/types.go:179
		// _ = "end of CoverTab[179260]"
//line /usr/local/go/src/database/sql/driver/types.go:179
	}
//line /usr/local/go/src/database/sql/driver/types.go:179
	// _ = "end of CoverTab[179256]"
//line /usr/local/go/src/database/sql/driver/types.go:179
	_go_fuzz_dep_.CoverTab[179257]++
								switch v.(type) {
	case []byte, bool, float64, int64, string, time.Time:
//line /usr/local/go/src/database/sql/driver/types.go:181
		_go_fuzz_dep_.CoverTab[179261]++
									return true
//line /usr/local/go/src/database/sql/driver/types.go:182
		// _ = "end of CoverTab[179261]"
	case decimalDecompose:
//line /usr/local/go/src/database/sql/driver/types.go:183
		_go_fuzz_dep_.CoverTab[179262]++
									return true
//line /usr/local/go/src/database/sql/driver/types.go:184
		// _ = "end of CoverTab[179262]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:185
	// _ = "end of CoverTab[179257]"
//line /usr/local/go/src/database/sql/driver/types.go:185
	_go_fuzz_dep_.CoverTab[179258]++
								return false
//line /usr/local/go/src/database/sql/driver/types.go:186
	// _ = "end of CoverTab[179258]"
}

// IsScanValue is equivalent to IsValue.
//line /usr/local/go/src/database/sql/driver/types.go:189
// It exists for compatibility.
//line /usr/local/go/src/database/sql/driver/types.go:191
func IsScanValue(v any) bool {
//line /usr/local/go/src/database/sql/driver/types.go:191
	_go_fuzz_dep_.CoverTab[179263]++
								return IsValue(v)
//line /usr/local/go/src/database/sql/driver/types.go:192
	// _ = "end of CoverTab[179263]"
}

// DefaultParameterConverter is the default implementation of
//line /usr/local/go/src/database/sql/driver/types.go:195
// ValueConverter that's used when a Stmt doesn't implement
//line /usr/local/go/src/database/sql/driver/types.go:195
// ColumnConverter.
//line /usr/local/go/src/database/sql/driver/types.go:195
//
//line /usr/local/go/src/database/sql/driver/types.go:195
// DefaultParameterConverter returns its argument directly if
//line /usr/local/go/src/database/sql/driver/types.go:195
// IsValue(arg). Otherwise, if the argument implements Valuer, its
//line /usr/local/go/src/database/sql/driver/types.go:195
// Value method is used to return a Value. As a fallback, the provided
//line /usr/local/go/src/database/sql/driver/types.go:195
// argument's underlying type is used to convert it to a Value:
//line /usr/local/go/src/database/sql/driver/types.go:195
// underlying integer types are converted to int64, floats to float64,
//line /usr/local/go/src/database/sql/driver/types.go:195
// bool, string, and []byte to themselves. If the argument is a nil
//line /usr/local/go/src/database/sql/driver/types.go:195
// pointer, ConvertValue returns a nil Value. If the argument is a
//line /usr/local/go/src/database/sql/driver/types.go:195
// non-nil pointer, it is dereferenced and ConvertValue is called
//line /usr/local/go/src/database/sql/driver/types.go:195
// recursively. Other types are an error.
//line /usr/local/go/src/database/sql/driver/types.go:208
var DefaultParameterConverter defaultConverter

type defaultConverter struct{}

var _ ValueConverter = defaultConverter{}

var valuerReflectType = reflect.TypeOf((*Valuer)(nil)).Elem()

// callValuerValue returns vr.Value(), with one exception:
//line /usr/local/go/src/database/sql/driver/types.go:216
// If vr.Value is an auto-generated method on a pointer type and the
//line /usr/local/go/src/database/sql/driver/types.go:216
// pointer is nil, it would panic at runtime in the panicwrap
//line /usr/local/go/src/database/sql/driver/types.go:216
// method. Treat it like nil instead.
//line /usr/local/go/src/database/sql/driver/types.go:216
// Issue 8415.
//line /usr/local/go/src/database/sql/driver/types.go:216
//
//line /usr/local/go/src/database/sql/driver/types.go:216
// This is so people can implement driver.Value on value types and
//line /usr/local/go/src/database/sql/driver/types.go:216
// still use nil pointers to those types to mean nil/NULL, just like
//line /usr/local/go/src/database/sql/driver/types.go:216
// string/*string.
//line /usr/local/go/src/database/sql/driver/types.go:216
//
//line /usr/local/go/src/database/sql/driver/types.go:216
// This function is mirrored in the database/sql package.
//line /usr/local/go/src/database/sql/driver/types.go:227
func callValuerValue(vr Valuer) (v Value, err error) {
//line /usr/local/go/src/database/sql/driver/types.go:227
	_go_fuzz_dep_.CoverTab[179264]++
								if rv := reflect.ValueOf(vr); rv.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/database/sql/driver/types.go:228
		_go_fuzz_dep_.CoverTab[179266]++
//line /usr/local/go/src/database/sql/driver/types.go:228
		return rv.IsNil()
									// _ = "end of CoverTab[179266]"
//line /usr/local/go/src/database/sql/driver/types.go:229
	}() && func() bool {
//line /usr/local/go/src/database/sql/driver/types.go:229
		_go_fuzz_dep_.CoverTab[179267]++
//line /usr/local/go/src/database/sql/driver/types.go:229
		return rv.Type().Elem().Implements(valuerReflectType)
									// _ = "end of CoverTab[179267]"
//line /usr/local/go/src/database/sql/driver/types.go:230
	}() {
//line /usr/local/go/src/database/sql/driver/types.go:230
		_go_fuzz_dep_.CoverTab[179268]++
									return nil, nil
//line /usr/local/go/src/database/sql/driver/types.go:231
		// _ = "end of CoverTab[179268]"
	} else {
//line /usr/local/go/src/database/sql/driver/types.go:232
		_go_fuzz_dep_.CoverTab[179269]++
//line /usr/local/go/src/database/sql/driver/types.go:232
		// _ = "end of CoverTab[179269]"
//line /usr/local/go/src/database/sql/driver/types.go:232
	}
//line /usr/local/go/src/database/sql/driver/types.go:232
	// _ = "end of CoverTab[179264]"
//line /usr/local/go/src/database/sql/driver/types.go:232
	_go_fuzz_dep_.CoverTab[179265]++
								return vr.Value()
//line /usr/local/go/src/database/sql/driver/types.go:233
	// _ = "end of CoverTab[179265]"
}

func (defaultConverter) ConvertValue(v any) (Value, error) {
//line /usr/local/go/src/database/sql/driver/types.go:236
	_go_fuzz_dep_.CoverTab[179270]++
								if IsValue(v) {
//line /usr/local/go/src/database/sql/driver/types.go:237
		_go_fuzz_dep_.CoverTab[179274]++
									return v, nil
//line /usr/local/go/src/database/sql/driver/types.go:238
		// _ = "end of CoverTab[179274]"
	} else {
//line /usr/local/go/src/database/sql/driver/types.go:239
		_go_fuzz_dep_.CoverTab[179275]++
//line /usr/local/go/src/database/sql/driver/types.go:239
		// _ = "end of CoverTab[179275]"
//line /usr/local/go/src/database/sql/driver/types.go:239
	}
//line /usr/local/go/src/database/sql/driver/types.go:239
	// _ = "end of CoverTab[179270]"
//line /usr/local/go/src/database/sql/driver/types.go:239
	_go_fuzz_dep_.CoverTab[179271]++

								switch vr := v.(type) {
	case Valuer:
//line /usr/local/go/src/database/sql/driver/types.go:242
		_go_fuzz_dep_.CoverTab[179276]++
									sv, err := callValuerValue(vr)
									if err != nil {
//line /usr/local/go/src/database/sql/driver/types.go:244
			_go_fuzz_dep_.CoverTab[179280]++
										return nil, err
//line /usr/local/go/src/database/sql/driver/types.go:245
			// _ = "end of CoverTab[179280]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:246
			_go_fuzz_dep_.CoverTab[179281]++
//line /usr/local/go/src/database/sql/driver/types.go:246
			// _ = "end of CoverTab[179281]"
//line /usr/local/go/src/database/sql/driver/types.go:246
		}
//line /usr/local/go/src/database/sql/driver/types.go:246
		// _ = "end of CoverTab[179276]"
//line /usr/local/go/src/database/sql/driver/types.go:246
		_go_fuzz_dep_.CoverTab[179277]++
									if !IsValue(sv) {
//line /usr/local/go/src/database/sql/driver/types.go:247
			_go_fuzz_dep_.CoverTab[179282]++
										return nil, fmt.Errorf("non-Value type %T returned from Value", sv)
//line /usr/local/go/src/database/sql/driver/types.go:248
			// _ = "end of CoverTab[179282]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:249
			_go_fuzz_dep_.CoverTab[179283]++
//line /usr/local/go/src/database/sql/driver/types.go:249
			// _ = "end of CoverTab[179283]"
//line /usr/local/go/src/database/sql/driver/types.go:249
		}
//line /usr/local/go/src/database/sql/driver/types.go:249
		// _ = "end of CoverTab[179277]"
//line /usr/local/go/src/database/sql/driver/types.go:249
		_go_fuzz_dep_.CoverTab[179278]++
									return sv, nil
//line /usr/local/go/src/database/sql/driver/types.go:250
		// _ = "end of CoverTab[179278]"

//line /usr/local/go/src/database/sql/driver/types.go:253
	case decimalDecompose:
//line /usr/local/go/src/database/sql/driver/types.go:253
		_go_fuzz_dep_.CoverTab[179279]++
									return vr, nil
//line /usr/local/go/src/database/sql/driver/types.go:254
		// _ = "end of CoverTab[179279]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:255
	// _ = "end of CoverTab[179271]"
//line /usr/local/go/src/database/sql/driver/types.go:255
	_go_fuzz_dep_.CoverTab[179272]++

								rv := reflect.ValueOf(v)
								switch rv.Kind() {
	case reflect.Pointer:
//line /usr/local/go/src/database/sql/driver/types.go:259
		_go_fuzz_dep_.CoverTab[179284]++

									if rv.IsNil() {
//line /usr/local/go/src/database/sql/driver/types.go:261
			_go_fuzz_dep_.CoverTab[179295]++
										return nil, nil
//line /usr/local/go/src/database/sql/driver/types.go:262
			// _ = "end of CoverTab[179295]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:263
			_go_fuzz_dep_.CoverTab[179296]++
										return defaultConverter{}.ConvertValue(rv.Elem().Interface())
//line /usr/local/go/src/database/sql/driver/types.go:264
			// _ = "end of CoverTab[179296]"
		}
//line /usr/local/go/src/database/sql/driver/types.go:265
		// _ = "end of CoverTab[179284]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/database/sql/driver/types.go:266
		_go_fuzz_dep_.CoverTab[179285]++
									return rv.Int(), nil
//line /usr/local/go/src/database/sql/driver/types.go:267
		// _ = "end of CoverTab[179285]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
//line /usr/local/go/src/database/sql/driver/types.go:268
		_go_fuzz_dep_.CoverTab[179286]++
									return int64(rv.Uint()), nil
//line /usr/local/go/src/database/sql/driver/types.go:269
		// _ = "end of CoverTab[179286]"
	case reflect.Uint64:
//line /usr/local/go/src/database/sql/driver/types.go:270
		_go_fuzz_dep_.CoverTab[179287]++
									u64 := rv.Uint()
									if u64 >= 1<<63 {
//line /usr/local/go/src/database/sql/driver/types.go:272
			_go_fuzz_dep_.CoverTab[179297]++
										return nil, fmt.Errorf("uint64 values with high bit set are not supported")
//line /usr/local/go/src/database/sql/driver/types.go:273
			// _ = "end of CoverTab[179297]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:274
			_go_fuzz_dep_.CoverTab[179298]++
//line /usr/local/go/src/database/sql/driver/types.go:274
			// _ = "end of CoverTab[179298]"
//line /usr/local/go/src/database/sql/driver/types.go:274
		}
//line /usr/local/go/src/database/sql/driver/types.go:274
		// _ = "end of CoverTab[179287]"
//line /usr/local/go/src/database/sql/driver/types.go:274
		_go_fuzz_dep_.CoverTab[179288]++
									return int64(u64), nil
//line /usr/local/go/src/database/sql/driver/types.go:275
		// _ = "end of CoverTab[179288]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/database/sql/driver/types.go:276
		_go_fuzz_dep_.CoverTab[179289]++
									return rv.Float(), nil
//line /usr/local/go/src/database/sql/driver/types.go:277
		// _ = "end of CoverTab[179289]"
	case reflect.Bool:
//line /usr/local/go/src/database/sql/driver/types.go:278
		_go_fuzz_dep_.CoverTab[179290]++
									return rv.Bool(), nil
//line /usr/local/go/src/database/sql/driver/types.go:279
		// _ = "end of CoverTab[179290]"
	case reflect.Slice:
//line /usr/local/go/src/database/sql/driver/types.go:280
		_go_fuzz_dep_.CoverTab[179291]++
									ek := rv.Type().Elem().Kind()
									if ek == reflect.Uint8 {
//line /usr/local/go/src/database/sql/driver/types.go:282
			_go_fuzz_dep_.CoverTab[179299]++
										return rv.Bytes(), nil
//line /usr/local/go/src/database/sql/driver/types.go:283
			// _ = "end of CoverTab[179299]"
		} else {
//line /usr/local/go/src/database/sql/driver/types.go:284
			_go_fuzz_dep_.CoverTab[179300]++
//line /usr/local/go/src/database/sql/driver/types.go:284
			// _ = "end of CoverTab[179300]"
//line /usr/local/go/src/database/sql/driver/types.go:284
		}
//line /usr/local/go/src/database/sql/driver/types.go:284
		// _ = "end of CoverTab[179291]"
//line /usr/local/go/src/database/sql/driver/types.go:284
		_go_fuzz_dep_.CoverTab[179292]++
									return nil, fmt.Errorf("unsupported type %T, a slice of %s", v, ek)
//line /usr/local/go/src/database/sql/driver/types.go:285
		// _ = "end of CoverTab[179292]"
	case reflect.String:
//line /usr/local/go/src/database/sql/driver/types.go:286
		_go_fuzz_dep_.CoverTab[179293]++
									return rv.String(), nil
//line /usr/local/go/src/database/sql/driver/types.go:287
		// _ = "end of CoverTab[179293]"
//line /usr/local/go/src/database/sql/driver/types.go:287
	default:
//line /usr/local/go/src/database/sql/driver/types.go:287
		_go_fuzz_dep_.CoverTab[179294]++
//line /usr/local/go/src/database/sql/driver/types.go:287
		// _ = "end of CoverTab[179294]"
	}
//line /usr/local/go/src/database/sql/driver/types.go:288
	// _ = "end of CoverTab[179272]"
//line /usr/local/go/src/database/sql/driver/types.go:288
	_go_fuzz_dep_.CoverTab[179273]++
								return nil, fmt.Errorf("unsupported type %T, a %s", v, rv.Kind())
//line /usr/local/go/src/database/sql/driver/types.go:289
	// _ = "end of CoverTab[179273]"
}

type decimalDecompose interface {
	// Decompose returns the internal decimal state into parts.
	// If the provided buf has sufficient capacity, buf may be returned as the coefficient with
	// the value set and length set as appropriate.
	Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32)
}

//line /usr/local/go/src/database/sql/driver/types.go:297
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/database/sql/driver/types.go:297
var _ = _go_fuzz_dep_.CoverTab
