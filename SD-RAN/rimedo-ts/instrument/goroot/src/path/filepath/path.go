// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/path/filepath/path.go:5
// Package filepath implements utility routines for manipulating filename paths
//line /usr/local/go/src/path/filepath/path.go:5
// in a way compatible with the target operating system-defined file paths.
//line /usr/local/go/src/path/filepath/path.go:5
//
//line /usr/local/go/src/path/filepath/path.go:5
// The filepath package uses either forward slashes or backslashes,
//line /usr/local/go/src/path/filepath/path.go:5
// depending on the operating system. To process paths such as URLs
//line /usr/local/go/src/path/filepath/path.go:5
// that always use forward slashes regardless of the operating
//line /usr/local/go/src/path/filepath/path.go:5
// system, see the path package.
//line /usr/local/go/src/path/filepath/path.go:12
package filepath

//line /usr/local/go/src/path/filepath/path.go:12
import (
//line /usr/local/go/src/path/filepath/path.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/path/filepath/path.go:12
)
//line /usr/local/go/src/path/filepath/path.go:12
import (
//line /usr/local/go/src/path/filepath/path.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/path/filepath/path.go:12
)

import (
	"errors"
	"io/fs"
	"os"
	"runtime"
	"sort"
	"strings"
)

// A lazybuf is a lazily constructed path buffer.
//line /usr/local/go/src/path/filepath/path.go:23
// It supports append, reading previously appended bytes,
//line /usr/local/go/src/path/filepath/path.go:23
// and retrieving the final string. It does not allocate a buffer
//line /usr/local/go/src/path/filepath/path.go:23
// to hold the output until that output diverges from s.
//line /usr/local/go/src/path/filepath/path.go:27
type lazybuf struct {
	path		string
	buf		[]byte
	w		int
	volAndPath	string
	volLen		int
}

func (b *lazybuf) index(i int) byte {
//line /usr/local/go/src/path/filepath/path.go:35
	_go_fuzz_dep_.CoverTab[17957]++
							if b.buf != nil {
//line /usr/local/go/src/path/filepath/path.go:36
		_go_fuzz_dep_.CoverTab[17959]++
								return b.buf[i]
//line /usr/local/go/src/path/filepath/path.go:37
		// _ = "end of CoverTab[17959]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:38
		_go_fuzz_dep_.CoverTab[17960]++
//line /usr/local/go/src/path/filepath/path.go:38
		// _ = "end of CoverTab[17960]"
//line /usr/local/go/src/path/filepath/path.go:38
	}
//line /usr/local/go/src/path/filepath/path.go:38
	// _ = "end of CoverTab[17957]"
//line /usr/local/go/src/path/filepath/path.go:38
	_go_fuzz_dep_.CoverTab[17958]++
							return b.path[i]
//line /usr/local/go/src/path/filepath/path.go:39
	// _ = "end of CoverTab[17958]"
}

func (b *lazybuf) append(c byte) {
//line /usr/local/go/src/path/filepath/path.go:42
	_go_fuzz_dep_.CoverTab[17961]++
							if b.buf == nil {
//line /usr/local/go/src/path/filepath/path.go:43
		_go_fuzz_dep_.CoverTab[17963]++
								if b.w < len(b.path) && func() bool {
//line /usr/local/go/src/path/filepath/path.go:44
			_go_fuzz_dep_.CoverTab[17965]++
//line /usr/local/go/src/path/filepath/path.go:44
			return b.path[b.w] == c
//line /usr/local/go/src/path/filepath/path.go:44
			// _ = "end of CoverTab[17965]"
//line /usr/local/go/src/path/filepath/path.go:44
		}() {
//line /usr/local/go/src/path/filepath/path.go:44
			_go_fuzz_dep_.CoverTab[17966]++
									b.w++
									return
//line /usr/local/go/src/path/filepath/path.go:46
			// _ = "end of CoverTab[17966]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:47
			_go_fuzz_dep_.CoverTab[17967]++
//line /usr/local/go/src/path/filepath/path.go:47
			// _ = "end of CoverTab[17967]"
//line /usr/local/go/src/path/filepath/path.go:47
		}
//line /usr/local/go/src/path/filepath/path.go:47
		// _ = "end of CoverTab[17963]"
//line /usr/local/go/src/path/filepath/path.go:47
		_go_fuzz_dep_.CoverTab[17964]++
								b.buf = make([]byte, len(b.path))
								copy(b.buf, b.path[:b.w])
//line /usr/local/go/src/path/filepath/path.go:49
		// _ = "end of CoverTab[17964]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:50
		_go_fuzz_dep_.CoverTab[17968]++
//line /usr/local/go/src/path/filepath/path.go:50
		// _ = "end of CoverTab[17968]"
//line /usr/local/go/src/path/filepath/path.go:50
	}
//line /usr/local/go/src/path/filepath/path.go:50
	// _ = "end of CoverTab[17961]"
//line /usr/local/go/src/path/filepath/path.go:50
	_go_fuzz_dep_.CoverTab[17962]++
							b.buf[b.w] = c
							b.w++
//line /usr/local/go/src/path/filepath/path.go:52
	// _ = "end of CoverTab[17962]"
}

func (b *lazybuf) string() string {
//line /usr/local/go/src/path/filepath/path.go:55
	_go_fuzz_dep_.CoverTab[17969]++
							if b.buf == nil {
//line /usr/local/go/src/path/filepath/path.go:56
		_go_fuzz_dep_.CoverTab[17971]++
								return b.volAndPath[:b.volLen+b.w]
//line /usr/local/go/src/path/filepath/path.go:57
		// _ = "end of CoverTab[17971]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:58
		_go_fuzz_dep_.CoverTab[17972]++
//line /usr/local/go/src/path/filepath/path.go:58
		// _ = "end of CoverTab[17972]"
//line /usr/local/go/src/path/filepath/path.go:58
	}
//line /usr/local/go/src/path/filepath/path.go:58
	// _ = "end of CoverTab[17969]"
//line /usr/local/go/src/path/filepath/path.go:58
	_go_fuzz_dep_.CoverTab[17970]++
							return b.volAndPath[:b.volLen] + string(b.buf[:b.w])
//line /usr/local/go/src/path/filepath/path.go:59
	// _ = "end of CoverTab[17970]"
}

const (
	Separator	= os.PathSeparator
	ListSeparator	= os.PathListSeparator
)

// Clean returns the shortest path name equivalent to path
//line /usr/local/go/src/path/filepath/path.go:67
// by purely lexical processing. It applies the following rules
//line /usr/local/go/src/path/filepath/path.go:67
// iteratively until no further processing can be done:
//line /usr/local/go/src/path/filepath/path.go:67
//
//line /usr/local/go/src/path/filepath/path.go:67
//  1. Replace multiple Separator elements with a single one.
//line /usr/local/go/src/path/filepath/path.go:67
//  2. Eliminate each . path name element (the current directory).
//line /usr/local/go/src/path/filepath/path.go:67
//  3. Eliminate each inner .. path name element (the parent directory)
//line /usr/local/go/src/path/filepath/path.go:67
//     along with the non-.. element that precedes it.
//line /usr/local/go/src/path/filepath/path.go:67
//  4. Eliminate .. elements that begin a rooted path:
//line /usr/local/go/src/path/filepath/path.go:67
//     that is, replace "/.." by "/" at the beginning of a path,
//line /usr/local/go/src/path/filepath/path.go:67
//     assuming Separator is '/'.
//line /usr/local/go/src/path/filepath/path.go:67
//
//line /usr/local/go/src/path/filepath/path.go:67
// The returned path ends in a slash only if it represents a root directory,
//line /usr/local/go/src/path/filepath/path.go:67
// such as "/" on Unix or `C:\` on Windows.
//line /usr/local/go/src/path/filepath/path.go:67
//
//line /usr/local/go/src/path/filepath/path.go:67
// Finally, any occurrences of slash are replaced by Separator.
//line /usr/local/go/src/path/filepath/path.go:67
//
//line /usr/local/go/src/path/filepath/path.go:67
// If the result of this process is an empty string, Clean
//line /usr/local/go/src/path/filepath/path.go:67
// returns the string ".".
//line /usr/local/go/src/path/filepath/path.go:67
//
//line /usr/local/go/src/path/filepath/path.go:67
// See also Rob Pike, “Lexical File Names in Plan 9 or
//line /usr/local/go/src/path/filepath/path.go:67
// Getting Dot-Dot Right,”
//line /usr/local/go/src/path/filepath/path.go:67
// https://9p.io/sys/doc/lexnames.html
//line /usr/local/go/src/path/filepath/path.go:90
func Clean(path string) string {
//line /usr/local/go/src/path/filepath/path.go:90
	_go_fuzz_dep_.CoverTab[17973]++
							originalPath := path
							volLen := volumeNameLen(path)
							path = path[volLen:]
							if path == "" {
//line /usr/local/go/src/path/filepath/path.go:94
		_go_fuzz_dep_.CoverTab[17978]++
								if volLen > 1 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:95
			_go_fuzz_dep_.CoverTab[17980]++
//line /usr/local/go/src/path/filepath/path.go:95
			return os.IsPathSeparator(originalPath[0])
//line /usr/local/go/src/path/filepath/path.go:95
			// _ = "end of CoverTab[17980]"
//line /usr/local/go/src/path/filepath/path.go:95
		}() && func() bool {
//line /usr/local/go/src/path/filepath/path.go:95
			_go_fuzz_dep_.CoverTab[17981]++
//line /usr/local/go/src/path/filepath/path.go:95
			return os.IsPathSeparator(originalPath[1])
//line /usr/local/go/src/path/filepath/path.go:95
			// _ = "end of CoverTab[17981]"
//line /usr/local/go/src/path/filepath/path.go:95
		}() {
//line /usr/local/go/src/path/filepath/path.go:95
			_go_fuzz_dep_.CoverTab[17982]++

									return FromSlash(originalPath)
//line /usr/local/go/src/path/filepath/path.go:97
			// _ = "end of CoverTab[17982]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:98
			_go_fuzz_dep_.CoverTab[17983]++
//line /usr/local/go/src/path/filepath/path.go:98
			// _ = "end of CoverTab[17983]"
//line /usr/local/go/src/path/filepath/path.go:98
		}
//line /usr/local/go/src/path/filepath/path.go:98
		// _ = "end of CoverTab[17978]"
//line /usr/local/go/src/path/filepath/path.go:98
		_go_fuzz_dep_.CoverTab[17979]++
								return originalPath + "."
//line /usr/local/go/src/path/filepath/path.go:99
		// _ = "end of CoverTab[17979]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:100
		_go_fuzz_dep_.CoverTab[17984]++
//line /usr/local/go/src/path/filepath/path.go:100
		// _ = "end of CoverTab[17984]"
//line /usr/local/go/src/path/filepath/path.go:100
	}
//line /usr/local/go/src/path/filepath/path.go:100
	// _ = "end of CoverTab[17973]"
//line /usr/local/go/src/path/filepath/path.go:100
	_go_fuzz_dep_.CoverTab[17974]++
							rooted := os.IsPathSeparator(path[0])

//line /usr/local/go/src/path/filepath/path.go:108
	n := len(path)
	out := lazybuf{path: path, volAndPath: originalPath, volLen: volLen}
	r, dotdot := 0, 0
	if rooted {
//line /usr/local/go/src/path/filepath/path.go:111
		_go_fuzz_dep_.CoverTab[17985]++
								out.append(Separator)
								r, dotdot = 1, 1
//line /usr/local/go/src/path/filepath/path.go:113
		// _ = "end of CoverTab[17985]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:114
		_go_fuzz_dep_.CoverTab[17986]++
//line /usr/local/go/src/path/filepath/path.go:114
		// _ = "end of CoverTab[17986]"
//line /usr/local/go/src/path/filepath/path.go:114
	}
//line /usr/local/go/src/path/filepath/path.go:114
	// _ = "end of CoverTab[17974]"
