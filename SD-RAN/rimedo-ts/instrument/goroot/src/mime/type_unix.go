// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

//line /usr/local/go/src/mime/type_unix.go:7
package mime

//line /usr/local/go/src/mime/type_unix.go:7
import (
//line /usr/local/go/src/mime/type_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/type_unix.go:7
)
//line /usr/local/go/src/mime/type_unix.go:7
import (
//line /usr/local/go/src/mime/type_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/type_unix.go:7
)

import (
	"bufio"
	"os"
	"strings"
)

func init() {
	osInitMime = initMimeUnix
}

// See https://specifications.freedesktop.org/shared-mime-info-spec/shared-mime-info-spec-0.21.html
//line /usr/local/go/src/mime/type_unix.go:19
// for the FreeDesktop Shared MIME-info Database specification.
//line /usr/local/go/src/mime/type_unix.go:21
var mimeGlobs = []string{
	"/usr/local/share/mime/globs2",
	"/usr/share/mime/globs2",
}

// Common locations for mime.types files on unix.
var typeFiles = []string{
	"/etc/mime.types",
	"/etc/apache2/mime.types",
	"/etc/apache/mime.types",
	"/etc/httpd/conf/mime.types",
}

func loadMimeGlobsFile(filename string) error {
//line /usr/local/go/src/mime/type_unix.go:34
	_go_fuzz_dep_.CoverTab[35865]++
						f, err := os.Open(filename)
						if err != nil {
//line /usr/local/go/src/mime/type_unix.go:36
		_go_fuzz_dep_.CoverTab[35869]++
							return err
//line /usr/local/go/src/mime/type_unix.go:37
		// _ = "end of CoverTab[35869]"
	} else {
//line /usr/local/go/src/mime/type_unix.go:38
		_go_fuzz_dep_.CoverTab[35870]++
//line /usr/local/go/src/mime/type_unix.go:38
		// _ = "end of CoverTab[35870]"
//line /usr/local/go/src/mime/type_unix.go:38
	}
//line /usr/local/go/src/mime/type_unix.go:38
	// _ = "end of CoverTab[35865]"
//line /usr/local/go/src/mime/type_unix.go:38
	_go_fuzz_dep_.CoverTab[35866]++
						defer f.Close()

						scanner := bufio.NewScanner(f)
						for scanner.Scan() {
//line /usr/local/go/src/mime/type_unix.go:42
		_go_fuzz_dep_.CoverTab[35871]++

							fields := strings.Split(scanner.Text(), ":")
							if len(fields) < 3 || func() bool {
//line /usr/local/go/src/mime/type_unix.go:45
			_go_fuzz_dep_.CoverTab[35875]++
//line /usr/local/go/src/mime/type_unix.go:45
			return len(fields[0]) < 1
//line /usr/local/go/src/mime/type_unix.go:45
			// _ = "end of CoverTab[35875]"
//line /usr/local/go/src/mime/type_unix.go:45
		}() || func() bool {
//line /usr/local/go/src/mime/type_unix.go:45
			_go_fuzz_dep_.CoverTab[35876]++
//line /usr/local/go/src/mime/type_unix.go:45
			return len(fields[2]) < 3
//line /usr/local/go/src/mime/type_unix.go:45
			// _ = "end of CoverTab[35876]"
//line /usr/local/go/src/mime/type_unix.go:45
		}() {
//line /usr/local/go/src/mime/type_unix.go:45
			_go_fuzz_dep_.CoverTab[35877]++
								continue
//line /usr/local/go/src/mime/type_unix.go:46
			// _ = "end of CoverTab[35877]"
		} else {
//line /usr/local/go/src/mime/type_unix.go:47
			_go_fuzz_dep_.CoverTab[35878]++
//line /usr/local/go/src/mime/type_unix.go:47
			if fields[0][0] == '#' || func() bool {
//line /usr/local/go/src/mime/type_unix.go:47
				_go_fuzz_dep_.CoverTab[35879]++
//line /usr/local/go/src/mime/type_unix.go:47
				return fields[2][0] != '*'
//line /usr/local/go/src/mime/type_unix.go:47
				// _ = "end of CoverTab[35879]"
//line /usr/local/go/src/mime/type_unix.go:47
			}() || func() bool {
//line /usr/local/go/src/mime/type_unix.go:47
				_go_fuzz_dep_.CoverTab[35880]++
//line /usr/local/go/src/mime/type_unix.go:47
				return fields[2][1] != '.'
//line /usr/local/go/src/mime/type_unix.go:47
				// _ = "end of CoverTab[35880]"
//line /usr/local/go/src/mime/type_unix.go:47
			}() {
//line /usr/local/go/src/mime/type_unix.go:47
				_go_fuzz_dep_.CoverTab[35881]++
									continue
//line /usr/local/go/src/mime/type_unix.go:48
				// _ = "end of CoverTab[35881]"
			} else {
//line /usr/local/go/src/mime/type_unix.go:49
				_go_fuzz_dep_.CoverTab[35882]++
//line /usr/local/go/src/mime/type_unix.go:49
				// _ = "end of CoverTab[35882]"
//line /usr/local/go/src/mime/type_unix.go:49
			}
//line /usr/local/go/src/mime/type_unix.go:49
			// _ = "end of CoverTab[35878]"
//line /usr/local/go/src/mime/type_unix.go:49
		}
//line /usr/local/go/src/mime/type_unix.go:49
		// _ = "end of CoverTab[35871]"
//line /usr/local/go/src/mime/type_unix.go:49
		_go_fuzz_dep_.CoverTab[35872]++

							extension := fields[2][1:]
							if strings.ContainsAny(extension, "?*[") {
//line /usr/local/go/src/mime/type_unix.go:52
			_go_fuzz_dep_.CoverTab[35883]++

//line /usr/local/go/src/mime/type_unix.go:62
			continue
//line /usr/local/go/src/mime/type_unix.go:62
			// _ = "end of CoverTab[35883]"
		} else {
//line /usr/local/go/src/mime/type_unix.go:63
			_go_fuzz_dep_.CoverTab[35884]++
//line /usr/local/go/src/mime/type_unix.go:63
			// _ = "end of CoverTab[35884]"
//line /usr/local/go/src/mime/type_unix.go:63
		}
//line /usr/local/go/src/mime/type_unix.go:63
		// _ = "end of CoverTab[35872]"
//line /usr/local/go/src/mime/type_unix.go:63
		_go_fuzz_dep_.CoverTab[35873]++
							if _, ok := mimeTypes.Load(extension); ok {
//line /usr/local/go/src/mime/type_unix.go:64
			_go_fuzz_dep_.CoverTab[35885]++

//line /usr/local/go/src/mime/type_unix.go:68
			continue
//line /usr/local/go/src/mime/type_unix.go:68
			// _ = "end of CoverTab[35885]"
		} else {
//line /usr/local/go/src/mime/type_unix.go:69
			_go_fuzz_dep_.CoverTab[35886]++
//line /usr/local/go/src/mime/type_unix.go:69
			// _ = "end of CoverTab[35886]"
//line /usr/local/go/src/mime/type_unix.go:69
		}
//line /usr/local/go/src/mime/type_unix.go:69
		// _ = "end of CoverTab[35873]"
//line /usr/local/go/src/mime/type_unix.go:69
		_go_fuzz_dep_.CoverTab[35874]++

							setExtensionType(extension, fields[1])
//line /usr/local/go/src/mime/type_unix.go:71
		// _ = "end of CoverTab[35874]"
	}
//line /usr/local/go/src/mime/type_unix.go:72
	// _ = "end of CoverTab[35866]"
//line /usr/local/go/src/mime/type_unix.go:72
	_go_fuzz_dep_.CoverTab[35867]++
						if err := scanner.Err(); err != nil {
//line /usr/local/go/src/mime/type_unix.go:73
		_go_fuzz_dep_.CoverTab[35887]++
							panic(err)
//line /usr/local/go/src/mime/type_unix.go:74
		// _ = "end of CoverTab[35887]"
	} else {
//line /usr/local/go/src/mime/type_unix.go:75
		_go_fuzz_dep_.CoverTab[35888]++
//line /usr/local/go/src/mime/type_unix.go:75
		// _ = "end of CoverTab[35888]"
//line /usr/local/go/src/mime/type_unix.go:75
	}
//line /usr/local/go/src/mime/type_unix.go:75
	// _ = "end of CoverTab[35867]"
//line /usr/local/go/src/mime/type_unix.go:75
	_go_fuzz_dep_.CoverTab[35868]++
						return nil
//line /usr/local/go/src/mime/type_unix.go:76
	// _ = "end of CoverTab[35868]"
}

