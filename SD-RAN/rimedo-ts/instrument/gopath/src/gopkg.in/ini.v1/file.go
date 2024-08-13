// Copyright 2017 Unknwon
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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:15
)

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// File represents a combination of one or more INI files in memory.
type File struct {
	options		LoadOptions
	dataSources	[]dataSource

	// Should make things safe, but sometimes doesn't matter.
	BlockMode	bool
	lock		sync.RWMutex

	// To keep data in order.
	sectionList	[]string
	// To keep track of the index of a section with same name.
	// This meta list is only used with non-unique section names are allowed.
	sectionIndexes	[]int

	// Actual data is stored here.
	sections	map[string][]*Section

	NameMapper
	ValueMapper
}

// newFile initializes File object with given data sources.
func newFile(dataSources []dataSource, opts LoadOptions) *File {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:51
	_go_fuzz_dep_.CoverTab[128247]++
									if len(opts.KeyValueDelimiters) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:52
		_go_fuzz_dep_.CoverTab[128251]++
										opts.KeyValueDelimiters = "=:"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:53
		// _ = "end of CoverTab[128251]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:54
		_go_fuzz_dep_.CoverTab[128252]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:54
		// _ = "end of CoverTab[128252]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:54
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:54
	// _ = "end of CoverTab[128247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:54
	_go_fuzz_dep_.CoverTab[128248]++
									if len(opts.KeyValueDelimiterOnWrite) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:55
		_go_fuzz_dep_.CoverTab[128253]++
										opts.KeyValueDelimiterOnWrite = "="
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:56
		// _ = "end of CoverTab[128253]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:57
		_go_fuzz_dep_.CoverTab[128254]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:57
		// _ = "end of CoverTab[128254]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:57
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:57
	// _ = "end of CoverTab[128248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:57
	_go_fuzz_dep_.CoverTab[128249]++
									if len(opts.ChildSectionDelimiter) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:58
		_go_fuzz_dep_.CoverTab[128255]++
										opts.ChildSectionDelimiter = "."
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:59
		// _ = "end of CoverTab[128255]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:60
		_go_fuzz_dep_.CoverTab[128256]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:60
		// _ = "end of CoverTab[128256]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:60
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:60
	// _ = "end of CoverTab[128249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:60
	_go_fuzz_dep_.CoverTab[128250]++

									return &File{
		BlockMode:	true,
		dataSources:	dataSources,
		sections:	make(map[string][]*Section),
		options:	opts,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:67
	// _ = "end of CoverTab[128250]"
}

// Empty returns an empty file object.
func Empty(opts ...LoadOptions) *File {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:71
	_go_fuzz_dep_.CoverTab[128257]++
									var opt LoadOptions
									if len(opts) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:73
		_go_fuzz_dep_.CoverTab[128259]++
										opt = opts[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:74
		// _ = "end of CoverTab[128259]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:75
		_go_fuzz_dep_.CoverTab[128260]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:75
		// _ = "end of CoverTab[128260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:75
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:75
	// _ = "end of CoverTab[128257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:75
	_go_fuzz_dep_.CoverTab[128258]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:78
	f, _ := LoadSources(opt, []byte(""))
									return f
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:79
	// _ = "end of CoverTab[128258]"
}

// NewSection creates a new section.
func (f *File) NewSection(name string) (*Section, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:83
	_go_fuzz_dep_.CoverTab[128261]++
									if len(name) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:84
		_go_fuzz_dep_.CoverTab[128266]++
										return nil, errors.New("empty section name")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:85
		// _ = "end of CoverTab[128266]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:86
		_go_fuzz_dep_.CoverTab[128267]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:86
		// _ = "end of CoverTab[128267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:86
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:86
	// _ = "end of CoverTab[128261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:86
	_go_fuzz_dep_.CoverTab[128262]++

									if (f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		_go_fuzz_dep_.CoverTab[128268]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		return f.options.InsensitiveSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		// _ = "end of CoverTab[128268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		_go_fuzz_dep_.CoverTab[128269]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		return name != DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		// _ = "end of CoverTab[128269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:88
		_go_fuzz_dep_.CoverTab[128270]++
										name = strings.ToLower(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:89
		// _ = "end of CoverTab[128270]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:90
		_go_fuzz_dep_.CoverTab[128271]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:90
		// _ = "end of CoverTab[128271]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:90
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:90
	// _ = "end of CoverTab[128262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:90
	_go_fuzz_dep_.CoverTab[128263]++

									if f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:92
		_go_fuzz_dep_.CoverTab[128272]++
										f.lock.Lock()
										defer f.lock.Unlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:94
		// _ = "end of CoverTab[128272]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:95
		_go_fuzz_dep_.CoverTab[128273]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:95
		// _ = "end of CoverTab[128273]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:95
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:95
	// _ = "end of CoverTab[128263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:95
	_go_fuzz_dep_.CoverTab[128264]++

									if !f.options.AllowNonUniqueSections && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:97
		_go_fuzz_dep_.CoverTab[128274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:97
		return inSlice(name, f.sectionList)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:97
		// _ = "end of CoverTab[128274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:97
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:97
		_go_fuzz_dep_.CoverTab[128275]++
										return f.sections[name][0], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:98
		// _ = "end of CoverTab[128275]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:99
		_go_fuzz_dep_.CoverTab[128276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:99
		// _ = "end of CoverTab[128276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:99
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:99
	// _ = "end of CoverTab[128264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:99
	_go_fuzz_dep_.CoverTab[128265]++

									f.sectionList = append(f.sectionList, name)

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:105
	f.sectionIndexes = append(f.sectionIndexes, len(f.sections[name]))

									sec := newSection(f, name)
									f.sections[name] = append(f.sections[name], sec)

									return sec, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:110
	// _ = "end of CoverTab[128265]"
}

// NewRawSection creates a new section with an unparseable body.
func (f *File) NewRawSection(name, body string) (*Section, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:114
	_go_fuzz_dep_.CoverTab[128277]++
									section, err := f.NewSection(name)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:116
		_go_fuzz_dep_.CoverTab[128279]++
										return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:117
		// _ = "end of CoverTab[128279]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:118
		_go_fuzz_dep_.CoverTab[128280]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:118
		// _ = "end of CoverTab[128280]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:118
	// _ = "end of CoverTab[128277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:118
	_go_fuzz_dep_.CoverTab[128278]++

									section.isRawSection = true
									section.rawBody = body
									return section, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:122
	// _ = "end of CoverTab[128278]"
}

// NewSections creates a list of sections.
func (f *File) NewSections(names ...string) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:126
	_go_fuzz_dep_.CoverTab[128281]++
									for _, name := range names {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:127
		_go_fuzz_dep_.CoverTab[128283]++
										if _, err = f.NewSection(name); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:128
			_go_fuzz_dep_.CoverTab[128284]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:129
			// _ = "end of CoverTab[128284]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:130
			_go_fuzz_dep_.CoverTab[128285]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:130
			// _ = "end of CoverTab[128285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:130
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:130
		// _ = "end of CoverTab[128283]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:131
	// _ = "end of CoverTab[128281]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:131
	_go_fuzz_dep_.CoverTab[128282]++
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:132
	// _ = "end of CoverTab[128282]"
}

// GetSection returns section by given name.
func (f *File) GetSection(name string) (*Section, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:136
	_go_fuzz_dep_.CoverTab[128286]++
									secs, err := f.SectionsByName(name)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:138
		_go_fuzz_dep_.CoverTab[128288]++
										return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:139
		// _ = "end of CoverTab[128288]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:140
		_go_fuzz_dep_.CoverTab[128289]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:140
		// _ = "end of CoverTab[128289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:140
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:140
	// _ = "end of CoverTab[128286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:140
	_go_fuzz_dep_.CoverTab[128287]++

									return secs[0], err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:142
	// _ = "end of CoverTab[128287]"
}

// SectionsByName returns all sections with given name.
func (f *File) SectionsByName(name string) ([]*Section, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:146
	_go_fuzz_dep_.CoverTab[128290]++
									if len(name) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:147
		_go_fuzz_dep_.CoverTab[128295]++
										name = DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:148
		// _ = "end of CoverTab[128295]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:149
		_go_fuzz_dep_.CoverTab[128296]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:149
		// _ = "end of CoverTab[128296]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:149
	// _ = "end of CoverTab[128290]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:149
	_go_fuzz_dep_.CoverTab[128291]++
									if f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:150
		_go_fuzz_dep_.CoverTab[128297]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:150
		return f.options.InsensitiveSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:150
		// _ = "end of CoverTab[128297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:150
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:150
		_go_fuzz_dep_.CoverTab[128298]++
										name = strings.ToLower(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:151
		// _ = "end of CoverTab[128298]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:152
		_go_fuzz_dep_.CoverTab[128299]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:152
		// _ = "end of CoverTab[128299]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:152
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:152
	// _ = "end of CoverTab[128291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:152
	_go_fuzz_dep_.CoverTab[128292]++

									if f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:154
		_go_fuzz_dep_.CoverTab[128300]++
										f.lock.RLock()
										defer f.lock.RUnlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:156
		// _ = "end of CoverTab[128300]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:157
		_go_fuzz_dep_.CoverTab[128301]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:157
		// _ = "end of CoverTab[128301]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:157
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:157
	// _ = "end of CoverTab[128292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:157
	_go_fuzz_dep_.CoverTab[128293]++

									secs := f.sections[name]
									if len(secs) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:160
		_go_fuzz_dep_.CoverTab[128302]++
										return nil, fmt.Errorf("section %q does not exist", name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:161
		// _ = "end of CoverTab[128302]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:162
		_go_fuzz_dep_.CoverTab[128303]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:162
		// _ = "end of CoverTab[128303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:162
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:162
	// _ = "end of CoverTab[128293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:162
	_go_fuzz_dep_.CoverTab[128294]++

									return secs, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:164
	// _ = "end of CoverTab[128294]"
}

// Section assumes named section exists and returns a zero-value when not.
func (f *File) Section(name string) *Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:168
	_go_fuzz_dep_.CoverTab[128304]++
									sec, err := f.GetSection(name)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:170
		_go_fuzz_dep_.CoverTab[128306]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:173
		sec, _ = f.NewSection(name)
										return sec
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:174
		// _ = "end of CoverTab[128306]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:175
		_go_fuzz_dep_.CoverTab[128307]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:175
		// _ = "end of CoverTab[128307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:175
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:175
	// _ = "end of CoverTab[128304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:175
	_go_fuzz_dep_.CoverTab[128305]++
									return sec
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:176
	// _ = "end of CoverTab[128305]"
}

// SectionWithIndex assumes named section exists and returns a new section when not.
func (f *File) SectionWithIndex(name string, index int) *Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:180
	_go_fuzz_dep_.CoverTab[128308]++
									secs, err := f.SectionsByName(name)
									if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:182
		_go_fuzz_dep_.CoverTab[128310]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:182
		return len(secs) <= index
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:182
		// _ = "end of CoverTab[128310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:182
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:182
		_go_fuzz_dep_.CoverTab[128311]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:185
		newSec, _ := f.NewSection(name)
										return newSec
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:186
		// _ = "end of CoverTab[128311]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:187
		_go_fuzz_dep_.CoverTab[128312]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:187
		// _ = "end of CoverTab[128312]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:187
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:187
	// _ = "end of CoverTab[128308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:187
	_go_fuzz_dep_.CoverTab[128309]++

									return secs[index]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:189
	// _ = "end of CoverTab[128309]"
}

// Sections returns a list of Section stored in the current instance.
func (f *File) Sections() []*Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:193
	_go_fuzz_dep_.CoverTab[128313]++
									if f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:194
		_go_fuzz_dep_.CoverTab[128316]++
										f.lock.RLock()
										defer f.lock.RUnlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:196
		// _ = "end of CoverTab[128316]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:197
		_go_fuzz_dep_.CoverTab[128317]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:197
		// _ = "end of CoverTab[128317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:197
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:197
	// _ = "end of CoverTab[128313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:197
	_go_fuzz_dep_.CoverTab[128314]++

									sections := make([]*Section, len(f.sectionList))
									for i, name := range f.sectionList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:200
		_go_fuzz_dep_.CoverTab[128318]++
										sections[i] = f.sections[name][f.sectionIndexes[i]]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:201
		// _ = "end of CoverTab[128318]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:202
	// _ = "end of CoverTab[128314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:202
	_go_fuzz_dep_.CoverTab[128315]++
									return sections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:203
	// _ = "end of CoverTab[128315]"
}

// ChildSections returns a list of child sections of given section name.
func (f *File) ChildSections(name string) []*Section {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:207
	_go_fuzz_dep_.CoverTab[128319]++
									return f.Section(name).ChildSections()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:208
	// _ = "end of CoverTab[128319]"
}

// SectionStrings returns list of section names.
func (f *File) SectionStrings() []string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:212
	_go_fuzz_dep_.CoverTab[128320]++
									list := make([]string, len(f.sectionList))
									copy(list, f.sectionList)
									return list
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:215
	// _ = "end of CoverTab[128320]"
}

// DeleteSection deletes a section or all sections with given name.
func (f *File) DeleteSection(name string) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:219
	_go_fuzz_dep_.CoverTab[128321]++
									secs, err := f.SectionsByName(name)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:221
		_go_fuzz_dep_.CoverTab[128323]++
										return
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:222
		// _ = "end of CoverTab[128323]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:223
		_go_fuzz_dep_.CoverTab[128324]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:223
		// _ = "end of CoverTab[128324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:223
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:223
	// _ = "end of CoverTab[128321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:223
	_go_fuzz_dep_.CoverTab[128322]++

									for i := 0; i < len(secs); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:225
		_go_fuzz_dep_.CoverTab[128325]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:229
		_ = f.DeleteSectionWithIndex(name, 0)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:229
		// _ = "end of CoverTab[128325]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:230
	// _ = "end of CoverTab[128322]"
}

// DeleteSectionWithIndex deletes a section with given name and index.
func (f *File) DeleteSectionWithIndex(name string, index int) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:234
	_go_fuzz_dep_.CoverTab[128326]++
									if !f.options.AllowNonUniqueSections && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:235
		_go_fuzz_dep_.CoverTab[128332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:235
		return index != 0
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:235
		// _ = "end of CoverTab[128332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:235
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:235
		_go_fuzz_dep_.CoverTab[128333]++
										return fmt.Errorf("delete section with non-zero index is only allowed when non-unique sections is enabled")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:236
		// _ = "end of CoverTab[128333]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:237
		_go_fuzz_dep_.CoverTab[128334]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:237
		// _ = "end of CoverTab[128334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:237
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:237
	// _ = "end of CoverTab[128326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:237
	_go_fuzz_dep_.CoverTab[128327]++

									if len(name) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:239
		_go_fuzz_dep_.CoverTab[128335]++
										name = DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:240
		// _ = "end of CoverTab[128335]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:241
		_go_fuzz_dep_.CoverTab[128336]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:241
		// _ = "end of CoverTab[128336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:241
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:241
	// _ = "end of CoverTab[128327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:241
	_go_fuzz_dep_.CoverTab[128328]++
									if f.options.Insensitive || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:242
		_go_fuzz_dep_.CoverTab[128337]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:242
		return f.options.InsensitiveSections
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:242
		// _ = "end of CoverTab[128337]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:242
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:242
		_go_fuzz_dep_.CoverTab[128338]++
										name = strings.ToLower(name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:243
		// _ = "end of CoverTab[128338]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:244
		_go_fuzz_dep_.CoverTab[128339]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:244
		// _ = "end of CoverTab[128339]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:244
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:244
	// _ = "end of CoverTab[128328]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:244
	_go_fuzz_dep_.CoverTab[128329]++

									if f.BlockMode {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:246
		_go_fuzz_dep_.CoverTab[128340]++
										f.lock.Lock()
										defer f.lock.Unlock()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:248
		// _ = "end of CoverTab[128340]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:249
		_go_fuzz_dep_.CoverTab[128341]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:249
		// _ = "end of CoverTab[128341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:249
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:249
	// _ = "end of CoverTab[128329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:249
	_go_fuzz_dep_.CoverTab[128330]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:252
	occurrences := 0

	sectionListCopy := make([]string, len(f.sectionList))
	copy(sectionListCopy, f.sectionList)

	for i, s := range sectionListCopy {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:257
		_go_fuzz_dep_.CoverTab[128342]++
										if s != name {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:258
			_go_fuzz_dep_.CoverTab[128345]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:259
			// _ = "end of CoverTab[128345]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:260
			_go_fuzz_dep_.CoverTab[128346]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:260
			// _ = "end of CoverTab[128346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:260
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:260
		// _ = "end of CoverTab[128342]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:260
		_go_fuzz_dep_.CoverTab[128343]++

										if occurrences == index {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:262
			_go_fuzz_dep_.CoverTab[128347]++
											if len(f.sections[name]) <= 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:263
				_go_fuzz_dep_.CoverTab[128349]++
												delete(f.sections, name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:264
				// _ = "end of CoverTab[128349]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:265
				_go_fuzz_dep_.CoverTab[128350]++
												f.sections[name] = append(f.sections[name][:index], f.sections[name][index+1:]...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:266
				// _ = "end of CoverTab[128350]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:267
			// _ = "end of CoverTab[128347]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:267
			_go_fuzz_dep_.CoverTab[128348]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:270
			f.sectionList = append(f.sectionList[:i], f.sectionList[i+1:]...)
											f.sectionIndexes = append(f.sectionIndexes[:i], f.sectionIndexes[i+1:]...)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:271
			// _ = "end of CoverTab[128348]"

		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:273
			_go_fuzz_dep_.CoverTab[128351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:273
			if occurrences > index {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:273
				_go_fuzz_dep_.CoverTab[128352]++

												f.sectionIndexes[i-1]--
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:275
				// _ = "end of CoverTab[128352]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
				_go_fuzz_dep_.CoverTab[128353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
				// _ = "end of CoverTab[128353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
			// _ = "end of CoverTab[128351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
		// _ = "end of CoverTab[128343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:276
		_go_fuzz_dep_.CoverTab[128344]++

										occurrences++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:278
		// _ = "end of CoverTab[128344]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:279
	// _ = "end of CoverTab[128330]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:279
	_go_fuzz_dep_.CoverTab[128331]++

									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:281
	// _ = "end of CoverTab[128331]"
}

func (f *File) reload(s dataSource) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:284
	_go_fuzz_dep_.CoverTab[128354]++
									r, err := s.ReadCloser()
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:286
		_go_fuzz_dep_.CoverTab[128356]++
										return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:287
		// _ = "end of CoverTab[128356]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:288
		_go_fuzz_dep_.CoverTab[128357]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:288
		// _ = "end of CoverTab[128357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:288
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:288
	// _ = "end of CoverTab[128354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:288
	_go_fuzz_dep_.CoverTab[128355]++
									defer r.Close()

									return f.parse(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:291
	// _ = "end of CoverTab[128355]"
}

// Reload reloads and parses all data sources.
func (f *File) Reload() (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:295
	_go_fuzz_dep_.CoverTab[128358]++
									for _, s := range f.dataSources {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:296
		_go_fuzz_dep_.CoverTab[128360]++
										if err = f.reload(s); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:297
			_go_fuzz_dep_.CoverTab[128362]++

											if os.IsNotExist(err) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:299
				_go_fuzz_dep_.CoverTab[128364]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:299
				return f.options.Loose
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:299
				// _ = "end of CoverTab[128364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:299
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:299
				_go_fuzz_dep_.CoverTab[128365]++
												_ = f.parse(bytes.NewBuffer(nil))
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:301
				// _ = "end of CoverTab[128365]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:302
				_go_fuzz_dep_.CoverTab[128366]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:302
				// _ = "end of CoverTab[128366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:302
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:302
			// _ = "end of CoverTab[128362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:302
			_go_fuzz_dep_.CoverTab[128363]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:303
			// _ = "end of CoverTab[128363]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:304
			_go_fuzz_dep_.CoverTab[128367]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:304
			// _ = "end of CoverTab[128367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:304
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:304
		// _ = "end of CoverTab[128360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:304
		_go_fuzz_dep_.CoverTab[128361]++
										if f.options.ShortCircuit {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:305
			_go_fuzz_dep_.CoverTab[128368]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:306
			// _ = "end of CoverTab[128368]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:307
			_go_fuzz_dep_.CoverTab[128369]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:307
			// _ = "end of CoverTab[128369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:307
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:307
		// _ = "end of CoverTab[128361]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:308
	// _ = "end of CoverTab[128358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:308
	_go_fuzz_dep_.CoverTab[128359]++
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:309
	// _ = "end of CoverTab[128359]"
}

// Append appends one or more data sources and reloads automatically.
func (f *File) Append(source interface{}, others ...interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:313
	_go_fuzz_dep_.CoverTab[128370]++
									ds, err := parseDataSource(source)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:315
		_go_fuzz_dep_.CoverTab[128373]++
										return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:316
		// _ = "end of CoverTab[128373]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:317
		_go_fuzz_dep_.CoverTab[128374]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:317
		// _ = "end of CoverTab[128374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:317
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:317
	// _ = "end of CoverTab[128370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:317
	_go_fuzz_dep_.CoverTab[128371]++
									f.dataSources = append(f.dataSources, ds)
									for _, s := range others {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:319
		_go_fuzz_dep_.CoverTab[128375]++
										ds, err = parseDataSource(s)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:321
			_go_fuzz_dep_.CoverTab[128377]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:322
			// _ = "end of CoverTab[128377]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:323
			_go_fuzz_dep_.CoverTab[128378]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:323
			// _ = "end of CoverTab[128378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:323
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:323
		// _ = "end of CoverTab[128375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:323
		_go_fuzz_dep_.CoverTab[128376]++
										f.dataSources = append(f.dataSources, ds)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:324
		// _ = "end of CoverTab[128376]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:325
	// _ = "end of CoverTab[128371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:325
	_go_fuzz_dep_.CoverTab[128372]++
									return f.Reload()
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:326
	// _ = "end of CoverTab[128372]"
}

func (f *File) writeToBuffer(indent string) (*bytes.Buffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:329
	_go_fuzz_dep_.CoverTab[128379]++
									equalSign := DefaultFormatLeft + f.options.KeyValueDelimiterOnWrite + DefaultFormatRight

									if PrettyFormat || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:332
		_go_fuzz_dep_.CoverTab[128382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:332
		return PrettyEqual
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:332
		// _ = "end of CoverTab[128382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:332
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:332
		_go_fuzz_dep_.CoverTab[128383]++
										equalSign = fmt.Sprintf(" %s ", f.options.KeyValueDelimiterOnWrite)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:333
		// _ = "end of CoverTab[128383]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:334
		_go_fuzz_dep_.CoverTab[128384]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:334
		// _ = "end of CoverTab[128384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:334
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:334
	// _ = "end of CoverTab[128379]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:334
	_go_fuzz_dep_.CoverTab[128380]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:337
	buf := bytes.NewBuffer(nil)
	for i, sname := range f.sectionList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:338
		_go_fuzz_dep_.CoverTab[128385]++
										sec := f.SectionWithIndex(sname, f.sectionIndexes[i])
										if len(sec.Comment) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:340
			_go_fuzz_dep_.CoverTab[128391]++

											lines := strings.Split(sec.Comment, LineBreak)
											for i := range lines {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:343
				_go_fuzz_dep_.CoverTab[128392]++
												if lines[i][0] != '#' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:344
					_go_fuzz_dep_.CoverTab[128394]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:344
					return lines[i][0] != ';'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:344
					// _ = "end of CoverTab[128394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:344
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:344
					_go_fuzz_dep_.CoverTab[128395]++
													lines[i] = "; " + lines[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:345
					// _ = "end of CoverTab[128395]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:346
					_go_fuzz_dep_.CoverTab[128396]++
													lines[i] = lines[i][:1] + " " + strings.TrimSpace(lines[i][1:])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:347
					// _ = "end of CoverTab[128396]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:348
				// _ = "end of CoverTab[128392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:348
				_go_fuzz_dep_.CoverTab[128393]++

												if _, err := buf.WriteString(lines[i] + LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:350
					_go_fuzz_dep_.CoverTab[128397]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:351
					// _ = "end of CoverTab[128397]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:352
					_go_fuzz_dep_.CoverTab[128398]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:352
					// _ = "end of CoverTab[128398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:352
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:352
				// _ = "end of CoverTab[128393]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:353
			// _ = "end of CoverTab[128391]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:354
			_go_fuzz_dep_.CoverTab[128399]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:354
			// _ = "end of CoverTab[128399]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:354
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:354
		// _ = "end of CoverTab[128385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:354
		_go_fuzz_dep_.CoverTab[128386]++

										if i > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			_go_fuzz_dep_.CoverTab[128400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			return DefaultHeader
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			// _ = "end of CoverTab[128400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			_go_fuzz_dep_.CoverTab[128401]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			return (i == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
				_go_fuzz_dep_.CoverTab[128402]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
				return strings.ToUpper(sec.name) != DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
				// _ = "end of CoverTab[128402]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			// _ = "end of CoverTab[128401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:356
			_go_fuzz_dep_.CoverTab[128403]++
											if _, err := buf.WriteString("[" + sname + "]" + LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:357
				_go_fuzz_dep_.CoverTab[128404]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:358
				// _ = "end of CoverTab[128404]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:359
				_go_fuzz_dep_.CoverTab[128405]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:359
				// _ = "end of CoverTab[128405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:359
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:359
			// _ = "end of CoverTab[128403]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:360
			_go_fuzz_dep_.CoverTab[128406]++

											if len(sec.keyList) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:362
				_go_fuzz_dep_.CoverTab[128407]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:363
				// _ = "end of CoverTab[128407]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:364
				_go_fuzz_dep_.CoverTab[128408]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:364
				// _ = "end of CoverTab[128408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:364
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:364
			// _ = "end of CoverTab[128406]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:365
		// _ = "end of CoverTab[128386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:365
		_go_fuzz_dep_.CoverTab[128387]++

										if sec.isRawSection {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:367
			_go_fuzz_dep_.CoverTab[128409]++
											if _, err := buf.WriteString(sec.rawBody); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:368
				_go_fuzz_dep_.CoverTab[128412]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:369
				// _ = "end of CoverTab[128412]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:370
				_go_fuzz_dep_.CoverTab[128413]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:370
				// _ = "end of CoverTab[128413]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:370
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:370
			// _ = "end of CoverTab[128409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:370
			_go_fuzz_dep_.CoverTab[128410]++

											if PrettySection {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:372
				_go_fuzz_dep_.CoverTab[128414]++

												if _, err := buf.WriteString(LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:374
					_go_fuzz_dep_.CoverTab[128415]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:375
					// _ = "end of CoverTab[128415]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:376
					_go_fuzz_dep_.CoverTab[128416]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:376
					// _ = "end of CoverTab[128416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:376
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:376
				// _ = "end of CoverTab[128414]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:377
				_go_fuzz_dep_.CoverTab[128417]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:377
				// _ = "end of CoverTab[128417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:377
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:377
			// _ = "end of CoverTab[128410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:377
			_go_fuzz_dep_.CoverTab[128411]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:378
			// _ = "end of CoverTab[128411]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:379
			_go_fuzz_dep_.CoverTab[128418]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:379
			// _ = "end of CoverTab[128418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:379
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:379
		// _ = "end of CoverTab[128387]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:379
		_go_fuzz_dep_.CoverTab[128388]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:384
		alignLength := 0
		if PrettyFormat {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:385
			_go_fuzz_dep_.CoverTab[128419]++
											for _, kname := range sec.keyList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:386
				_go_fuzz_dep_.CoverTab[128420]++
												keyLength := len(kname)

												if strings.Contains(kname, "\"") || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:389
					_go_fuzz_dep_.CoverTab[128422]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:389
					return strings.ContainsAny(kname, f.options.KeyValueDelimiters)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:389
					// _ = "end of CoverTab[128422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:389
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:389
					_go_fuzz_dep_.CoverTab[128423]++
													keyLength += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:390
					// _ = "end of CoverTab[128423]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:391
					_go_fuzz_dep_.CoverTab[128424]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:391
					if strings.Contains(kname, "`") {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:391
						_go_fuzz_dep_.CoverTab[128425]++
														keyLength += 6
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:392
						// _ = "end of CoverTab[128425]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
						_go_fuzz_dep_.CoverTab[128426]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
						// _ = "end of CoverTab[128426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
					// _ = "end of CoverTab[128424]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
				// _ = "end of CoverTab[128420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:393
				_go_fuzz_dep_.CoverTab[128421]++

												if keyLength > alignLength {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:395
					_go_fuzz_dep_.CoverTab[128427]++
													alignLength = keyLength
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:396
					// _ = "end of CoverTab[128427]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:397
					_go_fuzz_dep_.CoverTab[128428]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:397
					// _ = "end of CoverTab[128428]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:397
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:397
				// _ = "end of CoverTab[128421]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:398
			// _ = "end of CoverTab[128419]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:399
			_go_fuzz_dep_.CoverTab[128429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:399
			// _ = "end of CoverTab[128429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:399
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:399
		// _ = "end of CoverTab[128388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:399
		_go_fuzz_dep_.CoverTab[128389]++
										alignSpaces := bytes.Repeat([]byte(" "), alignLength)

	KeyList:
		for _, kname := range sec.keyList {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:403
			_go_fuzz_dep_.CoverTab[128430]++
											key := sec.Key(kname)
											if len(key.Comment) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:405
				_go_fuzz_dep_.CoverTab[128435]++
												if len(indent) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:406
					_go_fuzz_dep_.CoverTab[128437]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:406
					return sname != DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:406
					// _ = "end of CoverTab[128437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:406
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:406
					_go_fuzz_dep_.CoverTab[128438]++
													buf.WriteString(indent)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:407
					// _ = "end of CoverTab[128438]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:408
					_go_fuzz_dep_.CoverTab[128439]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:408
					// _ = "end of CoverTab[128439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:408
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:408
				// _ = "end of CoverTab[128435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:408
				_go_fuzz_dep_.CoverTab[128436]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:411
				lines := strings.Split(key.Comment, LineBreak)
				for i := range lines {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:412
					_go_fuzz_dep_.CoverTab[128440]++
													if lines[i][0] != '#' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:413
						_go_fuzz_dep_.CoverTab[128442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:413
						return lines[i][0] != ';'
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:413
						// _ = "end of CoverTab[128442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:413
					}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:413
						_go_fuzz_dep_.CoverTab[128443]++
														lines[i] = "; " + strings.TrimSpace(lines[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:414
						// _ = "end of CoverTab[128443]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:415
						_go_fuzz_dep_.CoverTab[128444]++
														lines[i] = lines[i][:1] + " " + strings.TrimSpace(lines[i][1:])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:416
						// _ = "end of CoverTab[128444]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:417
					// _ = "end of CoverTab[128440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:417
					_go_fuzz_dep_.CoverTab[128441]++

													if _, err := buf.WriteString(lines[i] + LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:419
						_go_fuzz_dep_.CoverTab[128445]++
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:420
						// _ = "end of CoverTab[128445]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:421
						_go_fuzz_dep_.CoverTab[128446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:421
						// _ = "end of CoverTab[128446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:421
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:421
					// _ = "end of CoverTab[128441]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:422
				// _ = "end of CoverTab[128436]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:423
				_go_fuzz_dep_.CoverTab[128447]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:423
				// _ = "end of CoverTab[128447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:423
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:423
			// _ = "end of CoverTab[128430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:423
			_go_fuzz_dep_.CoverTab[128431]++

											if len(indent) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:425
				_go_fuzz_dep_.CoverTab[128448]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:425
				return sname != DefaultSection
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:425
				// _ = "end of CoverTab[128448]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:425
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:425
				_go_fuzz_dep_.CoverTab[128449]++
												buf.WriteString(indent)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:426
				// _ = "end of CoverTab[128449]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:427
				_go_fuzz_dep_.CoverTab[128450]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:427
				// _ = "end of CoverTab[128450]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:427
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:427
			// _ = "end of CoverTab[128431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:427
			_go_fuzz_dep_.CoverTab[128432]++

											switch {
			case key.isAutoIncrement:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:430
				_go_fuzz_dep_.CoverTab[128451]++
												kname = "-"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:431
				// _ = "end of CoverTab[128451]"
			case strings.Contains(kname, "\"") || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:432
				_go_fuzz_dep_.CoverTab[128455]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:432
				return strings.ContainsAny(kname, f.options.KeyValueDelimiters)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:432
				// _ = "end of CoverTab[128455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:432
			}():
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:432
				_go_fuzz_dep_.CoverTab[128452]++
												kname = "`" + kname + "`"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:433
				// _ = "end of CoverTab[128452]"
			case strings.Contains(kname, "`"):
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:434
				_go_fuzz_dep_.CoverTab[128453]++
												kname = `"""` + kname + `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:435
				// _ = "end of CoverTab[128453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:435
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:435
				_go_fuzz_dep_.CoverTab[128454]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:435
				// _ = "end of CoverTab[128454]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:436
			// _ = "end of CoverTab[128432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:436
			_go_fuzz_dep_.CoverTab[128433]++

											for _, val := range key.ValueWithShadows() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:438
				_go_fuzz_dep_.CoverTab[128456]++
												if _, err := buf.WriteString(kname); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:439
					_go_fuzz_dep_.CoverTab[128461]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:440
					// _ = "end of CoverTab[128461]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:441
					_go_fuzz_dep_.CoverTab[128462]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:441
					// _ = "end of CoverTab[128462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:441
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:441
				// _ = "end of CoverTab[128456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:441
				_go_fuzz_dep_.CoverTab[128457]++

												if key.isBooleanType {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:443
					_go_fuzz_dep_.CoverTab[128463]++
													if kname != sec.keyList[len(sec.keyList)-1] {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:444
						_go_fuzz_dep_.CoverTab[128465]++
														buf.WriteString(LineBreak)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:445
						// _ = "end of CoverTab[128465]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:446
						_go_fuzz_dep_.CoverTab[128466]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:446
						// _ = "end of CoverTab[128466]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:446
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:446
					// _ = "end of CoverTab[128463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:446
					_go_fuzz_dep_.CoverTab[128464]++
													continue KeyList
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:447
					// _ = "end of CoverTab[128464]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:448
					_go_fuzz_dep_.CoverTab[128467]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:448
					// _ = "end of CoverTab[128467]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:448
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:448
				// _ = "end of CoverTab[128457]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:448
				_go_fuzz_dep_.CoverTab[128458]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:451
				if PrettyFormat {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:451
					_go_fuzz_dep_.CoverTab[128468]++
													buf.Write(alignSpaces[:alignLength-len(kname)])
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:452
					// _ = "end of CoverTab[128468]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:453
					_go_fuzz_dep_.CoverTab[128469]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:453
					// _ = "end of CoverTab[128469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:453
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:453
				// _ = "end of CoverTab[128458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:453
				_go_fuzz_dep_.CoverTab[128459]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:456
				if strings.ContainsAny(val, "\n`") {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:456
					_go_fuzz_dep_.CoverTab[128470]++
													val = `"""` + val + `"""`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:457
					// _ = "end of CoverTab[128470]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
					_go_fuzz_dep_.CoverTab[128471]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
					if !f.options.IgnoreInlineComment && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
						_go_fuzz_dep_.CoverTab[128472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
						return strings.ContainsAny(val, "#;")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
						// _ = "end of CoverTab[128472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
					}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:458
						_go_fuzz_dep_.CoverTab[128473]++
														val = "`" + val + "`"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:459
						// _ = "end of CoverTab[128473]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:460
						_go_fuzz_dep_.CoverTab[128474]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:460
						if len(strings.TrimSpace(val)) != len(val) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:460
							_go_fuzz_dep_.CoverTab[128475]++
															val = `"` + val + `"`
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:461
							// _ = "end of CoverTab[128475]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
							_go_fuzz_dep_.CoverTab[128476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
							// _ = "end of CoverTab[128476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
						// _ = "end of CoverTab[128474]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
					// _ = "end of CoverTab[128471]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
				// _ = "end of CoverTab[128459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:462
				_go_fuzz_dep_.CoverTab[128460]++
												if _, err := buf.WriteString(equalSign + val + LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:463
					_go_fuzz_dep_.CoverTab[128477]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:464
					// _ = "end of CoverTab[128477]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:465
					_go_fuzz_dep_.CoverTab[128478]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:465
					// _ = "end of CoverTab[128478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:465
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:465
				// _ = "end of CoverTab[128460]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:466
			// _ = "end of CoverTab[128433]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:466
			_go_fuzz_dep_.CoverTab[128434]++

											for _, val := range key.nestedValues {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:468
				_go_fuzz_dep_.CoverTab[128479]++
												if _, err := buf.WriteString(indent + "  " + val + LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:469
					_go_fuzz_dep_.CoverTab[128480]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:470
					// _ = "end of CoverTab[128480]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:471
					_go_fuzz_dep_.CoverTab[128481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:471
					// _ = "end of CoverTab[128481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:471
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:471
				// _ = "end of CoverTab[128479]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:472
			// _ = "end of CoverTab[128434]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:473
		// _ = "end of CoverTab[128389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:473
		_go_fuzz_dep_.CoverTab[128390]++

										if PrettySection {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:475
			_go_fuzz_dep_.CoverTab[128482]++

											if _, err := buf.WriteString(LineBreak); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:477
				_go_fuzz_dep_.CoverTab[128483]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:478
				// _ = "end of CoverTab[128483]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:479
				_go_fuzz_dep_.CoverTab[128484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:479
				// _ = "end of CoverTab[128484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:479
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:479
			// _ = "end of CoverTab[128482]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:480
			_go_fuzz_dep_.CoverTab[128485]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:480
			// _ = "end of CoverTab[128485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:480
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:480
		// _ = "end of CoverTab[128390]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:481
	// _ = "end of CoverTab[128380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:481
	_go_fuzz_dep_.CoverTab[128381]++

									return buf, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:483
	// _ = "end of CoverTab[128381]"
}

// WriteToIndent writes content into io.Writer with given indention.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:486
// If PrettyFormat has been set to be true,
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:486
// it will align "=" sign with spaces under each section.
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:489
func (f *File) WriteToIndent(w io.Writer, indent string) (int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:489
	_go_fuzz_dep_.CoverTab[128486]++
									buf, err := f.writeToBuffer(indent)
									if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:491
		_go_fuzz_dep_.CoverTab[128488]++
										return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:492
		// _ = "end of CoverTab[128488]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:493
		_go_fuzz_dep_.CoverTab[128489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:493
		// _ = "end of CoverTab[128489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:493
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:493
	// _ = "end of CoverTab[128486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:493
	_go_fuzz_dep_.CoverTab[128487]++
									return buf.WriteTo(w)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:494
	// _ = "end of CoverTab[128487]"
}

// WriteTo writes file content into io.Writer.
func (f *File) WriteTo(w io.Writer) (int64, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:498
	_go_fuzz_dep_.CoverTab[128490]++
									return f.WriteToIndent(w, "")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:499
	// _ = "end of CoverTab[128490]"
}

// SaveToIndent writes content to file system with given value indention.
func (f *File) SaveToIndent(filename, indent string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:503
	_go_fuzz_dep_.CoverTab[128491]++

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:506
	buf, err := f.writeToBuffer(indent)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:507
		_go_fuzz_dep_.CoverTab[128493]++
										return err
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:508
		// _ = "end of CoverTab[128493]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:509
		_go_fuzz_dep_.CoverTab[128494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:509
		// _ = "end of CoverTab[128494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:509
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:509
	// _ = "end of CoverTab[128491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:509
	_go_fuzz_dep_.CoverTab[128492]++

									return ioutil.WriteFile(filename, buf.Bytes(), 0666)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:511
	// _ = "end of CoverTab[128492]"
}

// SaveTo writes content to file system.
func (f *File) SaveTo(filename string) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:515
	_go_fuzz_dep_.CoverTab[128495]++
									return f.SaveToIndent(filename, "")
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:516
	// _ = "end of CoverTab[128495]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:517
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/file.go:517
var _ = _go_fuzz_dep_.CoverTab