//line /usr/local/go/src/path/filepath/path.go:114
	_go_fuzz_dep_.CoverTab[17975]++

							for r < n {
//line /usr/local/go/src/path/filepath/path.go:116
		_go_fuzz_dep_.CoverTab[17987]++
								switch {
		case os.IsPathSeparator(path[r]):
//line /usr/local/go/src/path/filepath/path.go:118
			_go_fuzz_dep_.CoverTab[17988]++

									r++
//line /usr/local/go/src/path/filepath/path.go:120
			// _ = "end of CoverTab[17988]"
		case path[r] == '.' && func() bool {
//line /usr/local/go/src/path/filepath/path.go:121
			_go_fuzz_dep_.CoverTab[17994]++
//line /usr/local/go/src/path/filepath/path.go:121
			return (r+1 == n || func() bool {
//line /usr/local/go/src/path/filepath/path.go:121
				_go_fuzz_dep_.CoverTab[17995]++
//line /usr/local/go/src/path/filepath/path.go:121
				return os.IsPathSeparator(path[r+1])
//line /usr/local/go/src/path/filepath/path.go:121
				// _ = "end of CoverTab[17995]"
//line /usr/local/go/src/path/filepath/path.go:121
			}())
//line /usr/local/go/src/path/filepath/path.go:121
			// _ = "end of CoverTab[17994]"
//line /usr/local/go/src/path/filepath/path.go:121
		}():
//line /usr/local/go/src/path/filepath/path.go:121
			_go_fuzz_dep_.CoverTab[17989]++

									r++
//line /usr/local/go/src/path/filepath/path.go:123
			// _ = "end of CoverTab[17989]"
		case path[r] == '.' && func() bool {
//line /usr/local/go/src/path/filepath/path.go:124
			_go_fuzz_dep_.CoverTab[17996]++
//line /usr/local/go/src/path/filepath/path.go:124
			return path[r+1] == '.'
//line /usr/local/go/src/path/filepath/path.go:124
			// _ = "end of CoverTab[17996]"
//line /usr/local/go/src/path/filepath/path.go:124
		}() && func() bool {
//line /usr/local/go/src/path/filepath/path.go:124
			_go_fuzz_dep_.CoverTab[17997]++
//line /usr/local/go/src/path/filepath/path.go:124
			return (r+2 == n || func() bool {
//line /usr/local/go/src/path/filepath/path.go:124
				_go_fuzz_dep_.CoverTab[17998]++
//line /usr/local/go/src/path/filepath/path.go:124
				return os.IsPathSeparator(path[r+2])
//line /usr/local/go/src/path/filepath/path.go:124
				// _ = "end of CoverTab[17998]"
//line /usr/local/go/src/path/filepath/path.go:124
			}())
//line /usr/local/go/src/path/filepath/path.go:124
			// _ = "end of CoverTab[17997]"
//line /usr/local/go/src/path/filepath/path.go:124
		}():
//line /usr/local/go/src/path/filepath/path.go:124
			_go_fuzz_dep_.CoverTab[17990]++

									r += 2
									switch {
			case out.w > dotdot:
//line /usr/local/go/src/path/filepath/path.go:128
				_go_fuzz_dep_.CoverTab[17999]++

										out.w--
										for out.w > dotdot && func() bool {
//line /usr/local/go/src/path/filepath/path.go:131
					_go_fuzz_dep_.CoverTab[18003]++
//line /usr/local/go/src/path/filepath/path.go:131
					return !os.IsPathSeparator(out.index(out.w))
//line /usr/local/go/src/path/filepath/path.go:131
					// _ = "end of CoverTab[18003]"
//line /usr/local/go/src/path/filepath/path.go:131
				}() {
//line /usr/local/go/src/path/filepath/path.go:131
					_go_fuzz_dep_.CoverTab[18004]++
											out.w--
//line /usr/local/go/src/path/filepath/path.go:132
					// _ = "end of CoverTab[18004]"
				}
//line /usr/local/go/src/path/filepath/path.go:133
				// _ = "end of CoverTab[17999]"
			case !rooted:
//line /usr/local/go/src/path/filepath/path.go:134
				_go_fuzz_dep_.CoverTab[18000]++

										if out.w > 0 {
//line /usr/local/go/src/path/filepath/path.go:136
					_go_fuzz_dep_.CoverTab[18005]++
											out.append(Separator)
//line /usr/local/go/src/path/filepath/path.go:137
					// _ = "end of CoverTab[18005]"
				} else {
//line /usr/local/go/src/path/filepath/path.go:138
					_go_fuzz_dep_.CoverTab[18006]++
//line /usr/local/go/src/path/filepath/path.go:138
					// _ = "end of CoverTab[18006]"
//line /usr/local/go/src/path/filepath/path.go:138
				}
//line /usr/local/go/src/path/filepath/path.go:138
				// _ = "end of CoverTab[18000]"
//line /usr/local/go/src/path/filepath/path.go:138
				_go_fuzz_dep_.CoverTab[18001]++
										out.append('.')
										out.append('.')
										dotdot = out.w
//line /usr/local/go/src/path/filepath/path.go:141
				// _ = "end of CoverTab[18001]"
//line /usr/local/go/src/path/filepath/path.go:141
			default:
//line /usr/local/go/src/path/filepath/path.go:141
				_go_fuzz_dep_.CoverTab[18002]++
//line /usr/local/go/src/path/filepath/path.go:141
				// _ = "end of CoverTab[18002]"
			}
//line /usr/local/go/src/path/filepath/path.go:142
			// _ = "end of CoverTab[17990]"
		default:
//line /usr/local/go/src/path/filepath/path.go:143
			_go_fuzz_dep_.CoverTab[17991]++

//line /usr/local/go/src/path/filepath/path.go:146
			if rooted && func() bool {
//line /usr/local/go/src/path/filepath/path.go:146
				_go_fuzz_dep_.CoverTab[18007]++
//line /usr/local/go/src/path/filepath/path.go:146
				return out.w != 1
//line /usr/local/go/src/path/filepath/path.go:146
				// _ = "end of CoverTab[18007]"
//line /usr/local/go/src/path/filepath/path.go:146
			}() || func() bool {
//line /usr/local/go/src/path/filepath/path.go:146
				_go_fuzz_dep_.CoverTab[18008]++
//line /usr/local/go/src/path/filepath/path.go:146
				return !rooted && func() bool {
//line /usr/local/go/src/path/filepath/path.go:146
					_go_fuzz_dep_.CoverTab[18009]++
//line /usr/local/go/src/path/filepath/path.go:146
					return out.w != 0
//line /usr/local/go/src/path/filepath/path.go:146
					// _ = "end of CoverTab[18009]"
//line /usr/local/go/src/path/filepath/path.go:146
				}()
//line /usr/local/go/src/path/filepath/path.go:146
				// _ = "end of CoverTab[18008]"
//line /usr/local/go/src/path/filepath/path.go:146
			}() {
//line /usr/local/go/src/path/filepath/path.go:146
				_go_fuzz_dep_.CoverTab[18010]++
										out.append(Separator)
//line /usr/local/go/src/path/filepath/path.go:147
				// _ = "end of CoverTab[18010]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:148
				_go_fuzz_dep_.CoverTab[18011]++
//line /usr/local/go/src/path/filepath/path.go:148
				// _ = "end of CoverTab[18011]"
//line /usr/local/go/src/path/filepath/path.go:148
			}
//line /usr/local/go/src/path/filepath/path.go:148
			// _ = "end of CoverTab[17991]"
//line /usr/local/go/src/path/filepath/path.go:148
			_go_fuzz_dep_.CoverTab[17992]++

//line /usr/local/go/src/path/filepath/path.go:152
			if runtime.GOOS == "windows" && func() bool {
//line /usr/local/go/src/path/filepath/path.go:152
				_go_fuzz_dep_.CoverTab[18012]++
//line /usr/local/go/src/path/filepath/path.go:152
				return out.w == 0
//line /usr/local/go/src/path/filepath/path.go:152
				// _ = "end of CoverTab[18012]"
//line /usr/local/go/src/path/filepath/path.go:152
			}() && func() bool {
//line /usr/local/go/src/path/filepath/path.go:152
				_go_fuzz_dep_.CoverTab[18013]++
//line /usr/local/go/src/path/filepath/path.go:152
				return out.volLen == 0
//line /usr/local/go/src/path/filepath/path.go:152
				// _ = "end of CoverTab[18013]"
//line /usr/local/go/src/path/filepath/path.go:152
			}() && func() bool {
//line /usr/local/go/src/path/filepath/path.go:152
				_go_fuzz_dep_.CoverTab[18014]++
//line /usr/local/go/src/path/filepath/path.go:152
				return r != 0
//line /usr/local/go/src/path/filepath/path.go:152
				// _ = "end of CoverTab[18014]"
//line /usr/local/go/src/path/filepath/path.go:152
			}() {
//line /usr/local/go/src/path/filepath/path.go:152
				_go_fuzz_dep_.CoverTab[18015]++
										for i := r; i < n && func() bool {
//line /usr/local/go/src/path/filepath/path.go:153
					_go_fuzz_dep_.CoverTab[18016]++
//line /usr/local/go/src/path/filepath/path.go:153
					return !os.IsPathSeparator(path[i])
//line /usr/local/go/src/path/filepath/path.go:153
					// _ = "end of CoverTab[18016]"
//line /usr/local/go/src/path/filepath/path.go:153
				}(); i++ {
//line /usr/local/go/src/path/filepath/path.go:153
					_go_fuzz_dep_.CoverTab[18017]++
											if path[i] == ':' {
//line /usr/local/go/src/path/filepath/path.go:154
						_go_fuzz_dep_.CoverTab[18018]++
												out.append('.')
												out.append(Separator)
												break
//line /usr/local/go/src/path/filepath/path.go:157
						// _ = "end of CoverTab[18018]"
					} else {
//line /usr/local/go/src/path/filepath/path.go:158
						_go_fuzz_dep_.CoverTab[18019]++
//line /usr/local/go/src/path/filepath/path.go:158
						// _ = "end of CoverTab[18019]"
//line /usr/local/go/src/path/filepath/path.go:158
					}
//line /usr/local/go/src/path/filepath/path.go:158
					// _ = "end of CoverTab[18017]"
				}
//line /usr/local/go/src/path/filepath/path.go:159
				// _ = "end of CoverTab[18015]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:160
				_go_fuzz_dep_.CoverTab[18020]++
//line /usr/local/go/src/path/filepath/path.go:160
				// _ = "end of CoverTab[18020]"
//line /usr/local/go/src/path/filepath/path.go:160
			}
//line /usr/local/go/src/path/filepath/path.go:160
			// _ = "end of CoverTab[17992]"
//line /usr/local/go/src/path/filepath/path.go:160
			_go_fuzz_dep_.CoverTab[17993]++

									for ; r < n && func() bool {
//line /usr/local/go/src/path/filepath/path.go:162
				_go_fuzz_dep_.CoverTab[18021]++
//line /usr/local/go/src/path/filepath/path.go:162
				return !os.IsPathSeparator(path[r])
//line /usr/local/go/src/path/filepath/path.go:162
				// _ = "end of CoverTab[18021]"
//line /usr/local/go/src/path/filepath/path.go:162
			}(); r++ {
//line /usr/local/go/src/path/filepath/path.go:162
				_go_fuzz_dep_.CoverTab[18022]++
										out.append(path[r])
//line /usr/local/go/src/path/filepath/path.go:163
				// _ = "end of CoverTab[18022]"
			}
//line /usr/local/go/src/path/filepath/path.go:164
			// _ = "end of CoverTab[17993]"
		}
//line /usr/local/go/src/path/filepath/path.go:165
		// _ = "end of CoverTab[17987]"
	}
//line /usr/local/go/src/path/filepath/path.go:166
	// _ = "end of CoverTab[17975]"
//line /usr/local/go/src/path/filepath/path.go:166
	_go_fuzz_dep_.CoverTab[17976]++

