// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/path/filepath/symlink.go:5
package filepath

//line /usr/local/go/src/path/filepath/symlink.go:5
import (
//line /usr/local/go/src/path/filepath/symlink.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/path/filepath/symlink.go:5
)
//line /usr/local/go/src/path/filepath/symlink.go:5
import (
//line /usr/local/go/src/path/filepath/symlink.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/path/filepath/symlink.go:5
)

import (
	"errors"
	"io/fs"
	"os"
	"runtime"
	"syscall"
)

func walkSymlinks(path string) (string, error) {
//line /usr/local/go/src/path/filepath/symlink.go:15
	_go_fuzz_dep_.CoverTab[18240]++
							volLen := volumeNameLen(path)
							pathSeparator := string(os.PathSeparator)

							if volLen < len(path) && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:19
		_go_fuzz_dep_.CoverTab[18243]++
//line /usr/local/go/src/path/filepath/symlink.go:19
		return os.IsPathSeparator(path[volLen])
//line /usr/local/go/src/path/filepath/symlink.go:19
		// _ = "end of CoverTab[18243]"
//line /usr/local/go/src/path/filepath/symlink.go:19
	}() {
//line /usr/local/go/src/path/filepath/symlink.go:19
		_go_fuzz_dep_.CoverTab[18244]++
								volLen++
//line /usr/local/go/src/path/filepath/symlink.go:20
		// _ = "end of CoverTab[18244]"
	} else {
//line /usr/local/go/src/path/filepath/symlink.go:21
		_go_fuzz_dep_.CoverTab[18245]++
//line /usr/local/go/src/path/filepath/symlink.go:21
		// _ = "end of CoverTab[18245]"
//line /usr/local/go/src/path/filepath/symlink.go:21
	}
//line /usr/local/go/src/path/filepath/symlink.go:21
	// _ = "end of CoverTab[18240]"
//line /usr/local/go/src/path/filepath/symlink.go:21
	_go_fuzz_dep_.CoverTab[18241]++
							vol := path[:volLen]
							dest := vol
							linksWalked := 0
							for start, end := volLen, volLen; start < len(path); start = end {
//line /usr/local/go/src/path/filepath/symlink.go:25
		_go_fuzz_dep_.CoverTab[18246]++
								for start < len(path) && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:26
			_go_fuzz_dep_.CoverTab[18256]++
//line /usr/local/go/src/path/filepath/symlink.go:26
			return os.IsPathSeparator(path[start])
//line /usr/local/go/src/path/filepath/symlink.go:26
			// _ = "end of CoverTab[18256]"
//line /usr/local/go/src/path/filepath/symlink.go:26
		}() {
//line /usr/local/go/src/path/filepath/symlink.go:26
			_go_fuzz_dep_.CoverTab[18257]++
									start++
//line /usr/local/go/src/path/filepath/symlink.go:27
			// _ = "end of CoverTab[18257]"
		}
//line /usr/local/go/src/path/filepath/symlink.go:28
		// _ = "end of CoverTab[18246]"
//line /usr/local/go/src/path/filepath/symlink.go:28
		_go_fuzz_dep_.CoverTab[18247]++
								end = start
								for end < len(path) && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:30
			_go_fuzz_dep_.CoverTab[18258]++
//line /usr/local/go/src/path/filepath/symlink.go:30
			return !os.IsPathSeparator(path[end])
//line /usr/local/go/src/path/filepath/symlink.go:30
			// _ = "end of CoverTab[18258]"
//line /usr/local/go/src/path/filepath/symlink.go:30
		}() {
//line /usr/local/go/src/path/filepath/symlink.go:30
			_go_fuzz_dep_.CoverTab[18259]++
									end++
//line /usr/local/go/src/path/filepath/symlink.go:31
			// _ = "end of CoverTab[18259]"
		}
//line /usr/local/go/src/path/filepath/symlink.go:32
		// _ = "end of CoverTab[18247]"
