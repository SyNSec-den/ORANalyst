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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:14
)

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/spf13/afero/mem"
)

const chmodBits = os.ModePerm | os.ModeSetuid | os.ModeSetgid | os.ModeSticky	// Only a subset of bits are allowed to be changed. Documented under os.Chmod()

type MemMapFs struct {
	mu	sync.RWMutex
	data	map[string]*mem.FileData
	init	sync.Once
}

func NewMemMapFs() Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:36
	_go_fuzz_dep_.CoverTab[117854]++
										return &MemMapFs{}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:37
	// _ = "end of CoverTab[117854]"
}

func (m *MemMapFs) getData() map[string]*mem.FileData {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:40
	_go_fuzz_dep_.CoverTab[117855]++
										m.init.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:41
		_go_fuzz_dep_.CoverTab[117857]++
											m.data = make(map[string]*mem.FileData)

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:45
		root := mem.CreateDir(FilePathSeparator)
											mem.SetMode(root, os.ModeDir|0755)
											m.data[FilePathSeparator] = root
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:47
		// _ = "end of CoverTab[117857]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:48
	// _ = "end of CoverTab[117855]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:48
	_go_fuzz_dep_.CoverTab[117856]++
										return m.data
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:49
	// _ = "end of CoverTab[117856]"
}

func (*MemMapFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:52
	_go_fuzz_dep_.CoverTab[117858]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:52
	return "MemMapFS"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:52
	// _ = "end of CoverTab[117858]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:52
}

func (m *MemMapFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:54
	_go_fuzz_dep_.CoverTab[117859]++
										name = normalizePath(name)
										m.mu.Lock()
										file := mem.CreateFile(name)
										m.getData()[name] = file
										m.registerWithParent(file, 0)
										m.mu.Unlock()
										return mem.NewFileHandle(file), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:61
	// _ = "end of CoverTab[117859]"
}

func (m *MemMapFs) unRegisterWithParent(fileName string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:64
	_go_fuzz_dep_.CoverTab[117860]++
										f, err := m.lockfreeOpen(fileName)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:66
		_go_fuzz_dep_.CoverTab[117863]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:67
		// _ = "end of CoverTab[117863]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:68
		_go_fuzz_dep_.CoverTab[117864]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:68
		// _ = "end of CoverTab[117864]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:68
	// _ = "end of CoverTab[117860]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:68
	_go_fuzz_dep_.CoverTab[117861]++
										parent := m.findParent(f)
										if parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:70
		_go_fuzz_dep_.CoverTab[117865]++
											log.Panic("parent of ", f.Name(), " is nil")
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:71
		// _ = "end of CoverTab[117865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:72
		_go_fuzz_dep_.CoverTab[117866]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:72
		// _ = "end of CoverTab[117866]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:72
	// _ = "end of CoverTab[117861]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:72
	_go_fuzz_dep_.CoverTab[117862]++

										parent.Lock()
										mem.RemoveFromMemDir(parent, f)
										parent.Unlock()
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:77
	// _ = "end of CoverTab[117862]"
}

func (m *MemMapFs) findParent(f *mem.FileData) *mem.FileData {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:80
	_go_fuzz_dep_.CoverTab[117867]++
										pdir, _ := filepath.Split(f.Name())
										pdir = filepath.Clean(pdir)
										pfile, err := m.lockfreeOpen(pdir)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:84
		_go_fuzz_dep_.CoverTab[117869]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:85
		// _ = "end of CoverTab[117869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:86
		_go_fuzz_dep_.CoverTab[117870]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:86
		// _ = "end of CoverTab[117870]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:86
	// _ = "end of CoverTab[117867]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:86
	_go_fuzz_dep_.CoverTab[117868]++
										return pfile
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:87
	// _ = "end of CoverTab[117868]"
}

