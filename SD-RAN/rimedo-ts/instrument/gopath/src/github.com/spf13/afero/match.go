// Copyright © 2014 Steve Francia <spf@spf13.com>.
// Copyright 2009 The Go Authors. All rights reserved.

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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:15
)

import (
	"path/filepath"
	"sort"
	"strings"
)

// Glob returns the names of all files matching pattern or nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// if there is no matching file. The syntax of patterns is the same
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// as in Match. The pattern may describe hierarchical names such as
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// /usr/*/bin/ed (assuming the Separator is '/').
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// Glob ignores file system errors such as I/O errors reading directories.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// The only possible returned error is ErrBadPattern, when pattern
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// is malformed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// This was adapted from (http://golang.org/pkg/path/filepath) and uses several
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:23
// built-ins from that package.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:34
func Glob(fs Fs, pattern string) (matches []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:34
	_go_fuzz_dep_.CoverTab[117815]++
										if !hasMeta(pattern) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:35
		_go_fuzz_dep_.CoverTab[117821]++

											if _, err = lstatIfPossible(fs, pattern); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:37
			_go_fuzz_dep_.CoverTab[117823]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:38
			// _ = "end of CoverTab[117823]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:39
			_go_fuzz_dep_.CoverTab[117824]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:39
			// _ = "end of CoverTab[117824]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:39
		// _ = "end of CoverTab[117821]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:39
		_go_fuzz_dep_.CoverTab[117822]++
											return []string{pattern}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:40
		// _ = "end of CoverTab[117822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:41
		_go_fuzz_dep_.CoverTab[117825]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:41
		// _ = "end of CoverTab[117825]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:41
	// _ = "end of CoverTab[117815]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:41
	_go_fuzz_dep_.CoverTab[117816]++

										dir, file := filepath.Split(pattern)
										switch dir {
	case "":
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:45
		_go_fuzz_dep_.CoverTab[117826]++
											dir = "."
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:46
		// _ = "end of CoverTab[117826]"
	case string(filepath.Separator):
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:47
		_go_fuzz_dep_.CoverTab[117827]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:47
		// _ = "end of CoverTab[117827]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:49
		_go_fuzz_dep_.CoverTab[117828]++
											dir = dir[0 : len(dir)-1]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:50
		// _ = "end of CoverTab[117828]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:51
	// _ = "end of CoverTab[117816]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:51
	_go_fuzz_dep_.CoverTab[117817]++

										if !hasMeta(dir) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:53
		_go_fuzz_dep_.CoverTab[117829]++
											return glob(fs, dir, file, nil)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:54
		// _ = "end of CoverTab[117829]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:55
		_go_fuzz_dep_.CoverTab[117830]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:55
		// _ = "end of CoverTab[117830]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:55
	// _ = "end of CoverTab[117817]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:55
	_go_fuzz_dep_.CoverTab[117818]++

										var m []string
										m, err = Glob(fs, dir)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:59
		_go_fuzz_dep_.CoverTab[117831]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:60
		// _ = "end of CoverTab[117831]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:61
		_go_fuzz_dep_.CoverTab[117832]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:61
		// _ = "end of CoverTab[117832]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:61
	// _ = "end of CoverTab[117818]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:61
	_go_fuzz_dep_.CoverTab[117819]++
										for _, d := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:62
		_go_fuzz_dep_.CoverTab[117833]++
											matches, err = glob(fs, d, file, matches)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:64
			_go_fuzz_dep_.CoverTab[117834]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:65
			// _ = "end of CoverTab[117834]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:66
			_go_fuzz_dep_.CoverTab[117835]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:66
			// _ = "end of CoverTab[117835]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:66
		// _ = "end of CoverTab[117833]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:67
	// _ = "end of CoverTab[117819]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:67
	_go_fuzz_dep_.CoverTab[117820]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:68
	// _ = "end of CoverTab[117820]"
}

// glob searches for files matching pattern in the directory dir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:71
// and appends them to matches. If the directory cannot be
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:71
// opened, it returns the existing matches. New matches are
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:71
// added in lexicographical order.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:75
func glob(fs Fs, dir, pattern string, matches []string) (m []string, e error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:75
	_go_fuzz_dep_.CoverTab[117836]++
										m = matches
										fi, err := fs.Stat(dir)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:78
		_go_fuzz_dep_.CoverTab[117841]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:79
		// _ = "end of CoverTab[117841]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:80
		_go_fuzz_dep_.CoverTab[117842]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:80
		// _ = "end of CoverTab[117842]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:80
	// _ = "end of CoverTab[117836]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:80
	_go_fuzz_dep_.CoverTab[117837]++
										if !fi.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:81
		_go_fuzz_dep_.CoverTab[117843]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:82
		// _ = "end of CoverTab[117843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:83
		_go_fuzz_dep_.CoverTab[117844]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:83
		// _ = "end of CoverTab[117844]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:83
	// _ = "end of CoverTab[117837]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:83
	_go_fuzz_dep_.CoverTab[117838]++
										d, err := fs.Open(dir)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:85
		_go_fuzz_dep_.CoverTab[117845]++
											return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:86
		// _ = "end of CoverTab[117845]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:87
		_go_fuzz_dep_.CoverTab[117846]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:87
		// _ = "end of CoverTab[117846]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:87
	// _ = "end of CoverTab[117838]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:87
	_go_fuzz_dep_.CoverTab[117839]++
										defer d.Close()

										names, _ := d.Readdirnames(-1)
										sort.Strings(names)

										for _, n := range names {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:93
		_go_fuzz_dep_.CoverTab[117847]++
											matched, err := filepath.Match(pattern, n)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:95
			_go_fuzz_dep_.CoverTab[117849]++
												return m, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:96
			// _ = "end of CoverTab[117849]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:97
			_go_fuzz_dep_.CoverTab[117850]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:97
			// _ = "end of CoverTab[117850]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:97
		// _ = "end of CoverTab[117847]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:97
		_go_fuzz_dep_.CoverTab[117848]++
											if matched {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:98
			_go_fuzz_dep_.CoverTab[117851]++
												m = append(m, filepath.Join(dir, n))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:99
			// _ = "end of CoverTab[117851]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:100
			_go_fuzz_dep_.CoverTab[117852]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:100
			// _ = "end of CoverTab[117852]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:100
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:100
		// _ = "end of CoverTab[117848]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:101
	// _ = "end of CoverTab[117839]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:101
	_go_fuzz_dep_.CoverTab[117840]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:102
	// _ = "end of CoverTab[117840]"
}

// hasMeta reports whether path contains any of the magic characters
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:105
// recognized by Match.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:107
func hasMeta(path string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:107
	_go_fuzz_dep_.CoverTab[117853]++

										return strings.ContainsAny(path, "*?[")
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:109
	// _ = "end of CoverTab[117853]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:110
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/match.go:110
var _ = _go_fuzz_dep_.CoverTab
