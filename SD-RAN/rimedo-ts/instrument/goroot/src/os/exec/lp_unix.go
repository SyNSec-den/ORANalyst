// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /usr/local/go/src/os/exec/lp_unix.go:7
package exec

//line /usr/local/go/src/os/exec/lp_unix.go:7
import (
//line /usr/local/go/src/os/exec/lp_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/exec/lp_unix.go:7
)
//line /usr/local/go/src/os/exec/lp_unix.go:7
import (
//line /usr/local/go/src/os/exec/lp_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/exec/lp_unix.go:7
)

import (
	"errors"
	"internal/syscall/unix"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// ErrNotFound is the error resulting if a path search failed to find an executable file.
var ErrNotFound = errors.New("executable file not found in $PATH")

func findExecutable(file string) error {
//line /usr/local/go/src/os/exec/lp_unix.go:22
	_go_fuzz_dep_.CoverTab[107474]++
						d, err := os.Stat(file)
						if err != nil {
//line /usr/local/go/src/os/exec/lp_unix.go:24
		_go_fuzz_dep_.CoverTab[107479]++
							return err
//line /usr/local/go/src/os/exec/lp_unix.go:25
		// _ = "end of CoverTab[107479]"
	} else {
//line /usr/local/go/src/os/exec/lp_unix.go:26
		_go_fuzz_dep_.CoverTab[107480]++
//line /usr/local/go/src/os/exec/lp_unix.go:26
		// _ = "end of CoverTab[107480]"
//line /usr/local/go/src/os/exec/lp_unix.go:26
	}
//line /usr/local/go/src/os/exec/lp_unix.go:26
	// _ = "end of CoverTab[107474]"
//line /usr/local/go/src/os/exec/lp_unix.go:26
	_go_fuzz_dep_.CoverTab[107475]++
						m := d.Mode()
						if m.IsDir() {
//line /usr/local/go/src/os/exec/lp_unix.go:28
		_go_fuzz_dep_.CoverTab[107481]++
							return syscall.EISDIR
//line /usr/local/go/src/os/exec/lp_unix.go:29
		// _ = "end of CoverTab[107481]"
	} else {
//line /usr/local/go/src/os/exec/lp_unix.go:30
		_go_fuzz_dep_.CoverTab[107482]++
//line /usr/local/go/src/os/exec/lp_unix.go:30
		// _ = "end of CoverTab[107482]"
//line /usr/local/go/src/os/exec/lp_unix.go:30
	}
//line /usr/local/go/src/os/exec/lp_unix.go:30
	// _ = "end of CoverTab[107475]"
//line /usr/local/go/src/os/exec/lp_unix.go:30
	_go_fuzz_dep_.CoverTab[107476]++
						err = unix.Eaccess(file, unix.X_OK)

//line /usr/local/go/src/os/exec/lp_unix.go:35
	if err == nil || func() bool {
//line /usr/local/go/src/os/exec/lp_unix.go:35
		_go_fuzz_dep_.CoverTab[107483]++
//line /usr/local/go/src/os/exec/lp_unix.go:35
		return (err != syscall.ENOSYS && func() bool {
//line /usr/local/go/src/os/exec/lp_unix.go:35
			_go_fuzz_dep_.CoverTab[107484]++
//line /usr/local/go/src/os/exec/lp_unix.go:35
			return err != syscall.EPERM
//line /usr/local/go/src/os/exec/lp_unix.go:35
			// _ = "end of CoverTab[107484]"
//line /usr/local/go/src/os/exec/lp_unix.go:35
		}())
//line /usr/local/go/src/os/exec/lp_unix.go:35
		// _ = "end of CoverTab[107483]"
//line /usr/local/go/src/os/exec/lp_unix.go:35
	}() {
//line /usr/local/go/src/os/exec/lp_unix.go:35
		_go_fuzz_dep_.CoverTab[107485]++
							return err
//line /usr/local/go/src/os/exec/lp_unix.go:36
		// _ = "end of CoverTab[107485]"
	} else {
//line /usr/local/go/src/os/exec/lp_unix.go:37
		_go_fuzz_dep_.CoverTab[107486]++
//line /usr/local/go/src/os/exec/lp_unix.go:37
		// _ = "end of CoverTab[107486]"
//line /usr/local/go/src/os/exec/lp_unix.go:37
	}
//line /usr/local/go/src/os/exec/lp_unix.go:37
	// _ = "end of CoverTab[107476]"
//line /usr/local/go/src/os/exec/lp_unix.go:37
	_go_fuzz_dep_.CoverTab[107477]++
						if m&0111 != 0 {
//line /usr/local/go/src/os/exec/lp_unix.go:38
		_go_fuzz_dep_.CoverTab[107487]++
							return nil
//line /usr/local/go/src/os/exec/lp_unix.go:39
		// _ = "end of CoverTab[107487]"
	} else {
//line /usr/local/go/src/os/exec/lp_unix.go:40
		_go_fuzz_dep_.CoverTab[107488]++
//line /usr/local/go/src/os/exec/lp_unix.go:40
		// _ = "end of CoverTab[107488]"
//line /usr/local/go/src/os/exec/lp_unix.go:40
	}
//line /usr/local/go/src/os/exec/lp_unix.go:40
	// _ = "end of CoverTab[107477]"
//line /usr/local/go/src/os/exec/lp_unix.go:40
	_go_fuzz_dep_.CoverTab[107478]++
						return fs.ErrPermission
//line /usr/local/go/src/os/exec/lp_unix.go:41
	// _ = "end of CoverTab[107478]"
}

// LookPath searches for an executable named file in the
//line /usr/local/go/src/os/exec/lp_unix.go:44
// directories named by the PATH environment variable.
//line /usr/local/go/src/os/exec/lp_unix.go:44
// If file contains a slash, it is tried directly and the PATH is not consulted.
//line /usr/local/go/src/os/exec/lp_unix.go:44
// Otherwise, on success, the result is an absolute path.
//line /usr/local/go/src/os/exec/lp_unix.go:44
//
//line /usr/local/go/src/os/exec/lp_unix.go:44
// In older versions of Go, LookPath could return a path relative to the current directory.
//line /usr/local/go/src/os/exec/lp_unix.go:44
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
//line /usr/local/go/src/os/exec/lp_unix.go:44
// errors.Is(err, ErrDot). See the package documentation for more details.
//line /usr/local/go/src/os/exec/lp_unix.go:52
func LookPath(file string) (string, error) {
//line /usr/local/go/src/os/exec/lp_unix.go:52
	_go_fuzz_dep_.CoverTab[107489]++

//line /usr/local/go/src/os/exec/lp_unix.go:57
	if strings.Contains(file, "/") {
//line /usr/local/go/src/os/exec/lp_unix.go:57
		_go_fuzz_dep_.CoverTab[107492]++
							err := findExecutable(file)
							if err == nil {
//line /usr/local/go/src/os/exec/lp_unix.go:59
			_go_fuzz_dep_.CoverTab[107494]++
								return file, nil
//line /usr/local/go/src/os/exec/lp_unix.go:60
			// _ = "end of CoverTab[107494]"
		} else {
//line /usr/local/go/src/os/exec/lp_unix.go:61
			_go_fuzz_dep_.CoverTab[107495]++
//line /usr/local/go/src/os/exec/lp_unix.go:61
			// _ = "end of CoverTab[107495]"
//line /usr/local/go/src/os/exec/lp_unix.go:61
		}
//line /usr/local/go/src/os/exec/lp_unix.go:61
		// _ = "end of CoverTab[107492]"
//line /usr/local/go/src/os/exec/lp_unix.go:61
		_go_fuzz_dep_.CoverTab[107493]++
							return "", &Error{file, err}
//line /usr/local/go/src/os/exec/lp_unix.go:62
		// _ = "end of CoverTab[107493]"
	} else {
//line /usr/local/go/src/os/exec/lp_unix.go:63
		_go_fuzz_dep_.CoverTab[107496]++
//line /usr/local/go/src/os/exec/lp_unix.go:63
		// _ = "end of CoverTab[107496]"
//line /usr/local/go/src/os/exec/lp_unix.go:63
	}
//line /usr/local/go/src/os/exec/lp_unix.go:63
	// _ = "end of CoverTab[107489]"
//line /usr/local/go/src/os/exec/lp_unix.go:63
	_go_fuzz_dep_.CoverTab[107490]++
						path := os.Getenv("PATH")
						for _, dir := range filepath.SplitList(path) {
//line /usr/local/go/src/os/exec/lp_unix.go:65
		_go_fuzz_dep_.CoverTab[107497]++
							if dir == "" {
//line /usr/local/go/src/os/exec/lp_unix.go:66
			_go_fuzz_dep_.CoverTab[107499]++

								dir = "."
//line /usr/local/go/src/os/exec/lp_unix.go:68
			// _ = "end of CoverTab[107499]"
		} else {
//line /usr/local/go/src/os/exec/lp_unix.go:69
			_go_fuzz_dep_.CoverTab[107500]++
//line /usr/local/go/src/os/exec/lp_unix.go:69
			// _ = "end of CoverTab[107500]"
//line /usr/local/go/src/os/exec/lp_unix.go:69
		}
//line /usr/local/go/src/os/exec/lp_unix.go:69
		// _ = "end of CoverTab[107497]"
//line /usr/local/go/src/os/exec/lp_unix.go:69
		_go_fuzz_dep_.CoverTab[107498]++
							path := filepath.Join(dir, file)
							if err := findExecutable(path); err == nil {
//line /usr/local/go/src/os/exec/lp_unix.go:71
			_go_fuzz_dep_.CoverTab[107501]++
								if !filepath.IsAbs(path) && func() bool {
//line /usr/local/go/src/os/exec/lp_unix.go:72
				_go_fuzz_dep_.CoverTab[107503]++
//line /usr/local/go/src/os/exec/lp_unix.go:72
				return execerrdot.Value() != "0"
//line /usr/local/go/src/os/exec/lp_unix.go:72
				// _ = "end of CoverTab[107503]"
//line /usr/local/go/src/os/exec/lp_unix.go:72
			}() {
//line /usr/local/go/src/os/exec/lp_unix.go:72
				_go_fuzz_dep_.CoverTab[107504]++
									return path, &Error{file, ErrDot}
//line /usr/local/go/src/os/exec/lp_unix.go:73
				// _ = "end of CoverTab[107504]"
			} else {
//line /usr/local/go/src/os/exec/lp_unix.go:74
				_go_fuzz_dep_.CoverTab[107505]++
//line /usr/local/go/src/os/exec/lp_unix.go:74
				// _ = "end of CoverTab[107505]"
//line /usr/local/go/src/os/exec/lp_unix.go:74
			}
//line /usr/local/go/src/os/exec/lp_unix.go:74
			// _ = "end of CoverTab[107501]"
//line /usr/local/go/src/os/exec/lp_unix.go:74
			_go_fuzz_dep_.CoverTab[107502]++
								return path, nil
//line /usr/local/go/src/os/exec/lp_unix.go:75
			// _ = "end of CoverTab[107502]"
		} else {
//line /usr/local/go/src/os/exec/lp_unix.go:76
			_go_fuzz_dep_.CoverTab[107506]++
//line /usr/local/go/src/os/exec/lp_unix.go:76
			// _ = "end of CoverTab[107506]"
//line /usr/local/go/src/os/exec/lp_unix.go:76
		}
//line /usr/local/go/src/os/exec/lp_unix.go:76
		// _ = "end of CoverTab[107498]"
	}
//line /usr/local/go/src/os/exec/lp_unix.go:77
	// _ = "end of CoverTab[107490]"
//line /usr/local/go/src/os/exec/lp_unix.go:77
	_go_fuzz_dep_.CoverTab[107491]++
						return "", &Error{file, ErrNotFound}
//line /usr/local/go/src/os/exec/lp_unix.go:78
	// _ = "end of CoverTab[107491]"
}

//line /usr/local/go/src/os/exec/lp_unix.go:79
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/exec/lp_unix.go:79
var _ = _go_fuzz_dep_.CoverTab
