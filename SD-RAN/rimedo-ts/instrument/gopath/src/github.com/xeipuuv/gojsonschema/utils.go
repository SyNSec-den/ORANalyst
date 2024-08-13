// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Various utility functions.
//
// created          26-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:26
)

import (
	"encoding/json"
	"math/big"
	"reflect"
)

func isKind(what interface{}, kinds ...reflect.Kind) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:34
	_go_fuzz_dep_.CoverTab[195729]++
											target := what
											if isJSONNumber(what) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:36
		_go_fuzz_dep_.CoverTab[195732]++

												target = *mustBeNumber(what)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:38
		// _ = "end of CoverTab[195732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:39
		_go_fuzz_dep_.CoverTab[195733]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:39
		// _ = "end of CoverTab[195733]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:39
	// _ = "end of CoverTab[195729]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:39
	_go_fuzz_dep_.CoverTab[195730]++
											targetKind := reflect.ValueOf(target).Kind()
											for _, kind := range kinds {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:41
		_go_fuzz_dep_.CoverTab[195734]++
												if targetKind == kind {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:42
			_go_fuzz_dep_.CoverTab[195735]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:43
			// _ = "end of CoverTab[195735]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:44
			_go_fuzz_dep_.CoverTab[195736]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:44
			// _ = "end of CoverTab[195736]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:44
		// _ = "end of CoverTab[195734]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:45
	// _ = "end of CoverTab[195730]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:45
	_go_fuzz_dep_.CoverTab[195731]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:46
	// _ = "end of CoverTab[195731]"
}

func existsMapKey(m map[string]interface{}, k string) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:49
	_go_fuzz_dep_.CoverTab[195737]++
											_, ok := m[k]
											return ok
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:51
	// _ = "end of CoverTab[195737]"
}