//line /usr/local/go/src/path/filepath/symlink.go:32
		_go_fuzz_dep_.CoverTab[18248]++

//line /usr/local/go/src/path/filepath/symlink.go:37
		isWindowsDot := runtime.GOOS == "windows" && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:37
			_go_fuzz_dep_.CoverTab[18260]++
//line /usr/local/go/src/path/filepath/symlink.go:37
			return path[volumeNameLen(path):] == "."
//line /usr/local/go/src/path/filepath/symlink.go:37
			// _ = "end of CoverTab[18260]"
//line /usr/local/go/src/path/filepath/symlink.go:37
		}()

//line /usr/local/go/src/path/filepath/symlink.go:40
		if end == start {
//line /usr/local/go/src/path/filepath/symlink.go:40
			_go_fuzz_dep_.CoverTab[18261]++

									break
//line /usr/local/go/src/path/filepath/symlink.go:42
			// _ = "end of CoverTab[18261]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:43
			_go_fuzz_dep_.CoverTab[18262]++
//line /usr/local/go/src/path/filepath/symlink.go:43
			if path[start:end] == "." && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:43
				_go_fuzz_dep_.CoverTab[18263]++
//line /usr/local/go/src/path/filepath/symlink.go:43
				return !isWindowsDot
//line /usr/local/go/src/path/filepath/symlink.go:43
				// _ = "end of CoverTab[18263]"
//line /usr/local/go/src/path/filepath/symlink.go:43
			}() {
//line /usr/local/go/src/path/filepath/symlink.go:43
				_go_fuzz_dep_.CoverTab[18264]++

										continue
//line /usr/local/go/src/path/filepath/symlink.go:45
				// _ = "end of CoverTab[18264]"
			} else {
//line /usr/local/go/src/path/filepath/symlink.go:46
				_go_fuzz_dep_.CoverTab[18265]++
//line /usr/local/go/src/path/filepath/symlink.go:46
				if path[start:end] == ".." {
//line /usr/local/go/src/path/filepath/symlink.go:46
					_go_fuzz_dep_.CoverTab[18266]++

//line /usr/local/go/src/path/filepath/symlink.go:50
					// Set r to the index of the last slash in dest,
					// after the volume.
					var r int
					for r = len(dest) - 1; r >= volLen; r-- {
//line /usr/local/go/src/path/filepath/symlink.go:53
						_go_fuzz_dep_.CoverTab[18269]++
												if os.IsPathSeparator(dest[r]) {
//line /usr/local/go/src/path/filepath/symlink.go:54
							_go_fuzz_dep_.CoverTab[18270]++
													break
//line /usr/local/go/src/path/filepath/symlink.go:55
							// _ = "end of CoverTab[18270]"
						} else {
//line /usr/local/go/src/path/filepath/symlink.go:56
							_go_fuzz_dep_.CoverTab[18271]++
//line /usr/local/go/src/path/filepath/symlink.go:56
							// _ = "end of CoverTab[18271]"
//line /usr/local/go/src/path/filepath/symlink.go:56
						}
//line /usr/local/go/src/path/filepath/symlink.go:56
						// _ = "end of CoverTab[18269]"
					}
//line /usr/local/go/src/path/filepath/symlink.go:57
					// _ = "end of CoverTab[18266]"
//line /usr/local/go/src/path/filepath/symlink.go:57
					_go_fuzz_dep_.CoverTab[18267]++
											if r < volLen || func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:58
						_go_fuzz_dep_.CoverTab[18272]++
//line /usr/local/go/src/path/filepath/symlink.go:58
						return dest[r+1:] == ".."
//line /usr/local/go/src/path/filepath/symlink.go:58
						// _ = "end of CoverTab[18272]"
//line /usr/local/go/src/path/filepath/symlink.go:58
					}() {
//line /usr/local/go/src/path/filepath/symlink.go:58
						_go_fuzz_dep_.CoverTab[18273]++

//line /usr/local/go/src/path/filepath/symlink.go:63
						if len(dest) > volLen {
//line /usr/local/go/src/path/filepath/symlink.go:63
							_go_fuzz_dep_.CoverTab[18275]++
													dest += pathSeparator
//line /usr/local/go/src/path/filepath/symlink.go:64
							// _ = "end of CoverTab[18275]"
						} else {
//line /usr/local/go/src/path/filepath/symlink.go:65
							_go_fuzz_dep_.CoverTab[18276]++
//line /usr/local/go/src/path/filepath/symlink.go:65
							// _ = "end of CoverTab[18276]"
//line /usr/local/go/src/path/filepath/symlink.go:65
						}
//line /usr/local/go/src/path/filepath/symlink.go:65
						// _ = "end of CoverTab[18273]"
//line /usr/local/go/src/path/filepath/symlink.go:65
						_go_fuzz_dep_.CoverTab[18274]++
												dest += ".."
//line /usr/local/go/src/path/filepath/symlink.go:66
						// _ = "end of CoverTab[18274]"
					} else {
//line /usr/local/go/src/path/filepath/symlink.go:67
						_go_fuzz_dep_.CoverTab[18277]++

												dest = dest[:r]
//line /usr/local/go/src/path/filepath/symlink.go:69
						// _ = "end of CoverTab[18277]"
					}
//line /usr/local/go/src/path/filepath/symlink.go:70
					// _ = "end of CoverTab[18267]"
//line /usr/local/go/src/path/filepath/symlink.go:70
					_go_fuzz_dep_.CoverTab[18268]++
											continue
//line /usr/local/go/src/path/filepath/symlink.go:71
					// _ = "end of CoverTab[18268]"
				} else {
//line /usr/local/go/src/path/filepath/symlink.go:72
					_go_fuzz_dep_.CoverTab[18278]++
//line /usr/local/go/src/path/filepath/symlink.go:72
					// _ = "end of CoverTab[18278]"
//line /usr/local/go/src/path/filepath/symlink.go:72
				}