func (m *MemMapFs) registerWithParent(f *mem.FileData, perm os.FileMode) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:90
	_go_fuzz_dep_.CoverTab[117871]++
										if f == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:91
		_go_fuzz_dep_.CoverTab[117874]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:92
		// _ = "end of CoverTab[117874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:93
		_go_fuzz_dep_.CoverTab[117875]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:93
		// _ = "end of CoverTab[117875]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:93
	// _ = "end of CoverTab[117871]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:93
	_go_fuzz_dep_.CoverTab[117872]++
										parent := m.findParent(f)
										if parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:95
		_go_fuzz_dep_.CoverTab[117876]++
											pdir := filepath.Dir(filepath.Clean(f.Name()))
											err := m.lockfreeMkdir(pdir, perm)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:98
			_go_fuzz_dep_.CoverTab[117878]++

												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:100
			// _ = "end of CoverTab[117878]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:101
			_go_fuzz_dep_.CoverTab[117879]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:101
			// _ = "end of CoverTab[117879]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:101
		// _ = "end of CoverTab[117876]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:101
		_go_fuzz_dep_.CoverTab[117877]++
											parent, err = m.lockfreeOpen(pdir)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:103
			_go_fuzz_dep_.CoverTab[117880]++

												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:105
			// _ = "end of CoverTab[117880]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:106
			_go_fuzz_dep_.CoverTab[117881]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:106
			// _ = "end of CoverTab[117881]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:106
		// _ = "end of CoverTab[117877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:107
		_go_fuzz_dep_.CoverTab[117882]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:107
		// _ = "end of CoverTab[117882]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:107
	// _ = "end of CoverTab[117872]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:107
	_go_fuzz_dep_.CoverTab[117873]++

										parent.Lock()
										mem.InitializeDir(parent)
										mem.AddToMemDir(parent, f)
										parent.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:112
	// _ = "end of CoverTab[117873]"
}

func (m *MemMapFs) lockfreeMkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:115
	_go_fuzz_dep_.CoverTab[117883]++
										name = normalizePath(name)
										x, ok := m.getData()[name]
										if ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:118
		_go_fuzz_dep_.CoverTab[117885]++

											i := mem.FileInfo{FileData: x}
											if !i.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:121
			_go_fuzz_dep_.CoverTab[117886]++
												return ErrFileExists
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:122
			// _ = "end of CoverTab[117886]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:123
			_go_fuzz_dep_.CoverTab[117887]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:123
			// _ = "end of CoverTab[117887]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:123
		// _ = "end of CoverTab[117885]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:124
		_go_fuzz_dep_.CoverTab[117888]++
											item := mem.CreateDir(name)
											mem.SetMode(item, os.ModeDir|perm)
											m.getData()[name] = item
											m.registerWithParent(item, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:128
		// _ = "end of CoverTab[117888]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:129
	// _ = "end of CoverTab[117883]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:129
	_go_fuzz_dep_.CoverTab[117884]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:130
	// _ = "end of CoverTab[117884]"
}

func (m *MemMapFs) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:133
	_go_fuzz_dep_.CoverTab[117889]++
										perm &= chmodBits
										name = normalizePath(name)

										m.mu.RLock()
										_, ok := m.getData()[name]
										m.mu.RUnlock()
										if ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:140
		_go_fuzz_dep_.CoverTab[117891]++
											return &os.PathError{Op: "mkdir", Path: name, Err: ErrFileExists}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:141
		// _ = "end of CoverTab[117891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:142
		_go_fuzz_dep_.CoverTab[117892]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:142
		// _ = "end of CoverTab[117892]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:142
	// _ = "end of CoverTab[117889]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:142
	_go_fuzz_dep_.CoverTab[117890]++

										m.mu.Lock()
										item := mem.CreateDir(name)
										mem.SetMode(item, os.ModeDir|perm)
										m.getData()[name] = item
										m.registerWithParent(item, perm)
										m.mu.Unlock()

										return m.setFileMode(name, perm|os.ModeDir)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:151
	// _ = "end of CoverTab[117890]"
}

func (m *MemMapFs) MkdirAll(path string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:154
	_go_fuzz_dep_.CoverTab[117893]++
										err := m.Mkdir(path, perm)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:156
		_go_fuzz_dep_.CoverTab[117895]++
											if err.(*os.PathError).Err == ErrFileExists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:157
			_go_fuzz_dep_.CoverTab[117897]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:158
			// _ = "end of CoverTab[117897]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:159
			_go_fuzz_dep_.CoverTab[117898]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:159
			// _ = "end of CoverTab[117898]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:159
		// _ = "end of CoverTab[117895]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:159
		_go_fuzz_dep_.CoverTab[117896]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:160
		// _ = "end of CoverTab[117896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:161
		_go_fuzz_dep_.CoverTab[117899]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:161
		// _ = "end of CoverTab[117899]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:161
	// _ = "end of CoverTab[117893]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:161
	_go_fuzz_dep_.CoverTab[117894]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:162
	// _ = "end of CoverTab[117894]"
}