func isStringInSlice(s []string, what string) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:54
	_go_fuzz_dep_.CoverTab[195738]++
											for i := range s {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:55
		_go_fuzz_dep_.CoverTab[195740]++
												if s[i] == what {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:56
			_go_fuzz_dep_.CoverTab[195741]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:57
			// _ = "end of CoverTab[195741]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:58
			_go_fuzz_dep_.CoverTab[195742]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:58
			// _ = "end of CoverTab[195742]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:58
		// _ = "end of CoverTab[195740]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:59
	// _ = "end of CoverTab[195738]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:59
	_go_fuzz_dep_.CoverTab[195739]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:60
	// _ = "end of CoverTab[195739]"
}

// indexStringInSlice returns the index of the first instance of 'what' in s or -1 if it is not found in s.
func indexStringInSlice(s []string, what string) int {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:64
	_go_fuzz_dep_.CoverTab[195743]++
											for i := range s {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:65
		_go_fuzz_dep_.CoverTab[195745]++
												if s[i] == what {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:66
			_go_fuzz_dep_.CoverTab[195746]++
													return i
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:67
			// _ = "end of CoverTab[195746]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:68
			_go_fuzz_dep_.CoverTab[195747]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:68
			// _ = "end of CoverTab[195747]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:68
		// _ = "end of CoverTab[195745]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:69
	// _ = "end of CoverTab[195743]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:69
	_go_fuzz_dep_.CoverTab[195744]++
											return -1
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:70
	// _ = "end of CoverTab[195744]"
}

func marshalToJSONString(value interface{}) (*string, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:73
	_go_fuzz_dep_.CoverTab[195748]++

											mBytes, err := json.Marshal(value)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:76
		_go_fuzz_dep_.CoverTab[195750]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:77
		// _ = "end of CoverTab[195750]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:78
		_go_fuzz_dep_.CoverTab[195751]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:78
		// _ = "end of CoverTab[195751]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:78
	// _ = "end of CoverTab[195748]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:78
	_go_fuzz_dep_.CoverTab[195749]++

											sBytes := string(mBytes)
											return &sBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:81
	// _ = "end of CoverTab[195749]"
}

func marshalWithoutNumber(value interface{}) (*string, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:84
	_go_fuzz_dep_.CoverTab[195752]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:91
	jsonString, err := marshalToJSONString(value)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:92
		_go_fuzz_dep_.CoverTab[195755]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:93
		// _ = "end of CoverTab[195755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:94
		_go_fuzz_dep_.CoverTab[195756]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:94
		// _ = "end of CoverTab[195756]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:94
	// _ = "end of CoverTab[195752]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:94
	_go_fuzz_dep_.CoverTab[195753]++

											var document interface{}

											err = json.Unmarshal([]byte(*jsonString), &document)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:99
		_go_fuzz_dep_.CoverTab[195757]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:100
		// _ = "end of CoverTab[195757]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:101
		_go_fuzz_dep_.CoverTab[195758]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:101
		// _ = "end of CoverTab[195758]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:101
	// _ = "end of CoverTab[195753]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:101
	_go_fuzz_dep_.CoverTab[195754]++

											return marshalToJSONString(document)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:103
	// _ = "end of CoverTab[195754]"
}

func isJSONNumber(what interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:106
	_go_fuzz_dep_.CoverTab[195759]++

											switch what.(type) {

	case json.Number:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:110
		_go_fuzz_dep_.CoverTab[195761]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:111
		// _ = "end of CoverTab[195761]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:112
	// _ = "end of CoverTab[195759]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:112
	_go_fuzz_dep_.CoverTab[195760]++

											return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:114
	// _ = "end of CoverTab[195760]"
}

func checkJSONInteger(what interface{}) (isInt bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:117
	_go_fuzz_dep_.CoverTab[195762]++

											jsonNumber := what.(json.Number)

											bigFloat, isValidNumber := new(big.Rat).SetString(string(jsonNumber))

											return isValidNumber && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:123
		_go_fuzz_dep_.CoverTab[195763]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:123
		return bigFloat.IsInt()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:123
		// _ = "end of CoverTab[195763]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:123
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:123
	// _ = "end of CoverTab[195762]"

}

// same as ECMA Number.MAX_SAFE_INTEGER and Number.MIN_SAFE_INTEGER
const (
	maxJSONFloat	= float64(1<<53 - 1)	// 9007199254740991.0 	 2^53 - 1
	minJSONFloat	= -float64(1<<53 - 1)	//-9007199254740991.0	-2^53 - 1
)

func mustBeInteger(what interface{}) *int {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:133
	_go_fuzz_dep_.CoverTab[195764]++

											if isJSONNumber(what) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:135
		_go_fuzz_dep_.CoverTab[195766]++

												number := what.(json.Number)

												isInt := checkJSONInteger(number)

												if isInt {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:141
			_go_fuzz_dep_.CoverTab[195767]++

													int64Value, err := number.Int64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:144
				_go_fuzz_dep_.CoverTab[195769]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:145
				// _ = "end of CoverTab[195769]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:146
				_go_fuzz_dep_.CoverTab[195770]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:146
				// _ = "end of CoverTab[195770]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:146
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:146
			// _ = "end of CoverTab[195767]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:146
			_go_fuzz_dep_.CoverTab[195768]++

													int32Value := int(int64Value)
													return &int32Value
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:149
			// _ = "end of CoverTab[195768]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:150
			_go_fuzz_dep_.CoverTab[195771]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:150
			// _ = "end of CoverTab[195771]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:150
		// _ = "end of CoverTab[195766]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:152
		_go_fuzz_dep_.CoverTab[195772]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:152
		// _ = "end of CoverTab[195772]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:152
	// _ = "end of CoverTab[195764]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:152
	_go_fuzz_dep_.CoverTab[195765]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:154
	// _ = "end of CoverTab[195765]"
}

func mustBeNumber(what interface{}) *big.Rat {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:157
	_go_fuzz_dep_.CoverTab[195773]++

											if isJSONNumber(what) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:159
		_go_fuzz_dep_.CoverTab[195775]++
												number := what.(json.Number)
												float64Value, success := new(big.Rat).SetString(string(number))
												if success {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:162
			_go_fuzz_dep_.CoverTab[195776]++
													return float64Value
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:163
			// _ = "end of CoverTab[195776]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:164
			_go_fuzz_dep_.CoverTab[195777]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:164
			// _ = "end of CoverTab[195777]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:164
		// _ = "end of CoverTab[195775]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:165
		_go_fuzz_dep_.CoverTab[195778]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:165
		// _ = "end of CoverTab[195778]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:165
	// _ = "end of CoverTab[195773]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:165
	_go_fuzz_dep_.CoverTab[195774]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:167
	// _ = "end of CoverTab[195774]"

}

func convertDocumentNode(val interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:171
	_go_fuzz_dep_.CoverTab[195779]++

											if lval, ok := val.([]interface{}); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:173
		_go_fuzz_dep_.CoverTab[195782]++

												res := []interface{}{}
												for _, v := range lval {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:176
			_go_fuzz_dep_.CoverTab[195784]++
													res = append(res, convertDocumentNode(v))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:177
			// _ = "end of CoverTab[195784]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:178
		// _ = "end of CoverTab[195782]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:178
		_go_fuzz_dep_.CoverTab[195783]++

												return res
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:180
		// _ = "end of CoverTab[195783]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:182
		_go_fuzz_dep_.CoverTab[195785]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:182
		// _ = "end of CoverTab[195785]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:182
	// _ = "end of CoverTab[195779]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:182
	_go_fuzz_dep_.CoverTab[195780]++

											if mval, ok := val.(map[interface{}]interface{}); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:184
		_go_fuzz_dep_.CoverTab[195786]++

												res := map[string]interface{}{}

												for k, v := range mval {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:188
			_go_fuzz_dep_.CoverTab[195788]++
													res[k.(string)] = convertDocumentNode(v)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:189
			// _ = "end of CoverTab[195788]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:190
		// _ = "end of CoverTab[195786]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:190
		_go_fuzz_dep_.CoverTab[195787]++

												return res
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:192
		// _ = "end of CoverTab[195787]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:194
		_go_fuzz_dep_.CoverTab[195789]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:194
		// _ = "end of CoverTab[195789]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:194
	// _ = "end of CoverTab[195780]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:194
	_go_fuzz_dep_.CoverTab[195781]++

											return val
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:196
	// _ = "end of CoverTab[195781]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:197
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/utils.go:197
var _ = _go_fuzz_dep_.CoverTab