//line /usr/local/go/src/path/filepath/path.go:169
	if out.w == 0 {
//line /usr/local/go/src/path/filepath/path.go:169
		_go_fuzz_dep_.CoverTab[18023]++
								out.append('.')
//line /usr/local/go/src/path/filepath/path.go:170
		// _ = "end of CoverTab[18023]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:171
		_go_fuzz_dep_.CoverTab[18024]++
//line /usr/local/go/src/path/filepath/path.go:171
		// _ = "end of CoverTab[18024]"
//line /usr/local/go/src/path/filepath/path.go:171
	}
//line /usr/local/go/src/path/filepath/path.go:171
	// _ = "end of CoverTab[17976]"
//line /usr/local/go/src/path/filepath/path.go:171
	_go_fuzz_dep_.CoverTab[17977]++

							return FromSlash(out.string())
//line /usr/local/go/src/path/filepath/path.go:173
	// _ = "end of CoverTab[17977]"
}

// IsLocal reports whether path, using lexical analysis only, has all of these properties:
//line /usr/local/go/src/path/filepath/path.go:176
//
//line /usr/local/go/src/path/filepath/path.go:176
//   - is within the subtree rooted at the directory in which path is evaluated
//line /usr/local/go/src/path/filepath/path.go:176
//   - is not an absolute path
//line /usr/local/go/src/path/filepath/path.go:176
//   - is not empty
//line /usr/local/go/src/path/filepath/path.go:176
//   - on Windows, is not a reserved name such as "NUL"
//line /usr/local/go/src/path/filepath/path.go:176
//
//line /usr/local/go/src/path/filepath/path.go:176
// If IsLocal(path) returns true, then
//line /usr/local/go/src/path/filepath/path.go:176
// Join(base, path) will always produce a path contained within base and
//line /usr/local/go/src/path/filepath/path.go:176
// Clean(path) will always produce an unrooted path with no ".." path elements.
//line /usr/local/go/src/path/filepath/path.go:176
//
//line /usr/local/go/src/path/filepath/path.go:176
// IsLocal is a purely lexical operation.
//line /usr/local/go/src/path/filepath/path.go:176
// In particular, it does not account for the effect of any symbolic links
//line /usr/local/go/src/path/filepath/path.go:176
// that may exist in the filesystem.
//line /usr/local/go/src/path/filepath/path.go:190
func IsLocal(path string) bool {
//line /usr/local/go/src/path/filepath/path.go:190
	_go_fuzz_dep_.CoverTab[18025]++
							return isLocal(path)
//line /usr/local/go/src/path/filepath/path.go:191
	// _ = "end of CoverTab[18025]"
}

func unixIsLocal(path string) bool {
//line /usr/local/go/src/path/filepath/path.go:194
	_go_fuzz_dep_.CoverTab[18026]++
							if IsAbs(path) || func() bool {
//line /usr/local/go/src/path/filepath/path.go:195
		_go_fuzz_dep_.CoverTab[18031]++
//line /usr/local/go/src/path/filepath/path.go:195
		return path == ""
//line /usr/local/go/src/path/filepath/path.go:195
		// _ = "end of CoverTab[18031]"
//line /usr/local/go/src/path/filepath/path.go:195
	}() {
//line /usr/local/go/src/path/filepath/path.go:195
		_go_fuzz_dep_.CoverTab[18032]++
								return false
//line /usr/local/go/src/path/filepath/path.go:196
		// _ = "end of CoverTab[18032]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:197
		_go_fuzz_dep_.CoverTab[18033]++
//line /usr/local/go/src/path/filepath/path.go:197
		// _ = "end of CoverTab[18033]"
//line /usr/local/go/src/path/filepath/path.go:197
	}
//line /usr/local/go/src/path/filepath/path.go:197
	// _ = "end of CoverTab[18026]"
//line /usr/local/go/src/path/filepath/path.go:197
	_go_fuzz_dep_.CoverTab[18027]++
							hasDots := false
							for p := path; p != ""; {
//line /usr/local/go/src/path/filepath/path.go:199
		_go_fuzz_dep_.CoverTab[18034]++
								var part string
								part, p, _ = strings.Cut(p, "/")
								if part == "." || func() bool {
//line /usr/local/go/src/path/filepath/path.go:202
			_go_fuzz_dep_.CoverTab[18035]++
//line /usr/local/go/src/path/filepath/path.go:202
			return part == ".."
//line /usr/local/go/src/path/filepath/path.go:202
			// _ = "end of CoverTab[18035]"
//line /usr/local/go/src/path/filepath/path.go:202
		}() {
//line /usr/local/go/src/path/filepath/path.go:202
			_go_fuzz_dep_.CoverTab[18036]++
									hasDots = true
									break
//line /usr/local/go/src/path/filepath/path.go:204
			// _ = "end of CoverTab[18036]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:205
			_go_fuzz_dep_.CoverTab[18037]++
//line /usr/local/go/src/path/filepath/path.go:205
			// _ = "end of CoverTab[18037]"
//line /usr/local/go/src/path/filepath/path.go:205
		}
//line /usr/local/go/src/path/filepath/path.go:205
		// _ = "end of CoverTab[18034]"
	}
//line /usr/local/go/src/path/filepath/path.go:206
	// _ = "end of CoverTab[18027]"
//line /usr/local/go/src/path/filepath/path.go:206
	_go_fuzz_dep_.CoverTab[18028]++
							if hasDots {
//line /usr/local/go/src/path/filepath/path.go:207
		_go_fuzz_dep_.CoverTab[18038]++
								path = Clean(path)
//line /usr/local/go/src/path/filepath/path.go:208
		// _ = "end of CoverTab[18038]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:209
		_go_fuzz_dep_.CoverTab[18039]++
//line /usr/local/go/src/path/filepath/path.go:209
		// _ = "end of CoverTab[18039]"
//line /usr/local/go/src/path/filepath/path.go:209
	}
//line /usr/local/go/src/path/filepath/path.go:209
	// _ = "end of CoverTab[18028]"
//line /usr/local/go/src/path/filepath/path.go:209
	_go_fuzz_dep_.CoverTab[18029]++
							if path == ".." || func() bool {
//line /usr/local/go/src/path/filepath/path.go:210
		_go_fuzz_dep_.CoverTab[18040]++
//line /usr/local/go/src/path/filepath/path.go:210
		return strings.HasPrefix(path, "../")
//line /usr/local/go/src/path/filepath/path.go:210
		// _ = "end of CoverTab[18040]"
//line /usr/local/go/src/path/filepath/path.go:210
	}() {
//line /usr/local/go/src/path/filepath/path.go:210
		_go_fuzz_dep_.CoverTab[18041]++
								return false
//line /usr/local/go/src/path/filepath/path.go:211
		// _ = "end of CoverTab[18041]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:212
		_go_fuzz_dep_.CoverTab[18042]++
//line /usr/local/go/src/path/filepath/path.go:212
		// _ = "end of CoverTab[18042]"
//line /usr/local/go/src/path/filepath/path.go:212
	}
//line /usr/local/go/src/path/filepath/path.go:212
	// _ = "end of CoverTab[18029]"
//line /usr/local/go/src/path/filepath/path.go:212
	_go_fuzz_dep_.CoverTab[18030]++
							return true
//line /usr/local/go/src/path/filepath/path.go:213
	// _ = "end of CoverTab[18030]"
}

// ToSlash returns the result of replacing each separator character
//line /usr/local/go/src/path/filepath/path.go:216
// in path with a slash ('/') character. Multiple separators are
//line /usr/local/go/src/path/filepath/path.go:216
// replaced by multiple slashes.
//line /usr/local/go/src/path/filepath/path.go:219
func ToSlash(path string) string {
//line /usr/local/go/src/path/filepath/path.go:219
	_go_fuzz_dep_.CoverTab[18043]++
							if Separator == '/' {
//line /usr/local/go/src/path/filepath/path.go:220
		_go_fuzz_dep_.CoverTab[18045]++
								return path
//line /usr/local/go/src/path/filepath/path.go:221
		// _ = "end of CoverTab[18045]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:222
		_go_fuzz_dep_.CoverTab[18046]++
//line /usr/local/go/src/path/filepath/path.go:222
		// _ = "end of CoverTab[18046]"
//line /usr/local/go/src/path/filepath/path.go:222
	}
//line /usr/local/go/src/path/filepath/path.go:222
	// _ = "end of CoverTab[18043]"
//line /usr/local/go/src/path/filepath/path.go:222
	_go_fuzz_dep_.CoverTab[18044]++
							return strings.ReplaceAll(path, string(Separator), "/")
//line /usr/local/go/src/path/filepath/path.go:223
	// _ = "end of CoverTab[18044]"
}

// FromSlash returns the result of replacing each slash ('/') character
//line /usr/local/go/src/path/filepath/path.go:226
// in path with a separator character. Multiple slashes are replaced
//line /usr/local/go/src/path/filepath/path.go:226
// by multiple separators.
//line /usr/local/go/src/path/filepath/path.go:229
func FromSlash(path string) string {
//line /usr/local/go/src/path/filepath/path.go:229
	_go_fuzz_dep_.CoverTab[18047]++
							if Separator == '/' {
//line /usr/local/go/src/path/filepath/path.go:230
		_go_fuzz_dep_.CoverTab[18049]++
								return path
//line /usr/local/go/src/path/filepath/path.go:231
		// _ = "end of CoverTab[18049]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:232
		_go_fuzz_dep_.CoverTab[18050]++
//line /usr/local/go/src/path/filepath/path.go:232
		// _ = "end of CoverTab[18050]"
//line /usr/local/go/src/path/filepath/path.go:232
	}
//line /usr/local/go/src/path/filepath/path.go:232
	// _ = "end of CoverTab[18047]"
//line /usr/local/go/src/path/filepath/path.go:232
	_go_fuzz_dep_.CoverTab[18048]++
							return strings.ReplaceAll(path, "/", string(Separator))
//line /usr/local/go/src/path/filepath/path.go:233
	// _ = "end of CoverTab[18048]"
}

// SplitList splits a list of paths joined by the OS-specific ListSeparator,
//line /usr/local/go/src/path/filepath/path.go:236
// usually found in PATH or GOPATH environment variables.
//line /usr/local/go/src/path/filepath/path.go:236
// Unlike strings.Split, SplitList returns an empty slice when passed an empty
//line /usr/local/go/src/path/filepath/path.go:236
// string.
//line /usr/local/go/src/path/filepath/path.go:240
func SplitList(path string) []string {
//line /usr/local/go/src/path/filepath/path.go:240
	_go_fuzz_dep_.CoverTab[18051]++
							return splitList(path)
//line /usr/local/go/src/path/filepath/path.go:241
	// _ = "end of CoverTab[18051]"
}

// Split splits path immediately following the final Separator,
//line /usr/local/go/src/path/filepath/path.go:244
// separating it into a directory and file name component.
//line /usr/local/go/src/path/filepath/path.go:244
// If there is no Separator in path, Split returns an empty dir
//line /usr/local/go/src/path/filepath/path.go:244
// and file set to path.
//line /usr/local/go/src/path/filepath/path.go:244
// The returned values have the property that path = dir+file.
//line /usr/local/go/src/path/filepath/path.go:249
func Split(path string) (dir, file string) {
//line /usr/local/go/src/path/filepath/path.go:249
	_go_fuzz_dep_.CoverTab[18052]++
							vol := VolumeName(path)
							i := len(path) - 1
							for i >= len(vol) && func() bool {
//line /usr/local/go/src/path/filepath/path.go:252
		_go_fuzz_dep_.CoverTab[18054]++
//line /usr/local/go/src/path/filepath/path.go:252
		return !os.IsPathSeparator(path[i])
//line /usr/local/go/src/path/filepath/path.go:252
		// _ = "end of CoverTab[18054]"
//line /usr/local/go/src/path/filepath/path.go:252
	}() {
//line /usr/local/go/src/path/filepath/path.go:252
		_go_fuzz_dep_.CoverTab[18055]++
								i--
//line /usr/local/go/src/path/filepath/path.go:253
		// _ = "end of CoverTab[18055]"
	}
//line /usr/local/go/src/path/filepath/path.go:254
	// _ = "end of CoverTab[18052]"
//line /usr/local/go/src/path/filepath/path.go:254
	_go_fuzz_dep_.CoverTab[18053]++
							return path[:i+1], path[i+1:]
//line /usr/local/go/src/path/filepath/path.go:255
	// _ = "end of CoverTab[18053]"
}

