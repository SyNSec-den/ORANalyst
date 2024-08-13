// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/io/ioutil/ioutil.go:5
// Package ioutil implements some I/O utility functions.
//line /snap/go/10455/src/io/ioutil/ioutil.go:5
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:5
// Deprecated: As of Go 1.16, the same functionality is now provided
//line /snap/go/10455/src/io/ioutil/ioutil.go:5
// by package [io] or package [os], and those implementations
//line /snap/go/10455/src/io/ioutil/ioutil.go:5
// should be preferred in new code.
//line /snap/go/10455/src/io/ioutil/ioutil.go:5
// See the specific function documentation for details.
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
package ioutil

//line /snap/go/10455/src/io/ioutil/ioutil.go:11
import (
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
)
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
import (
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/io/ioutil/ioutil.go:11
)

import (
	"io"
	"io/fs"
	"os"
	"sort"
)

// ReadAll reads from r until an error or EOF and returns the data it read.
//line /snap/go/10455/src/io/ioutil/ioutil.go:20
// A successful call returns err == nil, not err == EOF. Because ReadAll is
//line /snap/go/10455/src/io/ioutil/ioutil.go:20
// defined to read from src until EOF, it does not treat an EOF from Read
//line /snap/go/10455/src/io/ioutil/ioutil.go:20
// as an error to be reported.
//line /snap/go/10455/src/io/ioutil/ioutil.go:20
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:20
// Deprecated: As of Go 1.16, this function simply calls [io.ReadAll].
//line /snap/go/10455/src/io/ioutil/ioutil.go:26
func ReadAll(r io.Reader) ([]byte, error) {
//line /snap/go/10455/src/io/ioutil/ioutil.go:26
	_go_fuzz_dep_.CoverTab[2292]++
							return io.ReadAll(r)
//line /snap/go/10455/src/io/ioutil/ioutil.go:27
	// _ = "end of CoverTab[2292]"
}

// ReadFile reads the file named by filename and returns the contents.
//line /snap/go/10455/src/io/ioutil/ioutil.go:30
// A successful call returns err == nil, not err == EOF. Because ReadFile
//line /snap/go/10455/src/io/ioutil/ioutil.go:30
// reads the whole file, it does not treat an EOF from Read as an error
//line /snap/go/10455/src/io/ioutil/ioutil.go:30
// to be reported.
//line /snap/go/10455/src/io/ioutil/ioutil.go:30
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:30
// Deprecated: As of Go 1.16, this function simply calls [os.ReadFile].
//line /snap/go/10455/src/io/ioutil/ioutil.go:36
func ReadFile(filename string) ([]byte, error) {
//line /snap/go/10455/src/io/ioutil/ioutil.go:36
	_go_fuzz_dep_.CoverTab[2293]++
							return os.ReadFile(filename)
//line /snap/go/10455/src/io/ioutil/ioutil.go:37
	// _ = "end of CoverTab[2293]"
}

// WriteFile writes data to a file named by filename.
//line /snap/go/10455/src/io/ioutil/ioutil.go:40
// If the file does not exist, WriteFile creates it with permissions perm
//line /snap/go/10455/src/io/ioutil/ioutil.go:40
// (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
//line /snap/go/10455/src/io/ioutil/ioutil.go:40
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:40
// Deprecated: As of Go 1.16, this function simply calls [os.WriteFile].
//line /snap/go/10455/src/io/ioutil/ioutil.go:45
func WriteFile(filename string, data []byte, perm fs.FileMode) error {
//line /snap/go/10455/src/io/ioutil/ioutil.go:45
	_go_fuzz_dep_.CoverTab[2294]++
							return os.WriteFile(filename, data, perm)
//line /snap/go/10455/src/io/ioutil/ioutil.go:46
	// _ = "end of CoverTab[2294]"
}