func loadMimeFile(filename string) {
//line /usr/local/go/src/mime/type_unix.go:79
	_go_fuzz_dep_.CoverTab[35889]++
						f, err := os.Open(filename)
						if err != nil {
//line /usr/local/go/src/mime/type_unix.go:81
		_go_fuzz_dep_.CoverTab[35892]++
							return
//line /usr/local/go/src/mime/type_unix.go:82
		// _ = "end of CoverTab[35892]"
	} else {
//line /usr/local/go/src/mime/type_unix.go:83
		_go_fuzz_dep_.CoverTab[35893]++
//line /usr/local/go/src/mime/type_unix.go:83
		// _ = "end of CoverTab[35893]"
//line /usr/local/go/src/mime/type_unix.go:83
	}
//line /usr/local/go/src/mime/type_unix.go:83
	// _ = "end of CoverTab[35889]"
//line /usr/local/go/src/mime/type_unix.go:83
	_go_fuzz_dep_.CoverTab[35890]++
						defer f.Close()

						scanner := bufio.NewScanner(f)
						for scanner.Scan() {
//line /usr/local/go/src/mime/type_unix.go:87
		_go_fuzz_dep_.CoverTab[35894]++
							fields := strings.Fields(scanner.Text())
							if len(fields) <= 1 || func() bool {
//line /usr/local/go/src/mime/type_unix.go:89
			_go_fuzz_dep_.CoverTab[35896]++
//line /usr/local/go/src/mime/type_unix.go:89
			return fields[0][0] == '#'
//line /usr/local/go/src/mime/type_unix.go:89
			// _ = "end of CoverTab[35896]"
//line /usr/local/go/src/mime/type_unix.go:89
		}() {
//line /usr/local/go/src/mime/type_unix.go:89
			_go_fuzz_dep_.CoverTab[35897]++
								continue
//line /usr/local/go/src/mime/type_unix.go:90
			// _ = "end of CoverTab[35897]"
		} else {
//line /usr/local/go/src/mime/type_unix.go:91
			_go_fuzz_dep_.CoverTab[35898]++
//line /usr/local/go/src/mime/type_unix.go:91
			// _ = "end of CoverTab[35898]"
//line /usr/local/go/src/mime/type_unix.go:91
		}
//line /usr/local/go/src/mime/type_unix.go:91
		// _ = "end of CoverTab[35894]"
//line /usr/local/go/src/mime/type_unix.go:91
		_go_fuzz_dep_.CoverTab[35895]++
							mimeType := fields[0]
							for _, ext := range fields[1:] {
//line /usr/local/go/src/mime/type_unix.go:93
			_go_fuzz_dep_.CoverTab[35899]++
								if ext[0] == '#' {
//line /usr/local/go/src/mime/type_unix.go:94
				_go_fuzz_dep_.CoverTab[35901]++
									break
//line /usr/local/go/src/mime/type_unix.go:95
				// _ = "end of CoverTab[35901]"
			} else {
//line /usr/local/go/src/mime/type_unix.go:96
				_go_fuzz_dep_.CoverTab[35902]++
//line /usr/local/go/src/mime/type_unix.go:96
				// _ = "end of CoverTab[35902]"
//line /usr/local/go/src/mime/type_unix.go:96
			}
//line /usr/local/go/src/mime/type_unix.go:96
			// _ = "end of CoverTab[35899]"
//line /usr/local/go/src/mime/type_unix.go:96
			_go_fuzz_dep_.CoverTab[35900]++
								setExtensionType("."+ext, mimeType)
//line /usr/local/go/src/mime/type_unix.go:97
			// _ = "end of CoverTab[35900]"
		}
//line /usr/local/go/src/mime/type_unix.go:98
		// _ = "end of CoverTab[35895]"
	}
//line /usr/local/go/src/mime/type_unix.go:99
	// _ = "end of CoverTab[35890]"
//line /usr/local/go/src/mime/type_unix.go:99
	_go_fuzz_dep_.CoverTab[35891]++
						if err := scanner.Err(); err != nil {
//line /usr/local/go/src/mime/type_unix.go:100
		_go_fuzz_dep_.CoverTab[35903]++
							panic(err)
//line /usr/local/go/src/mime/type_unix.go:101
		// _ = "end of CoverTab[35903]"
	} else {
//line /usr/local/go/src/mime/type_unix.go:102
		_go_fuzz_dep_.CoverTab[35904]++
//line /usr/local/go/src/mime/type_unix.go:102
		// _ = "end of CoverTab[35904]"
//line /usr/local/go/src/mime/type_unix.go:102
	}
//line /usr/local/go/src/mime/type_unix.go:102
	// _ = "end of CoverTab[35891]"
}

