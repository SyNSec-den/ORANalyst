// Copyright ©2015 The Go Authors
// Copyright ©2015 Steve Francia <spf@spf13.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:16
)

import (
	"os"
	"path/filepath"
	"sort"
)

// readDirNames reads the directory named by dirname and returns
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:24
// a sorted list of directory entries.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:24
// adapted from https://golang.org/src/path/filepath/path.go
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:27
func readDirNames(fs Fs, dirname string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:27
	_go_fuzz_dep_.CoverTab[118017]++
										f, err := fs.Open(dirname)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:29
		_go_fuzz_dep_.CoverTab[118020]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:30
		// _ = "end of CoverTab[118020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:31
		_go_fuzz_dep_.CoverTab[118021]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:31
		// _ = "end of CoverTab[118021]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:31
	// _ = "end of CoverTab[118017]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:31
	_go_fuzz_dep_.CoverTab[118018]++
										names, err := f.Readdirnames(-1)
										f.Close()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:34
		_go_fuzz_dep_.CoverTab[118022]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:35
		// _ = "end of CoverTab[118022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:36
		_go_fuzz_dep_.CoverTab[118023]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:36
		// _ = "end of CoverTab[118023]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:36
	// _ = "end of CoverTab[118018]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:36
	_go_fuzz_dep_.CoverTab[118019]++
										sort.Strings(names)
										return names, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:38
	// _ = "end of CoverTab[118019]"
}

// walk recursively descends path, calling walkFn
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:41
// adapted from https://golang.org/src/path/filepath/path.go
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:43
func walk(fs Fs, path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:43
	_go_fuzz_dep_.CoverTab[118024]++
										err := walkFn(path, info, nil)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:45
		_go_fuzz_dep_.CoverTab[118029]++
											if info.IsDir() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:46
			_go_fuzz_dep_.CoverTab[118031]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:46
			return err == filepath.SkipDir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:46
			// _ = "end of CoverTab[118031]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:46
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:46
			_go_fuzz_dep_.CoverTab[118032]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:47
			// _ = "end of CoverTab[118032]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:48
			_go_fuzz_dep_.CoverTab[118033]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:48
			// _ = "end of CoverTab[118033]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:48
		// _ = "end of CoverTab[118029]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:48
		_go_fuzz_dep_.CoverTab[118030]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:49
		// _ = "end of CoverTab[118030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:50
		_go_fuzz_dep_.CoverTab[118034]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:50
		// _ = "end of CoverTab[118034]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:50
	// _ = "end of CoverTab[118024]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:50
	_go_fuzz_dep_.CoverTab[118025]++

										if !info.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:52
		_go_fuzz_dep_.CoverTab[118035]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:53
		// _ = "end of CoverTab[118035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:54
		_go_fuzz_dep_.CoverTab[118036]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:54
		// _ = "end of CoverTab[118036]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:54
	// _ = "end of CoverTab[118025]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:54
	_go_fuzz_dep_.CoverTab[118026]++

										names, err := readDirNames(fs, path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:57
		_go_fuzz_dep_.CoverTab[118037]++
											return walkFn(path, info, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:58
		// _ = "end of CoverTab[118037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:59
		_go_fuzz_dep_.CoverTab[118038]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:59
		// _ = "end of CoverTab[118038]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:59
	// _ = "end of CoverTab[118026]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:59
	_go_fuzz_dep_.CoverTab[118027]++

										for _, name := range names {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:61
		_go_fuzz_dep_.CoverTab[118039]++
											filename := filepath.Join(path, name)
											fileInfo, err := lstatIfPossible(fs, filename)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:64
			_go_fuzz_dep_.CoverTab[118040]++
												if err := walkFn(filename, fileInfo, err); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:65
				_go_fuzz_dep_.CoverTab[118041]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:65
				return err != filepath.SkipDir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:65
				// _ = "end of CoverTab[118041]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:65
			}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:65
				_go_fuzz_dep_.CoverTab[118042]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:66
				// _ = "end of CoverTab[118042]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:67
				_go_fuzz_dep_.CoverTab[118043]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:67
				// _ = "end of CoverTab[118043]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:67
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:67
			// _ = "end of CoverTab[118040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:68
			_go_fuzz_dep_.CoverTab[118044]++
												err = walk(fs, filename, fileInfo, walkFn)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:70
				_go_fuzz_dep_.CoverTab[118045]++
													if !fileInfo.IsDir() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:71
					_go_fuzz_dep_.CoverTab[118046]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:71
					return err != filepath.SkipDir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:71
					// _ = "end of CoverTab[118046]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:71
				}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:71
					_go_fuzz_dep_.CoverTab[118047]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:72
					// _ = "end of CoverTab[118047]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:73
					_go_fuzz_dep_.CoverTab[118048]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:73
					// _ = "end of CoverTab[118048]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:73
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:73
				// _ = "end of CoverTab[118045]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:74
				_go_fuzz_dep_.CoverTab[118049]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:74
				// _ = "end of CoverTab[118049]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:74
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:74
			// _ = "end of CoverTab[118044]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:75
		// _ = "end of CoverTab[118039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:76
	// _ = "end of CoverTab[118027]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:76
	_go_fuzz_dep_.CoverTab[118028]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:77
	// _ = "end of CoverTab[118028]"
}

// if the filesystem supports it, use Lstat, else use fs.Stat
func lstatIfPossible(fs Fs, path string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:81
	_go_fuzz_dep_.CoverTab[118050]++
										if lfs, ok := fs.(Lstater); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:82
		_go_fuzz_dep_.CoverTab[118052]++
											fi, _, err := lfs.LstatIfPossible(path)
											return fi, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:84
		// _ = "end of CoverTab[118052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:85
		_go_fuzz_dep_.CoverTab[118053]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:85
		// _ = "end of CoverTab[118053]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:85
	// _ = "end of CoverTab[118050]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:85
	_go_fuzz_dep_.CoverTab[118051]++
										return fs.Stat(path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:86
	// _ = "end of CoverTab[118051]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:96
func (a Afero) Walk(root string, walkFn filepath.WalkFunc) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:96
	_go_fuzz_dep_.CoverTab[118054]++
										return Walk(a.Fs, root, walkFn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:97
	// _ = "end of CoverTab[118054]"
}

func Walk(fs Fs, root string, walkFn filepath.WalkFunc) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:100
	_go_fuzz_dep_.CoverTab[118055]++
										info, err := lstatIfPossible(fs, root)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:102
		_go_fuzz_dep_.CoverTab[118057]++
											return walkFn(root, nil, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:103
		// _ = "end of CoverTab[118057]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:104
		_go_fuzz_dep_.CoverTab[118058]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:104
		// _ = "end of CoverTab[118058]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:104
	// _ = "end of CoverTab[118055]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:104
	_go_fuzz_dep_.CoverTab[118056]++
										return walk(fs, root, info, walkFn)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:105
	// _ = "end of CoverTab[118056]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:106
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/path.go:106
var _ = _go_fuzz_dep_.CoverTab