// ReadDir reads the directory named by dirname and returns
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// a list of fs.FileInfo for the directory's contents,
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// sorted by filename. If an error occurs reading the directory,
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// ReadDir returns no directory entries along with the error.
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// Deprecated: As of Go 1.16, [os.ReadDir] is a more efficient and correct choice:
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// it returns a list of [fs.DirEntry] instead of [fs.FileInfo],
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// and it returns partial results in the case of an error
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// midway through reading a directory.
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
// If you must continue obtaining a list of [fs.FileInfo], you still can:
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//	entries, err := os.ReadDir(dirname)
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//	if err != nil { ... }
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//	infos := make([]fs.FileInfo, 0, len(entries))
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//	for _, entry := range entries {
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//		info, err := entry.Info()
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//		if err != nil { ... }
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//		infos = append(infos, info)
//line /snap/go/10455/src/io/ioutil/ioutil.go:49
//	}
//line /snap/go/10455/src/io/ioutil/ioutil.go:69
func ReadDir(dirname string) ([]fs.FileInfo, error) {
//line /snap/go/10455/src/io/ioutil/ioutil.go:69
	_go_fuzz_dep_.CoverTab[2295]++
							f, err := os.Open(dirname)
							if err != nil {
//line /snap/go/10455/src/io/ioutil/ioutil.go:71
		_go_fuzz_dep_.CoverTab[525994]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:71
		_go_fuzz_dep_.CoverTab[2299]++
								return nil, err
//line /snap/go/10455/src/io/ioutil/ioutil.go:72
		// _ = "end of CoverTab[2299]"
	} else {
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
		_go_fuzz_dep_.CoverTab[525995]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
		_go_fuzz_dep_.CoverTab[2300]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
		// _ = "end of CoverTab[2300]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
	}
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
	// _ = "end of CoverTab[2295]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:73
	_go_fuzz_dep_.CoverTab[2296]++
							list, err := f.Readdir(-1)
							f.Close()
							if err != nil {
//line /snap/go/10455/src/io/ioutil/ioutil.go:76
		_go_fuzz_dep_.CoverTab[525996]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:76
		_go_fuzz_dep_.CoverTab[2301]++
								return nil, err
//line /snap/go/10455/src/io/ioutil/ioutil.go:77
		// _ = "end of CoverTab[2301]"
	} else {
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
		_go_fuzz_dep_.CoverTab[525997]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
		_go_fuzz_dep_.CoverTab[2302]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
		// _ = "end of CoverTab[2302]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
	}
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
	// _ = "end of CoverTab[2296]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:78
	_go_fuzz_dep_.CoverTab[2297]++
							sort.Slice(list, func(i, j int) bool {
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
		_go_fuzz_dep_.CoverTab[2303]++
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
		return list[i].Name() < list[j].Name()
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
		// _ = "end of CoverTab[2303]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
	})
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
	// _ = "end of CoverTab[2297]"
//line /snap/go/10455/src/io/ioutil/ioutil.go:79
	_go_fuzz_dep_.CoverTab[2298]++
							return list, nil
//line /snap/go/10455/src/io/ioutil/ioutil.go:80
	// _ = "end of CoverTab[2298]"
}

// NopCloser returns a ReadCloser with a no-op Close method wrapping
//line /snap/go/10455/src/io/ioutil/ioutil.go:83
// the provided Reader r.
//line /snap/go/10455/src/io/ioutil/ioutil.go:83
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:83
// Deprecated: As of Go 1.16, this function simply calls [io.NopCloser].
//line /snap/go/10455/src/io/ioutil/ioutil.go:87
func NopCloser(r io.Reader) io.ReadCloser {
//line /snap/go/10455/src/io/ioutil/ioutil.go:87
	_go_fuzz_dep_.CoverTab[2304]++
							return io.NopCloser(r)
//line /snap/go/10455/src/io/ioutil/ioutil.go:88
	// _ = "end of CoverTab[2304]"
}

// Discard is an io.Writer on which all Write calls succeed
//line /snap/go/10455/src/io/ioutil/ioutil.go:91
// without doing anything.
//line /snap/go/10455/src/io/ioutil/ioutil.go:91
//
//line /snap/go/10455/src/io/ioutil/ioutil.go:91
// Deprecated: As of Go 1.16, this value is simply [io.Discard].
//line /snap/go/10455/src/io/ioutil/ioutil.go:95
var Discard io.Writer = io.Discard
//line /snap/go/10455/src/io/ioutil/ioutil.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/io/ioutil/ioutil.go:95
var _ = _go_fuzz_dep_.CoverTab