//line /usr/local/go/src/path/filepath/symlink.go:72
				// _ = "end of CoverTab[18265]"
//line /usr/local/go/src/path/filepath/symlink.go:72
			}
//line /usr/local/go/src/path/filepath/symlink.go:72
			// _ = "end of CoverTab[18262]"
//line /usr/local/go/src/path/filepath/symlink.go:72
		}
//line /usr/local/go/src/path/filepath/symlink.go:72
		// _ = "end of CoverTab[18248]"
//line /usr/local/go/src/path/filepath/symlink.go:72
		_go_fuzz_dep_.CoverTab[18249]++

//line /usr/local/go/src/path/filepath/symlink.go:76
		if len(dest) > volumeNameLen(dest) && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:76
			_go_fuzz_dep_.CoverTab[18279]++
//line /usr/local/go/src/path/filepath/symlink.go:76
			return !os.IsPathSeparator(dest[len(dest)-1])
//line /usr/local/go/src/path/filepath/symlink.go:76
			// _ = "end of CoverTab[18279]"
//line /usr/local/go/src/path/filepath/symlink.go:76
		}() {
//line /usr/local/go/src/path/filepath/symlink.go:76
			_go_fuzz_dep_.CoverTab[18280]++
									dest += pathSeparator
//line /usr/local/go/src/path/filepath/symlink.go:77
			// _ = "end of CoverTab[18280]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:78
			_go_fuzz_dep_.CoverTab[18281]++
//line /usr/local/go/src/path/filepath/symlink.go:78
			// _ = "end of CoverTab[18281]"
//line /usr/local/go/src/path/filepath/symlink.go:78
		}
//line /usr/local/go/src/path/filepath/symlink.go:78
		// _ = "end of CoverTab[18249]"
