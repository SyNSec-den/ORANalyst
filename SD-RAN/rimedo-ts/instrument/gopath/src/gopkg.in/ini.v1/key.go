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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:15
)

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Key represents a key under a section.
type Key struct {
	s		*Section
	Comment		string
	name		string
	value		string
	isAutoIncrement	bool
	isBooleanType	bool

	isShadow	bool
	shadows		[]*Key

	nestedValues	[]string
}

// newKey simply return a key object with given values.
func newKey(s *Section, name, val string) *Key {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:42
	_go_fuzz_dep_.CoverTab[128517]++
									return &Key{
		s:	s,
		name:	name,
		value:	val,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:47
	// _ = "end of CoverTab[128517]"
}

func (k *Key) addShadow(val string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:50
	_go_fuzz_dep_.CoverTab[128518]++
									if k.isShadow {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:51
		_go_fuzz_dep_.CoverTab[128521]++
										return errors.New("cannot add shadow to another shadow key")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:52
		// _ = "end of CoverTab[128521]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
		_go_fuzz_dep_.CoverTab[128522]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
		if k.isAutoIncrement || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
			_go_fuzz_dep_.CoverTab[128523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
			return k.isBooleanType
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
			// _ = "end of CoverTab[128523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:53
			_go_fuzz_dep_.CoverTab[128524]++
											return errors.New("cannot add shadow to auto-increment or boolean key")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:54
			// _ = "end of CoverTab[128524]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
			_go_fuzz_dep_.CoverTab[128525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
			// _ = "end of CoverTab[128525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
		// _ = "end of CoverTab[128522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
	// _ = "end of CoverTab[128518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:55
	_go_fuzz_dep_.CoverTab[128519]++

									if !k.s.f.options.AllowDuplicateShadowValues {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:57
		_go_fuzz_dep_.CoverTab[128526]++

										if k.value == val {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:59
			_go_fuzz_dep_.CoverTab[128528]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:60
			// _ = "end of CoverTab[128528]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:61
			_go_fuzz_dep_.CoverTab[128529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:61
			// _ = "end of CoverTab[128529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:61
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:61
		// _ = "end of CoverTab[128526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:61
		_go_fuzz_dep_.CoverTab[128527]++
										for i := range k.shadows {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:62
			_go_fuzz_dep_.CoverTab[128530]++
											if k.shadows[i].value == val {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:63
				_go_fuzz_dep_.CoverTab[128531]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:64
				// _ = "end of CoverTab[128531]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:65
				_go_fuzz_dep_.CoverTab[128532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:65
				// _ = "end of CoverTab[128532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:65
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:65
			// _ = "end of CoverTab[128530]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:66
		// _ = "end of CoverTab[128527]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:67
		_go_fuzz_dep_.CoverTab[128533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:67
		// _ = "end of CoverTab[128533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:67
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:67
	// _ = "end of CoverTab[128519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:67
	_go_fuzz_dep_.CoverTab[128520]++

									shadow := newKey(k.s, k.name, val)
									shadow.isShadow = true
									k.shadows = append(k.shadows, shadow)
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:72
	// _ = "end of CoverTab[128520]"
}

// AddShadow adds a new shadow key to itself.
func (k *Key) AddShadow(val string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:76
	_go_fuzz_dep_.CoverTab[128534]++
									if !k.s.f.options.AllowShadows {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:77
		_go_fuzz_dep_.CoverTab[128536]++
										return errors.New("shadow key is not allowed")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:78
		// _ = "end of CoverTab[128536]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:79
		_go_fuzz_dep_.CoverTab[128537]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:79
		// _ = "end of CoverTab[128537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:79
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:79
	// _ = "end of CoverTab[128534]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:79
	_go_fuzz_dep_.CoverTab[128535]++
									return k.addShadow(val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:80
	// _ = "end of CoverTab[128535]"
}

func (k *Key) addNestedValue(val string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:83
	_go_fuzz_dep_.CoverTab[128538]++
									if k.isAutoIncrement || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:84
		_go_fuzz_dep_.CoverTab[128540]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:84
		return k.isBooleanType
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:84
		// _ = "end of CoverTab[128540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:84
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:84
		_go_fuzz_dep_.CoverTab[128541]++
										return errors.New("cannot add nested value to auto-increment or boolean key")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:85
		// _ = "end of CoverTab[128541]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:86
		_go_fuzz_dep_.CoverTab[128542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:86
		// _ = "end of CoverTab[128542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:86
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:86
	// _ = "end of CoverTab[128538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:86
	_go_fuzz_dep_.CoverTab[128539]++

									k.nestedValues = append(k.nestedValues, val)
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:89
	// _ = "end of CoverTab[128539]"
}

// AddNestedValue adds a nested value to the key.
func (k *Key) AddNestedValue(val string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:93
	_go_fuzz_dep_.CoverTab[128543]++
									if !k.s.f.options.AllowNestedValues {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:94
		_go_fuzz_dep_.CoverTab[128545]++
										return errors.New("nested value is not allowed")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:95
		// _ = "end of CoverTab[128545]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:96
		_go_fuzz_dep_.CoverTab[128546]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:96
		// _ = "end of CoverTab[128546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:96
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:96
	// _ = "end of CoverTab[128543]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:96
	_go_fuzz_dep_.CoverTab[128544]++
									return k.addNestedValue(val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:97
	// _ = "end of CoverTab[128544]"
}

// ValueMapper represents a mapping function for values, e.g. os.ExpandEnv
type ValueMapper func(string) string

// Name returns name of key.
func (k *Key) Name() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:104
	_go_fuzz_dep_.CoverTab[128547]++
									return k.name
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:105
	// _ = "end of CoverTab[128547]"
}

// Value returns raw value of key for performance purpose.
func (k *Key) Value() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:109
	_go_fuzz_dep_.CoverTab[128548]++
									return k.value
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:110
	// _ = "end of CoverTab[128548]"
}

// ValueWithShadows returns raw values of key and its shadows if any.
func (k *Key) ValueWithShadows() []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:114
	_go_fuzz_dep_.CoverTab[128549]++
									if len(k.shadows) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:115
		_go_fuzz_dep_.CoverTab[128552]++
										return []string{k.value}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:116
		// _ = "end of CoverTab[128552]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:117
		_go_fuzz_dep_.CoverTab[128553]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:117
		// _ = "end of CoverTab[128553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:117
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:117
	// _ = "end of CoverTab[128549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:117
	_go_fuzz_dep_.CoverTab[128550]++
									vals := make([]string, len(k.shadows)+1)
									vals[0] = k.value
									for i := range k.shadows {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:120
		_go_fuzz_dep_.CoverTab[128554]++
										vals[i+1] = k.shadows[i].value
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:121
		// _ = "end of CoverTab[128554]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:122
	// _ = "end of CoverTab[128550]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:122
	_go_fuzz_dep_.CoverTab[128551]++
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:123
	// _ = "end of CoverTab[128551]"
}

// NestedValues returns nested values stored in the key.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:126
// It is possible returned value is nil if no nested values stored in the key.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:128
func (k *Key) NestedValues() []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:128
	_go_fuzz_dep_.CoverTab[128555]++
									return k.nestedValues
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:129
	// _ = "end of CoverTab[128555]"
}

// transformValue takes a raw value and transforms to its final string.
func (k *Key) transformValue(val string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:133
	_go_fuzz_dep_.CoverTab[128556]++
									if k.s.f.ValueMapper != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:134
		_go_fuzz_dep_.CoverTab[128560]++
										val = k.s.f.ValueMapper(val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:135
		// _ = "end of CoverTab[128560]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:136
		_go_fuzz_dep_.CoverTab[128561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:136
		// _ = "end of CoverTab[128561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:136
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:136
	// _ = "end of CoverTab[128556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:136
	_go_fuzz_dep_.CoverTab[128557]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:139
	if !strings.Contains(val, "%") {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:139
		_go_fuzz_dep_.CoverTab[128562]++
										return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:140
		// _ = "end of CoverTab[128562]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:141
		_go_fuzz_dep_.CoverTab[128563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:141
		// _ = "end of CoverTab[128563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:141
	// _ = "end of CoverTab[128557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:141
	_go_fuzz_dep_.CoverTab[128558]++
									for i := 0; i < depthValues; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:142
		_go_fuzz_dep_.CoverTab[128564]++
										vr := varPattern.FindString(val)
										if len(vr) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:144
			_go_fuzz_dep_.CoverTab[128567]++
											break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:145
			// _ = "end of CoverTab[128567]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:146
			_go_fuzz_dep_.CoverTab[128568]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:146
			// _ = "end of CoverTab[128568]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:146
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:146
		// _ = "end of CoverTab[128564]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:146
		_go_fuzz_dep_.CoverTab[128565]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:149
		noption := vr[2 : len(vr)-2]

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:153
		nk, err := k.s.GetKey(noption)
		if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:154
			_go_fuzz_dep_.CoverTab[128569]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:154
			return k == nk
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:154
			// _ = "end of CoverTab[128569]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:154
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:154
			_go_fuzz_dep_.CoverTab[128570]++
											nk, _ = k.s.f.Section("").GetKey(noption)
											if nk == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:156
				_go_fuzz_dep_.CoverTab[128571]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:159
				break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:159
				// _ = "end of CoverTab[128571]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:160
				_go_fuzz_dep_.CoverTab[128572]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:160
				// _ = "end of CoverTab[128572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:160
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:160
			// _ = "end of CoverTab[128570]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:161
			_go_fuzz_dep_.CoverTab[128573]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:161
			// _ = "end of CoverTab[128573]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:161
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:161
		// _ = "end of CoverTab[128565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:161
		_go_fuzz_dep_.CoverTab[128566]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:164
		val = strings.Replace(val, vr, nk.value, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:164
		// _ = "end of CoverTab[128566]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:165
	// _ = "end of CoverTab[128558]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:165
	_go_fuzz_dep_.CoverTab[128559]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:166
	// _ = "end of CoverTab[128559]"
}

// String returns string representation of value.
func (k *Key) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:170
	_go_fuzz_dep_.CoverTab[128574]++
									return k.transformValue(k.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:171
	// _ = "end of CoverTab[128574]"
}

// Validate accepts a validate function which can
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:174
// return modifed result as key value.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:176
func (k *Key) Validate(fn func(string) string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:176
	_go_fuzz_dep_.CoverTab[128575]++
									return fn(k.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:177
	// _ = "end of CoverTab[128575]"
}

// parseBool returns the boolean value represented by the string.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:180
//
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:180
// It accepts 1, t, T, TRUE, true, True, YES, yes, Yes, y, ON, on, On,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:180
// 0, f, F, FALSE, false, False, NO, no, No, n, OFF, off, Off.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:180
// Any other value returns an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:185
func parseBool(str string) (value bool, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:185
	_go_fuzz_dep_.CoverTab[128576]++
									switch str {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:187
		_go_fuzz_dep_.CoverTab[128578]++
										return true, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:188
		// _ = "end of CoverTab[128578]"
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:189
		_go_fuzz_dep_.CoverTab[128579]++
										return false, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:190
		// _ = "end of CoverTab[128579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:190
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:190
		_go_fuzz_dep_.CoverTab[128580]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:190
		// _ = "end of CoverTab[128580]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:191
	// _ = "end of CoverTab[128576]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:191
	_go_fuzz_dep_.CoverTab[128577]++
									return false, fmt.Errorf("parsing \"%s\": invalid syntax", str)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:192
	// _ = "end of CoverTab[128577]"
}

// Bool returns bool type value.
func (k *Key) Bool() (bool, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:196
	_go_fuzz_dep_.CoverTab[128581]++
									return parseBool(k.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:197
	// _ = "end of CoverTab[128581]"
}

// Float64 returns float64 type value.
func (k *Key) Float64() (float64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:201
	_go_fuzz_dep_.CoverTab[128582]++
									return strconv.ParseFloat(k.String(), 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:202
	// _ = "end of CoverTab[128582]"
}

// Int returns int type value.
func (k *Key) Int() (int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:206
	_go_fuzz_dep_.CoverTab[128583]++
									v, err := strconv.ParseInt(k.String(), 0, 64)
									return int(v), err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:208
	// _ = "end of CoverTab[128583]"
}

// Int64 returns int64 type value.
func (k *Key) Int64() (int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:212
	_go_fuzz_dep_.CoverTab[128584]++
									return strconv.ParseInt(k.String(), 0, 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:213
	// _ = "end of CoverTab[128584]"
}

// Uint returns uint type valued.
func (k *Key) Uint() (uint, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:217
	_go_fuzz_dep_.CoverTab[128585]++
									u, e := strconv.ParseUint(k.String(), 0, 64)
									return uint(u), e
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:219
	// _ = "end of CoverTab[128585]"
}

// Uint64 returns uint64 type value.
func (k *Key) Uint64() (uint64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:223
	_go_fuzz_dep_.CoverTab[128586]++
									return strconv.ParseUint(k.String(), 0, 64)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:224
	// _ = "end of CoverTab[128586]"
}

// Duration returns time.Duration type value.
func (k *Key) Duration() (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:228
	_go_fuzz_dep_.CoverTab[128587]++
									return time.ParseDuration(k.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:229
	// _ = "end of CoverTab[128587]"
}

// TimeFormat parses with given format and returns time.Time type value.
func (k *Key) TimeFormat(format string) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:233
	_go_fuzz_dep_.CoverTab[128588]++
									return time.Parse(format, k.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:234
	// _ = "end of CoverTab[128588]"
}

// Time parses with RFC3339 format and returns time.Time type value.
func (k *Key) Time() (time.Time, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:238
	_go_fuzz_dep_.CoverTab[128589]++
									return k.TimeFormat(time.RFC3339)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:239
	// _ = "end of CoverTab[128589]"
}

// MustString returns default value if key value is empty.
func (k *Key) MustString(defaultVal string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:243
	_go_fuzz_dep_.CoverTab[128590]++
									val := k.String()
									if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:245
		_go_fuzz_dep_.CoverTab[128592]++
										k.value = defaultVal
										return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:247
		// _ = "end of CoverTab[128592]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:248
		_go_fuzz_dep_.CoverTab[128593]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:248
		// _ = "end of CoverTab[128593]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:248
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:248
	// _ = "end of CoverTab[128590]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:248
	_go_fuzz_dep_.CoverTab[128591]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:249
	// _ = "end of CoverTab[128591]"
}

// MustBool always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:252
// it returns false if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:254
func (k *Key) MustBool(defaultVal ...bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:254
	_go_fuzz_dep_.CoverTab[128594]++
									val, err := k.Bool()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:256
		_go_fuzz_dep_.CoverTab[128596]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:256
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:256
		// _ = "end of CoverTab[128596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:256
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:256
		_go_fuzz_dep_.CoverTab[128597]++
										k.value = strconv.FormatBool(defaultVal[0])
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:258
		// _ = "end of CoverTab[128597]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:259
		_go_fuzz_dep_.CoverTab[128598]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:259
		// _ = "end of CoverTab[128598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:259
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:259
	// _ = "end of CoverTab[128594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:259
	_go_fuzz_dep_.CoverTab[128595]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:260
	// _ = "end of CoverTab[128595]"
}

// MustFloat64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:263
// it returns 0.0 if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:265
func (k *Key) MustFloat64(defaultVal ...float64) float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:265
	_go_fuzz_dep_.CoverTab[128599]++
									val, err := k.Float64()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:267
		_go_fuzz_dep_.CoverTab[128601]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:267
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:267
		// _ = "end of CoverTab[128601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:267
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:267
		_go_fuzz_dep_.CoverTab[128602]++
										k.value = strconv.FormatFloat(defaultVal[0], 'f', -1, 64)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:269
		// _ = "end of CoverTab[128602]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:270
		_go_fuzz_dep_.CoverTab[128603]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:270
		// _ = "end of CoverTab[128603]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:270
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:270
	// _ = "end of CoverTab[128599]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:270
	_go_fuzz_dep_.CoverTab[128600]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:271
	// _ = "end of CoverTab[128600]"
}

// MustInt always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:274
// it returns 0 if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:276
func (k *Key) MustInt(defaultVal ...int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:276
	_go_fuzz_dep_.CoverTab[128604]++
									val, err := k.Int()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:278
		_go_fuzz_dep_.CoverTab[128606]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:278
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:278
		// _ = "end of CoverTab[128606]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:278
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:278
		_go_fuzz_dep_.CoverTab[128607]++
										k.value = strconv.FormatInt(int64(defaultVal[0]), 10)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:280
		// _ = "end of CoverTab[128607]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:281
		_go_fuzz_dep_.CoverTab[128608]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:281
		// _ = "end of CoverTab[128608]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:281
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:281
	// _ = "end of CoverTab[128604]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:281
	_go_fuzz_dep_.CoverTab[128605]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:282
	// _ = "end of CoverTab[128605]"
}

// MustInt64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:285
// it returns 0 if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:287
func (k *Key) MustInt64(defaultVal ...int64) int64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:287
	_go_fuzz_dep_.CoverTab[128609]++
									val, err := k.Int64()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:289
		_go_fuzz_dep_.CoverTab[128611]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:289
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:289
		// _ = "end of CoverTab[128611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:289
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:289
		_go_fuzz_dep_.CoverTab[128612]++
										k.value = strconv.FormatInt(defaultVal[0], 10)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:291
		// _ = "end of CoverTab[128612]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:292
		_go_fuzz_dep_.CoverTab[128613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:292
		// _ = "end of CoverTab[128613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:292
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:292
	// _ = "end of CoverTab[128609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:292
	_go_fuzz_dep_.CoverTab[128610]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:293
	// _ = "end of CoverTab[128610]"
}

// MustUint always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:296
// it returns 0 if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:298
func (k *Key) MustUint(defaultVal ...uint) uint {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:298
	_go_fuzz_dep_.CoverTab[128614]++
									val, err := k.Uint()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:300
		_go_fuzz_dep_.CoverTab[128616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:300
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:300
		// _ = "end of CoverTab[128616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:300
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:300
		_go_fuzz_dep_.CoverTab[128617]++
										k.value = strconv.FormatUint(uint64(defaultVal[0]), 10)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:302
		// _ = "end of CoverTab[128617]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:303
		_go_fuzz_dep_.CoverTab[128618]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:303
		// _ = "end of CoverTab[128618]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:303
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:303
	// _ = "end of CoverTab[128614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:303
	_go_fuzz_dep_.CoverTab[128615]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:304
	// _ = "end of CoverTab[128615]"
}

// MustUint64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:307
// it returns 0 if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:309
func (k *Key) MustUint64(defaultVal ...uint64) uint64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:309
	_go_fuzz_dep_.CoverTab[128619]++
									val, err := k.Uint64()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:311
		_go_fuzz_dep_.CoverTab[128621]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:311
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:311
		// _ = "end of CoverTab[128621]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:311
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:311
		_go_fuzz_dep_.CoverTab[128622]++
										k.value = strconv.FormatUint(defaultVal[0], 10)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:313
		// _ = "end of CoverTab[128622]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:314
		_go_fuzz_dep_.CoverTab[128623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:314
		// _ = "end of CoverTab[128623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:314
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:314
	// _ = "end of CoverTab[128619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:314
	_go_fuzz_dep_.CoverTab[128620]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:315
	// _ = "end of CoverTab[128620]"
}

// MustDuration always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:318
// it returns zero value if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:320
func (k *Key) MustDuration(defaultVal ...time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:320
	_go_fuzz_dep_.CoverTab[128624]++
									val, err := k.Duration()
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:322
		_go_fuzz_dep_.CoverTab[128626]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:322
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:322
		// _ = "end of CoverTab[128626]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:322
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:322
		_go_fuzz_dep_.CoverTab[128627]++
										k.value = defaultVal[0].String()
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:324
		// _ = "end of CoverTab[128627]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:325
		_go_fuzz_dep_.CoverTab[128628]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:325
		// _ = "end of CoverTab[128628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:325
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:325
	// _ = "end of CoverTab[128624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:325
	_go_fuzz_dep_.CoverTab[128625]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:326
	// _ = "end of CoverTab[128625]"
}

// MustTimeFormat always parses with given format and returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:329
// it returns zero value if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:331
func (k *Key) MustTimeFormat(format string, defaultVal ...time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:331
	_go_fuzz_dep_.CoverTab[128629]++
									val, err := k.TimeFormat(format)
									if len(defaultVal) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:333
		_go_fuzz_dep_.CoverTab[128631]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:333
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:333
		// _ = "end of CoverTab[128631]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:333
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:333
		_go_fuzz_dep_.CoverTab[128632]++
										k.value = defaultVal[0].Format(format)
										return defaultVal[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:335
		// _ = "end of CoverTab[128632]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:336
		_go_fuzz_dep_.CoverTab[128633]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:336
		// _ = "end of CoverTab[128633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:336
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:336
	// _ = "end of CoverTab[128629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:336
	_go_fuzz_dep_.CoverTab[128630]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:337
	// _ = "end of CoverTab[128630]"
}

// MustTime always parses with RFC3339 format and returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:340
// it returns zero value if error occurs.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:342
func (k *Key) MustTime(defaultVal ...time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:342
	_go_fuzz_dep_.CoverTab[128634]++
									return k.MustTimeFormat(time.RFC3339, defaultVal...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:343
	// _ = "end of CoverTab[128634]"
}

// In always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:346
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:348
func (k *Key) In(defaultVal string, candidates []string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:348
	_go_fuzz_dep_.CoverTab[128635]++
									val := k.String()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:350
		_go_fuzz_dep_.CoverTab[128637]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:351
			_go_fuzz_dep_.CoverTab[128638]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:352
			// _ = "end of CoverTab[128638]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:353
			_go_fuzz_dep_.CoverTab[128639]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:353
			// _ = "end of CoverTab[128639]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:353
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:353
		// _ = "end of CoverTab[128637]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:354
	// _ = "end of CoverTab[128635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:354
	_go_fuzz_dep_.CoverTab[128636]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:355
	// _ = "end of CoverTab[128636]"
}

// InFloat64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:358
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:360
func (k *Key) InFloat64(defaultVal float64, candidates []float64) float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:360
	_go_fuzz_dep_.CoverTab[128640]++
									val := k.MustFloat64()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:362
		_go_fuzz_dep_.CoverTab[128642]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:363
			_go_fuzz_dep_.CoverTab[128643]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:364
			// _ = "end of CoverTab[128643]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:365
			_go_fuzz_dep_.CoverTab[128644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:365
			// _ = "end of CoverTab[128644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:365
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:365
		// _ = "end of CoverTab[128642]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:366
	// _ = "end of CoverTab[128640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:366
	_go_fuzz_dep_.CoverTab[128641]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:367
	// _ = "end of CoverTab[128641]"
}

// InInt always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:370
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:372
func (k *Key) InInt(defaultVal int, candidates []int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:372
	_go_fuzz_dep_.CoverTab[128645]++
									val := k.MustInt()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:374
		_go_fuzz_dep_.CoverTab[128647]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:375
			_go_fuzz_dep_.CoverTab[128648]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:376
			// _ = "end of CoverTab[128648]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:377
			_go_fuzz_dep_.CoverTab[128649]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:377
			// _ = "end of CoverTab[128649]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:377
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:377
		// _ = "end of CoverTab[128647]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:378
	// _ = "end of CoverTab[128645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:378
	_go_fuzz_dep_.CoverTab[128646]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:379
	// _ = "end of CoverTab[128646]"
}

// InInt64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:382
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:384
func (k *Key) InInt64(defaultVal int64, candidates []int64) int64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:384
	_go_fuzz_dep_.CoverTab[128650]++
									val := k.MustInt64()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:386
		_go_fuzz_dep_.CoverTab[128652]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:387
			_go_fuzz_dep_.CoverTab[128653]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:388
			// _ = "end of CoverTab[128653]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:389
			_go_fuzz_dep_.CoverTab[128654]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:389
			// _ = "end of CoverTab[128654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:389
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:389
		// _ = "end of CoverTab[128652]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:390
	// _ = "end of CoverTab[128650]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:390
	_go_fuzz_dep_.CoverTab[128651]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:391
	// _ = "end of CoverTab[128651]"
}

// InUint always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:394
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:396
func (k *Key) InUint(defaultVal uint, candidates []uint) uint {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:396
	_go_fuzz_dep_.CoverTab[128655]++
									val := k.MustUint()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:398
		_go_fuzz_dep_.CoverTab[128657]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:399
			_go_fuzz_dep_.CoverTab[128658]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:400
			// _ = "end of CoverTab[128658]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:401
			_go_fuzz_dep_.CoverTab[128659]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:401
			// _ = "end of CoverTab[128659]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:401
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:401
		// _ = "end of CoverTab[128657]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:402
	// _ = "end of CoverTab[128655]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:402
	_go_fuzz_dep_.CoverTab[128656]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:403
	// _ = "end of CoverTab[128656]"
}

// InUint64 always returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:406
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:408
func (k *Key) InUint64(defaultVal uint64, candidates []uint64) uint64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:408
	_go_fuzz_dep_.CoverTab[128660]++
									val := k.MustUint64()
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:410
		_go_fuzz_dep_.CoverTab[128662]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:411
			_go_fuzz_dep_.CoverTab[128663]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:412
			// _ = "end of CoverTab[128663]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:413
			_go_fuzz_dep_.CoverTab[128664]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:413
			// _ = "end of CoverTab[128664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:413
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:413
		// _ = "end of CoverTab[128662]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:414
	// _ = "end of CoverTab[128660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:414
	_go_fuzz_dep_.CoverTab[128661]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:415
	// _ = "end of CoverTab[128661]"
}

// InTimeFormat always parses with given format and returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:418
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:420
func (k *Key) InTimeFormat(format string, defaultVal time.Time, candidates []time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:420
	_go_fuzz_dep_.CoverTab[128665]++
									val := k.MustTimeFormat(format)
									for _, cand := range candidates {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:422
		_go_fuzz_dep_.CoverTab[128667]++
										if val == cand {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:423
			_go_fuzz_dep_.CoverTab[128668]++
											return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:424
			// _ = "end of CoverTab[128668]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:425
			_go_fuzz_dep_.CoverTab[128669]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:425
			// _ = "end of CoverTab[128669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:425
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:425
		// _ = "end of CoverTab[128667]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:426
	// _ = "end of CoverTab[128665]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:426
	_go_fuzz_dep_.CoverTab[128666]++
									return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:427
	// _ = "end of CoverTab[128666]"
}

// InTime always parses with RFC3339 format and returns value without error,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:430
// it returns default value if error occurs or doesn't fit into candidates.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:432
func (k *Key) InTime(defaultVal time.Time, candidates []time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:432
	_go_fuzz_dep_.CoverTab[128670]++
									return k.InTimeFormat(time.RFC3339, defaultVal, candidates)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:433
	// _ = "end of CoverTab[128670]"
}

// RangeFloat64 checks if value is in given range inclusively,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:436
// and returns default value if it's not.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:438
func (k *Key) RangeFloat64(defaultVal, min, max float64) float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:438
	_go_fuzz_dep_.CoverTab[128671]++
									val := k.MustFloat64()
									if val < min || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:440
		_go_fuzz_dep_.CoverTab[128673]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:440
		return val > max
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:440
		// _ = "end of CoverTab[128673]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:440
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:440
		_go_fuzz_dep_.CoverTab[128674]++
										return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:441
		// _ = "end of CoverTab[128674]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:442
		_go_fuzz_dep_.CoverTab[128675]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:442
		// _ = "end of CoverTab[128675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:442
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:442
	// _ = "end of CoverTab[128671]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:442
	_go_fuzz_dep_.CoverTab[128672]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:443
	// _ = "end of CoverTab[128672]"
}

// RangeInt checks if value is in given range inclusively,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:446
// and returns default value if it's not.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:448
func (k *Key) RangeInt(defaultVal, min, max int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:448
	_go_fuzz_dep_.CoverTab[128676]++
									val := k.MustInt()
									if val < min || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:450
		_go_fuzz_dep_.CoverTab[128678]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:450
		return val > max
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:450
		// _ = "end of CoverTab[128678]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:450
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:450
		_go_fuzz_dep_.CoverTab[128679]++
										return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:451
		// _ = "end of CoverTab[128679]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:452
		_go_fuzz_dep_.CoverTab[128680]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:452
		// _ = "end of CoverTab[128680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:452
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:452
	// _ = "end of CoverTab[128676]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:452
	_go_fuzz_dep_.CoverTab[128677]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:453
	// _ = "end of CoverTab[128677]"
}

// RangeInt64 checks if value is in given range inclusively,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:456
// and returns default value if it's not.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:458
func (k *Key) RangeInt64(defaultVal, min, max int64) int64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:458
	_go_fuzz_dep_.CoverTab[128681]++
									val := k.MustInt64()
									if val < min || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:460
		_go_fuzz_dep_.CoverTab[128683]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:460
		return val > max
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:460
		// _ = "end of CoverTab[128683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:460
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:460
		_go_fuzz_dep_.CoverTab[128684]++
										return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:461
		// _ = "end of CoverTab[128684]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:462
		_go_fuzz_dep_.CoverTab[128685]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:462
		// _ = "end of CoverTab[128685]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:462
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:462
	// _ = "end of CoverTab[128681]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:462
	_go_fuzz_dep_.CoverTab[128682]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:463
	// _ = "end of CoverTab[128682]"
}

// RangeTimeFormat checks if value with given format is in given range inclusively,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:466
// and returns default value if it's not.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:468
func (k *Key) RangeTimeFormat(format string, defaultVal, min, max time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:468
	_go_fuzz_dep_.CoverTab[128686]++
									val := k.MustTimeFormat(format)
									if val.Unix() < min.Unix() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:470
		_go_fuzz_dep_.CoverTab[128688]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:470
		return val.Unix() > max.Unix()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:470
		// _ = "end of CoverTab[128688]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:470
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:470
		_go_fuzz_dep_.CoverTab[128689]++
										return defaultVal
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:471
		// _ = "end of CoverTab[128689]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:472
		_go_fuzz_dep_.CoverTab[128690]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:472
		// _ = "end of CoverTab[128690]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:472
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:472
	// _ = "end of CoverTab[128686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:472
	_go_fuzz_dep_.CoverTab[128687]++
									return val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:473
	// _ = "end of CoverTab[128687]"
}

// RangeTime checks if value with RFC3339 format is in given range inclusively,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:476
// and returns default value if it's not.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:478
func (k *Key) RangeTime(defaultVal, min, max time.Time) time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:478
	_go_fuzz_dep_.CoverTab[128691]++
									return k.RangeTimeFormat(time.RFC3339, defaultVal, min, max)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:479
	// _ = "end of CoverTab[128691]"
}

// Strings returns list of string divided by given delimiter.
func (k *Key) Strings(delim string) []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:483
	_go_fuzz_dep_.CoverTab[128692]++
									str := k.String()
									if len(str) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:485
		_go_fuzz_dep_.CoverTab[128696]++
										return []string{}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:486
		// _ = "end of CoverTab[128696]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:487
		_go_fuzz_dep_.CoverTab[128697]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:487
		// _ = "end of CoverTab[128697]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:487
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:487
	// _ = "end of CoverTab[128692]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:487
	_go_fuzz_dep_.CoverTab[128693]++

									runes := []rune(str)
									vals := make([]string, 0, 2)
									var buf bytes.Buffer
									escape := false
									idx := 0
									for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:494
		_go_fuzz_dep_.CoverTab[128698]++
										if escape {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:495
			_go_fuzz_dep_.CoverTab[128700]++
											escape = false
											if runes[idx] != '\\' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:497
				_go_fuzz_dep_.CoverTab[128702]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:497
				return !strings.HasPrefix(string(runes[idx:]), delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:497
				// _ = "end of CoverTab[128702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:497
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:497
				_go_fuzz_dep_.CoverTab[128703]++
												buf.WriteRune('\\')
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:498
				// _ = "end of CoverTab[128703]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:499
				_go_fuzz_dep_.CoverTab[128704]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:499
				// _ = "end of CoverTab[128704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:499
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:499
			// _ = "end of CoverTab[128700]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:499
			_go_fuzz_dep_.CoverTab[128701]++
											buf.WriteRune(runes[idx])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:500
			// _ = "end of CoverTab[128701]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:501
			_go_fuzz_dep_.CoverTab[128705]++
											if runes[idx] == '\\' {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:502
				_go_fuzz_dep_.CoverTab[128706]++
												escape = true
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:503
				// _ = "end of CoverTab[128706]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:504
				_go_fuzz_dep_.CoverTab[128707]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:504
				if strings.HasPrefix(string(runes[idx:]), delim) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:504
					_go_fuzz_dep_.CoverTab[128708]++
													idx += len(delim) - 1
													vals = append(vals, strings.TrimSpace(buf.String()))
													buf.Reset()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:507
					// _ = "end of CoverTab[128708]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:508
					_go_fuzz_dep_.CoverTab[128709]++
													buf.WriteRune(runes[idx])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:509
					// _ = "end of CoverTab[128709]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:510
				// _ = "end of CoverTab[128707]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:510
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:510
			// _ = "end of CoverTab[128705]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:511
		// _ = "end of CoverTab[128698]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:511
		_go_fuzz_dep_.CoverTab[128699]++
										idx++
										if idx == len(runes) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:513
			_go_fuzz_dep_.CoverTab[128710]++
											break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:514
			// _ = "end of CoverTab[128710]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:515
			_go_fuzz_dep_.CoverTab[128711]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:515
			// _ = "end of CoverTab[128711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:515
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:515
		// _ = "end of CoverTab[128699]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:516
	// _ = "end of CoverTab[128693]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:516
	_go_fuzz_dep_.CoverTab[128694]++

									if buf.Len() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:518
		_go_fuzz_dep_.CoverTab[128712]++
										vals = append(vals, strings.TrimSpace(buf.String()))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:519
		// _ = "end of CoverTab[128712]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:520
		_go_fuzz_dep_.CoverTab[128713]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:520
		// _ = "end of CoverTab[128713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:520
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:520
	// _ = "end of CoverTab[128694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:520
	_go_fuzz_dep_.CoverTab[128695]++

									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:522
	// _ = "end of CoverTab[128695]"
}

// StringsWithShadows returns list of string divided by given delimiter.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:525
// Shadows will also be appended if any.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:527
func (k *Key) StringsWithShadows(delim string) []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:527
	_go_fuzz_dep_.CoverTab[128714]++
									vals := k.ValueWithShadows()
									results := make([]string, 0, len(vals)*2)
									for i := range vals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:530
		_go_fuzz_dep_.CoverTab[128717]++
										if len(vals) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:531
			_go_fuzz_dep_.CoverTab[128719]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:532
			// _ = "end of CoverTab[128719]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:533
			_go_fuzz_dep_.CoverTab[128720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:533
			// _ = "end of CoverTab[128720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:533
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:533
		// _ = "end of CoverTab[128717]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:533
		_go_fuzz_dep_.CoverTab[128718]++

										results = append(results, strings.Split(vals[i], delim)...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:535
		// _ = "end of CoverTab[128718]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:536
	// _ = "end of CoverTab[128714]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:536
	_go_fuzz_dep_.CoverTab[128715]++

									for i := range results {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:538
		_go_fuzz_dep_.CoverTab[128721]++
										results[i] = k.transformValue(strings.TrimSpace(results[i]))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:539
		// _ = "end of CoverTab[128721]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:540
	// _ = "end of CoverTab[128715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:540
	_go_fuzz_dep_.CoverTab[128716]++
									return results
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:541
	// _ = "end of CoverTab[128716]"
}

// Float64s returns list of float64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Float64s(delim string) []float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:545
	_go_fuzz_dep_.CoverTab[128722]++
									vals, _ := k.parseFloat64s(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:547
	// _ = "end of CoverTab[128722]"
}

// Ints returns list of int divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Ints(delim string) []int {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:551
	_go_fuzz_dep_.CoverTab[128723]++
									vals, _ := k.parseInts(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:553
	// _ = "end of CoverTab[128723]"
}

// Int64s returns list of int64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Int64s(delim string) []int64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:557
	_go_fuzz_dep_.CoverTab[128724]++
									vals, _ := k.parseInt64s(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:559
	// _ = "end of CoverTab[128724]"
}

// Uints returns list of uint divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Uints(delim string) []uint {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:563
	_go_fuzz_dep_.CoverTab[128725]++
									vals, _ := k.parseUints(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:565
	// _ = "end of CoverTab[128725]"
}

// Uint64s returns list of uint64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Uint64s(delim string) []uint64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:569
	_go_fuzz_dep_.CoverTab[128726]++
									vals, _ := k.parseUint64s(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:571
	// _ = "end of CoverTab[128726]"
}

// Bools returns list of bool divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Bools(delim string) []bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:575
	_go_fuzz_dep_.CoverTab[128727]++
									vals, _ := k.parseBools(k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:577
	// _ = "end of CoverTab[128727]"
}

// TimesFormat parses with given format and returns list of time.Time divided by given delimiter.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:580
// Any invalid input will be treated as zero value (0001-01-01 00:00:00 +0000 UTC).
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:582
func (k *Key) TimesFormat(format, delim string) []time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:582
	_go_fuzz_dep_.CoverTab[128728]++
									vals, _ := k.parseTimesFormat(format, k.Strings(delim), true, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:584
	// _ = "end of CoverTab[128728]"
}

// Times parses with RFC3339 format and returns list of time.Time divided by given delimiter.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:587
// Any invalid input will be treated as zero value (0001-01-01 00:00:00 +0000 UTC).
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:589
func (k *Key) Times(delim string) []time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:589
	_go_fuzz_dep_.CoverTab[128729]++
									return k.TimesFormat(time.RFC3339, delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:590
	// _ = "end of CoverTab[128729]"
}

// ValidFloat64s returns list of float64 divided by given delimiter. If some value is not float, then
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:593
// it will not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:595
func (k *Key) ValidFloat64s(delim string) []float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:595
	_go_fuzz_dep_.CoverTab[128730]++
									vals, _ := k.parseFloat64s(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:597
	// _ = "end of CoverTab[128730]"
}

// ValidInts returns list of int divided by given delimiter. If some value is not integer, then it will
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:600
// not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:602
func (k *Key) ValidInts(delim string) []int {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:602
	_go_fuzz_dep_.CoverTab[128731]++
									vals, _ := k.parseInts(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:604
	// _ = "end of CoverTab[128731]"
}

// ValidInt64s returns list of int64 divided by given delimiter. If some value is not 64-bit integer,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:607
// then it will not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:609
func (k *Key) ValidInt64s(delim string) []int64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:609
	_go_fuzz_dep_.CoverTab[128732]++
									vals, _ := k.parseInt64s(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:611
	// _ = "end of CoverTab[128732]"
}

// ValidUints returns list of uint divided by given delimiter. If some value is not unsigned integer,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:614
// then it will not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:616
func (k *Key) ValidUints(delim string) []uint {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:616
	_go_fuzz_dep_.CoverTab[128733]++
									vals, _ := k.parseUints(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:618
	// _ = "end of CoverTab[128733]"
}

// ValidUint64s returns list of uint64 divided by given delimiter. If some value is not 64-bit unsigned
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:621
// integer, then it will not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:623
func (k *Key) ValidUint64s(delim string) []uint64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:623
	_go_fuzz_dep_.CoverTab[128734]++
									vals, _ := k.parseUint64s(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:625
	// _ = "end of CoverTab[128734]"
}

// ValidBools returns list of bool divided by given delimiter. If some value is not 64-bit unsigned
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:628
// integer, then it will not be included to result list.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:630
func (k *Key) ValidBools(delim string) []bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:630
	_go_fuzz_dep_.CoverTab[128735]++
									vals, _ := k.parseBools(k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:632
	// _ = "end of CoverTab[128735]"
}

// ValidTimesFormat parses with given format and returns list of time.Time divided by given delimiter.
func (k *Key) ValidTimesFormat(format, delim string) []time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:636
	_go_fuzz_dep_.CoverTab[128736]++
									vals, _ := k.parseTimesFormat(format, k.Strings(delim), false, false)
									return vals
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:638
	// _ = "end of CoverTab[128736]"
}

// ValidTimes parses with RFC3339 format and returns list of time.Time divided by given delimiter.
func (k *Key) ValidTimes(delim string) []time.Time {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:642
	_go_fuzz_dep_.CoverTab[128737]++
									return k.ValidTimesFormat(time.RFC3339, delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:643
	// _ = "end of CoverTab[128737]"
}

// StrictFloat64s returns list of float64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictFloat64s(delim string) ([]float64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:647
	_go_fuzz_dep_.CoverTab[128738]++
									return k.parseFloat64s(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:648
	// _ = "end of CoverTab[128738]"
}

// StrictInts returns list of int divided by given delimiter or error on first invalid input.
func (k *Key) StrictInts(delim string) ([]int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:652
	_go_fuzz_dep_.CoverTab[128739]++
									return k.parseInts(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:653
	// _ = "end of CoverTab[128739]"
}

// StrictInt64s returns list of int64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictInt64s(delim string) ([]int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:657
	_go_fuzz_dep_.CoverTab[128740]++
									return k.parseInt64s(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:658
	// _ = "end of CoverTab[128740]"
}

// StrictUints returns list of uint divided by given delimiter or error on first invalid input.
func (k *Key) StrictUints(delim string) ([]uint, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:662
	_go_fuzz_dep_.CoverTab[128741]++
									return k.parseUints(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:663
	// _ = "end of CoverTab[128741]"
}

// StrictUint64s returns list of uint64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictUint64s(delim string) ([]uint64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:667
	_go_fuzz_dep_.CoverTab[128742]++
									return k.parseUint64s(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:668
	// _ = "end of CoverTab[128742]"
}

// StrictBools returns list of bool divided by given delimiter or error on first invalid input.
func (k *Key) StrictBools(delim string) ([]bool, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:672
	_go_fuzz_dep_.CoverTab[128743]++
									return k.parseBools(k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:673
	// _ = "end of CoverTab[128743]"
}

// StrictTimesFormat parses with given format and returns list of time.Time divided by given delimiter
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:676
// or error on first invalid input.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:678
func (k *Key) StrictTimesFormat(format, delim string) ([]time.Time, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:678
	_go_fuzz_dep_.CoverTab[128744]++
									return k.parseTimesFormat(format, k.Strings(delim), false, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:679
	// _ = "end of CoverTab[128744]"
}

// StrictTimes parses with RFC3339 format and returns list of time.Time divided by given delimiter
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:682
// or error on first invalid input.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:684
func (k *Key) StrictTimes(delim string) ([]time.Time, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:684
	_go_fuzz_dep_.CoverTab[128745]++
									return k.StrictTimesFormat(time.RFC3339, delim)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:685
	// _ = "end of CoverTab[128745]"
}

// parseBools transforms strings to bools.
func (k *Key) parseBools(strs []string, addInvalid, returnOnInvalid bool) ([]bool, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:689
	_go_fuzz_dep_.CoverTab[128746]++
									vals := make([]bool, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:691
		_go_fuzz_dep_.CoverTab[128749]++
										val, err := parseBool(str)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:693
		// _ = "end of CoverTab[128749]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:694
	// _ = "end of CoverTab[128746]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:694
	_go_fuzz_dep_.CoverTab[128747]++
									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:696
		_go_fuzz_dep_.CoverTab[128750]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:697
			_go_fuzz_dep_.CoverTab[128751]++
											vals = append(vals, val.(bool))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:698
			// _ = "end of CoverTab[128751]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:699
		// _ = "end of CoverTab[128750]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:700
		_go_fuzz_dep_.CoverTab[128752]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:700
		// _ = "end of CoverTab[128752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:700
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:700
	// _ = "end of CoverTab[128747]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:700
	_go_fuzz_dep_.CoverTab[128748]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:701
	// _ = "end of CoverTab[128748]"
}

// parseFloat64s transforms strings to float64s.
func (k *Key) parseFloat64s(strs []string, addInvalid, returnOnInvalid bool) ([]float64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:705
	_go_fuzz_dep_.CoverTab[128753]++
									vals := make([]float64, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:707
		_go_fuzz_dep_.CoverTab[128756]++
										val, err := strconv.ParseFloat(str, 64)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:709
		// _ = "end of CoverTab[128756]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:710
	// _ = "end of CoverTab[128753]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:710
	_go_fuzz_dep_.CoverTab[128754]++
									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:712
		_go_fuzz_dep_.CoverTab[128757]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:713
			_go_fuzz_dep_.CoverTab[128758]++
											vals = append(vals, val.(float64))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:714
			// _ = "end of CoverTab[128758]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:715
		// _ = "end of CoverTab[128757]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:716
		_go_fuzz_dep_.CoverTab[128759]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:716
		// _ = "end of CoverTab[128759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:716
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:716
	// _ = "end of CoverTab[128754]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:716
	_go_fuzz_dep_.CoverTab[128755]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:717
	// _ = "end of CoverTab[128755]"
}

// parseInts transforms strings to ints.
func (k *Key) parseInts(strs []string, addInvalid, returnOnInvalid bool) ([]int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:721
	_go_fuzz_dep_.CoverTab[128760]++
									vals := make([]int, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:723
		_go_fuzz_dep_.CoverTab[128763]++
										val, err := strconv.ParseInt(str, 0, 64)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:725
		// _ = "end of CoverTab[128763]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:726
	// _ = "end of CoverTab[128760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:726
	_go_fuzz_dep_.CoverTab[128761]++
									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:728
		_go_fuzz_dep_.CoverTab[128764]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:729
			_go_fuzz_dep_.CoverTab[128765]++
											vals = append(vals, int(val.(int64)))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:730
			// _ = "end of CoverTab[128765]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:731
		// _ = "end of CoverTab[128764]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:732
		_go_fuzz_dep_.CoverTab[128766]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:732
		// _ = "end of CoverTab[128766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:732
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:732
	// _ = "end of CoverTab[128761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:732
	_go_fuzz_dep_.CoverTab[128762]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:733
	// _ = "end of CoverTab[128762]"
}

// parseInt64s transforms strings to int64s.
func (k *Key) parseInt64s(strs []string, addInvalid, returnOnInvalid bool) ([]int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:737
	_go_fuzz_dep_.CoverTab[128767]++
									vals := make([]int64, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:739
		_go_fuzz_dep_.CoverTab[128770]++
										val, err := strconv.ParseInt(str, 0, 64)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:741
		// _ = "end of CoverTab[128770]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:742
	// _ = "end of CoverTab[128767]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:742
	_go_fuzz_dep_.CoverTab[128768]++

									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:745
		_go_fuzz_dep_.CoverTab[128771]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:746
			_go_fuzz_dep_.CoverTab[128772]++
											vals = append(vals, val.(int64))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:747
			// _ = "end of CoverTab[128772]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:748
		// _ = "end of CoverTab[128771]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:749
		_go_fuzz_dep_.CoverTab[128773]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:749
		// _ = "end of CoverTab[128773]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:749
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:749
	// _ = "end of CoverTab[128768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:749
	_go_fuzz_dep_.CoverTab[128769]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:750
	// _ = "end of CoverTab[128769]"
}

// parseUints transforms strings to uints.
func (k *Key) parseUints(strs []string, addInvalid, returnOnInvalid bool) ([]uint, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:754
	_go_fuzz_dep_.CoverTab[128774]++
									vals := make([]uint, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:756
		_go_fuzz_dep_.CoverTab[128777]++
										val, err := strconv.ParseUint(str, 0, 64)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:758
		// _ = "end of CoverTab[128777]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:759
	// _ = "end of CoverTab[128774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:759
	_go_fuzz_dep_.CoverTab[128775]++

									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:762
		_go_fuzz_dep_.CoverTab[128778]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:763
			_go_fuzz_dep_.CoverTab[128779]++
											vals = append(vals, uint(val.(uint64)))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:764
			// _ = "end of CoverTab[128779]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:765
		// _ = "end of CoverTab[128778]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:766
		_go_fuzz_dep_.CoverTab[128780]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:766
		// _ = "end of CoverTab[128780]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:766
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:766
	// _ = "end of CoverTab[128775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:766
	_go_fuzz_dep_.CoverTab[128776]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:767
	// _ = "end of CoverTab[128776]"
}

// parseUint64s transforms strings to uint64s.
func (k *Key) parseUint64s(strs []string, addInvalid, returnOnInvalid bool) ([]uint64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:771
	_go_fuzz_dep_.CoverTab[128781]++
									vals := make([]uint64, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:773
		_go_fuzz_dep_.CoverTab[128784]++
										val, err := strconv.ParseUint(str, 0, 64)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:775
		// _ = "end of CoverTab[128784]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:776
	// _ = "end of CoverTab[128781]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:776
	_go_fuzz_dep_.CoverTab[128782]++
									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:778
		_go_fuzz_dep_.CoverTab[128785]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:779
			_go_fuzz_dep_.CoverTab[128786]++
											vals = append(vals, val.(uint64))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:780
			// _ = "end of CoverTab[128786]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:781
		// _ = "end of CoverTab[128785]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:782
		_go_fuzz_dep_.CoverTab[128787]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:782
		// _ = "end of CoverTab[128787]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:782
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:782
	// _ = "end of CoverTab[128782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:782
	_go_fuzz_dep_.CoverTab[128783]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:783
	// _ = "end of CoverTab[128783]"
}

type Parser func(str string) (interface{}, error)

// parseTimesFormat transforms strings to times in given format.
func (k *Key) parseTimesFormat(format string, strs []string, addInvalid, returnOnInvalid bool) ([]time.Time, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:789
	_go_fuzz_dep_.CoverTab[128788]++
									vals := make([]time.Time, 0, len(strs))
									parser := func(str string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:791
		_go_fuzz_dep_.CoverTab[128791]++
										val, err := time.Parse(format, str)
										return val, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:793
		// _ = "end of CoverTab[128791]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:794
	// _ = "end of CoverTab[128788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:794
	_go_fuzz_dep_.CoverTab[128789]++
									rawVals, err := k.doParse(strs, addInvalid, returnOnInvalid, parser)
									if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:796
		_go_fuzz_dep_.CoverTab[128792]++
										for _, val := range rawVals {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:797
			_go_fuzz_dep_.CoverTab[128793]++
											vals = append(vals, val.(time.Time))
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:798
			// _ = "end of CoverTab[128793]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:799
		// _ = "end of CoverTab[128792]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:800
		_go_fuzz_dep_.CoverTab[128794]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:800
		// _ = "end of CoverTab[128794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:800
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:800
	// _ = "end of CoverTab[128789]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:800
	_go_fuzz_dep_.CoverTab[128790]++
									return vals, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:801
	// _ = "end of CoverTab[128790]"
}

// doParse transforms strings to different types
func (k *Key) doParse(strs []string, addInvalid, returnOnInvalid bool, parser Parser) ([]interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:805
	_go_fuzz_dep_.CoverTab[128795]++
									vals := make([]interface{}, 0, len(strs))
									for _, str := range strs {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:807
		_go_fuzz_dep_.CoverTab[128797]++
										val, err := parser(str)
										if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:809
			_go_fuzz_dep_.CoverTab[128799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:809
			return returnOnInvalid
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:809
			// _ = "end of CoverTab[128799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:809
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:809
			_go_fuzz_dep_.CoverTab[128800]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:810
			// _ = "end of CoverTab[128800]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:811
			_go_fuzz_dep_.CoverTab[128801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:811
			// _ = "end of CoverTab[128801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:811
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:811
		// _ = "end of CoverTab[128797]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:811
		_go_fuzz_dep_.CoverTab[128798]++
										if err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:812
			_go_fuzz_dep_.CoverTab[128802]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:812
			return addInvalid
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:812
			// _ = "end of CoverTab[128802]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:812
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:812
			_go_fuzz_dep_.CoverTab[128803]++
											vals = append(vals, val)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:813
			// _ = "end of CoverTab[128803]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:814
			_go_fuzz_dep_.CoverTab[128804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:814
			// _ = "end of CoverTab[128804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:814
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:814
		// _ = "end of CoverTab[128798]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:815
	// _ = "end of CoverTab[128795]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:815
	_go_fuzz_dep_.CoverTab[128796]++
									return vals, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:816
	// _ = "end of CoverTab[128796]"
}

// SetValue changes key value.
func (k *Key) SetValue(v string) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:820
	_go_fuzz_dep_.CoverTab[128805]++
									if k.s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:821
		_go_fuzz_dep_.CoverTab[128807]++
										k.s.f.lock.Lock()
										defer k.s.f.lock.Unlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:823
		// _ = "end of CoverTab[128807]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:824
		_go_fuzz_dep_.CoverTab[128808]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:824
		// _ = "end of CoverTab[128808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:824
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:824
	// _ = "end of CoverTab[128805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:824
	_go_fuzz_dep_.CoverTab[128806]++

									k.value = v
									k.s.keysHash[k.name] = v
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:827
	// _ = "end of CoverTab[128806]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:828
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/key.go:828
var _ = _go_fuzz_dep_.CoverTab
