// Copyright ©2015 Steve Francia <spf@spf13.com>
// Portions Copyright ©2015 The Hugo Authors
// Portions Copyright 2016-present Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:17
)

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Filepath separator defined by os.Separator.
const FilePathSeparator = string(filepath.Separator)

// Takes a reader and a path and writes the content
func (a Afero) WriteReader(path string, r io.Reader) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:37
	_go_fuzz_dep_.CoverTab[118379]++
										return WriteReader(a.Fs, path, r)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:38
	// _ = "end of CoverTab[118379]"
}

func WriteReader(fs Fs, path string, r io.Reader) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:41
	_go_fuzz_dep_.CoverTab[118380]++
										dir, _ := filepath.Split(path)
										ospath := filepath.FromSlash(dir)

										if ospath != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:45
		_go_fuzz_dep_.CoverTab[118383]++
											err = fs.MkdirAll(ospath, 0777)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:47
			_go_fuzz_dep_.CoverTab[118384]++
												if err != os.ErrExist {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:48
				_go_fuzz_dep_.CoverTab[118385]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:49
				// _ = "end of CoverTab[118385]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:50
				_go_fuzz_dep_.CoverTab[118386]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:50
				// _ = "end of CoverTab[118386]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:50
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:50
			// _ = "end of CoverTab[118384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:51
			_go_fuzz_dep_.CoverTab[118387]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:51
			// _ = "end of CoverTab[118387]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:51
		// _ = "end of CoverTab[118383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:52
		_go_fuzz_dep_.CoverTab[118388]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:52
		// _ = "end of CoverTab[118388]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:52
	// _ = "end of CoverTab[118380]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:52
	_go_fuzz_dep_.CoverTab[118381]++

										file, err := fs.Create(path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:55
		_go_fuzz_dep_.CoverTab[118389]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:56
		// _ = "end of CoverTab[118389]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:57
		_go_fuzz_dep_.CoverTab[118390]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:57
		// _ = "end of CoverTab[118390]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:57
	// _ = "end of CoverTab[118381]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:57
	_go_fuzz_dep_.CoverTab[118382]++
										defer file.Close()

										_, err = io.Copy(file, r)
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:61
	// _ = "end of CoverTab[118382]"
}

// Same as WriteReader but checks to see if file/directory already exists.
func (a Afero) SafeWriteReader(path string, r io.Reader) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:65
	_go_fuzz_dep_.CoverTab[118391]++
										return SafeWriteReader(a.Fs, path, r)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:66
	// _ = "end of CoverTab[118391]"
}

