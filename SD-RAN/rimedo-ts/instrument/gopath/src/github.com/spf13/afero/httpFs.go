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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:14
)

import (
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type httpDir struct {
	basePath	string
	fs		HttpFs
}

func (d httpDir) Open(name string) (http.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:31
	_go_fuzz_dep_.CoverTab[117599]++
										if filepath.Separator != '/' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
		_go_fuzz_dep_.CoverTab[117603]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
		return strings.ContainsRune(name, filepath.Separator)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
		// _ = "end of CoverTab[117603]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
		_go_fuzz_dep_.CoverTab[117604]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:32
		return strings.Contains(name, "\x00")
											// _ = "end of CoverTab[117604]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:33
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:33
		_go_fuzz_dep_.CoverTab[117605]++
											return nil, errors.New("http: invalid character in file path")
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:34
		// _ = "end of CoverTab[117605]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:35
		_go_fuzz_dep_.CoverTab[117606]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:35
		// _ = "end of CoverTab[117606]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:35
	// _ = "end of CoverTab[117599]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:35
	_go_fuzz_dep_.CoverTab[117600]++
										dir := string(d.basePath)
										if dir == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:37
		_go_fuzz_dep_.CoverTab[117607]++
											dir = "."
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:38
		// _ = "end of CoverTab[117607]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:39
		_go_fuzz_dep_.CoverTab[117608]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:39
		// _ = "end of CoverTab[117608]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:39
	// _ = "end of CoverTab[117600]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:39
	_go_fuzz_dep_.CoverTab[117601]++

										f, err := d.fs.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:42
		_go_fuzz_dep_.CoverTab[117609]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:43
		// _ = "end of CoverTab[117609]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:44
		_go_fuzz_dep_.CoverTab[117610]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:44
		// _ = "end of CoverTab[117610]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:44
	// _ = "end of CoverTab[117601]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:44
	_go_fuzz_dep_.CoverTab[117602]++
										return f, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:45
	// _ = "end of CoverTab[117602]"
}

type HttpFs struct {
	source Fs
}

func NewHttpFs(source Fs) *HttpFs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:52
	_go_fuzz_dep_.CoverTab[117611]++
										return &HttpFs{source: source}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:53
	// _ = "end of CoverTab[117611]"
}

func (h HttpFs) Dir(s string) *httpDir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:56
	_go_fuzz_dep_.CoverTab[117612]++
										return &httpDir{basePath: s, fs: h}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:57
	// _ = "end of CoverTab[117612]"
}

func (h HttpFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:60
	_go_fuzz_dep_.CoverTab[117613]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:60
	return "h HttpFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:60
	// _ = "end of CoverTab[117613]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:60
}

func (h HttpFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:62
	_go_fuzz_dep_.CoverTab[117614]++
										return h.source.Create(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:63
	// _ = "end of CoverTab[117614]"
}

func (h HttpFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:66
	_go_fuzz_dep_.CoverTab[117615]++
										return h.source.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:67
	// _ = "end of CoverTab[117615]"
}

func (h HttpFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:70
	_go_fuzz_dep_.CoverTab[117616]++
										return h.source.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:71
	// _ = "end of CoverTab[117616]"
}

func (h HttpFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:74
	_go_fuzz_dep_.CoverTab[117617]++
										return h.source.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:75
	// _ = "end of CoverTab[117617]"
}

func (h HttpFs) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:78
	_go_fuzz_dep_.CoverTab[117618]++
										return h.source.Mkdir(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:79
	// _ = "end of CoverTab[117618]"
}

func (h HttpFs) MkdirAll(path string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:82
	_go_fuzz_dep_.CoverTab[117619]++
										return h.source.MkdirAll(path, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:83
	// _ = "end of CoverTab[117619]"
}

func (h HttpFs) Open(name string) (http.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:86
	_go_fuzz_dep_.CoverTab[117620]++
										f, err := h.source.Open(name)
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:88
		_go_fuzz_dep_.CoverTab[117622]++
											if httpfile, ok := f.(http.File); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:89
			_go_fuzz_dep_.CoverTab[117623]++
												return httpfile, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:90
			// _ = "end of CoverTab[117623]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:91
			_go_fuzz_dep_.CoverTab[117624]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:91
			// _ = "end of CoverTab[117624]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:91
		// _ = "end of CoverTab[117622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:92
		_go_fuzz_dep_.CoverTab[117625]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:92
		// _ = "end of CoverTab[117625]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:92
	// _ = "end of CoverTab[117620]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:92
	_go_fuzz_dep_.CoverTab[117621]++
										return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:93
	// _ = "end of CoverTab[117621]"
}

func (h HttpFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:96
	_go_fuzz_dep_.CoverTab[117626]++
										return h.source.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:97
	// _ = "end of CoverTab[117626]"
}

func (h HttpFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:100
	_go_fuzz_dep_.CoverTab[117627]++
										return h.source.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:101
	// _ = "end of CoverTab[117627]"
}

func (h HttpFs) RemoveAll(path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:104
	_go_fuzz_dep_.CoverTab[117628]++
										return h.source.RemoveAll(path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:105
	// _ = "end of CoverTab[117628]"
}

func (h HttpFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:108
	_go_fuzz_dep_.CoverTab[117629]++
										return h.source.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:109
	// _ = "end of CoverTab[117629]"
}

func (h HttpFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:112
	_go_fuzz_dep_.CoverTab[117630]++
										return h.source.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:113
	// _ = "end of CoverTab[117630]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/httpFs.go:114
var _ = _go_fuzz_dep_.CoverTab
