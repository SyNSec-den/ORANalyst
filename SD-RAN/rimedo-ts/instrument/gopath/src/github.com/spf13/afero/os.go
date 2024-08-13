// Copyright © 2014 Steve Francia <spf@spf13.com>.
// Copyright 2013 tsuru authors. All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:15
)

import (
	"os"
	"time"
)

var _ Lstater = (*OsFs)(nil)

// OsFs is a Fs implementation that uses functions provided by the os package.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:24
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:24
// For details in any method, check the documentation of the os package
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:24
// (http://golang.org/pkg/os/).
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:28
type OsFs struct{}

func NewOsFs() Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:30
	_go_fuzz_dep_.CoverTab[117991]++
										return &OsFs{}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:31
	// _ = "end of CoverTab[117991]"
}

func (OsFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:34
	_go_fuzz_dep_.CoverTab[117992]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:34
	return "OsFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:34
	// _ = "end of CoverTab[117992]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:34
}

func (OsFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:36
	_go_fuzz_dep_.CoverTab[117993]++
										f, e := os.Create(name)
										if f == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:38
		_go_fuzz_dep_.CoverTab[117995]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:41
		return nil, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:41
		// _ = "end of CoverTab[117995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:42
		_go_fuzz_dep_.CoverTab[117996]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:42
		// _ = "end of CoverTab[117996]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:42
	// _ = "end of CoverTab[117993]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:42
	_go_fuzz_dep_.CoverTab[117994]++
										return f, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:43
	// _ = "end of CoverTab[117994]"
}

func (OsFs) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:46
	_go_fuzz_dep_.CoverTab[117997]++
										return os.Mkdir(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:47
	// _ = "end of CoverTab[117997]"
}

func (OsFs) MkdirAll(path string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:50
	_go_fuzz_dep_.CoverTab[117998]++
										return os.MkdirAll(path, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:51
	// _ = "end of CoverTab[117998]"
}

func (OsFs) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:54
	_go_fuzz_dep_.CoverTab[117999]++
										f, e := os.Open(name)
										if f == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:56
		_go_fuzz_dep_.CoverTab[118001]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:59
		return nil, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:59
		// _ = "end of CoverTab[118001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:60
		_go_fuzz_dep_.CoverTab[118002]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:60
		// _ = "end of CoverTab[118002]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:60
	// _ = "end of CoverTab[117999]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:60
	_go_fuzz_dep_.CoverTab[118000]++
										return f, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:61
	// _ = "end of CoverTab[118000]"
}

func (OsFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:64
	_go_fuzz_dep_.CoverTab[118003]++
										f, e := os.OpenFile(name, flag, perm)
										if f == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:66
		_go_fuzz_dep_.CoverTab[118005]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:69
		return nil, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:69
		// _ = "end of CoverTab[118005]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:70
		_go_fuzz_dep_.CoverTab[118006]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:70
		// _ = "end of CoverTab[118006]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:70
	// _ = "end of CoverTab[118003]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:70
	_go_fuzz_dep_.CoverTab[118004]++
										return f, e
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:71
	// _ = "end of CoverTab[118004]"
}

func (OsFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:74
	_go_fuzz_dep_.CoverTab[118007]++
										return os.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:75
	// _ = "end of CoverTab[118007]"
}

func (OsFs) RemoveAll(path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:78
	_go_fuzz_dep_.CoverTab[118008]++
										return os.RemoveAll(path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:79
	// _ = "end of CoverTab[118008]"
}

func (OsFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:82
	_go_fuzz_dep_.CoverTab[118009]++
										return os.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:83
	// _ = "end of CoverTab[118009]"
}

func (OsFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:86
	_go_fuzz_dep_.CoverTab[118010]++
										return os.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:87
	// _ = "end of CoverTab[118010]"
}

func (OsFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:90
	_go_fuzz_dep_.CoverTab[118011]++
										return os.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:91
	// _ = "end of CoverTab[118011]"
}

func (OsFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:94
	_go_fuzz_dep_.CoverTab[118012]++
										return os.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:95
	// _ = "end of CoverTab[118012]"
}

func (OsFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:98
	_go_fuzz_dep_.CoverTab[118013]++
										return os.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:99
	// _ = "end of CoverTab[118013]"
}

func (OsFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:102
	_go_fuzz_dep_.CoverTab[118014]++
										fi, err := os.Lstat(name)
										return fi, true, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:104
	// _ = "end of CoverTab[118014]"
}

func (OsFs) SymlinkIfPossible(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:107
	_go_fuzz_dep_.CoverTab[118015]++
										return os.Symlink(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:108
	// _ = "end of CoverTab[118015]"
}

func (OsFs) ReadlinkIfPossible(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:111
	_go_fuzz_dep_.CoverTab[118016]++
										return os.Readlink(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:112
	// _ = "end of CoverTab[118016]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/os.go:113
var _ = _go_fuzz_dep_.CoverTab
