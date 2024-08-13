// Copyright © 2015 Steve Francia <spf@spf13.com>.
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
package mem

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:15
)

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/afero/internal/common"
)

const FilePathSeparator = string(filepath.Separator)

var _ fs.ReadDirFile = &File{}

type File struct {
	// atomic requires 64-bit alignment for struct field access
	at		int64
	readDirCount	int64
	closed		bool
	readOnly	bool
	fileData	*FileData
}

func NewFileHandle(data *FileData) *File {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:44
	_go_fuzz_dep_.CoverTab[116812]++
										return &File{fileData: data}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:45
	// _ = "end of CoverTab[116812]"
}

func NewReadOnlyFileHandle(data *FileData) *File {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:48
	_go_fuzz_dep_.CoverTab[116813]++
										return &File{fileData: data, readOnly: true}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:49
	// _ = "end of CoverTab[116813]"
}

func (f File) Data() *FileData {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:52
	_go_fuzz_dep_.CoverTab[116814]++
										return f.fileData
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:53
	// _ = "end of CoverTab[116814]"
}

type FileData struct {
	sync.Mutex
	name	string
	data	[]byte
	memDir	Dir
	dir	bool
	mode	os.FileMode
	modtime	time.Time
	uid	int
	gid	int
}

func (d *FileData) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:68
	_go_fuzz_dep_.CoverTab[116815]++
										d.Lock()
										defer d.Unlock()
										return d.name
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:71
	// _ = "end of CoverTab[116815]"
}

func CreateFile(name string) *FileData {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:74
	_go_fuzz_dep_.CoverTab[116816]++
										return &FileData{name: name, mode: os.ModeTemporary, modtime: time.Now()}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:75
	// _ = "end of CoverTab[116816]"
}

func CreateDir(name string) *FileData {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:78
	_go_fuzz_dep_.CoverTab[116817]++
										return &FileData{name: name, memDir: &DirMap{}, dir: true, modtime: time.Now()}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:79
	// _ = "end of CoverTab[116817]"
}

func ChangeFileName(f *FileData, newname string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:82
	_go_fuzz_dep_.CoverTab[116818]++
										f.Lock()
										f.name = newname
										f.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:85
	// _ = "end of CoverTab[116818]"
}

func SetMode(f *FileData, mode os.FileMode) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:88
	_go_fuzz_dep_.CoverTab[116819]++
										f.Lock()
										f.mode = mode
										f.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:91
	// _ = "end of CoverTab[116819]"
}

func SetModTime(f *FileData, mtime time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:94
	_go_fuzz_dep_.CoverTab[116820]++
										f.Lock()
										setModTime(f, mtime)
										f.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:97
	// _ = "end of CoverTab[116820]"
}

func setModTime(f *FileData, mtime time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:100
	_go_fuzz_dep_.CoverTab[116821]++
											f.modtime = mtime
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:101
	// _ = "end of CoverTab[116821]"
}

func SetUID(f *FileData, uid int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:104
	_go_fuzz_dep_.CoverTab[116822]++
											f.Lock()
											f.uid = uid
											f.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:107
	// _ = "end of CoverTab[116822]"
}

func SetGID(f *FileData, gid int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:110
	_go_fuzz_dep_.CoverTab[116823]++
											f.Lock()
											f.gid = gid
											f.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:113
	// _ = "end of CoverTab[116823]"
}

func GetFileInfo(f *FileData) *FileInfo {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:116
	_go_fuzz_dep_.CoverTab[116824]++
											return &FileInfo{f}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:117
	// _ = "end of CoverTab[116824]"
}

func (f *File) Open() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:120
	_go_fuzz_dep_.CoverTab[116825]++
											atomic.StoreInt64(&f.at, 0)
											atomic.StoreInt64(&f.readDirCount, 0)
											f.fileData.Lock()
											f.closed = false
											f.fileData.Unlock()
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:126
	// _ = "end of CoverTab[116825]"
}