//line /usr/local/go/src/path/filepath/symlink.go:78
		_go_fuzz_dep_.CoverTab[18250]++

								dest += path[start:end]

//line /usr/local/go/src/path/filepath/symlink.go:84
		fi, err := os.Lstat(dest)
		if err != nil {
//line /usr/local/go/src/path/filepath/symlink.go:85
			_go_fuzz_dep_.CoverTab[18282]++
									return "", err
//line /usr/local/go/src/path/filepath/symlink.go:86
			// _ = "end of CoverTab[18282]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:87
			_go_fuzz_dep_.CoverTab[18283]++
//line /usr/local/go/src/path/filepath/symlink.go:87
			// _ = "end of CoverTab[18283]"
//line /usr/local/go/src/path/filepath/symlink.go:87
		}
//line /usr/local/go/src/path/filepath/symlink.go:87
		// _ = "end of CoverTab[18250]"
//line /usr/local/go/src/path/filepath/symlink.go:87
		_go_fuzz_dep_.CoverTab[18251]++

								if fi.Mode()&fs.ModeSymlink == 0 {
//line /usr/local/go/src/path/filepath/symlink.go:89
			_go_fuzz_dep_.CoverTab[18284]++
									if !fi.Mode().IsDir() && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:90
				_go_fuzz_dep_.CoverTab[18286]++
//line /usr/local/go/src/path/filepath/symlink.go:90
				return end < len(path)
//line /usr/local/go/src/path/filepath/symlink.go:90
				// _ = "end of CoverTab[18286]"
//line /usr/local/go/src/path/filepath/symlink.go:90
			}() {
//line /usr/local/go/src/path/filepath/symlink.go:90
				_go_fuzz_dep_.CoverTab[18287]++
										return "", syscall.ENOTDIR
//line /usr/local/go/src/path/filepath/symlink.go:91
				// _ = "end of CoverTab[18287]"
			} else {
//line /usr/local/go/src/path/filepath/symlink.go:92
				_go_fuzz_dep_.CoverTab[18288]++
//line /usr/local/go/src/path/filepath/symlink.go:92
				// _ = "end of CoverTab[18288]"
//line /usr/local/go/src/path/filepath/symlink.go:92
			}
//line /usr/local/go/src/path/filepath/symlink.go:92
			// _ = "end of CoverTab[18284]"
//line /usr/local/go/src/path/filepath/symlink.go:92
			_go_fuzz_dep_.CoverTab[18285]++
									continue
//line /usr/local/go/src/path/filepath/symlink.go:93
			// _ = "end of CoverTab[18285]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:94
			_go_fuzz_dep_.CoverTab[18289]++
//line /usr/local/go/src/path/filepath/symlink.go:94
			// _ = "end of CoverTab[18289]"
//line /usr/local/go/src/path/filepath/symlink.go:94
		}
//line /usr/local/go/src/path/filepath/symlink.go:94
		// _ = "end of CoverTab[18251]"
//line /usr/local/go/src/path/filepath/symlink.go:94
		_go_fuzz_dep_.CoverTab[18252]++

//line /usr/local/go/src/path/filepath/symlink.go:98
		linksWalked++
		if linksWalked > 255 {
//line /usr/local/go/src/path/filepath/symlink.go:99
			_go_fuzz_dep_.CoverTab[18290]++
									return "", errors.New("EvalSymlinks: too many links")
//line /usr/local/go/src/path/filepath/symlink.go:100
			// _ = "end of CoverTab[18290]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:101
			_go_fuzz_dep_.CoverTab[18291]++
//line /usr/local/go/src/path/filepath/symlink.go:101
			// _ = "end of CoverTab[18291]"
//line /usr/local/go/src/path/filepath/symlink.go:101
		}
//line /usr/local/go/src/path/filepath/symlink.go:101
		// _ = "end of CoverTab[18252]"