// Join joins any number of path elements into a single path,
//line /usr/local/go/src/path/filepath/path.go:258
// separating them with an OS specific Separator. Empty elements
//line /usr/local/go/src/path/filepath/path.go:258
// are ignored. The result is Cleaned. However, if the argument
//line /usr/local/go/src/path/filepath/path.go:258
// list is empty or all its elements are empty, Join returns
//line /usr/local/go/src/path/filepath/path.go:258
// an empty string.
//line /usr/local/go/src/path/filepath/path.go:258
// On Windows, the result will only be a UNC path if the first
//line /usr/local/go/src/path/filepath/path.go:258
// non-empty element is a UNC path.
//line /usr/local/go/src/path/filepath/path.go:265
func Join(elem ...string) string {
//line /usr/local/go/src/path/filepath/path.go:265
	_go_fuzz_dep_.CoverTab[18056]++
							return join(elem)
//line /usr/local/go/src/path/filepath/path.go:266
	// _ = "end of CoverTab[18056]"
}

// Ext returns the file name extension used by path.
//line /usr/local/go/src/path/filepath/path.go:269
// The extension is the suffix beginning at the final dot
//line /usr/local/go/src/path/filepath/path.go:269
// in the final element of path; it is empty if there is
//line /usr/local/go/src/path/filepath/path.go:269
// no dot.
//line /usr/local/go/src/path/filepath/path.go:273
func Ext(path string) string {
//line /usr/local/go/src/path/filepath/path.go:273
	_go_fuzz_dep_.CoverTab[18057]++
							for i := len(path) - 1; i >= 0 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:274
		_go_fuzz_dep_.CoverTab[18059]++
//line /usr/local/go/src/path/filepath/path.go:274
		return !os.IsPathSeparator(path[i])
//line /usr/local/go/src/path/filepath/path.go:274
		// _ = "end of CoverTab[18059]"
//line /usr/local/go/src/path/filepath/path.go:274
	}(); i-- {
//line /usr/local/go/src/path/filepath/path.go:274
		_go_fuzz_dep_.CoverTab[18060]++
								if path[i] == '.' {
//line /usr/local/go/src/path/filepath/path.go:275
			_go_fuzz_dep_.CoverTab[18061]++
									return path[i:]
//line /usr/local/go/src/path/filepath/path.go:276
			// _ = "end of CoverTab[18061]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:277
			_go_fuzz_dep_.CoverTab[18062]++
//line /usr/local/go/src/path/filepath/path.go:277
			// _ = "end of CoverTab[18062]"
//line /usr/local/go/src/path/filepath/path.go:277
		}
//line /usr/local/go/src/path/filepath/path.go:277
		// _ = "end of CoverTab[18060]"
	}
//line /usr/local/go/src/path/filepath/path.go:278
	// _ = "end of CoverTab[18057]"
//line /usr/local/go/src/path/filepath/path.go:278
	_go_fuzz_dep_.CoverTab[18058]++
							return ""
//line /usr/local/go/src/path/filepath/path.go:279
	// _ = "end of CoverTab[18058]"
}

// EvalSymlinks returns the path name after the evaluation of any symbolic
//line /usr/local/go/src/path/filepath/path.go:282
// links.
//line /usr/local/go/src/path/filepath/path.go:282
// If path is relative the result will be relative to the current directory,
//line /usr/local/go/src/path/filepath/path.go:282
// unless one of the components is an absolute symbolic link.
//line /usr/local/go/src/path/filepath/path.go:282
// EvalSymlinks calls Clean on the result.
//line /usr/local/go/src/path/filepath/path.go:287
func EvalSymlinks(path string) (string, error) {
//line /usr/local/go/src/path/filepath/path.go:287
	_go_fuzz_dep_.CoverTab[18063]++
							return evalSymlinks(path)
//line /usr/local/go/src/path/filepath/path.go:288
	// _ = "end of CoverTab[18063]"
}

// Abs returns an absolute representation of path.
//line /usr/local/go/src/path/filepath/path.go:291
// If the path is not absolute it will be joined with the current
//line /usr/local/go/src/path/filepath/path.go:291
// working directory to turn it into an absolute path. The absolute
//line /usr/local/go/src/path/filepath/path.go:291
// path name for a given file is not guaranteed to be unique.
//line /usr/local/go/src/path/filepath/path.go:291
// Abs calls Clean on the result.
//line /usr/local/go/src/path/filepath/path.go:296
func Abs(path string) (string, error) {
//line /usr/local/go/src/path/filepath/path.go:296
	_go_fuzz_dep_.CoverTab[18064]++
							return abs(path)
//line /usr/local/go/src/path/filepath/path.go:297
	// _ = "end of CoverTab[18064]"
}

func unixAbs(path string) (string, error) {
//line /usr/local/go/src/path/filepath/path.go:300
	_go_fuzz_dep_.CoverTab[18065]++
							if IsAbs(path) {
//line /usr/local/go/src/path/filepath/path.go:301
		_go_fuzz_dep_.CoverTab[18068]++
								return Clean(path), nil
//line /usr/local/go/src/path/filepath/path.go:302
		// _ = "end of CoverTab[18068]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:303
		_go_fuzz_dep_.CoverTab[18069]++
//line /usr/local/go/src/path/filepath/path.go:303
		// _ = "end of CoverTab[18069]"
//line /usr/local/go/src/path/filepath/path.go:303
	}
//line /usr/local/go/src/path/filepath/path.go:303
	// _ = "end of CoverTab[18065]"
//line /usr/local/go/src/path/filepath/path.go:303
	_go_fuzz_dep_.CoverTab[18066]++
							wd, err := os.Getwd()
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:305
		_go_fuzz_dep_.CoverTab[18070]++
								return "", err
//line /usr/local/go/src/path/filepath/path.go:306
		// _ = "end of CoverTab[18070]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:307
		_go_fuzz_dep_.CoverTab[18071]++
//line /usr/local/go/src/path/filepath/path.go:307
		// _ = "end of CoverTab[18071]"
//line /usr/local/go/src/path/filepath/path.go:307
	}
//line /usr/local/go/src/path/filepath/path.go:307
	// _ = "end of CoverTab[18066]"
//line /usr/local/go/src/path/filepath/path.go:307
	_go_fuzz_dep_.CoverTab[18067]++
							return Join(wd, path), nil
//line /usr/local/go/src/path/filepath/path.go:308
	// _ = "end of CoverTab[18067]"
}

// Rel returns a relative path that is lexically equivalent to targpath when
//line /usr/local/go/src/path/filepath/path.go:311
// joined to basepath with an intervening separator. That is,
//line /usr/local/go/src/path/filepath/path.go:311
// Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
//line /usr/local/go/src/path/filepath/path.go:311
// On success, the returned path will always be relative to basepath,
//line /usr/local/go/src/path/filepath/path.go:311
// even if basepath and targpath share no elements.
//line /usr/local/go/src/path/filepath/path.go:311
// An error is returned if targpath can't be made relative to basepath or if
//line /usr/local/go/src/path/filepath/path.go:311
// knowing the current working directory would be necessary to compute it.
//line /usr/local/go/src/path/filepath/path.go:311
// Rel calls Clean on the result.
//line /usr/local/go/src/path/filepath/path.go:319
func Rel(basepath, targpath string) (string, error) {
//line /usr/local/go/src/path/filepath/path.go:319
	_go_fuzz_dep_.CoverTab[18072]++
							baseVol := VolumeName(basepath)
							targVol := VolumeName(targpath)
							base := Clean(basepath)
							targ := Clean(targpath)
							if sameWord(targ, base) {
//line /usr/local/go/src/path/filepath/path.go:324
		_go_fuzz_dep_.CoverTab[18079]++
								return ".", nil
//line /usr/local/go/src/path/filepath/path.go:325
		// _ = "end of CoverTab[18079]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:326
		_go_fuzz_dep_.CoverTab[18080]++
//line /usr/local/go/src/path/filepath/path.go:326
		// _ = "end of CoverTab[18080]"
//line /usr/local/go/src/path/filepath/path.go:326
	}
//line /usr/local/go/src/path/filepath/path.go:326
	// _ = "end of CoverTab[18072]"
//line /usr/local/go/src/path/filepath/path.go:326
	_go_fuzz_dep_.CoverTab[18073]++
							base = base[len(baseVol):]
							targ = targ[len(targVol):]
							if base == "." {
//line /usr/local/go/src/path/filepath/path.go:329
		_go_fuzz_dep_.CoverTab[18081]++
								base = ""
//line /usr/local/go/src/path/filepath/path.go:330
		// _ = "end of CoverTab[18081]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:331
		_go_fuzz_dep_.CoverTab[18082]++
//line /usr/local/go/src/path/filepath/path.go:331
		if base == "" && func() bool {
//line /usr/local/go/src/path/filepath/path.go:331
			_go_fuzz_dep_.CoverTab[18083]++
//line /usr/local/go/src/path/filepath/path.go:331
			return volumeNameLen(baseVol) > 2
//line /usr/local/go/src/path/filepath/path.go:331
			// _ = "end of CoverTab[18083]"
//line /usr/local/go/src/path/filepath/path.go:331
		}() {
//line /usr/local/go/src/path/filepath/path.go:331
			_go_fuzz_dep_.CoverTab[18084]++

									base = string(Separator)
//line /usr/local/go/src/path/filepath/path.go:333
			// _ = "end of CoverTab[18084]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:334
			_go_fuzz_dep_.CoverTab[18085]++
//line /usr/local/go/src/path/filepath/path.go:334
			// _ = "end of CoverTab[18085]"
//line /usr/local/go/src/path/filepath/path.go:334
		}
//line /usr/local/go/src/path/filepath/path.go:334
		// _ = "end of CoverTab[18082]"
//line /usr/local/go/src/path/filepath/path.go:334
	}
//line /usr/local/go/src/path/filepath/path.go:334
	// _ = "end of CoverTab[18073]"
//line /usr/local/go/src/path/filepath/path.go:334
	_go_fuzz_dep_.CoverTab[18074]++

//line /usr/local/go/src/path/filepath/path.go:337
	baseSlashed := len(base) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:337
		_go_fuzz_dep_.CoverTab[18086]++
//line /usr/local/go/src/path/filepath/path.go:337
		return base[0] == Separator
//line /usr/local/go/src/path/filepath/path.go:337
		// _ = "end of CoverTab[18086]"
//line /usr/local/go/src/path/filepath/path.go:337
	}()
	targSlashed := len(targ) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:338
		_go_fuzz_dep_.CoverTab[18087]++
//line /usr/local/go/src/path/filepath/path.go:338
		return targ[0] == Separator
//line /usr/local/go/src/path/filepath/path.go:338
		// _ = "end of CoverTab[18087]"
//line /usr/local/go/src/path/filepath/path.go:338
	}()
							if baseSlashed != targSlashed || func() bool {
//line /usr/local/go/src/path/filepath/path.go:339
		_go_fuzz_dep_.CoverTab[18088]++
//line /usr/local/go/src/path/filepath/path.go:339
		return !sameWord(baseVol, targVol)
//line /usr/local/go/src/path/filepath/path.go:339
		// _ = "end of CoverTab[18088]"
//line /usr/local/go/src/path/filepath/path.go:339
	}() {
//line /usr/local/go/src/path/filepath/path.go:339
		_go_fuzz_dep_.CoverTab[18089]++
								return "", errors.New("Rel: can't make " + targpath + " relative to " + basepath)
//line /usr/local/go/src/path/filepath/path.go:340
		// _ = "end of CoverTab[18089]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:341
		_go_fuzz_dep_.CoverTab[18090]++
//line /usr/local/go/src/path/filepath/path.go:341
		// _ = "end of CoverTab[18090]"
//line /usr/local/go/src/path/filepath/path.go:341
	}
