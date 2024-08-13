// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris

//line /usr/local/go/src/crypto/x509/root_unix.go:7
package x509

//line /usr/local/go/src/crypto/x509/root_unix.go:7
import (
//line /usr/local/go/src/crypto/x509/root_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/root_unix.go:7
)
//line /usr/local/go/src/crypto/x509/root_unix.go:7
import (
//line /usr/local/go/src/crypto/x509/root_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/root_unix.go:7
)

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	// certFileEnv is the environment variable which identifies where to locate
	// the SSL certificate file. If set this overrides the system default.
	certFileEnv	= "SSL_CERT_FILE"

	// certDirEnv is the environment variable which identifies which directory
	// to check for SSL certificate files. If set this overrides the system default.
	// It is a colon separated list of directories.
	// See https://www.openssl.org/docs/man1.0.2/man1/c_rehash.html.
	certDirEnv	= "SSL_CERT_DIR"
)

func (c *Certificate) systemVerify(opts *VerifyOptions) (chains [][]*Certificate, err error) {
//line /usr/local/go/src/crypto/x509/root_unix.go:28
	_go_fuzz_dep_.CoverTab[19267]++
							return nil, nil
//line /usr/local/go/src/crypto/x509/root_unix.go:29
	// _ = "end of CoverTab[19267]"
}

func loadSystemRoots() (*CertPool, error) {
//line /usr/local/go/src/crypto/x509/root_unix.go:32
	_go_fuzz_dep_.CoverTab[19268]++
							roots := NewCertPool()

							files := certFiles
							if f := os.Getenv(certFileEnv); f != "" {
//line /usr/local/go/src/crypto/x509/root_unix.go:36
		_go_fuzz_dep_.CoverTab[19274]++
								files = []string{f}
//line /usr/local/go/src/crypto/x509/root_unix.go:37
		// _ = "end of CoverTab[19274]"
	} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:38
		_go_fuzz_dep_.CoverTab[19275]++
//line /usr/local/go/src/crypto/x509/root_unix.go:38
		// _ = "end of CoverTab[19275]"
//line /usr/local/go/src/crypto/x509/root_unix.go:38
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:38
	// _ = "end of CoverTab[19268]"
//line /usr/local/go/src/crypto/x509/root_unix.go:38
	_go_fuzz_dep_.CoverTab[19269]++

							var firstErr error
							for _, file := range files {
//line /usr/local/go/src/crypto/x509/root_unix.go:41
		_go_fuzz_dep_.CoverTab[19276]++
								data, err := os.ReadFile(file)
								if err == nil {
//line /usr/local/go/src/crypto/x509/root_unix.go:43
			_go_fuzz_dep_.CoverTab[19278]++
									roots.AppendCertsFromPEM(data)
									break
//line /usr/local/go/src/crypto/x509/root_unix.go:45
			// _ = "end of CoverTab[19278]"
		} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:46
			_go_fuzz_dep_.CoverTab[19279]++
//line /usr/local/go/src/crypto/x509/root_unix.go:46
			// _ = "end of CoverTab[19279]"
//line /usr/local/go/src/crypto/x509/root_unix.go:46
		}
//line /usr/local/go/src/crypto/x509/root_unix.go:46
		// _ = "end of CoverTab[19276]"
//line /usr/local/go/src/crypto/x509/root_unix.go:46
		_go_fuzz_dep_.CoverTab[19277]++
								if firstErr == nil && func() bool {
//line /usr/local/go/src/crypto/x509/root_unix.go:47
			_go_fuzz_dep_.CoverTab[19280]++
//line /usr/local/go/src/crypto/x509/root_unix.go:47
			return !os.IsNotExist(err)
//line /usr/local/go/src/crypto/x509/root_unix.go:47
			// _ = "end of CoverTab[19280]"
//line /usr/local/go/src/crypto/x509/root_unix.go:47
		}() {
//line /usr/local/go/src/crypto/x509/root_unix.go:47
			_go_fuzz_dep_.CoverTab[19281]++
									firstErr = err
//line /usr/local/go/src/crypto/x509/root_unix.go:48
			// _ = "end of CoverTab[19281]"
		} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:49
			_go_fuzz_dep_.CoverTab[19282]++
//line /usr/local/go/src/crypto/x509/root_unix.go:49
			// _ = "end of CoverTab[19282]"
//line /usr/local/go/src/crypto/x509/root_unix.go:49
		}
//line /usr/local/go/src/crypto/x509/root_unix.go:49
		// _ = "end of CoverTab[19277]"
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:50
	// _ = "end of CoverTab[19269]"
//line /usr/local/go/src/crypto/x509/root_unix.go:50
	_go_fuzz_dep_.CoverTab[19270]++

							dirs := certDirectories
							if d := os.Getenv(certDirEnv); d != "" {
//line /usr/local/go/src/crypto/x509/root_unix.go:53
		_go_fuzz_dep_.CoverTab[19283]++

//line /usr/local/go/src/crypto/x509/root_unix.go:58
		dirs = strings.Split(d, ":")
//line /usr/local/go/src/crypto/x509/root_unix.go:58
		// _ = "end of CoverTab[19283]"
	} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:59
		_go_fuzz_dep_.CoverTab[19284]++
