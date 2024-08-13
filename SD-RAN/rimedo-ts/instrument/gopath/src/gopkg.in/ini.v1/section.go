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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:15
)

import (
	"errors"
	"fmt"
	"strings"
)

// Section represents a config section.
type Section struct {
	f		*File
	Comment		string
	name		string
	keys		map[string]*Key
	keyList		[]string
	keysHash	map[string]string

	isRawSection	bool
	rawBody		string
}

func newSection(f *File, name string) *Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:36
	_go_fuzz_dep_.CoverTab[129063]++
										return &Section{
		f:		f,
		name:		name,
		keys:		make(map[string]*Key),
		keyList:	make([]string, 0, 10),
		keysHash:	make(map[string]string),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:43
	// _ = "end of CoverTab[129063]"
}

// Name returns name of Section.
func (s *Section) Name() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:47
	_go_fuzz_dep_.CoverTab[129064]++
										return s.name
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:48
	// _ = "end of CoverTab[129064]"
}

// Body returns rawBody of Section if the section was marked as unparseable.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:51
// It still follows the other rules of the INI format surrounding leading/trailing whitespace.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:53
func (s *Section) Body() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:53
	_go_fuzz_dep_.CoverTab[129065]++
										return strings.TrimSpace(s.rawBody)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:54
	// _ = "end of CoverTab[129065]"
}

// SetBody updates body content only if section is raw.
func (s *Section) SetBody(body string) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:58
	_go_fuzz_dep_.CoverTab[129066]++
										if !s.isRawSection {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:59
		_go_fuzz_dep_.CoverTab[129068]++
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:60
		// _ = "end of CoverTab[129068]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:61
		_go_fuzz_dep_.CoverTab[129069]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:61
		// _ = "end of CoverTab[129069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:61
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:61
	// _ = "end of CoverTab[129066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:61
	_go_fuzz_dep_.CoverTab[129067]++
										s.rawBody = body
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:62
	// _ = "end of CoverTab[129067]"
}