//line /usr/local/go/src/path/filepath/path.go:341
	// _ = "end of CoverTab[18074]"
//line /usr/local/go/src/path/filepath/path.go:341
	_go_fuzz_dep_.CoverTab[18075]++

							bl := len(base)
							tl := len(targ)
							var b0, bi, t0, ti int
							for {
//line /usr/local/go/src/path/filepath/path.go:346
		_go_fuzz_dep_.CoverTab[18091]++
								for bi < bl && func() bool {
//line /usr/local/go/src/path/filepath/path.go:347
			_go_fuzz_dep_.CoverTab[18097]++
//line /usr/local/go/src/path/filepath/path.go:347
			return base[bi] != Separator
//line /usr/local/go/src/path/filepath/path.go:347
			// _ = "end of CoverTab[18097]"
//line /usr/local/go/src/path/filepath/path.go:347
		}() {
//line /usr/local/go/src/path/filepath/path.go:347
			_go_fuzz_dep_.CoverTab[18098]++
									bi++
//line /usr/local/go/src/path/filepath/path.go:348
			// _ = "end of CoverTab[18098]"
		}
//line /usr/local/go/src/path/filepath/path.go:349
		// _ = "end of CoverTab[18091]"
//line /usr/local/go/src/path/filepath/path.go:349
		_go_fuzz_dep_.CoverTab[18092]++
								for ti < tl && func() bool {
//line /usr/local/go/src/path/filepath/path.go:350
			_go_fuzz_dep_.CoverTab[18099]++
//line /usr/local/go/src/path/filepath/path.go:350
			return targ[ti] != Separator
//line /usr/local/go/src/path/filepath/path.go:350
			// _ = "end of CoverTab[18099]"
//line /usr/local/go/src/path/filepath/path.go:350
		}() {
//line /usr/local/go/src/path/filepath/path.go:350
			_go_fuzz_dep_.CoverTab[18100]++
									ti++
//line /usr/local/go/src/path/filepath/path.go:351
			// _ = "end of CoverTab[18100]"
		}
//line /usr/local/go/src/path/filepath/path.go:352
		// _ = "end of CoverTab[18092]"
//line /usr/local/go/src/path/filepath/path.go:352
		_go_fuzz_dep_.CoverTab[18093]++
								if !sameWord(targ[t0:ti], base[b0:bi]) {
//line /usr/local/go/src/path/filepath/path.go:353
			_go_fuzz_dep_.CoverTab[18101]++
									break
//line /usr/local/go/src/path/filepath/path.go:354
			// _ = "end of CoverTab[18101]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:355
			_go_fuzz_dep_.CoverTab[18102]++
//line /usr/local/go/src/path/filepath/path.go:355
			// _ = "end of CoverTab[18102]"
//line /usr/local/go/src/path/filepath/path.go:355
		}
//line /usr/local/go/src/path/filepath/path.go:355
		// _ = "end of CoverTab[18093]"
//line /usr/local/go/src/path/filepath/path.go:355
		_go_fuzz_dep_.CoverTab[18094]++
								if bi < bl {
//line /usr/local/go/src/path/filepath/path.go:356
			_go_fuzz_dep_.CoverTab[18103]++
									bi++
//line /usr/local/go/src/path/filepath/path.go:357
			// _ = "end of CoverTab[18103]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:358
			_go_fuzz_dep_.CoverTab[18104]++
//line /usr/local/go/src/path/filepath/path.go:358
			// _ = "end of CoverTab[18104]"
//line /usr/local/go/src/path/filepath/path.go:358
		}
//line /usr/local/go/src/path/filepath/path.go:358
		// _ = "end of CoverTab[18094]"
//line /usr/local/go/src/path/filepath/path.go:358
		_go_fuzz_dep_.CoverTab[18095]++
								if ti < tl {
//line /usr/local/go/src/path/filepath/path.go:359
			_go_fuzz_dep_.CoverTab[18105]++
									ti++
//line /usr/local/go/src/path/filepath/path.go:360
			// _ = "end of CoverTab[18105]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:361
			_go_fuzz_dep_.CoverTab[18106]++
//line /usr/local/go/src/path/filepath/path.go:361
			// _ = "end of CoverTab[18106]"
//line /usr/local/go/src/path/filepath/path.go:361
		}
//line /usr/local/go/src/path/filepath/path.go:361
		// _ = "end of CoverTab[18095]"
//line /usr/local/go/src/path/filepath/path.go:361
		_go_fuzz_dep_.CoverTab[18096]++
								b0 = bi
								t0 = ti
//line /usr/local/go/src/path/filepath/path.go:363
		// _ = "end of CoverTab[18096]"
	}
//line /usr/local/go/src/path/filepath/path.go:364
	// _ = "end of CoverTab[18075]"
//line /usr/local/go/src/path/filepath/path.go:364
	_go_fuzz_dep_.CoverTab[18076]++
							if base[b0:bi] == ".." {
//line /usr/local/go/src/path/filepath/path.go:365
		_go_fuzz_dep_.CoverTab[18107]++
								return "", errors.New("Rel: can't make " + targpath + " relative to " + basepath)
//line /usr/local/go/src/path/filepath/path.go:366
		// _ = "end of CoverTab[18107]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:367
		_go_fuzz_dep_.CoverTab[18108]++
//line /usr/local/go/src/path/filepath/path.go:367
		// _ = "end of CoverTab[18108]"
//line /usr/local/go/src/path/filepath/path.go:367
	}
//line /usr/local/go/src/path/filepath/path.go:367
	// _ = "end of CoverTab[18076]"
//line /usr/local/go/src/path/filepath/path.go:367
	_go_fuzz_dep_.CoverTab[18077]++
							if b0 != bl {
//line /usr/local/go/src/path/filepath/path.go:368
		_go_fuzz_dep_.CoverTab[18109]++

								seps := strings.Count(base[b0:bl], string(Separator))
								size := 2 + seps*3
								if tl != t0 {
//line /usr/local/go/src/path/filepath/path.go:372
			_go_fuzz_dep_.CoverTab[18113]++
									size += 1 + tl - t0
//line /usr/local/go/src/path/filepath/path.go:373
			// _ = "end of CoverTab[18113]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:374
			_go_fuzz_dep_.CoverTab[18114]++
//line /usr/local/go/src/path/filepath/path.go:374
			// _ = "end of CoverTab[18114]"
//line /usr/local/go/src/path/filepath/path.go:374
		}
//line /usr/local/go/src/path/filepath/path.go:374
		// _ = "end of CoverTab[18109]"
//line /usr/local/go/src/path/filepath/path.go:374
		_go_fuzz_dep_.CoverTab[18110]++
								buf := make([]byte, size)
								n := copy(buf, "..")
								for i := 0; i < seps; i++ {
//line /usr/local/go/src/path/filepath/path.go:377
			_go_fuzz_dep_.CoverTab[18115]++
									buf[n] = Separator
									copy(buf[n+1:], "..")
									n += 3
//line /usr/local/go/src/path/filepath/path.go:380
			// _ = "end of CoverTab[18115]"
		}
//line /usr/local/go/src/path/filepath/path.go:381
		// _ = "end of CoverTab[18110]"
//line /usr/local/go/src/path/filepath/path.go:381
		_go_fuzz_dep_.CoverTab[18111]++
								if t0 != tl {
//line /usr/local/go/src/path/filepath/path.go:382
			_go_fuzz_dep_.CoverTab[18116]++
									buf[n] = Separator
									copy(buf[n+1:], targ[t0:])
//line /usr/local/go/src/path/filepath/path.go:384
			// _ = "end of CoverTab[18116]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:385
			_go_fuzz_dep_.CoverTab[18117]++
//line /usr/local/go/src/path/filepath/path.go:385
			// _ = "end of CoverTab[18117]"
//line /usr/local/go/src/path/filepath/path.go:385
		}
//line /usr/local/go/src/path/filepath/path.go:385
		// _ = "end of CoverTab[18111]"
//line /usr/local/go/src/path/filepath/path.go:385
		_go_fuzz_dep_.CoverTab[18112]++
								return string(buf), nil
//line /usr/local/go/src/path/filepath/path.go:386
		// _ = "end of CoverTab[18112]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:387
		_go_fuzz_dep_.CoverTab[18118]++
//line /usr/local/go/src/path/filepath/path.go:387
		// _ = "end of CoverTab[18118]"
//line /usr/local/go/src/path/filepath/path.go:387
	}
//line /usr/local/go/src/path/filepath/path.go:387
	// _ = "end of CoverTab[18077]"
//line /usr/local/go/src/path/filepath/path.go:387
	_go_fuzz_dep_.CoverTab[18078]++
							return targ[t0:], nil
//line /usr/local/go/src/path/filepath/path.go:388
	// _ = "end of CoverTab[18078]"
}

// SkipDir is used as a return value from WalkFuncs to indicate that
//line /usr/local/go/src/path/filepath/path.go:391
// the directory named in the call is to be skipped. It is not returned
//line /usr/local/go/src/path/filepath/path.go:391
// as an error by any function.
//line /usr/local/go/src/path/filepath/path.go:394
var SkipDir error = fs.SkipDir

// SkipAll is used as a return value from WalkFuncs to indicate that
//line /usr/local/go/src/path/filepath/path.go:396
// all remaining files and directories are to be skipped. It is not returned
//line /usr/local/go/src/path/filepath/path.go:396
// as an error by any function.
//line /usr/local/go/src/path/filepath/path.go:399
var SkipAll error = fs.SkipAll

// WalkFunc is the type of the function called by Walk to visit each
//line /usr/local/go/src/path/filepath/path.go:401
// file or directory.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// The path argument contains the argument to Walk as a prefix.
//line /usr/local/go/src/path/filepath/path.go:401
// That is, if Walk is called with root argument "dir" and finds a file
//line /usr/local/go/src/path/filepath/path.go:401
// named "a" in that directory, the walk function will be called with
//line /usr/local/go/src/path/filepath/path.go:401
// argument "dir/a".
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// The directory and file are joined with Join, which may clean the
//line /usr/local/go/src/path/filepath/path.go:401
// directory name: if Walk is called with the root argument "x/../dir"
//line /usr/local/go/src/path/filepath/path.go:401
// and finds a file named "a" in that directory, the walk function will
//line /usr/local/go/src/path/filepath/path.go:401
// be called with argument "dir/a", not "x/../dir/a".
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// The info argument is the fs.FileInfo for the named path.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// The error result returned by the function controls how Walk continues.
//line /usr/local/go/src/path/filepath/path.go:401
// If the function returns the special value SkipDir, Walk skips the
//line /usr/local/go/src/path/filepath/path.go:401
// current directory (path if info.IsDir() is true, otherwise path's
//line /usr/local/go/src/path/filepath/path.go:401
// parent directory). If the function returns the special value SkipAll,
//line /usr/local/go/src/path/filepath/path.go:401
// Walk skips all remaining files and directories. Otherwise, if the function
//line /usr/local/go/src/path/filepath/path.go:401
// returns a non-nil error, Walk stops entirely and returns that error.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// The err argument reports an error related to path, signaling that Walk
//line /usr/local/go/src/path/filepath/path.go:401
// will not walk into that directory. The function can decide how to
//line /usr/local/go/src/path/filepath/path.go:401
// handle that error; as described earlier, returning the error will
//line /usr/local/go/src/path/filepath/path.go:401
// cause Walk to stop walking the entire tree.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// Walk calls the function with a non-nil err argument in two cases.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// First, if an os.Lstat on the root directory or any directory or file
//line /usr/local/go/src/path/filepath/path.go:401
// in the tree fails, Walk calls the function with path set to that
//line /usr/local/go/src/path/filepath/path.go:401
// directory or file's path, info set to nil, and err set to the error
//line /usr/local/go/src/path/filepath/path.go:401
// from os.Lstat.
//line /usr/local/go/src/path/filepath/path.go:401
//
//line /usr/local/go/src/path/filepath/path.go:401
// Second, if a directory's Readdirnames method fails, Walk calls the
//line /usr/local/go/src/path/filepath/path.go:401
// function with path set to the directory's path, info, set to an
//line /usr/local/go/src/path/filepath/path.go:401
// fs.FileInfo describing the directory, and err set to the error from
//line /usr/local/go/src/path/filepath/path.go:401
// Readdirnames.
//line /usr/local/go/src/path/filepath/path.go:439
type WalkFunc func(path string, info fs.FileInfo, err error) error