//line /usr/local/go/src/crypto/x509/root_unix.go:59
		// _ = "end of CoverTab[19284]"
//line /usr/local/go/src/crypto/x509/root_unix.go:59
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:59
	// _ = "end of CoverTab[19270]"
//line /usr/local/go/src/crypto/x509/root_unix.go:59
	_go_fuzz_dep_.CoverTab[19271]++

							for _, directory := range dirs {
//line /usr/local/go/src/crypto/x509/root_unix.go:61
		_go_fuzz_dep_.CoverTab[19285]++
								fis, err := readUniqueDirectoryEntries(directory)
								if err != nil {
//line /usr/local/go/src/crypto/x509/root_unix.go:63
			_go_fuzz_dep_.CoverTab[19287]++
									if firstErr == nil && func() bool {
//line /usr/local/go/src/crypto/x509/root_unix.go:64
				_go_fuzz_dep_.CoverTab[19289]++
//line /usr/local/go/src/crypto/x509/root_unix.go:64
				return !os.IsNotExist(err)
//line /usr/local/go/src/crypto/x509/root_unix.go:64
				// _ = "end of CoverTab[19289]"
//line /usr/local/go/src/crypto/x509/root_unix.go:64
			}() {
//line /usr/local/go/src/crypto/x509/root_unix.go:64
				_go_fuzz_dep_.CoverTab[19290]++
										firstErr = err
//line /usr/local/go/src/crypto/x509/root_unix.go:65
				// _ = "end of CoverTab[19290]"
			} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:66
				_go_fuzz_dep_.CoverTab[19291]++
//line /usr/local/go/src/crypto/x509/root_unix.go:66
				// _ = "end of CoverTab[19291]"
//line /usr/local/go/src/crypto/x509/root_unix.go:66
			}
//line /usr/local/go/src/crypto/x509/root_unix.go:66
			// _ = "end of CoverTab[19287]"
//line /usr/local/go/src/crypto/x509/root_unix.go:66
			_go_fuzz_dep_.CoverTab[19288]++
									continue
//line /usr/local/go/src/crypto/x509/root_unix.go:67
			// _ = "end of CoverTab[19288]"
		} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:68
			_go_fuzz_dep_.CoverTab[19292]++
//line /usr/local/go/src/crypto/x509/root_unix.go:68
			// _ = "end of CoverTab[19292]"
//line /usr/local/go/src/crypto/x509/root_unix.go:68
		}
//line /usr/local/go/src/crypto/x509/root_unix.go:68
		// _ = "end of CoverTab[19285]"
//line /usr/local/go/src/crypto/x509/root_unix.go:68
		_go_fuzz_dep_.CoverTab[19286]++
								for _, fi := range fis {
//line /usr/local/go/src/crypto/x509/root_unix.go:69
			_go_fuzz_dep_.CoverTab[19293]++
									data, err := os.ReadFile(directory + "/" + fi.Name())
									if err == nil {
//line /usr/local/go/src/crypto/x509/root_unix.go:71
				_go_fuzz_dep_.CoverTab[19294]++
										roots.AppendCertsFromPEM(data)
//line /usr/local/go/src/crypto/x509/root_unix.go:72
				// _ = "end of CoverTab[19294]"
			} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:73
				_go_fuzz_dep_.CoverTab[19295]++
//line /usr/local/go/src/crypto/x509/root_unix.go:73
				// _ = "end of CoverTab[19295]"
//line /usr/local/go/src/crypto/x509/root_unix.go:73
			}
//line /usr/local/go/src/crypto/x509/root_unix.go:73
			// _ = "end of CoverTab[19293]"
		}
//line /usr/local/go/src/crypto/x509/root_unix.go:74
		// _ = "end of CoverTab[19286]"
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:75
	// _ = "end of CoverTab[19271]"
//line /usr/local/go/src/crypto/x509/root_unix.go:75
	_go_fuzz_dep_.CoverTab[19272]++

							if roots.len() > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/root_unix.go:77
		_go_fuzz_dep_.CoverTab[19296]++
//line /usr/local/go/src/crypto/x509/root_unix.go:77
		return firstErr == nil
//line /usr/local/go/src/crypto/x509/root_unix.go:77
		// _ = "end of CoverTab[19296]"
//line /usr/local/go/src/crypto/x509/root_unix.go:77
	}() {
//line /usr/local/go/src/crypto/x509/root_unix.go:77
		_go_fuzz_dep_.CoverTab[19297]++
								return roots, nil
//line /usr/local/go/src/crypto/x509/root_unix.go:78
		// _ = "end of CoverTab[19297]"
	} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:79
		_go_fuzz_dep_.CoverTab[19298]++
//line /usr/local/go/src/crypto/x509/root_unix.go:79
		// _ = "end of CoverTab[19298]"