// Handle some relative paths
func normalizePath(path string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:166
	_go_fuzz_dep_.CoverTab[117900]++
										path = filepath.Clean(path)

										switch path {
	case ".":
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:170
		_go_fuzz_dep_.CoverTab[117901]++
											return FilePathSeparator
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:171
		// _ = "end of CoverTab[117901]"
	case "..":
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:172
		_go_fuzz_dep_.CoverTab[117902]++
											return FilePathSeparator
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:173
		// _ = "end of CoverTab[117902]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:174
		_go_fuzz_dep_.CoverTab[117903]++
											return path
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:175
		// _ = "end of CoverTab[117903]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:176
	// _ = "end of CoverTab[117900]"
}

func (m *MemMapFs) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:179
	_go_fuzz_dep_.CoverTab[117904]++
										f, err := m.open(name)
										if f != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:181
		_go_fuzz_dep_.CoverTab[117906]++
											return mem.NewReadOnlyFileHandle(f), err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:182
		// _ = "end of CoverTab[117906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:183
		_go_fuzz_dep_.CoverTab[117907]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:183
		// _ = "end of CoverTab[117907]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:183
	// _ = "end of CoverTab[117904]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:183
	_go_fuzz_dep_.CoverTab[117905]++
										return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:184
	// _ = "end of CoverTab[117905]"
}

func (m *MemMapFs) openWrite(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:187
	_go_fuzz_dep_.CoverTab[117908]++
										f, err := m.open(name)
										if f != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:189
		_go_fuzz_dep_.CoverTab[117910]++
											return mem.NewFileHandle(f), err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:190
		// _ = "end of CoverTab[117910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:191
		_go_fuzz_dep_.CoverTab[117911]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:191
		// _ = "end of CoverTab[117911]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:191
	// _ = "end of CoverTab[117908]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:191
	_go_fuzz_dep_.CoverTab[117909]++
										return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:192
	// _ = "end of CoverTab[117909]"
}