var lstat = os.Lstat	// for testing

// walkDir recursively descends path, calling walkDirFn.
func walkDir(path string, d fs.DirEntry, walkDirFn fs.WalkDirFunc) error {
//line /usr/local/go/src/path/filepath/path.go:444
	_go_fuzz_dep_.CoverTab[18119]++
							if err := walkDirFn(path, d, nil); err != nil || func() bool {
//line /usr/local/go/src/path/filepath/path.go:445
		_go_fuzz_dep_.CoverTab[18123]++
//line /usr/local/go/src/path/filepath/path.go:445
		return !d.IsDir()
//line /usr/local/go/src/path/filepath/path.go:445
		// _ = "end of CoverTab[18123]"
//line /usr/local/go/src/path/filepath/path.go:445
	}() {
//line /usr/local/go/src/path/filepath/path.go:445
		_go_fuzz_dep_.CoverTab[18124]++
								if err == SkipDir && func() bool {
//line /usr/local/go/src/path/filepath/path.go:446
			_go_fuzz_dep_.CoverTab[18126]++
//line /usr/local/go/src/path/filepath/path.go:446
			return d.IsDir()
//line /usr/local/go/src/path/filepath/path.go:446
			// _ = "end of CoverTab[18126]"
//line /usr/local/go/src/path/filepath/path.go:446
		}() {
//line /usr/local/go/src/path/filepath/path.go:446
			_go_fuzz_dep_.CoverTab[18127]++

									err = nil
//line /usr/local/go/src/path/filepath/path.go:448
			// _ = "end of CoverTab[18127]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:449
			_go_fuzz_dep_.CoverTab[18128]++
//line /usr/local/go/src/path/filepath/path.go:449
			// _ = "end of CoverTab[18128]"
//line /usr/local/go/src/path/filepath/path.go:449
		}
//line /usr/local/go/src/path/filepath/path.go:449
		// _ = "end of CoverTab[18124]"
//line /usr/local/go/src/path/filepath/path.go:449
		_go_fuzz_dep_.CoverTab[18125]++
								return err
//line /usr/local/go/src/path/filepath/path.go:450
		// _ = "end of CoverTab[18125]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:451
		_go_fuzz_dep_.CoverTab[18129]++
//line /usr/local/go/src/path/filepath/path.go:451
		// _ = "end of CoverTab[18129]"
//line /usr/local/go/src/path/filepath/path.go:451
	}
//line /usr/local/go/src/path/filepath/path.go:451
	// _ = "end of CoverTab[18119]"
//line /usr/local/go/src/path/filepath/path.go:451
	_go_fuzz_dep_.CoverTab[18120]++

							dirs, err := readDir(path)
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:454
		_go_fuzz_dep_.CoverTab[18130]++

								err = walkDirFn(path, d, err)
								if err != nil {
//line /usr/local/go/src/path/filepath/path.go:457
			_go_fuzz_dep_.CoverTab[18131]++
									if err == SkipDir && func() bool {
//line /usr/local/go/src/path/filepath/path.go:458
				_go_fuzz_dep_.CoverTab[18133]++
//line /usr/local/go/src/path/filepath/path.go:458
				return d.IsDir()
//line /usr/local/go/src/path/filepath/path.go:458
				// _ = "end of CoverTab[18133]"
//line /usr/local/go/src/path/filepath/path.go:458
			}() {
//line /usr/local/go/src/path/filepath/path.go:458
				_go_fuzz_dep_.CoverTab[18134]++
										err = nil
//line /usr/local/go/src/path/filepath/path.go:459
				// _ = "end of CoverTab[18134]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:460
				_go_fuzz_dep_.CoverTab[18135]++
//line /usr/local/go/src/path/filepath/path.go:460
				// _ = "end of CoverTab[18135]"
//line /usr/local/go/src/path/filepath/path.go:460
			}
//line /usr/local/go/src/path/filepath/path.go:460
			// _ = "end of CoverTab[18131]"
//line /usr/local/go/src/path/filepath/path.go:460
			_go_fuzz_dep_.CoverTab[18132]++
									return err
//line /usr/local/go/src/path/filepath/path.go:461
			// _ = "end of CoverTab[18132]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:462
			_go_fuzz_dep_.CoverTab[18136]++
//line /usr/local/go/src/path/filepath/path.go:462
			// _ = "end of CoverTab[18136]"
//line /usr/local/go/src/path/filepath/path.go:462
		}
//line /usr/local/go/src/path/filepath/path.go:462
		// _ = "end of CoverTab[18130]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:463
		_go_fuzz_dep_.CoverTab[18137]++
//line /usr/local/go/src/path/filepath/path.go:463
		// _ = "end of CoverTab[18137]"
//line /usr/local/go/src/path/filepath/path.go:463
	}
//line /usr/local/go/src/path/filepath/path.go:463
	// _ = "end of CoverTab[18120]"
//line /usr/local/go/src/path/filepath/path.go:463
	_go_fuzz_dep_.CoverTab[18121]++

							for _, d1 := range dirs {
//line /usr/local/go/src/path/filepath/path.go:465
		_go_fuzz_dep_.CoverTab[18138]++
								path1 := Join(path, d1.Name())
								if err := walkDir(path1, d1, walkDirFn); err != nil {
//line /usr/local/go/src/path/filepath/path.go:467
			_go_fuzz_dep_.CoverTab[18139]++
									if err == SkipDir {
//line /usr/local/go/src/path/filepath/path.go:468
				_go_fuzz_dep_.CoverTab[18141]++
										break
//line /usr/local/go/src/path/filepath/path.go:469
				// _ = "end of CoverTab[18141]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:470
				_go_fuzz_dep_.CoverTab[18142]++
//line /usr/local/go/src/path/filepath/path.go:470
				// _ = "end of CoverTab[18142]"
//line /usr/local/go/src/path/filepath/path.go:470
			}
//line /usr/local/go/src/path/filepath/path.go:470
			// _ = "end of CoverTab[18139]"
//line /usr/local/go/src/path/filepath/path.go:470
			_go_fuzz_dep_.CoverTab[18140]++
									return err
//line /usr/local/go/src/path/filepath/path.go:471
			// _ = "end of CoverTab[18140]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:472
			_go_fuzz_dep_.CoverTab[18143]++
//line /usr/local/go/src/path/filepath/path.go:472
			// _ = "end of CoverTab[18143]"
//line /usr/local/go/src/path/filepath/path.go:472
		}
//line /usr/local/go/src/path/filepath/path.go:472
		// _ = "end of CoverTab[18138]"
	}
//line /usr/local/go/src/path/filepath/path.go:473
	// _ = "end of CoverTab[18121]"
//line /usr/local/go/src/path/filepath/path.go:473
	_go_fuzz_dep_.CoverTab[18122]++
							return nil
//line /usr/local/go/src/path/filepath/path.go:474
	// _ = "end of CoverTab[18122]"
}

// walk recursively descends path, calling walkFn.
func walk(path string, info fs.FileInfo, walkFn WalkFunc) error {
//line /usr/local/go/src/path/filepath/path.go:478
	_go_fuzz_dep_.CoverTab[18144]++
							if !info.IsDir() {
//line /usr/local/go/src/path/filepath/path.go:479
		_go_fuzz_dep_.CoverTab[18148]++
								return walkFn(path, info, nil)
//line /usr/local/go/src/path/filepath/path.go:480
		// _ = "end of CoverTab[18148]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:481
		_go_fuzz_dep_.CoverTab[18149]++
//line /usr/local/go/src/path/filepath/path.go:481
		// _ = "end of CoverTab[18149]"
//line /usr/local/go/src/path/filepath/path.go:481
	}
//line /usr/local/go/src/path/filepath/path.go:481
	// _ = "end of CoverTab[18144]"
//line /usr/local/go/src/path/filepath/path.go:481
	_go_fuzz_dep_.CoverTab[18145]++

							names, err := readDirNames(path)
							err1 := walkFn(path, info, err)

//line /usr/local/go/src/path/filepath/path.go:488
	if err != nil || func() bool {
//line /usr/local/go/src/path/filepath/path.go:488
		_go_fuzz_dep_.CoverTab[18150]++
//line /usr/local/go/src/path/filepath/path.go:488
		return err1 != nil
//line /usr/local/go/src/path/filepath/path.go:488
		// _ = "end of CoverTab[18150]"
//line /usr/local/go/src/path/filepath/path.go:488
	}() {
//line /usr/local/go/src/path/filepath/path.go:488
		_go_fuzz_dep_.CoverTab[18151]++

//line /usr/local/go/src/path/filepath/path.go:493
		return err1
//line /usr/local/go/src/path/filepath/path.go:493
		// _ = "end of CoverTab[18151]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:494
		_go_fuzz_dep_.CoverTab[18152]++
//line /usr/local/go/src/path/filepath/path.go:494
		// _ = "end of CoverTab[18152]"
//line /usr/local/go/src/path/filepath/path.go:494
	}
//line /usr/local/go/src/path/filepath/path.go:494
	// _ = "end of CoverTab[18145]"
//line /usr/local/go/src/path/filepath/path.go:494
	_go_fuzz_dep_.CoverTab[18146]++

							for _, name := range names {
//line /usr/local/go/src/path/filepath/path.go:496
		_go_fuzz_dep_.CoverTab[18153]++
								filename := Join(path, name)
								fileInfo, err := lstat(filename)
								if err != nil {
//line /usr/local/go/src/path/filepath/path.go:499
			_go_fuzz_dep_.CoverTab[18154]++
									if err := walkFn(filename, fileInfo, err); err != nil && func() bool {
//line /usr/local/go/src/path/filepath/path.go:500
				_go_fuzz_dep_.CoverTab[18155]++
//line /usr/local/go/src/path/filepath/path.go:500
				return err != SkipDir
//line /usr/local/go/src/path/filepath/path.go:500
				// _ = "end of CoverTab[18155]"
//line /usr/local/go/src/path/filepath/path.go:500
			}() {
//line /usr/local/go/src/path/filepath/path.go:500
				_go_fuzz_dep_.CoverTab[18156]++
										return err
//line /usr/local/go/src/path/filepath/path.go:501
				// _ = "end of CoverTab[18156]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:502
				_go_fuzz_dep_.CoverTab[18157]++
//line /usr/local/go/src/path/filepath/path.go:502
				// _ = "end of CoverTab[18157]"
//line /usr/local/go/src/path/filepath/path.go:502
			}
//line /usr/local/go/src/path/filepath/path.go:502
			// _ = "end of CoverTab[18154]"
		} else {
//line /usr/local/go/src/path/filepath/path.go:503
			_go_fuzz_dep_.CoverTab[18158]++
									err = walk(filename, fileInfo, walkFn)
									if err != nil {
//line /usr/local/go/src/path/filepath/path.go:505
				_go_fuzz_dep_.CoverTab[18159]++
										if !fileInfo.IsDir() || func() bool {
//line /usr/local/go/src/path/filepath/path.go:506
					_go_fuzz_dep_.CoverTab[18160]++
//line /usr/local/go/src/path/filepath/path.go:506
					return err != SkipDir
//line /usr/local/go/src/path/filepath/path.go:506
					// _ = "end of CoverTab[18160]"
//line /usr/local/go/src/path/filepath/path.go:506
				}() {
//line /usr/local/go/src/path/filepath/path.go:506
					_go_fuzz_dep_.CoverTab[18161]++
											return err
//line /usr/local/go/src/path/filepath/path.go:507
					// _ = "end of CoverTab[18161]"
				} else {
//line /usr/local/go/src/path/filepath/path.go:508
					_go_fuzz_dep_.CoverTab[18162]++
//line /usr/local/go/src/path/filepath/path.go:508
					// _ = "end of CoverTab[18162]"
//line /usr/local/go/src/path/filepath/path.go:508
				}
//line /usr/local/go/src/path/filepath/path.go:508
				// _ = "end of CoverTab[18159]"
			} else {
//line /usr/local/go/src/path/filepath/path.go:509
				_go_fuzz_dep_.CoverTab[18163]++
//line /usr/local/go/src/path/filepath/path.go:509
				// _ = "end of CoverTab[18163]"
//line /usr/local/go/src/path/filepath/path.go:509
			}
//line /usr/local/go/src/path/filepath/path.go:509
			// _ = "end of CoverTab[18158]"
		}
//line /usr/local/go/src/path/filepath/path.go:510
		// _ = "end of CoverTab[18153]"
	}
//line /usr/local/go/src/path/filepath/path.go:511
	// _ = "end of CoverTab[18146]"
//line /usr/local/go/src/path/filepath/path.go:511
	_go_fuzz_dep_.CoverTab[18147]++
							return nil
//line /usr/local/go/src/path/filepath/path.go:512
	// _ = "end of CoverTab[18147]"
}