//line /usr/local/go/src/path/filepath/symlink.go:101
		_go_fuzz_dep_.CoverTab[18253]++

								link, err := os.Readlink(dest)
								if err != nil {
//line /usr/local/go/src/path/filepath/symlink.go:104
			_go_fuzz_dep_.CoverTab[18292]++
									return "", err
//line /usr/local/go/src/path/filepath/symlink.go:105
			// _ = "end of CoverTab[18292]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:106
			_go_fuzz_dep_.CoverTab[18293]++
//line /usr/local/go/src/path/filepath/symlink.go:106
			// _ = "end of CoverTab[18293]"
//line /usr/local/go/src/path/filepath/symlink.go:106
		}
//line /usr/local/go/src/path/filepath/symlink.go:106
		// _ = "end of CoverTab[18253]"
//line /usr/local/go/src/path/filepath/symlink.go:106
		_go_fuzz_dep_.CoverTab[18254]++

								if isWindowsDot && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:108
			_go_fuzz_dep_.CoverTab[18294]++
//line /usr/local/go/src/path/filepath/symlink.go:108
			return !IsAbs(link)
//line /usr/local/go/src/path/filepath/symlink.go:108
			// _ = "end of CoverTab[18294]"
//line /usr/local/go/src/path/filepath/symlink.go:108
		}() {
//line /usr/local/go/src/path/filepath/symlink.go:108
			_go_fuzz_dep_.CoverTab[18295]++

//line /usr/local/go/src/path/filepath/symlink.go:111
			break
//line /usr/local/go/src/path/filepath/symlink.go:111
			// _ = "end of CoverTab[18295]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:112
			_go_fuzz_dep_.CoverTab[18296]++
//line /usr/local/go/src/path/filepath/symlink.go:112
			// _ = "end of CoverTab[18296]"
//line /usr/local/go/src/path/filepath/symlink.go:112
		}
//line /usr/local/go/src/path/filepath/symlink.go:112
		// _ = "end of CoverTab[18254]"
//line /usr/local/go/src/path/filepath/symlink.go:112
		_go_fuzz_dep_.CoverTab[18255]++

								path = link + path[end:]

								v := volumeNameLen(link)
								if v > 0 {
//line /usr/local/go/src/path/filepath/symlink.go:117
			_go_fuzz_dep_.CoverTab[18297]++

									if v < len(link) && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:119
				_go_fuzz_dep_.CoverTab[18299]++
//line /usr/local/go/src/path/filepath/symlink.go:119
				return os.IsPathSeparator(link[v])
//line /usr/local/go/src/path/filepath/symlink.go:119
				// _ = "end of CoverTab[18299]"
//line /usr/local/go/src/path/filepath/symlink.go:119
			}() {
//line /usr/local/go/src/path/filepath/symlink.go:119
				_go_fuzz_dep_.CoverTab[18300]++
										v++
//line /usr/local/go/src/path/filepath/symlink.go:120
				// _ = "end of CoverTab[18300]"
			} else {
//line /usr/local/go/src/path/filepath/symlink.go:121
				_go_fuzz_dep_.CoverTab[18301]++
//line /usr/local/go/src/path/filepath/symlink.go:121
				// _ = "end of CoverTab[18301]"
//line /usr/local/go/src/path/filepath/symlink.go:121
			}
//line /usr/local/go/src/path/filepath/symlink.go:121
			// _ = "end of CoverTab[18297]"
//line /usr/local/go/src/path/filepath/symlink.go:121
			_go_fuzz_dep_.CoverTab[18298]++
									vol = link[:v]
									dest = vol
									end = len(vol)
