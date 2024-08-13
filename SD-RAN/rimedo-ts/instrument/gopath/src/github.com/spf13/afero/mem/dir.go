// Copyright © 2014 Steve Francia <spf@spf13.com>.
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
package mem

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:14
)

type Dir interface {
	Len() int
	Names() []string
	Files() []*FileData
	Add(*FileData)
	Remove(*FileData)
}

func RemoveFromMemDir(dir *FileData, f *FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:24
	_go_fuzz_dep_.CoverTab[116795]++
										dir.memDir.Remove(f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:25
	// _ = "end of CoverTab[116795]"
}

func AddToMemDir(dir *FileData, f *FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:28
	_go_fuzz_dep_.CoverTab[116796]++
										dir.memDir.Add(f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:29
	// _ = "end of CoverTab[116796]"
}

func InitializeDir(d *FileData) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:32
	_go_fuzz_dep_.CoverTab[116797]++
										if d.memDir == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:33
		_go_fuzz_dep_.CoverTab[116798]++
											d.dir = true
											d.memDir = &DirMap{}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:35
		// _ = "end of CoverTab[116798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:36
		_go_fuzz_dep_.CoverTab[116799]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:36
		// _ = "end of CoverTab[116799]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:36
	// _ = "end of CoverTab[116797]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/dir.go:37
var _ = _go_fuzz_dep_.CoverTab