// WalkDir walks the file tree rooted at root, calling fn for each file or
//line /usr/local/go/src/path/filepath/path.go:515
// directory in the tree, including root.
//line /usr/local/go/src/path/filepath/path.go:515
//
//line /usr/local/go/src/path/filepath/path.go:515
// All errors that arise visiting files and directories are filtered by fn:
//line /usr/local/go/src/path/filepath/path.go:515
// see the fs.WalkDirFunc documentation for details.
//line /usr/local/go/src/path/filepath/path.go:515
//
//line /usr/local/go/src/path/filepath/path.go:515
// The files are walked in lexical order, which makes the output deterministic
//line /usr/local/go/src/path/filepath/path.go:515
// but requires WalkDir to read an entire directory into memory before proceeding
//line /usr/local/go/src/path/filepath/path.go:515
// to walk that directory.
//line /usr/local/go/src/path/filepath/path.go:515
//
//line /usr/local/go/src/path/filepath/path.go:515
// WalkDir does not follow symbolic links.
//line /usr/local/go/src/path/filepath/path.go:515
//
//line /usr/local/go/src/path/filepath/path.go:515
// WalkDir calls fn with paths that use the separator character appropriate
//line /usr/local/go/src/path/filepath/path.go:515
// for the operating system. This is unlike [io/fs.WalkDir], which always
//line /usr/local/go/src/path/filepath/path.go:515
// uses slash separated paths.
//line /usr/local/go/src/path/filepath/path.go:530
func WalkDir(root string, fn fs.WalkDirFunc) error {
//line /usr/local/go/src/path/filepath/path.go:530
	_go_fuzz_dep_.CoverTab[18164]++
							info, err := os.Lstat(root)
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:532
		_go_fuzz_dep_.CoverTab[18167]++
								err = fn(root, nil, err)
//line /usr/local/go/src/path/filepath/path.go:533
		// _ = "end of CoverTab[18167]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:534
		_go_fuzz_dep_.CoverTab[18168]++
								err = walkDir(root, &statDirEntry{info}, fn)
//line /usr/local/go/src/path/filepath/path.go:535
		// _ = "end of CoverTab[18168]"
	}
//line /usr/local/go/src/path/filepath/path.go:536
	// _ = "end of CoverTab[18164]"
//line /usr/local/go/src/path/filepath/path.go:536
	_go_fuzz_dep_.CoverTab[18165]++
							if err == SkipDir || func() bool {
//line /usr/local/go/src/path/filepath/path.go:537
		_go_fuzz_dep_.CoverTab[18169]++
//line /usr/local/go/src/path/filepath/path.go:537
		return err == SkipAll
//line /usr/local/go/src/path/filepath/path.go:537
		// _ = "end of CoverTab[18169]"
//line /usr/local/go/src/path/filepath/path.go:537
	}() {
//line /usr/local/go/src/path/filepath/path.go:537
		_go_fuzz_dep_.CoverTab[18170]++
								return nil
//line /usr/local/go/src/path/filepath/path.go:538
		// _ = "end of CoverTab[18170]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:539
		_go_fuzz_dep_.CoverTab[18171]++
//line /usr/local/go/src/path/filepath/path.go:539
		// _ = "end of CoverTab[18171]"
//line /usr/local/go/src/path/filepath/path.go:539
	}
//line /usr/local/go/src/path/filepath/path.go:539
	// _ = "end of CoverTab[18165]"
//line /usr/local/go/src/path/filepath/path.go:539
	_go_fuzz_dep_.CoverTab[18166]++
							return err
//line /usr/local/go/src/path/filepath/path.go:540
	// _ = "end of CoverTab[18166]"
}

type statDirEntry struct {
	info fs.FileInfo
}

func (d *statDirEntry) Name() string {
//line /usr/local/go/src/path/filepath/path.go:547
	_go_fuzz_dep_.CoverTab[18172]++
//line /usr/local/go/src/path/filepath/path.go:547
	return d.info.Name()
//line /usr/local/go/src/path/filepath/path.go:547
	// _ = "end of CoverTab[18172]"
//line /usr/local/go/src/path/filepath/path.go:547
}
func (d *statDirEntry) IsDir() bool {
//line /usr/local/go/src/path/filepath/path.go:548
	_go_fuzz_dep_.CoverTab[18173]++
//line /usr/local/go/src/path/filepath/path.go:548
	return d.info.IsDir()
//line /usr/local/go/src/path/filepath/path.go:548
	// _ = "end of CoverTab[18173]"
//line /usr/local/go/src/path/filepath/path.go:548
}
func (d *statDirEntry) Type() fs.FileMode {
//line /usr/local/go/src/path/filepath/path.go:549
	_go_fuzz_dep_.CoverTab[18174]++
//line /usr/local/go/src/path/filepath/path.go:549
	return d.info.Mode().Type()
//line /usr/local/go/src/path/filepath/path.go:549
	// _ = "end of CoverTab[18174]"
//line /usr/local/go/src/path/filepath/path.go:549
}
func (d *statDirEntry) Info() (fs.FileInfo, error) {
//line /usr/local/go/src/path/filepath/path.go:550
	_go_fuzz_dep_.CoverTab[18175]++
//line /usr/local/go/src/path/filepath/path.go:550
	return d.info, nil
//line /usr/local/go/src/path/filepath/path.go:550
	// _ = "end of CoverTab[18175]"
//line /usr/local/go/src/path/filepath/path.go:550
}

// Walk walks the file tree rooted at root, calling fn for each file or
//line /usr/local/go/src/path/filepath/path.go:552
// directory in the tree, including root.
//line /usr/local/go/src/path/filepath/path.go:552
//
//line /usr/local/go/src/path/filepath/path.go:552
// All errors that arise visiting files and directories are filtered by fn:
//line /usr/local/go/src/path/filepath/path.go:552
// see the WalkFunc documentation for details.
//line /usr/local/go/src/path/filepath/path.go:552
//
//line /usr/local/go/src/path/filepath/path.go:552
// The files are walked in lexical order, which makes the output deterministic
//line /usr/local/go/src/path/filepath/path.go:552
// but requires Walk to read an entire directory into memory before proceeding
//line /usr/local/go/src/path/filepath/path.go:552
// to walk that directory.
//line /usr/local/go/src/path/filepath/path.go:552
//
//line /usr/local/go/src/path/filepath/path.go:552
// Walk does not follow symbolic links.
//line /usr/local/go/src/path/filepath/path.go:552
//
//line /usr/local/go/src/path/filepath/path.go:552
// Walk is less efficient than WalkDir, introduced in Go 1.16,
//line /usr/local/go/src/path/filepath/path.go:552
// which avoids calling os.Lstat on every visited file or directory.
//line /usr/local/go/src/path/filepath/path.go:566
func Walk(root string, fn WalkFunc) error {
//line /usr/local/go/src/path/filepath/path.go:566
	_go_fuzz_dep_.CoverTab[18176]++
							info, err := os.Lstat(root)
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:568
		_go_fuzz_dep_.CoverTab[18179]++
								err = fn(root, nil, err)
//line /usr/local/go/src/path/filepath/path.go:569
		// _ = "end of CoverTab[18179]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:570
		_go_fuzz_dep_.CoverTab[18180]++
								err = walk(root, info, fn)
//line /usr/local/go/src/path/filepath/path.go:571
		// _ = "end of CoverTab[18180]"
	}
//line /usr/local/go/src/path/filepath/path.go:572
	// _ = "end of CoverTab[18176]"
//line /usr/local/go/src/path/filepath/path.go:572
	_go_fuzz_dep_.CoverTab[18177]++
							if err == SkipDir || func() bool {
//line /usr/local/go/src/path/filepath/path.go:573
		_go_fuzz_dep_.CoverTab[18181]++
//line /usr/local/go/src/path/filepath/path.go:573
		return err == SkipAll
//line /usr/local/go/src/path/filepath/path.go:573
		// _ = "end of CoverTab[18181]"
//line /usr/local/go/src/path/filepath/path.go:573
	}() {
//line /usr/local/go/src/path/filepath/path.go:573
		_go_fuzz_dep_.CoverTab[18182]++
								return nil
//line /usr/local/go/src/path/filepath/path.go:574
		// _ = "end of CoverTab[18182]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:575
		_go_fuzz_dep_.CoverTab[18183]++
//line /usr/local/go/src/path/filepath/path.go:575
		// _ = "end of CoverTab[18183]"
//line /usr/local/go/src/path/filepath/path.go:575
	}
//line /usr/local/go/src/path/filepath/path.go:575
	// _ = "end of CoverTab[18177]"
//line /usr/local/go/src/path/filepath/path.go:575
	_go_fuzz_dep_.CoverTab[18178]++
							return err
//line /usr/local/go/src/path/filepath/path.go:576
	// _ = "end of CoverTab[18178]"
}

// readDir reads the directory named by dirname and returns
//line /usr/local/go/src/path/filepath/path.go:579
// a sorted list of directory entries.
//line /usr/local/go/src/path/filepath/path.go:581
func readDir(dirname string) ([]fs.DirEntry, error) {
//line /usr/local/go/src/path/filepath/path.go:581
	_go_fuzz_dep_.CoverTab[18184]++
							f, err := os.Open(dirname)
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:583
		_go_fuzz_dep_.CoverTab[18188]++
								return nil, err
//line /usr/local/go/src/path/filepath/path.go:584
		// _ = "end of CoverTab[18188]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:585
		_go_fuzz_dep_.CoverTab[18189]++
//line /usr/local/go/src/path/filepath/path.go:585
		// _ = "end of CoverTab[18189]"
//line /usr/local/go/src/path/filepath/path.go:585
	}
//line /usr/local/go/src/path/filepath/path.go:585
	// _ = "end of CoverTab[18184]"
//line /usr/local/go/src/path/filepath/path.go:585
	_go_fuzz_dep_.CoverTab[18185]++
							dirs, err := f.ReadDir(-1)
							f.Close()
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:588
		_go_fuzz_dep_.CoverTab[18190]++
								return nil, err
//line /usr/local/go/src/path/filepath/path.go:589
		// _ = "end of CoverTab[18190]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:590
		_go_fuzz_dep_.CoverTab[18191]++