// NewKey creates a new key to given section.
func (s *Section) NewKey(name, val string) (*Key, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:66
	_go_fuzz_dep_.CoverTab[129070]++
										if len(name) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:67
		_go_fuzz_dep_.CoverTab[129074]++
											return nil, errors.New("error creating new key: empty key name")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:68
		// _ = "end of CoverTab[129074]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
		_go_fuzz_dep_.CoverTab[129075]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
		if s.f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
			_go_fuzz_dep_.CoverTab[129076]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
			return s.f.options.InsensitiveKeys
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
			// _ = "end of CoverTab[129076]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:69
			_go_fuzz_dep_.CoverTab[129077]++
												name = strings.ToLower(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:70
			// _ = "end of CoverTab[129077]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
			_go_fuzz_dep_.CoverTab[129078]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
			// _ = "end of CoverTab[129078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
		// _ = "end of CoverTab[129075]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
	// _ = "end of CoverTab[129070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:71
	_go_fuzz_dep_.CoverTab[129071]++

										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:73
		_go_fuzz_dep_.CoverTab[129079]++
											s.f.lock.Lock()
											defer s.f.lock.Unlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:75
		// _ = "end of CoverTab[129079]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:76
		_go_fuzz_dep_.CoverTab[129080]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:76
		// _ = "end of CoverTab[129080]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:76
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:76
	// _ = "end of CoverTab[129071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:76
	_go_fuzz_dep_.CoverTab[129072]++

										if inSlice(name, s.keyList) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:78
		_go_fuzz_dep_.CoverTab[129081]++
											if s.f.options.AllowShadows {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:79
			_go_fuzz_dep_.CoverTab[129083]++
												if err := s.keys[name].addShadow(val); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:80
				_go_fuzz_dep_.CoverTab[129084]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:81
				// _ = "end of CoverTab[129084]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:82
				_go_fuzz_dep_.CoverTab[129085]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:82
				// _ = "end of CoverTab[129085]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:82
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:82
			// _ = "end of CoverTab[129083]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:83
			_go_fuzz_dep_.CoverTab[129086]++
												s.keys[name].value = val
												s.keysHash[name] = val
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:85
			// _ = "end of CoverTab[129086]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:86
		// _ = "end of CoverTab[129081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:86
		_go_fuzz_dep_.CoverTab[129082]++
											return s.keys[name], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:87
		// _ = "end of CoverTab[129082]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:88
		_go_fuzz_dep_.CoverTab[129087]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:88
		// _ = "end of CoverTab[129087]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:88
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:88
	// _ = "end of CoverTab[129072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:88
	_go_fuzz_dep_.CoverTab[129073]++

										s.keyList = append(s.keyList, name)
										s.keys[name] = newKey(s, name, val)
										s.keysHash[name] = val
										return s.keys[name], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:93
	// _ = "end of CoverTab[129073]"
}

// NewBooleanKey creates a new boolean type key to given section.
func (s *Section) NewBooleanKey(name string) (*Key, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:97
	_go_fuzz_dep_.CoverTab[129088]++
										key, err := s.NewKey(name, "true")
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:99
		_go_fuzz_dep_.CoverTab[129090]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:100
		// _ = "end of CoverTab[129090]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:101
		_go_fuzz_dep_.CoverTab[129091]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:101
		// _ = "end of CoverTab[129091]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:101
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:101
	// _ = "end of CoverTab[129088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:101
	_go_fuzz_dep_.CoverTab[129089]++

										key.isBooleanType = true
										return key, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:104
	// _ = "end of CoverTab[129089]"
}

// GetKey returns key in section by given name.
func (s *Section) GetKey(name string) (*Key, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:108
	_go_fuzz_dep_.CoverTab[129092]++
										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:109
		_go_fuzz_dep_.CoverTab[129097]++
											s.f.lock.RLock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:110
		// _ = "end of CoverTab[129097]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:111
		_go_fuzz_dep_.CoverTab[129098]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:111
		// _ = "end of CoverTab[129098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:111
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:111
	// _ = "end of CoverTab[129092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:111
	_go_fuzz_dep_.CoverTab[129093]++
										if s.f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:112
		_go_fuzz_dep_.CoverTab[129099]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:112
		return s.f.options.InsensitiveKeys
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:112
		// _ = "end of CoverTab[129099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:112
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:112
		_go_fuzz_dep_.CoverTab[129100]++
											name = strings.ToLower(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:113
		// _ = "end of CoverTab[129100]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:114
		_go_fuzz_dep_.CoverTab[129101]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:114
		// _ = "end of CoverTab[129101]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:114
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:114
	// _ = "end of CoverTab[129093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:114
	_go_fuzz_dep_.CoverTab[129094]++
										key := s.keys[name]
										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:116
		_go_fuzz_dep_.CoverTab[129102]++
											s.f.lock.RUnlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:117
		// _ = "end of CoverTab[129102]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:118
		_go_fuzz_dep_.CoverTab[129103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:118
		// _ = "end of CoverTab[129103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:118
	// _ = "end of CoverTab[129094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:118
	_go_fuzz_dep_.CoverTab[129095]++

										if key == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:120
		_go_fuzz_dep_.CoverTab[129104]++

											sname := s.name
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:123
			_go_fuzz_dep_.CoverTab[129106]++
												if i := strings.LastIndex(sname, s.f.options.ChildSectionDelimiter); i > -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:124
				_go_fuzz_dep_.CoverTab[129108]++
													sname = sname[:i]
													sec, err := s.f.GetSection(sname)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:127
					_go_fuzz_dep_.CoverTab[129110]++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:128
					// _ = "end of CoverTab[129110]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:129
					_go_fuzz_dep_.CoverTab[129111]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:129
					// _ = "end of CoverTab[129111]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:129
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:129
				// _ = "end of CoverTab[129108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:129
				_go_fuzz_dep_.CoverTab[129109]++
													return sec.GetKey(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:130
				// _ = "end of CoverTab[129109]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:131
				_go_fuzz_dep_.CoverTab[129112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:131
				// _ = "end of CoverTab[129112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:131
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:131
			// _ = "end of CoverTab[129106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:131
			_go_fuzz_dep_.CoverTab[129107]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:132
			// _ = "end of CoverTab[129107]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:133
		// _ = "end of CoverTab[129104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:133
		_go_fuzz_dep_.CoverTab[129105]++
											return nil, fmt.Errorf("error when getting key of section %q: key %q not exists", s.name, name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:134
		// _ = "end of CoverTab[129105]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:135
		_go_fuzz_dep_.CoverTab[129113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:135
		// _ = "end of CoverTab[129113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:135
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:135
	// _ = "end of CoverTab[129095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:135
	_go_fuzz_dep_.CoverTab[129096]++
										return key, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:136
	// _ = "end of CoverTab[129096]"
}

// HasKey returns true if section contains a key with given name.
func (s *Section) HasKey(name string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:140
	_go_fuzz_dep_.CoverTab[129114]++
										key, _ := s.GetKey(name)
										return key != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:142
	// _ = "end of CoverTab[129114]"
}

// Deprecated: Use "HasKey" instead.
func (s *Section) Haskey(name string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:146
	_go_fuzz_dep_.CoverTab[129115]++
										return s.HasKey(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:147
	// _ = "end of CoverTab[129115]"
}

// HasValue returns true if section contains given raw value.
func (s *Section) HasValue(value string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:151
	_go_fuzz_dep_.CoverTab[129116]++
										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:152
		_go_fuzz_dep_.CoverTab[129119]++
											s.f.lock.RLock()
											defer s.f.lock.RUnlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:154
		// _ = "end of CoverTab[129119]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:155
		_go_fuzz_dep_.CoverTab[129120]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:155
		// _ = "end of CoverTab[129120]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:155
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:155
	// _ = "end of CoverTab[129116]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:155
	_go_fuzz_dep_.CoverTab[129117]++

										for _, k := range s.keys {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:157
		_go_fuzz_dep_.CoverTab[129121]++
											if value == k.value {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:158
			_go_fuzz_dep_.CoverTab[129122]++
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:159
			// _ = "end of CoverTab[129122]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:160
			_go_fuzz_dep_.CoverTab[129123]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:160
			// _ = "end of CoverTab[129123]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:160
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:160
		// _ = "end of CoverTab[129121]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:161
	// _ = "end of CoverTab[129117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:161
	_go_fuzz_dep_.CoverTab[129118]++
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:162
	// _ = "end of CoverTab[129118]"
}

// Key assumes named Key exists in section and returns a zero-value when not.
func (s *Section) Key(name string) *Key {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:166
	_go_fuzz_dep_.CoverTab[129124]++
										key, err := s.GetKey(name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:168
		_go_fuzz_dep_.CoverTab[129126]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:171
		key, _ = s.NewKey(name, "")
											return key
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:172
		// _ = "end of CoverTab[129126]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:173
		_go_fuzz_dep_.CoverTab[129127]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:173
		// _ = "end of CoverTab[129127]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:173
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:173
	// _ = "end of CoverTab[129124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:173
	_go_fuzz_dep_.CoverTab[129125]++
										return key
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:174
	// _ = "end of CoverTab[129125]"
}

// Keys returns list of keys of section.
func (s *Section) Keys() []*Key {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:178
	_go_fuzz_dep_.CoverTab[129128]++
										keys := make([]*Key, len(s.keyList))
										for i := range s.keyList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:180
		_go_fuzz_dep_.CoverTab[129130]++
											keys[i] = s.Key(s.keyList[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:181
		// _ = "end of CoverTab[129130]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:182
	// _ = "end of CoverTab[129128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:182
	_go_fuzz_dep_.CoverTab[129129]++
										return keys
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:183
	// _ = "end of CoverTab[129129]"
}

// ParentKeys returns list of keys of parent section.
func (s *Section) ParentKeys() []*Key {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:187
	_go_fuzz_dep_.CoverTab[129131]++
										var parentKeys []*Key
										sname := s.name
										for {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:190
		_go_fuzz_dep_.CoverTab[129133]++
											if i := strings.LastIndex(sname, s.f.options.ChildSectionDelimiter); i > -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:191
			_go_fuzz_dep_.CoverTab[129134]++
												sname = sname[:i]
												sec, err := s.f.GetSection(sname)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:194
				_go_fuzz_dep_.CoverTab[129136]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:195
				// _ = "end of CoverTab[129136]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:196
				_go_fuzz_dep_.CoverTab[129137]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:196
				// _ = "end of CoverTab[129137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:196
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:196
			// _ = "end of CoverTab[129134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:196
			_go_fuzz_dep_.CoverTab[129135]++
												parentKeys = append(parentKeys, sec.Keys()...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:197
			// _ = "end of CoverTab[129135]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:198
			_go_fuzz_dep_.CoverTab[129138]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:199
			// _ = "end of CoverTab[129138]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:200
		// _ = "end of CoverTab[129133]"

	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:202
	// _ = "end of CoverTab[129131]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:202
	_go_fuzz_dep_.CoverTab[129132]++
										return parentKeys
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:203
	// _ = "end of CoverTab[129132]"
}

// KeyStrings returns list of key names of section.
func (s *Section) KeyStrings() []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:207
	_go_fuzz_dep_.CoverTab[129139]++
										list := make([]string, len(s.keyList))
										copy(list, s.keyList)
										return list
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:210
	// _ = "end of CoverTab[129139]"
}

// KeysHash returns keys hash consisting of names and values.
func (s *Section) KeysHash() map[string]string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:214
	_go_fuzz_dep_.CoverTab[129140]++
										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:215
		_go_fuzz_dep_.CoverTab[129143]++
											s.f.lock.RLock()
											defer s.f.lock.RUnlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:217
		// _ = "end of CoverTab[129143]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:218
		_go_fuzz_dep_.CoverTab[129144]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:218
		// _ = "end of CoverTab[129144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:218
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:218
	// _ = "end of CoverTab[129140]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:218
	_go_fuzz_dep_.CoverTab[129141]++

										hash := make(map[string]string, len(s.keysHash))
										for key, value := range s.keysHash {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:221
		_go_fuzz_dep_.CoverTab[129145]++
											hash[key] = value
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:222
		// _ = "end of CoverTab[129145]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:223
	// _ = "end of CoverTab[129141]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:223
	_go_fuzz_dep_.CoverTab[129142]++
										return hash
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:224
	// _ = "end of CoverTab[129142]"
}

// DeleteKey deletes a key from section.
func (s *Section) DeleteKey(name string) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:228
	_go_fuzz_dep_.CoverTab[129146]++
										if s.f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:229
		_go_fuzz_dep_.CoverTab[129148]++
											s.f.lock.Lock()
											defer s.f.lock.Unlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:231
		// _ = "end of CoverTab[129148]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:232
		_go_fuzz_dep_.CoverTab[129149]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:232
		// _ = "end of CoverTab[129149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:232
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:232
	// _ = "end of CoverTab[129146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:232
	_go_fuzz_dep_.CoverTab[129147]++

										for i, k := range s.keyList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:234
		_go_fuzz_dep_.CoverTab[129150]++
											if k == name {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:235
			_go_fuzz_dep_.CoverTab[129151]++
												s.keyList = append(s.keyList[:i], s.keyList[i+1:]...)
												delete(s.keys, name)
												delete(s.keysHash, name)
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:239
			// _ = "end of CoverTab[129151]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:240
			_go_fuzz_dep_.CoverTab[129152]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:240
			// _ = "end of CoverTab[129152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:240
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:240
		// _ = "end of CoverTab[129150]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:241
	// _ = "end of CoverTab[129147]"
}

// ChildSections returns a list of child sections of current section.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:244
// For example, "[parent.child1]" and "[parent.child12]" are child sections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:244
// of section "[parent]".
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:247
func (s *Section) ChildSections() []*Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:247
	_go_fuzz_dep_.CoverTab[129153]++
										prefix := s.name + s.f.options.ChildSectionDelimiter
										children := make([]*Section, 0, 3)
										for _, name := range s.f.sectionList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:250
		_go_fuzz_dep_.CoverTab[129155]++
											if strings.HasPrefix(name, prefix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:251
			_go_fuzz_dep_.CoverTab[129156]++
												children = append(children, s.f.sections[name]...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:252
			// _ = "end of CoverTab[129156]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:253
			_go_fuzz_dep_.CoverTab[129157]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:253
			// _ = "end of CoverTab[129157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:253
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:253
		// _ = "end of CoverTab[129155]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:254
	// _ = "end of CoverTab[129153]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:254
	_go_fuzz_dep_.CoverTab[129154]++
										return children
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:255
	// _ = "end of CoverTab[129154]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:256
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/section.go:256
var _ = _go_fuzz_dep_.CoverTab
