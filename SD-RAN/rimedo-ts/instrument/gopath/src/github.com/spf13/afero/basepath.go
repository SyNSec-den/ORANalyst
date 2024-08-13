//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:1
)

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	_	Lstater		= (*BasePathFs)(nil)
	_	fs.ReadDirFile	= (*BasePathFile)(nil)
)

// The BasePathFs restricts all operations to a given path within an Fs.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// The given file name to the operations on this Fs will be prepended with
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// the base path before calling the base Fs.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// Any file name (after filepath.Clean()) outside this base path will be
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// treated as non existing file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// Note that it does not clean the error messages on return, so you may
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:17
// reveal the real path on errors.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:25
type BasePathFs struct {
	source	Fs
	path	string
}

type BasePathFile struct {
	File
	path	string
}

func (f *BasePathFile) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:35
	_go_fuzz_dep_.CoverTab[117143]++
										sourcename := f.File.Name()
										return strings.TrimPrefix(sourcename, filepath.Clean(f.path))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:37
	// _ = "end of CoverTab[117143]"
}

func (f *BasePathFile) ReadDir(n int) ([]fs.DirEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:40
	_go_fuzz_dep_.CoverTab[117144]++
										if rdf, ok := f.File.(fs.ReadDirFile); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:41
		_go_fuzz_dep_.CoverTab[117146]++
											return rdf.ReadDir(n)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:42
		// _ = "end of CoverTab[117146]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:44
		_go_fuzz_dep_.CoverTab[117147]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:44
		// _ = "end of CoverTab[117147]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:44
	// _ = "end of CoverTab[117144]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:44
	_go_fuzz_dep_.CoverTab[117145]++
										return readDirFile{f.File}.ReadDir(n)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:45
	// _ = "end of CoverTab[117145]"
}

func NewBasePathFs(source Fs, path string) Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:48
	_go_fuzz_dep_.CoverTab[117148]++
										return &BasePathFs{source: source, path: path}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:49
	// _ = "end of CoverTab[117148]"
}