//line /usr/local/go/src/path/filepath/path.go:590
		// _ = "end of CoverTab[18191]"
//line /usr/local/go/src/path/filepath/path.go:590
	}
//line /usr/local/go/src/path/filepath/path.go:590
	// _ = "end of CoverTab[18185]"
//line /usr/local/go/src/path/filepath/path.go:590
	_go_fuzz_dep_.CoverTab[18186]++
							sort.Slice(dirs, func(i, j int) bool {
//line /usr/local/go/src/path/filepath/path.go:591
		_go_fuzz_dep_.CoverTab[18192]++
//line /usr/local/go/src/path/filepath/path.go:591
		return dirs[i].Name() < dirs[j].Name()
//line /usr/local/go/src/path/filepath/path.go:591
		// _ = "end of CoverTab[18192]"
//line /usr/local/go/src/path/filepath/path.go:591
	})
//line /usr/local/go/src/path/filepath/path.go:591
	// _ = "end of CoverTab[18186]"
//line /usr/local/go/src/path/filepath/path.go:591
	_go_fuzz_dep_.CoverTab[18187]++
							return dirs, nil
//line /usr/local/go/src/path/filepath/path.go:592
	// _ = "end of CoverTab[18187]"
}

// readDirNames reads the directory named by dirname and returns
//line /usr/local/go/src/path/filepath/path.go:595
// a sorted list of directory entry names.
//line /usr/local/go/src/path/filepath/path.go:597
func readDirNames(dirname string) ([]string, error) {
//line /usr/local/go/src/path/filepath/path.go:597
	_go_fuzz_dep_.CoverTab[18193]++
							f, err := os.Open(dirname)
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:599
		_go_fuzz_dep_.CoverTab[18196]++
								return nil, err
//line /usr/local/go/src/path/filepath/path.go:600
		// _ = "end of CoverTab[18196]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:601
		_go_fuzz_dep_.CoverTab[18197]++
//line /usr/local/go/src/path/filepath/path.go:601
		// _ = "end of CoverTab[18197]"
//line /usr/local/go/src/path/filepath/path.go:601
	}
//line /usr/local/go/src/path/filepath/path.go:601
	// _ = "end of CoverTab[18193]"
//line /usr/local/go/src/path/filepath/path.go:601
	_go_fuzz_dep_.CoverTab[18194]++
							names, err := f.Readdirnames(-1)
							f.Close()
							if err != nil {
//line /usr/local/go/src/path/filepath/path.go:604
		_go_fuzz_dep_.CoverTab[18198]++
								return nil, err
//line /usr/local/go/src/path/filepath/path.go:605
		// _ = "end of CoverTab[18198]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:606
		_go_fuzz_dep_.CoverTab[18199]++
//line /usr/local/go/src/path/filepath/path.go:606
		// _ = "end of CoverTab[18199]"
//line /usr/local/go/src/path/filepath/path.go:606
	}
//line /usr/local/go/src/path/filepath/path.go:606
	// _ = "end of CoverTab[18194]"
//line /usr/local/go/src/path/filepath/path.go:606
	_go_fuzz_dep_.CoverTab[18195]++
							sort.Strings(names)
							return names, nil
//line /usr/local/go/src/path/filepath/path.go:608
	// _ = "end of CoverTab[18195]"
}

// Base returns the last element of path.
//line /usr/local/go/src/path/filepath/path.go:611
// Trailing path separators are removed before extracting the last element.
//line /usr/local/go/src/path/filepath/path.go:611
// If the path is empty, Base returns ".".
//line /usr/local/go/src/path/filepath/path.go:611
// If the path consists entirely of separators, Base returns a single separator.
//line /usr/local/go/src/path/filepath/path.go:615
func Base(path string) string {
//line /usr/local/go/src/path/filepath/path.go:615
	_go_fuzz_dep_.CoverTab[18200]++
							if path == "" {
//line /usr/local/go/src/path/filepath/path.go:616
		_go_fuzz_dep_.CoverTab[18206]++
								return "."
//line /usr/local/go/src/path/filepath/path.go:617
		// _ = "end of CoverTab[18206]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:618
		_go_fuzz_dep_.CoverTab[18207]++
//line /usr/local/go/src/path/filepath/path.go:618
		// _ = "end of CoverTab[18207]"
//line /usr/local/go/src/path/filepath/path.go:618
	}
//line /usr/local/go/src/path/filepath/path.go:618
	// _ = "end of CoverTab[18200]"
//line /usr/local/go/src/path/filepath/path.go:618
	_go_fuzz_dep_.CoverTab[18201]++

							for len(path) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:620
		_go_fuzz_dep_.CoverTab[18208]++
//line /usr/local/go/src/path/filepath/path.go:620
		return os.IsPathSeparator(path[len(path)-1])
//line /usr/local/go/src/path/filepath/path.go:620
		// _ = "end of CoverTab[18208]"
//line /usr/local/go/src/path/filepath/path.go:620
	}() {
//line /usr/local/go/src/path/filepath/path.go:620
		_go_fuzz_dep_.CoverTab[18209]++
								path = path[0 : len(path)-1]
//line /usr/local/go/src/path/filepath/path.go:621
		// _ = "end of CoverTab[18209]"
	}
//line /usr/local/go/src/path/filepath/path.go:622
	// _ = "end of CoverTab[18201]"
//line /usr/local/go/src/path/filepath/path.go:622
	_go_fuzz_dep_.CoverTab[18202]++

							path = path[len(VolumeName(path)):]

							i := len(path) - 1
							for i >= 0 && func() bool {
//line /usr/local/go/src/path/filepath/path.go:627
		_go_fuzz_dep_.CoverTab[18210]++
//line /usr/local/go/src/path/filepath/path.go:627
		return !os.IsPathSeparator(path[i])
//line /usr/local/go/src/path/filepath/path.go:627
		// _ = "end of CoverTab[18210]"
//line /usr/local/go/src/path/filepath/path.go:627
	}() {
//line /usr/local/go/src/path/filepath/path.go:627
		_go_fuzz_dep_.CoverTab[18211]++
								i--
//line /usr/local/go/src/path/filepath/path.go:628
		// _ = "end of CoverTab[18211]"
	}
//line /usr/local/go/src/path/filepath/path.go:629
	// _ = "end of CoverTab[18202]"
//line /usr/local/go/src/path/filepath/path.go:629
	_go_fuzz_dep_.CoverTab[18203]++
							if i >= 0 {
//line /usr/local/go/src/path/filepath/path.go:630
		_go_fuzz_dep_.CoverTab[18212]++
								path = path[i+1:]
//line /usr/local/go/src/path/filepath/path.go:631
		// _ = "end of CoverTab[18212]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:632
		_go_fuzz_dep_.CoverTab[18213]++
//line /usr/local/go/src/path/filepath/path.go:632
		// _ = "end of CoverTab[18213]"
//line /usr/local/go/src/path/filepath/path.go:632
	}
//line /usr/local/go/src/path/filepath/path.go:632
	// _ = "end of CoverTab[18203]"
//line /usr/local/go/src/path/filepath/path.go:632
	_go_fuzz_dep_.CoverTab[18204]++

							if path == "" {
//line /usr/local/go/src/path/filepath/path.go:634
		_go_fuzz_dep_.CoverTab[18214]++
								return string(Separator)
//line /usr/local/go/src/path/filepath/path.go:635
		// _ = "end of CoverTab[18214]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:636
		_go_fuzz_dep_.CoverTab[18215]++
//line /usr/local/go/src/path/filepath/path.go:636
		// _ = "end of CoverTab[18215]"
//line /usr/local/go/src/path/filepath/path.go:636
	}
//line /usr/local/go/src/path/filepath/path.go:636
	// _ = "end of CoverTab[18204]"
//line /usr/local/go/src/path/filepath/path.go:636
	_go_fuzz_dep_.CoverTab[18205]++
							return path
//line /usr/local/go/src/path/filepath/path.go:637
	// _ = "end of CoverTab[18205]"
}

// Dir returns all but the last element of path, typically the path's directory.
//line /usr/local/go/src/path/filepath/path.go:640
// After dropping the final element, Dir calls Clean on the path and trailing
//line /usr/local/go/src/path/filepath/path.go:640
// slashes are removed.
//line /usr/local/go/src/path/filepath/path.go:640
// If the path is empty, Dir returns ".".
//line /usr/local/go/src/path/filepath/path.go:640
// If the path consists entirely of separators, Dir returns a single separator.
//line /usr/local/go/src/path/filepath/path.go:640
// The returned path does not end in a separator unless it is the root directory.
//line /usr/local/go/src/path/filepath/path.go:646
func Dir(path string) string {
//line /usr/local/go/src/path/filepath/path.go:646
	_go_fuzz_dep_.CoverTab[18216]++
							vol := VolumeName(path)
							i := len(path) - 1
							for i >= len(vol) && func() bool {
//line /usr/local/go/src/path/filepath/path.go:649
		_go_fuzz_dep_.CoverTab[18219]++
//line /usr/local/go/src/path/filepath/path.go:649
		return !os.IsPathSeparator(path[i])
//line /usr/local/go/src/path/filepath/path.go:649
		// _ = "end of CoverTab[18219]"
//line /usr/local/go/src/path/filepath/path.go:649
	}() {
//line /usr/local/go/src/path/filepath/path.go:649
		_go_fuzz_dep_.CoverTab[18220]++
								i--
//line /usr/local/go/src/path/filepath/path.go:650
		// _ = "end of CoverTab[18220]"
	}
//line /usr/local/go/src/path/filepath/path.go:651
	// _ = "end of CoverTab[18216]"
//line /usr/local/go/src/path/filepath/path.go:651
	_go_fuzz_dep_.CoverTab[18217]++
							dir := Clean(path[len(vol) : i+1])
							if dir == "." && func() bool {
//line /usr/local/go/src/path/filepath/path.go:653
		_go_fuzz_dep_.CoverTab[18221]++
//line /usr/local/go/src/path/filepath/path.go:653
		return len(vol) > 2
//line /usr/local/go/src/path/filepath/path.go:653
		// _ = "end of CoverTab[18221]"
//line /usr/local/go/src/path/filepath/path.go:653
	}() {
//line /usr/local/go/src/path/filepath/path.go:653
		_go_fuzz_dep_.CoverTab[18222]++

								return vol
//line /usr/local/go/src/path/filepath/path.go:655
		// _ = "end of CoverTab[18222]"
	} else {
//line /usr/local/go/src/path/filepath/path.go:656
		_go_fuzz_dep_.CoverTab[18223]++
//line /usr/local/go/src/path/filepath/path.go:656
		// _ = "end of CoverTab[18223]"
//line /usr/local/go/src/path/filepath/path.go:656
	}
//line /usr/local/go/src/path/filepath/path.go:656
	// _ = "end of CoverTab[18217]"
//line /usr/local/go/src/path/filepath/path.go:656
	_go_fuzz_dep_.CoverTab[18218]++
							return vol + dir
//line /usr/local/go/src/path/filepath/path.go:657
	// _ = "end of CoverTab[18218]"
}

// VolumeName returns leading volume name.
//line /usr/local/go/src/path/filepath/path.go:660
// Given "C:\foo\bar" it returns "C:" on Windows.
//line /usr/local/go/src/path/filepath/path.go:660
// Given "\\host\share\foo" it returns "\\host\share".
//line /usr/local/go/src/path/filepath/path.go:660
// On other platforms it returns "".
//line /usr/local/go/src/path/filepath/path.go:664
func VolumeName(path string) string {
//line /usr/local/go/src/path/filepath/path.go:664
	_go_fuzz_dep_.CoverTab[18224]++
							return FromSlash(path[:volumeNameLen(path)])
//line /usr/local/go/src/path/filepath/path.go:665
	// _ = "end of CoverTab[18224]"
}

//line /usr/local/go/src/path/filepath/path.go:666
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/path/filepath/path.go:666
var _ = _go_fuzz_dep_.CoverTab