func SafeWriteReader(fs Fs, path string, r io.Reader) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:69
	_go_fuzz_dep_.CoverTab[118392]++
										dir, _ := filepath.Split(path)
										ospath := filepath.FromSlash(dir)

										if ospath != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:73
		_go_fuzz_dep_.CoverTab[118397]++
											err = fs.MkdirAll(ospath, 0777)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:75
			_go_fuzz_dep_.CoverTab[118398]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:76
			// _ = "end of CoverTab[118398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:77
			_go_fuzz_dep_.CoverTab[118399]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:77
			// _ = "end of CoverTab[118399]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:77
		// _ = "end of CoverTab[118397]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:78
		_go_fuzz_dep_.CoverTab[118400]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:78
		// _ = "end of CoverTab[118400]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:78
	// _ = "end of CoverTab[118392]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:78
	_go_fuzz_dep_.CoverTab[118393]++

										exists, err := Exists(fs, path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:81
		_go_fuzz_dep_.CoverTab[118401]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:82
		// _ = "end of CoverTab[118401]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:83
		_go_fuzz_dep_.CoverTab[118402]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:83
		// _ = "end of CoverTab[118402]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:83
	// _ = "end of CoverTab[118393]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:83
	_go_fuzz_dep_.CoverTab[118394]++
										if exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:84
		_go_fuzz_dep_.CoverTab[118403]++
											return fmt.Errorf("%v already exists", path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:85
		// _ = "end of CoverTab[118403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:86
		_go_fuzz_dep_.CoverTab[118404]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:86
		// _ = "end of CoverTab[118404]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:86
	// _ = "end of CoverTab[118394]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:86
	_go_fuzz_dep_.CoverTab[118395]++

										file, err := fs.Create(path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:89
		_go_fuzz_dep_.CoverTab[118405]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:90
		// _ = "end of CoverTab[118405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:91
		_go_fuzz_dep_.CoverTab[118406]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:91
		// _ = "end of CoverTab[118406]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:91
	// _ = "end of CoverTab[118395]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:91
	_go_fuzz_dep_.CoverTab[118396]++
										defer file.Close()

										_, err = io.Copy(file, r)
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:95
	// _ = "end of CoverTab[118396]"
}

func (a Afero) GetTempDir(subPath string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:98
	_go_fuzz_dep_.CoverTab[118407]++
										return GetTempDir(a.Fs, subPath)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:99
	// _ = "end of CoverTab[118407]"
}

// GetTempDir returns the default temp directory with trailing slash
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:102
// if subPath is not empty then it will be created recursively with mode 777 rwx rwx rwx
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:104
func GetTempDir(fs Fs, subPath string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:104
	_go_fuzz_dep_.CoverTab[118408]++
										addSlash := func(p string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:105
		_go_fuzz_dep_.CoverTab[118411]++
											if FilePathSeparator != p[len(p)-1:] {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:106
			_go_fuzz_dep_.CoverTab[118413]++
												p = p + FilePathSeparator
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:107
			// _ = "end of CoverTab[118413]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:108
			_go_fuzz_dep_.CoverTab[118414]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:108
			// _ = "end of CoverTab[118414]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:108
		// _ = "end of CoverTab[118411]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:108
		_go_fuzz_dep_.CoverTab[118412]++
											return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:109
		// _ = "end of CoverTab[118412]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:110
	// _ = "end of CoverTab[118408]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:110
	_go_fuzz_dep_.CoverTab[118409]++
										dir := addSlash(os.TempDir())

										if subPath != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:113
		_go_fuzz_dep_.CoverTab[118415]++

											if FilePathSeparator == "\\" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:115
			_go_fuzz_dep_.CoverTab[118420]++
												subPath = strings.Replace(subPath, "\\", "____", -1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:116
			// _ = "end of CoverTab[118420]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:117
			_go_fuzz_dep_.CoverTab[118421]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:117
			// _ = "end of CoverTab[118421]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:117
		// _ = "end of CoverTab[118415]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:117
		_go_fuzz_dep_.CoverTab[118416]++
											dir = dir + UnicodeSanitize((subPath))
											if FilePathSeparator == "\\" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:119
			_go_fuzz_dep_.CoverTab[118422]++
												dir = strings.Replace(dir, "____", "\\", -1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:120
			// _ = "end of CoverTab[118422]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:121
			_go_fuzz_dep_.CoverTab[118423]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:121
			// _ = "end of CoverTab[118423]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:121
		// _ = "end of CoverTab[118416]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:121
		_go_fuzz_dep_.CoverTab[118417]++

											if exists, _ := Exists(fs, dir); exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:123
			_go_fuzz_dep_.CoverTab[118424]++
												return addSlash(dir)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:124
			// _ = "end of CoverTab[118424]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:125
			_go_fuzz_dep_.CoverTab[118425]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:125
			// _ = "end of CoverTab[118425]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:125
		// _ = "end of CoverTab[118417]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:125
		_go_fuzz_dep_.CoverTab[118418]++

											err := fs.MkdirAll(dir, 0777)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:128
			_go_fuzz_dep_.CoverTab[118426]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:129
			// _ = "end of CoverTab[118426]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:130
			_go_fuzz_dep_.CoverTab[118427]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:130
			// _ = "end of CoverTab[118427]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:130
		// _ = "end of CoverTab[118418]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:130
		_go_fuzz_dep_.CoverTab[118419]++
											dir = addSlash(dir)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:131
		// _ = "end of CoverTab[118419]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:132
		_go_fuzz_dep_.CoverTab[118428]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:132
		// _ = "end of CoverTab[118428]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:132
	// _ = "end of CoverTab[118409]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:132
	_go_fuzz_dep_.CoverTab[118410]++
										return dir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:133
	// _ = "end of CoverTab[118410]"
}

// Rewrite string to remove non-standard path characters
func UnicodeSanitize(s string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:137
	_go_fuzz_dep_.CoverTab[118429]++
										source := []rune(s)
										target := make([]rune, 0, len(source))

										for _, r := range source {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:141
		_go_fuzz_dep_.CoverTab[118431]++
											if unicode.IsLetter(r) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:142
			_go_fuzz_dep_.CoverTab[118432]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:142
			return unicode.IsDigit(r)
												// _ = "end of CoverTab[118432]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:143
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:143
			_go_fuzz_dep_.CoverTab[118433]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:143
			return unicode.IsMark(r)
												// _ = "end of CoverTab[118433]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:144
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:144
			_go_fuzz_dep_.CoverTab[118434]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:144
			return r == '.'
												// _ = "end of CoverTab[118434]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:145
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:145
			_go_fuzz_dep_.CoverTab[118435]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:145
			return r == '/'
												// _ = "end of CoverTab[118435]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:146
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:146
			_go_fuzz_dep_.CoverTab[118436]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:146
			return r == '\\'
												// _ = "end of CoverTab[118436]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:147
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:147
			_go_fuzz_dep_.CoverTab[118437]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:147
			return r == '_'
												// _ = "end of CoverTab[118437]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:148
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:148
			_go_fuzz_dep_.CoverTab[118438]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:148
			return r == '-'
												// _ = "end of CoverTab[118438]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:149
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:149
			_go_fuzz_dep_.CoverTab[118439]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:149
			return r == '%'
												// _ = "end of CoverTab[118439]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:150
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:150
			_go_fuzz_dep_.CoverTab[118440]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:150
			return r == ' '
												// _ = "end of CoverTab[118440]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:151
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:151
			_go_fuzz_dep_.CoverTab[118441]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:151
			return r == '#'
												// _ = "end of CoverTab[118441]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:152
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:152
			_go_fuzz_dep_.CoverTab[118442]++
												target = append(target, r)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:153
			// _ = "end of CoverTab[118442]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:154
			_go_fuzz_dep_.CoverTab[118443]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:154
			// _ = "end of CoverTab[118443]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:154
		// _ = "end of CoverTab[118431]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:155
	// _ = "end of CoverTab[118429]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:155
	_go_fuzz_dep_.CoverTab[118430]++

										return string(target)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:157
	// _ = "end of CoverTab[118430]"
}

// Transform characters with accents into plain forms.
func NeuterAccents(s string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:161
	_go_fuzz_dep_.CoverTab[118444]++
										t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
										result, _, _ := transform.String(t, string(s))

										return result
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:165
	// _ = "end of CoverTab[118444]"
}

func (a Afero) FileContainsBytes(filename string, subslice []byte) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:168
	_go_fuzz_dep_.CoverTab[118445]++
										return FileContainsBytes(a.Fs, filename, subslice)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:169
	// _ = "end of CoverTab[118445]"
}

// Check if a file contains a specified byte slice.
func FileContainsBytes(fs Fs, filename string, subslice []byte) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:173
	_go_fuzz_dep_.CoverTab[118446]++
										f, err := fs.Open(filename)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:175
		_go_fuzz_dep_.CoverTab[118448]++
											return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:176
		// _ = "end of CoverTab[118448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:177
		_go_fuzz_dep_.CoverTab[118449]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:177
		// _ = "end of CoverTab[118449]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:177
	// _ = "end of CoverTab[118446]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:177
	_go_fuzz_dep_.CoverTab[118447]++
										defer f.Close()

										return readerContainsAny(f, subslice), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:180
	// _ = "end of CoverTab[118447]"
}

func (a Afero) FileContainsAnyBytes(filename string, subslices [][]byte) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:183
	_go_fuzz_dep_.CoverTab[118450]++
										return FileContainsAnyBytes(a.Fs, filename, subslices)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:184
	// _ = "end of CoverTab[118450]"
}

// Check if a file contains any of the specified byte slices.
func FileContainsAnyBytes(fs Fs, filename string, subslices [][]byte) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:188
	_go_fuzz_dep_.CoverTab[118451]++
										f, err := fs.Open(filename)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:190
		_go_fuzz_dep_.CoverTab[118453]++
											return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:191
		// _ = "end of CoverTab[118453]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:192
		_go_fuzz_dep_.CoverTab[118454]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:192
		// _ = "end of CoverTab[118454]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:192
	// _ = "end of CoverTab[118451]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:192
	_go_fuzz_dep_.CoverTab[118452]++
										defer f.Close()

										return readerContainsAny(f, subslices...), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:195
	// _ = "end of CoverTab[118452]"
}

// readerContains reports whether any of the subslices is within r.
func readerContainsAny(r io.Reader, subslices ...[]byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:199
	_go_fuzz_dep_.CoverTab[118455]++

										if r == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:201
		_go_fuzz_dep_.CoverTab[118460]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:201
		return len(subslices) == 0
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:201
		// _ = "end of CoverTab[118460]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:201
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:201
		_go_fuzz_dep_.CoverTab[118461]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:202
		// _ = "end of CoverTab[118461]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:203
		_go_fuzz_dep_.CoverTab[118462]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:203
		// _ = "end of CoverTab[118462]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:203
	// _ = "end of CoverTab[118455]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:203
	_go_fuzz_dep_.CoverTab[118456]++

										largestSlice := 0

										for _, sl := range subslices {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:207
		_go_fuzz_dep_.CoverTab[118463]++
											if len(sl) > largestSlice {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:208
			_go_fuzz_dep_.CoverTab[118464]++
												largestSlice = len(sl)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:209
			// _ = "end of CoverTab[118464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:210
			_go_fuzz_dep_.CoverTab[118465]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:210
			// _ = "end of CoverTab[118465]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:210
		// _ = "end of CoverTab[118463]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:211
	// _ = "end of CoverTab[118456]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:211
	_go_fuzz_dep_.CoverTab[118457]++

										if largestSlice == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:213
		_go_fuzz_dep_.CoverTab[118466]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:214
		// _ = "end of CoverTab[118466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:215
		_go_fuzz_dep_.CoverTab[118467]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:215
		// _ = "end of CoverTab[118467]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:215
	// _ = "end of CoverTab[118457]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:215
	_go_fuzz_dep_.CoverTab[118458]++

										bufflen := largestSlice * 4
										halflen := bufflen / 2
										buff := make([]byte, bufflen)
										var err error
										var n, i int

										for {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:223
		_go_fuzz_dep_.CoverTab[118468]++
											i++
											if i == 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:225
			_go_fuzz_dep_.CoverTab[118471]++
												n, err = io.ReadAtLeast(r, buff[:halflen], halflen)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:226
			// _ = "end of CoverTab[118471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:227
			_go_fuzz_dep_.CoverTab[118472]++
												if i != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:228
				_go_fuzz_dep_.CoverTab[118474]++

													copy(buff[:], buff[halflen:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:230
				// _ = "end of CoverTab[118474]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:231
				_go_fuzz_dep_.CoverTab[118475]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:231
				// _ = "end of CoverTab[118475]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:231
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:231
			// _ = "end of CoverTab[118472]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:231
			_go_fuzz_dep_.CoverTab[118473]++
												n, err = io.ReadAtLeast(r, buff[halflen:], halflen)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:232
			// _ = "end of CoverTab[118473]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:233
		// _ = "end of CoverTab[118468]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:233
		_go_fuzz_dep_.CoverTab[118469]++

											if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:235
			_go_fuzz_dep_.CoverTab[118476]++
												for _, sl := range subslices {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:236
				_go_fuzz_dep_.CoverTab[118477]++
													if bytes.Contains(buff, sl) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:237
					_go_fuzz_dep_.CoverTab[118478]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:238
					// _ = "end of CoverTab[118478]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:239
					_go_fuzz_dep_.CoverTab[118479]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:239
					// _ = "end of CoverTab[118479]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:239
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:239
				// _ = "end of CoverTab[118477]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:240
			// _ = "end of CoverTab[118476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:241
			_go_fuzz_dep_.CoverTab[118480]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:241
			// _ = "end of CoverTab[118480]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:241
		// _ = "end of CoverTab[118469]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:241
		_go_fuzz_dep_.CoverTab[118470]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:243
			_go_fuzz_dep_.CoverTab[118481]++
												break
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:244
			// _ = "end of CoverTab[118481]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:245
			_go_fuzz_dep_.CoverTab[118482]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:245
			// _ = "end of CoverTab[118482]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:245
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:245
		// _ = "end of CoverTab[118470]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:246
	// _ = "end of CoverTab[118458]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:246
	_go_fuzz_dep_.CoverTab[118459]++
										return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:247
	// _ = "end of CoverTab[118459]"
}

func (a Afero) DirExists(path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:250
	_go_fuzz_dep_.CoverTab[118483]++
										return DirExists(a.Fs, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:251
	// _ = "end of CoverTab[118483]"
}

// DirExists checks if a path exists and is a directory.
func DirExists(fs Fs, path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:255
	_go_fuzz_dep_.CoverTab[118484]++
										fi, err := fs.Stat(path)
										if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:257
		_go_fuzz_dep_.CoverTab[118487]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:257
		return fi.IsDir()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:257
		// _ = "end of CoverTab[118487]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:257
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:257
		_go_fuzz_dep_.CoverTab[118488]++
											return true, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:258
		// _ = "end of CoverTab[118488]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:259
		_go_fuzz_dep_.CoverTab[118489]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:259
		// _ = "end of CoverTab[118489]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:259
	// _ = "end of CoverTab[118484]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:259
	_go_fuzz_dep_.CoverTab[118485]++
										if os.IsNotExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:260
		_go_fuzz_dep_.CoverTab[118490]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:261
		// _ = "end of CoverTab[118490]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:262
		_go_fuzz_dep_.CoverTab[118491]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:262
		// _ = "end of CoverTab[118491]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:262
	// _ = "end of CoverTab[118485]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:262
	_go_fuzz_dep_.CoverTab[118486]++
										return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:263
	// _ = "end of CoverTab[118486]"
}

func (a Afero) IsDir(path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:266
	_go_fuzz_dep_.CoverTab[118492]++
										return IsDir(a.Fs, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:267
	// _ = "end of CoverTab[118492]"
}

// IsDir checks if a given path is a directory.
func IsDir(fs Fs, path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:271
	_go_fuzz_dep_.CoverTab[118493]++
										fi, err := fs.Stat(path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:273
		_go_fuzz_dep_.CoverTab[118495]++
											return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:274
		// _ = "end of CoverTab[118495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:275
		_go_fuzz_dep_.CoverTab[118496]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:275
		// _ = "end of CoverTab[118496]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:275
	// _ = "end of CoverTab[118493]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:275
	_go_fuzz_dep_.CoverTab[118494]++
										return fi.IsDir(), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:276
	// _ = "end of CoverTab[118494]"
}

func (a Afero) IsEmpty(path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:279
	_go_fuzz_dep_.CoverTab[118497]++
										return IsEmpty(a.Fs, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:280
	// _ = "end of CoverTab[118497]"
}

// IsEmpty checks if a given file or directory is empty.
func IsEmpty(fs Fs, path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:284
	_go_fuzz_dep_.CoverTab[118498]++
										if b, _ := Exists(fs, path); !b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:285
		_go_fuzz_dep_.CoverTab[118502]++
											return false, fmt.Errorf("%q path does not exist", path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:286
		// _ = "end of CoverTab[118502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:287
		_go_fuzz_dep_.CoverTab[118503]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:287
		// _ = "end of CoverTab[118503]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:287
	// _ = "end of CoverTab[118498]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:287
	_go_fuzz_dep_.CoverTab[118499]++
										fi, err := fs.Stat(path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:289
		_go_fuzz_dep_.CoverTab[118504]++
											return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:290
		// _ = "end of CoverTab[118504]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:291
		_go_fuzz_dep_.CoverTab[118505]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:291
		// _ = "end of CoverTab[118505]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:291
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:291
	// _ = "end of CoverTab[118499]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:291
	_go_fuzz_dep_.CoverTab[118500]++
										if fi.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:292
		_go_fuzz_dep_.CoverTab[118506]++
											f, err := fs.Open(path)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:294
			_go_fuzz_dep_.CoverTab[118509]++
												return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:295
			// _ = "end of CoverTab[118509]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:296
			_go_fuzz_dep_.CoverTab[118510]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:296
			// _ = "end of CoverTab[118510]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:296
		// _ = "end of CoverTab[118506]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:296
		_go_fuzz_dep_.CoverTab[118507]++
											defer f.Close()
											list, err := f.Readdir(-1)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:299
			_go_fuzz_dep_.CoverTab[118511]++
												return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:300
			// _ = "end of CoverTab[118511]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:301
			_go_fuzz_dep_.CoverTab[118512]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:301
			// _ = "end of CoverTab[118512]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:301
		// _ = "end of CoverTab[118507]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:301
		_go_fuzz_dep_.CoverTab[118508]++
											return len(list) == 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:302
		// _ = "end of CoverTab[118508]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:303
		_go_fuzz_dep_.CoverTab[118513]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:303
		// _ = "end of CoverTab[118513]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:303
	// _ = "end of CoverTab[118500]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:303
	_go_fuzz_dep_.CoverTab[118501]++
										return fi.Size() == 0, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:304
	// _ = "end of CoverTab[118501]"
}

func (a Afero) Exists(path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:307
	_go_fuzz_dep_.CoverTab[118514]++
										return Exists(a.Fs, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:308
	// _ = "end of CoverTab[118514]"
}

// Check if a file or directory exists.
func Exists(fs Fs, path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:312
	_go_fuzz_dep_.CoverTab[118515]++
										_, err := fs.Stat(path)
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:314
		_go_fuzz_dep_.CoverTab[118518]++
											return true, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:315
		// _ = "end of CoverTab[118518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:316
		_go_fuzz_dep_.CoverTab[118519]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:316
		// _ = "end of CoverTab[118519]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:316
	// _ = "end of CoverTab[118515]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:316
	_go_fuzz_dep_.CoverTab[118516]++
										if os.IsNotExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:317
		_go_fuzz_dep_.CoverTab[118520]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:318
		// _ = "end of CoverTab[118520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:319
		_go_fuzz_dep_.CoverTab[118521]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:319
		// _ = "end of CoverTab[118521]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:319
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:319
	// _ = "end of CoverTab[118516]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:319
	_go_fuzz_dep_.CoverTab[118517]++
										return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:320
	// _ = "end of CoverTab[118517]"
}

func FullBaseFsPath(basePathFs *BasePathFs, relativePath string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:323
	_go_fuzz_dep_.CoverTab[118522]++
										combinedPath := filepath.Join(basePathFs.path, relativePath)
										if parent, ok := basePathFs.source.(*BasePathFs); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:325
		_go_fuzz_dep_.CoverTab[118524]++
											return FullBaseFsPath(parent, combinedPath)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:326
		// _ = "end of CoverTab[118524]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:327
		_go_fuzz_dep_.CoverTab[118525]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:327
		// _ = "end of CoverTab[118525]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:327
	// _ = "end of CoverTab[118522]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:327
	_go_fuzz_dep_.CoverTab[118523]++

										return combinedPath
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:329
	// _ = "end of CoverTab[118523]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:330
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/util.go:330
var _ = _go_fuzz_dep_.CoverTab