//line /usr/local/go/src/crypto/x509/root_unix.go:79
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:79
	// _ = "end of CoverTab[19272]"
//line /usr/local/go/src/crypto/x509/root_unix.go:79
	_go_fuzz_dep_.CoverTab[19273]++

							return nil, firstErr
//line /usr/local/go/src/crypto/x509/root_unix.go:81
	// _ = "end of CoverTab[19273]"
}

// readUniqueDirectoryEntries is like os.ReadDir but omits
//line /usr/local/go/src/crypto/x509/root_unix.go:84
// symlinks that point within the directory.
//line /usr/local/go/src/crypto/x509/root_unix.go:86
func readUniqueDirectoryEntries(dir string) ([]fs.DirEntry, error) {
//line /usr/local/go/src/crypto/x509/root_unix.go:86
	_go_fuzz_dep_.CoverTab[19299]++
							files, err := os.ReadDir(dir)
							if err != nil {
//line /usr/local/go/src/crypto/x509/root_unix.go:88
		_go_fuzz_dep_.CoverTab[19302]++
								return nil, err
//line /usr/local/go/src/crypto/x509/root_unix.go:89
		// _ = "end of CoverTab[19302]"
	} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:90
		_go_fuzz_dep_.CoverTab[19303]++
//line /usr/local/go/src/crypto/x509/root_unix.go:90
		// _ = "end of CoverTab[19303]"
//line /usr/local/go/src/crypto/x509/root_unix.go:90
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:90
	// _ = "end of CoverTab[19299]"
//line /usr/local/go/src/crypto/x509/root_unix.go:90
	_go_fuzz_dep_.CoverTab[19300]++
							uniq := files[:0]
							for _, f := range files {
//line /usr/local/go/src/crypto/x509/root_unix.go:92
		_go_fuzz_dep_.CoverTab[19304]++
								if !isSameDirSymlink(f, dir) {
//line /usr/local/go/src/crypto/x509/root_unix.go:93
			_go_fuzz_dep_.CoverTab[19305]++
									uniq = append(uniq, f)
//line /usr/local/go/src/crypto/x509/root_unix.go:94
			// _ = "end of CoverTab[19305]"
		} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:95
			_go_fuzz_dep_.CoverTab[19306]++
//line /usr/local/go/src/crypto/x509/root_unix.go:95
			// _ = "end of CoverTab[19306]"
//line /usr/local/go/src/crypto/x509/root_unix.go:95
		}
//line /usr/local/go/src/crypto/x509/root_unix.go:95
		// _ = "end of CoverTab[19304]"
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:96
	// _ = "end of CoverTab[19300]"
//line /usr/local/go/src/crypto/x509/root_unix.go:96
	_go_fuzz_dep_.CoverTab[19301]++
							return uniq, nil
//line /usr/local/go/src/crypto/x509/root_unix.go:97
	// _ = "end of CoverTab[19301]"
}

// isSameDirSymlink reports whether fi in dir is a symlink with a
//line /usr/local/go/src/crypto/x509/root_unix.go:100
// target not containing a slash.
//line /usr/local/go/src/crypto/x509/root_unix.go:102
func isSameDirSymlink(f fs.DirEntry, dir string) bool {
//line /usr/local/go/src/crypto/x509/root_unix.go:102
	_go_fuzz_dep_.CoverTab[19307]++
							if f.Type()&fs.ModeSymlink == 0 {
//line /usr/local/go/src/crypto/x509/root_unix.go:103
		_go_fuzz_dep_.CoverTab[19309]++
								return false
//line /usr/local/go/src/crypto/x509/root_unix.go:104
		// _ = "end of CoverTab[19309]"
	} else {
//line /usr/local/go/src/crypto/x509/root_unix.go:105
		_go_fuzz_dep_.CoverTab[19310]++
//line /usr/local/go/src/crypto/x509/root_unix.go:105
		// _ = "end of CoverTab[19310]"
//line /usr/local/go/src/crypto/x509/root_unix.go:105
	}
//line /usr/local/go/src/crypto/x509/root_unix.go:105
	// _ = "end of CoverTab[19307]"
//line /usr/local/go/src/crypto/x509/root_unix.go:105
	_go_fuzz_dep_.CoverTab[19308]++
							target, err := os.Readlink(filepath.Join(dir, f.Name()))
							return err == nil && func() bool {
//line /usr/local/go/src/crypto/x509/root_unix.go:107
		_go_fuzz_dep_.CoverTab[19311]++
//line /usr/local/go/src/crypto/x509/root_unix.go:107
		return !strings.Contains(target, "/")
//line /usr/local/go/src/crypto/x509/root_unix.go:107
		// _ = "end of CoverTab[19311]"
//line /usr/local/go/src/crypto/x509/root_unix.go:107
	}()
//line /usr/local/go/src/crypto/x509/root_unix.go:107
	// _ = "end of CoverTab[19308]"
}

//line /usr/local/go/src/crypto/x509/root_unix.go:108
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/root_unix.go:108
var _ = _go_fuzz_dep_.CoverTab