//line /usr/local/go/src/path/filepath/symlink.go:124
			// _ = "end of CoverTab[18298]"
		} else {
//line /usr/local/go/src/path/filepath/symlink.go:125
			_go_fuzz_dep_.CoverTab[18302]++
//line /usr/local/go/src/path/filepath/symlink.go:125
			if len(link) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/symlink.go:125
				_go_fuzz_dep_.CoverTab[18303]++
//line /usr/local/go/src/path/filepath/symlink.go:125
				return os.IsPathSeparator(link[0])
//line /usr/local/go/src/path/filepath/symlink.go:125
				// _ = "end of CoverTab[18303]"
//line /usr/local/go/src/path/filepath/symlink.go:125
			}() {
//line /usr/local/go/src/path/filepath/symlink.go:125
				_go_fuzz_dep_.CoverTab[18304]++

										dest = link[:1]
										end = 1
//line /usr/local/go/src/path/filepath/symlink.go:128
				// _ = "end of CoverTab[18304]"
			} else {
//line /usr/local/go/src/path/filepath/symlink.go:129
				_go_fuzz_dep_.CoverTab[18305]++
				// Symlink to relative path; replace last
				// path component in dest.
				var r int
				for r = len(dest) - 1; r >= volLen; r-- {
//line /usr/local/go/src/path/filepath/symlink.go:133
					_go_fuzz_dep_.CoverTab[18308]++
											if os.IsPathSeparator(dest[r]) {
//line /usr/local/go/src/path/filepath/symlink.go:134
						_go_fuzz_dep_.CoverTab[18309]++
												break
//line /usr/local/go/src/path/filepath/symlink.go:135
						// _ = "end of CoverTab[18309]"
					} else {
//line /usr/local/go/src/path/filepath/symlink.go:136
						_go_fuzz_dep_.CoverTab[18310]++
//line /usr/local/go/src/path/filepath/symlink.go:136
						// _ = "end of CoverTab[18310]"
//line /usr/local/go/src/path/filepath/symlink.go:136
					}
//line /usr/local/go/src/path/filepath/symlink.go:136
					// _ = "end of CoverTab[18308]"
				}
//line /usr/local/go/src/path/filepath/symlink.go:137
				// _ = "end of CoverTab[18305]"
//line /usr/local/go/src/path/filepath/symlink.go:137
				_go_fuzz_dep_.CoverTab[18306]++
										if r < volLen {
//line /usr/local/go/src/path/filepath/symlink.go:138
					_go_fuzz_dep_.CoverTab[18311]++
											dest = vol
//line /usr/local/go/src/path/filepath/symlink.go:139
					// _ = "end of CoverTab[18311]"
				} else {
//line /usr/local/go/src/path/filepath/symlink.go:140
					_go_fuzz_dep_.CoverTab[18312]++
											dest = dest[:r]
//line /usr/local/go/src/path/filepath/symlink.go:141
					// _ = "end of CoverTab[18312]"
				}
//line /usr/local/go/src/path/filepath/symlink.go:142
				// _ = "end of CoverTab[18306]"
//line /usr/local/go/src/path/filepath/symlink.go:142
				_go_fuzz_dep_.CoverTab[18307]++
										end = 0
//line /usr/local/go/src/path/filepath/symlink.go:143
				// _ = "end of CoverTab[18307]"
			}
//line /usr/local/go/src/path/filepath/symlink.go:144
			// _ = "end of CoverTab[18302]"
//line /usr/local/go/src/path/filepath/symlink.go:144
		}
//line /usr/local/go/src/path/filepath/symlink.go:144
		// _ = "end of CoverTab[18255]"
	}
//line /usr/local/go/src/path/filepath/symlink.go:145
	// _ = "end of CoverTab[18241]"
//line /usr/local/go/src/path/filepath/symlink.go:145
	_go_fuzz_dep_.CoverTab[18242]++
							return Clean(dest), nil
//line /usr/local/go/src/path/filepath/symlink.go:146
	// _ = "end of CoverTab[18242]"
}

//line /usr/local/go/src/path/filepath/symlink.go:147
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/path/filepath/symlink.go:147
var _ = _go_fuzz_dep_.CoverTab