func (m *MemMapFs) open(name string) (*mem.FileData, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:195
	_go_fuzz_dep_.CoverTab[117912]++
										name = normalizePath(name)

										m.mu.RLock()
										f, ok := m.getData()[name]
										m.mu.RUnlock()
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:201
		_go_fuzz_dep_.CoverTab[117914]++
											return nil, &os.PathError{Op: "open", Path: name, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:202
		// _ = "end of CoverTab[117914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:203
		_go_fuzz_dep_.CoverTab[117915]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:203
		// _ = "end of CoverTab[117915]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:203
	// _ = "end of CoverTab[117912]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:203
	_go_fuzz_dep_.CoverTab[117913]++
										return f, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:204
	// _ = "end of CoverTab[117913]"
}

func (m *MemMapFs) lockfreeOpen(name string) (*mem.FileData, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:207
	_go_fuzz_dep_.CoverTab[117916]++
										name = normalizePath(name)
										f, ok := m.getData()[name]
										if ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:210
		_go_fuzz_dep_.CoverTab[117917]++
											return f, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:211
		// _ = "end of CoverTab[117917]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:212
		_go_fuzz_dep_.CoverTab[117918]++
											return nil, ErrFileNotFound
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:213
		// _ = "end of CoverTab[117918]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:214
	// _ = "end of CoverTab[117916]"
}

func (m *MemMapFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:217
	_go_fuzz_dep_.CoverTab[117919]++
										perm &= chmodBits
										chmod := false
										file, err := m.openWrite(name)
										if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:221
		_go_fuzz_dep_.CoverTab[117927]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:221
		return (flag&os.O_EXCL > 0)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:221
		// _ = "end of CoverTab[117927]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:221
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:221
		_go_fuzz_dep_.CoverTab[117928]++
											return nil, &os.PathError{Op: "open", Path: name, Err: ErrFileExists}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:222
		// _ = "end of CoverTab[117928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:223
		_go_fuzz_dep_.CoverTab[117929]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:223
		// _ = "end of CoverTab[117929]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:223
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:223
	// _ = "end of CoverTab[117919]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:223
	_go_fuzz_dep_.CoverTab[117920]++
										if os.IsNotExist(err) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:224
		_go_fuzz_dep_.CoverTab[117930]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:224
		return (flag&os.O_CREATE > 0)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:224
		// _ = "end of CoverTab[117930]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:224
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:224
		_go_fuzz_dep_.CoverTab[117931]++
											file, err = m.Create(name)
											chmod = true
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:226
		// _ = "end of CoverTab[117931]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:227
		_go_fuzz_dep_.CoverTab[117932]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:227
		// _ = "end of CoverTab[117932]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:227
	// _ = "end of CoverTab[117920]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:227
	_go_fuzz_dep_.CoverTab[117921]++
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:228
		_go_fuzz_dep_.CoverTab[117933]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:229
		// _ = "end of CoverTab[117933]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:230
		_go_fuzz_dep_.CoverTab[117934]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:230
		// _ = "end of CoverTab[117934]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:230
	// _ = "end of CoverTab[117921]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:230
	_go_fuzz_dep_.CoverTab[117922]++
										if flag == os.O_RDONLY {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:231
		_go_fuzz_dep_.CoverTab[117935]++
											file = mem.NewReadOnlyFileHandle(file.(*mem.File).Data())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:232
		// _ = "end of CoverTab[117935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:233
		_go_fuzz_dep_.CoverTab[117936]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:233
		// _ = "end of CoverTab[117936]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:233
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:233
	// _ = "end of CoverTab[117922]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:233
	_go_fuzz_dep_.CoverTab[117923]++
										if flag&os.O_APPEND > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:234
		_go_fuzz_dep_.CoverTab[117937]++
											_, err = file.Seek(0, os.SEEK_END)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:236
			_go_fuzz_dep_.CoverTab[117938]++
												file.Close()
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:238
			// _ = "end of CoverTab[117938]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:239
			_go_fuzz_dep_.CoverTab[117939]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:239
			// _ = "end of CoverTab[117939]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:239
		// _ = "end of CoverTab[117937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:240
		_go_fuzz_dep_.CoverTab[117940]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:240
		// _ = "end of CoverTab[117940]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:240
	// _ = "end of CoverTab[117923]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:240
	_go_fuzz_dep_.CoverTab[117924]++
										if flag&os.O_TRUNC > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:241
		_go_fuzz_dep_.CoverTab[117941]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:241
		return flag&(os.O_RDWR|os.O_WRONLY) > 0
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:241
		// _ = "end of CoverTab[117941]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:241
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:241
		_go_fuzz_dep_.CoverTab[117942]++
											err = file.Truncate(0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:243
			_go_fuzz_dep_.CoverTab[117943]++
												file.Close()
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:245
			// _ = "end of CoverTab[117943]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:246
			_go_fuzz_dep_.CoverTab[117944]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:246
			// _ = "end of CoverTab[117944]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:246
		// _ = "end of CoverTab[117942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:247
		_go_fuzz_dep_.CoverTab[117945]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:247
		// _ = "end of CoverTab[117945]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:247
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:247
	// _ = "end of CoverTab[117924]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:247
	_go_fuzz_dep_.CoverTab[117925]++
										if chmod {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:248
		_go_fuzz_dep_.CoverTab[117946]++
											return file, m.setFileMode(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:249
		// _ = "end of CoverTab[117946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:250
		_go_fuzz_dep_.CoverTab[117947]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:250
		// _ = "end of CoverTab[117947]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:250
	// _ = "end of CoverTab[117925]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:250
	_go_fuzz_dep_.CoverTab[117926]++
										return file, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:251
	// _ = "end of CoverTab[117926]"
}

func (m *MemMapFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:254
	_go_fuzz_dep_.CoverTab[117948]++
										name = normalizePath(name)

										m.mu.Lock()
										defer m.mu.Unlock()

										if _, ok := m.getData()[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:260
		_go_fuzz_dep_.CoverTab[117950]++
											err := m.unRegisterWithParent(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:262
			_go_fuzz_dep_.CoverTab[117952]++
												return &os.PathError{Op: "remove", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:263
			// _ = "end of CoverTab[117952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:264
			_go_fuzz_dep_.CoverTab[117953]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:264
			// _ = "end of CoverTab[117953]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:264
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:264
		// _ = "end of CoverTab[117950]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:264
		_go_fuzz_dep_.CoverTab[117951]++
											delete(m.getData(), name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:265
		// _ = "end of CoverTab[117951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:266
		_go_fuzz_dep_.CoverTab[117954]++
											return &os.PathError{Op: "remove", Path: name, Err: os.ErrNotExist}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:267
		// _ = "end of CoverTab[117954]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:268
	// _ = "end of CoverTab[117948]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:268
	_go_fuzz_dep_.CoverTab[117949]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:269
	// _ = "end of CoverTab[117949]"
}

func (m *MemMapFs) RemoveAll(path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:272
	_go_fuzz_dep_.CoverTab[117955]++
										path = normalizePath(path)
										m.mu.Lock()
										m.unRegisterWithParent(path)
										m.mu.Unlock()

										m.mu.RLock()
										defer m.mu.RUnlock()

										for p := range m.getData() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:281
		_go_fuzz_dep_.CoverTab[117957]++
											if p == path || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:282
			_go_fuzz_dep_.CoverTab[117958]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:282
			return strings.HasPrefix(p, path+FilePathSeparator)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:282
			// _ = "end of CoverTab[117958]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:282
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:282
			_go_fuzz_dep_.CoverTab[117959]++
												m.mu.RUnlock()
												m.mu.Lock()
												delete(m.getData(), p)
												m.mu.Unlock()
												m.mu.RLock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:287
			// _ = "end of CoverTab[117959]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:288
			_go_fuzz_dep_.CoverTab[117960]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:288
			// _ = "end of CoverTab[117960]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:288
		// _ = "end of CoverTab[117957]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:289
	// _ = "end of CoverTab[117955]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:289
	_go_fuzz_dep_.CoverTab[117956]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:290
	// _ = "end of CoverTab[117956]"
}

func (m *MemMapFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:293
	_go_fuzz_dep_.CoverTab[117961]++
										oldname = normalizePath(oldname)
										newname = normalizePath(newname)

										if oldname == newname {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:297
		_go_fuzz_dep_.CoverTab[117964]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:298
		// _ = "end of CoverTab[117964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:299
		_go_fuzz_dep_.CoverTab[117965]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:299
		// _ = "end of CoverTab[117965]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:299
	// _ = "end of CoverTab[117961]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:299
	_go_fuzz_dep_.CoverTab[117962]++

										m.mu.RLock()
										defer m.mu.RUnlock()
										if _, ok := m.getData()[oldname]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:303
		_go_fuzz_dep_.CoverTab[117966]++
											m.mu.RUnlock()
											m.mu.Lock()
											m.unRegisterWithParent(oldname)
											fileData := m.getData()[oldname]
											delete(m.getData(), oldname)
											mem.ChangeFileName(fileData, newname)
											m.getData()[newname] = fileData
											m.registerWithParent(fileData, 0)
											m.mu.Unlock()
											m.mu.RLock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:313
		// _ = "end of CoverTab[117966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:314
		_go_fuzz_dep_.CoverTab[117967]++
											return &os.PathError{Op: "rename", Path: oldname, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:315
		// _ = "end of CoverTab[117967]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:316
	// _ = "end of CoverTab[117962]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:316
	_go_fuzz_dep_.CoverTab[117963]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:317
	// _ = "end of CoverTab[117963]"
}

func (m *MemMapFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:320
	_go_fuzz_dep_.CoverTab[117968]++
										fileInfo, err := m.Stat(name)
										return fileInfo, false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:322
	// _ = "end of CoverTab[117968]"
}

func (m *MemMapFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:325
	_go_fuzz_dep_.CoverTab[117969]++
										f, err := m.Open(name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:327
		_go_fuzz_dep_.CoverTab[117971]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:328
		// _ = "end of CoverTab[117971]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:329
		_go_fuzz_dep_.CoverTab[117972]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:329
		// _ = "end of CoverTab[117972]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:329
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:329
	// _ = "end of CoverTab[117969]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:329
	_go_fuzz_dep_.CoverTab[117970]++
										fi := mem.GetFileInfo(f.(*mem.File).Data())
										return fi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:331
	// _ = "end of CoverTab[117970]"
}

func (m *MemMapFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:334
	_go_fuzz_dep_.CoverTab[117973]++
										mode &= chmodBits

										m.mu.RLock()
										f, ok := m.getData()[name]
										m.mu.RUnlock()
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:340
		_go_fuzz_dep_.CoverTab[117975]++
											return &os.PathError{Op: "chmod", Path: name, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:341
		// _ = "end of CoverTab[117975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:342
		_go_fuzz_dep_.CoverTab[117976]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:342
		// _ = "end of CoverTab[117976]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:342
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:342
	// _ = "end of CoverTab[117973]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:342
	_go_fuzz_dep_.CoverTab[117974]++
										prevOtherBits := mem.GetFileInfo(f).Mode() & ^chmodBits

										mode = prevOtherBits | mode
										return m.setFileMode(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:346
	// _ = "end of CoverTab[117974]"
}

func (m *MemMapFs) setFileMode(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:349
	_go_fuzz_dep_.CoverTab[117977]++
										name = normalizePath(name)

										m.mu.RLock()
										f, ok := m.getData()[name]
										m.mu.RUnlock()
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:355
		_go_fuzz_dep_.CoverTab[117979]++
											return &os.PathError{Op: "chmod", Path: name, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:356
		// _ = "end of CoverTab[117979]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:357
		_go_fuzz_dep_.CoverTab[117980]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:357
		// _ = "end of CoverTab[117980]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:357
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:357
	// _ = "end of CoverTab[117977]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:357
	_go_fuzz_dep_.CoverTab[117978]++

										m.mu.Lock()
										mem.SetMode(f, mode)
										m.mu.Unlock()

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:363
	// _ = "end of CoverTab[117978]"
}

func (m *MemMapFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:366
	_go_fuzz_dep_.CoverTab[117981]++
										name = normalizePath(name)

										m.mu.RLock()
										f, ok := m.getData()[name]
										m.mu.RUnlock()
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:372
		_go_fuzz_dep_.CoverTab[117983]++
											return &os.PathError{Op: "chown", Path: name, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:373
		// _ = "end of CoverTab[117983]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:374
		_go_fuzz_dep_.CoverTab[117984]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:374
		// _ = "end of CoverTab[117984]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:374
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:374
	// _ = "end of CoverTab[117981]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:374
	_go_fuzz_dep_.CoverTab[117982]++

										mem.SetUID(f, uid)
										mem.SetGID(f, gid)

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:379
	// _ = "end of CoverTab[117982]"
}

func (m *MemMapFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:382
	_go_fuzz_dep_.CoverTab[117985]++
										name = normalizePath(name)

										m.mu.RLock()
										f, ok := m.getData()[name]
										m.mu.RUnlock()
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:388
		_go_fuzz_dep_.CoverTab[117987]++
											return &os.PathError{Op: "chtimes", Path: name, Err: ErrFileNotFound}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:389
		// _ = "end of CoverTab[117987]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:390
		_go_fuzz_dep_.CoverTab[117988]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:390
		// _ = "end of CoverTab[117988]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:390
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:390
	// _ = "end of CoverTab[117985]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:390
	_go_fuzz_dep_.CoverTab[117986]++

										m.mu.Lock()
										mem.SetModTime(f, mtime)
										m.mu.Unlock()

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:396
	// _ = "end of CoverTab[117986]"
}

func (m *MemMapFs) List() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:399
	_go_fuzz_dep_.CoverTab[117989]++
										for _, x := range m.data {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:400
		_go_fuzz_dep_.CoverTab[117990]++
											y := mem.FileInfo{FileData: x}
											fmt.Println(x.Name(), y.Size())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:402
		// _ = "end of CoverTab[117990]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:403
	// _ = "end of CoverTab[117989]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:404
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/memmap.go:404
var _ = _go_fuzz_dep_.CoverTab