// on a file outside the base path it returns the given file name and an error,
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:52
// else the given file with the base path prepended
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:54
func (b *BasePathFs) RealPath(name string) (path string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:54
	_go_fuzz_dep_.CoverTab[117149]++
										if err := validateBasePathName(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:55
		_go_fuzz_dep_.CoverTab[117152]++
											return name, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:56
		// _ = "end of CoverTab[117152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:57
		_go_fuzz_dep_.CoverTab[117153]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:57
		// _ = "end of CoverTab[117153]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:57
	// _ = "end of CoverTab[117149]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:57
	_go_fuzz_dep_.CoverTab[117150]++

										bpath := filepath.Clean(b.path)
										path = filepath.Clean(filepath.Join(bpath, name))
										if !strings.HasPrefix(path, bpath) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:61
		_go_fuzz_dep_.CoverTab[117154]++
											return name, os.ErrNotExist
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:62
		// _ = "end of CoverTab[117154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:63
		_go_fuzz_dep_.CoverTab[117155]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:63
		// _ = "end of CoverTab[117155]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:63
	// _ = "end of CoverTab[117150]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:63
	_go_fuzz_dep_.CoverTab[117151]++

										return path, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:65
	// _ = "end of CoverTab[117151]"
}

func validateBasePathName(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:68
	_go_fuzz_dep_.CoverTab[117156]++
										if runtime.GOOS != "windows" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:69
		_go_fuzz_dep_.CoverTab[117159]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:72
		return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:72
		// _ = "end of CoverTab[117159]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:73
		_go_fuzz_dep_.CoverTab[117160]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:73
		// _ = "end of CoverTab[117160]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:73
	// _ = "end of CoverTab[117156]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:73
	_go_fuzz_dep_.CoverTab[117157]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:77
	if filepath.IsAbs(name) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:77
		_go_fuzz_dep_.CoverTab[117161]++
											return os.ErrNotExist
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:78
		// _ = "end of CoverTab[117161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:79
		_go_fuzz_dep_.CoverTab[117162]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:79
		// _ = "end of CoverTab[117162]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:79
	// _ = "end of CoverTab[117157]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:79
	_go_fuzz_dep_.CoverTab[117158]++

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:81
	// _ = "end of CoverTab[117158]"
}

func (b *BasePathFs) Chtimes(name string, atime, mtime time.Time) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:84
	_go_fuzz_dep_.CoverTab[117163]++
										if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:85
		_go_fuzz_dep_.CoverTab[117165]++
											return &os.PathError{Op: "chtimes", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:86
		// _ = "end of CoverTab[117165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:87
		_go_fuzz_dep_.CoverTab[117166]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:87
		// _ = "end of CoverTab[117166]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:87
	// _ = "end of CoverTab[117163]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:87
	_go_fuzz_dep_.CoverTab[117164]++
										return b.source.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:88
	// _ = "end of CoverTab[117164]"
}

func (b *BasePathFs) Chmod(name string, mode os.FileMode) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:91
	_go_fuzz_dep_.CoverTab[117167]++
										if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:92
		_go_fuzz_dep_.CoverTab[117169]++
											return &os.PathError{Op: "chmod", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:93
		// _ = "end of CoverTab[117169]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:94
		_go_fuzz_dep_.CoverTab[117170]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:94
		// _ = "end of CoverTab[117170]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:94
	// _ = "end of CoverTab[117167]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:94
	_go_fuzz_dep_.CoverTab[117168]++
										return b.source.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:95
	// _ = "end of CoverTab[117168]"
}

func (b *BasePathFs) Chown(name string, uid, gid int) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:98
	_go_fuzz_dep_.CoverTab[117171]++
										if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:99
			_go_fuzz_dep_.CoverTab[117173]++
												return &os.PathError{Op: "chown", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:100
		// _ = "end of CoverTab[117173]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:101
		_go_fuzz_dep_.CoverTab[117174]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:101
		// _ = "end of CoverTab[117174]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:101
	// _ = "end of CoverTab[117171]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:101
	_go_fuzz_dep_.CoverTab[117172]++
											return b.source.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:102
	// _ = "end of CoverTab[117172]"
}

func (b *BasePathFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:105
	_go_fuzz_dep_.CoverTab[117175]++
											return "BasePathFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:106
	// _ = "end of CoverTab[117175]"
}

func (b *BasePathFs) Stat(name string) (fi os.FileInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:109
	_go_fuzz_dep_.CoverTab[117176]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:110
		_go_fuzz_dep_.CoverTab[117178]++
												return nil, &os.PathError{Op: "stat", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:111
		// _ = "end of CoverTab[117178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:112
		_go_fuzz_dep_.CoverTab[117179]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:112
		// _ = "end of CoverTab[117179]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:112
	// _ = "end of CoverTab[117176]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:112
	_go_fuzz_dep_.CoverTab[117177]++
											return b.source.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:113
	// _ = "end of CoverTab[117177]"
}

func (b *BasePathFs) Rename(oldname, newname string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:116
	_go_fuzz_dep_.CoverTab[117180]++
											if oldname, err = b.RealPath(oldname); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:117
		_go_fuzz_dep_.CoverTab[117183]++
												return &os.PathError{Op: "rename", Path: oldname, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:118
		// _ = "end of CoverTab[117183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:119
		_go_fuzz_dep_.CoverTab[117184]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:119
		// _ = "end of CoverTab[117184]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:119
	// _ = "end of CoverTab[117180]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:119
	_go_fuzz_dep_.CoverTab[117181]++
											if newname, err = b.RealPath(newname); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:120
		_go_fuzz_dep_.CoverTab[117185]++
												return &os.PathError{Op: "rename", Path: newname, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:121
		// _ = "end of CoverTab[117185]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:122
		_go_fuzz_dep_.CoverTab[117186]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:122
		// _ = "end of CoverTab[117186]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:122
	// _ = "end of CoverTab[117181]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:122
	_go_fuzz_dep_.CoverTab[117182]++
											return b.source.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:123
	// _ = "end of CoverTab[117182]"
}

func (b *BasePathFs) RemoveAll(name string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:126
	_go_fuzz_dep_.CoverTab[117187]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:127
		_go_fuzz_dep_.CoverTab[117189]++
												return &os.PathError{Op: "remove_all", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:128
		// _ = "end of CoverTab[117189]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:129
		_go_fuzz_dep_.CoverTab[117190]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:129
		// _ = "end of CoverTab[117190]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:129
	// _ = "end of CoverTab[117187]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:129
	_go_fuzz_dep_.CoverTab[117188]++
											return b.source.RemoveAll(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:130
	// _ = "end of CoverTab[117188]"
}

func (b *BasePathFs) Remove(name string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:133
	_go_fuzz_dep_.CoverTab[117191]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:134
		_go_fuzz_dep_.CoverTab[117193]++
												return &os.PathError{Op: "remove", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:135
		// _ = "end of CoverTab[117193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:136
		_go_fuzz_dep_.CoverTab[117194]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:136
		// _ = "end of CoverTab[117194]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:136
	// _ = "end of CoverTab[117191]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:136
	_go_fuzz_dep_.CoverTab[117192]++
											return b.source.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:137
	// _ = "end of CoverTab[117192]"
}

func (b *BasePathFs) OpenFile(name string, flag int, mode os.FileMode) (f File, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:140
	_go_fuzz_dep_.CoverTab[117195]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:141
		_go_fuzz_dep_.CoverTab[117198]++
												return nil, &os.PathError{Op: "openfile", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:142
		// _ = "end of CoverTab[117198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:143
		_go_fuzz_dep_.CoverTab[117199]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:143
		// _ = "end of CoverTab[117199]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:143
	// _ = "end of CoverTab[117195]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:143
	_go_fuzz_dep_.CoverTab[117196]++
											sourcef, err := b.source.OpenFile(name, flag, mode)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:145
		_go_fuzz_dep_.CoverTab[117200]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:146
		// _ = "end of CoverTab[117200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:147
		_go_fuzz_dep_.CoverTab[117201]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:147
		// _ = "end of CoverTab[117201]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:147
	// _ = "end of CoverTab[117196]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:147
	_go_fuzz_dep_.CoverTab[117197]++
											return &BasePathFile{sourcef, b.path}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:148
	// _ = "end of CoverTab[117197]"
}

func (b *BasePathFs) Open(name string) (f File, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:151
	_go_fuzz_dep_.CoverTab[117202]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:152
		_go_fuzz_dep_.CoverTab[117205]++
												return nil, &os.PathError{Op: "open", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:153
		// _ = "end of CoverTab[117205]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:154
		_go_fuzz_dep_.CoverTab[117206]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:154
		// _ = "end of CoverTab[117206]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:154
	// _ = "end of CoverTab[117202]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:154
	_go_fuzz_dep_.CoverTab[117203]++
											sourcef, err := b.source.Open(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:156
		_go_fuzz_dep_.CoverTab[117207]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:157
		// _ = "end of CoverTab[117207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:158
		_go_fuzz_dep_.CoverTab[117208]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:158
		// _ = "end of CoverTab[117208]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:158
	// _ = "end of CoverTab[117203]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:158
	_go_fuzz_dep_.CoverTab[117204]++
											return &BasePathFile{File: sourcef, path: b.path}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:159
	// _ = "end of CoverTab[117204]"
}

func (b *BasePathFs) Mkdir(name string, mode os.FileMode) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:162
	_go_fuzz_dep_.CoverTab[117209]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:163
		_go_fuzz_dep_.CoverTab[117211]++
												return &os.PathError{Op: "mkdir", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:164
		// _ = "end of CoverTab[117211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:165
		_go_fuzz_dep_.CoverTab[117212]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:165
		// _ = "end of CoverTab[117212]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:165
	// _ = "end of CoverTab[117209]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:165
	_go_fuzz_dep_.CoverTab[117210]++
											return b.source.Mkdir(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:166
	// _ = "end of CoverTab[117210]"
}

func (b *BasePathFs) MkdirAll(name string, mode os.FileMode) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:169
	_go_fuzz_dep_.CoverTab[117213]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:170
		_go_fuzz_dep_.CoverTab[117215]++
												return &os.PathError{Op: "mkdir", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:171
		// _ = "end of CoverTab[117215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:172
		_go_fuzz_dep_.CoverTab[117216]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:172
		// _ = "end of CoverTab[117216]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:172
	// _ = "end of CoverTab[117213]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:172
	_go_fuzz_dep_.CoverTab[117214]++
											return b.source.MkdirAll(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:173
	// _ = "end of CoverTab[117214]"
}

func (b *BasePathFs) Create(name string) (f File, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:176
	_go_fuzz_dep_.CoverTab[117217]++
											if name, err = b.RealPath(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:177
		_go_fuzz_dep_.CoverTab[117220]++
												return nil, &os.PathError{Op: "create", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:178
		// _ = "end of CoverTab[117220]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:179
		_go_fuzz_dep_.CoverTab[117221]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:179
		// _ = "end of CoverTab[117221]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:179
	// _ = "end of CoverTab[117217]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:179
	_go_fuzz_dep_.CoverTab[117218]++
											sourcef, err := b.source.Create(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:181
		_go_fuzz_dep_.CoverTab[117222]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:182
		// _ = "end of CoverTab[117222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:183
		_go_fuzz_dep_.CoverTab[117223]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:183
		// _ = "end of CoverTab[117223]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:183
	// _ = "end of CoverTab[117218]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:183
	_go_fuzz_dep_.CoverTab[117219]++
											return &BasePathFile{File: sourcef, path: b.path}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:184
	// _ = "end of CoverTab[117219]"
}

func (b *BasePathFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:187
	_go_fuzz_dep_.CoverTab[117224]++
											name, err := b.RealPath(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:189
		_go_fuzz_dep_.CoverTab[117227]++
												return nil, false, &os.PathError{Op: "lstat", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:190
		// _ = "end of CoverTab[117227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:191
		_go_fuzz_dep_.CoverTab[117228]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:191
		// _ = "end of CoverTab[117228]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:191
	// _ = "end of CoverTab[117224]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:191
	_go_fuzz_dep_.CoverTab[117225]++
											if lstater, ok := b.source.(Lstater); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:192
		_go_fuzz_dep_.CoverTab[117229]++
												return lstater.LstatIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:193
		// _ = "end of CoverTab[117229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:194
		_go_fuzz_dep_.CoverTab[117230]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:194
		// _ = "end of CoverTab[117230]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:194
	// _ = "end of CoverTab[117225]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:194
	_go_fuzz_dep_.CoverTab[117226]++
											fi, err := b.source.Stat(name)
											return fi, false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:196
	// _ = "end of CoverTab[117226]"
}

func (b *BasePathFs) SymlinkIfPossible(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:199
	_go_fuzz_dep_.CoverTab[117231]++
											oldname, err := b.RealPath(oldname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:201
		_go_fuzz_dep_.CoverTab[117235]++
												return &os.LinkError{Op: "symlink", Old: oldname, New: newname, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:202
		// _ = "end of CoverTab[117235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:203
		_go_fuzz_dep_.CoverTab[117236]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:203
		// _ = "end of CoverTab[117236]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:203
	// _ = "end of CoverTab[117231]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:203
	_go_fuzz_dep_.CoverTab[117232]++
											newname, err = b.RealPath(newname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:205
		_go_fuzz_dep_.CoverTab[117237]++
												return &os.LinkError{Op: "symlink", Old: oldname, New: newname, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:206
		// _ = "end of CoverTab[117237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:207
		_go_fuzz_dep_.CoverTab[117238]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:207
		// _ = "end of CoverTab[117238]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:207
	// _ = "end of CoverTab[117232]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:207
	_go_fuzz_dep_.CoverTab[117233]++
											if linker, ok := b.source.(Linker); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:208
		_go_fuzz_dep_.CoverTab[117239]++
												return linker.SymlinkIfPossible(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:209
		// _ = "end of CoverTab[117239]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:210
		_go_fuzz_dep_.CoverTab[117240]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:210
		// _ = "end of CoverTab[117240]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:210
	// _ = "end of CoverTab[117233]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:210
	_go_fuzz_dep_.CoverTab[117234]++
											return &os.LinkError{Op: "symlink", Old: oldname, New: newname, Err: ErrNoSymlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:211
	// _ = "end of CoverTab[117234]"
}

func (b *BasePathFs) ReadlinkIfPossible(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:214
	_go_fuzz_dep_.CoverTab[117241]++
											name, err := b.RealPath(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:216
		_go_fuzz_dep_.CoverTab[117244]++
												return "", &os.PathError{Op: "readlink", Path: name, Err: err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:217
		// _ = "end of CoverTab[117244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:218
		_go_fuzz_dep_.CoverTab[117245]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:218
		// _ = "end of CoverTab[117245]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:218
	// _ = "end of CoverTab[117241]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:218
	_go_fuzz_dep_.CoverTab[117242]++
											if reader, ok := b.source.(LinkReader); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:219
		_go_fuzz_dep_.CoverTab[117246]++
												return reader.ReadlinkIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:220
		// _ = "end of CoverTab[117246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:221
		_go_fuzz_dep_.CoverTab[117247]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:221
		// _ = "end of CoverTab[117247]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:221
	// _ = "end of CoverTab[117242]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:221
	_go_fuzz_dep_.CoverTab[117243]++
											return "", &os.PathError{Op: "readlink", Path: name, Err: ErrNoReadlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:222
	// _ = "end of CoverTab[117243]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:223
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/basepath.go:223
var _ = _go_fuzz_dep_.CoverTab