func initMimeUnix() {
//line /usr/local/go/src/mime/type_unix.go:105
	_go_fuzz_dep_.CoverTab[35905]++
						for _, filename := range mimeGlobs {
//line /usr/local/go/src/mime/type_unix.go:106
		_go_fuzz_dep_.CoverTab[35907]++
							if err := loadMimeGlobsFile(filename); err == nil {
//line /usr/local/go/src/mime/type_unix.go:107
			_go_fuzz_dep_.CoverTab[35908]++
								return
//line /usr/local/go/src/mime/type_unix.go:108
			// _ = "end of CoverTab[35908]"
		} else {
//line /usr/local/go/src/mime/type_unix.go:109
			_go_fuzz_dep_.CoverTab[35909]++
//line /usr/local/go/src/mime/type_unix.go:109
			// _ = "end of CoverTab[35909]"
//line /usr/local/go/src/mime/type_unix.go:109
		}
//line /usr/local/go/src/mime/type_unix.go:109
		// _ = "end of CoverTab[35907]"
	}
//line /usr/local/go/src/mime/type_unix.go:110
	// _ = "end of CoverTab[35905]"
//line /usr/local/go/src/mime/type_unix.go:110
	_go_fuzz_dep_.CoverTab[35906]++

//line /usr/local/go/src/mime/type_unix.go:113
	for _, filename := range typeFiles {
//line /usr/local/go/src/mime/type_unix.go:113
		_go_fuzz_dep_.CoverTab[35910]++
							loadMimeFile(filename)
//line /usr/local/go/src/mime/type_unix.go:114
		// _ = "end of CoverTab[35910]"
	}
//line /usr/local/go/src/mime/type_unix.go:115
	// _ = "end of CoverTab[35906]"
}

func initMimeForTests() map[string]string {
//line /usr/local/go/src/mime/type_unix.go:118
	_go_fuzz_dep_.CoverTab[35911]++
						mimeGlobs = []string{""}
						typeFiles = []string{"testdata/test.types"}
						return map[string]string{
		".T1":	"application/test",
		".t2":	"text/test; charset=utf-8",
		".png":	"image/png",
	}
//line /usr/local/go/src/mime/type_unix.go:125
	// _ = "end of CoverTab[35911]"
}

//line /usr/local/go/src/mime/type_unix.go:126
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/type_unix.go:126
var _ = _go_fuzz_dep_.CoverTab
