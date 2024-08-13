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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:15
)

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

// NameMapper represents a ini tag name mapper.
type NameMapper func(string) string

// Built-in name getters.
var (
	// SnackCase converts to format SNACK_CASE.
	SnackCase	NameMapper	= func(raw string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:33
		_go_fuzz_dep_.CoverTab[129158]++
										newstr := make([]rune, 0, len(raw))
										for i, chr := range raw {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:35
			_go_fuzz_dep_.CoverTab[129160]++
											if isUpper := 'A' <= chr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:36
				_go_fuzz_dep_.CoverTab[129162]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:36
				return chr <= 'Z'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:36
				// _ = "end of CoverTab[129162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:36
			}(); isUpper {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:36
				_go_fuzz_dep_.CoverTab[129163]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:37
					_go_fuzz_dep_.CoverTab[129164]++
													newstr = append(newstr, '_')
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:38
					// _ = "end of CoverTab[129164]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:39
					_go_fuzz_dep_.CoverTab[129165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:39
					// _ = "end of CoverTab[129165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:39
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:39
				// _ = "end of CoverTab[129163]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:40
				_go_fuzz_dep_.CoverTab[129166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:40
				// _ = "end of CoverTab[129166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:40
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:40
			// _ = "end of CoverTab[129160]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:40
			_go_fuzz_dep_.CoverTab[129161]++
											newstr = append(newstr, unicode.ToUpper(chr))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:41
			// _ = "end of CoverTab[129161]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:42
		// _ = "end of CoverTab[129158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:42
		_go_fuzz_dep_.CoverTab[129159]++
										return string(newstr)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:43
		// _ = "end of CoverTab[129159]"
	}
	// TitleUnderscore converts to format title_underscore.
	TitleUnderscore	NameMapper	= func(raw string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:46
		_go_fuzz_dep_.CoverTab[129167]++
										newstr := make([]rune, 0, len(raw))
										for i, chr := range raw {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:48
			_go_fuzz_dep_.CoverTab[129169]++
											if isUpper := 'A' <= chr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:49
				_go_fuzz_dep_.CoverTab[129171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:49
				return chr <= 'Z'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:49
				// _ = "end of CoverTab[129171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:49
			}(); isUpper {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:49
				_go_fuzz_dep_.CoverTab[129172]++
												if i > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:50
					_go_fuzz_dep_.CoverTab[129174]++
													newstr = append(newstr, '_')
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:51
					// _ = "end of CoverTab[129174]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:52
					_go_fuzz_dep_.CoverTab[129175]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:52
					// _ = "end of CoverTab[129175]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:52
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:52
				// _ = "end of CoverTab[129172]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:52
				_go_fuzz_dep_.CoverTab[129173]++
												chr -= 'A' - 'a'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:53
				// _ = "end of CoverTab[129173]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:54
				_go_fuzz_dep_.CoverTab[129176]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:54
				// _ = "end of CoverTab[129176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:54
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:54
			// _ = "end of CoverTab[129169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:54
			_go_fuzz_dep_.CoverTab[129170]++
											newstr = append(newstr, chr)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:55
			// _ = "end of CoverTab[129170]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:56
		// _ = "end of CoverTab[129167]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:56
		_go_fuzz_dep_.CoverTab[129168]++
										return string(newstr)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:57
		// _ = "end of CoverTab[129168]"
	}
)

func (s *Section) parseFieldName(raw, actual string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:61
	_go_fuzz_dep_.CoverTab[129177]++
									if len(actual) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:62
		_go_fuzz_dep_.CoverTab[129180]++
										return actual
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:63
		// _ = "end of CoverTab[129180]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:64
		_go_fuzz_dep_.CoverTab[129181]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:64
		// _ = "end of CoverTab[129181]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:64
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:64
	// _ = "end of CoverTab[129177]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:64
	_go_fuzz_dep_.CoverTab[129178]++
									if s.f.NameMapper != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:65
		_go_fuzz_dep_.CoverTab[129182]++
										return s.f.NameMapper(raw)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:66
		// _ = "end of CoverTab[129182]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:67
		_go_fuzz_dep_.CoverTab[129183]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:67
		// _ = "end of CoverTab[129183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:67
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:67
	// _ = "end of CoverTab[129178]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:67
	_go_fuzz_dep_.CoverTab[129179]++
									return raw
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:68
	// _ = "end of CoverTab[129179]"
}

func parseDelim(actual string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:71
	_go_fuzz_dep_.CoverTab[129184]++
									if len(actual) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:72
		_go_fuzz_dep_.CoverTab[129186]++
										return actual
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:73
		// _ = "end of CoverTab[129186]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:74
		_go_fuzz_dep_.CoverTab[129187]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:74
		// _ = "end of CoverTab[129187]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:74
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:74
	// _ = "end of CoverTab[129184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:74
	_go_fuzz_dep_.CoverTab[129185]++
									return ","
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:75
	// _ = "end of CoverTab[129185]"
}

var reflectTime = reflect.TypeOf(time.Now()).Kind()

// setSliceWithProperType sets proper values to slice based on its type.
func setSliceWithProperType(key *Key, field reflect.Value, delim string, allowShadow, isStrict bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:81
	_go_fuzz_dep_.CoverTab[129188]++
									var strs []string
									if allowShadow {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:83
		_go_fuzz_dep_.CoverTab[129194]++
										strs = key.StringsWithShadows(delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:84
		// _ = "end of CoverTab[129194]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:85
		_go_fuzz_dep_.CoverTab[129195]++
										strs = key.Strings(delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:86
		// _ = "end of CoverTab[129195]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:87
	// _ = "end of CoverTab[129188]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:87
	_go_fuzz_dep_.CoverTab[129189]++

									numVals := len(strs)
									if numVals == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:90
		_go_fuzz_dep_.CoverTab[129196]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:91
		// _ = "end of CoverTab[129196]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:92
		_go_fuzz_dep_.CoverTab[129197]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:92
		// _ = "end of CoverTab[129197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:92
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:92
	// _ = "end of CoverTab[129189]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:92
	_go_fuzz_dep_.CoverTab[129190]++

									var vals interface{}
									var err error

									sliceOf := field.Type().Elem().Kind()
									switch sliceOf {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:99
			_go_fuzz_dep_.CoverTab[129198]++
											vals = strs
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:100
		// _ = "end of CoverTab[129198]"
	case reflect.Int:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:101
		_go_fuzz_dep_.CoverTab[129199]++
											vals, err = key.parseInts(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:102
		// _ = "end of CoverTab[129199]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:103
		_go_fuzz_dep_.CoverTab[129200]++
											vals, err = key.parseInt64s(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:104
		// _ = "end of CoverTab[129200]"
	case reflect.Uint:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:105
		_go_fuzz_dep_.CoverTab[129201]++
											vals, err = key.parseUints(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:106
		// _ = "end of CoverTab[129201]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:107
		_go_fuzz_dep_.CoverTab[129202]++
											vals, err = key.parseUint64s(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:108
		// _ = "end of CoverTab[129202]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:109
		_go_fuzz_dep_.CoverTab[129203]++
											vals, err = key.parseFloat64s(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:110
		// _ = "end of CoverTab[129203]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:111
		_go_fuzz_dep_.CoverTab[129204]++
											vals, err = key.parseBools(strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:112
		// _ = "end of CoverTab[129204]"
	case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:113
		_go_fuzz_dep_.CoverTab[129205]++
											vals, err = key.parseTimesFormat(time.RFC3339, strs, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:114
		// _ = "end of CoverTab[129205]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:115
		_go_fuzz_dep_.CoverTab[129206]++
											return fmt.Errorf("unsupported type '[]%s'", sliceOf)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:116
		// _ = "end of CoverTab[129206]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:117
	// _ = "end of CoverTab[129190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:117
	_go_fuzz_dep_.CoverTab[129191]++
										if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:118
		_go_fuzz_dep_.CoverTab[129207]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:118
		return isStrict
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:118
		// _ = "end of CoverTab[129207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:118
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:118
		_go_fuzz_dep_.CoverTab[129208]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:119
		// _ = "end of CoverTab[129208]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:120
		_go_fuzz_dep_.CoverTab[129209]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:120
		// _ = "end of CoverTab[129209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:120
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:120
	// _ = "end of CoverTab[129191]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:120
	_go_fuzz_dep_.CoverTab[129192]++

										slice := reflect.MakeSlice(field.Type(), numVals, numVals)
										for i := 0; i < numVals; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:123
		_go_fuzz_dep_.CoverTab[129210]++
											switch sliceOf {
		case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:125
			_go_fuzz_dep_.CoverTab[129211]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]string)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:126
			// _ = "end of CoverTab[129211]"
		case reflect.Int:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:127
			_go_fuzz_dep_.CoverTab[129212]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]int)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:128
			// _ = "end of CoverTab[129212]"
		case reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:129
			_go_fuzz_dep_.CoverTab[129213]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]int64)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:130
			// _ = "end of CoverTab[129213]"
		case reflect.Uint:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:131
			_go_fuzz_dep_.CoverTab[129214]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]uint)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:132
			// _ = "end of CoverTab[129214]"
		case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:133
			_go_fuzz_dep_.CoverTab[129215]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]uint64)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:134
			// _ = "end of CoverTab[129215]"
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:135
			_go_fuzz_dep_.CoverTab[129216]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]float64)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:136
			// _ = "end of CoverTab[129216]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:137
			_go_fuzz_dep_.CoverTab[129217]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]bool)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:138
			// _ = "end of CoverTab[129217]"
		case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:139
			_go_fuzz_dep_.CoverTab[129218]++
												slice.Index(i).Set(reflect.ValueOf(vals.([]time.Time)[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:140
			// _ = "end of CoverTab[129218]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:140
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:140
			_go_fuzz_dep_.CoverTab[129219]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:140
			// _ = "end of CoverTab[129219]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:141
		// _ = "end of CoverTab[129210]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:142
	// _ = "end of CoverTab[129192]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:142
	_go_fuzz_dep_.CoverTab[129193]++
										field.Set(slice)
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:144
	// _ = "end of CoverTab[129193]"
}

func wrapStrictError(err error, isStrict bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:147
	_go_fuzz_dep_.CoverTab[129220]++
										if isStrict {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:148
		_go_fuzz_dep_.CoverTab[129222]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:149
		// _ = "end of CoverTab[129222]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:150
		_go_fuzz_dep_.CoverTab[129223]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:150
		// _ = "end of CoverTab[129223]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:150
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:150
	// _ = "end of CoverTab[129220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:150
	_go_fuzz_dep_.CoverTab[129221]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:151
	// _ = "end of CoverTab[129221]"
}

// setWithProperType sets proper value to field based on its type,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:154
// but it does not return error for failing parsing,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:154
// because we want to use default value that is already assigned to struct.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:157
func setWithProperType(t reflect.Type, key *Key, field reflect.Value, delim string, allowShadow, isStrict bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:157
	_go_fuzz_dep_.CoverTab[129224]++
										vt := t
										isPtr := t.Kind() == reflect.Ptr
										if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:160
		_go_fuzz_dep_.CoverTab[129227]++
											vt = t.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:161
		// _ = "end of CoverTab[129227]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:162
		_go_fuzz_dep_.CoverTab[129228]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:162
		// _ = "end of CoverTab[129228]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:162
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:162
	// _ = "end of CoverTab[129224]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:162
	_go_fuzz_dep_.CoverTab[129225]++
										switch vt.Kind() {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:164
		_go_fuzz_dep_.CoverTab[129229]++
											stringVal := key.String()
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:166
			_go_fuzz_dep_.CoverTab[129244]++
												field.Set(reflect.ValueOf(&stringVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:167
			// _ = "end of CoverTab[129244]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:168
			_go_fuzz_dep_.CoverTab[129245]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:168
			if len(stringVal) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:168
				_go_fuzz_dep_.CoverTab[129246]++
													field.SetString(key.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:169
				// _ = "end of CoverTab[129246]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
				_go_fuzz_dep_.CoverTab[129247]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
				// _ = "end of CoverTab[129247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
			// _ = "end of CoverTab[129245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:170
		// _ = "end of CoverTab[129229]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:171
		_go_fuzz_dep_.CoverTab[129230]++
											boolVal, err := key.Bool()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:173
			_go_fuzz_dep_.CoverTab[129248]++
												return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:174
			// _ = "end of CoverTab[129248]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:175
			_go_fuzz_dep_.CoverTab[129249]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:175
			// _ = "end of CoverTab[129249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:175
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:175
		// _ = "end of CoverTab[129230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:175
		_go_fuzz_dep_.CoverTab[129231]++
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:176
			_go_fuzz_dep_.CoverTab[129250]++
												field.Set(reflect.ValueOf(&boolVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:177
			// _ = "end of CoverTab[129250]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:178
			_go_fuzz_dep_.CoverTab[129251]++
												field.SetBool(boolVal)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:179
			// _ = "end of CoverTab[129251]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:180
		// _ = "end of CoverTab[129231]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:181
		_go_fuzz_dep_.CoverTab[129232]++

											if vt.Name() == "Duration" {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:183
			_go_fuzz_dep_.CoverTab[129252]++
												durationVal, err := key.Duration()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:185
				_go_fuzz_dep_.CoverTab[129255]++
													if intVal, err := key.Int64(); err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:186
					_go_fuzz_dep_.CoverTab[129257]++
														field.SetInt(intVal)
														return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:188
					// _ = "end of CoverTab[129257]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:189
					_go_fuzz_dep_.CoverTab[129258]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:189
					// _ = "end of CoverTab[129258]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:189
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:189
				// _ = "end of CoverTab[129255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:189
				_go_fuzz_dep_.CoverTab[129256]++
													return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:190
				// _ = "end of CoverTab[129256]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:191
				_go_fuzz_dep_.CoverTab[129259]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:191
				// _ = "end of CoverTab[129259]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:191
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:191
			// _ = "end of CoverTab[129252]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:191
			_go_fuzz_dep_.CoverTab[129253]++
												if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:192
				_go_fuzz_dep_.CoverTab[129260]++
													field.Set(reflect.ValueOf(&durationVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:193
				// _ = "end of CoverTab[129260]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:194
				_go_fuzz_dep_.CoverTab[129261]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:194
				if int64(durationVal) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:194
					_go_fuzz_dep_.CoverTab[129262]++
														field.Set(reflect.ValueOf(durationVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:195
					// _ = "end of CoverTab[129262]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
					_go_fuzz_dep_.CoverTab[129263]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
					// _ = "end of CoverTab[129263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
				// _ = "end of CoverTab[129261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
			// _ = "end of CoverTab[129253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:196
			_go_fuzz_dep_.CoverTab[129254]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:197
			// _ = "end of CoverTab[129254]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:198
			_go_fuzz_dep_.CoverTab[129264]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:198
			// _ = "end of CoverTab[129264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:198
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:198
		// _ = "end of CoverTab[129232]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:198
		_go_fuzz_dep_.CoverTab[129233]++

											intVal, err := key.Int64()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:201
			_go_fuzz_dep_.CoverTab[129265]++
												return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:202
			// _ = "end of CoverTab[129265]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:203
			_go_fuzz_dep_.CoverTab[129266]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:203
			// _ = "end of CoverTab[129266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:203
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:203
		// _ = "end of CoverTab[129233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:203
		_go_fuzz_dep_.CoverTab[129234]++
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:204
			_go_fuzz_dep_.CoverTab[129267]++
												pv := reflect.New(t.Elem())
												pv.Elem().SetInt(intVal)
												field.Set(pv)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:207
			// _ = "end of CoverTab[129267]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:208
			_go_fuzz_dep_.CoverTab[129268]++
												field.SetInt(intVal)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:209
			// _ = "end of CoverTab[129268]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:210
		// _ = "end of CoverTab[129234]"

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:212
		_go_fuzz_dep_.CoverTab[129235]++
											durationVal, err := key.Duration()

											if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:215
			_go_fuzz_dep_.CoverTab[129269]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:215
			return uint64(durationVal) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:215
			// _ = "end of CoverTab[129269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:215
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:215
			_go_fuzz_dep_.CoverTab[129270]++
												if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:216
				_go_fuzz_dep_.CoverTab[129272]++
													field.Set(reflect.ValueOf(&durationVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:217
				// _ = "end of CoverTab[129272]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:218
				_go_fuzz_dep_.CoverTab[129273]++
													field.Set(reflect.ValueOf(durationVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:219
				// _ = "end of CoverTab[129273]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:220
			// _ = "end of CoverTab[129270]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:220
			_go_fuzz_dep_.CoverTab[129271]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:221
			// _ = "end of CoverTab[129271]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:222
			_go_fuzz_dep_.CoverTab[129274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:222
			// _ = "end of CoverTab[129274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:222
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:222
		// _ = "end of CoverTab[129235]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:222
		_go_fuzz_dep_.CoverTab[129236]++

											uintVal, err := key.Uint64()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:225
			_go_fuzz_dep_.CoverTab[129275]++
												return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:226
			// _ = "end of CoverTab[129275]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:227
			_go_fuzz_dep_.CoverTab[129276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:227
			// _ = "end of CoverTab[129276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:227
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:227
		// _ = "end of CoverTab[129236]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:227
		_go_fuzz_dep_.CoverTab[129237]++
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:228
			_go_fuzz_dep_.CoverTab[129277]++
												pv := reflect.New(t.Elem())
												pv.Elem().SetUint(uintVal)
												field.Set(pv)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:231
			// _ = "end of CoverTab[129277]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:232
			_go_fuzz_dep_.CoverTab[129278]++
												field.SetUint(uintVal)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:233
			// _ = "end of CoverTab[129278]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:234
		// _ = "end of CoverTab[129237]"

	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:236
		_go_fuzz_dep_.CoverTab[129238]++
											floatVal, err := key.Float64()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:238
			_go_fuzz_dep_.CoverTab[129279]++
												return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:239
			// _ = "end of CoverTab[129279]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:240
			_go_fuzz_dep_.CoverTab[129280]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:240
			// _ = "end of CoverTab[129280]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:240
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:240
		// _ = "end of CoverTab[129238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:240
		_go_fuzz_dep_.CoverTab[129239]++
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:241
			_go_fuzz_dep_.CoverTab[129281]++
												pv := reflect.New(t.Elem())
												pv.Elem().SetFloat(floatVal)
												field.Set(pv)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:244
			// _ = "end of CoverTab[129281]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:245
			_go_fuzz_dep_.CoverTab[129282]++
												field.SetFloat(floatVal)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:246
			// _ = "end of CoverTab[129282]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:247
		// _ = "end of CoverTab[129239]"
	case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:248
		_go_fuzz_dep_.CoverTab[129240]++
											timeVal, err := key.Time()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:250
			_go_fuzz_dep_.CoverTab[129283]++
												return wrapStrictError(err, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:251
			// _ = "end of CoverTab[129283]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:252
			_go_fuzz_dep_.CoverTab[129284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:252
			// _ = "end of CoverTab[129284]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:252
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:252
		// _ = "end of CoverTab[129240]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:252
		_go_fuzz_dep_.CoverTab[129241]++
											if isPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:253
			_go_fuzz_dep_.CoverTab[129285]++
												field.Set(reflect.ValueOf(&timeVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:254
			// _ = "end of CoverTab[129285]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:255
			_go_fuzz_dep_.CoverTab[129286]++
												field.Set(reflect.ValueOf(timeVal))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:256
			// _ = "end of CoverTab[129286]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:257
		// _ = "end of CoverTab[129241]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:258
		_go_fuzz_dep_.CoverTab[129242]++
											return setSliceWithProperType(key, field, delim, allowShadow, isStrict)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:259
		// _ = "end of CoverTab[129242]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:260
		_go_fuzz_dep_.CoverTab[129243]++
											return fmt.Errorf("unsupported type %q", t)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:261
		// _ = "end of CoverTab[129243]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:262
	// _ = "end of CoverTab[129225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:262
	_go_fuzz_dep_.CoverTab[129226]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:263
	// _ = "end of CoverTab[129226]"
}

func parseTagOptions(tag string) (rawName string, omitEmpty bool, allowShadow bool, allowNonUnique bool, extends bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:266
	_go_fuzz_dep_.CoverTab[129287]++
										opts := strings.SplitN(tag, ",", 5)
										rawName = opts[0]
										for _, opt := range opts[1:] {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:269
		_go_fuzz_dep_.CoverTab[129289]++
											omitEmpty = omitEmpty || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:270
			_go_fuzz_dep_.CoverTab[129290]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:270
			return (opt == "omitempty")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:270
			// _ = "end of CoverTab[129290]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:270
		}()
											allowShadow = allowShadow || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:271
			_go_fuzz_dep_.CoverTab[129291]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:271
			return (opt == "allowshadow")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:271
			// _ = "end of CoverTab[129291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:271
		}()
											allowNonUnique = allowNonUnique || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:272
			_go_fuzz_dep_.CoverTab[129292]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:272
			return (opt == "nonunique")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:272
			// _ = "end of CoverTab[129292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:272
		}()
											extends = extends || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:273
			_go_fuzz_dep_.CoverTab[129293]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:273
			return (opt == "extends")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:273
			// _ = "end of CoverTab[129293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:273
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:273
		// _ = "end of CoverTab[129289]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:274
	// _ = "end of CoverTab[129287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:274
	_go_fuzz_dep_.CoverTab[129288]++
										return rawName, omitEmpty, allowShadow, allowNonUnique, extends
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:275
	// _ = "end of CoverTab[129288]"
}

// mapToField maps the given value to the matching field of the given section.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:278
// The sectionIndex is the index (if non unique sections are enabled) to which the value should be added.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:280
func (s *Section) mapToField(val reflect.Value, isStrict bool, sectionIndex int, sectionName string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:280
	_go_fuzz_dep_.CoverTab[129294]++
										if val.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:281
		_go_fuzz_dep_.CoverTab[129297]++
											val = val.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:282
		// _ = "end of CoverTab[129297]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:283
		_go_fuzz_dep_.CoverTab[129298]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:283
		// _ = "end of CoverTab[129298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:283
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:283
	// _ = "end of CoverTab[129294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:283
	_go_fuzz_dep_.CoverTab[129295]++
										typ := val.Type()

										for i := 0; i < typ.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:286
		_go_fuzz_dep_.CoverTab[129299]++
											field := val.Field(i)
											tpField := typ.Field(i)

											tag := tpField.Tag.Get("ini")
											if tag == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:291
			_go_fuzz_dep_.CoverTab[129305]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:292
			// _ = "end of CoverTab[129305]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:293
			_go_fuzz_dep_.CoverTab[129306]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:293
			// _ = "end of CoverTab[129306]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:293
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:293
		// _ = "end of CoverTab[129299]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:293
		_go_fuzz_dep_.CoverTab[129300]++

											rawName, _, allowShadow, allowNonUnique, extends := parseTagOptions(tag)
											fieldName := s.parseFieldName(tpField.Name, rawName)
											if len(fieldName) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:297
			_go_fuzz_dep_.CoverTab[129307]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:297
			return !field.CanSet()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:297
			// _ = "end of CoverTab[129307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:297
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:297
			_go_fuzz_dep_.CoverTab[129308]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:298
			// _ = "end of CoverTab[129308]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:299
			_go_fuzz_dep_.CoverTab[129309]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:299
			// _ = "end of CoverTab[129309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:299
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:299
		// _ = "end of CoverTab[129300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:299
		_go_fuzz_dep_.CoverTab[129301]++

											isStruct := tpField.Type.Kind() == reflect.Struct
											isStructPtr := tpField.Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:302
			_go_fuzz_dep_.CoverTab[129310]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:302
			return tpField.Type.Elem().Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:302
			// _ = "end of CoverTab[129310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:302
		}()
											isAnonymousPtr := tpField.Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:303
			_go_fuzz_dep_.CoverTab[129311]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:303
			return tpField.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:303
			// _ = "end of CoverTab[129311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:303
		}()
											if isAnonymousPtr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:304
			_go_fuzz_dep_.CoverTab[129312]++
												field.Set(reflect.New(tpField.Type.Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:305
			// _ = "end of CoverTab[129312]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:306
			_go_fuzz_dep_.CoverTab[129313]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:306
			// _ = "end of CoverTab[129313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:306
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:306
		// _ = "end of CoverTab[129301]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:306
		_go_fuzz_dep_.CoverTab[129302]++

											if extends && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
			_go_fuzz_dep_.CoverTab[129314]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
			return (isAnonymousPtr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
				_go_fuzz_dep_.CoverTab[129315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
				return (isStruct && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
					_go_fuzz_dep_.CoverTab[129316]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
					return tpField.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
					// _ = "end of CoverTab[129316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
				// _ = "end of CoverTab[129315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
			// _ = "end of CoverTab[129314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:308
			_go_fuzz_dep_.CoverTab[129317]++
												if isStructPtr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:309
				_go_fuzz_dep_.CoverTab[129320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:309
				return field.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:309
				// _ = "end of CoverTab[129320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:309
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:309
				_go_fuzz_dep_.CoverTab[129321]++
													field.Set(reflect.New(tpField.Type.Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:310
				// _ = "end of CoverTab[129321]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:311
				_go_fuzz_dep_.CoverTab[129322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:311
				// _ = "end of CoverTab[129322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:311
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:311
			// _ = "end of CoverTab[129317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:311
			_go_fuzz_dep_.CoverTab[129318]++
												fieldSection := s
												if rawName != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:313
				_go_fuzz_dep_.CoverTab[129323]++
													sectionName = s.name + s.f.options.ChildSectionDelimiter + rawName
													if secs, err := s.f.SectionsByName(sectionName); err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:315
					_go_fuzz_dep_.CoverTab[129324]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:315
					return sectionIndex < len(secs)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:315
					// _ = "end of CoverTab[129324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:315
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:315
					_go_fuzz_dep_.CoverTab[129325]++
														fieldSection = secs[sectionIndex]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:316
					// _ = "end of CoverTab[129325]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:317
					_go_fuzz_dep_.CoverTab[129326]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:317
					// _ = "end of CoverTab[129326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:317
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:317
				// _ = "end of CoverTab[129323]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:318
				_go_fuzz_dep_.CoverTab[129327]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:318
				// _ = "end of CoverTab[129327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:318
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:318
			// _ = "end of CoverTab[129318]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:318
			_go_fuzz_dep_.CoverTab[129319]++
												if err := fieldSection.mapToField(field, isStrict, sectionIndex, sectionName); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:319
				_go_fuzz_dep_.CoverTab[129328]++
													return fmt.Errorf("map to field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:320
				// _ = "end of CoverTab[129328]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:321
				_go_fuzz_dep_.CoverTab[129329]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:321
				// _ = "end of CoverTab[129329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:321
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:321
			// _ = "end of CoverTab[129319]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
			_go_fuzz_dep_.CoverTab[129330]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
			if isAnonymousPtr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				_go_fuzz_dep_.CoverTab[129331]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				return isStruct
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				// _ = "end of CoverTab[129331]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				_go_fuzz_dep_.CoverTab[129332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				return isStructPtr
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				// _ = "end of CoverTab[129332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:322
				_go_fuzz_dep_.CoverTab[129333]++
													if secs, err := s.f.SectionsByName(fieldName); err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:323
					_go_fuzz_dep_.CoverTab[129334]++
														if len(secs) <= sectionIndex {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:324
						_go_fuzz_dep_.CoverTab[129338]++
															return fmt.Errorf("there are not enough sections (%d <= %d) for the field %q", len(secs), sectionIndex, fieldName)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:325
						// _ = "end of CoverTab[129338]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:326
						_go_fuzz_dep_.CoverTab[129339]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:326
						// _ = "end of CoverTab[129339]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:326
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:326
					// _ = "end of CoverTab[129334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:326
					_go_fuzz_dep_.CoverTab[129335]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
					if isStructPtr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
						_go_fuzz_dep_.CoverTab[129340]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
						return field.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
						// _ = "end of CoverTab[129340]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
					}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:329
						_go_fuzz_dep_.CoverTab[129341]++
															field.Set(reflect.New(tpField.Type.Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:330
						// _ = "end of CoverTab[129341]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:331
						_go_fuzz_dep_.CoverTab[129342]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:331
						// _ = "end of CoverTab[129342]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:331
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:331
					// _ = "end of CoverTab[129335]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:331
					_go_fuzz_dep_.CoverTab[129336]++
														if err = secs[sectionIndex].mapToField(field, isStrict, sectionIndex, fieldName); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:332
						_go_fuzz_dep_.CoverTab[129343]++
															return fmt.Errorf("map to field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:333
						// _ = "end of CoverTab[129343]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:334
						_go_fuzz_dep_.CoverTab[129344]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:334
						// _ = "end of CoverTab[129344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:334
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:334
					// _ = "end of CoverTab[129336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:334
					_go_fuzz_dep_.CoverTab[129337]++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:335
					// _ = "end of CoverTab[129337]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:336
					_go_fuzz_dep_.CoverTab[129345]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:336
					// _ = "end of CoverTab[129345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:336
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:336
				// _ = "end of CoverTab[129333]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
				_go_fuzz_dep_.CoverTab[129346]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
				// _ = "end of CoverTab[129346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
			// _ = "end of CoverTab[129330]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
		// _ = "end of CoverTab[129302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:337
		_go_fuzz_dep_.CoverTab[129303]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
		if allowNonUnique && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
			_go_fuzz_dep_.CoverTab[129347]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
			return tpField.Type.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
			// _ = "end of CoverTab[129347]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:340
			_go_fuzz_dep_.CoverTab[129348]++
												newField, err := s.mapToSlice(fieldName, field, isStrict)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:342
				_go_fuzz_dep_.CoverTab[129350]++
													return fmt.Errorf("map to slice %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:343
				// _ = "end of CoverTab[129350]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:344
				_go_fuzz_dep_.CoverTab[129351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:344
				// _ = "end of CoverTab[129351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:344
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:344
			// _ = "end of CoverTab[129348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:344
			_go_fuzz_dep_.CoverTab[129349]++

												field.Set(newField)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:347
			// _ = "end of CoverTab[129349]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:348
			_go_fuzz_dep_.CoverTab[129352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:348
			// _ = "end of CoverTab[129352]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:348
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:348
		// _ = "end of CoverTab[129303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:348
		_go_fuzz_dep_.CoverTab[129304]++

											if key, err := s.GetKey(fieldName); err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:350
			_go_fuzz_dep_.CoverTab[129353]++
												delim := parseDelim(tpField.Tag.Get("delim"))
												if err = setWithProperType(tpField.Type, key, field, delim, allowShadow, isStrict); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:352
				_go_fuzz_dep_.CoverTab[129354]++
													return fmt.Errorf("set field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:353
				// _ = "end of CoverTab[129354]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:354
				_go_fuzz_dep_.CoverTab[129355]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:354
				// _ = "end of CoverTab[129355]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:354
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:354
			// _ = "end of CoverTab[129353]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:355
			_go_fuzz_dep_.CoverTab[129356]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:355
			// _ = "end of CoverTab[129356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:355
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:355
		// _ = "end of CoverTab[129304]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:356
	// _ = "end of CoverTab[129295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:356
	_go_fuzz_dep_.CoverTab[129296]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:357
	// _ = "end of CoverTab[129296]"
}

// mapToSlice maps all sections with the same name and returns the new value.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:360
// The type of the Value must be a slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:362
func (s *Section) mapToSlice(secName string, val reflect.Value, isStrict bool) (reflect.Value, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:362
	_go_fuzz_dep_.CoverTab[129357]++
										secs, err := s.f.SectionsByName(secName)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:364
		_go_fuzz_dep_.CoverTab[129360]++
											return reflect.Value{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:365
		// _ = "end of CoverTab[129360]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:366
		_go_fuzz_dep_.CoverTab[129361]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:366
		// _ = "end of CoverTab[129361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:366
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:366
	// _ = "end of CoverTab[129357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:366
	_go_fuzz_dep_.CoverTab[129358]++

										typ := val.Type().Elem()
										for i, sec := range secs {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:369
		_go_fuzz_dep_.CoverTab[129362]++
											elem := reflect.New(typ)
											if err = sec.mapToField(elem, isStrict, i, sec.name); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:371
			_go_fuzz_dep_.CoverTab[129364]++
												return reflect.Value{}, fmt.Errorf("map to field from section %q: %v", secName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:372
			// _ = "end of CoverTab[129364]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:373
			_go_fuzz_dep_.CoverTab[129365]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:373
			// _ = "end of CoverTab[129365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:373
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:373
		// _ = "end of CoverTab[129362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:373
		_go_fuzz_dep_.CoverTab[129363]++

											val = reflect.Append(val, elem.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:375
		// _ = "end of CoverTab[129363]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:376
	// _ = "end of CoverTab[129358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:376
	_go_fuzz_dep_.CoverTab[129359]++
										return val, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:377
	// _ = "end of CoverTab[129359]"
}

// mapTo maps a section to object v.
func (s *Section) mapTo(v interface{}, isStrict bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:381
	_go_fuzz_dep_.CoverTab[129366]++
										typ := reflect.TypeOf(v)
										val := reflect.ValueOf(v)
										if typ.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:384
		_go_fuzz_dep_.CoverTab[129369]++
											typ = typ.Elem()
											val = val.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:386
		// _ = "end of CoverTab[129369]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:387
		_go_fuzz_dep_.CoverTab[129370]++
											return errors.New("not a pointer to a struct")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:388
		// _ = "end of CoverTab[129370]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:389
	// _ = "end of CoverTab[129366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:389
	_go_fuzz_dep_.CoverTab[129367]++

										if typ.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:391
		_go_fuzz_dep_.CoverTab[129371]++
											newField, err := s.mapToSlice(s.name, val, isStrict)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:393
			_go_fuzz_dep_.CoverTab[129373]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:394
			// _ = "end of CoverTab[129373]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:395
			_go_fuzz_dep_.CoverTab[129374]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:395
			// _ = "end of CoverTab[129374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:395
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:395
		// _ = "end of CoverTab[129371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:395
		_go_fuzz_dep_.CoverTab[129372]++

											val.Set(newField)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:398
		// _ = "end of CoverTab[129372]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:399
		_go_fuzz_dep_.CoverTab[129375]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:399
		// _ = "end of CoverTab[129375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:399
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:399
	// _ = "end of CoverTab[129367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:399
	_go_fuzz_dep_.CoverTab[129368]++

										return s.mapToField(val, isStrict, 0, s.name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:401
	// _ = "end of CoverTab[129368]"
}

// MapTo maps section to given struct.
func (s *Section) MapTo(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:405
	_go_fuzz_dep_.CoverTab[129376]++
										return s.mapTo(v, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:406
	// _ = "end of CoverTab[129376]"
}

// StrictMapTo maps section to given struct in strict mode,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:409
// which returns all possible error including value parsing error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:411
func (s *Section) StrictMapTo(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:411
	_go_fuzz_dep_.CoverTab[129377]++
										return s.mapTo(v, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:412
	// _ = "end of CoverTab[129377]"
}

// MapTo maps file to given struct.
func (f *File) MapTo(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:416
	_go_fuzz_dep_.CoverTab[129378]++
										return f.Section("").MapTo(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:417
	// _ = "end of CoverTab[129378]"
}

// StrictMapTo maps file to given struct in strict mode,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:420
// which returns all possible error including value parsing error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:422
func (f *File) StrictMapTo(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:422
	_go_fuzz_dep_.CoverTab[129379]++
										return f.Section("").StrictMapTo(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:423
	// _ = "end of CoverTab[129379]"
}

// MapToWithMapper maps data sources to given struct with name mapper.
func MapToWithMapper(v interface{}, mapper NameMapper, source interface{}, others ...interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:427
	_go_fuzz_dep_.CoverTab[129380]++
										cfg, err := Load(source, others...)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:429
		_go_fuzz_dep_.CoverTab[129382]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:430
		// _ = "end of CoverTab[129382]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:431
		_go_fuzz_dep_.CoverTab[129383]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:431
		// _ = "end of CoverTab[129383]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:431
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:431
	// _ = "end of CoverTab[129380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:431
	_go_fuzz_dep_.CoverTab[129381]++
										cfg.NameMapper = mapper
										return cfg.MapTo(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:433
	// _ = "end of CoverTab[129381]"
}

// StrictMapToWithMapper maps data sources to given struct with name mapper in strict mode,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:436
// which returns all possible error including value parsing error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:438
func StrictMapToWithMapper(v interface{}, mapper NameMapper, source interface{}, others ...interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:438
	_go_fuzz_dep_.CoverTab[129384]++
										cfg, err := Load(source, others...)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:440
		_go_fuzz_dep_.CoverTab[129386]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:441
		// _ = "end of CoverTab[129386]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:442
		_go_fuzz_dep_.CoverTab[129387]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:442
		// _ = "end of CoverTab[129387]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:442
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:442
	// _ = "end of CoverTab[129384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:442
	_go_fuzz_dep_.CoverTab[129385]++
										cfg.NameMapper = mapper
										return cfg.StrictMapTo(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:444
	// _ = "end of CoverTab[129385]"
}

// MapTo maps data sources to given struct.
func MapTo(v, source interface{}, others ...interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:448
	_go_fuzz_dep_.CoverTab[129388]++
										return MapToWithMapper(v, nil, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:449
	// _ = "end of CoverTab[129388]"
}

// StrictMapTo maps data sources to given struct in strict mode,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:452
// which returns all possible error including value parsing error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:454
func StrictMapTo(v, source interface{}, others ...interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:454
	_go_fuzz_dep_.CoverTab[129389]++
										return StrictMapToWithMapper(v, nil, source, others...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:455
	// _ = "end of CoverTab[129389]"
}

// reflectSliceWithProperType does the opposite thing as setSliceWithProperType.
func reflectSliceWithProperType(key *Key, field reflect.Value, delim string, allowShadow bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:459
	_go_fuzz_dep_.CoverTab[129390]++
										slice := field.Slice(0, field.Len())
										if field.Len() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:461
		_go_fuzz_dep_.CoverTab[129394]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:462
		// _ = "end of CoverTab[129394]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:463
		_go_fuzz_dep_.CoverTab[129395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:463
		// _ = "end of CoverTab[129395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:463
	// _ = "end of CoverTab[129390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:463
	_go_fuzz_dep_.CoverTab[129391]++
										sliceOf := field.Type().Elem().Kind()

										if allowShadow {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:466
		_go_fuzz_dep_.CoverTab[129396]++
											var keyWithShadows *Key
											for i := 0; i < field.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:468
			_go_fuzz_dep_.CoverTab[129398]++
												var val string
												switch sliceOf {
			case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:471
				_go_fuzz_dep_.CoverTab[129400]++
													val = slice.Index(i).String()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:472
				// _ = "end of CoverTab[129400]"
			case reflect.Int, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:473
				_go_fuzz_dep_.CoverTab[129401]++
													val = fmt.Sprint(slice.Index(i).Int())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:474
				// _ = "end of CoverTab[129401]"
			case reflect.Uint, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:475
				_go_fuzz_dep_.CoverTab[129402]++
													val = fmt.Sprint(slice.Index(i).Uint())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:476
				// _ = "end of CoverTab[129402]"
			case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:477
				_go_fuzz_dep_.CoverTab[129403]++
													val = fmt.Sprint(slice.Index(i).Float())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:478
				// _ = "end of CoverTab[129403]"
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:479
				_go_fuzz_dep_.CoverTab[129404]++
													val = fmt.Sprint(slice.Index(i).Bool())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:480
				// _ = "end of CoverTab[129404]"
			case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:481
				_go_fuzz_dep_.CoverTab[129405]++
													val = slice.Index(i).Interface().(time.Time).Format(time.RFC3339)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:482
				// _ = "end of CoverTab[129405]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:483
				_go_fuzz_dep_.CoverTab[129406]++
													return fmt.Errorf("unsupported type '[]%s'", sliceOf)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:484
				// _ = "end of CoverTab[129406]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:485
			// _ = "end of CoverTab[129398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:485
			_go_fuzz_dep_.CoverTab[129399]++

												if i == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:487
				_go_fuzz_dep_.CoverTab[129407]++
													keyWithShadows = newKey(key.s, key.name, val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:488
				// _ = "end of CoverTab[129407]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:489
				_go_fuzz_dep_.CoverTab[129408]++
													_ = keyWithShadows.AddShadow(val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:490
				// _ = "end of CoverTab[129408]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:491
			// _ = "end of CoverTab[129399]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:492
		// _ = "end of CoverTab[129396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:492
		_go_fuzz_dep_.CoverTab[129397]++
											*key = *keyWithShadows
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:494
		// _ = "end of CoverTab[129397]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:495
		_go_fuzz_dep_.CoverTab[129409]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:495
		// _ = "end of CoverTab[129409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:495
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:495
	// _ = "end of CoverTab[129391]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:495
	_go_fuzz_dep_.CoverTab[129392]++

										var buf bytes.Buffer
										for i := 0; i < field.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:498
		_go_fuzz_dep_.CoverTab[129410]++
											switch sliceOf {
		case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:500
			_go_fuzz_dep_.CoverTab[129412]++
												buf.WriteString(slice.Index(i).String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:501
			// _ = "end of CoverTab[129412]"
		case reflect.Int, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:502
			_go_fuzz_dep_.CoverTab[129413]++
												buf.WriteString(fmt.Sprint(slice.Index(i).Int()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:503
			// _ = "end of CoverTab[129413]"
		case reflect.Uint, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:504
			_go_fuzz_dep_.CoverTab[129414]++
												buf.WriteString(fmt.Sprint(slice.Index(i).Uint()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:505
			// _ = "end of CoverTab[129414]"
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:506
			_go_fuzz_dep_.CoverTab[129415]++
												buf.WriteString(fmt.Sprint(slice.Index(i).Float()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:507
			// _ = "end of CoverTab[129415]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:508
			_go_fuzz_dep_.CoverTab[129416]++
												buf.WriteString(fmt.Sprint(slice.Index(i).Bool()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:509
			// _ = "end of CoverTab[129416]"
		case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:510
			_go_fuzz_dep_.CoverTab[129417]++
												buf.WriteString(slice.Index(i).Interface().(time.Time).Format(time.RFC3339))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:511
			// _ = "end of CoverTab[129417]"
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:512
			_go_fuzz_dep_.CoverTab[129418]++
												return fmt.Errorf("unsupported type '[]%s'", sliceOf)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:513
			// _ = "end of CoverTab[129418]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:514
		// _ = "end of CoverTab[129410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:514
		_go_fuzz_dep_.CoverTab[129411]++
											buf.WriteString(delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:515
		// _ = "end of CoverTab[129411]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:516
	// _ = "end of CoverTab[129392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:516
	_go_fuzz_dep_.CoverTab[129393]++
										key.SetValue(buf.String()[:buf.Len()-len(delim)])
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:518
	// _ = "end of CoverTab[129393]"
}

// reflectWithProperType does the opposite thing as setWithProperType.
func reflectWithProperType(t reflect.Type, key *Key, field reflect.Value, delim string, allowShadow bool) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:522
	_go_fuzz_dep_.CoverTab[129419]++
										switch t.Kind() {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:524
		_go_fuzz_dep_.CoverTab[129421]++
											key.SetValue(field.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:525
		// _ = "end of CoverTab[129421]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:526
		_go_fuzz_dep_.CoverTab[129422]++
											key.SetValue(fmt.Sprint(field.Bool()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:527
		// _ = "end of CoverTab[129422]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:528
		_go_fuzz_dep_.CoverTab[129423]++
											key.SetValue(fmt.Sprint(field.Int()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:529
		// _ = "end of CoverTab[129423]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:530
		_go_fuzz_dep_.CoverTab[129424]++
											key.SetValue(fmt.Sprint(field.Uint()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:531
		// _ = "end of CoverTab[129424]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:532
		_go_fuzz_dep_.CoverTab[129425]++
											key.SetValue(fmt.Sprint(field.Float()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:533
		// _ = "end of CoverTab[129425]"
	case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:534
		_go_fuzz_dep_.CoverTab[129426]++
											key.SetValue(fmt.Sprint(field.Interface().(time.Time).Format(time.RFC3339)))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:535
		// _ = "end of CoverTab[129426]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:536
		_go_fuzz_dep_.CoverTab[129427]++
											return reflectSliceWithProperType(key, field, delim, allowShadow)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:537
		// _ = "end of CoverTab[129427]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:538
		_go_fuzz_dep_.CoverTab[129428]++
											if !field.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:539
			_go_fuzz_dep_.CoverTab[129430]++
												return reflectWithProperType(t.Elem(), key, field.Elem(), delim, allowShadow)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:540
			// _ = "end of CoverTab[129430]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:541
			_go_fuzz_dep_.CoverTab[129431]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:541
			// _ = "end of CoverTab[129431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:541
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:541
		// _ = "end of CoverTab[129428]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:542
		_go_fuzz_dep_.CoverTab[129429]++
											return fmt.Errorf("unsupported type %q", t)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:543
		// _ = "end of CoverTab[129429]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:544
	// _ = "end of CoverTab[129419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:544
	_go_fuzz_dep_.CoverTab[129420]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:545
	// _ = "end of CoverTab[129420]"
}

// CR: copied from encoding/json/encode.go with modifications of time.Time support.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:548
// TODO: add more test coverage.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:550
func isEmptyValue(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:550
	_go_fuzz_dep_.CoverTab[129432]++
										switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:552
		_go_fuzz_dep_.CoverTab[129434]++
											return v.Len() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:553
		// _ = "end of CoverTab[129434]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:554
		_go_fuzz_dep_.CoverTab[129435]++
											return !v.Bool()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:555
		// _ = "end of CoverTab[129435]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:556
		_go_fuzz_dep_.CoverTab[129436]++
											return v.Int() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:557
		// _ = "end of CoverTab[129436]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:558
		_go_fuzz_dep_.CoverTab[129437]++
											return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:559
		// _ = "end of CoverTab[129437]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:560
		_go_fuzz_dep_.CoverTab[129438]++
											return v.Float() == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:561
		// _ = "end of CoverTab[129438]"
	case reflect.Interface, reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:562
		_go_fuzz_dep_.CoverTab[129439]++
											return v.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:563
		// _ = "end of CoverTab[129439]"
	case reflectTime:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:564
		_go_fuzz_dep_.CoverTab[129440]++
											t, ok := v.Interface().(time.Time)
											return ok && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
			_go_fuzz_dep_.CoverTab[129442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
			return t.IsZero()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
			// _ = "end of CoverTab[129442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
		// _ = "end of CoverTab[129440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
		_go_fuzz_dep_.CoverTab[129441]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:566
		// _ = "end of CoverTab[129441]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:567
	// _ = "end of CoverTab[129432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:567
	_go_fuzz_dep_.CoverTab[129433]++
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:568
	// _ = "end of CoverTab[129433]"
}

// StructReflector is the interface implemented by struct types that can extract themselves into INI objects.
type StructReflector interface {
	ReflectINIStruct(*File) error
}

func (s *Section) reflectFrom(val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:576
	_go_fuzz_dep_.CoverTab[129443]++
										if val.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:577
		_go_fuzz_dep_.CoverTab[129446]++
											val = val.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:578
		// _ = "end of CoverTab[129446]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:579
		_go_fuzz_dep_.CoverTab[129447]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:579
		// _ = "end of CoverTab[129447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:579
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:579
	// _ = "end of CoverTab[129443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:579
	_go_fuzz_dep_.CoverTab[129444]++
										typ := val.Type()

										for i := 0; i < typ.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:582
		_go_fuzz_dep_.CoverTab[129448]++
											if !val.Field(i).CanInterface() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:583
			_go_fuzz_dep_.CoverTab[129459]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:584
			// _ = "end of CoverTab[129459]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:585
			_go_fuzz_dep_.CoverTab[129460]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:585
			// _ = "end of CoverTab[129460]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:585
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:585
		// _ = "end of CoverTab[129448]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:585
		_go_fuzz_dep_.CoverTab[129449]++

											field := val.Field(i)
											tpField := typ.Field(i)

											tag := tpField.Tag.Get("ini")
											if tag == "-" {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:591
			_go_fuzz_dep_.CoverTab[129461]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:592
			// _ = "end of CoverTab[129461]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:593
			_go_fuzz_dep_.CoverTab[129462]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:593
			// _ = "end of CoverTab[129462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:593
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:593
		// _ = "end of CoverTab[129449]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:593
		_go_fuzz_dep_.CoverTab[129450]++

											rawName, omitEmpty, allowShadow, allowNonUnique, extends := parseTagOptions(tag)
											if omitEmpty && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:596
			_go_fuzz_dep_.CoverTab[129463]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:596
			return isEmptyValue(field)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:596
			// _ = "end of CoverTab[129463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:596
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:596
			_go_fuzz_dep_.CoverTab[129464]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:597
			// _ = "end of CoverTab[129464]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:598
			_go_fuzz_dep_.CoverTab[129465]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:598
			// _ = "end of CoverTab[129465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:598
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:598
		// _ = "end of CoverTab[129450]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:598
		_go_fuzz_dep_.CoverTab[129451]++

											if r, ok := field.Interface().(StructReflector); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:600
			_go_fuzz_dep_.CoverTab[129466]++
												return r.ReflectINIStruct(s.f)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:601
			// _ = "end of CoverTab[129466]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:602
			_go_fuzz_dep_.CoverTab[129467]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:602
			// _ = "end of CoverTab[129467]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:602
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:602
		// _ = "end of CoverTab[129451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:602
		_go_fuzz_dep_.CoverTab[129452]++

											fieldName := s.parseFieldName(tpField.Name, rawName)
											if len(fieldName) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:605
			_go_fuzz_dep_.CoverTab[129468]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:605
			return !field.CanSet()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:605
			// _ = "end of CoverTab[129468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:605
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:605
			_go_fuzz_dep_.CoverTab[129469]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:606
			// _ = "end of CoverTab[129469]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:607
			_go_fuzz_dep_.CoverTab[129470]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:607
			// _ = "end of CoverTab[129470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:607
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:607
		// _ = "end of CoverTab[129452]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:607
		_go_fuzz_dep_.CoverTab[129453]++

											if extends && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			_go_fuzz_dep_.CoverTab[129471]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			return tpField.Anonymous
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			// _ = "end of CoverTab[129471]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			_go_fuzz_dep_.CoverTab[129472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			return (tpField.Type.Kind() == reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
				_go_fuzz_dep_.CoverTab[129473]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
				return tpField.Type.Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
				// _ = "end of CoverTab[129473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			// _ = "end of CoverTab[129472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:609
			_go_fuzz_dep_.CoverTab[129474]++
												if err := s.reflectFrom(field); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:610
				_go_fuzz_dep_.CoverTab[129476]++
													return fmt.Errorf("reflect from field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:611
				// _ = "end of CoverTab[129476]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:612
				_go_fuzz_dep_.CoverTab[129477]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:612
				// _ = "end of CoverTab[129477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:612
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:612
			// _ = "end of CoverTab[129474]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:612
			_go_fuzz_dep_.CoverTab[129475]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:613
			// _ = "end of CoverTab[129475]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:614
			_go_fuzz_dep_.CoverTab[129478]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:614
			// _ = "end of CoverTab[129478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:614
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:614
		// _ = "end of CoverTab[129453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:614
		_go_fuzz_dep_.CoverTab[129454]++

											if (tpField.Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
			_go_fuzz_dep_.CoverTab[129479]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
			return tpField.Type.Elem().Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
			// _ = "end of CoverTab[129479]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
			_go_fuzz_dep_.CoverTab[129480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:616
			return (tpField.Type.Kind() == reflect.Struct && func() bool {
													_go_fuzz_dep_.CoverTab[129481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
				return tpField.Type.Name() != "Time"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
				// _ = "end of CoverTab[129481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
			// _ = "end of CoverTab[129480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:617
			_go_fuzz_dep_.CoverTab[129482]++

												sec, err := s.f.GetSection(fieldName)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:620
				_go_fuzz_dep_.CoverTab[129486]++

													sec, _ = s.f.NewSection(fieldName)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:622
				// _ = "end of CoverTab[129486]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:623
				_go_fuzz_dep_.CoverTab[129487]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:623
				// _ = "end of CoverTab[129487]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:623
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:623
			// _ = "end of CoverTab[129482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:623
			_go_fuzz_dep_.CoverTab[129483]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:626
			if len(sec.Comment) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:626
				_go_fuzz_dep_.CoverTab[129488]++
													sec.Comment = tpField.Tag.Get("comment")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:627
				// _ = "end of CoverTab[129488]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:628
				_go_fuzz_dep_.CoverTab[129489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:628
				// _ = "end of CoverTab[129489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:628
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:628
			// _ = "end of CoverTab[129483]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:628
			_go_fuzz_dep_.CoverTab[129484]++

												if err = sec.reflectFrom(field); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:630
				_go_fuzz_dep_.CoverTab[129490]++
													return fmt.Errorf("reflect from field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:631
				// _ = "end of CoverTab[129490]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:632
				_go_fuzz_dep_.CoverTab[129491]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:632
				// _ = "end of CoverTab[129491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:632
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:632
			// _ = "end of CoverTab[129484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:632
			_go_fuzz_dep_.CoverTab[129485]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:633
			// _ = "end of CoverTab[129485]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:634
			_go_fuzz_dep_.CoverTab[129492]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:634
			// _ = "end of CoverTab[129492]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:634
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:634
		// _ = "end of CoverTab[129454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:634
		_go_fuzz_dep_.CoverTab[129455]++

											if allowNonUnique && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:636
			_go_fuzz_dep_.CoverTab[129493]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:636
			return tpField.Type.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:636
			// _ = "end of CoverTab[129493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:636
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:636
			_go_fuzz_dep_.CoverTab[129494]++
												slice := field.Slice(0, field.Len())
												if field.Len() == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:638
				_go_fuzz_dep_.CoverTab[129497]++
													return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:639
				// _ = "end of CoverTab[129497]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:640
				_go_fuzz_dep_.CoverTab[129498]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:640
				// _ = "end of CoverTab[129498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:640
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:640
			// _ = "end of CoverTab[129494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:640
			_go_fuzz_dep_.CoverTab[129495]++
												sliceOf := field.Type().Elem().Kind()

												for i := 0; i < field.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:643
				_go_fuzz_dep_.CoverTab[129499]++
													if sliceOf != reflect.Struct && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:644
					_go_fuzz_dep_.CoverTab[129503]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:644
					return sliceOf != reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:644
					// _ = "end of CoverTab[129503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:644
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:644
					_go_fuzz_dep_.CoverTab[129504]++
														return fmt.Errorf("field %q is not a slice of pointer or struct", fieldName)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:645
					// _ = "end of CoverTab[129504]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:646
					_go_fuzz_dep_.CoverTab[129505]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:646
					// _ = "end of CoverTab[129505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:646
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:646
				// _ = "end of CoverTab[129499]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:646
				_go_fuzz_dep_.CoverTab[129500]++

													sec, err := s.f.NewSection(fieldName)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:649
					_go_fuzz_dep_.CoverTab[129506]++
														return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:650
					// _ = "end of CoverTab[129506]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:651
					_go_fuzz_dep_.CoverTab[129507]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:651
					// _ = "end of CoverTab[129507]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:651
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:651
				// _ = "end of CoverTab[129500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:651
				_go_fuzz_dep_.CoverTab[129501]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:654
				if len(sec.Comment) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:654
					_go_fuzz_dep_.CoverTab[129508]++
														sec.Comment = tpField.Tag.Get("comment")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:655
					// _ = "end of CoverTab[129508]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:656
					_go_fuzz_dep_.CoverTab[129509]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:656
					// _ = "end of CoverTab[129509]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:656
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:656
				// _ = "end of CoverTab[129501]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:656
				_go_fuzz_dep_.CoverTab[129502]++

													if err := sec.reflectFrom(slice.Index(i)); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:658
					_go_fuzz_dep_.CoverTab[129510]++
														return fmt.Errorf("reflect from field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:659
					// _ = "end of CoverTab[129510]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:660
					_go_fuzz_dep_.CoverTab[129511]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:660
					// _ = "end of CoverTab[129511]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:660
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:660
				// _ = "end of CoverTab[129502]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:661
			// _ = "end of CoverTab[129495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:661
			_go_fuzz_dep_.CoverTab[129496]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:662
			// _ = "end of CoverTab[129496]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:663
			_go_fuzz_dep_.CoverTab[129512]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:663
			// _ = "end of CoverTab[129512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:663
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:663
		// _ = "end of CoverTab[129455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:663
		_go_fuzz_dep_.CoverTab[129456]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:666
		key, err := s.GetKey(fieldName)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:667
			_go_fuzz_dep_.CoverTab[129513]++
												key, _ = s.NewKey(fieldName, "")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:668
			// _ = "end of CoverTab[129513]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:669
			_go_fuzz_dep_.CoverTab[129514]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:669
			// _ = "end of CoverTab[129514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:669
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:669
		// _ = "end of CoverTab[129456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:669
		_go_fuzz_dep_.CoverTab[129457]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:672
		if len(key.Comment) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:672
			_go_fuzz_dep_.CoverTab[129515]++
												key.Comment = tpField.Tag.Get("comment")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:673
			// _ = "end of CoverTab[129515]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:674
			_go_fuzz_dep_.CoverTab[129516]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:674
			// _ = "end of CoverTab[129516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:674
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:674
		// _ = "end of CoverTab[129457]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:674
		_go_fuzz_dep_.CoverTab[129458]++

											delim := parseDelim(tpField.Tag.Get("delim"))
											if err = reflectWithProperType(tpField.Type, key, field, delim, allowShadow); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:677
			_go_fuzz_dep_.CoverTab[129517]++
												return fmt.Errorf("reflect field %q: %v", fieldName, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:678
			// _ = "end of CoverTab[129517]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:679
			_go_fuzz_dep_.CoverTab[129518]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:679
			// _ = "end of CoverTab[129518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:679
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:679
		// _ = "end of CoverTab[129458]"

	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:681
	// _ = "end of CoverTab[129444]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:681
	_go_fuzz_dep_.CoverTab[129445]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:682
	// _ = "end of CoverTab[129445]"
}

// ReflectFrom reflects section from given struct. It overwrites existing ones.
func (s *Section) ReflectFrom(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:686
	_go_fuzz_dep_.CoverTab[129519]++
										typ := reflect.TypeOf(v)
										val := reflect.ValueOf(v)

										if s.name != DefaultSection && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
		_go_fuzz_dep_.CoverTab[129522]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
		return s.f.options.AllowNonUniqueSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
		// _ = "end of CoverTab[129522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
		_go_fuzz_dep_.CoverTab[129523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:690
		return (typ.Kind() == reflect.Slice || func() bool {
												_go_fuzz_dep_.CoverTab[129524]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
			return typ.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
			// _ = "end of CoverTab[129524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
		// _ = "end of CoverTab[129523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:691
		_go_fuzz_dep_.CoverTab[129525]++

											s.f.DeleteSection(s.name)

											if typ.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:695
			_go_fuzz_dep_.CoverTab[129529]++
												sec, err := s.f.NewSection(s.name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:697
				_go_fuzz_dep_.CoverTab[129531]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:698
				// _ = "end of CoverTab[129531]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:699
				_go_fuzz_dep_.CoverTab[129532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:699
				// _ = "end of CoverTab[129532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:699
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:699
			// _ = "end of CoverTab[129529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:699
			_go_fuzz_dep_.CoverTab[129530]++
												return sec.reflectFrom(val.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:700
			// _ = "end of CoverTab[129530]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:701
			_go_fuzz_dep_.CoverTab[129533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:701
			// _ = "end of CoverTab[129533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:701
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:701
		// _ = "end of CoverTab[129525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:701
		_go_fuzz_dep_.CoverTab[129526]++

											slice := val.Slice(0, val.Len())
											sliceOf := val.Type().Elem().Kind()
											if sliceOf != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:705
			_go_fuzz_dep_.CoverTab[129534]++
												return fmt.Errorf("not a slice of pointers")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:706
			// _ = "end of CoverTab[129534]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:707
			_go_fuzz_dep_.CoverTab[129535]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:707
			// _ = "end of CoverTab[129535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:707
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:707
		// _ = "end of CoverTab[129526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:707
		_go_fuzz_dep_.CoverTab[129527]++

											for i := 0; i < slice.Len(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:709
			_go_fuzz_dep_.CoverTab[129536]++
												sec, err := s.f.NewSection(s.name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:711
				_go_fuzz_dep_.CoverTab[129538]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:712
				// _ = "end of CoverTab[129538]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:713
				_go_fuzz_dep_.CoverTab[129539]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:713
				// _ = "end of CoverTab[129539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:713
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:713
			// _ = "end of CoverTab[129536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:713
			_go_fuzz_dep_.CoverTab[129537]++

												err = sec.reflectFrom(slice.Index(i))
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:716
				_go_fuzz_dep_.CoverTab[129540]++
													return fmt.Errorf("reflect from %dth field: %v", i, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:717
				// _ = "end of CoverTab[129540]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:718
				_go_fuzz_dep_.CoverTab[129541]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:718
				// _ = "end of CoverTab[129541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:718
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:718
			// _ = "end of CoverTab[129537]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:719
		// _ = "end of CoverTab[129527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:719
		_go_fuzz_dep_.CoverTab[129528]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:721
		// _ = "end of CoverTab[129528]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:722
		_go_fuzz_dep_.CoverTab[129542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:722
		// _ = "end of CoverTab[129542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:722
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:722
	// _ = "end of CoverTab[129519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:722
	_go_fuzz_dep_.CoverTab[129520]++

										if typ.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:724
		_go_fuzz_dep_.CoverTab[129543]++
											val = val.Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:725
		// _ = "end of CoverTab[129543]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:726
		_go_fuzz_dep_.CoverTab[129544]++
											return errors.New("not a pointer to a struct")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:727
		// _ = "end of CoverTab[129544]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:728
	// _ = "end of CoverTab[129520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:728
	_go_fuzz_dep_.CoverTab[129521]++

										return s.reflectFrom(val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:730
	// _ = "end of CoverTab[129521]"
}

// ReflectFrom reflects file from given struct.
func (f *File) ReflectFrom(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:734
	_go_fuzz_dep_.CoverTab[129545]++
										return f.Section("").ReflectFrom(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:735
	// _ = "end of CoverTab[129545]"
}

// ReflectFromWithMapper reflects data sources from given struct with name mapper.
func ReflectFromWithMapper(cfg *File, v interface{}, mapper NameMapper) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:739
	_go_fuzz_dep_.CoverTab[129546]++
										cfg.NameMapper = mapper
										return cfg.ReflectFrom(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:741
	// _ = "end of CoverTab[129546]"
}

// ReflectFrom reflects data sources from given struct.
func ReflectFrom(cfg *File, v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:745
	_go_fuzz_dep_.CoverTab[129547]++
										return ReflectFromWithMapper(cfg, v, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:746
	// _ = "end of CoverTab[129547]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:747
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/struct.go:747
var _ = _go_fuzz_dep_.CoverTab
