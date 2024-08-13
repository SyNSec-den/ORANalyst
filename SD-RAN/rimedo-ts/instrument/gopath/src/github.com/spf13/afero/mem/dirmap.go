// Copyright © 2015 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
package mem

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:14
)

import "sort"

type DirMap map[string]*FileData

func (m DirMap) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:20
	_go_fuzz_dep_.CoverTab[116800]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:20
	return len(m)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:20
	// _ = "end of CoverTab[116800]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:20
}
func (m DirMap) Add(f *FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:21
	_go_fuzz_dep_.CoverTab[116801]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:21
	m[f.name] = f
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:21
	// _ = "end of CoverTab[116801]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:21
}
func (m DirMap) Remove(f *FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:22
	_go_fuzz_dep_.CoverTab[116802]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:22
	delete(m, f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:22
	// _ = "end of CoverTab[116802]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:22
}
func (m DirMap) Files() (files []*FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:23
	_go_fuzz_dep_.CoverTab[116803]++
											for _, f := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:24
		_go_fuzz_dep_.CoverTab[116805]++
												files = append(files, f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:25
		// _ = "end of CoverTab[116805]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:26
	// _ = "end of CoverTab[116803]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:26
	_go_fuzz_dep_.CoverTab[116804]++
											sort.Sort(filesSorter(files))
											return files
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:28
	// _ = "end of CoverTab[116804]"
}

// implement sort.Interface for []*FileData
type filesSorter []*FileData

func (s filesSorter) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:34
	_go_fuzz_dep_.CoverTab[116806]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:34
	return len(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:34
	// _ = "end of CoverTab[116806]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:34
}
func (s filesSorter) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:35
	_go_fuzz_dep_.CoverTab[116807]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:35
	s[i], s[j] = s[j], s[i]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:35
	// _ = "end of CoverTab[116807]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:35
}
func (s filesSorter) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:36
	_go_fuzz_dep_.CoverTab[116808]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:36
	return s[i].name < s[j].name
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:36
	// _ = "end of CoverTab[116808]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:36
}

func (m DirMap) Names() (names []string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:38
	_go_fuzz_dep_.CoverTab[116809]++
											for x := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:39
		_go_fuzz_dep_.CoverTab[116811]++
												names = append(names, x)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:40
		// _ = "end of CoverTab[116811]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:41
	// _ = "end of CoverTab[116809]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:41
	_go_fuzz_dep_.CoverTab[116810]++
											return names
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:42
	// _ = "end of CoverTab[116810]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dirmap.go:43
var _ = _go_fuzz_dep_.CoverTab