func (f *File) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:129
	_go_fuzz_dep_.CoverTab[116826]++
											f.fileData.Lock()
											f.closed = true
											if !f.readOnly {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:132
		_go_fuzz_dep_.CoverTab[116828]++
												setModTime(f.fileData, time.Now())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:133
		// _ = "end of CoverTab[116828]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:134
		_go_fuzz_dep_.CoverTab[116829]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:134
		// _ = "end of CoverTab[116829]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:134
	// _ = "end of CoverTab[116826]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:134
	_go_fuzz_dep_.CoverTab[116827]++
											f.fileData.Unlock()
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:136
	// _ = "end of CoverTab[116827]"
}

func (f *File) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:139
	_go_fuzz_dep_.CoverTab[116830]++
											return f.fileData.Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:140
	// _ = "end of CoverTab[116830]"
}

func (f *File) Stat() (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:143
	_go_fuzz_dep_.CoverTab[116831]++
											return &FileInfo{f.fileData}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:144
	// _ = "end of CoverTab[116831]"
}

func (f *File) Sync() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:147
	_go_fuzz_dep_.CoverTab[116832]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:148
	// _ = "end of CoverTab[116832]"
}

func (f *File) Readdir(count int) (res []os.FileInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:151
	_go_fuzz_dep_.CoverTab[116833]++
											if !f.fileData.dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:152
		_go_fuzz_dep_.CoverTab[116837]++
												return nil, &os.PathError{Op: "readdir", Path: f.fileData.name, Err: errors.New("not a dir")}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:153
		// _ = "end of CoverTab[116837]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:154
		_go_fuzz_dep_.CoverTab[116838]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:154
		// _ = "end of CoverTab[116838]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:154
	// _ = "end of CoverTab[116833]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:154
	_go_fuzz_dep_.CoverTab[116834]++
											var outLength int64

											f.fileData.Lock()
											files := f.fileData.memDir.Files()[f.readDirCount:]
											if count > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:159
		_go_fuzz_dep_.CoverTab[116839]++
												if len(files) < count {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:160
			_go_fuzz_dep_.CoverTab[116841]++
													outLength = int64(len(files))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:161
			// _ = "end of CoverTab[116841]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:162
			_go_fuzz_dep_.CoverTab[116842]++
													outLength = int64(count)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:163
			// _ = "end of CoverTab[116842]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:164
		// _ = "end of CoverTab[116839]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:164
		_go_fuzz_dep_.CoverTab[116840]++
												if len(files) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:165
			_go_fuzz_dep_.CoverTab[116843]++
													err = io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:166
			// _ = "end of CoverTab[116843]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:167
			_go_fuzz_dep_.CoverTab[116844]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:167
			// _ = "end of CoverTab[116844]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:167
		// _ = "end of CoverTab[116840]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:168
		_go_fuzz_dep_.CoverTab[116845]++
												outLength = int64(len(files))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:169
		// _ = "end of CoverTab[116845]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:170
	// _ = "end of CoverTab[116834]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:170
	_go_fuzz_dep_.CoverTab[116835]++
											f.readDirCount += outLength
											f.fileData.Unlock()

											res = make([]os.FileInfo, outLength)
											for i := range res {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:175
		_go_fuzz_dep_.CoverTab[116846]++
												res[i] = &FileInfo{files[i]}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:176
		// _ = "end of CoverTab[116846]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:177
	// _ = "end of CoverTab[116835]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:177
	_go_fuzz_dep_.CoverTab[116836]++

											return res, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:179
	// _ = "end of CoverTab[116836]"
}

func (f *File) Readdirnames(n int) (names []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:182
	_go_fuzz_dep_.CoverTab[116847]++
											fi, err := f.Readdir(n)
											names = make([]string, len(fi))
											for i, f := range fi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:185
		_go_fuzz_dep_.CoverTab[116849]++
												_, names[i] = filepath.Split(f.Name())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:186
		// _ = "end of CoverTab[116849]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:187
	// _ = "end of CoverTab[116847]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:187
	_go_fuzz_dep_.CoverTab[116848]++
											return names, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:188
	// _ = "end of CoverTab[116848]"
}

// Implements fs.ReadDirFile
func (f *File) ReadDir(n int) ([]fs.DirEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:192
	_go_fuzz_dep_.CoverTab[116850]++
											fi, err := f.Readdir(n)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:194
		_go_fuzz_dep_.CoverTab[116853]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:195
		// _ = "end of CoverTab[116853]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:196
		_go_fuzz_dep_.CoverTab[116854]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:196
		// _ = "end of CoverTab[116854]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:196
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:196
	// _ = "end of CoverTab[116850]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:196
	_go_fuzz_dep_.CoverTab[116851]++
											di := make([]fs.DirEntry, len(fi))
											for i, f := range fi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:198
		_go_fuzz_dep_.CoverTab[116855]++
												di[i] = common.FileInfoDirEntry{FileInfo: f}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:199
		// _ = "end of CoverTab[116855]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:200
	// _ = "end of CoverTab[116851]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:200
	_go_fuzz_dep_.CoverTab[116852]++
											return di, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:201
	// _ = "end of CoverTab[116852]"
}

func (f *File) Read(b []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:204
	_go_fuzz_dep_.CoverTab[116856]++
											f.fileData.Lock()
											defer f.fileData.Unlock()
											if f.closed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:207
		_go_fuzz_dep_.CoverTab[116861]++
												return 0, ErrFileClosed
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:208
		// _ = "end of CoverTab[116861]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:209
		_go_fuzz_dep_.CoverTab[116862]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:209
		// _ = "end of CoverTab[116862]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:209
	// _ = "end of CoverTab[116856]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:209
	_go_fuzz_dep_.CoverTab[116857]++
											if len(b) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:210
		_go_fuzz_dep_.CoverTab[116863]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:210
		return int(f.at) == len(f.fileData.data)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:210
		// _ = "end of CoverTab[116863]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:210
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:210
		_go_fuzz_dep_.CoverTab[116864]++
												return 0, io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:211
		// _ = "end of CoverTab[116864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:212
		_go_fuzz_dep_.CoverTab[116865]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:212
		// _ = "end of CoverTab[116865]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:212
	// _ = "end of CoverTab[116857]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:212
	_go_fuzz_dep_.CoverTab[116858]++
											if int(f.at) > len(f.fileData.data) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:213
		_go_fuzz_dep_.CoverTab[116866]++
												return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:214
		// _ = "end of CoverTab[116866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:215
		_go_fuzz_dep_.CoverTab[116867]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:215
		// _ = "end of CoverTab[116867]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:215
	// _ = "end of CoverTab[116858]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:215
	_go_fuzz_dep_.CoverTab[116859]++
											if len(f.fileData.data)-int(f.at) >= len(b) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:216
		_go_fuzz_dep_.CoverTab[116868]++
												n = len(b)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:217
		// _ = "end of CoverTab[116868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:218
		_go_fuzz_dep_.CoverTab[116869]++
												n = len(f.fileData.data) - int(f.at)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:219
		// _ = "end of CoverTab[116869]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:220
	// _ = "end of CoverTab[116859]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:220
	_go_fuzz_dep_.CoverTab[116860]++
											copy(b, f.fileData.data[f.at:f.at+int64(n)])
											atomic.AddInt64(&f.at, int64(n))
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:223
	// _ = "end of CoverTab[116860]"
}

func (f *File) ReadAt(b []byte, off int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:226
	_go_fuzz_dep_.CoverTab[116870]++
											prev := atomic.LoadInt64(&f.at)
											atomic.StoreInt64(&f.at, off)
											n, err = f.Read(b)
											atomic.StoreInt64(&f.at, prev)
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:231
	// _ = "end of CoverTab[116870]"
}

func (f *File) Truncate(size int64) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:234
	_go_fuzz_dep_.CoverTab[116871]++
											if f.closed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:235
		_go_fuzz_dep_.CoverTab[116876]++
												return ErrFileClosed
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:236
		// _ = "end of CoverTab[116876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:237
		_go_fuzz_dep_.CoverTab[116877]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:237
		// _ = "end of CoverTab[116877]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:237
	// _ = "end of CoverTab[116871]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:237
	_go_fuzz_dep_.CoverTab[116872]++
											if f.readOnly {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:238
		_go_fuzz_dep_.CoverTab[116878]++
												return &os.PathError{Op: "truncate", Path: f.fileData.name, Err: errors.New("file handle is read only")}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:239
		// _ = "end of CoverTab[116878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:240
		_go_fuzz_dep_.CoverTab[116879]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:240
		// _ = "end of CoverTab[116879]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:240
	// _ = "end of CoverTab[116872]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:240
	_go_fuzz_dep_.CoverTab[116873]++
											if size < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:241
		_go_fuzz_dep_.CoverTab[116880]++
												return ErrOutOfRange
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:242
		// _ = "end of CoverTab[116880]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:243
		_go_fuzz_dep_.CoverTab[116881]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:243
		// _ = "end of CoverTab[116881]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:243
	// _ = "end of CoverTab[116873]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:243
	_go_fuzz_dep_.CoverTab[116874]++
											f.fileData.Lock()
											defer f.fileData.Unlock()
											if size > int64(len(f.fileData.data)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:246
		_go_fuzz_dep_.CoverTab[116882]++
												diff := size - int64(len(f.fileData.data))
												f.fileData.data = append(f.fileData.data, bytes.Repeat([]byte{00}, int(diff))...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:248
		// _ = "end of CoverTab[116882]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:249
		_go_fuzz_dep_.CoverTab[116883]++
												f.fileData.data = f.fileData.data[0:size]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:250
		// _ = "end of CoverTab[116883]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:251
	// _ = "end of CoverTab[116874]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:251
	_go_fuzz_dep_.CoverTab[116875]++
											setModTime(f.fileData, time.Now())
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:253
	// _ = "end of CoverTab[116875]"
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:256
	_go_fuzz_dep_.CoverTab[116884]++
											if f.closed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:257
		_go_fuzz_dep_.CoverTab[116887]++
												return 0, ErrFileClosed
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:258
		// _ = "end of CoverTab[116887]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:259
		_go_fuzz_dep_.CoverTab[116888]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:259
		// _ = "end of CoverTab[116888]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:259
	// _ = "end of CoverTab[116884]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:259
	_go_fuzz_dep_.CoverTab[116885]++
											switch whence {
	case io.SeekStart:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:261
		_go_fuzz_dep_.CoverTab[116889]++
												atomic.StoreInt64(&f.at, offset)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:262
		// _ = "end of CoverTab[116889]"
	case io.SeekCurrent:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:263
		_go_fuzz_dep_.CoverTab[116890]++
												atomic.AddInt64(&f.at, offset)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:264
		// _ = "end of CoverTab[116890]"
	case io.SeekEnd:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:265
		_go_fuzz_dep_.CoverTab[116891]++
												atomic.StoreInt64(&f.at, int64(len(f.fileData.data))+offset)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:266
		// _ = "end of CoverTab[116891]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:266
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:266
		_go_fuzz_dep_.CoverTab[116892]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:266
		// _ = "end of CoverTab[116892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:267
	// _ = "end of CoverTab[116885]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:267
	_go_fuzz_dep_.CoverTab[116886]++
											return f.at, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:268
	// _ = "end of CoverTab[116886]"
}

func (f *File) Write(b []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:271
	_go_fuzz_dep_.CoverTab[116893]++
											if f.closed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:272
		_go_fuzz_dep_.CoverTab[116898]++
												return 0, ErrFileClosed
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:273
		// _ = "end of CoverTab[116898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:274
		_go_fuzz_dep_.CoverTab[116899]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:274
		// _ = "end of CoverTab[116899]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:274
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:274
	// _ = "end of CoverTab[116893]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:274
	_go_fuzz_dep_.CoverTab[116894]++
											if f.readOnly {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:275
		_go_fuzz_dep_.CoverTab[116900]++
												return 0, &os.PathError{Op: "write", Path: f.fileData.name, Err: errors.New("file handle is read only")}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:276
		// _ = "end of CoverTab[116900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:277
		_go_fuzz_dep_.CoverTab[116901]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:277
		// _ = "end of CoverTab[116901]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:277
	// _ = "end of CoverTab[116894]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:277
	_go_fuzz_dep_.CoverTab[116895]++
											n = len(b)
											cur := atomic.LoadInt64(&f.at)
											f.fileData.Lock()
											defer f.fileData.Unlock()
											diff := cur - int64(len(f.fileData.data))
											var tail []byte
											if n+int(cur) < len(f.fileData.data) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:284
		_go_fuzz_dep_.CoverTab[116902]++
												tail = f.fileData.data[n+int(cur):]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:285
		// _ = "end of CoverTab[116902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:286
		_go_fuzz_dep_.CoverTab[116903]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:286
		// _ = "end of CoverTab[116903]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:286
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:286
	// _ = "end of CoverTab[116895]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:286
	_go_fuzz_dep_.CoverTab[116896]++
											if diff > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:287
		_go_fuzz_dep_.CoverTab[116904]++
												f.fileData.data = append(f.fileData.data, append(bytes.Repeat([]byte{00}, int(diff)), b...)...)
												f.fileData.data = append(f.fileData.data, tail...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:289
		// _ = "end of CoverTab[116904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:290
		_go_fuzz_dep_.CoverTab[116905]++
												f.fileData.data = append(f.fileData.data[:cur], b...)
												f.fileData.data = append(f.fileData.data, tail...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:292
		// _ = "end of CoverTab[116905]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:293
	// _ = "end of CoverTab[116896]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:293
	_go_fuzz_dep_.CoverTab[116897]++
											setModTime(f.fileData, time.Now())

											atomic.AddInt64(&f.at, int64(n))
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:297
	// _ = "end of CoverTab[116897]"
}

func (f *File) WriteAt(b []byte, off int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:300
	_go_fuzz_dep_.CoverTab[116906]++
											atomic.StoreInt64(&f.at, off)
											return f.Write(b)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:302
	// _ = "end of CoverTab[116906]"
}

func (f *File) WriteString(s string) (ret int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:305
	_go_fuzz_dep_.CoverTab[116907]++
											return f.Write([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:306
	// _ = "end of CoverTab[116907]"
}

func (f *File) Info() *FileInfo {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:309
	_go_fuzz_dep_.CoverTab[116908]++
											return &FileInfo{f.fileData}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:310
	// _ = "end of CoverTab[116908]"
}

type FileInfo struct {
	*FileData
}

// Implements os.FileInfo
func (s *FileInfo) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:318
	_go_fuzz_dep_.CoverTab[116909]++
											s.Lock()
											_, name := filepath.Split(s.name)
											s.Unlock()
											return name
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:322
	// _ = "end of CoverTab[116909]"
}
func (s *FileInfo) Mode() os.FileMode {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:324
	_go_fuzz_dep_.CoverTab[116910]++
											s.Lock()
											defer s.Unlock()
											return s.mode
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:327
	// _ = "end of CoverTab[116910]"
}
func (s *FileInfo) ModTime() time.Time {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:329
	_go_fuzz_dep_.CoverTab[116911]++
											s.Lock()
											defer s.Unlock()
											return s.modtime
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:332
	// _ = "end of CoverTab[116911]"
}
func (s *FileInfo) IsDir() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:334
	_go_fuzz_dep_.CoverTab[116912]++
											s.Lock()
											defer s.Unlock()
											return s.dir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:337
	// _ = "end of CoverTab[116912]"
}
func (s *FileInfo) Sys() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:339
	_go_fuzz_dep_.CoverTab[116913]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:339
	return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:339
	// _ = "end of CoverTab[116913]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:339
}
func (s *FileInfo) Size() int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:340
	_go_fuzz_dep_.CoverTab[116914]++
											if s.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:341
		_go_fuzz_dep_.CoverTab[116916]++
												return int64(42)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:342
		// _ = "end of CoverTab[116916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:343
		_go_fuzz_dep_.CoverTab[116917]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:343
		// _ = "end of CoverTab[116917]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:343
	// _ = "end of CoverTab[116914]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:343
	_go_fuzz_dep_.CoverTab[116915]++
											s.Lock()
											defer s.Unlock()
											return int64(len(s.data))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:346
	// _ = "end of CoverTab[116915]"
}

var (
	ErrFileClosed		= errors.New("File is closed")
	ErrOutOfRange		= errors.New("out of range")
	ErrTooLarge		= errors.New("too large")
	ErrFileNotFound		= os.ErrNotExist
	ErrFileExists		= os.ErrExist
	ErrDestinationExists	= os.ErrExist
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:356
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/mem/file.go:356
var _ = _go_fuzz_dep_.CoverTab
